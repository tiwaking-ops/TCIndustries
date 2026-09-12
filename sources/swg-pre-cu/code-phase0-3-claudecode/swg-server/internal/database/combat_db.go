package database

import (
	"database/sql"
	"fmt"
	"time"

	"swg-server/internal/combat"
	"swg-server/internal/creatures"
)

// GetCombatState reads the full combat-relevant HAM state for a character:
// current AND max for all three pools, wounds, battle fatigue, posture,
// stance, and incapacitation timestamp.
//
// This replaces the incomplete read in GetCharacterByID, which scanned
// health_current and health_max into the SAME Go field (silently losing
// current health) and discarded wounds/BF/posture/stance entirely via
// throwaway scan targets. Found while implementing Phase 3, since that's
// the first phase where current-vs-max actually matters -- Phase 0-2 never
// damaged a character, so the bug was latent and untested until now.
func (db *DB) GetCombatState(characterID string) (*combat.HAMPoolState, error) {
	s := &combat.HAMPoolState{}
	var posture, stance string
	var incapacitatedAt, lastCombatAt sql.NullTime

	err := db.conn.QueryRow(
		`SELECT health_current, health_max, action_current, action_max,
		        mind_current, mind_max, strength, constitution, quickness,
		        stamina, focus, willpower, wounds_health, wounds_action,
		        wounds_mind, battle_fatigue_pct, posture, stance,
		        incapacitated_at, last_combat_action_at
		 FROM ham_pool_states WHERE character_id = ?`,
		characterID,
	).Scan(
		&s.HealthCurrent, &s.HealthMax, &s.ActionCurrent, &s.ActionMax,
		&s.MindCurrent, &s.MindMax, &s.Strength, &s.Constitution, &s.Quickness,
		&s.Stamina, &s.Focus, &s.Willpower, &s.WoundsHealth, &s.WoundsAction,
		&s.WoundsMind, &s.BattleFatiguePct, &posture, &stance,
		&incapacitatedAt, &lastCombatAt,
	)
	if err != nil {
		return nil, fmt.Errorf("get combat state: %w", err)
	}

	s.Posture = combat.Posture(posture)
	s.Stance = combat.Stance(stance)
	if incapacitatedAt.Valid {
		t := incapacitatedAt.Time
		s.IncapacitatedAt = &t
	}
	if lastCombatAt.Valid {
		t := lastCombatAt.Time
		s.LastCombatActionAt = &t
	}
	return s, nil
}

// UpdateCombatState writes back current HAM, wounds, battle fatigue,
// posture, stance, and incapacitation state. Does not touch max pools
// (those are species-derived and fixed at character creation).
func (db *DB) UpdateCombatState(characterID string, s combat.HAMPoolState) error {
	var incapacitatedAt, lastCombatAt interface{}
	if s.IncapacitatedAt != nil {
		incapacitatedAt = *s.IncapacitatedAt
	}
	if s.LastCombatActionAt != nil {
		lastCombatAt = *s.LastCombatActionAt
	}

	_, err := db.conn.Exec(
		`UPDATE ham_pool_states SET
		   health_current = ?, action_current = ?, mind_current = ?,
		   wounds_health = ?, wounds_action = ?, wounds_mind = ?,
		   battle_fatigue_pct = ?, posture = ?, stance = ?,
		   incapacitated_at = ?, last_combat_action_at = ?,
		   updated_at = CURRENT_TIMESTAMP
		 WHERE character_id = ?`,
		s.HealthCurrent, s.ActionCurrent, s.MindCurrent,
		s.WoundsHealth, s.WoundsAction, s.WoundsMind,
		s.BattleFatiguePct, string(s.Posture), string(s.Stance),
		incapacitatedAt, lastCombatAt,
		characterID,
	)
	return err
}

// GetCloneBindPoint returns the planet/position a character respawns at on
// death. GDD 9.6.3: "Players bind to clone facility."
func (db *DB) GetCloneBindPoint(characterID string) (planet string, x, y, z float64, err error) {
	err = db.conn.QueryRow(
		`SELECT clone_bind_planet, clone_bind_x, clone_bind_y, clone_bind_z
		 FROM characters WHERE id = ?`,
		characterID,
	).Scan(&planet, &x, &y, &z)
	return
}

// SetCloneBindPoint updates a character's clone facility binding.
func (db *DB) SetCloneBindPoint(characterID, planet string, x, y, z float64) error {
	_, err := db.conn.Exec(
		`UPDATE characters SET clone_bind_planet = ?, clone_bind_x = ?,
		   clone_bind_y = ?, clone_bind_z = ? WHERE id = ?`,
		planet, x, y, z, characterID,
	)
	return err
}

// --- Lair operations (GDD 17.3, 9.5.3) ---

func (db *DB) CreateLair(l *creatures.Lair) error {
	_, err := db.conn.Exec(
		`INSERT INTO lairs (id, template_id, planet, pos_x, pos_z,
		   lair_hp, lair_hp_max, max_population)
		 VALUES (?, ?, ?, ?, ?, ?, ?, ?)`,
		l.ID, l.TemplateID, l.Planet, l.PosX, l.PosZ,
		l.LairHP, l.LairHPMax, l.MaxPopulation,
	)
	return err
}

func (db *DB) GetActiveLairs() ([]creatures.Lair, error) {
	rows, err := db.conn.Query(
		`SELECT id, template_id, planet, pos_x, pos_z, lair_hp, lair_hp_max,
		        max_population
		 FROM lairs WHERE destroyed_at IS NULL`,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var lairs []creatures.Lair
	for rows.Next() {
		var l creatures.Lair
		if err := rows.Scan(&l.ID, &l.TemplateID, &l.Planet, &l.PosX, &l.PosZ,
			&l.LairHP, &l.LairHPMax, &l.MaxPopulation); err != nil {
			return nil, err
		}
		lairs = append(lairs, l)
	}
	return lairs, nil
}

// --- Creature instance operations (GDD 17.3, 28.3) ---

func (db *DB) CreateCreatureInstance(c *creatures.Instance) error {
	var lairID interface{}
	if c.LairID != "" {
		lairID = c.LairID
	}
	_, err := db.conn.Exec(
		`INSERT INTO creature_instances (id, template_id, lair_id, planet,
		   pos_x, pos_y, pos_z, health_current, health_max, state)
		 VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`,
		c.ID, c.TemplateID, lairID, c.Planet,
		c.PosX, c.PosY, c.PosZ, c.HealthCurrent, c.HealthMax, c.State,
	)
	return err
}

// GetLivingCreatureInstances returns all creature instances not yet dead.
// Loaded at startup and kept in memory thereafter (GDD 29.2.1: real-time
// state is authoritative in-process, persisted periodically -- matching
// the same pattern already used for player position in world.go).
func (db *DB) GetLivingCreatureInstances() ([]creatures.Instance, error) {
	rows, err := db.conn.Query(
		`SELECT id, template_id, COALESCE(lair_id, ''), planet, pos_x, pos_y,
		        pos_z, health_current, health_max, state,
		        COALESCE(target_character_id, '')
		 FROM creature_instances WHERE state != 'dead'`,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var out []creatures.Instance
	for rows.Next() {
		var c creatures.Instance
		if err := rows.Scan(&c.ID, &c.TemplateID, &c.LairID, &c.Planet,
			&c.PosX, &c.PosY, &c.PosZ, &c.HealthCurrent, &c.HealthMax,
			&c.State, &c.TargetCharacterID); err != nil {
			return nil, err
		}
		out = append(out, c)
	}
	return out, nil
}

// UpdateCreatureInstance persists a creature's mutable state (health,
// position, AI state/target). Called on damage and on death, not every
// tick, to avoid excessive write load (same rationale as player position's
// 5-second persistence interval).
func (db *DB) UpdateCreatureInstance(c *creatures.Instance) error {
	var target interface{}
	if c.TargetCharacterID != "" {
		target = c.TargetCharacterID
	}
	var corpseExpires interface{}
	if c.CorpseExpiresAt != nil {
		corpseExpires = *c.CorpseExpiresAt
	}
	_, err := db.conn.Exec(
		`UPDATE creature_instances SET health_current = ?, pos_x = ?,
		   pos_y = ?, pos_z = ?, state = ?, target_character_id = ?,
		   corpse_expires_at = ?
		 WHERE id = ?`,
		c.HealthCurrent, c.PosX, c.PosY, c.PosZ, c.State, target,
		corpseExpires, c.ID,
	)
	return err
}

var _ = time.Now // keep time imported for future use (corpse expiry checks live in the handler)
