# HYGIENE NOTE — testbed/swg-phase3-combat (Phase 3 fork of code-phase-2)

**Status:** PROTOTYPE / EVIDENCE. No design authority. No TCIndustries canon.
**Base:** byte-copy of `sources/swg-pre-cu/code-phase-2/` (Perplexity Phase-2 line),
taken 2026-09-14; the SWG original is unmodified (verified: `git status` shows no
tracked-file changes under `sources/`).
**Author / Assessor:** Muse Spark (`opencode/muse-spark-1.3-contributor-free`) via OpenCode,
implementing owner-approved `proposals/SWG_Phase3_Combat_Core_Proposal_v0.1.md` (§10
decisions recorded 2026-09-14).

## Removed / replaced (file refs)

| # | Before (SWG-derived) | After (generic) | Files |
|---|---|---|---|
| 1 | Planet `tatooine` (spawn, schema defaults, seeds, assertions) | `zone-0001` | `internal/models/character.go` (DefaultPlanet), `internal/database/db.go` (column default), `internal/database/skills_db.go` (trainer seeds), `cmd/testclient/main.go` (planet assertion) |
| 2 | Spawn `(3528.0, 5.0, -4804.0)` ("near Mos Eisley") | `(20.0, 5.0, 0.0)` (owner C3: base location 10,0 area) | `character.go`, `db.go`, `cmd/phase1test/main.go` (all hardcoded offsets rebased: 3529→21, 3740→232, teleport 4529→1021, Z −4804→0) |
| 3 | City `Mos Eisley`, trainer IDs `*_mos` | City `Zone 0001`, IDs `*_z1` | `skills_db.go`, `cmd/phase2test/main.go` (4 trainer-ID refs), `swg-server/README.md` (2 doc mentions) |
| 4 | Species HAM differentiation (9-species modifier table, GDD 6.2.2 values) | Uniform 1000s, all modifiers zeroed (owner C2) | `internal/species/species.go` (baselines + table), `cmd/testclient/main.go` (expectedHAM table) |
| 5 | (new content) SWG creature/place names | Generic serials: templates `0001`/`0002` ("Creature 0001/0002"), lairs at (10,0)/(11,0) per owner C3 | `internal/creatures/creatures.go`, `internal/database/combat_db.go` (seeds) |

## Kept (non-Lucasfilm, recorded)

- Trainer personal names (Hilbi Bodé, Verloc Vryce, Trelos Voken, Kae'la Tiras,
  Dr. Kelm Uorin, Fhara Vex): invented, no SW canon match found; retained with new zone.
- Profession/skill-box IDs and names (marksman, brawler, artisan…): generic English
  profession vocabulary; retained (mechanics TBD/PROTOTYPE regardless).
- HAM/posture/stance/crate-free vocabulary: generic mechanics terms; retained.
- Fandom-derived accuracy model (posture mods, 0.5% scaling, 66 baseline): community
  research content (CC-BY-SA source, facts/values only, no creative expression reproduced).

## KNOWN-REMAINING SW-derived strings (not yet stripped — review gate to direct)

1. **Species IDs/display names** (`bothan`, `rodian`, `trandoshan`, `twilek`, `wookiee`,
   `zabrak`, `mon_calamari`, `sullustan`): `internal/species/species.go`,
   `cmd/testclient/main.go`, `cmd/phase1test/main.go`, `swg-server/README.md`,
   `GET /api/species`. Kept because they are structural keys across old tests + API;
   full rename deferred to the review gate (flagged here, not cleared).
2. **Godot client** (`swg-godot/`, untouched per proposal C8 deferral): `project.godot`
   description, `Login.tscn` strings, script comments retain SW references.
3. **Fork README** retains predecessor-GDD section references (6.2.2, 7.3, 9.x) as
   traceability pointers to HISTORICAL material, not as authority claims.

## Numbers provenance (all NON-CANONICAL placeholders unless GDD-given)

- GDD-given (no approval needed): incap 5-min timer; revive at 10% pools; lair 3–10
  population shape (MaxPopulation 3); corpse window 10 min ([ASSUMPTION]-tagged in GDD — implement flagged);
  death +50–200 wounds / +1–3% BF ranges (midpoints 125 / 2% used, Group-C convention).
- Owner-directed: fandom posture mods + 0.5%/66% accuracy model (C1; master equation
  image unrecoverable — documented analysis implemented, recorded in code);
  uniform HAM 1000s (C2); lair positions/serials (C3).
- Group-C provisionals (owner-approved §6 convention): unarmed acc 30 / dmg 5–15+2;
  creature stats (60HP 3–8/20acc/10m; 140HP 6–15/25acc/8m); 8 m attack / 5 m revive
  ranges; wound/BF min() stacking; 2%-per-5th-tick regen (no source formula exists —
  independently developed placeholder); flat XP 50×CLMax.

## Phase 4 additions (proposal SWG_Phase4_Resources_Crafting_Proposal_v0.1, approved 2026-09-14)

New strings are generic throughout: resource taxonomy
(`ferric_metal`, `conductive_alloy`, `structural_polymer`, `cultured_organic`,
`industrial_chemical`, `fibrous_flora`, `filtered_water`), schematics
(`basic_sidearm`, `basic_plate`, `harvester_deed`), stat lanes, harvester kinds —
no SW names introduced. Fandom-derived values are facts/numbers only.
GDD-given numbers implemented: gauss(500,150)/OQ stats, rarity bands, 10 s survey
cooldown, 1–3 sample units, ±500 m basic precision, harvester rates/hoppers/fees,
100 m exclusion, experimentation table + bands + ≤5 pts/attempt + 3–5 rounds,
base/final formulas, 10-min corpse window, per-corpse quality roll shape.
Owner-approved provisionals: TESTBED_FAST_CYCLE compression (lifespan 120–180 s,
harvest tick 20 s; defaults stay GDD-given), metre-unit radii, fixed 500 m survey
radius + exact-center waypoints (triangulation deferred), 15 m sample/empty ranges,
5 m corpse range, basic-tool yield curve, +5 deplete/sample, scouting/crafting XP
awards, schematic trio + complexities + XP500, skill-threshold lerp, +3/pt experiment
shift, 0.5/0.3/0.4/0.2 quality mults, corpse yields/types, fee ceil discretization,
equip mapping (damage-range replace, acc bonus, mitigation cap 50), guarantee test
spawns (`testspawn-*`, fast mode only). Multi-stack slot consumption (pool-wide,
weighted-mean quality) is a deliberate mechanic where the GDD is silent.
KNOWN-REMAINING from Phase 3 unchanged (species IDs, Godot client); factory (R8)
NOT built (default OUT, no code).

## Phase 6 additions (proposal SWG_Phase6_Social_Support_Proposal_v0.1, approved 2026-09-14)

New strings generic throughout: service action/message vocabulary, buff pools,
performance kinds (music/dance), stim schematic (`stim_pack`). No SW names.
Fandom-derived values: none (no fandom source used this phase).
GDD-given numbers implemented: buff band +500–3000 and 30–60 min durations
([ASSUMPTION]-tagged, band floor taken); revive-at-10% + wound/BF penalties
(pre-existing); 10-min corpse window (reused); AFK-watching preserved by omission
(GDD 23.2.4 — recorded, not an oversight).
Owner-approved provisionals: wound heal 25+25×tier with 60 s per-pair cooldown
(cast choreography deferred); buff 500+250×tier; BF heal (1+tiers)/5 s tick;
performer Mind drain 10/s with 2 m stationary rule; watch radius 20 m;
service XP (heal/buff 50, revive 100, tip 25, stim-use 25 medical);
stim schematic costs/props and potency/charges mapping; revive gated on
`medic_novice` (closes the Phase 3 any-player interim; Combat Medic elite gate
deferred); busking-anywhere at reduced-rate semantics deferred to venue-bonus
work (anywhere performs at full test rate — flagged simplification);
sessions/watch state in-memory (restart ends performances — flagged limit);
per-type guarantee spawns extended to all 7 types in fast-cycle mode
(test scaffolding; random spawns still drive lifecycle).
CORRECTION LOG: shared *rand.Rand raced across tick/request goroutines and killed
the resource tick loop mid-run (starved later surveys — diagnosed via quota
diagnostics); replaced with per-call fresh instances everywhere, shared field
removed. TickSpawns same-type relocation preserves seed-time type gaps, so
seeding now guarantees ≥1 spawn per type (GDD 11.2.3 shape).
Server bugs fixed by running: missing revive-XP award; corpse-expiry timestamp
parse shapes (Phase 4 carryover pattern reapplied).

## Phase 5 additions (proposal SWG_Phase5_Economy_Infrastructure_Proposal_v0.1, approved 2026-09-14)

New strings generic throughout: structure/vendor/ledger/bazaar vocabulary, deed
schematics (`structure_deed`, `vendor_deed`), tier/kind/status enums. No SW names.
GDD-given numbers implemented: kill-drop band, training-cost bands (left untouched,
retro-tagged), rarity/fee shapes, 50-slot cap [ASSUMPTION], 14-day grace + 180-day
reclaim [ASSUMPTION]s, 30-day telemetry window (literal semantics), 100 m harvester
rule (reused). Owner-approved provisionals: deed costs (house and vendor deeds
both 15M+8P as built — see correction log), house upkeep 200/500/1000, vendor upkeep
300+50/listing, 20 m no-build radius, 15 m purchase/collect ranges, 30 m terminal
rule, 10×CLMax drops (capped 200), fee ceil discretization, equip mapping, seeded
spawn terminal.
CORRECTION LOG (verify-don't-claim): vendor-deed costs as built are 15M+8P
(matching structure_deed, not the 10M+6P sketched during planning) — the built
values govern; tests consume accordingly. A mid-build reconstructive edit
corrupted harvester-deed quotas (8/4→10/6); caught by the test suite, reverted to
the verified 8/4, full suite re-greened. Deed schematics list the generic trio
extension (sidearm/plate/deed + structure/vendor deeds = 5 in registry).
Flake fixes (test-only, no mechanic change): phase3test posture-before-combat
reorder + fast kill-poll (retaliation race was flaking Test 4); phase5test
terminal-proximity sync (DB persist cadence) and incap-aware kill loops.

## Phase 7 additions (proposal SWG_Phase7_Civic_Systems_Proposal_v0.1, approved 2026-09-15)

New strings generic throughout: city/guild/group/mentorship/mail/waypoint/friend
vocabulary, deed schematic (`city_hall_deed`, 15M+8P like structure/vendor deeds),
rank/status/role/channel enums, 12-word standard emote library. No SW names.
Fandom-derived values: none (no fandom source used this phase).
GDD-given numbers implemented: founding ≥10 structures [HISTORICAL]; rank table
20+/40+/75+ with treasury conditions [AUTHORITATIVE-ASSUMPTION]; weekly elections +
7-day terms + 14-day inactivity trigger + earliest-candidacy tiebreak
([ASSUMPTION]-tagged, implemented flagged); 50-msg mail cap [ASSUMPTION]; group cap
20 [ASSUMPTION]; guild ≥3 founders [ASSUMPTION] + registrar sink shape;
30-day guild dissolution countdown; 2-min group succession; 7-day mentorship bond;
no-reputation rule (18.2.5, by omission — aligns with deferred provenance).
Owner-approved provisionals: city upkeep 500/1000/2000 per week (house-scale
anchored) with 20% mayor-status discount (owner decision 3, Politician-substitute);
4× treasury reserve for rank-ups; city radius 100 m; tax caps (flat 500/wk,
vendor 20%); dues cap 20%; registrar 1000; forming grace 7 d; election cooldown 7 d;
mail body 2000/subject 120; waypoint cap 50; mentorship newcomer threshold 5 boxes +
50 mentoring XP; TESTBED_FAST_CYCLE mapping (thresholds 4/6/8, elections 2 min,
term 60 min, grace 3 min, inactivity 45 min, cooldown 30 s — defaults GDD-given).
Deliberate mechanics where the GDD is silent: hall counts toward the founding total;
founder starts as mayor ("Mayor-Founder"); early election close on unanimous vote
(flagged acceleration — period expiry still closes); zoning as a binary mayor gate
(no per-request queue); banned-list enforced as purchase refusal (no interior
geometry for entry); one-guild/one-group membership; mail ownership sentinel
(`mail:<id>` — send-time transfer makes the 18.5 condemned-house edge impossible
by construction); waypoint share as direct copy + system note.
OUT as approved: Rank 4/specialization/full Politician, militia tools, faction
leaning/bases, travel movement + clone binding (flags only), stalls/escrow,
interiors/cells, built guild halls (nullable FK + ownership rule only), salaries,
holoemotes, reputation scores.
CORRECTION LOG (verify-don't-claim): UpdateCity bound `grace` into `status` and
`Status` into `grace_ends` — corrupted every city row (status filter broke the
vendor-tax split; election-close ticks failed loud); caught by the Test 6 ledger
assertion, fixed, suite re-greened. `structure_banned` table name vs generated
`structure_banneds` — caught by Test 10 perm-add. Test-side placedIDs staleness
(inline zoning placement never recorded) — caught by Test 10 fund 403. Per-character
fixed structure IDs replaced with unique IDs (one owner, many structures —
recorded behavior change, Phase 5 suite re-greened). `SetMaxOpenConns(1)` on the
SQLite pool after rapid mail sends hit SQLITE_BUSY against tick writes
(plumbing hardening, no mechanic change). Test-side `BuyerChar`→`CityBuyer` rename
(phase5 name collision on shared DB); main-city test geography shifted +100 m east
(phase5 leftovers at 45/70, 20 m no-build rule).
CODE-ONLY (flagged, durations untestable live): mayor-inactivity emergency election,
term-expiry turnover, 30-day guild succession countdown, 180-day reclaim reuse.
KNOWN PRE-EXISTING ISSUE (not introduced, not fixed — out of scope): no lair-respawn
mechanic exists, so shared-DB Phase 6 depends on earlier suites leaving survivors —
one full-suite run failed Test 1 with a single template-0001 survivor that could not
outpace regen (DB forensics: 5/6 instances dead); rerun passed. Respawn design
belongs to combat-system ownership (human decision), not this phase.
KNOWN-REMAINING from Phases 3–6 unchanged (species IDs, Godot client).
