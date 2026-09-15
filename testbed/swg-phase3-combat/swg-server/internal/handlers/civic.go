// Civic handlers for Phase 7: city founding (WS placement), membership-routed
// chat, group/guild presence helpers, and the city simulation tick
// (activation, upkeep, rank evaluation, elections, flat-tax collection).
// Testbed fork; generic content only. GDD-silent numbers [PROVISIONAL] per the
// owner-approved convention; GDD-given values cited.
package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"swg-server/internal/civic"
	"swg-server/internal/database"
	"swg-server/internal/protocol"
	"swg-server/internal/world"
)

// SeedCivicWorld prepares Phase 7 persistence. Idempotent.
func (h *WorldHandler) SeedCivicWorld() error {
	if err := h.db.EnsureCivicSchema(); err != nil {
		return err
	}
	if err := h.db.EnsureCivicSocialSchema(); err != nil {
		return err
	}
	log.Printf("Civic world seeded: cities, guilds, groups, mail, social")
	return nil
}

// isOnline reports world presence.
func (h *WorldHandler) isOnline(characterID string) bool {
	_, ok := h.findClient(characterID)
	return ok
}

// --- V1: city founding (WS positional placement, mirrors handlePlaceHouse) ---

func (h *WorldHandler) handlePlaceCityHall(client *Client, raw []byte) {
	if client.CharacterID == "" {
		h.sendError(client, "not in world")
		return
	}
	var msg protocol.WSMessage
	json.Unmarshal(raw, &msg)
	dataBytes, _ := json.Marshal(msg.Data)
	var pm protocol.PlaceCityHallMsg
	if err := json.Unmarshal(dataBytes, &pm); err != nil {
		h.sendError(client, "invalid place_city_hall data")
		return
	}
	if pm.Name == "" || len(pm.Name) > 40 {
		h.sendError(client, "city name required (max 40 characters)")
		return
	}
	deed, err := h.db.GetItem(client.CharacterID, pm.DeedItemID)
	if err != nil || deed.Schematic != "city_hall_deed" {
		h.sendError(client, "valid city hall deed required")
		return
	}
	if err := h.checkNoBuild(client.Pos.Planet, pm.X, pm.Z); err != nil {
		h.sendError(client, err.Error())
		return
	}
	if err := h.checkCityOverlap(client.Pos.Planet, pm.X, pm.Z, civic.CityRadiusM); err != nil {
		h.sendError(client, err.Error())
		return
	}
	hallID := fmt.Sprintf("st-%d", time.Now().UnixNano())
	if err := h.db.PlaceStructure(hallID, client.CharacterID, client.Pos.Planet,
		pm.X, pm.Z, "city_hall", "", 0); err != nil {
		h.sendError(client, "city hall placement failed")
		return
	}
	_ = h.db.DeleteItem(client.CharacterID, pm.DeedItemID)
	now := time.Now().Unix()
	cityID := fmt.Sprintf("city-%d", time.Now().UnixNano())
	base := civic.CityUpkeepWeekly[civic.RankOutpost]
	if err := h.db.CreateCity(cityID, pm.Name, client.Pos.Planet, pm.X, pm.Z,
		civic.CityRadiusM, base, client.CharacterID,
		now+civic.Secs(civic.MayorTerm), now+civic.Secs(civic.FormingGrace)); err != nil {
		h.sendError(client, "city founding failed")
		return
	}
	// The founder ("Mayor-Founder", GDD 14.2.1) starts as mayor + citizen.
	_ = h.db.AddCitizen(cityID, client.CharacterID)
	_ = h.db.TouchPresence(client.CharacterID, now)
	n, _ := h.db.CountStructuresInCity(client.Pos.Planet, pm.X, pm.Z, civic.CityRadiusM)
	status := civic.CityForming
	if n >= civic.FoundThreshold {
		h.activateCity(cityID)
		status = civic.CityActive
	}
	h.send(client, protocol.MsgCityFounded, protocol.CityFoundedMsg{
		CityID: cityID, Status: status, Structures: n, Threshold: civic.FoundThreshold,
	})
}

// checkCityOverlap rejects a hall whose radius would overlap a forming/active
// city (GDD 14.5 annexation-race guard).
func (h *WorldHandler) checkCityOverlap(zone string, x, z, radius float64) error {
	all, err := h.db.AllCities()
	if err != nil {
		return err
	}
	for _, c := range all {
		if c.Zone != zone || c.Status == civic.CityDissolved {
			continue
		}
		dx, dz := x-c.CenterX, z-c.CenterZ
		if dx*dx+dz*dz < (radius+c.RadiusM)*(radius+c.RadiusM) {
			return fmt.Errorf("too close to another city")
		}
	}
	return nil
}

// activateCity flips forming → active and auto-enrols in-radius owners
// (GDD 14.2.1: auto-enrol with opt-out).
func (h *WorldHandler) activateCity(cityID string) {
	c, err := h.db.GetCity(cityID)
	if err != nil {
		return
	}
	c.Status = civic.CityActive
	c.GraceEnds = sql.NullInt64{}
	_ = h.db.UpdateCity(c)
	h.enrolInRadius(c)
}

// enrolInRadius auto-enrols owners of in-radius structures (idempotent;
// opt-outs are never re-enrolled — GDD 14.2.1).
func (h *WorldHandler) enrolInRadius(c *database.CityRow) {
	structs, err := h.db.StructuresInCity(c.Zone, c.CenterX, c.CenterZ, c.RadiusM)
	if err != nil {
		return
	}
	for _, s := range structs {
		if h.db.IsOptedOut(c.ID, s.OwnerCharacterID) {
			continue
		}
		_ = h.db.AddCitizen(c.ID, s.OwnerCharacterID)
	}
}

// --- G2: membership-routed chat (called from handleChat) ---

func (h *WorldHandler) handleCivicChat(client *Client, chatMsg protocol.ChatMsg) {
	if chatMsg.Text == "" {
		h.sendError(client, "empty chat text")
		return
	}
	out := protocol.ChatMessageMsg{
		SenderName: client.Name, Channel: chatMsg.Channel, Text: chatMsg.Text,
	}
	switch world.ChatChannel(chatMsg.Channel) {
	case world.ChannelGroup:
		g, err := h.db.GroupOf(client.CharacterID)
		if err != nil {
			h.sendError(client, "not in a group")
			return
		}
		members, _ := h.db.GroupMembers(g.ID)
		h.send(client, protocol.MsgChatMessage, out) // echo
		for _, id := range members {
			if id == client.CharacterID {
				continue
			}
			if other, ok := h.findClient(id); ok {
				h.send(other, protocol.MsgChatMessage, out)
			}
		}
	case world.ChannelGuild:
		g, err := h.db.GuildOf(client.CharacterID)
		if err != nil {
			h.sendError(client, "not in a guild")
			return
		}
		members, _ := h.db.GuildMembers(g.ID)
		h.send(client, protocol.MsgChatMessage, out) // echo
		for _, m := range members {
			if m.CharacterID == client.CharacterID {
				continue
			}
			if other, ok := h.findClient(m.CharacterID); ok {
				h.send(other, protocol.MsgChatMessage, out)
			}
		}
	case world.ChannelTell:
		if chatMsg.Target == "" {
			h.sendError(client, "tell target required")
			return
		}
		targetID, err := h.db.CharacterIDByName(chatMsg.Target)
		if err != nil {
			h.sendError(client, "tell recipient unknown")
			return
		}
		other, ok := h.findClient(targetID)
		if !ok {
			// GDD 18.5: tell to an offline character errors clearly, never
			// silently fails (mail is the offline path).
			h.sendError(client, "recipient not in world (use mail)")
			return
		}
		h.send(other, protocol.MsgChatMessage, out)
		h.send(client, protocol.MsgChatMessage, out) // echo
	case world.ChannelFaction:
		// No faction/alignment system exists before Phase 8: the channel
		// exists in vocabulary but has no membership to route to.
		h.sendError(client, "faction alignment unavailable (later phase)")
	default:
		h.sendError(client, "unknown chat channel")
	}
}

// --- City simulation tick (driven from the resource tick loop, 1 s cadence;
// maintenance/upkeep accrue per harvest interval like tickStructures) ---

// tickCities runs each harvest interval: activation/grace, mayor
// term/inactivity, upkeep, rank evaluation, flat-tax collection.
func (h *WorldHandler) tickCities() {
	all, err := h.db.AllCities()
	if err != nil {
		log.Printf("city tick: list failed: %v", err)
		return
	}
	now := time.Now().Unix()
	for i := range all {
		c := &all[i]
		if c.Status == civic.CityDissolved {
			continue
		}
		if c.Status == civic.CityForming {
			h.tickForming(c, now)
			continue
		}
		h.tickMayor(c, now)
		h.tickUpkeep(c)
		h.tickRanks(c)
		h.tickFlatTax(c)
		h.tickElections(c, now)
	}
}

// tickForming activates on threshold or dissolves past grace.
func (h *WorldHandler) tickForming(c *database.CityRow, now int64) {
	n, err := h.db.CountStructuresInCity(c.Zone, c.CenterX, c.CenterZ, c.RadiusM)
	if err != nil {
		return
	}
	if n >= civic.FoundThreshold {
		h.activateCity(c.ID)
		return
	}
	if c.GraceEnds.Valid && now >= c.GraceEnds.Int64 {
		h.dissolveCity(c, "forming grace expired")
	}
}

// hasActiveMayor reports mayor set + term valid + still a citizen.
func (h *WorldHandler) hasActiveMayor(c *database.CityRow, now int64) bool {
	if !c.MayorID.Valid || c.MayorID.String == "" {
		return false
	}
	if c.TermEnd.Valid && now >= c.TermEnd.Int64 {
		return false
	}
	return h.db.IsCitizen(c.ID, c.MayorID.String)
}

// tickMayor handles term expiry and inactivity emergency elections.
func (h *WorldHandler) tickMayor(c *database.CityRow, now int64) {
	if !c.MayorID.Valid || c.MayorID.String == "" {
		h.ensureElection(c, now)
		return
	}
	expired := c.TermEnd.Valid && now >= c.TermEnd.Int64
	departed := !h.db.IsCitizen(c.ID, c.MayorID.String)
	inactive := false
	if seen := h.db.LastSeen(c.MayorID.String); seen > 0 {
		inactive = now-seen >= civic.Secs(civic.MayorInactivity)
	}
	if expired || departed || inactive {
		reason := "term expired"
		if departed {
			reason = "mayor left city"
		}
		if inactive {
			reason = "mayor inactive"
		}
		log.Printf("city %s: mayor vacancy (%s), opening emergency election", c.ID, reason)
		c.MayorID = sql.NullString{}
		c.TermEnd = sql.NullInt64{}
		_ = h.db.UpdateCity(c)
		h.ensureElection(c, now)
	}
}

// ensureElection opens an election when none is open (cooldown-gated).
func (h *WorldHandler) ensureElection(c *database.CityRow, now int64) {
	if _, err := h.db.GetOpenElection(c.ID); err == nil {
		return
	}
	if now-h.db.LastElectionEnd(c.ID) < civic.Secs(civic.ElectionCooldown) {
		return
	}
	_ = h.db.OpenElection(fmt.Sprintf("elec-%d", time.Now().UnixNano()),
		c.ID, now+civic.Secs(civic.ElectionPeriod))
}

// tickUpkeep deducts pro-rated weekly upkeep (mayor-discounted) from the
// treasury; shortfall downgrades, or dissolves a Rank-1 city (GDD 14.2.4 —
// citizen property is untouched: only the city row + citizenship dissolve).
func (h *WorldHandler) tickUpkeep(c *database.CityRow) {
	now := time.Now().Unix()
	fee := (civic.EffectiveUpkeep(c.Rank, h.hasActiveMayor(c, now)) + 167) / 168
	if fee <= 0 {
		return
	}
	if c.Treasury >= fee {
		c.Treasury -= fee
		_ = h.db.UpdateCity(c)
		_ = h.db.RecordLedger("city:"+c.ID, -fee, "city_upkeep", "sink", c.ID, c.Zone)
		return
	}
	if c.Rank == civic.RankOutpost {
		h.dissolveCity(c, "treasury exhausted")
		return
	}
	c.Rank = rankDown(c.Rank)
	c.UpkeepWeekly = civic.CityUpkeepWeekly[c.Rank]
	_ = h.db.UpdateCity(c)
	log.Printf("city %s: downgraded to %s (treasury shortfall)", c.ID, c.Rank)
}

func rankDown(rank string) string {
	switch rank {
	case civic.RankCity:
		return civic.RankTownship
	default:
		return civic.RankOutpost
	}
}

// dissolveCity marks dissolved and clears citizenship + open elections.
// Structures revert to standalone-owned (rows untouched — ownership stands).
func (h *WorldHandler) dissolveCity(c *database.CityRow, reason string) {
	log.Printf("city %s: dissolved (%s)", c.ID, reason)
	c.Status = civic.CityDissolved
	c.MayorID = sql.NullString{}
	c.TermEnd = sql.NullInt64{}
	_ = h.db.UpdateCity(c)
	citizens, _ := h.db.CitizensOf(c.ID)
	for _, cz := range citizens {
		_ = h.db.RemoveCitizen(c.ID, cz.CharacterID)
	}
	if e, err := h.db.GetOpenElection(c.ID); err == nil {
		_ = h.db.CloseElection(e.ID, "")
	}
}

// tickRanks promotes when structure count + treasury reserve (+ active mayor
// for Rank city) hold. Rank 4+ is never entered (owner decision 4).
func (h *WorldHandler) tickRanks(c *database.CityRow) {
	n, err := h.db.CountStructuresInCity(c.Zone, c.CenterX, c.CenterZ, c.RadiusM)
	if err != nil {
		return
	}
	now := time.Now().Unix()
	switch c.Rank {
	case civic.RankOutpost:
		need := civic.TreasuryReserveMult * civic.CityUpkeepWeekly[civic.RankTownship]
		if n >= civic.TownshipThreshold && c.Treasury >= need {
			c.Rank = civic.RankTownship
			c.UpkeepWeekly = civic.CityUpkeepWeekly[c.Rank]
			_ = h.db.UpdateCity(c)
			log.Printf("city %s: promoted to township", c.ID)
		}
	case civic.RankTownship:
		need := civic.TreasuryReserveMult * civic.CityUpkeepWeekly[civic.RankCity]
		if n >= civic.CityThreshold && c.Treasury >= need && h.hasActiveMayor(c, now) {
			c.Rank = civic.RankCity
			c.UpkeepWeekly = civic.CityUpkeepWeekly[c.Rank]
			_ = h.db.UpdateCity(c)
			log.Printf("city %s: promoted to city", c.ID)
		}
	}
	// Ongoing auto-enrol keeps citizenship aligned with in-radius owners.
	h.enrolInRadius(c)
}

// tickFlatTax collects the mayor-set flat citizen fee pro-rated per interval:
// solvent citizens pay into the treasury (ledger transfer pair); insolvent
// citizens accrue arrears (eviction evidence for the mayor).
func (h *WorldHandler) tickFlatTax(c *database.CityRow) {
	if c.TaxFlatWeekly <= 0 {
		return
	}
	share := (c.TaxFlatWeekly + 167) / 168
	if share <= 0 {
		return
	}
	citizens, err := h.db.CitizensOf(c.ID)
	if err != nil {
		return
	}
	for _, cz := range citizens {
		bal, err := h.db.GetCharacterCredits(cz.CharacterID)
		if err != nil || bal < share {
			_ = h.db.AddArrears(c.ID, cz.CharacterID, share)
			continue
		}
		if err := h.db.DeductCredits(cz.CharacterID, share); err != nil {
			_ = h.db.AddArrears(c.ID, cz.CharacterID, share)
			continue
		}
		c.Treasury += share
		_ = h.db.RecordLedger(cz.CharacterID, -share,
			"city_flat_tax", "transfer", "city:"+c.ID, c.Zone)
		_ = h.db.RecordLedger("city:"+c.ID, share,
			"city_flat_tax", "transfer", cz.CharacterID, c.Zone)
	}
	_ = h.db.UpdateCity(c)
}

// tickElections closes elections on unanimity (flagged test acceleration —
// GDD gives a weekly period with no early-close rule; defaults still close on
// period expiry) or period expiry; winner = most votes, tie → earliest ballot.
func (h *WorldHandler) tickElections(c *database.CityRow, now int64) {
	e, err := h.db.GetOpenElection(c.ID)
	if err != nil {
		return
	}
	ballots, err := h.db.BallotsOf(e.ID)
	if err != nil {
		return
	}
	citizens, err := h.db.CitizensOf(c.ID)
	if err != nil {
		return
	}
	closed := now >= e.EndsAt
	if len(citizens) > 0 && len(ballots) >= len(citizens) {
		closed = true // every citizen voted: no outcome left to await
	}
	if !closed {
		return
	}
	tally := map[string]int{}
	firstSeen := map[string]int{}
	for i, b := range ballots {
		tally[b.CandidateID]++
		if _, ok := firstSeen[b.CandidateID]; !ok {
			firstSeen[b.CandidateID] = i
		}
	}
	winner, best := "", -1
	for cand, votes := range tally {
		if votes > best || (votes == best && firstSeen[cand] < firstSeen[winner]) {
			winner, best = cand, votes
		}
	}
	if winner == "" || !h.db.IsCitizen(c.ID, winner) {
		// No valid winner (empty ballot or departed candidate): vacancy stands.
		_ = h.db.CloseElection(e.ID, "")
		return
	}
	_ = h.db.CloseElection(e.ID, winner)
	c.MayorID = sql.NullString{String: winner, Valid: true}
	c.TermEnd = sql.NullInt64{Int64: now + civic.Secs(civic.MayorTerm), Valid: true}
	_ = h.db.UpdateCity(c)
	log.Printf("city %s: elected mayor %s (%d votes)", c.ID, winner, best)
}

// CityUnlocks derives rank-gated convenience flags (GDD 14.2.2/14.6: flags are
// server-side unlocks; shuttleport TRAVEL movement stays later-phase scope).
func CityUnlocks(rank string) map[string]bool {
	return map[string]bool{
		"cantina_allowed":  rank == civic.RankTownship || rank == civic.RankCity,
		"medical_allowed":  rank == civic.RankTownship || rank == civic.RankCity,
		"shuttleport":      rank == civic.RankCity,
		"trainer_hosting":  rank == civic.RankCity,
		"clone_facility":   false, // Rank 4 Metropolis shape, unreachable
		"specialization":   false, // Rank 4 Expansion, unreachable
	}
}


