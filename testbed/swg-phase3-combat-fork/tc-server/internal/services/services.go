package services

import "time"

const (
	MedicNoviceBox       = "medic_novice"
	EntertainerNoviceBox = "entertainer_novice"
)

const (
	BuffHealth = "health"
	BuffAction = "action"
	BuffMind   = "mind"
)

func ValidBuffPool(pool string) bool {
	return pool == BuffHealth || pool == BuffAction || pool == BuffMind
}

func ValidKind(kind string) bool {
	return kind == "music" || kind == "dance"
}

func BuffMagnitude(healingTier int) int {
	if healingTier < 0 {
		healingTier = 0
	}
	if healingTier > 4 {
		healingTier = 4
	}
	return 500 + 250 * healingTier
}

const BuffDuration = 30 * time.Minute

func WoundHealAmount(healingTier int) int {
	if healingTier < 0 {
		healingTier = 0
	}
	if healingTier > 4 {
		healingTier = 4
	}
	return 25 + 25 * healingTier
}

const HealCooldown = 60 * time.Second
const HealRangeM = 8.0
const BuffRangeM = 8.0
const WatchRadiusM = 20.0
const MindDrainPerTick = 10

func BFHealPerTick(healingTier, perfTier int) float64 {
	return float64(1 + healingTier + perfTier)
}

const (
	HealXP   = 50
	BuffXP   = 50
	ReviveXP = 100
	TipXP    = 25
)
