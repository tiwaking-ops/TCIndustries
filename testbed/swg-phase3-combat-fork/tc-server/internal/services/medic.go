package services

import (
	"errors"
	"time"
)

// WoundHealResult represents the outcome of a medic healing wounds (GDD 24.2.1).
type WoundHealResult struct {
	TargetID    string `json:"target_id"`
	Pool        string `json:"pool"` // "health", "action", "mind"
	HealedAmount int   `json:"healed_amount"`
	CooldownSec  int   `json:"cooldown_sec"`
}

// MedicService manages medical actions, wounds healing, and revive gating (GDD 24.2.1, 24.2.2).
type MedicService struct {
	lastHealTime map[string]map[string]time.Time // healerID -> targetID -> timestamp
}

func NewMedicService() *MedicService {
	return &MedicService{
		lastHealTime: make(map[string]map[string]time.Time),
	}
}

// CanHealWounds checks per-pair cooldown (60s cooldown per GDD/provisional).
func (ms *MedicService) CanHealWounds(healerID, targetID string) (bool, time.Duration) {
	targets, ok := ms.lastHealTime[healerID]
	if !ok {
		return true, 0
	}
	last, ok := targets[targetID]
	if !ok {
		return true, 0
	}
	elapsed := time.Since(last)
	cooldown := 60 * time.Second
	if elapsed < cooldown {
		return false, cooldown - elapsed
	}
	return true, 0
}

// RecordHeal marks wound heal used for cooldown.
func (ms *MedicService) RecordHeal(healerID, targetID string) {
	targets, ok := ms.lastHealTime[healerID]
	if !ok {
		targets = make(map[string]time.Time)
		ms.lastHealTime[healerID] = targets
	}
	targets[targetID] = time.Now()
}

// HealWounds computes flat-plus-skill-scaled wound healing (GDD 24.2.1).
func (ms *MedicService) HealWounds(healerID, targetID string, pool string, currentWounds int, healingTier int) (int, error) {
	can, remaining := ms.CanHealWounds(healerID, targetID)
	if !can {
		return 0, errors.New("wound healing on cooldown for this target (" + remaining.Round(time.Second).String() + " remaining)")
	}

	if currentWounds <= 0 {
		return 0, errors.New("target has no wounds in this pool")
	}

	healAmount := 25 + (25 * healingTier)
	if healAmount > currentWounds {
		healAmount = currentWounds
	}

	ms.RecordHeal(healerID, targetID)
	return healAmount, nil
}
