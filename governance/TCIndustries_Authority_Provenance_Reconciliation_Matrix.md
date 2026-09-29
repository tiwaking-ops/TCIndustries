# TCIndustries — Authority & Provenance Reconciliation Matrix

**Artifact Type:** Governance / Provenance  
**Version:** 1.1  
**Date:** 2026-08-25 → **2026-09-28** (corrected; see change note)  
**Author:** OpenAI GPT-5.6 Luna (original artifact, 2026-08-25)  
**Revision authors:** GR-004 added 2026-09-17 (author not recorded — **UNVERIFIED** per GR-004); GR-006 added 2026-09-28 by OpenCode (space-bunny-free), per explicit owner approval  
**Scope:** Master GDD v1.1 Canonical Baseline only  
**Authority created by this artifact:** None. This document creates no new game-design decisions.  
**Purpose:** Determine what LOCKED actually means, distinguish human authority from inference, prevent future LLMs from treating canonical presence as human approval, and establish a clean provenance boundary before the Economic Interdependence Core.

**Change note (2026-09-28):** version and date corrected from v1.0 / 2026-08-25. That metadata had gone stale: GR-004 was added to §2 on 2026-09-17 without a corresponding version or date bump, leaving the header asserting a currency the document no longer had — the same defect class as F4 in GR-006. Corrected under GR-006.3's principle. **GR-006 added to §2** (owner-approved and in force 2026-09-28; its WINDOW, RETRO, REPORT, and EDGE implementation values remain undetermined and reserved to the owner). **GR-005 remains absent from this matrix** and is not in force here — it is recorded only in `proposals/TCIndustries_Manifest.md` as owner-confirmed 2026-08-25, never pasted into §2. Its merge is a separate owner act.

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

### GR-004 — Document Author / Assessor Attribution
**Status: GOVERNANCE RULE**

Every LLM-produced or LLM-assisted document committed to this repository **must** carry a metadata block at the top identifying, at minimum:

- **Author** — the model/agent identity that produced the document (e.g., "OpenCode (big-pickle)", "OpenAI GPT-5.6 Luna").
- **Version** — a document version identifier (e.g., "1.0") that changes on substantive revision.
- **Date** — creation date (ISO format).
- **Assessor** — where a reviewer, auditor, or second model/agent reviewed or verified the document, that identity is recorded separately from the Author. The same entity may hold both roles and should record both roles explicitly when it does.
- **Status / Authority class** — the document's own status (e.g., EVIDENCE / ANALYSIS / PROPOSAL / GOVERNANCE) and, where applicable, its authority class (per §1).

Attribution is **provenance metadata, not authority**. Recording an author/assessor never promotes a document's status or grants design authority.

**Rationale:** LLM sessions are independent and cannot be forced to report identity. Repository-level attribution rules make identity a documentation requirement, enabling provenance checks, audit trails, and the standing discipline that "authorial voice" is not "project decision."

**Enforcement (non-exhaustive):** operating rules in `AGENTS.md`; repository review/gates; document templates; standing convention of prefixed filenames where applicable. Attribution missing or ambiguous is recorded as **UNVERIFIED** in provenance tracking rather than silently assigned.

### GR-006 — Detective and Corrective Control of Project Decisions
**Status: GOVERNANCE RULE — IN FORCE from 2026-09-28, with GR-006.1's filing-window value UNDETERMINED.**

**Owner rationale (verbatim):** "Allows greater detective and corrective control of project decisions."

**Grounding — failures actually observed in this repository before the rule was drafted (not projections):**

| # | Observed failure | Evidence |
|---|---|---|
| F1 | Six owner rulings (HD-ITM-01, HD-PROF-01, HD-BLD-02, HD-PROF-02, HD-ITM-02, HD-BLD-03) were ruled 2026-09-19/20 but a grep of this repository's authority record for all six IDs returned **zero matches** for approximately nine days | Rulings existed only as decision records in `proposals/`; no register entry |
| F2 | Dependent documents actively contradicted ruled decisions — the canonical GDD and the Open Questions Register both still showed OQ-007/010/011/015/016 as `TBD` | `canonical/gdd/TCIndustries_Master_GDD_v1.1.1_Status_Patch.md` §31; `proposals/TCIndustries_Open_Questions_Register.md` §2 |
| F3 | The project memory recorded none of the six rulings | `governance/project_memory.md` ended 2026-09-18 |
| F4 | The authority register's own "Full Cumulative State" table was dated 2026-08-25 while the register body had reached v1.2 / 2026-09-19 — a summary silently disagreeing with the document it summarised | Pre-repair state, 2026-09-28 |
| F5 | Reserved filing-time acts were indistinguishable from forgotten ones — records "reserved" acts with no tracked state, so the gap was invisible in both directions | Filing-time act lists in the six decision records above |

**Common root cause:** the project has strong classification machinery (authority classes, status vocabulary, GR-001–GR-005) and weak detection/correction machinery. Nothing in the repository could tell an auditor, or the owner, whether a decision that *should* be in the authority record actually was.

**The rule.**

- **GR-006.1 — Authority-record presence.** Every ruling the owner issues must be present in full in `governance/TCIndustries_Human_Rulings_Register.md`. A ruling recorded only elsewhere is a **pending filing**, not a completed decision, until the register entry exists. **The filing window is UNDETERMINED:** the owner has not set a period, and no default is inferred, because an invented threshold would itself be an unauthorised value. The obligation to file is in force; the deadline is an **open implementation value reserved to the owner**.
- **GR-006.2 — Dependent-document consistency.** Where a filing changes a status appearing in another repository document, that document must be updated in the same filing action. The dependent set is fixed, not left to judgement: at minimum `canonical/gdd/` (status tables and TBD entries), the Open Questions Register, the Invariant Register, `governance/project_memory.md`, and the Continuous Queue Tracker. **Divergence between the authority record and a dependent document is a defect and is reported as one.**
- **GR-006.3 — Summary-table currency.** A document carrying a summary, index, or "current state" table that asserts its own currency must bring it current in the same action as any change it summarises. A dated summary knowingly behind its body must say so on its face.
- **GR-006.4 — Filing acts are tracked, not merely reserved.** A decision record may not reserve filing-time acts without recording them in a reportable state (pending / done / waived-with-reason). An act neither done nor waived is an **open obligation**, and open obligations are reportable.
- **GR-006.5 — Closure requires a mechanical check, not a judgement call.** A filing is not closed until a search for the ruling's ID returns a hit in the register **and** in every dependent document named in GR-006.2. This is the property that would have caught F1 on the day it happened.
- **GR-006.6 — Correction is filing, not re-litigation.** When a gap is found the remedy is to **complete the record**, not to revisit or reverse the decision. The ruling's authority is unaffected by how long it took to file.
- **GR-006.7 — A filing defect never voids a ruling.** Late, missing, or incomplete filing record **never** invalidates the underlying ruling and never removes a LOCKED status. This rule may not be used to argue that a decision the owner made is not a decision. *(Follows the precedent set by GR-004 and the HD-GDD-01 provenance-note pattern.)*
- **GR-006.8 — Detection is owed to the owner, not only to agents.** Any agent may and should surface a filing gap when it encounters one, as a **finding for the owner**, with specific IDs and file paths. **Reporting a gap is never authorisation to correct it unilaterally.** *(Reinforced by the 2026-09-17 duplicate Evidence Reconciliation Pass incident recorded in `governance/project_memory.md`.)*

**What this rule does not do.** It changes no design status and creates no game-design decision, invariant, or numeric value. It creates no new authority class — §1 classes A–G are unchanged, and a governance rule is not a source of design authority. It does not make any agent authoritative. It does not retroactively void or require re-ruling anything. It does not alter GR-001 through GR-005. It does not open, close, or substitute for HD-EIC-07. It does not require LLM consensus on anything.

**Open implementation values — reserved to the owner, not defaults:**

| Slot | Question | State |
|---|---|---|
| **WINDOW** | Period within which a ruling must be filed | **UNDETERMINED** — GR-006.1 in force, deadline unset |
| **RETRO** | Forward-only, or a one-off reconciliation sweep of all prior filings? | **NOT AUTHORISED** — no sweep performed or required |
| **REPORT** | Who runs the GR-006.5 check, and at what cadence? | **UNDETERMINED** |
| **EDGE** | If the owner rules in-session and an agent drafts the register entry immediately, is that "filing," or does filing require a separate later act? | **UNDETERMINED** — materially changes the burden; worth deciding explicitly |

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
