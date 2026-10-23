package crafting

import (
	"errors"
	"math/rand"
)

// CraftingSession represents an active crafting process (GDD 10.2, 10.7).
type CraftingSession struct {
	CharacterID    string             `json:"character_id"`
	SchematicID    string             `json:"schematic_id"`
	State          string             `json:"state"`           // "assembly", "experimenting", "finalized"
	ExperimentPts  int                `json:"experiment_pts"`  // remaining points
	ExpAllocations map[string]int     `json:"exp_allocations"` // slot -> points allocated
	ComputedStats  map[string]float64 `json:"computed_stats"`
	ResultItem     string             `json:"result_item"`
}

// Experimentation outcome table per GDD 10.7.1 (+5-10, +2-5, +0-2, 0, -2-5%).
var expTable = []float64{0.10, 0.50, 0.30, 0.08, 0.02} // success probabilities

// StartCraftingSession initializes a crafting session for a schematic.
func StartCraftingSession(characterID string, schem Schematic) *CraftingSession {
	allocs := make(map[string]int)
	for _, slot := range schem.ExpSlots {
		allocs[slot] = 0
	}
	return &CraftingSession{
		CharacterID:    characterID,
		SchematicID:    schem.ID,
		State:          "experimenting",
		ExperimentPts:  10, // starting experimentation points
		ExpAllocations: allocs,
		ComputedStats: map[string]float64{
			"quality": 500.0,
		},
		ResultItem: schem.ResultItem,
	}
}

// AllocateExperimentation allocates points to an experimentation lane (max 5 pts/round, GDD 10.7).
func (cs *CraftingSession) AllocateExperimentation(slot string, pts int) error {
	if cs.State != "experimenting" {
		return errors.New("crafting session is not in experimentation state")
	}
	if pts <= 0 || pts > 5 {
		return errors.New("experimentation points per attempt must be between 1 and 5")
	}
	if cs.ExperimentPts < pts {
		return errors.New("insufficient experimentation points remaining")
	}
	_, ok := cs.ExpAllocations[slot]
	if !ok {
		return errors.New("invalid experimentation slot")
	}

	cs.ExperimentPts -= pts
	cs.ExpAllocations[slot] += pts

	// Roll experimentation outcome per table 10.7.1
	roll := rand.Float64()
	cumulative := 0.0
	multiplier := 0.0
	for i, p := range expTable {
		cumulative += p
		if roll <= cumulative {
			switch i {
			case 0: // major success
				multiplier = 0.08
			case 1: // success
				multiplier = 0.04
			case 2: // marginal success
				multiplier = 0.01
			case 3: // neutral
				multiplier = 0.0
			case 4: // failure
				multiplier = -0.03
			}
			break
		}
	}

	cs.ComputedStats[slot] += float64(pts) * multiplier * 100.0
	return nil
}

// FinalizeCrafting completes the crafting session and produces the item stats.
func (cs *CraftingSession) FinalizeCrafting(crafterName string) map[string]interface{} {
	cs.State = "finalized"
	return map[string]interface{}{
		"item_template": cs.ResultItem,
		"crafter":       crafterName,
		"stats":         cs.ComputedStats,
		"quality":       cs.ComputedStats["quality"],
	}
}
