/**
 * Deterministic PRNG. mulberry32.
 *
 * Rationale: the galaxy must be reproducible from a 32-bit seed so that spawn
 * tables are regression-testable and a craft session can be replayed exactly.
 * We do NOT use Math.random() anywhere in /engine. The generator state is a
 * plain int32, which means it can be carried inside serializable state and
 * ported to C++/UE5 verbatim (Math.imul == 32-bit signed multiply).
 */

export interface RngStep {
  /** next generator state */
  s: number;
  /** uniform [0,1) */
  v: number;
}

/** Pure step function. No allocation beyond the returned tuple. */
export function step(state: number): RngStep {
  const t = (state + 0x6d2b79f5) | 0;
  let r = Math.imul(t ^ (t >>> 15), 1 | t);
  r = (r + Math.imul(r ^ (r >>> 7), 61 | r)) ^ r;
  return { s: t, v: ((r ^ (r >>> 14)) >>> 0) / 4294967296 };
}

/** FNV-1a. String seed -> int32. Stable across platforms. */
export function hashSeed(str: string): number {
  let h = 0x811c9dc5;
  for (let i = 0; i < str.length; i++) {
    h ^= str.charCodeAt(i);
    h = Math.imul(h, 0x01000193);
  }
  return h | 0;
}

/** Stateful cursor for generation code. Thin wrapper over step(). */
export class Rng {
  private s: number;

  constructor(seed: number | string) {
    this.s = typeof seed === 'string' ? hashSeed(seed) : seed | 0;
  }

  get state(): number {
    return this.s;
  }

  next(): number {
    const r = step(this.s);
    this.s = r.s;
    return r.v;
  }

  /** [min, max) */
  range(min: number, max: number): number {
    return min + this.next() * (max - min);
  }

  /** integer in [min, max] inclusive */
  int(min: number, max: number): number {
    return Math.floor(this.range(min, max + 1));
  }

  pick<T>(arr: readonly T[]): T {
    return arr[Math.floor(this.next() * arr.length)];
  }

  chance(p: number): boolean {
    return this.next() < p;
  }

  /**
   * Skewed roll. exp > 1 biases low (common), exp < 1 biases high (rare/good).
   * This is the single knob that makes a 950 OQ spawn feel like an event.
   */
  curve(min: number, max: number, exp: number): number {
    return min + (max - min) * Math.pow(this.next(), exp);
  }
}
