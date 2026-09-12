// ============================================================================
// Pre-CU Schematic Database.
// Only a small representative set — enough to validate the loop.
// Stat weights come from SWGEmu's published resource tester data and community
// reverse engineering. The key mechanic: every experimental property has a
// weighted blend of stats, and OQ is almost always a co-factor.
// ============================================================================

import type { Schematic } from "./types";

// Helper to keep entries compact.
const s = (
  id: string,
  name: string,
  description: string,
  profession: string,
  experimentPoints: number,
  complexity: number,
  slots: Schematic["slots"],
  experimentalProps: Schematic["experimentalProps"],
): Schematic => ({
  id, name, description, profession, experimentPoints, complexity,
  slots, experimentalProps,
});

type Slot = Schematic["slots"][number];
type Prop = Schematic["experimentalProps"][number];

const slot = (
  id: string,
  label: string,
  requiredClassId: string,
  requiredUnits: number,
  statWeights: Slot["statWeights"],
  experimentalContribution: Slot["experimentalContribution"],
): Slot => ({ id, label, requiredClassId, requiredUnits, statWeights, experimentalContribution });

const prop = (
  id: string,
  label: string,
  statWeights: Prop["statWeights"],
  minValue: number,
  maxValue: number,
  unit?: string,
): Prop => ({ id, label, statWeights, minValue, maxValue, unit });

// ----------------------------------------------------------------------------
// CDEF PISTOL — iconic starter weaponsmith weapon.
// 3 slots: barrel (ferrous metal + OQ/UT/SR), grip (non-ferrous metal OQ/CD/UT),
// and a power cell (chemical OQ/PE/CD).
// Experimental properties: Min Damage, Max Damage, Attack Speed, Wound Chance.
// ----------------------------------------------------------------------------
const cdefPistol = s(
  "cdef_pistol",
  "CDEF Pistol",
  "The ubiquitous starter blaster. Trash in combat, gold for calibrating crafting loops.",
  "Weaponsmith",
  8,        // experiment points at max skill
  25,
  [
    slot(
      "barrel",
      "Barrel",
      "mineral.ferrous",
      15,
      { OQ: 0.6, UT: 0.4, SR: 0.0 }, // barrel cares about OQ and UT for durability; SR is irrelevant to barrel assembly
      { min_dmg: 0.4, max_dmg: 0.4, attack_speed: 0.5, wound_chance: 0.3 },
    ),
    slot(
      "grip",
      "Grip & Receiver",
      "mineral.nonferrous",
      10,
      { OQ: 0.5, CD: 0.3, UT: 0.2 }, // copper/aluminum for conductivity
      { min_dmg: 0.3, max_dmg: 0.3, attack_speed: 0.3, wound_chance: 0.3 },
    ),
    slot(
      "power_cell",
      "Power Cell",
      "chemical.petrochem_fuel",
      8,
      { OQ: 0.5, PE: 0.5, CD: 0.0 },
      { min_dmg: 0.3, max_dmg: 0.3, attack_speed: 0.2, wound_chance: 0.4 },
    ),
  ],
  [
    prop("min_dmg",      "Minimum Damage",  { OQ: 0.5, PE: 0.5 }, 11, 25),
    prop("max_dmg",      "Maximum Damage",  { OQ: 0.6, PE: 0.4 }, 28, 75),
    prop("attack_speed", "Attack Speed",    { OQ: 0.7, UT: 0.3 }, 5.0, 2.0, "s/shot"),
      // speed is INVERTED: lower is better. Engine handles this at read time.
    prop("wound_chance", "Wound Chance",    { OQ: 0.4, PE: 0.6 }, 0, 8, "%"),
  ],
);

// ----------------------------------------------------------------------------
// STIM PACK B — medic consumable.
// Requires organic (flora.medicinal) and a chemical.adhesive/bacterio.
// Experimental: Healing (increases HP restored), Charges, Potency.
// ----------------------------------------------------------------------------
const stimPackB = s(
  "stim_b",
  "Stim Pack B",
  "Combat stim that restores health instantly. High-PE flora/chem = bigger heals.",
  "Medic",
  6,
  18,
  [
    slot(
      "bio",
      "Bio-Active Agent",
      "flora.medicinal",
      10,
      { OQ: 0.5, PE: 0.5 },
      { healing: 0.5, charges: 0.4, potency: 0.6 },
    ),
    slot(
      "binding",
      "Binding Agent",
      "chemical.adhesive",
      4,
      { OQ: 0.6, PE: 0.2, UT: 0.2 },
      { healing: 0.3, charges: 0.4, potency: 0.2 },
    ),
    slot(
      "bacterio",
      "Bacterio Boost",
      "chemical.bacterio",
      3,
      { OQ: 0.4, PE: 0.6 },
      { healing: 0.2, charges: 0.2, potency: 0.2 },
    ),
  ],
  [
    prop("healing",  "Healing Amount", { OQ: 0.5, PE: 0.5 }, 80, 250),
    prop("charges",  "Charges",        { OQ: 0.7, UT: 0.3 }, 3, 18),
    prop("potency",  "Wound Healing",  { OQ: 0.4, PE: 0.6 }, 0, 35),
  ],
);

// ----------------------------------------------------------------------------
// MINERAL HARVESTER — Architect / Structures.
// Uses steel, copper, polymer, energy. Experimental: Hopper Size, Extraction
// Rate, Decay Resistance, Power Cost.
// ----------------------------------------------------------------------------
const mineralHarvester = s(
  "harvester_mineral",
  "Mineral Harvester (Personal)",
  "A personal mineral extractor. High CD/UT metals = larger hopper; PE polymers = faster extraction.",
  "Architect",
  10,
  35,
  [
    slot(
      "frame",
      "Frame",
      "mineral.ferrous.steel",
      40,
      { OQ: 0.5, UT: 0.3, DR: 0.2 },
      { hopper: 0.35, extraction: 0.15, decay: 0.5, power: 0.15 },
    ),
    slot(
      "wiring",
      "Wiring & Motors",
      "mineral.nonferrous.copper",
      15,
      { OQ: 0.4, CD: 0.4, UT: 0.2 },
      { hopper: 0.2, extraction: 0.35, decay: 0.2, power: 0.3 },
    ),
    slot(
      "housing",
      "Polymer Housing",
      "chemical.polymer",
      10,
      { OQ: 0.5, MA: 0.3, DR: 0.2 },
      { hopper: 0.25, extraction: 0.25, decay: 0.2, power: 0.2 },
    ),
    slot(
      "power",
      "Power Source",
      "energy.renewable",
      20,
      { OQ: 0.4, PE: 0.6 },
      { hopper: 0.2, extraction: 0.25, decay: 0.1, power: 0.35 },
    ),
  ],
  [
    prop("hopper",     "Hopper Size",    { OQ: 0.5, UT: 0.3, CD: 0.2 }, 1000, 5500),
    prop("extraction", "Extraction Rate",{ OQ: 0.4, PE: 0.4, MA: 0.2 }, 2, 12, "u/cycle"),
    prop("decay",      "Decay Resistance",{ OQ: 0.7, DR: 0.3 }, 100, 850),
    prop("power",      "Power Efficiency",{ OQ: 0.3, PE: 0.5, CD: 0.2 }, 30, 6, "u/hr"),
  ],
);

// ----------------------------------------------------------------------------
// TRAVEL BISCUIT — Chef. Demonstrates flora + water + organic.
// Experimental: Nutrition (fills hunger), Duration, Filling.
// ----------------------------------------------------------------------------
const travelBiscuit = s(
  "travel_biscuit",
  "Travel Biscuit",
  "Dry, long-lasting survival ration. Ideal for scouts on long trips.",
  "Chef",
  5,
  10,
  [
    slot(
      "flour",
      "Flour Base",
      "flora.grain",
      5,
      { OQ: 0.6, PE: 0.4 },
      { nutrition: 0.4, duration: 0.5, filling: 0.6 },
    ),
    slot(
      "sweetener",
      "Sweetener",
      "flora.fruit",
      3,
      { OQ: 0.5, PE: 0.5 },
      { nutrition: 0.3, duration: 0.2, filling: 0.1 },
    ),
    slot(
      "water",
      "Mixing Water",
      "water",
      2,
      { OQ: 0.7, PE: 0.3 },
      { nutrition: 0.3, duration: 0.3, filling: 0.3 },
    ),
  ],
  [
    prop("nutrition", "Nutrition",    { OQ: 0.4, PE: 0.6 }, 100, 350),
    prop("duration",  "Duration",     { OQ: 0.5, PE: 0.5 }, 120, 900, "s"),
    prop("filling",   "Filling",      { OQ: 0.6, PE: 0.4 }, 10, 45),
  ],
);

export const SCHEMATICS: ReadonlyMap<string, Schematic> = new Map(
  [cdefPistol, stimPackB, mineralHarvester, travelBiscuit].map(s => [s.id, s] as const)
);

export const SCHEMATIC_LIST: ReadonlyArray<Schematic> = [cdefPistol, stimPackB, mineralHarvester, travelBiscuit];
