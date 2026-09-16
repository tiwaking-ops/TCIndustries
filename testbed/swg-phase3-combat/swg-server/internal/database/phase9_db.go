// Phase9DB extends the DB with Phase 9 persistence: mission terminals,
// missions, bounty contracts, pets, camps, and vendor-customer history, plus
// the lair damage log. DNA and tissue ride the crafted_items table as
// schematics (dna_sample / tissue_sample), reusing mail/vendor/claim
// machinery with no new tables (flagged reuse). Testbed fork; generic only.
//
// Mission sweep is lazy (no new tick): list/accept/turn-in expire old rows and
// refill to the listing target (flagged simplification — deterministic for
// tests). Camp expiry rides the harvest tick alongside structures.
package database

import (
	"database/sql"
	"fmt"
	"sync"
	"time"
)

// Row-ID minting: Windows wall clocks tick at ~15.6 ms granularity, so
// time.Now().UnixNano() alone is NOT unique within a process there — IDs
// minted in a tight loop (terminal refill, batch spawns) collided and failed
// their INSERTs. The counter disambiguates same-tick calls; callers keep the
// "<prefix>-<unixnano>" shape via the returned string.
var (
	rowIDMu   sync.Mutex
	rowIDLast int64
	rowIDSeq  uint64
)

// newRowID returns a unique "<prefix>-<unixnano>-<seq>" row ID.
func newRowID(prefix string) string {
	rowIDMu.Lock()
	defer rowIDMu.Unlock()
	ns := time.Now().UnixNano()
	if ns != rowIDLast {
		rowIDLast = ns
		rowIDSeq = 0
	}
	rowIDSeq++
	return fmt.Sprintf("%s-%d-%d", prefix, ns, rowIDSeq)
}

// NewRowID is the exported collision-proof row-ID mint (Phase 10 B6
// portability sweep): time.Now().UnixNano() alone is NOT unique on Windows
// (15.6 ms clock granularity), so any tight loop can mint duplicates.
// Format: "<prefix>-<unixnano>-<seq>". Mechanics unchanged; ID strings are
// unobservable to gameplay (proposal §3/B6).
func (db *DB) NewRowID(prefix string) string { return newRowID(prefix) }

// EnsurePhase9Schema creates Phase 9 tables if absent. Idempotent.
func (db *DB) EnsurePhase9Schema() error {
	stmts := []string{
		`CREATE TABLE IF NOT EXISTS mission_terminals (
			id TEXT PRIMARY KEY,
			zone TEXT NOT NULL,
			pos_x REAL NOT NULL,
			pos_z REAL NOT NULL,
			category TEXT NOT NULL
		)`,
		`CREATE TABLE IF NOT EXISTS missions (
			id TEXT PRIMARY KEY,
			terminal_id TEXT NOT NULL,
			type TEXT NOT NULL,
			target_ref TEXT NOT NULL DEFAULT '',
			schematic TEXT NOT NULL DEFAULT '',
			qty INTEGER NOT NULL DEFAULT 0,
			reward_credits INTEGER NOT NULL DEFAULT 0,
			reward_xp_type TEXT NOT NULL DEFAULT '',
			reward_xp INTEGER NOT NULL DEFAULT 0,
			status TEXT NOT NULL DEFAULT 'available',
			accepted_by TEXT NOT NULL DEFAULT '',
			posted_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
			expires_at INTEGER NOT NULL
		)`,
		`CREATE TABLE IF NOT EXISTS bounty_contracts (
			id TEXT PRIMARY KEY,
			poster_id TEXT NOT NULL,
			target_id TEXT NOT NULL,
			amount INTEGER NOT NULL,
			status TEXT NOT NULL DEFAULT 'open',
			created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
			expires_at INTEGER NOT NULL
		)`,
		`CREATE TABLE IF NOT EXISTS pets (
			id TEXT PRIMARY KEY,
			owner_character_id TEXT NOT NULL,
			template_id TEXT NOT NULL,
			name TEXT NOT NULL DEFAULT '',
			quality INTEGER NOT NULL DEFAULT 0,
			mode TEXT NOT NULL DEFAULT 'stay',
			target_id TEXT NOT NULL DEFAULT '',
			ability_flags TEXT NOT NULL DEFAULT '',
			zone TEXT NOT NULL DEFAULT '',
			pos_x REAL NOT NULL DEFAULT 0,
			pos_z REAL NOT NULL DEFAULT 0,
			created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
		)`,
		`CREATE TABLE IF NOT EXISTS camps (
			id TEXT PRIMARY KEY,
			deployed_by TEXT NOT NULL,
			zone TEXT NOT NULL,
			pos_x REAL NOT NULL,
			pos_z REAL NOT NULL,
			radius_m REAL NOT NULL,
			expires_at INTEGER NOT NULL,
			created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
		)`,
		`CREATE TABLE IF NOT EXISTS vendor_customers (
			vendor_id TEXT NOT NULL,
			character_id TEXT NOT NULL,
			first_bought_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
			PRIMARY KEY (vendor_id, character_id)
		)`,
		`CREATE TABLE IF NOT EXISTS lair_damage_log (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			lair_id TEXT NOT NULL,
			character_id TEXT NOT NULL,
			amount INTEGER NOT NULL,
			at INTEGER NOT NULL
		)`,
		`CREATE TABLE IF NOT EXISTS service_requests (
			id TEXT PRIMARY KEY,
			kind TEXT NOT NULL,
			requester_id TEXT NOT NULL,
			target_id TEXT NOT NULL,
			item_id TEXT NOT NULL DEFAULT '',
			detail TEXT NOT NULL DEFAULT '',
			fee INTEGER NOT NULL DEFAULT 0,
			status TEXT NOT NULL DEFAULT 'pending',
			created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
		)`,
		`CREATE TABLE IF NOT EXISTS holoemotes (
			id TEXT PRIMARY KEY,
			creator_id TEXT NOT NULL,
			name TEXT NOT NULL,
			text TEXT NOT NULL,
			price INTEGER NOT NULL DEFAULT 0,
			created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
		)`,
		`CREATE TABLE IF NOT EXISTS holo_rights (
			holo_id TEXT NOT NULL,
			character_id TEXT NOT NULL,
			PRIMARY KEY (holo_id, character_id)
		)`,
	}
	for i, s := range stmts {
		if _, err := db.conn.Exec(s); err != nil {
			return fmt.Errorf("phase9 schema %d failed: %w", i+1, err)
		}
	}
	// Seeded terminals at the zone spawn (bazaar-terminal precedent —
// single-zone MVP has no per-city terminals).
	for _, t := range []struct{ id, category string }{
		{"terminal-combat-01", "combat"},
		{"terminal-crafting-01", "crafting"},
		{"terminal-bounty-01", "bounty"},
	} {
		if _, err := db.conn.Exec(
			`INSERT OR IGNORE INTO mission_terminals (id, zone, pos_x, pos_z, category)
			VALUES (?, 'zone-0001', 20.0, 0.0, ?)`, t.id, t.category); err != nil {
			return fmt.Errorf("seed terminal %s failed: %w", t.id, err)
		}
	}
	return nil
}

// --- Mission terminals + missions ---

// MissionTerminal maps one mission terminal (bazaar terminals own the
// TerminalRow name — no collision by distinct naming).
type MissionTerminal struct {
	ID       string
	Zone     string
	X, Z     float64
	Category string
}

// TerminalsByCategory lists terminals of one category in a zone.
func (db *DB) TerminalsByCategory(zone, category string) ([]MissionTerminal, error) {
	rows, err := db.conn.Query(
		`SELECT id, zone, pos_x, pos_z, category FROM mission_terminals
		WHERE zone = ? AND category = ?`, zone, category)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var out []MissionTerminal
	for rows.Next() {
		var t MissionTerminal
		if err := rows.Scan(&t.ID, &t.Zone, &t.X, &t.Z, &t.Category); err != nil {
			return nil, err
		}
		out = append(out, t)
	}
	return out, rows.Err()
}

// MissionRow maps one missions row.
type MissionRow struct {
	ID           string
	TerminalID   string
	Type         string
	TargetRef    string
	Schematic    string
	Qty          int
	RewardCredit int
	RewardXPType string
	RewardXP     int
	Status       string
	AcceptedBy   string
	ExpiresAt    int64
}

// CreateMission inserts a mission row; returns its ID.
func (db *DB) CreateMission(m *MissionRow) (string, error) {
	id := newRowID("mission")
	_, err := db.conn.Exec(
		`INSERT INTO missions (id, terminal_id, type, target_ref, schematic, qty,
			reward_credits, reward_xp_type, reward_xp, status, accepted_by, expires_at)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, 'available', '', ?)`,
		id, m.TerminalID, m.Type, m.TargetRef, m.Schematic, m.Qty,
		m.RewardCredit, m.RewardXPType, m.RewardXP, m.ExpiresAt)
	if err != nil {
		return "", err
	}
	return id, nil
}

// GetMission loads one mission.
func (db *DB) GetMission(id string) (*MissionRow, error) {
	var m MissionRow
	err := db.conn.QueryRow(
		`SELECT id, terminal_id, type, target_ref, schematic, qty,
			reward_credits, reward_xp_type, reward_xp, status, accepted_by,
			expires_at FROM missions WHERE id = ?`, id,
	).Scan(&m.ID, &m.TerminalID, &m.Type, &m.TargetRef, &m.Schematic, &m.Qty,
		&m.RewardCredit, &m.RewardXPType, &m.RewardXP, &m.Status, &m.AcceptedBy,
		&m.ExpiresAt)
	if err != nil {
		return nil, err
	}
	return &m, nil
}

// AvailableMissions lists a terminal's available (unexpired-check by caller)
// missions.
func (db *DB) AvailableMissions(terminalID string) ([]MissionRow, error) {
	rows, err := db.conn.Query(
		`SELECT id, terminal_id, type, target_ref, schematic, qty,
			reward_credits, reward_xp_type, reward_xp, status, accepted_by,
			expires_at FROM missions
		WHERE terminal_id = ? AND status = 'available' ORDER BY rowid`, terminalID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var out []MissionRow
	for rows.Next() {
		var m MissionRow
		if err := rows.Scan(&m.ID, &m.TerminalID, &m.Type, &m.TargetRef,
			&m.Schematic, &m.Qty, &m.RewardCredit, &m.RewardXPType, &m.RewardXP,
			&m.Status, &m.AcceptedBy, &m.ExpiresAt); err != nil {
			return nil, err
		}
		out = append(out, m)
	}
	return out, rows.Err()
}

// MissionLog lists a character's accepted missions.
func (db *DB) MissionLog(charID string) ([]MissionRow, error) {
	rows, err := db.conn.Query(
		`SELECT id, terminal_id, type, target_ref, schematic, qty,
			reward_credits, reward_xp_type, reward_xp, status, accepted_by,
			expires_at FROM missions
		WHERE accepted_by = ? AND status = 'accepted' ORDER BY rowid`, charID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var out []MissionRow
	for rows.Next() {
		var m MissionRow
		if err := rows.Scan(&m.ID, &m.TerminalID, &m.Type, &m.TargetRef,
			&m.Schematic, &m.Qty, &m.RewardCredit, &m.RewardXPType, &m.RewardXP,
			&m.Status, &m.AcceptedBy, &m.ExpiresAt); err != nil {
			return nil, err
		}
		out = append(out, m)
	}
	return out, rows.Err()
}

// SetMissionStatus flips a mission (accepted/completed/expired/abandoned).
func (db *DB) SetMissionStatus(id, status, acceptedBy string) error {
	_, err := db.conn.Exec(
		`UPDATE missions SET status = ?, accepted_by = ? WHERE id = ?`,
		status, acceptedBy, id)
	return err
}

// ExpireMissions marks past-due available/accepted missions expired; returns
// the count (refund handling for contracts is separate).
func (db *DB) ExpireMissions(now int64) (int, error) {
	res, err := db.conn.Exec(
		`UPDATE missions SET status = 'expired'
		WHERE status IN ('available','accepted') AND expires_at <= ?`, now)
	if err != nil {
		return 0, err
	}
	n, _ := res.RowsAffected()
	return int(n), nil
}

// --- Bounty contracts (player-posted, escrowed) ---

// ContractRow maps one bounty contract.
type ContractRow struct {
	ID        string
	PosterID  string
	TargetID  string
	Amount    int
	Status    string
	ExpiresAt int64
}

// PostContract deducts the escrow upfront and opens the contract.
func (db *DB) PostContract(posterID, targetID string, amount int, expiresAt int64) (string, error) {
	tx, err := db.conn.Begin()
	if err != nil {
		return "", err
	}
	defer tx.Rollback()
	var bal int
	if err := tx.QueryRow(
		`SELECT credits FROM characters WHERE id = ?`, posterID).Scan(&bal); err != nil {
		return "", fmt.Errorf("poster unknown: %w", err)
	}
	if bal < amount {
		return "", fmt.Errorf("insufficient credits for escrow")
	}
	if _, err := tx.Exec(
		`UPDATE characters SET credits = credits - ? WHERE id = ?`,
		amount, posterID); err != nil {
		return "", err
	}
	id := newRowID("contract")
	if _, err := tx.Exec(
		`INSERT INTO bounty_contracts (id, poster_id, target_id, amount, status,
			expires_at) VALUES (?, ?, ?, ?, 'open', ?)`,
		id, posterID, targetID, amount, expiresAt); err != nil {
		return "", err
	}
	if err := RecordLedgerTx(tx, posterID, -amount,
		"bounty_escrow", "transfer", "contract:"+id, ""); err != nil {
		return "", err
	}
	if err := tx.Commit(); err != nil {
		return "", err
	}
	return id, nil
}

// GetContract loads one contract.
func (db *DB) GetContract(id string) (*ContractRow, error) {
	var c ContractRow
	err := db.conn.QueryRow(
		`SELECT id, poster_id, target_id, amount, status, expires_at
		FROM bounty_contracts WHERE id = ?`, id,
	).Scan(&c.ID, &c.PosterID, &c.TargetID, &c.Amount, &c.Status, &c.ExpiresAt)
	if err != nil {
		return nil, err
	}
	return &c, nil
}

// OpenContractsOn lists open contracts against a victim (stacked bounties all
// pay on a valid BH kill — flagged).
func (db *DB) OpenContractsOn(victimID string, now int64) ([]ContractRow, error) {
	rows, err := db.conn.Query(
		`SELECT id, poster_id, target_id, amount, status, expires_at
		FROM bounty_contracts
		WHERE target_id = ? AND status = 'open' AND expires_at > ?`, victimID, now)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var out []ContractRow
	for rows.Next() {
		var c ContractRow
		if err := rows.Scan(&c.ID, &c.PosterID, &c.TargetID, &c.Amount,
			&c.Status, &c.ExpiresAt); err != nil {
			return nil, err
		}
		out = append(out, c)
	}
	return out, rows.Err()
}

// OpenContracts lists all open contracts (BH-terminal visibility).
func (db *DB) OpenContracts(now int64) ([]ContractRow, error) {
	rows, err := db.conn.Query(
		`SELECT id, poster_id, target_id, amount, status, expires_at
		FROM bounty_contracts WHERE status = 'open' AND expires_at > ?
		ORDER BY rowid`, now)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var out []ContractRow
	for rows.Next() {
		var c ContractRow
		if err := rows.Scan(&c.ID, &c.PosterID, &c.TargetID, &c.Amount,
			&c.Status, &c.ExpiresAt); err != nil {
			return nil, err
		}
		out = append(out, c)
	}
	return out, rows.Err()
}

// CloseContract marks paid (escrow already left the poster at posting).
func (db *DB) CloseContract(id, status string) error {
	_, err := db.conn.Exec(
		`UPDATE bounty_contracts SET status = ? WHERE id = ? AND status = 'open'`,
		status, id)
	return err
}

// RefundContract returns escrow to the poster (expiry path).
func (db *DB) RefundContract(id string) error {
	tx, err := db.conn.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()
	var poster string
	var amount int
	if err := tx.QueryRow(
		`SELECT poster_id, amount FROM bounty_contracts
		WHERE id = ? AND status = 'open'`, id).Scan(&poster, &amount); err != nil {
		return err
	}
	if _, err := tx.Exec(
		`UPDATE characters SET credits = credits + ? WHERE id = ?`,
		amount, poster); err != nil {
		return err
	}
	if _, err := tx.Exec(
		`UPDATE bounty_contracts SET status = 'refunded' WHERE id = ?`, id); err != nil {
		return err
	}
	if err := RecordLedgerTx(tx, poster, amount,
		"bounty_refund", "transfer", "contract:"+id, ""); err != nil {
		return err
	}
	return tx.Commit()
}

// ExpiredContracts lists open contracts past expiry (sweep refunds them).
func (db *DB) ExpiredContracts(now int64) ([]ContractRow, error) {
	rows, err := db.conn.Query(
		`SELECT id, poster_id, target_id, amount, status, expires_at
		FROM bounty_contracts WHERE status = 'open' AND expires_at <= ?`, now)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var out []ContractRow
	for rows.Next() {
		var c ContractRow
		if err := rows.Scan(&c.ID, &c.PosterID, &c.TargetID, &c.Amount,
			&c.Status, &c.ExpiresAt); err != nil {
			return nil, err
		}
		out = append(out, c)
	}
	return out, rows.Err()
}

// --- Pets ---

// PetRow maps one pets row.
type PetRow struct {
	ID         string
	OwnerID    string
	TemplateID string
	Name       string
	Quality    int
	Mode       string // stay | follow | attack
	TargetID   string
	Abilities  string
	Zone       string
	PosX, PosZ float64
}

const petCols = `id, owner_character_id, template_id, name, quality, mode,
	target_id, ability_flags, zone, pos_x, pos_z`

func scanPetRow(row *sql.Rows, p *PetRow) error {
	return row.Scan(&p.ID, &p.OwnerID, &p.TemplateID, &p.Name, &p.Quality,
		&p.Mode, &p.TargetID, &p.Abilities, &p.Zone, &p.PosX, &p.PosZ)
}

// CreatePet inserts a pet; returns its ID.
func (db *DB) CreatePet(ownerID, templateID, name string, quality int, abilities, zone string, x, z float64) (string, error) {
	id := newRowID("pet")
	_, err := db.conn.Exec(
		`INSERT INTO pets (id, owner_character_id, template_id, name, quality,
			mode, abilities, zone, pos_x, pos_z)
		VALUES (?, ?, ?, ?, ?, 'stay', ?, ?, ?, ?)`,
		id, ownerID, templateID, name, quality, abilities, zone, x, z)
	if err != nil {
		return "", err
	}
	return id, nil
}

// GetPet loads one pet.
func (db *DB) GetPet(id string) (*PetRow, error) {
	var p PetRow
	err := db.conn.QueryRow(
		`SELECT `+petCols+` FROM pets WHERE id = ?`, id,
	).Scan(&p.ID, &p.OwnerID, &p.TemplateID, &p.Name, &p.Quality, &p.Mode,
		&p.TargetID, &p.Abilities, &p.Zone, &p.PosX, &p.PosZ)
	if err != nil {
		return nil, err
	}
	return &p, nil
}

// PetsOf lists an owner's pets.
func (db *DB) PetsOf(ownerID string) ([]PetRow, error) {
	rows, err := db.conn.Query(
		`SELECT `+petCols+` FROM pets WHERE owner_character_id = ?
		ORDER BY rowid`, ownerID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var out []PetRow
	for rows.Next() {
		var p PetRow
		if err := scanPetRow(rows, &p); err != nil {
			return nil, err
		}
		out = append(out, p)
	}
	return out, rows.Err()
}

// UpdatePetPos persists pet movement.
func (db *DB) UpdatePetPos(id, zone string, x, z float64) error {
	_, err := db.conn.Exec(
		`UPDATE pets SET zone = ?, pos_x = ?, pos_z = ? WHERE id = ?`,
		zone, x, z, id)
	return err
}

// SetPetCommand sets mode + target.
func (db *DB) SetPetCommand(id, mode, targetID string) error {
	_, err := db.conn.Exec(
		`UPDATE pets SET mode = ?, target_id = ? WHERE id = ?`,
		mode, targetID, id)
	return err
}

// DeletePet removes a pet (release).
func (db *DB) DeletePet(id string) error {
	_, err := db.conn.Exec(`DELETE FROM pets WHERE id = ?`, id)
	return err
}

// --- Camps ---

// CampRow maps one camps row.
type CampRow struct {
	ID        string
	DeployedBy string
	Zone      string
	PosX, PosZ float64
	RadiusM   float64
	ExpiresAt int64
}

// DeployCamp inserts a camp; returns its ID.
func (db *DB) DeployCamp(deployerID, zone string, x, z, radius float64, expiresAt int64) (string, error) {
	id := newRowID("camp")
	_, err := db.conn.Exec(
		`INSERT INTO camps (id, deployed_by, zone, pos_x, pos_z, radius_m,
			expires_at) VALUES (?, ?, ?, ?, ?, ?, ?)`,
		id, deployerID, zone, x, z, radius, expiresAt)
	if err != nil {
		return "", err
	}
	return id, nil
}

// ActiveCamps lists unexpired camps in a zone.
func (db *DB) ActiveCamps(zone string, now int64) ([]CampRow, error) {
	rows, err := db.conn.Query(
		`SELECT id, deployed_by, zone, pos_x, pos_z, radius_m, expires_at
		FROM camps WHERE zone = ? AND expires_at > ?`, zone, now)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var out []CampRow
	for rows.Next() {
		var c CampRow
		if err := rows.Scan(&c.ID, &c.DeployedBy, &c.Zone, &c.PosX, &c.PosZ,
			&c.RadiusM, &c.ExpiresAt); err != nil {
			return nil, err
		}
		out = append(out, c)
	}
	return out, rows.Err()
}

// DeleteCamp removes a camp (expiry sweep).
func (db *DB) DeleteCamp(id string) error {
	_, err := db.conn.Exec(`DELETE FROM camps WHERE id = ?`, id)
	return err
}

// ExpiredCamps lists camps past expiry (sweep deletes them).
func (db *DB) ExpiredCamps(now int64) ([]CampRow, error) {
	rows, err := db.conn.Query(
		`SELECT id, deployed_by, zone, pos_x, pos_z, radius_m, expires_at
		FROM camps WHERE expires_at <= ?`, now)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var out []CampRow
	for rows.Next() {
		var c CampRow
		if err := rows.Scan(&c.ID, &c.DeployedBy, &c.Zone, &c.PosX, &c.PosZ,
			&c.RadiusM, &c.ExpiresAt); err != nil {
			return nil, err
		}
		out = append(out, c)
	}
	return out, rows.Err()
}

// --- Service requests (Image Designer + Smuggler handshakes, invite pattern) ---

// ServiceRequest maps one service_requests row (kind: imagedesign | slice).
type ServiceRequest struct {
	ID          string
	Kind        string
	RequesterID string
	TargetID    string
	ItemID      string
	Detail      string
	Fee         int
}

// CreateServiceRequest inserts a pending request; returns its ID.
func (db *DB) CreateServiceRequest(kind, requesterID, targetID, itemID, detail string, fee int) (string, error) {
	id := newRowID("svcreq")
	_, err := db.conn.Exec(
		`INSERT INTO service_requests (id, kind, requester_id, target_id,
			item_id, detail, fee, status)
		VALUES (?, ?, ?, ?, ?, ?, ?, 'pending')`,
		id, kind, requesterID, targetID, itemID, detail, fee)
	if err != nil {
		return "", err
	}
	return id, nil
}

// GetServiceRequest loads one pending request.
func (db *DB) GetServiceRequest(id string) (*ServiceRequest, error) {
	var r ServiceRequest
	var status string
	err := db.conn.QueryRow(
		`SELECT id, kind, requester_id, target_id, item_id, detail, fee, status
		FROM service_requests WHERE id = ?`, id,
	).Scan(&r.ID, &r.Kind, &r.RequesterID, &r.TargetID, &r.ItemID, &r.Detail,
		&r.Fee, &status)
	if err != nil {
		return nil, err
	}
	if status != "pending" {
		return nil, fmt.Errorf("request no longer pending")
	}
	return &r, nil
}

// ResolveServiceRequest marks accepted/declined.
func (db *DB) ResolveServiceRequest(id, status string) error {
	_, err := db.conn.Exec(
		`UPDATE service_requests SET status = ? WHERE id = ? AND status = 'pending'`,
		status, id)
	return err
}

// --- Holoemotes (custom performer vocabulary, invite-pattern sale) ---

// CreateHoloemote inserts a custom + grants the creator rights.
func (db *DB) CreateHoloemote(creatorID, name, text string, price int) (string, error) {
	tx, err := db.conn.Begin()
	if err != nil {
		return "", err
	}
	defer tx.Rollback()
	id := newRowID("holo")
	if _, err := tx.Exec(
		`INSERT INTO holoemotes (id, creator_id, name, text, price)
		VALUES (?, ?, ?, ?, ?)`, id, creatorID, name, text, price); err != nil {
		return "", err
	}
	if _, err := tx.Exec(
		`INSERT INTO holo_rights (holo_id, character_id) VALUES (?, ?)`,
		id, creatorID); err != nil {
		return "", err
	}
	return id, tx.Commit()
}

// HoloemoteRow maps one custom.
type HoloemoteRow struct {
	ID        string
	CreatorID string
	Name      string
	Text      string
	Price     int
}

// HoloByName finds a custom by name (performance lookup).
func (db *DB) HoloByName(name string) (*HoloemoteRow, error) {
	var h HoloemoteRow
	err := db.conn.QueryRow(
		`SELECT id, creator_id, name, text, price FROM holoemotes WHERE name = ?`,
		name).Scan(&h.ID, &h.CreatorID, &h.Name, &h.Text, &h.Price)
	if err != nil {
		return nil, err
	}
	return &h, nil
}

// HasHoloRight reports perform rights.
func (db *DB) HasHoloRight(holoID, charID string) bool {
	var n int
	_ = db.conn.QueryRow(
		`SELECT COUNT(*) FROM holo_rights WHERE holo_id = ? AND character_id = ?`,
		holoID, charID).Scan(&n)
	return n > 0
}

// AllHoloemotes lists customs for browsing.
func (db *DB) AllHoloemotes() ([]HoloemoteRow, error) {
	rows, err := db.conn.Query(
		`SELECT id, creator_id, name, text, price FROM holoemotes ORDER BY rowid`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var out []HoloemoteRow
	for rows.Next() {
		var h HoloemoteRow
		if err := rows.Scan(&h.ID, &h.CreatorID, &h.Name, &h.Text, &h.Price); err != nil {
			return nil, err
		}
		out = append(out, h)
	}
	return out, rows.Err()
}

// BuyHoloemote transfers price to the creator and grants rights (atomic).
func (db *DB) BuyHoloemote(holoID, buyerID string) error {
	tx, err := db.conn.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()
	var creator string
	var price int
	if err := tx.QueryRow(
		`SELECT creator_id, price FROM holoemotes WHERE id = ?`,
		holoID).Scan(&creator, &price); err != nil {
		return fmt.Errorf("custom unknown: %w", err)
	}
	if creator == buyerID {
		return fmt.Errorf("already held")
	}
	var exists int
	_ = tx.QueryRow(
		`SELECT COUNT(*) FROM holo_rights WHERE holo_id = ? AND character_id = ?`,
		holoID, buyerID).Scan(&exists)
	if exists > 0 {
		return fmt.Errorf("already held")
	}
	if price > 0 {
		var bal int
		if err := tx.QueryRow(
			`SELECT credits FROM characters WHERE id = ?`, buyerID).Scan(&bal); err != nil {
			return err
		}
		if bal < price {
			return fmt.Errorf("insufficient credits")
		}
		if _, err := tx.Exec(
			`UPDATE characters SET credits = credits - ? WHERE id = ?`,
			price, buyerID); err != nil {
			return err
		}
		if _, err := tx.Exec(
			`UPDATE characters SET credits = credits + ? WHERE id = ?`,
			price, creator); err != nil {
			return err
		}
		if err := RecordLedgerTx(tx, buyerID, -price,
			"holoemote_sale", "transfer", creator, ""); err != nil {
			return err
		}
		if err := RecordLedgerTx(tx, creator, price,
			"holoemote_sale", "transfer", buyerID, ""); err != nil {
			return err
		}
	}
	if _, err := tx.Exec(
		`INSERT INTO holo_rights (holo_id, character_id) VALUES (?, ?)`,
		holoID, buyerID); err != nil {
		return err
	}
	return tx.Commit()
}

// UpdateAppearance applies Image Designer field edits (per-field tree mapping
// enforced by the caller).
func (db *DB) UpdateAppearance(charID string, fields map[string]interface{}) error {
	cols := map[string]string{
		"body_type": "body_type", "skin_color": "skin_color",
		"hair_style": "hair_style", "hair_color": "hair_color",
		"face_type": "face_type", "eye_color": "eye_color",
	}
	set := ""
	var args []interface{}
	for k, v := range fields {
		col, ok := cols[k]
		if !ok {
			continue
		}
		if set != "" {
			set += ", "
		}
		set += col + " = ?"
		args = append(args, v)
	}
	if set == "" {
		return fmt.Errorf("no editable fields")
	}
	if h, ok := fields["height"].(float64); ok {
		if h < 0.5 {
			h = 0.5
		}
		if h > 2.0 {
			h = 2.0
		}
		set += ", height = ?"
		args = append(args, h)
	}
	args = append(args, charID)
	_, err := db.conn.Exec(
		`UPDATE characters SET `+set+` WHERE id = ?`, args...)
	return err
}

// --- Craft counts (GDD 7.2.2 repeat diminishing: 50% XP after 10th craft) ---

// BumpCraftCount increments a character's per-schematic craft count and
// returns the new total (this craft included).
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

// --- Vendor customers (Merchant unique-buyer diminishing) ---

// RecordCustomer inserts an idempotent first-purchase row; returns true when
// this buyer is new to the vendor (full XP) vs repeat (diminished).
func (db *DB) RecordCustomer(vendorID, charID string) (bool, error) {
	res, err := db.conn.Exec(
		`INSERT OR IGNORE INTO vendor_customers (vendor_id, character_id)
		VALUES (?, ?)`, vendorID, charID)
	if err != nil {
		return false, err
	}
	n, _ := res.RowsAffected()
	return n == 1, nil
}

// --- Lair damage log (mission attribution, base_damage_log pattern) ---

// LogLairDamage records lair damage for attribution.
func (db *DB) LogLairDamage(lairID, charID string, amount int, now int64) error {
	_, err := db.conn.Exec(
		`INSERT INTO lair_damage_log (lair_id, character_id, amount, at)
		VALUES (?, ?, ?, ?)`, lairID, charID, amount, now)
	return err
}

// LairAttackers returns distinct attackers since a unix timestamp.
func (db *DB) LairAttackers(lairID string, since int64) ([]string, error) {
	rows, err := db.conn.Query(
		`SELECT DISTINCT character_id FROM lair_damage_log
		WHERE lair_id = ? AND at >= ?`, lairID, since)
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
