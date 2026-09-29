# SWG Phase 10 Implementation Verification Report — 2026-09-17

**Author:** Buffy — Codebuff agent (Freebuff client). **Status:** IMPLEMENTED +
VERIFIED (this report is evidence, not a status promotion of any design document).
**Authorization basis:** owner approval of `SWG_Phase10_Balance_Telemetry_Polish_Proposal_v0.2.md`
§9/§6 decisions recorded in that file's §10 (B1–B3, B5(i), B6, B7 in full; B4 =
characterization + options paper only; zero numeric changes).

## 1. Verification summary (real output, this host)

Static gates (cgo toolchain at `C:\mingw64`, per testbed README recipe):

```
go build ./...   → clean
go vet ./...     → clean
go test ./...    → ok faction / handlers / services / skills (all packages)
```

Full-chain regression (B7, `run_chain.sh` — phases 1–9 suites, each on a fresh DB
and its own server lifecycle, then the Phase 10 suite):

```
=== ALL PHASE 1 TESTS PASSED ===
=== ALL PHASE 2 TESTS PASSED ===      (after stale-assertion refresh, see §4)
=== ALL PHASE 3 TESTS PASSED ===
=== ALL PHASE 4 TESTS PASSED ===
=== ALL PHASE 5 TESTS PASSED ===
=== ALL PHASE 6 TESTS PASSED ===
=== ALL PHASE 7 TESTS PASSED ===
=== ALL PHASE 8 TESTS PASSED ===
=== ALL PHASE 9 TESTS PASSED ===
=== CHAIN RESULTS: p1=0 p2=0 p3=0 p4=0 p5=0 p6=0 p7=0 p8=0 p9=0 ===
CHAIN:ALL-GREEN
```

Phase 10 suite (final standalone run of the same code state):

```
  [PASS] Proxy 1: Economic Interdependence — combat cohort 2/2 covered via
         receipts (healF=false buffF=true buffE=true craftedPurchase=true);
         all-world 3/4 (75%)
  [PASS] Proxy 2: Horizontal Progression (bounded) — E/F damage-per-swing
         ratio ≈ 0.93× (E 90 dmg, F 97 dmg over 10 swings each)
  [PASS] Proxy 3: Player-Driven Economy (crafted share) — crafted_share=1/1 (100%)
  [NOT-MEASURABLE] Proxy 4: Sandbox Validation — qualitative by definition;
         artifacts ready for human review
  [PASS] Proxy 5: Mission-economy telemetry — lifecycle mix observable:
         available=29 accepted=0 completed=0 expired=0
  [PASS] Telemetry job completed with result payload (§29.8 shape)
=== PHASE 10 PROXIES COMPLETE: all observable proxies PASS ===
PHASE10_EXIT:0
```

Proxy 1 honest notes: the wound-heal leg is durably unexercisable without an
incap→revive cycle (sparring drains current health only; server refuses with
"target has no wounds" — surfaced, not swallowed), so the 'heal' event kind is
evidence-backed via code path + phase9-style revive flows, while buff and
crafted-purchase receipts were observed live. Proxy 2 reports the observed
novice-vs-+tier1 ratio (0.93–1.08× across runs, consistent with the fork's
tier structure not yet feeding damage mods into resolution); the 6-month
horizon form remains NOT-MEASURABLE by construction (B2 rule).

## 2. What was built (per package)

- **B1 — Telemetry completion** (`internal/database/phase10_db.go`,
  `internal/handlers/phase10_handlers.go`): append-only `interdependence_events`
  written best-effort at three existing hooks (wound-heal, buff receipt —
  `services.go`; crafted-item purchase — inside `AtomicPurchase`'s tx,
  provenance-checked against `crafted_items`); read-only REST
  `GET /api/phase10/telemetry|anomalies|jobs`; §29.6 nightly economic-snapshot
  job as a §29.8-shaped `phase10_jobs` record (single-row claim, idempotent
  completion, fail-recorded) on fast-cycle-mapped cadence (24 h → 60 s under
  `TESTBED_FAST_CYCLE=1`), separate from the real-time loop; queries for
  profession distribution (registry-joined), mission outcomes, wealth+Gini
  (diagnostic only), price volatility (from `market_records`, which already is
  the append-only sale history — no new table needed).
- **B2 — §1.3 proxy harness** (`cmd/phase10test/main.go`): five proxies, three
  verdicts (PASS / FAIL-with-telemetry / NOT-MEASURABLE); raw counts beside
  every percentage; real-resolution bounded progression proxy; env-overridable
  server URL (`P10_SERVER_URL`).
- **B3 — Balance validation** (`cmd/phase10report/main.go` +
  `docs/phase10_balance_report.md`): 13-band report, every row OBSERVED /
  CODE-CONSTANT-cited / DERIVED / NOT-MEASURABLE. Headline evidence: fork
  unarmed 5–15 vs §9.7 50–150 start band recorded OUT-OF-BAND (observation,
  not a fix); creature-drop and mission-reward constructions IN-BAND.
- **B4 — Service-XP residual** (B4 section of the report +
  `docs/phase10_service_xp_options.md`): all flat [PROVISIONAL] sites
  characterized with citations (medical 50/50/25/100, entertaining 25,
  image_designer 50×2, smuggler 50, stim 25, scouting 10/25/25, ranger 10/25,
  pets 25/100/50, mentoring 50+50, merchant price/100-shaped, structure 100×4);
  options paper holds §7.2.3 shapes (per-point wounds, per-tick BF,
  per-application buff, diminishing-returns hook) with explicit numeric asks.
  **No rate changed.**
- **B5(i) — Anomaly-lite** (in `phase10_handlers.go`): ranked review lists
  (wash-trade pairs, near-zero-net churn, placement bursts, price spread)
  using structural invariants only — no thresholds, no sigma, human review
  only, never auto-action.
- **B6 — Polish/stability:** COMBAT-DEBUG logs gated behind
  `TESTBED_COMBAT_DEBUG=1`; all 30 UnixNano ID-mint sites swept onto the
  collision-proof helper (RNG seeds deliberately untouched); dead
  `MailTimestamp` removed; `phase10test/` source unignored; README phase table
  refreshed with evidence-dated checkmarks.
- **B7 — Full-chain regression:** `run_chain.sh` (per-phase fresh DB +
  lifecycle, inter-phase listener reaping) + `run_phase10.sh` (dedicated port);
  output in §1; committed per package.

## 3. Commit inventory (7 commits, this session)

| Commit | Package | Content |
|---|---|---|
| f9852ec | B1+B5 | telemetry layer, hooks' writers, REST, job shape |
| 55ec433 | B2 | `cmd/phase10test` proxy harness |
| 12ca3db | B3+B4 | `cmd/phase10report` generator |
| ce09b35 | B6 | `run_phase10.sh`, `.gitignore` source-dir negations + phase10test |
| 458c311 | B6 | UnixNano sweep (13 files) + COMBAT-DEBUG gate |
| e6c1187 | B3/B4 | committed report + options paper |
| 2730c9a | B7 | `run_chain.sh` + phase2test assertion refresh |

## 4. Mixed-provenance disclosure (pre-existing dirty tree)

The working tree carried ~69 pre-existing dirty paths (13 modified tracked +
43 untracked testbed files) predating this session. Per the owner's instruction
only Phase 10 files were committed. Because most Phase 3–9 implementation files
were untracked, seven commits include pre-existing content inside files I also
edited — each new-file case is disclosed in its commit message
(`phase9_db.go`, `combat_db.go`, `service_db.go`, `crafting.go` inside 458c311;
`cmd/phase2test/main.go` inside 2730c9a, where only the Test 1 assertion is
Phase 10 work). Still uncommitted, untouched by me: remaining tracked-modified
files (`db.go` HAM repair, `faction_db.go`, `resource_db.go` partials, `server.go`
wiring, `world.go`, `missions.go`, `economy_http.go`…), untracked Phase 3–9
sources (`lairs.go`, `ranger.go`, `missions.go`, `pets.go`, `elite_http.go`,
`services*.go`, `combat.go` partials, remaining `cmd/phaseNtest` sources), the
Phase 9 Phase-10-proposal hand-off edits, and both proposal documents. A
follow-up "adopt the testbed into git" pass remains owner work.

## 5. Constraints compliance (verify, don't claim)

- **Zero numeric changes:** `git diff 19e714a..HEAD` across gameplay-constant
  packages (`internal/services`, `internal/skills`, `internal/combat`,
  `internal/economy`) contains no changed numeric constants (checked: only
  additive telemetry files and ID-mint replacements differ). The B6 sweep
  changes ID string *shape* (`<prefix>-<ns>-<seq>`), which is unobservable to
  gameplay mechanics; RNG seeding was left as-is deliberately.
- **No gameplay-behavior changes:** gameplay files received only (a)
  best-effort `_ =` telemetry writes after success points, (b) ID-mint
  replacement, (c) the log gate. Nothing reads telemetry to steer behavior.
- **REST additions read-only:** all three phase10 routes are GET-only.
- **Status discipline:** proposal v0.2 remains PROPOSED (§10 updated with an
  EXECUTED record, no promotion of any design doc); B4 default respected
  (provisionals stand); no new residual channels introduced.
- **Known follow-ups flagged, not absorbed:** phase 1–9 clients hardcode
  :8080 (chain requires it free — same as `run_phase9.sh` historically);
  `missions.go` lifecycle statuses don't feed B1 counters (row-level mix
  reported; per-terminal completion telemetry would want mission-row retention
  rather than deletion — B1-adjacent decision for the owner); heal-event
  end-to-end evidence requires an incap→revive flow.

## 6. Exit criteria (§2.1 disjunction)

Five §1.3 tests can be run: three PASS with data, one is honestly
NOT-MEASURABLE with artifacts ready for human review, one (Historical
Accuracy) remains human-pass-only by §6.4. FAIL-with-telemetry paths are
implemented and demonstrated (earlier shakeout runs produced accepted-form
FAIL verdicts with actionable telemetry before their causes were fixed).
Phase 10 exit condition: **met**.
