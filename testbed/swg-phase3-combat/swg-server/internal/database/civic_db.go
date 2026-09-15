// CivicDB extends the DB with Phase 7 persistence: cities, citizenship,
// elections, guilds, groups, mentorships, and presence. Mail, waypoints,
// friends, and structure permissions live in civic_social_db.go.
// Testbed fork; generic only. New tables store wall-clock deadlines as INTEGER
// unix seconds (no datetime parsing); audit columns keep DATETIME defaults.
package database

import (
	"database/sql"
	"fmt"
	"time"
)

// EnsureCivicSchema creates Phase 7 tables if absent. Called at server startup.
// Idempotent (CREATE TABLE IF NOT EXISTS + duplicate-column-tolerant ALTERs).
func (db *DB) EnsureCivicSchema() error {
	stmts := []string{
		`CREATE TABLE IF NOT EXISTS cities (
			id TEXT PRIMARY KEY,
			name TEXT NOT NULL,
			zone TEXT NOT NULL,
			center_x REAL NOT NULL,
			center_z REAL NOT NULL,
			radius_m REAL NOT NULL,
			rank TEXT NOT NULL DEFAULT 'outpost',
			treasury INTEGER NOT NULL DEFAULT 0,
			upkeep_weekly INTEGER NOT NULL DEFAULT 0,
			tax_flat_weekly INTEGER NOT NULL DEFAULT 0,
			tax_vendor_pct INTEGER NOT NULL DEFAULT 0,
			specialization TEXT NOT NULL DEFAULT 'none',
			mayor_character_id TEXT,
			term_end INTEGER,
			status TEXT NOT NULL DEFAULT 'forming',
			grace_ends INTEGER,
			zoning_open INTEGER NOT NULL DEFAULT 1,
			created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
		)`,
		`CREATE TABLE IF NOT EXISTS city_citizens (
			city_id TEXT NOT NULL,
			character_id TEXT NOT NULL,
			joined_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
			tax_arrears INTEGER NOT NULL DEFAULT 0,
			is_militia INTEGER NOT NULL DEFAULT 0,
			PRIMARY KEY (city_id, character_id)
		)`,
		`CREATE TABLE IF NOT EXISTS city_elections (
			id TEXT PRIMARY KEY,
			city_id TEXT NOT NULL,
			opened_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
			ends_at INTEGER NOT NULL,
			status TEXT NOT NULL DEFAULT 'open',
			winner_id TEXT
		)`,
		`CREATE TABLE IF NOT EXISTS city_ballots (
			election_id TEXT NOT NULL,
			voter_id TEXT NOT NULL,
			candidate_id TEXT NOT NULL,
			cast_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
			PRIMARY KEY (election_id, voter_id)
		)`,
		`CREATE TABLE IF NOT EXISTS guilds (
			id TEXT PRIMARY KEY,
			name TEXT NOT NULL UNIQUE,
			tag TEXT NOT NULL DEFAULT '',
			leader_id TEXT NOT NULL,
			treasury INTEGER NOT NULL DEFAULT 0,
			dues_pct INTEGER NOT NULL DEFAULT 0,
			faction TEXT NOT NULL DEFAULT 'neutral',
			hall_structure_id TEXT,
			dissolve_at INTEGER,
			founded_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
		)`,
		`CREATE TABLE IF NOT EXISTS guild_members (
			guild_id TEXT NOT NULL,
			character_id TEXT NOT NULL,
			role TEXT NOT NULL DEFAULT 'member',
			joined_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
			PRIMARY KEY (guild_id, character_id)
		)`,
		`CREATE TABLE IF NOT EXISTS guild_invites (
			id TEXT PRIMARY KEY,
			guild_id TEXT NOT NULL,
			inviter_id TEXT NOT NULL,
			invitee_id TEXT NOT NULL,
			status TEXT NOT NULL DEFAULT 'pending',
			created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
		)`,
		`CREATE TABLE IF NOT EXISTS groups (
			id TEXT PRIMARY KEY,
			leader_id TEXT NOT NULL,
			loot_rule TEXT NOT NULL DEFAULT 'round_robin',
			created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
		)`,
		`CREATE TABLE IF NOT EXISTS group_members (
			group_id TEXT NOT NULL,
			character_id TEXT NOT NULL,
			joined_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
			PRIMARY KEY (group_id, character_id)
		)`,
		`CREATE TABLE IF NOT EXISTS group_invites (
			id TEXT PRIMARY KEY,
			group_id TEXT NOT NULL,
			inviter_id TEXT NOT NULL,
			invitee_id TEXT NOT NULL,
			status TEXT NOT NULL DEFAULT 'pending',
			created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
		)`,
		`CREATE TABLE IF NOT EXISTS mentorships (
			mentor_id TEXT NOT NULL,
			protege_id TEXT NOT NULL,
			started_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
			expires_at INTEGER NOT NULL,
			PRIMARY KEY (mentor_id, protege_id)
		)`,
		`CREATE TABLE IF NOT EXISTS civic_presence (
			character_id TEXT PRIMARY KEY,
			last_seen INTEGER NOT NULL
		)`,
		// Opt-outs: leaving or evicted citizens are never auto-re-enrolled
		// (GDD 14.2.1 "citizens may opt out" would be void otherwise).
		`CREATE TABLE IF NOT EXISTS city_optouts (
			city_id TEXT NOT NULL,
			character_id TEXT NOT NULL,
			PRIMARY KEY (city_id, character_id)
		)`,
	}
	for i, s := range stmts {
		if _, err := db.conn.Exec(s); err != nil {
			return fmt.Errorf("civic schema %d failed: %w", i+1, err)
		}
	}
	// Section 13 depth: entry-permission flag on structures (lists below).
	if _, err := db.conn.Exec(
		`ALTER TABLE structures ADD COLUMN entry_perm TEXT NOT NULL DEFAULT 'public'`); err != nil {
		if !isDupColumn(err.Error()) {
			return fmt.Errorf("structures entry_perm column failed: %w", err)
		}
	}
	for _, s := range []string{
		`CREATE TABLE IF NOT EXISTS structure_admins (
			structure_id TEXT NOT NULL, character_id TEXT NOT NULL,
			PRIMARY KEY (structure_id, character_id))`,
		`CREATE TABLE IF NOT EXISTS structure_friends (
			structure_id TEXT NOT NULL, character_id TEXT NOT NULL,
			PRIMARY KEY (structure_id, character_id))`,
		`CREATE TABLE IF NOT EXISTS structure_banned (
			structure_id TEXT NOT NULL, character_id TEXT NOT NULL,
			PRIMARY KEY (structure_id, character_id))`,
	} {
		if _, err := db.conn.Exec(s); err != nil {
			return fmt.Errorf("civic perm schema failed: %w", err)
		}
	}
	return nil
}

// --- Cities ---

// CityRow maps one cities row.
type CityRow struct {
	ID              string
	Name            string
	Zone            string
	CenterX, CenterZ float64
	RadiusM         float64
	Rank            string
	Treasury        int
	UpkeepWeekly    int
	TaxFlatWeekly   int
	TaxVendorPct    int
	Specialization  string
	MayorID         sql.NullString
	TermEnd         sql.NullInt64
	Status          string
	GraceEnds       sql.NullInt64
	ZoningOpen      bool
}

func scanCity(row *sql.Row) (*CityRow, error) {
	var c CityRow
	var zoning int
	if err := row.Scan(&c.ID, &c.Name, &c.Zone, &c.CenterX, &c.CenterZ,
		&c.RadiusM, &c.Rank, &c.Treasury, &c.UpkeepWeekly, &c.TaxFlatWeekly,
		&c.TaxVendorPct, &c.Specialization, &c.MayorID, &c.TermEnd,
		&c.Status, &c.GraceEnds, &zoning); err != nil {
		return nil, err
	}
	c.ZoningOpen = zoning != 0
	return &c, nil
}

const cityCols = `id, name, zone, center_x, center_z, radius_m, rank, treasury,
	upkeep_weekly, tax_flat_weekly, tax_vendor_pct, specialization,
	mayor_character_id, term_end, status, grace_ends, zoning_open`

// CreateCity inserts a city row; returns its ID.
func (db *DB) CreateCity(id, name, zone string, cx, cz, radius float64,
	upkeep int, founderID string, termEnd, graceEnds int64) error {
	_, err := db.conn.Exec(
		`INSERT INTO cities (id, name, zone, center_x, center_z, radius_m, rank,
			treasury, upkeep_weekly, mayor_character_id, term_end, status, grace_ends)
		VALUES (?, ?, ?, ?, ?, ?, 'outpost', 0, ?, ?, ?, 'forming', ?)`,
		id, name, zone, cx, cz, radius, upkeep, founderID, termEnd, graceEnds)
	return err
}

// GetCity loads one city.
func (db *DB) GetCity(id string) (*CityRow, error) {
	return scanCity(db.conn.QueryRow(
		`SELECT `+cityCols+` FROM cities WHERE id = ?`, id))
}

// AllCities returns every city (tick + overlap driver).
func (db *DB) AllCities() ([]CityRow, error) {
	rows, err := db.conn.Query(`SELECT ` + cityCols + ` FROM cities`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var out []CityRow
	for rows.Next() {
		var c CityRow
		var zoning int
		if err := rows.Scan(&c.ID, &c.Name, &c.Zone, &c.CenterX, &c.CenterZ,
			&c.RadiusM, &c.Rank, &c.Treasury, &c.UpkeepWeekly, &c.TaxFlatWeekly,
			&c.TaxVendorPct, &c.Specialization, &c.MayorID, &c.TermEnd,
			&c.Status, &c.GraceEnds, &zoning); err != nil {
			return nil, err
		}
		c.ZoningOpen = zoning != 0
		out = append(out, c)
	}
	return out, rows.Err()
}

// UpdateCity persists treasury/rank/tax/mayor/status changes.
func (db *DB) UpdateCity(c *CityRow) error {
	var mayor interface{}
	if c.MayorID.Valid {
		mayor = c.MayorID.String
	}
	var term interface{}
	if c.TermEnd.Valid {
		term = c.TermEnd.Int64
	}
	var grace interface{}
	if c.GraceEnds.Valid {
		grace = c.GraceEnds.Int64
	}
	zoning := 0
	if c.ZoningOpen {
		zoning = 1
	}
	_, err := db.conn.Exec(
		`UPDATE cities SET rank = ?, treasury = ?, upkeep_weekly = ?,
			tax_flat_weekly = ?, tax_vendor_pct = ?, specialization = ?,
			mayor_character_id = ?, term_end = ?, status = ?, grace_ends = ?,
			zoning_open = ?
		WHERE id = ?`,
		c.Rank, c.Treasury, c.UpkeepWeekly, c.TaxFlatWeekly, c.TaxVendorPct,
		c.Specialization, mayor, term, c.Status, grace, zoning, c.ID)
	return err
}

// CountStructuresInCity counts non-destroyed structures-table rows inside a
// city radius. Harvesters are NOT counted: they bind to spawn geography, not
// civic geography (flagged scope reading of GDD 14.2.1's "any mix" —
// recorded, not silently dropped).
func (db *DB) CountStructuresInCity(zone string, cx, cz, radius float64) (int, error) {
	var n int
	err := db.conn.QueryRow(
		`SELECT COUNT(*) FROM structures
		WHERE zone = ? AND status != 'destroyed'
		AND ((pos_x - ?) * (pos_x - ?) + (pos_z - ?) * (pos_z - ?)) <= ? * ?`,
		zone, cx, cx, cz, cz, radius, radius).Scan(&n)
	return n, err
}

// StructuresInCity lists non-destroyed structures inside a city radius.
func (db *DB) StructuresInCity(zone string, cx, cz, radius float64) ([]StructureRow, error) {
	rows, err := db.conn.Query(
		`SELECT id, owner_character_id, zone, pos_x, pos_z, kind, tier,
			maintenance_pool, upkeep_weekly, status FROM structures
		WHERE zone = ? AND status != 'destroyed'
		AND ((pos_x - ?) * (pos_x - ?) + (pos_z - ?) * (pos_z - ?)) <= ? * ?`,
		zone, cx, cx, cz, cz, radius, radius)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var out []StructureRow
	for rows.Next() {
		s, err := scanStructure(rows)
		if err != nil {
			return nil, err
		}
		out = append(out, *s)
	}
	return out, rows.Err()
}

// CityContaining returns the first ACTIVE city containing a zone point, if any.
// Radii never overlap by placement validation, so at most one matches.
func (db *DB) CityContaining(zone string, x, z float64) (*CityRow, error) {
	rows, err := db.conn.Query(
		`SELECT `+cityCols+` FROM cities WHERE zone = ? AND status = 'active'`, zone)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var c CityRow
		var zoning int
		if err := rows.Scan(&c.ID, &c.Name, &c.Zone, &c.CenterX, &c.CenterZ,
			&c.RadiusM, &c.Rank, &c.Treasury, &c.UpkeepWeekly, &c.TaxFlatWeekly,
			&c.TaxVendorPct, &c.Specialization, &c.MayorID, &c.TermEnd,
			&c.Status, &c.GraceEnds, &zoning); err != nil {
			return nil, err
		}
		c.ZoningOpen = zoning != 0
		dx, dz := x-c.CenterX, z-c.CenterZ
		if dx*dx+dz*dz <= c.RadiusM*c.RadiusM {
			return &c, nil
		}
	}
	return nil, sql.ErrNoRows
}

// --- Citizenship ---

// CitizenRow maps one city_citizens row.
type CitizenRow struct {
	CityID     string
	CharacterID string
	TaxArrears int
	IsMilitia  bool
}

// AddCitizen enrolls a citizen (idempotent).
func (db *DB) AddCitizen(cityID, charID string) error {
	_, err := db.conn.Exec(
		`INSERT OR IGNORE INTO city_citizens (city_id, character_id) VALUES (?, ?)`,
		cityID, charID)
	return err
}

// RemoveCitizen unenrolls (opt-out / eviction).
func (db *DB) RemoveCitizen(cityID, charID string) error {
	_, err := db.conn.Exec(
		`DELETE FROM city_citizens WHERE city_id = ? AND character_id = ?`,
		cityID, charID)
	return err
}

// IsCitizen reports membership.
func (db *DB) IsCitizen(cityID, charID string) bool {
	var n int
	_ = db.conn.QueryRow(
		`SELECT COUNT(*) FROM city_citizens WHERE city_id = ? AND character_id = ?`,
		cityID, charID).Scan(&n)
	return n > 0
}

// CitizensOf lists a city's citizens.
func (db *DB) CitizensOf(cityID string) ([]CitizenRow, error) {
	rows, err := db.conn.Query(
		`SELECT city_id, character_id, tax_arrears, is_militia
		FROM city_citizens WHERE city_id = ? ORDER BY joined_at`, cityID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var out []CitizenRow
	for rows.Next() {
		var c CitizenRow
		var mil int
		if err := rows.Scan(&c.CityID, &c.CharacterID, &c.TaxArrears, &mil); err != nil {
			return nil, err
		}
		c.IsMilitia = mil != 0
		out = append(out, c)
	}
	return out, rows.Err()
}

// AddOptOut records an opt-out (leave/evict); auto-enrol skips these.
func (db *DB) AddOptOut(cityID, charID string) error {
	_, err := db.conn.Exec(
		`INSERT OR IGNORE INTO city_optouts (city_id, character_id) VALUES (?, ?)`,
		cityID, charID)
	return err
}

// IsOptedOut reports opt-out state.
func (db *DB) IsOptedOut(cityID, charID string) bool {
	var n int
	_ = db.conn.QueryRow(
		`SELECT COUNT(*) FROM city_optouts WHERE city_id = ? AND character_id = ?`,
		cityID, charID).Scan(&n)
	return n > 0
}

// AddArrears increments a citizen's flat-tax arrears (eviction evidence).
func (db *DB) AddArrears(cityID, charID string, amount int) error {
	_, err := db.conn.Exec(
		`UPDATE city_citizens SET tax_arrears = tax_arrears + ?
		WHERE city_id = ? AND character_id = ?`, amount, cityID, charID)
	return err
}

// --- Elections ---

// ElectionRow maps one city_elections row.
type ElectionRow struct {
	ID       string
	CityID   string
	EndsAt   int64
	Status   string
	WinnerID sql.NullString
}

// OpenElection starts an election; fails if one is already open.
func (db *DB) OpenElection(id, cityID string, endsAt int64) error {
	var n int
	if err := db.conn.QueryRow(
		`SELECT COUNT(*) FROM city_elections WHERE city_id = ? AND status = 'open'`,
		cityID).Scan(&n); err != nil {
		return err
	}
	if n > 0 {
		return fmt.Errorf("election already open")
	}
	_, err := db.conn.Exec(
		`INSERT INTO city_elections (id, city_id, ends_at, status)
		VALUES (?, ?, ?, 'open')`, id, cityID, endsAt)
	return err
}

// GetOpenElection returns the city's open election, if any.
func (db *DB) GetOpenElection(cityID string) (*ElectionRow, error) {
	var e ElectionRow
	err := db.conn.QueryRow(
		`SELECT id, city_id, ends_at, status, winner_id FROM city_elections
		WHERE city_id = ? AND status = 'open' LIMIT 1`, cityID,
	).Scan(&e.ID, &e.CityID, &e.EndsAt, &e.Status, &e.WinnerID)
	if err != nil {
		return nil, err
	}
	return &e, nil
}

// LastElectionEnd returns the latest closed-election close time (unix), or 0.
func (db *DB) LastElectionEnd(cityID string) int64 {
	var ends sql.NullInt64
	_ = db.conn.QueryRow(
		`SELECT MAX(ends_at) FROM city_elections
		WHERE city_id = ? AND status = 'closed'`, cityID).Scan(&ends)
	if !ends.Valid {
		return 0
	}
	return ends.Int64
}

// CastBallot records one vote (one-vote-per-citizen enforced by PK).
func (db *DB) CastBallot(electionID, voterID, candidateID string) error {
	_, err := db.conn.Exec(
		`INSERT INTO city_ballots (election_id, voter_id, candidate_id)
		VALUES (?, ?, ?)`, electionID, voterID, candidateID)
	return err
}

// BallotRow maps one ballot (cast_at ordering = tie-break source).
type BallotRow struct {
	VoterID     string
	CandidateID string
}

// BallotsOf returns ballots in cast order (rowid order = earliest first).
func (db *DB) BallotsOf(electionID string) ([]BallotRow, error) {
	rows, err := db.conn.Query(
		`SELECT voter_id, candidate_id FROM city_ballots
		WHERE election_id = ? ORDER BY rowid`, electionID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var out []BallotRow
	for rows.Next() {
		var b BallotRow
		if err := rows.Scan(&b.VoterID, &b.CandidateID); err != nil {
			return nil, err
		}
		out = append(out, b)
	}
	return out, rows.Err()
}

// CloseElection marks an election closed with its winner.
func (db *DB) CloseElection(electionID, winnerID string) error {
	_, err := db.conn.Exec(
		`UPDATE city_elections SET status = 'closed', winner_id = ? WHERE id = ?`,
		winnerID, electionID)
	return err
}

// --- Presence (mayor-inactivity evidence) ---

// TouchPresence records a character's world entry (upsert).
func (db *DB) TouchPresence(charID string, now int64) error {
	_, err := db.conn.Exec(
		`INSERT INTO civic_presence (character_id, last_seen) VALUES (?, ?)
		ON CONFLICT(character_id) DO UPDATE SET last_seen = ?`,
		charID, now, now)
	return err
}

// LastSeen returns the last world-entry time, or 0 if never.
func (db *DB) LastSeen(charID string) int64 {
	var t int64
	if err := db.conn.QueryRow(
		`SELECT last_seen FROM civic_presence WHERE character_id = ?`,
		charID).Scan(&t); err != nil {
		return 0
	}
	return t
}

// --- Treasury helpers (non-transactional; tx variants below for E2) ---

// FundCityTreasury moves wallet credits → city treasury (transfer, ledger-tagged
// by the caller convention: category city_fund).
func (db *DB) FundCityTreasury(cityID, charID string, amount int) error {
	tx, err := db.conn.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()
	var bal int
	if err := tx.QueryRow(
		`SELECT credits FROM characters WHERE id = ?`, charID).Scan(&bal); err != nil {
		return fmt.Errorf("funder unknown: %w", err)
	}
	if bal < amount {
		return fmt.Errorf("insufficient credits")
	}
	if _, err := tx.Exec(
		`UPDATE characters SET credits = credits - ? WHERE id = ?`, amount, charID); err != nil {
		return err
	}
	if _, err := tx.Exec(
		`UPDATE cities SET treasury = treasury + ? WHERE id = ?`, amount, cityID); err != nil {
		return err
	}
	if err := RecordLedgerTx(tx, charID, -amount, "city_fund", "transfer", "city:"+cityID, ""); err != nil {
		return err
	}
	if err := RecordLedgerTx(tx, "city:"+cityID, amount, "city_fund", "transfer", charID, ""); err != nil {
		return err
	}
	return tx.Commit()
}

// --- Guilds ---

// GuildRow maps one guilds row.
type GuildRow struct {
	ID              string
	Name            string
	Tag             string
	LeaderID        string
	Treasury        int
	DuesPct         int
	Faction         string
	HallStructureID sql.NullString
	DissolveAt      sql.NullInt64
}

// CreateGuild inserts a guild + founder-as-leader membership atomically,
// deducting the registrar sink from the founder.
func (db *DB) CreateGuild(id, name, tag, founderID string, registrarCost int) error {
	tx, err := db.conn.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()
	var bal int
	if err := tx.QueryRow(
		`SELECT credits FROM characters WHERE id = ?`, founderID).Scan(&bal); err != nil {
		return fmt.Errorf("founder unknown: %w", err)
	}
	if bal < registrarCost {
		return fmt.Errorf("insufficient credits for registrar fee")
	}
	if _, err := tx.Exec(
		`UPDATE characters SET credits = credits - ? WHERE id = ?`,
		registrarCost, founderID); err != nil {
		return err
	}
	if _, err := tx.Exec(
		`INSERT INTO guilds (id, name, tag, leader_id) VALUES (?, ?, ?, ?)`,
		id, name, tag, founderID); err != nil {
		return fmt.Errorf("guild name taken: %w", err)
	}
	if _, err := tx.Exec(
		`INSERT INTO guild_members (guild_id, character_id, role) VALUES (?, ?, 'leader')`,
		id, founderID); err != nil {
		return err
	}
	if err := RecordLedgerTx(tx, founderID, -registrarCost,
		"guild_registrar", "sink", "guild:"+id, ""); err != nil {
		return err
	}
	return tx.Commit()
}

// GetGuild loads one guild.
func (db *DB) GetGuild(id string) (*GuildRow, error) {
	var g GuildRow
	err := db.conn.QueryRow(
		`SELECT id, name, tag, leader_id, treasury, dues_pct, faction,
			hall_structure_id, dissolve_at FROM guilds WHERE id = ?`, id,
	).Scan(&g.ID, &g.Name, &g.Tag, &g.LeaderID, &g.Treasury, &g.DuesPct,
		&g.Faction, &g.HallStructureID, &g.DissolveAt)
	if err != nil {
		return nil, err
	}
	return &g, nil
}

// GuildOf returns the guild a character belongs to (one-guild-membership rule —
// flagged simplification; GDD is silent on multi-guild membership).
func (db *DB) GuildOf(charID string) (*GuildRow, error) {
	var g GuildRow
	err := db.conn.QueryRow(
		`SELECT g.id, g.name, g.tag, g.leader_id, g.treasury, g.dues_pct, g.faction,
			g.hall_structure_id, g.dissolve_at FROM guilds g
		JOIN guild_members m ON m.guild_id = g.id
		WHERE m.character_id = ? LIMIT 1`, charID,
	).Scan(&g.ID, &g.Name, &g.Tag, &g.LeaderID, &g.Treasury, &g.DuesPct,
		&g.Faction, &g.HallStructureID, &g.DissolveAt)
	if err != nil {
		return nil, err
	}
	return &g, nil
}

// GuildMemberRole returns a character's role in a guild.
func (db *DB) GuildMemberRole(guildID, charID string) (string, error) {
	var role string
	err := db.conn.QueryRow(
		`SELECT role FROM guild_members WHERE guild_id = ? AND character_id = ?`,
		guildID, charID).Scan(&role)
	return role, err
}

// GuildMembers lists members with roles (join order).
func (db *DB) GuildMembers(guildID string) ([]struct {
	CharacterID string
	Role        string
}, error) {
	rows, err := db.conn.Query(
		`SELECT character_id, role FROM guild_members WHERE guild_id = ?
		ORDER BY joined_at`, guildID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var out []struct {
		CharacterID string
		Role        string
	}
	for rows.Next() {
		var m struct {
			CharacterID string
			Role        string
		}
		if err := rows.Scan(&m.CharacterID, &m.Role); err != nil {
			return nil, err
		}
		out = append(out, m)
	}
	return out, rows.Err()
}

// AddGuildMember inserts a membership (caller enforces one-guild rule + role).
func (db *DB) AddGuildMember(guildID, charID, role string) error {
	_, err := db.conn.Exec(
		`INSERT INTO guild_members (guild_id, character_id, role) VALUES (?, ?, ?)`,
		guildID, charID, role)
	return err
}

// RemoveGuildMember deletes a membership.
func (db *DB) RemoveGuildMember(guildID, charID string) error {
	_, err := db.conn.Exec(
		`DELETE FROM guild_members WHERE guild_id = ? AND character_id = ?`,
		guildID, charID)
	return err
}

// SetGuildMemberRole changes a member's role.
func (db *DB) SetGuildMemberRole(guildID, charID, role string) error {
	_, err := db.conn.Exec(
		`UPDATE guild_members SET role = ? WHERE guild_id = ? AND character_id = ?`,
		role, guildID, charID)
	return err
}

// SetGuildLeader transfers leadership (role rows follow the ID).
func (db *DB) SetGuildLeader(guildID, newLeaderID string) error {
	tx, err := db.conn.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()
	if _, err := tx.Exec(
		`UPDATE guilds SET leader_id = ?, dissolve_at = NULL WHERE id = ?`,
		newLeaderID, guildID); err != nil {
		return err
	}
	if _, err := tx.Exec(
		`UPDATE guild_members SET role = 'member'
		WHERE guild_id = ? AND role = 'leader'`, guildID); err != nil {
		return err
	}
	if _, err := tx.Exec(
		`UPDATE guild_members SET role = 'leader'
		WHERE guild_id = ? AND character_id = ?`, guildID, newLeaderID); err != nil {
		return err
	}
	return tx.Commit()
}

// SetGuildDues sets the dues percentage (0–cap enforced by caller).
func (db *DB) SetGuildDues(guildID string, pct int) error {
	_, err := db.conn.Exec(`UPDATE guilds SET dues_pct = ? WHERE id = ?`, pct, guildID)
	return err
}

// DisbandGuild deletes a guild + memberships + pending invites.
func (db *DB) DisbandGuild(guildID string) error {
	tx, err := db.conn.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()
	for _, q := range []string{
		`DELETE FROM guild_invites WHERE guild_id = ?`,
		`DELETE FROM guild_members WHERE guild_id = ?`,
		`DELETE FROM guilds WHERE id = ?`,
	} {
		if _, err := tx.Exec(q, guildID); err != nil {
			return err
		}
	}
	return tx.Commit()
}

// GuildInviteRow maps one pending invite.
type GuildInviteRow struct {
	ID        string
	GuildID   string
	InviterID string
	InviteeID string
}

// CreateGuildInvite records a pending invite.
func (db *DB) CreateGuildInvite(id, guildID, inviterID, inviteeID string) error {
	_, err := db.conn.Exec(
		`INSERT INTO guild_invites (id, guild_id, inviter_id, invitee_id, status)
		VALUES (?, ?, ?, ?, 'pending')`, id, guildID, inviterID, inviteeID)
	return err
}

// PendingGuildInvites lists an invitee's pending invites.
func (db *DB) PendingGuildInvites(inviteeID string) ([]GuildInviteRow, error) {
	rows, err := db.conn.Query(
		`SELECT id, guild_id, inviter_id, invitee_id FROM guild_invites
		WHERE invitee_id = ? AND status = 'pending'`, inviteeID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var out []GuildInviteRow
	for rows.Next() {
		var iv GuildInviteRow
		if err := rows.Scan(&iv.ID, &iv.GuildID, &iv.InviterID, &iv.InviteeID); err != nil {
			return nil, err
		}
		out = append(out, iv)
	}
	return out, rows.Err()
}

// ResolveGuildInvite marks an invite accepted/declined.
func (db *DB) ResolveGuildInvite(id, status string) error {
	_, err := db.conn.Exec(
		`UPDATE guild_invites SET status = ? WHERE id = ? AND status = 'pending'`,
		status, id)
	return err
}

// GetGuildInvite loads one invite.
func (db *DB) GetGuildInvite(id string) (*GuildInviteRow, error) {
	var iv GuildInviteRow
	var status string
	err := db.conn.QueryRow(
		`SELECT id, guild_id, inviter_id, invitee_id, status FROM guild_invites
		WHERE id = ?`, id,
	).Scan(&iv.ID, &iv.GuildID, &iv.InviterID, &iv.InviteeID, &status)
	if err != nil {
		return nil, err
	}
	if status != "pending" {
		return nil, fmt.Errorf("invite no longer pending")
	}
	return &iv, nil
}

// FundGuildTreasury moves wallet credits → guild treasury (transfer).
func (db *DB) FundGuildTreasury(guildID, charID string, amount int) error {
	tx, err := db.conn.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()
	var bal int
	if err := tx.QueryRow(
		`SELECT credits FROM characters WHERE id = ?`, charID).Scan(&bal); err != nil {
		return fmt.Errorf("funder unknown: %w", err)
	}
	if bal < amount {
		return fmt.Errorf("insufficient credits")
	}
	if _, err := tx.Exec(
		`UPDATE characters SET credits = credits - ? WHERE id = ?`, amount, charID); err != nil {
		return err
	}
	if _, err := tx.Exec(
		`UPDATE guilds SET treasury = treasury + ? WHERE id = ?`, amount, guildID); err != nil {
		return err
	}
	if err := RecordLedgerTx(tx, charID, -amount, "guild_fund", "transfer", "guild:"+guildID, ""); err != nil {
		return err
	}
	if err := RecordLedgerTx(tx, "guild:"+guildID, amount, "guild_fund", "transfer", charID, ""); err != nil {
		return err
	}
	return tx.Commit()
}

// --- Groups ---

// GroupRow maps one groups row.
type GroupRow struct {
	ID       string
	LeaderID string
	LootRule string
}

// CreateGroup inserts a group + leader membership.
func (db *DB) CreateGroup(id, leaderID string) error {
	tx, err := db.conn.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()
	if _, err := tx.Exec(
		`INSERT INTO groups (id, leader_id, loot_rule) VALUES (?, ?, 'round_robin')`,
		id, leaderID); err != nil {
		return err
	}
	if _, err := tx.Exec(
		`INSERT INTO group_members (group_id, character_id) VALUES (?, ?)`,
		id, leaderID); err != nil {
		return err
	}
	return tx.Commit()
}

// GetGroup loads one group.
func (db *DB) GetGroup(id string) (*GroupRow, error) {
	var g GroupRow
	err := db.conn.QueryRow(
		`SELECT id, leader_id, loot_rule FROM groups WHERE id = ?`, id,
	).Scan(&g.ID, &g.LeaderID, &g.LootRule)
	if err != nil {
		return nil, err
	}
	return &g, nil
}

// GroupOf returns the group a character belongs to (one-group rule —
// flagged simplification; GDD is silent on multi-group membership).
func (db *DB) GroupOf(charID string) (*GroupRow, error) {
	var g GroupRow
	err := db.conn.QueryRow(
		`SELECT g.id, g.leader_id, g.loot_rule FROM groups g
		JOIN group_members m ON m.group_id = g.id
		WHERE m.character_id = ? LIMIT 1`, charID,
	).Scan(&g.ID, &g.LeaderID, &g.LootRule)
	if err != nil {
		return nil, err
	}
	return &g, nil
}

// GroupMembers lists members in join order (tenure source for succession).
func (db *DB) GroupMembers(groupID string) ([]string, error) {
	rows, err := db.conn.Query(
		`SELECT character_id FROM group_members WHERE group_id = ?
		ORDER BY joined_at`, groupID)
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

// AddGroupMember inserts a membership.
func (db *DB) AddGroupMember(groupID, charID string) error {
	_, err := db.conn.Exec(
		`INSERT INTO group_members (group_id, character_id) VALUES (?, ?)`,
		groupID, charID)
	return err
}

// RemoveGroupMember deletes a membership.
func (db *DB) RemoveGroupMember(groupID, charID string) error {
	_, err := db.conn.Exec(
		`DELETE FROM group_members WHERE group_id = ? AND character_id = ?`,
		groupID, charID)
	return err
}

// SetGroupLeader transfers leadership.
func (db *DB) SetGroupLeader(groupID, newLeaderID string) error {
	_, err := db.conn.Exec(
		`UPDATE groups SET leader_id = ? WHERE id = ?`, newLeaderID, groupID)
	return err
}

// SetLootRule changes the loot rule.
func (db *DB) SetLootRule(groupID, rule string) error {
	_, err := db.conn.Exec(
		`UPDATE groups SET loot_rule = ? WHERE id = ?`, rule, groupID)
	return err
}

// DisbandGroup deletes a group + memberships + pending invites.
func (db *DB) DisbandGroup(groupID string) error {
	tx, err := db.conn.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()
	for _, q := range []string{
		`DELETE FROM group_invites WHERE group_id = ?`,
		`DELETE FROM group_members WHERE group_id = ?`,
		`DELETE FROM groups WHERE id = ?`,
	} {
		if _, err := tx.Exec(q, groupID); err != nil {
			return err
		}
	}
	return tx.Commit()
}

// GroupInviteRow maps one pending group invite.
type GroupInviteRow struct {
	ID        string
	GroupID   string
	InviterID string
	InviteeID string
}

// CreateGroupInvite records a pending invite.
func (db *DB) CreateGroupInvite(id, groupID, inviterID, inviteeID string) error {
	_, err := db.conn.Exec(
		`INSERT INTO group_invites (id, group_id, inviter_id, invitee_id, status)
		VALUES (?, ?, ?, ?, 'pending')`, id, groupID, inviterID, inviteeID)
	return err
}

// PendingGroupInvites lists an invitee's pending invites.
func (db *DB) PendingGroupInvites(inviteeID string) ([]GroupInviteRow, error) {
	rows, err := db.conn.Query(
		`SELECT id, group_id, inviter_id, invitee_id FROM group_invites
		WHERE invitee_id = ? AND status = 'pending'`, inviteeID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var out []GroupInviteRow
	for rows.Next() {
		var iv GroupInviteRow
		if err := rows.Scan(&iv.ID, &iv.GroupID, &iv.InviterID, &iv.InviteeID); err != nil {
			return nil, err
		}
		out = append(out, iv)
	}
	return out, rows.Err()
}

// GetGroupInvite loads one pending invite.
func (db *DB) GetGroupInvite(id string) (*GroupInviteRow, error) {
	var iv GroupInviteRow
	var status string
	err := db.conn.QueryRow(
		`SELECT id, group_id, inviter_id, invitee_id, status FROM group_invites
		WHERE id = ?`, id,
	).Scan(&iv.ID, &iv.GroupID, &iv.InviterID, &iv.InviteeID, &status)
	if err != nil {
		return nil, err
	}
	if status != "pending" {
		return nil, fmt.Errorf("invite no longer pending")
	}
	return &iv, nil
}

// ResolveGroupInvite marks an invite accepted/declined.
func (db *DB) ResolveGroupInvite(id, status string) error {
	_, err := db.conn.Exec(
		`UPDATE group_invites SET status = ? WHERE id = ? AND status = 'pending'`,
		status, id)
	return err
}

// --- Mentorships ---

// AddBond records a mentorship bond.
func (db *DB) AddBond(mentorID, protegeID string, expiresAt int64) error {
	_, err := db.conn.Exec(
		`INSERT INTO mentorships (mentor_id, protege_id, expires_at)
		VALUES (?, ?, ?)`, mentorID, protegeID, expiresAt)
	return err
}

// BondRow maps one mentorship.
type BondRow struct {
	MentorID  string
	ProtegeID string
	ExpiresAt int64
}

// BondsOf returns bonds involving a character (either side).
func (db *DB) BondsOf(charID string) ([]BondRow, error) {
	rows, err := db.conn.Query(
		`SELECT mentor_id, protege_id, expires_at FROM mentorships
		WHERE mentor_id = ? OR protege_id = ?`, charID, charID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var out []BondRow
	for rows.Next() {
		var b BondRow
		if err := rows.Scan(&b.MentorID, &b.ProtegeID, &b.ExpiresAt); err != nil {
			return nil, err
		}
		out = append(out, b)
	}
	return out, rows.Err()
}

// CountSkillBoxes returns the number of owned skill boxes (threshold source).
func (db *DB) CountSkillBoxes(charID string) (int, error) {
	var n int
	err := db.conn.QueryRow(
		`SELECT COUNT(*) FROM character_skills WHERE character_id = ?`,
		charID).Scan(&n)
	return n, err
}

// OccupyTxTreasury applies a city-tax + guild-dues split inside the E2 purchase
// transaction: till keeps price − cityTax − dues; treasuries gain their shares;
// ledger transfer rows are written for every leg. Zero amounts are no-ops, so
// pre-civic behavior is byte-identical when no city/dues apply.
func OccupyTxTreasury(tx *sql.Tx, zone string, vendorX, vendorZ float64,
	sellerID string, price int, now time.Time) (till int, err error) {
	_ = now
	till = price
	// City vendor-% tax: vendor standing in an active taxed city. All active
	// zone cities are scanned (radii never overlap by placement validation, so
	// at most one contains the point).
	rows, err := tx.Query(
		`SELECT id, tax_vendor_pct, center_x, center_z, radius_m FROM cities
		WHERE zone = ? AND status = 'active'`, zone)
	if err == nil {
		type match struct {
			id     string
			taxPct int
		}
		var hits []match
		for rows.Next() {
			var m match
			var cx, cz, radius float64
			if err := rows.Scan(&m.id, &m.taxPct, &cx, &cz, &radius); err != nil {
				continue
			}
			dx, dz := vendorX-cx, vendorZ-cz
			if dx*dx+dz*dz <= radius*radius && m.taxPct > 0 {
				hits = append(hits, m)
			}
		}
		rows.Close()
		for _, m := range hits {
			tax := price * m.taxPct / 100
			if tax <= 0 {
				continue
			}
			till -= tax
			if _, err := tx.Exec(
				`UPDATE cities SET treasury = treasury + ? WHERE id = ?`,
				tax, m.id); err != nil {
				return price, err
			}
			if err := RecordLedgerTx(tx, sellerID, -tax,
				"city_vendor_tax", "transfer", "city:"+m.id, zone); err != nil {
				return price, err
			}
			if err := RecordLedgerTx(tx, "city:"+m.id, tax,
				"city_vendor_tax", "transfer", sellerID, zone); err != nil {
				return price, err
			}
		}
	}
	// Guild dues: seller's guild diverts dues% of the sale into its treasury.
	var guildID string
	var duesPct int
	if err := tx.QueryRow(
		`SELECT g.id, g.dues_pct FROM guilds g
		JOIN guild_members m ON m.guild_id = g.id
		WHERE m.character_id = ? LIMIT 1`, sellerID).Scan(&guildID, &duesPct); err == nil && duesPct > 0 {
		dues := price * duesPct / 100
		if dues > 0 {
			till -= dues
			if _, err := tx.Exec(
				`UPDATE guilds SET treasury = treasury + ? WHERE id = ?`,
				dues, guildID); err != nil {
				return price, err
			}
			if err := RecordLedgerTx(tx, sellerID, -dues,
				"guild_dues", "transfer", "guild:"+guildID, zone); err != nil {
				return price, err
			}
			if err := RecordLedgerTx(tx, "guild:"+guildID, dues,
				"guild_dues", "transfer", sellerID, zone); err != nil {
				return price, err
			}
		}
	}
	if till < 0 {
		till = 0
	}
	return till, nil
}

// CharacterName returns a character's display name (social UI co-display).
func (db *DB) CharacterName(charID string) string {
	var name string
	if err := db.conn.QueryRow(
		`SELECT name FROM characters WHERE id = ?`, charID).Scan(&name); err != nil {
		return ""
	}
	return name
}

// CharacterIDByName resolves a display name to a character ID.
func (db *DB) CharacterIDByName(name string) (string, error) {
	var id string
	err := db.conn.QueryRow(
		`SELECT id FROM characters WHERE name = ?`, name).Scan(&id)
	return id, err
}
