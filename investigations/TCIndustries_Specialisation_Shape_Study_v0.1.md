---
doc_id: SPEC-SHAPE-STUDY-T03
title: TCIndustries — Specialisation-Budget Shape Study (Task T-03)
version: "0.1"
date: 2026-10-23
status: PROPOSED — Analysis only. Creates no design authority.
authority_class: null   # investigation/analysis record; rules nothing, decides nothing
human_approved: false
controlling_documents: [MASTER-GDD-v1.1.1, HUMAN-RULINGS-REGISTER, EIC-PROGRAM-STATE]
reading_priority: advisory
archived_author_note: "Author: OpenCode documentarian agent under owner approval of Gap Analysis Task T-03 (2026-10-23). Assessor: none (single-session investigation)."
change_note: "v0.1 — Initial formal shape study for OQ-001 / PROF-005 / CRFT-006, executed under the newly established HD-EIC-10 structural boundaries (G2 gloss; C1, C2, C4 commitments; soft-residual exclusions by category)."
---

# TCIndustries — Specialisation-Budget Shape Study (Task T-03)

**Document Type:** Design Investigation / Comparative Shape Analysis (Non-Canonical)  
**Status:** PROPOSED — Analysis only. Creates no design authority, selects no shape, changes no GDD status, authorises no numeric values.

---

## 1. Objective & Scope

Task T-03 compares candidate specialisation-budget *shapes* (mechanisms governing how character capacity, skill points, or profession limits are bounded) against the project's LOCKED constraints and the newly established HD-EIC-10 structural boundaries:
- **PIL-003** (Interdependence must create opportunity rather than coercive inconvenience)
- **PROF-004 / HD-PROF-01 / HD-PROF-02** (Respecialisation: identity survives, free-drop cost instrument)
- **HD-PROF-03** (Trainer-gated acquisition with free player-teacher path)
- **HD-EIC-01** (Unit of validation: independent participant, including multi-account and organisational actors)
- **HD-EIC-02** (No hard factory-quality ceiling)
- **HD-EIC-10** (Durable commitments: C1 non-internalisable scarcity, C2 persistent info/discovery asymmetry, C4 non-economic institutional/social constraint; soft-residual exclusions by category)

**Explicitly out of scope:** Selecting a winning shape (that remains an open human decision for OQ-001); numeric parameterisation (BAL-001 gated); architectural residual-channel invention outside HD-EIC-10 boundaries; GDD text edits.

---

## 2. The Specialisation-Budget Problem (OQ-001)

The Core Question (OQ-001): *What is the exact specialisation budget / skill-cap model?*

Under Pre-CU SWG historical precedent, characters held a pool of **250 skill points** (or skill boxes costing varying points from 5 to 55). Under TCIndustries GDD design constraints, a pure skill-point ceiling proved vulnerable to multi-account internalisation (Package β falsification — HD-EIC-05) unless bounded by structural factors. 

With the HD-EIC-10 gate closed and HD-EIC-11 authorising bounded search under G2 ("literal effortlessness"), specialisation budgeting can now be analysed against the three adopted durable commitments (C1, C2, C4) rather than treating specialisation as a purely isolated character-sheet slider.

---

## 3. Candidate Specialisation Shapes

We examine four distinct structural shapes for character and participant capacity budgeting:

### Shape 1 — Global Fixed Skill-Point Pool (Pre-CU Historical Baseline)
- **Mechanism:** A hard numerical ceiling (e.g., 250 points) across all professions. To learn a new advanced profession, a player must drop an existing one (free-drop per HD-PROF-02).
- **Adversarial / Internalisation vulnerability:** Multi-account participants simply instantiate complementary characters (alt-farming) to cover all skill trees simultaneously, bypassing character-layer limits entirely (Package β failure mode).
- **Alignment with HD-EIC-10:** Fails to engage C1, C2, or C4 unless combined with external scarcity or information barriers; pure character-sheet caps are internally bypassable by organisational pooling.
- **Verdict:** Insufficient on its own; viable only as a character-layer hygiene rule subordinate to participant-layer constraints.

### Shape 2 — Per-Domain / Mutually Exclusive Specialisation Trees (Hard Bifurcation)
- **Mechanism:** Characters must select mutually exclusive branches (e.g., combat path vs. crafting path vs. service path) with permanent or high-friction barriers between domains.
- **Adversarial / Internalisation vulnerability:** Directly creates the "coercive inconvenience" PIL-003 forbids. Forces multi-accounting immediately because players refuse to be locked out of core systems on their main character.
- **Alignment with HD-EIC-10:** Violates PIL-003 ("opportunity rather than coercive inconvenience") and creates artificial friction rather than structural comparative advantage.
- **Verdict:** Fails PIL-003; strongly disfavored.

### Shape 3 — Activity-Based Decay of Unused Mastery / Maintenance Demand
- **Mechanism:** Skill effectiveness or capacity decays over time unless maintained through active use, specialized inputs (C1 scarcity), or ongoing social/civic participation (C4 institutional constraint).
- **Adversarial / Internalisation vulnerability:** Organisations can automate maintenance via multi-boxing scripts or dedicated alt-updaters, which was the exact mechanism that killed γ′-2 in the EIC retests. Under G2 ("literal effortlessness"), however, if the organizational effort required to maintain multiple active specialisations crosses the "effortless" threshold, it provides *sustained friction*.
- **Alignment with HD-EIC-10:** Strongly engages **C2** (discovery/knowledge maintenance as scale-resistant asymmetry) and **C4** (social/associational upkeep).
- **Verdict:** Viable candidate shape under HD-EIC-10 G2 framing, provided maintenance cannot be cheaply scripted away.

### Shape 4 — Organisational Complementarity & Interdependence Budget (Network Cap)
- **Mechanism:** Skill cap or efficiency bonuses scale based on direct, non-monetary interdependence linkages with other active participants (e.g., entertainer buffs, medic treatments, crafter-supplied tools required for peak output).
- **Adversarial / Internalisation vulnerability:** Organisations can internalise all links internally (the vertical integration failure tested in Architecture C). However, when combined with **C4** (non-economic institutional constraints, guild/association membership gates akin to the top-tier doctor buff rule in F-7), internalisation faces non-economic limits.
- **Alignment with HD-EIC-10:** Directly instantiates **PIL-003** and engages **C4** (associational/social constraints).
- **Verdict:** Highly aligned with project pillars, but depends entirely on the non-economic binding mechanism (C4).

---

## 4. Adversarial Walk-Through: Multi-Account & Organisational Stress Test

| Shape | Multi-Account Alt-Pooling | Organisational Vertical Integration | PIL-003 Compliance | HD-EIC-10 Alignment |
|---|---|---|---|---|
| **1. Global Point Pool** | Bypassed instantly (alt armies cover all trees) | Fully bypassed | Compliant (flexible respec) | Fails (soft-residual only) |
| **2. Per-Domain Bifurcation** | Bypassed via multi-boxing alts | Bypassed via alt division of labor | **Violates PIL-003** (coercive inconvenience) | Neutral / Artificial |
| **3. Maintenance Decay** | Resisted via updaters / scripts | Bypassed unless maintenance cost is structural (C1/C2) | Compliant if friction is moderate | Engages C2 / C4 |
| **4. Organisational Complementarity** | Bypassed if all roles held within guild | Bypassed if vertical integration is unconstrained | Compliant | Engages C4 (institutional constraint) |

---

## 5. What Evidence Would Discriminate (Discrimination Criteria)

To move OQ-001 from PROPOSED to a ruling candidate, future simulation or decision briefs must test:
1. **Effort Threshold (G2 test):** Does a candidate shape impose *costly organisational effort* (not mere capital expenditure) to maintain broad specialisation coverage across an organisation?
2. **Complementarity Binding:** Can organisational complementarity be sustained purely through market contracts, or does it require non-economic associational bonds (C4 / F-7 historical pattern)?
3. **Respec Friction:** Does the specialisation budget preserve HD-PROF-02's free-drop respec rule without creating degenerate instant-swapping exploitative loops?

---

## 6. Conclusion & Next Steps

This study concludes that **no pure character-sheet skill-point cap (Shape 1) can solve the participant-layer closure problem** on its own. Any viable specialisation budget model must anchor on at least one HD-EIC-10 durable commitment — specifically combining **Shape 3** (maintenance/decay) or **Shape 4** (complementarity) with **C2** (discovery asymmetry) or **C4** (institutional/associational constraints).

**Next Step:** This study is PROPOSED analysis. It feeds the upcoming decision-support and queue sessions without selecting a shape.
