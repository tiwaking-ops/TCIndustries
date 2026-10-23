package resources

import (
	"math/rand"
	"time"
)

// Spawn represents a resource spawn deposit in the world (GDD 11.2.1).
type Spawn struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Category    string    `json:"category"`     // e.g. "ferric_metal", "fibrous_flora"
	CenterX     float64   `json:"center_x"`     // world X coordinate
	CenterZ     float64   `json:"center_z"`     // world Z coordinate
	Radius      float64   `json:"radius"`       // spread radius (5000-20000m)
	PeakConc    int       `json:"peak_conc"`    // peak concentration (50-90%)
	QualityMean float64   `json:"quality_mean"` // Gauss(500, 150) quality stat
	ExpiresAt   time.Time `json:"expires_at"`   // 7-14 days lifespan
}

// GenerateInitialSpawns creates initial resource spawns for zone-0001 (GDD 11.2.1).
func GenerateInitialSpawns() []Spawn {
	categories := []string{
		"ferric_metal",
		"conductive_alloy",
		"structural_polymer",
		"cultured_organic",
		"industrial_chemical",
		"fibrous_flora",
		"filtered_water",
	}

	var spawns []Spawn
	rand.Seed(time.Now().UnixNano())

	for i, cat := range categories {
		spawns = append(spawns, Spawn{
			ID:          string(rune('a' + i)) + "-spawn-01",
			Name:        "Resource " + cat,
			Category:    cat,
			CenterX:     float64(rand.Intn(4000) - 2000),
			CenterZ:     float64(rand.Intn(4000) - 2000),
			Radius:      5000.0,
			PeakConc:    75,
			QualityMean: 500.0 + rand.NormFloat64()*150.0,
			ExpiresAt:   time.Now().Add(14 * 24 * time.Hour),
		})
	}
	return spawns
}
