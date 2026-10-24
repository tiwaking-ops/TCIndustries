package handlers

import (
	"encoding/json"
	"sync"
	"time"

	"tc-server/internal/combat"
	"tc-server/internal/protocol"
	"tc-server/internal/services"
	"tc-server/internal/world"
)

type performSession struct {
	PerformerID string
	Kind        string
	PosX, PosZ  float64
	StartedAt   time.Time
}

var (
	svcMu      sync.Mutex
	performing = make(map[string]*performSession)
	watching   = make(map[string]string)
	healMu     sync.Mutex
	lastHeal   = make(map[string]time.Time)
)

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

func (h *WorldHandler) SeedServiceWorld() error {
	return h.db.EnsureServiceSchema()
}

func (h *WorldHandler) findClient(characterID string) (*Client, bool) {
	h.mu.RLock()
	defer h.mu.RUnlock()
	c, ok := h.clients[characterID]
	return c, ok
}

func (h *WorldHandler) pushHAMUpdate(characterID string) {
	cl, ok := h.findClient(characterID)
	if !ok {
		return
	}
	st, err := h.db.GetCombatState(characterID)
	if err != nil {
		return
	}
	cl.Combat = *st
	h.sendHAMUpdate(characterID, *st)
}

// --- S1: Wound healing ---

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
	if !h.hasBox(client.CharacterID, services.MedicNoviceBox) {
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
	tst.WoundsHealth -= amount
	if tst.WoundsHealth < 0 {
		tst.WoundsHealth = 0
	}
	if err := h.db.UpdateCombatState(hm.TargetID, *tst); err != nil {
		h.sendError(client, "heal failed")
		return
	}
	_ = h.db.AddCharacterXP(client.CharacterID, "medical", services.HealXP)
	h.pushHAMUpdate(hm.TargetID)
	h.send(client, protocol.MsgWoundHealed, protocol.WoundHealedMsg{
		TargetID: hm.TargetID, Pool: "health", Amount: amount, WoundsLeft: tst.WoundsHealth,
	})
}

// --- S3: Buffs ---

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
	if !services.ValidBuffPool(bm.Pool) {
		h.sendError(client, "unknown buff pool")
		return
	}
	if !h.hasBox(client.CharacterID, services.MedicNoviceBox) {
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
	expires := time.Now().Add(services.BuffDuration)
	if err := h.db.AddBuff(bm.TargetID, client.CharacterID, bm.Pool, amount, expires); err != nil {
		h.sendError(client, "buff already active")
		return
	}
	_ = h.db.AddCharacterXP(client.CharacterID, "medical", services.BuffXP)
	h.pushHAMUpdate(bm.TargetID)
	h.send(client, protocol.MsgBuffApplied, protocol.BuffAppliedMsg{
		TargetID: bm.TargetID, Pool: bm.Pool, Amount: amount, ExpiresAt: expires.Format(time.RFC3339),
	})
}

// --- S4: Perform / Watch / BF healing ---

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
	if !h.hasBox(client.CharacterID, services.EntertainerNoviceBox) {
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
	h.endPerformance(client.CharacterID)
	h.send(client, protocol.MsgPerformanceStopped, protocol.PerformMsg{})
}

func (h *WorldHandler) endPerformance(performerID string) {
	svcMu.Lock()
	delete(performing, performerID)
	for watcher, perf := range watching {
		if perf == performerID {
			delete(watching, watcher)
			if wc, ok := h.findClient(watcher); ok {
				h.send(wc, protocol.MsgWatchStopped, protocol.WatchMsg{PerformerID: performerID})
			}
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
	perfSess, ok := performing[wm.PerformerID]
	if !ok {
		svcMu.Unlock()
		h.sendError(client, "performer is not active")
		return
	}
	perfClient, ok := h.findClient(wm.PerformerID)
	if !ok {
		svcMu.Unlock()
		h.sendError(client, "performer not in world")
		return
	}
	if world.Distance2D(client.Pos, perfClient.Pos) > services.WatchRadiusM {
		svcMu.Unlock()
		h.sendError(client, "performer out of range")
		return
	}
	// One-watch-target rule (GDD 23.5)
	watching[client.CharacterID] = wm.PerformerID
	svcMu.Unlock()
	h.send(client, protocol.MsgWatchStarted, protocol.WatchMsg{PerformerID: wm.PerformerID})
	_ = perfSess
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

// --- S6: Tips ---

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
		h.sendError(client, "tip amount must be positive")
		return
	}
	svcMu.Lock()
	perfID, watchingPerf := watching[client.CharacterID]
	svcMu.Unlock()
	if !watchingPerf || perfID != tm.TargetID {
		h.sendError(client, "must be watching the performer to tip them")
		return
	}
	if err := h.db.DeductCredits(client.CharacterID, tm.Amount); err != nil {
		h.sendError(client, "insufficient credits for tip")
		return
	}
	if err := h.db.AddCredits(tm.TargetID, tm.Amount); err != nil {
		_ = h.db.AddCredits(client.CharacterID, tm.Amount) // refund
		h.sendError(client, "tip failed")
		return
	}
	_ = h.db.RecordLedger(client.CharacterID, -tm.Amount, "tip", "transfer", tm.TargetID, client.Pos.Planet)
	_ = h.db.RecordLedger(tm.TargetID, tm.Amount, "tip", "transfer", client.CharacterID, client.Pos.Planet)
	_ = h.db.AddCharacterXP(tm.TargetID, "entertainer", services.TipXP)
	h.send(client, protocol.MsgTipReceipt, protocol.TipMsg{TargetID: tm.TargetID, Amount: tm.Amount})
	if perf, ok := h.findClient(tm.TargetID); ok {
		h.send(perf, protocol.MsgTipReceived, protocol.TipMsg{TargetID: tm.TargetID, Amount: tm.Amount})
	}
}

// --- S5: Stim use ---

func (h *WorldHandler) handleUseStim(client *Client, raw []byte) {
	if client.CharacterID == "" {
		h.sendError(client, "not in world")
		return
	}
	// For testbed MVP simplicity, stim use heals wounds/health directly.
	var msg protocol.WSMessage
	json.Unmarshal(raw, &msg)
	dataBytes, _ := json.Marshal(msg.Data)
	var sm protocol.StimMsg
	if err := json.Unmarshal(dataBytes, &sm); err != nil {
		h.sendError(client, "invalid stim data")
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
	healing := 50
	tst.WoundsHealth -= healing
	if tst.WoundsHealth < 0 {
		tst.WoundsHealth = 0
	}
	tst.HealthCurrent += healing
	effMax := combat.EffectiveMax(tst.HealthMax, tst.WoundsHealth, tst.BattleFatiguePct, tst.BuffHealth)
	if tst.HealthCurrent > effMax {
		tst.HealthCurrent = effMax
	}
	if err := h.db.UpdateCombatState(sm.TargetID, *tst); err != nil {
		h.sendError(client, "stim failed")
		return
	}
	_ = h.db.AddCharacterXP(client.CharacterID, "medical", 25)
	h.pushHAMUpdate(sm.TargetID)
	h.send(client, protocol.MsgStimUsed, protocol.StimUsedMsg{
		TargetID: sm.TargetID, ChargesLeft: 5, Healing: healing,
	})
}

// --- Service tick ---

func (h *WorldHandler) tickServices(tickCount int64) {
	h.tickPerformerDrain()
	if tickCount%5 != 0 {
		return
	}
	h.db.ExpireBuffs()
	h.tickPerformanceHealing()
}

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
		if dx*dx+dz*dz > 4 { // 2m stationary rule
			h.endPerformance(s.PerformerID)
			continue
		}
		st.MindCurrent -= services.MindDrainPerTick
		if st.MindCurrent < 0 {
			st.MindCurrent = 0
		}
		if err := h.db.UpdateCombatState(s.PerformerID, *st); err != nil {
			continue
		}
		h.pushHAMUpdate(s.PerformerID)
	}
}

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
		_, performing := performing[perfID]
		stillWatching := watching[watcherID] == perfID
		svcMu.Unlock()
		if !performing || !stillWatching {
			continue
		}
		dx := watcher.Pos.X - perf.Pos.X
		dz := watcher.Pos.Z - perf.Pos.Z
		if dx*dx+dz*dz > services.WatchRadiusM*services.WatchRadiusM {
			continue
		}
		wst, err := h.db.GetCombatState(watcherID)
		if err != nil || wst.BattleFatiguePct <= 0 {
			continue
		}
		healingTier := h.skillTier(perfID, "entertainer_healing_")
		perfTier := 1
		rate := services.BFHealPerTick(healingTier, perfTier)
		wst.BattleFatiguePct -= rate
		if wst.BattleFatiguePct < 0 {
			wst.BattleFatiguePct = 0
		}
		if err := h.db.UpdateCombatState(watcherID, *wst); err != nil {
			continue
		}
		h.pushHAMUpdate(watcherID)
	}
}

func (h *WorldHandler) endPerformerOnCombat(characterID string) {
	svcMu.Lock()
	_, isPerforming := performing[characterID]
	svcMu.Unlock()
	if isPerforming {
		h.endPerformance(characterID)
	}
}
