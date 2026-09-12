# TCIndustries — EIC Functional Shapes Refinement  
## Targeted γ Closure v0.2

**Document Type:** Narrow Provisional Refinement (Simulation Assumptions Only)  
**Version:** 0.2  
**Date:** 2026-08-25  
**Status:** PROVISIONAL — NOT HUMAN-LOCKED  
**Authority:** None. This document creates no canonical design.  

**Controlling References:**  
- `TCIndustries_EIC_Comparative_Simulation_Results_v0.1.md` (source of the four under-specified relationships)  
- `TCIndustries_EIC_Provisional_Functional_Shapes_v0.1.md` (parent shapes)  
- `TCIndustries_EIC_Comparative_Simulation_Specification_v0.1.md`  
- `TCIndustries_EIC_Human_Rulings_2026-08-25.md` (HD-EIC-01–04)  
- `TCIndustries_Master_GDD_v1.1.1_Status_Patch.md`

**Purpose:**  
Close only the four relationships that prevented Package γ from reaching a Phase-6 judgment.  

This is a targeted γ-closure patch. It is **not**:  
- a new architecture;  
- a general redesign of Functional Shapes;  
- an attempt to rescue γ;  
- a source of new residual channels;  
- a numeric tuning exercise;  
- a promotion of any shape to LOCKED.

**Outcome permission:**  
Each refinement is required to leave open the possibility that, after re-test:  
- γ′ SURVIVES Phase 6,  
- γ′ FAILS Phase 6, or  
- γ′ remains INCONCLUSIVE.  

No formulation may be written so that survival is the only possible result.

---

# 1. Scope Discipline

Only the four under-specified relationships identified in Comparative Simulation Results v0.1 §6 are in scope:

1. S4 — Facility opportunity cost after organisational ownership.  
2. S3 + R3 + R5 — Cost and completeness of organisational internalisation of discovery, regional, and temporal information.  
3. Demand residual weight — Whether discovery-linked / situational / novelty demand remains economically meaningful as average capability rises.  
4. C4 — Distinction between process knowledge that can be documented/transferred and knowledge that requires ongoing specialist involvement.

Everything else in Provisional Functional Shapes v0.1 remains unchanged. Packages α and β are not reopened. No new residual channels are introduced.

---

# 2. Refinement 1 — S4 Facility Opportunity Cost

**Existing shape (from Functional Shapes v0.1):**  
S4 — Asset / facility concentration. Elite production requires specialised, non-trivial facilities whose opportunity cost is high. Layer primarily constrained: Participant / organisation. Risk: Capital concentration may dominate; facilities can be owned by organisations.

**Unresolved relationship that blocked Phase 6:**  
What structural relationship makes facility ownership or concentration economically costly or bounded *after* an organisation is allowed to own facilities?

**Candidate structural formulations (PROVISIONAL):**

| ID | Formulation | Core structural idea | Falsification implication |
|----|-------------|----------------------|---------------------------|
| S4-a | Opportunity-cost of specialised capital lock-up | A facility optimised for one high-value domain cannot be cheaply re-purposed; capital committed to it has high switching cost and forgone alternative production. | If organisations can cheaply maintain parallel specialised facilities across domains, or if switching cost is negligible, S4-a fails to constrain ICR. |
| S4-b | Locality-bound facility advantage | High-value production is tightly coupled to regional resource or temporal conditions; facilities derive advantage from location that cannot be fully centralised. | If logistics or teleportation-like movement erase locality, or if organisations can site facilities optimally everywhere at low cost, S4-b collapses. |
| S4-c | Maintenance / utilisation threshold | Specialised facilities impose ongoing upkeep or minimum utilisation requirements that rise non-linearly with the number of distinct high-value domains an organisation attempts to cover. | If upkeep is linear or easily subsidised by scale, the pressure disappears. |
| S4-d | Information-facility coupling | The facility’s effective performance depends on continuous specialist information input; ownership of the physical asset does not automatically confer the information advantage. | If process settings or discovery data can be fully encoded into the facility, ownership alone closes the residual. |

**Explicit non-choices:**  
No numeric capital costs, upkeep rates, or switching penalties are selected.  
No claim is made that any of S4-a–d is required or sufficient.

**Re-test implication:**  
A γ′ that incorporates one or more of these formulations can be subjected to Phase 3 and Phase 6. Survival is possible only if the chosen formulation still generates positive ETA for independent participants after organisational facility ownership. Failure remains fully available.

---

# 3. Refinement 2 — S3 + R3 + R5 Information / Discovery Internalisation

**Existing shapes:**  
- S3 — Knowledge / information maintenance. Specialised knowledge decays or requires active practice / discovery.  
- R3 — Uncertainty / discovery-linked. Highest attributes require ongoing discovery or experimentation; static maps are incomplete.  
- R5 — Temporal / regional availability. Quality or attributes vary by time and region, rewarding local knowledge and logistics.

**Unresolved relationship that blocked Phase 6:**  
What happens to the cost and completeness of maintaining discovery, regional, and temporal information across an organisation? Can information be recorded, shared, delegated, or replicated so that the organisation fully internalises the advantage?

**Candidate structural formulations (PROVISIONAL):**

| ID | Formulation | Core structural idea | Falsification implication |
|----|-------------|----------------------|---------------------------|
| Info-a | Active maintenance requirement | High-value discovery and regional knowledge decays or becomes stale unless continuously refreshed by agents who are present and practising in the relevant domains/regions. Recording produces incomplete or time-limited value. | If recorded knowledge is permanent, complete, and costlessly shareable inside an organisation, Info-a fails. |
| Info-b | Coverage cost scales with domain × region × temporal resolution | An organisation that wishes to match independent specialist information advantage must maintain active presence or agents across the relevant combinatorial space; cost rises with the number of domains and regions covered. | If a small number of agents or automated sensors can cheaply cover the space, the cost pressure disappears. |
| Info-c | Delegation friction / principal-agent loss | Even when information can be shared, the organisation incurs coordination, trust, or incentive-alignment costs that reduce the effective value of delegated discovery relative to an independent specialist who owns the upside. | If delegation is frictionless and specialists can be perfectly contracted, Info-c provides no residual. |
| Info-d | Discovery residual is non-rival but time-sensitive | Information can be shared, but its highest-value window is short; first-mover or local-presence advantage persists even after later dissemination. | If dissemination is instantaneous and the value window is long, the residual collapses to pure first-mover (which organisations can also capture). |

**Explicit non-choices:**  
No claim that information markets, patents, or secrecy mechanics are required.  
No numeric decay rates or coverage costs.  
No new residual channel beyond the discovery / regional / temporal information already present in R3 + R5 + S3.

**Re-test implication:**  
γ′ can incorporate one or more of these formulations. Phase 2–3 and Phase 6 can then test whether independent specialists still retain positive ETA on discovery-linked activity after organisational information-sharing and multi-account coverage. Both survival and failure remain open.

---

# 4. Refinement 3 — Demand Residual Weight

**Existing treatment in Functional Shapes v0.1 and Package γ:**  
Demand includes elevated situational + organisational + discovery-linked components. The Comparative Simulation Results showed that advantage is sensitive to the weight placed on these residual-sensitive demand types.

**Unresolved relationship that blocked Phase 6:**  
What structural relationship determines whether discovery-linked, situational, novelty, or region-specific demand remains economically meaningful as average production capability rises?

**Candidate structural formulations (PROVISIONAL):**

| ID | Formulation | Core structural idea | Falsification implication |
|----|-------------|----------------------|---------------------------|
| Dem-a | Capability-expanding demand | Higher average capability unlocks new desirable attribute combinations or situational optimisations that did not previously exist; demand for the frontier therefore does not collapse merely because the previous frontier becomes common. | If all valuable attribute combinations are known and static, rising capability simply fills existing demand with cheaper supply and residual weight falls. |
| Dem-b | Situational / local optimality persists | Even when global best-in-class goods exist, local or situational conditions (resource availability, logistics, timing, organisational requirements) continue to make regionally or temporally matched goods higher-value than generic high-quality substitutes. | If logistics and information make global goods perfect substitutes for local needs, Dem-b fails. |
| Dem-c | Organisational and prestige demand remain differentiated | Organisations and status-seeking participants continue to value goods that signal discovery provenance, process novelty, or specific optimisation even after average quality rises. | If all demand optimises pure measured performance and cost, residual weight goes to zero. |
| Dem-d | Replacement and experimental demand | Ongoing experimentation, failure, and replacement cycles create continuous demand for process knowledge and discovery-linked inputs that pure volume manufacturing does not fully satisfy. | If production becomes highly reliable and experimental demand is negligible, this channel closes. |

**Explicit non-choices:**  
No numeric demand elasticities or preference weights.  
No claim that prestige or identity value is required (those remain deferred to later systems).  
No invention of new demand types beyond those already recognised in Functional Shapes §6.6.

**Re-test implication:**  
Phase 5 (Demand Neutralisation) and Phase 6 can apply progressive reduction of residual demand weight under each formulation. A package survives only if positive ETA remains under a non-trivial residual weight that the formulation itself makes structurally plausible. Failure remains available if residual weight collapses to zero under rising capability.

---

# 5. Refinement 4 — C4 Process Knowledge Replication

**Existing shape:**  
C4 — Process knowledge / information advantage. Specialists possess or generate process knowledge that improves yields or unlocks options.

**Unresolved relationship that blocked Phase 6:**  
What distinguishes knowledge that can be documented, transferred, or encoded from knowledge that continues to require ongoing specialist involvement?

**Candidate structural formulations (PROVISIONAL):**

| ID | Formulation | Core structural idea | Falsification implication |
|----|-------------|----------------------|---------------------------|
| C4-a | Codifiable vs tacit split | A portion of process knowledge can be fully documented and transferred; another portion remains tacit, context-dependent, or requires continuous practice and cannot be fully captured in recipes or facility settings. | If all economically relevant process knowledge proves fully codifiable, the residual disappears. |
| C4-b | Frontier knowledge decays into codified knowledge | Newly discovered process improvements begin as specialist-held advantage and gradually become codified and widely available; ongoing discovery is required to stay at the frontier. | If the frontier advances slowly or codification is near-instant, the window of specialist advantage collapses. |
| C4-c | Recipe + setting incompleteness | Documented recipes and facility settings leave residual degrees of freedom whose optimal resolution still depends on specialist judgment under live resource variation. | If recipes + settings fully determine optimal output under all resource conditions, C4-c fails. |
| C4-d | Knowledge ownership vs facility ownership | Even when knowledge can be taught or sold, the independent specialist retains residual advantage through continuous updating and personal experimentation history that is costly for an organisation to keep current across all domains. | If organisations can maintain equivalent continuous updating at low marginal cost, the distinction collapses. |

**Explicit non-choices:**  
No claim that “tacit knowledge” is a protected or necessary design element.  
No numeric rates of codification or decay.  
No new channel beyond the process-knowledge residual already present in C4.

**Re-test implication:**  
Phase 1 (quality parity), Phase 3 (organisational internalisation), and Phase 6 can test whether the chosen formulation still leaves independent specialists with positive ETA on process-sensitive high-value activity. Survival requires that the non-codifiable or frontier portion remains economically meaningful; failure is the expected result if codification is complete and cheap.

---

# 6. Revised Package γ′ Construction Rules

When a targeted re-test is performed, Package γ′ is formed by:

1. Retaining the original γ composition for all shapes not refined above.  
2. Selecting **one or more** candidate formulations from each of the four refinement sets (S4, Info, Dem, C4).  
3. Applying the identical adversarial sequence defined in Comparative Simulation Specification v0.1.  
4. Judging survival exclusively by Phase-6 combined closure under the Phase-6 rule (intermediate positive ETA is not survival).

**Permitted outcomes for any γ′:**  
- SURVIVES  
- FAILS  
- INCONCLUSIVE  

No formulation set may be written or selected in a way that makes SURVIVES the only coherent result.

**Forbidden during re-test construction:**  
- Adding residual channels not already present in Functional Shapes v0.1.  
- Altering Packages α or β.  
- Introducing numeric parameters as design decisions.  
- Treating any candidate formulation as LOCKED or preferred.

---

# 7. Falsification Implications at Architecture Level

| Possible re-test result | Interpretation for Architecture C |
|-------------------------|-----------------------------------|
| At least one coherent γ′ SURVIVES Phase 6 with identifiable load-bearing residuals | Architecture C remains viable under the refined shapes; the surviving residuals become the focus of the human validation gate. |
| All tested γ′ variants FAIL Phase 6 | Architecture C as expressible through the current Functional Shapes (including these refinements) is falsified. Human decision required; do not automatically invent Architecture D. |
| γ′ remains INCONCLUSIVE | Further narrowing is required on the still-blocking relationship(s); or a human ruling is needed on whether the project accepts the relevant structural commitment (e.g., persistent discovery scarcity, non-codifiable process knowledge, residual demand for novelty). |

This preserves the evidence chain: negative evidence from α and β stands; γ is given one controlled chance to become discriminable; the project does not invent new theory to force survival.

---

# 8. Explicit Non-Actions

This refinement does **not**:  
- Choose among S4-a–d, Info-a–d, Dem-a–d, or C4-a–d.  
- Assert that any formulation is necessary for interdependence.  
- Rescue Package γ by construction.  
- Reopen α or β.  
- Authorise numeric tuning.  
- Authorise Provenance & Reputation design.  
- Authorise a new Master GDD pass or Architecture D.

---

# 9. Recommended Immediate Next Step

**Targeted γ′ adversarial re-test** under the Comparative Simulation Specification sequence, using one or more combinations of the candidate formulations above.

The re-test report must again enforce:  
- Phase-6 combined closure as the sole survival criterion;  
- no invention of residuals;  
- explicit SURVIVES / FAILS / INCONCLUSIVE judgment;  
- Architecture-level statement;  
- under-specification report if discrimination remains impossible.

Only after that evidence exists should a human validation gate be convened.

---

**Document Control**

**Authority:** None. Provisional refinement of simulation assumptions only.  
**Status:** PROVISIONAL γ-closure patch. Does not modify Master GDD status tags or Human Rulings.  
**Lineage:** Issued in direct response to the under-specification findings of EIC Comparative Simulation Results v0.1, limited to the four relationships that blocked Package γ.

*End of EIC Functional Shapes Refinement — Targeted γ Closure v0.2*
