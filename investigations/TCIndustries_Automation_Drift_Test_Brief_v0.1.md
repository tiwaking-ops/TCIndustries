---
doc_id: AUTOMATION-DRIFT-BRIEF-T06
title: TCIndustries — Automation-Drift Structural Test Brief (Task T-06)
version: "0.1"
date: 2026-10-23
status: PROPOSED — Analysis only. Creates no design authority.
authority_class: null   # investigation/analysis record; rules nothing, decides nothing
human_approved: false
controlling_documents: [MASTER-GDD-v1.1.1, HUMAN-RULINGS-REGISTER, EIC-PROGRAM-STATE]
reading_priority: advisory
archived_author_note: "Author: OpenCode documentarian agent under owner approval of Gap Analysis Task T-06 (2026-10-23). Assessor: none (single-session investigation)."
change_note: "v0.1 — Initial formal automation-drift test brief for OQ-005 / MFG-001 / MFG-004, executed under the newly established HD-EIC-10/11 structural boundaries."
---

# TCIndustries — Automation-Drift Structural Test Brief (Task T-06)

**Document Type:** Design Investigation / Adversarial Test Brief (Non-Canonical)  
**Status:** PROPOSED — Analysis only. Creates no design authority, changes no GDD status, authorises no numeric values or automation mechanisms.

---

## 1. Objective & Scope

Task T-06 defines the adversarial test for the project's **DERIVED CONSTRAINT** governing manufacturing and automation:
> *"No configuration of factories/automation makes other players optional."* (MFG-001, MFG-004, GDD §10–11).

This brief specifies the actor model, automation configurations, and measurable interdependence signals required to test whether an economy with player-owned factories drifts into complete self-sufficiency (breaking PIL-003 interdependence) or maintains meaningful economic interdependence across independent participants (HD-EIC-01).

**Contraining Context:**
- **Non-canonical exclusion:** The historical claim that "factories can never produce Exceptional-tier output" is explicitly **non-canonical** (Open Questions Register §6). Factory capabilities must be tested structurally rather than dismissed by flat prohibition.
- **HD-EIC-10/11 boundaries:** Factory automation must respect the adopted durable commitments (C1 non-internalisable scarcity, C2 persistent info/discovery asymmetry, C4 non-economic institutional/social constraint) and the soft-residual exclusions by category.

---

## 2. Actor Model (per HD-EIC-01)

To ensure the test is actor-complete, four distinct participant configurations are modeled:

1. **Single Specialist (Independent Participant):** Focuses on a single profession, relying on market exchange for all other goods and services (the baseline interdependent actor).
2. **Generalist Solo Player:** Tries to harvest, craft, and service independently with zero market interaction.
3. **Multi-Account Unit (Multi-Boxer):** Operates 3–5 characters simultaneously across harvesting, crafting, and basic factory oversight.
4. **Organisational Vertical Cartel:** A large structured guild pooling capital, land, factories, discovery agents, and labor to achieve maximum internalisation.

---

## 3. Automation Configurations Under Test

Three distinct factory/automation configuration shapes (OQ-005) are pitted against the actor model:

- **Config 1 — Unbounded Passive Extraction & Production:** Factories run indefinitely upon placement, requiring zero ongoing input of scarce factors (C1) or process knowledge updates (C2), consuming only bulk base materials.
- **Config 2 — Attribute-Dependent / Scarcity-Gated Automation:** Factories require continuous input of non-internalisable scarce factors (C1) and decay in efficiency unless serviced by specialist crafters or refreshed with periodic discovery updates (C2).
- **Config 3 — Associational / Institutional-Gated Factories:** Factory ownership and operation rights are bound to civic/guild organizational structures (C4), preventing solo or multi-account exploitation without active community participation.

---

## 4. Measurable Interdependence Signals (Interdependence-Event Pattern)

Drawing from the testbed's telemetry patterns (E-09 ledger, interdependence events), three quantitative signals measure whether automation causes interdependence drift:

1. **Inter-Participant Transaction Ratio ($ITR$):** The proportion of intermediate manufacturing inputs sourced from *other* independent participants vs. produced internally within the actor's own alt/organisation network. 
   - *Failure threshold:* $ITR \to 0$ for organisational cartels (complete vertical internalisation).
2. **Service Reliance Index ($SRI$):** The frequency with which automated operators consume external medic, entertainer, and transport services.
   - *Failure threshold:* Factory operators successfully bypass all non-manufacturing service requirements.
3. **Automation-Induced Scarcity Pressure ($AISP$):** Whether automated factories drive competition for scarce natural factors (C1) or merely churn abundant base commodities.

---

## 5. Adversarial Test Specification: The Automation-Drift Stress Test

**Test Protocol:**
1. Initialize a simulated 100-tick economic environment with 40 independent specialists, 20 generalists, 10 multi-account units, and 5 organisational cartels.
2. Deploy Automation Config 1 (Unbounded) across the cartels and multi-account units.
3. Measure $ITR$ and $SRI$ across all actor tiers at tick 100.
4. **Falsification Criterion:** If organisational cartels and multi-account units achieve an $ITR < 0.05$ (meaning >95% of inputs are vertically integrated with zero external market touch) while maintaining parity or superiority over single specialists, **Config 1 is falsified** as violating the derived constraint ("factories make other players optional").
5. Repeat with Config 2 and Config 3 to determine which structural boundary successfully preserves $ITR > 0.20$.

---

## 6. Conclusion & Execution Status

This brief establishes a rigorous, falsifiable adversarial test for OQ-005 / MFG-001 / MFG-004. 

**Execution Status:** PROPOSED design investigation. Cannot be fully executed until a simulation host or testbed runner (such as the resource sim toolchain) implements the actor-network telemetry extensions specified in Task T-07.
