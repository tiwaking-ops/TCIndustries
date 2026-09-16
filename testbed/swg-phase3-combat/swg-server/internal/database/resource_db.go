// ResourceDB extends the DB with Phase 4 persistence: spawns, stacks,
// harvesters, and crafted items. Testbed fork; generic only.
package database

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"math"
	"math/rand"
	"strings"
	"time"

	"swg-server/internal/resources"
)

// EnsureResourceSchema creates Phase 4 tables (and item-equip columns) if absent.
// Called at server startup; base migrate() lists stay untouched.
func (db *DB) EnsureResourceSchema() error {
	stmts := []string{
		`CREATE TABLE IF NOT EXISTS resource_spawns (
			id TEXT PRIMARY KEY,
			type TEXT NOT NULL,
			zone TEXT NOT NULL,
			center_x REAL NOT NULL,
			center_z REAL NOT NULL,
			radius_m REAL NOT NULL,
			peak_concentration INTEGER NOT NULL,
			stats_json TEXT NOT NULL,
			spawned_at DATETIME NOT NULL,
			despawns_at DATETIME NOT NULL,
			active INTEGER NOT NULL DEFAULT 1,
			depleted_pct INTEGER NOT NULL DEFAULT 0
		)`,
		`CREATE TABLE IF NOT EXISTS resource_stacks (
			character_id TEXT NOT NULL,
			spawn_id TEXT NOT NULL,
			resource_type TEXT NOT NULL,
			quantity INTEGER NOT NULL DEFAULT 0,
			stats_json TEXT NOT NULL,
			PRIMARY KEY (character_id, spawn_id)
		)`,
		`CREATE TABLE IF NOT EXISTS harvesters (
			id TEXT PRIMARY KEY,
			owner_character_id TEXT NOT NULL,
			zone TEXT NOT NULL,
			pos_x REAL NOT NULL,
			pos_z REAL NOT NULL,
			kind TEXT NOT NULL,
			hopper_units INTEGER NOT NULL DEFAULT 0,
			maintenance_pool INTEGER NOT NULL DEFAULT 0,
			active INTEGER NOT NULL DEFAULT 1,
			last_tick_at DATETIME,
			spawn_id TEXT NOT NULL DEFAULT ''
		)`,
		`CREATE TABLE IF NOT EXISTS crafted_items (
			id TEXT PRIMARY KEY,
			owner_character_id TEXT NOT NULL,
			schematic_id TEXT NOT NULL,
			name TEXT NOT NULL,
			stats_json TEXT NOT NULL,
			equipped INTEGER NOT NULL DEFAULT 0,
			created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
		)`,
	}
	for i, s := range stmts {
		if _, err := db.conn.Exec(s); err != nil {
			return fmt.Errorf("resource schema %d failed: %w", i+1, err)
		}
	}
	// Equip columns on characters (idempotent: ignore duplicate-column errors).
	for _, col := range []string{
		`ALTER TABLE characters ADD COLUMN equipped_weapon_id TEXT`,
		`ALTER TABLE characters ADD COLUMN equipped_armor_id TEXT`,
	} {
		if _, err := db.conn.Exec(col); err != nil &&
			!strings.Contains(err.Error(), "duplicate column") {
			return fmt.Errorf("equip column failed: %w", err)
		}
	}
	return nil
}

// --- Spawns ---

// SpawnRow maps one resource_spawns row.
type SpawnRow struct {
	ID                string
	Type              string
	Zone              string
	CenterX, CenterZ  float64
	RadiusM           float64
	PeakConcentration int
	Stats             map[string]int
	SpawnedAt         time.Time
	DespawnsAt        time.Time
	Active            bool
	DepletedPct       int // local depletion from sampling (GDD 11.3.2)
}

// scanSpawnRow maps one resource_spawns row.
func scanSpawnRow(row interface {
	Scan(...interface{}) error
}) (*SpawnRow, error) {
	var r SpawnRow
	var statsJSON string
	var active int
	var spawned, despawns string
	if err := row.Scan(&r.ID, &r.Type, &r.Zone, &r.CenterX, &r.CenterZ,
		&r.RadiusM, &r.PeakConcentration, &statsJSON, &spawned, &despawns, &active,
		&r.DepletedPct); err != nil {
		return nil, err
	}
	if err := json.Unmarshal([]byte(statsJSON), &r.Stats); err != nil {
		return nil, err
	}
	r.SpawnedAt, _ = time.Parse("2006-01-02 15:04:05+00:00", spawned)
	r.DespawnsAt, _ = time.Parse("2006-01-02 15:04:05+00:00", despawns)
	if r.SpawnedAt.IsZero() {
		r.SpawnedAt, _ = time.Parse("2006-01-02 15:04:05", spawned)
	}
	if r.DespawnsAt.IsZero() {
		r.DespawnsAt, _ = time.Parse("2006-01-02 15:04:05", despawns)
	}
	r.Active = active != 0
	return &r, nil
}

// InsertSpawn persists one spawn.
func (db *DB) InsertSpawn(s *resources.Spawn) error {
	statsJSON, _ := json.Marshal(s.Stats)
	_, err := db.conn.Exec(
		`INSERT OR IGNORE INTO resource_spawns
			(id, type, zone, center_x, center_z, radius_m, peak_concentration,
			 stats_json, spawned_at, despawns_at, active)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, 1)`,
		s.ID, s.Type, s.Zone, s.CenterX, s.CenterZ, s.RadiusM,
		s.PeakConcentration, string(statsJSON),
		s.SpawnedAt.UTC().Format("2006-01-02 15:04:05"),
		s.DespawnsAt.UTC().Format("2006-01-02 15:04:05"),
	)
	return err
}

// ActiveSpawns returns all live spawns in a zone.
func (db *DB) ActiveSpawns(zone string) ([]SpawnRow, error) {
	rows, err := db.conn.Query(
		`SELECT id, type, zone, center_x, center_z, radius_m, peak_concentration,
			stats_json, spawned_at, despawns_at, active, depleted_pct
		FROM resource_spawns WHERE zone = ? AND active = 1`, zone)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var out []SpawnRow
	for rows.Next() {
		r, err := scanSpawnRow(rows)
		if err != nil {
			return nil, err
		}
		out = append(out, *r)
	}
	return out, rows.Err()
}

// SeedZoneSpawns tops the zone up to 20 active spawns (GDD 20–50 band minimum)
// across all 7 types, plus deterministic guarantee spawns in fast-test mode
// (documented test scaffolding: fixed IDs, high concentration, near spawn).
func (db *DB) SeedZoneSpawns(zone string, rng *rand.Rand, now time.Time) (int, error) {
	if resources.FastCycle() {
		// Deterministic guarantee spawns (documented test scaffolding): one
		// high-concentration spawn PER TYPE near the zone spawn, with a long
		// (30 min) life. Random spawns live 120–180 s while test walks take
		// ~1 min — without per-type guarantees, quota tests chase despawning
		// ghosts (diagnosed 2026-09-15). Random spawns still exercise the
		// lifecycle path; guarantees exercise the gather path.
		guarantees := []struct {
			id, rtype string
			x, z      float64
		}{
			{"testspawn-metal-01", "ferric_metal", 30, 0},
			{"testspawn-polymer-01", "structural_polymer", 40, 0},
			{"testspawn-organic-01", "cultured_organic", 20, 25},
			{"testspawn-chemical-01", "industrial_chemical", 45, 15},
			{"testspawn-alloy-01", "conductive_alloy", -5, 10},
			{"testspawn-water-01", "filtered_water", 35, -20},
			{"testspawn-flora-01", "fibrous_flora", 10, -25},
		}
		for _, g := range guarantees {
			t := resources.TypeByID(g.rtype)
			stats := map[string]int{resources.LaneOQ: 800}
			for _, lane := range t.RelevantLanes {
				stats[lane] = 800
			}
			statsJSON, _ := json.Marshal(stats)
			if _, err := db.conn.Exec(
				`INSERT OR IGNORE INTO resource_spawns
					(id, type, zone, center_x, center_z, radius_m, peak_concentration,
					 stats_json, spawned_at, despawns_at, active)
				VALUES (?, ?, ?, ?, ?, 15.0, 80, ?, ?, ?, 1)`,
				g.id, g.rtype, zone, g.x, g.z, string(statsJSON),
				now.UTC().Format("2006-01-02 15:04:05"),
				now.Add(30*time.Minute).UTC().Format("2006-01-02 15:04:05"),
			); err != nil {
				return 0, err
			}
		}
	}
	active, err := db.ActiveSpawns(zone)
	if err != nil {
		return 0, err
	}
	// Type coverage first (GDD 11.2.3: each zone has a defined type table —
	// every type must actually be present; pure-random seeding could starve a
	// type zone-wide indefinitely, and same-type relocation would preserve the
	// gap forever).
	present := map[string]bool{}
	for _, s := range active {
		present[s.Type] = true
	}
	for _, t := range resources.Types {
		if present[t.ID] {
			continue
		}
		s := resources.GenerateSpawn(rng,
			fmt.Sprintf("spawn-%d-cover-%s", now.UnixNano(), t.ID),
			t.ID, zone,
			(rng.Float64()-0.5)*600, (rng.Float64()-0.5)*600, now)
		if err := db.InsertSpawn(s); err != nil {
			return 0, err
		}
		present[t.ID] = true
	}
	active, err = db.ActiveSpawns(zone)
	if err != nil {
		return 0, err
	}
	need := 20 - len(active)
	for i := 0; i < need; i++ {
		t := resources.Types[rng.Intn(len(resources.Types))]
		s := resources.GenerateSpawn(rng,
			fmt.Sprintf("spawn-%d-%d", now.UnixNano(), i),
			t.ID, zone,
			(rng.Float64()-0.5)*1000, (rng.Float64()-0.5)*1000, now)
		if err := db.InsertSpawn(s); err != nil {
			return 0, err
		}
	}
	active, err = db.ActiveSpawns(zone)
	if err != nil {
		return 0, err
	}
	return len(active), nil
}

// TickSpawns retires expired spawns and relocates same-type replacements
// (GDD relocate-on-despawn). Returns (retired, created).
func (db *DB) TickSpawns(zone string, rng *rand.Rand, now time.Time) (int, int, error) {
	active, err := db.ActiveSpawns(zone)
	if err != nil {
		return 0, 0, err
	}
	retired, created := 0, 0
	for _, s := range active {
		if s.ID[:9] == "testspawn" {
			continue // guarantee spawns persist for the test session
		}
		if now.After(s.DespawnsAt) {
			if _, err := db.conn.Exec(
				`UPDATE resource_spawns SET active = 0 WHERE id = ?`, s.ID); err != nil {
				return retired, created, err
			}
			retired++
			ns := resources.GenerateSpawn(rng,
				fmt.Sprintf("spawn-%d-%s", now.UnixNano(), s.ID),
				s.Type, zone,
				(rng.Float64()-0.5)*1000, (rng.Float64()-0.5)*1000, now)
			if err := db.InsertSpawn(ns); err != nil {
				return retired, created, err
			}
			created++
		}
	}
	return retired, created, nil
}

// GetSpawn loads one spawn by ID.
func (db *DB) GetSpawn(id string) (*SpawnRow, error) {
	return scanSpawnRow(db.conn.QueryRow(
		`SELECT id, type, zone, center_x, center_z, radius_m, peak_concentration,
			stats_json, spawned_at, despawns_at, active, depleted_pct
		FROM resource_spawns WHERE id = ?`, id))
}

// BumpDepletion adds local depletion (GDD 11.3.2 diminishing returns).
func (db *DB) BumpDepletion(id string, pct int) error {
	_, err := db.conn.Exec(
		`UPDATE resource_spawns SET depleted_pct = depleted_pct + ? WHERE id = ?`, pct, id)
	return err
}

// EffectiveConcentration folds depletion into the gradient concentration.
func EffectiveConcentration(s *SpawnRow, x, z float64) int {
	base := concAt(s, x, z)
	eff := float64(base) * (1 - float64(s.DepletedPct)/100.0)
	if eff < 0 {
		eff = 0
	}
	return int(eff + 0.5)
}

func concAt(s *SpawnRow, x, z float64) int {
	dx := x - s.CenterX
	dz := z - s.CenterZ
	d := math.Sqrt(dx*dx + dz*dz)
	if d >= s.RadiusM {
		return 0
	}
	frac := d / s.RadiusM
	return int(float64(s.PeakConcentration)*(0.1+0.9*math.Exp(-2.3*frac*frac)) + 0.5)
}

// --- Stacks (inventory) ---

// StackRow maps one resource_stacks row.
type StackRow struct {
	SpawnID      string
	ResourceType string
	Quantity     int
	Stats        map[string]int
}

// AddToStack adds units to a character's stack for a spawn (upsert; stacks key
// by spawn ID — different spawns never merge, GDD 11.5.1).
func (db *DB) AddToStack(characterID, spawnID, rtype string, qty int, stats map[string]int) error {
	statsJSON, _ := json.Marshal(stats)
	_, err := db.conn.Exec(
		`INSERT INTO resource_stacks (character_id, spawn_id, resource_type, quantity, stats_json)
		VALUES (?, ?, ?, ?, ?)
		ON CONFLICT(character_id, spawn_id) DO UPDATE SET quantity = quantity + ?`,
		characterID, spawnID, rtype, qty, string(statsJSON), qty,
	)
	return err
}

// GetStacks returns all stacks for a character.
func (db *DB) GetStacks(characterID string) ([]StackRow, error) {
	rows, err := db.conn.Query(
		`SELECT spawn_id, resource_type, quantity, stats_json
		FROM resource_stacks WHERE character_id = ?`, characterID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var out []StackRow
	for rows.Next() {
		var r StackRow
		var statsJSON string
		if err := rows.Scan(&r.SpawnID, &r.ResourceType, &r.Quantity, &statsJSON); err != nil {
			return nil, err
		}
		_ = json.Unmarshal([]byte(statsJSON), &r.Stats)
		out = append(out, r)
	}
	return out, rows.Err()
}

// ConsumeStack deducts units; errors when insufficient (used at finalize).
func (db *DB) ConsumeStack(characterID, spawnID string, qty int) error {
	var have int
	err := db.conn.QueryRow(
		`SELECT quantity FROM resource_stacks WHERE character_id = ? AND spawn_id = ?`,
		characterID, spawnID).Scan(&have)
	if err == sql.ErrNoRows || have < qty {
		return fmt.Errorf("insufficient stack %s: have %d, need %d", spawnID, have, qty)
	}
	if err != nil {
		return err
	}
	_, err = db.conn.Exec(
		`UPDATE resource_stacks SET quantity = quantity - ? WHERE character_id = ? AND spawn_id = ?`,
		qty, characterID, spawnID)
	return err
}

// --- Harvesters ---

// HarvesterKinds: rates (units/hr) and hoppers per GDD 11.4.2; weekly fee per 11.4.3.
var HarvesterKinds = map[string]struct {
	RatePerHour int
	HopperCap   int
	WeeklyFee   int
}{
	"personal": {RatePerHour: 50, HopperCap: 10000, WeeklyFee: 500},
	"medium":   {RatePerHour: 100, HopperCap: 25000, WeeklyFee: 1500},
	"heavy":    {RatePerHour: 200, HopperCap: 50000, WeeklyFee: 5000},
}

// HarvesterRow maps one harvesters row.
type HarvesterRow struct {
	ID               string
	OwnerCharacterID string
	Zone             string
	PosX, PosZ       float64
	Kind             string
	HopperUnits      int
	MaintenancePool  int
	Active           bool
	SpawnID          string
}

// PlaceHarvester records a harvester deed placement.
func (db *DB) PlaceHarvester(id, owner, zone string, x, z float64, kind, spawnID string) error {
	_, err := db.conn.Exec(
		`INSERT INTO harvesters
			(id, owner_character_id, zone, pos_x, pos_z, kind, spawn_id, last_tick_at)
		VALUES (?, ?, ?, ?, ?, ?, ?, CURRENT_TIMESTAMP)`,
		id, owner, zone, x, z, kind, spawnID,
	)
	return err
}

// GetHarvesters returns harvesters owned by a character.
func (db *DB) GetHarvesters(owner string) ([]HarvesterRow, error) {
	return db.queryHarvesters(`WHERE owner_character_id = ?`, owner)
}

// AllHarvesters returns every harvester (tick driver).
func (db *DB) AllHarvesters() ([]HarvesterRow, error) {
	return db.queryHarvesters(``, nil)
}

func (db *DB) queryHarvesters(where string, arg interface{}) ([]HarvesterRow, error) {
	q := `SELECT id, owner_character_id, zone, pos_x, pos_z, kind, hopper_units,
		maintenance_pool, active, spawn_id FROM harvesters ` + where
	var rows *sql.Rows
	var err error
	if arg == nil {
		rows, err = db.conn.Query(q)
	} else {
		rows, err = db.conn.Query(q, arg)
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var out []HarvesterRow
	for rows.Next() {
		var r HarvesterRow
		var active int
		if err := rows.Scan(&r.ID, &r.OwnerCharacterID, &r.Zone, &r.PosX, &r.PosZ,
			&r.Kind, &r.HopperUnits, &r.MaintenancePool, &active, &r.SpawnID); err != nil {
			return nil, err
		}
		r.Active = active != 0
		out = append(out, r)
	}
	return out, rows.Err()
}

// FundHarvester deposits credits into the maintenance pool.
func (db *DB) FundHarvester(id string, credits int) error {
	_, err := db.conn.Exec(
		`UPDATE harvesters SET maintenance_pool = maintenance_pool + ? WHERE id = ?`, credits, id)
	return err
}

// UpdateHarvester persists per-tick hopper/pool/active state.
func (db *DB) UpdateHarvester(h *HarvesterRow) error {
	active := 0
	if h.Active {
		active = 1
	}
	_, err := db.conn.Exec(
		`UPDATE harvesters SET hopper_units = ?, maintenance_pool = ?, active = ?,
			last_tick_at = CURRENT_TIMESTAMP WHERE id = ?`,
		h.HopperUnits, h.MaintenancePool, active, h.ID)
	return err
}

// DeleteItem consumes an item (e.g. a deed on placement).
func (db *DB) DeleteItem(owner, id string) error {
	_, err := db.conn.Exec(
		`DELETE FROM crafted_items WHERE id = ? AND owner_character_id = ?`, id, owner)
	return err
}

// EmptyHopper transfers hopper units into the owner's stack (spawn-typed) and
// returns the transferred amount.
func (db *DB) EmptyHopper(h *HarvesterRow, spawnType string, stats map[string]int) (int, error) {
	units := h.HopperUnits
	if units <= 0 {
		return 0, nil
	}
	if _, err := db.conn.Exec(
		`UPDATE harvesters SET hopper_units = 0 WHERE id = ?`, h.ID); err != nil {
		return 0, err
	}
	if err := db.AddToStack(h.OwnerCharacterID, "hopper:"+h.SpawnID, spawnType, units, stats); err != nil {
		return 0, err
	}
	return units, nil
}

// --- Crafted items ---

// CraftedItemRow maps one crafted_items row.
type CraftedItemRow struct {
	ID         string
	Schematic  string
	Name       string
	Stats      map[string]float64
	Equipped   bool
	// Condition: 0–100 item condition (Phase 8 minimal surface for PvP death
	// penalties; no breakage, no repair — OQ-010 follow-ups).
	Condition int
}

// InsertItem records a crafted item; returns its ID.
func (db *DB) InsertItem(owner, schematic, name string, stats map[string]float64) (string, error) {
	id := fmt.Sprintf("item-%d", time.Now().UnixNano())
	statsJSON, _ := json.Marshal(stats)
	_, err := db.conn.Exec(
		`INSERT INTO crafted_items (id, owner_character_id, schematic_id, name, stats_json)
		VALUES (?, ?, ?, ?, ?)`, id, owner, schematic, name, string(statsJSON))
	if err != nil {
		return "", err
	}
	return id, nil
}

// itemCols lists crafted_items columns (condition added Phase 8).
const itemCols = `id, schematic_id, name, stats_json, equipped, condition_pct`

// GetItems returns a character's crafted items.
func (db *DB) GetItems(owner string) ([]CraftedItemRow, error) {
	rows, err := db.conn.Query(
		`SELECT `+itemCols+` FROM crafted_items
		WHERE owner_character_id = ?`, owner)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var out []CraftedItemRow
	for rows.Next() {
		var r CraftedItemRow
		var statsJSON string
		var equipped int
		if err := rows.Scan(&r.ID, &r.Schematic, &r.Name, &statsJSON,
			&equipped, &r.Condition); err != nil {
			return nil, err
		}
		_ = json.Unmarshal([]byte(statsJSON), &r.Stats)
		r.Equipped = equipped != 0
		out = append(out, r)
	}
	return out, rows.Err()
}

// GetItem loads one item (ownership-checked by caller via owner filter).
func (db *DB) GetItem(owner, id string) (*CraftedItemRow, error) {
	var r CraftedItemRow
	var statsJSON string
	var equipped int
	err := db.conn.QueryRow(
		`SELECT `+itemCols+` FROM crafted_items
		WHERE id = ? AND owner_character_id = ?`, id, owner,
	).Scan(&r.ID, &r.Schematic, &r.Name, &statsJSON, &equipped, &r.Condition)
	if err != nil {
		return nil, err
	}
	_ = json.Unmarshal([]byte(statsJSON), &r.Stats)
	r.Equipped = equipped != 0
	return &r, nil
}

// EquipItem marks an item equipped (slot: "weapon" or "armor") and records it on
// the character row; clears any previously equipped item of that slot.
func (db *DB) EquipItem(owner, id, slot string) error {
	col := "equipped_weapon_id"
	if slot == "armor" {
		col = "equipped_armor_id"
	}
	var prev sql.NullString
	if err := db.conn.QueryRow(
		`SELECT `+col+` FROM characters WHERE id = ?`, owner).Scan(&prev); err != nil {
		return err
	}
	tx, err := db.conn.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()
	if prev.Valid && prev.String != "" {
		if _, err := tx.Exec(
			`UPDATE crafted_items SET equipped = 0 WHERE id = ? AND owner_character_id = ?`,
			prev.String, owner); err != nil {
			return err
		}
	}
	if _, err := tx.Exec(
		`UPDATE crafted_items SET equipped = 1 WHERE id = ? AND owner_character_id = ?`,
		id, owner); err != nil {
		return err
	}
	if _, err := tx.Exec(
		`UPDATE characters SET `+col+` = ? WHERE id = ?`, id, owner); err != nil {
		return err
	}
	return tx.Commit()
}

// EquippedItems returns the character's equipped weapon/armor item rows (nil if none).
func (db *DB) EquippedItems(characterID string) (weapon, armor *CraftedItemRow, err error) {
	var wID, aID sql.NullString
	if err := db.conn.QueryRow(
		`SELECT equipped_weapon_id, equipped_armor_id FROM characters WHERE id = ?`,
		characterID).Scan(&wID, &aID); err != nil {
		return nil, nil, err
	}
	load := func(ns sql.NullString) (*CraftedItemRow, error) {
		if !ns.Valid || ns.String == "" {
			return nil, nil
		}
		return db.GetItem(characterID, ns.String)
	}
	if weapon, err = load(wID); err != nil {
		return nil, nil, err
	}
	if armor, err = load(aID); err != nil {
		return nil, nil, err
	}
	return weapon, armor, nil
}

// --- Corpse harvest support ---

// HarvestableCorpse is a dead, unexpired creature instance.
type HarvestableCorpse struct {
	InstanceID string
	TemplateID string
	Zone       string
	PosX, PosZ float64
}

// HarvestableCorpses returns dead instances in range whose corpse window holds.
func (db *DB) HarvestableCorpses(zone string, x, z, radiusM float64, now time.Time) ([]HarvestableCorpse, error) {
	rows, err := db.conn.Query(
		`SELECT id, template_id, zone, pos_x, pos_z, corpse_expires_at
		FROM creature_instances WHERE zone = ? AND state = 'dead'`, zone)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var out []HarvestableCorpse
	for rows.Next() {
		var c HarvestableCorpse
		var expStr sql.NullString
		var px, pz float64
		var id, tmpl, zn string
		if err := rows.Scan(&id, &tmpl, &zn, &px, &pz, &expStr); err != nil {
			return nil, err
		}
		if !expStr.Valid {
			continue
		}
		exp, ok := parseSQLiteTime(expStr.String)
		if !ok {
			continue // unparseable expiry: skip rather than misjudge
		}
		if now.After(exp) {
			continue
		}
		dx, dz := x-px, z-pz
		if dx*dx+dz*dz > radiusM*radiusM {
			continue
		}
		c = HarvestableCorpse{InstanceID: id, TemplateID: tmpl, Zone: zn, PosX: px, PosZ: pz}
		out = append(out, c)
	}
	return out, rows.Err()
}

// parseSQLiteTime parses the timestamp shapes the sqlite driver writes
// (time.Time values carry fractional seconds + zone offset; our own formatted
// values are plain "2006-01-02 15:04:05"). Reports ok=false when unparseable.
func parseSQLiteTime(s string) (time.Time, bool) {
	for _, layout := range []string{
		"2006-01-02 15:04:05.999999999-07:00",
		"2006-01-02T15:04:05.999999999-07:00",
		time.RFC3339Nano,
		time.RFC3339,
		"2006-01-02 15:04:05-07:00",
		"2006-01-02 15:04:05+00:00",
		"2006-01-02 15:04:05",
	} {
		if t, err := time.Parse(layout, s); err == nil {
			return t, true
		}
	}
	return time.Time{}, false
}

// ConsumeCorpse marks a harvested corpse fully consumed (state stays dead; expiry
// set to now so it cannot be harvested twice).
func (db *DB) ConsumeCorpse(id string, now time.Time) error {
	_, err := db.conn.Exec(
		`UPDATE creature_instances SET corpse_expires_at = ? WHERE id = ?`,
		now.UTC().Format("2006-01-02 15:04:05"), id)
	return err
}
