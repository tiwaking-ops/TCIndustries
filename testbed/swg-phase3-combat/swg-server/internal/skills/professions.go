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
// Phase 9: one pool per elite XP type (GDD §§8.3/7.2). Doctor/Combat Medic reuse
// medical; Musician/Dancer reuse entertaining (GDD-literal pool sharing).
type XPType string

const (
	XPCombat      XPType = "combat"
	XPCrafting    XPType = "crafting"
	XPScouting    XPType = "scouting"
	XPMedical     XPType = "medical"
	XPEntertainer XPType = "entertainer"
	XPMerchant    XPType = "merchant"

	// Phase 9 elite pools (free-form strings work in character_xp already;
	// constants formalize them for training validation).
	XPArmorCrafting    XPType = "armor_crafting"
	XPWeaponCrafting   XPType = "weapon_crafting"
	XPStructureCrafting XPType = "structure_crafting"
	XPBioEngineering   XPType = "bio_engineering"
	XPCreatureHandling XPType = "creature_handling"
	XPRanger           XPType = "ranger"
	XPPistolCombat     XPType = "pistol_combat"
	XPRifleCombat      XPType = "rifle_combat"
	XPCarbineCombat    XPType = "carbine_combat"
	XPFencingCombat    XPType = "fencing_combat"
	XPSwordCombat      XPType = "sword_combat"
	XPPolearmCombat    XPType = "polearm_combat"
	XPTerasKasi        XPType = "teras_kasi"
	XPImageDesigner    XPType = "image_designer"
	XPBountyHunter     XPType = "bounty_hunter"
	XPCommando         XPType = "commando"
	XPSmuggler         XPType = "smuggler"
)

// AllXPTypes lists every pool the earn-xp placeholder and training accept.
func AllXPTypes() []string {
	return []string{
		"combat", "crafting", "scouting", "medical", "entertainer", "merchant",
		"mentoring",
		"armor_crafting", "weapon_crafting", "structure_crafting",
		"bio_engineering", "creature_handling", "ranger",
		"pistol_combat", "rifle_combat", "carbine_combat",
		"fencing_combat", "sword_combat", "polearm_combat",
		"teras_kasi", "image_designer", "bounty_hunter", "commando", "smuggler",
	}
}

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

// Tier skill point costs (GDD Section 6.3.1 for basics).
var tierCosts = [4]int{2, 3, 4, 5}

const noviceCost = 0
const masterCost = 10

// CostTable parameterizes the SP ladder per profession class (Phase 9).
// Basics: 0 + 4×(2+3+4+5) + 10 = 66 (verified ladder, preserved).
// Standard 80-class elites: 0 + 4×(3+4+5+6) + 8 = 80 [PROVISIONAL distribution].
// 120-class elites (BH/Commando/Smuggler/TKA): 0 + 4×(4+6+8+10) + 8 = 120
// [PROVISIONAL distribution; TKA 120 follows §33.3 over §8.3.15's 80].
type CostTable struct {
	Tiers  [4]int
	Master int
}

// Standard80 totals 80; Elite120 totals 120.
var Standard80 = CostTable{Tiers: [4]int{3, 4, 5, 6}, Master: 8}
var Elite120 = CostTable{Tiers: [4]int{4, 6, 8, 10}, Master: 8}

// costTableFor returns the ladder for a profession (basics keep 66).
func costTableFor(category ProfessionCategory, id string) CostTable {
	if category != CategoryElite {
		return CostTable{Tiers: tierCosts, Master: masterCost}
	}
	switch id {
	case "bountyhunter", "commando", "smuggler", "teraskasi":
		return Elite120
	default:
		return Standard80
	}
}

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
	costs := costTableFor(def.Category, def.ID)

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
				SkillPointCost:   costs.Tiers[i],
				XPCost:           tierXPCosts[i],
				XPType:           def.XPType,
				CreditCost:       tierCreditCosts[i],
				GrantedAbilities: abilities,
			}
			tree.Boxes = append(tree.Boxes, box)
			totalSP += costs.Tiers[i]
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
		SkillPointCost: costs.Master,
		XPCost:         masterXPCost,
		XPType:         def.XPType,
		CreditCost:     masterCreditCost,
	}

	totalSP += costs.Master
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

	// Phase 9 elite roster (GDD §8.3). Profession IDs are single lowercase
	// words (slugify drops underscores); tree names are GDD-literal. XP and
	// credit ladders are reused from basics (flagged simplification — SP cost
	// is the differentiator). Droid Engineer is OUT (31.5 EXPANSION);
	// Chef/Tailor/Miner have no 8.3 entries (unspecifiable — OUT with reason).
	defs = append(defs, []ProfessionDef{
		{
			ID: "armorsmith", Name: "Armorsmith", Category: CategoryElite, XPType: XPArmorCrafting,
			Trees: []TreeDef{
				{Name: "Personal Armor", Abilities: []string{"craft_armor"}},
				{Name: "Armor Assembly", Abilities: []string{"armor_assembly"}},
				{Name: "Armor Experimentation", Abilities: []string{"armor_experimentation"}},
				{Name: "Armor Customization", Abilities: []string{"armor_customization"}},
			},
		},
		{
			ID: "weaponsmith", Name: "Weaponsmith", Category: CategoryElite, XPType: XPWeaponCrafting,
			Trees: []TreeDef{
				{Name: "Melee Weapons", Abilities: []string{"craft_melee"}},
				{Name: "Ranged Weapons", Abilities: []string{"craft_ranged"}},
				{Name: "Weapon Experimentation", Abilities: []string{"weapon_experimentation"}},
				{Name: "Weapon Assembly", Abilities: []string{"weapon_assembly"}},
			},
		},
		{
			ID: "architect", Name: "Architect", Category: CategoryElite, XPType: XPStructureCrafting,
			Trees: []TreeDef{
				{Name: "Structures", Abilities: []string{"craft_structures"}},
				{Name: "Furniture", Abilities: []string{"craft_furniture"}},
				{Name: "Installation", Abilities: []string{"craft_installations"}},
				{Name: "City Planning", Abilities: []string{"city_planning"}},
			},
		},
		{
			ID: "merchant", Name: "Merchant", Category: CategoryElite, XPType: XPMerchant,
			Trees: []TreeDef{
				{Name: "Vendor Management", Abilities: []string{"vendor_discount", "vendor_slots"}},
				{Name: "Advertising", Abilities: []string{"market_presence"}},
				{Name: "Wholesale", Abilities: []string{"bulk_trade"}},
				{Name: "Appraisal", Abilities: []string{"appraisal"}},
			},
		},
		{
			ID: "bioengineer", Name: "Bio-Engineer", Category: CategoryElite, XPType: XPBioEngineering,
			Trees: []TreeDef{
				{Name: "DNA Sampling", Abilities: []string{"dna_sample"}},
				{Name: "Creature Engineering", Abilities: []string{"dna_combine"}},
				{Name: "Cloning", Abilities: []string{"clone_pet"}},
				{Name: "Tissue Engineering", Abilities: []string{"tissue_sample"}},
			},
		},
		{
			ID: "creaturehandler", Name: "Creature Handler", Category: CategoryElite, XPType: XPCreatureHandling,
			Trees: []TreeDef{
				{Name: "Taming", Abilities: []string{"tame_creature"}},
				{Name: "Training", Abilities: []string{"train_pet"}},
				{Name: "Command", Abilities: []string{"command_pet"}},
				{Name: "Beast Mastery", Abilities: []string{"pet_mastery"}},
			},
		},
		{
			ID: "ranger", Name: "Ranger", Category: CategoryElite, XPType: XPRanger,
			Trees: []TreeDef{
				{Name: "Tracking", Abilities: []string{"track_creature"}},
				{Name: "Trapping", Abilities: []string{"trap_lore"}},
				{Name: "Camouflage", Abilities: []string{"wilderness_stealth"}},
				{Name: "Wilderness Survival", Abilities: []string{"camp_mastery", "survey_bonus"}},
			},
		},
		{
			ID: "pistoleer", Name: "Pistoleer", Category: CategoryElite, XPType: XPPistolCombat,
			Trees: []TreeDef{
				{Name: "Pistol Accuracy", Abilities: []string{"pistol_accuracy"}},
				{Name: "Pistol Speed", Abilities: []string{"pistol_speed"}},
				{Name: "Pistol Support", Abilities: []string{"pistol_support"}},
				{Name: "Pistol Finesse", Abilities: []string{"pistol_finesse"}},
			},
		},
		{
			ID: "rifleman", Name: "Rifleman", Category: CategoryElite, XPType: XPRifleCombat,
			Trees: []TreeDef{
				{Name: "Rifle Accuracy", Abilities: []string{"rifle_accuracy"}},
				{Name: "Rifle Speed", Abilities: []string{"rifle_speed"}},
				{Name: "Rifle Support", Abilities: []string{"rifle_support"}},
				{Name: "Rifle Finesse", Abilities: []string{"rifle_finesse"}},
			},
		},
		{
			ID: "carbineer", Name: "Carbineer", Category: CategoryElite, XPType: XPCarbineCombat,
			Trees: []TreeDef{
				{Name: "Carbine Accuracy", Abilities: []string{"carbine_accuracy"}},
				{Name: "Carbine Speed", Abilities: []string{"carbine_speed"}},
				{Name: "Carbine Support", Abilities: []string{"carbine_support"}},
				{Name: "Carbine Finesse", Abilities: []string{"carbine_finesse"}},
			},
		},
		{
			ID: "fencer", Name: "Fencer", Category: CategoryElite, XPType: XPFencingCombat,
			Trees: []TreeDef{
				{Name: "Fencing Accuracy", Abilities: []string{"fencing_accuracy"}},
				{Name: "Fencing Speed", Abilities: []string{"fencing_speed"}},
				{Name: "Fencing Defense", Abilities: []string{"fencing_defense"}},
				{Name: "Fencing Finesse", Abilities: []string{"fencing_finesse"}},
			},
		},
		{
			ID: "swordsman", Name: "Swordsman", Category: CategoryElite, XPType: XPSwordCombat,
			Trees: []TreeDef{
				{Name: "Sword Accuracy", Abilities: []string{"sword_accuracy"}},
				{Name: "Sword Speed", Abilities: []string{"sword_speed"}},
				{Name: "Sword Defense", Abilities: []string{"sword_defense"}},
				{Name: "Sword Finesse", Abilities: []string{"sword_finesse"}},
			},
		},
		{
			ID: "pikeman", Name: "Pikeman", Category: CategoryElite, XPType: XPPolearmCombat,
			Trees: []TreeDef{
				{Name: "Polearm Accuracy", Abilities: []string{"polearm_accuracy"}},
				{Name: "Polearm Speed", Abilities: []string{"polearm_speed"}},
				{Name: "Polearm Support", Abilities: []string{"polearm_support"}},
				{Name: "Polearm Finesse", Abilities: []string{"polearm_finesse"}},
			},
		},
		{
			ID: "teraskasi", Name: "Teras Kasi Artist", Category: CategoryElite, XPType: XPTerasKasi,
			Trees: []TreeDef{
				{Name: "Meditative Techniques", Abilities: []string{"meditation"}},
				{Name: "Power Strikes", Abilities: []string{"power_strike"}},
				{Name: "Defensive Techniques", Abilities: []string{"tka_defense"}},
				{Name: "Kata Mastery", Abilities: []string{"kata_mastery"}},
			},
		},
		{
			ID: "doctor", Name: "Doctor", Category: CategoryElite, XPType: XPMedical,
			Trees: []TreeDef{
				{Name: "Advanced Healing", Abilities: []string{"advanced_heal"}},
				{Name: "Enhancement", Abilities: []string{"craft_buff_pack"}},
				{Name: "Medicine Mastery", Abilities: []string{"medicine_mastery"}},
				{Name: "Diagnosis", Abilities: []string{"diagnosis"}},
			},
		},
		{
			ID: "combatmedic", Name: "Combat Medic", Category: CategoryElite, XPType: XPMedical,
			Trees: []TreeDef{
				{Name: "Field Triage", Abilities: []string{"field_triage"}},
				{Name: "Combat Medicine", Abilities: []string{"combat_heal"}},
				{Name: "Revivification", Abilities: []string{"swift_revive"}},
				{Name: "Battlefield Awareness", Abilities: []string{"battlefield_awareness"}},
			},
		},
		{
			ID: "musician", Name: "Musician", Category: CategoryElite, XPType: XPEntertainer,
			Trees: []TreeDef{
				{Name: "Instrumentation", Abilities: []string{"advanced_instruments"}},
				{Name: "Composition", Abilities: []string{"composition"}},
				{Name: "Inspiration", Abilities: []string{"performance_buff"}},
				{Name: "Showmanship", Abilities: []string{"stage_showmanship"}},
			},
		},
		{
			ID: "dancer", Name: "Dancer", Category: CategoryElite, XPType: XPEntertainer,
			Trees: []TreeDef{
				{Name: "Dance Mastery", Abilities: []string{"advanced_dances"}},
				{Name: "Choreography", Abilities: []string{"choreography"}},
				{Name: "Exotic Dances", Abilities: []string{"exotic_dance"}},
				{Name: "Stage Presence", Abilities: []string{"stage_presence"}},
			},
		},
		{
			ID: "imagedesigner", Name: "Image Designer", Category: CategoryElite, XPType: XPImageDesigner,
			Trees: []TreeDef{
				{Name: "Holoemote Design", Abilities: []string{"design_holoemote"}},
				{Name: "Body Modification", Abilities: []string{"modify_body"}},
				{Name: "Facial Modification", Abilities: []string{"modify_face"}},
				{Name: "Hair Styling", Abilities: []string{"style_hair"}},
			},
		},
		{
			ID: "bountyhunter", Name: "Bounty Hunter", Category: CategoryElite, XPType: XPBountyHunter,
			Trees: []TreeDef{
				{Name: "Investigation", Abilities: []string{"investigate_target"}},
				{Name: "Pursuit", Abilities: []string{"pursue_target"}},
				{Name: "Weapon Mastery", Abilities: []string{"bounty_arsenal"}},
				{Name: "Capture", Abilities: []string{"capture_target"}},
			},
		},
		{
			ID: "commando", Name: "Commando", Category: CategoryElite, XPType: XPCommando,
			Trees: []TreeDef{
				{Name: "Heavy Weapons", Abilities: []string{"heavy_weapons"}},
				{Name: "Demolitions", Abilities: []string{"demolitions"}},
				{Name: "Weapon Engineering", Abilities: []string{"heavy_engineering"}},
				{Name: "Tactical Assault", Abilities: []string{"tactical_assault"}},
			},
		},
		{
			ID: "smuggler", Name: "Smuggler", Category: CategoryElite, XPType: XPSmuggler,
			Trees: []TreeDef{
				{Name: "Spice Production", Abilities: []string{"craft_spice"}},
				{Name: "Slicing", Abilities: []string{"slice_item"}},
				{Name: "Underworld Contacts", Abilities: []string{"underworld_contacts"}},
				{Name: "Evasion", Abilities: []string{"evasion"}},
			},
		},
	}...)

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

// ElitePrereqs maps an elite profession to its Novice-entry requirements
// (GDD §8.3 prereqs; box IDs). Multi-master professions use the special cases
// below, not this list.
var ElitePrereqs = map[string][]string{
	"armorsmith":      {"artisan_master"},
	"weaponsmith":     {"artisan_master"},
	"architect":       {"artisan_master"},
	"merchant":        {"artisan_master"},
	"bioengineer":     {"scout_master", "medic_novice"},
	"creaturehandler": {"scout_master"},
	"ranger":          {"scout_master"},
	"pistoleer":       {"marksman_master"},
	"rifleman":        {"marksman_master"},
	"carbineer":       {"marksman_master"},
	"fencer":          {"brawler_master"},
	"swordsman":       {"brawler_master"},
	"pikeman":         {"brawler_master"},
	"doctor":          {"medic_master"},
	"combatmedic":     {"medic_master"},
	"musician":        {"entertainer_master"},
	"dancer":          {"entertainer_master"},
	"imagedesigner":   {"entertainer_master"},
	"smuggler":        {"pistoleer_master", "scout_master"},
}

// countMasterBoxes counts owned Master (Tier 5) boxes, excluding named IDs.
func countMasterBoxes(ownedBoxes map[string]bool, exclude ...string) int {
	excluded := map[string]bool{}
	for _, e := range exclude {
		excluded[e] = true
	}
	n := 0
	for id := range ownedBoxes {
		if excluded[id] {
			continue
		}
		if box := GetSkillBox(id); box != nil && box.Tier == 5 {
			n++
		}
	}
	return n
}

// eliteNoviceCheck enforces cross-profession entry gates (GDD §§8.3/6.3.3).
// TKA 120 follows §33.3 over §8.3.15's 80 (proposal tension recorded).
func eliteNoviceCheck(professionID string, ownedBoxes map[string]bool) (bool, string) {
	switch professionID {
	case "bountyhunter":
		// Master Marksman OR (Master Scout + 4 other Masters) — flagged reading
		// of "Master Marksman OR Master Scout + 4 other Master professions".
		if ownedBoxes["marksman_master"] {
			return true, ""
		}
		if ownedBoxes["scout_master"] && countMasterBoxes(ownedBoxes, "scout_master") >= 4 {
			return true, ""
		}
		return false, "requires Master Marksman, or Master Scout plus 4 other Masters"
	case "commando":
		if ownedBoxes["marksman_master"] && countMasterBoxes(ownedBoxes, "marksman_master") >= 3 {
			return true, ""
		}
		return false, "requires Master Marksman plus 3 other Masters"
	case "teraskasi":
		for _, need := range []string{"brawler_master", "fencer_master", "swordsman_master", "pikeman_master"} {
			if !ownedBoxes[need] {
				return false, "requires Master Brawler, Fencer, Swordsman and Pikeman"
			}
		}
		return true, ""
	default:
		if reqs, ok := ElitePrereqs[professionID]; ok {
			for _, need := range reqs {
				if !ownedBoxes[need] {
					return false, "requires " + need + " first"
				}
			}
		}
		return true, ""
	}
}

// EliteTrainer maps an elite profession to the base trainer profession that
// teaches it (flagged stand-in — no NPC authoring phase; registrar precedent).
// Elites train at the matching basic trainer (e.g. Pistoleer at Marksman).
var EliteTrainer = map[string]string{
	"armorsmith": "artisan", "weaponsmith": "artisan", "architect": "artisan",
	"merchant": "artisan",
	"bioengineer": "scout", "creaturehandler": "scout", "ranger": "scout",
	"pistoleer": "marksman", "rifleman": "marksman", "carbineer": "marksman",
	"commando": "marksman", "bountyhunter": "marksman", "smuggler": "marksman",
	"fencer": "brawler", "swordsman": "brawler", "pikeman": "brawler",
	"teraskasi": "brawler",
	"doctor": "medic", "combatmedic": "medic",
	"musician": "entertainer", "dancer": "entertainer", "imagedesigner": "entertainer",
}

// TrainerCovers reports whether a trainer profession may teach a box's
// profession (same profession, or a mapped elite).
func TrainerCovers(trainerProfessionID, boxProfessionID string) bool {
	if trainerProfessionID == boxProfessionID {
		return true
	}
	return EliteTrainer[boxProfessionID] == trainerProfessionID
}

// PrerequisiteCheck verifies that a character can train a specific skill box.
// Returns (canTrain, reason). GDD Section 7.3, 6.3.3.
//
// Rules:
// - Basic Novice: always trainable (0 XP, 0 credits, 0 skill points)
// - Elite Novice: cross-profession entry gates (eliteNoviceCheck)
// - Tier I-IV: must have Novice + previous tier in the same tree
// - Master: must have all 4 trees complete + Novice
func PrerequisiteCheck(boxID string, ownedBoxes map[string]bool) (bool, string) {
	box := GetSkillBox(boxID)
	if box == nil {
		return false, "skill box not found"
	}

	// Novice: basics always trainable; elites gated.
	if box.Tier == 0 {
		if prof := ProfessionByID(box.ProfessionID); prof != nil && prof.Category == CategoryElite {
			return eliteNoviceCheck(box.ProfessionID, ownedBoxes)
		}
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
