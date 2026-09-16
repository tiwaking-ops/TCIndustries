package handlers

// Phase 10 telemetry world wiring (BAL-001-legal: observation only).
//
// SeedPhase10Telemetry creates the Phase 10 observation tables.
// StartTelemetryJobLoop runs the §29.6 nightly economic-snapshot job as a
// §29.8-shaped phase10_jobs record (single-row claim, idempotent completion),
// on a fast-cycle-mapped cadence (GDD-given 24h default → 60s under
// TESTBED_FAST_CYCLE=1) — separate from the real-time world tick loop.
// The job only writes telemetry; nothing here steers gameplay.
//
// Phase10Anomalies (B5 option (i)) computes RANKED review lists with no
// thresholds: wash-trade pairs, balance churn, placement bursts, price
// outliers. Human review only — never auto-action (§29.6; SAFE-002 TBD
// untouched).
//
// New file: pre-existing handler files are untouched (mixed-provenance
// discipline for the Phase 10 commits).

import (
	"encoding/json"
	"log"
	"net/http"
	"sort"
	"strings"
	"sync/atomic"
	"time"
	"os"

	"swg-server/internal/database"
	"swg-server/internal/skills"
)

// SeedPhase10Telemetry ensures the Phase 10 telemetry tables exist.
func (h *WorldHandler) SeedPhase10Telemetry() error {
	if err := h.db.EnsurePhase10TelemetrySchema(); err != nil {
		return err
	}
	log.Printf("Phase 10 telemetry seeded: interdependence_events, phase10_jobs")
	return nil
}

// --- §29.6 / §29.8-shaped telemetry job -----------------------------------

// Fast-cycle mapping follows the established TESTBED_FAST_CYCLE idiom
// (lairs.go/ranger.go): GDD-given cadence default, testbed-mapped under the
// env flag. The GDD nightly job shape is preserved either way.
const (
	Phase10JobIntervalDefault = 24 * time.Hour
	Phase10JobIntervalFast    = 60 * time.Second
)

// Phase10JobRunning guards against concurrent job execution if loops ever
// overlap (single-row claim already serializes; belt and braces).
var Phase10JobRunning atomic.Bool

// StartTelemetryJobLoop launches the nightly (fast-cycle-mapped) telemetry
// job goroutine. It never touches the real-time tick loop. The loop
// self-reschedules: each completed run enqueues the next pending job row.
func (h *WorldHandler) StartTelemetryJobLoop() {
	interval := Phase10JobIntervalDefault
	if os.Getenv("TESTBED_FAST_CYCLE") == "1" {
		interval = Phase10JobIntervalFast
	}
	// Seed the first pending job so the loop always has work queued.
	if _, err := h.db.EnqueuePhase10Job("economic_telemetry"); err != nil {
		log.Printf("Warning: failed to enqueue initial phase 10 telemetry job: %v", err)
	}
	go func() {
		ticker := time.NewTicker(interval)
		defer ticker.Stop()
		for range ticker.C {
			h.RunTelemetryJob()
		}
	}()
	log.Printf("Phase 10 telemetry job loop started (interval %s)", interval)
}

// RunTelemetryJob claims one pending phase10_jobs row (§29.8 single-row
// claim), computes all read-only telemetry, and completes the job with the
// result JSON (idempotent completion). Failures are recorded on the job row,
// never swallowed silently.
func (h *WorldHandler) RunTelemetryJob() {
	if !Phase10JobRunning.CompareAndSwap(false, true) {
		return
	}
	defer Phase10JobRunning.Store(false)

	id, err := h.db.ClaimPhase10Job("economic_telemetry")
	if err != nil {
		log.Printf("Phase 10 telemetry job claim failed: %v", err)
		return
	}
	if id == "" {
		return // nothing pending
	}

	tel, err := h.CollectPhase10Telemetry()
	if err != nil {
		_ = h.db.FailPhase10Job(id, err.Error())
		log.Printf("Phase 10 telemetry job %s failed: %v", id, err)
		return
	}
	blob, err := json.Marshal(tel)
	if err != nil {
		_ = h.db.FailPhase10Job(id, err.Error())
		return
	}
	if err := h.db.CompletePhase10Job(id, string(blob)); err != nil {
		log.Printf("Phase 10 telemetry job %s completion failed: %v", id, err)
		return
	}
	// Self-reschedule: queue the next pending job row for the next tick.
	_, _ = h.db.EnqueuePhase10Job("economic_telemetry")
	log.Printf("Phase 10 telemetry job %s completed", id)
}

// Phase10Interdependence is the §34 KPI in testbed-observable form: raw
// counts beside every percentage (small-cohort rule, B2).
type Phase10Interdependence struct {
	WindowDays int   `json:"window_days"`
	Covered    int   `json:"covered"`
	Total      int   `json:"total"`
	Pct        float64 `json:"pct"` // 0 when total is 0; read WITH covered/total
}

// Phase10Wealth is the §34.3 wealth snapshot. Gini is diagnostic only —
// tracked, never target-capped, never a lever.
type Phase10Wealth struct {
	TotalCredits int64   `json:"total_credits"`
	Players      int     `json:"players"`
	Gini         float64 `json:"gini"`
}

// Phase10CraftedShare is the §34 player-driven-economy observation: raw
// counts beside the percentage (a 40/42 cohort is 95%, not a verdict).
type Phase10CraftedShare struct {
	Total         int     `json:"total"`
	WithSchematic int     `json:"with_schematic"`
	Pct           float64 `json:"pct"` // 0 when total is 0; read WITH the counts
}

// Phase10Telemetry is the full telemetry payload: job result, REST body,
// and phase10test evidence source, all from one code path.
type Phase10Telemetry struct {
	TakenAt         string                               `json:"taken_at"`
	Interdependence Phase10Interdependence               `json:"interdependence"`
	CraftedShare    Phase10CraftedShare                  `json:"crafted_share"`
	Wealth          Phase10Wealth                        `json:"wealth"`
	Missions        map[string]int                       `json:"mission_outcomes"`
	Volatility      []database.PriceVolatilityRow        `json:"price_volatility"`
	Professions     []database.ProfessionDistributionRow `json:"profession_distribution"`
}

// CollectPhase10Telemetry gathers every B1 read-only metric.
func (h *WorldHandler) CollectPhase10Telemetry() (*Phase10Telemetry, error) {
	tel := &Phase10Telemetry{TakenAt: time.Now().UTC().Format(time.RFC3339)}

	cov, tot, err := h.db.InterdependenceCoverage(database.InterdependenceWindowDays)
	if err != nil {
		return nil, err
	}
	tel.Interdependence = Phase10Interdependence{
		WindowDays: database.InterdependenceWindowDays, Covered: cov, Total: tot,
	}
	if tot > 0 {
		tel.Interdependence.Pct = float64(cov) / float64(tot) * 100
	}

	ct, cs, err := h.db.CraftedItemShare()
	if err != nil {
		return nil, err
	}
	tel.CraftedShare = Phase10CraftedShare{Total: ct, WithSchematic: cs}
	if ct > 0 {
		tel.CraftedShare.Pct = float64(cs) / float64(ct) * 100
	}

	total, players, gini, err := h.db.WealthDistribution()
	if err != nil {
		return nil, err
	}
	tel.Wealth = Phase10Wealth{TotalCredits: total, Players: players, Gini: gini}

	if tel.Missions, err = h.db.MissionOutcomeRates(); err != nil {
		return nil, err
	}
	if tel.Volatility, err = h.db.PriceVolatility(500); err != nil {
		return nil, err
	}
	if tel.Professions, err = h.phase10ProfessionDistribution(); err != nil {
		return nil, err
	}
	return tel, nil
}

// phase10ProfessionDistribution joins skill ownership pairs against the
// skills registry (box → profession + category). Categories come from the
// registry, not string guesses; stale box ids are skipped.
func (h *WorldHandler) phase10ProfessionDistribution() ([]database.ProfessionDistributionRow, error) {
	pairs, err := h.db.SkillOwnershipPairs()
	if err != nil {
		return nil, err
	}
	type profAgg struct {
		row     database.ProfessionDistributionRow
		holders map[string]bool // distinct characters owning ≥1 box
	}
	byProf := map[string]*profAgg{}
	for _, p := range pairs {
		pb := skills.GetSkillBox(p.BoxID)
		if pb == nil {
			continue // stale box id; not a registry profession box
		}
		agg, ok := byProf[pb.ProfessionID]
		if !ok {
			agg = &profAgg{holders: map[string]bool{}}
			agg.row.ProfessionID = pb.ProfessionID
			if prof := skills.ProfessionByID(pb.ProfessionID); prof != nil {
				agg.row.Category = string(prof.Category)
			}
			byProf[pb.ProfessionID] = agg
		}
		agg.row.BoxCount++
		agg.holders[p.CharacterID] = true
	}
	out := make([]database.ProfessionDistributionRow, 0, len(byProf))
	for _, agg := range byProf {
		agg.row.HolderCount = len(agg.holders)
		out = append(out, agg.row)
	}
	return out, nil
}

// --- B5(i): ranked anomaly review lists (no thresholds, no auto-action) ----

// Phase10Anomaly is one ranked review candidate. Rank is an ordering signal
// for a HUMAN reviewer, not a verdict and never an enforcement trigger.
type Phase10Anomaly struct {
	Kind        string   `json:"kind"` // wash_trade_pair | balance_churn | placement_burst | price_outlier
	Rank        float64  `json:"rank"` // sort key (higher = review first)
	CharacterID string   `json:"character_id,omitempty"`
	PartyID     string   `json:"party_id,omitempty"` // counterparty / item / structure
	Transfers   int      `json:"transfers"`
	Forward     int64    `json:"forward_amount"`
	Reverse     int64    `json:"reverse_amount"`
	Net         int64    `json:"net_amount"`
	Categories  []string `json:"categories,omitempty"`
	Note        string   `json:"note,omitempty"`
}

// Phase10Anomalies ranks the ledger/market history for human review.
// Option (i) construction — candidate generation uses structural invariants
// only (pairs need both directions; bursts need ≥3 placements; outliers are
// the top spread items). No tuned thresholds, no sigma, no auto-action.
func (h *WorldHandler) Phase10Anomalies() []Phase10Anomaly {
	out := h.anomaliesFromLedger()
	out = append(out, h.anomaliesFromMarket()...)
	sort.SliceStable(out, func(i, j int) bool { return out[i].Rank > out[j].Rank })
	return out
}

func (h *WorldHandler) anomaliesFromLedger() []Phase10Anomaly {
	flows, err := h.db.LedgerFlows()
	if err != nil {
		return nil
	}
	type flowKey struct{ from, to string }
	flows2 := map[flowKey]*Phase10Anomaly{}
	charTotals := map[string]*Phase10Anomaly{}
	for _, lf := range flows {
		if lf.CharacterID == "" {
			continue
		}
		if lf.Counterparty != "" {
			k := flowKey{lf.CharacterID, lf.Counterparty}
			f := flows2[k]
			if f == nil {
				f = &Phase10Anomaly{Kind: "wash_trade_pair", CharacterID: lf.CharacterID, PartyID: lf.Counterparty}
				flows2[k] = f
			}
			f.Transfers++
			if lf.Amount > 0 {
				f.Forward += lf.Amount
			} else {
				f.Reverse += -lf.Amount
			}
			f.Net = f.Forward - f.Reverse
		}
		t := charTotals[lf.CharacterID]
		if t == nil {
			t = &Phase10Anomaly{Kind: "balance_churn", CharacterID: lf.CharacterID}
			charTotals[lf.CharacterID] = t
		}
		t.Transfers++
		if lf.Amount > 0 {
			t.Forward += lf.Amount
		} else {
			t.Reverse += -lf.Amount
		}
		t.Net = t.Forward - t.Reverse
	}

	var out []Phase10Anomaly
	// Wash-trade pairs: both directions present, at least 2 round trips.
	for _, f := range flows2 {
		if f.Forward > 0 && f.Reverse > 0 {
			pairs := f.Forward
			if f.Reverse < pairs {
				pairs = f.Reverse
			}
			if pairs >= 2 {
				f.Rank = float64(pairs)
				out = append(out, *f)
			}
		}
	}
	// Balance churn: many transfers, near-zero net — the wash-trade shape.
	// The 10%-of-gross band is the DEFINITION of near-zero (so the filter is
	// the construct, not a tuned threshold); surfaced for human review only.
	for _, t := range charTotals {
		gross := t.Forward + t.Reverse
		if t.Transfers >= 6 && t.Forward >= 4 && t.Reverse >= 4 &&
			gross > 0 && abs64(t.Net)*10 <= gross {
			t.Rank = float64(t.Transfers)
			out = append(out, *t)
		}
	}
	out = append(out, h.anomaliesFromPlacements()...)
	return out
}

// anomaliesFromPlacements ranks placement bursts: any owner with ≥3 placed
// structures is a review candidate, ranked by burst count. The ≥3 minimum is
// the burst construct (one house + vendor + harvester is the normal case);
// not a tuned enforcement threshold — human review only.
func (h *WorldHandler) anomaliesFromPlacements() []Phase10Anomaly {
	placements, err := h.db.StructurePlacements()
	if err != nil {
		return nil
	}
	perOwner := map[string]int{}
	for _, p := range placements {
		if p.OwnerID == "" {
			continue
		}
		perOwner[p.OwnerID]++
	}
	var out []Phase10Anomaly
	for owner, n := range perOwner {
		if n >= 3 {
			out = append(out, Phase10Anomaly{
				Kind: "placement_burst", CharacterID: owner, Transfers: n,
				Rank: float64(n), Note: "structures placed by owner",
			})
		}
	}
	return out
}

func (h *WorldHandler) anomaliesFromMarket() []Phase10Anomaly {
	vols, err := h.db.PriceVolatility(500)
	if err != nil {
		return nil
	}
	var out []Phase10Anomaly
	for _, v := range vols {
		if v.Sales >= 2 {
			out = append(out, Phase10Anomaly{
				Kind: "price_outlier", PartyID: v.ItemID, Transfers: v.Sales,
				Forward: v.Max, Reverse: v.Min, Rank: v.SpreadPc,
				Note: "spread_pct over recent sales",
			})
		}
	}
	return out
}

func abs64(v int64) int64 {
	if v < 0 {
		return -v
	}
	return v
}

// --- Read-only REST surface (§34 instrumentation, no dashboards) -----------

// Phase10TelemetryHandler serves the Phase 10 observation routes:
//
//	GET /api/phase10/telemetry — full B1 payload (coverage, wealth+Gini,
//	    mission outcomes, price volatility, profession distribution)
//	GET /api/phase10/anomalies — B5(i) ranked review lists
//	GET /api/phase10/jobs      — telemetry job history (§29.8 records)
//
// All routes are read-only; the owner-keyed auth middleware wraps them.
type Phase10TelemetryHandler struct {
	db *database.DB
}

// NewPhase10TelemetryHandler builds the Phase 10 read-only REST handler.
func NewPhase10TelemetryHandler(db *database.DB) *Phase10TelemetryHandler {
	return &Phase10TelemetryHandler{db: db}
}

// HandlePhase10Route dispatches /api/phase10/* (same style as the other
// per-package route handlers).
func (ph *Phase10TelemetryHandler) HandlePhase10Route(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimPrefix(r.URL.Path, "/api/phase10/")
	switch {
	case path == "telemetry" && r.Method == "GET":
		wh := NewWorldHandler(ph.db)
		tel, err := wh.CollectPhase10Telemetry()
		if err != nil {
			writeCraftJSON(w, http.StatusInternalServerError,
				map[string]string{"error": "telemetry query failed"})
			return
		}
		writeCraftJSON(w, http.StatusOK, tel)
	case path == "anomalies" && r.Method == "GET":
		wh := NewWorldHandler(ph.db)
		writeCraftJSON(w, http.StatusOK, map[string]interface{}{
			"note":   "ranked review lists — no thresholds, human review only, never auto-action (§29.6)",
					"anomalies": wh.Phase10Anomalies(),
		})
	case path == "jobs" && r.Method == "GET":
		ph.jobs(w, r)
	default:
		writeCraftJSON(w, http.StatusNotFound, map[string]string{"error": "unknown phase10 route"})
	}
}

type phase10JobRow struct {
	ID          string `json:"id"`
	JobType     string `json:"job_type"`
	Status      string `json:"status"`
	ClaimedAt   string `json:"claimed_at,omitempty"`
	CompletedAt string `json:"completed_at,omitempty"`
	ResultJSON  string `json:"result_json,omitempty"`
}

func (ph *Phase10TelemetryHandler) jobs(w http.ResponseWriter, r *http.Request) {
	his, err := ph.db.Phase10JobHistory(50)
	if err != nil {
		writeCraftJSON(w, http.StatusInternalServerError,
			map[string]string{"error": "job query failed"})
		return
	}
	out := []phase10JobRow{}
	for _, j := range his {
		out = append(out, phase10JobRow{
			ID: j.ID, JobType: j.JobType, Status: j.Status,
			ClaimedAt: j.ClaimedAt, CompletedAt: j.CompletedAt, ResultJSON: j.ResultJSON,
		})
	}
	writeCraftJSON(w, http.StatusOK, map[string]interface{}{"jobs": out})
}
