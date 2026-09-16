// CombatDB extends the DB with Phase 3 combat persistence. All functions here
// operate on the same DB connection as the core DB. Testbed fork; generic only.
package database

import (
	"database/sql"
	"fmt"
	"time"

	"swg-server/internal/creatures"
)

// CombatState is the full combat-relevant persisted state for one character.
// Unlike the Phase-0 read path (which recomputes HAM from species), this reads
// current values, wounds, battle fatigue, posture, stance, and timers verbatim —
// Phase 3 is the first phase where current-vs-max actually matters.
type CombatState struct {
	CharacterID      string
	HealthCurrent    int
	HealthMax        int
	ActionCurrent    int
	ActionMax        int
	MindCurrent      int
	MindMax          int
	WoundsHealth     int
	WoundsAction     int
	WoundsMind       int
	BattleFatiguePct float64
	Posture          string
	Stance           string
	IncapacitatedAt  *time.Time
}

// EnsureCombatSchema creates Phase 3 tables if absent. Called at server startup
// (see server Run) so the fork's base migrate() list stays untouched.
func (db *DB) EnsureCombatSchema() error {
	stmts := []string{
		`CREATE TABLE IF NOT EXISTS lairs (
			id TEXT PRIMARY KEY,
			template_id TEXT NOT NULL,
			zone TEXT NOT NULL,
			pos_x REAL NOT NULL,
			pos_z REAL NOT NULL,
			lair_hp INTEGER NOT NULL,
			lair_hp_max INTEGER NOT NULL,
			max_population INTEGER NOT NULL,
			destroyed_at DATETIME
		)`,
		`CREATE TABLE IF NOT EXISTS creature_instances (
			id TEXT PRIMARY KEY,
			template_id TEXT NOT NULL,
			lair_id TEXT NOT NULL DEFAULT '',
			zone TEXT NOT NULL,
			pos_x REAL NOT NULL,
			pos_y REAL NOT NULL DEFAULT 5.0,
			pos_z REAL NOT NULL,
			health_current INTEGER NOT NULL,
			health_max INTEGER NOT NULL,
			state TEXT NOT NULL DEFAULT 'idle',
			target_character_id TEXT NOT NULL DEFAULT '',
			corpse_expires_at DATETIME
		)`,
		`CREATE TABLE IF NOT EXISTS clone_bindings (
			character_id TEXT PRIMARY KEY,
			zone TEXT NOT NULL,
			pos_x REAL NOT NULL,
			pos_y REAL NOT NULL,
			pos_z REAL NOT NULL
		)`,
	}
	for i, s := range stmts {
		if _, err := db.conn.Exec(s); err != nil {
			return fmt.Errorf("combat schema %d failed: %w", i+1, err)
		}
	}
	return nil
}

// GetCombatState loads the persisted combat state for a character.
func (db *DB) GetCombatState(characterID string) (*CombatState, error) {
	s := &CombatState{CharacterID: characterID}
	var incap sql.NullTime
	err := db.conn.QueryRow(
		`SELECT health_current, health_max, action_current, action_max,
			mind_current, mind_max, wounds_health, wounds_action, wounds_mind,
			battle_fatigue_pct, posture, stance, incapacitated_at
		FROM ham_pool_states WHERE character_id = ?`,
		characterID,
	).Scan(
		&s.HealthCurrent, &s.HealthMax, &s.ActionCurrent, &s.ActionMax,
		&s.MindCurrent, &s.MindMax, &s.WoundsHealth, &s.WoundsAction, &s.WoundsMind,
		&s.BattleFatiguePct, &s.Posture, &s.Stance, &incap,
	)
	if err != nil {
		return nil, err
	}
	if incap.Valid {
		t := incap.Time
		s.IncapacitatedAt = &t
	}
	return s, nil
}

// UpdateCombatState persists the combat state for a character.
func (db *DB) UpdateCombatState(s *CombatState) error {
	_, err := db.conn.Exec(
		`UPDATE ham_pool_states SET health_current = ?, health_max = ?,
			action_current = ?, action_max = ?, mind_current = ?, mind_max = ?,
			wounds_health = ?, wounds_action = ?, wounds_mind = ?,
			battle_fatigue_pct = ?, posture = ?, stance = ?,
			incapacitated_at = ?, updated_at = CURRENT_TIMESTAMP
		WHERE character_id = ?`,
		s.HealthCurrent, s.HealthMax, s.ActionCurrent, s.ActionMax,
		s.MindCurrent, s.MindMax, s.WoundsHealth, s.WoundsAction, s.WoundsMind,
		s.BattleFatiguePct, s.Posture, s.Stance, s.IncapacitatedAt, s.CharacterID,
	)
	return err
}

// SeedLairs inserts the starter lairs (idempotent) and tops up each lair's live
// population to its MaxPopulation. Returns the live instance count.
func (db *DB) SeedLairs() (int, error) {
	for _, l := range creatures.StarterLairs() {
		if _, err := db.conn.Exec(
			`INSERT OR IGNORE INTO lairs
				(id, template_id, zone, pos_x, pos_z, lair_hp, lair_hp_max, max_population)
			VALUES (?, ?, ?, ?, ?, ?, ?, ?)`,
			l.ID, l.TemplateID, l.Zone, l.PosX, l.PosZ, l.LairHP, l.LairHPMax, l.MaxPopulation,
		); err != nil {
			return 0, fmt.Errorf("seed lair %s: %w", l.ID, err)
		}
		var alive int
		if err := db.conn.QueryRow(
			`SELECT COUNT(*) FROM creature_instances
			WHERE lair_id = ? AND state != 'dead'`, l.ID,
		).Scan(&alive); err != nil {
			return 0, err
		}
		tmpl := creatures.TemplateByID(l.TemplateID)
		for i := alive; i < l.MaxPopulation; i++ {
			id := fmt.Sprintf("%s-%d", newRowID(l.ID), i)
			if _, err := db.conn.Exec(
				`INSERT INTO creature_instances
					(id, template_id, lair_id, zone, pos_x, pos_y, pos_z,
					 health_current, health_max, state)
				VALUES (?, ?, ?, ?, ?, 5.0, ?, ?, ?, 'idle')`,
				id, l.TemplateID, l.ID, l.Zone, l.PosX, l.PosZ,
				tmpl.HealthMax, tmpl.HealthMax,
			); err != nil {
				return 0, fmt.Errorf("spawn instance: %w", err)
			}
		}
	}
	insts, err := db.GetLivingCreatureInstances()
	if err != nil {
		return 0, err
	}
	return len(insts), nil
}

// CreatureInstanceRow maps one creature_instances row.
type CreatureInstanceRow struct {
	ID                string
	TemplateID        string
	LairID            string
	Zone              string
	PosX, PosY, PosZ  float64
	HealthCurrent     int
	HealthMax         int
	State             string
	TargetCharacterID string
}

// GetLivingCreatureInstances returns all non-dead instances.
func (db *DB) GetLivingCreatureInstances() ([]CreatureInstanceRow, error) {
	rows, err := db.conn.Query(
		`SELECT id, template_id, lair_id, zone, pos_x, pos_y, pos_z,
			health_current, health_max, state, target_character_id
		FROM creature_instances WHERE state != 'dead'`,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var out []CreatureInstanceRow
	for rows.Next() {
		var r CreatureInstanceRow
		if err := rows.Scan(&r.ID, &r.TemplateID, &r.LairID, &r.Zone,
			&r.PosX, &r.PosY, &r.PosZ, &r.HealthCurrent, &r.HealthMax,
			&r.State, &r.TargetCharacterID); err != nil {
			return nil, err
		}
		out = append(out, r)
	}
	return out, rows.Err()
}

// UpdateCreatureInstance persists an instance's combat-mutable fields.
func (db *DB) UpdateCreatureInstance(r *CreatureInstanceRow) error {
	_, err := db.conn.Exec(
		`UPDATE creature_instances SET health_current = ?, state = ?,
			target_character_id = ?, corpse_expires_at = ? WHERE id = ?`,
		r.HealthCurrent, r.State, r.TargetCharacterID, nil, r.ID,
	)
	return err
}

// KillCreatureInstance marks an instance dead with a corpse-expiry timestamp.
func (db *DB) KillCreatureInstance(id string, expires time.Time) error {
	_, err := db.conn.Exec(
		`UPDATE creature_instances SET health_current = 0, state = 'dead',
			target_character_id = '', corpse_expires_at = ? WHERE id = ?`,
		expires, id,
	)
	return err
}

// --- Phase 9 lair cycle (GDD 9.5.3/17.2.3: damage, regen, relocation) ---

// LairRow maps one lairs row.
type LairRow struct {
	ID            string
	TemplateID    string
	Zone          string
	PosX, PosZ    float64
	LairHP        int
	LairHPMax     int
	MaxPopulation int
	DestroyedAt   sql.NullString
}

// AllLairs returns every lair.
func (db *DB) AllLairs() ([]LairRow, error) {
	rows, err := db.conn.Query(
		`SELECT id, template_id, zone, pos_x, pos_z, lair_hp, lair_hp_max,
			max_population, destroyed_at FROM lairs`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var out []LairRow
	for rows.Next() {
		var l LairRow
		if err := rows.Scan(&l.ID, &l.TemplateID, &l.Zone, &l.PosX, &l.PosZ,
			&l.LairHP, &l.LairHPMax, &l.MaxPopulation, &l.DestroyedAt); err != nil {
			return nil, err
		}
		out = append(out, l)
	}
	return out, rows.Err()
}

// GetLair loads one lair.
func (db *DB) GetLair(id string) (*LairRow, error) {
	var l LairRow
	err := db.conn.QueryRow(
		`SELECT id, template_id, zone, pos_x, pos_z, lair_hp, lair_hp_max,
			max_population, destroyed_at FROM lairs WHERE id = ?`, id,
	).Scan(&l.ID, &l.TemplateID, &l.Zone, &l.PosX, &l.PosZ,
		&l.LairHP, &l.LairHPMax, &l.MaxPopulation, &l.DestroyedAt)
	if err != nil {
		return nil, err
	}
	return &l, nil
}

// UpdateLairHP persists lair damage.
func (db *DB) UpdateLairHP(id string, hp int) error {
	_, err := db.conn.Exec(
		`UPDATE lairs SET lair_hp = ? WHERE id = ?`, hp, id)
	return err
}

// DestroyLair marks a lair destroyed (spawning stops; relocation later).
func (db *DB) DestroyLair(id string) error {
	_, err := db.conn.Exec(
		`UPDATE lairs SET destroyed_at = CURRENT_TIMESTAMP WHERE id = ?`, id)
	return err
}

// LivingCountForLair counts non-dead instances of a lair.
func (db *DB) LivingCountForLair(lairID string) (int, error) {
	var n int
	err := db.conn.QueryRow(
		`SELECT COUNT(*) FROM creature_instances
		WHERE lair_id = ? AND state != 'dead'`, lairID).Scan(&n)
	return n, err
}

// SpawnLairInstance adds one live instance to a lair; returns its ID.
func (db *DB) SpawnLairInstance(lairID, zone, templateID string, x, z float64, hp int) (string, error) {
	id := fmt.Sprintf("%s-r", newRowID(lairID))
	_, err := db.conn.Exec(
		`INSERT INTO creature_instances
			(id, template_id, lair_id, zone, pos_x, pos_y, pos_z,
			 health_current, health_max, state)
		VALUES (?, ?, ?, ?, ?, 5.0, ?, ?, ?, 'idle')`,
		id, templateID, lairID, zone, x, z, hp, hp,
	)
	if err != nil {
		return "", err
	}
	return id, nil
}

// RelocatableLairs returns destroyed lairs older than the given SQLite
// datetime modifier (e.g. "-12 hours"; fast-cycle "-10 minutes").
func (db *DB) RelocatableLairs(modifier string) ([]LairRow, error) {
	rows, err := db.conn.Query(
		`SELECT id, template_id, zone, pos_x, pos_z, lair_hp, lair_hp_max,
			max_population, destroyed_at FROM lairs
		WHERE destroyed_at IS NOT NULL
		AND destroyed_at <= datetime('now', ?)`, modifier)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var out []LairRow
	for rows.Next() {
		var l LairRow
		if err := rows.Scan(&l.ID, &l.TemplateID, &l.Zone, &l.PosX, &l.PosZ,
			&l.LairHP, &l.LairHPMax, &l.MaxPopulation, &l.DestroyedAt); err != nil {
			return nil, err
		}
		out = append(out, l)
	}
	return out, rows.Err()
}

// InsertLair creates a relocated lair; returns its ID.
func (db *DB) InsertLair(templateID, zone string, x, z float64, hpMax, maxPop int) (string, error) {
	id := newRowID("lair")
	_, err := db.conn.Exec(
		`INSERT INTO lairs
			(id, template_id, zone, pos_x, pos_z, lair_hp, lair_hp_max, max_population)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?)`,
		id, templateID, zone, x, z, hpMax, hpMax, maxPop)
	if err != nil {
		return "", err
	}
	return id, nil
}

// DeleteLair removes a lair row (relocated-away husks).
func (db *DB) DeleteLair(id string) error {
	_, err := db.conn.Exec(`DELETE FROM lairs WHERE id = ?`, id)
	return err
}

// DeleteCreatureInstance removes an instance without a corpse (Phase 9
// taming path — taming isn't killing, so no corpse is left behind).
func (db *DB) DeleteCreatureInstance(id string) error {
	_, err := db.conn.Exec(`DELETE FROM creature_instances WHERE id = ?`, id)
	return err
}

// SetCloneBinding records a character's respawn point.
func (db *DB) SetCloneBinding(characterID, zone string, x, y, z float64) error {
	_, err := db.conn.Exec(
		`INSERT INTO clone_bindings (character_id, zone, pos_x, pos_y, pos_z)
		VALUES (?, ?, ?, ?, ?)
		ON CONFLICT(character_id) DO UPDATE SET zone = ?, pos_x = ?, pos_y = ?, pos_z = ?`,
		characterID, zone, x, y, z, zone, x, y, z,
	)
	return err
}

// CloneBinding holds a respawn point.
type CloneBinding struct {
	Zone       string
	X, Y, Z    float64
}

// GetCloneBinding returns the bound respawn point, or the zone default spawn when
// unbound (first death before any explicit bind — GDD 9.6.3 city facilities are
// later-phase scope, so the zone spawn is the only facility).
func (db *DB) GetCloneBinding(characterID string) (*CloneBinding, error) {
	var b CloneBinding
	err := db.conn.QueryRow(
		`SELECT zone, pos_x, pos_y, pos_z FROM clone_bindings WHERE character_id = ?`,
		characterID,
	).Scan(&b.Zone, &b.X, &b.Y, &b.Z)
	if err == sql.ErrNoRows {
		return &CloneBinding{Zone: "zone-0001", X: 20.0, Y: 5.0, Z: 0.0}, nil
	}
	if err != nil {
		return nil, err
	}
	return &b, nil
}
