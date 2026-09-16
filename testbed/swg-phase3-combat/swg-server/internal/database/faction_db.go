// FactionDB extends the DB with Phase 8 persistence: standings/flagging,
// pending PvP kills, point events, faction bases + damage log, item condition,
// city PvP permission, and guild alignment labels. Testbed fork; generic only.
//
// Kill-credit design (flagged): the killing blow writes a pending row at
// incapacitation time; cloning (death) consumes it into points + credit
// transfer + condition loss; a successful revive voids it (survived — no kill).
// Multi-attacker contribution is NOT split (GDD-silent; flagged omission).
// Wall-clock deadlines are INTEGER unix seconds (civic convention).
package database

import (
	"database/sql"
	"fmt"
	"time"
)

// EnsureFactionSchema creates Phase 8 tables if absent. Idempotent.
func (db *DB) EnsureFactionSchema() error {
	stmts := []string{
		`CREATE TABLE IF NOT EXISTS faction_standings (
			character_id TEXT PRIMARY KEY,
			alignment TEXT NOT NULL DEFAULT 'neutral',
			points INTEGER NOT NULL DEFAULT 0,
			overt INTEGER NOT NULL DEFAULT 0,
			overt_since INTEGER,
			last_change INTEGER
		)`,
		`CREATE TABLE IF NOT EXISTS pvp_pending (
			victim_id TEXT PRIMARY KEY,
			killer_id TEXT NOT NULL,
			at INTEGER NOT NULL
		)`,
		`CREATE TABLE IF NOT EXISTS pvp_point_events (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			character_id TEXT NOT NULL,
			amount INTEGER NOT NULL,
			reason TEXT NOT NULL,
			counterparty TEXT NOT NULL DEFAULT '',
			created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
		)`,
		`CREATE TABLE IF NOT EXISTS faction_bases (
			id TEXT PRIMARY KEY,
			guild_id TEXT NOT NULL,
			faction TEXT NOT NULL,
			zone TEXT NOT NULL,
			pos_x REAL NOT NULL,
			pos_z REAL NOT NULL,
			hp INTEGER NOT NULL,
			hp_max INTEGER NOT NULL,
			window_start INTEGER,
			window_end INTEGER,
			status TEXT NOT NULL DEFAULT 'active',
			placed_by TEXT NOT NULL,
			created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
		)`,
		`CREATE TABLE IF NOT EXISTS base_damage_log (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			base_id TEXT NOT NULL,
			character_id TEXT NOT NULL,
			amount INTEGER NOT NULL,
			at INTEGER NOT NULL
		)`,
	}
	for i, s := range stmts {
		if _, err := db.conn.Exec(s); err != nil {
			return fmt.Errorf("faction schema %d failed: %w", i+1, err)
		}
	}
	// PvP death condition (proposal §5.d minimal surface for OQ-010).
	if _, err := db.conn.Exec(
		`ALTER TABLE crafted_items ADD COLUMN condition_pct INTEGER NOT NULL DEFAULT 100`); err != nil {
		if !isDupColumn(err.Error()) {
			return fmt.Errorf("crafted_items condition column failed: %w", err)
		}
	}
	// City PvP permission (GDD 9.4.2 "player cities (if city allows)").
	if _, err := db.conn.Exec(
		`ALTER TABLE cities ADD COLUMN pvp_allowed INTEGER NOT NULL DEFAULT 1`); err != nil {
		if !isDupColumn(err.Error()) {
			return fmt.Errorf("cities pvp_allowed column failed: %w", err)
		}
	}
	return nil
}

// --- Standings ---

// StandingRow maps one faction_standings row.
type StandingRow struct {
	CharacterID string
	Alignment   string
	Points      int
	Overt       bool
	OvertSince  sql.NullInt64
	LastChange  sql.NullInt64
}

// GetStanding loads a standing, defaulting to neutral/covert for unknowns
// (neutral is the safe default per §15.1 — never a valid target).
func (db *DB) GetStanding(charID string) (*StandingRow, error) {
	var s StandingRow
	var overt int
	err := db.conn.QueryRow(
		`SELECT character_id, alignment, points, overt, overt_since, last_change
		FROM faction_standings WHERE character_id = ?`, charID,
	).Scan(&s.CharacterID, &s.Alignment, &s.Points, &overt, &s.OvertSince, &s.LastChange)
	if err == sql.ErrNoRows {
		return &StandingRow{CharacterID: charID, Alignment: "neutral"}, nil
	}
	if err != nil {
		return nil, err
	}
	s.Overt = overt != 0
	return &s, nil
}

// SetStanding inserts or replaces a standing row.
func (db *DB) SetStanding(s *StandingRow) error {
	overt := 0
	if s.Overt {
		overt = 1
	}
	var since, changed interface{}
	if s.OvertSince.Valid {
		since = s.OvertSince.Int64
	}
	if s.LastChange.Valid {
		changed = s.LastChange.Int64
	}
	_, err := db.conn.Exec(
		`INSERT INTO faction_standings
			(character_id, alignment, points, overt, overt_since, last_change)
		VALUES (?, ?, ?, ?, ?, ?)
		ON CONFLICT(character_id) DO UPDATE SET alignment = ?, points = ?,
			overt = ?, overt_since = ?, last_change = ?`,
		s.CharacterID, s.Alignment, s.Points, overt, since, changed,
		s.Alignment, s.Points, overt, since, changed)
	return err
}

// AddPoints adds faction points and records the event (observability per
// SAFE-003; no enforcement attached — OQ-012 stays TBD).
func (db *DB) AddPoints(charID string, amount int, reason, counterparty string) error {
	tx, err := db.conn.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()
	if _, err := tx.Exec(
		`INSERT INTO faction_standings (character_id, alignment, points)
		VALUES (?, 'neutral', ?)
		ON CONFLICT(character_id) DO UPDATE SET points = points + ?`,
		charID, amount, amount); err != nil {
		return err
	}
	if _, err := tx.Exec(
		`INSERT INTO pvp_point_events (character_id, amount, reason, counterparty)
		VALUES (?, ?, ?, ?)`, charID, amount, reason, counterparty); err != nil {
		return err
	}
	return tx.Commit()
}

// PointEvent maps one point-events row.
type PointEvent struct {
	ID           int64
	Amount       int
	Reason       string
	Counterparty string
}

// PointEvents returns a character's point history (newest last).
func (db *DB) PointEvents(charID string) ([]PointEvent, error) {
	rows, err := db.conn.Query(
		`SELECT id, amount, reason, counterparty FROM pvp_point_events
		WHERE character_id = ? ORDER BY id`, charID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var out []PointEvent
	for rows.Next() {
		var e PointEvent
		if err := rows.Scan(&e.ID, &e.Amount, &e.Reason, &e.Counterparty); err != nil {
			return nil, err
		}
		out = append(out, e)
	}
	return out, rows.Err()
}

// LatestPointEventAt returns the unix time of the newest point event, or an
// error when the character has no history (decay staleness source).
func (db *DB) LatestPointEventAt(charID string) (int64, error) {
	var created string
	err := db.conn.QueryRow(
		`SELECT created_at FROM pvp_point_events WHERE character_id = ?
		ORDER BY id DESC LIMIT 1`, charID).Scan(&created)
	if err != nil {
		return 0, err
	}
	for _, layout := range []string{
		"2006-01-02 15:04:05", "2006-01-02 15:04:05.999999999-07:00",
		"2006-01-02T15:04:05Z07:00", "2006-01-02 15:04:05-07:00",
	} {
		if t, perr := time.Parse(layout, created); perr == nil {
			return t.Unix(), nil
		}
	}
	return 0, fmt.Errorf("unparseable event time")
}

// OvertMembers returns online-candidate overt character IDs of one alignment
// (Phase 9 bounty input surface; presence filtered by the caller).
func (db *DB) OvertMembers(alignment string) ([]string, error) {
	rows, err := db.conn.Query(
		`SELECT character_id FROM faction_standings
		WHERE alignment = ? AND overt = 1`, alignment)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var out []string
	for rows.Next() {
		var id string
		if err := rows.Scan(&id); err != nil {
			return nil, err
		}
		out = append(out, id)
	}
	return out, rows.Err()
}

// --- Pending kills ---

// RecordPendingKill credits the killing blow at incap time (one row per
// victim; victims must be conscious to be attacked, so no overwrite race).
func (db *DB) RecordPendingKill(victimID, killerID string, now int64) error {
	_, err := db.conn.Exec(
		`INSERT OR REPLACE INTO pvp_pending (victim_id, killer_id, at)
		VALUES (?, ?, ?)`, victimID, killerID, now)
	return err
}

// TakePendingKill loads and deletes the pending row (death consumes it).
func (db *DB) TakePendingKill(victimID string) (string, bool) {
	tx, err := db.conn.Begin()
	if err != nil {
		return "", false
	}
	defer tx.Rollback()
	var killer string
	if err := tx.QueryRow(
		`SELECT killer_id FROM pvp_pending WHERE victim_id = ?`,
		victimID).Scan(&killer); err != nil {
		return "", false
	}
	if _, err := tx.Exec(
		`DELETE FROM pvp_pending WHERE victim_id = ?`, victimID); err != nil {
		return "", false
	}
	if err := tx.Commit(); err != nil {
		return "", false
	}
	return killer, true
}

// VoidPendingKill deletes the pending row (revive voids the kill credit).
func (db *DB) VoidPendingKill(victimID string) error {
	_, err := db.conn.Exec(
		`DELETE FROM pvp_pending WHERE victim_id = ?`, victimID)
	return err
}

// ApplyPvPDeath executes the death legs in ONE transaction: killer points +
// events, victim→killer 10% credit transfer + ledger pair, equipped condition
// loss. Called from the clone path (incap-timer deaths route through the same
// cloneCharacter, so both death kinds are covered).
func (db *DB) ApplyPvPDeath(victimID, killerID string, creditPct, condLoss, killAward int, zone string) (transferred int, err error) {
	tx, err := db.conn.Begin()
	if err != nil {
		return 0, err
	}
	defer tx.Rollback()
	var carried int
	if err := tx.QueryRow(
		`SELECT credits FROM characters WHERE id = ?`, victimID).Scan(&carried); err != nil {
		return 0, fmt.Errorf("victim unknown: %w", err)
	}
	transferred = carried * creditPct / 100
	if transferred > 0 {
		if _, err := tx.Exec(
			`UPDATE characters SET credits = credits - ? WHERE id = ?`,
			transferred, victimID); err != nil {
			return 0, err
		}
		if _, err := tx.Exec(
			`UPDATE characters SET credits = credits + ? WHERE id = ?`,
			transferred, killerID); err != nil {
			return 0, err
		}
		if err := RecordLedgerTx(tx, victimID, -transferred,
			"pvp_death_drop", "transfer", killerID, zone); err != nil {
			return 0, err
		}
		if err := RecordLedgerTx(tx, killerID, transferred,
			"pvp_death_drop", "transfer", victimID, zone); err != nil {
			return 0, err
		}
	}
	if _, err := tx.Exec(
		`UPDATE crafted_items SET condition_pct = MAX(0, condition_pct - ?)
		WHERE owner_character_id = ? AND equipped != 0`,
		condLoss, victimID); err != nil {
		return 0, err
	}
	if _, err := tx.Exec(
		`UPDATE faction_standings SET points = points + ? WHERE character_id = ?`,
		killAward, killerID); err != nil {
		return 0, err
	}
	if _, err := tx.Exec(
		`INSERT INTO pvp_point_events (character_id, amount, reason, counterparty)
		VALUES (?, ?, 'pvp_kill', ?)`, killerID, killAward, victimID); err != nil {
		return 0, err
	}
	if err := tx.Commit(); err != nil {
		return 0, err
	}
	return transferred, nil
}

// --- Bases ---

// BaseRow maps one faction_bases row.
type BaseRow struct {
	ID          string
	GuildID     string
	Faction     string
	Zone        string
	PosX, PosZ  float64
	HP          int
	HPMax       int
	WindowStart sql.NullInt64
	WindowEnd   sql.NullInt64
	Status      string
	PlacedBy    string
}

// PlaceBase inserts a base row.
func (db *DB) PlaceBase(b *BaseRow) error {
	var ws, we interface{}
	if b.WindowStart.Valid {
		ws = b.WindowStart.Int64
	}
	if b.WindowEnd.Valid {
		we = b.WindowEnd.Int64
	}
	_, err := db.conn.Exec(
		`INSERT INTO faction_bases (id, guild_id, faction, zone, pos_x, pos_z,
			hp, hp_max, window_start, window_end, status, placed_by)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, 'active', ?)`,
		b.ID, b.GuildID, b.Faction, b.Zone, b.PosX, b.PosZ,
		b.HP, b.HPMax, ws, we, b.PlacedBy)
	return err
}

// GetBase loads one base.
func (db *DB) GetBase(id string) (*BaseRow, error) {
	var b BaseRow
	err := db.conn.QueryRow(
		`SELECT id, guild_id, faction, zone, pos_x, pos_z, hp, hp_max,
			window_start, window_end, status, placed_by FROM faction_bases
		WHERE id = ?`, id,
	).Scan(&b.ID, &b.GuildID, &b.Faction, &b.Zone, &b.PosX, &b.PosZ,
		&b.HP, &b.HPMax, &b.WindowStart, &b.WindowEnd, &b.Status, &b.PlacedBy)
	if err != nil {
		return nil, err
	}
	return &b, nil
}

// AllBases returns every base (overlap + tick-free on-demand checks).
func (db *DB) AllBases() ([]BaseRow, error) {
	rows, err := db.conn.Query(
		`SELECT id, guild_id, faction, zone, pos_x, pos_z, hp, hp_max,
			window_start, window_end, status, placed_by FROM faction_bases`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var out []BaseRow
	for rows.Next() {
		var b BaseRow
		if err := rows.Scan(&b.ID, &b.GuildID, &b.Faction, &b.Zone, &b.PosX,
			&b.PosZ, &b.HP, &b.HPMax, &b.WindowStart, &b.WindowEnd,
			&b.Status, &b.PlacedBy); err != nil {
			return nil, err
		}
		out = append(out, b)
	}
	return out, rows.Err()
}

// UpdateBase persists HP/status/window changes.
func (db *DB) UpdateBase(b *BaseRow) error {
	var ws, we interface{}
	if b.WindowStart.Valid {
		ws = b.WindowStart.Int64
	}
	if b.WindowEnd.Valid {
		we = b.WindowEnd.Int64
	}
	_, err := db.conn.Exec(
		`UPDATE faction_bases SET hp = ?, status = ?, window_start = ?,
			window_end = ? WHERE id = ?`,
		b.HP, b.Status, ws, we, b.ID)
	return err
}

// GuildActiveBase returns a guild's active base, if any (one active base per
// guild — flagged simplification; GDD is silent on base counts).
func (db *DB) GuildActiveBase(guildID string) (*BaseRow, error) {
	var b BaseRow
	err := db.conn.QueryRow(
		`SELECT id, guild_id, faction, zone, pos_x, pos_z, hp, hp_max,
			window_start, window_end, status, placed_by FROM faction_bases
		WHERE guild_id = ? AND status = 'active' LIMIT 1`, guildID,
	).Scan(&b.ID, &b.GuildID, &b.Faction, &b.Zone, &b.PosX, &b.PosZ,
		&b.HP, &b.HPMax, &b.WindowStart, &b.WindowEnd, &b.Status, &b.PlacedBy)
	if err != nil {
		return nil, err
	}
	return &b, nil
}

// LogBaseDamage records siege participation (attribution source).
func (db *DB) LogBaseDamage(baseID, charID string, amount int, now int64) error {
	_, err := db.conn.Exec(
		`INSERT INTO base_damage_log (base_id, character_id, amount, at)
		VALUES (?, ?, ?, ?)`, baseID, charID, amount, now)
	return err
}

// BaseParticipants returns distinct attackers since a unix timestamp.
func (db *DB) BaseParticipants(baseID string, since int64) ([]string, error) {
	rows, err := db.conn.Query(
		`SELECT DISTINCT character_id FROM base_damage_log
		WHERE base_id = ? AND at >= ?`, baseID, since)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var out []string
	for rows.Next() {
		var id string
		if err := rows.Scan(&id); err != nil {
			return nil, err
		}
		out = append(out, id)
	}
	return out, rows.Err()
}

// AwardDestroy grants the destroy award to every participant.
func (db *DB) AwardDestroy(baseID, reason string, award int, participantIDs []string) error {
	tx, err := db.conn.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()
	for _, id := range participantIDs {
		if _, err := tx.Exec(
			`INSERT INTO faction_standings (character_id, alignment, points)
			VALUES (?, 'neutral', ?)
			ON CONFLICT(character_id) DO UPDATE SET points = points + ?`,
			id, award, award); err != nil {
			return err
		}
		if _, err := tx.Exec(
			`INSERT INTO pvp_point_events (character_id, amount, reason, counterparty)
			VALUES (?, ?, ?, ?)`, id, award, reason, baseID); err != nil {
			return err
		}
	}
	return tx.Commit()
}

// --- City PvP permission + guild alignment ---

// SetCityPvP sets the mayor's PvP permission flag.
func (db *DB) SetCityPvP(cityID string, allowed bool) error {
	v := 0
	if allowed {
		v = 1
	}
	_, err := db.conn.Exec(
		`UPDATE cities SET pvp_allowed = ? WHERE id = ?`, v, cityID)
	return err
}

// CityPvPAllowed reports a city's flag (default allow when unknown).
func (db *DB) CityPvPAllowed(cityID string) bool {
	var v int
	if err := db.conn.QueryRow(
		`SELECT pvp_allowed FROM cities WHERE id = ?`, cityID).Scan(&v); err != nil {
		return true
	}
	return v != 0
}

// SetGuildAlignment sets a guild's alignment label.
func (db *DB) SetGuildAlignment(guildID, alignment string) error {
	_, err := db.conn.Exec(
		`UPDATE guilds SET faction = ? WHERE id = ?`, alignment, guildID)
	return err
}

// GuildOfficerAlignments returns leader + officer personal alignments
// (base-placement gate source: all must match, non-neutral). Two-phase:
// IDs are collected and rows closed BEFORE per-ID standing lookups, because
// the single-connection pool (db.go hardening) cannot serve a nested query
// while outer rows are open — nesting self-deadlocks.
func (db *DB) GuildOfficerAlignments(guildID string) ([]string, error) {
	rows, err := db.conn.Query(
		`SELECT character_id FROM guild_members
		WHERE guild_id = ? AND role IN ('leader','officer')`, guildID)
	if err != nil {
		return nil, err
	}
	var ids []string
	for rows.Next() {
		var id string
		if err := rows.Scan(&id); err != nil {
			rows.Close()
			return nil, err
		}
		ids = append(ids, id)
	}
	if err := rows.Err(); err != nil {
		rows.Close()
		return nil, err
	}
	rows.Close()
	var out []string
	for _, id := range ids {
		st, err := db.GetStanding(id)
		if err != nil {
			return nil, err
		}
		out = append(out, st.Alignment)
	}
	return out, nil
}
