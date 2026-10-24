package database

import (
	"time"
)

// ActiveBuffs returns a character's unexpired buffs (used by the /api/services
// buffs route). The credit-ledger writer that used to live in this file moved to
// economy_db.go, which owns the ledger_entries schema.
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
		b.AppliedAt, _ = time.Parse("2006-01-02 15:04:05", applied)
		b.ExpiresAt, _ = time.Parse("2006-01-02 15:04:05", expires)
		out = append(out, b)
	}
	return out, rows.Err()
}
