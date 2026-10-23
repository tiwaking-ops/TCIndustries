package species

// Species represents one of the 9 playable species in the testbed game.
// Each species applies static bonuses/penalties to the baseline HAM attributes.
// Source: GDD Section 6.2.1 (historical reference; values retained as
// non-canonical placeholder tuning per HD-TST-01 patterns-only constraint).

type SpeciesID string

const (
	Human       SpeciesID = "human"
	SpeciesB    SpeciesID = "species_b"
	SpeciesC    SpeciesID = "species_c"
	SpeciesD    SpeciesID = "species_d"
	SpeciesE    SpeciesID = "species_e"
	SpeciesF    SpeciesID = "species_f"
	SpeciesG    SpeciesID = "species_g"
	SpeciesH    SpeciesID = "species_h"
	SpeciesI    SpeciesID = "species_i"
)

// SpeciesDef holds the display name and attribute modifiers for a species.
type SpeciesDef struct {
	ID          SpeciesID
	DisplayName string
	// Attribute modifiers applied to the Human baseline (GDD 6.2.2)
	HealthMod       int
	StrengthMod     int
	ConstitutionMod int
	ActionMod       int
	QuicknessMod    int
	StaminaMod      int
	MindMod         int
	FocusMod        int
	WillpowerMod    int
	// SpeciesF-specific: cannot use certain armor (cosmetic limitation)
	ArmorRestricted bool
}

// AllSpecies is the authoritative lookup table for all 9 species.
var AllSpecies = map[SpeciesID]SpeciesDef{
	Human: {
		ID:          Human,
		DisplayName: "Human",
	},
	SpeciesB: {
		ID:          SpeciesB,
		DisplayName: "Species B",
		StrengthMod: -10,
		// +10 Agility maps to Quickness in the HAM system
		QuicknessMod: 10,
	},
	SpeciesC: {
		ID:          SpeciesC,
		DisplayName: "Species C",
		// +10 Agility maps to Quickness
		QuicknessMod: 10,
		// -10 Constitution
		ConstitutionMod: -10,
	},
	SpeciesD: {
		ID:          SpeciesD,
		DisplayName: "Species D",
		// +20 Strength
		StrengthMod: 20,
		// -10 Quickness
		QuicknessMod: -10,
	},
	SpeciesE: {
		ID:          SpeciesE,
		DisplayName: "Species E",
		// +10 Charisma maps to Willpower (social/mental attribute)
		WillpowerMod: 10,
		// -10 Stamina
		StaminaMod: -10,
	},
	SpeciesF: {
		ID:          SpeciesF,
		DisplayName: "Species F",
		// +50 Strength
		StrengthMod: 50,
		// -50 Quickness
		QuicknessMod: -50,
		// +100 Health
		HealthMod: 100,
		// Cannot use certain armor
		ArmorRestricted: true,
	},
	SpeciesG: {
		ID:          SpeciesG,
		DisplayName: "Species G",
		// +10 Stamina
		StaminaMod: 10,
		// -10 Willpower
		WillpowerMod: -10,
	},
	SpeciesH: {
		ID:          SpeciesH,
		DisplayName: "Species H",
		// +10 Intelligence maps to Focus
		FocusMod: 10,
		// -10 Constitution
		ConstitutionMod: -10,
	},
	SpeciesI: {
		ID:          SpeciesI,
		DisplayName: "Species I",
		// +10 Agility maps to Quickness
		QuicknessMod: 10,
		// -10 Stamina
		StaminaMod: -10,
	},
}

// Baseline HAM values for a Human (GDD Section 6.2.2)
const (
	BaselineHealth       = 1000
	BaselineStrength     = 500
	BaselineConstitution = 500
	BaselineAction       = 1000
	BaselineQuickness    = 500
	BaselineStamina      = 500
	BaselineMind         = 1000
	BaselineFocus        = 500
	BaselineWillpower    = 500
)

// HAMState holds the computed HAM pool values for a character.
type HAMState struct {
	Health       int `json:"health"`
	Strength     int `json:"strength"`
	Constitution int `json:"constitution"`
	Action       int `json:"action"`
	Quickness    int `json:"quickness"`
	Stamina      int `json:"stamina"`
	Mind         int `json:"mind"`
	Focus        int `json:"focus"`
	Willpower    int `json:"willpower"`
}

// ComputeHAM calculates the starting HAM state for a character of the given species.
// Applies species modifiers to the Human baseline. (GDD 6.2.2)
// Wounds and Battle Fatigue start at zero (no combat has occurred yet).
func ComputeHAM(speciesID SpeciesID) HAMState {
	sp, ok := AllSpecies[speciesID]
	if !ok {
		// Default to Human baseline if unknown
		sp = AllSpecies[Human]
	}

	return HAMState{
		Health:       BaselineHealth + sp.HealthMod,
		Strength:     BaselineStrength + sp.StrengthMod,
		Constitution: BaselineConstitution + sp.ConstitutionMod,
		Action:       BaselineAction + sp.ActionMod,
		Quickness:    BaselineQuickness + sp.QuicknessMod,
		Stamina:      BaselineStamina + sp.StaminaMod,
		Mind:         BaselineMind + sp.MindMod,
		Focus:        BaselineFocus + sp.FocusMod,
		Willpower:    BaselineWillpower + sp.WillpowerMod,
	}
}

// IsValidSpecies checks if a species ID is one of the 9 playable species.
func IsValidSpecies(id SpeciesID) bool {
	_, ok := AllSpecies[id]
	return ok
}

// SpeciesList returns all species for client display (e.g., character creation dropdown).
func SpeciesList() []SpeciesDef {
	list := make([]SpeciesDef, 0, len(AllSpecies))
	for _, sp := range AllSpecies {
		list = append(list, sp)
	}
	return list
}
