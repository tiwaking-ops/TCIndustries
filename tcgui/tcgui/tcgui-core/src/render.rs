//! Rendering primitives and types shared between `tcgui-core` and `tcgui-render`.
//!
//! These types are intentionally minimal and GPU-free so that the core crate
//! can be tested without a display or a device context.

use crate::geometry::{Color, Rect};

/// Opaque identifier for a texture held by the renderer's resource tracker.
#[derive(Debug, Clone, Copy, PartialEq, Eq, Hash)]
pub struct TextureId {
    pub index: u32,
}

/// A single draw command for retained-mode software / test rendering.
#[derive(Debug, Clone, PartialEq)]
pub struct DrawCommand {
    pub rect: Rect,
    pub color: Color,
    pub z: i32,
    pub texture: Option<TextureId>,
}

/// A draw list collecting sorted draw commands.
#[derive(Debug, Clone, Default, PartialEq)]
pub struct DrawList {
    pub commands: Vec<DrawCommand>,
}

impl DrawList {
    pub fn new() -> Self {
        Self { commands: Vec::new() }
    }

    pub fn add(&mut self, rect: Rect, color: Color, z: i32, texture: Option<TextureId>) {
        self.commands.push(DrawCommand { rect, color, z, texture });
    }

    pub fn sort(&mut self) {
        self.commands.sort_by_key(|cmd| cmd.z);
    }
}
