// EconomyDB extends the DB with Phase 5 persistence: ledger, structures,
// vendor listings, snapshots, and bazaar terminals. Testbed fork; generic only.
package database

import (
	"database/sql"
	"fmt"
	"time"
)

// EnsureEconomySchema creates Phase 5 tables if absent. Called at server startup.
func (db *DB) EnsureEconomySchema() error {
	stmts := []string{
		`CREATE TABLE IF NOT EXISTS ledger_entries (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			character_id TEXT NOT NULL,
			amount INTEGER NOT NULL,
			category TEXT NOT NULL,
			flow TEXT NOT NULL,
			counterparty TEXT NOT NULL DEFAULT '',
			zone TEXT NOT NULL DEFAULT '',
			created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
		)`,
		`CREATE TABLE IF NOT EXISTS structures (
			id TEXT PRIMARY KEY,
			owner_character_id TEXT NOT NULL,
			zone TEXT NOT NULL,
			pos_x REAL NOT NULL,
			pos_z REAL NOT NULL,
			kind TEXT NOT NULL,
			tier TEXT NOT NULL DEFAULT '',
			maintenance_pool INTEGER NOT NULL DEFAULT 0,
			upkeep_weekly INTEGER NOT NULL DEFAULT 0,
			status TEXT NOT NULL DEFAULT 'active',
			lapsed_at DATETIME,
			created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
		)`,
		`CREATE TABLE IF NOT EXISTS vendor_listings (
			id TEXT PRIMARY KEY,
			vendor_id TEXT NOT NULL,
			item_id TEXT NOT NULL,
			price INTEGER NOT NULL,
			description TEXT NOT NULL DEFAULT '',
			active INTEGER NOT NULL DEFAULT 1,
			listed_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
		)`,
		`CREATE TABLE IF NOT EXISTS economic_snapshots (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			taken_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
			circulation BIGINT NOT NULL,
			faucet_30d BIGINT NOT NULL,
			sink_30d BIGINT NOT NULL,
			net_pct REAL NOT NULL
		)`,
		`CREATE TABLE IF NOT EXISTS bazaar_terminals (
			id TEXT PRIMARY KEY,
			zone TEXT NOT NULL,
			pos_x REAL NOT NULL,
			pos_z REAL NOT NULL
		)`,
	}
	for i, s := range stmts {
		if _, err := db.conn.Exec(s); err != nil {
			return fmt.Errorf("economy schema %d failed: %w", i+1, err)
		}
	}
	// Seeded terminal at the zone spawn (documented stand-in for GDD 22.2.1's
	// "terminals in every city" — single-zone MVP has no cities).
	if _, err := db.conn.Exec(
		`INSERT OR IGNORE INTO bazaar_terminals (id, zone, pos_x, pos_z)
		VALUES ('terminal-spawn-01', 'zone-0001', 20.0, 0.0)`); err != nil {
		return fmt.Errorf("seed terminal failed: %w", err)
	}
	return nil
}

// RecordLedger appends one ledger entry. Best-effort companion to the credit
// mutation it describes (the E2 purchase primitive records inside its own tx).
func (db *DB) RecordLedger(characterID string, amount int, category, flow, counterparty, zone string) error {
	_, err := db.conn.Exec(
		`INSERT INTO ledger_entries
			(character_id, amount, category, flow, counterparty, zone)
		VALUES (?, ?, ?, ?, ?, ?)`,
		characterID, amount, category, flow, counterparty, zone,
	)
	return err
}

// RecordLedgerTx is the transaction-scoped variant for the E2 primitive.
func RecordLedgerTx(tx *sql.Tx, characterID string, amount int, category, flow, counterparty, zone string) error {
	_, err := tx.Exec(
		`INSERT INTO ledger_entries
			(character_id, amount, category, flow, counterparty, zone)
		VALUES (?, ?, ?, ?, ?, ?)`,
		characterID, amount, category, flow, counterparty, zone,
	)
	return err
}

// LedgerEntry maps one ledger row.
type LedgerEntry struct {
	ID          int64
	Amount      int
	Category    string
	Flow        string
	Counterparty string
	Zone        string
}

// LedgerFor returns a character's ledger entries (newest last).
func (db *DB) LedgerFor(characterID string) ([]LedgerEntry, error) {
	rows, err := db.conn.Query(
		`SELECT id, amount, category, flow, counterparty, zone
		FROM ledger_entries WHERE character_id = ? ORDER BY id`, characterID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var out []LedgerEntry
	for rows.Next() {
		var e LedgerEntry
		if err := rows.Scan(&e.ID, &e.Amount, &e.Category, &e.Flow, &e.Counterparty, &e.Zone); err != nil {
			return nil, err
		}
		out = append(out, e)
	}
	return out, rows.Err()
}

// --- Structures ---

// StructureRow maps one structures row.
type StructureRow struct {
	ID               string
	OwnerCharacterID string
	Zone             string
	PosX, PosZ       float64
	Kind             string // "house" | "vendor"
	Tier             string
	MaintenancePool  int
	UpkeepWeekly     int
	Status           string
}

func scanStructure(rows *sql.Rows) (*StructureRow, error) {
	var s StructureRow
	if err := rows.Scan(&s.ID, &s.OwnerCharacterID, &s.Zone, &s.PosX, &s.PosZ,
		&s.Kind, &s.Tier, &s.MaintenancePool, &s.UpkeepWeekly, &s.Status); err != nil {
		return nil, err
	}
	return &s, nil
}

// PlaceStructure records a structure deed placement.
func (db *DB) PlaceStructure(id, owner, zone string, x, z float64, kind, tier string, upkeep int) error {
	_, err := db.conn.Exec(
		`INSERT INTO structures
			(id, owner_character_id, zone, pos_x, pos_z, kind, tier, upkeep_weekly, status)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, 'active')`,
		id, owner, zone, x, z, kind, tier, upkeep,
	)
	return err
}

// GetStructure loads one structure.
func (db *DB) GetStructure(id string) (*StructureRow, error) {
	rows, err := db.conn.Query(
		`SELECT id, owner_character_id, zone, pos_x, pos_z, kind, tier,
			maintenance_pool, upkeep_weekly, status FROM structures WHERE id = ?`, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if !rows.Next() {
		return nil, sql.ErrNoRows
	}
	return scanStructure(rows)
}

// AllStructures returns every structure (tick + collision driver).
func (db *DB) AllStructures() ([]StructureRow, error) {
	rows, err := db.conn.Query(
		`SELECT id, owner_character_id, zone, pos_x, pos_z, kind, tier,
			maintenance_pool, upkeep_weekly, status FROM structures`)
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

// UpdateStructure persists pool/status changes.
func (db *DB) UpdateStructure(s *StructureRow) error {
	_, err := db.conn.Exec(
		`UPDATE structures SET maintenance_pool = ?, status = ?,
			lapsed_at = CASE WHEN ? IN ('closed','condemned') AND lapsed_at IS NULL
				THEN CURRENT_TIMESTAMP ELSE lapsed_at END
		WHERE id = ?`,
		s.MaintenancePool, s.Status, s.Status, s.ID)
	return err
}

// FundStructure deposits wallet credits into a structure pool (transfer, not faucet).
func (db *DB) FundStructure(id string, credits int) error {
	_, err := db.conn.Exec(
		`UPDATE structures SET maintenance_pool = maintenance_pool + ? WHERE id = ?`,
		credits, id)
	return err
}

// StructureLapsedAt returns when a structure entered its lapsed state.
func (db *DB) StructureLapsedAt(id string) (time.Time, error) {
	var lapsed sql.NullString
	if err := db.conn.QueryRow(
		`SELECT lapsed_at FROM structures WHERE id = ?`, id).Scan(&lapsed); err != nil {
		return time.Time{}, err
	}
	if !lapsed.Valid {
		return time.Time{}, fmt.Errorf("never lapsed")
	}
	for _, layout := range []string{
		"2006-01-02 15:04:05", "2006-01-02 15:04:05.999999999-07:00",
		"2006-01-02T15:04:05Z07:00", time.RFC3339,
	} {
		if t, err := time.Parse(layout, lapsed.String); err == nil {
			return t, nil
		}
	}
	return time.Time{}, fmt.Errorf("unparseable lapsed_at")
}

// --- Vendor listings ---

// ListingRow maps one vendor_listings row.
type ListingRow struct {
	ID          string
	VendorID    string
	ItemID      string
	Price       int
	Description string
	Active      bool
}

// CreateListing inserts a vendor listing; returns its ID.
func (db *DB) CreateListing(vendorID, itemID string, price int, desc string) (string, error) {
	id := newRowID("listing")
	_, err := db.conn.Exec(
		`INSERT INTO vendor_listings (id, vendor_id, item_id, price, description, active)
		VALUES (?, ?, ?, ?, ?, 1)`, id, vendorID, itemID, price, desc)
	if err != nil {
		return "", err
	}
	return id, nil
}

// GetListing loads one listing.
func (db *DB) GetListing(id string) (*ListingRow, error) {
	var l ListingRow
	var active int
	err := db.conn.QueryRow(
		`SELECT id, vendor_id, item_id, price, description, active
		FROM vendor_listings WHERE id = ?`, id,
	).Scan(&l.ID, &l.VendorID, &l.ItemID, &l.Price, &l.Description, &active)
	if err != nil {
		return nil, err
	}
	l.Active = active != 0
	return &l, nil
}

// SetListingActive flips a listing (sold / pulled).
func (db *DB) SetListingActive(id string, active bool) error {
	v := 0
	if active {
		v = 1
	}
	_, err := db.conn.Exec(`UPDATE vendor_listings SET active = ? WHERE id = ?`, v, id)
	return err
}

// RepriceListing changes a live listing's price.
func (db *DB) RepriceListing(id string, price int) error {
	_, err := db.conn.Exec(`UPDATE vendor_listings SET price = ? WHERE id = ?`, price, id)
	return err
}

// MarketRecord maps one market_records row (GDD 22.3 shape, MVP fields).
type MarketRecord struct {
	ID        int64
	ListingID string
	VendorID  string
	ItemID    string
	SellerID  string
	BuyerID   string
	Price     int
	Zone      string
}

// RecentSales returns the last-N sales for one schematic (MVP market data).
func (db *DB) RecentSales(schematic string, n int) ([]MarketRecord, error) {
	rows, err := db.conn.Query(
		`SELECT m.id, m.listing_id, m.vendor_id, m.item_id, m.seller_id, m.buyer_id,
			m.price, m.zone
		FROM market_records m
		JOIN crafted_items i ON i.id = m.item_id
		WHERE i.schematic_id = ? ORDER BY m.id DESC LIMIT ?`, schematic, n)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var out []MarketRecord
	for rows.Next() {
		var r MarketRecord
		if err := rows.Scan(&r.ID, &r.ListingID, &r.VendorID, &r.ItemID,
			&r.SellerID, &r.BuyerID, &r.Price, &r.Zone); err != nil {
			return nil, err
		}
		out = append(out, r)
	}
	return out, rows.Err()
}

// ActiveListings returns live listings, optionally filtered by vendor.
func (db *DB) ActiveListings(vendorID string) ([]ListingRow, error) {
	q := `SELECT id, vendor_id, item_id, price, description, active
		FROM vendor_listings WHERE active = 1`
	var rows *sql.Rows
	var err error
	if vendorID == "" {
		rows, err = db.conn.Query(q)
	} else {
		rows, err = db.conn.Query(q+` AND vendor_id = ?`, vendorID)
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var out []ListingRow
	for rows.Next() {
		var l ListingRow
		var active int
		if err := rows.Scan(&l.ID, &l.VendorID, &l.ItemID, &l.Price,
			&l.Description, &active); err != nil {
			return nil, err
		}
		l.Active = active != 0
		out = append(out, l)
	}
	return out, rows.Err()
}

// --- Snapshots + terminals ---

// SnapshotRow maps one economic_snapshots row.
type SnapshotRow struct {
	ID          int64
	Circulation int64
	Faucet30d   int64
	Sink30d     int64
	NetPct      float64
}

// WriteSnapshot aggregates circulation + 30-day faucet/sink sums and stores one
// snapshot row (GDD 12.4/12.6 shape, literal 30-day window semantics).
func (db *DB) WriteSnapshot() (*SnapshotRow, error) {
	var circ sql.NullInt64
	if err := db.conn.QueryRow(`SELECT SUM(credits) FROM characters`).Scan(&circ); err != nil {
		return nil, err
	}
	var faucet, sink sql.NullInt64
	if err := db.conn.QueryRow(
		`SELECT COALESCE(SUM(amount),0) FROM ledger_entries
		WHERE flow = 'faucet' AND created_at >= datetime('now','-30 days')`).Scan(&faucet); err != nil {
		return nil, err
	}
	if err := db.conn.QueryRow(
		`SELECT COALESCE(SUM(-amount),0) FROM ledger_entries
		WHERE flow = 'sink' AND created_at >= datetime('now','-30 days')`).Scan(&sink); err != nil {
		return nil, err
	}
	net := 0.0
	if circ.Valid && circ.Int64 > 0 {
		net = float64(faucet.Int64-sink.Int64) / float64(circ.Int64) * 100.0
	}
	res, err := db.conn.Exec(
		`INSERT INTO economic_snapshots (circulation, faucet_30d, sink_30d, net_pct)
		VALUES (?, ?, ?, ?)`, circ.Int64, faucet.Int64, sink.Int64, net)
	if err != nil {
		return nil, err
	}
	id, _ := res.LastInsertId()
	return &SnapshotRow{ID: id, Circulation: circ.Int64,
		Faucet30d: faucet.Int64, Sink30d: sink.Int64, NetPct: net}, nil
}

// LatestSnapshot returns the newest snapshot row, if any.
func (db *DB) LatestSnapshot() (*SnapshotRow, error) {
	var s SnapshotRow
	var taken string
	err := db.conn.QueryRow(
		`SELECT id, circulation, faucet_30d, sink_30d, net_pct, taken_at
		FROM economic_snapshots ORDER BY id DESC LIMIT 1`,
	).Scan(&s.ID, &s.Circulation, &s.Faucet30d, &s.Sink30d, &s.NetPct, &taken)
	if err != nil {
		return nil, err
	}
	return &s, nil
}

// AddCredits adds credits to a wallet (faucet-side primitive; every call site
// must RecordLedger the matching faucet entry).
func (db *DB) AddCredits(characterID string, amount int) error {
	_, err := db.conn.Exec(
		`UPDATE characters SET credits = credits + ? WHERE id = ?`, amount, characterID)
	return err
}

// TerminalRow maps one bazaar terminal.
type TerminalRow struct {
	ID         string
	Zone       string
	X, Z       float64
}

// Terminals returns all bazaar terminals in a zone.
func (db *DB) Terminals(zone string) ([]TerminalRow, error) {
	rows, err := db.conn.Query(
		`SELECT id, zone, pos_x, pos_z FROM bazaar_terminals WHERE zone = ?`, zone)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var out []TerminalRow
	for rows.Next() {
		var t TerminalRow
		if err := rows.Scan(&t.ID, &t.Zone, &t.X, &t.Z); err != nil {
			return nil, err
		}
		out = append(out, t)
	}
	return out, rows.Err()
}

// --- E2 atomic purchase (the one way value moves) ---

// AtomicPurchase executes a vendor sale inside ONE database transaction
// (§§29.5/12.8/21.5 mandate): BEGIN IMMEDIATE serializes concurrent buyers, so
// the double-purchase race resolves to exactly one winner with no partial states.
// Caller supplies validated listing + vendor + buyer IDs; this function rechecks
// everything inside the tx (listing live, vendor open, buyer solvent, item present)
// and applies: buyer wallet −= price; vendor till += price (structures has no till
// column — till is tracked as maintenance_pool-neutral vendor credit via a
// vendor_tills companion: see below); item ownership → buyer; listing inactive;
// ledger transfer pair + market record are written by the caller-provided hook?
//
// Till design (recorded deviation, GDD-conformant in behavior): GDD 21.3 models
// till_credits on the Vendor entity. This build stores the till in the vendor
// structure's maintenance_pool-adjacent column vendor_till (added below) so till
// and pool never commingle — maintenance deductions must not eat sales revenue.
func (db *DB) EnsureVendorTillColumn() error {
	_, err := db.conn.Exec(`ALTER TABLE structures ADD COLUMN vendor_till INTEGER NOT NULL DEFAULT 0`)
	if err != nil && !isDupColumn(err.Error()) {
		return err
	}
	return nil
}

func isDupColumn(s string) bool {
	for i := 0; i+16 <= len(s); i++ {
		if s[i:i+16] == "duplicate column" {
			return true
		}
	}
	return false
}

// PurchaseResult reports an atomic purchase outcome.
type PurchaseResult struct {
	BuyerCharged bool
	TillCredited int
}

// AtomicPurchase runs the E2 primitive. On any validation failure it returns an
// error and applies NOTHING (rollback).
func (db *DB) AtomicPurchase(buyerID, listingID, zone string) (*PurchaseResult, error) {
	tx, err := db.conn.Begin()
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	var price, active int
	var itemID, vendorID string
	if err := tx.QueryRow(
		`SELECT price, active, item_id, vendor_id FROM vendor_listings WHERE id = ?`,
		listingID).Scan(&price, &active, &itemID, &vendorID); err != nil {
		return nil, fmt.Errorf("listing unknown: %w", err)
	}
	if active == 0 {
		return nil, fmt.Errorf("sold out")
	}
	var vStatus, vOwner, vZone string
	var vX, vZ float64
	if err := tx.QueryRow(
		`SELECT status, owner_character_id, zone, pos_x, pos_z FROM structures WHERE id = ?`, vendorID,
	).Scan(&vStatus, &vOwner, &vZone, &vX, &vZ); err != nil {
		return nil, fmt.Errorf("vendor unknown: %w", err)
	}
	if vStatus != "active" {
		return nil, fmt.Errorf("vendor closed")
	}
	if vOwner == buyerID {
		return nil, fmt.Errorf("owner cannot buy own listing")
	}
	var balance int
	if err := tx.QueryRow(
		`SELECT credits FROM characters WHERE id = ?`, buyerID).Scan(&balance); err != nil {
		return nil, fmt.Errorf("buyer unknown: %w", err)
	}
	if balance < price {
		return nil, fmt.Errorf("insufficient credits")
	}
	var itemOwner string
	var equipped int
	if err := tx.QueryRow(
		`SELECT owner_character_id, equipped FROM crafted_items WHERE id = ?`, itemID,
	).Scan(&itemOwner, &equipped); err != nil {
		return nil, fmt.Errorf("item missing")
	}
	if itemOwner != vOwner {
		return nil, fmt.Errorf("item not vendor-held")
	}
	if _, err := tx.Exec(
		`UPDATE characters SET credits = credits - ? WHERE id = ?`, price, buyerID); err != nil {
		return nil, err
	}
	// Phase 7: city vendor-% tax + guild dues split off inside the same tx
	// (OccupyTxTreasury; zero when no taxed city / no dues guild applies, so
	// pre-civic purchases resolve byte-identically).
	till, err := OccupyTxTreasury(tx, vZone, vX, vZ, vOwner, price, time.Now())
	if err != nil {
		return nil, err
	}
	if _, err := tx.Exec(
		`UPDATE structures SET vendor_till = vendor_till + ? WHERE id = ?`, till, vendorID); err != nil {
		return nil, err
	}
	if _, err := tx.Exec(
		`UPDATE crafted_items SET owner_character_id = ?, equipped = 0 WHERE id = ?`,
		buyerID, itemID); err != nil {
		return nil, err
	}
	if equipped != 0 {
		if _, err := tx.Exec(
			`UPDATE characters SET equipped_weapon_id = NULL WHERE id = ? AND equipped_weapon_id = ?`,
			vOwner, itemID); err != nil {
			return nil, err
		}
		if _, err := tx.Exec(
			`UPDATE characters SET equipped_armor_id = NULL WHERE id = ? AND equipped_armor_id = ?`,
			vOwner, itemID); err != nil {
			return nil, err
		}
	}
	if _, err := tx.Exec(
		`UPDATE vendor_listings SET active = 0 WHERE id = ?`, listingID); err != nil {
		return nil, err
	}
	if err := RecordLedgerTx(tx, buyerID, -price, "vendor_sale", "transfer", vendorID, zone); err != nil {
		return nil, err
	}
	if err := RecordLedgerTx(tx, vOwner, till, "vendor_sale", "transfer", buyerID, zone); err != nil {
		return nil, err
	}
	if _, err := tx.Exec(
		`INSERT INTO market_records (listing_id, vendor_id, item_id, seller_id, buyer_id,
			price, zone) VALUES (?, ?, ?, ?, ?, ?, ?)`,
		listingID, vendorID, itemID, vOwner, buyerID, price, zone); err != nil {
		return nil, err
	}
	// Phase 10 B1: interdependence telemetry — crafted-item purchase receipt
	// (provenance check on crafted_items is the §34 KPI's observation point).
	// Appended AFTER the market_records insert inside the same tx; best-effort
	// (a telemetry failure must never fail a gameplay purchase).
	var crafted int
	_ = tx.QueryRow(`SELECT COUNT(*) FROM crafted_items WHERE id = ?`,
		itemID).Scan(&crafted)
	if crafted > 0 {
		_, _ = tx.Exec(
			`INSERT INTO interdependence_events (kind, character_id, counterparty_id, zone)
			 VALUES ('crafted_purchase', ?, ?, ?)`, buyerID, vOwner, zone)
	}
	if err := tx.Commit(); err != nil {
		return nil, err
	}
	return &PurchaseResult{BuyerCharged: true, TillCredited: till}, nil
}

// EnsureMarketTable creates the transaction-history table (GDD 22.3 shape).
func (db *DB) EnsureMarketTable() error {
	_, err := db.conn.Exec(
		`CREATE TABLE IF NOT EXISTS market_records (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			listing_id TEXT NOT NULL,
			vendor_id TEXT NOT NULL,
			item_id TEXT NOT NULL,
			seller_id TEXT NOT NULL,
			buyer_id TEXT NOT NULL,
			price INTEGER NOT NULL,
			zone TEXT NOT NULL DEFAULT '',
			created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
		)`)
	return err
}

// VendorTill returns a vendor's uncollected till.
func (db *DB) VendorTill(vendorID string) (int, error) {
	var till int
	err := db.conn.QueryRow(
		`SELECT vendor_till FROM structures WHERE id = ?`, vendorID).Scan(&till)
	return till, err
}

// CollectTill moves till → owner wallet atomically; returns the amount.
func (db *DB) CollectTill(vendorID, ownerID string) (int, error) {
	tx, err := db.conn.Begin()
	if err != nil {
		return 0, err
	}
	defer tx.Rollback()
	var till int
	var owner string
	if err := tx.QueryRow(
		`SELECT vendor_till, owner_character_id FROM structures WHERE id = ?`,
		vendorID).Scan(&till, &owner); err != nil {
		return 0, err
	}
	if owner != ownerID {
		return 0, fmt.Errorf("not vendor owner")
	}
	if till <= 0 {
		return 0, nil
	}
	if _, err := tx.Exec(
		`UPDATE structures SET vendor_till = 0 WHERE id = ?`, vendorID); err != nil {
		return 0, err
	}
	if _, err := tx.Exec(
		`UPDATE characters SET credits = credits + ? WHERE id = ?`, till, ownerID); err != nil {
		return 0, err
	}
	if err := RecordLedgerTx(tx, ownerID, till, "vendor_sale", "transfer", vendorID, ""); err != nil {
		return 0, err
	}
	if err := tx.Commit(); err != nil {
		return 0, err
	}
	return till, nil
}
