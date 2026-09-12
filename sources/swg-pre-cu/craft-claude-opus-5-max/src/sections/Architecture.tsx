import { Bullet, CodeBlock, Panel, SectionHeader, Tag } from '@/components/ui';
import { TIERS } from '@/engine/craft';
import { cn } from '@/utils/cn';

export function Verdict() {
  return (
    <section id="verdict" className="scroll-mt-20">
      <SectionHeader
        index="01"
        kicker="Language Selection"
        title="Verdict: TypeScript, strict, running in a browser tab under Vite."
        lede={
          <>
            No engine. No scene graph. No actors. The prototype is a document that computes. You get
            hot-reload measured in tens of milliseconds, a UI layer you don't have to build, and a
            type system whose <em>structural</em> semantics happen to be an exact match for SWG's
            "this resource class exposes these attribute lanes" problem.
          </>
        }
      />

      <div className="grid gap-4 lg:grid-cols-[1.35fr_1fr]">
        <Panel title="The argument in one paragraph">
          <p className="text-[15px] leading-relaxed text-zinc-400">
            The Pre-CU core loop is not a game. Strip the fiction and it is{' '}
            <strong className="text-zinc-200">a taxonomy tree</strong> with inherited numeric bounds,{' '}
            <strong className="text-zinc-200">an append-only interval log</strong> of randomized
            spawns, <strong className="text-zinc-200">a sparse weight matrix</strong> mapping
            (slot, attribute) pairs onto output properties, and{' '}
            <strong className="text-zinc-200">a seven-state probability table</strong> applied
            iteratively under a scarce point budget. Total working set: a few hundred kilobytes.
            Total arithmetic: a few thousand multiply-adds per craft. There is no performance
            problem here and there never will be. The only cost that matters in this phase is{' '}
            <strong className="text-amber-300">the cost of changing the schema</strong>, and you are
            going to change it a hundred times before the loop feels right. Pick the language that
            makes schema churn free and puts a sortable table of 900 resource spawns on screen in an
            afternoon.
          </p>
          <div className="mt-4 flex flex-wrap gap-1.5 border-t border-zinc-800 pt-3">
            <Tag tone="emerald">runner-up: Python 3.12 + dataclasses</Tag>
            <Tag tone="rose">deferred: C++</Tag>
            <Tag tone="rose">rejected: Blueprint</Tag>
            <Tag tone="rose">rejected: Unity / Godot</Tag>
            <Tag tone="rose">wrong phase: Rust</Tag>
          </div>
        </Panel>

        <div className="grid grid-cols-2 gap-3">
          <Stat k="edit → observe" v="~30 ms" note="Vite HMR, state preserved" tone="text-emerald-300" />
          <Stat k="add an attribute lane" v="1 line" note="union member + cap entry" tone="text-emerald-300" />
          <Stat k="same in UE5 C++" v="~5 files" note="+ full rebuild + asset re-save" tone="text-rose-400" />
          <Stat k="debug UI cost" v="0" note="the DOM is your inspector" tone="text-emerald-300" />
          <Stat k="live dataset" v="< 1 MB" note="entire galaxy timeline, in L2" tone="text-cyan-300" />
          <Stat k="port surface" v="~600 LOC" note="pure functions; data goes as JSON" tone="text-cyan-300" />
        </div>
      </div>
    </section>
  );
}

function Stat({ k, v, note, tone }: { k: string; v: string; note: string; tone: string }) {
  return (
    <div className="rounded border border-zinc-800 bg-zinc-950/70 p-3">
      <div className="font-mono text-[9.5px] uppercase tracking-[0.14em] text-zinc-600">{k}</div>
      <div className={cn('mt-1 font-mono text-xl leading-none', tone)}>{v}</div>
      <div className="mt-1.5 text-[11px] leading-snug text-zinc-500">{note}</div>
    </div>
  );
}

export function WhyTypeScript() {
  return (
    <section id="why" className="scroll-mt-20">
      <SectionHeader
        index="02"
        kicker="Fit Analysis"
        title="Why this language, for this specific machine"
        lede="Not general advocacy. Seven properties of TypeScript that map one-to-one onto load-bearing requirements of the SWG resource/crafting loop."
      />

      <div className="grid gap-4 lg:grid-cols-2">
        <Panel title="01 — Structural typing is an attribute mask">
          <p className="mb-3 text-[14px] leading-relaxed text-zinc-400">
            Steel has no Potential Energy. Inert Gas has no Malleability. Wild Meat has Flavor and
            Unit Toughness but no Conductivity. Under nominal typing this becomes either a class
            explosion or a fat struct full of sentinel zeros — and sentinels poison every weighted
            average downstream. TypeScript expresses the real shape directly: a partial record over
            a closed lane vocabulary, where <code className="text-cyan-300">undefined</code> means{' '}
            <em>absent</em>, not <em>zero</em>.
          </p>
          <CodeBlock
            caption="engine/resources.ts"
            code={`export const STAT_KEYS = ['OQ','CD','CR','DR','FL','HR','MA','PE','SR','UT'] as const;
export type StatKey   = typeof STAT_KEYS[number];
export type CapTable  = Partial<Record<StatKey, readonly [number, number]>>;
export type StatVector = Partial<Record<StatKey, number>>;

// absence is semantic. a PE-weighted group fed Inert Gas scores 0 —
// a legal, terrible, and *representable* player decision.`}
          />
        </Panel>

        <Panel title="02 — Type erasure: the runtime object is the wire format">
          <p className="mb-3 text-[14px] leading-relaxed text-zinc-400">
            Every structure in this prototype is already a JSON DTO. There is no ORM, no reflection
            layer, no serializer to keep in sync. Dump the galaxy to a file, diff two seeds with{' '}
            <code className="text-cyan-300">jq</code>, paste a broken spawn into a unit test. When
            the port comes, the tuning tables cross the boundary untouched as UE5 DataTable rows;
            only the pure functions get transliterated.
          </p>
          <CodeBlock
            caption="the port boundary is free"
            code={`writeFileSync('galaxy.json', JSON.stringify(generateGalaxy({ seed, horizon })));
// -> UE5: FTableRowBase import, zero glue code.
// The prototype's persistence format IS the shipping data format.`}
          />
        </Panel>

        <Panel title="03 — Discriminated unions make the outcome table total">
          <p className="mb-3 text-[14px] leading-relaxed text-zinc-400">
            Assembly and experimentation both resolve to one of seven tiers, and the one you will
            forget to handle is <code className="text-rose-400">critical_failure</code> — the one
            that destroys the item and the player's afternoon. A literal union plus{' '}
            <code className="text-cyan-300">noFallthroughCasesInSwitch</code> makes forgetting it a
            compile error, not a Tuesday-night bug report.
          </p>
          <CodeBlock
            caption="engine/craft.ts"
            code={`export type Tier = 'critical_failure' | 'failure' | 'moderate'
                 | 'success' | 'good' | 'great' | 'amazing';

export type CraftPhase = 'design' | 'experimentation' | 'complete' | 'destroyed';
// every transition is (state, input) => state. no hidden globals.`}
          />
        </Panel>

        <Panel title="04 — The recipe is a sparse matrix, so write it as one">
          <p className="mb-3 text-[14px] leading-relaxed text-zinc-400">
            This is the heart of the "experimentation" demand in the brief. A schematic property is
            a weighted sum over (slot, attribute) pairs — a two-level sparse matrix, not an object
            graph. In TS the evaluator is a dozen lines of map/reduce over plain arrays, editable
            live. In C++ it is a template argument you will regret; in Blueprint it is a screen of
            nested ForEach nodes you cannot grep.
          </p>
          <CodeBlock
            caption="the entire quality evaluator"
            code={`export function propertyQuality(prop, resolved): number {
  let acc = 0, wsum = 0;
  for (const g of prop.groups) {
    const res = resolved[g.slotId];
    if (!res) continue;
    let sAcc = 0, sW = 0;
    for (const sw of g.stats) {
      sAcc += (res.stats[sw.stat] ?? 0) * sw.weight;  // ?? 0 == "lane absent"
      sW   += sw.weight;
    }
    acc  += (sW ? sAcc / sW : 0) * g.weight;
    wsum += g.weight;
  }
  return wsum ? clamp01(acc / wsum / 1000) : 0;
}`}
          />
        </Panel>

        <Panel title="05 — Iteration latency dominates everything else">
          <ul className="divide-y divide-zinc-900">
            <Bullet label="Vite HMR">
              module swap in ~30 ms with React state preserved. You retune a cap table and watch 900
              spawns re-colour without losing your selected row or your half-finished craft session.
            </Bullet>
            <Bullet label="UE5 C++">
              a header touch is a 30 s–4 min rebuild plus an editor reload. Live Coding helps until
              you change a USTRUCT, which is the change you make most.
            </Bullet>
            <Bullet label="Consequence">
              at 100 schema iterations, TS costs you an hour of waiting. C++ costs you a week. That
              week is the entire budget of a tech demo.
            </Bullet>
          </ul>
        </Panel>

        <Panel title="06 — Determinism you own, not the standard library's">
          <p className="mb-3 text-[14px] leading-relaxed text-zinc-400">
            A randomized global spawn pool is worthless if you cannot reproduce it. Do not use{' '}
            <code className="text-rose-400">Math.random()</code> or{' '}
            <code className="text-rose-400">std::uniform_int_distribution</code> (whose output is
            implementation-defined across libstdc++/libc++/MSVC). Ship a 6-line mulberry32 in
            userland. <code className="text-cyan-300">Math.imul</code> gives you exact int32
            multiply, so the same generator transliterates to C++ character-for-character and
            produces the identical galaxy.
          </p>
          <CodeBlock
            caption="engine/rng.ts — portable, pure, in-band"
            code={`export function step(state: number): { s: number; v: number } {
  const t = (state + 0x6d2b79f5) | 0;
  let r = Math.imul(t ^ (t >>> 15), 1 | t);
  r = (r + Math.imul(r ^ (r >>> 7), 61 | r)) ^ r;
  return { s: t, v: ((r ^ (r >>> 14)) >>> 0) / 4294967296 };
}
// generator state lives INSIDE CraftState -> a session is replayable
// from (initialState, inputs[]). bug reports become regression tests.`}
          />
        </Panel>

        <Panel title="07 — Memory management: there isn't any, and that is correct" className="lg:col-span-2">
          <div className="grid gap-5 lg:grid-cols-2">
            <p className="text-[14px] leading-relaxed text-zinc-400">
              The entire materialized galaxy — every spawn interval across a 240-day horizon — is
              roughly 2,000 records of ten small numbers, a few planet tuples, and a name. Under a
              megabyte. It fits in L2 and is regenerated in single-digit milliseconds. Any argument
              for manual memory control at this phase is an argument for paying a week of build
              latency to save microseconds you cannot measure. GC pressure is a non-issue because
              this is a <em>data loop</em>, not a frame loop: allocations happen on user action, not
              at 120 Hz.
            </p>
            <p className="text-[14px] leading-relaxed text-zinc-400">
              What you <strong className="text-zinc-200">should</strong> carry forward is{' '}
              <strong className="text-amber-300">layout discipline</strong>, and it costs nothing to
              adopt now: intern class identifiers as strings and never hold a pointer to a class
              node; keep the class tree a flat array with parent indices (flyweight); keep attribute
              lanes in a fixed, ordered vocabulary so the eventual UE5 representation can be a{' '}
              <code className="text-cyan-300">TStaticArray&lt;uint16, 10&gt;</code> or an SoA column
              store without touching a single formula. Write JS that is shaped like the C++ you will
              eventually write, and the port becomes transliteration instead of redesign.
            </p>
          </div>
        </Panel>
      </div>
    </section>
  );
}

const DEMARCATION = [
  {
    opt: 'C++ (UE5 or freestanding)',
    verdict: 'Defer',
    tone: 'rose' as const,
    cost:
      'The cost is not writing it — it is mutating it. Adding one attribute lane touches the USTRUCT, the cap tables, the serializer, the editor metadata, and rebuilds every translation unit that includes the header. Existing DataAssets silently invalidate. No REPL; your inspector is ImGui boilerplate or printf.',
    when:
      'Correct the moment the loop is frozen and you need it inside the gameplay framework, replication, or a profiler.',
  },
  {
    opt: 'Blueprint',
    verdict: 'Reject',
    tone: 'rose' as const,
    cost:
      'Structurally disqualified for schema work. Blueprint structs are binary assets: no textual diff, no grep, no code review. Renaming a member silently resets defaults on every asset referencing it. A schematic is an array-of-structs-of-arrays; iterating one becomes a screen of ForEach spaghetti with no way to unit test it. Blueprint is an orchestration language, not a schema language.',
    when:
      'Correct for wiring the finished, C++-owned system to actors and UMG — designers driving exposed parameters, never owning the model.',
  },
  {
    opt: 'Unity / Godot',
    verdict: 'Reject',
    tone: 'rose' as const,
    cost:
      'You would stand up a scene graph, a GameObject/Node hierarchy, and a serialization pipeline in order to test a spreadsheet. Every engine service here is pure tax. Explicitly out of scope per the brief.',
    when: 'Never, for this phase. The ultimate destination is UE5 regardless.',
  },
  {
    opt: 'Python 3.12 + dataclasses',
    verdict: 'Runner-up',
    tone: 'emerald' as const,
    cost:
      'Genuinely close. Better REPL, cleaner numeric story if you want numpy for the weight matrix. Loses on two axes: refactor safety (you will restructure this schema weekly; mypy is opt-in, slow, and will not catch a renamed dict key) and presentation (Textual or Dear PyGui costs you a day; a <table> costs ten minutes).',
    when:
      'Take it without regret if you would rather live in a REPL than a browser. The blueprint below is language-agnostic.',
  },
  {
    opt: 'Rust',
    verdict: 'Wrong phase',
    tone: 'amber' as const,
    cost:
      'The best model of the problem on this list — sum types, exhaustive matching, serde. But the borrow checker taxes exactly the exploratory, throw-it-away refactors this phase consists of, and compile times put you back in C++ territory.',
    when: 'Excellent choice if you were shipping a standalone deterministic simulation server.',
  },
];

export function Demarcation() {
  return (
    <section id="demarcation" className="scroll-mt-20">
      <SectionHeader
        index="03"
        kicker="Demarcation"
        title="Why not C++ or Blueprint — despite UE5 being the destination"
        lede="Choosing the destination language for the prototyping phase is the single most common way a systems prototype dies. The prototype's job is to answer design questions per hour. The engine's job is to ship. Different tools."
      />
      <div className="overflow-hidden rounded-lg border border-zinc-800">
        <table className="w-full border-collapse text-left">
          <thead className="bg-zinc-900/70">
            <tr className="font-mono text-[10px] uppercase tracking-widest text-zinc-500">
              <th className="px-4 py-2.5 font-medium">option</th>
              <th className="px-4 py-2.5 font-medium">call</th>
              <th className="px-4 py-2.5 font-medium">what it costs you in this phase</th>
              <th className="px-4 py-2.5 font-medium">when it becomes correct</th>
            </tr>
          </thead>
          <tbody>
            {DEMARCATION.map((d) => (
              <tr key={d.opt} className="border-t border-zinc-800 align-top">
                <td className="px-4 py-3 font-mono text-[12px] font-medium text-zinc-200">{d.opt}</td>
                <td className="px-4 py-3">
                  <Tag tone={d.tone}>{d.verdict}</Tag>
                </td>
                <td className="max-w-md px-4 py-3 text-[13px] leading-relaxed text-zinc-400">
                  {d.cost}
                </td>
                <td className="max-w-xs px-4 py-3 text-[13px] leading-relaxed text-zinc-500">
                  {d.when}
                </td>
              </tr>
            ))}
          </tbody>
        </table>
      </div>

      <Panel title="The rule" className="mt-4">
        <p className="text-[15px] leading-relaxed text-zinc-400">
          Prototype in the language that minimizes the cost of being wrong. You are going to be
          wrong about resource cap distributions, about how many experimentation points feel right,
          about whether Malleability should invert into encumbrance, and about whether a 950 OQ
          spawn should occur weekly or monthly. Every one of those is a five-second edit here and a
          five-minute rebuild there. The tech demo's deliverable is not code — it is{' '}
          <strong className="text-zinc-200">a validated set of tuning curves and a proven data
          model</strong>. Both cross the language boundary as JSON.
        </p>
      </Panel>
    </section>
  );
}

export function Blueprint() {
  return (
    <section id="blueprint" className="scroll-mt-20">
      <SectionHeader
        index="04"
        kicker="Data Architecture"
        title="High-level structural blueprint for the resource / crafting loop"
        lede="Five structures. Everything else in the demo below is derived state or presentation. This is the part that must survive the port to UE5 unchanged."
      />

      <div className="space-y-4">
        <Panel title="S1 — Resource class tree: flat array, parent ids, inherited caps">
          <div className="grid gap-5 lg:grid-cols-[1fr_1.1fr]">
            <div className="space-y-3 text-[14px] leading-relaxed text-zinc-400">
              <p>
                The taxonomy is a tree, but never store it as one. A flat array plus an id index
                gives you O(1) lookup, trivial serialization, and stable identity across a reload.
                Caps are declared as <em>partial overrides</em> and merged root→leaf, memoized on
                first resolve. Declaring a new tier is a one-line insert; nothing recompiles,
                nothing re-instances.
              </p>
              <p>
                Slot matching is a subtree predicate, not an equality test. A schematic slot that
                accepts <code className="text-cyan-300">ferrous</code> accepts Steel and Iron
                automatically — including classes you add tomorrow. This is the mechanism that lets
                the resource catalogue grow without touching a single schematic.
              </p>
            </div>
            <CodeBlock
              caption="engine/resources.ts"
              code={`interface ResourceClass {
  id: string;            // interned; nothing holds a pointer
  name: string;
  parent: string | null;
  caps: CapTable;        // partial override, merged down the chain
  spawnable: boolean;    // interior nodes exist only for slot matching
  scarcity: number;      // 0 ubiquitous .. 1 scarce
}

const capCache = new Map<string, CapTable>();
function resolveCaps(id: string): CapTable {          // memoized walk
  const c = CLASS_INDEX.get(id)!;
  return { ...(c.parent ? resolveCaps(c.parent) : {}), ...c.caps };
}

function isDescendant(cls: string, root: string): boolean { /* walk parents */ }`}
            />
          </div>
        </Panel>

        <Panel title="S2 — Spawn pool: an append-only interval log, not a tick loop">
          <div className="grid gap-5 lg:grid-cols-[1.1fr_1fr]">
            <CodeBlock
              caption="materialize once, query forever"
              code={`interface ResourceSpawn {
  id: string; name: string; classId: string;
  stats: StatVector;                       // rolled inside resolved caps
  planets: { planet: Planet; concentration: number }[];
  spawnTick: number;                       // inclusive, galactic day
  despawnTick: number;                     // exclusive
  purity: number;                          // cached mean cap-relative position
}

const activeAt = (g: Galaxy, t: number) =>
  g.spawns.filter(s => s.spawnTick <= t && t < s.despawnTick);`}
            />
            <ul className="divide-y divide-zinc-900">
              <Bullet label="No simulation step">
                Time is a query parameter. "What is spawned on day 137" is a filter over immutable
                intervals. There is no tick to desync, no state to corrupt, and time travel — the
                most useful debugging affordance you will build — is free.
              </Bullet>
              <Bullet label="Per-class RNG streams">
                Each spawnable class seeds its own generator from{' '}
                <code className="text-cyan-300">hash(seed + '::' + classId)</code>. Retuning Copper
                does not reshuffle the entire galaxy, so A/B comparisons stay meaningful.
              </Bullet>
              <Bullet label="Richness bias">
                Stats roll on <code className="text-cyan-300">min + (max-min)·rand^k</code>. One
                per-spawn roll picks k (1.7 typical, 0.45 for the ~3.5% jackpot), then per-lane
                jitter prevents uniformly-perfect spawns. That single exponent is the knob that
                makes a 964 OQ find feel like an event — and it is a live edit here.
              </Bullet>
            </ul>
          </div>
        </Panel>

        <Panel title="S3 — Schematic: a two-level sparse weight matrix">
          <div className="grid gap-5 lg:grid-cols-[1fr_1.1fr]">
            <div className="space-y-3 text-[14px] leading-relaxed text-zinc-400">
              <p>
                A schematic declares <strong className="text-zinc-200">slots</strong> (class-
                constrained, with unit counts) and{' '}
                <strong className="text-zinc-200">experimental properties</strong>. A property owns
                one or more <strong className="text-zinc-200">groups</strong>; a group binds a slot
                to a set of weighted attribute lanes. That nesting is the whole recipe language.
              </p>
              <p>
                <strong className="text-zinc-200">Experiment lines</strong> sit on top and are
                deliberately many-to-many with properties: one line drives both Capacity and
                Recharge, another drives Encumbrance alone. That coupling — plus a point budget
                smaller than the number of lines — is what turns experimentation into a decision
                rather than a formality.
              </p>
            </div>
            <CodeBlock
              caption="engine/schematics.ts"
              code={`interface ExperimentalProperty {
  id: string; label: string; unit: string;
  min: number; max: number;        // output range at 0% / 100%
  invert?: boolean;                // mass, encumbrance: lower is better
  groups: {
    slotId: string;                // which slot feeds this group
    weight: number;                // group weight
    stats: { stat: StatKey; weight: number }[];
  }[];
}

interface ExperimentLine { id: string; label: string; properties: string[]; }`}
            />
          </div>
        </Panel>

        <Panel title="S4 — The equation">
          <div className="grid gap-5 lg:grid-cols-[1.15fr_1fr]">
            <div>
              <CodeBlock
                caption="four lines that define the entire economy"
                code={`quality  = Σ_g w_g · ( Σ_s w_s · stat_s / Σ_s w_s ) / Σ_g w_g / 1000   // [0,1]
ceiling  = 0.5 + 0.5 · quality                    // resources gate the top end
base     = min(quality · assemblyMult, ceiling)   // assembly tier sets the start
pct_n+1  = pct_n + (ceiling - pct_n)·(1 - (1 - gain)^points)   // experimentation

value    = invert ? max - (max-min)·pct : min + (max-min)·pct`}
              />
              <p className="mt-3 text-[14px] leading-relaxed text-zinc-400">
                The load-bearing line is <code className="text-amber-300">ceiling</code>. Resources
                do not merely bias the result — they impose a hard wall that no amount of
                experimentation can cross. That single constraint is what creates resource scarcity,
                surveyor labour value, stockpiling, and a player economy. Delete it and crafting
                collapses into a slot machine.
              </p>
            </div>
            <div>
              <div className="mb-2 font-mono text-[10px] uppercase tracking-widest text-zinc-500">
                outcome table — shared by assembly and experimentation
              </div>
              <div className="overflow-hidden rounded border border-zinc-800">
                <table className="w-full border-collapse font-mono text-[11px]">
                  <thead className="bg-zinc-900/70 text-left text-zinc-500">
                    <tr>
                      <th className="px-3 py-1.5 font-medium">tier</th>
                      <th className="px-2 py-1.5 text-right font-medium">margin ≥</th>
                      <th className="px-2 py-1.5 text-right font-medium">asm ×</th>
                      <th className="px-3 py-1.5 text-right font-medium">exp gain/pt</th>
                    </tr>
                  </thead>
                  <tbody>
                    {TIERS.map((t) => (
                      <tr key={t.id} className="border-t border-zinc-900">
                        <td
                          className={cn(
                            'px-3 py-1.5',
                            t.tone === 'peak'
                              ? 'text-fuchsia-300'
                              : t.tone === 'good'
                                ? 'text-emerald-300'
                                : t.tone === 'ok'
                                  ? 'text-cyan-300'
                                  : t.tone === 'weak'
                                    ? 'text-amber-300'
                                    : 'text-rose-400',
                          )}
                        >
                          {t.label}
                        </td>
                        <td className="px-2 py-1.5 text-right tabular-nums text-zinc-400">
                          {t.margin === -Infinity ? '−∞' : t.margin}
                        </td>
                        <td className="px-2 py-1.5 text-right tabular-nums text-zinc-400">
                          {t.assemblyMult.toFixed(2)}
                        </td>
                        <td className="px-3 py-1.5 text-right tabular-nums text-zinc-400">
                          {t.gain < 0 ? (t.gain === -1 ? 'destroy' : t.gain.toFixed(2)) : `+${t.gain.toFixed(2)}`}
                        </td>
                      </tr>
                    ))}
                  </tbody>
                </table>
              </div>
              <p className="mt-3 text-[13px] leading-relaxed text-zinc-500">
                <code className="text-cyan-300">margin = skillMod − complexity·k − 7·(points−1) + U(−50,50)</code>.
                Gain is concave in points (compounding against shrinking headroom) while risk is
                linear in points. The optimum is interior — which is precisely the risk/reward
                decision the brief asks for, and it falls out of the math rather than being
                hand-authored.
              </p>
            </div>
          </div>
        </Panel>

        <Panel title="S5 — Craft session: a serializable FSM with in-band PRNG state">
          <div className="grid gap-5 lg:grid-cols-[1fr_1.1fr]">
            <ul className="divide-y divide-zinc-900">
              <Bullet label="design → experimentation → complete">
                with <code className="text-rose-400">destroyed</code> reachable from both. Four
                states, five transitions, all pure functions of{' '}
                <code className="text-cyan-300">(state, input)</code>.
              </Bullet>
              <Bullet label="rngState lives in the struct">
                so a session is a value, not a process. Serialize it, email it, replay it. A bug
                report is <code className="text-cyan-300">{'{ initialState, inputs[] }'}</code> and
                becomes a regression test verbatim.
              </Bullet>
              <Bullet label="base / ceiling frozen at assembly">
                Resource contribution is snapshotted once. Experimentation only mutates{' '}
                <code className="text-cyan-300">pct</code>. That separation is what makes the UI
                able to draw "where you started, where you are, where the rock will let you go" as
                three marks on one bar.
              </Bullet>
            </ul>
            <CodeBlock
              caption="engine/craft.ts"
              code={`interface CraftState {
  schematicId: string;
  phase: CraftPhase;
  rngState: number;                        // in-band. replayable.
  assignment: Record<string, string | null>;
  assemblyTier: Tier | null;
  base:    Record<string, number>;         // frozen resource contribution
  ceiling: Record<string, number>;         // hard wall from resources
  pct:     Record<string, number>;         // achieved, mutated by exp
  pointsTotal: number; pointsSpent: number;
  log: LogLine[];
}

assemble(s, schematic, resolved, skill): CraftState
experiment(s, schematic, lineId, points, skill): CraftState
finalize(s, schematic): CraftState`}
            />
          </div>
        </Panel>
      </div>
    </section>
  );
}

export function PortPath() {
  return (
    <section id="port" className="scroll-mt-20">
      <SectionHeader
        index="06"
        kicker="Exit Strategy"
        title="How this becomes UE5 without a rewrite"
        lede="The prototype is disposable; the model is not. Three rules keep the boundary clean from day one."
      />
      <div className="grid gap-4 lg:grid-cols-3">
        <Panel title="Rule 1 — /engine imports nothing">
          <p className="text-[14px] leading-relaxed text-zinc-400">
            No React, no DOM, no fetch, no <code className="text-rose-400">Date.now()</code>, no{' '}
            <code className="text-rose-400">Math.random()</code>. Enforce it with a single lint rule
            or a CI grep. Everything under <code className="text-cyan-300">/engine</code> is pure
            functions over plain data, which means it is simultaneously the design document, the
            test target, and the transliteration source. Roughly 600 lines cross the boundary.
          </p>
        </Panel>
        <Panel title="Rule 2 — Tuning data never gets rewritten">
          <p className="text-[14px] leading-relaxed text-zinc-400">
            Cap tables, class taxonomy, schematics and weight matrices are already declarative
            literals. Emit them as JSON and import as{' '}
            <code className="text-cyan-300">FTableRowBase</code> rows or{' '}
            <code className="text-cyan-300">UDataAsset</code>s. Designers keep editing the same
            tables. The only C++ you write is the evaluator, and it is arithmetic.
          </p>
        </Panel>
        <Panel title="Rule 3 — Golden-master the port">
          <p className="text-[14px] leading-relaxed text-zinc-400">
            Freeze 32 seeds, dump the full spawn timeline plus 1,000 scripted craft sessions to CSV.
            The C++ implementation must reproduce that file byte-for-byte.{' '}
            <code className="text-cyan-300">Math.imul</code> is int32 multiply and both sides are
            IEEE-754 doubles, so this is achievable — pin one rounding helper on both sides and the
            port stops being a leap of faith.
          </p>
        </Panel>
      </div>
    </section>
  );
}
