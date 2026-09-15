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
	Credits    int                    `json:"credits"` // Phase 5: wallet visibility
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

	// Phase 3: Combat messages (generic testbed set; no setting content)
	MsgCombatAction = "combat_action" // Client requests an attack on a target entity
	MsgSetPosture   = "set_posture"   // Client requests a posture change
	MsgSetStance    = "set_stance"    // Client requests a stance change
	MsgClone        = "clone"         // Client requests respawn at bound point after death
	MsgRevive       = "revive"        // Client attempts to revive a nearby incapacitated player

	// Phase 4: Resources & Crafting messages (generic testbed set)
	MsgSurvey          = "survey"           // Client surveys for a resource type with a tool
	MsgSample          = "sample"           // Client samples the nearest spawn in range
	MsgPlaceHarvester  = "place_harvester"  // Client places a harvester deed
	MsgEmptyHopper     = "empty_hopper"     // Client empties nearest owned harvester hopper
	MsgHarvestCorpse   = "harvest_corpse"   // Client harvests a nearby corpse

	// Phase 5: Economy Infrastructure messages (generic testbed set)
	MsgPlaceHouse    = "place_house"    // Client places a house deed
	MsgPlaceVendor   = "place_vendor"   // Client places a vendor deed
	MsgPurchase      = "purchase"       // Client buys a vendor listing in person
	MsgCollectTill   = "collect_till"   // Owner collects a vendor's till in person

	// Phase 7: Civic Systems messages (generic testbed set)
	MsgPlaceCityHall = "place_city_hall" // Client places a City Hall deed (founding)

	// Phase 6: Social Support Professions messages (generic testbed set)
	MsgHealWounds   = "heal_wounds"   // Medic heals a target's wounds
	MsgApplyBuff    = "apply_buff"    // Medic applies a HAM-pool buff
	MsgPerformStart = "perform_start" // Entertainer starts a performance
	MsgPerformStop  = "perform_stop"  // Entertainer stops performing
	MsgWatchStart   = "watch_start"   // Client starts watching a performer
	MsgWatchStop    = "watch_stop"    // Client stops watching
	MsgTip          = "tip"           // Client tips a watched performer
	MsgUseStim      = "use_stim"      // Client uses a stim pack on a target
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

	// Phase 3: Combat messages
	MsgCombatResult  = "combat_result"  // Result of a resolved attack (hit/miss/damage)
	MsgHAMUpdate     = "ham_update"     // Pushed whenever HAM/wounds/BF/posture/stance changes
	MsgCreatureDeath = "creature_death" // A creature instance died
	MsgIncapacitated = "incapacitated"  // A character was incapacitated
	MsgCloned        = "cloned"         // A character respawned at their bound point

	// Phase 4: Resources & Crafting messages
	MsgSurveyResult    = "survey_result"    // Nearest spawn of tool category + waypoint
	MsgSampleResult    = "sample_result"    // Units extracted into the character's stack
	MsgHarvesterPlaced = "harvester_placed" // Harvester deed placed
	MsgHopperEmptied   = "hopper_emptied"   // Hopper units moved into stack
	MsgCorpseHarvested = "corpse_harvested" // Corpse yield moved into stack

	// Phase 5: Economy Infrastructure messages
	MsgStructurePlaced = "structure_placed" // House/vendor deed placed
	MsgPurchaseReceipt = "purchase_receipt" // Atomic purchase completed
	MsgTillCollected   = "till_collected"   // Till moved to owner wallet

	// Phase 7 receipts and pushes.
	MsgCityFounded  = "city_founded"  // City Hall placed; city forming/active
	MsgGroupInvited = "group_invited" // An inviter invited this client to a group
	MsgGuildInvited = "guild_invited" // An inviter invited this client to a guild

	// Phase 6 receipts.
	MsgWoundHealed       = "wound_healed"
	MsgBuffApplied       = "buff_applied"
	MsgPerformanceStarted = "performance_started"
	MsgPerformanceStopped = "performance_stopped"
	MsgWatchStarted      = "watch_started"
	MsgWatchStopped      = "watch_stopped"
	MsgTipReceipt        = "tip_receipt"
	MsgTipReceived       = "tip_received"
	MsgStimUsed          = "stim_used"
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
	Channel string `json:"channel"` // spatial/shout/planet/group/guild/tell/faction/emote
	Text    string `json:"text"`
	Target  string `json:"target,omitempty"` // tell recipient character name
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
	// Phase 3: creature spawns carry type/template so clients can distinguish
	// creature entities from player entities sharing the spatial grid.
	EntityType string `json:"entity_type,omitempty"` // "player" or "creature"
	TemplateID string `json:"template_id,omitempty"` // creature template ID, empty for players
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

// --- Phase 3: Combat messages ---

// CombatActionMsg: client requests an attack. TargetID is a creature instance ID
// (player-vs-player targeting is later-phase scope — no PvP flagging exists yet).
type CombatActionMsg struct {
	TargetID string `json:"target_id"`
}

// PostureMsg / StanceMsg: client requests a posture or stance change.
type PostureMsg struct {
	Posture string `json:"posture"` // "standing", "kneeling", "prone", "crouched"
}

type StanceMsg struct {
	Stance string `json:"stance"` // "normal", "aggressive", "defensive", "berserk"
}

// ReviveMsg: client attempts to revive an incapacitated character in range.
type ReviveMsg struct {
	TargetID string `json:"target_id"`
}

// CombatResultMsg is broadcast to nearby clients on each resolved attack.
type CombatResultMsg struct {
	AttackerID string  `json:"attacker_id"`
	TargetID   string  `json:"target_id"`
	Hit        bool    `json:"hit"`
	Damage     int     `json:"damage"`
	HitChance  float64 `json:"hit_chance"`
}

// HAMUpdateMsg carries a character's full combat-relevant state. Sent after any
// change (damage, regen, revive, clone, posture/stance change).
type HAMUpdateMsg struct {
	CharacterID       string  `json:"character_id"`
	HealthCurrent     int     `json:"health_current"`
	HealthMax         int     `json:"health_max"`
	ActionCurrent     int     `json:"action_current"`
	MindCurrent       int     `json:"mind_current"`
	WoundsHealth      int     `json:"wounds_health"`
	BattleFatiguePct  float64 `json:"battle_fatigue_pct"`
	Posture           string  `json:"posture"`
	Stance            string  `json:"stance"`
	Incapacitated     bool    `json:"incapacitated"`
}

// CreatureDeathMsg notifies nearby clients that a creature instance died.
type CreatureDeathMsg struct {
	InstanceID string `json:"instance_id"`
	TemplateID string `json:"template_id"`
}

// IncapacitatedMsg notifies that a character was incapacitated.
type IncapacitatedMsg struct {
	CharacterID string `json:"character_id"`
}

// ClonedMsg confirms respawn at the bound point.
type ClonedMsg struct {
	CharacterID string   `json:"character_id"`
	Position    Position `json:"position"`
	Zone        string   `json:"zone"`
}

// --- Phase 4: Resources & Crafting messages ---

// SurveyMsg: client surveys with a category tool.
type SurveyMsg struct {
	Tool string `json:"tool"` // "mineral" | "chemical" | "flora" | "organic" | "water"
}

// WaypointMsg is an approximate spawn location (basic tier: exact center;
// triangulation gameplay deferred — flagged simplification).
type WaypointMsg struct {
	X float64 `json:"x"`
	Z float64 `json:"z"`
}

// SurveyResultMsg: nearest in-zone active spawn of the tool's types.
type SurveyResultMsg struct {
	SpawnID       string      `json:"spawn_id"`
	ResourceType  string      `json:"resource_type"`
	DistanceM     float64     `json:"distance_m"`
	Concentration int         `json:"concentration"`
	Waypoint      WaypointMsg `json:"waypoint"`
}

// SampleResultMsg: units extracted into the character's stack.
type SampleResultMsg struct {
	SpawnID      string `json:"spawn_id"`
	ResourceType string `json:"resource_type"`
	Units        int    `json:"units"`
}

// PlaceHarvesterMsg: client places a harvester (consumes a deed item).
type PlaceHarvesterMsg struct {
	X          float64 `json:"x"`
	Z          float64 `json:"z"`
	Kind       string  `json:"kind"` // "personal" | "medium" | "heavy"
	DeedItemID string  `json:"deed_item_id"`
}

// HarvesterPlacedMsg confirms placement.
type HarvesterPlacedMsg struct {
	HarvesterID string `json:"harvester_id"`
	SpawnID     string `json:"spawn_id"`
}

// HopperEmptiedMsg confirms hopper transfer.
type HopperEmptiedMsg struct {
	HarvesterID string `json:"harvester_id"`
	Units       int    `json:"units"`
}

// CorpseHarvestedMsg confirms corpse-yield transfer.
type CorpseHarvestedMsg struct {
	InstanceID   string `json:"instance_id"`
	ResourceType string `json:"resource_type"`
	Units        int    `json:"units"`
	Quality      int    `json:"quality"`
}

// --- Phase 5: Economy Infrastructure messages ---

// PlaceStructureMsg: client places a house or vendor deed at coordinates.
type PlaceStructureMsg struct {
	X          float64 `json:"x"`
	Z          float64 `json:"z"`
	Tier       string  `json:"tier,omitempty"` // houses: "small" | "medium" | "large"
	DeedItemID string  `json:"deed_item_id"`
}

// StructurePlacedMsg confirms placement.
type StructurePlacedMsg struct {
	StructureID string `json:"structure_id"`
	Kind        string `json:"kind"` // "house" | "vendor"
}

// PurchaseMsg: client buys a vendor listing (must be near the vendor).
type PurchaseMsg struct {
	ListingID string `json:"listing_id"`
}

// PurchaseReceiptMsg confirms the atomic transfer.
type PurchaseReceiptMsg struct {
	ListingID string `json:"listing_id"`
	ItemID    string `json:"item_id"`
}

// TillMsg: owner collects a vendor's till (must be near the vendor).
type TillMsg struct {
	VendorID string `json:"vendor_id"`
}

// TillCollectedMsg confirms till collection.
type TillCollectedMsg struct {
	VendorID string `json:"vendor_id"`
	Amount   int    `json:"amount"`
}

// --- Phase 7: Civic Systems messages ---

// PlaceCityHallMsg: client founds a city at coordinates (consumes a deed).
type PlaceCityHallMsg struct {
	X          float64 `json:"x"`
	Z          float64 `json:"z"`
	Name       string  `json:"name"`
	DeedItemID string  `json:"deed_item_id"`
}

// CityFoundedMsg confirms founding (forming or active).
type CityFoundedMsg struct {
	CityID    string `json:"city_id"`
	Status    string `json:"status"` // "forming" | "active"
	Structures int   `json:"structures"`
	Threshold  int   `json:"threshold"`
}

// GroupInvitedMsg notifies an online invitee of a group invite.
type GroupInvitedMsg struct {
	InviteID string `json:"invite_id"`
	GroupID  string `json:"group_id"`
	Inviter  string `json:"inviter"`
}

// GuildInvitedMsg notifies an online invitee of a guild invite.
type GuildInvitedMsg struct {
	InviteID string `json:"invite_id"`
	GuildID  string `json:"guild_id"`
	Inviter  string `json:"inviter"`
}

// --- Phase 6: Social Support Professions messages ---

// TargetMsg addresses another character (heal target).
type TargetMsg struct {
	TargetID string `json:"target_id"`
}

// WoundHealedMsg confirms wound healing.
type WoundHealedMsg struct {
	TargetID   string `json:"target_id"`
	Amount     int    `json:"amount"`
	WoundsLeft int    `json:"wounds_left"`
}

// BuffMsg requests a pool buff.
type BuffMsg struct {
	TargetID string `json:"target_id"`
	Pool     string `json:"pool"` // "health" | "action" | "mind"
}

// BuffAppliedMsg confirms a buff.
type BuffAppliedMsg struct {
	TargetID string `json:"target_id"`
	Pool     string `json:"pool"`
	Amount   int    `json:"amount"`
}

// PerformMsg starts a performance (kind) or carries session state.
type PerformMsg struct {
	Kind string `json:"kind,omitempty"` // "music" | "dance"
}

// WatchMsg starts/stops watching a performer.
type WatchMsg struct {
	PerformerID string `json:"performer_id,omitempty"`
}

// TipMsg tips a watched performer.
type TipMsg struct {
	TargetID string `json:"target_id"`
	Amount   int    `json:"amount"`
}

// StimMsg uses a stim pack on a target.
type StimMsg struct {
	ItemID   string `json:"item_id"`
	TargetID string `json:"target_id"`
}

// StimUsedMsg confirms stim use.
type StimUsedMsg struct {
	TargetID    string `json:"target_id"`
	ChargesLeft int    `json:"charges_left"`
}
