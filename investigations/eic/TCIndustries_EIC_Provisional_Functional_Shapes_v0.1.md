# TCIndustries — EIC Provisional Functional Shapes v0.1

**Document Type:** Provisional Functional Model (Simulation Assumptions Only)  
**Version:** 0.1  
**Date:** 2026-08-25  
**Status:** PROVISIONAL — NOT HUMAN-LOCKED  
**Authority:** None. This document creates no canonical design. All content is PROPOSED / simulation assumption unless explicitly restating a prior HUMAN-LOCKED ruling.  
**Controlling References:**  
- `TCIndustries_EIC_Human_Rulings_2026-08-25.md` (HD-EIC-01 to HD-EIC-04)  
- `TCIndustries_Master_GDD_v1.1.1_Status_Patch.md`  
- `TCIndustries_Authority_Provenance_Reconciliation_Matrix.md`  
- `TCIndustries_EIC_Candidate_v0.1.md`  
- `TCIndustries_EIC_Minimum_Simulation_Results_v0.1.md`  
- `TCIndustries_Economic_Interdependence_Core_Design_Investigation.md`

**Purpose of this document:**  
Convert the human-approved EIC direction (Architecture C as leading hypothesis + HD-EIC-01–04) into the minimum set of non-numeric structural relationships required to make the Economic Interdependence Core testable.  

This document does **not**:  
- rewrite or consolidate the Master GDD;  
- promote any mechanism to LOCKED or HUMAN-LOCKED;  
- select a specialisation-budget implementation;  
- introduce numeric balance values;  
- design Provenance & Reputation;  
- claim the EIC is solved;  
- collapse character specialisation into economic-participant independence.

It exists solely to enable the next adversarial validation or discriminating simulation step.

---

# 1. Purpose and Status

**What this document is deciding:**  
Nothing that is binding on the game. It defines provisional functional shapes (relationships among inputs, transformations, outputs, comparative advantage, dependency, substitution, and failure modes) so that simulation or structural analysis can discriminate among alternatives.

**What this document is not deciding:**  
- Any concrete skill system, budget, or mastery limit.  
- Any concrete resource attribute table.  
- Any concrete factory quality formula.  
- Any multi-account policy beyond the validation stance already ruled.  
- Any demand-generation formula.  
- Whether Architecture C survives contact with shaped evidence.

**Status discipline:**  
Every candidate shape, invariant, and failure condition below is labelled PROVISIONAL.  
Human rulings already recorded remain the only HUMAN-LOCKED elements.

---

# 2. Governing Principles

Only the following are treated as binding for this model (restated, not expanded):

**From Master GDD v1.1.1 (HUMAN-LOCKED principles relevant to EIC):**  
- Interdependence creates opportunity rather than coercive inconvenience (PIL-003 lineage).  
- Meaningful non-combat careers and player-driven economy.  
- One character should not efficiently perform every economically important role (design commitment, not yet a locked mechanism).  
- NPCs must not make meaningful player economic roles irrelevant.

**From HD-EIC-01 (HUMAN-LOCKED — unit of validation):**  
Character specialisation remains an intended design concern.  
EIC validation concerns comparative advantage among independent economic participants, including multi-account and organisational participants.  
Character specialisation and economic-participant comparative advantage are distinct layers and must not be collapsed.

**From HD-EIC-02 (HUMAN-LOCKED — principle only):**  
Specialists must retain economically meaningful comparative advantages that can arise from expertise, information, experimentation, discovery, responsiveness, customisation, production economics, reputation, provenance, or related player-created value.  
No permanent hard factory-quality ceiling is authorised.

**From HD-EIC-03 (HUMAN-LOCKED — validation stance):**  
Legitimate multi-accounting is an in-scope adversarial condition.  
Multi-account internalisation must not trivially eliminate economically meaningful independent participation.  
No account-level restriction is authorised by this ruling.

**From HD-EIC-04 (HUMAN-LOCKED — methodological permission):**  
Provisional non-numeric functional shapes may be defined for simulation.  
These remain PROPOSED / simulation assumptions.

No new LOCKED principles are created by this document.

---

# 3. EIC Problem Definition

**Operational problem being tested:**  
Can the system produce sustained economic reasons for independent economic participants to rely upon one another for high-value activity, such that specialised independent participation remains a rational strategy even when multi-accounting, organisational pooling, capital concentration, automation, and NPC substitution are available?

**What “economic interdependence” means in this model (precise distinctions):**

| Form | Definition | Status for EIC |
|------|------------|----------------|
| Transactional interdependence | Players trade because it is convenient or lower effort | Necessary but insufficient; can be fragile |
| Comparative interdependence | Another participant possesses a materially superior capability on a relevant dimension | Necessary for sustained specialist relevance |
| Structural interdependence | A participant cannot efficiently perform an important role without another participant’s output or capability | Desirable when opportunity-creating; dangerous when coercive |
| Economic interdependence | The system produces sustained expected-utility reasons for independent participants to specialise and transact rather than fully internalise | The target condition |

**Healthy vs artificial:**  
Healthy interdependence expands the set of high-value activities a player can access through others.  
Artificial interdependence merely raises the cost of self-sufficiency without expanding opportunity. The latter is treated as a failure mode under the existing pillar.

**What is not being tested here:**  
Whether characters feel “specialised” in a narrative sense. That is a character-layer concern. The EIC tests whether independent economic participants retain comparative advantage.

---

# 4. Economic Participant Model

The model explicitly distinguishes layers. Collapsing them is forbidden by HD-EIC-01.

| Layer | Description | Capabilities that can be combined | Constraints that may apply | Comparative advantage possible? | Independence possible? |
|-------|-------------|-----------------------------------|----------------------------|---------------------------------|-------------------------|
| 1. Character | Single in-world persona | Skills, knowledge, location, personal reputation, personal experimentation history | Character-level attention, skill progression shape (unresolved), personal capacity | Yes (specialist vs generalist) | Partial; limited by character constraints |
| 2. Account / multi-character participant | One human decision-maker controlling multiple characters (same or linked accounts) | Union of character capabilities + coordinated logistics | Account-level time, capital, coordination overhead (unresolved magnitude) | Yes (internal specialisation + external trade) | High; can approximate organisational behaviour |
| 3. Independent player economic participant | Decision unit treated as economically autonomous for validation | Whatever the player can command through characters/accounts | Whatever the system actually constrains at this layer | Yes (the primary validation unit) | By definition the unit under test |
| 4. Organisation | Formal or informal group of players pooling capacity | Specialists + capital + facilities + shared information + contracts | Internal governance, free-rider, trust, coordination costs | Yes (scale, specialisation depth, capital) | High; can internalise many roles |
| 5. NPC / system provider | Non-player supply of goods, services, or capacity | Baseline production, information, transport, etc. | System-defined limits on quality, cost, availability, customisation, responsiveness | Limited by design of substitution rules | N/A (not a player participant) |

**Why layers cannot be collapsed:**  
Preventing one character from mastering everything does not automatically prevent one human from mastering everything economically via multiple characters or an organisation.  
The EIC must therefore validate comparative advantage at the independent-participant layer, not merely at the character layer.

**Key implication for functional shapes:**  
Any shape that only constrains character-level concurrent mastery, without residual pressure or opportunity costs that survive multi-character coordination, is under-specified relative to HD-EIC-01 and HD-EIC-03.

---

# 5. Coupled EIC Functional Model

The EIC is not a linear pipeline. It is a set of mutually constraining and reinforcing relationships.

```
                    ┌─────────────────────────────────────┐
                    │     Independent Economic Participant │
                    │  (character / multi-char / org)      │
                    └─────────────────────────────────────┘
                                      │
          ┌───────────────────────────┼───────────────────────────┐
          ▼                           ▼                           ▼
   Specialisation              Resource Quality              Demand
   (capability                 / Attributes                  (existence,
    concentration,              (discovery,                   volume,
    opportunity cost,           scarcity,                     quality-
    information)                attributes,                   sensitivity,
                                location,                     recurrence,
                                uncertainty)                  prestige)
          │                           │                           │
          └─────────────┬─────────────┴─────────────┬─────────────┘
                        ▼                           ▼
                   Crafting                     Manufacturing
                   (experimentation,            (scale, repeatability,
                    differentiation,             consistency, throughput,
                    customisation,               flexibility,
                    quality channels)            residual specialist
                                                 channels)
                        │                           │
                        └─────────────┬─────────────┘
                                      ▼
                              NPC Substitution
                              (quality, cost,
                               availability,
                               customisation,
                               responsiveness,
                               locality)
                                      │
                                      ▼
                         Player Comparative Advantage
                         (sustained or eliminated)
```

**Feedback loops of primary interest:**

1. **Specialisation ↔ Resource quality ↔ Crafting**  
   Higher resource quality or rarer attributes increase the return to specialised experimentation and information. Specialisation that improves discovery or processing efficiency raises the effective value of scarce resources. Circular: better specialists make better use of quality resources, which makes specialisation more valuable.

2. **Crafting ↔ Manufacturing**  
   Manufacturing provides scale and consistency; crafting provides differentiation, experimentation, and residual quality or customisation channels. If manufacturing can fully replicate the high-value outputs of specialised crafting, the comparative advantage of the specialist collapses. If manufacturing is strictly inferior on all high-value dimensions, specialists may become bottlenecks rather than partners.

3. **Manufacturing / Crafting ↔ NPC substitution**  
   NPCs set a floor. If NPCs can supply adequate volume and quality at competitive cost, player production loses economic meaning. If NPCs are constrained on the dimensions that specialists excel at (customisation, responsiveness, discovery-linked quality, reputation-linked value), player roles remain viable.

4. **Demand ↔ all production channels**  
   Demand that is purely volume-driven and quality-insensitive favours the lowest-cost producer (often manufacturing or NPCs). Demand that is quality-sensitive, prestige-sensitive, situational, or customisation-sensitive preserves specialist channels. Demand itself is shaped by what players can achieve; higher player capability can expand the set of desired goods.

5. **Specialisation ↔ Organisation / multi-account**  
   Organisations and multi-account participants can internalise specialisation. The functional shape of concentration pressure must therefore be evaluated after internalisation, not before.

**Circular dependencies that must be acknowledged:**  
- Resource quality value depends on the existence of participants who can exploit it.  
- Specialist comparative value depends on demand for the dimensions specialists control.  
- Demand for high-quality / differentiated goods depends on the existence of reliable specialist supply.  
These loops are not bugs; they are the interdependence the system is trying to stabilise. They must be shaped carefully so the equilibrium is not “everyone is a generalist” or “only the largest organisation matters.”

---

# 6. Functional Shape Candidates

For each major subsystem, alternative provisional shapes are listed. None is selected. All are non-numeric.

## 6.1 Specialisation / Capability Concentration (Character Layer vs Participant Layer)

**Functional purpose:**  
Make simultaneous elite performance across the full set of high-value economic domains costly, incomplete, or lower-return at the relevant validation layer, while still allowing meaningful multi-role play.

**Inputs:**  
Character progression state, attention/time allocation, information state, facility ownership, organisational membership.

**Transformation:**  
Maps capability investment into effective performance; applies some form of concentration pressure or opportunity cost.

**Outputs:**  
Effective capability profile of a character and of an independent economic participant.

**Comparative advantage loci:**  
Depth in a domain, information advantage, experimentation history, residual quality channels, responsiveness, customisation.

**Dependency:**  
Requires that high-value domains exist and that demand values the dimensions specialists control.

**Substitution:**  
Another specialist, a generalist (lower performance), multi-character internalisation, organisation, purchased intermediates, NPCs, automation.

**Failure mode:**  
Pressure that only exists at character layer is fully internalised by multi-account or organisation → independent specialists become optional.

**Candidate shapes (PROVISIONAL):**

| ID | Shape | Core idea | Layer primarily constrained | Risk |
|----|-------|-----------|-----------------------------|------|
| S1 | Formal concurrent mastery limit | Hard or soft ceiling on simultaneous elite domains per character | Character | Easily bypassed by multi-character; may feel coercive |
| S2 | Soft capacity / attention / maintenance load | Maintaining elite performance consumes ongoing scarce attention or upkeep | Character + participant (if upkeep is real) | Magnitude of load is critical; too weak → irrelevant; too strong → coercive |
| S3 | Knowledge / information maintenance | Specialised knowledge decays or requires active practice / discovery | Character | Can be shared or recorded; information markets may form |
| S4 | Asset / facility concentration | Elite production requires specialised, non-trivial facilities whose opportunity cost is high | Participant / organisation | Capital concentration may dominate; facilities can be owned by organisations |
| S5 | Opportunity-cost / comparative-return shape | No hard limit; returns diminish or opportunity cost rises when spreading | Both | Softest; may not produce structural interdependence |

**Note:** HD-EIC-01 requires that any shape be evaluated for its residual effect after multi-character and organisational internalisation. A pure character-level S1 is therefore incomplete by itself.

## 6.2 Resource Quality / Attribute Model

**Functional purpose:**  
Create differentiated inputs whose value is not purely fungible volume, so that discovery, location, processing skill, and experimentation remain meaningful.

**Inputs:**  
World state, discovery actions, harvesting skill/information, location, temporal conditions.

**Transformation:**  
Produces resources possessing attributes, quality distributions, scarcity, or interaction effects.

**Outputs:**  
Inputs into crafting and manufacturing that are not perfect substitutes.

**Comparative advantage:**  
Better discovery, better attribute targeting, better processing yields, better information about where/when quality appears.

**Dependency:**  
Requires that crafting/manufacturing systems can actually express resource differences in final goods.

**Substitution:**  
Lower-quality substitutes, imported resources, synthetic/NPC sources, stockpiling.

**Failure mode:**  
All resources become pure volume commodities → specialist discovery and processing lose value; lowest-cost harvester dominates.

**Candidate shapes (PROVISIONAL):**

| ID | Shape | Core idea |
|----|-------|-----------|
| R1 | Multi-attribute vector | Resources carry multiple continuous or discrete attributes that interact in recipes |
| R2 | Quality tiers + scarcity | Discrete tiers with increasing scarcity / location dependence |
| R3 | Uncertainty / discovery-linked | Highest attributes require ongoing discovery or experimentation; static maps are incomplete |
| R4 | Interaction / recipe sensitivity | Certain attribute combinations unlock or strongly improve specific high-value outputs |
| R5 | Temporal / regional availability | Quality or attributes vary by time and region, rewarding local knowledge and logistics |

Combinations are possible. Pure volume commodities with no attribute expression are treated as a failure shape for interdependence.

## 6.3 Crafting Differentiation and Experimentation

**Functional purpose:**  
Provide channels through which specialised skill, information, and experimentation produce outputs that are not fully replicable by pure scale manufacturing or NPC baselines.

**Inputs:**  
Resources (with attributes), character skill/knowledge, experimentation history, tools, recipes, time.

**Transformation:**  
Converts inputs into finished or intermediate goods; may generate new knowledge or variants.

**Outputs:**  
Goods with quality, customisation, uniqueness, or residual properties; possible new recipes or process knowledge.

**Comparative advantage:**  
Higher success rates on difficult experiments, better attribute utilisation, customisation, responsiveness to specific demand, reputation-linked provenance.

**Dependency:**  
Requires resource differentiation and demand that values the residual channels.

**Substitution:**  
Manufacturing (if it can match residuals), other crafters, purchased goods, NPCs.

**Failure mode:**  
Crafting becomes a pure efficiency race with no residual differentiation → manufacturing or capital dominates.

**Candidate shapes (PROVISIONAL):**

| ID | Shape | Core idea |
|----|-------|-----------|
| C1 | Experimentation unlock / variance | Specialists can explore recipe space that generalists or factories cannot efficiently reach |
| C2 | Attribute utilisation efficiency | Specialists extract more value from high-attribute resources |
| C3 | Customisation / bespoke channel | Ability to produce goods tuned to specific buyer needs |
| C4 | Process knowledge / information advantage | Specialists possess or generate process knowledge that improves yields or unlocks options |
| C5 | Residual quality / consistency under difficulty | Specialists maintain higher effective quality or lower variance on demanding recipes |

Per HD-EIC-02, none of these may rely on a permanent hard factory-quality ceiling as the sole residual.

## 6.4 Manufacturing Constraints

**Functional purpose:**  
Provide scale, repeatability, and consistency while preserving economically meaningful residual channels for skilled participants.

**Inputs:**  
Intermediate goods, facilities, capital, labour (player or contracted), recipes, process settings.

**Transformation:**  
High-throughput conversion of inputs into standardised or semi-standardised outputs.

**Outputs:**  
Volume goods, components, or finished items with defined consistency and cost structure.

**Comparative advantage (for manufacturing operators):**  
Scale economies, logistics, capital efficiency, facility location, process optimisation.

**Residual specialist channels (must remain possible):**  
Customisation, discovery-linked quality, responsiveness, experimentation-derived processes, reputation/provenance value, or production economics that still favour specialists on some margin.

**Dependency:**  
Requires that demand exists for both volume and differentiated goods.

**Substitution:**  
Other manufacturers, specialists supplying high-end intermediates, NPCs, imports.

**Failure mode:**  
Manufacturing fully replicates all high-value specialist outputs at lower cost → specialists become optional suppliers of intermediates only, or irrelevant.

**Candidate shapes (PROVISIONAL):**

| ID | Shape | Core idea |
|----|-------|-----------|
| Mfg1 | Scale + consistency, open quality parity | Factories can match peak quality but specialists retain other channels (customisation, info, responsiveness, provenance) |
| Mfg2 | Scale + consistency with soft residual quality advantage | Specialists retain a probabilistic or situational quality edge that is not a hard ceiling |
| Mfg3 | Throughput vs flexibility trade-off | Factories excel at volume of known recipes; specialists excel at adaptation and new process discovery |
| Mfg4 | Input-quality sensitivity | Factory output quality is more tightly coupled to input quality; specialists can compensate for or better exploit inputs |

No shape may re-introduce a permanent hard “factories can never reach tier X” rule as a required mechanism.

## 6.5 NPC Substitution

**Functional purpose:**  
Provide a baseline economy floor without rendering player economic roles irrelevant.

**Inputs:**  
System parameters defining NPC capability on each relevant dimension.

**Transformation:**  
NPCs supply goods/services under defined constraints.

**Outputs:**  
Baseline availability that sets opportunity cost for player production.

**Comparative advantage (player vs NPC):**  
Must exist on at least some dimensions that matter to demand.

**Failure mode:**  
NPCs dominate on quality, cost, availability, customisation, and responsiveness simultaneously → player production collapses to niche or zero.

**Candidate dimensions (PROVISIONAL — treat as variables):**

| Dimension | Player-favouring setting | NPC-favouring setting | Notes |
|-----------|--------------------------|-----------------------|-------|
| Quality / attribute expression | Players can exceed NPC baseline on residual channels | NPCs match or exceed | Must not rely on hard ceiling alone |
| Cost | Players can be competitive after specialisation | NPCs always cheaper | Cost alone is fragile |
| Availability / volume | Players control scarce high-end supply | NPCs unlimited | Scarcity must be real |
| Customisation | Players can tailor | NPCs standardised only | Strong residual candidate |
| Responsiveness / locality | Players can react to local/situational demand | NPCs slow or global | Strong residual candidate |
| Reliability | Variable | High | Can cut either way |
| Information / discovery | Players generate new knowledge | NPCs static | Aligns with discovery pillar |
| Reputation / provenance | Player-created identity value | NPCs lack | Aligns with identity pillar |

**Useful combinations:**  
NPCs strong on volume + baseline quality + cost; weak on customisation, discovery-linked attributes, responsiveness, and player reputation. This preserves specialist roles without making NPCs artificially incompetent at ordinary goods.

## 6.6 Demand

**Functional purpose:**  
Generate sustained reasons for participants to value differentiated production rather than only the cheapest adequate output.

**Inputs:**  
Player goals, organisational goals, combat needs, non-combat needs, prestige, replacement cycles, situational events.

**Transformation:**  
Maps needs into willingness to pay for volume, quality, customisation, provenance, or other attributes.

**Outputs:**  
Demand signals that reward or ignore specialist channels.

**Failure mode:**  
All demand collapses to lowest-cost adequate goods → specialists and high-attribute resources lose economic meaning.

**Candidate demand types (PROVISIONAL):**

| Type | Description | Supports specialist residual? |
|------|-------------|-------------------------------|
| Existence | Basic need for a category of good | Weak |
| Volume | Quantity pressure | Favours manufacturing / NPCs |
| Quality-sensitive | Higher attributes or performance matter | Strong if attributes express |
| Replacement / recurring | Ongoing consumption | Supports sustained production |
| Situational | Event- or location-driven spikes | Favours responsiveness |
| Prestige / identity | Desire for recognised, provenance-linked, or unique goods | Strong residual candidate |
| Organisational | Group-scale needs (facilities, logistics, defence, etc.) | Can support both volume and specialist |
| Combat-driven | Performance under risk | Quality and customisation sensitive |
| Non-combat career | Tools, housing, services, entertainment, etc. | Broad support for diverse specialists |

**Critical question the shapes must answer:**  
If players become better at production, what continues to make differentiated output valuable rather than merely “good enough”?  
Possible answers (not selected): prestige/identity value, situational optimisation, organisational requirements that cannot be met by generic supply, discovery of new desirable attribute combinations, or social signalling. Pure arbitrary scarcity is insufficient and risks coercive feel.

---

# 7. Cross-System Interaction

**Reinforcing combinations (examples):**  
- R3/R4 (discovery-linked or interaction-sensitive resources) + C1/C4 (experimentation / process knowledge) + demand that values the resulting differentiation → strong specialist comparative advantage.  
- Mfg3 (throughput vs flexibility) + C3 (customisation) + situational/organisational demand → manufacturing and specialists become complements rather than substitutes.  
- NPC weak on customisation + responsiveness + provenance + player strong on those dimensions → independent specialists remain relevant even against capitalised organisations that use NPCs for volume.

**Undermining combinations (examples):**  
- Pure volume resources (no attributes) + manufacturing with full quality parity + quality-insensitive demand + strong NPC volume → specialists collapse.  
- Character-only concentration pressure (S1) + easy multi-character coordination + organisational facility ownership → independent single-character specialists become optional.  
- Hard factory quality ceiling (forbidden) would create artificial interdependence; the model must not rely on it.

**Key interaction under HD-EIC-01/03:**  
Any specialisation shape must be stress-tested after the participant has internalised multiple characters or formed an organisation. Shapes that only constrain the character layer are incomplete.

---

# 8. Comparative Advantage Model

Comparative advantage is defined relationally, not as an absolute skill number.

**Kinds of advantage that can exist without numeric implementation:**

1. **Exclusive or near-exclusive capability** — ability to produce a good or effect others cannot.  
2. **Efficiency** — higher output per input or per time.  
3. **Quality / attribute ceiling or distribution** — better statistical or peak outcomes (must not be a permanent hard factory lock).  
4. **Information / knowledge** — knowing where, when, how, or what combinations work.  
5. **Experimentation / discovery rate** — faster or more reliable exploration of design space.  
6. **Customisation / fit** — ability to match specific buyer requirements.  
7. **Responsiveness / locality** — speed or geographic advantage.  
8. **Opportunity cost structure** — lower cost of performing the role because of prior investment or specialisation.  
9. **Capacity / scale** — ability to supply volume (more natural to manufacturing/organisations).  
10. **Reputation / provenance / identity value** — buyer preference for the specific participant’s output.  
11. **Reliability under variance** — lower failure or more consistent high-end results.

**Provisional claim (testable):**  
For independent economic participants to remain relevant, at least some high-value demand must be sensitive to one or more of the residual channels that specialists (or specialist organisations) control more effectively than pure capital + NPC + generalist combinations.

---

# 9. Independence / Interdependence Conditions

**Conditions under which independent participants remain economically relevant (PROVISIONAL):**

1. There exist high-value activities whose best expected returns require capabilities that are costly for any single participant to fully internalise across all domains.  
2. Residual channels exist that manufacturing and NPCs do not fully close.  
3. Demand continues to value those residual channels even as overall production capability rises.  
4. Multi-account and organisational internalisation, while powerful, still leaves external comparative advantage on some margins (information boundaries, trust, local knowledge, specialisation depth, reputation portability limits, etc.).  
5. The cost of full internalisation (attention, coordination, capital opportunity cost, risk) is non-trivial relative to the gains from specialisation and trade.

**Conditions that would indicate failure of the interdependence goal:**  
- Independent specialists are systematically dominated by multi-account or organisational participants who simply replicate the same specialisations internally.  
- All high-value demand is satisfied by the lowest-cost adequate supplier (manufacturing or NPC).  
- The only remaining player roles are capital ownership and logistics, with no meaningful production specialisation.

---

# 10. Adversarial Cases

These are first-class test cases, not edge cases.

| Case | Description | Question the functional shapes must answer |
|------|-------------|-------------------------------------------|
| Multi-character (same account) | One player runs multiple specialised characters | Does concentration pressure or residual advantage survive internal coordination? |
| Multi-account | One player runs multiple accounts | Same as above; no account-level restriction is authorised |
| Organisation pooling specialists | Group assigns characters to roles | Can the organisation fully internalise high-value production, or do external specialists still offer residual value? |
| Organisation owns production capacity | Capital + facilities concentrated | Does capital concentration eliminate independent production advantage? |
| Highly capitalised buyer | Participant buys all intermediates and finishes in-house or via contracts | Is there still a reason to deal with independent specialists rather than pure vertical integration? |
| Organisation employs / contracts specialists | Labour market for specialists | Does this preserve or destroy independent comparative advantage? (Can be healthy interdependence if specialists retain bargaining power) |
| NPC replacement of weak roles | NPCs cover low-skill or high-volume roles | Do residual high-value roles remain for players? |
| Automation of repeated skilled actions | Systems reduce the skill content of repeated tasks | Does automation close residual channels or only the repetitive floor? |
| Generalist participant | Player spreads effort widely | Is the generalist systematically disadvantaged on high-value activity relative to specialists + trade? |

**Provisional stance required by rulings:**  
The system is not required to equalise multi-account and independent-player economics. It is required that multi-account internalisation does not *trivially* eliminate economically meaningful independent participation.

---

# 11. Provisional Invariants

All are PROVISIONAL and intended for later testing. None are LOCKED.

| ID | Provisional Invariant | Notes |
|----|----------------------|-------|
| PI-1 | At least one residual channel of specialist comparative advantage must remain economically meaningful after manufacturing quality parity and legitimate multi-account/organisational internalisation | Directly supports HD-EIC-02 and HD-EIC-03 |
| PI-2 | Resource differentiation (attributes, scarcity, discovery linkage, or interaction effects) must be expressible in final goods that demand can value | Without this, specialist processing collapses |
| PI-3 | NPC substitution must leave at least one high-value dimension (customisation, responsiveness, discovery-linked quality, provenance, or equivalent) on which independent players can systematically outperform | Prevents NPC dominance |
| PI-4 | Demand must contain a non-zero component that is sensitive to residual specialist channels even as average production capability rises | Prevents pure cost-minimisation equilibrium |
| PI-5 | Character-level specialisation pressure, whatever its form, must be evaluated for residual effect at the independent economic participant layer | Enforced by HD-EIC-01 |
| PI-6 | Interdependence generated by the shapes must expand opportunity sets rather than merely raise the cost of self-sufficiency | Aligns with PIL-003 lineage |

---

# 12. Failure Conditions

Observable conditions that would indicate the current EIC functional shape has failed (PROVISIONAL):

1. **Trivial internalisation:** After multi-account or organisational coordination, independent specialists have no remaining high-value comparative advantage on any residual channel.  
2. **Manufacturing closure:** Manufacturing (with or without capital concentration) fully closes all residual specialist channels that demand values.  
3. **NPC dominance:** NPCs supply adequate quality, volume, customisation, and responsiveness such that player production is optional for high-value activity.  
4. **Demand collapse to adequacy:** All economically significant demand becomes satisfied by lowest-cost adequate goods; differentiated production loses expected-utility justification.  
5. **Coercive interdependence:** The only way to access high-value activity is through forced reliance that does not expand the player’s opportunity set.  
6. **Generalist dominance:** Spreading capability widely is systematically equal or superior to specialisation + trade for high-value returns.  
7. **Under-specification persistence:** After shaping, simulation still cannot discriminate success from failure because critical functions remain undefined.

---

# 13. Evidence Requirements

**Claims that can be evaluated structurally (now):**  
- Whether a given combination of shapes is internally consistent.  
- Whether a shape only constrains the character layer or also the participant layer.  
- Whether a residual channel is logically closed by manufacturing or NPC assumptions.  
- Whether demand types exist that can value the residual channels.

**Claims that require comparative simulation:**  
- Whether residual specialist advantage survives multi-account and organisational internalisation under specific shape combinations.  
- Whether demand sensitivity to residual channels is sufficient to sustain specialist participation.  
- Relative strength of reinforcing vs undermining interactions.

**Claims that require prototype or empirical evidence:**  
- Actual player behaviour under opportunity-cost vs formal-limit specialisation shapes.  
- Whether information/knowledge advantages remain scarce once markets and organisations form.  
- Magnitude of coordination and attention costs in multi-character play (currently unresolved).  
- Realised value of provenance/reputation channels (out of scope for this document; future system).

**Claims currently impossible to discriminate without further shaping or evidence:**  
- Exact functional form of concentration pressure that survives multi-account (OQ-001 lineage remains open).  
- Which residual channels are load-bearing in practice.  
- Numeric thresholds (explicitly forbidden at this stage).

---

# 14. Open Decisions

The following remain unresolved and are explicitly flagged:

| ID | Open Decision | Status |
|----|---------------|--------|
| OQ-EIC-FS-01 | Which specialisation shape (or combination) is used, and at which layers | UNRESOLVED — REQUIRES HUMAN RULING or comparative evidence |
| OQ-EIC-FS-02 | Precise residual channels that remain after manufacturing parity | UNRESOLVED — REQUIRES EVIDENCE |
| OQ-EIC-FS-03 | Resource attribute model details | UNRESOLVED — REQUIRES EVIDENCE / later design |
| OQ-EIC-FS-04 | Strength and form of NPC constraints on residual dimensions | UNRESOLVED — REQUIRES HUMAN RULING or evidence |
| OQ-EIC-FS-05 | Demand composition (how much is quality-/prestige-/situational-sensitive) | UNRESOLVED — REQUIRES EVIDENCE |
| OQ-EIC-FS-06 | Magnitude of multi-character / organisational coordination costs | UNRESOLVED — REQUIRES EVIDENCE |
| OQ-EIC-FS-07 | Whether any soft quality residual for specialists is acceptable, and its form | UNRESOLVED — REQUIRES HUMAN RULING (consistent with HD-EIC-02) |

No attempt is made to resolve these here.

---

# 15. Recommended Next Validation Step

**Smallest useful next step:**  
Specify and run a discriminating comparative simulation that holds the following fixed under the ruled constraints:

- Layered validation unit (HD-EIC-01).  
- No hard factory-quality ceiling (HD-EIC-02).  
- Legitimate multi-accounting and organisational internalisation as adversarial conditions (HD-EIC-03).  
- Non-numeric functional shapes drawn from the candidates above (HD-EIC-04).

**Minimal simulation focus:**  
Test two or three coherent shape packages against the adversarial cases in §10, measuring whether residual specialist comparative advantage at the independent-participant layer survives.

**Success criterion for the simulation step:**  
It must be able to discriminate “residual advantage survives” from “residual advantage is trivially internalised or closed.” If it cannot, the shapes are still under-specified and further shaping or human ruling is required.

**Explicitly not next:**  
- Provenance & Reputation Engine design.  
- Numeric tuning.  
- Full economy specification.  
- Promotion of any shape to LOCKED.

---

# Decision Gate

| Major EIC Question | Classification |
|--------------------|----------------|
| Layered unit of validation (character vs independent participant) | Sufficiently shaped for testing (HD-EIC-01) |
| Existence of residual specialist comparative advantage without hard factory ceiling | Sufficiently shaped for testing (principle); specific channels require comparative simulation |
| Specialisation / concentration pressure form | Requires comparative simulation and/or human ruling |
| Resource quality / attribute expression | Requires comparative simulation; details still underspecified |
| Crafting residual channels | Requires comparative simulation |
| Manufacturing vs specialist complementarity | Requires comparative simulation |
| NPC substitution dimensions | Requires comparative simulation; strength requires human ruling or evidence |
| Demand sensitivity to residual channels | Requires comparative simulation |
| Multi-account / organisation survival of independent advantage | Requires comparative simulation (authorised adversarial condition) |
| Exact implementation of any budget or numeric parameter | Still too underspecified / forbidden at this stage |
| Provenance & Reputation coupling | Out of scope; deferred |

**Smallest next task required to move the EIC forward:**  
Produce a discriminating comparative simulation specification (or execute one) that packages 2–3 coherent non-numeric shape combinations from this document and tests them against multi-account, organisational, manufacturing-parity, and NPC-substitution adversarial cases, measuring residual independent-participant comparative advantage.

Architecture C remains PROPOSED / LEADING HYPOTHESIS and UNDERDETERMINED pending that evidence.

---

**Document Control**

**Authority:** None. PROVISIONAL functional shapes and simulation assumptions only.  
**Lineage:** Issued under HD-EIC-04 methodological permission in response to the EIC Human Rulings of 2026-08-25 and the under-specification finding of the Minimum Simulation Results.  
**No mechanical design is authorised beyond the provisional shapes stated.**  
**No status promotion has occurred.**

*End of EIC Provisional Functional Shapes v0.1*
