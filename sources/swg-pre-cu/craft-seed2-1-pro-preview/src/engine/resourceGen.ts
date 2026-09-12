// ============================================================================
// Random Resource Spawn Generator.
// In the real SWG, server ticker rolls new spawns periodically, old ones
// despawn. For this prototype we generate a single global "current spawn
// table" across all planets — enough to populate the economy loop.
// ============================================================================

import type { PlanetId, Resource, ResourceClass, StatMap, StatId } from "./types";
import { ALL_STATS } from "./types";
import { RESOURCE_CLASSES } from "./data/resourceTree";
import { makeRng, sampleStat, intBetween } from "./rng";

const FIRST_PARTS = [
  "Vatrisian", "Kolto", "Beryllius", "Kammrath", "Tatooinian", "Nabooan",
  "Dantari", "Corellian", "Alderaanian", "Rhodium", "Vandinite", "Sienari",
  "Kuat", "Kylantha", "Mandalorian", "Thyferran", "Ryll", "Glitterstim",
  "Kashyyykian", "Mon", "Bespin", "Endorian", "Dathomirian", "Yavinian",
  "Rorian", "Talusian", "Lokian", "Felucan", "Mygeetan", "Muunilinst",
];

const LAST_PARTS = [
  "Iron", "Steel", "Aluminum", "Copper", "Ore", "Petrochem", "Polymer",
  "Fiberplast", "Wood", "Water", "Gas", "Reactive", "Inert", "Radioactive",
  "Gemstone", "Carbonate", "Siliclastic", "Igneous", "Sedimentary",
  "Horn", "Ivory", "Hide", "Bone", "Meat", "Milk", "Scale",
  "Oil", "Adhesive", "Acid", "Bacterio", "Surfactant", "Polymer",
  "Fruit", "Flower", "Grain", "Greens", "Beans", "Hops", "Seeds", "Spice",
  "Energy", "Solar", "Wind", "Fusion", "Fission", "Geothermal",
];

function emptyStatMap(): { [S in StatId]: number } {
  return { OQ: 0, PE: 0, UT: 0, CR: 0, CD: 0, DR: 0, HR: 0, MA: 0, SR: 0 };
}

function randomName(rng: () => number): string {
  const first = FIRST_PARTS[Math.floor(rng() * FIRST_PARTS.length)];
  const last = LAST_PARTS[Math.floor(rng() * LAST_PARTS.length)];
  // Avoid duplicate first/last like "Iron Iron"
  if (first === last) return randomName(rng);
  return `${first} ${last}`;
}

/**
 * Generate stats for a resource instance. Stats not in the class's relevant
 * list stay at 0. OQ is always rolled. For flavor: high-OQ resources are
 * slightly rarer — we use a higher skew for OQ.
 */
function rollStats(rng: () => number, cls: ResourceClass): StatMap {
  const stats: { [S in StatId]: number } = emptyStatMap();
  // OQ always present, slightly higher skew (rarer to be 900+).
  stats.OQ = sampleStat(rng, 4);
  for (const s of cls.relevantStats) {
    if (s === "OQ") continue;
    stats[s] = sampleStat(rng, 3);
  }
  // Irrelevant stats remain 0 (signal to the crafter "this stat doesn't
  // apply to this resource class").
  return stats as StatMap;
}

const PLANETS: PlanetId[] = [
  "tatooine","naboo","corellia","dantooine","endor",
  "dathomir","yavin4","rori","talus","lok",
];

/**
 * Produce a global spawn table. Each planet gets ~6-10 active resources.
 * We bias spawns toward intermediate classes (not pure roots, not only leaves)
 * to give the economy realistic gating.
 */
export function generateSpawnTable(
  seed: number,
  perPlanetMin = 6,
  perPlanetMax = 10,
): Resource[] {
  const rng = makeRng(seed);
  const allClasses = Array.from(RESOURCE_CLASSES.values())
    // Exclude pure root classes — you never find "Mineral", you find Iron.
    .filter(c => c.parent !== undefined);

  const out: Resource[] = [];
  let idCounter = 0;

  for (const planet of PLANETS) {
    const count = intBetween(rng, perPlanetMin, perPlanetMax);
    // Each planet has a bias toward certain top-level classes (Tatooine has
    // lots of Mineral/Chemical, Naboo has Flora/Water, etc.). For brevity we
    // do uniform here — this is a demo knob.
    for (let i = 0; i < count; i++) {
      const cls = allClasses[Math.floor(rng() * allClasses.length)];
      out.push({
        id: `res_${++idCounter}`,
        name: randomName(rng),
        classId: cls.id,
        planetId: planet,
        stats: rollStats(rng, cls),
        availableUnits: intBetween(rng, 5000, 25000),
      });
    }
  }

  return out;
}

// Utility for the UI: describe which stats are non-zero.
export function relevantStatsOf(r: Resource): StatId[] {
  return ALL_STATS.filter(s => r.stats[s] > 0);
}
