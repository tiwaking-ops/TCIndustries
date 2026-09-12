# TCIndustries Documentation Consolidation and Cross-LLM Auditability Assessment

## Role

I am acting as an independent **documentation architect and evidence/governance auditor** specialising in long-lived complex software/game projects and multi-LLM documentation workflows.

I am **not** acting as the project design authority. I will not make game-design decisions or promote proposals into decisions.

**Author:** DeepSeek (per filename, unverified)

---

## 1. Executive Verdict

| Dimension | Rating | Notes |
|---|---|---|
| **Documentation discipline** | Very strong | Status vocabulary, authority rules, and governance principles are exceptionally well-developed |
| **Content redundancy** | High | Significant duplication exists between GDD versions and within the EIC investigation chain |
| **Authority clarity** | Good, but buried | Critical authority distinctions exist but are embedded in documents that may not be read first |
| **Cross-LLM auditability** | Moderate–Good | Possible but requires reading too many documents; no single entry point or manifest exists |
| **Historical containment** | Strong | Non-canonical boundaries and historical separation are well-defined |
| **State reconstruction** | Moderate | Current project state is inferable but requires parsing multiple documents and tracking version relationships |
| **Scalability** | Adequate | Current structure will not scale well as detailed system specs accumulate |

**Overall Assessment:**

The project possesses **excellent governance discipline** but suffers from **structural sprawl**. The underlying concepts (HUMAN-LOCKED, DERIVED CONSTRAINT, PROPOSED, canonical vs non-canonical, evidence vs authority) are sound and well-articulated. However, the documentation is distributed across too many overlapping files with unclear entry points for an independent LLM.

The immediate risk is that a new LLM (or human) attempting to understand the project state may:
1. Read the wrong document first and form an incorrect impression
2. Miss critical authority distinctions buried in supporting files
3. Mistake historical material for current canon
4. Be unable to determine which statements have been human-approved vs derived vs proposed

**Recommendation:** A modest consolidation into a **small, intentional active set** with a **single authoritative entry point** and a **stable state manifest**, preserving all historical material in a clearly separated archive.

---

## 2. Current Documentation Assessment

### 2.1 Overall Character

The TCIndustries documentation corpus is unusually rigorous for an early-stage game project. It demonstrates:

- **Authority hygiene**: Clear distinction between human decisions, LLM proposals, logical consequences, and historical material
- **Provenance consciousness**: Explicit tracking of where statements came from and what changed
- **Falsification orientation**: Design hypotheses are accompanied by explicit failure conditions and test criteria
- **Self-awareness**: Documents acknowledge their own limitations and under-specification
- **Change control**: Versioning and status discipline are consistently applied

However, this rigor has produced **document proliferation**. The EIC investigation alone generated approximately eight documents in a chain, many of which are now superseded or purely evidentiary.

### 2.2 Current Document Categories

| Category | Documents | Count |
|---|---|---|
| **Primary GDD** | v1.0 Consolidated, v1.1 Canonical Baseline, v1.1.1 Status Patch | 3 |
| **Audit/Governance** | Canonical Audit, Authority/Provenance Reconciliation Matrix | 2 |
| **EIC Investigation** | Design Investigation, Candidate v0.1, Falsification Pass, Comparative Simulation Specification, Simulation Results, Functional Shapes, γ Closure, γ′ Re-test Results | 8 |
| **Human Rulings** | Decision Brief, Human Rulings (HD-EIC-01–04), Human Rulings Architecture C Falsification | 3 |
| **Assessment/Review** | Documentation Consolidation Prompt (this analysis) | 1 |

**Total identified:** ~17 documents

### 2.3 Document Quality Assessment

| Document | Value | Authority | Status |
|---|---|---|---|
| **Master GDD v1.1.1** | High — current working reference | Human-approved (status patch) | Active |
| **Master GDD v1.1** | Moderate — superseded by v1.1.1 | Historical baseline | Superseded |
| **Master GDD v1.0** | Low — superseded | LLM consolidation | Superseded |
| **Canonical Audit** | High — provided corrections applied in v1.1 | Audit evidence | Archive |
| **Authority/Provenance Matrix** | High — governance framework | Governance | Active reference |
| **EIC Design Investigation** | High — conceptual analysis | PROPOSED | Evidence |
| **EIC Candidate v0.1** | Medium — superseded by simulation results | PROPOSED | Archive |
| **EIC Falsification Pass** | Medium — adversarial analysis | Evidence | Archive |
| **EIC Comparative Simulation Specification** | Medium — experimental design | Evidence | Archive |
| **EIC Simulation Results** | Medium — pre-functional-shapes | Evidence | Archive |
| **EIC Functional Shapes** | Medium — under-specified | PROVISIONAL | Archive |
| **EIC γ Closure** | Medium — refinement attempt | PROVISIONAL | Archive |
| **EIC γ′ Re-test Results** | Medium — final falsification | Evidence | Archive |
| **Human Decision Brief** | High — isolated open decisions | Human decision gate | Active reference |
| **Human Rulings (HD-EIC-01–04)** | High — human decisions | HUMAN-LOCKED | Active reference |
| **Human Rulings (Architecture C Falsification)** | High — human decisions | HUMAN-LOCKED | Active reference |
| **This analysis** | N/A — external review | No authority | Assessment |

---

## 3. Current Documentation Inventory

### 3.1 Active/Authoritative Documents

| Document | Version | Authority | Purpose |
|---|---|---|---|
| `TCIndustries_Master_GDD_v1.1.1_Status_Patch.md` | 1.1.1 | Working canonical reference | Current design baseline |
| `TCIndustries_Authority_Provenance_Reconciliation_Matrix.md` | 1.0 | Governance | Authority classification and rules |
| `TCIndustries_EIC_Human_Rulings_2026-08-25.md` | 1.0 | HUMAN-LOCKED | EIC validation principles |
| `TCIndustries_EIC_Human_Rulings_Architecture_C_Falsification.md` | 1.0 | HUMAN-LOCKED | Architecture C retirement |
| `TCIndustries_EIC_Human_Decision_Brief_v0.1.md` | 0.1 | Human gate record | Isolated open decisions |
| `TCIndustries_Master_GDD_v1.0_Canonical_Audit.md` | 1.0 | Audit evidence | Corrections applied; active as reference |

### 3.2 Superseded/Historical Documents

| Document | Version | Superseded By | Status |
|---|---|---|---|
| `TCIndustries_Master_GDD_v1.1_Canonical_Baseline.md` | 1.1 | v1.1.1 Status Patch | Superseded |
| `TCIndustries_Master_GDD_v1.0_Consolidated.md` | 1.0 | v1.1 | Superseded |
| `TCIndustries_Master_GDD_v1.0_Consolidated.md` (within merged file) | 1.0 | v1.1 | Superseded |

### 3.3 EIC Evidence/Archive Documents

| Document | Status | Recommendation |
|---|---|---|
| `EIC_Design_Investigation.md` | Evidence | Archive |
| `EIC_Candidate_v0.1.md` | PROPOSED (superseded) | Archive |
| `EIC_Falsification_Pass.md` | Evidence | Archive |
| `EIC_Comparative_Simulation_Specification_v0.1.md` | Evidence | Archive |
| `EIC_Minimum_Simulation_Results_v0.1.md` | Evidence | Archive |
| `EIC_Provisional_Functional_Shapes_v0.1.md` | PROVISIONAL (superseded by γ closure) | Archive |
| `EIC_Targeted_γ_Closure_v0.2.md` | PROVISIONAL (superseded by γ′ results) | Archive |
| `EIC_Targeted_γ′_Adversarial_Re-test_Results_v0.1.md` | Evidence | Archive |
| `EIC_Comparative_Simulation_Execution_&_Results_v0.1.md` | Evidence | Archive |

### 3.4 Duplicate/Overlapping Content

The following documents contain significant overlapping material:

| Overlap | Documents | Severity |
|---|---|---|
| GDD core vision/pillars | All three GDD versions | High — multiple versions of same content |
| Authority rules | Authority Matrix + GDD v1.1 + GDD v1.1.1 | Medium — consistent but redundant |
| EIC mechanism descriptions | Investigation, Candidate v0.1, Functional Shapes | High — concepts restated across chain |
| Simulation specifications | Specification + Results + γ Closure | Medium — specification content repeated |
| Human rulings | Human Rulings + Human Decision Brief + Architecture C Falsification | Low — distinct decisions but overlapping context |

---

## 4. Redundancy and Sprawl Analysis

### 4.1 Sprawl by Topic

| Topic Area | Documents | Count |
|---|---|---|
| EIC architecture investigation | 8 | Highest sprawl |
| GDD core/vision | 3 | High duplication |
| Human governance/rulings | 3 | Moderate duplication |
| Simulation specification/results | 4 | High sprawl |
| Authority/provenance | 2 | Low duplication |

### 4.2 Sources of Sprawl

1. **Sequential EIC investigation chain**: Each step generated a new document rather than updating a single living document
2. **Version proliferation**: Three GDD versions with only minor differences between them
3. **Evidence accumulation**: Falsification passes, simulation results, and re-tests generated separate documents
4. **Functional shape refinement**: Targeted γ closure created additional documents rather than integrating findings into a single EIC reference
5. **Human gate records**: Human decisions documented in multiple formats

### 4.3 What Could Be Eliminated or Consolidated

| Document | Action | Rationale |
|---|---|---|
| GDD v1.0 | Archive only | Fully superseded |
| GDD v1.1 | Archive only | Fully superseded by v1.1.1 |
| All EIC documents except Design Investigation + Final Results | Archive/Evidence | Historical record; not current design |
| Falsification Pass | Archive | Evidence of adversarial analysis; not current |
| γ Closure v0.2 | Archive | Superseded by γ′ results |
| Comparative Simulation Specification | Archive | Experimental design record only |
| Minimum Simulation Results | Archive | Superseded by γ′ results |

### 4.4 What Should Remain Active

| Document | Status | Rationale |
|---|---|---|
| Master GDD v1.1.1 | Active | Current working reference |
| Authority/Provenance Matrix | Active | Governance framework |
| Human Rulings (combined) | Active | Human decisions |
| EIC Design Investigation | Active | Conceptual foundation (no longer leading candidate, but valuable analysis) |
| EIC Final Evidence Report | Active (to be created) | Summary of what evidence showed |

---

## 5. Authority and Provenance Risks

### 5.1 Authority Risks

| Risk | Severity | Mitigation |
|---|---|---|
| **Multiple GDD versions** may lead to confusion about which is current | High | Reduce to single active GDD; archive others |
| **HUMAN-LOCKED rulings** are spread across multiple documents | Medium | Consolidate human rulings into a single authoritative record |
| **Authority Matrix** contains critical rules that may not be read by a new LLM | Medium | Elevate to a "first-read" entry point |
| **EIC investigation documents** may be mistaken for current design | High | Clearly mark as ARCHIVE/EVIDENCE; create summary of conclusions |
| **PROVISIONAL shapes** in Functional Shapes may be treated as design | Medium | Explicitly mark as superseded in an archive manifest |
| **Derived constraints** may be mistaken for human decisions | Low | Status tags are clear; risk is reader error |

### 5.2 Provenance Risks

| Risk | Severity | Mitigation |
|---|---|---|
| **Statement provenance** is traceable but requires cross-referencing | Medium | Add document-level provenance in metadata/headers |
| **Change history** between v1.0, v1.1, v1.1.1 is documented but not summarised | Low | Add change summary as a front section |
| **EIC decision chain** (which documents led to which conclusions) is complex | Medium | Create a traceability map or evidence chain summary |
| **Human ruling dates** and dependencies are not consistently highlighted | Low | Consistent date/authority headers |

### 5.3 Potential Conflicting Authority

| Conflict | Resolution |
|---|---|
| Three GDD versions all claiming authority | v1.1.1 is the current working reference; others are historical |
| EIC documents labelling themselves "leading candidate" | Architecture C is RETIRED as of HD-EIC-05 |
| Functional Shapes marked PROVISIONAL | This status is accurate; not current design |
| Authority Matrix ruling on CRFT-006 vs GDD v1.1 status | Human rulings (HD-EIC-05–08) supersede |

---

## 6. Cross-LLM Auditability Assessment

### 6.1 Current State Reconstruction (Criterion A)

**Rating: Moderate**

An independent LLM can determine the current project state, but it requires:

1. Reading the correct GDD version (v1.1.1) — but the document itself does not explicitly say "read me first"
2. Reading the Authority Matrix to understand status distinctions
3. Reading the Human Rulings to know what has been decided
4. Reading the EIC final results to understand that Architecture C failed
5. Inferring that earlier EIC documents are superseded

**Gap:** No single "entry point" or manifest exists.

### 6.2 Authority Reconstruction (Criterion B)

**Rating: Good**

The distinction between HUMAN-LOCKED, DERIVED CONSTRAINT, and PROPOSED is clearly marked in the GDD. However:

- Human rulings are in separate documents from the GDD
- Some rulings (e.g., Architecture C retirement) are not reflected in the GDD status tags
- A new LLM might read only the GDD and miss the fact that Architecture C is retired

### 6.3 Temporal Reconstruction (Criterion C)

**Rating: Moderate**

- GDD versions are dated and versioned
- EIC documents are dated but the version relationships are complex
- No consolidated change log exists across the entire corpus

### 6.4 Provenance (Criterion D)

**Rating: Good**

The Authority Matrix and audit documents provide strong provenance for key statements. However:

- Provenance is in separate documents from the statements themselves
- A full audit requires cross-referencing

### 6.5 Cross-LLM Comparability (Criterion E)

**Rating: Moderate**

Two independent LLMs reading the same corpus would likely:

- Agree on the GDD status tags
- Agree on the authority rules
- Possibly disagree on which EIC documents represent "current" design
- Possibly produce different summaries of the current state

**The main risk is reading order.** If LLM A reads the GDD first and LLM B reads the EIC investigation first, they may form different initial impressions. This is the most critical auditability gap.

### 6.6 Change Detection (Criterion F)

**Rating: Moderate**

- Version differences are documented in the GDD status patch
- EIC changes are documented within the EIC document chain
- No consolidated "what changed since last state" summary exists

### 6.7 Scope Control (Criterion G)

**Rating: Good**

The project has strong historical/non-canonical boundaries. Older proposals are explicitly marked as non-canonical. The risk of accidental contamination is lower than in most projects.

### 6.8 Scalability (Criterion H)

**Rating: Adequate but needs improvement**

The current structure will not scale well. If the project adds detailed combat, city, resource, and crafting specifications, the document count will grow without a clear organisational principle. The EIC chain already demonstrates what happens: a sequence of documents that are hard to navigate.

---

## 7. Recommended Documentation Architecture

### 7.1 Design Principles

1. **Single Entry Point**: One document that any independent LLM should read first to understand the project state
2. **Small Active Set**: ≤6 active documents that together define the current project
3. **Separate Archive**: All historical/superseded/evidence documents in a clearly marked archive
4. **State Manifest**: Machine-readable summary of the current project state
5. **Human Rulings Consolidation**: All human rulings in one place
6. **Evidence Summary**: What evidence exists and what it showed, in one place

### 7.2 Proposed Active Document Set

| Document | Purpose | Authority | Audience |
|---|---|---|---|
| **TCIndustries_Preamble.md** | Entry point; state summary; document map | Project overview | All readers |
| **TCIndustries_Master_GDD_v1.2.md** | Single current GDD | Working canonical reference | All readers |
| **TCIndustries_Human_Rulings.md** | All human decisions/rulings | HUMAN-LOCKED | Auditors, designers |
| **TCIndustries_Governance_Framework.md** | Authority/provenance rules (from Matrix) | Governance | Auditors, designers |
| **TCIndustries_Evidence_Summary.md** | What evidence exists and what it showed | Evidence | Auditors, designers |
| **TCIndustries_State_Manifest.json** | Machine-readable project state | Snapshot | LLMs, automation |

### 7.3 Proposed Archive Structure

```
archive/
├── 2026-08-25_pre-consolidation/
│   ├── GDD_v1.0_Consolidated.md
│   ├── GDD_v1.1_Canonical_Baseline.md
│   ├── Authority_Provenance_Reconciliation_Matrix.md
│   ├── Canonical_Audit.md
│   ├── EIC_Design_Investigation.md
│   ├── EIC_Candidate_v0.1.md
│   ├── EIC_Falsification_Pass.md
│   ├── EIC_Comparative_Simulation_Specification_v0.1.md
│   ├── EIC_Comparative_Simulation_Results_v0.1.md
│   ├── EIC_Functional_Shapes_v0.1.md
│   ├── EIC_Targeted_γ_Closure_v0.2.md
│   ├── EIC_Targeted_γ′_Re-test_Results_v0.1.md
│   ├── EIC_Comparative_Simulation_Execution_Results_v0.1.md
│   ├── EIC_Human_Decision_Brief_v0.1.md
│   └── EIC_Human_Rulings_Architecture_C_Falsification.md
└── manifest.json
```

### 7.4 Document Relationships

```
Preamble (entry point)
    │
    ├── Master GDD v1.2 (current design)
    │       │
    │       ├── references → Governance Framework
    │       ├── depends on → Human Rulings (for HUMAN-LOCKED items)
    │       └── evidence → Evidence Summary
    │
    ├── Human Rulings (all human decisions)
    │
    ├── Governance Framework (authority/provenance rules)
    │
    ├── Evidence Summary (EIC results and other evidence)
    │
    └── State Manifest (machine-readable snapshot)
            │
            └── references → Archive (for historical material)
```

---

## 8. Active Documentation Set (Detailed)

### 8.1 TCIndustries_Preamble.md

**Purpose**: Single entry point for anyone (human or LLM) starting with TCIndustries.

**Contents**:
- One-paragraph project description
- Document map (what to read in what order)
- Quick status summary (what's decided, what's in progress)
- Link to State Manifest
- Link to Archive

**Authority**: Informational only; not design authority

**Update Frequency**: When project state changes significantly

**LLM Reading Priority**: **MANDATORY FIRST READ**

**Must Not Contain**: Game design decisions; detailed specifications

---

### 8.2 TCIndustries_Master_GDD_v1.2.md

**Purpose**: Single, current, complete GDD.

**Contents**:
- All LOCKED, DERIVED CONSTRAINT, and PROPOSED design content
- Current status tags for all statements
- Explicit references to Human Rulings where relevant (not duplication)
- Open question register
- Assumption register

**Authority**: Working canonical reference

**Update Frequency**: When design decisions change or are added

**LLM Reading Priority**: **MANDATORY SECOND READ**

**Must Not Contain**: Historical material; evidence records; audit analysis

**Relationship to prior versions**: This is the **only** active GDD. All prior versions are archived.

---

### 8.3 TCIndustries_Human_Rulings.md

**Purpose**: Single authoritative record of all human project decisions.

**Contents**:
- Every human ruling with ID, date, decision text, and scope
- Grouped by topic (Vision, EIC, Combat, etc.)
- Explicit dependency tracking where rulings modify GDD statements
- Chronological log for historical context

**Authority**: HUMAN-LOCKED (decisions themselves)

**Update Frequency**: When new rulings are made

**LLM Reading Priority**: **MANDATORY** when auditing authority

**Must Not Contain**: LLM proposals; design analysis; evidence

**Relationship to GDD**: GDD references this document for HUMAN-LOCKED status; this document does not duplicate GDD content.

---

### 8.4 TCIndustries_Governance_Framework.md

**Purpose**: Authority and provenance rules.

**Contents**:
- Status vocabulary (LOCKED, DERIVED CONSTRAINT, PROPOSED, etc.)
- Authority classes (A–G from Authority Matrix)
- Governance rules (GR-001–003)
- Rules for distinguishing human approval from LLM inference
- Historical/non-canonical boundary rules

**Authority**: Governance (not design)

**Update Frequency**: When governance rules change (rare)

**LLM Reading Priority**: **OPTIONAL but recommended** before auditing

**Must Not Contain**: Game design decisions

**Relationship to GDD**: GDD references this framework; this document does not duplicate GDD status tags.

---

### 8.5 TCIndustries_Evidence_Summary.md

**Purpose**: Summary of all evidence relevant to current design.

**Contents**:
- EIC investigation summary (what was tested, what failed, what remains open)
- Prototype evidence status (when available)
- Simulation evidence summary
- Audit evidence summary
- Explicit "what this evidence means for current design" statements

**Authority**: Evidence (not design)

**Update Frequency**: When new evidence arrives

**LLM Reading Priority**: **OPTIONAL but recommended** for understanding design rationale

**Must Not Contain**: Design decisions

**Relationship to Archive**: This summarises evidence; the full evidence documents remain in the archive.

---

### 8.6 TCIndustries_State_Manifest.json

**Purpose**: Machine-readable snapshot of current project state.

**Contents**:

```json
{
  "manifestVersion": "1.0",
  "timestamp": "2026-08-25T...",
  "activeDocuments": [
    {
      "filename": "TCIndustries_Master_GDD_v1.2.md",
      "version": "1.2",
      "status": "Working Canonical Reference"
    }
    // ... more active documents
  ],
  "currentState": {
    "vision": { "locked": ["VIS-001", ...], "proposed": [...] },
    "eic": {
      "architectureC": "RETIRED",
      "humanRulings": ["HD-EIC-01", ...],
      "openQuestions": ["OQ-001", ...]
    },
    // ... other topic areas
  },
  "humanRulings": [
    { "id": "HD-EIC-01", "summary": "Layered unit of interdependence", "date": "2026-08-25" }
  ],
  "archiveLocation": "archive/2026-08-25_pre-consolidation/",
  "openQuestions": [
    { "id": "OQ-001", "question": "Exact specialisation budget / skill-cap model?" }
  ]
}
```

**Authority**: Informational snapshot; no design authority

**Update Frequency**: When project state changes

**LLM Reading Priority**: **OPTIONAL** for LLMs; **primary** for automation

**Must Not Contain**: Duplicate of GDD content

---

## 9. Archive / Provenance Structure

### 9.1 Archive Purpose

The archive preserves all historical, superseded, and evidence documents for provenance and auditability without allowing them to contaminate current canon.

### 9.2 Archive Organisation

```
archive/
├── 2026-08-25_pre-consolidation/    # Snapshot before consolidation
│   ├── GDD/
│   │   ├── Master_GDD_v1.0_Consolidated.md
│   │   ├── Master_GDD_v1.1_Canonical_Baseline.md
│   │   └── ...
│   ├── Governance/
│   │   ├── Authority_Provenance_Reconciliation_Matrix.md
│   │   └── ...
│   ├── EIC/
│   │   ├── Design_Investigation.md
│   │   ├── Candidate_v0.1.md
│   │   ├── Falsification_Pass.md
│   │   └── ...
│   ├── Human_Rulings/
│   │   ├── Decision_Brief_v0.1.md
│   │   └── ...
│   └── manifest.json
├── archive_manifest.json            # Overall archive index
└── README.md                        # Archive usage instructions
```

### 9.3 Archive Manifest

Each archive snapshot should include a manifest describing:
- Date of snapshot
- Which documents are included
- Relationship between documents
- Status of each document at time of archiving
- Supersession relationships

---

## 10. LLM Audit Reading Protocol

### 10.1 Standard Audit Reading Order

| Step | Document | Reading Time | Purpose |
|---|---|---|---|
| 1 | **Preamble** | ~1 min | Understand project, document map, entry point |
| 2 | **State Manifest** | ~2 min | Machine-readable state; quick summary |
| 3 | **Master GDD v1.2** | ~30–60 min | Current design content |
| 4 | **Human Rulings** | ~10 min | Authority verification for LOCKED items |
| 5 | **Evidence Summary** | ~10 min | Evidence status and rationale |
| 6 | **Governance Framework** | ~10 min | Authority rules (if needed for audit) |
| 7 | **Archive (as needed)** | Variable | Historical/provenance verification |

### 10.2 Authority Verification Protocol

1. **Identify the statement** in the GDD (get its ID, e.g., VIS-001, CRFT-006)
2. **Check its status** in the GDD (LOCKED / DERIVED CONSTRAINT / PROPOSED / etc.)
3. **If LOCKED**: Verify against Human Rulings document to confirm human approval
4. **If DERIVED CONSTRAINT**: Verify derivation trace to parent LOCKED principle
5. **If PROPOSED**: Note that no human approval exists
6. **Report uncertainty**: If a statement's authority cannot be traced, report it

### 10.3 Change Detection Protocol

1. Compare current State Manifest against previous version
2. Identify:
   - New documents
   - Changed documents
   - Status changes
   - New human rulings
   - Resolved open questions
3. Report delta

### 10.4 Audit Report Template

```markdown
# TCIndustries Audit Report
## Date: YYYY-MM-DD
## Auditor: [LLM/Person ID]

### 1. Current State Summary
- [one-paragraph summary]

### 2. Active Documents Inspected
- [list]

### 3. Authority Findings
- [statements with authority status]

### 4. Open Questions
- [currently unresolved]

### 5. Changes Since Previous Audit
- [if applicable]

### 6. Evidence Status
- [what evidence exists and what it shows]

### 7. Uncertainties / Items Requiring Human Verification
- [list]

### 8. Recommendations
- [if any]
```

---

## 11. Recommended Standard Metadata

### 11.1 Document Header Metadata

Each active document should include:

| Field | Required? | Purpose |
|---|---|---|
| `documentId` | Yes | Unique stable identifier |
| `title` | Yes | Full title |
| `version` | Yes | Semantic version |
| `status` | Yes | Working Canonical Reference / Governance / Evidence / etc. |
| `authority` | Yes | HUMAN-LOCKED / Governance / Evidence / etc. |
| `date` | Yes | Current document date |
| `supersedes` | No | Document(s) replaced by this one |
| `supersededBy` | No | Document(s) that replace this one (for active documents only) |
| `controllingReferences` | No | Documents that control/constrain this one |
| `audience` | Yes | Primary consumers |
| `readingPriority` | Yes | MANDATORY / RECOMMENDED / OPTIONAL / ARCHIVE |

### 11.2 Statement-Level Metadata

For important statements (LOCKED, DERIVED CONSTRAINT), include:

| Field | Required? | Purpose |
|---|---|---|
| `id` | Yes | Unique identifier |
| `status` | Yes | LOCKED / DERIVED CONSTRAINT / PROPOSED / etc. |
| `authority` | Yes | Human ruling ID or derivation source |
| `dateApproved` | For LOCKED | Date of human approval |
| `derivesFrom` | For DERIVED CONSTRAINT | Parent principle(s) |

### 11.3 Human Ruling Metadata

| Field | Required? | Purpose |
|---|---|---|
| `rulingId` | Yes | Unique identifier (e.g., HD-EIC-01) |
| `date` | Yes | Date of ruling |
| `decision` | Yes | The decision itself |
| `scope` | Yes | What the decision applies to |
| `modifies` | No | GDD statements modified by this ruling |
| `status` | Yes | HUMAN-LOCKED |

---

## 12. Migration / Consolidation Plan

### Phase 0: Preparation (Human Approval Required)

1. Human reviews and approves the consolidation plan
2. Decision is made on whether to proceed

### Phase 1: Archive Existing Documents

1. Create archive directory: `archive/2026-08-25_pre-consolidation/`
2. Copy all existing documents into the archive
3. Create archive manifest

**Documents to archive:**
- `TCIndustries_Master_GDD_v1.0_Consolidated.md`
- `TCIndustries_Master_GDD_v1.1_Canonical_Baseline.md`
- `TCIndustries_Authority_Provenance_Reconciliation_Matrix.md` (governance content to be extracted)
- `TCIndustries_Master_GDD_v1.0_Canonical_Audit.md`
- All EIC documents except those needed for Evidence Summary
- `TCIndustries_EIC_Human_Decision_Brief_v0.1.md`
- `TCIndustries_EIC_Human_Rulings_2026-08-25.md`
- `TCIndustries_EIC_Human_Rulings_Architecture_C_Falsification.md`

### Phase 2: Consolidate GDD

1. Confirm that v1.1.1 is the current working reference
2. Produce v1.2, which:
   - Renames to a clean filename (no version in filename)
   - Includes only current material (no historical notes in body)
   - Has a clean header with metadata
   - References Human Rulings document rather than duplicating rulings
   - Explicitly states that earlier GDD versions are archived

### Phase 3: Consolidate Human Rulings

1. Extract all human rulings from:
   - `TCIndustries_EIC_Human_Rulings_2026-08-25.md` (HD-EIC-01–04)
   - `TCIndustries_EIC_Human_Rulings_Architecture_C_Falsification.md` (HD-EIC-05–08)
2. Consolidate into a single `TCIndustries_Human_Rulings.md`
3. Order chronologically
4. Include explicit references to which GDD statements they modify

### Phase 4: Extract Governance Framework

1. Extract authority/governance content from:
   - `TCIndustries_Authority_Provenance_Reconciliation_Matrix.md`
   - GDD Sections 3–4 (How to Read, Authority Rules)
2. Consolidate into `TCIndustries_Governance_Framework.md`

### Phase 5: Create Evidence Summary

1. Extract EIC conclusions from:
   - `TCIndustries_EIC_Targeted_γ′_Adversarial_Re-test_Results_v0.1.md`
   - `TCIndustries_EIC_Human_Rulings_Architecture_C_Falsification.md`
2. Summarise what was tested, what failed, what remains open
3. Produce `TCIndustries_Evidence_Summary.md`

### Phase 6: Create Preamble

1. Write a short (1–2 page) entry document
2. Include project description, document map, reading order

### Phase 7: Create State Manifest

1. Produce machine-readable JSON manifest
2. Include current project state summary

### Phase 8: Update Archive Manifest

1. Document the consolidation in the archive manifest
2. Record that new active documents exist
3. Record supersession relationships

---

## 13. Risks and Trade-offs

### 13.1 Risks

| Risk | Likelihood | Impact | Mitigation |
|---|---|---|---|
| **Losing historical provenance** | Low | Medium | Archive all original documents; maintain manifest |
| **Human rulings misinterpreted** | Low | High | Consolidate rulings with clear scope statements |
| **Archive not consulted for provenance** | Medium | Low | Clear documentation of archive location and purpose |
| **New LLM reads archive first** | Medium | Medium | Preamble explicitly states reading order |
| **Consolidation introduces errors** | Low | High | Phase 0 approval; retain original documents |
| **EIC evidence summarised incorrectly** | Low | Medium | Cross-check summary against original evidence documents |

### 13.2 Trade-offs

| Trade-off | Choice | Rationale |
|---|---|---|
| **Single GDD vs multiple versions** | Single GDD; archive others | Single source of truth for current design |
| **EIC detail vs summary** | Summary active; full details archived | Active readers need conclusions, not full methodology |
| **Human rulings vs GDD** | Separate documents with references | Human rulings are decision records, not design content |
| **Machine-readable vs human-readable** | Both: manifest + human documents | Serves both audiences |
| **Provenance detail vs brevity** | Archive preserves detail; active documents are brief | Archive exists for auditors who need detail |
| **Single entry point vs distributed** | Single Preamble | Independent LLM must know where to start |

---

## 14. Human Authorization Gates

### Gate 1 — Consolidation Approval

| Question | Required |
|---|---|
| Does the project owner approve the consolidation approach? | **YES** |
| Does the project owner approve archiving all prior GDD versions? | **YES** |
| Does the project owner approve the active document set? | **YES** |

### Gate 2 — GDD v1.2 Authority

| Question | Required |
|---|---|
| Does the project owner approve GDD v1.2 as the single active GDD? | **YES** |
| Does the project owner confirm that v1.1.1 content is accurate? | **YES** |
| Does the project owner approve the active document metadata? | **YES** |

### Gate 3 — Human Rulings Consolidation

| Question | Required |
|---|---|
| Does the project owner approve consolidated Human Rulings document? | **YES** |
| Does the project owner confirm all rulings are correctly captured? | **YES** |
| Does the project owner approve ruling IDs and scope statements? | **YES** |

### Gate 4 — State Manifest

| Question | Required |
|---|---|
| Does the project owner approve the State Manifest schema? | **YES** |
| Does the project owner approve the current state summary? | **YES** |

### Gate 5 — Archive

| Question | Required |
|---|---|
| Does the project owner approve the archive organisation? | **YES** |
| Does the project owner approve the archive manifest format? | **YES** |

---

## 15. Target State

### 15.1 Active Documents (After Consolidation)

| Document | Purpose | Authority | Status |
|---|---|---|---|
| `TCIndustries_Preamble.md` | Entry point | Informational | Active |
| `TCIndustries_Master_GDD_v1.2.md` | Current design | Working canonical | Active |
| `TCIndustries_Human_Rulings.md` | Human decisions | HUMAN-LOCKED | Active |
| `TCIndustries_Governance_Framework.md` | Authority/provenance | Governance | Active |
| `TCIndustries_Evidence_Summary.md` | Evidence summary | Evidence | Active |
| `TCIndustries_State_Manifest.json` | Machine-readable state | Snapshot | Active |

### 15.2 Archive Structure

```
archive/
├── 2026-08-25_pre-consolidation/
│   ├── GDD/
│   ├── Governance/
│   ├── EIC/
│   ├── Human_Rulings/
│   └── manifest.json
├── archive_manifest.json
└── README.md
```

### 15.3 Document Count Reduction

| Category | Before | After | Reduction |
|---|---|---|---|
| Active documents | 17 | 6 | 65% |
| GDD versions | 3 | 1 | 67% |
| EIC documents | 8 | 1 (summary) + archive | 88% |
| Human rulings documents | 3 | 1 | 67% |

---

## 16. Open Questions / Information Not Established by the Corpus

The following questions are not answered by the supplied documentation:

### 16.1 About the Project

| Question | Status |
|---|---|
| Who is the project owner? | Not established in corpus |
| What is the project's current phase? | Design / pre-development (inferred) |
| What are the project's production timelines? | Not established |
| Is there a live prototype? | Referenced but not inspected (AS-001) |

### 16.2 About Documentation

| Question | Status |
|---|---|
| What version control system is used? | Not established |
| What is the document workflow (editing, review, approval)? | Governance rules exist; workflow not specified |
| Who is the primary maintainer of the GDD? | Not established |
| How are status changes documented in practice? | Status patches, but process not fully specified |

### 16.3 About Authority

| Question | Status |
|---|---|
| What is the human approval process for LOCKED items? | Described in principle; detailed process not specified |
| Who can make human rulings? | Project owner (inferred); not explicitly stated |
| How are human rulings communicated to the LLM system? | Via documents in this corpus |

### 16.4 About Next Steps

| Question | Status |
|---|---|
| What happens after Architecture C is retired? | Human boundaries statement pending (HD-EIC-07) |
| Is there a planned "Architecture D" investigation? | Not authorised; contingent on human boundaries |
| What is the EIC's next step? | Human must supply structural boundaries (HD-EIC-07) |

These open questions are not gaps to be filled; they represent information not established in the supplied corpus. Any documentation architecture should preserve these as known uncertainties.

---

## Appendices

### Appendix A: Document Dependency Map

```
Master GDD v1.1.1
    ├── derives authority from → Authority Matrix
    ├── applies → Human Rulings (HD-EIC-01–08)
    ├── implements → Canonical Audit recommendations
    └── supersedes → Master GDD v1.1

Human Rulings
    ├── HD-EIC-01–04 ← Human Decision Brief
    ├── HD-EIC-05–08 ← Architecture C Falsification Gate
    └── modify → GDD status tags

EIC Documents
    ├── Investigation → Candidate v0.1 → Falsification Pass
    ├── Functional Shapes → γ Closure → γ′ Re-test
    ├── Comparative Simulation Specification → Results
    └── All → Evidence Summary (to be created)

Authority Matrix
    ├── controls → GDD status vocabulary
    ├── GR-001–003 → Governance rules
    └── applied in → GDD v1.1 status reclassification
```

### Appendix B: Document Status Summary Table

| Document | Current Status | Recommended Status | Action |
|---|---|---|---|
| Master GDD v1.0 | Superseded | ARCHIVE | Move to archive |
| Master GDD v1.1 | Superseded | ARCHIVE | Move to archive |
| Master GDD v1.1.1 | Active | Consolidated into v1.2 | Extract; archive original |
| Canonical Audit | Active reference | ARCHIVE | Move to archive |
| Authority Matrix | Active governance | Extracted into Governance Framework; original archived | Extract; archive original |
| EIC Design Investigation | Evidence | EVIDENCE (archive) | Move to archive |
| EIC Candidate v0.1 | Superseded | ARCHIVE | Move to archive |
| EIC Falsification Pass | Evidence | ARCHIVE | Move to archive |
| EIC Simulation Specification | Evidence | ARCHIVE | Move to archive |
| EIC Simulation Results | Evidence | ARCHIVE | Move to archive |
| EIC Functional Shapes | PROVISIONAL (superseded) | ARCHIVE | Move to archive |
| EIC γ Closure | PROVISIONAL (superseded) | ARCHIVE | Move to archive |
| EIC γ′ Re-test Results | Evidence | ARCHIVE | Move to archive |
| Human Decision Brief | Human gate record | ARCHIVE | Move to archive |
| Human Rulings (HD-EIC-01–04) | HUMAN-LOCKED | Consolidated into Human Rulings | Extract; archive original |
| Human Rulings (Architecture C) | HUMAN-LOCKED | Consolidated into Human Rulings | Extract; archive original |

---

## Human Authorization Required

The following actions require explicit project owner approval before implementation:

1. **Archiving all GDD versions except the current one** — Preserves history while reducing active duplication. Approval required to ensure no loss of reference material.

2. **Consolidating all human rulings into a single document** — Ensures single source of truth for human decisions. Approval required to ensure no rulings are omitted or misinterpreted.

3. **Consolidating all EIC documents into an archive** — Preserves detailed evidence while reducing active sprawl. Approval required to ensure summary captures key evidence correctly.

4. **Creating the State Manifest** — New document type; requires approval of schema and content.

5. **Creating the Evidence Summary** — New document type; requires approval of content and conclusions.

6. **Creating the Preamble** — New entry point; requires approval of content and reading order.

7. **Adopting the LLM Audit Reading Protocol** — Changes how documents are consumed; requires approval if it affects human readers as well.

8. **Adopting the Standard Metadata** — Changes document headers; requires approval of metadata fields and values.

**No action should be taken without explicit human approval of each item.**