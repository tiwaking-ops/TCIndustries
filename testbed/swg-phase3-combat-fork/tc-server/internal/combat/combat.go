// Package combat implements HAM-driven combat resolution per GDD Section 9.2.
//
// Combat is deterministic and stat-based (GDD 9.1), not twitch-based, which
// is why this whole package has no client-side prediction concerns: every
// roll happens server-side (GDD 29.5) and the result is pushed to clients.
package combat

import "math/rand"

// Posture affects defense and movement. GDD Section 9.2.3.
type Posture string

const (
	PostureStanding Posture = "standing"
	PostureKneeling Posture = "kneeling"
	PostureProne    Posture = "prone"
	PostureCrouched Posture = "crouched"
)

// postureRangedAccuracyBonus: attacker's ranged accuracy bonus by posture.
// GDD 9.2.3: "Kneeling: +10% ranged accuracy", "Prone: +20% ranged accuracy".
var postureRangedAccuracyBonus = map[Posture]int{
	PostureStanding: 0,
	PostureKneeling: 10,
	PostureProne:    20,
	PostureCrouched: 0,
}

// postureMeleeDefenseBonus: target's melee defense modifier by posture.
// GDD 9.2.3: "Kneeling: -10% melee defense", "Prone: -30% melee defense",
// "Crouched: +5% melee defense".
var postureMeleeDefenseBonus = map[Posture]int{
	PostureStanding: 0,
	PostureKneeling: -10,
	PostureProne:    -30,
	PostureCrouched: 5,
}

// NOTE [gap, flagged rather than invented]: GDD 9.2.3 states Prone also
// gives the defender "harder to hit at range" in addition to its melee
// defense penalty, but does not give that ranged-defense effect an exact
// number the way it does every other posture modifier. This package
// implements only the two explicitly-quantified effects and leaves Prone's
// extra ranged-defense bonus at 0 rather than inventing a figure. Revisit
// if/when the GDD is amended with a specific value.

// Stance affects damage output and cannot be changed mid-combat. GDD 9.2.3.
type Stance string

const (
	StanceNormal     Stance = "normal"
	StanceAggressive Stance = "aggressive"
	StanceDefensive  Stance = "defensive"
	StanceBerserk    Stance = "berserk"
)

// stanceDamageModifier: GDD 9.2.3 "Aggressive: +10% damage, -10% defense",
// "Defensive: -10% damage, +10% defense", "Berserk: +20% damage, -20%
// defense, +10% attack speed". Attack-speed and defense trade-offs are not
// modeled yet in Phase 3's single-attack resolution; only the damage term
// is applied here.
var stanceDamageModifier = map[Stance]float64{
	StanceNormal:     0.0,
	StanceAggressive: 0.10,
	StanceDefensive:  -0.10,
	StanceBerserk:     0.20,
}

// WeaponType determines which posture rule set applies (ranged accuracy vs
// melee defense are two different posture effects per GDD 9.2.3).
type WeaponType string

const (
	WeaponRanged WeaponType = "ranged"
	WeaponMelee  WeaponType = "melee"
)

// Weapon is the subset of a crafted weapon's stats relevant to attack
// resolution. GDD Section 9.2.5 / 10.9.2.
type Weapon struct {
	Type          WeaponType
	BaseAccuracy  int
	MinDamage     int
	MaxDamage     int
	DamageType    string // "kinetic", "energy", "elemental", "stun" (GDD 9.2.5)
}

// Attacker is the subset of combat-relevant state needed to resolve an attack.
type Attacker struct {
	AccuracySkill   int
	DamageSkillBonus int
	Posture         Posture
	Stance          Stance
}

// Defender is the subset of combat-relevant state needed to resolve an attack.
type Defender struct {
	DefenseSkill       int
	Posture            Posture
	ArmorMitigationPct float64 // 0-100
	ResistanceMultiplier float64 // GDD 9.2.5 armor-type resistances; 1.0 = no bonus
}

// AttackResult is what gets sent to clients (combat_result message) and
// applied to the defender's HAM pool.
type AttackResult struct {
	Hit       bool
	HitChance int
	Roll      float64
	Damage    int // always targets Health for Phase 3's basic attacks
}

// ResolveAttack implements GDD Section 9.2.2's attack flow:
//
//	Hit Chance = Base Weapon Accuracy + Attacker Accuracy Skill
//	             - Target Defense Skill - Target Posture Penalty (melee only)
//	             + Attacker Posture Bonus (ranged only)
//	Damage = (Weapon Min-Max Damage + Skill Damage Bonus) * Stance Modifier
//	         * (1 - Armor Mitigation %) * Resistance Multiplier
//
// This is a direct port of the standalone reference implementation already
// verified (by actual execution, not just inspection) against sample
// inputs earlier in this project — see ham_combat_resolution.py.
func ResolveAttack(a Attacker, d Defender, w Weapon, rng *rand.Rand) AttackResult {
	postureBonus := 0
	if w.Type == WeaponRanged {
		postureBonus = postureRangedAccuracyBonus[a.Posture]
	}
	targetPosturePenalty := 0
	if w.Type == WeaponMelee {
		targetPosturePenalty = postureMeleeDefenseBonus[d.Posture]
	}

	hitChance := w.BaseAccuracy + a.AccuracySkill - d.DefenseSkill - targetPosturePenalty + postureBonus
	if hitChance < 5 {
		hitChance = 5 // sane floor, not GDD-specified -- prevents "unhittable" edge cases
	}
	if hitChance > 95 {
		hitChance = 95 // sane ceiling, not GDD-specified -- prevents "unmissable" edge cases
	}

	roll := rng.Float64() * 100
	if roll > float64(hitChance) {
		return AttackResult{Hit: false, HitChance: hitChance, Roll: roll, Damage: 0}
	}

	base := float64(w.MinDamage) + rng.Float64()*float64(w.MaxDamage-w.MinDamage)
	base += float64(a.DamageSkillBonus)
	stanceMod := 1 + stanceDamageModifier[a.Stance]
	mitigated := base * stanceMod * (1 - d.ArmorMitigationPct/100) * d.ResistanceMultiplier
	if mitigated < 0 {
		mitigated = 0
	}

	return AttackResult{Hit: true, HitChance: hitChance, Roll: roll, Damage: int(mitigated + 0.5)}
}
