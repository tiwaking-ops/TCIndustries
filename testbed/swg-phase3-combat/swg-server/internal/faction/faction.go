// Package faction implements Phase 8 (Faction & Advanced PvP) constants:
// generic alignments, flagging timers, rank thresholds, point awards, and
// base siege magnitudes. Testbed fork; generic only.
//
// Alignments are deliberately flavorless (alignment_a / alignment_b) per
// HD-TST-01: no Star Wars names, lore, or factions may enter the fork, so the
// testbed cannot be mistaken for canon setting content.
//
// GDD-given values are cited inline. Everything else is [PROVISIONAL] under the
// owner-approved convention (proposal §6, owner decisions 2026-09-15): magnitudes
// anchored to fork-local scales to avoid inventing fresh tuning; all
// NON-CANONICAL, flagged here and in HYGIENE_NOTE.md. Fast-cycle mappings are
// test-config-only; defaults stay GDD-given.
package faction

import (
	"os"
	"time"
)

// Alignments (GDD 15.2.1 Rebel/Imperial/Neutral, genericized for the fork).
const (
	AlignNeutral = "neutral"
	AlignA       = "alignment_a"
	AlignB       = "alignment_b"
)

// ValidAlignment reports whether a name is a known alignment.
func ValidAlignment(a string) bool {
	return a == AlignNeutral || a == AlignA || a == AlignB
}

// Opposing reports whether two non-neutral alignments oppose each other.
func Opposing(a, b string) bool {
	return (a == AlignA && b == AlignB) || (a == AlignB && b == AlignA)
}

// Ranks (GDD 15.2.3 titles are generic military English, kept verbatim;
// General + sponsorship are EXPANSION-OUT — promotion caps at Colonel).
const (
	RankRecruit  = "recruit"
	RankSergeant = "sergeant"
	RankMajor    = "major"
	RankColonel  = "colonel"
)

// RankThresholds maps rank → cumulative points (GDD 15.2.3 [HISTORICAL,
// ASSUMED] exact thresholds). Fast-cycle compresses them so Sergeant→Colonel
// execute live in tests; defaults stay GDD-given (same compression precedent
// as city founding counts).
var RankThresholds = map[string]int{
	RankRecruit:  0,
	RankSergeant: 2500,
	RankMajor:    10000,
	RankColonel:  30000,
}

// RankOrder lists promotable ranks ascending.
var RankOrder = []string{RankRecruit, RankSergeant, RankMajor, RankColonel}

// RankFor returns the highest rank whose threshold the points meet.
func RankFor(points int) string {
	rank := RankRecruit
	for _, r := range RankOrder {
		if points >= RankThresholds[r] {
			rank = r
		}
	}
	return rank
}

// Point awards [PROVISIONAL — GDD gives no per-kill/destroy amounts].
const (
	// KillAward: faction points per PvP kill (incap credit runs through death;
	// see pending-kill design in faction_db.go).
	KillAward = 100
	// DestroyAward: faction points per base-destroy participant.
	DestroyAward = 250
)

// BaseHP: faction base structure HP [PROVISIONAL — GDD gives the HP mechanic
// without numbers; ≈90 s solo siege at crafted-sidearm scale].
const BaseHP = 3000

// Flagging timers.
var (
	// CovertDelay: Overt→Covert switch delay (GDD 9.4.1 5-min anti-combat-log
	// timer). Fast-cycle compresses so the refusal + success paths both
	// execute live.
	CovertDelay = 5 * time.Minute
	// SwitchCooldown: alignment-change cooldown (GDD 15.2.1 [ASSUMPTION] 30
	// days). NOT compressed: the live-tested path is the refusal itself, so
	// the cooldown must exceed any test session. Moving TO neutral is always
	// allowed (leaving is safe by §15.1 opt-in spirit); declaring a
	// non-neutral alignment within the cooldown of the last change is refused.
	SwitchCooldown = 30 * 24 * time.Hour
	// PointsDecayAfter: rank-point inactivity decay (GDD 15.2.3 [ASSUMPTION]
	// 60+ days). Code-only path (durations untestable live).
	PointsDecayAfter = 60 * 24 * time.Hour
)

var fastCycle = os.Getenv("TESTBED_FAST_CYCLE") == "1"

func init() {
	if fastCycle {
		RankThresholds[RankSergeant] = 100
		RankThresholds[RankMajor] = 200
		RankThresholds[RankColonel] = 300
		CovertDelay = 30 * time.Second
	}
}

// FastCycle reports whether test time compression is active.
func FastCycle() bool { return fastCycle }

// Secs converts a wall-clock duration to whole seconds for INTEGER-unix columns.
func Secs(d time.Duration) int64 { return int64(d / time.Second) }

// Base placement.
const (
	// BaseNoBuildM: separation between bases and other siege/civic objects
	// [PROVISIONAL — reuses the 20 m structure rule; cities use radius overlap].
	BaseNoBuildM = 20.0
	// BaseAttackRangeM: siege range gate [PROVISIONAL — same basis as the 8 m
	// melee gate].
	BaseAttackRangeM = 8.0
)

// PvPCreditPct: share of carried credits transferred victim→killer on PvP
// death (GDD 9.4.2/9.6.2 "drop ... 10% of carried credits", resolved as a
// ledger-tagged transfer — proposal §5.d).
const PvPCreditPct = 10

// ConditionLoss: item-condition points lost on equipped items per PvP death
// (GDD 9.4.2 "double PvE death penalty" shape; PvE has no condition loss yet,
// so 10 points is the full MVP representation). Floor 0, no breakage, no
// repair (OQ-010 follow-ups, flagged). [PROVISIONAL]
const ConditionLoss = 10
