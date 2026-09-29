# SWG Pre-CU Phase 3 — Combat Core: Proposal Plan v0.1

**Status:** PROPOSED plan. No authority. No design decision. Human approval required
before any implementation, and separately before any numeric value is treated as tuned.
**Author / Assessor:** Muse Spark (`opencode/muse-spark-1.3-contributor-free`) via OpenCode —
research and planning only; all authority decisions remain human.
**Date drafted:** 2026-09-14.
**Controlling references:** `AGENTS.md`; `sources/swg-pre-cu/README.md`;
`proposals/SWG_Code_Due_Diligence_Checklist_v0.1.md`;
`investigations/SWG_Code_Due_Diligence_Evidence_Log_2026-09-14.md` (§§1–7);
`governance/TCIndustries_Testbed_Reuse_Human_Ruling_HD-TST-01_2026-09-14.md`;
predecessor GDD `sources/swg-pre-cu/SWG_PreCU_GDD.md` (§§9, 17, 30 — HISTORICAL, never TCIndustries canon).

---

## 1. Objective

Implement **Phase 3 — Combat Core** on the Perplexity server line (`code-phase-2/` state:
auth, movement/spatial, skills/professions), satisfying the predecessor-GDD §30 exit
criteria verbatim: *"A player can fight and defeat a spawned creature using a skill-gated
ability, take damage against the correct HAM pool, and be incapacitated/revived-or-cloned
correctly (Section 9.6)."*

Entry criteria (§30: Phases 1+2 complete — needs movement + Marksman/Brawler trees) are
MET on the Perplexity line (verified §§3.1: phase1test 7/7, phase2test 11/11 on the
phase-2 tree).

## 2. Research basis (what exists and what it says)

### 2.1 Predecessor-GDD combat map (HISTORICAL input, not authority)

| GDD section | Specifies | Phase-3 relevance |
|---|---|---|
| 9.1 philosophy | Tactical, HAM-driven, deterministic; death has consequences | Guiding shape only |
| 9.2.1 HAM pools | Health/Action/Mind + 6 attributes; ANY-pool-zero incap; 5-min timer; wounds (flat max reduction, Medic-healed); BF (% max reduction, Entertainer-healed) | Core — implement fully |
| 9.2.2 resolution | 8-step attack flow; hit-chance and damage formulas (quoted in §2.3) | Core — implement flow; formula values per §6 |
| 9.2.3 postures/stances | 4 postures (accuracy/defense/movespeed mods) + 4 stances (dmg/def mods); change locks | Core data + basic modifiers; animation locks out |
| 9.2.4 abilities | 6 categories + example abilities w/ HAM costs, cooldowns, effects | DEFERRED — basic attacks only (no crafted-weapon system until Phase 4 to hang them on) |
| 9.2.5 weapons/damage | Melee/ranged/heavy categories; kinetic/energy/elemental/stun; armor resistances | DEFERRED except unarmed baseline (no items exist yet) |
| 9.2.6 states/DoTs | Stun/KD/snare/blind/dizzy; 4 DoTs + cures | DEFERRED (needs Combat Medic, Phase 9 profession) |
| 9.3 group combat | XP bonus, threat/aggro math, healer/DPS roles | DEFERRED (needs groups + Medic) |
| 9.4 PvP | Overt/covert flagging, factional points, base warfare | EXCLUDED — Phase 8 scope |
| 9.5 creatures | CL 1–90+, con colors, 4 AI classes, lair HP/spawns 3–10/patrol/aggro 5–80 m | Core — MVP subset (below) |
| 9.6 death/respawn | Incap 5 min → revive at 10% pools + wounds/BF; death → clone bind, −5% item condition (no items: noted, skipped), +1–3% BF, +50–200 wounds, credit drop (PvP-scaled) | Core — implement minus item-condition/credit-drop (nothing to apply them to) |
| 9.7 balance | Weapon damage bands; HAM scaling; TTK targets (20–40 s solo white-con etc.) | TARGETS ONLY — numbers per §6, no tuning in proposal |
| 17 creature system | Templates, taxonomy, 10-min harvestable corpse [ASSUMPTION], lairs | Templates + lairs core; corpse despawn timer core; harvesting Phase 4 |

### 2.2 Existing implementations (evidence, not authority)

- **Group C `code-phase0-3-claudecode/`** implements this exact phase against GDD v2 and
  passes 6/6 live tests (evidence log §3.1): pure `combat` package (resolution + HAM state
  machine with explicitly-flagged [ASSUMPTION]s on wound/BF stacking and revive midpoints),
  `creatures` package (templates, instances, lairs, `DistanceSq`), handlers (skill gate on
  `marksman_novice`/`brawler_novice`, 8 m attack range, retaliation AI, 1 s tick loop),
  protocol messages (`combat_action`, `set_posture/stance`, `clone`, `revive`,
  `combat_result`, `ham_update`, `creature_death`, `incapacitated`, `cloned`), persistence
  (`combat_db.go`), known simplifications (ungated revive, static creatures, instant
  despawn, flat CL XP). **Recommended reference implementation**: adapt its PATTERNS, do not
  copy its SWG content (names, Tatooine, tuning) without the §4 hygiene pass.
- **Track A/B `combat_resolution.py`** (12 lines): hit/damage arithmetic reference; consistent
  with GDD 9.2.2 shape. Consult, don't depend on.
- **Perplexity chat export**: 12 combat / 14 HAM / 4 Phase-3 mentions in 377 lines —
  minor Group A provenance color only; no design content relied upon.

## 3. Base and working location (RECOMMENDED — human confirms)

1. Fork `sources/swg-pre-cu/code-phase-2/` to a human-named new location; the SWG original
   stays unmodified (same fork discipline as HD-TST-01 preconditions).
2. All NEW content uses generic IDs from the start (no SW names, per HD-TST-01 direction);
   pre-existing base strings touched by the work (e.g. the `tatooine` spawn-planet literal
   in `character.go`, trainer-seed locale) are renamed to a generic zone ID within the fork,
   recorded file-by-file in the hygiene note (evidence-log §4 pattern).
3. Alternative (NOT recommended): extend the Group C tree — rejected because its Phase 3 is
   already complete there; re-doing it adds no testbed value and its SWG content would need
   stripping first anyway.

## 4. Scope IN (work packages, each verify-by-running before the next)

- **C1 — Pure combat package** (`internal/combat/`): hit-chance/damage resolution per 9.2.2
  formula shape; posture/stance modifier tables per 9.2.3; unarmed baseline weapon
  (stats per §6). No DB, no handlers — unit-verifiable logic.
- **C2 — HAM state machine** (`ham_state`-equivalent): 3 pools + 6 attributes from species
  base; effective-max under wounds/BF (stacking rule: ADOPT Group C's flagged min() convention
  OR human picks — see §6); ANY-zero incap; 5-min timer; revive at 10% + wounds/BF;
  clone at bound facility + wounds/BF; out-of-combat regen tick rate per §6.
- **C3 — Creatures, lairs, AI** (`creatures`-equivalent): template/instance/lair types;
  2+ generic tutorial-tier templates (stats per §6); lair-seeded instances persisted in
  `lairs`/`creature_instances`; MVP retaliation AI (proximity aggro, in-range attack,
  out-of-range disengage — Group C pattern); 1 s world tick (AI + incap timeouts + regen);
  corpse-despawn timer (harvest interaction itself Phase 4).
- **C4 — Protocol**: `combat_action`, `set_posture`, `set_stance`, `revive`, `clone`
  inbound; `combat_result`, `ham_update`, `creature_death`, `incapacitated`, `cloned`
  outbound; creature-bearing `entity_spawn` (`entity_type`/`template_id`, Group C pattern).
- **C5 — Persistence**: combat state columns, clone binding, lair/instance tables, idempotent
  seeding; ACID training-style transactions where state changes span tables.
- **C6 — Handlers**: skill gate (Novice of a combat profession); target validation; range
  check; retaliation wiring; revive/clone flows; combat XP award (flat CL-tier, Group C
  pattern — full 7.2.1 formula needs effective-level, not yet built).
- **C7 — `phase3test`** (6 tests mirroring the verified Group C set): skill-gate rejection;
  unknown-target rejection; kill a spawned creature (real resolution); posture reflected;
  real-time retaliation → incap (accept ~1–2 min runtime, no mocks); clone with wounds/BF.
  Plus re-run of ALL Phase 0–2 tests per the verify-don't-claim rule.
- **C8 — Godot**: DEFERRED to display-only follow-up. Rationale: Group C's own Godot tree is
  byte-identical to Phase 2's (evidence §2.3) — i.e. combat was verified server-side via
  tests, and the existing `World.gd` HAM display already surfaces pool state. No new client
  logic in this plan.

## 5. Scope OUT (explicitly deferred, with homes)

PvP/flagging/factions (§9.4 → Phase 8); special attacks, DoTs, conditions (§9.2.4/9.2.6 →
later phases, needs Medic/weapons); crafted weapons/armor/damage types beyond unarmed
(Phase 4); corpse harvesting (Phase 4, §17.2.2); Combat-Medic-gated revive (later
profession — interim: any-nearby-player revive, Group C precedent, flagged); group
XP/threat/roles (§9.3 → needs groups); credit-drop/item-decay on death (nothing to apply
to yet); creature pursuit/patrol movement (MVP AI is stationary retaliation).

## 6. Numeric authorization requests (values NOT set by this proposal)

Governance boundary: no numeric tuning without explicit human authorisation. GDD-GIVEN
numbers need no approval (implement as written): incap 5-min timer; revive at 10% pools;
lair spawns 3–10; corpse window 10 min ([ASSUMPTION]-tagged in GDD — implement, flagged);
death +1–3% BF / +50–200 wounds; balance BANDS/TTK in 9.7.1 as acceptance targets, not inputs.
GDD-SILENT numbers (each needs human approval OR interim [ASSUMPTION] sign-off): unarmed
accuracy/damage; creature HP/damage/accuracy/aggro radii; attack/revive ranges; regen
%/tick; combat XP per CL tier; wound/BF stacking order; revive wound/BF gain (Group C used
midpoints 125 / 2%). PROPOSED interim convention (approval requested): adopt Group C's
run-verified numbers as PROVISIONAL placeholders marked NON-CANONICAL in code comments —
this avoids inventing new tuning while keeping the build testable. No placeholder survives
contact with real TCIndustries design (all combat numerics stay TBD/PROTOTYPE).

## 7. Acceptance criteria (verify, don't claim)

`go build ./...` + `go vet ./...` clean; fresh DB; `testclient` + `phase1test` + `phase2test`
+ new `phase3test` all pass against the same running server with real pasted output;
hygiene note with file refs for every renamed/new string; commit per verified package
(C1→C7). Godot import remains inconclusive environment-wide (evidence §2.3) — server-side
verification governs.

## 8. Risks and dependencies

- Numeric gate (§6) is the critical path: without either approvals or provisional-sign-off,
  C1–C3 cannot be built honestly (every formula needs constants).
- Base-location + generic-ID decisions (§3) precede all code.
- Human review gate before any Phase-4-equivalent follow-up (HD-TST-01 pattern).
- Scope-creep magnets: DoTs, pursuit AI, PvP — all §5-excluded; enforce at review.

## 9. Decisions requested of the human (the plan asks; it does not decide)

1. Base/location + generic-ID approach (§3) — confirm or amend.
2. Numeric path (§6) — approve listed values, or approve the Group-C-provisional convention,
   or supply replacements.
3. Scope (§§4–5) — confirm inclusions/exclusions (esp. server-only verification, C8 deferral).
4. Authorisation to implement on the approved base once 1–3 are settled (see §10).

## 10. Implementation status

APPROVED AND EXECUTED 2026-09-14. Owner decisions (§9) received and implemented:
(1) base/location + generic IDs confirmed — fork at `testbed/swg-phase3-combat/`
(byte-copy of `sources/swg-pre-cu/code-phase-2/`, original unmodified);
(2) numerics — fandom Weapon_Accuracy values adopted where recoverable (posture mods,
0.5%/point, 66 baseline; master equation image unrecoverable, documented analysis
implemented and recorded), HAM uniform 1000s all species, Group-C provisional
convention for the remainder (all flagged NON-CANONICAL in code + HYGIENE_NOTE.md);
(3) scope confirmed incl. C8 Godot deferral (client untouched);
(4) implementation authorised — DONE this session.
Results: `go build` + `go vet` clean; fresh-DB full suite
testclient 9/9 + phase1test 7/7 + phase2test 11/11 + phase3test 6/6, ALL PASS
(real output; Test 5 genuine tick-driven incapacitation). De-SWG renames recorded in
`testbed/swg-phase3-combat/HYGIENE_NOTE.md`; known-remaining species IDs + Godot
strings flagged there for the review gate. Human review gate (per §8/HD-TST-01 pattern)
is now the next step before any Phase-4-equivalent work.

*End of proposal v0.1. To enact: owner approves (or amends) §§3/6/9–10; approval and date
recorded in `governance/TCIndustries_Human_Rulings_Register.md` before any code is written.*
