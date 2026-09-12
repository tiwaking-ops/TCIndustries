/**
 * Runtime invariant harness. Not a test framework — a tripwire that runs in
 * the page so a bad tuning edit is visible immediately rather than three
 * sessions later. Cheap enough to run on every galaxy regeneration.
 */

import { Rng } from './rng';
import {
  CLASS_INDEX,
  RESOURCE_CLASSES,
  generateGalaxy,
  isDescendant,
  resolveCaps,
  statLanes,
  type Galaxy,
  type ResourceSpawn,
} from './resources';
import { SCHEMATICS } from './schematics';
import {
  assemble,
  ceilingFor,
  experiment,
  newSession,
  propertyQuality,
  type CraftState,
} from './craft';

export interface Assertion {
  name: string;
  pass: boolean;
  detail: string;
}

function fnv(str: string): string {
  let h = 0x811c9dc5;
  for (let i = 0; i < str.length; i++) {
    h ^= str.charCodeAt(i);
    h = Math.imul(h, 0x01000193);
  }
  return (h >>> 0).toString(16).padStart(8, '0');
}

export function runSelfTests(galaxy: Galaxy): Assertion[] {
  const out: Assertion[] = [];

  // 1. Determinism: same seed => byte-identical timeline.
  const twin = generateGalaxy({ seed: galaxy.seed, horizon: galaxy.horizon });
  const a = fnv(JSON.stringify(galaxy.spawns));
  const b = fnv(JSON.stringify(twin.spawns));
  out.push({
    name: 'determinism',
    pass: a === b,
    detail: `fnv(timeline) ${a} == ${b}`,
  });

  // 2. Cap containment: no stat may escape its resolved class window.
  let violations = 0;
  let checked = 0;
  for (const s of galaxy.spawns) {
    const caps = resolveCaps(s.classId);
    for (const [k, v] of Object.entries(s.stats)) {
      const cap = caps[k as keyof typeof caps];
      checked++;
      if (!cap || v < cap[0] || v > cap[1]) violations++;
    }
  }
  out.push({
    name: 'cap containment',
    pass: violations === 0,
    detail: `${checked} lane values inside inherited windows · ${violations} violations`,
  });

  // 3. Lane fidelity: a spawn exposes exactly the lanes its class declares.
  let laneErr = 0;
  for (const s of galaxy.spawns) {
    const want = statLanes(s.classId).join(',');
    const got = Object.keys(s.stats).sort().join(',');
    if (want.split(',').sort().join(',') !== got) laneErr++;
  }
  out.push({
    name: 'lane fidelity',
    pass: laneErr === 0,
    detail: `${galaxy.spawns.length} spawns match declared lane sets · ${laneErr} mismatches`,
  });

  // 4. Schematic reachability: every slot must be satisfiable by real spawns.
  const unreachable: string[] = [];
  for (const sc of SCHEMATICS) {
    for (const slot of sc.slots) {
      const ok = RESOURCE_CLASSES.some((c) => c.spawnable && isDescendant(c.id, slot.classId));
      if (!ok) unreachable.push(`${sc.id}.${slot.id}`);
    }
    for (const p of sc.properties) {
      for (const g of p.groups) {
        if (!sc.slots.some((s) => s.id === g.slotId)) unreachable.push(`${sc.id}.${p.id}→${g.slotId}`);
      }
    }
  }
  out.push({
    name: 'schematic reachability',
    pass: unreachable.length === 0,
    detail: unreachable.length
      ? `dangling: ${unreachable.join(', ')}`
      : `${SCHEMATICS.length} schematics · all slots satisfiable, all groups bound`,
  });

  // 5. Ceiling invariant: 400 randomized sessions, pct must never exceed ceiling.
  const rng = new Rng(`selftest::${galaxy.seed}`);
  let breach = 0;
  let sessions = 0;
  let destroyed = 0;

  // Hoist slot pools out of the session loop: one pass over the log per slot
  // class instead of 400. Same discipline the bench UI uses.
  const pools = new Map<string, ResourceSpawn[]>();
  for (const sc of SCHEMATICS) {
    for (const slot of sc.slots) {
      if (pools.has(slot.classId)) continue;
      pools.set(
        slot.classId,
        galaxy.spawns.filter((s) => isDescendant(s.classId, slot.classId)),
      );
    }
  }

  for (let i = 0; i < 400; i++) {
    const sc = SCHEMATICS[i % SCHEMATICS.length];
    const resolved: Record<string, ResourceSpawn | null> = {};
    let complete = true;
    for (const slot of sc.slots) {
      const pool = pools.get(slot.classId)!;
      if (!pool.length) {
        complete = false;
        break;
      }
      resolved[slot.id] = pick(pool, rng);
    }
    if (!complete) continue;
    const skill = { assembly: rng.int(20, 100), experimentation: rng.int(20, 100) };
    let st: CraftState = assemble(newSession(sc.id, rng.int(-2e9, 2e9)), sc, resolved, skill);
    sessions++;
    while (st.phase === 'experimentation' && st.pointsSpent < st.pointsTotal) {
      st = experiment(st, sc, pick(sc.lines, rng).id, rng.int(1, 3), skill);
    }
    if (st.phase === 'destroyed') destroyed++;
    for (const p of sc.properties) {
      const q = propertyQuality(p, resolved);
      const cap = ceilingFor(q);
      const v = st.pct[p.id] ?? 0;
      if (v > cap + 1e-9 || v < -1e-9) breach++;
    }
  }
  out.push({
    name: 'resource ceiling invariant',
    pass: breach === 0,
    detail: `${sessions} randomized sessions · ${destroyed} destroyed · ${breach} ceiling breaches`,
  });

  // 6. Taxonomy integrity: no orphans, no cycles.
  let orphan = 0;
  for (const c of RESOURCE_CLASSES) {
    if (c.parent && !CLASS_INDEX.has(c.parent)) orphan++;
    let depth = 0;
    let cur = c.parent;
    while (cur && depth < 32) {
      cur = CLASS_INDEX.get(cur)?.parent ?? null;
      depth++;
    }
    if (depth >= 32) orphan++;
  }
  out.push({
    name: 'taxonomy integrity',
    pass: orphan === 0,
    detail: `${RESOURCE_CLASSES.length} classes · acyclic · ${orphan} orphans`,
  });

  return out;
}

function pick<T>(arr: readonly T[], rng: Rng): T {
  return arr[rng.int(0, arr.length - 1)];
}
