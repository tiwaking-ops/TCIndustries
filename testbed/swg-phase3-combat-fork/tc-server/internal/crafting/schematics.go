package crafting

// ResourceRequirement defines what resource category and amount a schematic needs.
type ResourceRequirement struct {
	Category string `json:"category"` // e.g. "ferric_metal"
	Amount   int    `json:"amount"`   // quantity required
}

// Schematic represents a crafting recipe (GDD 10.3).
type Schematic struct {
	ID           string                `json:"id"`
	Name         string                `json:"name"`
	Category     string                `json:"category"`     // "weapon", "armor", "structure"
	Complexity   int                   `json:"complexity"`   // skill complexity requirement
	Requirements []ResourceRequirement `json:"requirements"` // resource slots
	ExpSlots     []string              `json:"exp_slots"`    // experimentation attribute lanes
	ResultItem   string                `json:"result_item"`  // item template produced
}

// GetDefaultSchematics returns the baseline generic schematics for Phase 4.
func GetDefaultSchematics() map[string]Schematic {
	return map[string]Schematic{
		"basic_sidearm": {
			ID:       "basic_sidearm",
			Name:     "Basic Sidearm",
			Category: "weapon",
			Complexity: 5,
			Requirements: []ResourceRequirement{
				{Category: "ferric_metal", Amount: 50},
				{Category: "conductive_alloy", Amount: 20},
			},
			ExpSlots:   []string{"damage", "accuracy", "speed"},
			ResultItem: "weapon_sidearm_basic",
		},
		"basic_plate": {
			ID:       "basic_plate",
			Name:     "Basic Armor Plate",
			Category: "armor",
			Complexity: 5,
			Requirements: []ResourceRequirement{
				{Category: "structural_polymer", Amount: 60},
				{Category: "industrial_chemical", Amount: 20},
			},
			ExpSlots:   []string{"mitigation", "encumbrance"},
			ResultItem: "armor_plate_basic",
		},
		"harvester_deed": {
			ID:       "harvester_deed",
			Name:     "Personal Harvester Deed",
			Category: "structure",
			Complexity: 10,
			Requirements: []ResourceRequirement{
				{Category: "structural_polymer", Amount: 100},
				{Category: "ferric_metal", Amount: 150},
			},
			ExpSlots:   []string{"hopper_capacity", "extraction_rate"},
			ResultItem: "deed_harvester_personal",
		},
	}
}
