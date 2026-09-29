---
doc_id: EVIDENCE-RECON-PASS-2026-09-12
title: TCIndustries — Evidence Reconciliation Pass (GDD §34 Step 0)
version: "1.0"
date: 2026-09-12
author: UNVERIFIED
assessor: UNVERIFIED
status: EVIDENCE / ANALYSIS (no design authority)
authority_class: null
human_approved: false
supersedes: []
superseded_by: null
controlling_documents: [MASTER-GDD-v1.1.1, HUMAN-RULINGS-REGISTER, SWG-PRE-CU-README]
reading_priority: mandatory
---

# TCIndustries — Evidence Reconciliation Pass

**This document performs GDD §34 Step 0 ("Evidence Reconciliation Pass — Inspect any existing prototype; produce GDD ↔ prototype ↔ decision evidence matrix").**

It is an **evidence and analysis** artifact. It changes **no** GDD status, promotes **nothing**, and makes **no** game-design decision. All prototype findings are evidence only, per PROT-001. Where this document reports a tension, it records the tension; it does not resolve it (resolution requires human decision).

---

## 1. Purpose and Scope

| Item | Detail |
|---|---|
| Trigger | GDD §34 Step 0 — open process requirement §36 item 10; also flagged in `TCIndustries-Master_GDD_v1.0_Canonical_Audit` as "Phase 0 — Evidence Reconciliation". Unblocked 2026-09-12 by PROT-002 identity confirmation. |
| Deliverable | GDD ↔ prototype ↔ decision evidence matrix (below). |
| Primary object of inspection | `sources/swg-pre-cu/craft-seed2-1-pro-preview/` — the Seed-2.1 TypeScript/Vite/React prototype, confirmed by the project owner as the prototype referenced by GDD PROT-002. |
| Scope boundary | The tree is **SWG Pre-CU crafting content** (zero TCIndustries-specific strings). Reconciliation must therefore judge what the prototype proves for TCIndustries **through scope accounting**, not by treating its content as TCIndustries design. |
| Two-track boundary | `proposals/SWG_Code_Due_Diligence_Checklist_v0.1.md` (PROPOSED procedure, separate track) verifies the demo **runs as claimed**. This pass judges only **what it proves**. This pass does **not** build, run, or execute the prototype; it inspects the source. |
| Status effect | None. PROT-002 remains PROTOTYPE / ASSUMPTION. AS-001 remains ASSUMPTION. No GDD element changes status as a result of this pass. |

---

## 2. Objects of Inspection

All paths are under `sources/swg-pre-cu/`. All are HISTORICAL predecessor material; none inherit TCIndustries authority (SWG-001, SWG-002, `swg-pre-cu/README.md` governance notes).

| Artifact | Path | Provenance | Role in this pass |
|---|---|---|---|
| **Seed-2.1 crafting core prototype (PRIMARY)** | `craft-seed2-1-pro-preview/` | Group B; arena.ai; author seed-2.1-pro-preview LLM; deployed demo per `swg-pre-cu/README.md` | Confirmed identity of GDD PROT-002 prototype. Full source inspection. |
| Claude-opus-5-max architecture demo (SECONDARY) | `craft-claude-opus-5-max/` | Group B; author claude-opus-5-max LLM | Related predecessor demo (FSM craft bench, spawn browser, self-test). Not PROT-002. Inspected `engine/craft.ts`. |
| Phase 0–2 server (SECONDARY) | `code-phase-0/1/2/` | Group A lineage (Perplexity — GPT-5.2 GDD) | Go server: auth, spatial, movement, skills/professions/XP. Inspected Phase 2 README + structure. |
| Phase 0–3 server (SECONDARY) | `code-phase0-3-claudecode/` | Group C; author unknown; embedded GDD v2 | Adds Phase 3 Combat, verify-by-running methodology. Inspected `CLAUDE.md`. |
| Track A/B scaffold (SECONDARY) | `track-a-b-code/` | Group C; author unknown | Mostly class stubs + small reference functions (combat resolution, resource spawn, inflation sim, experimentation outcomes, XP curve). Inspected representative files. |
| Java skill calculator (SECONDARY) | `java-swg-skillcalc/` | Group C; author unknown | Working Swing skill calculator (250-pt pool, prereq chains, refunds, master lock). Inspected README. |

**Inspection depth (primary):** all engine files read in full —
`types.ts`, `crafting.ts`, `resourceGen.ts`, `schematics.ts`, `rng.ts`, `data/resourceTree.ts`, `index.ts`, `App.tsx`, `components/StatBar.tsx`, plus project metadata (`package.json`, `tsconfig.json`, `vite.config.ts`).

**Verification note:** this pass is source-inspection only. Whether the demo builds/runs as claimed is the due-diligence checklist's question (`proposals/SWG_Code_Due_Diligence_Checklist_v0.1.md`), not this pass's.

---

## 3. Seed-2.1 Prototype — Behaviour Observed (EVIDENCE)

The prototype is an interactive tech demo of an SWG Pre-CU-style crafting core. Everything below is **EVIDENCE** (behaviour demonstrated by the source). It is SWG-flavoured content and is recorded as HISTORICAL-adjacent evidence, **not** TCIndustries design.

| # | Observed behaviour | Where (file:line) |
|---|---|---|
| E1 | Resource class hierarchy with subtype inheritance; schematic slots accept exact match or subtypes. | `resourceTree.ts:131-139` (`isSubtypeOf`); `crafting.ts:44` |
| E2 | Resource instance model: unique id, two-part random name, class, planet, 9 attributes (OQ/PE/UT/CR/CD/DR/HR/MA/SR), `availableUnits` "vein" quantity. | `types.ts:35-43`; `resourceGen.ts:70-103` |
| E3 | Attributes sampled at spawn and immutable for the spawn's life. | `types.ts:41` |
| E4 | Stat generation uses a left-skewed distribution (max of N uniform rolls); OQ gets extra skew ("rarer to be 900+"); irrelevant class stats stay 0. | `rng.ts:50-56`; `resourceGen.ts:48-59` |
| E5 | Deterministic, replayable seeded RNG (mulberry32) for spawn tables and experiment rolls. | `rng.ts:12-21`; `App.tsx:654` |
| E6 | A global spawn table across 10 planets, 6–10 resources per planet. Spawn distribution across planets is explicitly uniform ("we do uniform here — this is a demo knob"). | `resourceGen.ts:61-103` |
| E7 | Schematic model: named slots requiring a resource class + units + per-slot stat weights + per-slot contribution to experimental properties; schematics carry an experiment-point budget, complexity, and a profession tag. | `types.ts:66-102`; `schematics.ts:53-93,100-138,145-192,198-236` |
| E8 | Assembly validation: rejects empty slots, type-incompatible resources, insufficient units. | `crafting.ts:36-58` |
| E9 | Per-property weighted resource score from assigned slots (weights normalised; missing weights treated as 0). | `crafting.ts:70-84` |
| E10 | Resource cap ("garbage in, garbage out"): the resource-derived ceiling that experimentation can never exceed. | `crafting.ts:86-115`; UI `App.tsx:529-533` |
| E11 | Experimentation engine: d100 roll vs. skill thresholds; outcomes amazing / great / success / failure / critical; box-delta gains; up to 5 points per attempt; critical failures can subtract boxes; over-experimenting is permitted. | `crafting.ts:120-167,231-263` |
| E12 | Two skill profiles (novice vs master) change thresholds — a coarse two-band skill expression. | `crafting.ts:135-147`; `App.tsx:368` |
| E13 | Finalization: value lerped from min..max by realized fraction (cap × box ratio); inverted stats (lower is better) handled; overall quality scalar 0..1 derived; quality labels Trash…Exceptional…LEGENDARY. | `crafting.ts:286-344`; `App.tsx:60-76` |
| E14 | CraftedItem is an immutable record: id, schematic id, serial number, resource bill of materials, final property values, experiment log. | `types.ts:134-142`; `App.tsx:613,626` |
| E15 | UI: global spawn survey table with "+500 units" action, inventory stacks, schematic picker, craft bench, crafted-item history. Thin "debug shell" over the engine. | `App.tsx:252-345,398-645` |
| E16 | Engine structured as pure, dependency-free TypeScript functions explicitly designed to port 1:1 to a C++ engine module. | `crafting.ts:1-6`; `index.ts:1-3`; `App.tsx:1-6` |
| E17 | No manufacturing/factory/automation systems. No economy/trade/vendor layer. No skill-progression or profession/spending system beyond the two static profiles. No combat, services, cities, or organisations. No time-based lifecycle/despawn ticking (despawn exists only as a type comment). | Absence across `src/engine/*` |

**Scope accounting:** the tree contains zero TCIndustries-specific strings; all names, planets, stats, schematics (CDEF Pistol, Stim Pack B, Mineral Harvester, Travel Biscuit), and professions (Weaponsmith/Medic/Architect/Chef) are SWG Pre-CU content. Consequences:

- The prototype **demonstrates a family of SWG-era crafting mechanics is implementable and interactive**. It does **not** demonstrate any TCIndustries mechanic.
- Any mechanical resemblance to TCIndustries GDD text is because both draw on the same SWG-era inspiration. Under SWG-001/SWG-002 and PROT-001, that resemblance confers **no** authority.
- The prototype's numeric values are SWG-historical demo values, not TCIndustries-tuned values, and are **not** candidates for lifting (BAL-001).

---

## 4. GDD ↔ Prototype ↔ Decision Evidence Matrix

Statuses and decision citations are taken from `canonical/gdd/TCIndustries_Master_GDD_v1.1.1_Status_Patch.md` and `governance/TCIndustries_Human_Rulings_Register.md`.

Verdict legend — **EV**＝behavior evidenced by prototype; **PA**＝partially evidenced; **NE**＝not evidenced; **TN**＝tension surfaced (recorded, not resolved); **OS**＝out of scope / no bearing.

### 4.1 Pillars

| GDD element | Status / decision | Prototype observation | Verdict & note |
|---|---|---|---|
| PIL-001 Discovery & Gold Rush | LOCKED (explicit pillar) | Resources are variable-quality (E4), regional via planets (E6), finite via `availableUnits` (E2). No despawn/time cycle, no discovery *event* mechanics — spawn table is static per seed (E17). | **PA.** Scarce/variable/regional resource modelling is demonstrated; temporary-lifecycle and event-generation are not. No status change. |
| PIL-002 Identity & Reputation | LOCKED (explicit pillar) | Crafted items carry a serial number, BOM, and experiment log (E14); an item-identity family is demonstrated. No crafter/organisation/brand association; no reputation emergence. | **PA.** Item-level identity evidenced; creator/brand/provenance association and reputation are not. |
| PIL-003 Interdependence | LOCKED (explicit pillar; HD-EIC-01–09 governing) | Single-character crafting loop only; no professions-interaction, no economy layer, no multi-account dimension (E17). | **NE / OS.** The prototype exercises none of the residual channels adjudicated in the Architecture C falsification chain (α/β/γ′). It is not evidence for or against HD-EIC-05. HD-EIC-07 gate untouched. |

### 4.2 Resources

| GDD element | Status / decision | Prototype observation | Verdict & note |
|---|---|---|---|
| RES-002 Resource design goal | LOCKED (explicit) | Differentiation (attributes), scarcity (vein quantity), regional placement (planets) demonstrated. Trade/events absent. | **PA.** |
| RES-003 Taxonomy | PROPOSED | Tree of family → class → subtype with per-class applicable attributes (E1). Processing options/material state absent. | **PA.** Taxonomy category subset demonstrated (SWG form); TCIndustries taxonomy TBD. |
| RES-004 Resource instances | PROPOSED | Instances: id, name, class, planet, attributes, remaining quantity; lifecycle/time fields absent (E2, E3). | **PA.** Individualised, distinguishable instances demonstrated. |
| RES-005 Resource attributes | PROPOSED (sets/ranges TBD) | Nine fixed attributes mapped to product outcomes via stat weights/caps (E4, E10); exact TCIndustries attribute set remains TBD. | **PA.** Attribute→outcome mapping family demonstrated; numbers SWG-historical. |
| RES-006 Lifecycle & scarcity | PROPOSED (math TBD) | `availableUnits` (extraction-limited vein) present; cycling/depletion/despawn not implemented (E17). No spawn-rate/duration maths (BAL-001 forbids inventing them). | **NE** for lifecycle cycling; **PA** for finite extraction quantity. |
| RES-007 Surveying & information | PROPOSED | Global spawn table is visible to the player; "survey" is a flat "+500 units" button (E15). No information asymmetry, no surveying skill, no information value. | **PA/NE.** Trivialised relative to the GDD intent. Noted as a tension (4.4 T2). |

### 4.3 Crafting

| GDD element | Status / decision | Prototype observation | Verdict & note |
|---|---|---|---|
| CRFT-001 Differentiated products | LOCKED — principle only (HD-GDD-01; quality tiers/experimentation/Exceptional/factory rules open) | Same schematic with different resources/experimentation produces different final values and quality labels (E10–E14). | **EV** at the principle level: differentiated, non-commodity outcomes are demonstrably runnable. Implementation details remain open. |
| CRFT-002 Schematics & recipes | PROPOSED (structure TBD) | Slot-based schematic model with class-typed requirements, units, stat weights (E7, E8). | **EV** of an SWG-flavoured schematic family; TCIndustries structure TBD. |
| CRFT-003 Experimentation | PROPOSED (model TBD) | d100 skill-profile box model with resource caps bounding outcomes (E9–E11). | **EV** of an SWG-style experimentation family; TCIndustries model TBD. |
| CRFT-004 Provenance | PROPOSED | Serial number + BOM + experiment log retained on items; no crafter/org/resource-source chain, no visibility rules (E14). | **PA/NE.** Item record lineage evidenced; crafter provenance chain not. |
| CRFT-005 Skill & tool requirements | PROPOSED (TBD) | No tool/station requirements; skill expressed only via two static profiles (E12, E17). | **NE** for tools/requirements; **PA** for skill affecting outcomes. |
| CRFT-006 No instant mastery of everything | PROPOSED (demoted 2026-08-25; mechanism unresolved) | No spending/budget layer; nothing to evidence mastery limits. | **NE.** Consistent with CRFT-006 being an unresolved mechanism question. |

### 4.4 Manufacturing, Items, Economy, Commerce, Services

| GDD element | Status / decision | Prototype observation | Verdict & note |
|---|---|---|---|
| MFG-001 Controlled automation | LOCKED (explicit) | No manufacturing/automation systems of any kind. | **NE.** |
| MFG-002/003 Manufacturing | PROPOSED (TBD) | None. | **NE.** |
| MFG-004 Anti-Factorio drift | DERIVED CONSTRAINT (VIS-006, MFG-001) | None. | **NE.** |
| ITM-001 Item identity/provenance | PROPOSED | Quality attributes + serial number + records (E13–E14); no decay; no creator/owner association. | **PA.** |
| ITM-002 Durability/maintenance | PROPOSED / TBD | None. | **NE.** |
| ITM-003 Stacking/uniqueness | PROPOSED | Resources stack by spawn reference; crafted items are individual records (E14, E15). | **PA.** |
| ECO-001 Player-driven economy | LOCKED (explicit) | No economy layer. | **NE.** |
| ECO-002 Regional economies | PROPOSED | Planets exist but spawn distribution is uniform; regional bias is an unconfigured "demo knob" (E6). | **PA/NE.** Regional differentiation not implemented. |
| ECO-003 No pure spreadsheet optimisation | DERIVED CONSTRAINT (VIS-006, VIS-003) | No economy. | **NE.** |
| ECO-004 Currency/pricing | TBD | None. | **NE.** |
| RET-001/002/003 Retail/commerce | LOCKED (RET-001) / PROPOSED | None. | **NE.** |
| SERV-001/002 Services | LOCKED (SERV-001) / PROPOSED | None. | **NE.** |

### 4.5 Loops, Balance, Validation, Prototype

| GDD element | Status / decision | Prototype observation | Verdict & note |
|---|---|---|---|
| LOOP-001 Interconnected loops | PROPOSED | Survey → extract → craft → experiment → item loop is runnable (E15); retail/service/combat/org loop legs absent (E17). | **PA.** A non-combat crafting leg is technically demonstrable. |
| LOOP-002 Non-combat viability | LOCKED (explicit direction) | Medic/Chef/Architect schematics show non-combat crafting execution (E7). Evidence of a technically working non-combat loop family; not TCIndustries content. | **PA** (technical): corroborates implementability only. |
| BAL-001 No premature numeric tuning | LOCKED (consistent philosophy) | Prototype contains many demo numbers (thresholds, costs, caps, quality ranges) — all SWG-historical values (E10–E13). | **TN / no conflict** recorded: none of these numbers are TCIndustries tuning; none are lifted. See 4.6 T1. |
| TEST-001 Design invariants | PROPOSED | Pure-function engine with explicit arithmetic (E16); no invariant/test harness in the tree (self-test exists in the claude-opus-5-max sibling, not here). | **PA** for simulatability architecture. |
| TEST-002 Simulation preference | PROPOSED | Deterministic, seeded, pure-TS engine is well-suited to offline simulation (E5, E16); not itself a TCIndustries simulation. | **PA** (architectural affinity; observation only). |
| PROT-001 Prototype authority | LOCKED (explicit principle) | This pass operates under it: prototype = evidence, GDD governs. | **OS** — honoured. No change. |
| PROT-002 Existing prototype | PROTOTYPE / ASSUMPTION (reconcile in future pass) | Identity now confirmed (owner, 2026-09-12) and behaviours matrixed above; "must be reconciled in a future pass" requirement is this document. | **EV** — this pass performs the reconciliation. Text/status stands; no change. |
| AS-001 Functional prototype exists | ASSUMPTION | Complete, self-consistent source tree present (E16, metadata); runnable-claim verification deferred to due-diligence track. | **EV** (existence/completeness). Run-verification deferred. No status change. |
| SWG-001/002 Inspiration/strict separation | LOCKED (explicit) | Prototype is entirely SWG Pre-CU content; zero TCIndustries promotion here. | **OS** — honoured. No change. |
| §36 "fully frozen canon requires prototype evidence reconciliation" | Inherited note | This pass is the reconciliation step. It does not itself validate canon. | **EV** — executed; nothing validated by virtue of this pass. |

---

## 4.6 Tensions Surfaced (Recorded, Not Resolved)

These are observations. Resolution is a human decision and is **not** proposed here.

- **T1 — SWG demo numbers vs BAL-001.** The prototype is full of concrete values (experiment thresholds, point budgets, complexity deltas, quality-band ranges: `crafting.ts:120-167`, `schematics.ts`, `rng.ts:50-56`). BAL-001 (LOCKED) prohibits premature TCIndustries numeric tuning. These coexist without conflict **only** because the values are SWG-historical demo values with no TCIndustries status. Any future lifting of these values into TCIndustries design would cross the BAL-001 boundary and require explicit human authorisation.
- **T2 — Flat "survey" vs RES-007 information asymmetry.** The GDD proposes surveying as meaningful information economics (RES-007). The demo exposes a global spawn table with a one-click "+500 units" action. The demo does not implement the feature the GDD intends; the demo's form cannot be read as evidence for or against RES-007.
- **T3 — Exceptional-quality vocabulary.** The demo applies SWG-style quality labels including "Exceptional"/"LEGENDARY" (`App.tsx:60-76`). GDD §36 lists "Factories can never produce Exceptional-tier output" among explicitly non-canonical historical items, and CRFT-001's authority note leaves "whether Exceptional exists" open. The demo's vocabulary is historical evidence only and does not open or close the TCIndustries question.
- **T4 — Regional uniformity.** The generator comment acknowledges a real-SWG regional bias ("Tatooine has lots of Mineral… for brevity we do uniform here" — `resourceGen.ts:86-88`). ECO-002 (PROPOSED) anticipates regional economies. The demo neither implements nor refutes regional differentiation.
- **T5 — Standing coherence gap (pre-existing, not caused by this prototype).** GDD §34 frames the Economic Interdependence Core as upcoming, unstarted work, with no reference to Architecture C's retirement (documented in `TCIndustries_Documentation_Architecture_Audit_v1.md` §5.1 as the corpus's largest coherence risk). This pass neither widens nor fixes that gap.

---

## 5. EIC Cross-Reference (Explicit Non-Interference Statement)

- The Architecture C falsification evidence chain (`investigations/eic/`) adjudicated residual channels (α soft-pressure+residuals, β formal character limits, γ facility/information concentration) under the HD-EIC-01–04 authorisations, using simulated functional shapes. The Seed-2.1 prototype was **not** part of that evidence and exercises **none** of those channels.
- Consequently the prototype is **not** evidence for or against HD-EIC-05 (falsification acceptance), and its reconciliation here does **not** interact with HD-EIC-09 or the HD-EIC-07 structural-boundary gate. The next authority-bearing action for EIC remains human-supplied boundaries (HD-EIC-07).

---

## 6. Outcomes and Documentation Consequences

1. **Reconciliation performed.** PROT-002's earlier note ("must be reconciled in a future pass") is satisfied by this document; the GDD text and statuses are untouched.
2. **No status changes.** Every verdict above is EVIDENCE; no PROPOSED/TBD/ASSUMPTION/PROTOTYPE element was promoted, demoted, or re-classified, and no LOCKED element was altered.
3. **No design decisions made.** Notable: no verdict recommends adopting or rejecting any prototype mechanic for TCIndustries.
4. **Runnable-claim verification is deferred** to the SWG Code Due Diligence checklist track (`proposals/SWG_Code_Due_Diligence_Checklist_v0.1.md`).
5. **Evidence value captured.** The primary durable observation for future system design is *architectural*: a pure, deterministic, C++-portable crafting-core engine with resource-derived caps and skill-profile experimentation is a demonstrably implementable shape for a craft loop (relevant to OQ-003/OQ-004 context, and to TEST-002-style simulation) — recorded as evidence, not as an approved design.

---

## 7. Method and Traceability

- **Artifacts inspected:** all files under `sources/swg-pre-cu/craft-seed2-1-pro-preview/` (all engine + UI + metadata); `craft-claude-opus-5-max/src/engine/craft.ts`; Phase 2 README; `code-phase0-3-claudecode/swg-server/CLAUDE.md`; Phase 0/1/2 READMEs; `track-a-b-code` representative files (`crafting_system.py`, `resource_generation.py`, `resource_spawn_generator.py`, `crafting_experimentation.py`, `economy_inflation_sim.py`, `xp_progression_curve.py`, `combat_resolution.py`, `phase_backlog.md`); `java-swg-skillcalc/README.md`. Source inspection only; nothing executed.
- **GDD source:** `canonical/gdd/TCIndustries_Master_GDD_v1.1.1_Status_Patch.md` (§11–17, §25–29, §31–36).
- **Decision source:** `governance/TCIndustries_Human_Rulings_Register.md` (HD-EIC-01–09, HD-GDD-01); `governance/eic/TCIndustries_EIC_Human_Ruling_HD-EIC-09_2026-09-12.md`.
- **Identity confirmation source:** `sources/swg-pre-cu/README.md` ("PROT-002 identity — CONFIRMED 2026-09-12").
- **Formal classification:** This document is EVIDENCE / analysis. It is not a proposal, ruling, or authority-bearing document.

---

*End of Evidence Reconciliation Pass.*