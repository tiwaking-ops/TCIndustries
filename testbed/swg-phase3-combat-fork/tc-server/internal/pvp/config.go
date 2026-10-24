// Package pvp holds Phase 8 (Faction & PvP) domain logic: the attack-legality
// engine, flag/timer state, duels, wars, and battlefields. Testbed fork;
// generic placeholder content only (R-03: FACTION_A / FACTION_B).
//
// Design authority: proposals/proposal-claude-tcindustries-phase8-faction-pvp-v1.md
//   R-01..R-08 — human rulings, final.
//   D-01..D-16 — approved defaults, implemented as configurable values here so a
//               tuning change never needs a refactor.
package pvp

import (
	"os"
	"time"
)

// PvPConfig centralises every Phase 8 tunable (proposal D-01..D-16). Durations
// compress under TESTBED_FAST_CYCLE=1 so the integration suite runs in minutes
// instead of 24 hours (same convention as the Phase 4/5/7 test compression).
// Values are computed once at package init; tests can override via
// withOverriddenConfig for fast, deterministic sub-second assertions.
type PvPConfig struct {
	// D-01: delay before a covert → overt request takes effect.
	OvertDelay time.Duration
	// D-02: delay before an overt → covert request takes effect, allowed only
	// once the aggression timer has expired.
	CovertDelay time.Duration
	// D-03: aggression timer length; refreshed to full on every PvP attack
	// made or received (any mode) and on leaving a battlefield zone.
	AggressionTimer time.Duration
	// D-04: faction switch/leave is blocked while overt or while the aggression
	// timer is active.
	BlockFactionChangeWhileOvert bool
	BlockFactionChangeWhileAggro  bool
	// D-05: standing resets to 0 on switch/leave.
	ResetStandingOnFactionChange bool
	// D-06: no standing awards in Phase 8; kills are logged only.
	AwardStanding bool

	// D-07: duel challenge expiry.
	DuelChallengeExpiry time.Duration
	// D-08: duel start range, and separation beyond which a duel ends.
	DuelStartRangeM       float64
	DuelEndRangeM         float64
	DuelSeparationGrace   time.Duration
	// D-09: duel timeout.
	DuelTimeout            time.Duration
	// D-10: guild roles permitted to declare/end war.
	WarDeclareRoles        []string
	// D-11: delay before a declared war becomes active.
	WarActivationDelay     time.Duration
	// D-13: war combat is legal anywhere in the world.
	WarAnywhere            bool
	// D-12: either leader may end a war; no automatic expiry.
	WarAnyLeaderMayEnd     bool
	WarAutomaticExpiry     time.Duration // 0 = never

	// D-15: recruiter NPC interaction radius.
	RecruiterInteractM     float64
}

// defaultConfig returns the approved (Section 3) production defaults.
func defaultConfig() PvPConfig {
	return PvPConfig{
		OvertDelay:                    30 * time.Second,
		CovertDelay:                   300 * time.Second,
		AggressionTimer:               120 * time.Second,
		BlockFactionChangeWhileOvert:  true,
		BlockFactionChangeWhileAggro:  true,
		ResetStandingOnFactionChange:  true,
		AwardStanding:                 false, // D-06
		DuelChallengeExpiry:           30 * time.Second,
		DuelStartRangeM:               64.0,
		DuelEndRangeM:                 128.0,
		DuelSeparationGrace:           10 * time.Second,
		DuelTimeout:                   10 * time.Minute,
		WarDeclareRoles:               []string{"leader", "officer"},
		WarActivationDelay:            24 * time.Hour,
		WarAnywhere:                   true,
		WarAnyLeaderMayEnd:            true,
		WarAutomaticExpiry:            0, // never
		RecruiterInteractM:            8.0,
	}
}

// fastConfig returns the compressed test variant used when
// TESTBED_FAST_CYCLE=1. Same semantics, scaled durations.
func fastConfig() PvPConfig {
	c := defaultConfig()
	c.OvertDelay = 2 * time.Second
	c.CovertDelay = 3 * time.Second
	c.AggressionTimer = 2 * time.Second
	c.DuelChallengeExpiry = 2 * time.Second
	c.DuelSeparationGrace = 2 * time.Second
	c.DuelTimeout = 10 * time.Second
	c.WarActivationDelay = 3 * time.Second
	return c
}

// fastCycle reports whether test time compression is active. Mirrors the
// convention used by internal/resources, internal/economy, and internal/civic.
func fastCycle() bool { return os.Getenv("TESTBED_FAST_CYCLE") == "1" }

// Cfg is the active configuration. It is set once at init from the environment
// and may be swapped in tests via withOverriddenConfig.
var Cfg = func() PvPConfig {
	if fastCycle() {
		return fastConfig()
	}
	return defaultConfig()
}()

// withOverriddenConfig swaps Cfg for the duration of fn (test-only helper).
func withOverriddenConfig(c PvPConfig, fn func()) {
	prev := Cfg
	Cfg = c
	defer func() { Cfg = prev }()
	fn()
}
