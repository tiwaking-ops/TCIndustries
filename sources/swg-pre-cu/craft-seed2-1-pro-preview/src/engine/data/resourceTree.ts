// ============================================================================
// Pre-CU Resource Class Hierarchy
// Source: SWGEmu published schematics + community decompiled client data.
// We represent it as a flat map with parent pointers for O(1) lookup and a
// recursive subtype check (easily ported to a hash-set cache in C++).
// ============================================================================

import type { ResourceClass, StatId } from "../types";

// Helper to build entries compactly.
const rc = (
  id: string,
  name: string,
  parent: string | undefined,
  relevantStats: StatId[]
): ResourceClass => ({ id, name, parent, relevantStats });

// In Pre-CU, every resource class has a specific subset of stats it can spawn
// with. OQ is universal. Other stats only appear on relevant classes. This
// directly affects schematic gating — e.g. a Chemical gives you PE and CD but
// not UT or SR.
const ALL_CLASSES: ResourceClass[] = [
  // --- TOP-LEVEL -----------------------------------------------------------
  rc("mineral",         "Mineral",         undefined, ["OQ","CR","DR","HR","MA","SR","UT"]),
  rc("chemical",        "Chemical",        undefined, ["OQ","CD","DR","PE","UT"]),
  rc("flora",           "Flora",           undefined, ["OQ","PE","UT"]),
  rc("gas",             "Gas",             undefined, ["OQ","CD","CR","DR","HR","PE","UT"]),
  rc("water",           "Water",           undefined, ["OQ","PE","UT"]),
  rc("energy",          "Energy",          undefined, ["OQ","PE"]),
  rc("creature",        "Creature Resource", undefined, ["OQ","PE","UT","CR","DR","HR","SR","MA","CD"]),

  // --- MINERAL BRANCHES ----------------------------------------------------
  rc("mineral.ferrous",     "Ferrous Metal",     "mineral", ["OQ","CR","DR","HR","MA","SR","UT"]),
  rc("mineral.nonferrous",  "Non-Ferrous Metal", "mineral", ["OQ","CD","CR","DR","HR","MA","SR","UT"]),
  rc("mineral.ferrous.iron",     "Iron",     "mineral.ferrous",     ["OQ","CR","DR","HR","MA","SR","UT"]),
  rc("mineral.ferrous.steel",    "Steel",    "mineral.ferrous",     ["OQ","CR","DR","HR","MA","SR","UT"]),
  rc("mineral.nonferrous.aluminum", "Aluminum", "mineral.nonferrous", ["OQ","CD","MA","SR","UT"]),
  rc("mineral.nonferrous.copper",   "Copper",   "mineral.nonferrous", ["OQ","CD","MA","SR","UT"]),
  rc("mineral.gemstone",    "Gemstone",    "mineral", ["OQ","CR","DR","HR","PE","SR","UT"]),
  rc("mineral.radioactive", "Radioactive", "mineral", ["OQ","CD","DR","HR","PE","UT"]),
  rc("mineral.ore",         "Ore",         "mineral", ["OQ","CR","DR","HR","MA","SR","UT"]),
  rc("mineral.ore.iron_ore",    "Iron Ore",    "mineral.ore", ["OQ","CR","DR","HR","MA","SR","UT"]),
  rc("mineral.ore.alu_ore",     "Aluminum Ore", "mineral.ore", ["OQ","CR","DR","HR","MA","SR","UT"]),
  rc("mineral.ore.carbonate",   "Carbonate Ore","mineral.ore", ["OQ","CR","DR","HR","MA","SR","UT"]),
  rc("mineral.ore.siliclastic", "Siliclastic Ore","mineral.ore", ["OQ","CR","DR","HR","MA","SR","UT"]),
  rc("mineral.ore.igneous",     "Igneous Ore",  "mineral.ore", ["OQ","CR","DR","HR","MA","SR","UT"]),
  rc("mineral.ore.sedimentary", "Sedimentary Ore","mineral.ore", ["OQ","CR","DR","HR","MA","SR","UT"]),
  rc("mineral.ore.desh",        "Desh Copper Ore","mineral.ore", ["OQ","CD","CR","DR","HR","MA","SR","UT"]),
  rc("mineral.ore.ferrous",     "Ferrous Ore",  "mineral.ore", ["OQ","CR","DR","HR","MA","SR","UT"]),

  // --- CHEMICAL ------------------------------------------------------------
  rc("chemical.lubricating",    "Lubricating Oil",    "chemical", ["OQ","DR","PE","UT"]),
  rc("chemical.petrochem_fuel", "Petrochem Fuel",     "chemical", ["OQ","CD","DR","PE","UT"]),
  rc("chemical.polymer",        "Polymer",            "chemical", ["OQ","CD","DR","MA","PE","UT"]),
  rc("chemical.adhesive",       "Adhesive",           "chemical", ["OQ","CD","DR","MA","PE","UT"]),
  rc("chemical.surfactant",     "Surfactant",         "chemical", ["OQ","CD","DR","PE","UT"]),
  rc("chemical.acid",           "Acid",               "chemical", ["OQ","CD","DR","PE","UT"]),
  rc("chemical.bacterio",       "Bacterio Agent",     "chemical", ["OQ","PE","UT"]),

  // --- FLORA ---------------------------------------------------------------
  rc("flora.food",          "Edible Flora",     "flora", ["OQ","PE","UT"]),
  rc("flora.structural",    "Structural Flora", "flora", ["OQ","UT"]),
  rc("flora.medicinal",     "Medicinal Flora",  "flora", ["OQ","PE","UT"]),
  rc("flora.wood",          "Wood",             "flora.structural", ["OQ","UT"]),
  rc("flora.fiberplast",    "Fiberplast",       "flora.structural", ["OQ","UT"]),
  rc("flora.fruit",         "Fruit",            "flora.food",       ["OQ","PE","UT"]),
  rc("flora.flower",        "Flower",           "flora.medicinal",  ["OQ","PE","UT"]),
  rc("flora.greens",        "Greens",           "flora.food",       ["OQ","PE","UT"]),
  rc("flora.beans",         "Beans",            "flora.food",       ["OQ","PE","UT"]),
  rc("flora.grain",         "Grain",            "flora.food",       ["OQ","PE","UT"]),
  rc("flora.hops",          "Hops",             "flora.food",       ["OQ","PE","UT"]),
  rc("flora.seeds",         "Seeds",            "flora.food",       ["OQ","PE","UT"]),
  rc("flora.spice",         "Spice",            "flora",            ["OQ","PE","UT"]),
  rc("flora.reactive",      "Reactive Gas Flora","flora",          ["OQ","PE","UT"]),

  // --- GAS -----------------------------------------------------------------
  rc("gas.reactive",      "Reactive Gas",    "gas", ["OQ","CD","CR","DR","HR","PE","UT"]),
  rc("gas.inert",         "Inert Gas",       "gas", ["OQ","CD","CR","DR","HR","PE","UT"]),
  rc("gas.known_reactive","Known Reactive",  "gas.reactive", ["OQ","CD","CR","DR","HR","PE","UT"]),
  rc("gas.known_inert",   "Known Inert",     "gas.inert",    ["OQ","CD","CR","DR","HR","PE","UT"]),

  // --- WATER ---------------------------------------------------------------
  rc("water.pure",       "Pure Water",      "water", ["OQ","PE","UT"]),
  rc("water.mineral",    "Mineral Water",   "water", ["OQ","PE","UT"]),
  rc("water.river",      "River Water",     "water", ["OQ","PE","UT"]),
  rc("water.swamp",      "Swamp Water",     "water", ["OQ","PE","UT"]),
  rc("water.ocean",      "Ocean Water",     "water", ["OQ","PE","UT"]),

  // --- ENERGY --------------------------------------------------------------
  rc("energy.renewable",  "Renewable Energy", "energy", ["OQ","PE"]),
  rc("energy.nonrenewable","Non-Renewable Energy","energy", ["OQ","PE"]),
  rc("energy.solar",     "Solar Energy",    "energy.renewable",    ["OQ","PE"]),
  rc("energy.wind",      "Wind Energy",     "energy.renewable",    ["OQ","PE"]),
  rc("energy.hydro",     "Hydro Energy",    "energy.renewable",    ["OQ","PE"]),
  rc("energy.fusion",    "Fusion Energy",   "energy.nonrenewable", ["OQ","PE"]),
  rc("energy.chemical",  "Chemical Energy", "energy.nonrenewable", ["OQ","PE"]),
  rc("energy.fission",   "Fission Energy",  "energy.nonrenewable", ["OQ","PE"]),
  rc("energy.geothermal","Geothermal Energy","energy.renewable",   ["OQ","PE"]),

  // --- CREATURE RESOURCES --------------------------------------------------
  rc("creature.bone",       "Creature Bone",     "creature", ["OQ","CR","DR","HR","MA","SR","UT"]),
  rc("creature.hide",       "Creature Hide",     "creature", ["OQ","CR","DR","HR","MA","SR","UT"]),
  rc("creature.meat",       "Creature Meat",     "creature", ["OQ","PE","UT"]),
  rc("creature.milk",       "Creature Milk",     "creature", ["OQ","PE","UT"]),
  rc("creature.ivory",      "Ivory",             "creature.bone", ["OQ","CR","DR","HR","MA","SR","UT"]),
  rc("creature.horn",       "Horn",              "creature.bone", ["OQ","CR","DR","HR","MA","SR","UT"]),
  rc("creature.avian_bone", "Avian Bone",        "creature.bone", ["OQ","CR","DR","HR","MA","SR","UT"]),
  rc("creature.scale_hide", "Scaley Hide",       "creature.hide", ["OQ","CR","DR","HR","MA","SR","UT"]),
  rc("creature.wooly_hide", "Wooly Hide",        "creature.hide", ["OQ","CR","DR","HR","MA","SR","UT"]),
  rc("creature.leathery_hide","Leathery Hide",   "creature.hide", ["OQ","CR","DR","HR","MA","SR","UT"]),
  rc("creature.soft_tissue","Soft Tissue",      "creature.meat", ["OQ","PE","UT"]),
  rc("creature.seafood",    "Seafood",           "creature.meat", ["OQ","PE","UT"]),
];

export const RESOURCE_CLASSES: ReadonlyMap<string, ResourceClass> = new Map(
  ALL_CLASSES.map(c => [c.id, c] as const)
);

// ------- Inheritance utilities (portable to C++ with precomputed sets) -----

export function getClass(id: string): ResourceClass {
  const c = RESOURCE_CLASSES.get(id);
  if (!c) throw new Error(`Unknown resource class: ${id}`);
  return c;
}

/**
 * Returns true if `candidateId` equals `ancestorId` or descends from it.
 * Cached in hot paths via a Map<string, Set<string>> of "all subtype ids".
 */
export function isSubtypeOf(candidateId: string, ancestorId: string): boolean {
  if (candidateId === ancestorId) return true;
  let c = RESOURCE_CLASSES.get(candidateId);
  while (c && c.parent) {
    if (c.parent === ancestorId) return true;
    c = RESOURCE_CLASSES.get(c.parent);
  }
  return false;
}

/** Return the leaf-most root ancestor (e.g. Iron -> mineral). */
export function getRootClass(id: string): ResourceClass {
  let c = getClass(id);
  while (c.parent) {
    const p = RESOURCE_CLASSES.get(c.parent);
    if (!p) break;
    c = p;
  }
  return c;
}
