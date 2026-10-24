// Phase 8 battlefield zone entry/exit detection (R-01c, D-14). Zone membership
// is derived from authoritative server position only — never a client claim.
package handlers

import (
	"time"

	"tc-server/internal/pvp"
)

// updateBattlefieldPresence tracks which battlefield zone (if any) a client is
// standing in and applies D-14 for faction members:
//   - entering forces overt (and drops any pending covert),
//   - leaving refreshes the aggression timer,
//   - a zone change pushes the new state to the client.
//
// Zone membership is POSITIONAL, so it is tracked and pushed for every client
// — including neutrals, who are affected by the zone even though the zone rule
// never lets them fight there (D-14). Only the flag side effects are gated on
// faction membership.
func (h *WorldHandler) updateBattlefieldPresence(client *Client) {
	if client.CharacterID == "" {
		return
	}

	zoneID := h.zoneAt(client.Pos.Planet, client.Pos.X, client.Pos.Z)

	// Only faction members carry flags, so only they get flag side effects.
	cf, err := h.db.GetCharacterFaction(client.CharacterID)
	factioned := err == nil && cf != nil

	prev := client.pvpZoneID
	if zoneID == prev {
		h.sendPvPStateChange(client, zoneID)
		return
	}

	// The flag row carries no zone column in the proposal's schema, so zone
	// membership is session state recorded here and re-derived on login.
	client.pvpZoneID = zoneID

	if factioned {
		v := h.pvpFlagView(client.CharacterID)
		now := time.Now()
		if zoneID != "" && prev == "" {
			// Entering: force overt (D-14), cancel any pending covert.
			v.State = pvp.FlagOvert
			v.Pending = ""
			v.ReadyAt = time.Time{}
			_ = h.savePvPFlag(client.CharacterID, v)
		} else if zoneID == "" && prev != "" {
			// Leaving: aggression timer applies before covert is allowed (D-14).
			_ = h.savePvPFlag(client.CharacterID, v.RefreshAggro(now))
		}
	}

	h.sendPvPStateChange(client, client.pvpZoneID)
}

// pvpZoneID tracks the battlefield zone a client is currently inside, derived
// from authoritative position. Kept out of Client's struct documentation on
// purpose: it is session state, not persisted design state.
// (Declared as a field addition on Client in world.go.)
