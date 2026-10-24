// Phase 8 (Faction & PvP) world handlers: faction NPC interaction, flag
// toggles, duels, battlefield zones, and the PvP attack gate wired into
// handleCombatAction. Testbed fork; generic placeholder content only.
//
// Design authority: proposals/proposal-claude-tcindustries-phase8-faction-pvp-v1.md
//   R-01..R-08 — human rulings. D-01..D-16 — approved defaults (internal/pvp.PvPConfig).
package handlers

import (
	"database/sql"
	"encoding/json"
	"math"
	"time"

	"tc-server/internal/database"
	"tc-server/internal/pvp"
	"tc-server/internal/protocol"
)

// --- Helpers ---

// pvpFlagView loads a character's flag state, defaulting to covert for a
// faction member and returning an empty view for a neutral (no row).
func (h *WorldHandler) pvpFlagView(charID string) pvp.FlagView {
	row, err := h.db.GetPvPFlag(charID)
	if err != nil || row == nil {
		return pvp.NewFlagView()
	}
	v := pvp.FlagView{State: pvp.FlagState(row.State)}
	v.Pending = pvp.FlagState(row.PendingState.String)
	if row.TransitionAt.Valid {
		v.ReadyAt = time.Unix(row.TransitionAt.Int64, 0)
	}
	if row.AggressionExpiresAt > 0 {
		v.AggressionExpiresAt = time.Unix(row.AggressionExpiresAt, 0)
	}
	return v
}

// savePvPFlag persists a flag view.
func (h *WorldHandler) savePvPFlag(charID string, v pvp.FlagView) error {
	row := &database.PvPFlagRow{
		CharacterID:         charID,
		State:               string(v.State),
		PendingState:        sql.NullString{String: string(v.Pending), Valid: v.Pending != ""},
		AggressionExpiresAt: v.AggressionExpiresAt.Unix(),
	}
	if !v.ReadyAt.IsZero() {
		row.TransitionAt = sql.NullInt64{Int64: v.ReadyAt.Unix(), Valid: true}
	}
	if v.AggressionExpiresAt.IsZero() {
		row.AggressionExpiresAt = 0
	}
	return h.db.UpsertPvPFlag(row)
}

// sendPvPStateChange pushes the current flag/timer/zone state to a client.
func (h *WorldHandler) sendPvPStateChange(client *Client, zoneID string) {
	if client.CharacterID == "" {
		return
	}
	now := time.Now()
	v := h.pvpFlagView(client.CharacterID)
	out := protocol.PvPStateChangeMsg{
		CharacterID:        client.CharacterID,
		State:              string(v.State),
		TransitionInMS:     v.TransitionIn(now).Milliseconds(),
		AggressionRemainMS: v.AggressionIn(now).Milliseconds(),
		InBattlefield:      zoneID,
	}
	if v.HasPending() {
		out.PendingState = string(v.Pending)
	}
	h.send(client, protocol.MsgPvPStateChange, out)
}

// sendFactionUpdate pushes faction membership to a client (login + on change).
func (h *WorldHandler) sendFactionUpdate(client *Client) {
	if client.CharacterID == "" {
		return
	}
	out := protocol.FactionUpdateMsg{CharacterID: client.CharacterID}
	if cf, err := h.db.GetCharacterFaction(client.CharacterID); err == nil && cf != nil {
		out.FactionID = cf.FactionID
		out.Standing = cf.Standing
		if f, err := h.db.GetFactionByID(cf.FactionID); err == nil && f != nil {
			out.FactionCode = f.Code
			out.FactionName = f.DisplayName
		}
	}
	h.send(client, protocol.MsgFactionUpdate, out)
}

// zoneAt returns the battlefield zone containing a position ("" if none).
func (h *WorldHandler) zoneAt(planet string, x, z float64) string {
	zones, err := h.db.ListZones()
	if err != nil {
		return ""
	}
	pvpZones := make([]pvp.Zone, 0, len(zones))
	for _, zr := range zones {
		if zr.Zone != planet {
			continue
		}
		pvpZones = append(pvpZones, pvp.Zone{
			ID: zr.ID, Shape: pvp.ZoneShape(zr.Shape),
			CenterX: zr.CenterX, CenterZ: zr.CenterZ, Radius: zr.RadiusM,
			MinX: zr.MinX, MinZ: zr.MinZ, MaxX: zr.MaxX, MaxZ: zr.MaxZ, Active: true,
		})
	}
	return pvp.ZoneContaining(pvpZones, x, z)
}

// --- Faction NPC interaction (R-04, Step 4) ---

func (h *WorldHandler) handleSetFaction(client *Client, raw []byte) {
	if client.CharacterID == "" {
		h.sendError(client, "not in world")
		return
	}
	var msg protocol.WSMessage
	json.Unmarshal(raw, &msg)
	dataBytes, _ := json.Marshal(msg.Data)
	var fm protocol.SetFactionMsg
	if err := json.Unmarshal(dataBytes, &fm); err != nil {
		h.sendPvPError(client, "invalid set_faction data")
		return
	}

	// Validate the NPC exists and the client is within its interaction radius
	// against the authoritative server position (R-04).
	recruiters, err := h.db.ListFactionRecruiters()
	if err != nil {
		h.sendPvPError(client, "recruiter_lookup_failed")
		return
	}
	var npc *database.RecruiterRow
	for i := range recruiters {
		if recruiters[i].ID == fm.NPCID {
			npc = &recruiters[i]
			break
		}
	}
	if npc == nil {
		h.sendPvPError(client, "recruiter_unknown")
		return
	}
	radius := pvp.Cfg.RecruiterInteractM
	if npc.InteractM > 0 {
		radius = npc.InteractM
	}
	dx := client.Pos.X - npc.PosX
	dz := client.Pos.Z - npc.PosZ
	if dx*dx+dz*dz > radius*radius {
		h.sendPvPError(client, "recruiter_out_of_range")
		return
	}

	now := time.Now()
	existing, err := h.db.GetCharacterFaction(client.CharacterID)
	hasFaction := err == nil && existing != nil

	switch fm.Action {
	case "join":
		if hasFaction {
			h.sendPvPError(client, "already_in_faction")
			return
		}
		if fm.FactionID == "" {
			h.sendPvPError(client, "faction_id_required")
			return
		}
		if _, err := h.db.GetFactionByID(fm.FactionID); err != nil {
			h.sendPvPError(client, "faction_unknown")
			return
		}
		if err := h.db.SetCharacterFaction(client.CharacterID, fm.FactionID); err != nil {
			h.sendPvPError(client, "join_failed")
			return
		}
		// New members start covert (proposal 4.2).
		v := pvp.NewFlagView()
		_ = h.savePvPFlag(client.CharacterID, v)
		h.sendFactionUpdate(client)
		h.sendPvPStateChange(client, h.zoneAt(client.Pos.Planet, client.Pos.X, client.Pos.Z))

	case "switch":
		if !hasFaction {
			h.sendPvPError(client, "not_in_faction")
			return
		}
		if fm.FactionID == "" {
			h.sendPvPError(client, "faction_id_required")
			return
		}
		if existing.FactionID == fm.FactionID {
			h.sendPvPError(client, "same_faction")
			return
		}
		if _, err := h.db.GetFactionByID(fm.FactionID); err != nil {
			h.sendPvPError(client, "faction_unknown")
			return
		}
		// D-04: blocked while overt or while aggression is active.
		v := h.pvpFlagView(client.CharacterID)
		if ok, reason := v.CanChangeFaction(now); !ok {
			h.sendPvPError(client, reason)
			return
		}
		if err := h.db.SetCharacterFaction(client.CharacterID, fm.FactionID); err != nil {
			h.sendPvPError(client, "switch_failed")
			return
		}
		// D-05: standing resets (handled by SetCharacterFaction); flag → covert.
		_ = h.savePvPFlag(client.CharacterID, pvp.NewFlagView())
		h.sendFactionUpdate(client)
		h.sendPvPStateChange(client, h.zoneAt(client.Pos.Planet, client.Pos.X, client.Pos.Z))

	case "leave":
		if !hasFaction {
			h.sendPvPError(client, "not_in_faction")
			return
		}
		v := h.pvpFlagView(client.CharacterID)
		if ok, reason := v.CanChangeFaction(now); !ok {
			h.sendPvPError(client, reason)
			return
		}
		if err := h.db.RemoveCharacterFaction(client.CharacterID); err != nil {
			h.sendPvPError(client, "leave_failed")
			return
		}
		_ = h.db.DeletePvPFlag(client.CharacterID)
		h.sendFactionUpdate(client)

	default:
		h.sendPvPError(client, "unknown_action")
	}
}

// --- Flag toggle (R-05, Step 5) ---

func (h *WorldHandler) handleTogglePvP(client *Client, raw []byte) {
	if client.CharacterID == "" {
		h.sendError(client, "not in world")
		return
	}
	var msg protocol.WSMessage
	json.Unmarshal(raw, &msg)
	dataBytes, _ := json.Marshal(msg.Data)
	var tm protocol.TogglePvPMsg
	if err := json.Unmarshal(dataBytes, &tm); err != nil {
		h.sendPvPError(client, "invalid toggle_pvp data")
		return
	}

	// Only faction members have flags.
	cf, err := h.db.GetCharacterFaction(client.CharacterID)
	if err != nil || cf == nil {
		h.sendPvPError(client, "no_faction")
		return
	}

	now := time.Now()
	v := h.pvpFlagView(client.CharacterID)

	switch pvp.FlagState(tm.State) {
	case pvp.FlagOvert:
		trans := v.RequestOvert(now)
		v.State = trans.State
		v.Pending = trans.Pending
		v.ReadyAt = trans.ReadyAt
	case pvp.FlagCovert:
		trans, ok := v.RequestCovert(now)
		if !ok {
			h.sendPvPError(client, "covert_blocked_aggression")
			return
		}
		v.State = trans.State
		v.Pending = trans.Pending
		v.ReadyAt = trans.ReadyAt
	default:
		h.sendPvPError(client, "invalid_state")
		return
	}

	if err := h.savePvPFlag(client.CharacterID, v); err != nil {
		h.sendPvPError(client, "save_failed")
		return
	}
	h.sendPvPStateChange(client, h.zoneAt(client.Pos.Planet, client.Pos.X, client.Pos.Z))
}

// --- Duels (R-01a, Step 8) ---

func (h *WorldHandler) handleDuelChallenge(client *Client, raw []byte) {
	if client.CharacterID == "" {
		h.sendError(client, "not in world")
		return
	}
	var msg protocol.WSMessage
	json.Unmarshal(raw, &msg)
	dataBytes, _ := json.Marshal(msg.Data)
	var dm protocol.DuelChallengeMsg
	if err := json.Unmarshal(dataBytes, &dm); err != nil {
		h.sendPvPError(client, "invalid duel_challenge data")
		return
	}
	if dm.TargetID == client.CharacterID {
		h.sendPvPError(client, "self")
		return
	}

	h.mu.RLock()
	target, ok := h.clients[dm.TargetID]
	h.mu.RUnlock()
	if !ok {
		h.sendPvPError(client, "target_not_found")
		return
	}

	// D-08: both participants within DuelStartRangeM.
	dist := math.Hypot(client.Pos.X-target.Pos.X, client.Pos.Z-target.Pos.Z)
	if dist > pvp.Cfg.DuelStartRangeM {
		h.sendPvPError(client, "duel_out_of_range")
		return
	}

	// Neither may already be in a pending/active duel.
	if _, err := h.db.ActiveDuelFor(client.CharacterID); err == nil {
		h.sendPvPError(client, "already_in_duel")
		return
	}

	duelID := h.db.NewRowID("duel")
	if err := h.db.CreateDuel(duelID, client.CharacterID, dm.TargetID); err != nil {
		h.sendPvPError(client, "duel_create_failed")
		return
	}
	h.send(client, protocol.MsgDuelUpdate, protocol.DuelUpdateMsg{
		DuelID: duelID, Status: "pending", Opponent: dm.TargetID,
	})
	h.send(target, protocol.MsgDuelUpdate, protocol.DuelUpdateMsg{
		DuelID: duelID, Status: "pending", Opponent: client.CharacterID,
	})
}

func (h *WorldHandler) handleDuelResponse(client *Client, raw []byte) {
	if client.CharacterID == "" {
		h.sendError(client, "not in world")
		return
	}
	var msg protocol.WSMessage
	json.Unmarshal(raw, &msg)
	dataBytes, _ := json.Marshal(msg.Data)
	var rm protocol.DuelResponseMsg
	if err := json.Unmarshal(dataBytes, &rm); err != nil {
		h.sendPvPError(client, "invalid duel_response data")
		return
	}
	duel, err := h.db.GetDuel(rm.DuelID)
	if err != nil || duel == nil {
		h.sendPvPError(client, "duel_unknown")
		return
	}
	if duel.TargetID != client.CharacterID || duel.Status != "pending" {
		h.sendPvPError(client, "duel_not_pending")
		return
	}

	now := sql.NullInt64{Int64: time.Now().Unix(), Valid: true}
	var win, reason sql.NullString
	if rm.Accept {
		_ = h.db.UpdateDuelStatus(duel.ID, "active", now, sql.NullInt64{}, win, reason)
		// Notify both participants.
		h.mu.RLock()
		challenger, cOK := h.clients[duel.ChallengerID]
		h.mu.RUnlock()
		if cOK {
			h.send(challenger, protocol.MsgDuelUpdate, protocol.DuelUpdateMsg{
				DuelID: duel.ID, Status: "active", Opponent: duel.TargetID,
			})
		}
		h.send(client, protocol.MsgDuelUpdate, protocol.DuelUpdateMsg{
			DuelID: duel.ID, Status: "active", Opponent: duel.ChallengerID,
		})
	} else {
		_ = h.db.UpdateDuelStatus(duel.ID, "declined", sql.NullInt64{}, now, win, reason)
		h.mu.RLock()
		challenger, cOK := h.clients[duel.ChallengerID]
		h.mu.RUnlock()
		if cOK {
			h.send(challenger, protocol.MsgDuelUpdate, protocol.DuelUpdateMsg{
				DuelID: duel.ID, Status: "declined", Opponent: duel.TargetID, Reason: "declined",
			})
		}
		h.send(client, protocol.MsgDuelUpdate, protocol.DuelUpdateMsg{
			DuelID: duel.ID, Status: "declined", Opponent: duel.ChallengerID,
		})
	}
}

func (h *WorldHandler) handleDuelForfeit(client *Client, raw []byte) {
	if client.CharacterID == "" {
		h.sendError(client, "not in world")
		return
	}
	var msg protocol.WSMessage
	json.Unmarshal(raw, &msg)
	dataBytes, _ := json.Marshal(msg.Data)
	var fm protocol.DuelForfeitMsg
	if err := json.Unmarshal(dataBytes, &fm); err != nil {
		h.sendPvPError(client, "invalid duel_forfeit data")
		return
	}
	duel, err := h.db.GetDuel(fm.DuelID)
	if err != nil || duel == nil {
		h.sendPvPError(client, "duel_unknown")
		return
	}
	if duel.Status != "active" && duel.Status != "pending" {
		h.sendPvPError(client, "duel_not_active")
		return
	}
	if duel.ChallengerID != client.CharacterID && duel.TargetID != client.CharacterID {
		h.sendPvPError(client, "not_participant")
		return
	}

	// The other participant is the winner.
	winner := duel.ChallengerID
	if client.CharacterID == duel.ChallengerID {
		winner = duel.TargetID
	}
	now := sql.NullInt64{Int64: time.Now().Unix(), Valid: true}
	_ = h.db.UpdateDuelStatus(duel.ID, "ended", sql.NullInt64{}, now,
		sql.NullString{String: winner, Valid: true}, sql.NullString{String: "forfeit", Valid: true})

	h.mu.RLock()
	opp, ok := h.clients[winner]
	h.mu.RUnlock()
	if ok {
		h.send(opp, protocol.MsgDuelUpdate, protocol.DuelUpdateMsg{
			DuelID: duel.ID, Status: "ended", Opponent: client.CharacterID,
			Reason: "forfeit", WasVictor: true,
		})
	}
	h.send(client, protocol.MsgDuelUpdate, protocol.DuelUpdateMsg{
		DuelID: duel.ID, Status: "ended", Opponent: winner, Reason: "forfeit",
	})
}

// sendPvPError pushes a machine-readable PvP rejection reason.
func (h *WorldHandler) sendPvPError(client *Client, reason string) {
	h.send(client, protocol.MsgPvPError, protocol.PvPErrorMsg{Reason: reason})
}

// buildPvPContext assembles the legality context for a character from the
// database and authoritative position. Zone lookup is supplied by the caller
// only when a zone matters (it is already known at the call site).
func (h *WorldHandler) buildPvPContext(charID, zoneID string) pvp.CharacterPvPContext {
	ctx := pvp.CharacterPvPContext{CharacterID: charID, ZoneID: zoneID}
	if cf, err := h.db.GetCharacterFaction(charID); err == nil && cf != nil {
		ctx.FactionID = cf.FactionID
		v := h.pvpFlagView(charID)
		ctx.Flag = v.State
	}
	if d, err := h.db.ActiveDuelFor(charID); err == nil && d != nil && d.Status == "active" {
		ctx.DuelID = d.ID
	}
	return ctx
}

// guildIDFor returns the character's guild ID, or "" when unguilded.
func (h *WorldHandler) guildIDFor(charID string) string {
	if g, err := h.db.GuildOf(charID); err == nil && g != nil {
		return g.ID
	}
	return ""
}

// logPvPDefeat records a pvp_events row for a PvP defeat (D-06: logged, never
// awarded).
func (h *WorldHandler) logPvPDefeat(attackerID, targetID string, basis pvp.PvPBasis, zoneID string) {
	id := h.db.NewRowID("pvpev")
	_ = h.db.RecordPvPEvent(id, "defeat", string(basis), attackerID, targetID,
		sql.NullString{String: zoneID, Valid: zoneID != ""})
}
