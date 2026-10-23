//! Health, Action, Mind (HAM) bar widget implementation.
//!
//! Renders the three classic pre-CU Star Wars Galaxies status bars (Health,
//! Action, Mind) stacked vertically with a white background, colored fills, and
//! black vertical wound notches (R-001, visual-spec.md section 3).

use crate::geometry::{Color, Rect, Vec2};
use crate::render::DrawList;

/// A composite HAM (Health, Action, Mind) bar widget.
#[derive(Debug, Clone)]
pub struct HamBarWidget {
    pub health_current: f32,
    pub health_max: f32,
    pub health_wounds: f32,
    pub action_current: f32,
    pub action_max: f32,
    pub action_wounds: f32,
    pub mind_current: f32,
    pub mind_max: f32,
    pub mind_wounds: f32,
    pub bounds: Rect,
}

impl HamBarWidget {
    #[allow(clippy::too_many_arguments)]
    pub fn new(
        health_current: f32, health_max: f32, health_wounds: f32,
        action_current: f32, action_max: f32, action_wounds: f32,
        mind_current: f32, mind_max: f32, mind_wounds: f32,
        bounds: Rect,
    ) -> Self {
        Self {
            health_current, health_max, health_wounds,
            action_current, action_max, action_wounds,
            mind_current, mind_max, mind_wounds,
            bounds,
        }
    }

    /// Single bar ratio (current / max), clamped 0.0..=1.0.
    pub fn health_ratio(&self) -> f32 {
        if self.health_max <= 0.0 { 0.0 } else { (self.health_current / self.health_max).clamp(0.0, 1.0) }
    }

    pub fn action_ratio(&self) -> f32 {
        if self.action_max <= 0.0 { 0.0 } else { (self.action_current / self.action_max).clamp(0.0, 1.0) }
    }

    pub fn mind_ratio(&self) -> f32 {
        if self.mind_max <= 0.0 { 0.0 } else { (self.mind_current / self.mind_max).clamp(0.0, 1.0) }
    }

    /// Emit draw commands into `draw_list`.
    ///
    /// Per `docs/visual-spec.md` Section 3:
    /// - Bar Height: 12px per bar
    /// - Spacing: 2px vertical gap between bars
    /// - Wound Notch Width: 2px black vertical line indicating wound threshold
    /// - Background: White (`#FFFFFF`) with colored fill overlay from left to right.
    /// - Colors: Health (`#CC2222`), Action (`#22CC22`), Mind (`#2266CC`).
    pub fn draw(&self, draw_list: &mut DrawList, z: i32) {
        let bar_h = 12.0;
        let gap = 2.0;
        let total_w = self.bounds.width();
        let mut y = self.bounds.min.y;

        let bars = [
            (self.health_current, self.health_max, self.health_wounds, Color::from_rgba(0.8, 0.133, 0.133, 1.0)), // Health #CC2222
            (self.action_current, self.action_max, self.action_wounds, Color::from_rgba(0.133, 0.8, 0.133, 1.0)), // Action #22CC22
            (self.mind_current, self.mind_max, self.mind_wounds, Color::from_rgba(0.133, 0.4, 0.8, 1.0)),   // Mind #2266CC
        ];

        for (curr, max, wounds, color) in bars {
            let bar_rect = Rect::new(
                Vec2::new(self.bounds.min.x, y),
                Vec2::new(self.bounds.min.x + total_w, y + bar_h),
            );

            // 1. White background (#FFFFFF)
            draw_list.add(bar_rect, Color::WHITE, z, None);

            // 2. Colored fill overlay
            let eff_max = max.max(1.0);
            let ratio = (curr / eff_max).clamp(0.0, 1.0);
            let fill_w = total_w * ratio;
            if fill_w > 0.0 {
                let fill_rect = Rect::new(
                    bar_rect.min,
                    Vec2::new(bar_rect.min.x + fill_w, bar_rect.max.y),
                );
                draw_list.add(fill_rect, color, z + 1, None);
            }

            // 3. Wound notch (2px black vertical line at wound threshold position from right)
            if wounds > 0.0 {
                let wound_ratio = (wounds / eff_max).clamp(0.0, 1.0);
                let notch_x = bar_rect.min.x + total_w * (1.0 - wound_ratio);
                let notch_rect = Rect::new(
                    Vec2::new(notch_x, bar_rect.min.y),
                    Vec2::new(notch_x + 2.0, bar_rect.max.y),
                );
                draw_list.add(notch_rect, Color::BLACK, z + 2, None);
            }

            y += bar_h + gap;
        }

        draw_list.sort();
    }

    pub fn draw_list(&self, z: i32) -> DrawList {
        let mut dl = DrawList::new();
        self.draw(&mut dl, z);
        dl
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    fn sample_bounds() -> Rect {
        Rect::new(Vec2::ZERO, Vec2::new(100.0, 40.0))
    }

    #[test]
    fn test_ham_bar_ratios() {
        let ham = HamBarWidget::new(
            50.0, 100.0, 10.0,
            75.0, 100.0, 0.0,
            25.0, 100.0, 25.0,
            sample_bounds(),
        );
        assert_eq!(ham.health_ratio(), 0.5);
        assert_eq!(ham.action_ratio(), 0.75);
        assert_eq!(ham.mind_ratio(), 0.25);
    }

    // Snapshot tests at 0%, 50%, 100% with wounds

    #[test]
    fn test_ham_snapshot_100_percent_with_wounds() {
        let ham = HamBarWidget::new(
            1000.0, 1000.0, 200.0,
            1000.0, 1000.0, 100.0,
            1000.0, 1000.0, 0.0,
            sample_bounds(),
        );
        let dl = ham.draw_list(0);
        // Health (bg + fill + notch = 3), Action (bg + fill + notch = 3), Mind (bg + fill = 2) -> total 8 commands
        assert_eq!(dl.commands.len(), 8);

        // Verify white backgrounds are present
        let white_count = dl.commands.iter().filter(|c| c.color == Color::WHITE).count();
        assert_eq!(white_count, 3);

        // Verify wound notch count (health and action have wounds > 0, mind = 0)
        let black_notches = dl.commands.iter().filter(|c| c.color == Color::BLACK).count();
        assert_eq!(black_notches, 2);
    }

    #[test]
    fn test_ham_snapshot_50_percent_with_wounds() {
        let ham = HamBarWidget::new(
            500.0, 1000.0, 100.0,
            500.0, 1000.0, 100.0,
            500.0, 1000.0, 100.0,
            sample_bounds(),
        );
        let dl = ham.draw_list(0);
        assert_eq!(dl.commands.len(), 9); // 3 bg + 3 fill + 3 notches
    }

    #[test]
    fn test_ham_snapshot_0_percent_no_wounds() {
        let ham = HamBarWidget::new(
            0.0, 1000.0, 0.0,
            0.0, 1000.0, 0.0,
            0.0, 1000.0, 0.0,
            sample_bounds(),
        );
        let dl = ham.draw_list(0);
        // 0% current means no fill quads (3 bg + 0 fill + 0 notches = 3 commands)
        assert_eq!(dl.commands.len(), 3);
        for cmd in &dl.commands {
            assert_eq!(cmd.color, Color::WHITE);
        }
    }
}
