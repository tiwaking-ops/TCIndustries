//! Input event processing, hit testing, focus management, and pointer capture.
//!
//! Provides core input types (`InputEvent`, `PointerButton`, `Key`, `Modifiers`),
//! hit testing by z-order, pointer capture during drag operations, focus management,
//! Tab order cycling, Esc handling, and World/Data interaction modes with consumed
//! event reporting.

use crate::geometry::{Rect, Vec2};
use crate::arena::{WidgetArena, WidgetId};

/// Mouse or pointer buttons.
#[derive(Debug, Clone, Copy, PartialEq, Eq, Hash)]
pub enum PointerButton {
    Left,
    Right,
    Middle,
}

/// Keyboard modifier flags.
#[derive(Debug, Clone, Copy, Default, PartialEq, Eq, Hash)]
pub struct Modifiers {
    pub shift: bool,
    pub ctrl: bool,
    pub alt: bool,
}

/// Logical keys for focus and navigation.
#[derive(Debug, Clone, Copy, PartialEq, Eq, Hash)]
pub enum Key {
    Tab,
    Escape,
    Enter,
    Space,
    Backspace,
    ArrowUp,
    ArrowDown,
    ArrowLeft,
    ArrowRight,
    Other(u32),
}

/// Unified input event representation.
#[derive(Debug, Clone, PartialEq)]
pub enum InputEvent {
    PointerDown { pos: Vec2, button: PointerButton },
    PointerUp { pos: Vec2, button: PointerButton },
    PointerMove { pos: Vec2 },
    KeyDown { key: Key, modifiers: Modifiers },
    KeyUp { key: Key, modifiers: Modifiers },
    Char { ch: char },
}

/// Operating mode of the input dispatcher.
#[derive(Debug, Clone, Copy, PartialEq, Eq, Hash)]
pub enum InteractionMode {
    /// World mode: input can pass through to world/scene if not consumed by UI.
    World,
    /// Data mode: UI/modal input has strict focus; scene input is suppressed.
    Data,
}

/// Result of dispatching an input event.
#[derive(Debug, Clone, Copy, PartialEq, Eq)]
pub struct DispatchResult {
    /// Whether the event was consumed by a widget or system handler.
    pub consumed: bool,
    /// The widget ID that handled or captured the event, if any.
    pub target: Option<WidgetId>,
}

/// Central input manager handling focus, pointer capture, hit testing, and tab order.
#[derive(Debug, Clone)]
pub struct InputManager {
    /// Current interaction mode (World vs Data).
    pub mode: InteractionMode,
    /// Currently focused widget ID.
    pub focused_widget: Option<WidgetId>,
    /// Widget currently holding pointer capture (e.g. during drag).
    pub captured_widget: Option<WidgetId>,
    /// Ordered list of focusable widget IDs (for Tab order cycling).
    pub tab_order: Vec<WidgetId>,
    /// Last known pointer position.
    pub pointer_pos: Vec2,
}

impl InputManager {
    pub fn new() -> Self {
        Self {
            mode: InteractionMode::World,
            focused_widget: None,
            captured_widget: None,
            tab_order: Vec::new(),
            pointer_pos: Vec2::ZERO,
        }
    }

    /// Hit-tests a list of candidate widgets with their bounding rects and z-orders,
    /// returning the topmost live widget containing `pos`.
    ///
    /// Higher z-order takes precedence. Ties are broken by reverse order (last drawn / frontmost).
    pub fn hit_test(
        &self,
        arena: &WidgetArena,
        candidates: &[(WidgetId, Rect, i32)],
        pos: Vec2,
    ) -> Option<WidgetId> {
        let mut best: Option<(WidgetId, i32, usize)> = None;

        for (idx, &(id, rect, z)) in candidates.iter().enumerate() {
            // Verify node is live in arena
            if arena.get(id).is_none() {
                continue;
            }
            if rect.contains(pos) {
                match best {
                    None => best = Some((id, z, idx)),
                    Some((_, best_z, best_idx)) => {
                        // Higher z wins, or if equal z, higher index (topmost in draw order) wins
                        if z > best_z || (z == best_z && idx >= best_idx) {
                            best = Some((id, z, idx));
                        }
                    }
                }
            }
        }

        best.map(|(id, _, _)| id)
    }

    /// Cycles focus forward (Tab) or backward (Shift+Tab) through `tab_order`.
    pub fn cycle_tab(&mut self, reverse: bool) {
        if self.tab_order.is_empty() {
            return;
        }

        let current_pos = self.focused_widget.and_then(|id| {
            self.tab_order.iter().position(|&w| w == id)
        });

        let next_pos = match current_pos {
            Some(pos) => {
                if reverse {
                    if pos == 0 { self.tab_order.len() - 1 } else { pos - 1 }
                } else {
                    (pos + 1) % self.tab_order.len()
                }
            }
            None => {
                if reverse { self.tab_order.len() - 1 } else { 0 }
            }
        };

        self.focused_widget = self.tab_order.get(next_pos).copied();
    }

    /// Handles Escape key presses: clears focus or releases pointer capture.
    pub fn handle_escape(&mut self) -> bool {
        if self.captured_widget.is_some() {
            self.captured_widget = None;
            true
        } else if self.focused_widget.is_some() {
            self.focused_widget = None;
            true
        } else {
            false
        }
    }

    /// Dispatches an [`InputEvent`] through the interaction system, updating
    /// capture, focus, and returning a [`DispatchResult`].
    pub fn dispatch(
        &mut self,
        arena: &WidgetArena,
        candidates: &[(WidgetId, Rect, i32)],
        event: &InputEvent,
    ) -> DispatchResult {
        match *event {
            InputEvent::PointerMove { pos } => {
                self.pointer_pos = pos;
                if let Some(cap) = self.captured_widget {
                    if arena.get(cap).is_some() {
                        return DispatchResult { consumed: true, target: Some(cap) };
                    } else {
                        self.captured_widget = None;
                    }
                }
                let hit = self.hit_test(arena, candidates, pos);
                DispatchResult {
                    consumed: hit.is_some() && self.mode == InteractionMode::Data,
                    target: hit,
                }
            }
            InputEvent::PointerDown { pos, button } => {
                self.pointer_pos = pos;
                let hit = self.hit_test(arena, candidates, pos);
                if let Some(id) = hit {
                    self.focused_widget = Some(id);
                    if button == PointerButton::Left {
                        self.captured_widget = Some(id);
                    }
                    DispatchResult { consumed: true, target: Some(id) }
                } else {
                    // Clicked outside: clear focus if in data mode
                    if self.mode == InteractionMode::Data {
                        self.focused_widget = None;
                    }
                    DispatchResult { consumed: false, target: None }
                }
            }
            InputEvent::PointerUp { pos, button } => {
                self.pointer_pos = pos;
                let target = self.captured_widget.or_else(|| self.hit_test(arena, candidates, pos));
                if button == PointerButton::Left {
                    self.captured_widget = None;
                }
                DispatchResult {
                    consumed: target.is_some(),
                    target,
                }
            }
            InputEvent::KeyDown { key, modifiers } => {
                if key == Key::Escape {
                    let handled = self.handle_escape();
                    return DispatchResult { consumed: handled, target: self.focused_widget };
                }
                if key == Key::Tab {
                    self.cycle_tab(modifiers.shift);
                    return DispatchResult { consumed: true, target: self.focused_widget };
                }
                DispatchResult {
                    consumed: self.focused_widget.is_some(),
                    target: self.focused_widget,
                }
            }
            InputEvent::KeyUp { .. } => {
                DispatchResult {
                    consumed: self.focused_widget.is_some(),
                    target: self.focused_widget,
                }
            }
            InputEvent::Char { .. } => {
                DispatchResult {
                    consumed: self.focused_widget.is_some(),
                    target: self.focused_widget,
                }
            }
        }
    }
}

impl Default for InputManager {
    fn default() -> Self {
        Self::new()
    }
}

// InputManager unit tests can be added here or inline if needed.
