// Combat handlers for Phase 3 — Combat Core. Testbed fork; generic content only.
// Adapts Group-C verified patterns (skill gate, ranges, tick loop, retaliation AI)
// to this fork's owner-directed numbers: fandom posture/0.5% accuracy model
// (internal/combat), uniform-1000 HAM (C2), generic creatures at (10,0)/(11,0)
// (C3), Group-C provisionals elsewhere (marked NON-CANONICAL).
package handlers

import (
	"encoding/json"
	"log"
	"math/rand"
	"time"

	"swg-server/internal/combat"
	"swg-server/internal/creatures"
	"swg-server/internal/database"
	"swg-server/internal/protocol"
	"swg-server/internal/services"
	"swg-server/internal/world"
)

// Unarmed baseline weapon. No crafted-weapon system exists until Phase 4, so every
// character fights unarmed for now.
// BaseAccuracy 30 / 5-15 damage / +2 skill bonus: [PROVISIONAL] Group-C convention
// (the fandom accuracy page gives no unarmed values). NON-CANONICAL placeholder.
var unarmedWeapon = combat.Weapon{
	Name: "unarmed", Melee: true, BaseAccuracy: 30,
	MinDamage: 5, MaxDamage: 15, DamageType: "kinetic",
}

// UnarmedDamageBonus: [PROVISIONAL] flat skill-gate bonus (Group-C convention).
const UnarmedDamageBonus = 2

// AttackRangeM: [PROVISIONAL] melee range gate (Group-C convention). The fandom
// ideal-range interpolation model needs per-weapon range tables that do not exist;
// a flat gate is the honest MVP. NON-CANONICAL.
const AttackRangeM = 8.0

// ReviveRangeM: [PROVISIONAL] same basis as AttackRangeM.
const ReviveRangeM = 5.0

// TickInterval drives creature AI, incap timeouts, and (every 5th tick) HAM regen.
const TickInterval = 1 * time.Second

// requiredCombatSkillBoxes: exit criteria require "a skill-gated ability" —
// combat_action is gated on a Novice combat-profession box.
var requiredCombatSkillBoxes = []string{"marksman_novice", "brawler_novice"}

// LiveCreature is the in-memory runtime mirror of a live creature instance.
type LiveCreature struct {
	ID         string
	TemplateID string
	Name       string
	Zone       string
	PosX, PosZ float64
}

// liveCreatures maps instance ID → runtime creature for connected-world AI.
// Guarded by h.mu (shared with the clients map).
func (h *WorldHandler) getLiveCreature(id string) (*LiveCreature, bool) {
	h.mu.RLock()
	defer h.mu.RUnlock()
	c, ok := h.liveCreatures[id]
	return c, ok
}

// SeedCombatWorld prepares the combat world: schema, starter lairs/instances,
// in-memory mirrors, and grid entities. Idempotent (safe on restart).
func (h *WorldHandler) SeedCombatWorld() error {
	if err := h.db.EnsureCombatSchema(); err != nil {
		return err
	}
	count, err := h.db.SeedLairs()
	if err != nil {
		return err
	}
	insts, err := h.db.GetLivingCreatureInstances()
	if err != nil {
		return err
	}
	h.mu.Lock()
	if h.liveCreatures == nil {
		h.liveCreatures = make(map[string]*LiveCreature)
	}
	for _, in := range insts {
		tmpl := creatures.TemplateByID(in.TemplateID)
		name := in.TemplateID
		if tmpl != nil {
			name = tmpl.Name
		}
		h.liveCreatures[in.ID] = &LiveCreature{
			ID: in.ID, TemplateID: in.TemplateID, Name: name, Zone: in.Zone,
			PosX: in.PosX, PosZ: in.PosZ,
		}
		h.grid.AddEntity(&world.Entity{
			CharacterID: in.ID,
			Name:        name,
			Species:     "creature",
			Pos:         world.Position{Planet: in.Zone, X: in.PosX, Y: 5.0, Z: in.PosZ},
		})
	}
	h.mu.Unlock()
	log.Printf("Combat world seeded: %d live creature instances", count)
	return nil
}

// StartTickLoop runs the world simulation tick: creature AI, incap timeouts,
// HAM regen (every 5th tick), and Phase 6 service ticks (drain every tick,
// BF healing + buff expiry every 5th).
func (h *WorldHandler) StartTickLoop() {
	ticker := time.NewTicker(TickInterval)
	var tickCount int64
	go func() {
		for range ticker.C {
			tickCount++
			h.tickCreatureAI()
			h.tickIncapacitationTimeouts()
			h.tickServices(tickCount)
			if tickCount%combat.RegenTickModulo == 0 {
				h.tickHAMRegen(tickCount)
			}
		}
	}()
}

// --- Client message handlers ---

func (h *WorldHandler) handleCombatAction(client *Client, raw []byte) {
	if client.CharacterID == "" {
		h.sendError(client, "not in world")
		return
	}
	st, err := h.db.GetCombatState(client.CharacterID)
	if err != nil {
		h.sendError(client, "combat state unavailable")
		return
	}
	if st.IncapacitatedAt != nil {
		h.sendError(client, "incapacitated, cannot act")
		return
	}
	h.endPerformerOnCombat(client.CharacterID) // GDD 23.2.1: no fighting while performing
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

	target, ok := h.getLiveCreature(action.TargetID)
	if !ok {
		// Phase 8: fall through to player targets, then faction bases.
		// Skill/incap/performer gates above already apply uniformly.
		if victim, vok := h.findClient(action.TargetID); vok {
			if deny := h.pvpDeny(client, victim); deny != "" {
				h.sendError(client, deny)
				return
			}
			h.flagAttackerOvert(client)
			h.handlePvPAttack(client, victim)
			return
		}
		if base, berr := h.db.GetBase(action.TargetID); berr == nil {
			if deny := h.baseDeny(client, base); deny != "" {
				h.sendError(client, deny)
				return
			}
			h.flagAttackerOvert(client)
			h.handleBaseAttack(client, base)
			return
		}
		log.Printf("COMBAT-DEBUG char=%s target %s not in live map (%d live)",
			client.CharacterID, action.TargetID, len(h.liveCreatures))
		h.sendError(client, "target not found")
		return
	}
	tmpl := creatures.TemplateByID(target.TemplateID)
	if tmpl == nil {
		h.sendError(client, "unknown creature template")
		return
	}
	if creatures.DistanceSq(client.Pos.X, client.Pos.Z, target.PosX, target.PosZ) > AttackRangeM*AttackRangeM {
		log.Printf("COMBAT-DEBUG char=%s at (%.1f,%.1f) target %s at (%.1f,%.1f): out of range",
			client.CharacterID, client.Pos.X, client.Pos.Z, target.ID, target.PosX, target.PosZ)
		h.sendError(client, "target out of range")
		return
	}

	// Equipped-weapon link (Phase 4 minimal-delta convention, proposal §6.8):
	// a crafted sidearm replaces the unarmed damage range and adds its accuracy
	// maximum as a BaseAccuracy bonus. [PROVISIONAL mapping, flagged.]
	weapon := unarmedWeapon
	if wItem, _, err := h.db.EquippedItems(client.CharacterID); err == nil && wItem != nil {
		if dmin, ok := wItem.Stats["damage_min"]; ok {
			weapon.MinDamage = int(dmin)
		}
		if dmax, ok := wItem.Stats["damage_max"]; ok {
			weapon.MaxDamage = int(dmax)
		}
		if acc, ok := wItem.Stats["accuracy_max"]; ok {
			weapon.BaseAccuracy = unarmedWeapon.BaseAccuracy + int(acc)
		}
		weapon.Name = "crafted:" + wItem.ID
	}
	// Unarmed proficiency is folded into the weapon row (BaseAccuracy 30,
	// provisional) per the resolution signature; AccuracySkill carries
	// stance/posture-independent skill mods (none in MVP).
	attacker := combat.Attacker{
		AccuracySkill: 0, DamageBonus: UnarmedDamageBonus,
		Posture: combat.Posture(st.Posture), Stance: combat.Stance(st.Stance),
	}
	defender := combat.Defender{
		DefenseSkill: 0, Posture: combat.PostureStanding,
		ArmorMitigationPct: 0, ResistanceMultiplier: 1.0,
	}
	// Resolve with a fresh RNG per call: the old shared *rand.Rand was a data
	// race across the world tick, resource tick, and request handlers (a race
	// here once panicked a tick loop mid-run). Time-seeded per call preserves
	// existing behavior without shared state.
	result := combat.ResolveAttack(attacker, defender, weapon, rand.New(rand.NewSource(time.Now().UnixNano())))

	h.broadcastToInterested(client, protocol.MsgCombatResult, protocol.CombatResultMsg{
		AttackerID: client.CharacterID, TargetID: target.ID,
		Hit: result.Hit, Damage: result.Damage, HitChance: result.HitChance,
	})
	// The attacker always sees their own result (interest broadcast excludes
	// the sender by design — nearby observers get it above).
	h.send(client, protocol.MsgCombatResult, protocol.CombatResultMsg{
		AttackerID: client.CharacterID, TargetID: target.ID,
		Hit: result.Hit, Damage: result.Damage, HitChance: result.HitChance,
	})
	if !result.Hit {
		return
	}

	// Apply damage to the instance (DB is the authority; mirror follows).
	insts, err := h.db.GetLivingCreatureInstances()
	if err != nil {
		return
	}
	for _, in := range insts {
		if in.ID != target.ID {
			continue
		}
		in.HealthCurrent -= result.Damage
		if in.HealthCurrent <= 0 {
			if err := h.db.KillCreatureInstance(in.ID, time.Now().Add(creatures.CorpseWindow)); err != nil {
				return
			}
			h.mu.Lock()
			delete(h.liveCreatures, in.ID)
			h.mu.Unlock()
			h.grid.RemoveEntity(in.ID, world.Position{Planet: in.Zone, X: in.PosX, Y: in.PosY, Z: in.PosZ})
			h.broadcastToInterested(client, protocol.MsgCreatureDeath, protocol.CreatureDeathMsg{
				InstanceID: in.ID, TemplateID: in.TemplateID,
			})
			h.send(client, protocol.MsgCreatureDeath, protocol.CreatureDeathMsg{
				InstanceID: in.ID, TemplateID: in.TemplateID,
			})
			// Combat XP: [PROVISIONAL] flat per-CLMax reward (Group-C convention:
			// full 7.2.1 formula needs an effective-level concept not yet built).
			xpReward := 50 * tmpl.CLMax
			_ = h.db.AddCharacterXP(client.CharacterID, "combat", xpReward)
			// Phase 5 creature-drop faucet (GDD 12.2.2 band 10–200, CL-scaled):
			// 10 × CLMax, ledger-tagged faucet. [PROVISIONAL mapping.]
			drop := 10 * tmpl.CLMax
			if drop > 200 {
				drop = 200
			}
			_ = h.db.AddCredits(client.CharacterID, drop)
			_ = h.db.RecordLedger(client.CharacterID, drop, "creature_drop", "faucet",
				"creature:"+in.ID, client.Pos.Planet)
		} else {
			row := in
			if err := h.db.UpdateCreatureInstance(&row); err != nil {
				return
			}
			// Retaliation hook: a wounded idle creature acquires the attacker.
			if row.TargetCharacterID == "" {
				row.TargetCharacterID = client.CharacterID
				row.State = "aggro"
				_ = h.db.UpdateCreatureInstance(&row)
			}
		}
	}
}

func (h *WorldHandler) hasCombatSkill(characterID string) bool {
	for _, boxID := range requiredCombatSkillBoxes {
		if ok, err := h.db.HasSkillBox(characterID, boxID); err == nil && ok {
			return true
		}
	}
	return false
}

func (h *WorldHandler) handleSetPosture(client *Client, raw []byte) {
	if client.CharacterID == "" {
		h.sendError(client, "not in world")
		return
	}
	var msg protocol.WSMessage
	json.Unmarshal(raw, &msg)
	dataBytes, _ := json.Marshal(msg.Data)
	var pm protocol.PostureMsg
	if err := json.Unmarshal(dataBytes, &pm); err != nil {
		h.sendError(client, "invalid set_posture data")
		return
	}
	switch combat.Posture(pm.Posture) {
	case combat.PostureStanding, combat.PostureKneeling, combat.PostureProne, combat.PostureCrouched:
	default:
		h.sendError(client, "unknown posture")
		return
	}
	st, err := h.db.GetCombatState(client.CharacterID)
	if err != nil {
		h.sendError(client, "combat state unavailable")
		return
	}
	if st.IncapacitatedAt != nil {
		h.sendError(client, "incapacitated, cannot act")
		return
	}
	st.Posture = pm.Posture
	if err := h.db.UpdateCombatState(st); err != nil {
		h.sendError(client, "posture update failed")
		return
	}
	h.pushHAMUpdate(client.CharacterID)
}

func (h *WorldHandler) handleSetStance(client *Client, raw []byte) {
	if client.CharacterID == "" {
		h.sendError(client, "not in world")
		return
	}
	var msg protocol.WSMessage
	json.Unmarshal(raw, &msg)
	dataBytes, _ := json.Marshal(msg.Data)
	var sm protocol.StanceMsg
	if err := json.Unmarshal(dataBytes, &sm); err != nil {
		h.sendError(client, "invalid set_stance data")
		return
	}
	switch combat.Stance(sm.Stance) {
	case combat.StanceNormal, combat.StanceAggressive, combat.StanceDefensive, combat.StanceBerserk:
	default:
		h.sendError(client, "unknown stance")
		return
	}
	st, err := h.db.GetCombatState(client.CharacterID)
	if err != nil {
		h.sendError(client, "combat state unavailable")
		return
	}
	if st.IncapacitatedAt != nil {
		h.sendError(client, "incapacitated, cannot act")
		return
	}
	st.Stance = sm.Stance
	if err := h.db.UpdateCombatState(st); err != nil {
		h.sendError(client, "stance update failed")
		return
	}
	h.pushHAMUpdate(client.CharacterID)
}

func (h *WorldHandler) handleRevive(client *Client, raw []byte) {
	if client.CharacterID == "" {
		h.sendError(client, "not in world")
		return
	}
	var msg protocol.WSMessage
	json.Unmarshal(raw, &msg)
	dataBytes, _ := json.Marshal(msg.Data)
	var rm protocol.ReviveMsg
	if err := json.Unmarshal(dataBytes, &rm); err != nil {
		h.sendError(client, "invalid revive data")
		return
	}
	// Phase 6 gate (proposal S2): revive requires Novice Medic, closing the
	// Phase 3 any-nearby-player interim. Base-Medic per GDD 8.2.5's own benefits
	// list; Combat Medic elite gating deferred with elite professions.
	target, ok := h.findClient(rm.TargetID)
	if !ok {
		h.sendError(client, "revive target not in world")
		return
	}
	if creatures.DistanceSq(client.Pos.X, client.Pos.Z, target.Pos.X, target.Pos.Z) > ReviveRangeM*ReviveRangeM {
		h.sendError(client, "revive target out of range")
		return
	}
	if ok, err := h.db.HasSkillBox(client.CharacterID, "medic_novice"); err != nil || !ok {
		h.sendError(client, "requires Novice Medic")
		return
	}
	tst, cur, err := h.loadCombatState(rm.TargetID)
	if err != nil || tst.IncapacitatedAt == nil {
		h.sendError(client, "target is not incapacitated")
		return
	}
	now := time.Now()
	revived := combat.Revive(cur, now)
	applyHAMPoolState(tst, revived)
	if err := h.db.UpdateCombatState(tst); err != nil {
		h.sendError(client, "revive failed")
		return
	}
	// Phase 8: surviving voids the PvP kill credit (kill requires death).
	_ = h.db.VoidPendingKill(rm.TargetID)
	_ = h.db.AddCharacterXP(client.CharacterID, "medical", services.ReviveXP)
	h.pushHAMUpdate(rm.TargetID)
}

func (h *WorldHandler) handleClone(client *Client, raw []byte) {
	if client.CharacterID == "" {
		h.sendError(client, "not in world")
		return
	}
	st, err := h.db.GetCombatState(client.CharacterID)
	if err != nil {
		h.sendError(client, "combat state unavailable")
		return
	}
	if st.IncapacitatedAt == nil {
		h.sendError(client, "clone available only while incapacitated")
		return
	}
	h.cloneCharacter(client, st, time.Now())
}

// cloneCharacter applies clone respawn: pools restored (+provisional wounds/BF),
// character moved to the bound point, grid + visibility refreshed.
func (h *WorldHandler) cloneCharacter(client *Client, st *database.CombatState, now time.Time) {
	cloned := combat.Clone(toHAMPoolState(st,
		h.db.BuffBonus(client.CharacterID, "health"),
		h.db.BuffBonus(client.CharacterID, "action"),
		h.db.BuffBonus(client.CharacterID, "mind")), now)
	applyHAMPoolState(st, cloned)
	if err := h.db.UpdateCombatState(st); err != nil {
		h.sendError(client, "clone failed")
		return
	}
	// Phase 8: death consumes a pending PvP kill into points + victim→killer
	// credit transfer + equipped condition loss (F3). Timer-expiry deaths
	// route through this same function, so both death kinds are covered.
	h.applyPvPDeathLegs(client)
	bind, err := h.db.GetCloneBinding(client.CharacterID)
	if err != nil {
		h.sendError(client, "clone binding unavailable")
		return
	}
	oldPos := client.Pos
	newPos := world.Position{Planet: bind.Zone, X: bind.X, Y: bind.Y, Z: bind.Z}
	client.Pos = newPos
	_ = h.db.UpdateCharacterPosition(client.CharacterID, bind.X, bind.Y, bind.Z, 0, bind.Zone)
	h.grid.MoveEntity(client.CharacterID, oldPos, newPos, 0)
	h.updateVisibility(client)
	h.updateVisibilityForOthers(client)
	h.send(client, protocol.MsgCloned, protocol.ClonedMsg{
		CharacterID: client.CharacterID,
		Position:    protocol.Position{X: bind.X, Y: bind.Y, Z: bind.Z},
		Zone:        bind.Zone,
	})
	h.pushHAMUpdate(client.CharacterID)
}

func (h *WorldHandler) findClient(characterID string) (*Client, bool) {
	h.mu.RLock()
	defer h.mu.RUnlock()
	c, ok := h.clients[characterID]
	return c, ok
}

// pushHAMUpdate broadcasts a character's current combat state to interested clients.
// HealthMax shown is the EFFECTIVE max (wounds/BF reductions + buff bonus), so
// buffs are observable in-band; currents were already exposed.
func (h *WorldHandler) pushHAMUpdate(characterID string) {
	_, cur, err := h.loadCombatState(characterID)
	if err != nil {
		return
	}
	incap := cur.IncapacitatedAt != nil
	msg := protocol.HAMUpdateMsg{
		CharacterID: characterID, HealthCurrent: cur.HealthCurrent,
		HealthMax: combat.EffectiveMax(cur.HealthMax, cur.WoundsHealth, cur.BattleFatiguePct, cur.BuffHealth),
		ActionCurrent: cur.ActionCurrent,
		MindCurrent: cur.MindCurrent, WoundsHealth: cur.WoundsHealth,
		BattleFatiguePct: cur.BattleFatiguePct, Posture: string(cur.Posture),
		Stance: string(cur.Stance), Incapacitated: incap,
	}
	h.mu.RLock()
	sender, ok := h.clients[characterID]
	h.mu.RUnlock()
	if !ok {
		return
	}
	h.send(sender, protocol.MsgHAMUpdate, msg)
	for _, e := range h.grid.Nearby(sender.Pos, world.InterestRadius, characterID) {
		h.mu.RLock()
		other, ok := h.clients[e.CharacterID]
		h.mu.RUnlock()
		if ok {
			h.send(other, protocol.MsgHAMUpdate, msg)
		}
	}
}

// --- World tick ---

// tickCreatureAI: MVP retaliation AI. Idle creatures scan for connected players
// within aggro radius; aggro'd creatures attack each tick while the target stays
// in attack range, and disengage out of range. Creatures do not move (later scope).
func (h *WorldHandler) tickCreatureAI() {
	h.mu.RLock()
	insts := make([]*LiveCreature, 0, len(h.liveCreatures))
	for _, c := range h.liveCreatures {
		insts = append(insts, &LiveCreature{ID: c.ID, TemplateID: c.TemplateID, Name: c.Name, Zone: c.Zone, PosX: c.PosX, PosZ: c.PosZ})
	}
	clients := make([]*Client, 0, len(h.clients))
	for _, c := range h.clients {
		if c.CharacterID != "" {
			clients = append(clients, c)
		}
	}
	h.mu.RUnlock()

	rows, err := h.db.GetLivingCreatureInstances()
	if err != nil {
		return
	}
	byID := make(map[string]*database.CreatureInstanceRow)
	for i := range rows {
		byID[rows[i].ID] = &rows[i]
	}

	for _, c := range insts {
		tmpl := creatures.TemplateByID(c.TemplateID)
		if tmpl == nil {
			continue
		}
		row, ok := byID[c.ID]
		if !ok {
			continue
		}
		// Acquire: nearest connected player in aggro radius (same zone).
		if row.TargetCharacterID == "" {
			var best *Client
			bestD := tmpl.AggroRadiusM * tmpl.AggroRadiusM
			for _, cl := range clients {
				if cl.Pos.Planet != c.Zone {
					continue
				}
				d := creatures.DistanceSq(cl.Pos.X, cl.Pos.Z, c.PosX, c.PosZ)
				if d <= bestD {
					bestD = d
					best = cl
				}
			}
			if best != nil {
				row.TargetCharacterID = best.CharacterID
				row.State = "aggro"
				_ = h.db.UpdateCreatureInstance(row)
			}
			continue
		}
		// Attack: target must still be connected, in zone, in attack range, alive.
		h.mu.RLock()
		tgt, ok := h.clients[row.TargetCharacterID]
		h.mu.RUnlock()
		if !ok || tgt.Pos.Planet != c.Zone ||
			creatures.DistanceSq(tgt.Pos.X, tgt.Pos.Z, c.PosX, c.PosZ) > AttackRangeM*AttackRangeM {
			row.TargetCharacterID = ""
			row.State = "idle"
			_ = h.db.UpdateCreatureInstance(row)
			continue
		}
		tst, err := h.db.GetCombatState(tgt.CharacterID)
		if err != nil || tst.IncapacitatedAt != nil {
			if err == nil {
				row.TargetCharacterID = ""
				row.State = "idle"
				_ = h.db.UpdateCreatureInstance(row)
			}
			continue
		}
		att := combat.Attacker{
			AccuracySkill: tmpl.AccuracyBase, DamageBonus: 0,
			Posture: combat.PostureStanding, Stance: combat.StanceNormal,
		}
		wpn := combat.Weapon{
			Name: "claws", Melee: true, BaseAccuracy: 0,
			MinDamage: tmpl.DamageMin, MaxDamage: tmpl.DamageMax,
			DamageType: "kinetic",
		}
		// Equipped-armor link (Phase 4 minimal-delta convention): crafted plate
		// protection maximum becomes mitigation %, capped at 50. [PROVISIONAL]
		mitigation := 0.0
		if _, armor, err := h.db.EquippedItems(tgt.CharacterID); err == nil && armor != nil {
			if prot, ok := armor.Stats["protection_max"]; ok {
				mitigation = prot
				if mitigation > 50 {
					mitigation = 50
				}
			}
		}
		def := combat.Defender{
			DefenseSkill: 0, Posture: combat.Posture(tst.Posture),
			ArmorMitigationPct: mitigation, ResistanceMultiplier: 1.0,
		}
		res := combat.ResolveAttack(att, def, wpn,
			rand.New(rand.NewSource(time.Now().UnixNano())))
		crMsg := protocol.CombatResultMsg{
			AttackerID: c.ID, TargetID: tgt.CharacterID,
			Hit: res.Hit, Damage: res.Damage, HitChance: res.HitChance,
		}
		h.broadcastToInterested(tgt, protocol.MsgCombatResult, crMsg)
		h.send(tgt, protocol.MsgCombatResult, crMsg) // victim always sees own hits
		if !res.Hit {
			continue
		}
		now := time.Now()
		_, cur, err := h.loadCombatState(tgt.CharacterID)
		if err != nil {
			continue
		}
		updated, justIncap := combat.ApplyDamage(cur, res.Damage, now)
		applyHAMPoolState(tst, updated)
		_ = h.db.UpdateCombatState(tst)
		h.pushHAMUpdate(tgt.CharacterID)
		if justIncap {
			h.send(tgt, protocol.MsgIncapacitated, protocol.IncapacitatedMsg{CharacterID: tgt.CharacterID})
		}
	}
}

// tickIncapacitationTimeouts auto-clones characters whose 5-minute timer expired.
func (h *WorldHandler) tickIncapacitationTimeouts() {
	now := time.Now()
	h.mu.RLock()
	clients := make([]*Client, 0, len(h.clients))
	for _, c := range h.clients {
		if c.CharacterID != "" {
			clients = append(clients, c)
		}
	}
	h.mu.RUnlock()
	for _, cl := range clients {
		st, cur, err := h.loadCombatState(cl.CharacterID)
		if err != nil || st.IncapacitatedAt == nil {
			continue
		}
		if combat.IsDead(cur, now) {
			h.cloneCharacter(cl, st, now)
		}
	}
}

// tickHAMRegen applies out-of-combat regen per the provisional convention.
func (h *WorldHandler) tickHAMRegen(tickCount int64) {
	h.mu.RLock()
	clients := make([]*Client, 0, len(h.clients))
	for _, c := range h.clients {
		if c.CharacterID != "" {
			clients = append(clients, c)
		}
	}
	h.mu.RUnlock()
	for _, cl := range clients {
		st, cur, err := h.loadCombatState(cl.CharacterID)
		if err != nil {
			continue
		}
		updated := combat.RegenTick(cur, tickCount)
		applyHAMPoolState(st, updated)
		_ = h.db.UpdateCombatState(st)
	}
}

// toHAMPoolState / applyHAMPoolState bridge the DB row and the pure-logic state.
// Buff bonuses are overlaid by the caller via loadCombatState (buffs live in
// their own table with wall-clock expiry, not in ham_pool_states).
func toHAMPoolState(s *database.CombatState, bh, ba, bm int) combat.HAMPoolState {
	return combat.HAMPoolState{
		HealthCurrent: s.HealthCurrent, HealthMax: s.HealthMax,
		ActionCurrent: s.ActionCurrent, ActionMax: s.ActionMax,
		MindCurrent: s.MindCurrent, MindMax: s.MindMax,
		WoundsHealth: s.WoundsHealth, WoundsAction: s.WoundsAction,
		WoundsMind: s.WoundsMind, BattleFatiguePct: s.BattleFatiguePct,
		BuffHealth: bh, BuffAction: ba, BuffMind: bm,
		Posture: combat.Posture(s.Posture), Stance: combat.Stance(s.Stance),
		IncapacitatedAt: s.IncapacitatedAt,
	}
}

// loadCombatState loads the persisted row plus active buff bonuses.
func (h *WorldHandler) loadCombatState(characterID string) (*database.CombatState, combat.HAMPoolState, error) {
	st, err := h.db.GetCombatState(characterID)
	if err != nil {
		return nil, combat.HAMPoolState{}, err
	}
	bh := h.db.BuffBonus(characterID, "health")
	ba := h.db.BuffBonus(characterID, "action")
	bm := h.db.BuffBonus(characterID, "mind")
	return st, toHAMPoolState(st, bh, ba, bm), nil
}

func applyHAMPoolState(s *database.CombatState, u combat.HAMPoolState) {
	s.HealthCurrent = u.HealthCurrent
	s.ActionCurrent = u.ActionCurrent
	s.MindCurrent = u.MindCurrent
	s.WoundsHealth = u.WoundsHealth
	s.WoundsAction = u.WoundsAction
	s.WoundsMind = u.WoundsMind
	s.BattleFatiguePct = u.BattleFatiguePct
	s.Posture = string(u.Posture)
	s.Stance = string(u.Stance)
	s.IncapacitatedAt = u.IncapacitatedAt
}

// creatureInfo reports template info for grid entities that are creatures.
func (h *WorldHandler) creatureInfo(id string) (templateID string, ok bool) {
	h.mu.RLock()
	defer h.mu.RUnlock()
	c, ok := h.liveCreatures[id]
	if !ok {
		return "", false
	}
	return c.TemplateID, true
}
