# `tcgui` Usage Guide (`docs/usage.md`)

**Status:** Complete (Task T-91)  
**Scope:** Guide for adding a screen, adding a widget, binding to the model, and connecting a real backend.

---

## 1. Adding a Screen

Screens in `tcgui` are managed by `GuiScreens` in `tcgui-core::screens`. Each screen handles its own state updates and event responses.

1. Define a new variant or state struct for your screen (e.g., `MyScreenState`).
2. Add a handling method in `GuiScreens` for your screen's lifecycle events:
   - `on_enter(&mut self)`
   - `update(&mut self, dt: f32)`
   - `handle_event(&mut self, event: BackendEvent)`
3. Render the screen by emitting `DrawCommand`s or loading a declarative layout (`.ui.json`).

---

## 2. Adding a Widget

To add a new custom widget:

1. Implement the `Widget` trait (`tcgui_core::widget::Widget`):
   - `layout(&mut self, constraints)`
   - `draw(&self, draw_list)`
   - `handle_event(&mut self, event)`
2. Register the widget type string in the widget registry (`WidgetKind` / `layout_loader.rs`) so it can be instantiated declaratively from `.ui.json` layouts.
3. Add unit tests for click handling, state changes, and zero-leak cleanup in the widget's test module.

---

## 3. Data Binding via `UiModel`

`tcgui-core::model::UiModel` provides a reactive key-value store with typed bindings:

```rust
use tcgui_core::model::{UiModel, UiValue};

let mut model = UiModel::new();
model.set("player.health", UiValue::Number(100.0));

// Subscribe to changes
let subscription_id = model.subscribe("player.health", |val| {
    println!("Health updated: {:?}", val);
});

model.set("player.health", UiValue::Number(85.0));
```

When widgets are destroyed, their model subscriptions are automatically unsubscribed (`stats().bindings` decreases).

---

## 4. Connecting a Real Backend

`tcgui-core::backend` defines the asynchronous client-server communication contract (`BackendConnection`):

- **Client Commands (`ClientCommand`):** Sent from UI actions (e.g. login, movement, chat message, crafting start).
- **Backend Events (`BackendEvent`):** Received from server (e.g. inventory updates, HAM changes, combat state).
- **`MockBackend`:** A deterministic reference implementation used for testing screens without a live server connection.

To wire up a real network connection:
1. Implement the `BackendConnection` trait wrapping your TCP/WebSockets transport.
2. Feed received `BackendEvent`s into `GuiScreens` / `UiModel`.
3. Dispatch `ClientCommand`s to the server transport layer upon user interaction.
