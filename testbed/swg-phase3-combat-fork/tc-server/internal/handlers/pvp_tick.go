// Phase 8 background tick processing: flag transitions, aggression timer
// expiry, and guild war activation (D-11). Driven by StartTickLoop in combat.go.
package handlers

import (
	"log"
	"time"
)

// tickPvP processes background PvP state changes each second:
//  1. Promotes pending flag transitions whose ready_at timestamp has passed.
//  2. Clears expired aggression timers.
//  3. Activates guild wars whose active_from timestamp has passed (D-11).
func (h *WorldHandler) tickPvP() {
	now := time.Now()

	// 3. Activate due guild wars (D-11).
	if n, err := h.db.ActivateDueWars(now); err == nil && n > 0 {
		log.Printf("pvp: activated %d guild war(s)", n)
	}

	// 1 & 2. Process online clients for flag transitions and aggro expiry.
	h.mu.RLock()
	clients := make([]*Client, 0, len(h.clients))
	for _, c := range h.clients {
		if c.CharacterID != "" {
			clients = append(clients, c)
		}
	}
	h.mu.RUnlock()

	for _, client := range clients {
		row, err := h.db.GetPvPFlag(client.CharacterID)
		if err != nil || row == nil {
			continue
		}
		v := h.pvpFlagView(client.CharacterID)
		changed := false

		// Promote due pending state transition.
		if v.HasPending() && !v.ReadyAt.IsZero() && !now.Before(v.ReadyAt) {
			v.State = v.Pending
			v.Pending = ""
			v.ReadyAt = time.Time{}
			changed = true
		}

		// Expire aggression timer.
		if v.AggroActive(now) && now.After(v.AggressionExpiresAt) {
			v.AggressionExpiresAt = time.Time{}
			changed = true
		}

		if changed {
			_ = h.savePvPFlag(client.CharacterID, v)
			h.sendPvPStateChange(client, h.zoneAt(client.Pos.Planet, client.Pos.X, client.Pos.Z))
		}
	}
}
