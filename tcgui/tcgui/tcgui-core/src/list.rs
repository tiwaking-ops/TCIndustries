//! List and Grid widgets implementation with selection, scrolling,
//! typed drag-and-drop payloads, and `can_drop` validation support.

use crate::geometry::{Rect, Vec2};
use crate::widget::WidgetState;

/// Typed drag-and-drop payload for list items, inventory slots, or hotbars.
#[derive(Debug, Clone, PartialEq, Eq)]
pub enum DragPayload {
    /// An inventory item with slot index and item ID / name.
    InventoryItem { slot_index: usize, item_id: u64, name: String },
    /// A hotbar slot reference.
    HotbarSlot { slot_index: usize },
    /// A generic list item.
    ListItem { index: usize, text: String },
    /// Custom text or id payload.
    Custom { tag: String, value: u64 },
}

/// A generic list widget supporting selection, scrolling, drag start, and drop acceptance.
pub struct ListWidget {
    /// Bounding rectangle in logical pixels.
    pub bounds: Rect,
    /// List items (label strings).
    pub items: Vec<String>,
    /// Index of currently selected item, if any.
    pub selected_index: Option<usize>,
    /// Height of each row in pixels.
    pub row_height: f32,
    /// Vertical scroll offset in pixels.
    pub scroll_offset: f32,
    /// Current drag payload being dragged from this list.
    pub drag_payload: Option<DragPayload>,
    /// Predicate function determining if a dropped payload is accepted.
    #[allow(clippy::type_complexity)]
    pub can_drop: Option<Box<dyn Fn(&DragPayload) -> bool>>,
    /// Current widget state.
    pub state: WidgetState,
}

impl std::fmt::Debug for ListWidget {
    fn fmt(&self, f: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
        f.debug_struct("ListWidget")
            .field("bounds", &self.bounds)
            .field("items", &self.items)
            .field("selected_index", &self.selected_index)
            .field("row_height", &self.row_height)
            .field("scroll_offset", &self.scroll_offset)
            .field("drag_payload", &self.drag_payload)
            .field("can_drop", &self.can_drop.as_ref().map(|_| "<callback>"))
            .field("state", &self.state)
            .finish()
    }
}

impl Clone for ListWidget {
    fn clone(&self) -> Self {
        Self {
            bounds: self.bounds,
            items: self.items.clone(),
            selected_index: self.selected_index,
            row_height: self.row_height,
            scroll_offset: self.scroll_offset,
            drag_payload: self.drag_payload.clone(),
            can_drop: None, // closures cannot be cloned; reset on clone
            state: self.state,
        }
    }
}

impl ListWidget {
    pub fn new(bounds: Rect, row_height: f32) -> Self {
        Self {
            bounds,
            items: Vec::new(),
            selected_index: None,
            row_height,
            scroll_offset: 0.0,
            drag_payload: None,
            can_drop: None,
            state: WidgetState::Normal,
        }
    }

    pub fn with_items(mut self, items: Vec<String>) -> Self {
        self.items = items;
        self
    }

    /// Handle mouse click to select an item.
    pub fn on_click(&mut self, pos: Vec2) -> Option<usize> {
        if !self.bounds.contains(pos) {
            return None;
        }
        let local_y = pos.y - self.bounds.min.y + self.scroll_offset;
        let index = (local_y / self.row_height) as usize;
        if index < self.items.len() {
            self.selected_index = Some(index);
            Some(index)
        } else {
            None
        }
    }

    /// Begin dragging an item at `index`, creating a typed [`DragPayload`].
    pub fn start_drag(&mut self, index: usize) -> Option<DragPayload> {
        if index < self.items.len() {
            let payload = DragPayload::ListItem {
                index,
                text: self.items[index].clone(),
            };
            self.drag_payload = Some(payload.clone());
            Some(payload)
        } else {
            None
        }
    }

    /// Check if a dropped payload is accepted by `can_drop`.
    pub fn evaluate_drop(&self, payload: &DragPayload) -> bool {
        if let Some(ref cb) = self.can_drop {
            cb(payload)
        } else {
            true // default accept
        }
    }

    pub fn bounds(&self) -> Rect {
        self.bounds
    }
}

/// A 2D grid widget (rows x columns) with cell selection and drag-and-drop support.
pub struct GridWidget {
    /// Bounding rectangle in logical pixels.
    pub bounds: Rect,
    /// Number of rows.
    pub rows: usize,
    /// Number of columns.
    pub cols: usize,
    /// Cell contents (row-major order).
    pub cells: Vec<String>,
    /// Currently selected cell coordinate `(row, col)`, if any.
    pub selected_cell: Option<(usize, usize)>,
    /// Current drag payload originating from this grid cell.
    pub drag_payload: Option<DragPayload>,
    /// Drop validation callback.
    #[allow(clippy::type_complexity)]
    pub can_drop: Option<Box<dyn Fn(&DragPayload) -> bool>>,
    /// Current widget state.
    pub state: WidgetState,
}

impl std::fmt::Debug for GridWidget {
    fn fmt(&self, f: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
        f.debug_struct("GridWidget")
            .field("bounds", &self.bounds)
            .field("rows", &self.rows)
            .field("cols", &self.cols)
            .field("cells", &self.cells)
            .field("selected_cell", &self.selected_cell)
            .field("drag_payload", &self.drag_payload)
            .field("can_drop", &self.can_drop.as_ref().map(|_| "<callback>"))
            .field("state", &self.state)
            .finish()
    }
}

impl Clone for GridWidget {
    fn clone(&self) -> Self {
        Self {
            bounds: self.bounds,
            rows: self.rows,
            cols: self.cols,
            cells: self.cells.clone(),
            selected_cell: self.selected_cell,
            drag_payload: self.drag_payload.clone(),
            can_drop: None,
            state: self.state,
        }
    }
}

impl GridWidget {
    pub fn new(bounds: Rect, rows: usize, cols: usize) -> Self {
        let total_cells = rows * cols;
        Self {
            bounds,
            rows,
            cols,
            cells: vec![String::new(); total_cells],
            selected_cell: None,
            drag_payload: None,
            can_drop: None,
            state: WidgetState::Normal,
        }
    }

    /// Cell width in logical pixels.
    pub fn cell_width(&self) -> f32 {
        if self.cols == 0 { 0.0 } else { self.bounds.width() / self.cols as f32 }
    }

    /// Cell height in logical pixels.
    pub fn cell_height(&self) -> f32 {
        if self.rows == 0 { 0.0 } else { self.bounds.height() / self.rows as f32 }
    }

    /// Handle click to select a grid cell.
    pub fn on_click(&mut self, pos: Vec2) -> Option<(usize, usize)> {
        if !self.bounds.contains(pos) {
            return None;
        }
        let cw = self.cell_width();
        let ch = self.cell_height();
        if cw <= 0.0 || ch <= 0.0 {
            return None;
        }
        let col = ((pos.x - self.bounds.min.x) / cw) as usize;
        let row = ((pos.y - self.bounds.min.y) / ch) as usize;

        if row < self.rows && col < self.cols {
            self.selected_cell = Some((row, col));
            Some((row, col))
        } else {
            None
        }
    }

    /// Start drag from a cell `(row, col)`.
    pub fn start_drag(&mut self, row: usize, col: usize) -> Option<DragPayload> {
        if row < self.rows && col < self.cols {
            let cell_idx = row * self.cols + col;
            let payload = DragPayload::InventoryItem {
                slot_index: cell_idx,
                item_id: (cell_idx + 1000) as u64,
                name: self.cells.get(cell_idx).cloned().unwrap_or_default(),
            };
            self.drag_payload = Some(payload.clone());
            Some(payload)
        } else {
            None
        }
    }

    /// Evaluate drop acceptance.
    pub fn evaluate_drop(&self, payload: &DragPayload) -> bool {
        if let Some(ref cb) = self.can_drop {
            cb(payload)
        } else {
            true
        }
    }

    pub fn bounds(&self) -> Rect {
        self.bounds
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_list_widget_selection_and_click() {
        let bounds = Rect::from_origin_size(Vec2::ZERO, Vec2::new(150.0, 100.0));
        let mut list = ListWidget::new(bounds, 20.0).with_items(vec![
            "Alpha".to_string(),
            "Beta".to_string(),
            "Gamma".to_string(),
        ]);

        // Click on second item (y = 25.0)
        let clicked = list.on_click(Vec2::new(10.0, 25.0));
        assert_eq!(clicked, Some(1));
        assert_eq!(list.selected_index, Some(1));
    }

    #[test]
    fn test_list_widget_drag_and_drop_validation() {
        let bounds = Rect::from_origin_size(Vec2::ZERO, Vec2::new(150.0, 100.0));
        let mut list = ListWidget::new(bounds, 20.0).with_items(vec!["Item 1".to_string()]);

        // Set can_drop rule: only accept InventoryItem
        list.can_drop = Some(Box::new(|payload| {
            matches!(payload, DragPayload::InventoryItem { .. })
        }));

        let accepted_payload = DragPayload::InventoryItem {
            slot_index: 0,
            item_id: 42,
            name: "Blaster".to_string(),
        };
        let rejected_payload = DragPayload::HotbarSlot { slot_index: 1 };

        assert!(list.evaluate_drop(&accepted_payload));
        assert!(!list.evaluate_drop(&rejected_payload));

        let started = list.start_drag(0);
        assert!(started.is_some());
    }

    #[test]
    fn test_grid_widget_selection_and_drag() {
        let bounds = Rect::from_origin_size(Vec2::ZERO, Vec2::new(200.0, 200.0));
        let mut grid = GridWidget::new(bounds, 2, 2);
        grid.cells[0] = "Sword".to_string();

        // Click top-right cell (col 1, row 0) -> x=150, y=50
        let clicked = grid.on_click(Vec2::new(150.0, 50.0));
        assert_eq!(clicked, Some((0, 1)));
        assert_eq!(grid.selected_cell, Some((0, 1)));

        // Start drag on top-left cell (row 0, col 0)
        let drag = grid.start_drag(0, 0);
        assert!(drag.is_some());
        if let Some(DragPayload::InventoryItem { slot_index, item_id, name }) = drag {
            assert_eq!(slot_index, 0);
            assert_eq!(item_id, 1000);
            assert_eq!(name, "Sword");
        } else {
            panic!("Expected InventoryItem drag payload");
        }
    }
}
