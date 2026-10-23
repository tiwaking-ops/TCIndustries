package economy

import (
	"database/sql"
	"time"
)

// EconomicSnapshot represents macro-economic circulation metrics (GDD 12.6).
type EconomicSnapshot struct {
	ID             int64     `json:"id"`
	TotalCirculation int     `json:"total_circulation"`
	TotalFaucets   int       `json:"total_faucets"`
	TotalSinks     int       `json:"total_sinks"`
	InflationPct   float64   `json:"inflation_pct"`
	Timestamp      time.Time `json:"timestamp"`
}

// GenerateEconomicSnapshot computes current macro metrics from the credit ledger and characters table.
func GenerateEconomicSnapshot(db *sql.DB) (EconomicSnapshot, error) {
	var totalCirculation int
	err := db.QueryRow("SELECT COALESCE(SUM(credits), 0) FROM characters").Scan(&totalCirculation)
	if err != nil {
		return EconomicSnapshot{}, err
	}

	var totalFaucets, totalSinks int
	db.QueryRow("SELECT COALESCE(SUM(amount), 0) FROM credit_ledger WHERE category = 'faucet'").Scan(&totalFaucets)
	db.QueryRow("SELECT COALESCE(SUM(amount), 0) FROM credit_ledger WHERE category = 'sink'").Scan(&totalSinks)

	// Simple rolling net inflation calculation
	inflationPct := 0.0
	if totalCirculation > 0 {
		net := float64(totalFaucets - totalSinks)
		inflationPct = (net / float64(totalCirculation)) * 100.0
	}

	snapshot := EconomicSnapshot{
		TotalCirculation: totalCirculation,
		TotalFaucets:     totalFaucets,
		TotalSinks:       totalSinks,
		InflationPct:     inflationPct,
		Timestamp:        time.Now(),
	}

	// Persist snapshot
	query := `INSERT INTO economic_snapshots (total_circulation, total_faucets, total_sinks, inflation_pct, timestamp)
	          VALUES (?, ?, ?, ?, ?)`
	_, err = db.Exec(query, snapshot.TotalCirculation, snapshot.TotalFaucets, snapshot.TotalSinks, snapshot.InflationPct, snapshot.Timestamp)
	if err != nil {
		return snapshot, err
	}

	return snapshot, nil
}
