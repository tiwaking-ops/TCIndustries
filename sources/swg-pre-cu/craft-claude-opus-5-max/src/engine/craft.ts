/**
 * Craft session: a pure finite state machine.
 *
 *   design -> (assemble) -> experimentation -> (finalize) -> complete
 *                    \-> destroyed            \-> destroyed
 *
 * Every transition is `(state, input) => state`. The PRNG state lives INSIDE
 * the state object, so a session is fully serializable and replayable: dump
 * the initial state + the input list and you can reproduce any bug report
 * bit-for-bit. No hidden globals, no Math.random(), no side effects.
 */

import { step } from './rng';
import type { ResourceSpawn, StatVector } from './resources';
import type { ExperimentalProperty, Schematic } from './schematics';

export type Tier =
  | 'critical_failure'
  | 'failure'
  | 'moderate'
  | 'success'
  | 'good'
  | 'great'
  | 'amazing';

export interface TierDef {
  id: Tier;
  label: string;
  /** inclusive lower bound on the roll margin */
  margin: number;
  /** multiplier applied to the resource-derived base at assembly */
  assemblyMult: number;
  /** fraction of remaining headroom captured PER POINT during experimentation */
  gain: number;
  tone: 'bad' | 'weak' | 'ok' | 'good' | 'peak';
}

/** Ordered high -> low. First match wins. Exhaustive by construction. */
export const TIERS: readonly TierDef[] = [
  { id: 'amazing', label: 'Amazing Success', margin: 45, assemblyMult: 1.1, gain: 0.24, tone: 'peak' },
  { id: 'great', label: 'Great Success', margin: 30, assemblyMult: 1.06, gain: 0.18, tone: 'good' },
  { id: 'good', label: 'Good Success', margin: 15, assemblyMult: 1.03, gain: 0.14, tone: 'good' },
  { id: 'success', label: 'Success', margin: 0, assemblyMult: 1.0, gain: 0.1, tone: 'ok' },
  { id: 'moderate', label: 'Moderate Success', margin: -18, assemblyMult: 0.92, gain: 0.06, tone: 'weak' },
  { id: 'failure', label: 'Failure', margin: -40, assemblyMult: 0.74, gain: -0.05, tone: 'bad' },
  { id: 'critical_failure', label: 'Critical Failure', margin: -Infinity, assemblyMult: 0, gain: -1, tone: 'bad' },
];

export function tierFor(margin: number): TierDef {
  for (const t of TIERS) if (margin >= t.margin) return t;
  return TIERS[TIERS.length - 1];
}

export type CraftPhase = 'design' | 'experimentation' | 'complete' | 'destroyed';

export interface LogLine {
  seq: number;
  kind: 'info' | 'roll' | 'gain' | 'loss' | 'fatal' | 'done';
  text: string;
  detail?: string;
}

export interface CraftState {
  schematicId: string;
  phase: CraftPhase;
  rngState: number;
  /** slotId -> spawn id */
  assignment: Record<string, string | null>;
  assemblyTier: Tier | null;
  /** propId -> resource-derived quality in [0,1] (frozen at assembly) */
  base: Record<string, number>;
  /** propId -> hard ceiling implied by resources */
  ceiling: Record<string, number>;
  /** propId -> current achieved fraction in [0,1] */
  pct: Record<string, number>;
  pointsTotal: number;
  pointsSpent: number;
  seq: number;
  log: LogLine[];
}

export interface CrafterSkill {
  /** assembly skill mod, 0..100 */
  assembly: number;
  /** experimentation skill mod, 0..100 */
  experimentation: number;
}

// ---- quality math --------------------------------------------------------

const clamp01 = (n: number) => (n < 0 ? 0 : n > 1 ? 1 : n);

/**
 * Weighted sum over (slot, stat). Returns [0,1]. Missing lanes contribute 0 —
 * this is intentional: putting Inert Gas (no PE) into a PE-weighted group is a
 * legal but terrible decision, and the model must let the player make it.
 */
export function propertyQuality(
  prop: ExperimentalProperty,
  resolved: Record<string, ResourceSpawn | null>,
): number {
  let acc = 0;
  let wsum = 0;
  for (const g of prop.groups) {
    const res = resolved[g.slotId];
    if (!res) continue;
    const stats: StatVector = res.stats;
    let sAcc = 0;
    let sW = 0;
    for (const sw of g.stats) {
      sAcc += (stats[sw.stat] ?? 0) * sw.weight;
      sW += sw.weight;
    }
    acc += (sW ? sAcc / sW : 0) * g.weight;
    wsum += g.weight;
  }
  if (!wsum) return 0;
  return clamp01(acc / wsum / 1000);
}

/**
 * Resources set the CEILING; experimentation only walks you toward it.
 * You cannot experiment your way out of bad rock. This is the load-bearing
 * economic rule of the entire game.
 */
export function ceilingFor(quality: number): number {
  return clamp01(0.5 + 0.5 * quality);
}

export function propertyValue(prop: ExperimentalProperty, pct: number): number {
  const t = clamp01(pct);
  return prop.invert ? prop.max - (prop.max - prop.min) * t : prop.min + (prop.max - prop.min) * t;
}

export function formatValue(prop: ExperimentalProperty, pct: number): string {
  const v = propertyValue(prop, pct);
  return v.toFixed(prop.decimals ?? 0);
}

export function pointBudget(schematic: Schematic, skill: CrafterSkill): number {
  return Math.max(2, Math.round(4 + skill.experimentation / 12 - schematic.complexity / 14));
}

// ---- transitions ---------------------------------------------------------

export function newSession(schematicId: string, seed: number): CraftState {
  return {
    schematicId,
    phase: 'design',
    rngState: seed | 0,
    assignment: {},
    assemblyTier: null,
    base: {},
    ceiling: {},
    pct: {},
    pointsTotal: 0,
    pointsSpent: 0,
    seq: 0,
    log: [],
  };
}

function push(state: CraftState, kind: LogLine['kind'], text: string, detail?: string): CraftState {
  const seq = state.seq + 1;
  return { ...state, seq, log: [{ seq, kind, text, detail }, ...state.log].slice(0, 120) };
}

function roll(state: CraftState): { state: CraftState; v: number } {
  const r = step(state.rngState);
  return { state: { ...state, rngState: r.s }, v: r.v };
}

export function assemble(
  prev: CraftState,
  schematic: Schematic,
  resolved: Record<string, ResourceSpawn | null>,
  skill: CrafterSkill,
): CraftState {
  const r = roll(prev);
  let state = r.state;
  const margin = skill.assembly - schematic.complexity * 2 + (r.v * 100 - 50);
  const tier = tierFor(margin);

  if (tier.id === 'critical_failure') {
    state = push(state, 'fatal', 'CRITICAL FAILURE — assembly', 'Resources consumed. Session terminated.');
    return { ...state, phase: 'destroyed', assemblyTier: tier.id };
  }

  const base: Record<string, number> = {};
  const ceiling: Record<string, number> = {};
  const pct: Record<string, number> = {};
  for (const p of schematic.properties) {
    const q = propertyQuality(p, resolved);
    base[p.id] = q;
    ceiling[p.id] = ceilingFor(q);
    pct[p.id] = clamp01(Math.min(q * tier.assemblyMult, ceiling[p.id]));
  }

  const points = pointBudget(schematic, skill);
  state = push(
    state,
    'roll',
    `Assembly: ${tier.label}`,
    `margin ${margin.toFixed(1)} · base x${tier.assemblyMult.toFixed(2)} · ${points} experimentation points granted`,
  );
  return {
    ...state,
    phase: 'experimentation',
    assemblyTier: tier.id,
    base,
    ceiling,
    pct,
    pointsTotal: points,
    pointsSpent: 0,
  };
}

/**
 * Risk/reward: committing N points to one line in a single pass applies a
 * -7 margin penalty per extra point. Gain compounds per point against the
 * remaining headroom, so the marginal value of point N is strictly decreasing
 * while its risk is strictly increasing. That tension IS the minigame.
 */
export function experiment(
  prev: CraftState,
  schematic: Schematic,
  lineId: string,
  points: number,
  skill: CrafterSkill,
): CraftState {
  if (prev.phase !== 'experimentation') return prev;
  const line = schematic.lines.find((l) => l.id === lineId);
  if (!line) return prev;
  const remaining = prev.pointsTotal - prev.pointsSpent;
  const spend = Math.max(1, Math.min(points, remaining));
  if (remaining <= 0) return prev;

  const r = roll(prev);
  let state = { ...r.state, pointsSpent: prev.pointsSpent + spend };
  const margin =
    skill.experimentation - schematic.complexity * 1.6 - (spend - 1) * 7 + (r.v * 100 - 50);
  const tier = tierFor(margin);

  if (tier.id === 'critical_failure') {
    state = push(state, 'fatal', `CRITICAL FAILURE — ${line.label}`, 'Item destroyed on the bench.');
    return { ...state, phase: 'destroyed' };
  }

  const pct = { ...state.pct };
  const deltas: string[] = [];
  for (const propId of line.properties) {
    const cap = state.ceiling[propId] ?? 1;
    const cur = pct[propId] ?? 0;
    let next: number;
    if (tier.gain < 0) {
      next = clamp01(cur + tier.gain * spend);
    } else {
      const headroom = Math.max(0, cap - cur);
      next = cur + headroom * (1 - Math.pow(1 - tier.gain, spend));
    }
    pct[propId] = clamp01(next);
    deltas.push(`${propId} ${(cur * 100).toFixed(1)}% → ${(pct[propId] * 100).toFixed(1)}%`);
  }

  state = push(
    state,
    tier.gain < 0 ? 'loss' : 'gain',
    `${line.label}: ${tier.label} (${spend}pt)`,
    `margin ${margin.toFixed(1)} · ${deltas.join(' · ')}`,
  );
  return { ...state, pct };
}

export function finalize(prev: CraftState, schematic: Schematic): CraftState {
  if (prev.phase !== 'experimentation') return prev;
  const summary = schematic.properties
    .map((p) => `${p.label} ${formatValue(p, prev.pct[p.id] ?? 0)}${p.unit}`)
    .join(' · ');
  const state = push(prev, 'done', `Crafted: ${schematic.name}`, summary);
  return { ...state, phase: 'complete' };
}

/** Aggregate 0-100 "item rating" — the number a player would actually compare. */
export function itemRating(schematic: Schematic, pct: Record<string, number>): number {
  const vals = schematic.properties.map((p) => pct[p.id] ?? 0);
  if (!vals.length) return 0;
  return (vals.reduce((a, b) => a + b, 0) / vals.length) * 100;
}
