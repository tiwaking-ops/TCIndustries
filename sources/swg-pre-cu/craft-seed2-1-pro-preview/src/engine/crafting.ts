// ============================================================================
// Crafting Engine — Assembly + Experimentation
// Pure functions. No React. No side effects beyond RNG. Port verbatim to C++.
// The only thing that changes when we move to UE5 is the RNG source and the
// inventory persistence.
// ============================================================================

import type {
  Schematic, SchematicSlot, Resource,
  SlotAssignment, ExperimentLogEntry, ExperimentOutcome,
  ExperimentProgress, CraftedItem, StatId,
} from "./types";

// Working copies used during the active craft — fields are mutable because
// the player is actively iterating on resource selection and experiment
// progress before committing the final item.
type MutableAssignment = {
  -readonly [K in keyof SlotAssignment]: SlotAssignment[K];
};

type MutableProgress = {
  -readonly [K in keyof ExperimentProgress]: ExperimentProgress[K];
};
import { RESOURCE_CLASSES, isSubtypeOf } from "./data/resourceTree";
import { rollDie } from "./rng";

// ============================================================================
// Assembly validation
// ============================================================================

export interface AssignmentError {
  slotId: string;
  message: string;
}

export function validateSlot(
  slot: SchematicSlot,
  assignment: SlotAssignment,
  resource: Resource | null,
): AssignmentError | null {
  if (!resource || assignment.resourceId === null) {
    return { slotId: slot.id, message: `Slot "${slot.label}" is empty.` };
  }
  if (!isSubtypeOf(resource.classId, slot.requiredClassId)) {
    const cls = RESOURCE_CLASSES.get(resource.classId);
    return {
      slotId: slot.id,
      message: `Slot requires ${RESOURCE_CLASSES.get(slot.requiredClassId)?.name}, got ${cls?.name}.`,
    };
  }
  if (assignment.unitsUsed < slot.requiredUnits) {
    return {
      slotId: slot.id,
      message: `Need ${slot.requiredUnits} units, assigned ${assignment.unitsUsed}.`,
    };
  }
  return null;
}

// ============================================================================
// Weighted resource average for a given property.
// ============================================================================

/**
 * Compute the weighted average of a resource's stats for a set of stat weights.
 * Weights are *not* required to sum to 1 — we normalize so that the result is
 * independent of weight magnitudes (SWG quirk: OQ 0.5, PE 0.5 is same avg as
 * OQ 50, PE 50). Missing weights are treated as 0.
 */
export function weightedResourceScore(
  resource: Resource,
  weights: Readonly<Partial<Record<StatId, number>>>,
): number {
  let totalWeight = 0;
  let weightedSum = 0;
  for (const stat of Object.keys(weights) as StatId[]) {
    const w = weights[stat] ?? 0;
    if (w <= 0) continue;
    weightedSum += resource.stats[stat] * w;
    totalWeight += w;
  }
  if (totalWeight === 0) return 0;
  return weightedSum / totalWeight;
}

/**
 * Compute the per-property resource cap (0..1 scalar) across all filled slots.
 * This is the ceiling that experimentation can never exceed. It encodes the
 * entire Pre-CU "garbage in, garbage out" philosophy.
 */
export function computeResourceCaps(
  schematic: Schematic,
  assignments: ReadonlyArray<SlotAssignment>,
  resourcesById: ReadonlyMap<string, Resource>,
): ReadonlyMap<string, number> {
  const caps = new Map<string, number>();
  for (const prop of schematic.experimentalProps) {
    let totalContribution = 0;
    let weightedScore = 0;
    for (const slot of schematic.slots) {
      const contrib = slot.experimentalContribution[prop.id] ?? 0;
      if (contrib <= 0) continue;
      const a = assignments.find(x => x.slotId === slot.id);
      if (!a || !a.resourceId) continue;
      const r = resourcesById.get(a.resourceId);
      if (!r) continue;
      const score = weightedResourceScore(r, prop.statWeights);
      weightedScore += score * contrib;
      totalContribution += contrib;
    }
    const cap = totalContribution > 0 ? weightedScore / totalContribution / 1000 : 0;
    caps.set(prop.id, Math.max(0, Math.min(1, cap)));
  }
  return caps;
}

// ============================================================================
// Skill thresholds for experimentation d100 roll.
// Ranges are inclusive lower-bounds on a d100. Outcome is resolved from top
// (amazing) down. Default = master craftsman (~80% success/great/amazing).
// ============================================================================

export interface SkillProfile {
  // roll >= amazingThreshold -> amazing success
  readonly amazingThreshold: number;   // e.g. 95 for master
  // roll >= greatThreshold -> great
  readonly greatThreshold: number;     // e.g. 75
  // roll >= successThreshold -> success
  readonly successThreshold: number;   // e.g. 30
  // roll >= critFailThreshold -> failure
  // below -> critical failure
  readonly critFailThreshold: number;  // e.g. 5
}

export const NOVICE_SKILL: SkillProfile = {
  amazingThreshold: 99,
  greatThreshold: 85,
  successThreshold: 35,
  critFailThreshold: 15,
};

export const MASTER_SKILL: SkillProfile = {
  amazingThreshold: 92,
  greatThreshold: 65,
  successThreshold: 22,
  critFailThreshold: 4,
};

export function resolveOutcome(roll: number, skill: SkillProfile): ExperimentOutcome {
  if (roll >= skill.amazingThreshold) return "amazing";
  if (roll >= skill.greatThreshold) return "great";
  if (roll >= skill.successThreshold) return "success";
  if (roll >= skill.critFailThreshold) return "failure";
  return "critical";
}

// Box deltas per point spent. In Pre-CU you could dump up to 5 points in one
// attempt — high risk, high reward.
export function boxDelta(outcome: ExperimentOutcome, points: number): number {
  switch (outcome) {
    case "amazing":  return points;                  // 1:1 perfect
    case "great":    return Math.ceil(points / 2);   // half progress
    case "success":  return 0;                       // no gain no pain
    case "failure":  return 0;                       // nothing happens
    case "critical": return -Math.ceil(points / 2);  // real penalty
  }
}

// ============================================================================
// Crafting session factory
// ============================================================================

export interface CraftingSession {
  readonly schematic: Schematic;
  assignments: MutableAssignment[];
  progress: MutableProgress[];
  pointsRemaining: number;
  log: ExperimentLogEntry[];
  skill: SkillProfile;
  finalized: boolean;
}

export function initSession(schematic: Schematic, skill: SkillProfile = MASTER_SKILL): CraftingSession {
  return {
    schematic,
    assignments: schematic.slots.map(s => ({ slotId: s.id, resourceId: null, unitsUsed: 0 })),
    progress: schematic.experimentalProps.map(p => ({ propertyId: p.id, boxes: 0, maxBoxes: 10 })),
    pointsRemaining: schematic.experimentPoints,
    log: [],
    skill,
    finalized: false,
  };
}

export function assignToSlot(
  session: CraftingSession,
  slotId: string,
  resourceId: string,
  units: number,
): void {
  const a = session.assignments.find(x => x.slotId === slotId);
  if (!a) return;
  a.resourceId = resourceId;
  a.unitsUsed = units;
}

export function clearSlot(session: CraftingSession, slotId: string): void {
  const a = session.assignments.find(x => x.slotId === slotId);
  if (!a) return;
  a.resourceId = null;
  a.unitsUsed = 0;
}

// Pre-experimentation assembly phase. In Pre-CU this is a separate d100 roll
// that gives you a "starting quality" modifier. For simplicity we treat it as
// a single bonus/penalty to all starting boxes. A master gets +1 starting box
// on every prop, essentially skipping the "warmed up" phase.
export function runAssemblyPhase(session: CraftingSession): void {
  // Master assembly bonus: +1 box across the board; capped at 1.
  for (const p of session.progress) {
    p.boxes = Math.min(1, p.maxBoxes);
  }
}

/**
 * Execute one experimentation attempt. Returns the log entry. Does NOT prevent
 * "wasting" points on already-capped properties — that's a player decision in
 * Pre-CU (you can over-experiment and crit-fail your way back down on purpose,
// though it's almost never correct).
 */
export function experiment(
  session: CraftingSession,
  propertyId: string,
  pointsToSpend: number,
  rng: () => number,
): ExperimentLogEntry | null {
  if (session.finalized) return null;
  const spend = Math.max(1, Math.min(pointsToSpend, session.pointsRemaining, 5));
  if (spend <= 0) return null;

  const prop = session.schematic.experimentalProps.find(p => p.id === propertyId);
  if (!prop) return null;

  const prog = session.progress.find(p => p.propertyId === propertyId);
  if (!prog) return null;

  const dieRoll = rollDie(rng, 100);
  const outcome = resolveOutcome(dieRoll, session.skill);
  const delta = boxDelta(outcome, spend);
  prog.boxes = Math.max(0, Math.min(prog.maxBoxes, prog.boxes + delta));

  session.pointsRemaining -= spend;
  const entry: ExperimentLogEntry = {
    attempt: session.log.length + 1,
    pointsSpent: spend,
    propertyId,
    outcome,
    boxDelta: delta,
    dieRoll,
  };
  session.log.push(entry);
  return entry;
}

/**
 * Auto-experiment helper: greedily spends remaining points evenly across all
 * properties, simulating a "best effort" crafter. For demo / bulk-testing.
 */
export function autoExperiment(session: CraftingSession, rng: () => number): void {
  while (session.pointsRemaining > 0) {
    // Find the prop with the lowest box ratio.
    const prog = [...session.progress]
      .sort((a, b) => (a.boxes / a.maxBoxes) - (b.boxes / b.maxBoxes))[0];
    if (!prog) break;
    // Spend up to 3 points per attempt — balanced risk.
    const spend = Math.min(3, session.pointsRemaining);
    experiment(session, prog.propertyId, spend, rng);
    if (session.progress.every(p => p.boxes >= p.maxBoxes)) break;
  }
}

// ============================================================================
// Finalization — convert progress into a concrete CraftedItem.
// ============================================================================

export function finalize(
  session: CraftingSession,
  resourcesById: ReadonlyMap<string, Resource>,
  rng: () => number,
): CraftedItem {
  const caps = computeResourceCaps(session.schematic, session.assignments, resourcesById);
  const finalPropertyValues: Record<string, number> = {};

  for (const prop of session.schematic.experimentalProps) {
    const prog = session.progress.find(p => p.propertyId === prop.id);
    const boxRatio = prog ? prog.boxes / prog.maxBoxes : 0;
    const cap = caps.get(prop.id) ?? 0;
    // Attenuate the cap by boxRatio: boxRatio determines how much of the
    // resource cap you actually realize. Then lerp min..max by that.
    const realized = cap * boxRatio;
    // Standard lerp: minValue = worst case, maxValue = best case. For stats
    // where lower is better (attack_speed, power_efficiency), the schematic
    // simply encodes maxValue < minValue, and this lerp naturally yields a
    // decreasing value as realized → 1.
    const value = prop.minValue + (prop.maxValue - prop.minValue) * realized;
    finalPropertyValues[prop.id] = Math.round(value * 100) / 100;
  }

  // Compute an overall item quality scalar (0..1) for display sorting.
  // For "inverted" stats (maxValue < minValue, e.g. attack speed = 5..2)
  // the ratio (v - min) / range is naturally negative-to-1 when normalized
  // against the ABSOLUTE quality — so we flip the normal case.
  const propRatios: number[] = session.schematic.experimentalProps.map(p => {
    const v = finalPropertyValues[p.id];
    const min = Math.min(p.minValue, p.maxValue);
    const max = Math.max(p.minValue, p.maxValue);
    const range = max - min;
    if (range === 0) return 1;
    // Position from worst (0) to best (1):
    // If minValue > maxValue, lower is better; else higher is better.
    if (p.minValue > p.maxValue) {
      // inverted: best = min (small), worst = max (large)
      return (p.minValue - v) / range;
    } else {
      return (v - p.minValue) / range;
    }
  });
  const overall = propRatios.reduce((a, b) => a + b, 0) / propRatios.length;

  const usedResources = session.assignments
    .filter(a => a.resourceId)
    .map(a => ({ resourceId: a.resourceId!, units: a.unitsUsed }));

  session.finalized = true;
  return {
    id: `item_${Math.floor(rng() * 1e9)}`,
    schematicId: session.schematic.id,
    serialNumber: Math.floor(rng() * 1000000),
    resourcesUsed: usedResources,
    finalPropertyValues,
    overallQuality: overall,
    experimentLog: [...session.log],
  };
}
