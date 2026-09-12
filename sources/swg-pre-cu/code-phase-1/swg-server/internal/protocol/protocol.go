package protocol

// Protocol defines the JSON message format for client-server communication.
// Real-time state sync uses WebSocket; non-real-time operations use HTTP.
// (GDD Section 29.2.4)

// --- HTTP API Request/Response types ---

type RegisterRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type RegisterResponse struct {
	AccountID string `json:"account_id"`
	Username  string `json:"username"`
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	AccountID string `json:"account_id"`
	Token     string `json:"token"`
	Username  string `json:"username"`
}

type CreateCharacterRequest struct {
	Species    string                 `json:"species"`
	Name       string                 `json:"name"`
	Appearance CharacterAppearanceReq `json:"appearance"`
}

type CharacterAppearanceReq struct {
	BodyType  string  `json:"body_type"`
	SkinColor string  `json:"skin_color"`
	HairStyle string  `json:"hair_style"`
	HairColor string  `json:"hair_color"`
	FaceType  string  `json:"face_type"`
	EyeColor  string  `json:"eye_color"`
	Height    float64 `json:"height"`
}

type CreateCharacterResponse struct {
	CharacterID string `json:"character_id"`
	Name        string `json:"name"`
	Species     string `json:"species"`
}

type ListCharactersResponse struct {
	Characters []CharacterSummary `json:"characters"`
}

type CharacterSummary struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Species string `json:"species"`
	Planet  string `json:"planet"`
}

type GetCharacterResponse struct {
	ID         string                 `json:"id"`
	Name       string                 `json:"name"`
	Species    string                 `json:"species"`
	Appearance CharacterAppearanceReq `json:"appearance"`
	Position   Position               `json:"position"`
	Planet     string                 `json:"planet"`
	HAM        HAMState               `json:"ham"`
}

type Position struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
	Z float64 `json:"z"`
}

type HAMState struct {
	Health       int `json:"health"`
	Strength     int `json:"strength"`
	Constitution int `json:"constitution"`
	Action       int `json:"action"`
	Quickness    int `json:"quickness"`
	Stamina      int `json:"stamina"`
	Mind         int `json:"mind"`
	Focus        int `json:"focus"`
	Willpower    int `json:"willpower"`
}

type SpeciesListResponse struct {
	Species []SpeciesInfo `json:"species"`
}

type SpeciesInfo struct {
	ID          string `json:"id"`
	DisplayName string `json:"display_name"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

// --- WebSocket message types ---
// (GDD Section 29.2.4: real-time state sync over persistent connection)

type WSMessage struct {
	Type string      `json:"type"`
	Data interface{} `json:"data,omitempty"`
}

// Client → Server messages
const (
	MsgEnterWorld = "enter_world" // Client requests to enter the world with a character
	MsgMove       = "move"        // Client sends position update (predicted movement)
	MsgChat       = "chat"        // Client sends a spatial chat message
)

// Server → Client messages
const (
	MsgWorldEnter         = "world_enter"         // Server confirms world entry with full character state
	MsgEntitySpawn        = "entity_spawn"        // Server tells client about another player nearby
	MsgEntityMove         = "entity_move"         // Server updates entity position
	MsgEntityDespawn      = "entity_despawn"      // Server tells client an entity left
	MsgChatMessage        = "chat_message"        // Server relays a chat message
	MsgPositionCorrection = "position_correction" // Server corrects client position (anti-cheat)
	MsgError              = "error"               // Server sends an error
)

type EnterWorldMsg struct {
	CharacterID string `json:"character_id"`
}

type WorldEnterMsg struct {
	CharacterID string   `json:"character_id"`
	Name        string   `json:"name"`
	Species     string   `json:"species"`
	Position    Position `json:"position"`
	Planet      string   `json:"planet"`
	HAM         HAMState `json:"ham"`
}

type MoveMsg struct {
	X       float64 `json:"x"`
	Y       float64 `json:"y"`
	Z       float64 `json:"z"`
	Heading float64 `json:"heading"`
}

type ChatMsg struct {
	Channel string `json:"channel"` // "spatial", "system", etc.
	Text    string `json:"text"`
}

type ChatMessageMsg struct {
	SenderName string `json:"sender_name"`
	Channel    string `json:"channel"`
	Text       string `json:"text"`
}

type EntitySpawnMsg struct {
	EntityID string   `json:"entity_id"`
	Name     string   `json:"name"`
	Species  string   `json:"species"`
	Position Position `json:"position"`
}

type EntityMoveMsg struct {
	EntityID string   `json:"entity_id"`
	Position Position `json:"position"`
	Heading  float64  `json:"heading"`
}

type EntityDespawnMsg struct {
	EntityID string `json:"entity_id"`
}

type ErrorMsg struct {
	Message string `json:"message"`
}

// PositionCorrectionMsg tells the client to snap to a server-authoritative position.
// Sent when the server rejects a movement (speed exceeded, out of bounds, etc.).
// GDD 29.2.1: server is authoritative. (GDD 29.5: anti-exploit)
type PositionCorrectionMsg struct {
	Position Position `json:"position"`
	Reason   string   `json:"reason"`
}
