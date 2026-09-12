// Package combat (continued): HAM pool state machine.
// GDD Section 9.2.1 (HAM pools, wounds, battle fatigue) and
// Section 9.6 (incapacitation, death, clone facilities).
package combat

import "time"

// IncapacitationTimeout is how long an incapacitated player can be revived
// before dying and needing to clone. GDD 9.2.1: "Incapacitation timer: 5 minutes".
const IncapacitationTimeout = 5 * time.Minute

// HAMPoolState is the full combat-relevant state for one character.
// Mirrors the columns already present in ham_pool_states (health_current,
// health_max, etc. -- these existed in the schema before Phase 3 but were
// never fully read back by application code; see the GetCharacterByID fix
// in combat_db.go).
type HAMPoolState struct {
	HealthCurrent int
	HealthMax     int
	ActionCurrent int
	ActionMax     int
	MindCurrent   int
	MindMax       int

	Strength, Constitution     int
	Quickness, Stamina         int
	Focus, Willpower           int

	WoundsHealth, WoundsAction, WoundsMind int
	BattleFatiguePct                        float64

	Posture Posture
	Stance  Stance

	IncapacitatedAt     *time.Time
	LastCombatActionAt  *time.Time
}

// EffectiveMax computes a pool's actual max after Wounds and Battle Fatigue,
// per GDD 9.2.1: "Wounds reduce maximum HAM pool size" (flat) and
// "Battle Fatigue reduces maximum HAM pool size" (percentage), applied on
// top of wounds. Example from the GDD: 500 Health Wounds reduce max Health
// from 5000 to 4500 (flat); 20% Battle Fatigue reduces max Health from
// 5000 to 4000 (this example implies BF is applied to the *un-wounded*
// base, not compounded on top of the wounds-adjusted value -- so that's
// the order implemented here: wounds subtract first, then BF percentage
// is applied to the original base, and the smaller result of the two
// reductions... [ASSUMPTION] Re-reading 9.2.1's two examples, both are
// given independently off the same 5000 base rather than stacked, and the
// GDD does not specify the combined-reduction case. This implementation
// applies both reductions independently against the species base and takes
// whichever is lower (wounds and BF both hurt you; the pool reflects the
// worse of the two, not a naive double-penalty) since that keeps very
// wounded + very fatigued characters from having a max of zero or negative
// from double-counting. Flagged here rather than silently picked.
func EffectiveMax(speciesMax, wounds int, battleFatiguePct float64) int {
	woundedMax := speciesMax - wounds
	fatiguedMax := int(float64(speciesMax) * (1 - battleFatiguePct/100))
	effective := woundedMax
	if fatiguedMax < effective {
		effective = fatiguedMax
	}
	if effective < 1 {
		effective = 1 // a pool can be nearly exhausted but never non-positive
	}
	return effective
}

// ApplyDamage subtracts damage from the Health pool (Phase 3's basic
// attacks always target Health; Action/Mind-targeting abilities are a
// later-phase special-attack feature per GDD 9.2.4). Returns the updated
// state and whether this damage caused incapacitation.
func ApplyDamage(state HAMPoolState, damage int, now time.Time) (HAMPoolState, bool) {
	effMax := EffectiveMax(state.HealthMax, state.WoundsHealth, state.BattleFatiguePct)
	state.HealthCurrent -= damage
	if state.HealthCurrent > effMax {
		state.HealthCurrent = effMax
	}

	justIncapacitated := false
	if state.HealthCurrent <= 0 && state.IncapacitatedAt == nil {
		state.HealthCurrent = 0
		state.IncapacitatedAt = &now
		justIncapacitated = true
	}
	state.LastCombatActionAt = &now
	return state, justIncapacitated
}

// IsDead reports whether the incapacitation timer has expired without a
// revive. GDD 9.6.2: "Incapacitation timer expires without revive."
func IsDead(state HAMPoolState, now time.Time) bool {
	if state.IncapacitatedAt == nil {
		return false
	}
	return now.Sub(*state.IncapacitatedAt) >= IncapacitationTimeout
}

// Revive implements GDD 9.6.1's revive process: "Player returns to
// standing with 10% HAM pools. Player gains Wounds and Battle Fatigue."
// [ASSUMPTION] exact wound/BF gain on revive isn't numerically specified
// beyond "gains Wounds and Battle Fatigue" -- using the same 50-200 Wounds
// and 1-3% Battle Fatigue range GDD 9.6.2 gives for ordinary death, since
// that's the only concrete range the document provides for this kind of
// event, applied at its midpoint rather than randomized for reproducibility.
func Revive(state HAMPoolState, now time.Time) HAMPoolState {
	effMax := EffectiveMax(state.HealthMax, state.WoundsHealth, state.BattleFatiguePct)
	state.HealthCurrent = effMax / 10
	if state.HealthCurrent < 1 {
		state.HealthCurrent = 1
	}
	state.IncapacitatedAt = nil
	state.WoundsHealth += 125 // midpoint of GDD 9.6.2's 50-200 range
	state.BattleFatiguePct += 2 // midpoint of GDD 9.6.2's 1-3% range
	if state.BattleFatiguePct > 100 {
		state.BattleFatiguePct = 100
	}
	state.Posture = PostureStanding
	return state
}

// Clone implements GDD 9.6.2/9.6.3: death after the incapacitation timer
// expires, respawn at the bound clone facility, lose 5% item condition
// (item-condition side effects are out of Phase 3's scope -- no items yet),
// gain 1-3% Battle Fatigue and 50-200 Wounds (same midpoint convention as
// Revive above, for the same reason).
func Clone(state HAMPoolState, now time.Time) HAMPoolState {
	state.HealthCurrent = state.HealthMax
	state.ActionCurrent = state.ActionMax
	state.MindCurrent = state.MindMax
	state.IncapacitatedAt = nil
	state.WoundsHealth += 125
	state.BattleFatiguePct += 2
	if state.BattleFatiguePct > 100 {
		state.BattleFatiguePct = 100
	}
	state.Posture = PostureStanding
	return state
}
