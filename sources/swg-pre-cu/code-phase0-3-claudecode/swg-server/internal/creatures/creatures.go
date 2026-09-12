// Package creatures implements creature templates, live instances, and
// lairs per GDD Section 17 (Creature System) and Section 9.5 (Creature
// Combat). Templates are static reference data (same pattern as
// internal/skills/professions.go); instances and lairs are dynamic world
// state persisted via internal/database/combat_db.go.
package creatures

import "time"

// AIClass per GDD 17.2.1's behavior classes.
type AIClass string

const (
	AIPassive    AIClass = "passive"
	AIAggressive AIClass = "aggressive"
	AIScavenger  AIClass = "scavenger"
	AIHumanoid   AIClass = "humanoid"
)

// Template is static reference data for a creature type. GDD 28.3
// CreatureTemplate (simplified for Phase 3: full harvest-yield and
// taming-eligibility fields are Phase 4/8.3.7 concerns, not needed yet).
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
}

// Phase 3 starter bestiary. Two low-CL Tatooine creatures, both named in
// GDD Section 5.2.1's Tatooine creature list ("Worrt, Dewback, Tusken
// Raider") or in-universe-consistent with it (Womp Rat is the canonical
// lowest-tier Tatooine pest, in the same spirit as the GDD's examples,
// used here as the CL-1 tutorial-tier target the Phase 3 exit criteria
// needs: "fight and defeat a spawned creature").
// [ASSUMPTION] Exact HP/damage numbers for named SW creatures aren't given
// anywhere in the GDD (Section 9.7.1 only gives tier-level damage bands for
// player *weapons*); these are picked to sit clearly within a "trivial for
// a fresh character" range, consistent with Section 9.5.1's CL/difficulty
// framing.
func AllTemplates() []Template {
	return []Template{
		{
			ID: "womp_rat", Name: "Womp Rat", AIClass: AIAggressive,
			CLMin: 1, CLMax: 1, HealthMax: 60,
			DamageMin: 3, DamageMax: 8, AccuracyBase: 20,
			AggroRadiusM: 10,
		},
		{
			ID: "worrt", Name: "Worrt", AIClass: AIScavenger,
			CLMin: 2, CLMax: 3, HealthMax: 140,
			DamageMin: 6, DamageMax: 15, AccuracyBase: 25,
			AggroRadiusM: 8,
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

// Instance is a live, spawned creature. GDD 28.3 CreatureInstance.
type Instance struct {
	ID                string
	TemplateID        string
	LairID            string // "" if free-roaming (not lair-bound)
	Planet            string
	PosX, PosY, PosZ  float64
	HealthCurrent     int
	HealthMax         int
	State             string // "idle", "aggro", "dead"
	TargetCharacterID string // who it's retaliating against, "" if none
	CorpseExpiresAt   *time.Time
}

// CorpseWindow: how long a dead creature can be interacted with before
// despawning. GDD 17.2.2: "[ASSUMPTION] 10 minutes" (harvesting isn't
// implemented until Phase 4, but the despawn timer itself is a Phase 3
// world-simulation concern, so it's included now).
const CorpseWindow = 10 * time.Minute

// Lair spawns creature instances and can be destroyed. GDD 28.3 Lair,
// GDD 9.5.3.
type Lair struct {
	ID            string
	TemplateID    string
	Planet        string
	PosX, PosZ    float64
	LairHP        int
	LairHPMax     int
	MaxPopulation int
	DestroyedAt   *time.Time
}

// DistanceSq returns squared 2D distance (x/z plane) -- used for aggro-
// radius checks without a sqrt, matching the spatial grid's own convention
// in internal/world/spatial.go.
func DistanceSq(x1, z1, x2, z2 float64) float64 {
	dx := x1 - x2
	dz := z1 - z2
	return dx*dx + dz*dz
}
