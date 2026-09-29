// Package creatures implements generic creature templates, live instances, and
// lairs for Phase 3. Testbed fork; de-SWG'd — template IDs/names are generic
// serials (owner decision 2026-09-14, C3), NOT predecessor-setting creatures.
//
// PLACEMENT (owner decision 2026-09-14, C3): first lair at (10,0) ID 0001;
// each subsequent creature/lair increments x by 1 and serial by 1
// (e.g. second at (11,0) ID 0002). Serials are zero-padded 4-digit strings.
//
// STATS [PROVISIONAL] Group-C convention (owner-approved where owner supplied no
// values): tutorial-tier numbers only, NON-CANONICAL placeholders for review.
package creatures

import "time"

// AIClass behavior classes (generic behavioral terms, not setting content).
type AIClass string

const (
	AIPassive    AIClass = "passive"
	AIAggressive AIClass = "aggressive"
	AIScavenger  AIClass = "scavenger"
	AIHumanoid   AIClass = "humanoid"
)

// Template is static reference data for a creature type.
type Template struct {
	ID           string
	Name         string
	AIClass      AIClass
	CLMin, CLMax int
	HealthMax    int
	DamageMin    int
	DamageMax    int
	AccuracyBase int
	AggroRadiusM float64
	// Tamable flags Creature Handler eligibility (GDD 17.2.4; Phase 9).
	// Flagged test content: 0001 tamable, 0002 not.
	Tamable bool
}

// AllTemplates is the Phase 3 starter bestiary: two low-CL tutorial-tier
// generics. Stats are Group-C provisionals (see package note).
func AllTemplates() []Template {
	return []Template{
		{
			ID: "0001", Name: "Creature 0001", AIClass: AIAggressive,
			CLMin: 1, CLMax: 1, HealthMax: 60,
			DamageMin: 3, DamageMax: 8, AccuracyBase: 20,
			AggroRadiusM: 10, Tamable: true,
		},
		{
			ID: "0002", Name: "Creature 0002", AIClass: AIScavenger,
			CLMin: 2, CLMax: 3, HealthMax: 140,
			DamageMin: 6, DamageMax: 15, AccuracyBase: 25,
			AggroRadiusM: 8, Tamable: false,
		},
	}
}

// TemplateByID looks up a template, or nil if unknown.
func TemplateByID(id string) *Template {
	for _, t := range AllTemplates() {
		if t.ID == id {
			return &t
		}
	}
	return nil
}

// Instance is a live, spawned creature.
type Instance struct {
	ID                string
	TemplateID        string
	LairID            string // "" if free-roaming (not lair-bound)
	Zone              string
	PosX, PosY, PosZ  float64
	HealthCurrent     int
	HealthMax         int
	State             string // "idle", "aggro", "dead"
	TargetCharacterID string // who it's retaliating against, "" if none
	CorpseExpiresAt   *time.Time
}

// CorpseWindow: GDD-given 10 minutes (predecessor §17.2.2 [ASSUMPTION]-tagged);
// the despawn timer is Phase 3 world-simulation scope, harvesting is Phase 4.
const CorpseWindow = 10 * time.Minute

// Lair spawns creature instances and can be destroyed.
type Lair struct {
	ID            string
	TemplateID    string
	Zone          string
	PosX, PosZ    float64
	LairHP        int
	LairHPMax     int
	MaxPopulation int
	DestroyedAt   *time.Time
}

// StarterLairs returns the two tutorial lairs at the owner-directed positions:
// (10,0) serial 0001, (11,0) serial 0002, in zone-0001.
func StarterLairs() []Lair {
	return []Lair{
		{ID: "0001", TemplateID: "0001", Zone: "zone-0001", PosX: 10, PosZ: 0,
			LairHP: 200, LairHPMax: 200, MaxPopulation: 3},
		{ID: "0002", TemplateID: "0002", Zone: "zone-0001", PosX: 11, PosZ: 0,
			LairHP: 300, LairHPMax: 300, MaxPopulation: 3},
	}
}

// DistanceSq returns squared 2D distance (x/z plane), matching the spatial
// grid's convention in internal/world/spatial.go.
func DistanceSq(x1, z1, x2, z2 float64) float64 {
	dx := x1 - x2
	dz := z1 - z2
	return dx*dx + dz*dz
}
