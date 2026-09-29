// Ranger handlers for Phase 9: camp deployment + expiry sweep, creature
// tracking, and pet ticks. Testbed fork; generic content only. GDD-silent
// numbers [PROVISIONAL] per the owner-approved convention; GDD-given cited.
package handlers

import (
	"encoding/json"
	"math"
	"os"
	"time"

	"swg-server/internal/database"
	"swg-server/internal/protocol"
)

// Camp tuning (GDD 25.2.2: 30–60 min life [ASSUMPTION]; fast-cycle 5 min).
var CampLife = 30 * time.Minute

// CampRadiusM: buff radius [PROVISIONAL — GDD gives shape without metres].
const CampRadiusM = 20.0

// Camp regen ticks (provisional §6.8-adjacent magnitudes, gate-visible):
// Ranger-deployed camps restore more than Scout-deployed (reduced
// effectiveness, GDD 25.2.2).
const (
	CampRegenRanger = 50
	CampRegenScout  = 20
	CampXPPerMember = 10 // ranger XP per occupied camp per harvest tick
)

func init() {
	if os.Getenv("TESTBED_FAST_CYCLE") == "1" {
		CampLife = 5 * time.Minute
	}
}

// --- Camp deployment (WS positional, wilderness-only) ---

func (h *WorldHandler) handleDeployCamp(client *Client, raw []byte) {
	if client.CharacterID == "" {
		h.sendError(client, "not in world")
		return
	}
	rangerTier := h.skillTier(client.CharacterID, "ranger_wilderness_survival_")
	scoutTier := h.skillTier(client.CharacterID, "scout_survival_")
	full := rangerTier >= 1
	if !full && scoutTier < 1 {
		h.sendError(client, "requires Scout Survival or Ranger")
		return
	}
	// Wilderness-only (GDD 25.2.2 — not a persistent Structure).
	if _, err := h.db.CityContaining(client.Pos.Planet, client.Pos.X, client.Pos.Z); err == nil {
		h.sendError(client, "camps cannot be placed inside cities")
		return
	}
	now := time.Now().Unix()
	id, err := h.db.DeployCamp(client.CharacterID, client.Pos.Planet,
		client.Pos.X, client.Pos.Z, CampRadiusM, now+int64(CampLife/time.Second))
	if err != nil {
		h.sendError(client, "camp deployment failed")
		return
	}
	h.send(client, protocol.MsgCampDeployed, protocol.CampDeployedMsg{CampID: id})
}

// tickCamps runs each harvest interval: expiry sweep + occupant regen +
// camping XP (GDD 7.2.4 "Camping (XP per camp use by group members)" —
// awarded to the deployer per occupant present, provisional rate).
func (h *WorldHandler) tickCamps() {
	now := time.Now().Unix()
	for _, c := range expiredCamps(h) {
		_ = h.db.DeleteCamp(c.ID)
	}
	camps, err := h.db.ActiveCamps("zone-0001", now)
	if err != nil {
		return
	}
	for i := range camps {
		c := &camps[i]
		full := h.skillTier(c.DeployedBy, "ranger_wilderness_survival_") >= 1
		regen := CampRegenScout
		if full {
			regen = CampRegenRanger
		}
		for _, cl := range clientsInRadius(h, c.Zone, c.PosX, c.PosZ, c.RadiusM) {
			st, err := h.db.GetCombatState(cl)
			if err != nil || st.IncapacitatedAt != nil {
				continue
			}
			st.HealthCurrent += regen
			if st.HealthCurrent > st.HealthMax {
				st.HealthCurrent = st.HealthMax
			}
			st.ActionCurrent += regen
			if st.ActionCurrent > st.ActionMax {
				st.ActionCurrent = st.ActionMax
			}
			_ = h.db.UpdateCombatState(st)
			h.pushHAMUpdate(cl)
		}
		_ = h.db.AddCharacterXP(c.DeployedBy, "ranger", CampXPPerMember)
	}
}

func expiredCamps(h *WorldHandler) []database.CampRow {
	out, _ := h.db.ExpiredCamps(time.Now().Unix())
	return out
}

func clientsInRadius(h *WorldHandler, zone string, x, z, radius float64) []string {
	h.mu.RLock()
	defer h.mu.RUnlock()
	var out []string
	for _, c := range h.clients {
		if c.CharacterID == "" || c.Pos.Planet != zone {
			continue
		}
		dx, dz := c.Pos.X-x, c.Pos.Z-z
		if dx*dx+dz*dz <= radius*radius {
			out = append(out, c.CharacterID)
		}
	}
	return out
}

// --- Creature tracking (WS; Ranger + Scout Hunting tiers) ---

func (h *WorldHandler) handleTrack(client *Client, raw []byte) {
	if client.CharacterID == "" {
		h.sendError(client, "not in world")
		return
	}
	var msg protocol.WSMessage
	json.Unmarshal(raw, &msg)
	dataBytes, _ := json.Marshal(msg.Data)
	var tm protocol.TrackMsg
	if err := json.Unmarshal(dataBytes, &tm); err != nil {
		h.sendError(client, "invalid track data")
		return
	}
	hunting := h.skillTier(client.CharacterID, "scout_hunting_")
	ranger := h.skillTier(client.CharacterID, "ranger_tracking_")
	if hunting < 2 && ranger < 1 {
		h.sendError(client, "requires Scout Hunting II or Ranger Tracking")
		return
	}
	rows, err := h.db.GetLivingCreatureInstances()
	if err != nil {
		h.sendError(client, "tracking failed")
		return
	}
	var best *database.CreatureInstanceRow
	bestD := math.MaxFloat64
	for i := range rows {
		if tm.TemplateID != "" && rows[i].TemplateID != tm.TemplateID {
			continue
		}
		dx := client.Pos.X - rows[i].PosX
		dz := client.Pos.Z - rows[i].PosZ
		if rows[i].Zone != client.Pos.Planet {
			continue
		}
		if d := dx*dx + dz*dz; d < bestD {
			bestD = d
			best = &rows[i]
		}
	}
	if best == nil {
		h.sendError(client, "no trail found")
		return
	}
	pool := "scouting"
	if ranger >= 1 {
		pool = "ranger"
	}
	_ = h.db.AddCharacterXP(client.CharacterID, pool, 25) // [PROVISIONAL]
	h.send(client, protocol.MsgSurveyResult, protocol.SurveyResultMsg{
		SpawnID:       best.ID,
		ResourceType:  "creature:" + best.TemplateID,
		DistanceM:     math.Sqrt(bestD),
		Concentration: 100,
		Waypoint:      protocol.WaypointMsg{X: best.PosX, Z: best.PosZ},
		Refined:       true,
	})
}

// --- Pet ticks (1 s cadence, driven from StartTickLoop) ---

// tickPets runs every tick: follow movement + attack resolution for pets in
// attack mode. Pets are invulnerable and non-targetable in MVP (flagged;
// pet durability design is human-directed).
func (h *WorldHandler) tickPets() {
	h.mu.RLock()
	pets := make([]PetRuntime, 0, len(h.livePets))
	for _, p := range h.livePets {
		pets = append(pets, *p)
	}
	h.mu.RUnlock()
	for i := range pets {
		p := &pets[i]
		owner, ok := h.findClient(p.OwnerID)
		if !ok {
			continue
		}
		prow, err := h.db.GetPet(p.ID)
		if err != nil {
			continue
		}
		switch prow.Mode {
		case "follow":
			dx := owner.Pos.X - p.PosX
			dz := owner.Pos.Z - p.PosZ
			if dx*dx+dz*dz > PetFollowRadiusM*PetFollowRadiusM {
				dist := math.Sqrt(dx*dx + dz*dz)
				step := PetStepM
				if dist < step {
					step = dist
				}
				p.PosX += dx / dist * step
				p.PosZ += dz / dist * step
				h.movePet(p)
			}
		case "attack":
			h.tickPetAttack(p, prow)
		}
	}
}

// movePet persists + republishes pet movement.
func (h *WorldHandler) movePet(p *PetRuntime) {
	_ = h.db.UpdatePetPos(p.ID, p.Zone, p.PosX, p.PosZ)
	h.mu.Lock()
	if live, ok := h.livePets[p.ID]; ok {
		live.PosX, live.PosZ = p.PosX, p.PosZ
	}
	h.mu.Unlock()
}

// tickPetAttack damages the pet's target instance (creatures only — PvP pets
// need targeting-validity design, flagged OUT).
func (h *WorldHandler) tickPetAttack(p *PetRuntime, prow *database.PetRow) {
	rows, err := h.db.GetLivingCreatureInstances()
	if err != nil {
		return
	}
	var tgt *database.CreatureInstanceRow
	for i := range rows {
		if rows[i].ID == prow.TargetID {
			tgt = &rows[i]
			break
		}
	}
	if tgt == nil {
		_ = h.db.SetPetCommand(p.ID, "stay", "")
		return
	}
	dx := tgt.PosX - p.PosX
	dz := tgt.PosZ - p.PosZ
	if tgt.Zone != p.Zone || dx*dx+dz*dz > PetSightM*PetSightM {
		return
	}
	dmg := petDamage(p)
	tgt.HealthCurrent -= dmg
	if tgt.HealthCurrent <= 0 {
		_ = h.db.KillCreatureInstance(tgt.ID, time.Now().Add(10*time.Minute))
		_ = h.db.AddCharacterXP(p.OwnerID, "combat", 50) // [PROVISIONAL flat]
		_ = h.db.SetPetCommand(p.ID, "stay", "")
		return
	}
	_ = h.db.UpdateCreatureInstance(tgt)
}

// petDamage is a flat roll with a quality kicker (provisional §6.8-add:
// +1 per 100 quality — gate-visible; template-scaled derivation deferred).
func petDamage(p *PetRuntime) int {
	return 5 + p.Quality/100
}