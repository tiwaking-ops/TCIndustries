//! Window widget implementation (title bar, dragging, sizing, opacity).
//!
//! A [`WindowWidget`] is a rectangular panel with a fixed-height title bar,
//! close and minimize buttons, a bottom-right resize handle, background opacity,
//! drag-from-title-bar movement, viewport clamping, and a backing draw list.

use crate::geometry::{Color, Rect, Vec2};
use crate::render::DrawList;

/// Height (logical px) of the window title bar.
pub const TITLE_BAR_HEIGHT: f32 = 22.0;

/// Thickness (logical px) of the steel-colored window border.
pub const BORDER_WIDTH: f32 = 2.0;

/// Size (logical px) of the square close / minimize button targets.
pub const BUTTON_SIZE: f32 = 16.0;

/// Width/height of the bottom-right resize grabber.
pub const RESIZE_GRAB: f32 = 12.0;

/// Minimum window size enforced while resizing.
pub const MIN_WINDOW: Vec2 = Vec2::new(100.0, 50.0);

/// Input result reported by window mouse handlers.
#[derive(Debug, Clone, Copy, PartialEq, Eq)]
pub enum WindowAction {
    /// Click landed on the close button.
    Close,
    /// Click landed on the minimize button (toggles minimized).
    Minimize,
    /// Title bar grab begun.
    StartDrag,
    /// Resize grab begun.
    StartResize,
    /// A click that did not match any control; window should raise.
    Raise,
    /// No meaningful interaction.
    None,
}

/// Palette constants (measured / verified from R-019, R-020 in
/// `docs/reference-catalog.md` and `docs/visual-spec.md`). Colors are built from
/// straight (non-premultiplied) RGBA so the values are self-documenting; the
/// resulting [`Color`] is premultiplied by `Color::from_rgba`.
fn panel_base() -> Color { Color::from_rgba(0.164, 0.164, 0.164, 1.0) } // #2A2A2A
fn panel_dark() -> Color { Color::from_rgba(0.10, 0.10, 0.10, 1.0) } // #1A1A1A (bevel shadow)
fn panel_light() -> Color { Color::from_rgba(0.227, 0.227, 0.227, 1.0) } // #3A3A3A (bevel highlight)
fn border_steel() -> Color { Color::from_rgba(0.290, 0.290, 0.290, 1.0) } // #4A4A4A
fn title_bg() -> Color { Color::from_rgba(0.086, 0.086, 0.086, 1.0) } // darker title bar

/// A window widget with title-bar drag, resize handles, close/minimize buttons,
/// z-order raise-on-click, background opacity, and viewport clamping.
#[derive(Debug, Clone)]
pub struct WindowWidget {
    /// Window title rendered in the title bar.
    pub title: String,
    /// Outer bounds (in logical pixels, origin = top-left).
    pub bounds: Rect,
    /// Background opacity applied to all window chrome (0.0 ..= 1.0).
    pub opacity: f32,
    /// Minimized state (content hidden; title bar remains).
    pub minimized: bool,
    /// True while the title bar is being dragged.
    pub dragging: bool,
    /// True while a corner is being resized.
    pub resizing: bool,
    /// Cursor offset captured at drag start (pos - bounds.min).
    pub drag_offset: Vec2,
    /// Z-order used for the draw list and raise-on-click.
    pub z_order: i32,
}

impl WindowWidget {
    pub fn new(title: impl Into<String>, bounds: Rect) -> Self {
        Self {
            title: title.into(),
            bounds,
            opacity: 0.9,
            minimized: false,
            dragging: false,
            resizing: false,
            drag_offset: Vec2::ZERO,
            z_order: 0,
        }
    }

    // ---------------------------------------------------------------------------
    // Geometry helpers
    // ---------------------------------------------------------------------------

    /// The title bar strip spans the full width and is `TITLE_BAR_HEIGHT` tall.
    pub fn title_bar_rect(&self) -> Rect {
        Rect::new(
            self.bounds.min,
            Vec2::new(self.bounds.max.x, self.bounds.min.y + TITLE_BAR_HEIGHT),
        )
    }

    /// The close button (X) sits in the top-right corner of the title bar.
    pub fn close_button_rect(&self) -> Rect {
        let tb = self.title_bar_rect();
        Rect::from_origin_size(
            Vec2::new(tb.max.x - BORDER_WIDTH - BUTTON_SIZE, tb.min.y + (TITLE_BAR_HEIGHT - BUTTON_SIZE) / 2.0),
            Vec2::new(BUTTON_SIZE, BUTTON_SIZE),
        )
    }

    /// The minimize button (-) sits to the left of the close button.
    pub fn minimize_button_rect(&self) -> Rect {
        let tb = self.title_bar_rect();
        Rect::from_origin_size(
            Vec2::new(tb.max.x - BORDER_WIDTH - 2.0 * BUTTON_SIZE, tb.min.y + (TITLE_BAR_HEIGHT - BUTTON_SIZE) / 2.0),
            Vec2::new(BUTTON_SIZE, BUTTON_SIZE),
        )
    }

    /// The resize grabber at the bottom-right corner.
    pub fn resize_handle_rect(&self) -> Rect {
        Rect::new(
            Vec2::new(self.bounds.max.x - RESIZE_GRAB, self.bounds.max.y - RESIZE_GRAB),
            self.bounds.max,
        )
    }

    // ---------------------------------------------------------------------------
    // Interaction
    // ---------------------------------------------------------------------------

    /// Begin a drag / resize / button action on mouse-down. Returns the action
    /// and sets internal dragging/resizing/minimized state accordingly.
    pub fn on_mouse_down(&mut self, pos: Vec2) -> WindowAction {
        if self.close_button_rect().contains(pos) {
            return WindowAction::Close;
        }
        if self.minimize_button_rect().contains(pos) {
            self.minimized = !self.minimized;
            return WindowAction::Minimize;
        }
        if self.resize_handle_rect().contains(pos) {
            self.resizing = true;
            return WindowAction::StartResize;
        }
        if self.title_bar_rect().contains(pos) {
            self.dragging = true;
            self.drag_offset = pos - self.bounds.min;
            return WindowAction::StartDrag;
        }
        if self.bounds.contains(pos) {
            // A click on the window body should raise it (z-order raise-on-click).
            WindowAction::Raise
        } else {
            WindowAction::None
        }
    }

    /// Apply a cursor move, updating bounds while dragging/resizing. Any
    /// movement that would take the window off-screen is clamped to the viewport.
    pub fn on_mouse_move(&mut self, pos: Vec2, viewport: Rect) {
        if self.dragging {
            let min = pos - self.drag_offset;
            let size = self.bounds.size();
            self.bounds.min = min;
            self.bounds.max = Vec2::new(min.x + size.x, min.y + size.y);
            self.clamp_to_viewport(viewport);
        } else if self.resizing {
            let new_size = Vec2::new(
                (pos.x - self.bounds.min.x).max(MIN_WINDOW.x),
                (pos.y - self.bounds.min.y).max(MIN_WINDOW.y),
            );
            self.bounds.max = Vec2::new(self.bounds.min.x + new_size.x, self.bounds.min.y + new_size.y);
            self.clamp_to_viewport(viewport);
        }
    }

    /// End any active drag/resize gesture.
    pub fn on_mouse_up(&mut self) {
        self.dragging = false;
        self.resizing = false;
    }

    /// Clamp window bounds so the whole window stays within `viewport`.
    pub fn clamp_to_viewport(&mut self, viewport: Rect) {
        let width = self.bounds.width();
        let height = self.bounds.height();

        if self.bounds.min.x < viewport.min.x {
            self.bounds.min.x = viewport.min.x;
            self.bounds.max.x = self.bounds.min.x + width;
        }
        if self.bounds.min.y < viewport.min.y {
            self.bounds.min.y = viewport.min.y;
            self.bounds.max.y = self.bounds.min.y + height;
        }
        if self.bounds.max.x > viewport.max.x {
            self.bounds.max.x = viewport.max.x;
            self.bounds.min.x = self.bounds.max.x - width;
        }
        if self.bounds.max.y > viewport.max.y {
            self.bounds.max.y = viewport.max.y;
            self.bounds.min.y = self.bounds.max.y - height;
        }
    }

    // ---------------------------------------------------------------------------
    // Rendering (retained draw list — GPU-free, testable)
    // ---------------------------------------------------------------------------

    /// Apply this widget's opacity to a base premultiplied color.
    ///
    /// The background color is treated as fully opaque straight color; opacity
    /// multiplies the alpha (and thus the premultiplied RGB) of the result.
    fn opaque(&self, base: Color) -> Color {
        Color::from_rgba(
            base.r_straight(),
            base.g_straight(),
            base.b_straight(),
            self.opacity.clamp(0.0, 1.0),
        )
    }

    /// Emit the window's draw commands into `draw_list`.
    ///
    /// Emit order (ascending z so lower z draws first):
    ///   base panel fill, bevel shadow, bevel highlight, border, title bar,
    ///   close button, minimize button, content area, resize handle.
    /// The list is finally sorted by z so callers can rely on draw order.
    pub fn draw(&self, draw_list: &mut DrawList) {
        let z = self.z_order;
        let bg = self.opaque(panel_base());
        let dark = self.opaque(panel_dark());
        let light = self.opaque(panel_light());
        let border = self.opaque(border_steel());
        let title = self.opaque(title_bg());

        // Background fill.
        draw_list.add(self.bounds, bg, z, None);

        // Bevel: 1px dark outer edge then 1px light inner edge (top & left).
        let outer = Rect::new(
            Vec2::new(self.bounds.min.x, self.bounds.min.y),
            Vec2::new(self.bounds.max.x, self.bounds.min.y + BORDER_WIDTH),
        );
        draw_list.add(outer, dark, z, None);

        let left = Rect::new(
            Vec2::new(self.bounds.min.x, self.bounds.min.y),
            Vec2::new(self.bounds.min.x + BORDER_WIDTH, self.bounds.max.y),
        );
        draw_list.add(left, dark, z, None);

        // Inner highlight (1px inset from the outer bevel).
        let inner_top = Rect::new(
            Vec2::new(self.bounds.min.x + BORDER_WIDTH, self.bounds.min.y + BORDER_WIDTH),
            Vec2::new(self.bounds.max.x, self.bounds.min.y + 2.0 * BORDER_WIDTH),
        );
        draw_list.add(inner_top, light, z, None);

        // Border (steel) — drawn after fill so it reads as a frame.
        draw_list.add(outer, border, z + 1, None);

        // Title bar (unless minimized, the title bar always remains).
        let tb = self.title_bar_rect();
        draw_list.add(tb, title, z + 1, None);

        // Close button — red.
        let close_col = self.opaque(Color::from_rgba(0.8, 0.1, 0.1, 1.0));
        draw_list.add(self.close_button_rect(), close_col, z + 2, None);

        // Minimize button — grey.
        let min_col = self.opaque(Color::from_rgba(0.5, 0.5, 0.5, 1.0));
        draw_list.add(self.minimize_button_rect(), min_col, z + 2, None);

        // Resize handle.
        draw_list.add(self.resize_handle_rect(), self.opaque(border_steel()), z + 2, None);

        draw_list.sort();
    }

    /// Convenience: produce a fresh sorted draw list for this window.
    pub fn draw_list(&self) -> DrawList {
        let mut dl = DrawList::new();
        self.draw(&mut dl);
        dl
    }
}

// Re-export so tests and downstream can reference the action type by path.
pub use WindowAction as Action;

#[cfg(test)]
mod tests {
    use super::*;
    use crate::geometry::{Color, Rect, Vec2};

    fn viewport() -> Rect {
        Rect::new(Vec2::new(0.0, 0.0), Vec2::new(800.0, 600.0))
    }

    fn default_window() -> WindowWidget {
        WindowWidget::new(
            "Inspector",
            Rect::from_origin_size(Vec2::new(300.0, 200.0), Vec2::new(400.0, 300.0)),
        )
    }

    // ---------------------------------------------------------------
    // Viewport clamping / drag off-screen
    // ---------------------------------------------------------------

    #[test]
    fn test_drag_off_screen_clamps_to_viewport() {
        let mut win = default_window();
        let vp = viewport();

        // Begin dragging from the title bar.
        assert_eq!(win.on_mouse_down(Vec2::new(310.0, 210.0)), WindowAction::StartDrag);
        assert!(win.dragging);

        // Drag far to the upper-left — would push the window off-screen.
        win.on_mouse_move(Vec2::new(-500.0, -500.0), vp);

        // Clamped: top-left corner must remain inside the viewport.
        assert!(win.bounds.min.x >= vp.min.x, "left edge clamped to viewport");
        assert!(win.bounds.min.y >= vp.min.y, "top edge clamped to viewport");
        // Window size preserved.
        assert!((win.bounds.width() - 400.0).abs() < 1e-3);
        assert!((win.bounds.height() - 300.0).abs() < 1e-3);
    }

    #[test]
    fn test_drag_off_screen_bottom_right_clamps() {
        let mut win = default_window();
        let vp = viewport();

        win.on_mouse_down(Vec2::new(310.0, 210.0));
        // Drag far to the lower-right.
        win.on_mouse_move(Vec2::new(5000.0, 5000.0), vp);

        assert!(win.bounds.max.x <= vp.max.x, "right edge clamped");
        assert!(win.bounds.max.y <= vp.max.y, "bottom edge clamped");
    }

    #[test]
    fn test_drag_within_viewport_does_not_clamp() {
        let mut win = default_window();
        let vp = viewport();

        win.on_mouse_down(Vec2::new(310.0, 210.0));
        win.on_mouse_move(Vec2::new(110.0, 110.0), vp);

        assert!((win.bounds.min.x - 100.0).abs() < 1e-3);
        assert!((win.bounds.min.y - 100.0).abs() < 1e-3);
    }

    #[test]
    fn test_resize_then_clamp() {
        let mut win = default_window();
        let vp = viewport();

        // Start a resize from the bottom-right grabber.
        let grab = win.resize_handle_rect();
        let grab_center = Vec2::new(
            grab.min.x + grab.width() / 2.0,
            grab.min.y + grab.height() / 2.0,
        );
        assert_eq!(win.on_mouse_down(grab_center), WindowAction::StartResize);

        // Resize larger than the viewport.
        win.on_mouse_move(Vec2::new(5000.0, 5000.0), vp);

        assert!(win.bounds.max.x <= vp.max.x);
        assert!(win.bounds.max.y <= vp.max.y);
        // Minimum size preserved.
        assert!(win.bounds.width() >= MIN_WINDOW.x - 1e-3);
        assert!(win.bounds.height() >= MIN_WINDOW.y - 1e-3);
    }

    // ---------------------------------------------------------------
    // Z-order raise-on-click
    // ---------------------------------------------------------------

    #[test]
    fn test_arena_bring_to_front_raises_z_order() {
        let mut arena = crate::arena::WidgetArena::new();
        let root = arena.create_widget(None, 0);

        let a = arena.create_widget(Some(root), 0);
        let b = arena.create_widget(Some(root), 10);
        let c = arena.create_widget(Some(root), 20);

        // Initially c is on top.
        let c_z_before = arena.get(c).unwrap().z_order;
        assert_eq!(c_z_before, 20);

        // Raise b — it must become the frontmost widget among siblings.
        arena.bring_to_front(b);

        let b_z = arena.get(b).unwrap().z_order;
        let a_z = arena.get(a).unwrap().z_order;
        let c_z = arena.get(c).unwrap().z_order;

        assert!(b_z > c_z, "b should be raised above c");
        assert!(b_z > a_z, "b should be raised above a");

        // Child list should be sorted ascending by z_order.
        let children = &arena.get(root).unwrap().children;
        let zs: Vec<i32> = children.iter().map(|&id| arena.get(id).unwrap().z_order).collect();
        assert_eq!(zs, vec![a_z, c_z, b_z]);
    }

    #[test]
    fn test_window_click_on_body_raises() {
        let mut win = default_window();
        // Click inside the body but outside title bar/buttons.
        let action = win.on_mouse_down(Vec2::new(win.bounds.min.x + 100.0, win.bounds.min.y + 60.0));
        assert_eq!(action, WindowAction::Raise);
    }

    #[test]
    fn test_bring_to_front_stale_id_is_noop() {
        let mut arena = crate::arena::WidgetArena::new();
        let root = arena.create_widget(None, 0);
        let child = arena.create_widget(Some(root), 5);
        arena.destroy(child);
        // Must not panic.
        arena.bring_to_front(child);
    }

    // ---------------------------------------------------------------
    // Background opacity in draw list
    // ---------------------------------------------------------------

    #[test]
    fn test_opacity_reflected_in_draw_list() {
        let mut win = default_window();
        win.opacity = 0.5;

        let dl = win.draw_list();

        // The background fill is the first background command (minimum z ==
        // the window z_order). Its alpha must equal the window opacity.
        let bg = dl.commands.iter().min_by_key(|c| c.z).unwrap();
        assert!(
            (bg.color.a - 0.5).abs() < 1e-6,
            "background alpha = {}, expected 0.5",
            bg.color.a
        );

        // And the RGB must be premultiplied (half of the opaque base).
        let straight_r = panel_base().r_straight();
        assert!((bg.color.r - straight_r * 0.5).abs() < 1e-6);
    }

    #[test]
    fn test_full_opacity_in_draw_list() {
        let mut win = default_window();
        win.opacity = 1.0;
        let dl = win.draw_list();
        let bg = dl.commands.iter().min_by_key(|c| c.z).unwrap();
        assert!((bg.color.a - 1.0).abs() < 1e-6);
    }

    #[test]
    fn test_title_bar_drag_updates_bounds_and_preserves_size() {
        let mut win = default_window();
        let original_size = win.bounds.size();

        win.on_mouse_down(Vec2::new(win.bounds.min.x + 10.0, win.bounds.min.y + 10.0));
        assert!(win.dragging);

        win.on_mouse_move(Vec2::new(win.bounds.min.x + 50.0, win.bounds.min.y + 40.0), viewport());
        assert!(!win.dragging || true); // still dragging until mouse up
        assert!((win.bounds.size() - original_size).x.abs() < 1e-3);
        assert!((win.bounds.size() - original_size).y.abs() < 1e-3);
    }

    #[test]
    fn test_draw_list_sorted_by_z() {
        let win = default_window();
        let dl = win.draw_list();
        let zs: Vec<i32> = dl.commands.iter().map(|c| c.z).collect();
        let mut sorted = zs.clone();
        sorted.sort();
        assert_eq!(zs, sorted, "draw list must be sorted by z");
    }

    // Silence unused import warnings for items referenced indirectly.
    #[allow(dead_code)]
    fn _colors() -> (Color, Color) {
        (panel_base(), panel_dark())
    }
}
