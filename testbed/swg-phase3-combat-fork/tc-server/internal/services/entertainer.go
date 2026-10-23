package services

import (
	"errors"
	"time"
)

// PerformanceSession represents an active entertainer performance (GDD 23.2.1).
type PerformanceSession struct {
	PerformerID string    `json:"performer_id"`
	PerformanceKind string `json:"performance_kind"` // "music" or "dance"
	Tier        int       `json:"tier"`
	StartedAt   time.Time `json:"started_at"`
	Watchers    map[string]bool // watcherID -> true
}

// EntertainerService manages performances, watch sessions, and battle fatigue healing (GDD 23.2).
type EntertainerService struct {
	performances map[string]*PerformanceSession // performerID -> session
	watcherToPerformer map[string]string         // watcherID -> performerID
}

func NewEntertainerService() *EntertainerService {
	return &EntertainerService{
		performances:       make(map[string]*PerformanceSession),
		watcherToPerformer: make(map[string]string),
	}
}

// StartPerformance starts music or dance performance.
func (es *EntertainerService) StartPerformance(performerID, kind string, tier int) (*PerformanceSession, error) {
	if _, ok := es.performances[performerID]; ok {
		return nil, errors.New("already performing")
	}

	session := &PerformanceSession{
		PerformerID:     performerID,
		PerformanceKind: kind,
		Tier:            tier,
		StartedAt:       time.Now(),
		Watchers:        make(map[string]bool),
	}

	es.performances[performerID] = session
	return session, nil
}

// StopPerformance ends performance.
func (es *EntertainerService) StopPerformance(performerID string) {
	session, ok := es.performances[performerID]
	if !ok {
		return
	}
	for watcherID := range session.Watchers {
		delete(es.watcherToPerformer, watcherID)
	}
	delete(es.performances, performerID)
}

// WatchPerformance registers a player watching a performer (GDD 23.2.1).
func (es *EntertainerService) WatchPerformance(watcherID, performerID string) error {
	session, ok := es.performances[performerID]
	if !ok {
		return errors.New("performer is not active")
	}

	// Enforce one-watch-target rule (GDD 23.5)
	if oldPerformer, ok := es.watcherToPerformer[watcherID]; ok {
		if oldPerformer != performerID {
			if oldSession, ok := es.performances[oldPerformer]; ok {
				delete(oldSession.Watchers, watcherID)
			}
		}
	}

	session.Watchers[watcherID] = true
	es.watcherToPerformer[watcherID] = performerID
	return nil
}

// CalculateBFHeal computes battle fatigue healing per tick for watchers (GDD 23.2.2).
func CalculateBFHeal(entertainerTier int) float64 {
	// 1-2% BF reduction per tick per tier
	return 1.0 + float64(entertainerTier)*0.5
}
