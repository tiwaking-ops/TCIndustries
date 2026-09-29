// Service handlers for Phase 6: wound healing, buffs, perform/watch,
// tips, and stim use, plus the service tick (BF healing, Mind drain, buff
// expiry, session interruption). Testbed fork; generic content only.
// GDD-silent numbers [PROVISIONAL] per the owner-approved convention.
package handlers

import (
	"encoding/json"
	"sync"
	"time"

	"swg-server/internal/protocol"
	"swg-server/internal/services"
	"swg-server/internal/world"
)

// performSession is one live performance (in-memory; restart/disconnect ends it,
// like disconnects end world presence — flagged scope limit).
type performSession struct {
	PerformerID string
	Kind        string // "music" | "dance"
	PosX, PosZ  float64
	StartedAt   time.Time
}

var (
	svcMu      sync.Mutex
	performing = make(map[string]*performSession) // performerID → session
	watching   = make(map[string]string)          // watcherID → performerID
	healMu     sync.Mutex
	lastHeal   = make(map[string]time.Time) // healerID+"\x00"+targetID → last heal
	// meditating tracks TKA meditation (Phase 9): performer-style stationary
	// state granting accelerated currents (gate-add magnitudes, flagged).
	meditating = make(map[string]medState) // characterID → start snapshot
)

// medState snapshots meditation start (movement breaks it, performer rule).
type medState struct {
	At     time.Time
	PosX   float64
	PosZ   float64
}

// medicTier counts owned boxes in a Medic tree (e.g. tree "healing" matches
// medic_healing_*). Entertainer tiers work the same way.
func (h *WorldHandler) skillTier(characterID, prefix string) int {
	owned, err := h.db.GetCharacterSkillBoxes(characterID)
	if err != nil {
		return 0
	}
	n := 0
	for id := range owned {
		if len(id) > len(prefix) && id[:len(prefix)] == prefix {
			n++
		}
	}
	if n > 4 {
		n = 4
	}
	return n
}

func (h *WorldHandler) hasBox(characterID, boxID string) bool {
	ok, err := h.db.HasSkillBox(characterID, boxID)
	return err == nil && ok
}

// SeedServiceWorld prepares Phase 6 persistence. Idempotent.
func (h *WorldHandler) SeedServiceWorld() error {
	return h.db.EnsureServiceSchema()
}

// --- S1: wound healing ---

func (h *WorldHandler) handleHealWounds(client *Client, raw []byte) {
	if client.CharacterID == "" {
		h.sendError(client, "not in world")
		return
	}
	var msg protocol.WSMessage
	json.Unmarshal(raw, &msg)
	dataBytes, _ := json.Marshal(msg.Data)
	var hm protocol.TargetMsg
	if err := json.Unmarshal(dataBytes, &hm); err != nil {
		h.sendError(client, "invalid heal_wounds data")
		return
	}
	if !h.hasBox(client.CharacterID, "medic_novice") {
		h.sendError(client, "requires Novice Medic")
		return
	}
	target, ok := h.findClient(hm.TargetID)
	if !ok {
		h.sendError(client, "heal target not in world")
		return
	}
	if world.Distance2D(client.Pos, target.Pos) > services.HealRangeM {
		h.sendError(client, "heal target out of range")
		return
	}
	tst, err := h.db.GetCombatState(hm.TargetID)
	if err != nil {
		h.sendError(client, "target combat state unavailable")
		return
	}
	if tst.WoundsHealth <= 0 {
		h.sendError(client, "target has no wounds")
		return
	}
	key := client.CharacterID + "\x00" + hm.TargetID
	healMu.Lock()
	last, seen := lastHeal[key]
	if seen && time.Since(last) < services.HealCooldown {
		healMu.Unlock()
		h.sendError(client, "heal cooling down")
		return
	}
	lastHeal[key] = time.Now()
	healMu.Unlock()
	tier := h.skillTier(client.CharacterID, "medic_healing_")
	amount := services.WoundHealAmount(tier)
	// Phase 9: Doctor (Advanced Healing) + Combat Medic (Field Triage) tiers
	// scale wound healing (gate-add, flagged — the fork has no stance-break
	// rule, so in-combat healing already works for everyone; numbers are the
	// only expressible elite benefit).
	amount += 10 * h.skillTier(client.CharacterID, "doctor_advanced_healing_")
	amount += 10 * h.skillTier(client.CharacterID, "combatmedic_field_triage_")
	tst.WoundsHealth -= amount
	if tst.WoundsHealth < 0 {
		tst.WoundsHealth = 0
	}
	if err := h.db.UpdateCombatState(tst); err != nil {
		h.sendError(client, "heal failed")
		return
	}
	_ = h.db.AddCharacterXP(client.CharacterID, "medical", services.HealXP)
	h.pushHAMUpdate(hm.TargetID)
	// Phase 10 B1: append-only interdependence telemetry (§34 economic-
	// interdependence instrumentation). Best-effort, observation only.
	_ = h.db.RecordInterdependenceEvent("heal", hm.TargetID, client.CharacterID, client.Pos.Planet)
	h.send(client, protocol.MsgWoundHealed, protocol.WoundHealedMsg{
		TargetID: hm.TargetID, Amount: amount, WoundsLeft: tst.WoundsHealth,
	})
}

// --- S3: buffs ---

func (h *WorldHandler) handleApplyBuff(client *Client, raw []byte) {
	if client.CharacterID == "" {
		h.sendError(client, "not in world")
		return
	}
	var msg protocol.WSMessage
	json.Unmarshal(raw, &msg)
	dataBytes, _ := json.Marshal(msg.Data)
	var bm protocol.BuffMsg
	if err := json.Unmarshal(dataBytes, &bm); err != nil {
		h.sendError(client, "invalid apply_buff data")
		return
	}
	if bm.Pool != services.BuffHealth && bm.Pool != services.BuffAction && bm.Pool != services.BuffMind {
		h.sendError(client, "unknown buff pool")
		return
	}
	if !h.hasBox(client.CharacterID, "medic_novice") {
		h.sendError(client, "requires Novice Medic")
		return
	}
	target, ok := h.findClient(bm.TargetID)
	if !ok {
		h.sendError(client, "buff target not in world")
		return
	}
	if world.Distance2D(client.Pos, target.Pos) > services.BuffRangeM {
		h.sendError(client, "buff target out of range")
		return
	}
	tst, err := h.db.GetCombatState(bm.TargetID)
	if err != nil || tst.IncapacitatedAt != nil {
		h.sendError(client, "target must be conscious")
		return
	}
	tier := h.skillTier(client.CharacterID, "medic_healing_")
	amount := services.BuffMagnitude(tier)
	if err := h.db.AddBuff(bm.TargetID, client.CharacterID, bm.Pool, amount,
		time.Now().Add(services.BuffDuration)); err != nil {
		h.sendError(client, "buff already active")
		return
	}
	_ = h.db.AddCharacterXP(client.CharacterID, "medical", services.BuffXP)
	h.pushHAMUpdate(bm.TargetID)
	// Phase 10 B1: interdependence telemetry (buff receipt), best-effort.
	_ = h.db.RecordInterdependenceEvent("buff", bm.TargetID, client.CharacterID, client.Pos.Planet)
	h.send(client, protocol.MsgBuffApplied, protocol.BuffAppliedMsg{
		TargetID: bm.TargetID, Pool: bm.Pool, Amount: amount,
	})
}

// --- S4: perform / watch ---

func (h *WorldHandler) handlePerformStart(client *Client, raw []byte) {
	if client.CharacterID == "" {
		h.sendError(client, "not in world")
		return
	}
	var msg protocol.WSMessage
	json.Unmarshal(raw, &msg)
	dataBytes, _ := json.Marshal(msg.Data)
	var pm protocol.PerformMsg
	if err := json.Unmarshal(dataBytes, &pm); err != nil {
		h.sendError(client, "invalid perform data")
		return
	}
	if !services.ValidKind(pm.Kind) {
		h.sendError(client, "unknown performance kind (music/dance)")
		return
	}
	if !h.hasBox(client.CharacterID, "entertainer_novice") {
		h.sendError(client, "requires Novice Entertainer")
		return
	}
	svcMu.Lock()
	performing[client.CharacterID] = &performSession{
		PerformerID: client.CharacterID, Kind: pm.Kind,
		PosX: client.Pos.X, PosZ: client.Pos.Z, StartedAt: time.Now(),
	}
	svcMu.Unlock()
	h.send(client, protocol.MsgPerformanceStarted, protocol.PerformMsg{Kind: pm.Kind})
}

func (h *WorldHandler) handlePerformStop(client *Client, raw []byte) {
	if client.CharacterID == "" {
		h.sendError(client, "not in world")
		return
	}
	svcMu.Lock()
	delete(performing, client.CharacterID)
	for watcher, perf := range watching {
		if perf == client.CharacterID {
			delete(watching, watcher)
		}
	}
	svcMu.Unlock()
	h.send(client, protocol.MsgPerformanceStopped, protocol.PerformMsg{})
}

// endPerformance is the tick-side stop (interrupt, incap, movement, disconnect).
func (h *WorldHandler) endPerformance(performerID string) {
	svcMu.Lock()
	delete(performing, performerID)
	for watcher, perf := range watching {
		if perf == performerID {
			delete(watching, watcher)
		}
	}
	svcMu.Unlock()
}

func (h *WorldHandler) handleWatchStart(client *Client, raw []byte) {
	if client.CharacterID == "" {
		h.sendError(client, "not in world")
		return
	}
	var msg protocol.WSMessage
	json.Unmarshal(raw, &msg)
	dataBytes, _ := json.Marshal(msg.Data)
	var wm protocol.WatchMsg
	if err := json.Unmarshal(dataBytes, &wm); err != nil {
		h.sendError(client, "invalid watch data")
		return
	}
	svcMu.Lock()
	_, performing := performing[wm.PerformerID]
	svcMu.Unlock()
	if !performing {
		h.sendError(client, "target is not performing")
		return
	}
	perf, ok := h.findClient(wm.PerformerID)
	if !ok {
		h.sendError(client, "performer not in world")
		return
	}
	dx := client.Pos.X - perf.Pos.X
	dz := client.Pos.Z - perf.Pos.Z
	if dx*dx+dz*dz > services.WatchRadiusM*services.WatchRadiusM {
		h.sendError(client, "too far from performer")
		return
	}
	svcMu.Lock()
	watching[client.CharacterID] = wm.PerformerID // re-watch switches (no double-dip)
	svcMu.Unlock()
	h.send(client, protocol.MsgWatchStarted, protocol.WatchMsg{PerformerID: wm.PerformerID})
}

func (h *WorldHandler) handleWatchStop(client *Client, raw []byte) {
	if client.CharacterID == "" {
		h.sendError(client, "not in world")
		return
	}
	svcMu.Lock()
	delete(watching, client.CharacterID)
	svcMu.Unlock()
	h.send(client, protocol.MsgWatchStopped, protocol.WatchMsg{})
}

// --- S6: tips ---

func (h *WorldHandler) handleTip(client *Client, raw []byte) {
	if client.CharacterID == "" {
		h.sendError(client, "not in world")
		return
	}
	var msg protocol.WSMessage
	json.Unmarshal(raw, &msg)
	dataBytes, _ := json.Marshal(msg.Data)
	var tm protocol.TipMsg
	if err := json.Unmarshal(dataBytes, &tm); err != nil {
		h.sendError(client, "invalid tip data")
		return
	}
	if tm.Amount <= 0 {
		h.sendError(client, "tip must be positive")
		return
	}
	// Tips go to a currently-watched performer (no remote tipping).
	svcMu.Lock()
	watched, watching := watching[client.CharacterID]
	svcMu.Unlock()
	if !watching || watched != tm.TargetID {
		h.sendError(client, "tip requires watching the performer")
		return
	}
	bal, err := h.db.GetCharacterCredits(client.CharacterID)
	if err != nil || bal < tm.Amount {
		h.sendError(client, "insufficient credits")
		return
	}
	if err := h.db.DeductCredits(client.CharacterID, tm.Amount); err != nil {
		h.sendError(client, "tip failed")
		return
	}
	if err := h.db.AddCredits(tm.TargetID, tm.Amount); err != nil {
		_ = h.db.AddCredits(client.CharacterID, tm.Amount) // refund on partial failure
		h.sendError(client, "tip failed")
		return
	}
	_ = h.db.RecordLedger(client.CharacterID, -tm.Amount, "tip", "transfer", tm.TargetID, client.Pos.Planet)
	_ = h.db.RecordLedger(tm.TargetID, tm.Amount, "tip", "transfer", client.CharacterID, client.Pos.Planet)
	_ = h.db.AddCharacterXP(tm.TargetID, "entertaining", services.TipXP)
	h.send(client, protocol.MsgTipReceipt, protocol.TipMsg{TargetID: tm.TargetID, Amount: tm.Amount})
	if perf, ok := h.findClient(tm.TargetID); ok {
		h.send(perf, protocol.MsgTipReceived, protocol.TipMsg{TargetID: tm.TargetID, Amount: tm.Amount})
	}
}

// --- S5: stim use ---

func (h *WorldHandler) handleUseStim(client *Client, raw []byte) {
	if client.CharacterID == "" {
		h.sendError(client, "not in world")
		return
	}
	var msg protocol.WSMessage
	json.Unmarshal(raw, &msg)
	dataBytes, _ := json.Marshal(msg.Data)
	var sm protocol.StimMsg
	if err := json.Unmarshal(dataBytes, &sm); err != nil {
		h.sendError(client, "invalid stim data")
		return
	}
	item, err := h.db.GetItem(client.CharacterID, sm.ItemID)
	if err != nil || item.Schematic != "stim_pack" {
		h.sendError(client, "stim pack required")
		return
	}
	charges := 0
	if v, ok := item.Stats["charges_remaining"]; ok {
		charges = int(v)
	} else if v, ok := item.Stats["charges_max"]; ok {
		charges = int(v)
	}
	if charges <= 0 {
		h.sendError(client, "stim pack depleted")
		return
	}
	target, ok := h.findClient(sm.TargetID)
	if !ok {
		h.sendError(client, "stim target not in world")
		return
	}
	if world.Distance2D(client.Pos, target.Pos) > services.HealRangeM {
		h.sendError(client, "stim target out of range")
		return
	}
	tst, err := h.db.GetCombatState(sm.TargetID)
	if err != nil {
		h.sendError(client, "target combat state unavailable")
		return
	}
	potency := 0.0
	if v, ok := item.Stats["potency_max"]; ok {
		potency = v
	}
	healing := 0.0
	if v, ok := item.Stats["healing_max"]; ok {
		healing = v
	}
	tst.WoundsHealth -= int(potency)
	if tst.WoundsHealth < 0 {
		tst.WoundsHealth = 0
	}
	tst.HealthCurrent += int(healing)
	if err := h.db.UpdateCombatState(tst); err != nil {
		h.sendError(client, "stim failed")
		return
	}
	charges--
	if charges <= 0 {
		_ = h.db.DeleteItem(client.CharacterID, sm.ItemID)
	} else {
		item.Stats["charges_remaining"] = float64(charges)
		if blob, err := json.Marshal(item.Stats); err == nil {
			_ = h.db.UpdateItemStats(client.CharacterID, sm.ItemID, string(blob))
		}
	}
	_ = h.db.AddCharacterXP(client.CharacterID, "medical", 25) // [PROVISIONAL]
	h.pushHAMUpdate(sm.TargetID)
	h.send(client, protocol.MsgStimUsed, protocol.StimUsedMsg{
		TargetID: sm.TargetID, ChargesLeft: charges,
	})
}

// --- Service tick (called from the world tick loop) ---

// tickServices runs every tick: Mind drain for performers, interruption checks.
// BF healing + buff expiry run every 5th tick (regen cadence).
func (h *WorldHandler) tickServices(tickCount int64) {
	h.tickPerformerDrain()
	if tickCount%5 != 0 {
		return
	}
	h.db.ExpireBuffs()
	h.tickPerformanceHealing()
}

// tickPerformerDrain charges Mind per tick and ends sessions on interrupt
// (incap), movement (>2 m, GDD stationary rule), or disconnect.
func (h *WorldHandler) tickPerformerDrain() {
	svcMu.Lock()
	sessions := make([]*performSession, 0, len(performing))
	for _, s := range performing {
		sessions = append(sessions, &performSession{
			PerformerID: s.PerformerID, Kind: s.Kind,
			PosX: s.PosX, PosZ: s.PosZ, StartedAt: s.StartedAt,
		})
	}
	svcMu.Unlock()
	for _, s := range sessions {
		cl, ok := h.findClient(s.PerformerID)
		if !ok {
			h.endPerformance(s.PerformerID)
			continue
		}
		st, err := h.db.GetCombatState(s.PerformerID)
		if err != nil || st.IncapacitatedAt != nil {
			h.endPerformance(s.PerformerID)
			continue
		}
		dx := cl.Pos.X - s.PosX
		dz := cl.Pos.Z - s.PosZ
		if dx*dx+dz*dz > 4 { // 2 m stationary rule (GDD 23.2.1)
			h.endPerformance(s.PerformerID)
			continue
		}
		st.MindCurrent -= services.MindDrainPerTick
		if st.MindCurrent < 0 {
			st.MindCurrent = 0
		}
		if err := h.db.UpdateCombatState(st); err != nil {
			continue
		}
		if tickPush(cl) {
			h.pushHAMUpdate(s.PerformerID)
		}
	}
}

// tickPush throttles performer ham_update pushes to every 5th drain tick.
var (
	pushMu      sync.Mutex
	pushCounter = map[string]int{}
)

func tickPush(cl *Client) bool {
	pushMu.Lock()
	defer pushMu.Unlock()
	pushCounter[cl.CharacterID]++
	if pushCounter[cl.CharacterID] >= 5 {
		pushCounter[cl.CharacterID] = 0
		return true
	}
	return false
}

// tickPerformanceHealing applies per-tick BF healing to watchers in radius.
func (h *WorldHandler) tickPerformanceHealing() {
	svcMu.Lock()
	pairs := make(map[string]string)
	for watcher, perf := range watching {
		pairs[watcher] = perf
	}
	svcMu.Unlock()
	for watcherID, perfID := range pairs {
		watcher, ok := h.findClient(watcherID)
		if !ok {
			continue
		}
		perf, ok := h.findClient(perfID)
		if !ok {
			continue
		}
		svcMu.Lock()
		sess, performing := performing[perfID]
		stillWatching := watching[watcherID] == perfID
		svcMu.Unlock()
		if !performing || !stillWatching {
			continue
		}
		dx := watcher.Pos.X - perf.Pos.X
		dz := watcher.Pos.Z - perf.Pos.Z
		if dx*dx+dz*dz > services.WatchRadiusM*services.WatchRadiusM {
			continue // left radius: session ends cleanly next watch action (23.5)
		}
		_ = sess
		wst, err := h.db.GetCombatState(watcherID)
		if err != nil || wst.BattleFatiguePct <= 0 {
			continue
		}
		rate := services.BFHealPerTick(
			h.skillTier(perfID, "entertainer_healing_"),
			h.perfTier(perfID, sess.Kind),
		)
		wst.BattleFatiguePct -= rate
		if wst.BattleFatiguePct < 0 {
			wst.BattleFatiguePct = 0
		}
		if err := h.db.UpdateCombatState(wst); err != nil {
			continue
		}
		h.pushHAMUpdate(watcherID)
		// Phase 9: elite Musician/Dancer performers grant watchers a small
		// health buff (+100, same-type machinery throttles duplicates —
		// gate-add magnitude, flagged).
		if h.elitePerformer(perfID) {
			_ = h.db.AddBuff(watcherID, perfID, "health", 100,
				time.Now().Add(5*time.Minute))
		}
	}
}

// perfTier returns the performer's Music or Dance tree tier for the given kind
// (Phase 9: elite Musician/Dancer boxes extend the base trees — flagged).
func (h *WorldHandler) perfTier(performerID, kind string) int {
	if kind == "music" {
		n := h.skillTier(performerID, "entertainer_music_")
		if e := h.skillTier(performerID, "musician_instrumentation_"); e > n {
			n = e
		}
		return n
	}
	n := h.skillTier(performerID, "entertainer_dance_")
	if e := h.skillTier(performerID, "dancer_dance_mastery_"); e > n {
		n = e
	}
	return n
}

// elitePerformer reports Musician/Dancer elite boxes (performance group-buff
// source, gate-add magnitudes flagged).
func (h *WorldHandler) elitePerformer(performerID string) bool {
	for _, prefix := range []string{
		"musician_instrumentation_", "musician_composition_",
		"musician_inspiration_", "musician_showmanship_",
		"dancer_dance_mastery_", "dancer_choreography_",
		"dancer_exotic_dances_", "dancer_stage_presence_",
	} {
		if h.skillTier(performerID, prefix) >= 1 {
			return true
		}
	}
	return false
}

// handleMeditate toggles TKA meditation (WS stateful presence, performing
// precedent). Stationary; ends on combat/incap/movement like performances.
func (h *WorldHandler) handleMeditate(client *Client, raw []byte) {
	if client.CharacterID == "" {
		h.sendError(client, "not in world")
		return
	}
	if h.skillTier(client.CharacterID, "teraskasi_meditative_techniques_") < 1 {
		h.sendError(client, "requires Meditative Techniques training")
		return
	}
	svcMu.Lock()
	if _, ok := meditating[client.CharacterID]; ok {
		delete(meditating, client.CharacterID)
		svcMu.Unlock()
		h.send(client, protocol.MsgMeditationStopped, protocol.PerformMsg{})
		return
	}
	meditating[client.CharacterID] = medState{
		At: time.Now(), PosX: client.Pos.X, PosZ: client.Pos.Z,
	}
	svcMu.Unlock()
	h.send(client, protocol.MsgMeditationStarted, protocol.PerformMsg{})
}

// tickMeditation grants accelerated currents to stationary meditators
// (+5/tick + 5/Meditative-tier — gate-add, flagged). Movement beyond 2 m,
// combat, incap, or disconnect ends it (performer stationary rule).
func (h *WorldHandler) tickMeditation() {
	svcMu.Lock()
	snap := make(map[string]medState, len(meditating))
	for id, ms := range meditating {
		snap[id] = ms
	}
	svcMu.Unlock()
	end := func(id string) {
		svcMu.Lock()
		delete(meditating, id)
		svcMu.Unlock()
	}
	for id, ms := range snap {
		cl, ok := h.findClient(id)
		if !ok {
			end(id)
			continue
		}
		dx := cl.Pos.X - ms.PosX
		dz := cl.Pos.Z - ms.PosZ
		if dx*dx+dz*dz > 4 {
			end(id)
			continue
		}
		st, err := h.db.GetCombatState(id)
		if err != nil || st.IncapacitatedAt != nil {
			end(id)
			continue
		}
		gain := 5 + 5*h.skillTier(id, "teraskasi_meditative_techniques_")
		st.HealthCurrent += gain
		if st.HealthCurrent > st.HealthMax {
			st.HealthCurrent = st.HealthMax
		}
		st.ActionCurrent += gain
		if st.ActionCurrent > st.ActionMax {
			st.ActionCurrent = st.ActionMax
		}
		st.MindCurrent += gain
		if st.MindCurrent > st.MindMax {
			st.MindCurrent = st.MindMax
		}
		if err := h.db.UpdateCombatState(st); err != nil {
			continue
		}
		h.pushHAMUpdate(id)
	}
}

// endPerformerOnCombat ends a performer's session when they take combat action
// (GDD 23.2.1: performers cannot fight while performing). Meditation ends too.
func (h *WorldHandler) endPerformerOnCombat(characterID string) {
	svcMu.Lock()
	_, isPerforming := performing[characterID]
	if _, med := meditating[characterID]; med {
		delete(meditating, characterID)
	}
	svcMu.Unlock()
	if isPerforming {
		h.endPerformance(characterID)
	}
}
