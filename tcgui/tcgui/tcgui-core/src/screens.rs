//! GUI Screens implementation (Login, Character Select, HUD, Inventory, Chat).
//!
//! Manages active screen states and UI composition for the playable demo.

use crate::backend::{BackendEvent, ClientCommand, MockBackend, BackendConnection};
use crate::geometry::{Rect, Vec2};
use crate::window::WindowWidget;

#[derive(Debug, Clone, Copy, PartialEq, Eq)]
pub enum ScreenState {
    Login,
    CharacterSelect,
    Hud,
}

#[derive(Debug, Clone)]
pub struct GuiScreens {
    pub current_screen: ScreenState,
    pub backend: MockBackend,
    pub username_input: String,
    pub password_input: String,
    pub selected_character: Option<u64>,
    pub chat_input: String,
    pub chat_messages: Vec<String>,
    pub health: f32,
    pub max_health: f32,
    pub action: f32,
    pub max_action: f32,
    pub mind: f32,
    pub max_mind: f32,
    pub credits: u64,
    pub inventory_open: bool,
    pub chat_open: bool,
    pub login_window: WindowWidget,
    pub hud_window: WindowWidget,
}

impl GuiScreens {
    pub fn new() -> Self {
        let mut screens = Self {
            current_screen: ScreenState::Login,
            backend: MockBackend::new(),
            username_input: String::new(),
            password_input: String::new(),
            selected_character: None,
            chat_input: String::new(),
            chat_messages: Vec::new(),
            health: 1000.0,
            max_health: 1000.0,
            action: 1000.0,
            max_action: 1000.0,
            mind: 1000.0,
            max_mind: 1000.0,
            credits: 0,
            inventory_open: false,
            chat_open: true,
            login_window: WindowWidget::new("TCIndustries — Login", Rect::new(Vec2::new(312.0, 234.0), Vec2::new(712.0, 534.0))),
            hud_window: WindowWidget::new("HUD & Status", Rect::new(Vec2::new(20.0, 20.0), Vec2::new(320.0, 220.0))),
        };
        screens.poll_backend();
        screens
    }

    pub fn poll_backend(&mut self) {
        let events = self.backend.poll_events();
        for ev in events {
            match ev {
                BackendEvent::Connected => {
                    self.chat_messages.push("[System]: Connected to backend.".to_string());
                }
                BackendEvent::PlayerStateUpdate { health, action, mind, credits } => {
                    self.health = health;
                    self.action = action;
                    self.mind = mind;
                    self.credits = credits;
                }
                BackendEvent::ChatMessage { channel, sender, text } => {
                    self.chat_messages.push(format!("[{}][{}]: {}", channel, sender, text));
                }
                _ => {}
            }
        }
    }

    pub fn handle_login_submit(&mut self) {
        if !self.username_input.is_empty() {
            self.backend.send_command(ClientCommand::Login {
                username: self.username_input.clone(),
                token: "mock-token-12345".to_string(),
            });
            self.current_screen = ScreenState::Hud;
            self.chat_messages.push(format!("[System]: User '{}' logged in successfully.", self.username_input));
        }
    }

    pub fn send_chat(&mut self, text: String) {
        if !text.is_empty() {
            self.backend.send_command(ClientCommand::SendChat {
                channel: "General".to_string(),
                text: text.clone(),
            });
            self.chat_messages.push(format!("[General][You]: {}", text));
            self.chat_input.clear();
        }
    }
}

impl Default for GuiScreens {
    fn default() -> Self {
        Self::new()
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_screens_initial_state() {
        let screens = GuiScreens::new();
        assert_eq!(screens.current_screen, ScreenState::Login);
        assert!(!screens.chat_messages.is_empty());
    }

    #[test]
    fn test_login_transition() {
        let mut screens = GuiScreens::new();
        screens.username_input = "Tiwaking".to_string();
        screens.handle_login_submit();
        assert_eq!(screens.current_screen, ScreenState::Hud);
    }

    #[test]
    fn test_chat_sending() {
        let mut screens = GuiScreens::new();
        screens.send_chat("Hello Pre-CU world!".to_string());
        assert_eq!(screens.chat_messages.last().unwrap(), "[General][You]: Hello Pre-CU world!");
    }
}
