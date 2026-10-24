// Phase 7/8 entity-spawn faction extension (proposal Section 7.3). Attaches
// faction_id and pvp_state to player spawns only; creature spawns stay blank.
package handlers

import "tc-server/internal/protocol"

// withFactionFields enriches an EntitySpawnMsg with faction_id and pvp_state
// when the entity is a player carrying a faction. Creatures and neutrals stay
// unchanged, so older clients that ignore unknown fields keep working.
func (h *WorldHandler) withFactionFields(msg protocol.EntitySpawnMsg, charID string) protocol.EntitySpawnMsg {
	if msg.EntityType == "creature" || msg.EntityType == "" {
		return msg
	}
	if cf, err := h.db.GetCharacterFaction(charID); err == nil && cf != nil {
		if f, err := h.db.GetFactionByID(cf.FactionID); err == nil && f != nil {
			msg.FactionID = cf.FactionID
		}
	}
	if fl, err := h.db.GetPvPFlag(charID); err == nil && fl != nil {
		msg.PvPState = fl.State
	}
	return msg
}
