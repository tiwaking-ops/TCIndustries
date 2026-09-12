import { useMemo, useState } from 'react';
import { activeAt, generateGalaxy } from '@/engine/resources';
import { SectionHeader, Tag } from '@/components/ui';
import { Blueprint, Demarcation, PortPath, Verdict, WhyTypeScript } from '@/sections/Architecture';
import { SpawnBrowser } from '@/sections/SpawnBrowser';
import { CraftBench } from '@/sections/CraftBench';
import { SelfTest } from '@/sections/SelfTest';

const HORIZON = 240;

const NAV = [
  { id: 'verdict', label: '01 Verdict' },
  { id: 'why', label: '02 Fit' },
  { id: 'demarcation', label: '03 Demarcation' },
  { id: 'blueprint', label: '04 Blueprint' },
  { id: 'demo', label: '05 Live Loop' },
  { id: 'port', label: '06 Port' },
];

export default function App() {
  const [seed, setSeed] = useState('ANCHORHEAD');
  const [day, setDay] = useState(120);

  const galaxy = useMemo(() => generateGalaxy({ seed, horizon: HORIZON }), [seed]);
  const active = useMemo(() => activeAt(galaxy, day), [galaxy, day]);

  return (
    <div className="min-h-screen bg-zinc-950 text-zinc-300">
      {/* header */}
      <header className="grid-backdrop relative overflow-hidden border-b border-zinc-800">
        <div className="pointer-events-none absolute -top-40 left-1/3 h-96 w-96 rounded-full bg-amber-600/10 blur-[120px]" />
        <div className="pointer-events-none absolute -bottom-40 right-1/4 h-96 w-96 rounded-full bg-cyan-600/10 blur-[120px]" />
        <div className="relative mx-auto max-w-[1400px] px-6 py-14 lg:py-20">
          <div className="flex flex-wrap items-center gap-2">
            <Tag tone="amber">architecture memo</Tag>
            <Tag tone="zinc">pre-cu swg core loop</Tag>
            <Tag tone="cyan">single-player · offline · pre-engine</Tag>
          </div>
          <h1 className="mt-5 max-w-4xl text-4xl font-bold leading-[1.08] tracking-tight text-zinc-50 sm:text-5xl lg:text-6xl">
            Build the economy in{' '}
            <span className="bg-gradient-to-r from-amber-300 to-amber-500 bg-clip-text text-transparent">
              TypeScript
            </span>
            .<br className="hidden sm:block" /> Port the arithmetic to UE5 later.
          </h1>
          <p className="mt-5 max-w-3xl text-[16px] leading-relaxed text-zinc-400">
            Language selection, justification, and a data-structural blueprint for randomized global
            resource spawns, attribute-driven schematics, and player experimentation — followed by a
            working implementation of the loop, running in this page. Nothing below is a mockup: the
            spawn timeline, the weight matrices, the assembly rolls and the risk/reward
            experimentation curve are all live engine code.
          </p>
          <div className="mt-8 flex flex-wrap gap-2">
            {NAV.map((n) => (
              <a
                key={n.id}
                href={`#${n.id}`}
                className="rounded border border-zinc-800 bg-zinc-900/60 px-3 py-1.5 font-mono text-[11px] uppercase tracking-wider text-zinc-400 transition hover:border-amber-500/60 hover:text-amber-300"
              >
                {n.label}
              </a>
            ))}
          </div>
        </div>
      </header>

      <main className="mx-auto max-w-[1400px] space-y-20 px-6 py-16">
        <Verdict />
        <WhyTypeScript />
        <Demarcation />
        <Blueprint />

        <section id="demo" className="scroll-mt-20">
          <SectionHeader
            index="05"
            kicker="Tech Demo"
            title="The loop, running"
            lede={
              <>
                Everything here is driven by <code className="text-cyan-300">/src/engine</code> —
                zero DOM dependencies, zero engine dependencies, ~600 lines. Reroll the galaxy to
                regenerate {HORIZON} days of spawn intervals from a 32-bit seed; scrub the clock to
                watch the pool rotate; take the best available rock to the bench and find out
                whether your resources or your dice are the constraint.
              </>
            }
          />
          <div className="space-y-6">
            <SpawnBrowser
              galaxy={galaxy}
              day={day}
              setDay={setDay}
              seed={seed}
              setSeed={setSeed}
              active={active}
            />
            <CraftBench active={active} />
            <SelfTest galaxy={galaxy} />
          </div>
        </section>

        <PortPath />
      </main>

      <footer className="border-t border-zinc-800 bg-black/40">
        <div className="mx-auto flex max-w-[1400px] flex-wrap items-center justify-between gap-4 px-6 py-8 font-mono text-[11px] text-zinc-600">
          <span>
            engine: pure TS · prng: mulberry32 · state: value semantics · deps: none beyond React
          </span>
          <span className="text-zinc-700">
            phase 1 of N — prove the loops, freeze the tables, then transliterate.
          </span>
        </div>
      </footer>
    </div>
  );
}
