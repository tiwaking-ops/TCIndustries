// Phase 8 battlefield zone entry/exit detection (R-01c, D-14). Zone membership
// is derived from authoritative server position only — never a client claim.
package handlers

import (
	"time"

	"tc-server/internal/pvp"
)

// updateBattlefieldPresence compares the client's current zone against the
// zone recorded on its flag row and applies D-14:
//   - entering forces overt (and drops any pending covert),
//   - leaving refreshes the aggression timer,
//   - a zone change pushes the new state to the client.
func (h *WorldHandler) updateBattlefieldPresence(client *Client) {
	if client.CharacterID == "" {
		return
	}
	// Only faction members carry flags and zone state.
	cf, err := h.db.GetCharacterFaction(client.CharacterID)
	if err != nil || cf == nil {
		return
	}

	zoneID := h.zoneAt(client.Pos.Planet, client.Pos.X, client.Pos.Z)
	v := h.pvpFlagView(client.CharacterID)

	// The flag row itself carries no zone column in the proposal's schema, so
	// membership is tracked per-session and pushed; persistence of the zone is
	// derived on load from the last known position. Record it on the client.
	prev := client.pvpZoneID
	if zoneID == prev {
		h.sendPvPStateChange(client, zoneID)
		return
	}

	now := time.Now()
	if zoneID != "" && prev == "" {
		// Entering: force overt (D-14), cancel any pending covert.
		v.State = pvp.FlagOvert
		v.Pending = ""
		v.ReadyAt = time.Time{}
		if err := h.savePvPFlag(client.CharacterID, v); err == nil {
			client.pvpZoneID = zoneID
		}
	} else if zoneID == "" && prev != "" {
		// Leaving: aggression timer applies before covert is allowed (D-14).
		v = v.RefreshAggro(now)
		if err := h.savePvPFlag(client.CharacterID, v); err == nil {
			client.pvpZoneID = ""
		}
	} else {
		client.pvpZoneID = zoneID
	}

	h.sendPvPStateChange(client, client.pvpZoneID)
}

// pvpZoneID tracks the battlefield zone a client is currently inside, derived
// from authoritative position. Kept out of Client's struct documentation on
// purpose: it is session state, not persisted design state.
// (Declared as a field addition on Client in world.go.)
