//! Progress bar widget implementation.
//!
//! Renders a horizontal progress bar with a customizable fill color and
//! light background.

use crate::geometry::{Color, Rect, Vec2};
use crate::render::DrawList;

/// A progress bar widget displaying a fractional completion (0.0 ..= 1.0).
#[derive(Debug, Clone)]
pub struct ProgressBar {
    pub progress: f32,
    pub bounds: Rect,
    pub fill_color: Color,
    pub bg_color: Color,
}

impl ProgressBar {
    pub fn new(progress: f32, bounds: Rect, fill_color: Color) -> Self {
        Self {
            progress: progress.clamp(0.0, 1.0),
            bounds,
            fill_color,
            bg_color: Color::from_rgba(0.9, 0.9, 0.9, 1.0), // white/light background per spec
        }
    }

    /// Emit draw commands for the progress bar.
    pub fn draw(&self, draw_list: &mut DrawList, z: i32) {
        // Background
        draw_list.add(self.bounds, self.bg_color, z, None);

        // Fill bar proportional to progress ratio
        let w = self.bounds.width() * self.progress.clamp(0.0, 1.0);
        if w > 0.0 {
            let fill_rect = Rect::new(
                self.bounds.min,
                Vec2::new(self.bounds.min.x + w, self.bounds.max.y),
            );
            draw_list.add(fill_rect, self.fill_color, z + 1, None);
        }
    }

    pub fn draw_list(&self, z: i32) -> DrawList {
        let mut dl = DrawList::new();
        self.draw(&mut dl, z);
        dl.sort();
        dl
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_progress_bar_clamping() {
        let pb_over = ProgressBar::new(1.5, Rect::new(Vec2::ZERO, Vec2::new(100.0, 10.0)), Color::WHITE);
        assert_eq!(pb_over.progress, 1.0);

        let pb_under = ProgressBar::new(-0.5, Rect::new(Vec2::ZERO, Vec2::new(100.0, 10.0)), Color::WHITE);
        assert_eq!(pb_under.progress, 0.0);
    }

    #[test]
    fn test_progress_bar_draw_list_commands() {
        let pb = ProgressBar::new(0.5, Rect::new(Vec2::ZERO, Vec2::new(200.0, 20.0)), Color::BLACK);
        let dl = pb.draw_list(0);
        assert_eq!(dl.commands.len(), 2); // background + fill
    }
}
