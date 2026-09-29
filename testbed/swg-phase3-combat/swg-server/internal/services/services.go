// Package services implements Phase 6 (Social Support Professions) tuning and
// pure formulas: wound healing, buffs, performance/battle-fatigue healing, tips.
// Testbed fork; generic content only.
//
// GDD-given values are cited inline. Everything else is [PROVISIONAL] under the
// owner-approved convention (proposal §6): magnitudes anchored to fork-local
// scales (combat damage bands, fee bands, XP award bands) to avoid inventing
// fresh tuning; all NON-CANONICAL, flagged here and in HYGIENE_NOTE.md.
package services

import "time"

// Buff pools (GDD 23.3 BuffInstance shared shape: one type per HAM pool).
const (
	BuffHealth = "health"
	BuffAction = "action"
	BuffMind   = "mind"
)

// BuffMagnitude returns the HAM-pool increase for a medic of the given Healing
// tree tier (owned medic_healing_* boxes, 0–4). GDD band +500–3000 (24.2.3);
// MVP uses the band floor scaled modestly by tier. [PROVISIONAL]
func BuffMagnitude(healingTier int) int {
	if healingTier < 0 {
		healingTier = 0
	}
	if healingTier > 4 {
		healingTier = 4
	}
	return 500 + 250*healingTier // 500–1500
}

// BuffDuration: GDD 30–60+ min band ([ASSUMPTION]-tagged); MVP takes the floor.
const BuffDuration = 30 * time.Minute

// WoundHealAmount returns wounds cleared per heal action: flat 25 plus 25 per
// Healing tier (GDD 24.2.1 "flat-plus-skill-scaled"; split itself provisional).
func WoundHealAmount(healingTier int) int {
	if healingTier < 0 {
		healingTier = 0
	}
	if healingTier > 4 {
		healingTier = 4
	}
	return 25 + 25*healingTier // 25–125
}

// HealCooldown: per-healer-target diminishing window (GDD 24.2.1 shape;
// duration provisional).
const HealCooldown = 60 * time.Second

// HealRangeM: wound-heal application range [PROVISIONAL — GDD silent; mirrors
// the 8 m attack-range analog rather than the 5 m revive gate].
const HealRangeM = 8.0

// BFHealPerTick returns Battle Fatigue % cleared per 5 s performance tick:
// base 1 plus Healing tier plus the performed discipline tier (Music xor Dance),
// capturing 23.6's Musician/Dancer differentiation through existing tree data.
// [PROVISIONAL rate]
func BFHealPerTick(healingTier, perfTier int) float64 {
	return float64(1 + healingTier + perfTier)
}

// WatchRadiusM: audience pairing radius [PROVISIONAL — GDD says "nearby /
/// venue watch radius" without metres].
const WatchRadiusM = 20.0

// MindDrainPerTick: performer Mind cost per 1 s tick. Must exceed regen
// (~4/s at 1000 Mind) or Mind never visibly drains; 10/s gives ~100 s of
// performance per full pool at test scale. [PROVISIONAL]
const MindDrainPerTick = 10

// BuffRangeM: buff application range [PROVISIONAL, same basis as HealRangeM].
const BuffRangeM = 8.0

// Service XP awards (GDD tips-XP per 7.2.3 exists; magnitudes provisional).
const (
	HealXP   = 50  // medical, per wound-heal action
	BuffXP   = 50  // medical, per buff application
	ReviveXP = 100 // medical, per successful revive
	TipXP    = 25  // entertaining, per tip received
)

// Valid performance kinds.
func ValidKind(kind string) bool {
	return kind == "music" || kind == "dance"
}
