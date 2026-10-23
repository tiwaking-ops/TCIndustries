package database

import (
	"fmt"
	"time"
)

// AddCredits adds credits to a character's wallet (faucet-side primitive).
func (db *DB) AddCredits(characterID string, amount int) error {
	if amount <= 0 {
		return fmt.Errorf("amount must be positive")
	}
	_, err := db.conn.Exec(
		`UPDATE characters SET credits = credits + ? WHERE id = ?`,
		amount, characterID,
	)
	return err
}

// EnsureServiceSchema creates Phase 6 tables if absent (buffs).
func (db *DB) EnsureServiceSchema() error {
	_, err := db.conn.Exec(
		`CREATE TABLE IF NOT EXISTS buffs (
			id TEXT PRIMARY KEY,
			target_character_id TEXT NOT NULL,
			source_character_id TEXT NOT NULL,
			buff_type TEXT NOT NULL,
			amount INTEGER NOT NULL,
			applied_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
			expires_at DATETIME NOT NULL
		)`)
	if err != nil {
		return fmt.Errorf("service schema failed: %w", err)
	}
	_, err = db.conn.Exec(
		`CREATE INDEX IF NOT EXISTS idx_buffs_target ON buffs(target_character_id)`)
	return err
}

// BuffRow maps one buffs row.
type BuffRow struct {
	ID         string
	TargetID   string
	SourceID   string
	BuffType   string
	Amount     int
	AppliedAt  time.Time
	ExpiresAt  time.Time
}

// AddBuff records a buff, rejecting a second active buff of the same pool type.
func (db *DB) AddBuff(target, source, buffType string, amount int, expires time.Time) error {
	var n int
	if err := db.conn.QueryRow(
		`SELECT COUNT(*) FROM buffs WHERE target_character_id = ? AND buff_type = ?
		AND expires_at > CURRENT_TIMESTAMP`,
		target, buffType).Scan(&n); err != nil {
		return err
	}
	if n > 0 {
		return fmt.Errorf("buff already active")
	}
	id := fmt.Sprintf("buff_%d", time.Now().UnixNano())
	_, err := db.conn.Exec(
		`INSERT INTO buffs (id, target_character_id, source_character_id,
			buff_type, amount, expires_at)
		VALUES (?, ?, ?, ?, ?, ?)`,
		id, target, source, buffType, amount,
		expires.UTC().Format("2006-01-02 15:04:05"))
	return err
}

// BuffBonus sums active buffs for one pool ("health" | "action" | "mind").
func (db *DB) BuffBonus(target, pool string) int {
	var total int
	_ = db.conn.QueryRow(
		`SELECT COALESCE(SUM(amount),0) FROM buffs
		WHERE target_character_id = ? AND buff_type = ?
		AND expires_at > CURRENT_TIMESTAMP`,
		target, pool).Scan(&total)
	return total
}

// ExpireBuffs deletes expired rows; returns the count removed.
func (db *DB) ExpireBuffs() int {
	res, err := db.conn.Exec(`DELETE FROM buffs WHERE expires_at <= CURRENT_TIMESTAMP`)
	if err != nil {
		return 0
	}
	n, _ := res.RowsAffected()
	return int(n)
}

// UpdateItemStats rewrites an item's stats JSON (e.g. stim charge tracking).
// Ported from the upstream tree; required by resource_db.go's ammo decrement.
func (db *DB) UpdateItemStats(owner, itemID, statsJSON string) error {
	_, err := db.conn.Exec(
		`UPDATE crafted_items SET stats_json = ? WHERE id = ? AND owner_character_id = ?`,
		statsJSON, itemID, owner)
	return err
}

// BumpCraftCount increments a character's per-schematic craft count and
// returns the new total (this craft included).
// Drives the repeat-craft XP diminishing in the crafting handler (GDD 7.2.2:
// 50% XP after the 10th craft of one schematic).
// Ported from the upstream tree; required by handlers/crafting.go.
func (db *DB) BumpCraftCount(charID, schematic string) (int, error) {
	if _, err := db.conn.Exec(
		`CREATE TABLE IF NOT EXISTS craft_counts (
			character_id TEXT NOT NULL,
			schematic TEXT NOT NULL,
			n INTEGER NOT NULL DEFAULT 0,
			PRIMARY KEY (character_id, schematic)
		)`); err != nil {
		return 0, err
	}
	if _, err := db.conn.Exec(
		`INSERT INTO craft_counts (character_id, schematic, n) VALUES (?, ?, 1)
		ON CONFLICT(character_id, schematic) DO UPDATE SET n = n + 1`,
		charID, schematic); err != nil {
		return 0, err
	}
	var n int
	if err := db.conn.QueryRow(
		`SELECT n FROM craft_counts WHERE character_id = ? AND schematic = ?`,
		charID, schematic).Scan(&n); err != nil {
		return 0, err
	}
	return n, nil
}
