// Package combat implements attack resolution for Phase 3 — Combat Core.
// Testbed fork of the Perplexity Phase-2 line; de-SWG'd (generic content only).
//
// ACCURACY MODEL (owner decision 2026-09-14, C1): values from
// https://swgemulator.fandom.com/wiki/Weapon_Accuracy "where possible":
//   - Attacker posture mods: standing +0, crouching +16, prone +50, running -50.
//   - Target posture mods: standing +0; crouching -16 vs ranged / +16 vs melee;
//     prone -25 vs ranged / +25 vs melee.
//   - Scaling: +1 accuracy (skill or weapon) = +0.5% hit; +1 defense = -0.5% hit;
//     accuracy-above-defense guarantees at least ~66% standing-vs-standing.
//   - The page's master ToHit equation is an IMAGE (no readable text), and state
//     mods / secondary defenses (dodge/counter/block) are listed as UNKNOWN, so the
//     exact original equation could NOT be recovered — the assembly below (66-base
//     + 0.5%/point differential, 0..100 clamp) is the documented analysis on that
//     page, not a recovered original. Recorded, not claimed.
// GDD-silent values below (unarmed accuracy/damage, clamp, range model) follow the
// Group-C provisional convention (owner decision 2026-09-14, proposal §6): adopted as
// NON-CANONICAL placeholders, flagged [PROVISIONAL].
package combat

import "math/rand"

// Posture identifies a character's combat posture.
type Posture string

const (
	PostureStanding Posture = "standing"
	PostureKneeling Posture = "kneeling"
	PostureProne    Posture = "prone"
	PostureCrouched Posture = "crouched"
)

// Stance identifies a character's combat stance.
type Stance string

const (
	StanceNormal     Stance = "normal"
	StanceAggressive Stance = "aggressive"
	StanceDefensive  Stance = "defensive"
	StanceBerserk    Stance = "berserk"
)

// Weapon is the attacker's weapon. Only the unarmed baseline exists until a
// crafted-weapon system exists (Phase 4 scope).
type Weapon struct {
	Name          string
	Melee         bool // true = melee (target-posture melee column applies)
	BaseAccuracy  int
	MinDamage     int
	MaxDamage     int
	DamageType    string
}

// Attacker carries the attacker's combat-relevant state for one resolution.
type Attacker struct {
	AccuracySkill int
	DamageBonus   int
	Posture       Posture
	Stance        Stance
}

// Defender carries the defender's combat-relevant state for one resolution.
type Defender struct {
	DefenseSkill         int
	Posture              Posture
	ArmorMitigationPct   float64
	ResistanceMultiplier float64
}

// AttackResult is the outcome of one resolved attack.
type AttackResult struct {
	HitChance float64
	Roll      int
	Hit       bool
	Damage    int
}

// attackerPostureMod maps base-game postures onto the fandom page's posture mods
// (standing +0, crouching +16, prone +50, running -50).
// [ASSUMPTION] the fandom table has no kneeling row: kneeling maps to crouching
// (+16); crouched maps to crouching (+16). No running posture exists in this build.
func attackerPostureMod(p Posture) int {
	switch p {
	case PostureProne:
		return 50
	case PostureKneeling, PostureCrouched:
		return 16
	default:
		return 0
	}
}

// targetPostureMod maps the defender's posture onto the fandom page's target mods.
// Phase 3 combat is melee-only (unarmed), so the melee column applies:
// standing +0, crouching +16, prone +25.
// [ASSUMPTION] kneeling maps to the crouching row (+16), same basis as above.
func targetPostureMod(p Posture, melee bool) int {
	switch p {
	case PostureProne:
		if melee {
			return 25
		}
		return -25
	case PostureKneeling, PostureCrouched:
		if melee {
			return 16
		}
		return -16
	default:
		return 0
	}
}

// stanceDamageMult applies stance damage modifiers (GDD 9.2.3: aggressive +10%,
// defensive -10%, berserk +20%). Defense-side stance effects are out of scope for
// the MVP (no defender stances on creatures); recorded, not silently dropped.
func stanceDamageMult(s Stance) float64 {
	switch s {
	case StanceAggressive:
		return 1.10
	case StanceDefensive:
		return 0.90
	case StanceBerserk:
		return 1.20
	default:
		return 1.0
	}
}

// ResolveAttack resolves one attack: hit chance per the fandom documented analysis
// (66% tied-standing baseline, 0.5% per point of accuracy-vs-defense differential),
// d100 roll, then damage per the GDD 9.2.2 shape with Group-C provisional weapon
// numbers. Pure function — deterministic given rng.
func ResolveAttack(att Attacker, def Defender, w Weapon, rng *rand.Rand) AttackResult {
	attTotal := w.BaseAccuracy + att.AccuracySkill + attackerPostureMod(att.Posture)
	defTotal := def.DefenseSkill + targetPostureMod(def.Posture, w.Melee)
	hitChance := 66.0 + 0.5*float64(attTotal-defTotal)
	if hitChance < 0 {
		hitChance = 0
	}
	if hitChance > 100 {
		hitChance = 100
	}
	roll := rng.Intn(100) + 1
	hit := float64(roll) <= hitChance
	damage := 0
	if hit {
		span := w.MaxDamage - w.MinDamage + 1
		raw := w.MinDamage + rng.Intn(span) + att.DamageBonus
		dmg := float64(raw) * stanceDamageMult(att.Stance)
		dmg = dmg * (1 - def.ArmorMitigationPct/100.0) * def.ResistanceMultiplier
		damage = int(dmg)
		if damage < 0 {
			damage = 0
		}
	}
	return AttackResult{HitChance: hitChance, Roll: roll, Hit: hit, Damage: damage}
}
