# TCIndustries — Authority & Provenance Reconciliation Matrix

**Artifact Type:** Governance / Provenance  
**Version:** 1.0  
**Date:** 2026-08-25  
**Author:** OpenAI GPT-5.6 Luna  
**Scope:** Master GDD v1.1 Canonical Baseline only  
**Authority created by this artifact:** None. This document creates no new game-design decisions.  
**Purpose:** Determine what LOCKED actually means, distinguish human authority from inference, prevent future LLMs from treating canonical presence as human approval, and establish a clean provenance boundary before the Economic Interdependence Core.

**Controlling references:**
- `TCIndustries_Master_GDD_v1.1_Canonical_Baseline.md`
- `TCIndustries_Master_GDD_v1.0_Canonical_Audit.md`
- `TCIndustries_Master_GDD_v1.0_Consolidated.md`

---

## 1. Authority Classes

| Code | Authority Type | Meaning |
|---|---|---|
| **A** | Explicit Human / Project Ruling | Directly established by the project owner or brief. Strongest authority. |
| **B** | Explicit Project Principle / Pillar | Explicitly established as a project pillar or objective. Not necessarily a detailed implementation rule. |
| **C** | Derived from Locked Principle | Logical consequence of an authoritative principle. Not independently human-approved. |
| **D** | Audit / Governance Reclassification | Status assigned by governance reasoning (audit), not by a new game-design decision. |
| **E** | Multi-Source Inference | Consistent across source documents, but not independently human-approved. Insufficient alone for LOCKED. |
| **F** | Historical / External Reference | SWG or other external material. No TCIndustries design authority. |
| **G** | Prototype / Evidence | Behaviour demonstrated by prototype or test. Evidence only; does not grant design authority. |

**Critical distinction:** A statement may remain **LOCKED** only where its underlying authority is **A** or **B**. Classes **C–E** are not substitutes for human approval.

---

## 2. Governance Rules (Adopted)

### GR-001 — Canonical Presence Is Not Human Approval
**Status: GOVERNANCE RULE**

Inclusion of a statement in the Master GDD does not establish that the statement was independently approved by the project owner. A statement is HUMAN-LOCKED only where its authority can be traced to an explicit human/project ruling or an explicitly designated project-level principle. Logical consequences, repeated proposals, LLM consensus, source-document frequency, and inclusion in a canonical document do **not** independently constitute human approval.

### GR-002 — Historical Proposal Boundary
**Status: GOVERNANCE RULE**

Design proposals appearing in superseded or first-pass documents remain non-canonical unless explicitly promoted through the project’s change-control process. Their existence in source history does not constitute approval, inheritance, or implicit reintroduction into the current Canonical Baseline.

### GR-003 — Derived Constraint Strictness
**Status: GOVERNANCE RULE**

A statement may be marked **DERIVED CONSTRAINT** only when it is a genuine logical consequence of one or more HUMAN-LOCKED principles. Preferable implementation choices, design hypotheses, and “highly consistent with the pillars” statements are **PROPOSED**, not DERIVED CONSTRAINT.

---

## 3. LOCKED Rules — Reconciliation

| ID | Statement | v1.1 Status | Authority Type | Source / Basis | Human-confirmed? | Assessment |
|---|---|---|---|---|---|---|
| VIS-001 | Living persistent player-driven virtual world | LOCKED | A/B | Explicit project objective / core vision | Subject to final confirmation | Legitimate core canon |
| VIS-002 | Citizen rather than chosen hero | LOCKED | A/B | Explicit player fantasy | Subject to final confirmation | Legitimate core canon |
| VIS-003 | Player-driven economy | LOCKED | A/B | Explicit project direction | Subject to final confirmation | Legitimate core canon |
| VIS-004 | Interdependent professions | LOCKED | A/B | Explicit project pillar | Subject to final confirmation | Legitimate principle |
| VIS-005 | Emergence over scripted experiences | LOCKED | A/B | Explicit project direction | Subject to final confirmation | Legitimate principle |
| VIS-006 | Anti-goals / drift prevention | LOCKED | A/B | Explicit anti-goal direction | Subject to final confirmation | Legitimate principle |
| PIL-001 | Discovery / Gold Rush | LOCKED | B | Explicit design pillar | Subject to final confirmation | Legitimate pillar |
| PIL-002 | Identity and reputation | LOCKED | B | Explicit design pillar | Subject to final confirmation | Legitimate pillar |
| PIL-003 | Interdependence | LOCKED | B | Explicit design pillar | Subject to final confirmation | Legitimate pillar |
| PIL-004 | Player-created society | LOCKED | B | Explicit design pillar | Subject to final confirmation | Legitimate pillar |
| LOOP-002 | Meaningful long-term non-combat career | LOCKED | A/B | Explicit project direction | Subject to final confirmation | Legitimate high-level requirement |
| LOOP-003 | Progression expands choices rather than one linear ladder | LOCKED | B | Explicit project philosophy | Subject to final confirmation | Legitimate principle |
| WRLD-001 | Persistent shared world | LOCKED | B/C | “Required by vision” | No independent ruling identified | **Authority review required** |
| PLR-001 | Persistent character identity | LOCKED | B/C | “Required by pillars” | No independent ruling identified | **Authority review required** |
| PLR-003 | Persistent meaningful ownership | LOCKED | B/C | Required by society/economy pillars | No independent ruling identified | **Authority review required** |
| PROF-001 | Flexible skill-based profession architecture | LOCKED | A/B | Explicit project requirement | Subject to final confirmation | Legitimate high-level rule |
| PROF-004 | Respecialisation must be possible | LOCKED | A/B | Explicit project philosophy | Subject to final confirmation | Legitimate principle (exact costs remain TBD) |
| RES-002 | Resources exist to create discovery/scarcity/economic events | LOCKED | B | Resource pillar implementation objective | Subject to final confirmation | Legitimate design objective |
| CRFT-001 | Crafting must produce differentiated products | LOCKED | B/C | Identity/reputation + resource principles | No independent ruling identified | **Authority review required** |
| MFG-001 | Automation must not eliminate player relevance | LOCKED | A/B | Explicit anti-automation direction | Subject to final confirmation | Legitimate high-level rule |
| ECO-001 | Economy primarily player-driven | LOCKED | A/B | Explicit project direction | Subject to final confirmation | Legitimate core rule |
| RET-001 | Player shops/vendors/commercial spaces | LOCKED | A/B | Explicit project direction | Subject to final confirmation | Legitimate core rule |
| SERV-001 | Meaningful non-combat services are viable careers | LOCKED | A/B | Explicit project direction | Subject to final confirmation | Legitimate core rule |
| CMBT-001 | Combat is one lifestyle among many | LOCKED | A/B | Explicit project direction | Subject to final confirmation | Legitimate principle |
| BLD-001 | Player/organisation-owned meaningful structures | LOCKED | A/B | Explicit project direction | Subject to final confirmation | Legitimate core rule |
| SOC-001 | Player organisations | LOCKED | A/B | Explicit project direction | Subject to final confirmation | Legitimate core rule |
| EMRG-001 | Design for emergent gameplay | LOCKED | B | Explicit project direction | Subject to final confirmation | Legitimate principle |
| SAFE-001 | Ownership/transfer safeguards | LOCKED | B/C | Consequence of ownership principle | No independent ruling identified | **Authority review required** |
| SWG-001 | SWG is inspiration, not reproduction | LOCKED | A/B | Explicit project direction | Yes | Legitimate project boundary |
| SWG-002 | Strict separation from SWG IP | LOCKED | A/B | Explicit project/IP boundary | Yes | Legitimate project boundary |
| PROT-001 | GDD has authority over prototype behaviour | LOCKED | D | Governance / audit decision | Governance-approved | Valid governance rule (label as such) |
| EXP-001 (principle) | Data-driven extensibility is desired | LOCKED (principle) | B | Explicit project direction | Subject to final confirmation | Principle legitimate; implementation remains PROPOSED |
| BAL-001 | No premature numeric precision | LOCKED | A/B | Explicit project philosophy | Subject to final confirmation | Legitimate governance/design principle |

### Authority-review queue (not demoted; flagged)

These remain LOCKED in v1.1 but rest partly on “required by pillars / vision” language rather than independently documented human rulings. They require explicit human confirmation before being treated as frozen:

- WRLD-001
- PLR-001
- PLR-003
- CRFT-001
- SAFE-001

---

## 4. DERIVED CONSTRAINTS — Reconciliation

| ID | Statement | v1.1 Status | Claimed Derivation | Authority Type | Assessment | Recommended Disposition |
|---|---|---|---|---|---|---|
| CRFT-006 | Crafting depth must prevent one character efficiently dominating all high-value production | DERIVED CONSTRAINT | PIL-003 + specialisation principles | C/D | Interdependence does **not** logically require a character skill-cap style rule. Other mechanisms (resource, service, geographic, organisational, reputation, capacity) could achieve interdependence. | **DEMOTE → PROPOSED** |
| MFG-004 | No infinite self-sufficient personal-factory automation | DERIVED CONSTRAINT | VIS-006 + MFG-001 | C | Reasonable logical consequence of the anti-goal and controlled-automation principle. | **Retain DERIVED CONSTRAINT** (wording may be narrowed) |
| ECO-003 | No pure spreadsheet optimisation | DERIVED CONSTRAINT | VIS-006 + VIS-003 | C | Reasonable anti-goal consequence. | **Retain DERIVED CONSTRAINT** |
| CMBT-003 | No forced combat path | DERIVED CONSTRAINT | LOOP-002 + VIS-006 | C | Reasonable logical consequence of non-combat viability + anti-goals. | **Retain DERIVED CONSTRAINT** |
| PROG-002 | Multiple viable long-term paths | DERIVED CONSTRAINT | LOOP-002 | C/D | LOOP-002 establishes at least one meaningful non-combat path. “Multiple viable paths” is a stronger design proposition. | **DEMOTE → PROPOSED** |

---

## 5. Explicit Non-Canonical Boundary

The following proposals appear in older first-pass source material. They are **NOT** part of the v1.1 Canonical Baseline. Per GR-002 they remain non-canonical unless and until explicitly promoted by human/project decision.

| Proposal | Canonical Status |
|---|---|
| Factories can never produce Exceptional-tier output | **NOT CANONICAL** |
| Vitality / Stamina / Focus pools | **NOT CANONICAL** |
| Skill Discipline / Skill Box architecture | **NOT CANONICAL** |
| Use-based Skill Point acquisition | **NOT CANONICAL** |
| Two-to-three concurrent profession target | **NOT CANONICAL** |
| Toxicity attribute (as defined resource attribute) | **NOT CANONICAL** |
| Time-limited vs extraction-limited resource pool model (as chosen model) | **NOT CANONICAL** |
| Harvesters as installed structures (specific implementation) | **NOT CANONICAL** |
| Experimentation point allocation / success-critical-failure model | **NOT CANONICAL** |
| Single unified currency as the default | **NOT CANONICAL** |
| Commerce Directory | **NOT CANONICAL** |
| Fully tradeable by default / no account or character soul-binding | **NOT CANONICAL** |
| Services required to be recurring/consumable by design | **NOT CANONICAL** |

“It existed in an earlier TCIndustries document” ≠ “it exists in the v1.1 canon.”

---

## 6. Recommended Status Model

```
HUMAN-LOCKED (A/B)
        │
        ├── explicit project principle or ruling
        │
        ▼
DERIVED CONSTRAINT (C)
        │
        │  only where the consequence is genuinely logical
        │
        ▼
PROPOSED
        │
        ├── design recommendation
        ├── implementation candidate
        └── design constraint candidate
```

Separate evidence/provenance track (does not grant design authority):

```
HISTORICAL (F) ──────────────┐
                             │
PROTOTYPE / EVIDENCE (G) ────┼──> informs design; does not grant authority
                             │
OLDER PROPOSAL ──────────────┘
```

---

## 7. Immediate Governance Disposition

| Action | Disposition |
|---|---|
| CRFT-006 | **DEMOTE → PROPOSED** (to be applied in next status patch or recorded as pending correction) |
| PROG-002 | **DEMOTE → PROPOSED** (same) |
| WRLD-001, PLR-001, PLR-003, CRFT-001, SAFE-001 | **Authority review** — human confirmation required |
| MFG-004, ECO-003, CMBT-003 | Retain as DERIVED CONSTRAINT |
| Older mechanical proposals listed in §5 | Explicitly non-canonical |
| Master GDD body | Do **not** expand or regenerate |
| Numerical values / skill trees / combat systems | Do **not** begin |

---

## 8. Project State After This Reconciliation

| Area | State |
|---|---|
| Vision & pillars | Strong; subject to final human confirmation of HUMAN-LOCKED set |
| Authority framework | Strong (with this matrix) |
| Status vocabulary | Strong |
| Provenance discipline | Strengthened by this artifact |
| Mechanical design maturity | Early — intentionally |
| Economic interdependence | Not yet demonstrated |
| Prototype evidence | Unreconciled |
| Ready for major system design | After human confirmation of HUMAN-LOCKED set + optional prototype reconciliation |
| Ready for canonical freeze | No |

**Clean sequence going forward:**

1. **This artifact** — Authority & Provenance Reconciliation Matrix  
2. **Human confirmation** of the genuine HUMAN-LOCKED set (and disposition of the authority-review queue)  
3. **Prototype evidence reconciliation** (when prototype is available)  
4. **Economic Interdependence Core** package (specialisation ↔ resources ↔ crafting ↔ manufacturing ↔ NPC substitution ↔ demand), expressed as Principle → Constraint → Invariant → Failure Condition → Test  
5. Only then deeper mechanical specification and numerical modelling  

No further LLM consolidation of the Master GDD should occur.

---

## 9. Notes on CRFT-006 and PROG-002 (Rationale)

**CRFT-006:**  
PIL-003 requires that important professions depend upon one another. That does not logically entail a particular character-level “no instant mastery” rule. Interdependence can be produced by resource scarcity, specialised equipment, geography, logistics, reputation, production capacity limits, service dependencies, organisational structure, information asymmetry, opportunity cost, or combinations of these. The statement remains a valuable design objective to test; it is not presently a strict derived constraint.

**PROG-002:**  
LOOP-002 establishes that a meaningful long-term non-combat career must remain viable. That supports “at least one meaningful non-combat path.” The stronger claim that multiple paths must simultaneously be viable is highly desirable for the sandbox philosophy but is a design proposition, not a strict logical implication of LOOP-002.

---

**End of Authority & Provenance Reconciliation Matrix**

This artifact is the governance checkpoint immediately preceding human confirmation of the HUMAN-LOCKED set and the Economic Interdependence Core. It creates no new game design.
