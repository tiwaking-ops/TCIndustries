package pvp

import (
	"time"
)

// PvPBasis names the rule that permitted an attack (proposal Section 5).
type PvPBasis string

const (
	BasisNone       PvPBasis = ""
	BasisDuel       PvPBasis = "duel"
	BasisWar        PvPBasis = "war"
	BasisFaction    PvPBasis = "faction"
	BasisBattlefield PvPBasis = "battlefield"
)

// FlagState is a faction member's PvP visibility flag (R-05).
type FlagState string

const (
	FlagCovert FlagState = "covert"
	FlagOvert  FlagState = "overt"
)

// CharacterPvPContext is everything the legality engine needs about one side
// of a potential attack. The handler builds it from the database and the
// authoritative spatial position; the engine itself does no I/O.
type CharacterPvPContext struct {
	CharacterID string
	FactionID   string // "" = neutral
	Flag        FlagState
	DuelID      string // active duel this character participates in, if any
	GuildID     string // "" = no guild
	ZoneID      string // battlefield zone containing the character, if any
}

// PvPDecision is the result of a legality check.
type PvPDecision struct {
	Allowed bool
	Basis   PvPBasis
	Reason  string // machine-readable rejection reason when !Allowed
}

// WarActive reports whether guild g1 and guild g2 are at active war. The
// handler injects this from the database so the engine stays I/O-free.
type WarActiveFunc func(g1, g2 string) bool

// CanAttackPlayer decides whether attacker a may attack target t right now.
//
// Evaluation order (proposal Section 5 — first match that allows wins):
//  1. self / same-account → reject ("self")
//  2. duel       → allow ("duel")
//  3. war        → allow ("war")
//  4. battlefield→ allow ("battlefield")
//  5. faction    → allow ("faction")
//  6. otherwise → reject, with a specific reason where one applies.
//
// warActive is the injected guild-war lookup; it may be nil (no wars).
//
// PvE is unchanged and never reaches this function.
func CanAttackPlayer(now time.Time, a, t CharacterPvPContext, warActive WarActiveFunc) PvPDecision {
	// 1. Self.
	if a.CharacterID == t.CharacterID {
		return PvPDecision{Reason: "self"}
	}

	// 2. Duel — both participants of the same active duel.
	if a.DuelID != "" && a.DuelID == t.DuelID {
		return PvPDecision{Allowed: true, Basis: BasisDuel}
	}

	// 3. War — guilds actively at war with each other.
	if a.GuildID != "" && t.GuildID != "" && a.GuildID != t.GuildID &&
		warActive != nil && warActive(a.GuildID, t.GuildID) {
		return PvPDecision{Allowed: true, Basis: BasisWar}
	}

	// 4. Battlefield — same active zone, both factioned, factions differ.
	if a.ZoneID != "" && a.ZoneID == t.ZoneID {
		if a.FactionID != "" && t.FactionID != "" && a.FactionID != t.FactionID {
			return PvPDecision{Allowed: true, Basis: BasisBattlefield}
		}
		return PvPDecision{Reason: "battlefield_faction_rule"}
	}

	// 5. Faction — different factions, both overt.
	if a.FactionID != "" && t.FactionID != "" {
		if a.FactionID == t.FactionID {
			return PvPDecision{Reason: "same_faction"}
		}
		if a.Flag != FlagOvert {
			return PvPDecision{Reason: "attacker_covert"}
		}
		if t.Flag != FlagOvert {
			return PvPDecision{Reason: "target_covert"}
		}
		return PvPDecision{Allowed: true, Basis: BasisFaction}
	}

	// 6. No rule applied.
	if a.FactionID == "" || t.FactionID == "" {
		return PvPDecision{Reason: "neutral_target"}
	}
	return PvPDecision{Reason: "not_permitted"}
}

// RefreshAggression reports whether a character's aggression timer should be
// refreshed by this interaction: any faction member gains (or renews) the
// timer on a PvP attack made or received (D-03). Neutrals have no flag and so
// never carry a timer.
func RefreshAggression(now time.Time, c CharacterPvPContext) (time.Time, bool) {
	if c.FactionID == "" {
		return time.Time{}, false
	}
	return now.Add(Cfg.AggressionTimer), true
}
