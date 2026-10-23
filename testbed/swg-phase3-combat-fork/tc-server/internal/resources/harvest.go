package resources

import (
	"errors"
	"time"
)

// CorpseHarvestResult represents materials harvested from a creature corpse (GDD 17.2.2).
type CorpseHarvestResult struct {
	CorpseID     string `json:"corpse_id"`
	ResourceName string `json:"resource_name"`
	Category     string `json:"category"`
	Amount       int    `json:"amount"`
	Quality      int    `json:"quality"`
}

// HarvestManager tracks creature corpse windows and harvesting state.
type HarvestManager struct {
	harvestedCorpses map[string]bool // corpseID -> true if already harvested
}

func NewHarvestManager() *HarvestManager {
	return &HarvestManager{
		harvestedCorpses: make(map[string]bool),
	}
}

// HarvestCorpse harvests materials from a dead creature instance within the 10-minute window (GDD 17.2.2).
func (hm *HarvestManager) HarvestCorpse(corpseID string, diedAt time.Time, templateName string) (CorpseHarvestResult, error) {
	if hm.harvestedCorpses[corpseID] {
		return CorpseHarvestResult{}, errors.New("corpse has already been harvested")
	}

	// Check 10-minute window
	if time.Since(diedAt) > 10*time.Minute {
		return CorpseHarvestResult{}, errors.New("corpse has expired and cannot be harvested")
	}

	hm.harvestedCorpses[corpseID] = true

	// Yield meat/hide/bone based on template
	return CorpseHarvestResult{
		CorpseID:     corpseID,
		ResourceName: templateName + " Meat",
		Category:     "cultured_organic",
		Amount:       5 + (time.Now().Nanosecond() % 10), // 5-14 units
		Quality:      400 + (time.Now().Nanosecond() % 200),
	}, nil
}
