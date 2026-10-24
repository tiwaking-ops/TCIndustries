// Package resources implements the Phase 4 resource-spawn service: generic
// taxonomy, stat generation, and lifecycle management. Testbed fork; generic only.
//
// SOURCE MODEL: predecessor-GDD 11.2.1 spawn pseudocode, implemented literally
// (gauss mean 500 / stddev 150, clamp 0–1000, OQ-as-mean, radius band, peak
// concentration band, relocate-on-despawn). UNIT ADAPTATION [ASSUMPTION]: the GDD
// specifies radii in KILOMETRES on full planets; this single-zone testbed works in
// METRES, so the 5–20 band is read as metres. Flagged, not silently converted.
// TIME COMPRESSION [ASSUMPTION, test scaffolding]: 7–14-day lifecycles and hourly
// harvest ticks are untestable live — package vars default to GDD values and are
// overridden ONLY when env TESTBED_FAST_CYCLE=1 (lifespan 120–180 s, harvest tick
// 20 s). Defaults stay GDD-given; compression never ships in default constants.
package resources

import (
	"math"
	"math/rand"
	"os"
	"time"
)

// Stat lanes (generic vocabulary). OQ is always the mean of the relevant lanes.
const (
	LaneOQ           = "OQ"
	LaneDurability   = "durability"
	LaneConductivity = "conductivity"
	LaneMalleability = "malleability"
	LanePotency      = "potency"
	LanePurity       = "purity"
	LaneTexture      = "texture"
	LaneNutrition    = "nutrition"
	LaneVolatility   = "volatility"
	LaneSolubility   = "solubility"
)

// ResourceType is one spawnable generic resource class.
type ResourceType struct {
	ID            string
	Name          string
	Category      string // metal | polymer | organic | chemical | flora | water
	RelevantLanes []string
}

// Types is the MVP generic taxonomy (owner-approved derivation, proposal §6.3:
// GDD 10.4.1 categories minus setting names; 7 types within the 5–10 band).
var Types = []ResourceType{
	{ID: "ferric_metal", Name: "Ferric Metal", Category: "metal",
		RelevantLanes: []string{LaneDurability, LaneMalleability, LaneConductivity, LanePurity}},
	{ID: "conductive_alloy", Name: "Conductive Alloy", Category: "metal",
		RelevantLanes: []string{LaneConductivity, LaneDurability, LanePurity, LaneMalleability}},
	{ID: "structural_polymer", Name: "Structural Polymer", Category: "polymer",
		RelevantLanes: []string{LaneDurability, LaneTexture, LaneMalleability, LaneVolatility}},
	{ID: "cultured_organic", Name: "Cultured Organic", Category: "organic",
		RelevantLanes: []string{LaneNutrition, LanePotency, LaneTexture, LanePurity}},
	{ID: "industrial_chemical", Name: "Industrial Chemical", Category: "chemical",
		RelevantLanes: []string{LanePotency, LaneVolatility, LanePurity, LaneSolubility}},
	{ID: "fibrous_flora", Name: "Fibrous Flora", Category: "flora",
		RelevantLanes: []string{LaneTexture, LaneDurability, LaneNutrition, LaneSolubility}},
	{ID: "filtered_water", Name: "Filtered Water", Category: "water",
		RelevantLanes: []string{LanePurity, LaneSolubility, LaneNutrition, LanePotency}},
}

// TypeByID looks up a resource type, or nil if unknown.
func TypeByID(id string) *ResourceType {
	for _, t := range Types {
		if t.ID == id {
			return &t
		}
	}
	return nil
}

// ToolCategories lists the 5 GDD survey-tool categories.
func ToolCategories() []string {
	return []string{"mineral", "chemical", "flora", "organic", "water"}
}

// ToolCategory maps the 5 GDD survey-tool categories onto generic types.
func ToolCategory(tool string) []string {
	switch tool {
	case "mineral":
		return []string{"ferric_metal", "conductive_alloy"}
	case "chemical":
		return []string{"industrial_chemical", "structural_polymer"}
	case "flora":
		return []string{"fibrous_flora"}
	case "organic":
		return []string{"cultured_organic"}
	case "water":
		return []string{"filtered_water"}
	default:
		return nil
	}
}

// --- Time configuration (GDD defaults; test compression via env) ---

var (
	// SpawnLifeMin/Max: GDD 7–14 days.
	SpawnLifeMin = 7 * 24 * time.Hour
	SpawnLifeMax = 14 * 24 * time.Hour
	// HarvestTickInterval: GDD hourly extraction.
	HarvestTickInterval = 1 * time.Hour
	// SurveyCooldown: GDD-given 10 s (NOT compressed).
	SurveyCooldown = 10 * time.Second
	// GuaranteeSpawn: extra deterministic test spawn when compressed (see SeedZone).
	fastCycle = os.Getenv("TESTBED_FAST_CYCLE") == "1"
)

func init() {
	if fastCycle {
		SpawnLifeMin = 120 * time.Second
		SpawnLifeMax = 180 * time.Second
		HarvestTickInterval = 20 * time.Second
	}
}

// FastCycle reports whether test time compression is active.
func FastCycle() bool { return fastCycle }

// Spawn is one live resource spawn.
type Spawn struct {
	ID                string
	Type              string
	Zone              string
	CenterX, CenterZ  float64
	RadiusM           float64
	PeakConcentration int // 50–90 (GDD 11.2.1)
	Stats             map[string]int
	SpawnedAt         time.Time
	DespawnsAt        time.Time
}

// gauss returns a N(mean, stddev) sample via Box–Muller.
func gauss(rng *rand.Rand, mean, stddev float64) float64 {
	u1 := rng.Float64()
	u2 := rng.Float64()
	if u1 < 1e-9 {
		u1 = 1e-9
	}
	z := math.Sqrt(-2*math.Log(u1)) * math.Cos(2*math.Pi*u2)
	return mean + stddev*z
}

func clamp0_1000(v float64) int {
	if v < 0 {
		return 0
	}
	if v > 1000 {
		return 1000
	}
	return int(v + 0.5)
}

// GenerateSpawn creates one spawn of the given type: Gaussian stats per lane,
// OQ as the lane mean, radius/peak/lifespan per the GDD bands.
func GenerateSpawn(rng *rand.Rand, id, rtype, zone string, cx, cz float64, now time.Time) *Spawn {
	t := TypeByID(rtype)
	stats := make(map[string]int)
	sum := 0
	for _, lane := range t.RelevantLanes {
		v := clamp0_1000(gauss(rng, 500, 150))
		stats[lane] = v
		sum += v
	}
	stats[LaneOQ] = sum / len(t.RelevantLanes)
	life := SpawnLifeMin + time.Duration(rng.Int63n(int64(SpawnLifeMax-SpawnLifeMin)))
	return &Spawn{
		ID: id, Type: rtype, Zone: zone, CenterX: cx, CenterZ: cz,
		RadiusM:           5 + rng.Float64()*15, // 5–20 band (metres, see package note)
		PeakConcentration: 50 + rng.Intn(41),   // 50–90
		Stats:             stats,
		SpawnedAt:         now,
		DespawnsAt:        now.Add(life),
	}
}

// ConcentrationAt returns the local concentration % at (x,z): peak at center,
// tapering toward ~10% of peak at the rim (GDD gradient shape, provisional curve).
func ConcentrationAt(s *Spawn, x, z float64) int {
	dx := x - s.CenterX
	dz := z - s.CenterZ
	d := math.Sqrt(dx*dx + dz*dz)
	if d >= s.RadiusM {
		return 0
	}
	frac := d / s.RadiusM
	conc := float64(s.PeakConcentration) * (0.1 + 0.9*math.Exp(-2.3*frac*frac))
	return int(conc + 0.5)
}
