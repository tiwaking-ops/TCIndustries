package economy

import (
	"database/sql"
	"time"
)

// LedgerEntry represents an append-only credit ledger record (GDD 12.6).
type LedgerEntry struct {
	ID            int64     `json:"id"`
	CharacterID   string    `json:"character_id"`
	Category      string    `json:"category"` // "faucet", "sink", "transfer"
	Subtype       string    `json:"subtype"`  // e.g. "creature_drop", "training_cost", "vendor_purchase"
	Amount        int       `json:"amount"`
	Counterparty  string    `json:"counterparty"`
	Planet        string    `json:"planet"`
	PosX          float64   `json:"pos_x"`
	PosZ          float64   `json:"pos_z"`
	Timestamp     time.Time `json:"timestamp"`
}

// RecordLedgerEntry writes an append-only entry to the database.
func RecordLedgerEntry(db *sql.DB, characterID, category, subtype string, amount int, counterparty, planet string, posX, posZ float64) error {
	query := `INSERT INTO credit_ledger (character_id, category, subtype, amount, counterparty, planet, pos_x, pos_z, timestamp)
	          VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)`
	_, err := db.Exec(query, characterID, category, subtype, amount, counterparty, planet, posX, posZ, time.Now())
	return err
}
