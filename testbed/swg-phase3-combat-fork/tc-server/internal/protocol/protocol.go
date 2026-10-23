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

	// Phase 3 (GDD Section 9)
	MsgCombatAction = "combat_action" // Client requests an attack on a target entity
	MsgSetPosture   = "set_posture"   // Client requests a posture change (GDD 9.2.3)
	MsgSetStance    = "set_stance"    // Client requests a stance change (GDD 9.2.3)
	MsgClone        = "clone"         // Client requests respawn at bound clone facility after death
	MsgRevive       = "revive"        // Client attempts to revive a nearby incapacitated player
)

// Server → Client messages
const (
	MsgWorldEnter         = "world_enter"         // Server confirms world entry with full character state
	MsgEntitySpawn        = "entity_spawn"        // Server tells client about another player or creature nearby
	MsgEntityMove         = "entity_move"         // Server updates entity position
	MsgEntityDespawn      = "entity_despawn"      // Server tells client an entity left
	MsgChatMessage        = "chat_message"        // Server relays a chat message
	MsgPositionCorrection = "position_correction" // Server corrects client position (anti-cheat)
	MsgError              = "error"               // Server sends an error

	// Phase 3 (GDD Section 9)
	MsgCombatResult   = "combat_result"   // Result of a resolved attack (hit/miss/damage)
	MsgHAMUpdate      = "ham_update"      // Pushed whenever a character's HAM/wounds/BF/posture/stance changes
	MsgCreatureDeath  = "creature_death"  // A creature instance died
	MsgIncapacitated  = "incapacitated"   // A character was incapacitated
	MsgCloned         = "cloned"          // A character respawned at their clone facility
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

// EntitySpawnMsg. EntityType/TemplateID are additive fields introduced in
// Phase 3 to distinguish creatures from players sharing the same spatial
// grid and spawn/despawn mechanism; both are empty/"player" for the
// existing Phase 0-2 player-spawn case, so the wire format stays backward
// compatible.
type EntitySpawnMsg struct {
	EntityID   string   `json:"entity_id"`
	Name       string   `json:"name"`
	Species    string   `json:"species"`
	Position   Position `json:"position"`
	EntityType string   `json:"entity_type,omitempty"` // "player" or "creature"
	TemplateID string   `json:"template_id,omitempty"` // creature template ID, empty for players
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

// --- Phase 3: Combat messages (GDD Section 9) ---

// CombatActionMsg: client requests an attack. EntityID is the target
// (another creature or, in a later phase once PvP flagging exists,
// another player).
type CombatActionMsg struct {
	TargetID string `json:"target_id"`
}

type SetPostureMsg struct {
	Posture string `json:"posture"` // "standing", "kneeling", "prone", "crouched"
}

type SetStanceMsg struct {
	Stance string `json:"stance"` // "normal", "aggressive", "defensive", "berserk"
}

type ReviveMsg struct {
	TargetID string `json:"target_id"` // incapacitated character to revive
}

// CombatResultMsg is sent to the attacker, the defender, and nearby
// observers whenever an attack resolves. GDD 9.2.2.
type CombatResultMsg struct {
	AttackerID string `json:"attacker_id"`
	TargetID   string `json:"target_id"`
	Hit        bool   `json:"hit"`
	Damage     int    `json:"damage"`
	HitChance  int    `json:"hit_chance"`
}

// HAMUpdateMsg carries a character's full combat-relevant state. Sent after
// any change (damage, regen, revive, clone, posture/stance change) rather
// than requiring the client to infer state from combat_result alone.
type HAMUpdateMsg struct {
	CharacterID       string  `json:"character_id"`
	HealthCurrent     int     `json:"health_current"`
	HealthMax         int     `json:"health_max"`
	ActionCurrent     int     `json:"action_current"`
	ActionMax         int     `json:"action_max"`
	MindCurrent       int     `json:"mind_current"`
	MindMax           int     `json:"mind_max"`
	WoundsHealth      int     `json:"wounds_health"`
	BattleFatiguePct  float64 `json:"battle_fatigue_pct"`
	Posture           string  `json:"posture"`
	Stance            string  `json:"stance"`
	Incapacitated     bool    `json:"incapacitated"`
}

type CreatureDeathMsg struct {
	EntityID string `json:"entity_id"`
}

type IncapacitatedMsg struct {
	CharacterID string `json:"character_id"`
}

type ClonedMsg struct {
	CharacterID string   `json:"character_id"`
	Position    Position `json:"position"`
	Planet      string   `json:"planet"`
}
