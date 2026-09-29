# TCIndustries — Resource Lifecycle Sandbox Simulation Results v0.1

**Status:** EVIDENCE. Reports what was observed in a NON-CANONICAL sandbox; no
design decisions, no tuning recommendations, no parameter proposals. All
parameters come from `TCIndustries_Resource_Sim_Spec_v0.1.md` §3 and carry no
design authority (BAL-001).

**Author:** Buffy — Codebuff agent (Freebuff), T-04 deliverable, 2026-09-17.
**Task authorization:** owner approval of T-04 (Gap-Analysis proposal §9,
recorded 2026-09-17).
**Invariant under test:** INV-002 (decomposed R1a–d, R2, R3, R4) from
`proposals/TCIndustries_Invariant_Register_v0.1.md`.
**Telemetry:** raw per-run JSON at `investigations/resource-sim/telemetry_runs.json`
(324 records); simulator source `investigations/resource-sim/main.go`
(vet-clean, deterministic, stdlib-only).

---

## 1. Run summary

- **324 runs** (108 parameter points × 3 seeds), 2,000 ticks each; wall time ~52 s.
- Every run produced all required telemetry fields; no run crashed or aborted.

## 2. Verdicts (three-form taxonomy per spec §5)

| Req | Meaning | Threshold | Observed | Verdict |
|---|---|---|---|---|
| **R1a** | Discovery → event (price/traffic within 60 ticks) | ≥90% of runs at ≥90% of discoveries | **15.1%** of runs met the 90%-of-discoveries bar (median per-run follow rate 75.1%) | **FAIL-with-telemetry** |
| **R1b** | Drought window (demand present, zero supply ≥30 ticks) with ≥5% scarcity-era price rise | ≥90% of runs show ≥1 | **0 of 324 runs** recorded a counted drought win | **FAIL-with-telemetry** |
| **R1c** | Full cycle: discovery → depletion → expiry → rediscovery (same region-class+kind) | ≥90% of runs show ≥1 | **100%** of runs (median 38.5 cycles/run) | **PASS** |
| **R1d** | Not inventory churn: ≥3 price events AND ≥3 traffic events per run | >90% of runs | **48.8%** (median 69 price events; median 3 traffic events; 123/324 runs had zero traffic events) | **FAIL-with-telemetry** |
| **R2** | Information value: early extractors (≤20 ticks from discovery) earn ≥1.2× late extractors (≥60 ticks) per unit value | ≥90% of runs | **74.4%** of runs (median early:late value ratio ≈ 1,870×; 83 runs had zero late-extraction, auto-failing) | **FAIL-with-telemetry** |
| **R3** | Anti-monopoly: no zone >60% of trade volume | 100% of runs | **100%** (median max zone share 38.4%) | **PASS** |
| **R4** | No quality-blind strategy: QB value yield ≤0.9× QS | ≥90% of runs | **48.1%** (median QB 654 vs QS 718 — a real but sub-threshold gap ≈0.91×) | **FAIL-with-telemetry** |

**Overall INV-002 verdict: FAIL-with-telemetry** — the lifecycle *cycles* and
*anti-monopoly structure* hold (R1c, R3), while the *event-visibility*,
*scarcity-signal*, and *quality-sensitivity* properties did not reach their
spec'd acceptance bars under this sandbox's agent population and detection
rules.

## 3. What the telemetry says (mechanics of each failure)

1. **R1a/R1d — events happen but not reliably *at* discoveries.** Price events
   are frequent (median 69/run) and discovery-follow rates are high in bulk
   (median 75%), but the strict coupling "each discovery should precede a
   detectable event within 60 ticks" fails because discovery clustering
   (interval=20 points oversupply spawns into the same windows) and
   simultaneous multi-spawn depletion smear event attribution. Traffic events
   are the weaker leg: 123/324 runs recorded none — harvester migration is
   rare when discovered stock is evenly spread across four zones.
2. **R1b — droughts never counted, by construction of the population.** With
   40 harvesters + interval∈{20,40,80} + 9 kinds across 4 zones, every
   zone-kind with demand>0.6 receives discovered stock quickly; supply==0
   windows ≥30 ticks effectively require spawn droughts co-occurring with
   surveyor absence — the agent population is too dense relative to the
   NON-CANONICAL spawn cadence for scarcity windows to form.
3. **R2 — early/late asymmetry is enormous where measurable** (median ratio
   ~1,870×), but 83 runs have *zero* late extraction (spawns fully extracted
   within 60 ticks of discovery at high harvester density), so the per-run
   threshold fails despite the structural effect being present and strong.
4. **R4 — quality sensitivity is real but small at this scale.** The QB-vs-QS
   gap (≈0.91×) arises only from spawn-selection preference, not from any
   price/return channel available to the agents (value delivered feeds zone
   price quality-weighting, but agents are not paid per-quality by the sim).
   The signal exists; the agent economics do not amplify it past the 0.9× bar.

## 4. Implementation deviations from spec (full disclosure)

| ID | Deviation | Why |
|---|---|---|
| D1 | Price-event detector: 20-tick window deltas vs. median of window deltas (not 1-tick deltas) | The spec'd formula mixed time constants and fired ~always (44k events/run smoke test) |
| D2 | Traffic events = zone relocations (traders + harvesters) vs. 3× run-median relocation | Original "inflow" metric was structurally always zero |
| D3 | Extraction delivers into the spawn's local zone `delivered` stock | Without it the trader/arbitrage layer had nothing to carry (chain was broken) |
| D4 | Added a demand-scaled consumption sink per zone-kind | Without a sink, supply monotonically accumulates; droughts and price cycles are impossible |
| D5 | Traders scan all kinds × zones for best positive-margin route | Random-kind selection never found carryable stock |
| D6 | QB/QS pair placed in the same zone; QS scores by attribute sum, QB is quality-indifferent | Deconfound location from the quality comparison |
| D7 | Drought "win" compares last zero-supply price vs. entry price (≥5% rise over ≥30 ticks) | Comparing post-resupply price inverts the signal (resupply lowers price) |
| D8 | Transport cost parameter set to 0.05 (per-unit-distance penalty 0.05 vs. demand spread ≤2.0) | Initial value 5.0 made every route unprofitable; zero trades |
| D9 | Harvester population = full 40 (the QB/QS pair are *additional* instrumented agents) | Initial off-by-one halved the population and starved R2/R4 statistics |
| — | Envelope size: 108 points × 3 seeds = 324 runs (task text estimated "16 points") | Full factorial per spec §3; more runs, same structure |

All deviations are recorder/agent-mechanics corrections made during
shakeout; the parameter envelope and its NON-CANONICAL values are unchanged
from the spec except D8 (one scaffolding constant, with rationale above).
Spec §3's table is amended by this document; the spec file itself is retained
unmodified for provenance.

## 5. Honest interpretation (what this run does and does not establish)

- **Does establish (structurally, in this sandbox):** full lifecycle cycling
  works (R1c 100%); decentralized arbitrage keeps any zone from dominating
  trade (R3 100%); early information confers large extraction advantage where
  measurable (R2's strong ratios); quality-selection produces measurable
  separation between agent strategies (R4's gap direction is consistent).
- **Does NOT establish:** that TCIndustries resources *should* behave this way;
  that any parameter here is a design value; that the failing requirements are
  properties of the *design shape* rather than of this sandbox's dense agent
  population and strict per-run detection bars. The FAIL verdicts are
  **findings about the test rig as much as about the pattern** — they
  identify which detection/agent-density questions a future iteration must
  resolve (e.g., event-attribution windows, population/scarcity balance,
  quality-denominated agent returns).
- **No tuning recommendations** are made, per the task's acceptance criteria.
  The FAIL-with-telemetry rows carry their diagnosis inline; next steps are a
  human question (rerun with amended detection rules, extend the sandbox, or
  accept the structural results as sufficient at shape level).

## 6. Acceptance self-audit (T-04 acceptance criteria)

| Criterion | Result |
|---|---|
| Every parameter marked NON-CANONICAL with rationale | ✅ spec §3 table (rationale column); D8 amendment documented with rationale; results doc restates no value as design |
| Results state which invariants held/failed/NOT-MEASURABLE | ✅ §2 verdict table (2 PASS, 5 FAIL-with-telemetry, 0 NOT-MEASURABLE at run level; §5 states the interpretation limits) |
| No tuning recommendations | ✅ §5 explicitly routes next steps to the owner; no parameter proposal anywhere |

---

*End of results v0.1. Telemetry: `telemetry_runs.json` (324 records, one per
run). Reproduce: `go build -o resourcesim . && ./resourcesim -out telemetry_runs.json`.*
