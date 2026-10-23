//! Core library for `tcgui-core`: widgets, layout, input, theming, data binding, screens.

pub mod geometry;
pub mod arena;
pub mod render;
pub mod widget;
pub mod window;
pub mod ham;
pub mod progressbar;
pub mod backend;
pub mod screens;
pub mod theme;
pub mod layout_loader;
pub mod textinput;
pub mod list;
pub mod stats;

pub use geometry::{Vec2, Rect, Color, blend_colors};
pub use arena::{WidgetId, WidgetNode, WidgetArena, ArenaStats};
pub use render::{TextureId, DrawList, DrawCommand};
pub use widget::{
    Widget, WidgetState, Command, WidgetKind, Button, Checkbox, Toggle,
    Label, Panel, ImageIcon
};
pub use window::{WindowWidget, WindowAction};
pub use ham::HamBarWidget;
pub use progressbar::ProgressBar;
pub use textinput::{TextInput, TextInputAction};
pub use list::{ListWidget, GridWidget, DragPayload};
pub use backend::{BackendEvent, ClientCommand, InventoryItemInfo, BackendConnection, MockBackend};
pub use screens::{GuiScreens, ScreenState};
pub use theme::{
    PaletteFile, PaletteToken, ThemeFile, FontConfig, NineSliceConfig, StyleConfig,
    ThemeError, load_palette_from_str, load_theme_from_str, generate_schemas,
};
pub use layout_loader::{
    UiWidgetJson, UiAxisSpec, UiLayoutSpec, LayoutLoadError,
    load_layout_from_str, generate_layout_schema,
};
pub use stats::SoakTestHarness;

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_arena_lifecycle_and_stale_ids() {
        let mut arena = WidgetArena::new();
        let id1 = arena.create_widget(None, 0);
        assert!(arena.get(id1).is_some());

        arena.destroy(id1);
        assert!(arena.get(id1).is_none());

        for _ in 0..10_000 {
            let id = arena.create_widget(None, 0);
            arena.destroy(id);
        }
        assert_eq!(arena.stats().widgets, 0);
    }
}
