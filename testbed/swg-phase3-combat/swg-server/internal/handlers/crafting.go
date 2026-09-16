// Crafting HTTP handlers for Phase 4: schematic registry, server-side crafting
// sessions (start/assign/assemble/experiment/finalize), stacks/items queries,
// and equip. Transactional request/response shape (matches the skills-handler
// pattern). Testbed fork; generic content only.
package handlers

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"strings"
	"sync"
	"time"

	"swg-server/internal/crafting"
	"swg-server/internal/database"
)

// serverRNG returns a fresh RNG per call (no shared-state races).
func serverRNG() *rand.Rand {
	return rand.New(rand.NewSource(time.Now().UnixNano()))
}

// CraftHandler owns the in-memory crafting session store.
type CraftHandler struct {
	db       *database.DB
	mu       sync.Mutex
	sessions map[string]*crafting.Session
}

// NewCraftHandler creates the handler.
func NewCraftHandler(db *database.DB) *CraftHandler {
	return &CraftHandler{db: db, sessions: make(map[string]*crafting.Session)}
}

// HandleCraftRoute dispatches /api/craft/* sub-paths (AuthMiddleware-wrapped).
func (ch *CraftHandler) HandleCraftRoute(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimPrefix(r.URL.Path, "/api/craft/")
	charID := r.URL.Query().Get("character_id")
	switch {
	case path == "schematics" && r.Method == "GET":
		ch.listSchematics(w, r)
	case path == "sessions/start" && r.Method == "POST":
		ch.startSession(w, r)
	case path == "sessions/assign" && r.Method == "POST":
		ch.assignSlot(w, r)
	case path == "sessions/assemble" && r.Method == "POST":
		ch.assemble(w, r)
	case path == "sessions/experiment" && r.Method == "POST":
		ch.experiment(w, r)
	case path == "sessions/finalize" && r.Method == "POST":
		ch.finalize(w, r)
	case path == "resources" && r.Method == "GET":
		ch.listStacks(w, r, charID)
	case path == "spawns" && r.Method == "GET":
		ch.listSpawns(w, r, charID)
	case path == "items" && r.Method == "GET":
		ch.listItems(w, r, charID)
	case path == "equip" && r.Method == "POST":
		ch.equip(w, r)
	case path == "harvester/fund" && r.Method == "POST":
		ch.fundHarvester(w, r)
	default:
		writeCraftJSON(w, http.StatusNotFound, map[string]string{"error": "unknown craft route"})
	}
}

func writeCraftJSON(w http.ResponseWriter, code int, v interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	_ = json.NewEncoder(w).Encode(v)
}

func decodeBody(w http.ResponseWriter, r *http.Request) (map[string]interface{}, bool) {
	var body map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		writeCraftJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid JSON"})
		return nil, false
	}
	return body, true
}

func strOf(body map[string]interface{}, key string) string {
	if v, ok := body[key].(string); ok {
		return v
	}
	return ""
}

func (ch *CraftHandler) artisanBoxes(characterID string) int {
	owned, err := ch.db.GetCharacterSkillBoxes(characterID)
	if err != nil {
		return 0
	}
	n := 0
	for id := range owned {
		if strings.HasPrefix(id, "artisan_") {
			n++
		}
	}
	return n
}

func (ch *CraftHandler) getSession(characterID, sessionID string) (*crafting.Session, bool) {
	ch.mu.Lock()
	defer ch.mu.Unlock()
	s, ok := ch.sessions[sessionID]
	if !ok || s.OwnerCharID != characterID {
		return nil, false
	}
	return s, true
}

func (ch *CraftHandler) listSchematics(w http.ResponseWriter, r *http.Request) {
	type slotView struct {
		ID string `json:"id"`
		Label string `json:"label"`
		AcceptedTypes []string `json:"accepted_types"`
		UnitsRequired int `json:"units_required"`
	}
	type propView struct {
		ID string `json:"id"`
		Label string `json:"label"`
		BaseMin float64 `json:"base_min"`
		BaseMax float64 `json:"base_max"`
	}
	type schemView struct {
		ID string `json:"id"`
		Name string `json:"name"`
		ProfessionGate string `json:"profession_gate"`
		Complexity int `json:"complexity"`
		ExperimentPoints int `json:"experiment_points"`
		MaxRounds int `json:"max_rounds"`
		EquipSlot string `json:"equip_slot"`
		XPPool string `json:"xp_pool"`
		XPReward int `json:"xp_reward"`
		WeaponStyle string `json:"weapon_style"`
		EquipGate string `json:"equip_gate"`
		Slots []slotView `json:"slots"`
		Properties []propView `json:"properties"`
	}
	out := []schemView{}
	for _, sc := range crafting.Schematics {
		v := schemView{ID: sc.ID, Name: sc.Name, ProfessionGate: sc.ProfessionGate,
			Complexity: sc.Complexity, ExperimentPoints: sc.ExperimentPoints,
			MaxRounds: sc.MaxRounds, EquipSlot: sc.EquipSlot,
			XPPool: sc.XPPool, XPReward: sc.XPReward,
			WeaponStyle: sc.WeaponStyle, EquipGate: sc.EquipGate}
		for _, s := range sc.Slots {
			v.Slots = append(v.Slots, slotView{ID: s.ID, Label: s.Label,
				AcceptedTypes: s.AcceptedTypes, UnitsRequired: s.UnitsRequired})
		}
		for _, p := range sc.Properties {
			v.Properties = append(v.Properties, propView{ID: p.ID, Label: p.Label,
				BaseMin: p.BaseMin, BaseMax: p.BaseMax})
		}
		out = append(out, v)
	}
	writeCraftJSON(w, http.StatusOK, map[string]interface{}{"schematics": out})
}

func (ch *CraftHandler) startSession(w http.ResponseWriter, r *http.Request) {
	body, ok := decodeBody(w, r)
	if !ok {
		return
	}
	charID, schemID := strOf(body, "character_id"), strOf(body, "schematic_id")
	sc := crafting.SchematicByID(schemID)
	if sc == nil {
		writeCraftJSON(w, http.StatusNotFound, map[string]string{"error": "unknown schematic"})
		return
	}
	// Profession gate (MVP: Artisan Novice auto-grants knowledge of all MVP
	// schematics — owner decision §9.4; elite gating deferred per §9.7).
	has, err := ch.db.HasSkillBox(charID, sc.ProfessionGate)
	if err != nil || !has {
		writeCraftJSON(w, http.StatusForbidden,
			map[string]string{"error": "requires " + sc.ProfessionGate})
		return
	}
	ch.mu.Lock()
	sid := ch.db.NewRowID("sess")
	s := crafting.NewSession(sid, sc.ID, charID, sc.ExperimentPoints)
	ch.sessions[sid] = s
	ch.mu.Unlock()
	writeCraftJSON(w, http.StatusCreated, map[string]interface{}{
		"session_id": sid, "phase": string(s.Phase), "points": s.PointsRemaining,
	})
}

func (ch *CraftHandler) assignSlot(w http.ResponseWriter, r *http.Request) {
	body, ok := decodeBody(w, r)
	if !ok {
		return
	}
	charID, sid := strOf(body, "character_id"), strOf(body, "session_id")
	s, ok := ch.getSession(charID, sid)
	if !ok {
		writeCraftJSON(w, http.StatusNotFound, map[string]string{"error": "unknown session"})
		return
	}
	sc := crafting.SchematicByID(s.SchematicID)
	slotID := strOf(body, "slot_id")
	var slot *crafting.SchematicSlot
	for _, sl := range sc.Slots {
		if sl.ID == slotID {
			slot = &sl
			break
		}
	}
	if slot == nil {
		writeCraftJSON(w, http.StatusBadRequest, map[string]string{"error": "unknown slot"})
		return
	}
	stacks, err := ch.db.GetStacks(charID)
	if err != nil {
		writeCraftJSON(w, http.StatusInternalServerError, map[string]string{"error": "stack lookup failed"})
		return
	}
	// A slot draws from the character's whole pool of accepted-type stacks
	// (different spawn IDs never merge in inventory, but one slot may consume
	// from several stacks — the preferred stack is recorded for quality).
	acceptSet := map[string]bool{}
	for _, t := range slot.AcceptedTypes {
		acceptSet[t] = true
	}
	total, primary, primaryQty := 0, "", -1
	for _, st := range stacks {
		if !acceptSet[st.ResourceType] {
			continue
		}
		total += st.Quantity
		if st.Quantity > primaryQty {
			primary, primaryQty = st.SpawnID, st.Quantity
		}
	}
	if total < slot.UnitsRequired {
		writeCraftJSON(w, http.StatusBadRequest, map[string]string{"error": "insufficient units in stack"})
		return
	}
	ch.mu.Lock()
	s.Assignments[slotID] = crafting.Assignment{SlotID: slotID, SpawnID: primary, Units: slot.UnitsRequired}
	ch.mu.Unlock()
	writeCraftJSON(w, http.StatusOK, map[string]string{"status": "assigned"})
}

func (ch *CraftHandler) assemble(w http.ResponseWriter, r *http.Request) {
	body, ok := decodeBody(w, r)
	if !ok {
		return
	}
	charID, sid := strOf(body, "character_id"), strOf(body, "session_id")
	s, ok := ch.getSession(charID, sid)
	if !ok {
		writeCraftJSON(w, http.StatusNotFound, map[string]string{"error": "unknown session"})
		return
	}
	sc := crafting.SchematicByID(s.SchematicID)
	stacks, _ := ch.db.GetStacks(charID)
	have := map[string]int{}
	for _, st := range stacks {
		have[st.SpawnID] = st.Quantity
	}
	ch.mu.Lock()
	err := crafting.Assemble(s, sc, have)
	ch.mu.Unlock()
	if err != nil {
		writeCraftJSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}
	writeCraftJSON(w, http.StatusOK, map[string]string{"phase": string(s.Phase)})
}

func (ch *CraftHandler) experiment(w http.ResponseWriter, r *http.Request) {
	body, ok := decodeBody(w, r)
	if !ok {
		return
	}
	charID, sid := strOf(body, "character_id"), strOf(body, "session_id")
	s, ok := ch.getSession(charID, sid)
	if !ok {
		writeCraftJSON(w, http.StatusNotFound, map[string]string{"error": "unknown session"})
		return
	}
	sc := crafting.SchematicByID(s.SchematicID)
	propID := strOf(body, "property_id")
	points := 0
	if v, ok := body["points"].(float64); ok {
		points = int(v)
	}
	th := crafting.ThresholdsFor(ch.artisanBoxes(charID))
	ch.mu.Lock()
	out, delta, err := crafting.Experiment(s, sc, propID, points, th, serverRNG())
	if err == nil && out == crafting.OutcomeCriticalSuccess {
		s.Crits++ // Phase 9: GDD 7.2.2 critical-success 2x XP tracking
	}
	phase := s.Phase
	remaining := s.PointsRemaining
	ch.mu.Unlock()
	if err != nil {
		writeCraftJSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}
	writeCraftJSON(w, http.StatusOK, map[string]interface{}{
		"outcome": string(out), "delta_pct": delta,
		"phase": string(phase), "points_remaining": remaining,
	})
}

func (ch *CraftHandler) finalize(w http.ResponseWriter, r *http.Request) {
	body, ok := decodeBody(w, r)
	if !ok {
		return
	}
	charID, sid := strOf(body, "character_id"), strOf(body, "session_id")
	ch.mu.Lock()
	s, ok := ch.sessions[sid]
	if !ok || s.OwnerCharID != charID {
		ch.mu.Unlock()
		writeCraftJSON(w, http.StatusNotFound, map[string]string{"error": "unknown session"})
		return
	}
	sc := crafting.SchematicByID(s.SchematicID)
	if s.Phase != crafting.PhaseExperimentation {
		ch.mu.Unlock()
		writeCraftJSON(w, http.StatusBadRequest, map[string]string{"error": "session not ready"})
		return
	}
	name := strOf(body, "name")
	if name == "" {
		name = sc.Name + " (custom)"
	}
	// Snapshot under lock; DB work outside it.
	assigns := make(map[string]crafting.Assignment)
	for k, v := range s.Assignments {
		assigns[k] = v
	}
	bonuses := make(map[string]float64)
	for k, v := range s.Bonuses {
		bonuses[k] = v
	}
	crits := s.Crits
	ch.mu.Unlock()

	qualities := make(map[string]float64)
	consumes := []consumeOp{}
	stacks, err := ch.db.GetStacks(charID)
	if err != nil {
		writeCraftJSON(w, http.StatusBadRequest,
			map[string]string{"error": "stack lookup failed"})
		return
	}
	for slotID, a := range assigns {
		var slotUnits int
		var acceptSet map[string]bool
		for _, sl := range sc.Slots {
			if sl.ID == slotID {
				slotUnits = sl.UnitsRequired
				acceptSet = map[string]bool{}
				for _, t := range sl.AcceptedTypes {
					acceptSet[t] = true
				}
			}
		}
		// Greedy consume across accepted stacks, richest first; quality is the
		// quantity-weighted mean OQ of what was actually consumed.
		var cands []stackCand
		for _, st := range stacks {
			if acceptSet[st.ResourceType] && st.Quantity > 0 {
				oq := 0
				if st.Stats != nil {
					oq = st.Stats["OQ"]
				}
				cands = append(cands, stackCand{st.SpawnID, st.Quantity, oq})
			}
		}
		sortByQtyDesc(cands)
		need := a.Units
		if need <= 0 {
			need = slotUnits
		}
		qSum, qN := 0, 0
		for _, c := range cands {
			if need <= 0 {
				break
			}
			take := c.qty
			if take > need {
				take = need
			}
			consumes = append(consumes, consumeOp{spawnID: c.id, qty: take})
			qSum += c.oq * take
			qN += take
			need -= take
		}
		if need > 0 {
			writeCraftJSON(w, http.StatusBadRequest,
				map[string]string{"error": "resource shortfall at finalize: " + slotID})
			return
		}
		if qN > 0 {
			qualities[slotID] = float64(qSum) / float64(qN)
		}
	}
	for _, c := range consumes {
		if err := ch.db.ConsumeStack(charID, c.spawnID, c.qty); err != nil {
			writeCraftJSON(w, http.StatusBadRequest,
				map[string]string{"error": "resource shortfall at finalize: " + err.Error()})
			return
		}
	}
	final := crafting.FinalStats(sc, bonuses, qualities)
	flat := map[string]float64{}
	for pid, mm := range final {
		flat[pid+"_min"] = mm[0]
		flat[pid+"_max"] = mm[1]
	}
	// Charge-based consumables (e.g. stim packs) start full: charges_remaining
	// tracks live charges separately from the crafted charges_max stat.
	if _, ok := flat["charges_max"]; ok {
		flat["charges_remaining"] = flat["charges_max"]
	}
	// Phase 9: optional tissue ingredient (Bio-Engineer trade good) — a
	// tissue_sample item named in the request adds +5 protection_max to armor
	// and is consumed on success (flagged; avoids session-model surgery for
	// GDD 8.4's "optional" tissue).
	tissueID := strOf(body, "tissue_item_id")
	if tissueID != "" {
		if sc.EquipSlot != "armor" {
			writeCraftJSON(w, http.StatusBadRequest, map[string]string{"error": "tissue only applies to armor"})
			return
		}
		tissue, err := ch.db.GetItem(charID, tissueID)
		if err != nil || tissue.Schematic != "tissue_sample" {
			writeCraftJSON(w, http.StatusBadRequest, map[string]string{"error": "valid tissue sample required"})
			return
		}
		flat["protection_max"] += 5
	}
	itemID, err := ch.db.InsertItem(charID, sc.ID, name, flat)
	if err != nil {
		writeCraftJSON(w, http.StatusInternalServerError, map[string]string{"error": "item persist failed"})
		return
	}
	if tissueID != "" {
		_ = ch.db.DeleteItem(charID, tissueID)
	}
	// Phase 9: schematic XP pools (elite pools; "" = generic crafting,
	// preserving all existing schematics byte-for-byte), GDD 7.2.2 repeat
	// diminishing (50% from the 11th craft of one schematic) and critical
	// doubling (any critical success this session).
	pool := sc.XPPool
	if pool == "" {
		pool = "crafting"
	}
	xp := sc.XPReward
	if crits > 0 {
		xp *= 2
	}
	if n, err := ch.db.BumpCraftCount(charID, sc.ID); err == nil && n > 10 {
		xp /= 2
	}
	_ = ch.db.AddCharacterXP(charID, pool, xp)
	ch.mu.Lock()
	s.Phase = crafting.PhaseComplete
	delete(ch.sessions, sid)
	ch.mu.Unlock()
	writeCraftJSON(w, http.StatusCreated, map[string]interface{}{
		"item_id": itemID, "stats": flat,
	})
}

// stackCand is one candidate stack for greedy finalize consumption.
type stackCand struct {
	id  string
	qty int
	oq  int
}

// consumeOp is one stack deduction at finalize.
type consumeOp struct {
	spawnID string
	qty     int
}

// sortByQtyDesc orders candidates richest-first (insertion sort; tiny N).
func sortByQtyDesc(c []stackCand) {
	for i := 1; i < len(c); i++ {
		for j := i; j > 0 && c[j].qty > c[j-1].qty; j-- {
			c[j], c[j-1] = c[j-1], c[j]
		}
	}
}

func (ch *CraftHandler) listStacks(w http.ResponseWriter, r *http.Request, charID string) {
	if charID == "" {
		writeCraftJSON(w, http.StatusBadRequest, map[string]string{"error": "character_id required"})
		return
	}
	stacks, err := ch.db.GetStacks(charID)
	if err != nil {
		writeCraftJSON(w, http.StatusInternalServerError, map[string]string{"error": "stack lookup failed"})
		return
	}
	writeCraftJSON(w, http.StatusOK, map[string]interface{}{"stacks": stacks})
}

func (ch *CraftHandler) listItems(w http.ResponseWriter, r *http.Request, charID string) {
	if charID == "" {
		writeCraftJSON(w, http.StatusBadRequest, map[string]string{"error": "character_id required"})
		return
	}
	items, err := ch.db.GetItems(charID)
	if err != nil {
		writeCraftJSON(w, http.StatusInternalServerError, map[string]string{"error": "item lookup failed"})
		return
	}
	writeCraftJSON(w, http.StatusOK, map[string]interface{}{"items": items})
}

// listSpawns returns active spawn summaries (test/introspection aid).
func (ch *CraftHandler) listSpawns(w http.ResponseWriter, r *http.Request, charID string) {
	if charID == "" {
		writeCraftJSON(w, http.StatusBadRequest, map[string]string{"error": "character_id required"})
		return
	}
	char, err := ch.db.GetCharacterByID(charID)
	if err != nil {
		writeCraftJSON(w, http.StatusNotFound, map[string]string{"error": "character unknown"})
		return
	}
	spawns, err := ch.db.ActiveSpawns(char.Planet)
	if err != nil {
		writeCraftJSON(w, http.StatusInternalServerError, map[string]string{"error": "spawn lookup failed"})
		return
	}
	type spawnView struct {
		ID   string `json:"id"`
		Type string `json:"type"`
	}
	out := []spawnView{}
	for _, s := range spawns {
		out = append(out, spawnView{ID: s.ID, Type: s.Type})
	}
	writeCraftJSON(w, http.StatusOK, map[string]interface{}{"spawns": out})
}

func (ch *CraftHandler) equip(w http.ResponseWriter, r *http.Request) {
	body, ok := decodeBody(w, r)
	if !ok {
		return
	}
	charID, itemID := strOf(body, "character_id"), strOf(body, "item_id")
	item, err := ch.db.GetItem(charID, itemID)
	if err != nil {
		writeCraftJSON(w, http.StatusNotFound, map[string]string{"error": "item not found"})
		return
	}
	sc := crafting.SchematicByID(item.Schematic)
	if sc == nil || (sc.EquipSlot != "weapon" && sc.EquipSlot != "armor") {
		writeCraftJSON(w, http.StatusBadRequest, map[string]string{"error": "item not equippable"})
		return
	}
	if err := ch.db.EquipItem(charID, itemID, sc.EquipSlot); err != nil {
		writeCraftJSON(w, http.StatusInternalServerError, map[string]string{"error": "equip failed"})
		return
	}
	writeCraftJSON(w, http.StatusOK, map[string]string{"status": "equipped", "slot": sc.EquipSlot})
}

// fundHarvester moves wallet credits into a harvester's maintenance pool.
func (ch *CraftHandler) fundHarvester(w http.ResponseWriter, r *http.Request) {
	body, ok := decodeBody(w, r)
	if !ok {
		return
	}
	charID, hvid := strOf(body, "character_id"), strOf(body, "harvester_id")
	credits := 0
	if v, ok := body["credits"].(float64); ok {
		credits = int(v)
	}
	if credits <= 0 {
		writeCraftJSON(w, http.StatusBadRequest, map[string]string{"error": "credits must be positive"})
		return
	}
	if err := ch.db.DeductCredits(charID, credits); err != nil {
		writeCraftJSON(w, http.StatusBadRequest,
			map[string]string{"error": "wallet deduction failed: " + err.Error()})
		return
	}
	if err := ch.db.FundHarvester(hvid, credits); err != nil {
		writeCraftJSON(w, http.StatusInternalServerError, map[string]string{"error": "funding failed"})
		return
	}
	// Phase 5 ledger: wallet→pool funding is a transfer (not faucet/sink).
	_ = ch.db.RecordLedger(charID, -credits, "harvester_fund", "transfer", hvid, "")
	writeCraftJSON(w, http.StatusOK, map[string]string{"status": "funded"})
}
