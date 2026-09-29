// Elite HTTP handlers for Phase 9: pet commands/combining, Image Designer +
// Smuggler handshakes, holoemotes, spice/buff-pack use. All routes authed;
// acting character travels in body/query (civic-route conventions). Testbed
// fork; generic content only. GDD-silent numbers [PROVISIONAL], flagged.
package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"swg-server/internal/database"
)

// EliteHandler serves /api/elite/*. It holds the world handler for grid
// operations (pet release) and presence checks.
type EliteHandler struct {
	db *database.DB
	w  *WorldHandler
}

// NewEliteHandler creates the handler.
func NewEliteHandler(db *database.DB, w *WorldHandler) *EliteHandler {
	return &EliteHandler{db: db, w: w}
}

// HandleEliteRoute dispatches /api/elite/* sub-paths.
func (eh *EliteHandler) HandleEliteRoute(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimPrefix(r.URL.Path, "/api/elite/")
	charID := r.URL.Query().Get("character_id")
	switch {
	case path == "pet/list" && r.Method == "GET":
		eh.petList(w, r, charID)
	case path == "pet/command" && r.Method == "POST":
		eh.petCommand(w, r)
	case path == "pet/release" && r.Method == "POST":
		eh.petRelease(w, r)
	case path == "camp/list" && r.Method == "GET":
		eh.campList(w, r, charID)
	case path == "dna/combine" && r.Method == "POST":
		eh.dnaCombine(w, r)
	case path == "id/request" && r.Method == "POST":
		eh.idRequest(w, r)
	case path == "id/respond" && r.Method == "POST":
		eh.idRespond(w, r)
	case path == "holo/create" && r.Method == "POST":
		eh.holoCreate(w, r)
	case path == "holo/buy" && r.Method == "POST":
		eh.holoBuy(w, r)
	case path == "holo/list" && r.Method == "GET":
		eh.holoList(w, r, charID)
	case path == "slice/request" && r.Method == "POST":
		eh.sliceRequest(w, r)
	case path == "slice/respond" && r.Method == "POST":
		eh.sliceRespond(w, r)
	case path == "spice/use" && r.Method == "POST":
		eh.spiceUse(w, r)
	case path == "pack/use" && r.Method == "POST":
		eh.packUse(w, r)
	default:
		writeCraftJSON(w, http.StatusNotFound, map[string]string{"error": "unknown elite route"})
	}
}

// --- Pets ---

func (eh *EliteHandler) petList(w http.ResponseWriter, r *http.Request, charID string) {
	if charID == "" {
		writeCraftJSON(w, http.StatusBadRequest, map[string]string{"error": "character_id required"})
		return
	}
	list, err := eh.db.PetsOf(charID)
	if err != nil {
		writeCraftJSON(w, http.StatusInternalServerError, map[string]string{"error": "lookup failed"})
		return
	}
	if list == nil {
		list = []database.PetRow{}
	}
	writeCraftJSON(w, http.StatusOK, map[string]interface{}{"pets": list})
}

// petCommand sets stay/follow/attack(+target creature). Commands need no
// Creature Handler boxes beyond ownership (uniform ownership model — flagged;
// CH gating lives at tame time).
func (eh *EliteHandler) petCommand(w http.ResponseWriter, r *http.Request) {
	body, ok := decodeBody(w, r)
	if !ok {
		return
	}
	charID, petID, mode := strOf(body, "character_id"), strOf(body, "pet_id"), strOf(body, "mode")
	p, err := eh.db.GetPet(petID)
	if err != nil || p.OwnerID != charID {
		writeCraftJSON(w, http.StatusNotFound, map[string]string{"error": "pet unknown"})
		return
	}
	targetID := strOf(body, "target_id")
	switch mode {
	case "stay", "follow":
		targetID = ""
	case "attack":
		if targetID == "" {
			writeCraftJSON(w, http.StatusBadRequest, map[string]string{"error": "attack target required"})
			return
		}
		// Creatures only (PvP pets need targeting-validity design — OUT).
		live := false
		if rows, err := eh.db.GetLivingCreatureInstances(); err == nil {
			for i := range rows {
				if rows[i].ID == targetID {
					live = true
				}
			}
		}
		if !live {
			writeCraftJSON(w, http.StatusBadRequest, map[string]string{"error": "target not a live creature"})
			return
		}
	default:
		writeCraftJSON(w, http.StatusBadRequest, map[string]string{"error": "mode must be stay/follow/attack"})
		return
	}
	if err := eh.db.SetPetCommand(petID, mode, targetID); err != nil {
		writeCraftJSON(w, http.StatusInternalServerError, map[string]string{"error": "command failed"})
		return
	}
	writeCraftJSON(w, http.StatusOK, map[string]string{"status": "command set"})
}

func (eh *EliteHandler) campList(w http.ResponseWriter, r *http.Request, _ string) {
	list, err := eh.db.ActiveCamps("zone-0001", time.Now().Unix())
	if err != nil {
		writeCraftJSON(w, http.StatusInternalServerError, map[string]string{"error": "lookup failed"})
		return
	}
	if list == nil {
		list = []database.CampRow{}
	}
	writeCraftJSON(w, http.StatusOK, map[string]interface{}{"camps": list})
}

func (eh *EliteHandler) petRelease(w http.ResponseWriter, r *http.Request) {
	body, ok := decodeBody(w, r)
	if !ok {
		return
	}
	charID, petID := strOf(body, "character_id"), strOf(body, "pet_id")
	p, err := eh.db.GetPet(petID)
	if err != nil || p.OwnerID != charID {
		writeCraftJSON(w, http.StatusNotFound, map[string]string{"error": "pet unknown"})
		return
	}
	_ = eh.db.DeletePet(petID)
	eh.w.removePetWorld(petID)
	writeCraftJSON(w, http.StatusOK, map[string]string{"status": "released"})
}

// dnaCombine engineers an enhanced pet from two DNA samples (amended §9.3
// scope). Outcome quality = max(parent Q) + 5×Engineering-tier, capped at
// 1000 (PROVISIONAL formula — confirmation rides item 7). Template = first
// parent (flagged). Both samples consumed.
func (eh *EliteHandler) dnaCombine(w http.ResponseWriter, r *http.Request) {
	body, ok := decodeBody(w, r)
	if !ok {
		return
	}
	charID := strOf(body, "character_id")
	if has, _ := eh.db.HasSkillBox(charID, "bioengineer_novice"); !has {
		writeCraftJSON(w, http.StatusForbidden, map[string]string{"error": "requires Novice Bio-Engineer"})
		return
	}
	a, err := eh.db.GetItem(charID, strOf(body, "dna_a"))
	if err != nil || a.Schematic != "dna_sample" {
		writeCraftJSON(w, http.StatusBadRequest, map[string]string{"error": "valid DNA A required"})
		return
	}
	b, err := eh.db.GetItem(charID, strOf(body, "dna_b"))
	if err != nil || b.Schematic != "dna_sample" {
		writeCraftJSON(w, http.StatusBadRequest, map[string]string{"error": "valid DNA B required"})
		return
	}
	tmplA := dnaTemplateOf(a)
	qa, qb := dnaQualityOf(a), dnaQualityOf(b)
	best := qa
	if qb > best {
		best = qb
	}
	tier := eh.w.skillTier(charID, "bioengineer_creature_engineering_")
	quality := best + 5*tier
	if quality > 1000 {
		quality = 1000
	}
	name := strOf(body, "name")
	if name == "" || len(name) > 40 {
		name = "Enhanced Specimen"
	}
	petID, err := eh.db.CreatePet(charID, tmplA, name, quality, "engineered",
		"zone-0001", 20.0, 0.0)
	if err != nil {
		writeCraftJSON(w, http.StatusInternalServerError, map[string]string{"error": "engineering failed"})
		return
	}
	_ = eh.db.DeleteItem(charID, a.ID)
	_ = eh.db.DeleteItem(charID, b.ID)
	eh.w.addPetWorld(petID, charID, tmplA, name, quality, "zone-0001", 20.0, 0.0)
	writeCraftJSON(w, http.StatusCreated, map[string]string{"pet_id": petID})
}

func dnaQualityOf(item *database.CraftedItemRow) int {
	if item.Stats == nil {
		return 0
	}
	return int(item.Stats["dna_quality"])
}

func dnaTemplateOf(item *database.CraftedItemRow) string {
	if item.Stats == nil {
		return "0001"
	}
	n := int(item.Stats["dna_template"])
	if n <= 0 {
		return "0001"
	}
	return fmt.Sprintf("%04d", n)
}

// --- Image Designer handshake (invite pattern) ---

var idFieldGates = map[string]string{
	"hair_style": "imagedesigner_hair_styling_",
	"hair_color": "imagedesigner_hair_styling_",
	"face_type":  "imagedesigner_facial_modification_",
	"eye_color":  "imagedesigner_facial_modification_",
	"body_type":  "imagedesigner_body_modification_",
	"skin_color": "imagedesigner_body_modification_",
	"height":     "imagedesigner_body_modification_",
}

func (eh *EliteHandler) idRequest(w http.ResponseWriter, r *http.Request) {
	body, ok := decodeBody(w, r)
	if !ok {
		return
	}
	charID := strOf(body, "character_id")
	if has, _ := eh.db.HasSkillBox(charID, "imagedesigner_novice"); !has {
		writeCraftJSON(w, http.StatusForbidden, map[string]string{"error": "requires Novice Image Designer"})
		return
	}
	targetID, ok := eh.charByName(w, strOf(body, "target_name"))
	if !ok {
		return
	}
	fee := intOf(body, "fee")
	if fee < 0 {
		writeCraftJSON(w, http.StatusBadRequest, map[string]string{"error": "fee must not be negative"})
		return
	}
	id, err := eh.db.CreateServiceRequest("imagedesign", charID, targetID, "",
		strOf(body, "changes"), fee)
	if err != nil {
		writeCraftJSON(w, http.StatusInternalServerError, map[string]string{"error": "request failed"})
		return
	}
	writeCraftJSON(w, http.StatusCreated, map[string]string{"request_id": id})
}

// idRespond applies field edits atomically with the fee transfer. Fields are
// JSON {"hair_style": "...", ...}; each field needs its tree box (provisional
// per-field mapping — §6.11).
func (eh *EliteHandler) idRespond(w http.ResponseWriter, r *http.Request) {
	body, ok := decodeBody(w, r)
	if !ok {
		return
	}
	charID := strOf(body, "character_id")
	req, err := eh.db.GetServiceRequest(strOf(body, "request_id"))
	if err != nil || req.TargetID != charID || req.Kind != "imagedesign" {
		writeCraftJSON(w, http.StatusNotFound, map[string]string{"error": "request unknown"})
		return
	}
	if strOf(body, "accept") != "yes" {
		_ = eh.db.ResolveServiceRequest(req.ID, "declined")
		writeCraftJSON(w, http.StatusOK, map[string]string{"status": "declined"})
		return
	}
	var changes map[string]interface{}
	if err := json.Unmarshal([]byte(req.Detail), &changes); err != nil || len(changes) == 0 {
		writeCraftJSON(w, http.StatusBadRequest, map[string]string{"error": "empty change set"})
		return
	}
	for field := range changes {
		gate, known := idFieldGates[field]
		if !known {
			writeCraftJSON(w, http.StatusBadRequest, map[string]string{"error": "field not editable: " + field})
			return
		}
		if eh.w.skillTier(req.RequesterID, gate) < 1 {
			writeCraftJSON(w, http.StatusForbidden, map[string]string{"error": "designer lacks " + field + " training"})
			return
		}
	}
	if req.Fee > 0 {
		bal, err := eh.db.GetCharacterCredits(charID)
		if err != nil || bal < req.Fee {
			writeCraftJSON(w, http.StatusBadRequest, map[string]string{"error": "insufficient credits for fee"})
			return
		}
	}
	// Apply via sequenced DB calls (appearance + wallet + resolve); each is
	// atomic, the trio best-effort ordered (fee moves only after a clean edit).
	if err := eh.db.UpdateAppearance(charID, changes); err != nil {
		writeCraftJSON(w, http.StatusBadRequest, map[string]string{"error": "edit failed: " + err.Error()})
		return
	}
	if req.Fee > 0 {
		if err := eh.db.DeductCredits(charID, req.Fee); err != nil {
			writeCraftJSON(w, http.StatusBadRequest, map[string]string{"error": "fee failed"})
			return
		}
		_ = eh.db.AddCredits(req.RequesterID, req.Fee)
		_ = eh.db.RecordLedger(charID, -req.Fee, "image_design", "transfer", req.RequesterID, "")
		_ = eh.db.RecordLedger(req.RequesterID, req.Fee, "image_design", "transfer", charID, "")
	}
	_ = eh.db.ResolveServiceRequest(req.ID, "accepted")
	_ = eh.db.AddCharacterXP(req.RequesterID, "image_designer", 50) // [PROVISIONAL]
	writeCraftJSON(w, http.StatusOK, map[string]string{"status": "applied"})
}

func (eh *EliteHandler) charByName(w http.ResponseWriter, name string) (string, bool) {
	id, err := eh.db.CharacterIDByName(name)
	if err != nil || id == "" {
		writeCraftJSON(w, http.StatusNotFound, map[string]string{"error": "character unknown"})
		return "", false
	}
	return id, true
}

// --- Holoemotes ---

func (eh *EliteHandler) holoCreate(w http.ResponseWriter, r *http.Request) {
	body, ok := decodeBody(w, r)
	if !ok {
		return
	}
	charID := strOf(body, "character_id")
	if eh.w.skillTier(charID, "imagedesigner_holoemote_design_") < 1 {
		writeCraftJSON(w, http.StatusForbidden, map[string]string{"error": "requires Holoemote Design training"})
		return
	}
	name, text := strOf(body, "name"), strOf(body, "text")
	if !validHoloName(name) || len(text) == 0 || len(text) > 200 {
		writeCraftJSON(w, http.StatusBadRequest, map[string]string{"error": "bad name or text"})
		return
	}
	price := intOf(body, "price")
	if price < 0 {
		writeCraftJSON(w, http.StatusBadRequest, map[string]string{"error": "price must not be negative"})
		return
	}
	id, err := eh.db.CreateHoloemote(charID, name, text, price)
	if err != nil {
		writeCraftJSON(w, http.StatusBadRequest, map[string]string{"error": "create failed: " + err.Error()})
		return
	}
	_ = eh.db.AddCharacterXP(charID, "image_designer", 50) // [PROVISIONAL]
	writeCraftJSON(w, http.StatusCreated, map[string]string{"holo_id": id})
}

func validHoloName(name string) bool {
	if len(name) < 2 || len(name) > 24 {
		return false
	}
	for _, c := range name {
		if !(c >= 'a' && c <= 'z') && c != '_' {
			return false
		}
	}
	return true
}

func (eh *EliteHandler) holoBuy(w http.ResponseWriter, r *http.Request) {
	body, ok := decodeBody(w, r)
	if !ok {
		return
	}
	charID, holoID := strOf(body, "character_id"), strOf(body, "holo_id")
	if err := eh.db.BuyHoloemote(holoID, charID); err != nil {
		writeCraftJSON(w, http.StatusBadRequest, map[string]string{"error": "buy failed: " + err.Error()})
		return
	}
	writeCraftJSON(w, http.StatusOK, map[string]string{"status": "rights granted"})
}

func (eh *EliteHandler) holoList(w http.ResponseWriter, r *http.Request, charID string) {
	list, err := eh.db.AllHoloemotes()
	if err != nil {
		writeCraftJSON(w, http.StatusInternalServerError, map[string]string{"error": "lookup failed"})
		return
	}
	if list == nil {
		list = []database.HoloemoteRow{}
	}
	writeCraftJSON(w, http.StatusOK, map[string]interface{}{"holoemotes": list})
}

// --- Slicing handshake (Smuggler service on others' items) ---

func (eh *EliteHandler) sliceRequest(w http.ResponseWriter, r *http.Request) {
	body, ok := decodeBody(w, r)
	if !ok {
		return
	}
	charID := strOf(body, "character_id")
	if eh.w.skillTier(charID, "smuggler_slicing_") < 1 {
		writeCraftJSON(w, http.StatusForbidden, map[string]string{"error": "requires Slicing training"})
		return
	}
	targetID, ok := eh.charByName(w, strOf(body, "target_name"))
	if !ok {
		return
	}
	itemID := strOf(body, "item_id")
	item, err := eh.db.GetItem(targetID, itemID)
	if err != nil {
		writeCraftJSON(w, http.StatusNotFound, map[string]string{"error": "target item unknown"})
		return
	}
	if item.Schematic != "duelist_pistol" && item.Schematic != "marksman_rifle" &&
		item.Schematic != "patrol_carbine" && item.Schematic != "dueling_blade" &&
		item.Schematic != "battle_sword" && item.Schematic != "war_pike" &&
		item.Schematic != "basic_sidearm" && item.Schematic != "basic_plate" &&
		item.Schematic != "armor_composite" && item.Schematic != "armor_bone" &&
		item.Schematic != "armor_chitin" {
		writeCraftJSON(w, http.StatusBadRequest, map[string]string{"error": "item not sliceable"})
		return
	}
	if v, ok := item.Stats["sliced"]; ok && v != 0 {
		writeCraftJSON(w, http.StatusBadRequest, map[string]string{"error": "already sliced (one-time)"})
		return
	}
	fee := intOf(body, "fee")
	if fee < 0 {
		writeCraftJSON(w, http.StatusBadRequest, map[string]string{"error": "fee must not be negative"})
		return
	}
	id, err := eh.db.CreateServiceRequest("slice", charID, targetID, itemID, "", fee)
	if err != nil {
		writeCraftJSON(w, http.StatusInternalServerError, map[string]string{"error": "request failed"})
		return
	}
	writeCraftJSON(w, http.StatusCreated, map[string]string{"request_id": id})
}

// sliceRespond applies the one-time boost (+10% damage/protection max —
// provisional §6.9) with the fee transfer.
func (eh *EliteHandler) sliceRespond(w http.ResponseWriter, r *http.Request) {
	body, ok := decodeBody(w, r)
	if !ok {
		return
	}
	charID := strOf(body, "character_id")
	req, err := eh.db.GetServiceRequest(strOf(body, "request_id"))
	if err != nil || req.TargetID != charID || req.Kind != "slice" {
		writeCraftJSON(w, http.StatusNotFound, map[string]string{"error": "request unknown"})
		return
	}
	if strOf(body, "accept") != "yes" {
		_ = eh.db.ResolveServiceRequest(req.ID, "declined")
		writeCraftJSON(w, http.StatusOK, map[string]string{"status": "declined"})
		return
	}
	item, err := eh.db.GetItem(charID, req.ItemID)
	if err != nil {
		writeCraftJSON(w, http.StatusBadRequest, map[string]string{"error": "item gone"})
		return
	}
	if v, ok := item.Stats["sliced"]; ok && v != 0 {
		writeCraftJSON(w, http.StatusBadRequest, map[string]string{"error": "already sliced"})
		return
	}
	if req.Fee > 0 {
		bal, err := eh.db.GetCharacterCredits(charID)
		if err != nil || bal < req.Fee {
			writeCraftJSON(w, http.StatusBadRequest, map[string]string{"error": "insufficient credits for fee"})
			return
		}
	}
	stats := item.Stats
	if stats == nil {
		stats = map[string]float64{}
	}
	for _, k := range []string{"damage_max", "protection_max"} {
		if v, ok := stats[k]; ok {
			stats[k] = v * 1.1
		}
	}
	stats["sliced"] = 1
	blob, _ := json.Marshal(stats)
	if err := eh.db.UpdateItemStats(charID, req.ItemID, string(blob)); err != nil {
		writeCraftJSON(w, http.StatusInternalServerError, map[string]string{"error": "slice failed"})
		return
	}
	if req.Fee > 0 {
		_ = eh.db.DeductCredits(charID, req.Fee)
		_ = eh.db.AddCredits(req.RequesterID, req.Fee)
		_ = eh.db.RecordLedger(charID, -req.Fee, "slicing_fee", "transfer", req.RequesterID, "")
		_ = eh.db.RecordLedger(req.RequesterID, req.Fee, "slicing_fee", "transfer", charID, "")
	}
	_ = eh.db.ResolveServiceRequest(req.ID, "accepted")
	_ = eh.db.AddCharacterXP(req.RequesterID, "smuggler", 50) // [PROVISIONAL]
	writeCraftJSON(w, http.StatusOK, map[string]string{"status": "sliced"})
}

// --- Spice + buff-pack use (stim-pattern consumables) ---

func (eh *EliteHandler) spiceUse(w http.ResponseWriter, r *http.Request) {
	body, ok := decodeBody(w, r)
	if !ok {
		return
	}
	charID := strOf(body, "character_id")
	eh.useChargedBuff(w, charID, charID, strOf(body, "item_id"), "smuggler")
}

func (eh *EliteHandler) packUse(w http.ResponseWriter, r *http.Request) {
	body, ok := decodeBody(w, r)
	if !ok {
		return
	}
	charID := strOf(body, "character_id")
	targetID := strOf(body, "target_id")
	if targetID == "" {
		targetID = charID
	}
	eh.useChargedBuff(w, charID, targetID, strOf(body, "item_id"), "medical")
}

// useChargedBuff applies a charged buff consumable to a conscious in-range
// target (stim-use precedent: range + conscious, no skill gate — access is via
// crafting, i.e. interdependence through production not permission).
// xpPool is smuggler (spice) or medical (packs, stim parity).
func (eh *EliteHandler) useChargedBuff(w http.ResponseWriter, charID, targetID, itemID, xpPool string) {
	item, err := eh.db.GetItem(charID, itemID)
	if err != nil {
		writeCraftJSON(w, http.StatusNotFound, map[string]string{"error": "item not found"})
		return
	}
	poolBySchem := map[string]string{
		"spice_rush": "action", "spice_calm": "mind",
		"buff_pack_health": "health", "buff_pack_action": "action", "buff_pack_mind": "mind",
	}
	pool, known := poolBySchem[item.Schematic]
	if !known {
		writeCraftJSON(w, http.StatusBadRequest, map[string]string{"error": "not a buff consumable"})
		return
	}
	charges := 0
	if v, ok := item.Stats["charges_remaining"]; ok {
		charges = int(v)
	} else if v, ok := item.Stats["charges_max"]; ok {
		charges = int(v)
	}
	if charges <= 0 {
		writeCraftJSON(w, http.StatusBadRequest, map[string]string{"error": "consumable depleted"})
		return
	}
	target, ok := eh.w.findClient(targetID)
	if !ok {
		writeCraftJSON(w, http.StatusBadRequest, map[string]string{"error": "target not in world"})
		return
	}
	me, ok := eh.w.findClient(charID)
	if !ok {
		writeCraftJSON(w, http.StatusBadRequest, map[string]string{"error": "not in world"})
		return
	}
	dx := me.Pos.X - target.Pos.X
	dz := me.Pos.Z - target.Pos.Z
	if dx*dx+dz*dz > 64 { // 8 m application range (HealRangeM analogue)
		writeCraftJSON(w, http.StatusBadRequest, map[string]string{"error": "target out of range"})
		return
	}
	tst, err := eh.db.GetCombatState(targetID)
	if err != nil || tst.IncapacitatedAt != nil {
		writeCraftJSON(w, http.StatusBadRequest, map[string]string{"error": "target must be conscious"})
		return
	}
	amount := 0.0
	if v, ok := item.Stats["potency_max"]; ok {
		amount = v
	}
	if amount <= 0 {
		writeCraftJSON(w, http.StatusBadRequest, map[string]string{"error": "consumable inert"})
		return
	}
	if err := eh.db.AddBuff(targetID, charID, pool, int(amount),
		time.Now().Add(30*time.Minute)); err != nil {
		writeCraftJSON(w, http.StatusBadRequest, map[string]string{"error": "buff already active"})
		return
	}
	charges--
	if charges <= 0 {
		_ = eh.db.DeleteItem(charID, itemID)
	} else {
		item.Stats["charges_remaining"] = float64(charges)
		if blob, err := json.Marshal(item.Stats); err == nil {
			_ = eh.db.UpdateItemStats(charID, itemID, string(blob))
		}
	}
	eh.w.pushHAMUpdate(targetID)
	_ = eh.db.AddCharacterXP(charID, xpPool, 25) // [PROVISIONAL]
	writeCraftJSON(w, http.StatusOK, map[string]interface{}{
		"status": "applied", "charges_left": charges,
	})
}
