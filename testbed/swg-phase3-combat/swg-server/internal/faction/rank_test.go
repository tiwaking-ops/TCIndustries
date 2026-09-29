// Unit tests for faction rank thresholds (Phase 8/9 shared logic).
// Run: go test ./internal/faction/ (without TESTBED_FAST_CYCLE for defaults).
package faction

import "testing"

func TestRankForDefaults(t *testing.T) {
	if FastCycle() {
		t.Skip("fast-cycle thresholds active; defaults asserted in normal env")
	}
	cases := []struct {
		points int
		want   string
	}{
		{0, RankRecruit}, {100, RankRecruit}, {2499, RankRecruit},
		{2500, RankSergeant}, {9999, RankSergeant},
		{10000, RankMajor}, {29999, RankMajor},
		{30000, RankColonel}, {100000, RankColonel}, // General never assigned
	}
	for _, c := range cases {
		if got := RankFor(c.points); got != c.want {
			t.Errorf("RankFor(%d) = %q, want %q", c.points, got, c.want)
		}
	}
}
