// Package civic implements Phase 7 (Civic Systems) constants: city ranks and
// thresholds, upkeep magnitudes, tax/dues bounds, mail caps, group caps, chat
// channel vocabulary, and the standard emote library. Testbed fork; generic only.
//
// GDD-given values are cited inline. Everything else is [PROVISIONAL] under the
// owner-approved convention (proposal §6, owner decisions 2026-09-15): magnitudes
// anchored to fork-local scales (house/vendor upkeep, training-cost bands) to
// avoid inventing fresh tuning; all NON-CANONICAL, flagged here and in
// HYGIENE_NOTE.md. Fast-cycle mappings are test-config-only; defaults stay
// GDD-given (same convention as resources.HarvestTickInterval).
package civic

import (
	"os"
	"time"
)

// City ranks (GDD 14.2.2). Metropolis exists as a data shape only: rank
// evaluation caps at city per the approved plan (owner decision 4).
const (
	RankOutpost    = "outpost"
	RankTownship   = "township"
	RankCity       = "city"
	RankMetropolis = "metropolis" // data shape only; unreachable (Expansion)
)

// City statuses.
const (
	CityForming   = "forming"   // hall placed, structure threshold unmet, grace running
	CityActive    = "active"    // threshold met
	CityDissolved = "dissolved" // treasury failure (rank 1) or grace expiry (forming)
)

// RankStructureThresholds: structures-in-radius needed (GDD 14.2.2: 10/20/40,
// [ASSUMPTION]-authoritative; the hall itself counts toward the total —
// flagged simplification of 14.2.1's "10 structures plus hall" reading).
// Fast-cycle compresses counts the way lifespans were compressed in Phases 4–5:
// the evaluation machinery is identical, only constants shrink (test-only).
var (
	FoundThreshold   = 10
	TownshipThreshold = 20
	CityThreshold    = 40
)

// CityRadiusM: structures within this radius of the hall count as city
// structures [PROVISIONAL — GDD gives a "defined radius" without metres].
const CityRadiusM = 100.0

// City upkeep per week by rank [PROVISIONAL — anchored to the house upkeep
// scale (200/500/1000): a Rank-1 city costs about a medium house to run].
var CityUpkeepWeekly = map[string]int{
	RankOutpost:  500,
	RankTownship: 1000,
	RankCity:     2000,
}

// MayorDiscountPct: upkeep reduction while the city has an active mayor.
// Owner decision 3 (mayor-status-gated upkeep discount — the lightweight
// Politician-perk substitute; full Politician profession stays deferred).
// [PROVISIONAL percentage]
const MayorDiscountPct = 20

// EffectiveUpkeep returns the weekly upkeep after the mayor discount.
func EffectiveUpkeep(rank string, hasMayor bool) int {
	base, ok := CityUpkeepWeekly[rank]
	if !ok {
		return 0
	}
	if hasMayor {
		return base * (100 - MayorDiscountPct) / 100
	}
	return base
}

// TreasuryReserveMult: rank-up requires treasury >= mult × target upkeep
// (GDD 14.2.2 "treasury funded 30 days" ≈ 4× weekly — the window itself is
// GDD-silent; 4× encodes the one-month-funded shape). [PROVISIONAL]
const TreasuryReserveMult = 4

// Tax bounds [PROVISIONAL — GDD gives flat-or-percent shape without numbers].
const (
	MaxFlatTaxWeekly  = 500
	MaxVendorTaxPct   = 20
	MaxDuesPct        = 20
)

// Election timing: weekly voting period + 7-day term (both [ASSUMPTION]-tagged
// in GDD 14.2.3). Fast-cycle compresses the period so tests can observe
// expiry-close; the term stays long in fast-cycle too so an elected test mayor
// does not expire mid-suite (flagged asymmetry, test-only).
var (
	ElectionPeriod = 7 * 24 * time.Hour
	MayorTerm      = 7 * 24 * time.Hour
	// FormingGrace: founder-placed hall with too few structures activates if
	// the count is reached within this window (GDD 14.2.1 "grace period" —
	// duration GDD-silent). [PROVISIONAL]
	FormingGrace = 7 * 24 * time.Hour
	// MayorInactivity: mayor unseen this long triggers an emergency election
	// (GDD 14.5, 14-day [ASSUMPTION]). Presence = world entry (last_seen).
	MayorInactivity = 14 * 24 * time.Hour
	// ElectionCooldown: minimum gap between elections in one city.
	// [PROVISIONAL — GDD-silent; prevents election spam]
	ElectionCooldown = 7 * 24 * time.Hour
)

var fastCycle = os.Getenv("TESTBED_FAST_CYCLE") == "1"

func init() {
	if fastCycle {
		FoundThreshold = 4
		TownshipThreshold = 6
		CityThreshold = 8
		ElectionPeriod = 2 * time.Minute
		MayorTerm = 60 * time.Minute
		FormingGrace = 3 * time.Minute
		MayorInactivity = 45 * time.Minute
		ElectionCooldown = 30 * time.Second
	}
}

// FastCycle reports whether test time compression is active.
func FastCycle() bool { return fastCycle }

// Secs converts a wall-clock duration to whole seconds for INTEGER-unix
// deadline columns.
func Secs(d time.Duration) int64 { return int64(d / time.Second) }

// Guild rules.
const (
	// GuildMinFounders: minimum distinct founders (GDD 19.2.4 [ASSUMPTION] 3).
	GuildMinFounders = 3
	// GuildRegistrarCost: founding fee credit sink (GDD 19.2.4 "credit cost" —
	// value GDD-silent). [PROVISIONAL — training-cost scale: above early
	// 10–50, below Master 5k–20k]
	GuildRegistrarCost = 1000
	// GuildDissolutionDays: countdown when no officers remain to inherit
	// leadership (GDD 19.5: 30-day countdown). Wall-clock; code-only path.
	GuildDissolutionDays = 30 * 24 * time.Hour
	// GroupCap: max group members (GDD 19.2.1 [ASSUMPTION] 20).
	GroupCap = 20
	// GroupLeaderTimeout: disconnect-without-transfer succession delay
	// (GDD 19.5: 2 minutes). Leave-action transfers are immediate (same rule
	// applied to an explicit leave — flagged); disconnect-timeout is code-only.
	GroupLeaderTimeout = 2 * time.Minute
)

// Loot rules (GDD 19.2.2). Stored per group; round_robin is the effective
// default (no loot pipeline exists yet beyond corpse/credit paths — recorded,
// the enum is data + selection, not distribution).
const (
	LootRoundRobin  = "round_robin"
	LootMasterLooter = "master_looter"
	LootNeedGreed   = "need_greed"
)

// ValidLootRule reports whether a loot rule name is known.
func ValidLootRule(r string) bool {
	return r == LootRoundRobin || r == LootMasterLooter || r == LootNeedGreed
}

// Guild roles.
const (
	GuildLeader = "leader"
	GuildOfficer = "officer"
	GuildMember = "member"
)

// Mail rules.
const (
	// MailCap: mailbox capacity (GDD 18.2.2 [ASSUMPTION] 50).
	MailCap = 50
	// MailBodyMax: max body characters [PROVISIONAL — GDD-silent].
	MailBodyMax = 2000
)

// Mentorship: 7-day bond (GDD 19.2.5 [ASSUMPTION]); protege threshold =
// fewer than 5 owned skill boxes at bond time (playtime/skill-point threshold
// exists in GDD without a number). [PROVISIONAL threshold]
const (
	MentorshipDays       = 7
	MentorshipMaxBoxes   = 5
	MentorshipXP         = 50 // both sides, "mentoring" pool [PROVISIONAL]
)

// Waypoint cap per character [PROVISIONAL — GDD-silent].
const WaypointCap = 50

// Chat channels (GDD 18.2.1; group/guild/tell added to the existing
// spatial/shout/planet set; faction exists with empty membership — Phase 8
// owns alignment; emote rides the spatial radius per 18.2.3 baseline).
const (
	ChatGroup   = "group"
	ChatGuild   = "guild"
	ChatTell    = "tell"
	ChatFaction = "faction"
	ChatEmote   = "emote"
)

// Standard emote library (GDD 18.2.3 baseline expression; generic English,
// server-validated; holoemotes/Image Designer stay EXPANSION-OUT).
var emotes = map[string]bool{
	"wave": true, "bow": true, "dance": true, "sit": true,
	"laugh": true, "cheer": true, "salute": true, "shrug": true,
	"point": true, "nod": true, "applaud": true, "yawn": true,
}

// ValidEmote reports whether a name is in the standard library.
func ValidEmote(e string) bool { return emotes[e] }
