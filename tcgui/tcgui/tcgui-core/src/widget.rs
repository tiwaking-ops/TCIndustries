//! Widget trait and basic widget implementations.
//!
//! Defines the `Widget` trait, a `WidgetKind` enumeration for the widget type
//! registry, and concrete widgets: `Panel`, `Label`, `ImageIcon`, `Button`,
//! `Checkbox`, `Toggle`.

use crate::geometry::{Color, Rect, Vec2};

#[derive(Debug, Clone, Copy, PartialEq, Eq, Hash)]
pub enum WidgetState {
    Normal,
    Hover,
    Pressed,
    Disabled,
}

#[derive(Debug, Clone, PartialEq, Eq, Hash)]
pub struct Command {
    pub target: crate::arena::WidgetId,
    pub action: String,
}

pub trait Widget: std::fmt::Debug {
    fn bounds(&self) -> Rect;
    fn state(&self) -> WidgetState {
        WidgetState::Normal
    }
    fn update(&mut self, _dt: f32) {}
}

#[derive(Debug, Clone, Copy, PartialEq, Eq, Hash)]
pub enum WidgetKind {
    Panel,
    Label,
    Image,
    Button,
    Checkbox,
    Toggle,
    Window,
    ProgressBar,
    HamBar,
    Slider,
    TextInput,
    ScrollView,
    TabBar,
    List,
    Grid,
    Tooltip,
    Radar,
    ChatLog,
}

impl WidgetKind {
    pub fn registered_names() -> &'static [(&'static str, WidgetKind)] {
        &[
            ("panel", WidgetKind::Panel),
            ("label", WidgetKind::Label),
            ("image", WidgetKind::Image),
            ("button", WidgetKind::Button),
            ("checkbox", WidgetKind::Checkbox),
            ("toggle", WidgetKind::Toggle),
            ("window", WidgetKind::Window),
            ("progress_bar", WidgetKind::ProgressBar),
            ("ham_bar", WidgetKind::HamBar),
            ("slider", WidgetKind::Slider),
            ("text_input", WidgetKind::TextInput),
            ("scroll_view", WidgetKind::ScrollView),
            ("tab_bar", WidgetKind::TabBar),
            ("list", WidgetKind::List),
            ("grid", WidgetKind::Grid),
            ("tooltip", WidgetKind::Tooltip),
            ("radar", WidgetKind::Radar),
            ("chat_log", WidgetKind::ChatLog),
        ]
    }

    pub fn from_tag(tag: &str) -> Option<WidgetKind> {
        Self::registered_names()
            .iter()
            .find(|(name, _)| *name == tag)
            .map(|(_, kind)| *kind)
    }
}

#[derive(Debug, Clone)]
pub struct Button {
    pub label: String,
    pub state: WidgetState,
    pub bounds: Rect,
}

impl Button {
    pub fn new(label: impl Into<String>, bounds: Rect) -> Self {
        Self {
            label: label.into(),
            state: WidgetState::Normal,
            bounds,
        }
    }

    pub fn on_click(&self, mouse: Vec2) -> Option<Command> {
        if self.state != WidgetState::Disabled && self.bounds.contains(mouse) {
            Some(Command {
                target: crate::arena::WidgetId { index: 0, generation: 0 },
                action: "clicked".to_string(),
            })
        } else {
            None
        }
    }
}

impl Widget for Button {
    fn bounds(&self) -> Rect {
        self.bounds
    }
    fn state(&self) -> WidgetState {
        self.state
    }
}

#[derive(Debug, Clone)]
pub struct Checkbox {
    pub checked: bool,
    pub state: WidgetState,
    pub bounds: Rect,
}

impl Checkbox {
    pub fn new(checked: bool, bounds: Rect) -> Self {
        Self {
            checked,
            state: WidgetState::Normal,
            bounds,
        }
    }
}

impl Widget for Checkbox {
    fn bounds(&self) -> Rect {
        self.bounds
    }
    fn state(&self) -> WidgetState {
        self.state
    }
}

#[derive(Debug, Clone)]
pub struct Toggle {
    pub on: bool,
    pub state: WidgetState,
    pub bounds: Rect,
}

impl Toggle {
    pub fn new(on: bool, bounds: Rect) -> Self {
        Self {
            on,
            state: WidgetState::Normal,
            bounds,
        }
    }
}

impl Widget for Toggle {
    fn bounds(&self) -> Rect {
        self.bounds
    }
    fn state(&self) -> WidgetState {
        self.state
    }
}

#[derive(Debug, Clone)]
pub struct Label {
    pub text: String,
    pub color: Color,
    pub bounds: Rect,
}

impl Label {
    pub fn new(text: impl Into<String>, color: Color, bounds: Rect) -> Self {
        Self {
            text: text.into(),
            color,
            bounds,
        }
    }
}

impl Widget for Label {
    fn bounds(&self) -> Rect {
        self.bounds
    }
}

#[derive(Debug, Clone)]
pub struct Panel {
    pub bounds: Rect,
    pub background: Color,
}

impl Panel {
    pub fn new(bounds: Rect, background: Color) -> Self {
        Self { bounds, background }
    }
}

impl Widget for Panel {
    fn bounds(&self) -> Rect {
        self.bounds
    }
}

#[derive(Debug, Clone)]
pub struct ImageIcon {
    pub bounds: Rect,
    pub tint: Color,
}

impl ImageIcon {
    pub fn new(bounds: Rect, tint: Color) -> Self {
        Self { bounds, tint }
    }
}

impl Widget for ImageIcon {
    fn bounds(&self) -> Rect {
        self.bounds
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_button_click_returns_command_when_enabled() {
        let btn = Button::new("OK", Rect::new(Vec2::new(0.0, 0.0), Vec2::new(10.0, 10.0)));
        assert!(btn.on_click(Vec2::new(5.0, 5.0)).is_some());
        assert!(btn.on_click(Vec2::new(20.0, 20.0)).is_none());
    }

    #[test]
    fn test_disabled_button_never_fires() {
        let mut btn = Button::new("OK", Rect::new(Vec2::new(0.0, 0.0), Vec2::new(10.0, 10.0)));
        btn.state = WidgetState::Disabled;
        assert!(btn.on_click(Vec2::new(5.0, 5.0)).is_none());
    }

    #[test]
    fn test_widget_kind_from_tag() {
        assert_eq!(WidgetKind::from_tag("button"), Some(WidgetKind::Button));
        assert_eq!(WidgetKind::from_tag("ham_bar"), Some(WidgetKind::HamBar));
        assert_eq!(WidgetKind::from_tag("unknown"), None);
    }
}
