package database

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/mattn/go-sqlite3"

	"swg-server/internal/models"
	"swg-server/internal/species"
)

// DB wraps the database connection and provides data access methods.
type DB struct {
	conn *sql.DB
}

// New creates a new database connection. Uses SQLite for development.
// The SQL is compatible with PostgreSQL — to switch, replace the driver and
// connection string in New(). (GDD Section 29.4)
func New(dbPath string) (*DB, error) {
	if dbPath == "" {
		dbPath = "swg.db"
	}

	conn, err := sql.Open("sqlite3", dbPath+"?_foreign_keys=on")
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	// Enable WAL mode for better concurrent read performance
	if _, err := conn.Exec("PRAGMA journal_mode=WAL"); err != nil {
		return nil, fmt.Errorf("failed to set WAL mode: %w", err)
	}

	db := &DB{conn: conn}

	if err := db.migrate(); err != nil {
		return nil, fmt.Errorf("failed to run migrations: %w", err)
	}

	return db, nil
}

// migrate runs all database migrations in order.
func (db *DB) migrate() error {
	migrations := []string{
		// 001: accounts table (GDD 27.3)
		`CREATE TABLE IF NOT EXISTS accounts (
			id TEXT PRIMARY KEY,
			username TEXT UNIQUE NOT NULL,
			password_hash TEXT NOT NULL,
			created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
		)`,

		// 002: characters table (GDD 6, 27.3)
		`CREATE TABLE IF NOT EXISTS characters (
			id TEXT PRIMARY KEY,
			account_id TEXT NOT NULL REFERENCES accounts(id) ON DELETE CASCADE,
			name TEXT UNIQUE NOT NULL,
			species TEXT NOT NULL,
			body_type TEXT DEFAULT 'average',
			skin_color TEXT DEFAULT 'default',
			hair_style TEXT DEFAULT 'none',
			hair_color TEXT DEFAULT 'none',
			face_type TEXT DEFAULT 'default',
			eye_color TEXT DEFAULT 'brown',
			height REAL DEFAULT 1.0,
			pos_x REAL DEFAULT 3528.0,
			pos_y REAL DEFAULT 5.0,
			pos_z REAL DEFAULT -4804.0,
			heading REAL DEFAULT 0.0,
			planet TEXT DEFAULT 'tatooine',
			credits INTEGER NOT NULL DEFAULT 5000,
			created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
		)`,

		// 003: ham_pool_states table (GDD 6.2.2, 28.4)
		// Stores the persistent HAM state per character.
		// Wounds and Battle Fatigue start at zero.
		`CREATE TABLE IF NOT EXISTS ham_pool_states (
			character_id TEXT PRIMARY KEY REFERENCES characters(id) ON DELETE CASCADE,
			health_current INTEGER NOT NULL,
			health_max INTEGER NOT NULL,
			action_current INTEGER NOT NULL,
			action_max INTEGER NOT NULL,
			mind_current INTEGER NOT NULL,
			mind_max INTEGER NOT NULL,
			strength INTEGER NOT NULL,
			constitution INTEGER NOT NULL,
			quickness INTEGER NOT NULL,
			stamina INTEGER NOT NULL,
			focus INTEGER NOT NULL,
			willpower INTEGER NOT NULL,
			wounds_health INTEGER NOT NULL DEFAULT 0,
			wounds_action INTEGER NOT NULL DEFAULT 0,
			wounds_mind INTEGER NOT NULL DEFAULT 0,
			battle_fatigue_pct REAL NOT NULL DEFAULT 0.0,
			posture TEXT NOT NULL DEFAULT 'standing',
			stance TEXT NOT NULL DEFAULT 'normal',
			incapacitated_at DATETIME,
			last_combat_action_at DATETIME,
			updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
		)`,

		// Index for looking up characters by account
		`CREATE INDEX IF NOT EXISTS idx_characters_account_id ON characters(account_id)`,

		// 005: character_skills table (SkillBoxOwnership, GDD 28.3)
		`CREATE TABLE IF NOT EXISTS character_skills (
				character_id TEXT NOT NULL REFERENCES characters(id) ON DELETE CASCADE,
				skill_box_id TEXT NOT NULL,
				acquired_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
				PRIMARY KEY (character_id, skill_box_id)
			)`,

		// 006: character_xp table (typed XP pools, GDD 6.3.2)
		`CREATE TABLE IF NOT EXISTS character_xp (
				character_id TEXT NOT NULL REFERENCES characters(id) ON DELETE CASCADE,
				xp_type TEXT NOT NULL,
				amount INTEGER NOT NULL DEFAULT 0,
				PRIMARY KEY (character_id, xp_type)
			)`,

		// 007: trainers table (GDD 7.3)
		`CREATE TABLE IF NOT EXISTS trainers (
				id TEXT PRIMARY KEY,
				profession_id TEXT NOT NULL,
				name TEXT NOT NULL,
				planet TEXT NOT NULL,
				pos_x REAL NOT NULL,
				pos_z REAL NOT NULL,
				city TEXT NOT NULL
			)`,
	}

	for i, m := range migrations {
		if _, err := db.conn.Exec(m); err != nil {
			return fmt.Errorf("migration %d failed: %w", i+1, err)
		}
	}

	return nil
}

// Close closes the database connection.
func (db *DB) Close() error {
	return db.conn.Close()
}

// --- Account operations ---

// CreateAccount inserts a new account into the database.
func (db *DB) CreateAccount(id, username, passwordHash string) error {
	_, err := db.conn.Exec(
		`INSERT INTO accounts (id, username, password_hash) VALUES (?, ?, ?)`,
		id, username, passwordHash,
	)
	return err
}

// GetAccountByUsername retrieves an account by username.
func (db *DB) GetAccountByUsername(username string) (*models.Account, error) {
	a := &models.Account{}
	err := db.conn.QueryRow(
		`SELECT id, username, password_hash, created_at FROM accounts WHERE username = ?`,
		username,
	).Scan(&a.ID, &a.Username, &a.PasswordHash, &a.CreatedAt)
	if err != nil {
		return nil, err
	}
	return a, nil
}

// GetAccountByID retrieves an account by ID.
func (db *DB) GetAccountByID(id string) (*models.Account, error) {
	a := &models.Account{}
	err := db.conn.QueryRow(
		`SELECT id, username, password_hash, created_at FROM accounts WHERE id = ?`,
		id,
	).Scan(&a.ID, &a.Username, &a.PasswordHash, &a.CreatedAt)
	if err != nil {
		return nil, err
	}
	return a, nil
}

// --- Character operations ---

// CreateCharacter inserts a new character and its HAM state in a single transaction.
// This enforces the atomicity requirement from GDD Section 29.5.
func (db *DB) CreateCharacter(char *models.Character) error {
	tx, err := db.conn.Begin()
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	// Insert character
	_, err = tx.Exec(
		`INSERT INTO characters (id, account_id, name, species, body_type, skin_color,
			hair_style, hair_color, face_type, eye_color, height,
			pos_x, pos_y, pos_z, heading, planet)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`,
		char.ID, char.AccountID, char.Name, char.Species,
		char.Appearance.BodyType, char.Appearance.SkinColor,
		char.Appearance.HairStyle, char.Appearance.HairColor,
		char.Appearance.FaceType, char.Appearance.EyeColor, char.Appearance.Height,
		char.PosX, char.PosY, char.PosZ, char.Heading, char.Planet,
	)
	if err != nil {
		return fmt.Errorf("failed to insert character: %w", err)
	}

	// Compute and insert HAM state from species (GDD 6.2.2)
	ham := species.ComputeHAM(species.SpeciesID(char.Species))

	_, err = tx.Exec(
		`INSERT INTO ham_pool_states (
			character_id, health_current, health_max,
			action_current, action_max, mind_current, mind_max,
			strength, constitution, quickness, stamina, focus, willpower
		) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`,
		char.ID,
		ham.Health, ham.Health,
		ham.Action, ham.Action,
		ham.Mind, ham.Mind,
		ham.Strength, ham.Constitution,
		ham.Quickness, ham.Stamina,
		ham.Focus, ham.Willpower,
	)
	if err != nil {
		return fmt.Errorf("failed to insert HAM state: %w", err)
	}

	return tx.Commit()
}

// GetCharactersByAccount retrieves all characters for an account.
func (db *DB) GetCharactersByAccount(accountID string) ([]models.Character, error) {
	rows, err := db.conn.Query(
		`SELECT id, account_id, name, species, body_type, skin_color,
			hair_style, hair_color, face_type, eye_color, height,
			pos_x, pos_y, pos_z, heading, planet, credits, created_at
		FROM characters WHERE account_id = ? ORDER BY created_at`,
		accountID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var chars []models.Character
	for rows.Next() {
		var c models.Character
		err := rows.Scan(
			&c.ID, &c.AccountID, &c.Name, &c.Species,
			&c.Appearance.BodyType, &c.Appearance.SkinColor,
			&c.Appearance.HairStyle, &c.Appearance.HairColor,
			&c.Appearance.FaceType, &c.Appearance.EyeColor, &c.Appearance.Height,
			&c.PosX, &c.PosY, &c.PosZ, &c.Heading, &c.Planet, &c.Credits, &c.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		chars = append(chars, c)
	}
	return chars, nil
}

// GetCharacterByID retrieves a single character by ID, including its HAM state.
func (db *DB) GetCharacterByID(id string) (*models.CharacterWithHAM, error) {
	c := &models.CharacterWithHAM{}

	err := db.conn.QueryRow(
		`SELECT c.id, c.account_id, c.name, c.species,
			c.body_type, c.skin_color, c.hair_style, c.hair_color,
			c.face_type, c.eye_color, c.height,
			c.pos_x, c.pos_y, c.pos_z, c.heading, c.planet, c.credits, c.created_at,
			h.health_current, h.health_max,
			h.action_current, h.action_max,
			h.mind_current, h.mind_max,
			h.strength, h.constitution, h.quickness, h.stamina, h.focus, h.willpower,
			h.wounds_health, h.wounds_action, h.wounds_mind,
			h.battle_fatigue_pct, h.posture, h.stance
		FROM characters c
		JOIN ham_pool_states h ON c.id = h.character_id
		WHERE c.id = ?`,
		id,
	).Scan(
		&c.ID, &c.AccountID, &c.Name, &c.Species,
		&c.Appearance.BodyType, &c.Appearance.SkinColor,
		&c.Appearance.HairStyle, &c.Appearance.HairColor,
		&c.Appearance.FaceType, &c.Appearance.EyeColor, &c.Appearance.Height,
		&c.PosX, &c.PosY, &c.PosZ, &c.Heading, &c.Planet, &c.Credits, &c.CreatedAt,
		&c.HAM.Health, &c.HAM.Health,
		&c.HAM.Action, &c.HAM.Action,
		&c.HAM.Mind, &c.HAM.Mind,
		&c.HAM.Strength, &c.HAM.Constitution,
		&c.HAM.Quickness, &c.HAM.Stamina,
		&c.HAM.Focus, &c.HAM.Willpower,
		// wounds and BF are stored but not yet in the HAMState struct for Phase 0
		// They'll be needed in Phase 3 (Combat) and Phase 6 (Entertainer/Medic)
		new(int), new(int), new(int), new(float64), new(string), new(string),
	)
	if err != nil {
		return nil, err
	}

	// Recompute HAM max values from species to ensure consistency
	// (current values may differ from max once wounds/BF are implemented)
	ham := species.ComputeHAM(species.SpeciesID(c.Species))
	c.HAM = ham

	return c, nil
}

// GetCharacterNameExists checks if a character name is already taken.
func (db *DB) GetCharacterNameExists(name string) (bool, error) {
	var count int
	err := db.conn.QueryRow(
		`SELECT COUNT(*) FROM characters WHERE name = ?`,
		name,
	).Scan(&count)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

// Silence unused import warning for os (used in production builds for env vars)
var _ = os.Getenv

// UpdateCharacterPosition persists a character's position to the database.
// Called periodically (every 5-10s) and on disconnect. (GDD 29.4 persistence)
func (db *DB) UpdateCharacterPosition(characterID string, x, y, z, heading float64, planet string) error {
	_, err := db.conn.Exec(
		`UPDATE characters SET pos_x = ?, pos_y = ?, pos_z = ?, heading = ?, planet = ? WHERE id = ?`,
		x, y, z, heading, planet, characterID,
	)
	return err
}
