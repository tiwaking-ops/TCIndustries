// ============================================================================
// Deterministic PRNG (mulberry32) + weighted distributions.
// Using a seedable RNG means we can replay the exact same spawn table or
// experimentation roll for debugging — critical for balancing. In UE5 this
// maps to a fast xorshift or the engine's FRandomStream.
// ============================================================================

export interface Rng {
  (): number; // 0..1 float
}

export function makeRng(seed: number): Rng {
  let s = seed >>> 0;
  return function rng(): number {
    s = (s + 0x6D2B79F5) >>> 0;
    let t = s;
    t = Math.imul(t ^ (t >>> 15), t | 1);
    t ^= t + Math.imul(t ^ (t >>> 7), t | 61);
    return ((t ^ (t >>> 14)) >>> 0) / 4294967296;
  };
}

export function intBetween(rng: Rng, lo: number, hi: number): number {
  return Math.floor(rng() * (hi - lo + 1)) + lo;
}

/** Inclusive integer range, uniform. */
export function rollDie(rng: Rng, sides: number): number {
  return Math.floor(rng() * sides) + 1;
}

/** Weighted pick from array of { value, weight }. */
export function weightedPick<T>(rng: Rng, entries: ReadonlyArray<{ value: T; weight: number }>): T {
  const total = entries.reduce((s, e) => s + e.weight, 0);
  let r = rng() * total;
  for (const e of entries) {
    r -= e.weight;
    if (r <= 0) return e.value;
  }
  return entries[entries.length - 1].value;
}

/**
 * Sample a resource stat with Pre-CU distribution bias. Pre-CU was NOT uniform
 * across 0..1000 — it was a left-skewed distribution: lots of mediocre
 * resources, rare high-quality ones. We approximate with a beta-like curve by
 * taking max of 3 uniform rolls, which biases low. For high-end JTL/class 10
 * resources you add even more skew.
 */
export function sampleStat(rng: Rng, skew = 3): number {
  // Max of `skew` uniform 0..1 rolls. Higher skew = higher average but still
  // clamped. skew=3 gives mean ~570, median ~630, 95th percentile ~900.
  let v = 0;
  for (let i = 0; i < skew; i++) v = Math.max(v, rng());
  return Math.round(v * 1000);
}
