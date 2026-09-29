# SWG Phase 10 — Balance, Telemetry & Polish — Proposal Plan v0.2

**Status:** PROPOSED plan. Candidate material only — not canonical, not TCIndustries canon.
No design decision. Human approval required before any implementation, and separately
before any numeric value is treated as tuned.
**Author / Assessor:** Buffy — Codebuff agent (`codebuff/freebuff`) via Freebuff, **Proposal v0.2**, 2026-09-17 —
research, verification, and planning only; all authority decisions remain human.
**Relationship to prior version:** revises and extends
`SWG_Phase10_Balance_Telemetry_Polish_Proposal_v0.1.md` (Muse Spark, 2026-09-16, PROPOSED,
decision-pending — retained unaltered as provenance). Every load-bearing v0.1 claim was
re-verified first-hand against the SWG Pre-CU GDD, the Master GDD v1.1.1, the governance
rulings, and the testbed source (evidence noted inline); corrections and new findings from
the completed Phase 9 verification pass are integrated. No status changed anywhere.
**Base:** `testbed/swg-phase3-combat/` (patterns-only fork per HD-TST-01; all fork behaviour
PROTOTYPE/EVIDENCE).
**Authority basis:** predecessor GDD `sources/swg-pre-cu/SWG_PreCU_GDD.md` §§ 1.3, 7.2.3,
9.7, 11.2.1, 12.2–12.6, 15.2.3, 22.5, 29.6, 30 (Phase 10), 31, 33, 34 (HISTORICAL — never
TCIndustries canon per AGENTS.md / HD-TST-01); Master GDD v1.1.1 §§ 25–27 (BAL-001 LOCKED;
BAL-002, SAFE-003, TEST-001, TEST-002 PROPOSED; SAFE-002 PROPOSED/TBD), PROT-001 LOCKED;
`governance/TCIndustries_Testbed_Reuse_Human_Ruling_HD-TST-01_2026-09-14.md`;
HD-EIC-01–09; HD-GDD-01; `proposals/SWG_Phase9_Elite_Missions_Proposal_v0.1.md` §10
decision 6c (service-XP scaling residual routed here — verified verbatim).

---

## 1. Status discipline (read first)

1. This proposal invents no TCIndustries design and promotes nothing. Every predecessor-GDD
   citation below is HISTORICAL (SWG Pre-CU predecessor) unless explicitly tagged with a
   Master GDD v1.1.1 status.
2. The single most load-bearing constraint: **BAL-001 is LOCKED** ("do not invent skill
   caps, spawn rates, prices, decay curves, combat coefficients, or other numeric balance
   values until the structural design is stable") and **HD-TST-01 explicitly forbids numeric
   tuning** in the fork. Therefore "Balance" in this phase means **measure, compare against
   HISTORICAL targets, and report gaps** — never retune numbers to hit bands. Any rate or
   value change (including the Phase 9 service-XP scaling residual) requires a separate
   explicit numeric authorization; §6 lists each such number as a decision, not a default.
3. No EIC shaping (HD-EIC-05/07 in force): interdependence KPIs below are **measurement
   only**. If the ≥90% interdependence proxy fails, the output is a telemetry report, not a
   new residual or mechanism.
4. Provenance / formal reputation stays deferred: the ≥95% crafted-share proxy must use the
   existing `crafted_items.schematic_id` query path (verified present, §2.5), never a new
   provenance system. SAFE-002 (PROPOSED/TBD) is untouched: B5 surfaces signal for human
   review and performs no account actions.

## 2. Research findings (all verified 2026-09-17 unless noted)

### 2.1 Phase 10 definition (predecessor GDD §30.2 — read first-hand, HISTORICAL)

- **Implements:** §33 (balance target validation), §34 (success metrics instrumentation),
  §29.6 (monitoring).
- **Entry criteria:** Phase 9 complete. **Status as of 2026-09-17:** the Phase 9 E2E suite
  (`run_phase9.sh`, `cmd/phase9test`) passes all 8 tests — `=== ALL PHASE 9 TESTS PASSED
  ===`, `PHASE9_EXIT:0` — verified on this Windows host today (evidence: `/tmp/p9run.log`
  output pasted in the Phase 9 verification report; working-tree commits still pending
  owner action). Implementation-side entry criterion is met; this file remains
  planning-only until §9 authorization.
- **Exit criteria:** "§1.3's five success-criteria tests can be run and pass (or produce
  actionable telemetry showing where they don't)." The disjunction is load-bearing:
  instrumentation + an honest failure signal satisfies exit.

### 2.2 §34 Success Metrics → §1.3 tests (HISTORICAL targets, testbed-observable form)

| §1.3 Test | §34 KPI | Target | Testbed data source |
|---|---|---|---|
| Economic Interdependence | % active combat players receiving a crafted item / entertainer heal / medic heal in last 7 d | ≥90% | new append-only `interdependence_events` log written at existing service/purchase hooks (B1) |
| Horizontal Progression | 6-month vs 2-week character effectiveness ratio, equivalent gear | 2–3x | NOT-MEASURABLE live (no 6-month horizon) — bounded proxy: novice-boxed vs elite-boxed effectiveness ratio at sidearm scale (§6.3) |
| Player-Driven Economy | % items in inventories/vendors with `crafted_from_schematic_id != null` | ≥95% | `crafted_items.schematic_id` — verified present (`resource_db.go:580` insert, `civic_social_db.go` ownership paths) |
| Sandbox Validation | player-organized events / guilds / schemes not authored by design | qualitative | NOT automatable — observation-log template (§6.4) |
| Historical Accuracy | core loops vs documented Pre-CU accounts | human checklist | HUMAN pass only — out of testbed scope (§6.4) |

### 2.3 §33 Balance Targets (HISTORICAL bands — validate-against, never tune-to)

- **Combat (§9.7):** damage 50–150 start / 150–400 mid / 300–700 end-game / 1000–1500 max
  crit; HAM 3000–6000 unbuffed / 6000–12000 buffed; TTK 20–40 s solo white-con, 60–120 s
  group-vs-elite, 30–60 s PvP 1v1; cooldown/HAM-cost bands per ability tier (specials
  unbuilt in fork — ability-system residual, Phase 9 §5/§10). Priority [MVP] — gates the
  Horizontal Progression Test (§33.7).
- **Progression time (§33.3):** basic mastery 40–60 h; elite 80–120 h; multi-prereq elite
  150–250 h; full 250-point cap 200–300 h cumulative. Unmeasurable live — projection method
  (§6.3): observed XP/hour per pool from phase-test timing logs extrapolated to the
  §33.3 milestones; NOT-MEASURABLE where combat XP derives from fast-cycle-mapped creature
  density and projection would be dishonest.
- **Economic (§33.4, §12.4):** target inflation ~5–10% annualized on a rolling 30-day
  window; novice-crafter ≈ novice-combat income; master 2–3× novice; maintenance 5–15% of
  weekly income; stall commission 4–6% (stalls dormant in fork — recorded, no value).
- **Housing/civic/faction (§33.5–33.6):** cost bands + faction thresholds
  2,500/10,000/30,000/75,000 (§15.2.3). Priority [EXPANSION] — Phase 10 records current
  fork values beside targets; no changes.

### 2.4 §29.6 Monitoring + §29.8 job shape (HISTORICAL)

- Nightly economic snapshot job (faucet/sink aggregation → net inflation) separate from the
  real-time loop; `WorldTickJob` record with idempotent resumption (pending/running/
  completed/failed, §29.8 crash-mid-tick rule).
- Lightweight anomaly job flags wash trading / price outliers / placement bursts **for
  human review, never auto-action** ("surface signal, not auto-ban" — no trust & safety
  team exists). Telemetry events structured to populate the §34 dashboard directly.
- Testbed status (verified): `economic_snapshots` table + `WriteSnapshot`/`LatestSnapshot`
  with literal 30-day faucet/sink window (`internal/database/economy_db.go:371–410`) and a
  per-harvest-tick hook already call it under `TESTBED_FAST_CYCLE=1`
  (`internal/handlers/resources.go:132`). MISSING: scheduled nightly-job *shape*, any
  anomaly detection, Gini, volatility, profession distribution, interdependence-event
  counters, mission completion rates, any dashboard.

### 2.5 Testbed inventory (verified against source 2026-09-17, PROTOTYPE)

- **Ledger:** `internal/economy/economy.go` flow constants `faucet`/`sink` + category tags
  (training_cost, maintenance_fee, creature_drop, starting_credits, mission_reward,
  mission_delivery, mission_sample); `RecordLedger` calls at mission turn-in
  (`missions.go:366–379` — reward faucet + escrow-refund sinks).
- **Snapshots:** as §2.4.
- **Provenance query path:** `crafted_items.schematic_id` (no new system needed).
- **Provisional service-XP rates (flagged `[PROVISIONAL]` in code, verified locations):**
  image-designer session ×2 and smuggler slice flat 50 (`elite_http.go:327, 367, 512`);
  stim-use 25 (`elite_http.go:613`); pet-tick combat +50 (`ranger.go:266`); camp/track 25
  (`ranger.go:101` via `CampXPPerMember`, `ranger.go:176`); medical heal 50 / buff 50 /
  revive 100 / tip 25 (`internal/services/services.go:78–81`); stim-use 25
  (`services.go:422`). GDD §7.2.3 (HISTORICAL) specifies per-point wounds XP, per-tick
  BF XP scaled with BF healed, per-application buff XP, revive bonus, and diminishing
  returns on repeat healing — the fork implements none of the scaling; this is the Phase 9
  decision 6c residual (B4).
- **Fast-cycle compression convention:** `TESTBED_FAST_CYCLE=1` maps GDD cadences
  (harvest 1 h → 20 s, spawn lives 120–180 s, camp/harvest interval, lair relocation
  window 317 s observed in Phase 9) — wired in `internal/resources/resources.go:107–114`,
  `internal/economy/economy.go:75`, `internal/civic/civic.go:106`,
  `internal/faction/faction.go:102`, `internal/handlers/lairs.go:35`,
  `internal/handlers/ranger.go:32`.
- **Phase 9 verification findings (new since v0.1, all fixed in the working tree):**
  `GetCharacterByID` returned species-baseline HAM instead of live combat state (duplicate
  scan target + `ComputeHAM` overwrite; fixed — REST `/api/characters/{id}` now reflects
  live damage; relevant to B1 because character-state instrumentation reads this surface);
  mission-ID minting via `time.Now().UnixNano()` collided on Windows clock granularity →
  `newRowID` collision-proof  helper introduced for missions; **45 further `UnixNano()`
  ID-minting sites** across 17 files inventoried (civic-family ×11, phase9_db ×7,
  resource_db ×4, resources/faction/combat_db ×3 each, and 10 further sites across
  9 files) — portability polish candidate (B6), non-numeric, non-mechanical.
- **Polish debt:** `COMBAT-DEBUG` log lines (`combat.go:207, 218`); `.gitignore`
  `phase*test` pattern previously ignored the `cmd/phase1test–9test` source dirs (negated
  in working tree); `run_phase9.sh` requires cgo (Windows recipe now documented in the
  fork README).

### 2.6 Master GDD constraints that shape this phase

- BAL-001 LOCKED + Master GDD "Inherited from prior versions" item 3 ("Numeric values
  remain universally avoided / TBD", verified at line 983; v0.1's "§34 step 7" was an
  imprecise citation) — no fork number becomes a TCIndustries number.
- BAL-002 PROPOSED — the lens for reading all validation output: a DPS-parity gap is not a
  failure; a collapsed profession category is.
- SAFE-003 PROPOSED (economic observability), TEST-001 PROPOSED (design invariants as
  tests), TEST-002 PROPOSED (simulation preference) — joint authority basis for B1–B3
  alongside the HISTORICAL §§33/34/29.6.
- PROT-001 LOCKED — fork behaviour is evidence, never design authority.
- OQ-001–017 remain TBD; Phase 10 answers none of them. It may attach numbers-adjacent
  *observations* (e.g. "fork value X sits outside HISTORICAL band Y") without resolving
  anything.

## 3. Scope IN

- **B1 — Telemetry completion (read-only instrumentation):** nightly-snapshot job shape
  (idempotent `WorldTickJob`-style record per §29.8; fast-cycle-mapped cadence, GDD-given
  defaults); append-only `interdependence_events` log (crafted purchase, entertainer/medic
  heal receipt, resource purchase — per character, counterparty, timestamp) written at
  existing hooks only; profession-distribution query (boxes by combat/crafting/social
  category); mission completion/cancellation rates from the existing missions table;
  wealth-distribution query (Gini — diagnostic only per §34.3, tracked, never
  target-capped); resource price-volatility per rotation cycle (if no sale-price history
  table exists, add an append-only telemetry record at sale time — structure, not value).
  All observation; no levers, no steering.
- **B2 — §1.3 proxy harness (`cmd/phase10test/`):** five proxies runnable against a live
  fresh-DB world, each reporting PASS / FAIL-with-telemetry / NOT-MEASURABLE-with-reason.
  FAIL is an accepted exit outcome provided the telemetry is actionable (§2.1
  disjunction). With test-population cohorts of ~3 characters, percentages are reported
  alongside raw counts (a 2/3 cohort is 67%, not a verdict).
- **B3 — Balance validation report:** scripted comparison of fork-observable values against
  the §33 bands — unarmed 5–15 damage (sidearm scale) vs §9.7 50–150 start band; PvP TTK
  vs 30–60 s; creature-drop 10×CLMax cap 200 vs §12.2.2 10–200 band; mission rewards vs
  §12.2.1 500–10k band; maintenance share vs 5–15%; faction thresholds vs §15.2.3; XP-rate
  projection vs §33.3 horizons. Output is a REPORT (fork value vs HISTORICAL band vs gap)
  committed as evidence. Zero value changes.
- **B4 — Service-XP scaling residual (Phase 9 decision 6c):** telemetry characterization of
  the flat-50/25/100 provisionals (§2.5 list) + a rate-redesign OPTIONS paper (per-point
  wounds, per-tick BF, per-application buff, diminishing-returns hook — §7.2.3 shapes with
  explicit numeric asks). Implementation of any new rate waits on §6 authorization;
  default if declined: provisionals stand, characterized.
- **B5 — Anomaly-lite (monitoring, never enforcement):** wash-trade heuristic (rapid
  A→B→A transfer pairs), price-outlier flag, placement-burst flag over ledger/market
  history. Two output options offered in §6: (i) ranked review lists with no thresholds
  (no numeric authorization needed), or (ii) sigma/velocity thresholds (numeric, needs
  authorization). Either way: human review only, no auto-action (§29.6; SAFE-002 TBD
  untouched).
- **B6 — Polish / stability:** flake-harden the 0–9 suite chain (shared-DB coexistence,
  fast-cycle maps, REST point-in-time reads over WS-soak baselines — pattern proven in
  Phase 9 Test 4); retire or env-gate `COMBAT-DEBUG` logging; portability sweep of the 45
  `UnixNano()` ID sites onto the `newRowID` helper (mechanics unchanged — ID format
  unobservable to gameplay); README phase-table Phase 10 row; hygiene note extended per
  package; verify the `.gitignore` source-dir negation lands with the commits. No mechanic
  changes under this package.
- **B7 — Full-chain regression:** phases 0–9 suites re-run green on fresh DB plus new
  `phase10test`, real output pasted, commit per verified package (verify-don't-claim,
  HD-TST-01 methodology).

## 4. Scope OUT (explicitly deferred, with homes)

- Any numeric retuning (BAL-001 / HD-TST-01 — needs its own human numeric authorization,
  never implied). Includes "helpful" in-pass fixes to any B3-observed gap.
- EIC shaping or new residuals (HD-EIC-07 gate — human structural boundaries still pending).
- Provenance / reputation systems (still deferred); auto-enforcement or trust-and-safety
  actions (no mandate; SAFE-002 stays TBD).
- Player-facing dashboards/charting ([EXPANSION] per §§34.6/33.7 — server-side query
  endpoints + human-readable output only; GUI dashboards are a later pass).
- Structured live playtests with human cohorts (§34.5 — human-run activity; the testbed
  provides only the instrumentation they would read).
- Stalls/escrow follow-up and PvP-combat-XP amendment (separately sequenced per Phase 9
  §10 decisions 6a/6b — not absorbed here).
- Ability system / special attacks (Phase 9 §10 major residual — human-directed future
  phase, not polish).

## 5. Tensions and honest limits (human resolves, proposal does not)

1. **"Balance validation" vs BAL-001:** resolved by construction (§1.2) — validate =
   observe + report. If any B3 gap tempts a tuning fix, that fix is a NEW numeric
   proposal, not part of this phase.
2. **Time-horizon KPIs are unmeasurable live** (6-month ratios, 40–300 h mastery, 30-day
   rolling inflation on a fresh DB, 90-day city survival): method is rate-projection +
   fast-cycle mapping + explicit NOT-MEASURABLE where projection is dishonest. No
   simulated-player fabrication to fake a PASS.
3. **Qualitative tests stay qualitative:** Sandbox Validation = observation-log template
   (event, date, witnesses); Historical Accuracy = human checklist — the harness reports
   their existence/absence, never a fabricated score.
4. **Gini is diagnostic** (§34.3: cartel concentration is historically accurate emergence)
   — B1 computes it; no threshold, no gate, no lever.
5. **Mission faucet dominates Phase 9+ economies:** B3 reports faucet mix honestly (mission
   rewards live via turn-in, verified §2.5) against the §12.4 5–10% band as observation
   only.
6. **Small-cohort statistics:** proxies run over the E2E test cohort (≈3 combat
   characters), so "≥90% of combat players" is a coarse signal; raw counts are reported
   beside every percentage (B2).

## 6. Numeric authorization requests (values NOT set by this proposal)

**GDD-GIVEN (implement/observe as written, no approval needed):** 30-day window + 5–10%
band (observed, never steered); §9.7/§33.2/33.3/33.4/33.5/33.6 bands (compared-against
only); 10 s survey cooldown (reused); nightly snapshot cadence shape (fast-cycle-mapped,
defaults GDD-given); ≥90% / ≥95% / 2–3× KPI targets (measured-against only).

**NEEDS APPROVAL (each: approve, amend, or decline-to-stand-pat):**
1. Service-XP redesign rates (per-point / per-tick / per-application magnitudes per
   service type, §7.2.3 shapes) — B4 implementation only on approval.
2. Anomaly heuristics — choose output option: **(i) ranked lists, no thresholds**
   (recommended; requires no numeric authorization), or **(ii) explicit
   sigma/velocity thresholds** (numbers, authorized item-by-item).
3. Snapshot/aggregation cadence + retention under fast-cycle.
4. Any B3 follow-up value change — explicitly a separate decision, default NO.

**PROPOSED interim convention (approval requested):** Group-C-style NON-CANONICAL
flagging on every non-GDD number (code comments + `HYGIENE_NOTE.md`), fork-local
magnitudes reused where a magnitude must exist.

## 7. Acceptance criteria (verify, don't claim)

`go build ./...` + `go vet ./...` clean (cgo toolchain per fork README recipe on Windows);
fresh DB; phase0–9 suites green PLUS new `phase10test` (five proxies report PASS /
FAIL-with-telemetry / NOT-MEASURABLE; at least the three automatable proxies —
interdependence, crafted-share, combat-TTK spot — produce real numbers with raw counts);
B3 report committed as a file with fork-value vs HISTORICAL-band table; B5 demonstrated
on a seeded wash-trade fixture; **no numeric diff in any mechanic file** (enforced by
`git diff` review — instrumentation-only proof); hygiene note extended; commit per
verified package.

## 8. Sequencing (non-binding)

B1 → B2 → B3 → B5 → B4-options-paper → B6 → B7 — measurement exists before any redesign
talk. B4 implementation (if authorized) becomes a separately authorized micro-package
after this phase exits.

## 9. Decisions requested (no code until answered)

1. Base/location — confirm `testbed/swg-phase3-combat/`.
2. Scope — confirm B1–B7 IN / §4 OUT as written, or amend.
3. Observe-don't-steer rule (§1.2) — confirm, or authorize specific tuning levers now
   (not recommended).
4. Time-horizon method (§5.2) — confirm projection + NOT-MEASURABLE honesty, or direct an
   alternative.
5. Qualitative-test method (§5.3) — confirm log-template + human-checklist approach.
6. Numeric set (§6) — approve/decline each of (1)–(4); if (1) declined, provisionals stand
   characterized; pick B5 option (i) or (ii).
7. Authorization to implement on approval of 1–6.

## 10. Implementation status

**AUTHORIZED 2026-09-17** — owner approved the v0.2 plan with these decision
outcomes: §9 items 1–5 and 7 **APPROVED** as written; §6 (1) **DECLINED as
authorization to change rates** (provisionals stand, characterized per decision 5
outcome: B4 = characterization telemetry + options paper only, no rate redesign);
§6 (2) **option (i) ranked lists, no thresholds**; §6 (3) fast-cycle-mapped
cadence per §8 conventions (GDD-given nightly shape, compression test-config-only);
§6 (4) default NO stands. Constraints binding on execution: zero numeric changes
(BAL-001/HD-TST-01), zero gameplay-behavior changes (instrumentation + hook-side
telemetry writes only, best-effort so failures cannot alter outcomes), REST
additions read-only (`GET /api/telemetry/consolidated`), `go vet` + `go test`
clean, commit per verified package with verify-don't-claim output. Execution
follows §8 sequence; every package's completion is reported with real pasted
evidence in the Phase 10 verification report.

Prior status record: PENDING — no authorization given, no Phase 10 code written. Entry precondition (Phase 9
suite green) is met as of 2026-09-17 (PHASE9_EXIT:0, all 8 tests, Windows host; commits
pending owner action). Next step is an explicit item-7 authorization covering §§3–6 as
amended; until then this file is planning-only. On authorization, execution follows the
verify-don't-claim methodology (`go build` + `go vet` clean, fresh DB, all phase tests
re-run old+new with real pasted output, hygiene note extended per package, commit per
verified package), with the human review gate before any post-Phase-10-equivalent work
(HD-TST-01 pattern).

**EXECUTED 2026-09-17 (same day, after the authorization above).** All packages
implemented and verified on the Windows host; `go vet` clean, `go test ./...` green,
full-chain regression `p1=0 … p9=0` (all nine pre-existing suites green on fresh DBs)
followed by `PHASE10_EXIT:0` (all observable §1.3 proxies PASS; Proxy 4 honestly
NOT-MEASURABLE). Zero numeric changes: the B6 sweep converts only ID *minting* to the
collision-proof `newRowID`/`DB.NewRowID` helper (Windows-clock portability; mechanics
unchanged; RNG seeds untouched), and the only gameplay-adjacent code paths touched are
best-effort telemetry writes at existing hooks. B1 REST landed at `/api/phase10/*`
(telemetry, anomalies, jobs) — a route-prefix refinement of the authorized
`/api/telemetry/consolidated` sketch, same read-only content. B5 is option (i) ranked
lists (no thresholds). B4 is characterization + options paper only, per the §6 (1)
DECLINED-as-rate-change outcome. Committed per package (commits f9852ec, 55ec433,
12ca3db, ce09b35, 458c311, e6c1187, 2730c9a); full evidence and the mixed-provenance
file inventory live in the Phase 10 verification report. Additions outside this file's
original enumeration: `run_phase10.sh`/`run_chain.sh` (B6/B7 harnesses), the B7 flake
fixes (chain port reaping; phase2test stale 6-profession assertion refreshed to the
Phase 9 registry — test-hygiene, not tuning).

*End of proposal v0.2. To enact: owner approves (or amends) §§3–6/9–10; approval and date
recorded before any code is written. Supersedes v0.1 only upon human adoption; until then
v0.1 remains the prior candidate on record.*
