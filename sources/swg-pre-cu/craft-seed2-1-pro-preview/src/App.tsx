// ============================================================================
// SWG Pre-CU Crafting Core Tech Demo
// UI is deliberately thin: a debug shell over the pure TS engine.
// When porting to UE5, the entire src/engine/ directory maps 1:1 to C++ USTRUCTs
// and UBlueprintFunctionLibrary statics. The React layer is replaced by UMG.
// ============================================================================

import { useMemo, useState, useCallback } from "react";
import { cn } from "./utils/cn";
import {
  generateSpawnTable,
  initSession,
  assignToSlot,
  runAssemblyPhase,
  experiment,
  autoExperiment,
  finalize,
  validateSlot,
  computeResourceCaps,
  makeRng,
  getClass,
  isSubtypeOf,
  SCHEMATIC_LIST,
  MASTER_SKILL,
  NOVICE_SKILL,
  ALL_STATS,
  type Resource,
  type Schematic,
  type CraftingSession,
  type CraftedItem,
  type SlotAssignment,
  type SkillProfile,
  type ExperimentOutcome,
} from "./engine";
import { StatPill } from "./components/StatBar";

// ---------------------------------------------------------------------------
// Utilities
// ---------------------------------------------------------------------------

function formatPlanet(p: string) {
  const names: Record<string, string> = {
    tatooine: "Tatooine", naboo: "Naboo", corellia: "Corellia",
    dantooine: "Dantooine", endor: "Endor", dathomir: "Dathomir",
    yavin4: "Yavin IV", rori: "Rori", talus: "Talus", lok: "Lok",
  };
  return names[p] ?? p;
}

function outcomeColor(o: ExperimentOutcome) {
  switch (o) {
    case "amazing":  return "text-cyan-300 bg-cyan-950 border-cyan-700";
    case "great":    return "text-emerald-300 bg-emerald-950 border-emerald-700";
    case "success":  return "text-zinc-300 bg-zinc-800 border-zinc-700";
    case "failure":  return "text-orange-300 bg-orange-950 border-orange-800";
    case "critical": return "text-rose-300 bg-rose-950 border-rose-700";
  }
}

function qualityLabel(q: number) {
  if (q >= 0.9) return "LEGENDARY";
  if (q >= 0.75) return "Exceptional";
  if (q >= 0.55) return "High Quality";
  if (q >= 0.35) return "Average";
  if (q >= 0.2) return "Below Average";
  return "Trash";
}

function qualityColor(q: number) {
  if (q >= 0.9) return "text-cyan-300";
  if (q >= 0.75) return "text-emerald-300";
  if (q >= 0.55) return "text-lime-300";
  if (q >= 0.35) return "text-yellow-300";
  if (q >= 0.2) return "text-orange-300";
  return "text-rose-400";
}

// ---------------------------------------------------------------------------
// Main App
// ---------------------------------------------------------------------------

export default function App() {
  const [seed, setSeed] = useState<number>(1337);
  const [surveyFilter, setSurveyFilter] = useState({ planet: "all", classFilter: "", search: "" });
  const [selectedResourceId, setSelectedResourceId] = useState<string | null>(null);
  const [inventory, setInventory] = useState<ReadonlyMap<string, number>>(new Map());
  const [selectedSchematic, setSelectedSchematic] = useState<Schematic | null>(null);
  const [session, setSession] = useState<CraftingSession | null>(null);
  const [sessionSeed, setSessionSeed] = useState<number>(42);
  const [crafted, setCrafted] = useState<CraftedItem[]>([]);
  const [useNovice, setUseNovice] = useState(false);
  const [finalItem, setFinalItem] = useState<CraftedItem | null>(null);
  const [spawnTable, setSpawnTable] = useState<Resource[]>(() => generateSpawnTable(1337));
  const [resourcesById, setResourcesById] = useState<ReadonlyMap<string, Resource>>(() => {
    const t = generateSpawnTable(1337);
    return new Map(t.map(r => [r.id, r]));
  });

  const regenerate = useCallback((newSeed: number) => {
    const t = generateSpawnTable(newSeed);
    setSeed(newSeed);
    setSpawnTable(t);
    setResourcesById(new Map(t.map(r => [r.id, r])));
    setInventory(new Map());
    setSession(null);
    setFinalItem(null);
    setSelectedSchematic(null);
  }, []);

  const surveyResource = useCallback((res: Resource, units: number) => {
    setInventory(prev => {
      const next = new Map(prev);
      const held = next.get(res.id) ?? 0;
      const toTake = Math.min(units, res.availableUnits);
      next.set(res.id, held + toTake);
      return next;
    });
  }, []);

  const ejectResource = useCallback((resId: string, units?: number) => {
    setInventory(prev => {
      const next = new Map(prev);
      const held = next.get(resId) ?? 0;
      if (units === undefined || units >= held) {
        next.delete(resId);
      } else {
        next.set(resId, held - units);
      }
      return next;
    });
  }, []);

  const startCraft = useCallback((sch: Schematic) => {
    const skill: SkillProfile = useNovice ? NOVICE_SKILL : MASTER_SKILL;
    const s = initSession(sch, skill);
    // Assembly phase automatically runs (master craftsman bonus in our model).
    runAssemblyPhase(s);
    setSelectedSchematic(sch);
    setSession(s);
    setFinalItem(null);
    setSessionSeed(Math.floor(Math.random() * 1e9));
  }, [useNovice]);

  const handleAssign = useCallback((slotId: string, resourceId: string) => {
    if (!session) return;
    const s = { ...session, assignments: session.assignments.map(a => ({ ...a })), progress: session.progress.map(p => ({ ...p })), log: [...session.log] };
    const slot = session.schematic.slots.find(x => x.id === slotId);
    const need = slot?.requiredUnits ?? 1;
    const have = inventory.get(resourceId) ?? 0;
    const useUnits = Math.min(need, have);
    assignToSlot(s, slotId, resourceId, useUnits);
    setSession(s);
  }, [session, inventory]);

  const handleUnassign = useCallback((slotId: string) => {
    if (!session) return;
    const s = { ...session, assignments: session.assignments.map(a => ({ ...a })), progress: session.progress.map(p => ({ ...p })), log: [...session.log] };
    const a = s.assignments.find(x => x.slotId === slotId);
    if (a) { a.resourceId = null; a.unitsUsed = 0; }
    setSession(s);
  }, [session]);

  const handleExperiment = useCallback((propertyId: string, points: number) => {
    if (!session) return;
    const s = { ...session, assignments: [...session.assignments], progress: session.progress.map(p => ({ ...p })), log: [...session.log] };
    const rng = makeRng(sessionSeed + s.log.length * 7919 + propertyId.length * 31);
    experiment(s, propertyId, points, rng);
    setSession(s);
  }, [session, sessionSeed]);

  const handleAuto = useCallback(() => {
    if (!session) return;
    const s = { ...session, assignments: [...session.assignments], progress: session.progress.map(p => ({ ...p })), log: [...session.log] };
    const rng = makeRng(sessionSeed);
    autoExperiment(s, rng);
    setSession(s);
  }, [session, sessionSeed]);

  const handleFinalize = useCallback(() => {
    if (!session) return;
    const rng = makeRng(sessionSeed + 999991);
    const item = finalize(session, resourcesById, rng);
    setFinalItem(item);
    // Deduct resources from inventory.
    setInventory(prev => {
      const next = new Map(prev);
      for (const used of item.resourcesUsed) {
        const cur = next.get(used.resourceId) ?? 0;
        const remaining = cur - used.units;
        if (remaining <= 0) next.delete(used.resourceId);
        else next.set(used.resourceId, remaining);
      }
      return next;
    });
    setCrafted(prev => [item, ...prev]);
  }, [session, resourcesById, sessionSeed]);

  const handleResetSession = useCallback(() => {
    if (!selectedSchematic) return;
    startCraft(selectedSchematic);
  }, [selectedSchematic, startCraft]);

  // ----- Derived data -----
  const filteredSpawns = useMemo(() => {
    return spawnTable.filter(r => {
      if (surveyFilter.planet !== "all" && r.planetId !== surveyFilter.planet) return false;
      if (surveyFilter.search && !r.name.toLowerCase().includes(surveyFilter.search.toLowerCase())) return false;
      if (surveyFilter.classFilter) {
        const cls = getClass(r.classId);
        if (!isSubtypeOf(r.classId, surveyFilter.classFilter)) {
          if (!cls.name.toLowerCase().includes(surveyFilter.classFilter.toLowerCase())) return false;
        }
      }
      return true;
    });
  }, [spawnTable, surveyFilter]);

  const inventoryResources = useMemo(() => {
    const out: Array<{ res: Resource; units: number }> = [];
    for (const [id, units] of inventory) {
      const res = resourcesById.get(id);
      if (res) out.push({ res, units });
    }
    return out;
  }, [inventory, resourcesById]);

  const sessionErrors = useMemo(() => {
    if (!session) return [] as string[];
    const errs: string[] = [];
    for (const slot of session.schematic.slots) {
      const a = session.assignments.find(x => x.slotId === slot.id);
      const res = a?.resourceId ? resourcesById.get(a.resourceId) ?? null : null;
      const err = validateSlot(slot, a as SlotAssignment, res);
      if (err) errs.push(err.message);
    }
    return errs;
  }, [session, resourcesById]);

  const resourceCaps = useMemo(() => {
    if (!session) return new Map<string, number>();
    return computeResourceCaps(session.schematic, session.assignments, resourcesById);
  }, [session, resourcesById]);

  const selectedResource = selectedResourceId ? resourcesById.get(selectedResourceId) : null;

  return (
    <div className="min-h-screen bg-zinc-950 text-zinc-200 font-mono text-[13px]">
      <TopBar seed={seed} onRegen={regenerate} />

      <div className="grid grid-cols-12 gap-2 p-2 max-w-[1800px] mx-auto">
        {/* LEFT: Spawn Survey */}
        <section className="col-span-4 border border-zinc-800 bg-zinc-900/40 rounded">
          <div className="px-2 py-1 border-b border-zinc-800 bg-zinc-900 flex items-center justify-between">
            <h2 className="text-[11px] uppercase tracking-widest text-zinc-400">Global Resource Spawns</h2>
            <span className="text-[10px] text-zinc-500">{filteredSpawns.length} / {spawnTable.length}</span>
          </div>
          <div className="p-2 border-b border-zinc-800 grid grid-cols-2 gap-1">
            <select
              className="col-span-1 bg-zinc-800 border border-zinc-700 px-1 py-0.5 text-[11px]"
              value={surveyFilter.planet}
              onChange={e => setSurveyFilter(f => ({ ...f, planet: e.target.value }))}
            >
              <option value="all">All Planets</option>
              {["tatooine","naboo","corellia","dantooine","endor","dathomir","yavin4","rori","talus","lok"].map(p =>
                <option key={p} value={p}>{formatPlanet(p)}</option>
              )}
            </select>
            <input
              className="col-span-1 bg-zinc-800 border border-zinc-700 px-1 py-0.5 text-[11px] placeholder:text-zinc-600"
              placeholder="Filter by class..."
              value={surveyFilter.classFilter}
              onChange={e => setSurveyFilter(f => ({ ...f, classFilter: e.target.value }))}
            />
            <input
              className="col-span-2 bg-zinc-800 border border-zinc-700 px-1 py-0.5 text-[11px] placeholder:text-zinc-600"
              placeholder="Search by name..."
              value={surveyFilter.search}
              onChange={e => setSurveyFilter(f => ({ ...f, search: e.target.value }))}
            />
          </div>
          <div className="max-h-[calc(100vh-230px)] overflow-auto">
            <table className="w-full text-[11px]">
              <thead className="sticky top-0 bg-zinc-900 text-zinc-500 text-[10px] uppercase">
                <tr>
                  <th className="text-left px-1 py-0.5 w-24">Planet</th>
                  <th className="text-left px-1 py-0.5">Name / Class</th>
                  <th className="text-left px-1 py-0.5">OQ</th>
                  <th className="text-left px-1 py-0.5 w-8"></th>
                </tr>
              </thead>
              <tbody>
                {filteredSpawns.map(r => (
                  <tr
                    key={r.id}
                    onClick={() => setSelectedResourceId(r.id)}
                    className={cn(
                      "border-t border-zinc-800/60 hover:bg-zinc-800/50 cursor-pointer",
                      selectedResourceId === r.id && "bg-amber-950/40"
                    )}
                  >
                    <td className="px-1 py-0.5 text-zinc-500">{formatPlanet(r.planetId).slice(0, 6)}</td>
                    <td className="px-1 py-0.5">
                      <div className="text-zinc-100">{r.name}</div>
                      <div className="text-[10px] text-zinc-500">{getClass(r.classId).name}</div>
                    </td>
                    <td className="px-1 py-0.5">
                      <span className={cn(
                        "tabular-nums",
                        r.stats.OQ >= 900 ? "text-cyan-300" :
                        r.stats.OQ >= 800 ? "text-emerald-300" :
                        r.stats.OQ >= 600 ? "text-yellow-300" : "text-zinc-500"
                      )}>{r.stats.OQ}</span>
                    </td>
                    <td className="px-1 py-0.5">
                      <button
                        onClick={(e) => { e.stopPropagation(); surveyResource(r, 500); }}
                        className="text-[10px] px-1 border border-zinc-700 hover:border-emerald-500 hover:text-emerald-400"
                        title="Survey 500 units"
                      >+500</button>
                    </td>
                  </tr>
                ))}
              </tbody>
            </table>
          </div>
          {selectedResource && <ResourceInspector res={selectedResource} />}
        </section>

        {/* MIDDLE: Inventory + Schematic picker */}
        <section className="col-span-3 flex flex-col gap-2">
          <div className="border border-zinc-800 bg-zinc-900/40 rounded">
            <div className="px-2 py-1 border-b border-zinc-800 bg-zinc-900 flex items-center justify-between">
              <h2 className="text-[11px] uppercase tracking-widest text-zinc-400">Inventory</h2>
              <span className="text-[10px] text-zinc-500">{inventoryResources.length} stacks</span>
            </div>
            <div className="max-h-64 overflow-auto">
              {inventoryResources.length === 0 && (
                <div className="p-3 text-[11px] text-zinc-500 italic">
                  Survey resources from the spawn list to begin.
                </div>
              )}
              {inventoryResources.map(({ res, units }) => (
                <div key={res.id} className="border-t border-zinc-800/60 px-2 py-1 text-[11px]">
                  <div className="flex justify-between items-start">
                    <div className="flex-1 min-w-0">
                      <div className="text-zinc-100 truncate">{res.name}</div>
                      <div className="text-[10px] text-zinc-500">
                        {getClass(res.classId).name} · <span className="text-amber-300">OQ {res.stats.OQ}</span>
                      </div>
                    </div>
                    <div className="text-right pl-2">
                      <div className="text-zinc-300 tabular-nums">{units}u</div>
                      <button
                        onClick={() => ejectResource(res.id)}
                        className="text-[9px] text-rose-400 hover:text-rose-300"
                      >drop</button>
                    </div>
                  </div>
                </div>
              ))}
            </div>
          </div>

          <div className="border border-zinc-800 bg-zinc-900/40 rounded flex-1">
            <div className="px-2 py-1 border-b border-zinc-800 bg-zinc-900 flex items-center justify-between">
              <h2 className="text-[11px] uppercase tracking-widest text-zinc-400">Schematics</h2>
              <label className="flex items-center gap-1 text-[10px] text-zinc-400">
                <input type="checkbox" checked={useNovice} onChange={e => setUseNovice(e.target.checked)} />
                Novice
              </label>
            </div>
            <div className="p-1 space-y-1">
              {SCHEMATIC_LIST.map(sch => (
                <button
                  key={sch.id}
                  onClick={() => startCraft(sch)}
                  className={cn(
                    "w-full text-left p-2 border rounded",
                    selectedSchematic?.id === sch.id
                      ? "border-amber-600 bg-amber-950/30"
                      : "border-zinc-800 bg-zinc-900 hover:border-zinc-600"
                  )}
                >
                  <div className="flex justify-between items-baseline">
                    <div className="text-zinc-100 text-[12px]">{sch.name}</div>
                    <div className="text-[9px] text-zinc-500">{sch.profession}</div>
                  </div>
                  <div className="text-[10px] text-zinc-500 mt-0.5">{sch.description}</div>
                  <div className="text-[9px] text-zinc-600 mt-0.5">
                    {sch.slots.length} slots · {sch.experimentPoints} exp pts · complexity {sch.complexity}
                  </div>
                </button>
              ))}
            </div>
          </div>
        </section>

        {/* RIGHT: Workstation */}
        <section className="col-span-5 border border-zinc-800 bg-zinc-900/40 rounded flex flex-col">
          <div className="px-2 py-1 border-b border-zinc-800 bg-zinc-900 flex items-center justify-between">
            <h2 className="text-[11px] uppercase tracking-widest text-zinc-400">Crafting Workstation</h2>
            {session && (
              <div className="flex gap-2 items-center">
                <span className="text-[10px] text-zinc-500">
                  Points: <span className="text-amber-300">{session.pointsRemaining}</span> / {session.schematic.experimentPoints}
                </span>
                <button
                  onClick={handleResetSession}
                  className="text-[10px] px-1.5 py-0.5 border border-zinc-700 hover:border-rose-500 hover:text-rose-400"
                >Reset</button>
              </div>
            )}
          </div>

          {!session && (
            <div className="p-6 text-center text-zinc-500 text-xs italic flex-1 flex items-center justify-center">
              Select a schematic to begin crafting.
            </div>
          )}

          {session && !finalItem && (
            <div className="flex-1 flex flex-col overflow-hidden">
              {/* Slots */}
              <div className="p-2 border-b border-zinc-800">
                <div className="text-[10px] uppercase text-zinc-500 mb-1">Assembly</div>
                {sessionErrors.length > 0 && (
                  <div className="mb-2 p-1.5 border border-rose-800 bg-rose-950/40 text-[10px] text-rose-300">
                    {sessionErrors.map((e, i) => <div key={i}>• {e}</div>)}
                  </div>
                )}
                <div className="space-y-1.5">
                  {session.schematic.slots.map(slot => {
                    const a = session.assignments.find(x => x.slotId === slot.id)!;
                    const assignedRes = a.resourceId ? resourcesById.get(a.resourceId) : null;
                    const compatible = inventoryResources
                      .filter(({ res }) => isSubtypeOf(res.classId, slot.requiredClassId) && (inventory.get(res.id) ?? 0) >= slot.requiredUnits);
                    return (
                      <div key={slot.id} className="border border-zinc-800 bg-zinc-900/60 p-1.5">
                        <div className="flex justify-between items-center">
                          <div>
                            <div className="text-[11px] text-zinc-100">{slot.label}</div>
                            <div className="text-[9px] text-zinc-500">
                              Requires {slot.requiredUnits}u {getClass(slot.requiredClassId).name}
                            </div>
                          </div>
                          {assignedRes ? (
                            <button
                              onClick={() => handleUnassign(slot.id)}
                              className="text-[9px] text-rose-400 hover:text-rose-300"
                            >clear</button>
                          ) : null}
                        </div>
                        {assignedRes ? (
                          <div className="mt-1 flex items-center justify-between">
                            <div>
                              <div className="text-[11px] text-emerald-300">{assignedRes.name}</div>
                              <div className="flex gap-2 mt-0.5 flex-wrap">
                                {ALL_STATS.map(s => assignedRes.stats[s] > 0 && (
                                  <span key={s} className="text-[9px]">
                                    <span className="text-zinc-500">{s}</span>
                                    <span className="ml-0.5 text-zinc-300 tabular-nums">{assignedRes.stats[s]}</span>
                                  </span>
                                ))}
                              </div>
                            </div>
                            <div className="text-[10px] text-zinc-400">{a.unitsUsed}u</div>
                          </div>
                        ) : (
                          <div className="mt-1">
                            {compatible.length === 0 ? (
                              <div className="text-[10px] text-zinc-600 italic">
                                No matching resources in inventory. Survey some first.
                              </div>
                            ) : (
                              <select
                                className="w-full bg-zinc-800 border border-zinc-700 px-1 py-0.5 text-[10px]"
                                value=""
                                onChange={e => { if (e.target.value) handleAssign(slot.id, e.target.value); }}
                              >
                                <option value="">-- assign resource --</option>
                                {compatible.map(({ res, units }) => (
                                  <option key={res.id} value={res.id}>
                                    {res.name} (OQ {res.stats.OQ}, {units}u available)
                                  </option>
                                ))}
                              </select>
                            )}
                          </div>
                        )}
                      </div>
                    );
                  })}
                </div>
              </div>

              {/* Experimentation */}
              <div className="p-2 border-b border-zinc-800 flex-1 overflow-auto">
                <div className="text-[10px] uppercase text-zinc-500 mb-2 flex items-center justify-between">
                  <span>Experimentation</span>
                  <button
                    onClick={handleAuto}
                    disabled={session.pointsRemaining === 0 || sessionErrors.length > 0}
                    className="px-1.5 py-0.5 border border-zinc-700 hover:border-sky-500 hover:text-sky-400 disabled:opacity-40 disabled:cursor-not-allowed"
                  >Auto-Experiment</button>
                </div>
                <div className="space-y-2">
                  {session.schematic.experimentalProps.map(prop => {
                    const prog = session.progress.find(p => p.propertyId === prop.id)!;
                    const cap = resourceCaps.get(prop.id) ?? 0;
                    const pct = Math.round((prog.boxes / prog.maxBoxes) * 100);
                    const isInverted = prop.minValue > prop.maxValue;
                    return (
                      <div key={prop.id} className="border border-zinc-800 bg-zinc-900/60 p-1.5">
                        <div className="flex items-center justify-between mb-1">
                          <div className="text-[11px] text-zinc-100">{prop.label}</div>
                          <div className="flex items-center gap-2 text-[10px]">
                            <span className="text-zinc-500">
                              cap <span className={cn(cap > 0.75 ? "text-emerald-300" : cap > 0.5 ? "text-yellow-300" : "text-rose-400")}>{Math.round(cap * 100)}%</span>
                            </span>
                            <span className="text-amber-300">{prog.boxes}/{prog.maxBoxes}</span>
                          </div>
                        </div>
                        <div className="h-2 bg-zinc-800 relative overflow-hidden mb-1">
                          <div
                            className={cn("h-full bg-gradient-to-r",
                              isInverted ? "from-rose-500 to-emerald-400" : "from-emerald-500 to-cyan-400")}
                            style={{ width: `${pct}%` }}
                          />
                          <div
                            className="absolute top-0 h-full w-px bg-white/60"
                            style={{ left: `${cap * 100}%` }}
                            title="Resource cap — you cannot exceed this regardless of experiment luck"
                          />
                        </div>
                        <div className="flex gap-1 mt-1">
                          {[1, 2, 3, 4, 5].map(n => (
                            <button
                              key={n}
                              disabled={session.pointsRemaining < n || sessionErrors.length > 0}
                              onClick={() => handleExperiment(prop.id, n)}
                              className={cn(
                                "flex-1 py-0.5 text-[10px] border",
                                n === 5 ? "border-rose-700 text-rose-300 hover:bg-rose-950"
                                       : "border-zinc-700 text-zinc-300 hover:bg-zinc-800",
                                "disabled:opacity-30 disabled:cursor-not-allowed"
                              )}
                            >
                              +{n}
                            </button>
                          ))}
                        </div>
                        <div className="text-[9px] text-zinc-600 mt-1 flex gap-2 flex-wrap">
                          Stats: {Object.entries(prop.statWeights).map(([s, w]) => (
                            <span key={s}>{s} <span className="text-zinc-500">×{w}</span></span>
                          ))}
                        </div>
                      </div>
                    );
                  })}
                </div>
              </div>

              {/* Experiment log */}
              <div className="p-2 max-h-40 overflow-auto">
                <div className="text-[10px] uppercase text-zinc-500 mb-1">Experiment Log</div>
                {session.log.length === 0 ? (
                  <div className="text-[10px] text-zinc-600 italic">No experiments yet.</div>
                ) : (
                  <div className="space-y-0.5">
                    {session.log.slice().reverse().map(e => {
                      const prop = session.schematic.experimentalProps.find(p => p.id === e.propertyId);
                      return (
                        <div key={e.attempt} className="flex items-center gap-2 text-[10px]">
                          <span className="text-zinc-600 w-6">#{e.attempt}</span>
                          <span className={cn(
                            "px-1 border text-[9px]",
                            outcomeColor(e.outcome)
                          )}>{e.outcome}</span>
                          <span className="text-zinc-400 flex-1 truncate">{prop?.label}</span>
                          <span className="text-zinc-500">{e.pointsSpent}pt</span>
                          <span className={cn(
                            "tabular-nums w-10 text-right",
                            e.boxDelta > 0 ? "text-emerald-300" :
                            e.boxDelta < 0 ? "text-rose-400" : "text-zinc-500"
                          )}>{e.boxDelta > 0 ? "+" : ""}{e.boxDelta}</span>
                          <span className="text-zinc-600 w-8 text-right">d{e.dieRoll}</span>
                        </div>
                      );
                    })}
                  </div>
                )}
              </div>

              <div className="p-2 border-t border-zinc-800 bg-zinc-900">
                <button
                  onClick={handleFinalize}
                  disabled={sessionErrors.length > 0 || session.pointsRemaining === session.schematic.experimentPoints}
                  className="w-full py-1.5 border border-emerald-700 bg-emerald-950/50 text-emerald-300 text-[11px] uppercase tracking-wider hover:bg-emerald-900/50 disabled:opacity-40 disabled:cursor-not-allowed"
                >
                  Finalize Craft
                </button>
              </div>
            </div>
          )}

          {finalItem && session && <FinalItemView item={finalItem} schematic={session.schematic} resourcesById={resourcesById} />}
        </section>

        {/* Far right: crafted items history */}
        <section className="col-span-12 lg:col-span-12 mt-2 border border-zinc-800 bg-zinc-900/40 rounded">
          <div className="px-2 py-1 border-b border-zinc-800 bg-zinc-900">
            <h2 className="text-[11px] uppercase tracking-widest text-zinc-400">
              Crafted Items ({crafted.length}) — <span className="text-zinc-600 normal-case">Engine returns immutable CraftedItem records, serializable over network or to DB</span>
            </h2>
          </div>
          <div className="p-2 flex gap-2 overflow-auto">
            {crafted.length === 0 && (
              <div className="text-[11px] text-zinc-500 italic">No items crafted yet.</div>
            )}
            {crafted.map(item => {
              const sch = SCHEMATIC_LIST.find(s => s.id === item.schematicId);
              return (
                <div key={item.id} className="border border-zinc-800 bg-zinc-950 p-2 min-w-[200px]">
                  <div className="flex justify-between items-baseline">
                    <div className="text-[11px] text-zinc-100">{sch?.name ?? item.schematicId}</div>
                    <div className="text-[9px] text-zinc-600">#{item.serialNumber}</div>
                  </div>
                  <div className={cn("text-[10px] uppercase mt-0.5", qualityColor(item.overallQuality))}>
                    {qualityLabel(item.overallQuality)} · {Math.round(item.overallQuality * 100)}%
                  </div>
                  <div className="mt-1 space-y-0.5">
                    {sch?.experimentalProps.map(p => (
                      <div key={p.id} className="flex justify-between text-[10px]">
                        <span className="text-zinc-500">{p.label}</span>
                        <span className="text-zinc-200 tabular-nums">
                          {item.finalPropertyValues[p.id]}{p.unit ? ` ${p.unit}` : ""}
                        </span>
                      </div>
                    ))}
                  </div>
                </div>
              );
            })}
          </div>
        </section>
      </div>

      <footer className="px-3 py-2 text-[10px] text-zinc-600 max-w-[1800px] mx-auto">
        <p>
          <span className="text-amber-400">Engine:</span> Pure TypeScript, zero dependencies.
          <span className="text-zinc-700 mx-2">|</span>
          <span className="text-amber-400">Data:</span> {spawnTable.length} resources across 10 planets, {SCHEMATIC_LIST.length} schematics.
          <span className="text-zinc-700 mx-2">|</span>
          Seed: {seed}. All RNG is mulberry32, replayable.
        </p>
      </footer>
    </div>
  );
}

// ---------------------------------------------------------------------------
// Sub-components
// ---------------------------------------------------------------------------

function TopBar({ seed, onRegen }: { seed: number; onRegen: (s: number) => void }) {
  const [input, setInput] = useState(String(seed));
  return (
    <header className="border-b border-zinc-800 bg-zinc-900 px-3 py-2">
      <div className="max-w-[1800px] mx-auto flex items-center justify-between">
        <div className="flex items-center gap-3">
          <div className="text-amber-400 text-sm font-bold tracking-tight">
            SWG PRE-CU CRAFTING CORE
          </div>
          <div className="text-[10px] text-zinc-500 uppercase tracking-wider">
            Offline Tech Demo / Engine Prototype
          </div>
        </div>
        <div className="flex items-center gap-2">
          <span className="text-[10px] text-zinc-500">Spawn Seed:</span>
          <input
            className="bg-zinc-800 border border-zinc-700 px-2 py-0.5 text-[11px] w-24 font-mono"
            value={input}
            onChange={e => setInput(e.target.value)}
          />
          <button
            onClick={() => {
              const n = parseInt(input, 10);
              if (!isNaN(n)) onRegen(n);
            }}
            className="px-2 py-0.5 text-[11px] border border-zinc-700 hover:border-amber-500 hover:text-amber-400"
          >Respawn</button>
          <button
            onClick={() => {
              const n = Math.floor(Math.random() * 999999);
              setInput(String(n));
              onRegen(n);
            }}
            className="px-2 py-0.5 text-[11px] border border-zinc-700 hover:border-sky-500 hover:text-sky-400"
          >Random Seed</button>
        </div>
      </div>
    </header>
  );
}

function ResourceInspector({ res }: { res: Resource }) {
  const cls = getClass(res.classId);
  return (
    <div className="border-t border-zinc-800 p-2 bg-zinc-950/50">
      <div className="text-[11px] text-zinc-100 font-bold">{res.name}</div>
      <div className="text-[10px] text-zinc-500">
        {cls.name} · {formatPlanet(res.planetId)} · {res.availableUnits.toLocaleString()}u in vein
      </div>
      <div className="mt-1.5 grid grid-cols-2 gap-x-2 gap-y-0.5">
        {ALL_STATS.map(s => (
          <StatPill key={s} stat={s} value={res.stats[s]} />
        ))}
      </div>
    </div>
  );
}

function FinalItemView({ item, schematic, resourcesById }: { item: CraftedItem; schematic: Schematic; resourcesById: ReadonlyMap<string, Resource> }) {
  return (
    <div className="flex-1 overflow-auto p-2">
      <div className="text-[10px] uppercase text-emerald-400 mb-2">
        Crafting Complete — Item Serialized
      </div>
      <div className="border border-emerald-800 bg-emerald-950/20 p-3 mb-2">
        <div className="flex items-baseline justify-between">
          <div className="text-xl font-bold text-zinc-100">{schematic.name}</div>
          <div className="text-[10px] text-zinc-500">#{item.serialNumber}</div>
        </div>
        <div className={cn("text-sm mt-1 font-bold uppercase tracking-wide", qualityColor(item.overallQuality))}>
          {qualityLabel(item.overallQuality)} ({Math.round(item.overallQuality * 100)}%)
        </div>
      </div>

      <div className="grid grid-cols-2 gap-2 mb-2">
        {schematic.experimentalProps.map(p => {
          const v = item.finalPropertyValues[p.id];
          return (
            <div key={p.id} className="border border-zinc-800 bg-zinc-900/60 p-2">
              <div className="text-[10px] uppercase text-zinc-500">{p.label}</div>
              <div className="text-2xl text-cyan-300 tabular-nums font-bold">
                {v}{p.unit ? <span className="text-xs text-zinc-500 ml-1">{p.unit}</span> : null}
              </div>
              <div className="text-[9px] text-zinc-600 mt-0.5">
                range {p.minValue} – {p.maxValue}
              </div>
            </div>
          );
        })}
      </div>

      <div className="text-[10px] uppercase text-zinc-500 mb-1">Resource Bill of Materials</div>
      <div className="border border-zinc-800 bg-zinc-900/30 p-1.5 mb-2 space-y-0.5">
        {item.resourcesUsed.map(u => {
          const r = resourcesById.get(u.resourceId);
          return (
            <div key={u.resourceId} className="text-[10px] text-zinc-300 flex justify-between">
              <span>
                <span className="text-zinc-100">{r?.name ?? u.resourceId}</span>
                <span className="text-zinc-600 ml-1">({r ? getClass(r.classId).name : "?"})</span>
              </span>
              <span className="text-zinc-400 tabular-nums">{u.units}u</span>
            </div>
          );
        })}
      </div>

      <div className="text-[10px] uppercase text-zinc-500 mb-1">Experiment Roll History</div>
      <div className="border border-zinc-800 bg-zinc-900/30 p-1.5 max-h-40 overflow-auto space-y-0.5">
        {item.experimentLog.map(e => {
          const prop = schematic.experimentalProps.find(p => p.id === e.propertyId);
          return (
            <div key={e.attempt} className="flex items-center gap-2 text-[10px]">
              <span className="text-zinc-600 w-6">#{e.attempt}</span>
              <span className={cn("px-1 border text-[9px]", outcomeColor(e.outcome))}>{e.outcome}</span>
              <span className="text-zinc-400 flex-1 truncate">{prop?.label}</span>
              <span className="text-zinc-500">{e.pointsSpent}pt</span>
              <span className={cn(
                "tabular-nums w-10 text-right",
                e.boxDelta > 0 ? "text-emerald-300" :
                e.boxDelta < 0 ? "text-rose-400" : "text-zinc-500"
              )}>{e.boxDelta > 0 ? "+" : ""}{e.boxDelta}</span>
              <span className="text-zinc-600 w-8 text-right">d{e.dieRoll}</span>
            </div>
          );
        })}
      </div>
    </div>
  );
}
