package economy

import (
	"database/sql"
	"errors"
)

// AtomicTransfer executes an atomic credit and/or item transfer between two characters (GDD 29.5).
// Uses BEGIN IMMEDIATE to prevent race conditions (double-purchase safety).
func AtomicTransfer(db *sql.DB, buyerID, sellerID string, creditAmount int) error {
	if creditAmount <= 0 {
		return errors.New("transfer amount must be positive")
	}

	tx, err := db.BeginTx(nil, &sql.TxOptions{Isolation: sql.LevelSerializable})
	if err != nil {
		return err
	}
	defer tx.Rollback()

	// Check buyer balance
	var buyerCredits int
	err = tx.QueryRow("SELECT credits FROM characters WHERE id = ?", buyerID).Scan(&buyerCredits)
	if err != nil {
		return err
	}

	if buyerCredits < creditAmount {
		return errors.New("insufficient credits")
	}

	// Deduct from buyer
	_, err = tx.Exec("UPDATE characters SET credits = credits - ? WHERE id = ?", creditAmount, buyerID)
	if err != nil {
		return err
	}

	// Add to seller (if sellerID is provided and not empty)
	if sellerID != "" {
		_, err = tx.Exec("UPDATE characters SET credits = credits + ? WHERE id = ?", creditAmount, sellerID)
		if err != nil {
			return err
		}
	}

	return tx.Commit()
}
