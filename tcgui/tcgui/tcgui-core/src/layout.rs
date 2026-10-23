//! Layout engine: table-driven, constraint-based box layout.
//!
//! The engine resolves a tree of [`LayoutNode`] specs into concrete physical
//! rectangles (`Rect`) measured in physical pixels. Layout is computed in
//! *logical* pixels first, then a uniform [`UiScale`] factor is applied to
//! produce physical bounds.
//!
//! ## Axis rules
//! - [`Axis::Horizontal`] is a row: the main axis runs left-to-right, the
//!   cross axis runs top-to-bottom.
//! - [`Axis::Vertical`] is a column: the main axis runs top-to-bottom, the
//!   cross axis runs left-to-right.
//!
//! ## Size values
//! A node declares its desired size along each axis with [`SizeValue`]:
//! - `Fixed(px)` &mdash; absolute logical pixels.
//! - `Percent(f)` &mdash; fraction of the *available* space along that axis.
//! - `Fill` &mdash; consume all remaining free space (main axis only; behaves
//!   like `Percent(1.0)`).
//! - `WrapContent` &mdash; size to intrinsic content (main axis); the node
//!   contributes its `intrinsic_size` or `0` if none is given.
//!
//! ## Constraints
//! [`Constraints`] apply hard `min`/`max` bounds (logical pixels) *after* a
//! raw size is resolved and *before* the cross-axis fill. If `min > max`,
//! resolution produces a [`LayoutError::ContradictoryConstraints`].
//!
//! ## Containers
//! - `Row` / `Column` lay children out along a single axis (`Axis::Horizontal`
//!   / `Axis::Vertical`). Children with no explicit size shrink to their
//!   `WrapContent` (or intrinsic) contribution.
//! - `Grid` lays children into a `rows x cols` cell matrix. Each child occupies
//!   one cell (row-major). The cell count is implicit from `rows`/`cols` and
//!   must equal the number of children, otherwise
//!   [`LayoutError::GridMismatch`] is raised.
//!
//! ## UI scale factor
//! A single `ui_scale: f32` (must be finite and `> 0`, otherwise
//! [`LayoutError::InvalidScale`]) multiplies every resolved logical rect to
//! yield the final physical rect.
//!
//! ## Error handling
//! [`LayoutEngine::layout`] returns a [`LayoutResult`] carrying the resolved
//! rects *and* any error. On error the rect map contains entries resolved up to
//! the point of failure, which keeps the engine total and deterministic.

use crate::geometry::{Rect, Vec2};

/// Identifier for a node in a [`LayoutTree`].
pub type LayoutNodeId = usize;

/// The direction of a single-axis container.
#[derive(Debug, Clone, Copy, PartialEq, Eq, Hash)]
pub enum Axis {
    /// Left-to-right main axis (a row).
    Horizontal,
    /// Top-to-bottom main axis (a column).
    Vertical,
}

impl Axis {
    /// The perpendicular axis.
    pub fn cross(self) -> Axis {
        match self {
            Axis::Horizontal => Axis::Vertical,
            Axis::Vertical => Axis::Horizontal,
        }
    }

    /// Index along the underlying `Vec2` (0 = x, 1 = y).
    pub fn as_usize(self) -> usize {
        match self {
            Axis::Horizontal => 0,
            Axis::Vertical => 1,
        }
    }
}

/// A desired size along a single axis.
#[derive(Debug, Clone, Copy, PartialEq)]
pub enum SizeValue {
    /// Absolute size in logical pixels.
    Fixed(f32),
    /// Fraction (0.0..=1.0) of the available space along this axis.
    Percent(f32),
    /// Consume all remaining free space (main axis only).
    Fill,
    /// Size to intrinsic / content contribution.
    WrapContent,
}

impl Default for SizeValue {
    fn default() -> Self {
        Self::WrapContent
    }
}

/// Hard minimum/maximum size constraints (logical pixels), per axis.
#[derive(Debug, Clone, Copy, Default, PartialEq)]
pub struct Constraints {
    pub min_width: f32,
    pub min_height: f32,
    pub max_width: f32,
    pub max_height: f32,
}

/// The kind of node layout container.
#[derive(Debug, Clone, Copy, PartialEq)]
pub enum NodeKind {
    /// A single-axis container (`Row`/`Column`).
    Stack { axis: Axis, gap: f32 },
    /// A grid of `rows x cols` cells.
    Grid { rows: usize, cols: usize, gap: f32 },
    /// A leaf widget sized by [`LayoutNode::size`] / `intrinsic_size`.
    Leaf,
}

/// Specification for a single node in the layout tree.
#[derive(Debug, Clone, PartialEq)]
pub struct LayoutNode {
    /// User handle that is mapped back in the result (typically a `WidgetId.index`).
    pub user_id: u64,
    /// Container kind (leaf vs. stack vs. grid).
    pub kind: NodeKind,
    /// Desired size along the main/cross axes. `None` means "auto".
    pub size: (SizeValue, SizeValue),
    /// Hard min/max constraints.
    pub constraints: Constraints,
    /// Intrinsic (content) size for `WrapContent`/`Leaf` resolution.
    /// `(0, 0)` means "no intrinsic content".
    pub intrinsic_size: crate::geometry::Vec2,
    /// Direct children (indices into the same [`LayoutTree`]).
    pub children: Vec<LayoutNodeId>,
    /// Per-child overrides: size value + constraints keyed by the child's
    /// *position* in `children`.
    pub child_overrides: Vec<(SizeValue, SizeValue, Constraints)>,
}

impl LayoutNode {
    pub fn leaf(user_id: u64) -> Self {
        Self {
            user_id,
            kind: NodeKind::Leaf,
            size: (SizeValue::default(), SizeValue::default()),
            constraints: Constraints::default(),
            intrinsic_size: crate::geometry::Vec2::ZERO,
            children: Vec::new(),
            child_overrides: Vec::new(),
        }
    }

    pub fn row(user_id: u64, axis: Axis) -> Self {
        Self {
            user_id,
            kind: NodeKind::Stack { axis, gap: 0.0 },
            size: (SizeValue::default(), SizeValue::default()),
            constraints: Constraints::default(),
            intrinsic_size: crate::geometry::Vec2::ZERO,
            children: Vec::new(),
            child_overrides: Vec::new(),
        }
    }

    pub fn column(user_id: u64) -> Self {
        Self::row(user_id, Axis::Vertical)
    }

    pub fn grid(user_id: u64, rows: usize, cols: usize, gap: f32) -> Self {
        Self {
            user_id,
            kind: NodeKind::Grid { rows, cols, gap },
            size: (SizeValue::default(), SizeValue::default()),
            constraints: Constraints::default(),
            intrinsic_size: crate::geometry::Vec2::ZERO,
            children: Vec::new(),
            child_overrides: Vec::new(),
        }
    }

    /// Adds a child and a corresponding size/constraint override.
    pub fn with_child(mut self, size: (SizeValue, SizeValue), constraints: Constraints, child: LayoutNodeId) -> Self {
        self.children.push(child);
        self.child_overrides.push((size.0, size.1, constraints));
        self
    }
}

/// Errors that can occur during layout resolution.
#[derive(Debug, Clone, PartialEq)]
pub enum LayoutError {
    /// The UI scale factor was zero, negative, NaN or infinity.
    InvalidScale(f32),
    /// A single-axis container had no children.
    EmptyContainer(LayoutNodeId),
    /// A `min` constraint exceeds the corresponding `max` constraint.
    ContradictoryConstraints {
        node: LayoutNodeId,
        axis: Axis,
        min: f32,
        max: f32,
    },
    /// A [`SizeValue::Percent`] was outside the `0.0..=1.0` range.
    InvalidPercent { node: LayoutNodeId, value: f32 },
    /// Grid child count did not match `rows * cols`.
    GridMismatch { node: LayoutNodeId, expected: usize, actual: usize },
    /// A percent-based size left the container size undetermined (under-constrained).
    UnderconstrainedSize { node: LayoutNodeId },
    /// A fixed-size child plus the gap exceeded the available space (over-constrained
    /// along the main axis with no flexible sibling to absorb the overflow).
    OverconstrainedMain { node: LayoutNodeId, overflow: f32 },
}

impl std::fmt::Display for LayoutError {
    fn fmt(&self, f: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
        match self {
            LayoutError::InvalidScale(v) => write!(f, "invalid ui scale factor: {v}"),
            LayoutError::EmptyContainer(n) => write!(f, "container node {n} has no children"),
            LayoutError::ContradictoryConstraints { node, axis, min, max } => {
                write!(f, "node {node} on {axis:?} has min ({min}) > max ({max})")
            }
            LayoutError::InvalidPercent { node, value } => {
                write!(f, "node {node} has percent value {value} outside 0..=1")
            }
            LayoutError::GridMismatch { node, expected, actual } => {
                write!(f, "grid node {node} expects {expected} children but has {actual}")
            }
            LayoutError::UnderconstrainedSize { node } => {
                write!(f, "node {node} size is under-constrained (percent with no container size)")
            }
            LayoutError::OverconstrainedMain { node, overflow } => {
                write!(f, "node {node} overflows main axis by {overflow} with no flexible child")
            }
        }
    }
}

impl std::error::Error for LayoutError {}

/// Output of a layout pass: a map from node index to physical `Rect`, plus any
/// errors encountered. The rect map always contains every successfully resolved
/// node; on error it contains entries resolved up to the failure point.
#[derive(Debug, Clone, Default)]
pub struct LayoutResult {
    pub rects: Vec<(LayoutNodeId, Rect)>,
    pub errors: Vec<(LayoutNodeId, LayoutError)>,
}

impl LayoutResult {
    pub fn get(&self, id: LayoutNodeId) -> Option<Rect> {
        self.rects.iter().find(|(i, _)| *i == id).map(|(_, r)| *r)
    }

    pub fn has_errors(&self) -> bool {
        !self.errors.is_empty()
    }
}

/// A table of nodes addressable by [`LayoutNodeId`].
#[derive(Debug, Clone, Default)]
pub struct LayoutTree {
    pub nodes: Vec<LayoutNode>,
}

impl LayoutTree {
    pub fn new() -> Self {
        Self { nodes: Vec::new() }
    }

    pub fn add(&mut self, node: LayoutNode) -> LayoutNodeId {
        let id = self.nodes.len();
        self.nodes.push(node);
        id
    }

    pub fn get(&self, id: LayoutNodeId) -> Option<&LayoutNode> {
        self.nodes.get(id)
    }
}

/// Resolved (logical) size and physical size for diagnostic output.
#[derive(Debug, Clone, Copy, PartialEq)]
struct ResolvedSize {
    logical: crate::geometry::Vec2,
    physical: crate::geometry::Vec2,
}

/// The layout engine. Configures a uniform `ui_scale` and performs layout.
#[derive(Debug, Clone, Copy)]
pub struct LayoutEngine {
    pub ui_scale: f32,
}

impl LayoutEngine {
    pub fn new(ui_scale: f32) -> Self {
        Self { ui_scale }
    }

    /// Validates the scale factor; returns an error if it is invalid.
    fn validate_scale(&self) -> Result<(), LayoutError> {
        let s = self.ui_scale;
        if s.is_nan() || s.is_infinite() || s <= 0.0 {
            return Err(LayoutError::InvalidScale(s));
        }
        Ok(())
    }

    /// Run a full layout pass over `tree` rooted at `root`, within the given
    /// logical `viewport` (the root's available space).
    ///
    /// Returns a [`LayoutResult`] containing resolved physical rects and any
    /// errors. Layout is deterministic and idempotent: two calls with identical
    /// inputs produce identical outputs.
    pub fn layout(&self, tree: &LayoutTree, root: LayoutNodeId, viewport: crate::geometry::Vec2) -> LayoutResult {
        let mut result = LayoutResult::default();

        if let Err(e) = self.validate_scale() {
            result.errors.push((root, e));
            return result;
        }

        // Pre-validate per-node constraints (cheap, catches contradictions early).
        for (i, node) in tree.nodes.iter().enumerate() {
            if let Err(e) = self.validate_node(i, node) {
                result.errors.push((i, e));
                // continue so we report as many as possible
            }
        }

        self.resolve_node(tree, root, crate::geometry::Vec2::ZERO, viewport, &mut result);
        result
    }

    /// Validate structural constraints of a single node.
    fn validate_node(&self, id: LayoutNodeId, node: &LayoutNode) -> Result<(), LayoutError> {
        // percent range
        let (sx, sy) = &node.size;
        if let SizeValue::Percent(v) = sx {
            if !(0.0..=1.0).contains(v) {
                return Err(LayoutError::InvalidPercent { node: id, value: *v });
            }
        }
        if let SizeValue::Percent(v) = sy {
            if !(0.0..=1.0).contains(v) {
                return Err(LayoutError::InvalidPercent { node: id, value: *v });
            }
        }
        // min/max contradictions
        let c = &node.constraints;
        if c.min_width > c.max_width && c.max_width > 0.0 || c.min_width > c.max_width {
            return Err(LayoutError::ContradictoryConstraints {
                node: id, axis: Axis::Horizontal, min: c.min_width, max: c.max_width,
            });
        }
        if c.min_height > c.max_height {
            return Err(LayoutError::ContradictoryConstraints {
                node: id, axis: Axis::Vertical, min: c.min_height, max: c.max_height,
            });
        }
        Ok(())
    }

    /// Recursively resolve a node. `origin` is the logical origin of the node's
    /// available space; `available` is the available size along each axis.
    fn resolve_node(
        &self,
        tree: &LayoutTree,
        id: LayoutNodeId,
        origin: crate::geometry::Vec2,
        available: crate::geometry::Vec2,
        result: &mut LayoutResult,
    ) {
        let Some(node) = tree.get(id) else {
            result.errors.push((id, LayoutError::EmptyContainer(id)));
            return;
        };

        let scale = self.ui_scale;
        let resolved = self.resolve_size(tree, id, node, node.size, available, result);

        // Apply constraints.
        let c = node.constraints;
        let raw = resolved.logical;
        let clamped = Vec2::new(
            apply_constraint(raw.x, c.min_width, c.max_width),
            apply_constraint(raw.y, c.min_height, c.max_height),
        );

        let physical_origin = origin * scale;
        let physical_size = clamped * scale;
        let rect = Rect::from_origin_size(physical_origin, physical_size);
        result.rects.push((id, rect));

        match node.kind {
            NodeKind::Leaf => {
                // no children
            }
            NodeKind::Stack { axis, gap } => {
                if node.children.is_empty() {
                    result.errors.push((id, LayoutError::EmptyContainer(id)));
                    return;
                }
                self.resolve_stack(tree, id, node, axis, gap, origin, clamped, result);
            }
            NodeKind::Grid { rows, cols, gap } => {
                let expected = rows.saturating_mul(cols);
                let actual = node.children.len();
                if expected != actual && actual != 0 {
                    result.errors.push((
                        id,
                        LayoutError::GridMismatch { node: id, expected, actual },
                    ));
                }
                self.resolve_grid(tree, id, node, rows, cols, gap, origin, clamped, result);
            }
        }
    }

    /// Resolve a node's raw size given the available space and its `SizeValue`.
    fn resolve_size(
        &self,
        tree: &LayoutTree,
        id: LayoutNodeId,
        node: &LayoutNode,
        sv: (SizeValue, SizeValue),
        available: crate::geometry::Vec2,
        result: &mut LayoutResult,
    ) -> ResolvedSize {
        let mut logical = Vec2::ZERO;

        for axis_idx in 0..2 {
            let avail = if axis_idx == 0 { available.x } else { available.y };
            let sv = if axis_idx == 0 { sv.0 } else { sv.1 };
            let val = match sv {
                SizeValue::Fixed(px) => px,
                SizeValue::Fill => avail,
                SizeValue::Percent(f) => {
                    if avail <= 0.0 {
                        result.errors.push((id, LayoutError::UnderconstrainedSize { node: id }));
                        0.0
                    } else {
                        avail * f
                    }
                }
                SizeValue::WrapContent => {
                    // Intrinsic content size along this axis.
                    let intrinsic = if axis_idx == 0 {
                        node.intrinsic_size.x
                    } else {
                        node.intrinsic_size.y
                    };
                    intrinsic.max(0.0)
                }
            };
            if axis_idx == 0 {
                logical.x = val;
            } else {
                logical.y = val;
            }
        }

        ResolvedSize {
            logical,
            physical: logical * self.ui_scale,
        }
    }

    /// Resolve a single-axis (row/column) container.
    fn resolve_stack(
        &self,
        tree: &LayoutTree,
        id: LayoutNodeId,
        node: &LayoutNode,
        axis: Axis,
        gap: f32,
        origin: crate::geometry::Vec2,
        size: crate::geometry::Vec2,
        result: &mut LayoutResult,
    ) {
        let n = node.children.len() as f32;
        let gaps = (n - 1.0).max(0.0) * gap;
        let main_dim = if axis == Axis::Horizontal { size.x } else { size.y };
        let cross_dim = if axis == Axis::Horizontal { size.y } else { size.x };
        let main_avail = (main_dim - gaps).max(0.0);

        // Two-pass: first assign fixed/wrap sizes, then distribute free space to
        // Fill children.
        let axis_idx = axis.as_usize();
        let cross_idx = axis.cross().as_usize();

        // First collect raw main-axis sizes.
        struct Slot {
            fixed_main: f32,
            is_fill: bool,
            child: LayoutNodeId,
            // resolved cross-axis size (available minus gaps)
            cross_size: f32,
        }
        let mut slots: Vec<Slot> = Vec::with_capacity(node.children.len());
        let cross_avail = (cross_dim - gaps).max(0.0);
        for (pos, &child_id) in node.children.iter().enumerate() {
            let child_node = match tree.get(child_id) {
                Some(n) => n,
                None => {
                    result.errors.push((child_id, LayoutError::EmptyContainer(child_id)));
                    continue;
                }
            };
            let (sv_main, sv_cross) = self.child_size(node, child_node, pos);
            let avail_vec = Vec2::new(
                if axis_idx == 0 { main_avail } else { cross_avail },
                if axis_idx == 1 { main_avail } else { cross_avail },
            );
            let resolved = self.resolve_size(tree, child_id, child_node, (sv_main, sv_cross), avail_vec, result);
            let c = child_node.constraints;
            let min_m = if axis_idx == 0 { c.min_width } else { c.min_height };
            let max_m = if axis_idx == 0 { c.max_width } else { c.max_height };
            let fixed_main = apply_constraint(resolved.logical.as_arr()[axis_idx], min_m, max_m);
            let cross_idx = axis.cross().as_usize();
            let min_c = if cross_idx == 0 { c.min_width } else { c.min_height };
            let max_c = if cross_idx == 0 { c.max_width } else { c.max_height };
            slots.push(Slot {
                fixed_main,
                is_fill: matches!(sv_main, SizeValue::Fill),
                child: child_id,
                cross_size: apply_constraint(resolved.logical.as_arr()[cross_idx], min_c, max_c),
            });
        }

        let fill_count = slots.iter().filter(|s| s.is_fill).count() as f32;
        let used_fixed: f32 = slots.iter().map(|s| s.fixed_main).sum::<f32>() + gaps;
        let free = (main_avail - slots.iter().map(|s| s.fixed_main).sum::<f32>()).max(0.0);

        let mut cursor = if axis == Axis::Horizontal { origin.x } else { origin.y };
        for (i, slot) in slots.iter().enumerate() {
            let child_main = if slot.is_fill {
                if fill_count > 0.0 { free / fill_count } else { 0.0 }
            } else {
                slot.fixed_main
            };
            // overflow check
            let total_allocated: f32 = slots.iter().map(|s| if s.is_fill { 0.0 } else { s.fixed_main }).sum::<f32>() + gaps + free;
            if total_allocated > main_avail + 1e-4 && i == slots.len() - 1 && fill_count == 0.0 {
                let overflow = total_allocated - main_avail;
                result.errors.push((id, LayoutError::OverconstrainedMain { node: id, overflow }));
            }

            let child_cross = slot.cross_size;
            let child_size = Vec2::new(
                if axis_idx == 0 { child_main } else { child_cross },
                if axis_idx == 1 { child_main } else { child_cross },
            );
            let child_origin = if axis == Axis::Horizontal {
                Vec2::new(cursor, origin.y + (cross_avail - child_cross) / 2.0)
            } else {
                Vec2::new(origin.x + (cross_avail - child_cross) / 2.0, cursor)
            };

            // Recurse.
            self.resolve_node(tree, slot.child, child_origin, child_size, result);

            cursor += child_main + gap;
        }
    }

    /// Resolve a grid container (row-major cell placement).
    fn resolve_grid(
        &self,
        tree: &LayoutTree,
        id: LayoutNodeId,
        node: &LayoutNode,
        rows: usize,
        cols: usize,
        gap: f32,
        origin: crate::geometry::Vec2,
        size: crate::geometry::Vec2,
        result: &mut LayoutResult,
    ) {
        let cols_f = cols as f32;
        let rows_f = rows as f32;
        let cell_w = (size.x - gap * (cols_f - 1.0).max(0.0)) / cols_f;
        let cell_h = (size.y - gap * (rows_f - 1.0).max(0.0)) / rows_f;
        for (pos, &child_id) in node.children.iter().enumerate() {
            let col = (pos as f32 % cols_f) as usize;
            let row = (pos as f32 / cols_f) as usize;
            let cx = origin.x + col as f32 * (cell_w + gap);
            let cy = origin.y + row as f32 * (cell_h + gap);
            let child_size = Vec2::new(cell_w, cell_h);
            self.resolve_node(tree, child_id, Vec2::new(cx, cy), child_size, result);
        }
    }

    /// Resolve a child's size override: explicit override takes priority over
    /// the child node's own `size`.
    fn child_size(&self, parent: &LayoutNode, child: &LayoutNode, pos: usize) -> (SizeValue, SizeValue) {
        if pos < parent.child_overrides.len() {
            let (sx, sy, _c) = &parent.child_overrides[pos];
            (*sx, *sy)
        } else {
            child.size
        }
    }
}

/// Extension trait so `Vec2` can be indexed by axis for convenience.
trait Vec2Ext {
    fn as_arr(&self) -> [f32; 2];
}

impl Vec2Ext for crate::geometry::Vec2 {
    fn as_arr(&self) -> [f32; 2] {
        [self.x, self.y]
    }
}

/// Apply a min/max clamp; if both bounds are active and `min > max`, clamp to
/// the min (the contradiction is reported separately by [`validate_node`]).
fn apply_constraint(val: f32, min: f32, max: f32) -> f32 {
    let v = if min > 0.0 { val.max(min) } else { val };
    if max > 0.0 { v.min(max) } else { v }
}

/// Apply an axis-indexed constraint pair.
fn apply_constraint_idx(val: f32, c: Constraints, axis_idx: usize) -> f32 {
    match axis_idx {
        0 => apply_constraint(val, c.min_width, c.max_width),
        _ => apply_constraint(val, c.min_height, c.max_height),
    }
}

#[cfg(test)]
mod tests;
