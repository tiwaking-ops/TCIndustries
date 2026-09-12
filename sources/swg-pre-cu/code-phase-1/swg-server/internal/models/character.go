package models

import (
	"time"

	"swg-server/internal/species"
)

// Account represents a player account. (GDD Section 27.3)
type Account struct {
	ID           string    `json:"id"`
	Username     string    `json:"username"`
	PasswordHash string    `json:"-"` // never serialized
	CreatedAt    time.Time `json:"created_at"`
}

// CharacterAppearance holds the visual customization data for a character.
// This is the appearance data selected during character creation. (GDD Section 6.3)
type CharacterAppearance struct {
	BodyType  string  `json:"body_type"`
	SkinColor string  `json:"skin_color"`
	HairStyle string  `json:"hair_style"`
	HairColor string  `json:"hair_color"`
	FaceType  string  `json:"face_type"`
	EyeColor  string  `json:"eye_color"`
	Height    float64 `json:"height"` // 0.8 to 1.2 scale factor
}

// Character represents a player character. (GDD Section 6, 27.3)
type Character struct {
	ID         string              `json:"id"`
	AccountID  string              `json:"account_id"`
	Name       string              `json:"name"`
	Species    string              `json:"species"`
	Appearance CharacterAppearance `json:"appearance"`

	// Position in the world (starts at a default spawn point)
	PosX float64 `json:"pos_x"`
	PosY float64 `json:"pos_y"`
	PosZ float64 `json:"pos_z"`
	// Heading/rotation in degrees
	Heading float64 `json:"heading"`

	// Starting planet (Tatooine for MVP — GDD 5.2.1)
	Planet string `json:"planet"`

	CreatedAt time.Time `json:"created_at"`
}

// CharacterWithHAM combines character data with computed HAM state for client display.
type CharacterWithHAM struct {
	Character
	HAM species.HAMState `json:"ham"`
}

// Default spawn position on Tatooine (near Mos Eisley)
const (
	DefaultSpawnX = 3528.0
	DefaultSpawnY = 5.0
	DefaultSpawnZ = -4804.0
	DefaultPlanet = "tatooine"
)
