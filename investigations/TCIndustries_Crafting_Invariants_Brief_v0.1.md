---
doc_id: CRFT-INVARIANTS-BRIEF-T05
title: TCIndustries — Crafting-Differentiation Invariant Test Brief (Task T-05)
version: "0.1"
date: 2026-10-23
status: PROPOSED — Analysis only. Creates no design authority.
authority_class: null   # investigation/analysis record; rules nothing, decides nothing
human_approved: false
controlling_documents: [MASTER-GDD-v1.1.1, HUMAN-RULINGS-REGISTER, EIC-PROGRAM-STATE]
reading_priority: advisory
archived_author_note: "Author: OpenCode documentarian agent under owner approval of Gap Analysis Task T-05 (2026-10-23). Assessor: none (single-session investigation)."
change_note: "v0.1 — Initial formal crafting differentiation invariant test brief for OQ-003 / OQ-004 / CRFT-001, executed under the newly established HD-EIC-10/11 structural boundaries."
---

# TCIndustries — Crafting-Differentiation Invariant Test Brief (Task T-05)

**Document Type:** Design Investigation / Test Brief (Non-Canonical)  
**Status:** PROPOSED — Analysis only. Creates no design authority, changes no GDD status, authorises no numeric values or experimentation mechanisms.

---

## 1. Objective & Scope

Task T-05 derives testable design invariants from the LOCKED principle **CRFT-001** (Differentiated Products) and specifies the simulation or test structure required to evaluate candidate crafting experimentation *shapes* (OQ-004), resource attribute mappings (OQ-003), and manufacturing constraints (OQ-005), without implementing any game mechanics.

**Contraining Context:**
- **CRFT-001** is HUMAN-LOCKED (principle only, per HD-GDD-01): crafting must produce differentiated products rather than identical commodities, but quality, experimentation, Exceptional tier, attributes, factories, and tiers remain open.
- **Non-canonical exclusions:** The older historical proposals defining "experimentation point allocation / success-critical-failure model" and "factories can never produce Exceptional-tier output" are explicitly **non-canonical** (Open Questions Register §6). This brief analyzes why those historical models failed or were excluded and specifies how alternative shapes must be tested.
- **HD-EIC-10/11 boundaries:** Any crafting differentiation mechanism must be evaluated against the durable commitments (C1 non-internalisable scarcity, C2 persistent info/discovery asymmetry, C4 non-economic institutional/social constraint) and the soft-residual exclusions by category.

---

## 2. Invariant Derivation from CRFT-001

From CRFT-001, we derive three fundamental testable invariants (Principle → Constraint → Invariant → Failure Condition → Test):

### Invariant I-CRFT-1 — Deterministic Recipe Response (Input-Output Fidelity)
- **Principle:** A crafter's skill and ingredient choices must have predictable, non-random structural effects on product attributes.
- **Constraint:** Equivalent inputs plus equivalent skill must yield equivalent product outputs; differing inputs or higher skill must yield measurable, non-trivial differences in product attributes.
- **Invariant (INV-CRFT-01):** *For any recipe class $C$, input attribute set $A$, and skill level $S$: output product attributes $O = f(A, S)$ are continuous and monotonic with respect to attribute quality and skill, with zero random noise in base attribute mapping.*
- **Failure Condition:** Identical inputs and skill producing wildly variant outputs without explanatory variation, or differing input quality producing identical output attributes (quality-blind collapse).
- **Test:** Input-Output Coverage Matrix (automated test harness checking input attribute variations against resulting product attribute vectors across all recipe classes).

### Invariant I-CRFT-2 — Non-Commoditisation (Attribute Differentiation Durability)
- **Principle:** Differentiated products must resist commoditisation by mass production or organisational pooling.
- **Constraint:** Mass production (manufacturing/automation) must not erase the value of crafter experimentation, resource attribute variance, or specialized process knowledge (C2).
- **Invariant (INV-CRFT-02):** *The marginal value of crafting output scales with input attribute quality and crafter skill such that no automated factory configuration can produce items of maximal attribute tier without consuming scarce, attribute-bearing inputs or non-codified process knowledge.*
- **Failure Condition:** Factory automation producing maximal-quality goods from homogenous, ubiquitous base inputs, reducing all crafted items to flat commodities.
- **Test:** Adversarial Automation-Drift Simulation (pitting automated factories against skilled specialist crafters under C1/C2 constraint checks).

### Invariant I-CRFT-3 — Provenance & Skill Attribution
- **Principle:** Crafted items reflect their maker's expertise and provenance.
- **Constraint:** The output must encode the crafter's identity or brand attribute in a way that is verifiable by consumers (supporting PLR-003 ownership and market reputation).
- **Invariant (INV-CRFT-03):** *Every crafted item retains an immutable reference to its crafting author and constituent resource provenance batch, enabling consumer-facing differentiation based on maker reputation and resource origin.*
- **Failure Condition:** Items losing maker attribution upon transfer, or market pricing failing to reflect crafter reputation differentials.
- **Test:** Provenance Traceability Audit (verifying metadata persistence across transfer and bazaar transactions).

---

## 3. Evaluation of Candidate Experimentation Shapes (OQ-004)

The old model of "experimentation point allocation / success-critical-failure" is on the non-canonical list because it encouraged degenerate point-grinding and macro automation. Alternative shapes for OQ-004 to be evaluated under this test brief include:

1. **Information-Asymmetric Recipe Discovery (C2-Coupled):** Experimentation is not point-spending, but *discovery of hidden attribute interactions* through trial and material consumption, where discovery efficiency depends on process knowledge that cannot be instantly copied across an organisation.
2. **Resource-Yield Scaling (C1-Coupled):** Experimentation outcomes are bounded directly by the quality and rarity of harvested raw material attributes, making craftsmanship an interpretive translation of scarce natural inputs rather than an arbitrary min-max minigame.
3. **Associational / Apprenticeship Crafting Gates (C4-Coupled):** Advanced experimentation tiers require guild/associational transfer of techniques (mirroring the F-7 doctor buff rule), embedding crafting mastery in social and institutional structures rather than solo grinding.

---

## 4. Specification of the Crafting Differentiation Test

To evaluate any candidate crafting shape without implementing game code, the test framework must execute a **Crafting Differentiation Audit**:
1. **Matrix Generation:** Define a matrix of 10 resource attribute profiles × 5 skill levels × 3 automation tiers.
2. **Execution Simulation:** Run deterministic attribute mapping functions across the matrix.
3. **Metric Calculation:**
   - *Differentiation Index ($DI$):* Variance of output attributes divided by variance of input attributes. Must remain $> 0$ across all valid inputs.
   - *Automation Parity Ratio ($APR$):* Maximum output quality achieved by automated factories divided by maximum output quality achieved by skilled crafters using scarce inputs. Must remain $< 1.0$ when C1/C2 constraints are active.
4. **Verdicts:** Pass / Fail-with-telemetry / Not-Measurable.

---

## 5. Conclusion

CRFT-001 is a LOCKED principle, but its implementation mechanics (OQ-003, OQ-004, OQ-005) are wide open. This brief establishes that any future crafting and experimentation model must satisfy **INV-CRFT-01** (deterministic fidelity), **INV-CRFT-02** (anti-commoditisation under C1/C2), and **INV-CRFT-03** (provenance persistence), while explicitly avoiding the non-canonical point-allocation experimentation trap.
