// FactionDB extends the DB with Phase 8 persistence: factions, character
// faction associations, PvP flags, recruiters, battlefield zones, duels, wars,
// and PvP events. Testbed fork; generic placeholder content only (R-03:
// FACTION_A / FACTION_B).
//
// Design authority: proposals/proposal-claude-tcindustries-phase8-faction-pvp-v1.md
package database

import (
	"database/sql"
	"fmt"
	"time"
)

// EnsureFactionSchema creates Phase 8 tables and seeds defaults if absent.
// Called at server startup; idempotent.
func (db *DB) EnsureFactionSchema() error {
	stmts := []string{
		`CREATE TABLE IF NOT EXISTS factions (
			id TEXT PRIMARY KEY,
			code TEXT NOT NULL UNIQUE,
			display_name TEXT NOT NULL,
			created_at INTEGER NOT NULL
		)`,
		`CREATE TABLE IF NOT EXISTS character_factions (
			character_id TEXT PRIMARY KEY REFERENCES characters(id),
			faction_id TEXT NOT NULL REFERENCES factions(id),
			standing INTEGER NOT NULL DEFAULT 0,
			joined_at INTEGER NOT NULL,
			changed_at INTEGER NOT NULL
		)`,
		`CREATE TABLE IF NOT EXISTS pvp_flags (
			character_id TEXT PRIMARY KEY REFERENCES characters(id),
			state TEXT NOT NULL CHECK (state IN ('covert','overt')),
			pending_state TEXT CHECK (pending_state IN ('covert','overt')),
			transition_at INTEGER,
			aggression_expires_at INTEGER NOT NULL DEFAULT 0,
			last_aggressor_id TEXT,
			updated_at INTEGER NOT NULL
		)`,
		`CREATE TABLE IF NOT EXISTS faction_recruiters (
			id TEXT PRIMARY KEY,
			faction_id TEXT NOT NULL REFERENCES factions(id),
			name TEXT NOT NULL,
			zone TEXT NOT NULL,
			pos_x REAL NOT NULL,
			pos_z REAL NOT NULL,
			interact_m REAL NOT NULL DEFAULT 8.0
		)`,
		`CREATE TABLE IF NOT EXISTS pvp_zones (
			id TEXT PRIMARY KEY,
			name TEXT NOT NULL,
			zone TEXT NOT NULL,
			shape TEXT NOT NULL CHECK (shape IN ('circle','rect')),
			center_x REAL, center_z REAL, radius_m REAL,
			min_x REAL, min_z REAL, max_x REAL, max_z REAL,
			active INTEGER NOT NULL DEFAULT 1
		)`,
		`CREATE TABLE IF NOT EXISTS pvp_duels (
			id TEXT PRIMARY KEY,
			challenger_id TEXT NOT NULL REFERENCES characters(id),
			target_id TEXT NOT NULL REFERENCES characters(id),
			status TEXT NOT NULL CHECK (status IN ('pending','active','ended','declined','expired')),
			created_at INTEGER NOT NULL,
			started_at INTEGER,
			ended_at INTEGER,
			winner_id TEXT,
			end_reason TEXT
		)`,
		`CREATE TABLE IF NOT EXISTS guild_wars (
			id TEXT PRIMARY KEY,
			declaring_guild_id TEXT NOT NULL REFERENCES guilds(id),
			target_guild_id TEXT NOT NULL REFERENCES guilds(id),
			declared_by_character_id TEXT NOT NULL REFERENCES characters(id),
			status TEXT NOT NULL CHECK (status IN ('declared','active','ended')),
			declared_at INTEGER NOT NULL,
			active_from INTEGER NOT NULL,
			ended_at INTEGER,
			ended_by_character_id TEXT,
			end_reason TEXT
		)`,
		`CREATE TABLE IF NOT EXISTS pvp_events (
			id TEXT PRIMARY KEY,
			kind TEXT NOT NULL,
			basis TEXT NOT NULL,
			attacker_id TEXT NOT NULL,
			target_id TEXT NOT NULL,
			zone_id TEXT,
			created_at INTEGER NOT NULL
		)`,
	}
	for i, s := range stmts {
		if _, err := db.conn.Exec(s); err != nil {
			return fmt.Errorf("faction schema %d failed: %w", i+1, err)
		}
	}

	for _, idx := range []string{
		`CREATE INDEX IF NOT EXISTS idx_guild_wars_status ON guild_wars(status)`,
		`CREATE INDEX IF NOT EXISTS idx_pvp_duels_status ON pvp_duels(status)`,
		`CREATE INDEX IF NOT EXISTS idx_pvp_events_attacker ON pvp_events(attacker_id)`,
	} {
		if _, err := db.conn.Exec(idx); err != nil {
			return fmt.Errorf("faction index failed: %w", err)
		}
	}

	// Seed data (R-03, D-15, D-16): idempotent.
	now := time.Now().Unix()
	for _, f := range []struct{ id, code, name string }{
		{"faction-a-01", "FACTION_A", "Faction Alpha"},
		{"faction-b-01", "FACTION_B", "Faction Bravo"},
	} {
		if _, err := db.conn.Exec(
			`INSERT OR IGNORE INTO factions (id, code, display_name, created_at) VALUES (?, ?, ?, ?)`,
			f.id, f.code, f.name, now,
		); err != nil {
			return fmt.Errorf("seed faction failed: %w", err)
		}
	}
	for _, r := range []struct{ id, fid, name, zone string; x, z float64 }{
		{"recruiter-a-01", "faction-a-01", "Alpha Recruiter", "zone-0001", 3568.0, -4804.0},
		{"recruiter-b-01", "faction-b-01", "Bravo Recruiter", "zone-0001", 3488.0, -4804.0},
	} {
		if _, err := db.conn.Exec(
			`INSERT OR IGNORE INTO faction_recruiters (id, faction_id, name, zone, pos_x, pos_z, interact_m)
			VALUES (?, ?, ?, ?, ?, ?, 8.0)`, r.id, r.fid, r.name, r.zone, r.x, r.z,
		); err != nil {
			return fmt.Errorf("seed recruiter failed: %w", err)
		}
	}
	for _, z := range []struct {
		id, name, zone, shape string
		cx, cz, r, minx, minz, maxx, maxz float64
	}{
		{"zone-circle-01", "Outpost Battlefield", "zone-0001", "circle", 3800.0, -4804.0, 150.0, 0, 0, 0, 0},
		{"zone-rect-01", "Canyon Arena", "zone-0001", "rect", 0, 0, 0, 3900.0, -4900.0, 4200.0, -4700.0},
	} {
		if _, err := db.conn.Exec(
			`INSERT OR IGNORE INTO pvp_zones
			(id, name, zone, shape, center_x, center_z, radius_m, min_x, min_z, max_x, max_z, active)
			VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, 1)`,
			z.id, z.name, z.zone, z.shape, z.cx, z.cz, z.r, z.minx, z.minz, z.maxx, z.maxz,
		); err != nil {
			return fmt.Errorf("seed pvp zone failed: %w", err)
		}
	}

	return nil
}

// --- Factions ---

type FactionRow struct {
	ID          string
	Code        string
	DisplayName string
}

// GetFactionByID loads a faction by its ID.
func (db *DB) GetFactionByID(id string) (*FactionRow, error) {
	var f FactionRow
	err := db.conn.QueryRow(
		`SELECT id, code, display_name FROM factions WHERE id = ?`, id,
	).Scan(&f.ID, &f.Code, &f.DisplayName)
	if err != nil {
		return nil, err
	}
	return &f, nil
}

// ListFactionRecruiters loads all seeded recruiters.
func (db *DB) ListFactionRecruiters() ([]RecruiterRow, error) {
	rows, err := db.conn.Query(
		`SELECT id, faction_id, name, zone, pos_x, pos_z, interact_m FROM faction_recruiters`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var out []RecruiterRow
	for rows.Next() {
		var r RecruiterRow
		if err := rows.Scan(&r.ID, &r.FactionID, &r.Name, &r.Zone, &r.PosX, &r.PosZ, &r.InteractM); err != nil {
			return nil, err
		}
		out = append(out, r)
	}
	return out, rows.Err()
}

type RecruiterRow struct {
	ID        string
	FactionID string
	Name      string
	Zone      string
	PosX      float64
	PosZ      float64
	InteractM float64
}

// --- Character faction membership (R-04) ---

type CharacterFactionRow struct {
	CharacterID string
	FactionID   string
	Standing    int
	JoinedAt    int64
	ChangedAt   int64
}

// GetCharacterFaction loads a character's membership (sql.ErrNoRows = neutral).
func (db *DB) GetCharacterFaction(charID string) (*CharacterFactionRow, error) {
	var r CharacterFactionRow
	err := db.conn.QueryRow(
		`SELECT character_id, faction_id, standing, joined_at, changed_at
		FROM character_factions WHERE character_id = ?`, charID,
	).Scan(&r.CharacterID, &r.FactionID, &r.Standing, &r.JoinedAt, &r.ChangedAt)
	if err != nil {
		return nil, err
	}
	return &r, nil
}

// SetCharacterFaction joins or switches a faction. D-05 resets standing to 0.
func (db *DB) SetCharacterFaction(charID, factionID string) error {
	now := time.Now().Unix()
	_, err := db.conn.Exec(
		`INSERT INTO character_factions (character_id, faction_id, standing, joined_at, changed_at)
		VALUES (?, ?, 0, ?, ?)
		ON CONFLICT(character_id) DO UPDATE SET faction_id = ?, standing = 0, changed_at = ?`,
		charID, factionID, now, now, factionID, now,
	)
	return err
}

// RemoveCharacterFaction leaves a faction, returning the character to neutral.
func (db *DB) RemoveCharacterFaction(charID string) error {
	tx, err := db.conn.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()
	if _, err := tx.Exec(`DELETE FROM character_factions WHERE character_id = ?`, charID); err != nil {
		return err
	}
	if _, err := tx.Exec(`DELETE FROM pvp_flags WHERE character_id = ?`, charID); err != nil {
		return err
	}
	return tx.Commit()
}

// --- PvP flags (R-05) ---

type PvPFlagRow struct {
	CharacterID         string
	State               string
	PendingState        sql.NullString
	TransitionAt        sql.NullInt64
	AggressionExpiresAt int64
	LastAggressorID     sql.NullString
	UpdatedAt           int64
}

// GetPvPFlag loads a character's flag (sql.ErrNoRows if never created).
func (db *DB) GetPvPFlag(charID string) (*PvPFlagRow, error) {
	var r PvPFlagRow
	err := db.conn.QueryRow(
		`SELECT character_id, state, pending_state, transition_at, aggression_expires_at, last_aggressor_id, updated_at
		FROM pvp_flags WHERE character_id = ?`, charID,
	).Scan(&r.CharacterID, &r.State, &r.PendingState, &r.TransitionAt, &r.AggressionExpiresAt, &r.LastAggressorID, &r.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &r, nil
}

// UpsertPvPFlag creates or updates a character's flag row.
func (db *DB) UpsertPvPFlag(r *PvPFlagRow) error {
	now := time.Now().Unix()
	_, err := db.conn.Exec(
		`INSERT INTO pvp_flags
		(character_id, state, pending_state, transition_at, aggression_expires_at, last_aggressor_id, updated_at)
		VALUES (?, ?, ?, ?, ?, ?, ?)
		ON CONFLICT(character_id) DO UPDATE SET state = ?, pending_state = ?, transition_at = ?,
			aggression_expires_at = ?, last_aggressor_id = ?, updated_at = ?`,
		r.CharacterID, r.State, r.PendingState, r.TransitionAt, r.AggressionExpiresAt, r.LastAggressorID, now,
		r.State, r.PendingState, r.TransitionAt, r.AggressionExpiresAt, r.LastAggressorID, now,
	)
	return err
}

// DeletePvPFlag clears a character's flag row.
func (db *DB) DeletePvPFlag(charID string) error {
	_, err := db.conn.Exec(`DELETE FROM pvp_flags WHERE character_id = ?`, charID)
	return err
}

// --- Battlefield zones (R-08) ---

type ZoneRow struct {
	ID                        string
	Name, Zone, Shape         string
	CenterX, CenterZ, RadiusM float64
	MinX, MinZ, MaxX, MaxZ    float64
	Active                    bool
}

// ListZones returns all active battlefield zones.
func (db *DB) ListZones() ([]ZoneRow, error) {
	rows, err := db.conn.Query(
		`SELECT id, name, zone, shape, center_x, center_z, radius_m, min_x, min_z, max_x, max_z, active
		FROM pvp_zones WHERE active = 1`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var out []ZoneRow
	for rows.Next() {
		var z ZoneRow
		var cx, cz, r, minx, minz, maxx, maxz sql.NullFloat64
		var active int
		if err := rows.Scan(&z.ID, &z.Name, &z.Zone, &z.Shape, &cx, &cz, &r, &minx, &minz, &maxx, &maxz, &active); err != nil {
			return nil, err
		}
		z.CenterX, z.CenterZ, z.RadiusM = cx.Float64, cz.Float64, r.Float64
		z.MinX, z.MinZ, z.MaxX, z.MaxZ = minx.Float64, minz.Float64, maxx.Float64, maxz.Float64
		z.Active = active != 0
		out = append(out, z)
	}
	return out, rows.Err()
}

// --- Duels (R-01a) ---

type DuelRow struct {
	ID           string
	ChallengerID string
	TargetID     string
	Status       string // pending|active|ended|declined|expired
	CreatedAt    int64
	StartedAt    sql.NullInt64
	EndedAt      sql.NullInt64
	WinnerID     sql.NullString
	EndReason    sql.NullString
}

// CreateDuel inserts a pending duel challenge.
func (db *DB) CreateDuel(id, challengerID, targetID string) error {
	now := time.Now().Unix()
	_, err := db.conn.Exec(
		`INSERT INTO pvp_duels (id, challenger_id, target_id, status, created_at)
		VALUES (?, ?, ?, 'pending', ?)`, id, challengerID, targetID, now,
	)
	return err
}

// GetDuel loads a duel by ID.
func (db *DB) GetDuel(id string) (*DuelRow, error) {
	var d DuelRow
	err := db.conn.QueryRow(
		`SELECT id, challenger_id, target_id, status, created_at, started_at, ended_at, winner_id, end_reason
		FROM pvp_duels WHERE id = ?`, id,
	).Scan(&d.ID, &d.ChallengerID, &d.TargetID, &d.Status, &d.CreatedAt, &d.StartedAt, &d.EndedAt, &d.WinnerID, &d.EndReason)
	if err != nil {
		return nil, err
	}
	return &d, nil
}

// ActiveDuelFor returns a character's pending/active duel, if any.
func (db *DB) ActiveDuelFor(charID string) (*DuelRow, error) {
	var d DuelRow
	err := db.conn.QueryRow(
		`SELECT id, challenger_id, target_id, status, created_at, started_at, ended_at, winner_id, end_reason
		FROM pvp_duels
		WHERE (challenger_id = ? OR target_id = ?) AND status IN ('pending', 'active')
		LIMIT 1`, charID, charID,
	).Scan(&d.ID, &d.ChallengerID, &d.TargetID, &d.Status, &d.CreatedAt, &d.StartedAt, &d.EndedAt, &d.WinnerID, &d.EndReason)
	if err != nil {
		return nil, err
	}
	return &d, nil
}

// UpdateDuelStatus updates a duel's state.
func (db *DB) UpdateDuelStatus(id, status string, startedAt, endedAt sql.NullInt64, winnerID, endReason sql.NullString) error {
	_, err := db.conn.Exec(
		`UPDATE pvp_duels SET status = ?, started_at = ?, ended_at = ?, winner_id = ?, end_reason = ?
		WHERE id = ?`, status, startedAt, endedAt, winnerID, endReason, id,
	)
	return err
}

// --- Guild wars (R-01a, R-07) ---

type GuildWarRow struct {
	ID                    string
	DeclaringGuildID      string
	TargetGuildID         string
	DeclaredByCharacterID string
	Status                string // declared|active|ended
	DeclaredAt            int64
	ActiveFrom            int64
	EndedAt               sql.NullInt64
	EndedByCharacterID    sql.NullString
	EndReason             sql.NullString
}

// CreateGuildWar records a guild war declaration.
func (db *DB) CreateGuildWar(id, declaringGuildID, targetGuildID, declaredBy string, activeFrom int64) error {
	now := time.Now().Unix()
	_, err := db.conn.Exec(
		`INSERT INTO guild_wars
		(id, declaring_guild_id, target_guild_id, declared_by_character_id, status, declared_at, active_from)
		VALUES (?, ?, ?, ?, 'declared', ?, ?)`,
		id, declaringGuildID, targetGuildID, declaredBy, now, activeFrom,
	)
	return err
}

// AnyActiveWar reports whether g1 and g2 are at active war.
func (db *DB) AnyActiveWar(g1, g2 string) bool {
	var n int
	_ = db.conn.QueryRow(
		`SELECT COUNT(*) FROM guild_wars
		WHERE status = 'active'
		AND ((declaring_guild_id = ? AND target_guild_id = ?)
		OR (declaring_guild_id = ? AND target_guild_id = ?))`,
		g1, g2, g2, g1,
	).Scan(&n)
	return n > 0
}

// AnyPendingOrActiveWar blocks duplicate declarations on the same pair.
func (db *DB) AnyPendingOrActiveWar(g1, g2 string) bool {
	var n int
	_ = db.conn.QueryRow(
		`SELECT COUNT(*) FROM guild_wars
		WHERE status IN ('declared','active')
		AND ((declaring_guild_id = ? AND target_guild_id = ?)
		OR (declaring_guild_id = ? AND target_guild_id = ?))`,
		g1, g2, g2, g1,
	).Scan(&n)
	return n > 0
}

// GuildWarsFor lists wars involving a guild (either side).
func (db *DB) GuildWarsFor(guildID string) ([]GuildWarRow, error) {
	rows, err := db.conn.Query(
		`SELECT id, declaring_guild_id, target_guild_id, declared_by_character_id,
			status, declared_at, active_from, ended_at, ended_by_character_id, end_reason
		FROM guild_wars
		WHERE declaring_guild_id = ? OR target_guild_id = ?`, guildID, guildID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var out []GuildWarRow
	for rows.Next() {
		var w GuildWarRow
		if err := rows.Scan(&w.ID, &w.DeclaringGuildID, &w.TargetGuildID, &w.DeclaredByCharacterID,
			&w.Status, &w.DeclaredAt, &w.ActiveFrom, &w.EndedAt, &w.EndedByCharacterID, &w.EndReason); err != nil {
			return nil, err
		}
		out = append(out, w)
	}
	return out, rows.Err()
}

// GetGuildWar loads one war by ID.
func (db *DB) GetGuildWar(id string) (*GuildWarRow, error) {
	var w GuildWarRow
	err := db.conn.QueryRow(
		`SELECT id, declaring_guild_id, target_guild_id, declared_by_character_id,
			status, declared_at, active_from, ended_at, ended_by_character_id, end_reason
		FROM guild_wars WHERE id = ?`, id,
	).Scan(&w.ID, &w.DeclaringGuildID, &w.TargetGuildID, &w.DeclaredByCharacterID,
		&w.Status, &w.DeclaredAt, &w.ActiveFrom, &w.EndedAt, &w.EndedByCharacterID, &w.EndReason)
	if err != nil {
		return nil, err
	}
	return &w, nil
}

// EndGuildWar marks a war ended with the surrendering character and reason.
func (db *DB) EndGuildWar(id, endedBy, reason string) error {
	now := time.Now().Unix()
	_, err := db.conn.Exec(
		`UPDATE guild_wars SET status = 'ended', ended_at = ?, ended_by_character_id = ?, end_reason = ?
		WHERE id = ? AND status != 'ended'`, now, endedBy, reason, id)
	return err
}

// ActivateDueWars flips declared → active once active_from passes (D-11).
// Returns the number activated.
func (db *DB) ActivateDueWars(now time.Time) (int, error) {
	res, err := db.conn.Exec(
		`UPDATE guild_wars SET status = 'active' WHERE status = 'declared' AND active_from <= ?`,
		now.Unix())
	if err != nil {
		return 0, err
	}
	n, _ := res.RowsAffected()
	return int(n), nil
}

// --- PvP telemetry (D-06) ---

// RecordPvPEvent inserts a pvp_events row (kills are logged, never awarded).
func (db *DB) RecordPvPEvent(id, kind, basis, attackerID, targetID string, zoneID sql.NullString) error {
	now := time.Now().Unix()
	_, err := db.conn.Exec(
		`INSERT INTO pvp_events (id, kind, basis, attacker_id, target_id, zone_id, created_at)
		VALUES (?, ?, ?, ?, ?, ?, ?)`, id, kind, basis, attackerID, targetID, zoneID, now,
	)
	return err
}

// PvPEventRow maps one pvp_events row.
type PvPEventRow struct {
	ID         string
	Kind       string
	Basis      string
	AttackerID string
	TargetID   string
	ZoneID     sql.NullString
	CreatedAt  int64
}

// PvPEventsFor returns a character's defeat events (as attacker or target),
// newest first. Without a reader this table is write-only and unverifiable.
func (db *DB) PvPEventsFor(charID string) ([]PvPEventRow, error) {
	rows, err := db.conn.Query(
		`SELECT id, kind, basis, attacker_id, target_id, zone_id, created_at
		FROM pvp_events
		WHERE attacker_id = ? OR target_id = ?
		ORDER BY id DESC`, charID, charID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var out []PvPEventRow
	for rows.Next() {
		var e PvPEventRow
		if err := rows.Scan(&e.ID, &e.Kind, &e.Basis, &e.AttackerID, &e.TargetID, &e.ZoneID, &e.CreatedAt); err != nil {
			return nil, err
		}
		out = append(out, e)
	}
	return out, rows.Err()
}
