# SWG Pre-CU Phase 4 — Resources & Crafting: Proposal Plan v0.1

**Status:** PROPOSED plan. No authority. No design decision. Human approval required
before any implementation, and separately before any numeric value is treated as tuned.
**Author / Assessor:** Muse Spark (`opencode/muse-spark-1.3-contributor-free`) via OpenCode —
research and planning only; all authority decisions remain human.
**Date drafted:** 2026-09-14.
**Controlling references:** `AGENTS.md`; `proposals/SWG_Phase3_Combat_Core_Proposal_v0.1.md`
(executed — fork at `testbed/swg-phase3-combat/`, review gate pending);
`investigations/SWG_Code_Due_Diligence_Evidence_Log_2026-09-14.md` (§§5–6, gap rows 2–4);
`governance/TCIndustries_Testbed_Reuse_Human_Ruling_HD-TST-01_2026-09-14.md`;
predecessor GDD `sources/swg-pre-cu/SWG_PreCU_GDD.md` (§§7, 8.2.1, 8.3.1–8.3.5, 10, 11,
12, 17.2.2, 26, 28, 30 — HISTORICAL, never TCIndustries canon).

---

## 1. Objective

Implement **Phase 4 — Resources & Crafting** on the testbed fork (Phases 0–3 complete
there, subject to the pending Phase 3 review gate), satisfying the predecessor-GDD §30
exit criteria verbatim: *"A player can survey, harvest a resource, craft an item from a
schematic with experimentation, and equip a crafted weapon/armor with correctly-computed
stats."* This is the HD-TST-01-authorised pipeline exercise
(survey → harvest → craft-with-experimentation → equip), now as a build plan.

Entry criteria (§30: Phase 3 complete — "crafted gear needs the combat/item system to
matter") are MET in the fork (Phase 3 verified 6/6 this session), conditional on the
review gate passing first — see §10.

## 2. Research basis

### 2.1 Predecessor-GDD map (HISTORICAL input, not authority)

| GDD section | Specifies | Phase-4 relevance |
|---|---|---|
| 10.2 core loop | 9-step loop: schematic → gather → tool → select → allocate → customize → experiment → craft → prototype/factory | Backbone of work packages R1–R7 |
| 10.3 schematics | Sources (trainer/loot/quest/reverse-engineer); attributes (resources, complexity, exp attributes, category, profession lock); DL-44 worked example | Registry + generic schematics (sources MVP: trainer-grant; §9 decision) |
| 10.4 resources | 8 categories; 10 stats 0–1000 + OQ-as-average; rarity snap (900+ ~5%, <500 ~40%) | Taxonomy + stat model (generic categories; §9) |
| 10.4.3 spawning | Regions 5–20 km, 7–14 day life, relocate-on-despawn, overlap, unique stats, procedural names, no telegraphing | Spawn service (time-compressed; §6) |
| 10.5/11.3/26 surveying | 5 tool categories; waypoint ±500m → triangulate; 10 s cooldown; 1–3 units/sample, 2–5 s animation; diminishing returns + 1%/hr recovery; tool tiers (±500/±250/±100 m); Scout/Ranger bonuses; hot-spot competition (no queues) | Survey/sample mechanics + tool tiers |
| 10.6/11.4 harvesters | Personal/Medium/Heavy (50/100/200 units/hr × concentration); hoppers 10/25/50k; weekly maintenance 500–5000; power generators; 100 m no-clustering; permissions; tied to spawn life | Harvester MVP (power: §9 decision) |
| 10.7 experimentation | 10/50/30/8/2 outcome table (+5–10/+2–5/+0–2/0/−2–5%); ≤5 points/attempt; 3–5 rounds; tool bonuses +10/+20%; base = schematic + quality×multiplier; final = base × (1+bonus) | Session FSM (demos already prove the shape) |
| 10.8 factories | Prototype → batch (identical stats); 1 item/hr/slot, 5–10 slots; maintenance + power | OPTIONAL R8 only (EIC/MFG caution; §9) |
| 10.9/8.2.1/8.3.1–5 | Artisan foundation (Engineering/Business/Experimentation/Complexity); 6 elite crafts + Merchant listed in §30 scope | Schematic gating (elite-tension note §5) |
| 11.2.1 spawn pseudocode | Executable spec: gauss(500,150) stats, OQ average, 5–20 km radius, 50–90 peak concentration, 7+rand(0–7) days, 20–50 spawns/planet | Direct implementable shape (compressed) |
| 11.2.2–11.2.3 rarity/distribution | OQ bands 10/30/40/15/5%; concentration 30/50/20%; per-planet type tables (SW planets — NOT reusable, develop independently) | Curves adoptable; planet tables must be re-authored generic |
| 11.5 storage/transport | 80-slot inventory, 100k stacks by spawn ID (no cross-ID stacking), weight, house chests, bazaar/direct trade | Minimal stack-store MVP (houses/bazaar OUT) |
| 11.6 market | Price bands as AI-agent guidance; "do NOT hardcode prices" | Guidance only — no price logic in Phase 4 |
| 17.2.2 corpse harvesting | Any-player harvest in 10-min [ASSUMPTION] window; fixed yield type per template; per-corpse quality roll (NOT regional system) | Corpse-harvest hook (fork has kill path, no corpse timer — add) |
| 28 Item/Schematic | Item + Schematic + Profession/SkillTree/SkillBox entity shapes (stats maps, condition, crafter attribution) | Persistence shapes |
| 7 XP / 12 sinks | Crafting-XP types; harvester/factory maintenance as credit sinks | XP awards (numbers §6); maintenance MVP |

### 2.2 Existing implementations (evidence, not authority)

- **Seed demo engine** (`craft-seed2-1-pro-preview/src/engine/`): full session FSM
  (init → assign → assembly → experiment → finalize, `crafting.ts`); 4 schematics with
  slots/points/complexity; spawn tables 6–10/planet, 5000–25000 units; class taxonomy;
  seeded RNG; Novice/Master skill threshold profiles. Direct PATTERN source for R6.
- **Claude demo engine** (`craft-claude-opus-5-max/src/engine/`): deterministic galaxy
  timeline + cap windows + lane fidelity (SelfTest 6/6); points-budget FSM + ceilings.
  PATTERN source for R1 determinism discipline + R6 ceiling checks.
- **Track A/B**: `resource_spawn_generator.gen()` is near-verbatim the GDD 11.2.1
  pseudocode (gauss 500/150, OQ-as-mean, radius 5–20) — recorded as DERIVED-FROM or
  CONVERGENT, not asserted either way; usable as cross-check, not dependency.
  `crafting_experimentation` OUTCOMES matches the 10.7.1 table exactly (data, no roll fn);
  `crafting_system.py` / `resource_generation.py` are one-line stubs (nothing transfers).
- **Java Artisan** (`ProfessionFactory.buildArtisan`): 18-box Artisan WITH a Surveying
  branch (range/concentration bonuses) — GDD 8.2.1's Artisan has NO Surveying branch
  (Engineering/Business/Experimentation/Complexity instead). Non-canonical variant either
  way; the survey-bonus SHAPE is the transferable observation (links §26 Scout-bonus logic).
- **Chat export**: negligible (1 Phase-4 mention, 1 schematic) — provenance color only.
- **Fork state**: Phases 0–3 green; NO resource/crafting code anywhere in Go trees
  (evidence §6 rows 2–4); Go-side session/FSM/timeline work starts from zero, informed by
  the demo engines above. Corpse path exists only as instant-despawn on kill (fork
  `KillCreatureInstance` analogue needed: add corpse-expiry + harvest interaction).

## 3. Base and working location (RECOMMENDED — human confirms)

Build in place on `testbed/swg-phase3-combat/` (new packages `internal/resources/`,
`internal/crafting/`, DB extensions, handlers, `cmd/phase4test/`); no new fork, no
touching `sources/`. Rationale: Phase 4's entry criterion IS the fork's Phase 3
(skill gates via `artisan_novice`, HAM-adjacent buff-less baseline, creature-harvest hook
into the Phase 3 kill path).

## 4. Scope IN (work packages, each verify-by-running before the next)

- **R1 — Spawn service** (`internal/resources/`): seeded lifecycle manager over the GDD
  11.2.1 shape — N active spawns in zone-0001, Gaussian stats + OQ, concentration
  gradient, rarity curves, relocate-on-despawn, idempotent seeding; lifecycle driven by
  the world tick under the §6 time-compression convention. Deterministic-by-seed
  (claude-demo discipline) so tests are reproducible.
- **R2 — Survey/sample**: category tools (5 generic categories), `/survey`-equivalent WS
  action with 10 s cooldown, nearest-spawn + distance + concentration response, waypoint
  approximation (±500 m basic tier), sample action yielding 1–3 units with diminishing
  returns; tool tiers (basic/crafted) + Artisan/Scout-bonus hooks (hooks only — Scout
  trees exist as data, Ranger doesn't).
- **R3 — Harvesters**: deed → place (100 m exclusion) → hourly-tick extraction
  (rate × concentration) → hopper (caps) → empty/redeed/permission-minimal (owner-only
  MVP); weekly maintenance pool (credit sink, §12-consistent). Power generators:
  §9 decision (recommend: deferred, harvester runs on maintenance pool, flagged).
- **R4 — Resource stacks**: minimal stack-store keyed by spawn ID (no cross-ID stacking);
  quantities as integers; NO full inventory/weight/houses (OUT).
- **R5 — Schematic registry**: seeded generic schematics (recommend 3: basic sidearm,
  basic plate, harvester deed — de-SWG'd analogues of the demo set; NO SW names).
  Acquisition MVP: auto-granted with Artisan Novice (trainer-sale needs Phase 5 vendors;
  §9 confirms).
- **R6 — Crafting sessions**: server-side session FSM (assign → assemble → experiment →
  finalize) adapted from the seed-demo pattern; Artisan-Novice gate; experimentation
  points + 10/50/30/8/2 table + base/final formulas; skill-threshold profiles
  (Novice/Master shaping); crafting-XP awards; output item with computed stats +
  crafter attribution (§28 shape); equip path for weapon/armor analogues (stats must
  demonstrably change combat-relevant values — close the Phase 4 exit-criteria loop).
- **R7 — Corpse harvest**: 10-min [ASSUMPTION] corpse window on the Phase 3 kill path;
  sample-tool interaction; fixed yield type per template; per-corpse quality roll
  (separate from regional system, per 17.2.2).
- **R8 — Factory MVP (OPTIONAL, default OUT)**: prototype → identical-stat batch,
  run slots, maintenance fee. Included ONLY on explicit approval given MFG-001/MFG-004
  automation caution and HD-EIC-07 (no EIC shaping); factory throughput numbers would
  need separate numeric sign-off.

## 5. Scope OUT (explicitly deferred, with homes)

Bazaar/vendor/retail/direct-trade UI (Phase 5); houses/storage/power structures;
droid/tailor/chef content; loot/quest/reverse-engineer schematic sources (need loot
tables/missions); PvP harvester destruction (no PvP); market/price logic (11.6:
emergent-only, never hardcoded); weight/encumbrance; Scout/Ranger bonus numbers
(hooks without values).
**GDD-internal tension recorded (human resolves):** §30 Phase 4 scope lists elites
8.3.1–8.3.5 (incl. Merchant) while the roadmap puts "remaining elite professions" in
Phase 9. Recommendation: MVP gates on Artisan Novice ONLY; elite trees deferred —
confirm or amend (§9).

## 6. Numeric authorization requests (values NOT set by this proposal)

GDD-GIVEN (implement as written, no approval needed): 10 stats 0–1000 + OQ-as-average;
rarity bands 10/30/40/15/5% and concentration 30/50/20%; survey cooldown 10 s; sample
1–3 units, 2–5 s action; tool precision ±500/±250/±100 m; harvester rates 50/100/200
× concentration, hoppers 10/25/50k, maintenance 500–5000/wk, 100 m exclusion;
experimentation table + bonus bands + ≤5 pts/attempt + 3–5 rounds + tool +10/+20%;
base/final stat formulas; corpse window 10 min ([ASSUMPTION]-tagged, implement flagged);
stack/weight/inventory figures as STORAGE SHAPES only (80 slots, 100k stacks);
11.6.1 price bands as GUIDANCE, never inputs.
GDD-SILENT / INDEPENDENT (each needs human approval OR interim-[ASSUMPTION] sign-off):
(1) **time-compression convention** — 7–14-day lifecycles, hourly ticks, weekly fees are
untestable live; propose lifecycle-in-minutes + tick-in-seconds mapping as provisional
test scaffolding, GDD ratios preserved; (2) single-zone MVP (per-planet tables collapse
to zone-0001); (3) generic resource taxonomy — recommended derivation: metal, polymer,
organic, chemical, flora, water (GDD 10.4.1 categories minus SW names), 5–10 types,
stat lanes per type; confirm or re-author; (4) generic schematic set (R5 trio) +
complexity/skill mapping; (5) XP reward numbers (survey/sample/harvest/craft/kill-adjacent);
(6) harvester power decision; (7) R8 factory numbers if approved; (8) equip-stat →
combat-effect mapping magnitudes (must move real combat values without tuning combat
itself — propose minimal delta convention, flagged).
PROPOSED interim convention (approval requested): Group-C-style provisional marking —
every non-GDD number ships flagged NON-CANONICAL in code comments + HYGIENE_NOTE.md,
reusing demo-engine values (seed spawn counts, skill thresholds) where they exist to
avoid inventing fresh tuning.

## 7. Acceptance criteria (verify, don't claim)

`go build ./...` + `go vet ./...` clean; fresh DB; phase0–3 tests green PLUS new
`phase4test` end-to-end (survey → sample → harvest-tick → craft-with-experimentation →
equip-with-stat-delta, experimentation variance observed across runs, corpse harvest
within window) with real pasted output; time-compression ONLY inside test config, never
in default server constants; hygiene note extended (every new string with file refs);
commit per verified package (R1→R7[, R8]).

## 8. Risks and dependencies

- Time-compression vs GDD fidelity: tests must prove LOGIC (lifecycle transitions,
  depletion, relocation), not durations — durations stay GDD-given in defaults.
- Factory/default-scope tension (§§4/R8, 5/MFG): R8 slightest scope creep touches
  automation-constraint design (TBD + MFG-004) — hard gate, no silent inclusion.
- Bazaar-less demand: with Phase 5 OUT, crafted items have no market — acceptance is
  mechanical (stats compute + equip), never economic.
- Ordering: Phase 3 review gate MUST pass before R1 starts (HD-TST-01 pattern);
  EIC/provenance/reputation untouched throughout.

## 9. Decisions requested of the human (the plan asks; it does not decide)

1. Base/location (§3) — confirm building on `testbed/swg-phase3-combat/`.
2. Time-compression convention (§6.1) — approve mapping or supply one.
3. Generic taxonomy + schematic trio (§§6.3–6.4) — confirm, amend, or re-author.
4. Schematic acquisition (§4/R5) — auto-grant with Novice vs alternative.
5. Harvester power (§4/R3) — defer (recommended) vs minimal pool.
6. Factory R8 — IN (with numeric sign-off) vs OUT (default).
7. Elite-profession tension (§5) — Artisan-Novice-only MVP vs alternative.
8. Equip-effect magnitudes (§6.8) — approve minimal-delta convention or supply.
9. Authorisation to implement on approval of 1–8 (cf. §10).

## 10. Implementation status

APPROVED AND EXECUTED 2026-09-14. Owner decisions (§9) received: base confirmed,
all-recommended-options approved (time-compression mapping, generic taxonomy +
schematic trio, auto-grant with Novice, power deferred, R8 OUT, Artisan-Novice-only,
minimal-delta equip), implementation authorised — treated as covering the Phase 3
review-gate precondition for this line (owner may still call the gate separately).
Results on `testbed/swg-phase3-combat/`: `go build` + `go vet` clean; fresh-DB full
suite testclient 9/9 + phase1test 7/7 + phase2test 11/11 + phase3test 6/6 +
phase4test 9/9, ALL PASS (real output; R1–R7 built, R8 excluded per approval).
Observed: sidearm damage ~420–450 from OQ-800 quality (GDD formula literal);
funded harvester 32 units/tick vs unfunded shutdown; lifecycle 18→18 zero-overlap
rotation; corpse yields feeding plate craft; multi-stack slot consumption adopted
where the GDD is silent (recorded in code + HYGIENE_NOTE.md). Bugs found by running:
attacker-excluded broadcasts (Phase 3 carryover pattern reapplied), sqlite timestamp
parse shapes, corpse/hopper pseudo-stack assignment, single-stack quota impossibility.
HYGIENE_NOTE.md extended (§Phase 4 additions); species-ID + Godot known-remainings
unchanged. Next: human review gate before any Phase-5-equivalent work.

*End of proposal v0.1. To enact: owner approves (or amends) §§3–6/9–10; approval and date
recorded before any code is written.*
