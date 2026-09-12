# TCIndustries — EIC Comparative Simulation Specification v0.1

**Document Type:** Adversarial Simulation Specification (Non-Canonical)  
**Version:** 0.1 (post-Functional-Shapes revision)  
**Date:** 2026-08-25  
**Status:** Specification only. Creates no design authority. Does not promote Architecture C or any functional shape.  
**Authority:** None.  

**Controlling References (in order of relevance for this specification):**  
- `TCIndustries_EIC_Provisional_Functional_Shapes_v0.1.md` (hypotheses under test)  
- `TCIndustries_EIC_Human_Rulings_2026-08-25.md` (HD-EIC-01 to HD-EIC-04)  
- `TCIndustries_EIC_Candidate_v0.1.md`  
- `TCIndustries_EIC_Minimum_Simulation_Results_v0.1.md` (prior under-specification finding)  
- `TCIndustries_Master_GDD_v1.1.1_Status_Patch.md`  
- `TCIndustries_Authority_Provenance_Reconciliation_Matrix.md`

**Lineage note:**  
This document supersedes the earlier Comparative Simulation Specification that preceded the Functional Shapes and Human Rulings. The prior minimum simulation was found UNDER-SPECIFIED. This specification is the controlled next step authorised after HD-EIC-01–04 and the completion of Provisional Functional Shapes v0.1.

**Hard constraints:**  
- Do not invent new residual channels to rescue Architecture C.  
- Do not select or invent numeric balance values as design decisions.  
- Do not promote any shape or package to LOCKED.  
- Do not design Provenance & Reputation.  
- The simulation must be capable of producing FAIL results for Architecture C as currently shaped.  
- Falsification is the goal, not defence.

---

# 1. Purpose

Define the smallest controlled experiment that can discriminate among the coherent functional-shape packages defined in Provisional Functional Shapes v0.1.

**What this specification decides:**  
How the hypotheses will be subjected to adversarial pressure.  
What constitutes evidence of survival, failure, or continued under-specification.

**What this specification does not decide:**  
Which functional shapes are correct.  
Whether Architecture C is valid.  
Any mechanical implementation.

**Governing question (the only question that matters):**  
Does there exist a coherent package of functional shapes, drawn from those defined in Provisional Functional Shapes v0.1, under which independent economic participants retain economically meaningful comparative advantage after:

- legitimate multi-account internalisation,
- organisational pooling and capital concentration,
- manufacturing quality parity (no hard specialist-only quality tier),
- NPC baseline substitution, and
- progressive neutralisation of demand for residual channels?

If every tested package fails under this pressure, that is evidence against Architecture C as currently shaped.  
If some packages survive on identifiable residual channels, those channels become candidates for further human scrutiny.  
If the simulation cannot discriminate because critical dimensions remain undefined, the result is INCONCLUSIVE and further shaping or human ruling is required.

---

# 2. Governing Constraints from Human Rulings

These are binding on the experimental design:

| Ruling | Constraint on simulation |
|--------|--------------------------|
| HD-EIC-01 | Validation unit is layered. Character specialisation is tracked but the decisive metric is comparative advantage at the independent economic participant layer (including multi-account and organisational actors). |
| HD-EIC-02 | No hard factory-quality ceiling is permitted. Specialist residual advantage must be tested under quality parity. |
| HD-EIC-03 | Legitimate multi-accounting is an in-scope adversarial condition. No account-level restrictions may be introduced. |
| HD-EIC-04 | Only non-numeric provisional functional shapes from the Functional Shapes document may be used as simulation assumptions. |

**Explicit prohibition:**  
Do not add new residual channels, soft quality edges, or concentration pressures that were not already listed as candidates in Provisional Functional Shapes v0.1 merely to make a package survive.

---

# 3. Core Metrics

Two primary metrics. Exact computational form is an implementation detail of the simulator; the conceptual definitions are fixed.

## 3.1 External Transaction Advantage (ETA)

For a given independent economic participant and a high-value activity:

```
ETA = Expected value obtained by transacting with an external specialist
    − Expected value obtained by full internalisation / self-supply
```

(adjusted for attention, coordination, capital opportunity cost, and risk)

**Interpretation bands (qualitative):**  
- ETA ≤ 0 across high-value activity → independent specialists are economically redundant for that participant.  
- ETA moderately and stably positive → voluntary comparative advantage (target condition).  
- ETA extremely high and unavoidable → risk of coercive dependence (failure relative to PIL-003 lineage).

## 3.2 Internalisation Closure Rate (ICR)

Fraction of high-value production stages or residual channels that a rational multi-account or organisational participant can economically close through internalisation, such that external specialists become optional.

**Interpretation:**  
- High ICR under optimiser pressure → Architecture C fails to preserve independent comparative advantage.  
- Persistently incomplete ICR on identifiable residual channels → those channels are candidates for load-bearing status.

Secondary observational metrics (reported but not primary decision criteria):  
- Specialist market share of high-value goods under each condition.  
- Multi-account / organisation advantage ratio relative to independent single-character specialists.  
- Self-supply rate of high-value activity.

---

# 4. Shape Packages Under Test

The simulation does **not** test individual mechanics in isolation. It tests coherent packages drawn from Provisional Functional Shapes v0.1.

Three packages are defined. They are deliberately contrasting so that survival or failure can be attributed.

### Package α — Soft Pressure + Residual Differentiation (baseline leading candidate)

| Component | Shape selected from Functional Shapes |
|-----------|---------------------------------------|
| Specialisation | S2 (soft capacity / attention / maintenance load) + elements of S5 (opportunity-cost return) |
| Resources | R1 (multi-attribute) + R3 (uncertainty / discovery-linked) |
| Crafting | C1 (experimentation unlock) + C3 (customisation) + C4 (process knowledge) |
| Manufacturing | Mfg1 (scale + consistency, open quality parity) + Mfg3 (throughput vs flexibility) |
| NPC | Strong on volume / baseline quality / cost; weak on customisation, discovery-linked attributes, responsiveness, provenance |
| Demand | Mixed: volume + quality-sensitive + situational + moderate prestige/identity component |

**Intent:** The package that most closely tracks the current Architecture C leading hypothesis under the Functional Shapes.

### Package β — Formal Character Limit + Weaker Residuals

| Component | Shape selected |
|-----------|----------------|
| Specialisation | S1 (formal concurrent mastery limit at character layer) only |
| Resources | R2 (quality tiers + scarcity) |
| Crafting | C2 (attribute utilisation efficiency) primarily |
| Manufacturing | Mfg1 (open quality parity) |
| NPC | Same as α |
| Demand | Primarily volume + quality-sensitive; reduced prestige and customisation weight |

**Intent:** Stress-test whether a pure character-layer formal limit survives multi-account and organisational internalisation. Expected to be weaker on the independent-participant layer.

### Package γ — Asset / Facility Concentration + Strong Information Residuals

| Component | Shape selected |
|-----------|----------------|
| Specialisation | S4 (asset / facility concentration) + S3 (knowledge / information maintenance) |
| Resources | R3 + R4 (interaction / recipe sensitivity) + R5 (temporal / regional) |
| Crafting | C1 + C4 (experimentation and process knowledge) |
| Manufacturing | Mfg3 (throughput vs flexibility) + Mfg4 (input-quality sensitivity) |
| NPC | Same baseline as α |
| Demand | Elevated situational + organisational + discovery-linked demand; moderate prestige |

**Intent:** Test whether concentration pressure located at facilities and information, rather than character skill caps, produces more robust independent-participant advantage.

**Note:** No package may introduce a hard factory-quality ceiling. All operate under quality parity.

---

# 5. Adversarial Test Sequence

Each package is subjected to the same progressive adversarial sequence. The sequence is designed to close residual channels systematically.

### Phase 0 — Baseline
- All residual channels available as defined by the package.
- Independent specialists, multi-account participants, small organisations, and manufacturing agents present.
- Demand includes the residual-sensitive components defined by the package.
- Measure ETA and ICR.

### Phase 1 — Factory Quality Parity
- Manufacturing path can match specialist measured quality given equivalent high-attribute inputs and capital.
- No hard specialist-only tier.
- Measure whether residual channels (customisation, experimentation-derived processes, responsiveness, information, etc.) still generate positive ETA for independent specialists.

### Phase 2 — Multi-Account Internalisation
- Introduce rational multi-account decision units that can allocate characters across the specialisation domains defined by the package.
- Character-layer constraints (if any) remain, but the participant can internalise multiple domains.
- Measure residual ETA for true independent (non-multi-account) specialists and the ICR of the multi-account units.

### Phase 3 — Organisational Internalisation + Capital Concentration
- Organisations can pool specialists, own facilities, concentrate capital, and contract or employ production capacity.
- Logistics and intermediate-goods markets are available.
- Measure whether independent specialists retain any high-value comparative advantage or whether ICR approaches closure.

### Phase 4 — NPC Substitution Pressure
- NPCs supply baseline volume and quality at competitive cost.
- Progressively strengthen NPC capability on the residual dimensions the package claims are player-favouring (customisation, responsiveness, discovery-linked quality, etc.), one dimension at a time and then in combination.
- Identify which residual dimensions, if any, remain load-bearing against NPC competition.

### Phase 5 — Demand Neutralisation
- Progressively reduce or remove demand weight on prestige, customisation, responsiveness, provenance, and discovery-linked attributes.
- Observe at which point (if any) independent specialist ETA collapses to ≤ 0.
- This phase determines whether the claimed residual channels are actually load-bearing or merely asserted.

### Phase 6 — Combined Closure
- All prior pressures applied simultaneously under rational optimiser agents.
- Final measurement of ETA distribution and ICR for independent economic participants.

**Critical instruction:**  
At no point may the simulation operator invent an additional residual channel that was not already present in the package definition drawn from Provisional Functional Shapes v0.1.

---

# 6. Agent Set (Minimal)

| Agent | Role in experiment |
|-------|--------------------|
| Independent focused specialist | Primary unit whose comparative advantage is under test |
| Independent generalist | Control for opportunity-cost of specialisation |
| Multi-account decision unit | Adversarial internaliser (HD-EIC-03) |
| Small organisation | Adversarial pooler of specialists + capital |
| Manufacturing / industrial agent | Quality-parity and scale competitor |
| NPC baseline provider | Substitution floor |
| Rational economic optimiser (meta) | Searches for dominant strategies under each phase |
| Consumer / demand population | Tunable preference weights for residual channels |

Large multi-organisation ecosystems may be deferred if Phase 3 already produces clear FAIL results.

---

# 7. Decision Rules

For each package, after the full adversarial sequence:

| Outcome | Condition | Interpretation |
|---------|-----------|----------------|
| **SURVIVES** | Independent economic participants retain stably positive ETA on at least one residual channel that demand continues to value, even under combined multi-account, organisational, manufacturing-parity, NPC, and partial demand-neutralisation pressure | Package remains a viable candidate for Architecture C; identify the load-bearing residual(s) |
| **FAILS** | ETA for independent specialists collapses to ≤ 0 across high-value activity, or ICR approaches full closure under optimiser pressure | Package is insufficient; Architecture C as expressed by this package does not preserve independent comparative advantage |
| **INCONCLUSIVE** | Results are dominated by undefined parameters, or discrimination is impossible because a critical functional relationship was left underspecified in the package | Further shaping or human ruling required before the package can be evaluated |

**Architecture-level judgment:**  
- If all three packages FAIL → strong evidence against Architecture C under the current Functional Shapes.  
- If one or more SURVIVE on identifiable residuals → those residuals become the focus of the next human gate.  
- If the majority of results are INCONCLUSIVE → the Functional Shapes remain under-specified for discrimination; return to shaping or request human ruling on the blocking dimensions.

**Do not:**  
Rescue a failing package by adding new mechanisms.  
Treat survival of a residual under light pressure as success if it collapses under combined adversarial pressure.

---

# 8. Required Simulation Outputs

The results report must contain, for each package and each phase:

1. ETA distribution (or qualitative band) for independent specialists.  
2. ICR for multi-account and organisational agents.  
3. Identification of which residual channel(s), if any, continued to generate positive ETA.  
4. Point in the adversarial sequence at which (if ever) independent specialist advantage collapsed.  
5. Explicit PASS / FAIL / INCONCLUSIVE judgment per package with supporting metric summary.  
6. Statement of any dimensions that prevented discrimination (under-specification report).

**Mandatory final statements:**  
- “Under the tested packages, independent economic participants [did / did not] retain meaningful comparative advantage after legitimate internalisation and quality parity.”  
- “The load-bearing residual channel(s), if any, were: …”  
- “Architecture C as currently shaped is [supported as still viable / weakened / not supported] by this evidence.”  
- “Unresolved dimensions that still prevent full discrimination: …”

---

# 9. Explicit Out of Scope

- Numeric tuning of any parameter as a design decision.  
- Full reputation or provenance system dynamics (deferred).  
- Exact skill trees, point costs, or formal budget numbers.  
- Combat, cities, governance, currency design.  
- Any mechanism or residual channel not already listed as a candidate in Provisional Functional Shapes v0.1.  
- Promotion of any package or residual to LOCKED status.

---

# 10. Relationship to Prior Work

| Prior artefact | Relationship |
|----------------|--------------|
| EIC-Candidate v0.1 | Source of original M1–M8 mechanisms; still the conceptual parent of Architecture C |
| Falsification Pass | Identified the attack surfaces now being tested under shaped conditions |
| Minimum Simulation Results | Demonstrated under-specification; this specification exists to remove that under-specification |
| Human Rulings HD-EIC-01–04 | Binding constraints on unit of validation, quality parity, multi-account stance, and permission for provisional shapes |
| Provisional Functional Shapes v0.1 | Direct source of the shape packages and residual channels under test |

---

# 11. Success Criterion for This Specification

This specification is successful if a competent simulation (or rigorous structural analysis) run against it can produce a clear SURVIVES / FAILS / INCONCLUSIVE judgment for each package and an Architecture-level statement of the form required in §8.

It is unsuccessful if the experiment remains unable to discriminate because the packages are still too loosely defined. In that case the correct next action is further Functional Shapes refinement or a human ruling on the blocking open decisions (see Functional Shapes §14), not invention of new residuals.

---

# Decision Gate (for the simulation results that will follow)

| Question | Classification after this specification |
|----------|----------------------------------------|
| Can Architecture C be subjected to controlled adversarial closure? | Yes — this specification defines the experiment |
| Which residual channels are load-bearing? | Requires execution of the simulation |
| Does any coherent package survive multi-account + organisation + quality parity + NPC + demand neutralisation? | Requires execution of the simulation |
| Numeric values | Still forbidden |
| Provenance & Reputation | Still deferred |
| Promotion of Architecture C | Forbidden until after human validation of evidence |

**Smallest next task after this specification:**  
Execute (or rigorously analyse) the comparative simulation defined herein and produce the results report with the mandatory statements in §8.

Only after that evidence exists should a human decision gate be convened on the survival or failure of Architecture C under the Functional Shapes.

---

**Document Control**

**Authority:** None. Simulation specification and experimental design only.  
**Status:** PROVISIONAL specification for adversarial evidence generation.  
**No design is authorised or promoted by this document.**  

*End of EIC Comparative Simulation Specification v0.1*
