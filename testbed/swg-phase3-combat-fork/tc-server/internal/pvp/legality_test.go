package pvp

import (
	"testing"
	"time"
)

func TestCanAttackPlayer_LegalityMatrix(t *testing.T) {
	now := time.Now()

	cases := []struct {
		name        string
		a, t        CharacterPvPContext
		wantAllowed bool
		wantBasis   PvPBasis
		wantReason  string
	}{
		{
			name:        "self attack rejected",
			a:           CharacterPvPContext{CharacterID: "c1"},
			t:           CharacterPvPContext{CharacterID: "c1"},
			wantAllowed: false, wantReason: "self",
		},
		{
			name: "neutral vs neutral rejected",
			a:    CharacterPvPContext{CharacterID: "c1"},
			t:    CharacterPvPContext{CharacterID: "c2"},
			wantAllowed: false, wantReason: "neutral_target",
		},
		{
			name: "same faction rejected",
			a:    CharacterPvPContext{CharacterID: "c1", FactionID: "f-a", Flag: FlagOvert},
			t:    CharacterPvPContext{CharacterID: "c2", FactionID: "f-a", Flag: FlagOvert},
			wantAllowed: false, wantReason: "same_faction",
		},
		{
			name: "covert attacker rejected",
			a:    CharacterPvPContext{CharacterID: "c1", FactionID: "f-a", Flag: FlagCovert},
			t:    CharacterPvPContext{CharacterID: "c2", FactionID: "f-b", Flag: FlagOvert},
			wantAllowed: false, wantReason: "attacker_covert",
		},
		{
			name: "covert target rejected",
			a:    CharacterPvPContext{CharacterID: "c1", FactionID: "f-a", Flag: FlagOvert},
			t:    CharacterPvPContext{CharacterID: "c2", FactionID: "f-b", Flag: FlagCovert},
			wantAllowed: false, wantReason: "target_covert",
		},
		{
			name: "overt vs overt different factions allowed",
			a:    CharacterPvPContext{CharacterID: "c1", FactionID: "f-a", Flag: FlagOvert},
			t:    CharacterPvPContext{CharacterID: "c2", FactionID: "f-b", Flag: FlagOvert},
			wantAllowed: true, wantBasis: BasisFaction,
		},
		{
			name: "duel allows both participants",
			a:    CharacterPvPContext{CharacterID: "c1", DuelID: "d1"},
			t:    CharacterPvPContext{CharacterID: "c2", DuelID: "d1"},
			wantAllowed: true, wantBasis: BasisDuel,
		},
		{
			name: "duel does not grant third party permissions",
			a:    CharacterPvPContext{CharacterID: "c1", DuelID: "d1"},
			t:    CharacterPvPContext{CharacterID: "c3", DuelID: "d2"},
			wantAllowed: false, wantReason: "neutral_target",
		},
		{
			name: "war allows cross-guild combat",
			a:    CharacterPvPContext{CharacterID: "c1", GuildID: "g1"},
			t:    CharacterPvPContext{CharacterID: "c2", GuildID: "g2"},
			wantAllowed: true, wantBasis: BasisWar,
		},
		{
			name: "battlefield allows opposing factions in same zone",
			a:    CharacterPvPContext{CharacterID: "c1", FactionID: "f-a", Flag: FlagOvert, ZoneID: "z1"},
			t:    CharacterPvPContext{CharacterID: "c2", FactionID: "f-b", Flag: FlagOvert, ZoneID: "z1"},
			wantAllowed: true, wantBasis: BasisBattlefield,
		},
		{
			name: "battlefield same faction rejected",
			a:    CharacterPvPContext{CharacterID: "c1", FactionID: "f-a", Flag: FlagOvert, ZoneID: "z1"},
			t:    CharacterPvPContext{CharacterID: "c2", FactionID: "f-a", Flag: FlagOvert, ZoneID: "z1"},
			wantAllowed: false, wantReason: "battlefield_faction_rule",
		},
		{
			name: "battlefield neutral target rejected",
			a:    CharacterPvPContext{CharacterID: "c1", FactionID: "f-a", Flag: FlagOvert, ZoneID: "z1"},
			t:    CharacterPvPContext{CharacterID: "c2", ZoneID: "z1"},
			wantAllowed: false, wantReason: "battlefield_faction_rule",
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			warActive := func(g1, g2 string) bool {
				return g1 == "g1" && g2 == "g2"
			}
			got := CanAttackPlayer(now, tc.a, tc.t, warActive)
			if got.Allowed != tc.wantAllowed {
				t.Fatalf("Allowed = %v, want %v (reason %q)", got.Allowed, tc.wantAllowed, got.Reason)
			}
			if tc.wantAllowed && got.Basis != tc.wantBasis {
				t.Fatalf("Basis = %q, want %q", got.Basis, tc.wantBasis)
			}
			if !tc.wantAllowed && got.Reason != tc.wantReason {
				t.Fatalf("Reason = %q, want %q", got.Reason, tc.wantReason)
			}
		})
	}
}

func TestRefreshAggression(t *testing.T) {
	now := time.Now()

	// Faction member: timer refreshes to now + AggressionTimer.
	member := CharacterPvPContext{CharacterID: "c1", FactionID: "f-a"}
	exp, ok := RefreshAggression(now, member)
	if !ok {
		t.Fatal("faction member should refresh aggression")
	}
	want := now.Add(Cfg.AggressionTimer)
	if !exp.Equal(want) {
		t.Fatalf("expiry = %v, want %v", exp, want)
	}

	// Neutral: never carries a timer.
	neutral := CharacterPvPContext{CharacterID: "c2"}
	if _, ok := RefreshAggression(now, neutral); ok {
		t.Fatal("neutral character must not refresh aggression")
	}
}

func TestFastCycleCompression(t *testing.T) {
	fc := fastConfig()
	def := defaultConfig()
	if fc.OvertDelay >= def.OvertDelay {
		t.Fatalf("fast-cycle OvertDelay %v not compressed below default %v", fc.OvertDelay, def.OvertDelay)
	}
	if fc.AggressionTimer >= def.AggressionTimer {
		t.Fatalf("fast-cycle AggressionTimer %v not compressed", fc.AggressionTimer)
	}
	if fc.WarActivationDelay >= def.WarActivationDelay {
		t.Fatalf("fast-cycle WarActivationDelay %v not compressed", fc.WarActivationDelay)
	}
}

func TestContainmentTest(t *testing.T) {
	circle := Zone{Shape: ZoneCircle, CenterX: 100, CenterZ: 100, Radius: 50, Active: true}
	if !circle.Contains(120, 100) {
		t.Error("point inside circle should be contained")
	}
	if circle.Contains(200, 100) {
		t.Error("point outside circle should not be contained")
	}

	rect := Zone{Shape: ZoneRect, MinX: 0, MinZ: 0, MaxX: 100, MaxZ: 100, Active: true}
	if !rect.Contains(50, 50) {
		t.Error("point inside rect should be contained")
	}
	if rect.Contains(150, 50) {
		t.Error("point outside rect should not be contained")
	}
}
