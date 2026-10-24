// Phase 8 PvP attack resolution and the combat gate wired into
// handleCombatAction. PvE (creature) combat bypasses this file entirely.
//
// Design authority: proposals/proposal-claude-tcindustries-phase8-faction-pvp-v1.md
//   Section 5 (legality engine + side effects), R-06 (no extra penalty).
package handlers

import (
	"database/sql"
	"log"
	"time"

	"tc-server/internal/combat"
	"tc-server/internal/pvp"
	"tc-server/internal/protocol"
	"tc-server/internal/world"
)

// resolvePvPAttack decides whether attacker may damage target right now and,
// when allowed, applies the attack identically to creature retaliation:
// combat.ResolveAttack → combat.ApplyDamage → HAM update → incidental
// incapacitation. No PvP-specific penalties are ever applied (R-06).
//
// Returns the decision; the caller (handleCombatAction) sends the error on
// rejection so message framing stays in one place.
func (h *WorldHandler) resolvePvPAttack(attacker, target *Client) pvp.PvPDecision {
	if attacker == nil || target == nil {
		return pvp.PvPDecision{Reason: "target_not_found"}
	}

	// Battlefield membership from authoritative positions (D-14).
	zoneID := h.zoneAt(attacker.Pos.Planet, attacker.Pos.X, attacker.Pos.Z)
	tZoneID := h.zoneAt(target.Pos.Planet, target.Pos.X, target.Pos.Z)
	if zoneID == "" {
		zoneID = tZoneID
	}

	aCtx := h.buildPvPContext(attacker.CharacterID, zoneID)
	txCtx := h.buildPvPContext(target.CharacterID, tZoneID)
	aCtx.GuildID = h.guildIDFor(attacker.CharacterID)
	txCtx.GuildID = h.guildIDFor(target.CharacterID)

	warActive := func(g1, g2 string) bool { return h.db.AnyActiveWar(g1, g2) }
	decision := pvp.CanAttackPlayer(time.Now(), aCtx, txCtx, warActive)
	if !decision.Allowed {
		return decision
	}

	// Range gate: same on-foot reach as creature combat.
	dist := world.Distance2D(attacker.Pos, target.Pos)
	if dist > AttackRangeM {
		return pvp.PvPDecision{Reason: "target_out_of_range"}
	}

	attackerSide := combat.Attacker{
		AccuracySkill: 5, DamageSkillBonus: 2, // flat skill-gate bonus, as PvE
		Posture: attacker.Combat.Posture, Stance: attacker.Combat.Stance,
	}
	defender := combat.Defender{
		DefenseSkill: 0, Posture: target.Combat.Posture,
		ArmorMitigationPct: 0, ResistanceMultiplier: 1.0,
	}
	result := combat.ResolveAttack(attackerSide, defender, unarmedWeapon, combatRNG)

	h.send(attacker, protocol.MsgCombatResult, protocol.CombatResultMsg{
		AttackerID: attacker.CharacterID, TargetID: target.CharacterID,
		Hit: result.Hit, Damage: result.Damage, HitChance: result.HitChance,
	})
	h.broadcastToInterested(attacker, protocol.MsgCombatResult, protocol.CombatResultMsg{
		AttackerID: attacker.CharacterID, TargetID: target.CharacterID,
		Hit: result.Hit, Damage: result.Damage, HitChance: result.HitChance,
	})

	if target.Combat.IncapacitatedAt != nil {
		// Existing rules already stop here; no PvP-specific behaviour.
		return decision
	}

	// D-03: any faction member involved refreshes its aggression timer.
	now := time.Now()
	for _, c := range []*Client{attacker, target} {
		cf, err := h.db.GetCharacterFaction(c.CharacterID)
		if err != nil || cf == nil {
			continue
		}
		v := h.pvpFlagView(c.CharacterID)
		v = v.RefreshAggro(now)
		if err := h.savePvPFlag(c.CharacterID, v); err != nil {
			log.Printf("pvp: aggro refresh failed for %s: %v", c.CharacterID, err)
		}
		h.sendPvPStateChange(c, h.zoneAt(c.Pos.Planet, c.Pos.X, c.Pos.Z))
	}

	if !result.Hit {
		return decision
	}

	h.mu.Lock()
	newState, justIncap := combat.ApplyDamage(target.Combat, result.Damage, now)
	target.Combat = newState
	updateErr := h.db.UpdateCombatState(target.CharacterID, target.Combat)
	h.mu.Unlock()
	if updateErr != nil {
		log.Printf("pvp: combat-state persist failed: %v", updateErr)
	}

	h.sendHAMUpdateLocked(target, target.Combat)

	if justIncap {
		h.send(target, protocol.MsgIncapacitated, protocol.IncapacitatedMsg{CharacterID: target.CharacterID})
		// D-06: logged, never awarded. Basis follows the eval order.
		h.logPvPDefeat(attacker.CharacterID, target.CharacterID, decision.Basis, tZoneID)
		h.endDuelOnDefeat(target.CharacterID)
	}

	return decision
}

// endDuelOnDefeat closes an active duel when a participant is defeated
// (D-09: ends on incapacitation), naming the survivor winner.
func (h *WorldHandler) endDuelOnDefeat(loserID string) {
	duel, err := h.db.ActiveDuelFor(loserID)
	if err != nil || duel == nil || duel.Status != "active" {
		return
	}
	winnerID := duel.ChallengerID
	if loserID == duel.ChallengerID {
		winnerID = duel.TargetID
	}
	now := sql.NullInt64{Int64: time.Now().Unix(), Valid: true}
	_ = h.db.UpdateDuelStatus(duel.ID, "ended", sql.NullInt64{}, now,
		sql.NullString{String: winnerID, Valid: true},
		sql.NullString{String: "incapacitated", Valid: true})

	h.mu.RLock()
	winner, ok := h.clients[winnerID]
	loser, lok := h.clients[loserID]
	h.mu.RUnlock()
	if ok {
		h.send(winner, protocol.MsgDuelUpdate, protocol.DuelUpdateMsg{
			DuelID: duel.ID, Status: "ended", Opponent: loserID,
			Reason: "incapacitated", WasVictor: true,
		})
	}
	if lok {
		h.send(loser, protocol.MsgDuelUpdate, protocol.DuelUpdateMsg{
			DuelID: duel.ID, Status: "ended", Opponent: winnerID, Reason: "incapacitated",
		})
	}
}
