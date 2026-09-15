// Resource world handlers for Phase 4: survey, sample, harvesters, corpse
// harvest, spawn lifecycle + harvest ticks, and zone seeding. Testbed fork;
// generic content only. All GDD-silent numbers are [PROVISIONAL] placeholders
// under the owner-approved convention (proposal §6); GDD-given values are cited.
package handlers

import (
	"encoding/json"
	"log"
	"math"
	"math/rand"
	"sync"
	"time"

	"swg-server/internal/database"
	"swg-server/internal/protocol"
	"swg-server/internal/resources"
)

// Survey/action ranges (provisional where the GDD is silent):
const (
	// SurveyRadiusM: fixed search radius (GDD has a 5–100 km player slider;
	// single-zone MVP uses a fixed radius — flagged simplification).
	SurveyRadiusM = 500.0
	// SampleRangeM: must stand near the spawn (GDD: "stand at location").
	SampleRangeM = 15.0
	// EmptyRangeM: must be near the harvester to empty it (GDD silent).
	EmptyRangeM = 15.0
	// CorpseHarvestRangeM: sample-tool interaction range (GDD silent).
	CorpseHarvestRangeM = 5.0
	// HarvesterExclusionM: GDD-given 100 m no-clustering rule.
	HarvesterExclusionM = 100.0
)

// corpseYieldType: fixed yield type per creature template (GDD 17.2.2 shape).
// [PROVISIONAL] generic assignments, flagged.
var corpseYieldType = map[string]string{
	"0001": "cultured_organic",
	"0002": "fibrous_flora",
}

var (
	surveyMu       sync.Mutex
	lastSurveyByCh = make(map[string]time.Time)
)

// SeedResourceWorld prepares Phase 4 persistence + zone spawns. Idempotent.
func (h *WorldHandler) SeedResourceWorld() error {
	if err := h.db.EnsureResourceSchema(); err != nil {
		return err
	}
	// Fresh RNG per call (see combat.go note: shared *rand.Rand raced).
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	count, err := h.db.SeedZoneSpawns("zone-0001", rng, time.Now())
	if err != nil {
		return err
	}
	log.Printf("Resource world seeded: %d active spawns in zone-0001", count)
	return nil
}

// StartResourceTickLoop drives spawn lifecycle + harvester ticks (1 s cadence;
// spawn checks every 5th tick; harvest work when the harvest interval elapsed).
func (h *WorldHandler) StartResourceTickLoop() {
	ticker := time.NewTicker(1 * time.Second)
	var tickCount int64
	lastHarvest := time.Now()
	go func() {
		for range ticker.C {
			tickCount++
			if tickCount%5 == 0 {
				rng := rand.New(rand.NewSource(time.Now().UnixNano()))
				if _, _, err := h.db.TickSpawns("zone-0001", rng, time.Now()); err != nil {
					log.Printf("spawn tick failed: %v", err)
				}
			}
			if time.Since(lastHarvest) >= resources.HarvestTickInterval {
				lastHarvest = time.Now()
				h.tickHarvesters()
				h.tickStructures() // Phase 5: structure maintenance + grace/destroy
				h.tickCities()      // Phase 7: activation, upkeep, ranks, elections, flat tax
				if _, err := h.db.WriteSnapshot(); err != nil {
					log.Printf("snapshot failed: %v", err)
				}
			}
		}
	}()
}

// tickHarvesters runs one harvest interval (= one game-hour by construction, so
// rate × concentration applies directly) for every harvester: funded + bound to
// a live spawn → extract into the hopper (capped, no overflow per GDD 11.4.2);
// deduct the pro-rated weekly fee (ceiled per tick — flagged discretization);
// unfunded or spawn-dead harvesters go inactive.
func (h *WorldHandler) tickHarvesters() {
	all, err := h.db.AllHarvesters()
	if err != nil {
		log.Printf("harvest tick: list failed: %v", err)
		return
	}
	for _, hv := range all {
		kind, ok := database.HarvesterKinds[hv.Kind]
		if !ok {
			continue
		}
		if hv.MaintenancePool <= 0 {
			if hv.Active {
				hv.Active = false
				_ = h.db.UpdateHarvester(&hv)
			}
			continue
		}
		spawn, err := h.db.GetSpawn(hv.SpawnID)
		if err != nil || !spawn.Active {
			continue // spawn despawned: extracts nothing (GDD 11.4.3)
		}
		conc := database.EffectiveConcentration(spawn, hv.PosX, hv.PosZ)
		units := int(float64(kind.RatePerHour) * float64(conc) / 100.0)
		if hopperSpace := kind.HopperCap - hv.HopperUnits; hopperSpace < units {
			units = hopperSpace // no overflow (GDD 11.4.2)
		}
		if units < 0 {
			units = 0
		}
		fee := (kind.WeeklyFee + 167) / 168
		hv.MaintenancePool -= fee
		if hv.MaintenancePool < 0 {
			hv.MaintenancePool = 0
		}
		// Phase 5 ledger: maintenance-fee sink (rate UNCHANGED — retro-tagged).
		_ = h.db.RecordLedger(hv.OwnerCharacterID, -fee, "maintenance_fee", "sink", hv.ID, hv.Zone)
		hv.HopperUnits += units
		hv.Active = true
		if err := h.db.UpdateHarvester(&hv); err != nil {
			log.Printf("harvest tick: update %s failed: %v", hv.ID, err)
		}
	}
}

// --- Client message handlers ---

func (h *WorldHandler) handleSurvey(client *Client, raw []byte) {
	if client.CharacterID == "" {
		h.sendError(client, "not in world")
		return
	}
	var msg protocol.WSMessage
	json.Unmarshal(raw, &msg)
	dataBytes, _ := json.Marshal(msg.Data)
	var sm protocol.SurveyMsg
	if err := json.Unmarshal(dataBytes, &sm); err != nil {
		h.sendError(client, "invalid survey data")
		return
	}
	types := resources.ToolCategory(sm.Tool)
	if types == nil {
		h.sendError(client, "unknown survey tool")
		return
	}
	surveyMu.Lock()
	if last, ok := lastSurveyByCh[client.CharacterID]; ok &&
		time.Since(last) < resources.SurveyCooldown {
		surveyMu.Unlock()
		h.sendError(client, "survey cooling down")
		return
	}
	lastSurveyByCh[client.CharacterID] = time.Now()
	surveyMu.Unlock()

	spawns, err := h.db.ActiveSpawns(client.Pos.Planet)
	if err != nil {
		h.sendError(client, "survey failed")
		return
	}
	want := make(map[string]bool)
	for _, t := range types {
		want[t] = true
	}
	var best *database.SpawnRow
	bestD := SurveyRadiusM * SurveyRadiusM
	first := true
	for i := range spawns {
		s := &spawns[i]
		if !want[s.Type] {
			continue
		}
		dx := client.Pos.X - s.CenterX
		dz := client.Pos.Z - s.CenterZ
		d := dx*dx + dz*dz
		if d > SurveyRadiusM*SurveyRadiusM {
			continue
		}
		if first || d < bestD {
			best, bestD, first = s, d, false
		}
	}
	if best == nil {
		h.sendError(client, "no matching spawn in range")
		return
	}
	_ = h.db.AddCharacterXP(client.CharacterID, "scouting", 10) // [PROVISIONAL]
	h.send(client, protocol.MsgSurveyResult, protocol.SurveyResultMsg{
		SpawnID:       best.ID,
		ResourceType:  best.Type,
		DistanceM:     math.Sqrt(bestD),
		Concentration: database.EffectiveConcentration(best, client.Pos.X, client.Pos.Z),
		Waypoint:      protocol.WaypointMsg{X: best.CenterX, Z: best.CenterZ},
	})
}

func (h *WorldHandler) handleSample(client *Client, raw []byte) {
	if client.CharacterID == "" {
		h.sendError(client, "not in world")
		return
	}
	spawns, err := h.db.ActiveSpawns(client.Pos.Planet)
	if err != nil {
		h.sendError(client, "sample failed")
		return
	}
	var best *database.SpawnRow
	bestD := SampleRangeM * SampleRangeM
	first := true
	for i := range spawns {
		s := &spawns[i]
		dx := client.Pos.X - s.CenterX
		dz := client.Pos.Z - s.CenterZ
		if d := dx*dx + dz*dz; d <= SampleRangeM*SampleRangeM && (first || d < bestD) {
			best, bestD, first = s, d, false
		}
	}
	if best == nil {
		h.sendError(client, "no spawn in sample range")
		return
	}
	conc := database.EffectiveConcentration(best, client.Pos.X, client.Pos.Z)
	units := 1 + conc/50 // [PROVISIONAL] basic-tool yield curve: 80% → 2 units
	if units > 3 {
		units = 3 // GDD 1–3 band
	}
	if err := h.db.AddToStack(client.CharacterID, best.ID, best.Type, units, best.Stats); err != nil {
		h.sendError(client, "sample storage failed")
		return
	}
	_ = h.db.BumpDepletion(best.ID, 5)                          // GDD 11.3.2 diminishing returns
	_ = h.db.AddCharacterXP(client.CharacterID, "scouting", 25) // [PROVISIONAL]
	h.send(client, protocol.MsgSampleResult, protocol.SampleResultMsg{
		SpawnID: best.ID, ResourceType: best.Type, Units: units,
	})
}

func (h *WorldHandler) handlePlaceHarvester(client *Client, raw []byte) {
	if client.CharacterID == "" {
		h.sendError(client, "not in world")
		return
	}
	var msg protocol.WSMessage
	json.Unmarshal(raw, &msg)
	dataBytes, _ := json.Marshal(msg.Data)
	var pm protocol.PlaceHarvesterMsg
	if err := json.Unmarshal(dataBytes, &pm); err != nil {
		h.sendError(client, "invalid place_harvester data")
		return
	}
	if _, ok := database.HarvesterKinds[pm.Kind]; !ok {
		h.sendError(client, "unknown harvester kind")
		return
	}
	deed, err := h.db.GetItem(client.CharacterID, pm.DeedItemID)
	if err != nil || deed.Schematic != "harvester_deed" {
		h.sendError(client, "valid harvester deed required")
		return
	}
	spawns, err := h.db.ActiveSpawns(client.Pos.Planet)
	if err != nil {
		h.sendError(client, "placement failed")
		return
	}
	var best *database.SpawnRow
	bestConc := 0
	for i := range spawns {
		s := &spawns[i]
		dx := pm.X - s.CenterX
		dz := pm.Z - s.CenterZ
		if dx*dx+dz*dz > s.RadiusM*s.RadiusM {
			continue // must be within the spawn radius (GDD 11.4.1)
		}
		if c := database.EffectiveConcentration(s, pm.X, pm.Z); c > bestConc {
			best, bestConc = s, c
		}
	}
	if best == nil {
		h.sendError(client, "no live spawn at placement")
		return
	}
	all, err := h.db.AllHarvesters()
	if err != nil {
		h.sendError(client, "placement failed")
		return
	}
	for _, hv := range all {
		if hv.Zone != client.Pos.Planet {
			continue
		}
		dx := pm.X - hv.PosX
		dz := pm.Z - hv.PosZ
		if dx*dx+dz*dz < HarvesterExclusionM*HarvesterExclusionM {
			h.sendError(client, "too close to another harvester (100 m rule)")
			return
		}
	}
	id := "hv-" + client.CharacterID[:8] + "-" + pm.Kind
	if err := h.db.PlaceHarvester(id, client.CharacterID, client.Pos.Planet,
		pm.X, pm.Z, pm.Kind, best.ID); err != nil {
		h.sendError(client, "placement failed")
		return
	}
	_ = h.db.DeleteItem(client.CharacterID, pm.DeedItemID) // deed consumed
	h.send(client, protocol.MsgHarvesterPlaced, protocol.HarvesterPlacedMsg{
		HarvesterID: id, SpawnID: best.ID,
	})
}

func (h *WorldHandler) handleEmptyHopper(client *Client, raw []byte) {
	if client.CharacterID == "" {
		h.sendError(client, "not in world")
		return
	}
	owned, err := h.db.GetHarvesters(client.CharacterID)
	if err != nil {
		h.sendError(client, "hopper lookup failed")
		return
	}
	var best *database.HarvesterRow
	bestD := EmptyRangeM * EmptyRangeM
	first := true
	for i := range owned {
		hv := &owned[i]
		if hv.Zone != client.Pos.Planet {
			continue
		}
		dx := client.Pos.X - hv.PosX
		dz := client.Pos.Z - hv.PosZ
		if d := dx*dx + dz*dz; d <= EmptyRangeM*EmptyRangeM && (first || d < bestD) {
			best, bestD, first = hv, d, false
		}
	}
	if best == nil {
		h.sendError(client, "no owned harvester in range")
		return
	}
	spawn, err := h.db.GetSpawn(best.SpawnID)
	if err != nil {
		h.sendError(client, "harvester spawn unknown")
		return
	}
	units, err := h.db.EmptyHopper(best, spawn.Type, spawn.Stats)
	if err != nil {
		h.sendError(client, "hopper transfer failed")
		return
	}
	h.send(client, protocol.MsgHopperEmptied, protocol.HopperEmptiedMsg{
		HarvesterID: best.ID, Units: units,
	})
}

func (h *WorldHandler) handleHarvestCorpse(client *Client, raw []byte) {
	if client.CharacterID == "" {
		h.sendError(client, "not in world")
		return
	}
	corpses, err := h.db.HarvestableCorpses(
		client.Pos.Planet, client.Pos.X, client.Pos.Z,
		CorpseHarvestRangeM, time.Now())
	if err != nil || len(corpses) == 0 {
		h.sendError(client, "no harvestable corpse in range")
		return
	}
	c := corpses[0]
	ytype, ok := corpseYieldType[c.TemplateID]
	if !ok {
		ytype = "cultured_organic" // fallback generic (flagged)
	}
	t := resources.TypeByID(ytype)
	// Per-corpse quality roll (GDD 17.2.2: separate from the regional system).
	// Fresh RNG per call (shared *rand.Rand raced across goroutines).
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	stats := map[string]int{}
	sum := 0
	for _, lane := range t.RelevantLanes {
		v := 300 + rng.Intn(501) // 300–800 band [PROVISIONAL]
		stats[lane] = v
		sum += v
	}
	stats[resources.LaneOQ] = sum / len(t.RelevantLanes)
	units := 2 + rng.Intn(3) // 2–4 [PROVISIONAL: GDD gives no corpse yield counts]
	if err := h.db.AddToStack(client.CharacterID, "corpse:"+c.InstanceID, ytype, units, stats); err != nil {
		h.sendError(client, "corpse harvest storage failed")
		return
	}
	_ = h.db.ConsumeCorpse(c.InstanceID, time.Now())
	_ = h.db.AddCharacterXP(client.CharacterID, "scouting", 25) // [PROVISIONAL]
	h.send(client, protocol.MsgCorpseHarvested, protocol.CorpseHarvestedMsg{
		InstanceID: c.InstanceID, ResourceType: ytype,
		Units: units, Quality: stats[resources.LaneOQ],
	})
}
