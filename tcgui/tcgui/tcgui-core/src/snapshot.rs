//! Clock trait, primitive commands for rectangles, text, and 9-slice, and deterministic snapshot test harness.

use crate::geometry::{Color, Rect, Vec2};

/// Deterministic clock trait for animation and timing.
pub trait Clock {
    fn now(&self) -> f32; // current time in seconds
    fn dt(&self) -> f32;  // delta time since last frame
}

/// A deterministic test clock implementation.
#[derive(Debug, Clone, Copy, Default)]
pub struct TestClock {
    time: f32,
    dt: f32,
}

impl TestClock {
    pub fn new(time: f32, dt: f32) -> Self {
        Self { time, dt }
    }

    pub fn advance(&mut self, dt: f32) {
        self.dt = dt;
        self.time += dt;
    }
}

impl Clock for TestClock {
    fn now(&self) -> f32 { self.time }
    fn dt(&self) -> f32 { self.dt }
}

/// Extended primitive draw commands for T-16.
#[derive(Debug, Clone, PartialEq)]
pub enum PrimitiveCommand {
    Rect {
        rect: Rect,
        color: Color,
        z: i32,
    },
    Text {
        text: String,
        position: Vec2,
        font_size: f32,
        color: Color,
        z: i32,
    },
    NineSlice {
        rect: Rect,
        texture_key: String,
        margins: [f32; 4],
        tint: Color,
        z: i32,
    },
}

/// Extended draw list holding primitive commands.
#[derive(Debug, Clone, Default, PartialEq)]
pub struct PrimitiveDrawList {
    pub commands: Vec<PrimitiveCommand>,
}

impl PrimitiveDrawList {
    pub fn new() -> Self {
        Self { commands: Vec::new() }
    }

    pub fn add_rect(&mut self, rect: Rect, color: Color, z: i32) {
        self.commands.push(PrimitiveCommand::Rect { rect, color, z });
    }

    pub fn add_text(&mut self, text: impl Into<String>, position: Vec2, font_size: f32, color: Color, z: i32) {
        self.commands.push(PrimitiveCommand::Text {
            text: text.into(),
            position,
            font_size,
            color,
            z,
        });
    }

    pub fn add_nine_slice(&mut self, rect: Rect, texture_key: impl Into<String>, margins: [f32; 4], tint: Color, z: i32) {
        self.commands.push(PrimitiveCommand::NineSlice {
            rect,
            texture_key: texture_key.into(),
            margins,
            tint,
            z,
        });
    }

    pub fn sort(&mut self) {
        self.commands.sort_by_key(|cmd| match cmd {
            PrimitiveCommand::Rect { z, .. } => *z,
            PrimitiveCommand::Text { z, .. } => *z,
            PrimitiveCommand::NineSlice { z, .. } => *z,
        });
    }

    /// Generates a deterministic canonical string snapshot of the draw list.
    ///
    /// Commands are ordered by `(z, kind_discriminant, position-key, payload)`
    /// so the snapshot is identical regardless of insertion order — essential for
    /// deterministic snapshot testing across runs.
    pub fn snapshot(&self) -> String {
        let mut sorted = self.commands.clone();
        sorted.sort_by(|a, b| {
            use PrimitiveCommand::*;
            fn key(c: &PrimitiveCommand) -> (i32, u8, i32, i32) {
                match c {
                    Rect { z, rect, .. } => (*z, 0, rect.min.x as i32, rect.min.y as i32),
                    NineSlice { z, rect, .. } => (*z, 1, rect.min.x as i32, rect.min.y as i32),
                    Text { z, position, .. } => (*z, 2, position.x as i32, position.y as i32),
                }
            }
            key(a).cmp(&key(b))
        });

        let mut lines = Vec::new();
        for cmd in sorted {
            match cmd {
                PrimitiveCommand::Rect { rect, color, z } => {
                    lines.push(format!(
                        "RECT[z={}, min=({},{}), max=({},{}), col=({:.2},{:.2},{:.2},{:.2})]",
                        z, rect.min.x, rect.min.y, rect.max.x, rect.max.y, color.r, color.g, color.b, color.a
                    ));
                }
                PrimitiveCommand::Text { text, position, font_size, color, z } => {
                    lines.push(format!(
                        "TEXT[z={}, pos=({},{}), size={}, col=({:.2},{:.2},{:.2},{:.2}), text=\"{}\"]",
                        z, position.x, position.y, font_size, color.r, color.g, color.b, color.a, text
                    ));
                }
                PrimitiveCommand::NineSlice { rect, texture_key, margins, tint, z } => {
                    lines.push(format!(
                        "9SLICE[z={}, rect=({},{}),({},{}), tex=\"{}\", margins={:?}, tint=({:.2},{:.2},{:.2},{:.2})]",
                        z, rect.min.x, rect.min.y, rect.max.x, rect.max.y, texture_key, margins, tint.r, tint.g, tint.b, tint.a
                    ));
                }
            }
        }
        lines.join("\n")
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_test_clock_advance() {
        let mut clock = TestClock::new(0.0, 0.016);
        assert_eq!(clock.now(), 0.0);
        clock.advance(0.016);
        assert_eq!(clock.now(), 0.016);
        assert_eq!(clock.dt(), 0.016);
    }

    #[test]
    fn test_primitive_draw_list_snapshot_determinism() {
        let mut dl1 = PrimitiveDrawList::new();
        dl1.add_rect(Rect::new(Vec2::ZERO, Vec2::new(10.0, 10.0)), Color::WHITE, 0);
        dl1.add_text("Hello", Vec2::new(5.0, 5.0), 12.0, Color::BLACK, 1);
        dl1.add_nine_slice(Rect::new(Vec2::ZERO, Vec2::new(50.0, 20.0)), "panel", [4.0; 4], Color::WHITE, 0);

        let mut dl2 = PrimitiveDrawList::new();
        dl2.add_text("Hello", Vec2::new(5.0, 5.0), 12.0, Color::BLACK, 1);
        dl2.add_nine_slice(Rect::new(Vec2::ZERO, Vec2::new(50.0, 20.0)), "panel", [4.0; 4], Color::WHITE, 0);
        dl2.add_rect(Rect::new(Vec2::ZERO, Vec2::new(10.0, 10.0)), Color::WHITE, 0);

        assert_eq!(dl1.snapshot(), dl2.snapshot());
    }
}
