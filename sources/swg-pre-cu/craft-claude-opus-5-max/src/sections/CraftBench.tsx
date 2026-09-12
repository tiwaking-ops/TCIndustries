import { useMemo, useState } from 'react';
import {
  CLASS_INDEX,
  STAT_KEYS,
  capPct,
  isDescendant,
  type ResourceSpawn,
  type StatKey,
} from '@/engine/resources';
import { SCHEMATICS, type ExperimentalProperty, type Schematic } from '@/engine/schematics';
import {
  assemble,
  ceilingFor,
  experiment,
  finalize,
  formatValue,
  itemRating,
  newSession,
  pointBudget,
  propertyQuality,
  type CraftState,
} from '@/engine/craft';
import { hashSeed } from '@/engine/rng';
import { Panel, Tag, heat } from '@/components/ui';
import { cn } from '@/utils/cn';

/** Aggregate every (property, group) weight that reads a given slot. */
function slotWeights(schematic: Schematic, slotId: string): Partial<Record<StatKey, number>> {
  const w: Partial<Record<StatKey, number>> = {};
  for (const p of schematic.properties) {
    for (const g of p.groups) {
      if (g.slotId !== slotId) continue;
      for (const sw of g.stats) w[sw.stat] = (w[sw.stat] ?? 0) + sw.weight * g.weight;
    }
  }
  return w;
}

function scoreResource(res: ResourceSpawn, w: Partial<Record<StatKey, number>>): number {
  let acc = 0;
  let tot = 0;
  for (const k of STAT_KEYS) {
    const weight = w[k];
    if (!weight) continue;
    acc += (res.stats[k] ?? 0) * weight;
    tot += weight;
  }
  return tot ? acc / tot : 0;
}

export function CraftBench({ active }: { active: ResourceSpawn[] }) {
  const [schematicId, setSchematicId] = useState(SCHEMATICS[0].id);
  const [assignment, setAssignment] = useState<Record<string, string | null>>({});
  const [skill, setSkill] = useState({ assembly: 62, experimentation: 58 });
  const [alloc, setAlloc] = useState<Record<string, number>>({});
  const [session, setSession] = useState<CraftState>(() =>
    newSession(SCHEMATICS[0].id, hashSeed('bench-0')),
  );
  const [runs, setRuns] = useState(0);

  const schematic = useMemo(
    () => SCHEMATICS.find((s) => s.id === schematicId)!,
    [schematicId],
  );

  const byId = useMemo(() => new Map(active.map((r) => [r.id, r])), [active]);

  const candidates = useMemo(() => {
    const m: Record<string, ResourceSpawn[]> = {};
    for (const slot of schematic.slots) {
      const w = slotWeights(schematic, slot.id);
      m[slot.id] = active
        .filter((r) => isDescendant(r.classId, slot.classId))
        .sort((a, b) => scoreResource(b, w) - scoreResource(a, w));
    }
    return m;
  }, [active, schematic]);

  const resolved = useMemo(() => {
    const m: Record<string, ResourceSpawn | null> = {};
    for (const slot of schematic.slots) m[slot.id] = byId.get(assignment[slot.id] ?? '') ?? null;
    return m;
  }, [assignment, byId, schematic]);

  const filled = schematic.slots.every((s) => resolved[s.id]);
  const preview = useMemo(() => {
    const m: Record<string, number> = {};
    for (const p of schematic.properties) m[p.id] = propertyQuality(p, resolved);
    return m;
  }, [resolved, schematic]);

  const remaining = session.pointsTotal - session.pointsSpent;
  const inDesign = session.phase === 'design';
  const live = session.phase === 'experimentation';

  function reset(nextSchematic = schematicId) {
    setSession(newSession(nextSchematic, hashSeed(`bench-${runs + 1}`)));
    setRuns((r) => r + 1);
    setAlloc({});
  }

  function autoAssign() {
    const next: Record<string, string | null> = {};
    for (const slot of schematic.slots) next[slot.id] = candidates[slot.id]?.[0]?.id ?? null;
    setAssignment(next);
    reset();
  }

  function pickSchematic(id: string) {
    setSchematicId(id);
    setAssignment({});
    setSession(newSession(id, hashSeed(`bench-${runs + 1}`)));
    setRuns((r) => r + 1);
    setAlloc({});
  }

  const budget = pointBudget(schematic, skill);
  const rating = itemRating(schematic, live || session.phase === 'complete' ? session.pct : preview);

  return (
    <div className="space-y-4">
      <Panel
        title="Crafting Bench"
        subtitle="assembly → experimentation → finalize · pure FSM, PRNG state carried in-band"
        right={
          <div className="flex flex-wrap items-center gap-2">
            <Tag tone={session.phase === 'destroyed' ? 'rose' : live ? 'amber' : 'zinc'}>
              phase: {session.phase}
            </Tag>
            <Tag tone="cyan">rating {rating.toFixed(1)}</Tag>
          </div>
        }
      >
        <div className="flex flex-wrap items-center gap-2">
          {SCHEMATICS.map((s) => (
            <button
              key={s.id}
              onClick={() => pickSchematic(s.id)}
              className={cn(
                'rounded border px-3 py-1.5 text-left transition',
                s.id === schematicId
                  ? 'border-amber-500/70 bg-amber-500/10'
                  : 'border-zinc-800 bg-zinc-900/50 hover:border-zinc-600',
              )}
            >
              <div
                className={cn(
                  'text-[12px] font-semibold',
                  s.id === schematicId ? 'text-amber-200' : 'text-zinc-300',
                )}
              >
                {s.name}
              </div>
              <div className="font-mono text-[10px] text-zinc-500">
                cx {s.complexity} · {s.slots.length} slots · {s.properties.length} props ·{' '}
                {s.xpType}
              </div>
            </button>
          ))}
          <div className="ml-auto flex items-center gap-4">
            <SkillSlider
              label="assembly mod"
              value={skill.assembly}
              onChange={(v) => setSkill((s) => ({ ...s, assembly: v }))}
            />
            <SkillSlider
              label="experiment mod"
              value={skill.experimentation}
              onChange={(v) => setSkill((s) => ({ ...s, experimentation: v }))}
            />
            <div className="rounded border border-zinc-800 bg-black/40 px-2 py-1 text-center">
              <div className="font-mono text-[9px] uppercase tracking-widest text-zinc-600">
                points
              </div>
              <div className="font-mono text-sm text-amber-300">{budget}</div>
            </div>
          </div>
        </div>
      </Panel>

      <div className="grid gap-4 xl:grid-cols-[1fr_1fr_0.9fr]">
        {/* ---- SLOTS ---- */}
        <Panel
          title="Resource Slots"
          subtitle="class-subtree constrained · ranked by aggregated group weights"
          right={
            <button
              onClick={autoAssign}
              className="rounded border border-cyan-700/60 bg-cyan-500/10 px-2 py-1 font-mono text-[10px] uppercase tracking-wider text-cyan-300 transition hover:bg-cyan-500/20"
            >
              solve best
            </button>
          }
        >
          <div className="space-y-3">
            {schematic.slots.map((slot) => {
              const list = candidates[slot.id] ?? [];
              const res = resolved[slot.id];
              const w = slotWeights(schematic, slot.id);
              const lanes = STAT_KEYS.filter((k) => w[k]);
              return (
                <div key={slot.id} className="rounded border border-zinc-800 bg-black/40 p-2.5">
                  <div className="flex items-baseline justify-between gap-2">
                    <span className="text-[12px] font-semibold text-zinc-200">{slot.label}</span>
                    <span className="font-mono text-[10px] text-zinc-500">
                      {slot.units}u · {CLASS_INDEX.get(slot.classId)?.name}
                    </span>
                  </div>
                  <div className="mt-1 flex flex-wrap gap-1">
                    {lanes.map((k) => (
                      <span
                        key={k}
                        className="rounded bg-zinc-800/70 px-1 font-mono text-[9px] text-zinc-400"
                      >
                        {k}·{w[k]}
                      </span>
                    ))}
                  </div>
                  <select
                    disabled={!inDesign}
                    value={assignment[slot.id] ?? ''}
                    onChange={(e) => {
                      setAssignment((a) => ({ ...a, [slot.id]: e.target.value || null }));
                      reset();
                    }}
                    className="mt-2 w-full rounded border border-zinc-700 bg-zinc-950 px-2 py-1 font-mono text-[11px] text-zinc-200 outline-none focus:border-amber-500 disabled:opacity-40"
                  >
                    <option value="">— empty ({list.length} candidates) —</option>
                    {list.slice(0, 40).map((r) => (
                      <option key={r.id} value={r.id}>
                        {r.name} · {CLASS_INDEX.get(r.classId)?.name} · score{' '}
                        {scoreResource(r, w).toFixed(0)}
                      </option>
                    ))}
                  </select>
                  {res && (
                    <div className="mt-1.5 flex flex-wrap gap-x-2.5 gap-y-0.5">
                      {lanes.map((k) => {
                        const v = res.stats[k];
                        const h = heat(v === undefined ? 0 : capPct(res.classId, k, v));
                        return (
                          <span key={k} className="font-mono text-[10px]">
                            <span className="text-zinc-600">{k}</span>{' '}
                            <span className={v === undefined ? 'text-rose-500' : h.text}>
                              {v ?? 'n/a'}
                            </span>
                          </span>
                        );
                      })}
                    </div>
                  )}
                </div>
              );
            })}

            <div className="flex gap-2 pt-1">
              <button
                disabled={!filled || !inDesign}
                onClick={() => setSession((s) => assemble(s, schematic, resolved, skill))}
                className="flex-1 rounded border border-amber-600/70 bg-amber-500/15 px-3 py-2 font-mono text-[11px] uppercase tracking-widest text-amber-200 transition hover:bg-amber-500/25 disabled:cursor-not-allowed disabled:border-zinc-800 disabled:bg-zinc-900 disabled:text-zinc-600"
              >
                assemble
              </button>
              <button
                onClick={() => reset()}
                className="rounded border border-zinc-700 px-3 py-2 font-mono text-[11px] uppercase tracking-widest text-zinc-400 transition hover:border-zinc-500 hover:text-zinc-200"
              >
                reset
              </button>
            </div>
          </div>
        </Panel>

        {/* ---- PROPERTIES ---- */}
        <Panel
          title="Experimental Properties"
          subtitle="fill = achieved · notch = resource ceiling · tick = assembly base"
        >
          <div className="space-y-3.5">
            {schematic.properties.map((p) => {
              const q = preview[p.id] ?? 0;
              const pct = live || session.phase === 'complete' ? (session.pct[p.id] ?? 0) : q;
              const cap = session.ceiling[p.id] ?? ceilingFor(q);
              const base = session.base[p.id] ?? q;
              return <PropRow key={p.id} prop={p} pct={pct} cap={cap} base={base} dim={inDesign} />;
            })}
          </div>
          <div className="mt-4 border-t border-zinc-800 pt-3 font-mono text-[10px] leading-relaxed text-zinc-600">
            quality = Σ<sub>g</sub> w<sub>g</sub> · (Σ<sub>s</sub> w<sub>s</sub> · stat<sub>s</sub> /
            Σw<sub>s</sub>) / Σw<sub>g</sub> / 1000
            <br />
            ceiling = 0.5 + 0.5 · quality &nbsp;→&nbsp; bad rock is a hard wall, not a slow climb
          </div>
        </Panel>

        {/* ---- EXPERIMENTATION ---- */}
        <div className="space-y-4">
          <Panel
            title="Experimentation"
            subtitle={live ? `${remaining} / ${session.pointsTotal} points remaining` : 'assemble first'}
          >
            {live ? (
              <div className="space-y-2.5">
                {schematic.lines.map((line) => {
                  const spend = Math.min(alloc[line.id] ?? 1, Math.max(1, remaining));
                  const k = skill.experimentation - schematic.complexity * 1.6 - (spend - 1) * 7;
                  const pDestroy = Math.max(0, Math.min(1, (10 - k) / 100));
                  const pSetback = Math.max(0, Math.min(1, (32 - k) / 100));
                  return (
                    <div key={line.id} className="rounded border border-zinc-800 bg-black/40 p-2.5">
                      <div className="text-[12px] font-semibold text-zinc-200">{line.label}</div>
                      <div className="font-mono text-[10px] text-zinc-500">
                        drives {line.properties.join(', ')}
                      </div>
                      <div className="mt-2 flex items-center gap-2">
                        <div className="flex items-center rounded border border-zinc-700">
                          <button
                            onClick={() =>
                              setAlloc((a) => ({ ...a, [line.id]: Math.max(1, spend - 1) }))
                            }
                            className="px-2 py-0.5 font-mono text-zinc-400 hover:text-amber-300"
                          >
                            −
                          </button>
                          <span className="w-6 text-center font-mono text-[12px] text-amber-300">
                            {spend}
                          </span>
                          <button
                            onClick={() =>
                              setAlloc((a) => ({
                                ...a,
                                [line.id]: Math.min(remaining, spend + 1),
                              }))
                            }
                            className="px-2 py-0.5 font-mono text-zinc-400 hover:text-amber-300"
                          >
                            +
                          </button>
                        </div>
                        <button
                          onClick={() =>
                            setSession((s) => experiment(s, schematic, line.id, spend, skill))
                          }
                          className="flex-1 rounded border border-cyan-700/60 bg-cyan-500/10 px-2 py-1 font-mono text-[10px] uppercase tracking-wider text-cyan-300 transition hover:bg-cyan-500/20"
                        >
                          commit
                        </button>
                      </div>
                      <div className="mt-1.5 flex gap-3 font-mono text-[9.5px]">
                        <span className="text-rose-400">
                          destroy {(pDestroy * 100).toFixed(1)}%
                        </span>
                        <span className="text-amber-400/80">
                          setback {(pSetback * 100).toFixed(1)}%
                        </span>
                      </div>
                    </div>
                  );
                })}
                <button
                  onClick={() => setSession((s) => finalize(s, schematic))}
                  className="w-full rounded border border-emerald-700/60 bg-emerald-500/10 px-3 py-2 font-mono text-[11px] uppercase tracking-widest text-emerald-300 transition hover:bg-emerald-500/20"
                >
                  finalize prototype
                </button>
              </div>
            ) : (
              <div className="py-6 text-center font-mono text-[11px] text-zinc-600">
                {session.phase === 'destroyed'
                  ? 'session terminated — resources lost'
                  : session.phase === 'complete'
                    ? 'item committed to inventory'
                    : filled
                      ? 'slots ready — run assembly'
                      : 'assign every slot to continue'}
              </div>
            )}
          </Panel>

          <Panel title="Session Log" bodyClassName="p-0">
            <div className="max-h-56 overflow-auto p-3">
              {session.log.length === 0 && (
                <div className="py-4 text-center font-mono text-[11px] text-zinc-700">
                  no transitions recorded
                </div>
              )}
              {session.log.map((l) => (
                <div key={l.seq} className="border-b border-zinc-900 py-1.5 last:border-0">
                  <div
                    className={cn(
                      'font-mono text-[11px]',
                      l.kind === 'fatal'
                        ? 'text-rose-400'
                        : l.kind === 'loss'
                          ? 'text-amber-400'
                          : l.kind === 'gain'
                            ? 'text-emerald-300'
                            : l.kind === 'done'
                              ? 'text-cyan-300'
                              : 'text-zinc-300',
                    )}
                  >
                    <span className="text-zinc-700">
                      [{String(l.seq).padStart(3, '0')}]
                    </span>{' '}
                    {l.text}
                  </div>
                  {l.detail && (
                    <div className="pl-9 font-mono text-[10px] leading-snug text-zinc-600">
                      {l.detail}
                    </div>
                  )}
                </div>
              ))}
            </div>
          </Panel>
        </div>
      </div>
    </div>
  );
}

function PropRow({
  prop,
  pct,
  cap,
  base,
  dim,
}: {
  prop: ExperimentalProperty;
  pct: number;
  cap: number;
  base: number;
  dim: boolean;
}) {
  const h = heat(pct);
  return (
    <div>
      <div className="flex items-baseline justify-between gap-2">
        <span className="text-[12px] text-zinc-300">
          {prop.label}
          {prop.invert && <span className="ml-1 font-mono text-[9px] text-zinc-600">(lower=better)</span>}
        </span>
        <span className={cn('font-mono text-[13px] tabular-nums', dim ? 'text-zinc-500' : h.text)}>
          {formatValue(prop, pct)}
          <span className="ml-0.5 text-[10px] text-zinc-600">{prop.unit}</span>
        </span>
      </div>
      <div className="relative mt-1 h-2.5 overflow-hidden rounded bg-zinc-900 ring-1 ring-inset ring-zinc-800">
        <div
          className="absolute inset-y-0 left-0 bg-zinc-800/80"
          style={{ width: `${cap * 100}%` }}
        />
        <div
          className={cn('absolute inset-y-0 left-0 transition-all duration-300', h.bar, dim && 'opacity-50')}
          style={{ width: `${pct * 100}%` }}
        />
        <div
          className="absolute inset-y-0 w-px bg-zinc-100/70"
          style={{ left: `${cap * 100}%` }}
          title="resource ceiling"
        />
        <div
          className="absolute inset-y-0 w-px bg-amber-400/60"
          style={{ left: `${base * 100}%` }}
          title="assembly base"
        />
      </div>
      <div className="mt-0.5 flex justify-between font-mono text-[9px] text-zinc-700">
        <span>
          {prop.min}–{prop.max}
          {prop.unit}
        </span>
        <span>
          achieved {(pct * 100).toFixed(1)}% · ceiling {(cap * 100).toFixed(1)}%
        </span>
      </div>
    </div>
  );
}

function SkillSlider({
  label,
  value,
  onChange,
}: {
  label: string;
  value: number;
  onChange: (v: number) => void;
}) {
  return (
    <label className="block w-32">
      <div className="flex justify-between font-mono text-[9px] uppercase tracking-widest text-zinc-600">
        <span>{label}</span>
        <span className="text-zinc-400">{value}</span>
      </div>
      <input
        type="range"
        min={0}
        max={100}
        value={value}
        onChange={(e) => onChange(Number(e.target.value))}
        className="mt-1 h-1 w-full cursor-pointer appearance-none rounded bg-zinc-800"
      />
    </label>
  );
}
