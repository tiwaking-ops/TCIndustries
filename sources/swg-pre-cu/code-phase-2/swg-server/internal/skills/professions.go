package skills

// Skill system data model for SWG Pre-CU.
// GDD Sections 6.3, 7, 8.2, 28.3.
//
// Every character has 250 skill points. Skills are organized into profession
// trees. Each profession has 4 trees, each tree has 4 tiers (I–IV), plus a
// Novice box (free entry) and a Master box (capstone, requires all 4 trees).
//
// Skill box costs follow the GDD example (Section 6.3.1):
//   Novice: 0 points (free entry into profession)
//   Tier I: 2 points
//   Tier II: 3 points
//   Tier III: 4 points
//   Tier IV: 5 points
//   Master: 10 points
// Total per profession: 0 + 4×(2+3+4+5) + 10 = 66 points

// ProfessionCategory classifies a profession's tier in the progression system.
type ProfessionCategory string

const (
	CategoryBasic  ProfessionCategory = "basic"
	CategoryElite  ProfessionCategory = "elite"
	CategoryHybrid ProfessionCategory = "hybrid"
)

// XPType represents a typed experience point pool.
// GDD Section 6.3.2: XP is typed — combat XP can't buy crafting skills.
type XPType string

const (
	XPCombat      XPType = "combat"
	XPCrafting    XPType = "crafting"
	XPScouting    XPType = "scouting"
	XPMedical     XPType = "medical"
	XPEntertainer XPType = "entertainer"
	XPMerchant    XPType = "merchant"
)

// ProfessionDef is a compact definition used to generate full skill tree data.
type ProfessionDef struct {
	ID       string
	Name     string
	Category ProfessionCategory
	XPType   XPType
	Trees    []TreeDef
}

// TreeDef defines one skill tree within a profession.
type TreeDef struct {
	Name string
	// Abilities granted at each tier (index 0=I, 1=II, 2=III, 3=IV)
	Abilities []string
}

// SkillBox represents a single skill box in a profession tree.
type SkillBox struct {
	ID               string // stable string ID, e.g. "marksman_novice"
	ProfessionID     string
	TreeID           string // nullable for Novice/Master
	Name             string
	Tier             int // 0=Novice, 1-4=tiers, 5=Master
	SkillPointCost   int
	XPCost           int
	XPType           XPType
	CreditCost       int
	GrantedAbilities []string
}

// SkillTree represents one tree within a profession.
type SkillTree struct {
	ID           string
	ProfessionID string
	Name         string
	Boxes        []*SkillBox // ordered: tier 1, 2, 3, 4
}

// Profession represents a full profession with all its trees and boxes.
type Profession struct {
	ID               string
	Name             string
	Category         ProfessionCategory
	XPType           XPType
	Trees            []*SkillTree
	NoviceBox        *SkillBox
	MasterBox        *SkillBox
	TotalSkillPoints int
}

// Tier skill point costs (GDD Section 6.3.1)
var tierCosts = [4]int{2, 3, 4, 5}

const noviceCost = 0
const masterCost = 10

// XP costs scale per tier (GDD 7.3: "early boxes ~1k XP, master boxes ~100k XP")
var tierXPCosts = [4]int{1000, 5000, 15000, 40000}

const masterXPCost = 100000

// Credit costs scale with tier (GDD: 10-50 credits for early, 5000-20000 for master)
var tierCreditCosts = [4]int{50, 200, 1000, 5000}

const masterCreditCost = 10000

// slugify converts a profession+tree name to a snake_case ID.
func slugify(parts ...string) string {
	result := ""
	for i, p := range parts {
		if i > 0 {
			result += "_"
		}
		for _, c := range p {
			if (c >= 'a' && c <= 'z') || (c >= '0' && c <= '9') {
				result += string(c)
			} else if c >= 'A' && c <= 'Z' {
				result += string(c + 32) // to lowercase
			} else if c == ' ' || c == '-' {
				result += "_"
			}
		}
	}
	return result
}

// romanNumeral converts 1-4 to I, II, III, IV.
func romanNumeral(n int) string {
	return []string{"I", "II", "III", "IV"}[n-1]
}

// generateProfession builds a full Profession from a ProfessionDef.
func generateProfession(def ProfessionDef) *Profession {
	p := &Profession{
		ID:       def.ID,
		Name:     def.Name,
		Category: def.Category,
		XPType:   def.XPType,
	}

	// Novice box (profession-level, free entry)
	p.NoviceBox = &SkillBox{
		ID:             slugify(def.ID, "novice"),
		ProfessionID:   def.ID,
		TreeID:         "",
		Name:           "Novice " + def.Name,
		Tier:           0,
		SkillPointCost: noviceCost,
		XPCost:         0,
		XPType:         def.XPType,
		CreditCost:     0,
	}

	totalSP := 0

	// Generate 4 trees, each with 4 tiers
	for _, treeDef := range def.Trees {
		treeID := slugify(def.ID, treeDef.Name)
		tree := &SkillTree{
			ID:           treeID,
			ProfessionID: def.ID,
			Name:         treeDef.Name,
		}

		for i := 0; i < 4; i++ {
			tier := i + 1
			abilities := []string{}
			if i < len(treeDef.Abilities) {
				abilities = treeDef.Abilities
			}
			box := &SkillBox{
				ID:               slugify(def.ID, treeDef.Name, romanNumeral(tier)),
				ProfessionID:     def.ID,
				TreeID:           treeID,
				Name:             treeDef.Name + " " + romanNumeral(tier),
				Tier:             tier,
				SkillPointCost:   tierCosts[i],
				XPCost:           tierXPCosts[i],
				XPType:           def.XPType,
				CreditCost:       tierCreditCosts[i],
				GrantedAbilities: abilities,
			}
			tree.Boxes = append(tree.Boxes, box)
			totalSP += tierCosts[i]
		}

		p.Trees = append(p.Trees, tree)
	}

	// Master box (profession-level, requires all 4 trees complete)
	p.MasterBox = &SkillBox{
		ID:             slugify(def.ID, "master"),
		ProfessionID:   def.ID,
		TreeID:         "",
		Name:           "Master " + def.Name,
		Tier:           5,
		SkillPointCost: masterCost,
		XPCost:         masterXPCost,
		XPType:         def.XPType,
		CreditCost:     masterCreditCost,
	}

	totalSP += masterCost
	p.TotalSkillPoints = totalSP

	return p
}

// AllProfessions returns the 6 basic professions with full skill tree data.
// GDD Section 8.2.
func AllProfessions() []*Profession {
	defs := []ProfessionDef{
		{
			ID:       "artisan",
			Name:     "Artisan",
			Category: CategoryBasic,
			XPType:   XPCrafting,
			Trees: []TreeDef{
				{Name: "Engineering", Abilities: []string{"survey_tools", "basic_components"}},
				{Name: "Business", Abilities: []string{"resource_efficiency", "factory_use"}},
				{Name: "Experimentation", Abilities: []string{"experimentation_bonus"}},
				{Name: "Complexity", Abilities: []string{"advanced_schematics"}},
			},
		},
		{
			ID:       "brawler",
			Name:     "Brawler",
			Category: CategoryBasic,
			XPType:   XPCombat,
			Trees: []TreeDef{
				{Name: "Unarmed Damage", Abilities: []string{"unarmed_strike"}},
				{Name: "Unarmed Speed", Abilities: []string{"unarmed_combo"}},
				{Name: "Melee Support", Abilities: []string{"knockdown", "defensive_stance"}},
				{Name: "Melee Finesse", Abilities: []string{"special_attack", "posture_change"}},
			},
		},
		{
			ID:       "marksman",
			Name:     "Marksman",
			Category: CategoryBasic,
			XPType:   XPCombat,
			Trees: []TreeDef{
				{Name: "Ranged Accuracy", Abilities: []string{"aimed_shot"}},
				{Name: "Ranged Speed", Abilities: []string{"rapid_fire"}},
				{Name: "Ranged Support", Abilities: []string{"suppression_fire", "defensive_posture"}},
				{Name: "Ranged Finesse", Abilities: []string{"critical_shot", "special_ammo"}},
			},
		},
		{
			ID:       "scout",
			Name:     "Scout",
			Category: CategoryBasic,
			XPType:   XPScouting,
			Trees: []TreeDef{
				{Name: "Exploration", Abilities: []string{"terrain_navigation", "movement_speed"}},
				{Name: "Hunting", Abilities: []string{"creature_tracking", "harvest_bonus"}},
				{Name: "Trapping", Abilities: []string{"trap_placement", "creature_knowledge"}},
				{Name: "Survival", Abilities: []string{"camp_placement", "group_buff"}},
			},
		},
		{
			ID:       "medic",
			Name:     "Medic",
			Category: CategoryBasic,
			XPType:   XPMedical,
			Trees: []TreeDef{
				{Name: "Healing", Abilities: []string{"wound_heal"}},
				{Name: "Injury Treatment", Abilities: []string{"revive_player"}},
				{Name: "Medicine Crafting", Abilities: []string{"craft_stimpack", "craft_medical_supply"}},
				{Name: "Support", Abilities: []string{"diagnosis", "disease_cure"}},
			},
		},
		{
			ID:       "entertainer",
			Name:     "Entertainer",
			Category: CategoryBasic,
			XPType:   XPEntertainer,
			Trees: []TreeDef{
				{Name: "Music", Abilities: []string{"instrument_performance", "bf_heal"}},
				{Name: "Dance", Abilities: []string{"dance_performance", "bf_heal"}},
				{Name: "Healing", Abilities: []string{"bf_heal_rate"}},
				{Name: "Showmanship", Abilities: []string{"group_performance", "flourish"}},
			},
		},
	}

	var professions []*Profession
	for _, def := range defs {
		professions = append(professions, generateProfession(def))
	}
	return professions
}

// AllSkillBoxes returns every skill box across all professions as a flat list.
func AllSkillBoxes() []*SkillBox {
	var boxes []*SkillBox
	for _, p := range AllProfessions() {
		boxes = append(boxes, p.NoviceBox)
		for _, tree := range p.Trees {
			boxes = append(boxes, tree.Boxes...)
		}
		boxes = append(boxes, p.MasterBox)
	}
	return boxes
}

// SkillBoxByID returns the skill box with the given ID, or nil if not found.
var boxIndex map[string]*SkillBox

func init() {
	boxIndex = make(map[string]*SkillBox)
	for _, box := range AllSkillBoxes() {
		boxIndex[box.ID] = box
	}
}

// GetSkillBox returns the skill box with the given ID, or nil if not found.
func GetSkillBox(id string) *SkillBox {
	return boxIndex[id]
}

// ProfessionByID returns the profession with the given ID, or nil if not found.
func ProfessionByID(id string) *Profession {
	for _, p := range AllProfessions() {
		if p.ID == id {
			return p
		}
	}
	return nil
}

// PrerequisiteCheck verifies that a character can train a specific skill box.
// Returns (canTrain, reason). GDD Section 7.3, 6.3.3.
//
// Rules:
// - Novice: always trainable (0 XP, 0 credits, 0 skill points)
// - Tier I-IV: must have Novice + previous tier in the same tree
// - Master: must have all 4 trees complete + Novice
func PrerequisiteCheck(boxID string, ownedBoxes map[string]bool) (bool, string) {
	box := GetSkillBox(boxID)
	if box == nil {
		return false, "skill box not found"
	}

	// Novice is always trainable
	if box.Tier == 0 {
		return true, ""
	}

	// Must own the Novice box of this profession
	noviceID := slugify(box.ProfessionID, "novice")
	if !ownedBoxes[noviceID] {
		return false, "must train Novice " + box.ProfessionID + " first"
	}

	// Master box: must have all 4 trees complete
	if box.Tier == 5 {
		prof := ProfessionByID(box.ProfessionID)
		if prof == nil {
			return false, "profession not found"
		}
		for _, tree := range prof.Trees {
			for _, tbox := range tree.Boxes {
				if !ownedBoxes[tbox.ID] {
					return false, "must complete " + tree.Name + " tree before Master"
				}
			}
		}
		return true, ""
	}

	// Tier I-IV: must have previous tier in the same tree
	if box.Tier > 1 {
		prevTier := box.Tier - 1
		prof := ProfessionByID(box.ProfessionID)
		for _, tree := range prof.Trees {
			if tree.ID == box.TreeID {
				if prevTier <= len(tree.Boxes) {
					prevBox := tree.Boxes[prevTier-1]
					if !ownedBoxes[prevBox.ID] {
						return false, "must train " + prevBox.Name + " first"
					}
				}
				break
			}
		}
	}

	return true, ""
}

// DependentBoxes returns skill boxes that must be dropped if the given box is dropped.
// GDD Section 7.4.1: "Dropping a skill drops all dependent skills."
func DependentBoxes(boxID string, ownedBoxes map[string]bool) []string {
	box := GetSkillBox(boxID)
	if box == nil {
		return nil
	}

	var dependents []string

	// If dropping Novice, all boxes in the profession must be dropped
	if box.Tier == 0 {
		for _, b := range AllSkillBoxes() {
			if b.ProfessionID == box.ProfessionID && b.ID != box.ID && ownedBoxes[b.ID] {
				dependents = append(dependents, b.ID)
			}
		}
		return dependents
	}

	// If dropping a tier box, all higher tiers in the same tree must be dropped
	if box.Tier >= 1 && box.Tier <= 4 {
		prof := ProfessionByID(box.ProfessionID)
		for _, tree := range prof.Trees {
			if tree.ID == box.TreeID {
				for _, tbox := range tree.Boxes {
					if tbox.Tier > box.Tier && ownedBoxes[tbox.ID] {
						dependents = append(dependents, tbox.ID)
					}
				}
				break
			}
		}
		// Also drop Master if we're dropping a tree box (since Master requires all trees)
		if ownedBoxes[slugify(box.ProfessionID, "master")] {
			dependents = append(dependents, slugify(box.ProfessionID, "master"))
		}
	}

	return dependents
}

// TotalSkillPointsUsed computes the total skill points spent on owned boxes.
func TotalSkillPointsUsed(ownedBoxes map[string]bool) int {
	total := 0
	for boxID := range ownedBoxes {
		if box := GetSkillBox(boxID); box != nil {
			total += box.SkillPointCost
		}
	}
	return total
}

// MaxSkillPoints is the total skill points every character has. GDD 6.3.1.
const MaxSkillPoints = 250
