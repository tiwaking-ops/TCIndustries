package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"

	"swg-server/internal/database"
	"swg-server/internal/models"
	"swg-server/internal/protocol"
	"swg-server/internal/species"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // Allow all origins in dev. Restrict in production.
	},
}

// Client represents a connected WebSocket client in the world.
type Client struct {
	Conn        *websocket.Conn
	CharacterID string
	Name        string
	Species     string
	AccountID   string
	Send        chan []byte
}

// WorldHandler manages all connected clients and real-time world state.
// This implements the authoritative server model (GDD Section 29.2.1).
type WorldHandler struct {
	db      *database.DB
	mu      sync.RWMutex
	clients map[string]*Client // characterID → Client
}

func NewWorldHandler(db *database.DB) *WorldHandler {
	return &WorldHandler{
		db:      db,
		clients: make(map[string]*Client),
	}
}

// HandleWebSocket upgrades HTTP to WebSocket and manages the client lifecycle.
// GDD Section 29.2.4: real-time state sync over persistent WebSocket connection.
func (h *WorldHandler) HandleWebSocket(w http.ResponseWriter, r *http.Request) {
	// Extract account ID from token (passed as query param for WS)
	token := r.URL.Query().Get("token")
	if token == "" {
		http.Error(w, "missing token", http.StatusUnauthorized)
		return
	}

	accountID, err := ParseToken(token)
	if err != nil {
		http.Error(w, "invalid token", http.StatusUnauthorized)
		return
	}

	_, err = h.db.GetAccountByID(accountID)
	if err != nil {
		http.Error(w, "invalid token", http.StatusUnauthorized)
		return
	}

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("WebSocket upgrade failed: %v", err)
		return
	}

	client := &Client{
		Conn:      conn,
		AccountID: accountID,
		Send:      make(chan []byte, 256),
	}

	// Start read and write pumps
	go h.writePump(client)
	go h.readPump(client)
}

// readPump reads messages from the client WebSocket.
func (h *WorldHandler) readPump(client *Client) {
	defer func() {
		h.disconnectClient(client)
		client.Conn.Close()
	}()

	for {
		_, message, err := client.Conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseNormalClosure) {
				log.Printf("WebSocket read error: %v", err)
			}
			break
		}

		var msg protocol.WSMessage
		if err := json.Unmarshal(message, &msg); err != nil {
			h.sendError(client, "invalid message format")
			continue
		}

		switch msg.Type {
		case protocol.MsgEnterWorld:
			h.handleEnterWorld(client, message)
		case protocol.MsgMove:
			h.handleMove(client, message)
		case protocol.MsgChat:
			h.handleChat(client, message)
		default:
			h.sendError(client, "unknown message type: "+msg.Type)
		}
	}
}

// writePump writes messages to the client WebSocket.
func (h *WorldHandler) writePump(client *Client) {
	defer client.Conn.Close()

	for message := range client.Send {
		if err := client.Conn.WriteMessage(websocket.TextMessage, message); err != nil {
			log.Printf("WebSocket write error: %v", err)
			break
		}
	}
}

// handleEnterWorld processes an enter_world request from the client.
// Loads the character from DB, registers the client in the world, and sends
// the initial world state including HAM values. (GDD Phase 0 exit criteria)
func (h *WorldHandler) handleEnterWorld(client *Client, raw []byte) {
	var msg protocol.WSMessage
	if err := json.Unmarshal(raw, &msg); err != nil {
		h.sendError(client, "invalid message")
		return
	}

	dataBytes, err := json.Marshal(msg.Data)
	if err != nil {
		h.sendError(client, "invalid message data")
		return
	}

	var enterMsg protocol.EnterWorldMsg
	if err := json.Unmarshal(dataBytes, &enterMsg); err != nil {
		h.sendError(client, "invalid enter_world data")
		return
	}

	// Load character from DB
	charWithHAM, err := h.db.GetCharacterByID(enterMsg.CharacterID)
	if err != nil {
		h.sendError(client, "character not found")
		return
	}

	// Verify this character belongs to the authenticated account
	if charWithHAM.AccountID != client.AccountID {
		h.sendError(client, "character does not belong to this account")
		return
	}

	// Register client in world
	client.CharacterID = charWithHAM.ID
	client.Name = charWithHAM.Name
	client.Species = charWithHAM.Species

	h.mu.Lock()
	h.clients[charWithHAM.ID] = client
	h.mu.Unlock()

	// Compute HAM from species (authoritative server-side computation)
	ham := species.ComputeHAM(species.SpeciesID(charWithHAM.Species))

	// Send world_enter confirmation with full character state
	worldEnter := protocol.WorldEnterMsg{
		CharacterID: charWithHAM.ID,
		Name:        charWithHAM.Name,
		Species:     charWithHAM.Species,
		Position: protocol.Position{
			X: charWithHAM.PosX,
			Y: charWithHAM.PosY,
			Z: charWithHAM.PosZ,
		},
		Planet: charWithHAM.Planet,
		HAM: protocol.HAMState{
			Health:       ham.Health,
			Strength:     ham.Strength,
			Constitution: ham.Constitution,
			Action:       ham.Action,
			Quickness:    ham.Quickness,
			Stamina:      ham.Stamina,
			Mind:         ham.Mind,
			Focus:        ham.Focus,
			Willpower:    ham.Willpower,
		},
	}

	h.send(client, protocol.MsgWorldEnter, worldEnter)

	// Notify other clients of this new entity
	h.broadcastEntitySpawn(client)

	// Tell this client about existing entities
	h.sendExistingEntities(client)

	log.Printf("Player %s (%s) entered world on %s at (%.1f, %.1f, %.1f)",
		charWithHAM.Name, charWithHAM.Species, charWithHAM.Planet,
		charWithHAM.PosX, charWithHAM.PosY, charWithHAM.PosZ)
}

// handleMove processes a movement update from the client.
// In Phase 0, the server accepts client position updates but does not
// validate them (server-side movement validation comes in Phase 1).
// GDD Section 29.2.1: server is authoritative, but Phase 0 allows predicted movement.
func (h *WorldHandler) handleMove(client *Client, raw []byte) {
	var msg protocol.WSMessage
	json.Unmarshal(raw, &msg)

	dataBytes, _ := json.Marshal(msg.Data)
	var moveMsg protocol.MoveMsg
	if err := json.Unmarshal(dataBytes, &moveMsg); err != nil {
		h.sendError(client, "invalid move data")
		return
	}

	if client.CharacterID == "" {
		h.sendError(client, "not in world")
		return
	}

	// Broadcast movement to other clients
	entityMove := protocol.EntityMoveMsg{
		EntityID: client.CharacterID,
		Position: protocol.Position{
			X: moveMsg.X,
			Y: moveMsg.Y,
			Z: moveMsg.Z,
		},
		Heading: moveMsg.Heading,
	}

	h.broadcastToOthers(client, protocol.MsgEntityMove, entityMove)
}

// handleChat processes a spatial chat message and relays it to nearby clients.
// GDD Section 18.2.1: spatial chat.
func (h *WorldHandler) handleChat(client *Client, raw []byte) {
	var msg protocol.WSMessage
	json.Unmarshal(raw, &msg)

	dataBytes, _ := json.Marshal(msg.Data)
	var chatMsg protocol.ChatMsg
	if err := json.Unmarshal(dataBytes, &chatMsg); err != nil {
		h.sendError(client, "invalid chat data")
		return
	}

	if client.CharacterID == "" {
		h.sendError(client, "not in world")
		return
	}

	chatMessage := protocol.ChatMessageMsg{
		SenderName: client.Name,
		Channel:    chatMsg.Channel,
		Text:       chatMsg.Text,
	}

	// Broadcast to all clients (Phase 0: no spatial filtering yet)
	h.broadcastToAll(protocol.MsgChatMessage, chatMessage)
}

// broadcastEntitySpawn tells all other clients about a newly spawned entity.
func (h *WorldHandler) broadcastEntitySpawn(client *Client) {
	spawn := protocol.EntitySpawnMsg{
		EntityID: client.CharacterID,
		Name:     client.Name,
		Species:  client.Species,
		Position: protocol.Position{
			X: models.DefaultSpawnX,
			Y: models.DefaultSpawnY,
			Z: models.DefaultSpawnZ,
		},
	}
	h.broadcastToOthers(client, protocol.MsgEntitySpawn, spawn)
}

// sendExistingEntities tells a client about all already-connected entities.
func (h *WorldHandler) sendExistingEntities(client *Client) {
	h.mu.RLock()
	defer h.mu.RUnlock()

	for _, other := range h.clients {
		if other.CharacterID == client.CharacterID {
			continue
		}
		spawn := protocol.EntitySpawnMsg{
			EntityID: other.CharacterID,
			Name:     other.Name,
			Species:  other.Species,
			Position: protocol.Position{
				X: models.DefaultSpawnX,
				Y: models.DefaultSpawnY,
				Z: models.DefaultSpawnZ,
			},
		}
		h.send(client, protocol.MsgEntitySpawn, spawn)
	}
}

// disconnectClient removes a client from the world and notifies others.
func (h *WorldHandler) disconnectClient(client *Client) {
	if client.CharacterID == "" {
		return
	}

	h.mu.Lock()
	delete(h.clients, client.CharacterID)
	h.mu.Unlock()

	// Notify other clients
	despawn := protocol.EntityDespawnMsg{EntityID: client.CharacterID}
	h.broadcastToOthers(client, protocol.MsgEntityDespawn, despawn)

	close(client.Send)
	log.Printf("Player %s disconnected", client.Name)
}

// send sends a typed message to a single client.
func (h *WorldHandler) send(client *Client, msgType string, data interface{}) {
	msg := protocol.WSMessage{Type: msgType, Data: data}
	payload, err := json.Marshal(msg)
	if err != nil {
		log.Printf("Failed to marshal message: %v", err)
		return
	}
	select {
	case client.Send <- payload:
	default:
		log.Printf("Client %s send buffer full, dropping message", client.Name)
	}
}

// sendError sends an error message to a client.
func (h *WorldHandler) sendError(client *Client, message string) {
	h.send(client, protocol.MsgError, protocol.ErrorMsg{Message: message})
}

// broadcastToAll sends a message to all connected clients.
func (h *WorldHandler) broadcastToAll(msgType string, data interface{}) {
	msg := protocol.WSMessage{Type: msgType, Data: data}
	payload, _ := json.Marshal(msg)

	h.mu.RLock()
	defer h.mu.RUnlock()

	for _, client := range h.clients {
		select {
		case client.Send <- payload:
		default:
			log.Printf("Client %s send buffer full", client.Name)
		}
	}
}

// broadcastToOthers sends a message to all clients except the sender.
func (h *WorldHandler) broadcastToOthers(sender *Client, msgType string, data interface{}) {
	msg := protocol.WSMessage{Type: msgType, Data: data}
	payload, _ := json.Marshal(msg)

	h.mu.RLock()
	defer h.mu.RUnlock()

	for _, client := range h.clients {
		if client.CharacterID == sender.CharacterID {
			continue
		}
		select {
		case client.Send <- payload:
		default:
			log.Printf("Client %s send buffer full", client.Name)
		}
	}
}
