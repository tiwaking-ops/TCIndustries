// Faction handlers for Phase 8: alignment flagging, PvP targeting validity,
// player-vs-player resolution, base placement/siege, faction chat routing, and
// the PvP death legs (points, credit transfer, condition). Testbed fork;
// generic content only. GDD-silent numbers [PROVISIONAL] per the owner-approved
// convention; GDD-given values cited.
package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"time"

	"swg-server/internal/combat"
	"swg-server/internal/database"
	"swg-server/internal/faction"
	"swg-server/internal/protocol"
	"swg-server/internal/world"
)

// SeedFactionWorld prepares Phase 8 persistence. Idempotent.
func (h *WorldHandler) SeedFactionWorld() error {
	if err := h.db.EnsureFactionSchema(); err != nil {
		return err
	}
	log.Printf("Faction world seeded: standings, pending kills, bases")
	return nil
}

// --- F1: overt/covert flagging (WS) ---

func (h *WorldHandler) handleGoOvert(client *Client, raw []byte) {
	if client.CharacterID == "" {
		h.sendError(client, "not in world")
		return
	}
	st, err := h.db.GetStanding(client.CharacterID)
	if err != nil {
		h.sendError(client, "standing unavailable")
		return
	}
	if st.Alignment == faction.AlignNeutral {
		h.sendError(client, "declare an alignment first")
		return
	}
	if st.Overt {
		h.sendError(client, "already overt")
		return
	}
	now := time.Now().Unix()
	st.Overt = true
	st.OvertSince = sql.NullInt64{Int64: now, Valid: true}
	if err := h.db.SetStanding(st); err != nil {
		h.sendError(client, "flagging failed")
		return
	}
	h.broadcastFlag(client, st)
}

func (h *WorldHandler) handleGoCovert(client *Client, raw []byte) {
	if client.CharacterID == "" {
		h.sendError(client, "not in world")
		return
	}
	st, err := h.db.GetStanding(client.CharacterID)
	if err != nil {
		h.sendError(client, "standing unavailable")
		return
	}
	if !st.Overt {
		h.sendError(client, "already covert")
		return
	}
	now := time.Now()
	wait := faction.CovertDelay
	if st.OvertSince.Valid {
		elapsed := now.Unix() - st.OvertSince.Int64
		if elapsed < faction.Secs(wait) {
			h.sendError(client, "covert delay not elapsed (anti-combat-log)")
			return
		}
	}
	st.Overt = false
	st.OvertSince = sql.NullInt64{}
	if err := h.db.SetStanding(st); err != nil {
		h.sendError(client, "flagging failed")
		return
	}
	h.broadcastFlag(client, st)
}

// broadcastFlag pushes flag state to the client and interested observers
// (overt must be readable at a glance — GDD 9.4.1).
func (h *WorldHandler) broadcastFlag(client *Client, st *database.StandingRow) {
	msg := protocol.FlagChangedMsg{Overt: st.Overt, Faction: st.Alignment}
	h.send(client, protocol.MsgFlagChanged, msg)
	h.broadcastToInterested(client, protocol.MsgFlagChanged, msg)
}

// flagAttackerOvert applies "attacking enemy immediately flags Overt"
// (GDD 9.4.1): set only when not already overt, so fighting never extends a
// player's own covert-delay clock (flagged reading).
func (h *WorldHandler) flagAttackerOvert(client *Client) {
	st, err := h.db.GetStanding(client.CharacterID)
	if err != nil || st.Overt || st.Alignment == faction.AlignNeutral {
		return
	}
	st.Overt = true
	st.OvertSince = sql.NullInt64{Int64: time.Now().Unix(), Valid: true}
	if err := h.db.SetStanding(st); err != nil {
		return
	}
	h.broadcastFlag(client, st)
}

// --- F2: PvP validity + resolution ---

// pvpDeny reasons a player-vs-player attack. Empty string = valid.
//
// GDD 9.4.1 tension resolved (proposal §5): "Covert ... cannot attack enemies"
// is read as "cannot attack WHILE REMAINING covert" — a covert attacker whose
// target is otherwise valid is auto-flagged overt by the attempt (second GDD
// sentence, operative). Rejected attempts never flag. Victims must always be
// overt (no symmetrical exception — safety is one-sided by design).
func (h *WorldHandler) pvpDeny(attacker *Client, victim *Client) string {
	if victim.CharacterID == attacker.CharacterID {
		return "cannot attack yourself"
	}
	ast, err := h.db.GetStanding(attacker.CharacterID)
	if err != nil {
		return "standing unavailable"
	}
	vst, err := h.db.GetStanding(victim.CharacterID)
	if err != nil {
		return "target standing unavailable"
	}
	if ast.Alignment == faction.AlignNeutral || !faction.Opposing(ast.Alignment, vst.Alignment) {
		return "no valid enemy alignment"
	}
	if !vst.Overt {
		return "target is not overt"
	}
	if victim.Pos.Planet != attacker.Pos.Planet {
		return "target not in this zone"
	}
	if world.Distance2D(attacker.Pos, victim.Pos) > AttackRangeM {
		return "target out of range"
	}
	vtst, err := h.db.GetCombatState(victim.CharacterID)
	if err != nil || vtst.IncapacitatedAt != nil {
		return "target must be conscious"
	}
	// City permission at the defender's ground (GDD 9.4.2 "if city allows").
	if c, err := h.db.CityContaining(victim.Pos.Planet, victim.Pos.X, victim.Pos.Z); err == nil {
		if !h.db.CityPvPAllowed(c.ID) {
			return "pvp denied by city"
		}
	}
	return ""
}

// stylePool maps a weapon style tag to its elite XP pool (Phase 9).
func stylePool(style string) string {
	switch style {
	case "pistol":
		return "pistol_combat"
	case "rifle":
		return "rifle_combat"
	case "carbine":
		return "carbine_combat"
	case "fencing":
		return "fencing_combat"
	case "sword":
		return "sword_combat"
	case "polearm":
		return "polearm_combat"
	case "bounty":
		return "bounty_hunter"
	case "commando":
		return "commando"
	default:
		return "combat"
	}
}

// attackerWeapon mirrors the PvE weapon link (crafted sidearm replaces the
// unarmed range and adds accuracy max as a BaseAccuracy bonus). [PROVISIONAL mapping]
// Phase 9: TKA Power Strikes tiers add unarmed damage (+2/tier, gate-add,
// flagged) when fighting bare-handed.
func (h *WorldHandler) attackerWeapon(client *Client) combat.Weapon {
	weapon := unarmedWeapon
	wItem, _, werr := h.db.EquippedItems(client.CharacterID)
	if werr != nil {
		return weapon
	}
	if wItem == nil {
		// Bare-handed TKA scaling (gate-add, flagged).
		tier := h.skillTier(client.CharacterID, "teraskasi_power_strikes_")
		weapon.MinDamage += 2 * tier
		weapon.MaxDamage += 2 * tier
		return weapon
	}
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
	return weapon
}

// handlePvPAttack resolves one accepted player-vs-player attack.
func (h *WorldHandler) handlePvPAttack(client *Client, victim *Client) {
	attacker := combat.Attacker{
		AccuracySkill: 0, DamageBonus: UnarmedDamageBonus,
		Posture: combat.PostureStanding, Stance: combat.StanceNormal,
	}
	// Attacker posture/stance shape mirrors the PvE path.
	if st, err := h.db.GetCombatState(client.CharacterID); err == nil {
		attacker.Posture = combat.Posture(st.Posture)
		attacker.Stance = combat.Stance(st.Stance)
	}
	weapon := h.attackerWeapon(client)
	// Defender mitigation mirrors the creature-tick path (crafted plate
	// protection_max, capped at 50). [PROVISIONAL]
	mitigation := 0.0
	if _, armor, err := h.db.EquippedItems(victim.CharacterID); err == nil && armor != nil {
		if prot, ok := armor.Stats["protection_max"]; ok {
			mitigation = prot
			if mitigation > 50 {
				mitigation = 50
			}
		}
	}
	vtst, vcur, err := h.loadCombatState(victim.CharacterID)
	if err != nil {
		h.sendError(client, "target combat state unavailable")
		return
	}
	defender := combat.Defender{
		DefenseSkill: 0, Posture: combat.Posture(vtst.Posture),
		ArmorMitigationPct: mitigation, ResistanceMultiplier: 1.0,
	}
	result := combat.ResolveAttack(attacker, defender, weapon,
		rand.New(rand.NewSource(time.Now().UnixNano())))
	crMsg := protocol.CombatResultMsg{
		AttackerID: client.CharacterID, TargetID: victim.CharacterID,
		Hit: result.Hit, Damage: result.Damage, HitChance: result.HitChance,
	}
	h.broadcastToInterested(client, protocol.MsgCombatResult, crMsg)
	h.send(client, protocol.MsgCombatResult, crMsg)
	if tv, ok := h.findClient(victim.CharacterID); ok {
		h.send(tv, protocol.MsgCombatResult, crMsg) // victim always sees own hits
	}
	if !result.Hit {
		return
	}
	updated, justIncap := combat.ApplyDamage(vcur, result.Damage, time.Now())
	applyHAMPoolState(vtst, updated)
	if err := h.db.UpdateCombatState(vtst); err != nil {
		h.sendError(client, "damage failed")
		return
	}
	h.pushHAMUpdate(victim.CharacterID)
	if justIncap {
		if tv, ok := h.findClient(victim.CharacterID); ok {
			h.send(tv, protocol.MsgIncapacitated,
				protocol.IncapacitatedMsg{CharacterID: victim.CharacterID})
		}
		// Kill credit runs through death: pending row consumed at clone,
		// voided at revive (faction_db.go design note).
		_ = h.db.RecordPendingKill(victim.CharacterID, client.CharacterID, time.Now().Unix())
	}
}

// --- F5: base placement (WS positional, hall-placement precedent) ---

func (h *WorldHandler) handlePlaceBase(client *Client, raw []byte) {
	if client.CharacterID == "" {
		h.sendError(client, "not in world")
		return
	}
	var msg protocol.WSMessage
	json.Unmarshal(raw, &msg)
	dataBytes, _ := json.Marshal(msg.Data)
	var pm protocol.PlaceBaseMsg
	if err := json.Unmarshal(dataBytes, &pm); err != nil {
		h.sendError(client, "invalid place_base data")
		return
	}
	deed, err := h.db.GetItem(client.CharacterID, pm.DeedItemID)
	if err != nil || deed.Schematic != "base_deed" {
		h.sendError(client, "valid base deed required")
		return
	}
	g, err := h.db.GetGuild(pm.GuildID)
	if err != nil {
		h.sendError(client, "guild unknown")
		return
	}
	role, err := h.db.GuildMemberRole(pm.GuildID, client.CharacterID)
	if err != nil || (role != "leader" && role != "officer") {
		h.sendError(client, "requires guild officer")
		return
	}
	// Officer alignment gate (GDD 15.2.4 [ASSUMPTION]): every officer-slot
	// holder shares one non-neutral alignment; the placer must hold Major+.
	officerAligns, err := h.db.GuildOfficerAlignments(pm.GuildID)
	if err != nil || len(officerAligns) == 0 {
		h.sendError(client, "officer alignment unavailable")
		return
	}
	for _, a := range officerAligns {
		if a == faction.AlignNeutral || a != g.Faction {
			h.sendError(client, "officers must share the guild alignment")
			return
		}
	}
	pst, err := h.db.GetStanding(client.CharacterID)
	if err != nil {
		h.sendError(client, "standing unavailable")
		return
	}
	if faction.RankFor(pst.Points) != faction.RankMajor &&
		faction.RankFor(pst.Points) != faction.RankColonel {
		h.sendError(client, "requires Major rank or better")
		return
	}
	if pst.Alignment != g.Faction {
		h.sendError(client, "placer must hold the guild alignment")
		return
	}
	if _, err := h.db.GuildActiveBase(pm.GuildID); err == nil {
		h.sendError(client, "guild already fields a base")
		return
	}
	// Wilderness-equivalent: outside every active city radius (point check),
	// 20 m from structures and standing bases.
	if _, err := h.db.CityContaining(client.Pos.Planet, pm.X, pm.Z); err == nil {
		h.sendError(client, "bases cannot be placed inside cities")
		return
	}
	if err := h.checkNoBuild(client.Pos.Planet, pm.X, pm.Z); err != nil {
		h.sendError(client, err.Error())
		return
	}
	if err := h.checkBaseSeparation(client.Pos.Planet, pm.X, pm.Z); err != nil {
		h.sendError(client, err.Error())
		return
	}
	baseID := h.db.NewRowID("base")
	if err := h.db.PlaceBase(&database.BaseRow{
		ID: baseID, GuildID: pm.GuildID, Faction: g.Faction,
		Zone: client.Pos.Planet, PosX: pm.X, PosZ: pm.Z,
		HP: faction.BaseHP, HPMax: faction.BaseHP, PlacedBy: client.CharacterID,
	}); err != nil {
		h.sendError(client, "base placement failed")
		return
	}
	_ = h.db.DeleteItem(client.CharacterID, pm.DeedItemID)
	_ = h.db.AddCharacterXP(client.CharacterID, "structure_crafting", 100)
	h.send(client, protocol.MsgBasePlaced, protocol.BasePlacedMsg{BaseID: baseID})
}

// checkBaseSeparation enforces base-vs-base spacing [PROVISIONAL 20 m].
func (h *WorldHandler) checkBaseSeparation(zone string, x, z float64) error {
	all, err := h.db.AllBases()
	if err != nil {
		return err
	}
	for _, b := range all {
		if b.Zone != zone || b.Status != "active" {
			continue
		}
		dx, dz := x-b.PosX, z-b.PosZ
		if dx*dx+dz*dz < faction.BaseNoBuildM*faction.BaseNoBuildM {
			return fmt.Errorf("too close to another base")
		}
	}
	return nil
}

// handleBaseAttack resolves one accepted siege hit.
func (h *WorldHandler) handleBaseAttack(client *Client, base *database.BaseRow) {
	now := time.Now().Unix()
	attacker := combat.Attacker{
		AccuracySkill: 0, DamageBonus: UnarmedDamageBonus,
		Posture: combat.PostureStanding, Stance: combat.StanceNormal,
	}
	if st, err := h.db.GetCombatState(client.CharacterID); err == nil {
		attacker.Posture = combat.Posture(st.Posture)
		attacker.Stance = combat.Stance(st.Stance)
	}
	weapon := h.attackerWeapon(client)
	// Bases do not dodge: dummy defender (reuses the resolution shape against
	// zero defense — flagged; turret/evasion mechanics are EXPANSION).
	defender := combat.Defender{
		DefenseSkill: 0, Posture: combat.PostureStanding,
		ArmorMitigationPct: 0, ResistanceMultiplier: 1.0,
	}
	result := combat.ResolveAttack(attacker, defender, weapon,
		rand.New(rand.NewSource(time.Now().UnixNano())))
	crMsg := protocol.CombatResultMsg{
		AttackerID: client.CharacterID, TargetID: base.ID,
		Hit: result.Hit, Damage: result.Damage, HitChance: result.HitChance,
	}
	h.broadcastToInterested(client, protocol.MsgCombatResult, crMsg)
	h.send(client, protocol.MsgCombatResult, crMsg)
	if !result.Hit {
		return
	}
	base.HP -= result.Damage
	if base.HP < 0 {
		base.HP = 0
	}
	_ = h.db.LogBaseDamage(base.ID, client.CharacterID, result.Damage, now)
	if base.HP <= 0 {
		base.Status = "destroyed"
		_ = h.db.UpdateBase(base)
		windowStart := int64(0)
		if base.WindowStart.Valid {
			windowStart = base.WindowStart.Int64
		}
		participants, _ := h.db.BaseParticipants(base.ID, windowStart)
		_ = h.db.AwardDestroy(base.ID, "base_destroy", faction.DestroyAward, participants)
		h.broadcastToInterested(client, protocol.MsgBaseDestroyed,
			protocol.BaseDestroyedMsg{BaseID: base.ID, Destroyed: true})
		h.send(client, protocol.MsgBaseDestroyed,
			protocol.BaseDestroyedMsg{BaseID: base.ID, Destroyed: true})
		for _, pid := range participants {
			if pc, ok := h.findClient(pid); ok {
				st, _ := h.db.GetStanding(pid)
				h.send(pc, protocol.MsgPointsAwarded, protocol.PointsAwardedMsg{
					Amount: faction.DestroyAward, Reason: "base_destroy",
					Total: st.Points, Rank: faction.RankFor(st.Points),
				})
			}
		}
		return
	}
	_ = h.db.UpdateBase(base)
}

// baseDeny reasons a siege hit. Empty string = valid.
func (h *WorldHandler) baseDeny(attacker *Client, base *database.BaseRow) string {
	if base.Status != "active" {
		return "base already destroyed"
	}
	ast, err := h.db.GetStanding(attacker.CharacterID)
	if err != nil {
		return "standing unavailable"
	}
	if !ast.Overt {
		return "you must be overt to siege"
	}
	if ast.Alignment == faction.AlignNeutral || ast.Alignment == base.Faction {
		return "no valid enemy alignment"
	}
	now := time.Now().Unix()
	if !base.WindowStart.Valid || !base.WindowEnd.Valid ||
		now < base.WindowStart.Int64 || now > base.WindowEnd.Int64 {
		return "base outside vulnerability window"
	}
	if base.Zone != attacker.Pos.Planet {
		return "base not in this zone"
	}
	dx := attacker.Pos.X - base.PosX
	dz := attacker.Pos.Z - base.PosZ
	if dx*dx+dz*dz > faction.BaseAttackRangeM*faction.BaseAttackRangeM {
		return "base out of range"
	}
	return ""
}

// --- Faction chat (fills the Phase 7 stub) ---

func (h *WorldHandler) handleFactionChat(client *Client, text string) {
	if text == "" {
		h.sendError(client, "empty chat text")
		return
	}
	st, err := h.db.GetStanding(client.CharacterID)
	if err != nil || st.Alignment == faction.AlignNeutral {
		h.sendError(client, "no alignment")
		return
	}
	out := protocol.ChatMessageMsg{
		SenderName: client.Name, Channel: "faction", Text: text,
	}
	// Chat reaches all online same-side members (overt state is irrelevant to
	// chat; the overt roster is a separate HTTP surface for Phase 9 input).
	h.send(client, protocol.MsgChatMessage, out) // echo
	h.mu.RLock()
	defer h.mu.RUnlock()
	for _, c := range h.clients {
		if c.CharacterID == client.CharacterID || c.CharacterID == "" {
			continue
		}
		ost, err := h.db.GetStanding(c.CharacterID)
		if err != nil || ost.Alignment != st.Alignment {
			continue
		}
		h.send(c, protocol.MsgChatMessage, out)
	}
}

// --- F3: death legs (clone path; timer deaths route through cloneCharacter) ---

// applyPvPDeathLegs consumes a pending kill into points + credit transfer +
// condition loss, and notifies the killer when online.
func (h *WorldHandler) applyPvPDeathLegs(victim *Client) {
	killerID, ok := h.db.TakePendingKill(victim.CharacterID)
	if !ok {
		return
	}
	transferred, err := h.db.ApplyPvPDeath(victim.CharacterID, killerID,
		faction.PvPCreditPct, faction.ConditionLoss, faction.KillAward,
		victim.Pos.Planet)
	if err != nil {
		log.Printf("pvp death legs failed: %v", err)
		return
	}
	_ = transferred
	// Phase 9: open player-bounty contracts on the victim pay out to the
	// killer from poster escrow (stacked contracts all pay — flagged).
	_, _ = h.db.PayBounties(victim.CharacterID, killerID, victim.Pos.Planet, BountyHunterXP)
	if kc, online := h.findClient(killerID); online {
		st, _ := h.db.GetStanding(killerID)
		h.send(kc, protocol.MsgPointsAwarded, protocol.PointsAwardedMsg{
			Amount: faction.KillAward, Reason: "pvp_kill",
			Total: st.Points, Rank: faction.RankFor(st.Points),
		})
	}
}

// maybeDecayPoints applies lazy rank-point decay when the latest activity is
// older than the window (GDD 15.2.3 [ASSUMPTION]; code-only path — 60-day
// durations are untestable live). At most one halving per window: the decay
// event itself refreshes the activity timestamp (flagged semantics).
func (h *WorldHandler) maybeDecayPoints(charID string) {
	st, err := h.db.GetStanding(charID)
	if err != nil || st.Points <= 0 {
		return
	}
	// Freshness signal is the newest point-event row.
	latest, err := h.db.LatestPointEventAt(charID)
	if err != nil || time.Now().Unix()-latest < faction.Secs(faction.PointsDecayAfter) {
		return
	}
	halved := st.Points / 2
	st.Points = halved
	_ = h.db.SetStanding(st)
	_ = h.db.AddPoints(charID, 0, "decay", "")
	log.Printf("faction decay: %s halved to %d (stale activity)", charID, halved)
}
