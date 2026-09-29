// Unit tests for elite entry gates (Phase 9). Pure logic — no server needed.
// Run: go test ./internal/skills/
package skills

import "testing"

// own builds an owned-boxes set from IDs.
func own(ids ...string) map[string]bool {
	m := map[string]bool{}
	for _, id := range ids {
		m[id] = true
	}
	return m
}

func TestEliteNoviceGates(t *testing.T) {
	cases := []struct {
		name string
		box  string
		have []string
		want bool
	}{
		{"armorsmith needs artisan master", "armorsmith_novice", nil, false},
		{"armorsmith open with master", "armorsmith_novice", []string{"artisan_master"}, true},
		{"weaponsmith open", "weaponsmith_novice", []string{"artisan_master"}, true},
		{"architect open", "architect_novice", []string{"artisan_master"}, true},
		{"merchant open", "merchant_novice", []string{"artisan_master"}, true},
		{"bio needs scout+medic-novice", "bioengineer_novice", []string{"scout_master"}, false},
		{"bio open", "bioengineer_novice", []string{"scout_master", "medic_novice"}, true},
		{"handler needs scout", "creaturehandler_novice", nil, false},
		{"handler open", "creaturehandler_novice", []string{"scout_master"}, true},
		{"ranger open", "ranger_novice", []string{"scout_master"}, true},
		{"pistoleer needs marksman", "pistoleer_novice", []string{"brawler_master"}, false},
		{"fencer needs brawler", "fencer_novice", []string{"marksman_master"}, false},
		{"doctor needs medic", "doctor_novice", []string{"medic_master"}, true},
		{"musician needs entertainer", "musician_novice", []string{"entertainer_master"}, true},
		{"imagedesigner needs entertainer", "imagedesigner_novice", []string{"entertainer_master"}, true},
		{"smuggler needs pistol+scout masters", "smuggler_novice", []string{"pistoleer_master"}, false},
		{"smuggler open", "smuggler_novice", []string{"pistoleer_master", "scout_master"}, true},
		{"bounty option1 marksman", "bountyhunter_novice", []string{"marksman_master"}, true},
		{
			"bounty option2 scout+4",
			"bountyhunter_novice",
			[]string{"scout_master", "artisan_master", "medic_master", "entertainer_master", "brawler_master"},
			true,
		},
		{
			"bounty option2 short (3 others)",
			"bountyhunter_novice",
			[]string{"scout_master", "artisan_master", "medic_master", "entertainer_master"},
			false,
		},
		{
			"commando short (2 others)",
			"commando_novice",
			[]string{"marksman_master", "artisan_master", "medic_master"},
			false,
		},
		{
			"commando open",
			"commando_novice",
			[]string{"marksman_master", "artisan_master", "medic_master", "scout_master"},
			true,
		},
		{
			"tka needs all four melee masters",
			"teraskasi_novice",
			[]string{"brawler_master", "fencer_master", "swordsman_master"},
			false,
		},
		{
			"tka open",
			"teraskasi_novice",
			[]string{"brawler_master", "fencer_master", "swordsman_master", "pikeman_master"},
			true,
		},
	}
	for _, c := range cases {
		ok, reason := PrerequisiteCheck(c.box, own(c.have...))
		if ok != c.want {
			t.Errorf("%s: got %v (%q), want %v", c.name, ok, reason, c.want)
		}
	}
}

func TestEliteCosts(t *testing.T) {
	for _, p := range AllProfessions() {
		if p.Category != CategoryElite {
			continue
		}
		want := 80
		if p.ID == "bountyhunter" || p.ID == "commando" || p.ID == "smuggler" || p.ID == "teraskasi" {
			want = 120
		}
		if p.TotalSkillPoints != want {
			t.Errorf("%s totals %d, want %d", p.ID, p.TotalSkillPoints, want)
		}
		if len(p.Trees) != 4 {
			t.Errorf("%s has %d trees, want 4", p.ID, len(p.Trees))
		}
	}
	n := 0
	for _, p := range AllProfessions() {
		if p.Category == CategoryElite {
			n++
		}
	}
	if n != 22 {
		t.Errorf("elite roster = %d, want 22", n)
	}
}

func TestTrainerCoverage(t *testing.T) {
	for _, p := range AllProfessions() {
		if p.Category != CategoryElite {
			continue
		}
		base, ok := EliteTrainer[p.ID]
		if !ok || base == "" {
			t.Errorf("%s has no trainer coverage", p.ID)
		}
		if !TrainerCovers(base, p.ID) {
			t.Errorf("trainer %s does not cover %s", base, p.ID)
		}
		if TrainerCovers("medic", p.ID) && base != "medic" {
			t.Errorf("medic over-covers %s", p.ID)
		}
	}
}
