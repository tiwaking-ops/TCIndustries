import { useMemo, useState } from 'react';
import {
  CLASS_INDEX,
  PLANETS,
  RESOURCE_CLASSES,
  STAT_KEYS,
  capPct,
  classChain,
  isDescendant,
  resolveCaps,
  type Galaxy,
  type ResourceSpawn,
} from '@/engine/resources';
import { Panel, Tag, heat } from '@/components/ui';
import { cn } from '@/utils/cn';

const ROOT_FILTERS = [
  { id: 'resource', label: 'All' },
  { id: 'mineral', label: 'Mineral' },
  { id: 'metal', label: 'Metal' },
  { id: 'radioactive', label: 'Radioactive' },
  { id: 'chemical', label: 'Chemical' },
  { id: 'gas', label: 'Gas' },
  { id: 'flora', label: 'Flora' },
  { id: 'creature', label: 'Creature' },
];

export function SpawnBrowser({
  galaxy,
  day,
  setDay,
  seed,
  setSeed,
  active,
}: {
  galaxy: Galaxy;
  day: number;
  setDay: (d: number) => void;
  seed: string;
  setSeed: (s: string) => void;
  active: ResourceSpawn[];
}) {
  const [root, setRoot] = useState('resource');
  const [planet, setPlanet] = useState<string>('any');
  const [q, setQ] = useState('');
  const [sel, setSel] = useState<string | null>(null);

  const rows = useMemo(() => {
    const term = q.trim().toLowerCase();
    return active
      .filter((s) => isDescendant(s.classId, root))
      .filter((s) => planet === 'any' || s.planets.some((p) => p.planet === planet))
      .filter(
        (s) =>
          !term ||
          s.name.toLowerCase().includes(term) ||
          (CLASS_INDEX.get(s.classId)?.name.toLowerCase().includes(term) ?? false),
      )
      .sort((a, b) => b.purity - a.purity);
  }, [active, root, planet, q]);

  const selected = rows.find((r) => r.id === sel) ?? null;
  const spawnable = RESOURCE_CLASSES.filter((c) => c.spawnable).length;

  return (
    <div className="space-y-4">
      <Panel
        title="Galactic Spawn Authority"
        subtitle={`${galaxy.spawns.length} intervals materialized across ${galaxy.horizon} days · ${spawnable} spawnable leaf classes`}
        right={
          <div className="flex items-center gap-2">
            <Tag tone="amber">seed {seed}</Tag>
            <Tag tone="cyan">day {day}</Tag>
          </div>
        }
      >
        <div className="grid gap-4 lg:grid-cols-[1fr_auto]">
          <div className="space-y-3">
            <div className="flex flex-wrap items-center gap-2">
              <label className="font-mono text-[10px] uppercase tracking-widest text-zinc-500">seed</label>
              <input
                value={seed}
                onChange={(e) => setSeed(e.target.value)}
                className="w-40 rounded border border-zinc-700 bg-black/60 px-2 py-1 font-mono text-[12px] text-amber-300 outline-none focus:border-amber-500"
              />
              <button
                onClick={() => setSeed(Math.random().toString(36).slice(2, 8))}
                className="rounded border border-zinc-700 bg-zinc-800/60 px-2.5 py-1 font-mono text-[11px] text-zinc-300 transition hover:border-amber-500/60 hover:text-amber-300"
              >
                reroll galaxy
              </button>
              <span className="font-mono text-[11px] text-zinc-600">
                same seed ⇒ byte-identical timeline
              </span>
            </div>

            <div className="flex items-center gap-3">
              <span className="font-mono text-[10px] uppercase tracking-widest text-zinc-500">clock</span>
              <input
                type="range"
                min={0}
                max={galaxy.horizon}
                value={day}
                onChange={(e) => setDay(Number(e.target.value))}
                className="h-1 flex-1 cursor-pointer appearance-none rounded bg-zinc-800"
              />
              <button
                onClick={() => setDay(Math.min(galaxy.horizon, day + 7))}
                className="rounded border border-zinc-700 bg-zinc-800/60 px-2.5 py-1 font-mono text-[11px] text-zinc-300 transition hover:border-cyan-500/60 hover:text-cyan-300"
              >
                +7d
              </button>
            </div>
          </div>

          <div className="grid grid-cols-3 gap-3 self-start lg:w-64">
            <Metric label="active" value={String(active.length)} tone="text-cyan-300" />
            <Metric label="in filter" value={String(rows.length)} tone="text-zinc-200" />
            <Metric
              label="peak purity"
              value={rows.length ? `${(rows[0].purity * 100).toFixed(0)}%` : '—'}
              tone="text-fuchsia-300"
            />
          </div>
        </div>

        <div className="mt-4 flex flex-wrap items-center gap-1.5 border-t border-zinc-800 pt-3">
          {ROOT_FILTERS.map((f) => (
            <button
              key={f.id}
              onClick={() => setRoot(f.id)}
              className={cn(
                'rounded border px-2 py-1 font-mono text-[10px] uppercase tracking-wider transition',
                root === f.id
                  ? 'border-amber-500/70 bg-amber-500/10 text-amber-300'
                  : 'border-zinc-800 bg-zinc-900/60 text-zinc-500 hover:text-zinc-300',
              )}
            >
              {f.label}
            </button>
          ))}
          <select
            value={planet}
            onChange={(e) => setPlanet(e.target.value)}
            className="ml-auto rounded border border-zinc-700 bg-black/60 px-2 py-1 font-mono text-[11px] text-zinc-300 outline-none"
          >
            <option value="any">any planet</option>
            {PLANETS.map((p) => (
              <option key={p} value={p}>
                {p}
              </option>
            ))}
          </select>
          <input
            value={q}
            onChange={(e) => setQ(e.target.value)}
            placeholder="filter by name / class"
            className="w-48 rounded border border-zinc-700 bg-black/60 px-2 py-1 font-mono text-[11px] text-zinc-300 outline-none placeholder:text-zinc-600 focus:border-cyan-500"
          />
        </div>
      </Panel>

      <div className="grid gap-4 xl:grid-cols-[1.6fr_1fr]">
        <Panel title="Active Pool" bodyClassName="p-0" subtitle="sorted by cap-relative purity">
          <div className="max-h-[26rem] overflow-auto">
            <table className="w-full border-collapse font-mono text-[11px]">
              <thead className="sticky top-0 z-10 bg-zinc-900/95 backdrop-blur">
                <tr className="text-left text-zinc-500">
                  <th className="px-3 py-2 font-medium">resource</th>
                  <th className="px-2 py-2 font-medium">class</th>
                  {STAT_KEYS.map((k) => (
                    <th key={k} className="px-1 py-2 text-right font-medium">
                      {k}
                    </th>
                  ))}
                  <th className="px-2 py-2 text-right font-medium">ttl</th>
                </tr>
              </thead>
              <tbody>
                {rows.map((s) => (
                  <tr
                    key={s.id}
                    onClick={() => setSel(s.id)}
                    className={cn(
                      'cursor-pointer border-t border-zinc-900 transition hover:bg-zinc-900/60',
                      sel === s.id && 'bg-amber-500/5',
                    )}
                  >
                    <td className="whitespace-nowrap px-3 py-1.5 text-zinc-200">{s.name}</td>
                    <td className="whitespace-nowrap px-2 py-1.5 text-zinc-500">
                      {CLASS_INDEX.get(s.classId)?.name}
                    </td>
                    {STAT_KEYS.map((k) => {
                      const v = s.stats[k];
                      if (v === undefined)
                        return (
                          <td key={k} className="px-1 py-1.5 text-right text-zinc-800">
                            ·
                          </td>
                        );
                      const h = heat(capPct(s.classId, k, v));
                      return (
                        <td key={k} className={cn('px-1 py-1.5 text-right tabular-nums', h.text)}>
                          {v}
                        </td>
                      );
                    })}
                    <td className="px-2 py-1.5 text-right text-zinc-600">{s.despawnTick - day}d</td>
                  </tr>
                ))}
                {!rows.length && (
                  <tr>
                    <td colSpan={13} className="px-3 py-8 text-center text-zinc-600">
                      no spawns match the predicate at day {day}
                    </td>
                  </tr>
                )}
              </tbody>
            </table>
          </div>
        </Panel>

        <Panel
          title="Struct Inspector"
          subtitle={selected ? 'the runtime object is the wire format' : 'select a row'}
          bodyClassName="p-0"
        >
          {selected ? (
            <div className="space-y-3 p-4">
              <div>
                <div className="text-lg font-semibold text-zinc-100">{selected.name}</div>
                <div className="mt-1 flex flex-wrap items-center gap-1 font-mono text-[10px] text-zinc-500">
                  {classChain(selected.classId).map((c, i) => (
                    <span key={c.id} className={i === 0 ? '' : 'before:mr-1 before:content-["›"]'}>
                      {c.name}
                    </span>
                  ))}
                </div>
              </div>

              <div className="space-y-1.5">
                {STAT_KEYS.filter((k) => selected.stats[k] !== undefined).map((k) => {
                  const v = selected.stats[k]!;
                  const cap = resolveCaps(selected.classId)[k]!;
                  const p = capPct(selected.classId, k, v);
                  const h = heat(p);
                  return (
                    <div key={k} className="grid grid-cols-[2rem_1fr_3rem] items-center gap-2">
                      <span className="font-mono text-[10px] text-zinc-500">{k}</span>
                      <div className="relative h-1.5 overflow-hidden rounded bg-zinc-800">
                        <div className={cn('h-full', h.bar)} style={{ width: `${p * 100}%` }} />
                      </div>
                      <span className={cn('text-right font-mono text-[11px] tabular-nums', h.text)}>
                        {v}
                      </span>
                      <span className="col-span-3 -mt-1 text-right font-mono text-[9px] text-zinc-700">
                        cap [{cap[0]}–{cap[1]}] · {(p * 100).toFixed(0)}% of window
                      </span>
                    </div>
                  );
                })}
              </div>

              <div className="rounded border border-zinc-800 bg-black/50 p-2.5">
                <div className="mb-1 font-mono text-[9px] uppercase tracking-widest text-zinc-600">
                  distribution
                </div>
                {selected.planets.map((p) => (
                  <div key={p.planet} className="flex items-center gap-2 py-0.5">
                    <span className="w-20 font-mono text-[11px] text-zinc-400">{p.planet}</span>
                    <div className="h-1 flex-1 overflow-hidden rounded bg-zinc-800">
                      <div className="h-full bg-cyan-500/70" style={{ width: `${p.concentration}%` }} />
                    </div>
                    <span className="font-mono text-[10px] tabular-nums text-zinc-500">
                      {p.concentration}%
                    </span>
                  </div>
                ))}
                <div className="mt-2 font-mono text-[10px] text-zinc-600">
                  window [{selected.spawnTick}, {selected.despawnTick}) · id {selected.id}
                </div>
              </div>
            </div>
          ) : (
            <div className="p-8 text-center font-mono text-[11px] text-zinc-600">
              no selection
            </div>
          )}
        </Panel>
      </div>
    </div>
  );
}

function Metric({ label, value, tone }: { label: string; value: string; tone: string }) {
  return (
    <div className="rounded border border-zinc-800 bg-black/40 px-2 py-1.5">
      <div className="font-mono text-[9px] uppercase tracking-widest text-zinc-600">{label}</div>
      <div className={cn('font-mono text-lg leading-tight tabular-nums', tone)}>{value}</div>
    </div>
  );
}
