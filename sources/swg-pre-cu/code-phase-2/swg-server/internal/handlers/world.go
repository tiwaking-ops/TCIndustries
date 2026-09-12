package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/gorilla/websocket"

	"swg-server/internal/database"
	"swg-server/internal/models"
	"swg-server/internal/protocol"
	"swg-server/internal/species"
	"swg-server/internal/world"
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

	// Phase 1: Server-authoritative position tracking
	Pos          world.Position
	Heading      float64
	LastMoveTime time.Time

	// Phase 1: Visibility tracking — which entities this client can currently see
	Visible map[string]bool // set of character IDs visible to this client

	// Phase 1: Position persistence
	LastPersistTime time.Time
}

// WorldHandler manages all connected clients, the spatial grid, and real-time
// world state. Implements the authoritative server model (GDD 29.2.1) with
// spatial partitioning for interest management (GDD 29.2.2).
type WorldHandler struct {
	db   *database.DB
	mu   sync.RWMutex
	grid *world.Grid

	// clients maps characterID → *Client for all connected players
	clients map[string]*Client
}

func NewWorldHandler(db *database.DB) *WorldHandler {
	return &WorldHandler{
		db:      db,
		grid:    world.NewGrid(64.0), // 64m cells
		clients: make(map[string]*Client),
	}
}

// HandleWebSocket upgrades HTTP to WebSocket and manages the client lifecycle.
// GDD Section 29.2.4: real-time state sync over persistent WebSocket connection.
func (h *WorldHandler) HandleWebSocket(w http.ResponseWriter, r *http.Request) {
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
		Visible:   make(map[string]bool),
	}

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
// Loads the character from DB, registers the client in the spatial grid,
// and sends the initial world state. (GDD Phase 0/1 exit criteria)
func (h *WorldHandler) handleEnterWorld(client *Client, raw []byte) {
	var msg protocol.WSMessage
	if err := json.Unmarshal(raw, &msg); err != nil {
		h.sendError(client, "invalid message")
		return
	}

	dataBytes, _ := json.Marshal(msg.Data)
	var enterMsg protocol.EnterWorldMsg
	if err := json.Unmarshal(dataBytes, &enterMsg); err != nil {
		h.sendError(client, "invalid enter_world data")
		return
	}

	charWithHAM, err := h.db.GetCharacterByID(enterMsg.CharacterID)
	if err != nil {
		h.sendError(client, "character not found")
		return
	}

	if charWithHAM.AccountID != client.AccountID {
		h.sendError(client, "character does not belong to this account")
		return
	}

	// Set client state
	client.CharacterID = charWithHAM.ID
	client.Name = charWithHAM.Name
	client.Species = charWithHAM.Species
	client.Pos = world.Position{
		Planet: charWithHAM.Planet,
		X:      charWithHAM.PosX,
		Y:      charWithHAM.PosY,
		Z:      charWithHAM.PosZ,
	}
	client.Heading = charWithHAM.Heading
	client.LastMoveTime = time.Now()
	client.LastPersistTime = time.Now()

	// Register in the clients map
	h.mu.Lock()
	h.clients[charWithHAM.ID] = client
	h.mu.Unlock()

	// Add to spatial grid
	h.grid.AddEntity(&world.Entity{
		CharacterID: client.CharacterID,
		Name:        client.Name,
		Species:     client.Species,
		Pos:         client.Pos,
		Heading:     client.Heading,
	})

	// Compute HAM from species (authoritative server-side computation)
	ham := species.ComputeHAM(species.SpeciesID(charWithHAM.Species))

	// Send world_enter confirmation with full character state
	worldEnter := protocol.WorldEnterMsg{
		CharacterID: charWithHAM.ID,
		Name:        charWithHAM.Name,
		Species:     charWithHAM.Species,
		Position: protocol.Position{
			X: client.Pos.X,
			Y: client.Pos.Y,
			Z: client.Pos.Z,
		},
		Planet: client.Pos.Planet,
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

	// Phase 1: Compute initial visibility and send spawns
	h.updateVisibility(client)

	log.Printf("Player %s (%s) entered world on %s at (%.1f, %.1f, %.1f)",
		client.Name, client.Species, client.Pos.Planet,
		client.Pos.X, client.Pos.Y, client.Pos.Z)
}

// handleMove processes a movement update from the client.
// Phase 1: Server-validated movement with anti-cheat (GDD 29.2.1, 29.5).
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

	newPos := world.Position{
		Planet: client.Pos.Planet,
		X:      moveMsg.X,
		Y:      moveMsg.Y,
		Z:      moveMsg.Z,
	}
	now := time.Now()
	dt := now.Sub(client.LastMoveTime)

	// Validate movement
	result := world.ValidateMovement(client.Pos, newPos, dt)
	if !result.Accepted {
		// Reject: send position correction
		correction := protocol.PositionCorrectionMsg{
			Position: protocol.Position{
				X: client.Pos.X,
				Y: client.Pos.Y,
				Z: client.Pos.Z,
			},
			Reason: result.Reason,
		}
		h.send(client, protocol.MsgPositionCorrection, correction)
		log.Printf("Movement rejected for %s: %s (dist=%.1f, max=%.1f)",
			client.Name, result.Reason, result.ActualDist, result.MaxAllowedDist)
		return
	}
	if result.Coalesced {
		// Move accepted but not applied (dt too small). Do not update position,
		// grid, visibility, or persistence. (Anti-bypass: a client can't send
		// a huge move within 50ms to skip validation.)
		return
	}

	// Accepted: update server-side state
	oldPos := client.Pos
	client.Pos = newPos
	client.Heading = moveMsg.Heading
	client.LastMoveTime = now

	// Update entity position in the spatial grid
	h.grid.MoveEntity(client.CharacterID, oldPos, newPos, client.Heading)

	// Broadcast movement to interested clients (within interest radius)
	entityMove := protocol.EntityMoveMsg{
		EntityID: client.CharacterID,
		Position: protocol.Position{
			X: newPos.X,
			Y: newPos.Y,
			Z: newPos.Z,
		},
		Heading: moveMsg.Heading,
	}
	h.broadcastToInterested(client, protocol.MsgEntityMove, entityMove)

	// Update visibility (entity may have entered/left someone's range)
	h.updateVisibility(client)

	// Also update other clients' visibility of the mover
	h.updateVisibilityForOthers(client)

	// Persist position periodically (every 5 seconds while moving)
	if now.Sub(client.LastPersistTime) >= 5*time.Second {
		h.persistPosition(client)
	}
}

// handleChat processes a chat message and delivers it based on channel type.
// Phase 1: Spatial chat with radius filtering (GDD 18.2.1).
// - spatial (/say): 20m radius
// - shout: 50m radius
// - planet: all clients on same planet
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

	channel := world.ChatChannel(chatMsg.Channel)
	radius := world.ChatRadius(channel)

	if radius > 0 {
		// Spatial chat: deliver to nearby clients within radius
		h.mu.RLock()
		nearby := h.grid.Nearby(client.Pos, radius, client.CharacterID)
		h.mu.RUnlock()

		// Also include the sender (echo back)
		h.send(client, protocol.MsgChatMessage, chatMessage)

		// Deliver to nearby clients
		h.mu.RLock()
		for _, e := range nearby {
			if other, ok := h.clients[e.CharacterID]; ok {
				h.send(other, protocol.MsgChatMessage, chatMessage)
			}
		}
		h.mu.RUnlock()
	} else {
		// Planet-wide or system chat
		h.mu.RLock()
		for _, c := range h.clients {
			if c.Pos.Planet == client.Pos.Planet {
				h.send(c, protocol.MsgChatMessage, chatMessage)
			}
		}
		h.mu.RUnlock()
	}
}

// updateVisibility recomputes which entities this client can see, and sends
// spawn/despawn messages for entities that entered/left the interest radius.
// Phase 1: Proximity-based entity sync (GDD 29.2.2).
func (h *WorldHandler) updateVisibility(client *Client) {
	// Query nearby entities within interest radius
	nearby := h.grid.Nearby(client.Pos, world.InterestRadius, client.CharacterID)

	newVisible := make(map[string]bool)
	for _, e := range nearby {
		newVisible[e.CharacterID] = true
	}

	// Diff: send despawn for entities that left visibility
	for id := range client.Visible {
		if !newVisible[id] {
			h.send(client, protocol.MsgEntityDespawn, protocol.EntityDespawnMsg{EntityID: id})
		}
	}

	// Diff: send spawn for entities that newly entered visibility
	for id := range newVisible {
		if !client.Visible[id] {
			// Find the entity to get its info
			if e := h.findEntity(id); e != nil {
				h.send(client, protocol.MsgEntitySpawn, protocol.EntitySpawnMsg{
					EntityID: e.CharacterID,
					Name:     e.Name,
					Species:  e.Species,
					Position: protocol.Position{
						X: e.Pos.X,
						Y: e.Pos.Y,
						Z: e.Pos.Z,
					},
				})
			}
		}
	}

	client.Visible = newVisible

	// Symmetric visibility: if we just started seeing someone, make sure
	// they also see us (if they didn't already)
	h.mu.Lock()
	for _, e := range nearby {
		if other, ok := h.clients[e.CharacterID]; ok {
			if !other.Visible[client.CharacterID] {
				other.Visible[client.CharacterID] = true
				h.send(other, protocol.MsgEntitySpawn, protocol.EntitySpawnMsg{
					EntityID: client.CharacterID,
					Name:     client.Name,
					Species:  client.Species,
					Position: protocol.Position{
						X: client.Pos.X,
						Y: client.Pos.Y,
						Z: client.Pos.Z,
					},
				})
			}
		}
	}
	h.mu.Unlock()
}

// updateVisibilityForOthers checks all clients who can potentially see the mover.
// If the mover has left their interest radius, they get a despawn.
// If the mover has entered their interest radius, they get a spawn.
// This is the symmetric counterpart to updateVisibility — it handles the case
// where the MOVER changes position and affects OTHER clients' visibility.
func (h *WorldHandler) updateVisibilityForOthers(mover *Client) {
	h.mu.Lock()
	defer h.mu.Unlock()

	// Find all clients within 2x interest radius (max range that could be affected)
	// This is an optimization — we only need to check clients near the mover.
	nearby := h.grid.Nearby(mover.Pos, world.InterestRadius*1.5, mover.CharacterID)
	nearbySet := make(map[string]bool)
	for _, e := range nearby {
		nearbySet[e.CharacterID] = true
	}

	for _, other := range h.clients {
		if other.CharacterID == mover.CharacterID || other.CharacterID == "" {
			continue
		}

		// Check if other could previously see mover
		wasVisible := other.Visible[mover.CharacterID]

		// Check if other can now see mover (within interest radius)
		nowVisible := nearbySet[other.CharacterID] &&
			world.Distance2D(other.Pos, mover.Pos) <= world.InterestRadius

		if wasVisible && !nowVisible {
			// Mover left other's range
			delete(other.Visible, mover.CharacterID)
			h.send(other, protocol.MsgEntityDespawn, protocol.EntityDespawnMsg{EntityID: mover.CharacterID})
		} else if !wasVisible && nowVisible {
			// Mover entered other's range
			other.Visible[mover.CharacterID] = true
			h.send(other, protocol.MsgEntitySpawn, protocol.EntitySpawnMsg{
				EntityID: mover.CharacterID,
				Name:     mover.Name,
				Species:  mover.Species,
				Position: protocol.Position{
					X: mover.Pos.X,
					Y: mover.Pos.Y,
					Z: mover.Pos.Z,
				},
			})
		}
	}
}

// findEntity looks up an entity in the grid by character ID.
func (h *WorldHandler) findEntity(characterID string) *world.Entity {
	// Search all cells — this is only called on visibility change, not every frame
	h.mu.RLock()
	defer h.mu.RUnlock()

	if client, ok := h.clients[characterID]; ok {
		return &world.Entity{
			CharacterID: client.CharacterID,
			Name:        client.Name,
			Species:     client.Species,
			Pos:         client.Pos,
			Heading:     client.Heading,
		}
	}
	return nil
}

// broadcastToInterested sends a message to all clients within the interest radius
// of the sender. Used for movement updates. (GDD 29.2.2)
func (h *WorldHandler) broadcastToInterested(sender *Client, msgType string, data interface{}) {
	msg := protocol.WSMessage{Type: msgType, Data: data}
	payload, _ := json.Marshal(msg)

	h.mu.RLock()
	defer h.mu.RUnlock()

	nearby := h.grid.Nearby(sender.Pos, world.InterestRadius, sender.CharacterID)
	for _, e := range nearby {
		if client, ok := h.clients[e.CharacterID]; ok {
			select {
			case client.Send <- payload:
			default:
				log.Printf("Client %s send buffer full", client.Name)
			}
		}
	}
}

// disconnectClient removes a client from the world and notifies nearby players.
func (h *WorldHandler) disconnectClient(client *Client) {
	if client.CharacterID == "" {
		return
	}

	// Persist final position
	h.persistPosition(client)

	// Remove from spatial grid
	h.grid.RemoveEntity(client.CharacterID, client.Pos)

	// Remove from clients map
	h.mu.Lock()
	delete(h.clients, client.CharacterID)
	h.mu.Unlock()

	// Notify clients who could see this entity that it despawned
	despawn := protocol.EntityDespawnMsg{EntityID: client.CharacterID}
	h.broadcastToInterested(client, protocol.MsgEntityDespawn, despawn)

	// Also remove from other clients' visibility sets
	h.mu.Lock()
	for _, other := range h.clients {
		if other.Visible[client.CharacterID] {
			delete(other.Visible, client.CharacterID)
			h.send(other, protocol.MsgEntityDespawn, despawn)
		}
	}
	h.mu.Unlock()

	close(client.Send)
	log.Printf("Player %s disconnected", client.Name)
}

// persistPosition saves the client's position to the database.
// Called every 5 seconds while moving and on disconnect. (GDD 29.4)
func (h *WorldHandler) persistPosition(client *Client) {
	if client.CharacterID == "" {
		return
	}
	if err := h.db.UpdateCharacterPosition(
		client.CharacterID,
		client.Pos.X, client.Pos.Y, client.Pos.Z,
		client.Heading, client.Pos.Planet,
	); err != nil {
		log.Printf("Failed to persist position for %s: %v", client.Name, err)
	} else {
		client.LastPersistTime = time.Now()
	}
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

// Unused but kept for reference — broadcasts to all clients on the same planet.
func (h *WorldHandler) broadcastToPlanet(sender *Client, msgType string, data interface{}) {
	msg := protocol.WSMessage{Type: msgType, Data: data}
	payload, _ := json.Marshal(msg)

	h.mu.RLock()
	defer h.mu.RUnlock()

	for _, client := range h.clients {
		if client.Pos.Planet != sender.Pos.Planet {
			continue
		}
		select {
		case client.Send <- payload:
		default:
			log.Printf("Client %s send buffer full", client.Name)
		}
	}
}

// models import used for DefaultSpawn constants in enter world flow
var _ = models.DefaultSpawnX
