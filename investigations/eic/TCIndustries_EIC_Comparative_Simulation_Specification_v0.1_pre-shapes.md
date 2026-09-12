# TCIndustries — EIC Comparative Simulation Specification v0.1

**Document Type:** Simulation Specification (Non-Canonical)  
**Version:** 0.1 (pre-Functional-Shapes original; superseded by the post-shapes revision filed as `TCIndustries_EIC_Comparative_Simulation_Specification_v0.1.md`)  
**Author:** Author LLM Unknown
**Date:** 2026-08-25  
**Status:** Specification only. Creates no design authority. Does not modify Architecture C.  
**Controlling References:**  
- `TCIndustries_EIC_Candidate_v0.1.md`  
- `TCIndustries_EIC_Candidate_v0.1_Falsification_Pass.md`  
- `TCIndustries_Master_GDD_v1.1.1_Status_Patch.md`  

**Hard constraints on this document:**  
- Do not modify Architecture C.  
- Do not introduce new mechanisms.  
- Do not patch weak mechanisms.  
- Do not resolve HD-1 through HD-9.  
- Build the smallest simulation capable of answering the two unresolved questions identified by the Falsification Pass.  
- If both tested variants fail, say so. That outcome must remain possible.

---

# 1. Purpose

Produce evidence, not design.

The Falsification Pass identified two critical unresolved attack surfaces:

1. **Factory quality parity vs residual specialist value (M4)**  
2. **Formal concentration budget vs soft-pressure alternatives (M1)**

This specification defines the minimum comparative simulation required to generate PASS / FAIL / INCONCLUSIVE results on the relevant invariants under adversarial optimisation.

No other EIC questions are in scope.

---

# 2. Core Metrics

Two primary metrics are required. Both are conceptual; exact formulas are implementation details of the simulation, not design decisions.

## 2.1 External Transaction Advantage (ETA)

For any agent and any production stage or high-value activity:

```
ETA = Value obtained by transacting with an external specialist
    − Value obtained by self-supply / internalisation
```

(adjusted for attention cost, capital cost, and risk)

**Interpretation:**
- ETA ≈ 0 → specialist is economically redundant for that agent  
- ETA moderately positive → voluntary comparative advantage (target band)  
- ETA extremely high and unavoidable → risk of coercive dependence  

Healthy interdependence requires a positive but non-coercive ETA band for a significant fraction of high-value activity across the agent population.

## 2.2 Internalisation Pressure (IP)

For any agent (especially organisations and multi-account decision units):

```
IP = Benefit of owning/controlling an additional production stage
   − Cost of doing so
```

**Interpretation:**
- Persistently high IP across stages → strong pressure toward closed vertical integration  
- Failure condition: every economically successful organisation inevitably internalises every economically important function, rendering independent specialists optional  

Vertical integration is allowed. Total economic closure is not.

---

# 3. Test A — Factory Quality Parity

## 3.1 Question

If factories / industrial processes can match or exceed the raw measured quality of individual specialist output, what economically valuable activity remains that makes specialist players necessary?

## 3.2 Experimental Design

**Condition:** No hard quality ceiling. Factories may achieve measured quality equal to or greater than specialist hand production when given equivalent high-quality inputs and sufficient capital.

**Progressive removal of proposed residual advantages:**

Run sequential conditions in which the following specialist advantages are progressively neutralised or made available to the factory/industrial path:

| Step | Advantage under test | Neutralisation method in simulation |
|---|---|---|
| A0 | Baseline | Both specialist and factory have equivalent resource access and capital |
| A1 | Setup / tuning knowledge | Factory path receives equivalent setup templates or one-time specialist hire that persists |
| A2 | Experimentation / design expertise | Optimal configurations become known and reproducible by factory path |
| A3 | Responsiveness to new resources | Industrial path receives equivalent monitoring / retuning capacity |
| A4 | Customisation | Most demand is for standardised high-quality goods; customisation is a thin niche |
| A5 | Provenance / identity | Consumers optimise pure measured performance; provenance premium set to zero or near-zero |
| A6 | Reputation | New high-quality factory output is treated as equivalent to established specialist reputation |
| A7 | Complementary services | Repair / modification / certification are themselves organisable or low-friction |
| A8 | Full residual collapse | All of the above simultaneously neutralised |

**Agents present in Test A:**  
- Human specialist (focused)  
- Factory / industrial organisation  
- Rational economic optimiser (meta-agent)  
- Consumer population with tunable preference weights (performance vs provenance vs customisation)

**Primary outputs:**  
- Specialist market share of high-value goods at each step  
- ETA for consumers choosing specialist vs factory  
- Whether any residual specialist demand remains economically significant after A5–A8  
- PASS / FAIL / INCONCLUSIVE on M4 invariant

**M4 Invariant under test:**  
Even when automation can produce high measured quality, there remain economically significant categories of demand that continue to reward specialised human players.

**Decision rule:**  
- If specialist demand collapses to near-zero by A5–A8 under optimiser pressure → **FAIL** for current residual-channel theory of M4  
- If material specialist demand survives even after aggressive neutralisation → **PASS** (and identify which residual actually carried the value)  
- If results are sensitive to untested parameters → **INCONCLUSIVE**

**Critical instruction:** Do not invent new residual channels during the simulation to rescue the result. Only test the residual channels already listed in EIC-Candidate v0.1 M4.

---

# 4. Test B — Formal Concentration Budget vs Soft Pressure

## 4.1 Question

Does a formal concentration budget produce an economically significant improvement in interdependence that a pure soft-pressure architecture cannot achieve?

## 4.2 Experimental Design

**Variant B-Formal:**  
A formal concentration budget / concurrent mastery limit exists. Exact numerical form is a simulation parameter, not a design decision. The budget must be binding enough that a single character cannot be elite across the full set of high-value production domains.

**Variant B-Soft:**  
No formal budget. Concentration pressure is produced only by:  
- capacity / attention / maintenance load  
- knowledge and information maintenance cost  
- asset / facility concentration requirements  
- opportunity-cost gradients arising from quality/throughput gaps  
- resource access and logistics constraints  
- reputation compounding on focused play  

(These are the non-budget alternatives already listed under M1 in the candidate. No new ones are added.)

**Agents and comparison matrix:**

| Agent | B-Formal | B-Soft |
|---|---|---|
| Single character (generalist attempt) | ✓ | ✓ |
| Single character (focused specialist) | ✓ | ✓ |
| Multi-account decision unit | ✓ | ✓ |
| Small organisation | ✓ | ✓ |
| Large organisation | later / secondary | later / secondary |
| Rational economic optimiser | ✓ | ✓ |

**Primary outputs for each variant:**  
- Self-supply rate of high-value activity  
- ETA (external transaction advantage) distribution  
- Multi-account advantage relative to single-character specialist  
- Specialist premium  
- Rate at which organisations internalise previously external specialist functions (IP trajectory)  
- New-entrant time/cost to first meaningful economic participation (secondary)

**M1 Invariant under test:**  
Under normal play, the marginal economic return to deepening a primary specialisation exceeds the return to adding another elite-level domain beyond the point at which concentration pressure becomes binding.

**Whole-candidate interdependence invariant under test:**  
High-value activity preferentially involves specialised players or organisations rather than pure self-supply; cooperation is not mandatory.

**Decision rules:**  
- If B-Formal produces materially higher ETA / lower self-supply / lower multi-account advantage than B-Soft under the same optimiser pressure → formal budget **earns consideration**  
- If B-Soft matches or exceeds B-Formal on interdependence metrics → formal budget **has not earned existence**  
- If both variants produce self-supply or closed internalisation as the dominant high-value strategy → **both fail**; Architecture C’s current concentration theory is insufficient  
- If results are dominated by untested parameters → **INCONCLUSIVE**

**Critical instruction:** The formal budget must not be given artificial advantages (e.g., also receiving stronger soft pressures) that the soft variant lacks. Comparison must be clean.

---

# 5. Shared Simulation Elements (Minimal)

To keep the simulation as small as possible while still answering the two questions, the following shared elements are required. They are not new design; they are the minimum representation of existing candidate mechanisms needed for the tests to be meaningful.

## 5.1 Production Stages (simplified)

Represent a short chain only:  
Resource access → Processing → Crafting / Finishing → Manufacturing (volume) → Consumption

Exact number of stages is a simulation parameter. Keep minimal (3–5).

## 5.2 Resources (minimal M2 representation)

- Temporary instances with 2–3 attributes  
- Short lifecycle  
- Geographic or access variation sufficient to create temporary comparative advantage  
- No full resource taxonomy required  

## 5.3 Differentiation (minimal M3 representation)

- Product quality is a function of resource attributes + skill/knowledge depth + (optional) experimentation outcome  
- Factory path and specialist path can be given equivalent or divergent access to this function depending on Test A step  

## 5.4 Demand (minimal M6 representation)

- At least two sinks: one consumption/decay style, one other (construction, status, or service)  
- Demand must be sufficient that production has somewhere to go; exact balance is not the object of these tests  

## 5.5 Agents (required set)

1. Solo generalist  
2. Focused specialist  
3. Multi-account decision unit (N characters under one optimiser)  
4. Small organisation  
5. Factory / industrial agent  
6. Rational economic optimiser (meta-agent that searches for dominant strategies)  
7. Consumer population with tunable preference weights  
8. New entrant (secondary, for exclusion risk observation)

Large organisation can be deferred to a second wave if the first wave already produces clear FAIL results.

## 5.6 What is deliberately out of scope

- Full reputation system dynamics  
- Full provenance depth  
- Exact skill trees or point costs  
- Exact resource spawn mathematics  
- Currency design  
- Combat  
- City / governance systems  
- Any mechanism not already present in EIC-Candidate v0.1  

---

# 6. Output Requirements

The simulation must produce, for each test and each relevant invariant:

```
PASS
FAIL
INCONCLUSIVE
```

accompanied by the metric values that support the judgment (ETA distributions, self-supply rates, specialist market share, IP trajectories, multi-account advantage ratios).

**Explicit required statements in the results report:**

1. Does specialist demand survive factory quality parity after progressive neutralisation of residual channels? (Test A)  
2. Does a formal concentration budget produce material interdependence gains that soft pressure cannot? (Test B)  
3. Under the better-performing concentration variant (or both), does self-supply or closed internalisation become the dominant high-value strategy?  
4. Is there evidence of coercive interdependence (ETA so high that independent function becomes impractical)?  
5. If both Test A residual theory and both concentration variants fail, state clearly: **Architecture C as currently formulated fails the comparative simulation.**

No redesign recommendations are to be included in the results report. Evidence only.

---

# 7. Success and Failure Conditions for the Specification Itself

This specification succeeds if it enables a simulation that can return a clear PASS, FAIL, or INCONCLUSIVE on the two critical questions without requiring additional design decisions.

This specification fails if:
- The tests cannot be run without inventing new mechanisms  
- The metrics cannot distinguish the variants  
- The only way to obtain a PASS is to add untested residual channels or artificial constraints  

In that case the correct report is that the candidate is not yet simulatable in a form that can resolve the open attacks.

---

# 8. Governance and Process Position

**Current formal state (unchanged by this document):**

| Item | Status |
|---|---|
| EIC investigation | COMPLETE |
| Architecture C | PROPOSED / LEADING HYPOTHESIS |
| Adversarial falsification | COMPLETE |
| Validation | NOT ACHIEVED |
| Canonical | NO |
| Human decisions (HD-1–HD-9) | DEFERRED |
| Next operation | Comparative structural simulation (this specification) |

**Process rule reinforced:**  
Candidate → attack → simulation → evidence → human ruling → detailed mechanics  

Do not reverse this order.  
Do not resolve specialisation budget, manufacturing residual theory, or multi-account policy before evidence exists.

---

# 9. Immediate Execution Instruction

Implement the smallest simulation that satisfies this specification.

Prefer clarity and discriminability over realism or completeness.

If implementation reveals that a test cannot be run without adding mechanisms not present in EIC-Candidate v0.1, stop and report that the candidate is under-specified for that test. Do not fill the gap with new design.

---

**Document Control**

**Status:** Simulation specification only.  
**No design authority created.**  
**Architecture C unmodified.**  
**No new mechanisms introduced.**  

**Lineage:** Produced in direct response to the requirement for a narrow comparative simulation specification capable of answering the two unresolved questions from the Falsification Pass. Respects all status and authority rules of Master GDD v1.1.1 and the Authority & Provenance Reconciliation Matrix.

---

*End of Comparative Simulation Specification v0.1*
