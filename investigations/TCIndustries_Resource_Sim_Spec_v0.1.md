# TCIndustries — Resource Lifecycle Sandbox Simulation Specification v0.1

**Status:** EVIDENCE (test specification). Proposes no design. Sets no TCIndustries
value. Every parameter below is **NON-CANONICAL test scaffolding** (BAL-001:
false precision is worse than explicit unresolved questions — these numbers
exist only so the *structural invariants* can be exercised, and none is a
proposal, recommendation, or precedent).

**Author:** Buffy — Codebuff agent (Freebuff), T-04 deliverable, 2026-09-17.
**Task authorization:** owner approval of T-04 under
`proposals/TCIndustries_Gap_Analysis_SWG_Reuse_and_Test_Program_Proposal_v0.1.md`
§9 (recorded there 2026-09-17).
**Invariants under test:** INV-002 (event-generating lifecycle regimes) from
`proposals/TCIndustries_Invariant_Register_v0.1.md`, which anchors to PIL-001
("a resource discovery should be an **event**, not merely an inventory update"),
RES-002, RES-004/006/007. Method authorized by TEST-002 (simulation preference).

## 1. What this simulation tests — and what it does not

**Tested (structural):** whether a temporary, geographically distributed,
attribute-variable resource system with survey-gated information exhibits —
across a parameter envelope — (R1) all three lifecycle regimes as *observable
economic events*, (R2) information-asymmetry value, (R3) anti-monopoly
behavior, and (R4) no stable extract-while-ignoring-quality degenerate strategy.

**Not tested:** real players; TCIndustries attribute sets or mappings (none
exist); any numeric claim ("spawns should last N hours" is never asserted);
any other GDD system (no crafting, cities, or combat exist in the sim).

## 2. World model

- A 40×30 grid of regions (1,200 cells). Regions carry a **terrain class**
  (5 classes) and belong to one of **4 trading zones** (clusters of cells).
- **Resources:** 3 families (metal, organic, gas); each family has 3 subtypes;
  each subtype has 4 measurable attributes, generated per spawn-instance in
  fixed envelopes (below). A spawn-instance occupies one region, has a spawn
  tick, a finite accessible quantity, and an expiry tick.
- **Economy:** each zone has a **demand signal** per resource family that
  fluctuates with its own seeded noise; zone price for a subtype is a function
  of (delivered quantity vs. demand, weighted by delivered attribute quality).
  Prices are local per zone — no global pool — so regional differentials can
  exist and arbitrage is possible for agents.
- **Agents (simulated, deliberately simple):**
  - *Surveyors* (25): pick a region per tick; if it holds an undiscovered
    spawn, discover it (discovery is public — the *location + attributes*
    become known to all; the sim's "information" layer is discovery timing).
  - *Harvesters* (40): travel to discovered spawns, extract per tick at a rate
    scaled by the spawn's remaining quantity (finite pool: extraction depletes).
  - *Traders* (15): buy where price + transport cost is low, sell in the
    highest-margin zone; their activity equalizes prices over time (this is
    the anti-monopoly pressure R3 must survive).

## 3. NON-CANONICAL parameter envelope

Central values were chosen to make every regime reachable; the envelope varies
each around its center. **None of these values is a TCIndustries design
proposal; they exist only to exercise structure.**

| Parameter | Center | Envelope | NON-CANONICAL rationale |
|---|---|---|---|
| Spawn lifetime | 200 ticks | {100, 200, 400} | Short enough to deplete within a run; long enough to trade around. Tests regime visibility vs. churn. |
| Spawn quantity | 5,000 units | {2,000, 5,000, 10,000} | Relative to extraction rate (below) this spans scarce-to-glutted. Tests scarcity-pressure existence. |
| Spawn interval | 40 ticks | {20, 40, 80} per family-subtype | Spans oversupply ↔ drought. Tests discovery cadence as an event trigger. |
| Extraction rate | 50 units/tick/agent | fixed | Calibration constant (agents must be able to exhaust a low-quantity spawn). Not varied; not a design value. |
| Attribute envelope | each attribute uniform in [20, 100] | fixed | Wide enough that quality differentials are large; band ends NON-CANONICAL. |
| Surveyor count | 25 | {10, 25} | Tests information-asymmetry decay under thin vs. dense survey coverage (R2). |
| Trader count | 15 | {8, 15} | Tests anti-monopoly pressure strength (R3). |
| Demand noise σ | 0.15 | fixed | Small relative fluctuation; demand level itself derives from the seeded walk, not a tuned target. |
| Horizon | 2,000 ticks | fixed | Chosen so center-valued spawns cycle ≥5 times. |
| Seeds | 3 per point | {11, 23, 47} | Reproducibility; verdicts require consistency across seeds. |

**Envelope size:** 3 lifetimes × 3 quantities × 3 intervals × 2 surveyor counts × 2 trader counts = **108 parameter points × 3 seeds = 324 runs**. (The T-04 task text said "16 points × 3 seeds" as an order-of-magnitude estimate; the full factorial is 108×3 — reported honestly here and run in full.)

## 4. Event detection (what counts as "observable")

Per run, the recorder tracks per-tick: zone prices per subtype, discovery
events (region, subtype, attributes), quantity-depletion curves, agent
migration between zones (harvester/trader relocations), and trade volumes.
An **event** is detected when a metric moves beyond its own run-local baseline:
- **Price event:** |Δprice| over a 20-tick window > 3× that run's median |Δprice| (run-relative, self-calibrating — no absolute price rule).
- **Traffic event:** zone-level harvester+trader inflow over 20 ticks > 3× run median inflow.
- **Discovery event:** every first-discovery of a spawn-instance (recorded; classified by whether a price/traffic event follows within 60 ticks).

## 5. Verdict rules (mapped to the invariant)

| Requirement (INV-002 decomposition) | Rule for PASS | FAIL | NOT-MEASURABLE |
|---|---|---|---|
| R1a discovery regime visible | ≥90% of runs have ≥1 discovery whose 60-tick window contains a price or traffic event | <90% | recorder failure |
| R1b scarcity regime exists | ≥90% of runs show ≥1 subtype-zone entering a "drought window" (demand present, zero deliveries ≥30 ticks) with a subsequent price rise | <90% | — |
| R1c depletion/rediscovery regime | ≥90% of runs show ≥1 full cycle: discovery → depletion → expiry → rediscovery of same subtype in same region-class | <90% | — |
| R1d not-inventory-churn | median run shows ≥3 detected price events AND ≥3 traffic events | fewer in >10% of runs | — |
| R2 information value | early discoverers (agents extracting within 20 ticks of first discovery) earn ≥1.2× the per-tick yield of late extractors (>100 ticks after), in ≥90% of runs | <90% | — |
| R3 anti-monopoly | no single zone holds >60% of total trade volume in the final 25% of any run | any run violates | — |
| R4 no quality-blind strategy | a quality-blind harvester variant (extracts anything) earns ≤0.9× the yield of a quality-sensitive variant in ≥90% of runs | <90% | — |

All thresholds (90%, 1.2×, 0.9×, 60%, 3×) are **test-acceptance scaffolding**
— they define what this experiment counts as structural success and carry no
design meaning.

## 6. Simulator

Standalone Go program at `investigations/resource-sim/` (stdlib-only; no
testbed code modified; deterministic seeded RNG; JSON telemetry per run +
summary). Envelope run: all 324 runs; results and verdicts in
`TCIndustries_Resource_Sim_Results_v0.1.md`.

## 7. Acceptance criteria (from the task)

1. ✅ every parameter marked NON-CANONICAL with rationale (§3 table, column 4);
2. results state which invariants held/failed/NOT-MEASURABLE (results doc §verdicts);
3. no tuning recommendations anywhere (results doc constraint; violations = rework).

---

*End of spec. The simulator implements §2–§5 exactly; deviations are listed in
the results doc if any arise.*
