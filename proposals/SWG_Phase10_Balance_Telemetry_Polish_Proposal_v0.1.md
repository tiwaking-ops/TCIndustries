# SWG Phase 10 — Balance, Telemetry & Polish — Proposal v0.1

**Status:** PROPOSED. Candidate material only — not canonical, not TCIndustries canon.
**Author / Assessor:** Muse Spark (`opencode/muse-spark-1.3-contributor-free`) via OpenCode, 2026-09-16.
**Base:** `testbed/swg-phase3-combat/` (patterns-only fork per HD-TST-01; all fork behaviour PROTOTYPE/EVIDENCE).
**Authority basis:** SWG Pre-CU GDD v2 §§ 29.6, 30 (Phase 10), 31, 33, 34 (HISTORICAL — never TCIndustries canon per SWG-001/002);
Master GDD v1.1.1 §§ 25–27 (BAL-001 LOCKED, BAL-002 PROPOSED, SAFE-003 PROPOSED, TEST-001/002 PROPOSED);
HD-TST-01; HD-EIC-01–09; HD-GDD-01; Phase 9 proposal §10 decision 6 (residual placements).

---

## 1. Status discipline (read first)

1. This proposal invents no TCIndustries design and promotes nothing. Every GDD citation below is HISTORICAL
   (SWG Pre-CU predecessor) unless explicitly tagged with a Master GDD v1.1.1 status.
2. The single most load-bearing constraint: **BAL-001 is LOCKED** ("do not invent skill caps, spawn rates,
   prices, decay curves, combat coefficients, or other numeric balance values until the structural design is
   stable") and **HD-TST-01 explicitly forbids numeric tuning** in the fork. Therefore "Balance" in this
   phase means **measure, compare against targets, and report gaps** — never retune numbers to hit bands.
   Any rate/number change (including the Phase 9 service-XP scaling residual) requires a separate explicit
   numeric authorization; §7 lists each such number as a decision, not a default.
3. No EIC shaping (HD-EIC-05/07 in force): interdependence KPIs below are **measurement only**. If the
   ≥90% interdependence proxy fails, the output is a telemetry report, not a new residual or mechanism.
4. Provenance / formal reputation stays deferred: the ≥95% crafted-share proxy must use the existing
   `crafted_items.schematic_id` query path (verified present, §2.5), never a new provenance system.

## 2. Research findings (sources consulted)

### 2.1 SWG GDD Phase 10 definition (§30.2, HISTORICAL)

- **Implements:** §33 (balance target validation), §34 (success metrics instrumentation), §29.6 (monitoring).
- **Entry criteria:** Phase 9 complete. **NOT MET YET** — Phase 9 implementation is in progress (Test 3
  triangulation debugging, September 2026). Phase 10 work must not start until the Phase 9 suite is green
  and committed; this proposal is planning-only until then.
- **Exit criteria:** "§1.3's five success-criteria tests can be run and pass (or produce actionable telemetry
  showing where they don't)." Note the disjunction: instrumentation + honest failure signal satisfies exit.

### 2.2 §34 Success Metrics → §1.3 tests (HISTORICAL targets, testbed-observable form)

| §1.3 Test | §34 KPI | Target | Testbed data source |
|---|---|---|---|
| Economic Interdependence | % active combat players receiving crafted item / entertainer heal / medic heal in last 7 d | ≥90% | §29.6-style event log across combat/heal/purchase paths (TO BUILD — see B2) |
| Horizontal Progression | 6-month vs 2-week character effectiveness ratio, equivalent gear | 2–3x | CANNOT measure live (no 6-month horizon) — proxy required (see §6.3) |
| Player-Driven Economy | % items in inventories/vendors with `crafted_from_schematic_id != null` | ≥95% | `crafted_items.schematic_id` query (EXISTS — §2.5) vs loot/NPC-sourced items |
| Sandbox Validation | player-organized events / guilds / schemes not authored by design | qualitative signal | NOT automatable — observation log method (§6.4) |
| Historical Accuracy | core loops vs documented Pre-CU accounts | design review checklist | HUMAN pass only — out of testbed scope (§6.4) |

### 2.3 §33 Balance Targets (HISTORICAL bands — validate-against, never tune-to)

- **Combat (§9.7):** damage 50–150 start / 150–400 mid / 300–700 end-game / 1000–1500 max crit; HAM 3000–6000
  unbuffed, 6000–12000 buffed; TTK 20–40 s solo white-con, 60–120 s group-vs-elite, 30–60 s PvP 1v1;
  cooldown/HAM-cost bands per ability tier. Priority [MVP] — gates the Horizontal Progression Test.
- **Progression time (§33.3):** basic mastery 40–60 h; elite 80–120 h; multi-prereq elite 150–250 h; full
  250-point cap 200–300 h cumulative. Testbed cannot observe these horizons live — projection method only.
- **Economic (§33.4):** novice crafter ≈ novice combat mission income; master 2–3x novice; maintenance
  5–15% of weekly income; stall commission 4–6% ([ASSUMPTION], stalls OUT — dormant).
- **Housing/civic/faction (§33.5–33.6):** cost bands + rank thresholds (2500/10k/30k/75k).
  Priority [EXPANSION] — iterate post-MVP; Phase 10 records current fork values beside targets, no changes.

### 2.4 §29.6 Monitoring + §29.2.3/29.3 jobs (HISTORICAL shapes)

- Nightly economic snapshot job (faucet/sink aggregation → net inflation) separate from the real-time loop;
  `WorldTickJob` record with idempotent resumption (pending/running/completed/failed).
- Lightweight anomaly job flags wash trading / price outliers / placement bursts **for human review, never
  auto-action**. Telemetry events structured to populate the §34 dashboard directly.
- Testbed status: snapshot + ledger tagging EXIST (Phase 5 E7: `economic_snapshots`, faucet/sink/transfer
  ledger, per-tick hook, `/api/economy/snapshot`). MISSING: scheduled nightly job shape, anomaly detection
  of any kind, Gini, volatility, profession-distribution, interdependence-event counters, mission-completion
  rates, any dashboard.

### 2.5 Testbed inventory (verified 2026-09-16, PROTOTYPE)

- EXISTS: ledger with faucet/sink/transfer tags (`internal/economy/economy.go`, `economy_db.go`);
  `WriteSnapshot`/`LatestSnapshot` (literal 30-day window); `crafted_items.schematic_id` (provenance query
  path, no new system needed); mission reward faucet hook (`missions.go` turnIn); kill-drop faucet mapping
  10×CLMax capped 200 [PROVISIONAL]; TESTBED_FAST_CYCLE compression convention.
- PROVISIONAL (flagged, untouched since): flat-50 service XP (`elite_http.go` ×3: image-designer session,
  holo-performer, smuggler slice; `ranger.go` pet-tick combat +50); service XP heal/buff 50, revive 100,
  tip 25, stim-use 25 (Phase 6); sidearm-scale PvP TTK inside 30–60 s band (Phase 8, observed not tuned).
- Phase 9 §10 decision 6c routes the **service-XP scaling residual** (flat-50s vs per-point/per-tick) HERE
  as a rate redesign — gated on §7 numeric authorization; provisionals stand until then.

### 2.6 Master GDD constraints that shape this phase

- BAL-001 LOCKED (no premature numerics) + §34 step 7 ("simulation and numeric tuning only after structural
  shapes exist") + §36 lineage ("numeric values remain universally avoided / TBD").
- BAL-002 PROPOSED (balance for interdependence/viability, not DPS parity) — the lens for reading all
  validation output: a DPS-parity gap is not a failure; a collapsed profession category is.
- SAFE-003 PROPOSED (economic observability), TEST-001 PROPOSED (design invariants as tests), TEST-002
  PROPOSED (simulation preference) — joint authority basis for B1–B3 alongside the HISTORICAL §§33/34/29.6.
- OQ-001–OQ-017 remain TBD; Phase 10 answers none of them. It may attach numbers-adjacent *observations*
  (e.g. "fork value X sits outside HISTORICAL band Y") without resolving anything.

## 3. Scope IN

- **B1 — Telemetry completion (read-only instrumentation):** nightly-snapshot job shape (fast-cycle-mapped,
  defaults GDD-given cadences); interdependence-event log (crafted purchase, entertainer/medic heal receipt
  per character per 7 d); profession-distribution query (boxes by category); mission completion/cancellation
  rates; wealth-distribution query (Gini, diagnostic only per §34.3 — tracked, never target-capped);
  resource price volatility per rotation cycle. All append-only observation; no levers, no steering.
- **B2 — §1.3 proxy harness (`phase10test`):** five proxies runnable against a live fresh-DB world, each
  reporting PASS / FAIL-with-telemetry / NOT-MEASURABLE-with-reason. FAIL is an accepted exit outcome
  provided the telemetry is actionable (per the Phase 10 exit disjunction).
- **B3 — Balance validation report:** scripted comparison of fork-observable values against the §33 bands
  (combat damage/TTK spot-checks at sidearm scale, XP-rate projection vs §33.3 horizons, faucet/sink mix vs
  §33.4, maintenance share vs 5–15%). Output is a REPORT (table of fork value vs HISTORICAL band vs gap),
  committed as evidence. Zero value changes.
- **B4 — Service-XP scaling residual (Phase 9 → here):** telemetry characterization of flat-50 awards +
  a rate-redesign OPTIONS paper (per-point / per-tick shapes with explicit numeric asks). Implementation of
  any new rate waits on §7 authorization; default if declined: provisionals stand, characterized.
- **B5 — Anomaly-lite (monitoring, never enforcement):** wash-trade heuristic + price-outlier flag query
  over ledger/market history, output to human-readable review list. No auto-action, no account states
  (per §29.6 "surface signal, not auto-ban"; SAFE-002 TBD untouched).
- **B6 — Polish / stability:** flake-harden the 0–10 suite chain (shared-DB coexistence, fast-cycle maps);
  retire temporary debug logging left by Phase 9; hygiene note extended per package; fork README phase
  table updated (Phase 10 row); no mechanic changes under this package.
- **B7 — Full-chain regression:** phases 0–9 suites re-run green on fresh DB plus new `phase10test`, real
  output pasted, commit per verified package (verify-don't-claim, HD-TST-01 methodology).

## 4. Scope OUT (explicitly deferred, with homes)

- Any numeric retuning (BAL-001/HD-TST-01 — needs its own human numeric authorization, never implied).
- EIC shaping or new residuals (HD-EIC-07 gate — human structural boundaries still pending).
- Provenance / reputation systems (still deferred); auto-enforcement or trust-and-safety actions (no mandate).
- Player-facing full dashboards/charting ([EXPANSION] per §§34.6/33.7 — server-side query endpoints +
  human-readable output only; GUI dashboards are a later pass).
- Structured live playtests with human cohorts (§34.5 recommendations — human-run activity, testbed only
  provides the instrumentation they would read).
- Stalls/escrow follow-up and PvP-combat-XP amendment (separately sequenced per Phase 9 §10 6a/6b — not
  absorbed here).

## 5. Tensions and honest limits (human resolves, proposal does not)

1. **"Balance validation" vs BAL-001:** resolved by construction (§1.2) — validate = observe + report.
   If any B3 gap tempts a tuning fix, that fix is a NEW numeric proposal, not part of this phase.
2. **Time-horizon KPIs are unmeasurable live** (6-month ratios, 40–300 h mastery, 30-day rolling inflation
   on a fresh DB, 90-day city survival): method is rate-projection + fast-cycle mapping + explicit
   NOT-MEASURABLE where projection is dishonest. No simulated-player fabrication to fake a PASS.
3. **Qualitative tests stay qualitative:** Sandbox Validation = observation log template (event, date,
   witnesses); Historical Accuracy = human checklist — the harness reports their existence/absence, never
   a fabricated score.
4. **Gini is diagnostic** (§34.3: cartel concentration is historically accurate emergence) — B1 computes it;
   no threshold, no gate, no lever.
5. **Mission faucet dominates Phase 9+ economies:** B3 must report faucet mix honestly (mission rewards now
   live via turnIn) against the §12.4 5–10% band as observation only.

## 6. Numeric authorization requests (values NOT set by this proposal)

GDD-GIVEN (implement/observe as written, no approval needed): 30-day window + 5–10% band (observed, never
steered); §33.2/33.3/33.4/33.5/33.6 bands (compared-against only); 10 s survey cooldown (reused);
nightly snapshot cadence shape (fast-cycle-mapped, defaults GDD-given); ≥90% / ≥95% / 2–3x KPI targets
(measured-against only).
NEEDS APPROVAL (each: approve, amend, or decline-to-stand-pat):
(1) service-XP redesign rates (per-point/per-tick magnitudes per service type);
(2) anomaly-heuristic thresholds (wash-trade velocity, outlier sigma — flag-only, but numbers nonetheless);
(3) snapshot/aggregation cadence + retention under fast-cycle;
(4) any B3 follow-up value change (explicitly a separate decision, default NO).
PROPOSED interim convention (approval requested): Group-C-style NON-CANONICAL flagging on every
non-GDD number (code comments + HYGIENE_NOTE.md), fork-local magnitudes reused where a magnitude must
exist.

## 7. Acceptance criteria (verify, don't claim)

`go build ./...` + `go vet ./...` clean; fresh DB; phase0–9 suites green PLUS new `phase10test`
(five proxies report PASS / FAIL-with-telemetry / NOT-MEASURABLE; at least the three automatable proxies
— interdependence, crafted-share, combat-TTK spot — produce real numbers); B3 report committed as a file
with fork-value vs HISTORICAL-band table; B5 review-list query demonstrated on seeded wash-trade fixture;
no numeric diff in any mechanic file (enforced by `git diff` review — instrumentation-only proof);
hygiene note extended; commit per verified package.

## 8. Open serving suggestion (non-binding)

Sequence B1 → B2 → B3 → B5 → B4-options-paper → B6 → B7, so measurement exists before any redesign talk.
B4 implementation (if authorized) becomes a separately authorized micro-package after this phase exits.

## 9. Decisions requested (no code until all answered)

1. Base/location — confirm `testbed/swg-phase3-combat/`.
2. Scope — confirm B1–B7 IN / §4 OUT as written, or amend.
3. Observe-don't-steer rule (§1.2) — confirm, or authorize specific tuning levers now (not recommended).
4. Time-horizon method (§5.2) — confirm projection + NOT-MEASURABLE honesty, or direct an alternative.
5. Qualitative-test method (§5.3) — confirm log-template + human-checklist approach.
6. Numeric set (§6) — approve/decline each of (1)–(4); if (1) declined, provisionals stand characterized.
7. Authorisation to implement on approval of 1–6 (cf. §10).

## 10. Implementation status

PENDING — no authorization given, no code written. Entry precondition (Phase 9 suite green + committed)
unmet at time of writing. Next step is an explicit item-7 authorization covering §§3–6 as amended; until
then this file is planning-only. On authorization, execution follows the verify-don't-claim methodology
(`go build` + `go vet` clean, fresh DB, all phase tests re-run old+new with real pasted output, hygiene
note extended per package, commit per verified package), with the human review gate before any
post-Phase-10-equivalent work (HD-TST-01 pattern).

*End of proposal v0.1. To enact: owner approves (or amends) §§3–6/9–10; approval and date recorded
before any code is written.*
