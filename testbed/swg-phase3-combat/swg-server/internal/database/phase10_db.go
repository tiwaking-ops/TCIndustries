package database

// Phase 10 telemetry layer (BAL-001-legal: observation only).
//
// Design notes, per proposal §3/§2.4:
//   - interdependence_events is APPEND-ONLY; writes occur at existing hook
//     sites (service heal/buff receipt, crafted-item purchase). No retro
//     backfill: pre-log history is genuinely unobservable.
//   - market_records (economy_db.go) already IS the append-only sale-price
//     history; volatility reads it directly. No new history table.
//   - All queries are read-only SELECTs; nothing here mutates gameplay state.
//   - The telemetry job follows the §29.8 WorldTickJob shape (single-row
//     claim + idempotent completion) but lives in the `phase10_jobs` table
//     and only ever writes telemetry — it is separate from the real-time
//     world loop.
//   - Row IDs use NewRowID (collision-proof; Windows clock granularity).
//   - No numeric tuning: no gameplay value is read, compared, or changed.

import (
	"database/sql"
	"fmt"
	"math"
)

// EnsurePhase10TelemetrySchema creates the Phase 10 observation tables.
// Called from SeedPhase10Telemetry during server bootstrap.
func (db *DB) EnsurePhase10TelemetrySchema() error {
	stmts := []string{
		// B1: append-only interdependence log (§34 economic-interdependence
		// instrumentation). kind ∈ crafted_purchase | heal | buff.
		// counterparty = vendor owner (purchases) or service provider (heals).
		`CREATE TABLE IF NOT EXISTS interdependence_events (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			kind TEXT NOT NULL,
			character_id TEXT NOT NULL,
			counterparty_id TEXT NOT NULL DEFAULT '',
			zone TEXT NOT NULL DEFAULT '',
			created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
		)`,
		`CREATE INDEX IF NOT EXISTS idx_interdependence_char_time
			ON interdependence_events(character_id, created_at)`,

		// §29.8 WorldTickJob-shaped telemetry job record (single-row claim,
		// idempotent completion). Cadence is fast-cycle-mapped in the handler
		// (GDD-given 24h default maps to 60s under TESTBED_FAST_CYCLE=1).
		`CREATE TABLE IF NOT EXISTS phase10_jobs (
			id TEXT PRIMARY KEY,
			job_type TEXT NOT NULL,
			status TEXT NOT NULL DEFAULT 'pending',
			claimed_at DATETIME,
			completed_at DATETIME,
			result_json TEXT NOT NULL DEFAULT ''
		)`,

		// B5(i) ranked anomaly review lists are computed from ledger_entries +
		// market_records + structures at query time — no threshold state and
		// no extra tables needed (option (i): ranked lists, no thresholds).
	}
	for i, s := range stmts {
		if _, err := db.conn.Exec(s); err != nil {
			return fmt.Errorf("phase10 telemetry schema %d failed: %w", i+1, err)
		}
	}
	return nil
}

// --- B1: interdependence event writers (hook-callable) -------------------

// RecordInterdependenceEvent appends one interdependence observation.
// Best-effort by contract: hook sites call it with `_ =` exactly like the
// existing ledger/XP hook calls, so a telemetry failure can never fail a
// gameplay action. Only ever INSERTS.
func (db *DB) RecordInterdependenceEvent(kind, characterID, counterpartyID, zone string) error {
	_, err := db.conn.Exec(
		`INSERT INTO interdependence_events (kind, character_id, counterparty_id, zone)
		 VALUES (?, ?, ?, ?)`,
		kind, characterID, counterpartyID, zone,
	)
	return err
}

// InterdependenceWindow is the §34 KPI lookback ("last 7 days"). Fast-cycle
// test runs substitute a shorter window at the query layer via days.
const InterdependenceWindowDays = 7

// InterdependenceCoverage returns, per §34 KPI form, the share of characters
// who received a crafted item / entertainer-music / medic heal in the window.
// Small cohorts are expected (testbed): the percentage is meaningless without
// the raw counts, so both are returned and reported side by side (B2 rule).
func (db *DB) InterdependenceCoverage(days int) (covered, total int, err error) {
	var nullTotal sql.NullInt64
	err = db.conn.QueryRow(
		`SELECT COUNT(*) FROM characters`).Scan(&nullTotal)
	if err != nil {
		return 0, 0, err
	}
	total = int(nullTotal.Int64)
	err = db.conn.QueryRow(
		`SELECT COUNT(DISTINCT character_id) FROM interdependence_events
		 WHERE created_at >= datetime('now', ?)`,
		fmt.Sprintf("-%d days", days),
	).Scan(&covered)
	return covered, total, err
}

// --- B1: profession distribution -----------------------------------------
//
// The box→profession/category mapping lives in the skills registry
// (internal/skills), which this package cannot import (layering). The DB
// layer exposes raw ownership pairs; the handlers layer joins them against
// the registry (Phase10ProfessionDistribution in phase10_handlers.go).

// ProfessionDistributionRow is one profession's box-ownership footprint.
// Populated by the handlers layer's registry join (the box→profession
// mapping lives in internal/skills, which this package cannot import).
type ProfessionDistributionRow struct {
	ProfessionID string `json:"profession_id"`
	Category     string `json:"category"` // basic | elite | hybrid
	BoxCount     int    `json:"box_count"`
	HolderCount  int    `json:"holder_count"` // characters owning ≥1 box
}

// SkillOwnershipPair is one (character, box) ownership row.
type SkillOwnershipPair struct {
	CharacterID string
	BoxID       string
}

// SkillOwnershipPairs returns all skill-box ownership rows.
func (db *DB) SkillOwnershipPairs() ([]SkillOwnershipPair, error) {
	rows, err := db.conn.Query(
		`SELECT character_id, skill_box_id FROM character_skills`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var out []SkillOwnershipPair
	for rows.Next() {
		var p SkillOwnershipPair
		if err := rows.Scan(&p.CharacterID, &p.BoxID); err != nil {
			return nil, err
		}
		out = append(out, p)
	}
	return out, rows.Err()
}

// CraftedItemShare reports the §34 player-driven-economy observation: how
// many items exist and how many carry a schematic provenance. The §1.3 KPI
// reads "crafted_from_schematic_id != null ≥95%" — reported with raw counts
// beside the percentage (small-cohort rule).
func (db *DB) CraftedItemShare() (total, withSchematic int, err error) {
	err = db.conn.QueryRow(
		`SELECT COUNT(*), COALESCE(SUM(CASE WHEN schematic_id != '' THEN 1 ELSE 0 END),0)
		 FROM crafted_items`).Scan(&total, &withSchematic)
	return total, withSchematic, err
}

// --- B1: mission outcome rates --------------------------------------------

// MissionOutcomeRates reports the lifecycle mix of posted missions. Statuses
// are exactly those the missions table uses (available/accepted/completed/
// expired — set by phase9_db.go SetMissionStatus/ExpireMissions).
func (db *DB) MissionOutcomeRates() (map[string]int, error) {
	rows, err := db.conn.Query(
		`SELECT status, COUNT(*) FROM missions GROUP BY status`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	out := map[string]int{}
	for rows.Next() {
		var status string
		var n int
		if err := rows.Scan(&status, &n); err != nil {
			return nil, err
		}
		out[status] = n
	}
	return out, rows.Err()
}

// --- B1: wealth distribution (diagnostic only, §34.3) ----------------------

// WealthDistribution reports raw credits statistics plus the Gini coefficient.
// Gini is DIAGNOSTIC ONLY per §34.3 — tracked, never target-capped, never a
// balance lever (BAL-001: observation, not tuning).
func (db *DB) WealthDistribution() (total int64, players int, gini float64, err error) {
	rows, err := db.conn.Query(
		`SELECT credits FROM characters WHERE credits >= 0 ORDER BY credits ASC`)
	if err != nil {
		return 0, 0, 0, err
	}
	defer rows.Close()
	var vals []int64
	for rows.Next() {
		var v int64
		if err := rows.Scan(&v); err != nil {
			return 0, 0, 0, err
		}
		vals = append(vals, v)
	}
	if err := rows.Err(); err != nil {
		return 0, 0, 0, err
	}
	n := len(vals)
	if n == 0 {
		return 0, 0, 0, nil
	}
	var sum int64
	for _, v := range vals {
		sum += v
	}
	// Gini for a sorted sample: G = (2*Σ i*x_i)/(n*Σ x_i) − (n+1)/n, i = 1..n
	var weighted int64
	for i, v := range vals {
		weighted += int64(i+1) * v
	}
	if sum == 0 {
		return 0, n, 0, nil // perfect equality when nobody holds credits
	}
	g := (2*float64(weighted))/(float64(n)*float64(sum)) - (float64(n)+1)/float64(n)
	return sum, n, math.Max(0, math.Min(1, g)), nil
}

// --- B1: resource price volatility (per rotation cycle) -------------------

// PriceVolatilityRow summarizes one item_id's sale-price spread from the
// append-only market_records history (GDD 22.3). Spread (max−min over mean)
// is a robust small-sample volatility measure; stddev-based sigma is
// meaningless at testbed cohort sizes.
type PriceVolatilityRow struct {
	ItemID   string  `json:"item_id"`
	Sales    int     `json:"sales"`
	Min      int64   `json:"min_price"`
	Max      int64   `json:"max_price"`
	Mean     float64 `json:"mean_price"`
	SpreadPc float64 `json:"spread_pct"` // (max-min)/mean*100, 0 when mean 0
}

// PriceVolatility reads the last `limitCycles` sale observations per item.
// "Rotation cycle" in the testbed maps to the recorded sale stream — the
// per-item series IS the per-cycle history (no scheduler exists to sample).
func (db *DB) PriceVolatility(limitCycles int) ([]PriceVolatilityRow, error) {
	rows, err := db.conn.Query(
		`SELECT item_id, COUNT(*), MIN(price), MAX(price), AVG(price)
		 FROM (
		   SELECT item_id, price FROM market_records
		   ORDER BY id DESC LIMIT ?
		 )
		 GROUP BY item_id`, limitCycles)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var out []PriceVolatilityRow
	for rows.Next() {
		var r PriceVolatilityRow
		if err := rows.Scan(&r.ItemID, &r.Sales, &r.Min, &r.Max, &r.Mean); err != nil {
			return nil, err
		}
		if r.Mean > 0 {
			r.SpreadPc = float64(r.Max-r.Min) / r.Mean * 100
		}
		out = append(out, r)
	}
	return out, rows.Err()
}

// --- B5(i) anomaly source queries (read-only; ranking lives in handlers) ---

// LedgerFlowRow is one ledger_entries observation for anomaly analysis.
type LedgerFlowRow struct {
	CharacterID   string
	Counterparty  string
	Category      string
	Amount        int64
}

// LedgerFlows returns the full ledger stream in insertion order (append-only
// source: ledger_entries is written by RecordLedger/RecordLedgerTx hooks).
func (db *DB) LedgerFlows() ([]LedgerFlowRow, error) {
	rows, err := db.conn.Query(
		`SELECT character_id, counterparty_id, category, amount
		 FROM ledger_entries ORDER BY id`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var out []LedgerFlowRow
	for rows.Next() {
		var r LedgerFlowRow
		if err := rows.Scan(&r.CharacterID, &r.Counterparty, &r.Category, &r.Amount); err != nil {
			return nil, err
		}
		out = append(out, r)
	}
	return out, rows.Err()
}

// StructurePlacementRow is one structure placement observation.
type StructurePlacementRow struct {
	ID        string
	OwnerID   string
	Zone      string
	CreatedAt string
}

// StructurePlacements returns structure rows for placement-burst analysis.
func (db *DB) StructurePlacements() ([]StructurePlacementRow, error) {
	rows, err := db.conn.Query(
		`SELECT id, owner_character_id, zone, created_at
		 FROM structures ORDER BY created_at`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var out []StructurePlacementRow
	for rows.Next() {
		var r StructurePlacementRow
		if err := rows.Scan(&r.ID, &r.OwnerID, &r.Zone, &r.CreatedAt); err != nil {
			return nil, err
		}
		out = append(out, r)
	}
	return out, rows.Err()
}

// --- §29.8-shaped telemetry job -------------------------------------------

// EnqueuePhase10Job creates one pending telemetry job row (§29.8 pending
// state). The telemetry loop self-reschedules: each completed run enqueues
// the next pending row.
func (db *DB) EnqueuePhase10Job(jobType string) (string, error) {
	id := newRowID("p10job")
	_, err := db.conn.Exec(
		`INSERT INTO phase10_jobs (id, job_type, status) VALUES (?, ?, 'pending')`,
		id, jobType)
	return id, err
}

// ClaimPhase10Job atomically claims one pending job of jobType (single-row
// claim rule, §29.8). Returns ("", nil) when nothing is pending.
func (db *DB) ClaimPhase10Job(jobType string) (string, error) {
	res, err := db.conn.Exec(
		`UPDATE phase10_jobs SET status='running', claimed_at=CURRENT_TIMESTAMP
		 WHERE id = (
		   SELECT id FROM phase10_jobs
		   WHERE job_type = ? AND status = 'pending'
		   ORDER BY rowid LIMIT 1
		 ) AND status = 'pending'`, jobType)
	if err != nil {
		return "", err
	}
	n, err := res.RowsAffected()
	if err != nil || n == 0 {
		return "", err
	}
	var id string
	err = db.conn.QueryRow(
		`SELECT id FROM phase10_jobs WHERE job_type = ? AND status = 'running'
		 ORDER BY claimed_at DESC, rowid DESC LIMIT 1`, jobType).Scan(&id)
	return id, err
}

// CompletePhase10Job marks a claimed job completed, storing its result JSON
// (idempotent completion, §29.8: only transitions running → completed).
func (db *DB) CompletePhase10Job(id, resultJSON string) error {
	_, err := db.conn.Exec(
		`UPDATE phase10_jobs
		 SET status='completed', completed_at=CURRENT_TIMESTAMP, result_json=?
		 WHERE id = ? AND status = 'running'`, resultJSON, id)
	return err
}

// FailPhase10Job marks a claimed job failed (§29.8 crash-mid-tick rule — the
// job row records the failure for resumption/inspection instead of vanishing).
func (db *DB) FailPhase10Job(id, reason string) error {
	_, err := db.conn.Exec(
		`UPDATE phase10_jobs
		 SET status='failed', completed_at=CURRENT_TIMESTAMP, result_json=?
		 WHERE id = ? AND status = 'running'`, reason, id)
	return err
}

// Phase10JobHistoryRow is one job-history observation for the REST surface.
type Phase10JobHistoryRow struct {
	ID          string
	JobType     string
	Status      string
	ClaimedAt   string
	CompletedAt string
	ResultJSON  string
}

// Phase10JobHistory returns the most recent job records (newest first).
func (db *DB) Phase10JobHistory(limit int) ([]Phase10JobHistoryRow, error) {
	rows, err := db.conn.Query(
		`SELECT id, job_type, status, COALESCE(claimed_at,''),
		        COALESCE(completed_at,''), result_json
		 FROM phase10_jobs ORDER BY rowid DESC LIMIT ?`, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var out []Phase10JobHistoryRow
	for rows.Next() {
		var j Phase10JobHistoryRow
		if err := rows.Scan(&j.ID, &j.JobType, &j.Status, &j.ClaimedAt,
			&j.CompletedAt, &j.ResultJSON); err != nil {
			return nil, err
		}
		out = append(out, j)
	}
	return out, rows.Err()
}
