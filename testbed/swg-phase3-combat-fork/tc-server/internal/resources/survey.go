package resources

import (
	"math"
	"math/rand"
	"time"
)

// SurveyTool represents a surveying tool category (GDD 10.5).
type SurveyTool struct {
	Category  string  // e.g., "metal", "organic", "energy"
	Precision float64 // radius in meters (basic: 500.0)
}

// SurveyResult represents the output of a survey scan.
type SurveyResult struct {
	ResourceName string  `json:"resource_name"`
	Category     string  `json:"category"`
	Concentration int    `json:"concentration"` // 0-100%
	Distance     float64 `json:"distance"`      // meters from player
	WaypointX    float64 `json:"waypoint_x"`    // approximate center with precision error
	WaypointZ    float64 `json:"waypoint_z"`
}

// SampleResult represents harvested resource sample units.
type SampleResult struct {
	ResourceID   string `json:"resource_id"`
	ResourceName string `json:"resource_name"`
	Amount       int    `json:"amount"`
	Quality      int    `json:"quality"`
}

// SurveyManager manages survey cooldowns and scan resolution.
type SurveyManager struct {
	lastSurveyTime map[string]time.Time // characterID -> timestamp
}

func NewSurveyManager() *SurveyManager {
	return &SurveyManager{
		lastSurveyTime: make(map[string]time.Time),
	}
}

// CanSurvey checks if survey is off cooldown (10s per GDD 10.5).
func (sm *SurveyManager) CanSurvey(characterID string) (bool, time.Duration) {
	last, ok := sm.lastSurveyTime[characterID]
	if !ok {
		return true, 0
	}
	elapsed := time.Since(last)
	cooldown := 10 * time.Second
	if elapsed < cooldown {
		return false, cooldown - elapsed
	}
	return true, 0
}

// RecordSurvey marks survey used.
func (sm *SurveyManager) RecordSurvey(characterID string) {
	sm.lastSurveyTime[characterID] = time.Now()
}

// FindNearestSpawn calculates survey results for a player at (px, pz) given active spawns.
func FindNearestSpawn(px, pz float64, spawns []Spawn, precision float64) *SurveyResult {
	if len(spawns) == 0 {
		return nil
	}

	var bestSpawn *Spawn
	minDist := math.MaxFloat64

	for i := range spawns {
		s := &spawns[i]
		dx := s.CenterX - px
		dz := s.CenterZ - pz
		dist := math.Sqrt(dx*dx + dz*dz)
		if dist < minDist {
			minDist = dist
			bestSpawn = s
		}
	}

	if bestSpawn == nil || minDist > 1000.0 { // outside basic survey range
		return nil
	}

	// Add precision error to waypoint per GDD 10.5 (±500m basic precision)
	errX := (rand.Float64() - 0.5) * precision
	errZ := (rand.Float64() - 0.5) * precision

	// Calculate concentration based on distance to spawn center
	// Max concentration at center, falling off to 0 at Radius
	conc := 0
	if minDist <= bestSpawn.Radius {
		factor := 1.0 - (minDist / bestSpawn.Radius)
		conc = int(float64(bestSpawn.PeakConc) * factor)
		if conc < 1 {
			conc = 1
		}
	} else {
		conc = 5 // trace reading at edge
	}

	return &SurveyResult{
		ResourceName:  bestSpawn.Name,
		Category:      bestSpawn.Category,
		Concentration: conc,
		Distance:      math.Round(minDist*10) / 10,
		WaypointX:     math.Round((bestSpawn.CenterX+errX)*10) / 10,
		WaypointZ:     math.Round((bestSpawn.CenterZ+errZ)*10) / 10,
	}
}

// SampleSpawn yields 1-3 resource units (GDD 10.5).
func SampleSpawn(spawn *Spawn) SampleResult {
	amount := 1 + rand.Intn(3) // 1 to 3 units
	return SampleResult{
		ResourceID:   spawn.ID,
		ResourceName: spawn.Name,
		Amount:       amount,
		Quality:      int(spawn.QualityMean),
	}
}
