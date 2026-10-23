package economy

import (
	"database/sql"
	"math/rand"
)

// CalculateCreatureDropCredits computes credit drop reward scaled by creature level (GDD 12.2).
// Band: 10 to 200 credits.
func CalculateCreatureDropCredits(creatureLevel int) int {
	if creatureLevel < 1 {
		creatureLevel = 1
	}
	base := 10 + (creatureLevel * 15)
	if base > 200 {
		base = 200
	}
	// Add minor variance (±20%)
	variance := float64(base) * 0.2 * (rand.Float64()*2 - 1)
	result := int(float64(base) + variance)
	if result < 10 {
		result = 10
	}
	return result
}

// AwardCreatureDrop credits a character for defeating a creature and logs a faucet ledger entry (GDD 12.2, 12.6).
func AwardCreatureDrop(db *sql.DB, characterID string, creatureLevel int, planet string, posX, posZ float64) (int, error) {
	credits := CalculateCreatureDropCredits(creatureLevel)

	tx, err := db.Begin()
	if err != nil {
		return 0, err
	}
	defer tx.Rollback()

	_, err = tx.Exec("UPDATE characters SET credits = credits + ? WHERE id = ?", credits, characterID)
	if err != nil {
		return 0, err
	}

	err = RecordLedgerEntry(db, characterID, "faucet", "creature_drop", credits, "environment", planet, posX, posZ)
	if err != nil {
		return 0, err
	}

	return credits, tx.Commit()
}
