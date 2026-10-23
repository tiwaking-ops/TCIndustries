package resources

import (
	"errors"
	"math"
	"time"
)

// HarvesterType defines harvester size and rates (GDD 10.6).
type HarvesterType struct {
	Name        string
	ExtractRate int // units per hour at 100% concentration
	HopperCap   int // max units stored
	Maintenance int // credits per week
}

var HarvesterTypes = map[string]HarvesterType{
	"personal": {Name: "Personal Harvester", ExtractRate: 50, HopperCap: 10000, Maintenance: 500},
	"medium":   {Name: "Medium Harvester", ExtractRate: 100, HopperCap: 25000, Maintenance: 2000},
	"heavy":    {Name: "Heavy Harvester", ExtractRate: 200, HopperCap: 50000, Maintenance: 5000},
}

// HarvesterInstance represents a placed harvester structure in the world.
type HarvesterInstance struct {
	ID           string    `json:"id"`
	OwnerID      string    `json:"owner_id"`
	TypeName     string    `json:"type_name"`
	ResourceID   string    `json:"resource_id"`
	ResourceName string    `json:"resource_name"`
	X            float64   `json:"x"`
	Z            float64   `json:"z"`
	HopperAmount int       `json:"hopper_amount"`
	Maintenance  int       `json:"maintenance_pool"` // credits remaining
	LastTick     time.Time `json:"last_tick"`
}

// HarvesterManager manages placed harvesters and hourly extraction ticks.
type HarvesterManager struct {
	harvesters map[string]*HarvesterInstance
}

func NewHarvesterManager() *HarvesterManager {
	return &HarvesterManager{
		harvesters: make(map[string]*HarvesterInstance),
	}
}

// PlaceHarvester validates 100m exclusion zone and places harvester.
func (hm *HarvesterManager) PlaceHarvester(id, ownerID, typeName, resourceID, resourceName string, x, z float64, spawn *Spawn) (*HarvesterInstance, error) {
	ht, ok := HarvesterTypes[typeName]
	if !ok {
		return nil, errors.New("unknown harvester type")
	}

	// Check 100m exclusion zone from other harvesters
	for _, h := range hm.harvesters {
		dx := h.X - x
		dz := h.Z - z
		if math.Sqrt(dx*dx+dz*dz) < 100.0 {
			return nil, errors.New("too close to another harvester (100m exclusion required)")
		}
	}

	h := &HarvesterInstance{
		ID:           id,
		OwnerID:      ownerID,
		TypeName:     typeName,
		ResourceID:   resourceID,
		ResourceName: resourceName,
		X:            x,
		Z:            z,
		HopperAmount: 0,
		Maintenance:  ht.Maintenance,
		LastTick:     time.Now(),
	}

	hm.harvesters[id] = h
	return h, nil
}

// TickHarvesters simulates hourly extraction (or fast-cycle rate).
func (hm *HarvesterManager) TickHarvesters(spawns []Spawn) {
	now := time.Now()
	for _, h := range hm.harvesters {
		ht := HarvesterTypes[h.TypeName]
		elapsed := now.Sub(h.LastTick).Hours()
		if elapsed < (1.0 / 60.0) { // minimum tick interval (1 min compressed)
			continue
		}

		// Find resource concentration at (h.X, h.Z)
		var targetSpawn *Spawn
		for i := range spawns {
			if spawns[i].ID == h.ResourceID {
				targetSpawn = &spawns[i]
				break
			}
		}

		if targetSpawn == nil {
			continue // resource depleted / expired
		}

		dx := targetSpawn.CenterX - h.X
		dz := targetSpawn.CenterZ - h.Z
		dist := math.Sqrt(dx*dx + dz*dz)
		if dist > targetSpawn.Radius {
			continue // outside spawn radius
		}

		factor := 1.0 - (dist / targetSpawn.Radius)
		conc := float64(targetSpawn.PeakConc) * factor
		if conc < 0 {
			conc = 0
		}

		extracted := int(float64(ht.ExtractRate) * (conc / 100.0) * elapsed)
		if extracted > 0 {
			h.HopperAmount += extracted
			if h.HopperAmount > ht.HopperCap {
				h.HopperAmount = ht.HopperCap
			}
		}
		h.LastTick = now
	}
}
