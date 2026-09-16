// ServiceDB extends the DB with Phase 6 persistence: buffs. Performance and
// watch sessions are intentionally in-memory (handler state): a restart ends
// performances, like disconnects do — flagged scope limit, recorded here.
// Testbed fork; generic only.
package database

import (
	"fmt"
	"time"
)

// EnsureServiceSchema creates Phase 6 tables if absent.
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

// AddBuff records a buff, rejecting a second active buff of the same pool type
// (GDD 23.3/24.2.3 same-type non-stacking, cross-source included).
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
	_, err := db.conn.Exec(
		`INSERT INTO buffs (id, target_character_id, source_character_id,
			buff_type, amount, expires_at)
		VALUES (?, ?, ?, ?, ?, ?)`,
		newRowID("buff"),
		target, source, buffType, amount,
		expires.UTC().Format("2006-01-02 15:04:05"))
	return err
}

// ActiveBuffs returns a character's unexpired buffs.
func (db *DB) ActiveBuffs(target string) ([]BuffRow, error) {
	rows, err := db.conn.Query(
		`SELECT id, target_character_id, source_character_id, buff_type, amount,
			applied_at, expires_at FROM buffs
		WHERE target_character_id = ? AND expires_at > CURRENT_TIMESTAMP`,
		target)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var out []BuffRow
	for rows.Next() {
		var b BuffRow
		var applied, expires string
		if err := rows.Scan(&b.ID, &b.TargetID, &b.SourceID, &b.BuffType,
			&b.Amount, &applied, &expires); err != nil {
			return nil, err
		}
		b.AppliedAt, _ = parseServiceTime(applied)
		b.ExpiresAt, _ = parseServiceTime(expires)
		out = append(out, b)
	}
	return out, rows.Err()
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
func (db *DB) UpdateItemStats(owner, itemID, statsJSON string) error {
	_, err := db.conn.Exec(
		`UPDATE crafted_items SET stats_json = ? WHERE id = ? AND owner_character_id = ?`,
		statsJSON, itemID, owner)
	return err
}

func parseServiceTime(s string) (time.Time, bool) {
	for _, layout := range []string{
		"2006-01-02 15:04:05",
		"2006-01-02 15:04:05.999999999-07:00",
		"2006-01-02T15:04:05Z07:00",
		time.RFC3339,
	} {
		if t, err := time.Parse(layout, s); err == nil {
			return t, true
		}
	}
	return time.Time{}, false
}
