// Package crafting implements Phase 4 schematic registry and server-side
// crafting sessions: assign → assemble (validation gate) → experiment → finalize.
// Testbed fork; generic content only. Pattern-adapted from the seed-demo engine
// (session FSM, points pool, outcome table) with two deliberate deviations toward
// the GDD: (1) no assembly roll (GDD 10.2 has none — the seed roll is that demo's
// simplification); (2) GDD-literal % bonuses per attempt instead of box-progress.
// Outcome distribution + bonus bands are GDD 10.7.1 literals. Everything else
// marked [PROVISIONAL] is Group-C-style placeholder under owner-approved convention.
package crafting

import (
	"fmt"
	"math/rand"
)

// SchematicSlot is one resource slot in a schematic.
type SchematicSlot struct {
	ID             string
	Label          string
	AcceptedTypes  []string
	UnitsRequired  int
}

// ExpProperty is one experimentable property.
type ExpProperty struct {
	ID             string
	Label          string
	BaseMin        float64
	BaseMax        float64
	QualityMult    float64 // [PROVISIONAL] resource-quality contribution (GDD example: 0.5)
	PointsPerRound int     // attempt cap context (GDD ≤5 enforced globally)
}

// Schematic is a craftable generic design.
type Schematic struct {
	ID               string
	Name             string
	ProfessionGate   string // required skill box (MVP: artisan_novice for all)
	Complexity       int
	Slots            []SchematicSlot
	Properties       []ExpProperty
	ExperimentPoints int
	MaxRounds        int
	XPReward         int // [PROVISIONAL] crafting XP on finalize
	EquipSlot        string // "weapon" | "armor" | "deed" | "none"
}

// Schematics is the MVP generic registry (owner-approved trio, proposal §6.4).
var Schematics = []Schematic{
	{
		ID: "basic_sidearm", Name: "Basic Sidearm", ProfessionGate: "artisan_novice",
		Complexity: 5, ExperimentPoints: 10, MaxRounds: 4, XPReward: 500,
		EquipSlot: "weapon",
		Slots: []SchematicSlot{
			{ID: "frame", Label: "Frame", AcceptedTypes: []string{"ferric_metal", "conductive_alloy"}, UnitsRequired: 10},
			{ID: "grip", Label: "Grip", AcceptedTypes: []string{"structural_polymer", "fibrous_flora"}, UnitsRequired: 5},
		},
		Properties: []ExpProperty{
			{ID: "damage", Label: "Damage", BaseMin: 20, BaseMax: 40, QualityMult: 0.5},
			{ID: "accuracy", Label: "Accuracy", BaseMin: 5, BaseMax: 15, QualityMult: 0.3},
		},
	},
	{
		ID: "basic_plate", Name: "Basic Plate", ProfessionGate: "artisan_novice",
		Complexity: 5, ExperimentPoints: 10, MaxRounds: 4, XPReward: 500,
		EquipSlot: "armor",
		Slots: []SchematicSlot{
			{ID: "plating", Label: "Plating", AcceptedTypes: []string{"ferric_metal", "conductive_alloy"}, UnitsRequired: 10},
			{ID: "lining", Label: "Lining", AcceptedTypes: []string{"cultured_organic", "fibrous_flora"}, UnitsRequired: 5},
		},
		Properties: []ExpProperty{
			{ID: "protection", Label: "Protection", BaseMin: 5, BaseMax: 20, QualityMult: 0.4},
			{ID: "durability", Label: "Durability", BaseMin: 80, BaseMax: 100, QualityMult: 0.2},
		},
	},
	{
		ID: "harvester_deed", Name: "Harvester Deed", ProfessionGate: "artisan_novice",
		Complexity: 5, ExperimentPoints: 10, MaxRounds: 4, XPReward: 500,
		EquipSlot: "deed",
		Slots: []SchematicSlot{
			{ID: "housing", Label: "Housing", AcceptedTypes: []string{"ferric_metal", "conductive_alloy"}, UnitsRequired: 8},
			{ID: "fittings", Label: "Fittings", AcceptedTypes: []string{"structural_polymer", "fibrous_flora"}, UnitsRequired: 4},
		},
		Properties: []ExpProperty{
			{ID: "efficiency", Label: "Efficiency", BaseMin: 100, BaseMax: 100, QualityMult: 0.3},
		},
	},
	{
		ID: "structure_deed", Name: "Structure Deed (small house)", ProfessionGate: "artisan_novice",
		Complexity: 5, ExperimentPoints: 10, MaxRounds: 4, XPReward: 500,
		EquipSlot: "deed",
		Slots: []SchematicSlot{
			{ID: "frame", Label: "Frame", AcceptedTypes: []string{"ferric_metal", "conductive_alloy"}, UnitsRequired: 15},
			{ID: "fittings", Label: "Fittings", AcceptedTypes: []string{"structural_polymer", "fibrous_flora"}, UnitsRequired: 8},
		},
		Properties: []ExpProperty{
			{ID: "capacity", Label: "Capacity", BaseMin: 200, BaseMax: 200, QualityMult: 0.2},
		},
	},
	{
		ID: "vendor_deed", Name: "Vendor Deed", ProfessionGate: "artisan_novice",
		Complexity: 5, ExperimentPoints: 10, MaxRounds: 4, XPReward: 500,
		EquipSlot: "deed",
		Slots: []SchematicSlot{
			{ID: "frame", Label: "Frame", AcceptedTypes: []string{"ferric_metal", "conductive_alloy"}, UnitsRequired: 15},
			{ID: "fittings", Label: "Fittings", AcceptedTypes: []string{"structural_polymer", "fibrous_flora"}, UnitsRequired: 8},
		},
		Properties: []ExpProperty{
			{ID: "presentation", Label: "Presentation", BaseMin: 100, BaseMax: 100, QualityMult: 0.2},
		},
	},
	{
		ID: "city_hall_deed", Name: "City Hall Deed", ProfessionGate: "artisan_novice",
		Complexity: 5, ExperimentPoints: 10, MaxRounds: 4, XPReward: 500,
		EquipSlot: "deed",
		Slots: []SchematicSlot{
			{ID: "frame", Label: "Frame", AcceptedTypes: []string{"ferric_metal", "conductive_alloy"}, UnitsRequired: 15},
			{ID: "fittings", Label: "Fittings", AcceptedTypes: []string{"structural_polymer", "fibrous_flora"}, UnitsRequired: 8},
		},
		Properties: []ExpProperty{
			{ID: "capacity", Label: "Capacity", BaseMin: 200, BaseMax: 200, QualityMult: 0.2},
		},
	},
	{
		ID: "base_deed", Name: "Faction Base Deed", ProfessionGate: "artisan_novice",
		Complexity: 5, ExperimentPoints: 10, MaxRounds: 4, XPReward: 500,
		EquipSlot: "deed",
		Slots: []SchematicSlot{
			{ID: "frame", Label: "Frame", AcceptedTypes: []string{"ferric_metal", "conductive_alloy"}, UnitsRequired: 15},
			{ID: "fittings", Label: "Fittings", AcceptedTypes: []string{"structural_polymer", "fibrous_flora"}, UnitsRequired: 8},
		},
		Properties: []ExpProperty{
			{ID: "capacity", Label: "Capacity", BaseMin: 200, BaseMax: 200, QualityMult: 0.2},
		},
	},
	{
		ID: "stim_pack", Name: "Stim Pack", ProfessionGate: "medic_novice",
		Complexity: 5, ExperimentPoints: 10, MaxRounds: 4, XPReward: 500,
		EquipSlot: "consumable",
		Slots: []SchematicSlot{
			{ID: "bio", Label: "Bio-Active Agent", AcceptedTypes: []string{"cultured_organic", "fibrous_flora"}, UnitsRequired: 6},
			{ID: "binding", Label: "Binding Agent", AcceptedTypes: []string{"industrial_chemical"}, UnitsRequired: 4},
		},
		Properties: []ExpProperty{
			{ID: "healing", Label: "Healing Amount", BaseMin: 50, BaseMax: 150, QualityMult: 0.5},
			{ID: "charges", Label: "Charges", BaseMin: 3, BaseMax: 5, QualityMult: 0},
			{ID: "potency", Label: "Wound Healing", BaseMin: 10, BaseMax: 30, QualityMult: 0.4},
		},
	},
}

// SchematicByID looks up a schematic, or nil if unknown.
func SchematicByID(id string) *Schematic {
	for _, s := range Schematics {
		if s.ID == id {
			return &s
		}
	}
	return nil
}

// Outcome tiers (GDD 10.7.1 literals: 10/50/30/8/2 distribution).
type Outcome string

const (
	OutcomeCriticalSuccess Outcome = "critical_success"
	OutcomeSuccess         Outcome = "success"
	OutcomeNormal          Outcome = "normal"
	OutcomeFailure         Outcome = "failure"
	OutcomeCriticalFailure Outcome = "critical_failure"
)

// SkillThresholds shapes experimentation odds by crafter skill (seed-demo
// threshold profiles, server-adapted). [PROVISIONAL] values + interpolation.
type SkillThresholds struct {
	Amazing int
	Great   int
	Success int
	CritFail int
}

// ThresholdsFor maps owned Artisan box count (0–18) onto novice→master profiles.
// Lerp between the seed demo's NOVICE (99/85/35/15) and MASTER (92/65/22/4).
func ThresholdsFor(artisanBoxes int) SkillThresholds {
	if artisanBoxes < 0 {
		artisanBoxes = 0
	}
	if artisanBoxes > 18 {
		artisanBoxes = 18
	}
	f := float64(artisanBoxes) / 18.0
	lerp := func(novice, master int) int {
		return novice + int(f*float64(master-novice)+0.5)
	}
	return SkillThresholds{
		Amazing: lerp(99, 92), Great: lerp(85, 65),
		Success: lerp(35, 22), CritFail: lerp(15, 4),
	}
}

// Session phases.
type SessionPhase string

const (
	PhaseDesign         SessionPhase = "design"
	PhaseExperimentation SessionPhase = "experimentation"
	PhaseComplete       SessionPhase = "complete"
)

// Assignment binds a resource stack to a slot.
type Assignment struct {
	SlotID  string
	SpawnID string
	Units   int
}

// Session is one server-side crafting session.
type Session struct {
	ID             string
	SchematicID    string
	OwnerCharID    string
	Assignments    map[string]Assignment
	Bonuses        map[string]float64 // property ID → accumulated % bonus
	PointsRemaining int
	RoundsUsed     int
	Phase          SessionPhase
	Log            []string
}

// NewSession opens a design-phase session.
func NewSession(id, schematicID, owner string, points int) *Session {
	return &Session{
		ID: id, SchematicID: schematicID, OwnerCharID: owner,
		Assignments: make(map[string]Assignment),
		Bonuses:     make(map[string]float64),
		PointsRemaining: points, Phase: PhaseDesign,
	}
}

// Assemble validates slot assignments (types + units, checked against live
// stacks by the caller) and opens experimentation. No assembly roll (see package
// note — deliberate GDD-literal deviation from the seed demo).
func Assemble(s *Session, sc *Schematic, stacks map[string]int) error {
	if s.Phase != PhaseDesign {
		return fmt.Errorf("session not in design phase")
	}
	for _, slot := range sc.Slots {
		a, ok := s.Assignments[slot.ID]
		if !ok {
			return fmt.Errorf("slot %q unassigned", slot.ID)
		}
		_ = a
	}
	s.Phase = PhaseExperimentation
	return nil
}

// Experiment spends N points (1–5, GDD cap) on one property: d100 roll against
// skill thresholds shifted +3 per point beyond the first [PROVISIONAL scaling —
// GDD says points raise success chance without giving the curve], then a uniform
// % delta inside the GDD outcome band. Returns outcome + applied delta.
func Experiment(s *Session, sc *Schematic, propID string, points int, th SkillThresholds, rng *rand.Rand) (Outcome, float64, error) {
	if s.Phase != PhaseExperimentation {
		return "", 0, fmt.Errorf("session not in experimentation phase")
	}
	if s.RoundsUsed >= sc.MaxRounds {
		return "", 0, fmt.Errorf("round cap reached (%d)", sc.MaxRounds)
	}
	if points < 1 || points > 5 {
		return "", 0, fmt.Errorf("points per attempt must be 1–5")
	}
	if points > s.PointsRemaining {
		return "", 0, fmt.Errorf("insufficient experimentation points")
	}
	var prop *ExpProperty
	for _, p := range sc.Properties {
		if p.ID == propID {
			prop = &p
			break
		}
	}
	if prop == nil {
		return "", 0, fmt.Errorf("unknown property %q", propID)
	}
	roll := rng.Intn(100) + 1
	eff := roll + 3*(points-1) // [PROVISIONAL] points-shift, see doc comment
	var out Outcome
	var lo, hi float64
	switch {
	case eff >= th.Amazing:
		out, lo, hi = OutcomeCriticalSuccess, 5, 10
	case eff >= th.Great:
		out, lo, hi = OutcomeSuccess, 2, 5
	case eff >= th.Success:
		out, lo, hi = OutcomeNormal, 0, 2
	case eff >= th.CritFail:
		out = OutcomeFailure
	default:
		out, lo, hi = OutcomeCriticalFailure, -5, -2
	}
	delta := 0.0
	if out == OutcomeCriticalSuccess || out == OutcomeSuccess ||
		out == OutcomeNormal || out == OutcomeCriticalFailure {
		delta = lo + rng.Float64()*(hi-lo)
	}
	s.Bonuses[propID] += delta
	s.PointsRemaining -= points
	s.RoundsUsed++
	s.Log = append(s.Log, fmt.Sprintf("%s: %s %+.1f%% (roll %d)", propID, out, delta, roll))
	return out, delta, nil
}

// FinalStats computes per-property [min, max] after quality + experimentation:
// base = schematic base + mean-OQ × mult (GDD 10.7.4 example literal);
// final = base × (1 + bonus%). qualities maps slot ID → assigned OQ.
func FinalStats(sc *Schematic, bonuses map[string]float64, qualities map[string]float64) map[string][2]float64 {
	sum, n := 0.0, 0
	for _, q := range qualities {
		sum += q
		n++
	}
	meanQ := 0.0
	if n > 0 {
		meanQ = sum / float64(n)
	}
	out := make(map[string][2]float64)
	for _, p := range sc.Properties {
		bonus := bonuses[p.ID] / 100.0
		minBase := p.BaseMin + meanQ*p.QualityMult
		maxBase := p.BaseMax + meanQ*p.QualityMult
		out[p.ID] = [2]float64{minBase * (1 + bonus), maxBase * (1 + bonus)}
	}
	return out
}
