// Unit tests for Phase 9 service formulas (no server needed).
// Run: go test ./internal/services/
package services

import "testing"

func TestWoundHealAmount(t *testing.T) {
	// Base formula: 25 + 25×tier (elite boosts apply at the call site).
	cases := []struct {
		tier int
		want int
	}{
		{0, 25}, {1, 50}, {2, 75}, {4, 125}, {9, 125}, {-2, 25},
	}
	for _, c := range cases {
		if got := WoundHealAmount(c.tier); got != c.want {
			t.Errorf("WoundHealAmount(%d) = %d, want %d", c.tier, got, c.want)
		}
	}
}

func TestBuffMagnitude(t *testing.T) {
	// GDD band floor scaled: 500 + 250×tier.
	if got := BuffMagnitude(0); got != 500 {
		t.Errorf("BuffMagnitude(0) = %d, want 500", got)
	}
	if got := BuffMagnitude(4); got != 1500 {
		t.Errorf("BuffMagnitude(4) = %d, want 1500", got)
	}
}
