package species

// Species represents one of the 9 playable species in the base game.
// OWNER DECISION 2026-09-14 (C2): species differentiation is REMOVED in this
// testbed fork — every species computes to the uniform human baseline (all pools
// and all six attributes = 1000). The species table is retained as an ID registry
// (character creation + API compatibility); all modifiers are zero.
// NOTE: species IDs/names below are SW-derived and REMAIN in this fork — recorded as
// known-remaining in HYGIENE_NOTE.md (full species rename deferred to review gate).

type SpeciesID string

const (
	Human       SpeciesID = "human"
	Bothan      SpeciesID = "bothan"
	Rodian      SpeciesID = "rodian"
	Trandoshan  SpeciesID = "trandoshan"
	Twilek      SpeciesID = "twilek"
	Wookiee     SpeciesID = "wookiee"
	Zabrak      SpeciesID = "zabrak"
	MonCalamari SpeciesID = "mon_calamari"
	Sullustan   SpeciesID = "sullustan"
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
	// Wookiee-specific: cannot use certain armor (cosmetic limitation)
	ArmorRestricted bool
}

// AllSpecies is the authoritative lookup table for all 9 species.
// Owner decision 2026-09-14: all modifiers zeroed — uniform 1000s for every species.
var AllSpecies = map[SpeciesID]SpeciesDef{
	Human: {
		ID:          Human,
		DisplayName: "Human",
	},
	Bothan: {
		ID:          Bothan,
		DisplayName: "Bothan",
	},
	Rodian: {
		ID:          Rodian,
		DisplayName: "Rodian",
	},
	Trandoshan: {
		ID:          Trandoshan,
		DisplayName: "Trandoshan",
	},
	Twilek: {
		ID:          Twilek,
		DisplayName: "Twi'lek",
	},
	Wookiee: {
		ID:          Wookiee,
		DisplayName: "Wookiee",
	},
	Zabrak: {
		ID:          Zabrak,
		DisplayName: "Zabrak",
	},
	MonCalamari: {
		ID:          MonCalamari,
		DisplayName: "Mon Calamari",
	},
	Sullustan: {
		ID:          Sullustan,
		DisplayName: "Sullustan",
	},
}

// Baseline HAM values — owner decision 2026-09-14 (C2): all nine pools/attributes
// start at 1000 (the human starting pool value), uniformly, for every species.
const (
	BaselineHealth       = 1000
	BaselineStrength     = 1000
	BaselineConstitution = 1000
	BaselineAction       = 1000
	BaselineQuickness    = 1000
	BaselineStamina      = 1000
	BaselineMind         = 1000
	BaselineFocus        = 1000
	BaselineWillpower    = 1000
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
