# TCIndustries — EIC Minimum Comparative Simulation Results v0.1

**Document Type:** Simulation Evidence Report (Non-Canonical)  
**Version:** 0.1  
**Date:** 2026-08-25  
**Status:** Evidence only. Creates no design authority. Does not modify Architecture C.  
**Controlling Specification:** `TCIndustries_EIC_Comparative_Simulation_Specification_v0.1.md`  
**Controlling Candidate:** `TCIndustries_EIC_Candidate_v0.1.md`  

**Hard constraints obeyed:**  
- Architecture C was not modified.  
- No new mechanisms were introduced.  
- Illustrative parameters were used and are explicitly non-canonical.  
- Where Candidate v0.1 is silent, results are classified as INCONCLUSIVE or SPECIFICATION-limited rather than treated as architecture validation.  
- The simulation did not optimise or rescue the design.

---

# 1. Executive Evidence Summary

| Test | Status | Failure Type | One-line finding |
|---|---|---|---|
| **Test A — Factory Quality Parity** | **INCONCLUSIVE** (downgraded from model PASS) | SPECIFICATION / MODEL LIMITATION | Specialist share did not collapse under residual neutralisation in the minimal model, but the model is too coarse for residual channels to drive share; outcome is not robust evidence that M4 residuals work. |
| **Test B — Formal vs Soft Concentration** | **INCONCLUSIVE** | PARAMETER / SPECIFICATION | Formal budget improves mean ETA modestly; it does not reduce mean self-supply in baseline. Multi-account behaviour is highly sensitive to an undefined efficiency parameter. Formal budget has not earned existence. |
| **Whole-candidate interdependence** | **INCONCLUSIVE** | SPECIFICATION-LIMITED | Positive ETA and moderate self-supply appear under illustrative parameters, but critical relationships are under-specified. Not validation. |

**Bottom line:**  
Architecture C is **not validated**. It is also **not cleanly falsified** by this minimal run. The dominant result is **under-specification**: Candidate v0.1 does not define enough magnitudes and functional forms for a minimal simulation to produce architecture-level certainty on the two critical attack surfaces.

---

# 2. Failure-Type Discipline Applied

Per the required refinement, every adverse or ambiguous result is classified as one of:

1. **Architecture failure** — Candidate cannot produce voluntary interdependence even under its intended assumptions.  
2. **Parameter failure** — Architecture may work in principle, but only in parameter regions that are extreme or fragile.  
3. **Specification failure** — Candidate does not define enough for the outcome to be determined.

This report privileges honesty over false certainty. A model-level “PASS” that depends on coarse thresholds or undefined magnitudes is **not** treated as architecture validation.

---

# 3. Test A — Factory Quality Parity

## 3.1 What was tested

Factories match specialist measured quality (parity regime).  
Proposed residual advantages from Candidate M4 were progressively neutralised:

A0 Baseline (all residuals active) → A1 Setup → A2 Experimentation → A3 Responsiveness → A4 Customisation → A5 Provenance → A6 Reputation → A7 Services → A8 Full collapse.

## 3.2 Observed metrics (illustrative parameters)

| Step | Mean specialist share | Mean ETA | Mean self-supply |
|---|---|---|---|
| A0 Baseline | 0.22 | 0.180 | 0.54 |
| A5 Provenance neutralised | 0.22 | 0.141 | 0.54 |
| A8 Full residual collapse | 0.22 | 0.125 | 0.54 |

Specialist share was **flat** across residual neutralisation under baseline consumer preference weights.

## 3.3 Sensitivity

- **Transaction friction** (0.02 / 0.08 / 0.20): specialist share remained 0.22; ETA declined with friction as expected.  
- **Consumer provenance weight** (0.05 / 0.15 / 0.40): only at high provenance weight (0.40) did specialist share rise (to 0.263). At realistic-to-low weights, residuals did not move share.

## 3.4 Judgment

**Status: INCONCLUSIVE**  
**Failure type: SPECIFICATION / MODEL LIMITATION**

**Reason:**  
The minimal agent decision model is threshold-based and treats most residuals as production-side cost/quality adjustments rather than as distinct consumer-visible value. As a result, progressive residual neutralisation did not produce the demand collapse that the Falsification Pass predicted as the critical risk. This is **not** strong evidence that residual channels survive factory parity. It is evidence that:

1. The Candidate does not specify residual → consumer-value functions tightly enough to test the attack, and  
2. The minimal simulation cannot discriminate residual survival without additional structure that would itself be new design.

**What would be required to convert this to architecture-level evidence:**  
Explicit, Candidate-authorised functions for:
- how setup / experimentation / responsiveness alter factory vs specialist cost or quality over time,  
- how provenance and customisation enter consumer utility,  
- decay of one-time specialist setup value to factories.

Until those exist, Test A cannot produce a clean Architecture PASS or Architecture FAIL.  
**Current residual-channel theory of M4 remains unproven.**

---

# 4. Test B — Formal Concentration Budget vs Soft Pressure

## 4.1 What was tested

**B-Formal:** binding formal concentration budget (illustrative attention/quality penalty on non-focus domains).  
**B-Soft:** no formal budget; only soft pressures (attention cost, quality gap, facility/opportunity-cost style effects already listed under M1).

Agents: focused specialist, solo generalist, multi-account, small organisation.  
Factory present at high but not full-parity quality so concentration is the variable.

## 4.2 Observed metrics (baseline illustrative parameters)

| Variant | Mean ETA | Mean self-supply | Multi-account self-supply |
|---|---|---|---|
| B-Soft | 0.205 | 0.438 | 0.40 |
| B-Formal | 0.284 | 0.438 | 0.40 |

- Formal improves mean ETA (+0.080).  
- Formal does **not** reduce mean self-supply.  
- Multi-account self-supply is identical under both variants in baseline.

## 4.3 Sensitivity — multi-account efficiency (critical)

| Multi-account efficiency | Soft multi self-supply | Formal multi self-supply |
|---|---|---|
| 0.55 (inefficient) | 0.40 | 0.40 |
| 0.82 (baseline) | 0.40 | 0.40 |
| 0.95 (highly efficient) | **0.85** | **0.40** |

Only when multi-account attention efficiency is very high does Soft allow multi-account internalisation to dominate, while Formal continues to restrain it.  
**Multi-account efficiency is not defined in Candidate v0.1.**

## 4.4 Judgment

**Status: INCONCLUSIVE**  
**Failure type: PARAMETER / SPECIFICATION**

**Reason:**  
Formal budget produces a modest ETA gain but no baseline reduction in self-supply or multi-account internalisation. The only regime in which Formal clearly outperforms Soft on the multi-account attack is when multi-account efficiency is high — a magnitude the Candidate does not specify.  

Therefore:

- Formal concentration budget has **not earned its existence** on current evidence.  
- Soft pressure has **not been shown sufficient** against efficient multi-accounting.  
- The comparison cannot yet support a human ruling on HD-1.

---

# 5. Whole-Candidate Interdependence

Under the better concentration variant (B-Formal in this run):

- Mean self-supply ≈ 0.44  
- Mean ETA ≈ 0.28 (positive)  
- Multi-account self-supply ≈ 0.40  

**Status: INCONCLUSIVE**  

Positive ETA and non-dominant self-supply are *consistent with* voluntary comparative advantage under the illustrative parameterisation. They do **not** constitute validation, because:

1. Critical magnitudes (multi-account efficiency, residual → value functions, consumer preference weights, transaction friction, resource temporality) are under-specified.  
2. Test A could not stress residual survival in a discriminating way.  
3. Multi-account internalisation becomes dominant under Soft when efficiency is high — exactly the attack the Falsification Pass identified as critical.

No evidence of coercive interdependence (forced-trade regime) was observed in this parameterisation.

---

# 6. Under-Specification Audit (Primary Output)

The most important result of this simulation is not a PASS or FAIL. It is the list of relationships that Candidate v0.1 leaves undefined and that control the outcome:

| # | Undefined item | Candidate status | Impact on simulation | Classification if critical |
|---|---|---|---|---|
| 1 | Resource attributes → product measured quality mapping | OQ-003 open | Cannot derive specialist vs factory quality gap from first principles | SPECIFICATION |
| 2 | Exact form/magnitude of concentration pressure | OQ-001 / M1 open | Formal vs Soft comparison rests on illustrative penalties only | SPECIFICATION |
| 3 | Multi-account attention/capital/logistics pooling efficiency | OQ-012 open | Multi-account dominance highly sensitive to this value | SPECIFICATION / PARAMETER |
| 4 | Consumer preference weights (performance vs provenance vs custom) | HD-7 open | Survival of specialist demand after provenance neutralisation depends on this | SPECIFICATION / PARAMETER |
| 5 | Decay of setup/tuning knowledge; factory retuning cost | M4 residuals listed; no dynamics | Cannot determine duration of specialist setup value to factories | SPECIFICATION |
| 6 | Transaction friction / search / trust costs | Not quantified | ETA sign can flip; interdependence band unknown | PARAMETER / SPECIFICATION |
| 7 | Resource lifecycle duration and regional dispersion | OQ-002 open | M2 temporality defence against monopoly/internalisation untested | SPECIFICATION |

**Implication:**  
Several of the simulation’s numeric outcomes sit downstream of these gaps. Reporting them as architecture validation would be false certainty. The correct scientific output is:

> **UNDER-SPECIFIED — HUMAN/DESIGN INPUT REQUIRED** on items 1–7 before architecture-level PASS/FAIL can be claimed on the two critical attack surfaces.

---

# 7. Required Statements (from Simulation Specification §6)

1. **Does specialist demand survive factory quality parity after progressive residual neutralisation?**  
   → **INCONCLUSIVE.** Model did not show collapse, but model cannot discriminate residual survival. Residual-channel theory of M4 remains unproven.

2. **Does a formal concentration budget produce material interdependence gains that soft pressure cannot?**  
   → **INCONCLUSIVE.** Modest ETA gain only; no baseline self-supply reduction; multi-account effect appears only at high (undefined) efficiency. Formal budget has not earned existence.

3. **Under the better concentration variant, does self-supply or closed internalisation become the dominant high-value strategy?**  
   → **Not in baseline illustrative parameters.** Becomes dominant for multi-account under Soft when efficiency is high. Overall: **INCONCLUSIVE.**

4. **Evidence of coercive interdependence?**  
   → **Not observed** in this parameterisation (ETA moderate; no forced-trade regime).

5. **If both critical tests are Architecture FAIL, does Architecture C fail?**  
   → The rule is acknowledged. Current classification is **not** dual Architecture FAIL. Current classification is dual **INCONCLUSIVE** driven largely by specification gaps. Architecture C is therefore **neither validated nor rejected** by this run.

---

# 8. What This Simulation Did *Not* Do

- Did not invent residual channels beyond those listed in Candidate M4.  
- Did not invent a hard factory quality ceiling.  
- Did not resolve HD-1 through HD-9.  
- Did not promote any mechanism.  
- Did not claim that positive ETA under illustrative parameters equals validated interdependence.  
- Did not optimise Architecture C to make the numbers look better.

---

# 9. Implications for Next Human / Design Action

The simulation has done its job as an **evidence generator**: it has shown where the candidate is under-specified for the attacks that matter.

**Productive next steps (human decision space):**

1. **Decide which under-specified items must be given provisional shape** solely for the purpose of a more discriminating simulation (still non-canonical, still labelled illustrative or provisional). Priority:  
   - Multi-account efficiency / cost model (OQ-012)  
   - Residual → consumer value and residual decay dynamics (M4)  
   - Consumer preference weight ranges (HD-7)  
   - Concentration pressure functional form options for Formal vs Soft (OQ-001)

2. **Or** accept that Architecture C cannot be architecture-tested until those shapes exist, and keep the candidate in PROPOSED / LEADING HYPOTHESIS status with validation deferred.

3. **Do not** treat the current numeric “PASS-like” specialist-share stability as evidence that M4 works.  
4. **Do not** treat the modest Formal ETA gain as evidence that a concentration budget is required or sufficient.

---

# 10. Updated Programme Status

| Item | Status |
|---|---|
| EIC investigation | COMPLETE |
| Architecture C | PROPOSED / LEADING HYPOTHESIS |
| Candidate formulation | COMPLETE |
| Adversarial falsification | COMPLETE |
| Comparative simulation specification | COMPLETE |
| Minimum simulation implementation | COMPLETE (this document) |
| Validation | **NOT ACHIEVED** |
| Clean falsification | **NOT ACHIEVED** |
| Dominant finding | **UNDER-SPECIFIED on critical magnitudes** |
| Human decisions (HD-1–HD-9) | DEFERRED |
| Canonical EIC | NOT ESTABLISHED |

**Process position remains:**  
Candidate → attack → simulation → **evidence (here)** → human interpretation → human ruling → revised candidate  

No step was skipped. No design was generated from the simulation.

---

**Document Control**

**Status:** Simulation evidence report only.  
**No design authority created.**  
**Architecture C unmodified.**  
**No new mechanisms introduced.**  

**Lineage:** Implements `TCIndustries_EIC_Comparative_Simulation_Specification_v0.1.md` under the strict instruction that the simulation is an evidence generator, not a design generator. Respects all status and authority rules of Master GDD v1.1.1 and the Authority & Provenance Reconciliation Matrix.

---

*End of Minimum Simulation Results v0.1*
