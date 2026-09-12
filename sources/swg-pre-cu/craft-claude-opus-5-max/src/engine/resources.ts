/**
 * Resource domain: attribute lanes, the class tree, cap resolution, and the
 * global spawn timeline.
 *
 * Design notes
 * ------------
 * - Attributes are a fixed 10-lane vocabulary. A resource class exposes a
 *   SUBSET of lanes. Absence is semantic ("Steel has no Potential Energy"),
 *   not zero. Modeled as a sparse map keyed by lane, NOT a 10-float struct
 *   with sentinels, because sentinels leak into every weighted average.
 * - The class tree is a flat array + id index (flyweight). Nothing holds a
 *   pointer to a class; everything holds an interned string id. This is the
 *   shape that survives serialization into a UE5 DataTable row.
 * - Caps are INHERITED AND NARROWED down the tree. resolveCaps() walks the
 *   parent chain once and memoizes. Adding a tier to the taxonomy is a
 *   one-line edit, not a class hierarchy refactor.
 */

import { Rng } from './rng';

export const STAT_KEYS = ['OQ', 'CD', 'CR', 'DR', 'FL', 'HR', 'MA', 'PE', 'SR', 'UT'] as const;
export type StatKey = (typeof STAT_KEYS)[number];

export const STAT_NAMES: Record<StatKey, string> = {
  OQ: 'Overall Quality',
  CD: 'Conductivity',
  CR: 'Cold Resistance',
  DR: 'Decay Resistance',
  FL: 'Flavor',
  HR: 'Heat Resistance',
  MA: 'Malleability',
  PE: 'Potential Energy',
  SR: 'Shock Resistance',
  UT: 'Unit Toughness',
};

export type StatRange = readonly [number, number];
export type CapTable = Partial<Record<StatKey, StatRange>>;
export type StatVector = Partial<Record<StatKey, number>>;

export interface ResourceClass {
  id: string;
  name: string;
  parent: string | null;
  /** partial override; merged over the resolved parent table */
  caps: CapTable;
  /** only leaves spawn; interior nodes exist purely for schematic matching */
  spawnable: boolean;
  /** 0 = ubiquitous, 1 = scarce. Drives gap length in the spawn timeline. */
  scarcity: number;
}

const C = (
  id: string,
  name: string,
  parent: string | null,
  caps: CapTable,
  spawnable = false,
  scarcity = 0.2,
): ResourceClass => ({ id, name, parent, caps, spawnable, scarcity });

export const RESOURCE_CLASSES: readonly ResourceClass[] = [
  C('resource', 'Resource', null, {}),

  // ---- INORGANIC -------------------------------------------------------
  C('inorganic', 'Inorganic', 'resource', {}),

  C('mineral', 'Mineral', 'inorganic', {
    CD: [1, 500], CR: [1, 500], DR: [1, 700], HR: [1, 600],
    MA: [1, 700], OQ: [1, 800], SR: [1, 700], UT: [1, 600],
  }),
  C('igneous', 'Igneous', 'mineral', { HR: [300, 1000], MA: [1, 500], SR: [300, 1000] }),
  C('extrusive_ore', 'Extrusive Ore', 'igneous', { OQ: [1, 900], CD: [1, 600] }, true, 0.15),
  C('intrusive_ore', 'Intrusive Ore', 'igneous', { OQ: [200, 1000], DR: [300, 900] }, true, 0.25),
  C('carbonate_ore', 'Carbonate Ore', 'mineral', { MA: [400, 1000], SR: [1, 500], OQ: [1, 750] }, true, 0.15),
  C('siliclastic_ore', 'Siliclastic Ore', 'mineral', { CD: [300, 900], HR: [1, 500], OQ: [1, 800] }, true, 0.15),

  C('metal', 'Metal', 'inorganic', {
    CD: [1, 700], CR: [1, 700], DR: [1, 800], HR: [1, 800],
    MA: [1, 800], OQ: [1, 1000], SR: [1, 800], UT: [1, 800],
  }),
  C('ferrous', 'Ferrous Metal', 'metal', { SR: [400, 1000], CD: [1, 400] }),
  C('steel', 'Steel', 'ferrous', { MA: [1, 600], HR: [300, 1000], DR: [200, 1000], OQ: [1, 1000] }, true, 0.3),
  C('iron', 'Iron', 'ferrous', { MA: [200, 800], DR: [1, 700], HR: [1, 700], OQ: [1, 900] }, true, 0.2),
  C('nonferrous', 'Non-Ferrous Metal', 'metal', { SR: [1, 600] }),
  C('copper', 'Copper', 'nonferrous', { CD: [600, 1000], HR: [1, 700], MA: [400, 1000], OQ: [1, 1000] }, true, 0.25),
  C('aluminum', 'Aluminum', 'nonferrous', { CD: [300, 800], MA: [500, 1000], DR: [1, 600], OQ: [1, 1000] }, true, 0.25),

  C('radioactive', 'Radioactive', 'inorganic', {
    CD: [600, 1000], CR: [1, 600], DR: [1, 600], HR: [1, 600], MA: [1, 400],
    OQ: [1, 1000], PE: [600, 1000], SR: [1, 500], UT: [1, 600],
  }),
  C('known_radioactive', 'Known Radioactive', 'radioactive', { PE: [700, 1000], OQ: [1, 900] }, true, 0.5),
  C('polymetric', 'Polymetric Radioactive', 'radioactive', { PE: [600, 900], CD: [800, 1000], OQ: [200, 1000] }, true, 0.6),

  C('chemical', 'Chemical', 'inorganic', {
    CD: [1, 700], DR: [1, 800], OQ: [1, 1000], PE: [1, 800], UT: [1, 600],
  }),
  C('petrochem_solid', 'Petrochem Fuel (Solid)', 'chemical', { PE: [400, 1000], DR: [1, 700], OQ: [1, 900] }, true, 0.2),
  C('petrochem_liquid', 'Petrochem Fuel (Liquid)', 'chemical', { PE: [500, 1000], CD: [1, 500], OQ: [1, 900] }, true, 0.2),
  C('lubricating_oil', 'Lubricating Oil', 'chemical', { CD: [1, 400], DR: [300, 900], OQ: [1, 800], UT: [200, 900] }, true, 0.3),

  C('gas', 'Gas', 'inorganic', { CD: [1, 700], DR: [1, 700], OQ: [1, 1000], PE: [1, 900] }),
  C('reactive_gas', 'Reactive Gas', 'gas', { PE: [500, 1000], CD: [300, 900], DR: [1, 600] }, true, 0.35),
  C('inert_gas', 'Inert Gas', 'gas', { PE: [1, 500], DR: [400, 1000], CD: [1, 400], OQ: [1, 900] }, true, 0.35),

  // ---- ORGANIC ---------------------------------------------------------
  C('organic', 'Organic', 'resource', {}),

  C('flora', 'Flora', 'organic', {
    DR: [1, 800], FL: [1, 800], OQ: [1, 1000], PE: [1, 800], UT: [1, 800],
  }),
  C('cereal', 'Cereal', 'flora', { FL: [1, 900], PE: [300, 1000] }),
  C('corn', 'Corn', 'cereal', { FL: [200, 1000], DR: [1, 700], UT: [1, 700] }, true, 0.1),
  C('wheat', 'Wheat', 'cereal', { PE: [400, 1000], FL: [1, 800] }, true, 0.1),
  C('oats', 'Oats', 'cereal', { UT: [300, 1000], DR: [200, 900], OQ: [1, 900] }, true, 0.1),
  C('vegetable', 'Vegetable', 'flora', { FL: [1, 800], PE: [1, 700], DR: [1, 900], UT: [1, 900] }),
  C('tuber', 'Tuber Vegetable', 'vegetable', { UT: [300, 1000], FL: [1, 700] }, true, 0.1),
  C('beans', 'Beans', 'vegetable', { PE: [300, 900], FL: [200, 1000] }, true, 0.1),

  C('creature', 'Creature', 'organic', {
    DR: [1, 800], MA: [1, 800], OQ: [1, 1000], SR: [1, 800], UT: [1, 800],
  }),
  C('hide', 'Hide', 'creature', { MA: [300, 1000], SR: [1, 900], DR: [1, 900] }),
  C('bristley_hide', 'Bristley Hide', 'hide', { SR: [400, 1000], MA: [1, 600] }, true, 0.2),
  C('leathery_hide', 'Leathery Hide', 'hide', { MA: [500, 1000], DR: [300, 1000] }, true, 0.2),
  C('scaley_hide', 'Scaley Hide', 'hide', { SR: [600, 1000], MA: [1, 500], UT: [300, 1000] }, true, 0.3),
  C('wooly_hide', 'Wooly Hide', 'hide', { UT: [500, 1000], MA: [400, 900], SR: [1, 600] }, true, 0.2),
  C('bone', 'Bone', 'creature', { MA: [1, 600], SR: [300, 1000], DR: [300, 1000], UT: [1, 700] }),
  C('animal_bone', 'Animal Bone', 'bone', { SR: [400, 1000] }, true, 0.15),
  C('horn', 'Horn', 'bone', { MA: [200, 800], DR: [500, 1000] }, true, 0.35),
  C('meat', 'Meat', 'creature', { DR: [1, 700], FL: [1, 1000], PE: [1, 700], UT: [1, 700], MA: [1, 400] }),
  C('domesticated_meat', 'Domesticated Meat', 'meat', { FL: [200, 900] }, true, 0.1),
  C('wild_meat', 'Wild Meat', 'meat', { FL: [1, 1000], UT: [200, 900] }, true, 0.15),
];

export const CLASS_INDEX: ReadonlyMap<string, ResourceClass> = new Map(
  RESOURCE_CLASSES.map((c) => [c.id, c]),
);

export const CHILDREN_INDEX: ReadonlyMap<string, string[]> = (() => {
  const m = new Map<string, string[]>();
  for (const c of RESOURCE_CLASSES) {
    if (!c.parent) continue;
    const arr = m.get(c.parent);
    if (arr) arr.push(c.id);
    else m.set(c.parent, [c.id]);
  }
  return m;
})();

const capCache = new Map<string, CapTable>();

/** Walk root -> leaf, merging cap overrides. Memoized. */
export function resolveCaps(classId: string): CapTable {
  const hit = capCache.get(classId);
  if (hit) return hit;
  const cls = CLASS_INDEX.get(classId);
  if (!cls) return {};
  const merged: CapTable = cls.parent ? { ...resolveCaps(cls.parent), ...cls.caps } : { ...cls.caps };
  capCache.set(classId, merged);
  return merged;
}

export function statLanes(classId: string): StatKey[] {
  const caps = resolveCaps(classId);
  return STAT_KEYS.filter((k) => caps[k] !== undefined);
}

/** Is `classId` inside the subtree rooted at `rootId`? Schematic slot matching. */
export function isDescendant(classId: string, rootId: string): boolean {
  let cur: string | null = classId;
  while (cur) {
    if (cur === rootId) return true;
    cur = CLASS_INDEX.get(cur)?.parent ?? null;
  }
  return false;
}

export function classChain(classId: string): ResourceClass[] {
  const out: ResourceClass[] = [];
  let cur: string | null = classId;
  while (cur) {
    const c = CLASS_INDEX.get(cur);
    if (!c) break;
    out.unshift(c);
    cur = c.parent;
  }
  return out;
}

// ---- Spawn timeline ------------------------------------------------------

export const PLANETS = [
  'Corellia', 'Dantooine', 'Dathomir', 'Endor', 'Lok',
  'Naboo', 'Rori', 'Talus', 'Tatooine', 'Yavin IV',
] as const;
export type Planet = (typeof PLANETS)[number];

export interface ResourceSpawn {
  id: string;
  name: string;
  classId: string;
  stats: StatVector;
  planets: { planet: Planet; concentration: number }[];
  /** galactic day index, inclusive */
  spawnTick: number;
  /** galactic day index, exclusive */
  despawnTick: number;
  /** cached mean of (stat - min) / (max - min) across present lanes */
  purity: number;
}

const ONSET = ['ba', 'de', 'ka', 'thu', 'vor', 'zen', 'qua', 'mi', 'so', 'lar', 'nu', 'gre', 'py', 'ta', 'os', 'ir', 've', 'do', 'ash', 'kel'];
const CODA = ['nite', 'ium', 'ar', 'os', 'ex', 'ite', 'on', 'us', 'al', 'yx', 'or', 'esh', 'ax', 'in', 'eth', 'ol'];
const ROMAN = ['II', 'III', 'IV', 'V', 'VI', 'IX', 'X'];

function genName(rng: Rng): string {
  let s = rng.pick(ONSET) + (rng.chance(0.55) ? rng.pick(ONSET) : '') + rng.pick(CODA);
  s = s[0].toUpperCase() + s.slice(1);
  if (rng.chance(0.22)) s += ' ' + rng.pick(ROMAN);
  return s;
}

function rollStats(rng: Rng, classId: string): { stats: StatVector; purity: number } {
  const caps = resolveCaps(classId);
  // Per-spawn "richness" bias. Rare high-yield spawns are the entire economy.
  const richness = rng.next();
  const exp = richness > 0.965 ? 0.45 : richness > 0.85 ? 0.9 : 1.7;
  const stats: StatVector = {};
  let acc = 0;
  let n = 0;
  for (const k of STAT_KEYS) {
    const cap = caps[k];
    if (!cap) continue;
    // per-lane jitter so a spawn is rarely uniformly good
    const laneExp = Math.max(0.3, exp * rng.range(0.7, 1.45));
    const v = Math.round(rng.curve(cap[0], cap[1], laneExp));
    stats[k] = v;
    acc += cap[1] > cap[0] ? (v - cap[0]) / (cap[1] - cap[0]) : 1;
    n++;
  }
  return { stats, purity: n ? acc / n : 0 };
}

export interface GalaxyConfig {
  seed: string;
  /** days of history+future to materialize */
  horizon: number;
}

export interface Galaxy {
  seed: string;
  horizon: number;
  /** append-only. Never mutated. Query by tick window. */
  spawns: readonly ResourceSpawn[];
}

/**
 * Materialize the entire spawn timeline up front as an append-only log of
 * [spawnTick, despawnTick) intervals. "What is active on day T" is a filter,
 * not a simulation step. Consequences: time travel is free, the pool is
 * diffable, and there is no tick loop to desync.
 */
export function generateGalaxy(cfg: GalaxyConfig): Galaxy {
  const spawns: ResourceSpawn[] = [];
  let serial = 0;

  for (const cls of RESOURCE_CLASSES) {
    if (!cls.spawnable) continue;
    // Deterministic per-class stream: changing one class's tuning does not
    // reshuffle the entire galaxy.
    const rng = new Rng(`${cfg.seed}::${cls.id}`);
    let cursor = -rng.int(0, 30);
    while (cursor < cfg.horizon) {
      const duration = rng.int(6, 22);
      const { stats, purity } = rollStats(rng, cls.id);
      const planetCount = rng.int(1, 3);
      const pool = [...PLANETS];
      const planets: ResourceSpawn['planets'] = [];
      for (let i = 0; i < planetCount; i++) {
        const idx = rng.int(0, pool.length - 1);
        planets.push({ planet: pool.splice(idx, 1)[0], concentration: rng.int(18, 96) });
      }
      spawns.push({
        id: `sp_${cls.id}_${serial++}`,
        name: genName(rng),
        classId: cls.id,
        stats,
        planets,
        spawnTick: cursor,
        despawnTick: cursor + duration,
        purity,
      });
      cursor += duration + rng.int(0, Math.round(cls.scarcity * 40));
    }
  }

  spawns.sort((a, b) => a.spawnTick - b.spawnTick || a.name.localeCompare(b.name));
  return { seed: cfg.seed, horizon: cfg.horizon, spawns };
}

export function activeAt(galaxy: Galaxy, tick: number): ResourceSpawn[] {
  return galaxy.spawns.filter((s) => s.spawnTick <= tick && tick < s.despawnTick);
}

/** % of the class cap window this value occupies. Drives all UI heat. */
export function capPct(classId: string, stat: StatKey, value: number): number {
  const cap = resolveCaps(classId)[stat];
  if (!cap) return 0;
  const [lo, hi] = cap;
  return hi > lo ? Math.max(0, Math.min(1, (value - lo) / (hi - lo))) : 1;
}
