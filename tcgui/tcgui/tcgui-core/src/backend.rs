//! Backend interface trait and deterministic mock implementation.
//!
//! Defines the communication bridge between the GUI client and the game server
//! or simulation backend.

use serde::{Serialize, Deserialize};

/// Events sent from the backend/server to the GUI client.
#[derive(Debug, Clone, Serialize, Deserialize, PartialEq)]
pub enum BackendEvent {
    Connected,
    Disconnected { reason: String },
    PlayerStateUpdate {
        health: f32,
        action: f32,
        mind: f32,
        credits: u64,
    },
    ChatMessage {
        channel: String,
        sender: String,
        text: String,
    },
    InventoryUpdate {
        items: Vec<InventoryItemInfo>,
    },
}

#[derive(Debug, Clone, Serialize, Deserialize, PartialEq)]
pub struct InventoryItemInfo {
    pub slot_id: u32,
    pub item_name: String,
    pub quantity: u32,
}

/// Commands/actions sent from the GUI client to the backend/server.
#[derive(Debug, Clone, Serialize, Deserialize, PartialEq)]
pub enum ClientCommand {
    Login { username: String, token: String },
    SelectCharacter { character_id: u64 },
    SendChat { channel: String, text: String },
    UseItem { slot_id: u32 },
}

/// Trait defining the backend communication contract.
pub trait BackendConnection: std::fmt::Debug {
    /// Send a command to the backend.
    fn send_command(&mut self, cmd: ClientCommand);
    /// Poll for pending events from the backend.
    fn poll_events(&mut self) -> Vec<BackendEvent>;
}

/// A deterministic mock backend for standalone GUI testing and demo mode.
#[derive(Debug, Clone)]
pub struct MockBackend {
    pub connected: bool,
    pub outbound_commands: Vec<ClientCommand>,
    pub inbound_events: Vec<BackendEvent>,
}

impl MockBackend {
    pub fn new() -> Self {
        let mut backend = Self {
            connected: true,
            outbound_commands: Vec::new(),
            inbound_events: Vec::new(),
        };
        backend.inbound_events.push(BackendEvent::Connected);
        backend.inbound_events.push(BackendEvent::PlayerStateUpdate {
            health: 850.0,
            action: 900.0,
            mind: 1000.0,
            credits: 15420,
        });
        backend.inbound_events.push(BackendEvent::ChatMessage {
            channel: "System".to_string(),
            sender: "Server".to_string(),
            text: "Welcome to TCIndustries Pre-CU SWG Sandbox.".to_string(),
        });
        backend
    }
}

impl BackendConnection for MockBackend {
    fn send_command(&mut self, cmd: ClientCommand) {
        self.outbound_commands.push(cmd);
    }

    fn poll_events(&mut self) -> Vec<BackendEvent> {
        std::mem::take(&mut self.inbound_events)
    }
}

impl Default for MockBackend {
    fn default() -> Self {
        Self::new()
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_mock_backend_flow() {
        let mut backend = MockBackend::new();
        let events = backend.poll_events();
        assert!(!events.is_empty());

        backend.send_command(ClientCommand::SendChat {
            channel: "General".to_string(),
            text: "Hello World!".to_string(),
        });
        assert_eq!(backend.outbound_commands.len(), 1);
    }
}
