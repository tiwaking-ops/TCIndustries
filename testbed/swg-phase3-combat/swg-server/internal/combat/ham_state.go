// Package combat (continued): HAM pool state machine for Phase 3.
// Testbed fork; de-SWG'd (generic content only).
//
// HAM VALUES (owner decision 2026-09-14, C2): Health/Action/Mind pools and all six
// attributes start at 1000 for every species (uniform human baseline). No
// attribute-regeneration formula exists in the cited HAM reference, so regen is
// independently developed HERE as a provisional: 2% of max per 5th 1-second tick
// (Group-C provisional convention, NON-CANONICAL, flagged for review).
package combat

import "time"

// IncapacitationTimeout: GDD-given (9.2.1/9.6.1) — 5 minutes, no approval needed.
const IncapacitationTimeout = 5 * time.Minute

// RegenTickModulo + RegenPctPerTick: [PROVISIONAL] Group-C convention — HAM regen
// fires every 5th 1-second tick at 2% of effective max. No source formula exists;
// independently developed placeholder, flagged for review, NOT tuned.
const (
	RegenTickModulo  = 5
	RegenPctPerTick  = 2
)

// HAMPoolState is the full combat-relevant state for one character.
// BuffHealth/BuffAction/BuffMind are Phase 6 temporary pool increases (GDD
// 24.2.3 shape); they raise effective maxima and are NOT persisted in
// ham_pool_states (they live in the buffs table and are overlaid by callers).
type HAMPoolState struct {
	HealthCurrent int
	HealthMax     int
	ActionCurrent int
	ActionMax     int
	MindCurrent   int
	MindMax       int

	WoundsHealth, WoundsAction, WoundsMind int
	BattleFatiguePct                        float64

	BuffHealth, BuffAction, BuffMind int

	Posture Posture
	Stance  Stance

	IncapacitatedAt    *time.Time
	LastCombatActionAt *time.Time
}

// EffectiveMax computes a pool's max after Wounds (flat), Battle Fatigue (%),
// and buff bonus (flat add, Phase 6). Wound/BF stacking keeps the Group-C min()
// convention; the buff applies on top of whichever reduction wins.
func EffectiveMax(base, wounds int, battleFatiguePct float64, buff int) int {
	woundedMax := base - wounds
	fatiguedMax := int(float64(base) * (1 - battleFatiguePct/100))
	effective := woundedMax
	if fatiguedMax < effective {
		effective = fatiguedMax
	}
	effective += buff
	if effective < 1 {
		effective = 1
	}
	return effective
}

// ApplyDamage subtracts damage from the Health pool (Phase 3 basic attacks always
// target Health; Action/Mind-targeting abilities are later-phase scope).
func ApplyDamage(state HAMPoolState, damage int, now time.Time) (HAMPoolState, bool) {
	effMax := EffectiveMax(state.HealthMax, state.WoundsHealth, state.BattleFatiguePct, state.BuffHealth)
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

// IsDead reports whether the incapacitation timer expired without a revive.
func IsDead(state HAMPoolState, now time.Time) bool {
	if state.IncapacitatedAt == nil {
		return false
	}
	return now.Sub(*state.IncapacitatedAt) >= IncapacitationTimeout
}

// Revive implements GDD 9.6.1: standing at 10% pools, gains Wounds + Battle Fatigue.
// [PROVISIONAL] GDD gives no revive-specific wound/BF numbers; Group-C convention
// reuses the death range midpoints (50–200 → 125; 1–3% → 2%), flagged for review.
func Revive(state HAMPoolState, now time.Time) HAMPoolState {
	effMax := EffectiveMax(state.HealthMax, state.WoundsHealth, state.BattleFatiguePct, state.BuffHealth)
	state.HealthCurrent = effMax / 10
	if state.HealthCurrent < 1 {
		state.HealthCurrent = 1
	}
	state.IncapacitatedAt = nil
	state.WoundsHealth += 125
	state.BattleFatiguePct += 2
	if state.BattleFatiguePct > 100 {
		state.BattleFatiguePct = 100
	}
	state.Posture = PostureStanding
	return state
}

// Clone implements GDD 9.6.2/9.6.3: full pools at the bound facility plus the same
// provisional wound/BF gain. Item-condition loss is out of scope (no items exist).
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

// RegenTick restores out-of-combat HAM per the provisional convention above.
// now is passed in so tests can drive time without sleeping the tick loop.
func RegenTick(state HAMPoolState, tickCount int64) HAMPoolState {
	if state.IncapacitatedAt != nil {
		return state // no regen while incapacitated
	}
	if tickCount%RegenTickModulo != 0 {
		return state
	}
	heal := func(cur, base, wounds, buff int) int {
		effMax := EffectiveMax(base, wounds, state.BattleFatiguePct, buff)
		cur += effMax * RegenPctPerTick / 100
		if cur > effMax {
			cur = effMax
		}
		return cur
	}
	state.HealthCurrent = heal(state.HealthCurrent, state.HealthMax, state.WoundsHealth, state.BuffHealth)
	state.ActionCurrent = heal(state.ActionCurrent, state.ActionMax, state.WoundsAction, state.BuffAction)
	state.MindCurrent = heal(state.MindCurrent, state.MindMax, state.WoundsMind, state.BuffMind)
	return state
}
