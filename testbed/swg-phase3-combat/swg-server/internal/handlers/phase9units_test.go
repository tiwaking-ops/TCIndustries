// Unit tests for Phase 9 pure helpers (no server needed).
// Run: go test ./internal/handlers/
package handlers

import "testing"

func TestMerchantDiscountedUpkeep(t *testing.T) {
	// GDD 8.3.5: up to 50% reduction at tier IV.
	cases := []struct {
		weekly int
		tier   int
		want   int
	}{
		{800, 0, 800},
		{800, 1, 700},
		{800, 2, 600},
		{800, 4, 400},
		{800, 9, 400},  // clamped
		{800, -3, 800}, // clamped
	}
	for _, c := range cases {
		if got := merchantDiscountedUpkeep(c.weekly, c.tier); got != c.want {
			t.Errorf("upkeep(%d, tier %d) = %d, want %d", c.weekly, c.tier, got, c.want)
		}
	}
}

func TestStylePool(t *testing.T) {
	cases := map[string]string{
		"pistol": "pistol_combat", "rifle": "rifle_combat",
		"carbine": "carbine_combat", "fencing": "fencing_combat",
		"sword": "sword_combat", "polearm": "polearm_combat",
		"bounty": "bounty_hunter", "commando": "commando",
		"unknown": "combat", "": "combat",
	}
	for style, want := range cases {
		if got := stylePool(style); got != want {
			t.Errorf("stylePool(%q) = %q, want %q", style, got, want)
		}
	}
}

func TestTameCaps(t *testing.T) {
	wantCL := [5]int{3, 8, 12, 16, 20}
	for i, w := range wantCL {
		if tameCLCaps[i] != w {
			t.Errorf("tameCLCaps[%d] = %d, want %d", i, tameCLCaps[i], w)
		}
	}
	if tameCLCaps[4] != 20 {
		t.Errorf("master CL cap must be 20 (GDD 8.3.7)")
	}
	wantPets := [5]int{1, 1, 2, 2, 3}
	for i, w := range wantPets {
		if petCapByTier[i] != w {
			t.Errorf("petCapByTier[%d] = %d, want %d", i, petCapByTier[i], w)
		}
	}
}
