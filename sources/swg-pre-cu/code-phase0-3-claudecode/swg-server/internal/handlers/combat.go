package handlers

import (
	"encoding/json"
	"log"
	"math/rand"
	"time"

	"swg-server/internal/combat"
	"swg-server/internal/creatures"
	"swg-server/internal/protocol"
	"swg-server/internal/world"
)

// Unarmed baseline weapon. GDD Section 8.2.2 (Brawler) grants "unlock
// unarmed combat abilities"; no crafted-weapon system exists until Phase 4
// (Section 10.9.2), so every character fights with their fists for now.
// [ASSUMPTION] Exact unarmed base stats aren't given anywhere in the GDD
// (Section 9.7.1 gives *crafted weapon* damage bands only); picked to sit
// clearly below a starting crafted weapon's 50-150 dmg/hit range, since
// fists should be worse than even the cheapest crafted weapon.
var unarmedWeapon = combat.Weapon{
	Type: combat.WeaponMelee, BaseAccuracy: 30,
	MinDamage: 5, MaxDamage: 15, DamageType: "kinetic",
}

// AttackRangeM: [ASSUMPTION] not given an exact figure in the GDD; melee
// range, generously sized for testability over strict realism at this
// stage.
const AttackRangeM = 8.0

// ReviveRangeM: [ASSUMPTION] same basis as AttackRangeM.
const ReviveRangeM = 5.0

var combatRNG = rand.New(rand.NewSource(time.Now().UnixNano()))

// requiredCombatSkillBoxes: GDD Section 30 Phase 3 exit criteria requires
// "a skill-gated ability" -- combat_action is gated on owning at least one
// of the two Basic combat professions' Novice box (GDD 8.2.2, 8.2.3).
var requiredCombatSkillBoxes = []string{"marksman_novice", "brawler_novice"}

func (h *WorldHandler) hasCombatSkill(characterID string) bool {
	for _, boxID := range requiredCombatSkillBoxes {
		ok, err := h.db.HasSkillBox(characterID, boxID)
		if err == nil && ok {
			return true
		}
	}
	return false
}

// --- Client message handlers ---

func (h *WorldHandler) handleCombatAction(client *Client, raw []byte) {
	if client.CharacterID == "" {
		h.sendError(client, "not in world")
		return
	}
	if client.Combat.IncapacitatedAt != nil {
		h.sendError(client, "incapacitated, cannot act")
		return
	}
	if !h.hasCombatSkill(client.CharacterID) {
		h.sendError(client, "requires Novice Marksman or Novice Brawler")
		return
	}

	var msg protocol.WSMessage
	json.Unmarshal(raw, &msg)
	dataBytes, _ := json.Marshal(msg.Data)
	var action protocol.CombatActionMsg
	if err := json.Unmarshal(dataBytes, &action); err != nil {
		h.sendError(client, "invalid combat_action data")
		return
	}

	h.mu.Lock()
	target, ok := h.creatures[action.TargetID]
	h.mu.Unlock()
	if !ok || target.State == "dead" {
		h.sendError(client, "target not found")
		return
	}

	dist := creatures.DistanceSq(client.Pos.X, client.Pos.Z, target.PosX, target.PosZ)
	if dist > AttackRangeM*AttackRangeM {
		h.sendError(client, "target out of range")
		return
	}

	tmpl := creatures.TemplateByID(target.TemplateID)
	if tmpl == nil {
		h.sendError(client, "unknown creature template")
		return
	}

	attacker := combat.Attacker{
		AccuracySkill: 5, DamageSkillBonus: 2, // flat skill-gate bonus, see combat.go doc note
		Posture: client.Combat.Posture, Stance: client.Combat.Stance,
	}
	defender := combat.Defender{
		DefenseSkill: 0, Posture: combat.PostureStanding,
		ArmorMitigationPct: 0, ResistanceMultiplier: 1.0,
	}
	result := combat.ResolveAttack(attacker, defender, unarmedWeapon, combatRNG)

	h.broadcastCombatResult(client.CharacterID, action.TargetID, result)

	if !result.Hit {
		return
	}

	h.mu.Lock()
	target.HealthCurrent -= result.Damage
	died := target.HealthCurrent <= 0
	if died {
		target.HealthCurrent = 0
		target.State = "dead"
		target.TargetCharacterID = ""
	} else if target.State == "idle" {
		target.State = "aggro"
		target.TargetCharacterID = client.CharacterID
	}
	h.db.UpdateCreatureInstance(target)
	h.mu.Unlock()

	if died {
		h.grid.RemoveEntity(target.ID, world.Position{Planet: target.Planet, X: target.PosX, Y: target.PosY, Z: target.PosZ})
		h.broadcastToInterested(client, protocol.MsgEntityDespawn, protocol.EntityDespawnMsg{EntityID: target.ID})
		h.send(client, protocol.MsgCreatureDeath, protocol.CreatureDeathMsg{EntityID: target.ID})
		// Combat XP reward (GDD 7.2.1: "Killing creatures (scales with
		// creature level)"). [ASSUMPTION] flat reward per CL tier rather
		// than the full formula in 7.2.1 (which needs a player "effective
		// level" concept not yet built) -- scoped simplification for Phase 3.
		h.db.AddCharacterXP(client.CharacterID, "combat", 50*tmpl.CLMax)
	}
}

func (h *WorldHandler) handleSetPosture(client *Client, raw []byte) {
	if client.CharacterID == "" {
		h.sendError(client, "not in world")
		return
	}
	var msg protocol.WSMessage
	json.Unmarshal(raw, &msg)
	dataBytes, _ := json.Marshal(msg.Data)
	var p protocol.SetPostureMsg
	if err := json.Unmarshal(dataBytes, &p); err != nil {
		h.sendError(client, "invalid set_posture data")
		return
	}
	switch combat.Posture(p.Posture) {
	case combat.PostureStanding, combat.PostureKneeling, combat.PostureProne, combat.PostureCrouched:
		client.Combat.Posture = combat.Posture(p.Posture)
	default:
		h.sendError(client, "invalid posture")
		return
	}
	h.db.UpdateCombatState(client.CharacterID, client.Combat)
	h.sendHAMUpdate(client.CharacterID, client.Combat)
}

func (h *WorldHandler) handleSetStance(client *Client, raw []byte) {
	if client.CharacterID == "" {
		h.sendError(client, "not in world")
		return
	}
	var msg protocol.WSMessage
	json.Unmarshal(raw, &msg)
	dataBytes, _ := json.Marshal(msg.Data)
	var s protocol.SetStanceMsg
	if err := json.Unmarshal(dataBytes, &s); err != nil {
		h.sendError(client, "invalid set_stance data")
		return
	}
	switch combat.Stance(s.Stance) {
	case combat.StanceNormal, combat.StanceAggressive, combat.StanceDefensive, combat.StanceBerserk:
		client.Combat.Stance = combat.Stance(s.Stance)
	default:
		h.sendError(client, "invalid stance")
		return
	}
	h.db.UpdateCombatState(client.CharacterID, client.Combat)
	h.sendHAMUpdate(client.CharacterID, client.Combat)
}

func (h *WorldHandler) handleClone(client *Client) {
	if client.CharacterID == "" {
		h.sendError(client, "not in world")
		return
	}
	if client.Combat.IncapacitatedAt == nil {
		h.sendError(client, "not incapacitated")
		return
	}

	client.Combat = combat.Clone(client.Combat, time.Now())

	planet, x, y, z, err := h.db.GetCloneBindPoint(client.CharacterID)
	if err != nil {
		h.sendError(client, "no clone facility bound")
		return
	}

	oldPos := client.Pos
	client.Pos = world.Position{Planet: planet, X: x, Y: y, Z: z}
	client.LastMoveTime = time.Now()
	h.grid.MoveEntity(client.CharacterID, oldPos, client.Pos, client.Heading)

	h.db.UpdateCombatState(client.CharacterID, client.Combat)
	h.db.UpdateCharacterPosition(client.CharacterID, x, y, z, client.Heading, planet)

	h.send(client, protocol.MsgCloned, protocol.ClonedMsg{
		CharacterID: client.CharacterID,
		Position:    protocol.Position{X: x, Y: y, Z: z},
		Planet:      planet,
	})
	h.sendHAMUpdate(client.CharacterID, client.Combat)
	h.updateVisibility(client)
	log.Printf("Player %s cloned at %s (%.1f, %.1f, %.1f)", client.Name, planet, x, y, z)
}

func (h *WorldHandler) handleRevive(client *Client, raw []byte) {
	if client.CharacterID == "" {
		h.sendError(client, "not in world")
		return
	}
	var msg protocol.WSMessage
	json.Unmarshal(raw, &msg)
	dataBytes, _ := json.Marshal(msg.Data)
	var r protocol.ReviveMsg
	if err := json.Unmarshal(dataBytes, &r); err != nil {
		h.sendError(client, "invalid revive data")
		return
	}

	h.mu.RLock()
	target, ok := h.clients[r.TargetID]
	h.mu.RUnlock()
	if !ok {
		h.sendError(client, "target not found")
		return
	}
	if target.Combat.IncapacitatedAt == nil {
		h.sendError(client, "target is not incapacitated")
		return
	}
	if world.Distance2D(client.Pos, target.Pos) > ReviveRangeM {
		h.sendError(client, "target out of range")
		return
	}

	// [ASSUMPTION, Phase 3 scope note] GDD 9.6.1 gates revive behind Combat
	// Medic (an elite profession not yet built -- Phase 9). Any nearby
	// player can revive for now; add a Combat Medic skill-box check here
	// once that profession exists, the same way handleCombatAction gates
	// on Marksman/Brawler today.
	target.Combat = combat.Revive(target.Combat, time.Now())
	h.db.UpdateCombatState(target.CharacterID, target.Combat)
	h.sendHAMUpdate(target.CharacterID, target.Combat)
	log.Printf("%s revived %s", client.Name, target.Name)
}

// --- Broadcast helpers ---

func (h *WorldHandler) broadcastCombatResult(attackerID, targetID string, result combat.AttackResult) {
	h.mu.RLock()
	attacker, ok := h.clients[attackerID]
	h.mu.RUnlock()
	if !ok {
		return
	}
	msg := protocol.CombatResultMsg{
		AttackerID: attackerID, TargetID: targetID,
		Hit: result.Hit, Damage: result.Damage, HitChance: result.HitChance,
	}
	h.send(attacker, protocol.MsgCombatResult, msg)
	h.broadcastToInterested(attacker, protocol.MsgCombatResult, msg)
}

func (h *WorldHandler) sendHAMUpdate(characterID string, s combat.HAMPoolState) {
	h.mu.RLock()
	client, ok := h.clients[characterID]
	h.mu.RUnlock()
	if !ok {
		return
	}
	h.send(client, protocol.MsgHAMUpdate, protocol.HAMUpdateMsg{
		CharacterID:      characterID,
		HealthCurrent:    s.HealthCurrent,
		HealthMax:        s.HealthMax,
		ActionCurrent:    s.ActionCurrent,
		ActionMax:        s.ActionMax,
		MindCurrent:      s.MindCurrent,
		MindMax:          s.MindMax,
		WoundsHealth:     s.WoundsHealth,
		BattleFatiguePct: s.BattleFatiguePct,
		Posture:          string(s.Posture),
		Stance:           string(s.Stance),
		Incapacitated:    s.IncapacitatedAt != nil,
	})
}

// --- World tick: creature AI, incapacitation timeout, HAM regen ---

// StartTickLoop runs the Phase 3 world simulation tick. GDD Section
// 29.2.3's server-architecture table specifies HAM regeneration at a
// ~1-5 second cadence; 1 second is used here for both regen and AI so a
// single ticker drives both (regen is additionally rate-limited internally
// to every 5th tick, matching the table's upper bound, rather than firing
// every second).
func (h *WorldHandler) StartTickLoop() {
	ticker := time.NewTicker(1 * time.Second)
	tickCount := 0
	go func() {
		for range ticker.C {
			tickCount++
			h.tickCreatureAI()
			h.tickIncapacitationTimeouts()
			if tickCount%5 == 0 {
				h.tickHAMRegen()
			}
		}
	}()
}

// tickCreatureAI: MVP retaliation AI. A creature with no target scans for
// a nearby player within its aggro radius and targets them (all AI classes
// treated as aggressive-on-proximity for this pass -- GDD 17.2.1's
// passive/scavenger/humanoid distinctions are richer behavior flagged as a
// later refinement, not required by Phase 3's exit criteria). A creature
// with a target attacks it each tick if still in range and alive.
func (h *WorldHandler) tickCreatureAI() {
	h.mu.Lock()
	defer h.mu.Unlock()

	for _, c := range h.creatures {
		if c.State == "dead" {
			continue
		}
		tmpl := creatures.TemplateByID(c.TemplateID)
		if tmpl == nil {
			continue
		}

		if c.TargetCharacterID == "" {
			nearby := h.grid.Nearby(world.Position{Planet: c.Planet, X: c.PosX, Y: c.PosY, Z: c.PosZ}, tmpl.AggroRadiusM, "")
			for _, e := range nearby {
				if client, ok := h.clients[e.CharacterID]; ok && client.Combat.IncapacitatedAt == nil {
					c.TargetCharacterID = e.CharacterID
					c.State = "aggro"
					break
				}
			}
			continue
		}

		target, ok := h.clients[c.TargetCharacterID]
		distSq := 0.0
		if ok {
			distSq = creatures.DistanceSq(c.PosX, c.PosZ, target.Pos.X, target.Pos.Z)
		}
		// A stationary creature (no chase/pursuit AI in Phase 3 -- flagged
		// simplification, see AllTemplates doc comment) can only fight a
		// target that's actually within its reach right now. If the
		// target is gone, incapacitated, or has moved out of range, drop
		// the target -- there's no "give up after a while" case to model
		// separately, since a creature that can't move has nothing to
		// give up on: in-range-or-not is the whole story.
		if !ok || target.Combat.IncapacitatedAt != nil || distSq > tmpl.AggroRadiusM*tmpl.AggroRadiusM {
			c.TargetCharacterID = ""
			c.State = "idle"
			continue
		}

		attacker := combat.Attacker{AccuracySkill: tmpl.AccuracyBase, Stance: combat.StanceNormal, Posture: combat.PostureStanding}
		defender := combat.Defender{DefenseSkill: 0, Posture: target.Combat.Posture, ArmorMitigationPct: 0, ResistanceMultiplier: 1.0}
		weapon := combat.Weapon{Type: combat.WeaponMelee, BaseAccuracy: 20, MinDamage: tmpl.DamageMin, MaxDamage: tmpl.DamageMax}
		result := combat.ResolveAttack(attacker, defender, weapon, combatRNG)

		h.send(target, protocol.MsgCombatResult, protocol.CombatResultMsg{
			AttackerID: c.ID, TargetID: c.TargetCharacterID,
			Hit: result.Hit, Damage: result.Damage, HitChance: result.HitChance,
		})

		if result.Hit {
			newState, justIncap := combat.ApplyDamage(target.Combat, result.Damage, time.Now())
			target.Combat = newState
			h.db.UpdateCombatState(target.CharacterID, target.Combat)
			h.sendHAMUpdateLocked(target, target.Combat)
			if justIncap {
				h.send(target, protocol.MsgIncapacitated, protocol.IncapacitatedMsg{CharacterID: target.CharacterID})
			}
		}
	}
}

// sendHAMUpdateLocked is sendHAMUpdate's counterpart for call sites that
// already hold h.mu (avoids a self-deadlock on the RWMutex).
func (h *WorldHandler) sendHAMUpdateLocked(client *Client, s combat.HAMPoolState) {
	h.send(client, protocol.MsgHAMUpdate, protocol.HAMUpdateMsg{
		CharacterID:      client.CharacterID,
		HealthCurrent:    s.HealthCurrent,
		HealthMax:        s.HealthMax,
		ActionCurrent:    s.ActionCurrent,
		ActionMax:        s.ActionMax,
		MindCurrent:      s.MindCurrent,
		MindMax:          s.MindMax,
		WoundsHealth:     s.WoundsHealth,
		BattleFatiguePct: s.BattleFatiguePct,
		Posture:          string(s.Posture),
		Stance:           string(s.Stance),
		Incapacitated:    s.IncapacitatedAt != nil,
	})
}

// tickIncapacitationTimeouts: GDD 9.6.2 "Incapacitation timer expires
// without revive" -> death -> auto-clone at bound facility.
func (h *WorldHandler) tickIncapacitationTimeouts() {
	h.mu.RLock()
	var expired []*Client
	now := time.Now()
	for _, c := range h.clients {
		if c.Combat.IncapacitatedAt != nil && combat.IsDead(c.Combat, now) {
			expired = append(expired, c)
		}
	}
	h.mu.RUnlock()

	for _, client := range expired {
		h.handleClone(client)
	}
}

// tickHAMRegen: GDD 29.2.3 "HAM regeneration (out of combat) ~1-5 second
// tick". [ASSUMPTION] no exact regen rate is given in the GDD; 2% of
// effective max per interval is used, applied only when the player hasn't
// taken or dealt damage in the last 10 seconds ("out of combat").
func (h *WorldHandler) tickHAMRegen() {
	h.mu.RLock()
	defer h.mu.RUnlock()
	now := time.Now()
	for _, client := range h.clients {
		if client.Combat.IncapacitatedAt != nil {
			continue
		}
		if client.Combat.LastCombatActionAt != nil && now.Sub(*client.Combat.LastCombatActionAt) < 10*time.Second {
			continue
		}
		effMax := combat.EffectiveMax(client.Combat.HealthMax, client.Combat.WoundsHealth, client.Combat.BattleFatiguePct)
		if client.Combat.HealthCurrent >= effMax {
			continue
		}
		client.Combat.HealthCurrent += effMax / 50 // 2%
		if client.Combat.HealthCurrent > effMax {
			client.Combat.HealthCurrent = effMax
		}
		h.db.UpdateCombatState(client.CharacterID, client.Combat)
		h.sendHAMUpdateLocked(client, client.Combat)
	}
}

// --- World seeding ---

// SeedCombatWorld creates starter lairs and creature instances near the
// Tatooine spawn point (GDD 9.5.3, 17.2.3), unless lairs already exist
// from a previous run (idempotent across restarts).
func (h *WorldHandler) SeedCombatWorld() error {
	existing, err := h.db.GetActiveLairs()
	if err != nil {
		return err
	}
	if len(existing) > 0 {
		return h.loadCreaturesFromDB()
	}

	seedLairs := []struct {
		templateID     string
		x, z           float64
		lairHP, maxPop int
	}{
		{"womp_rat", 3560, -4790, 200, 3},
		{"worrt", 3560, -4798, 350, 2}, // ~8m from the womp rat lair, not 74m -- both clusters need to sit within a fresh character's reach of each other, not scattered across the map
	}

	for _, sl := range seedLairs {
		lair := &creatures.Lair{
			ID: "lair_" + sl.templateID, TemplateID: sl.templateID,
			Planet: "tatooine", PosX: sl.x, PosZ: sl.z,
			LairHP: sl.lairHP, LairHPMax: sl.lairHP, MaxPopulation: sl.maxPop,
		}
		if err := h.db.CreateLair(lair); err != nil {
			return err
		}

		tmpl := creatures.TemplateByID(sl.templateID)
		for i := 0; i < sl.maxPop; i++ {
			inst := &creatures.Instance{
				ID: lair.ID + "_c" + itoa(i), TemplateID: sl.templateID, LairID: lair.ID,
				Planet:        "tatooine",
				PosX:          sl.x + float64(i)*3, PosY: 5, PosZ: sl.z + float64(i)*2,
				HealthCurrent: tmpl.HealthMax, HealthMax: tmpl.HealthMax,
				State: "idle",
			}
			if err := h.db.CreateCreatureInstance(inst); err != nil {
				return err
			}
		}
	}
	return h.loadCreaturesFromDB()
}

func (h *WorldHandler) loadCreaturesFromDB() error {
	instances, err := h.db.GetLivingCreatureInstances()
	if err != nil {
		return err
	}
	h.mu.Lock()
	defer h.mu.Unlock()
	for i := range instances {
		inst := instances[i]
		h.creatures[inst.ID] = &inst
		tmpl := creatures.TemplateByID(inst.TemplateID)
		name := inst.TemplateID
		if tmpl != nil {
			name = tmpl.Name
		}
		h.grid.AddEntity(&world.Entity{
			CharacterID: inst.ID, // grid's field name predates creatures sharing it; holds the instance ID here
			Name:        name,
			Species:     "", // unused for creatures; EntityType/TemplateID on the wire carry the distinction
			Pos:         world.Position{Planet: inst.Planet, X: inst.PosX, Y: inst.PosY, Z: inst.PosZ},
		})
	}
	log.Printf("Loaded %d creature instances into the world", len(instances))
	return nil
}

func itoa(i int) string {
	if i == 0 {
		return "0"
	}
	digits := "0123456789"
	var b []byte
	for i > 0 {
		b = append([]byte{digits[i%10]}, b...)
		i /= 10
	}
	return string(b)
}
