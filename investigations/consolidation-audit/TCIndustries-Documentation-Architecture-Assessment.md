# TCIndustries Documentation Consolidation and Cross-LLM Auditability Assessment

**Assessment Date:** 2026-08-26  
**Auditor Role:** Independent Documentation Architect and Evidence/Governance Auditor  
**Corpus Analyzed:** `TCIndustries-project-files-v1-merged.txt` (461,694 chars, 8,037 lines, 19 document boundaries identified)  
**Authority Created by This Assessment:** None. This report creates no new game-design decisions, does not promote proposals, and does not modify canon.
**Author:** Author LLM Unknown

---

## 1. Executive Verdict

TCIndustries possesses unusually strong governance discipline for a long-lived design project. The introduction of **HUMAN-LOCKED versus DERIVED CONSTRAINT**, the explicit governance rules GR-001 to GR-003, and the Authority & Provenance Reconciliation Matrix provide a genuine foundation for cross-LLM auditability. No other project at this stage typically achieves this level of authority hygiene.

The current documentation nevertheless fails the three success tests in its present sprawled form:

- **Test 1 — State Reconstruction: MARGINAL PASS with high cost.** A new LLM *can* determine current state, but only by reconciling three overlapping Master GDD versions (v1.0 Consolidated, v1.1 Canonical Baseline, v1.1.1 Status Patch) plus two Human Rulings records. The active set is not machine-identifiable without prior knowledge.

- **Test 2 — Audit Reproducibility: FAIL without protocol.** Two independent LLMs will produce substantially different findings because there is no stable reading order, no audit manifest, no standardized machine-readable metadata envelope, and 14 EIC investigation/evidence documents that look canonical to a naive parser.

- **Test 3 — Historical Containment: PASS by discipline, FAIL by structure.** The governance *rules* prevent historical contamination (GR-002 is explicit), but the *structure* invites contamination: duplicate simulation specifications, three Master GDD copies in the same corpus, and no archival boundary.

The remedy is not to rewrite the Master GDD for style. It is to establish a small, explicitly ordered active set, move all other material into a preserved archive with containment markers, and introduce a machine-readable manifest and metadata envelope that enforces the existing governance vocabulary.

---

## 2. Current Documentation Assessment

### 2.1 What exists

The corpus contains 19 logical documents across four functional families:

1.  **Governance / Provenance (1 doc):** Authority & Provenance Reconciliation Matrix v1.0
2.  **Canonical Design Reference (4 docs):** v1.0 Consolidated (Consolidation Candidate), v1.0 Canonical Audit, v1.1 Canonical Baseline, v1.1.1 Status Patch
3.  **Human Authority Records (2 docs):** Human Rulings 2026-08-25 (HD-EIC-01 to HD-EIC-04) and Human Rulings Architecture C Falsification Gate (HD-EIC-05 to HD-EIC-08)
4.  **EIC Investigation & Evidence (12 docs):** Economic Interdependence Core Design Investigation, EIC Candidate v0.1, Falsification Pass, Comparative Simulation Results v0.1.1, Execution & Results v0.1, Simulation Specification v0.1 (two versions), Functional Shapes Refinement v0.2, Human Decision Brief v0.1, Minimum Simulation Results v0.1, Provisional Functional Shapes v0.1, Targeted γ′ Re-test Results v0.1

Direct evidence from file analysis:
- Total LOCKED mentions: 514; HUMAN-LOCKED: 98; DERIVED CONSTRAINT: 70; PROPOSED: 215; TBD: 209. This confirms status vocabulary is actively used.
- All 17 OQ identifiers (OQ-001 through OQ-017) appear, with OQ-001 appearing in 16 documents.
- No audit manifest exists (search for "manifest" returned zero results).
- Metadata fields are inconsistent: Document Type appears in 13/19 docs, Version in 18/19, Authority in 16/19.

### 2.2 What is genuinely authoritative

Based on explicit authority statements and controlling references (not self-declaration):

- **Current working canonical reference:** `TCIndustries_Master_GDD_v1.1.1_Status_Patch.md` (Doc 17). Direct evidence: Authority Statement "This document is the current working canonical design reference" and lineage "Prior Version 1.1 — Canonical Baseline". It applies human rulings from the Reconciliation Matrix. It supersedes v1.1 Baseline, which superseded v1.0 Consolidated.
- **Governance foundation:** `TCIndustries_Authority_Provenance_Reconciliation_Matrix.md` (Doc 0). Authority created: None, but defines GR-001, GR-002, GR-003 and Authority Classes A-G. This is the only document that defines what LOCKED may mean.
- **Human authority:** `TCIndustries_EIC_Human_Rulings_2026-08-25.md` (Doc 10) and `Architecture C Falsification Gate` (Doc 11). Direct evidence: "Authority: Explicit human project rulings" and "HUMAN-LOCKED (unit of validation only)" with explicit "Explicitly does not determine" scoping. HD-EIC-01 through HD-EIC-08 are the only HUMAN-LOCKED decisions for the EIC program.
- **Supporting governance:** `TCIndustries_Master_GDD_v1.0_Canonical_Audit.md` (Doc 15) — the audit that identified over-promoted LOCKED and recommended DERIVED CONSTRAINT vocabulary. Status: Audit record, not design authority.

All other 12 EIC documents explicitly state **Authority: None** and **PROPOSED / Non-Canonical / Simulation Assumptions Only**. Direct evidence: "This document creates no new LOCKED design decisions", "All content is PROPOSED / simulation assumption unless explicitly restating a prior HUMAN-LOCKED ruling".

### 2.3 What should become archive/provenance

- v1.0 Consolidated Master GDD (Doc 16) — superseded working reference, now provenance only.
- v1.1 Canonical Baseline (Doc 18) — superseded by v1.1.1 Status Patch, but needed for lineage. Should become archive.
- v1.0 Canonical Audit (Doc 15) — audit record, remains valuable but not active canon.
- All 12 EIC investigation and evidence documents — simulation evidence, falsification analysis, functional shapes. Must be preserved as evidence/provenance with containment.

---

## 3. Current Documentation Inventory

| Index | Inferred Filename | Document Type (as declared) | Version | Authority Claim | Actual Role per Evidence |
|---|---|---|---|---|---|
| 0 | Authority_Provenance_Reconciliation_Matrix.md | Governance / Provenance | 1.0 | None — governance only | **Governance root** |
| 1 | Economic_Interdependence_Core_Design_Investigation.md | Design Investigation (Non-Canonical) | 1.0 | None | Investigation / evidence |
| 2 | EIC_Candidate_v0.1.md | Structural Candidate / Falsification Spec | 0.1 | None | Proposed candidate |
| 3 | EIC_Candidate_v0.1_Falsification_Pass.md | Falsification / Adversarial Analysis | 1.0 | None | Evidence |
| 4 | EIC_Comparative_Simulation_Results_v0.1.1.md | Qualitative Adversarial Structural Evaluation (Non-Canonical) | 0.1.1 | None | Evidence |
| 5 | EIC_Comparative_Simulation_Execution_Results_v0.1.md | Adversarial Evidence Report (Non-Canonical) | 0.1 | None | Evidence |
| 6 | EIC_Comparative_Simulation_Specification_v0.1 (post-shapes) | Adversarial Simulation Specification (Non-Canonical) | 0.1 post-Functional-Shapes | None | Evidence spec — supersedes earlier spec |
| 7 | EIC_Comparative_Simulation_Specification_v0.1 (pre-shapes) | Simulation Specification (Non-Canonical) | 0.1 | None | Superseded spec — duplicate |
| 8 | EIC_Functional_Shapes_Refinement_v0.2.md | Narrow Provisional Refinement | 0.2 | None | Provisional refinement |
| 9 | EIC_Human_Decision_Brief_v0.1.md | Human Decision Gate (Non-Canonical) | 0.1 | None | Decision brief — not authority |
| 10 | EIC_Human_Rulings_2026-08-25.md | Human Authority Record | 1.0 | Explicit human rulings | **Human authority** HD-EIC-01 to 04 |
| 11 | EIC_Human_Rulings_Architecture_C_Falsification.md | Human Authority Record | 1.0 | Explicit human rulings | **Human authority** HD-EIC-05 to 08 |
| 12 | EIC_Minimum_Comparative_Simulation_Results_v0.1.md | Simulation Evidence Report (Non-Canonical) | 0.1 | None | Evidence — inconclusive/underdetermined |
| 13 | EIC_Provisional_Functional_Shapes_v0.1.md | Provisional Functional Model (Simulation Assumptions Only) | 0.1 | None | Provisional shapes — simulation assumptions |
| 14 | EIC_Targeted_γ′_Adversarial_Retest_Results_v0.1.md | Adversarial Evidence Report (Non-Canonical) | 0.1 | None | Evidence — falsification |
| 15 | Master_GDD_v1.0_Canonical_Audit.md | Canonical GDD Audit | — | None — audit | **Audit record** |
| 16 | Master_GDD_v1.0_Consolidated.md | Consolidated Master GDD (first-pass) | 1.0 | "current primary design authority" (now superseded) | Superseded — provenance |
| 17 | Master_GDD_v1.1.1_Status_Patch.md | Working Canonical Reference | 1.1.1 | Current working canonical (human rulings applied) | **Current canonical** |
| 18 | Master_GDD_v1.1_Canonical_Baseline.md | Working Canonical Reference | 1.1 | Current working canonical (now superseded by 1.1.1) | Superseded — provenance |

Key inference vs direct evidence: The role assignments above are based on explicit "Authority Statement", "Status", "Controlling References", and "This document supersedes" sentences, not on file name or self-declared canonicity. Doc 16 claimed to supersede first-pass docs at the time of its creation, but Doc 18 and Doc 17 explicitly supersede it in turn — temporal chain is directly evidenced in lineage tables.

---

## 4. Redundancy and Sprawl Analysis

### 4.1 Duplication

- **Master GDD triplication:** Docs 16, 17, 18 share 425-490 overlapping lines (84-91% line overlap by set intersection). They represent the same document at three stages of governance hygiene. For an LLM without temporal ordering, this is the single largest source of authority confusion. A naive LLM will treat all three as equally current and may select the oldest (v1.0) because it contains more PROPOSED detail that looks useful.

- **Simulation specification duplication:** Doc 6 and Doc 7 both titled "EIC Comparative Simulation Specification v0.1" with identical version numbers but different dates/contexts. Doc 6 explicitly states "This document supersedes the earlier Comparative Simulation Specification". Without that sentence being machine-parsed, two specifications compete.

- **Human Rulings duplication appearance:** Docs 10 and 11 share identical date (2026-08-25) and similar metadata but are distinct sequential rulings. Doc 11 correctly lists Doc 10 as controlling prior document. However, a new LLM may treat them as conflicting duplicates unless the HD-EIC identifier sequence is understood as cumulative.

### 4.2 Sprawl mechanisms observed

1.  **Evidence accumulation without archival boundary:** Each EIC falsification cycle produced a new spec, results, and refinement. All remain in the same directory as canonical material. The investigation produced valuable scientific practice (falsification over rhetoric), but left 12 non-canonical documents that outnumber canonical documents 3:1.

2.  **Version proliferation without manifest:** 19 documents share date 2026-08-25, making temporal ordering impossible from date alone. Only lineage sentences establish order. An LLM that relies on file modification time will fail.

3.  **Open questions scattered:** OQ-001 appears in 16 documents, OQ-012 in 9, but no single authoritative OQ registry exists in the active set. Master GDD contains "Remaining open design decisions" list (10 items) and "Remaining Human Decisions" list (11 items), but EIC docs add OQ-001 to OQ-017 without a unified status table.

4.  **Provenance mentions everywhere:** Every document contains "provenance" 2-22 times, but only three documents (16,17,18) have a formal Provenance and Change Notes section with change scope. This dilutes provenance discipline.

### 4.3 Information difficult to locate for an independent LLM

Direct evidence from corpus scan:

- **What is currently HUMAN-LOCKED?** Requires reading Doc 0 (authority classes), Doc 17 (status patch table of 7 confirmations), and Docs 10-11 (HD-EIC-01 to 08). No single file lists all HUMAN-LOCKED statements across GDD and EIC program.

- **What is explicitly NOT CANONICAL?** The historical non-canonical boundary list (12 items: "Factories can never produce Exceptional-tier output", "Vitality / Stamina / Focus pools", etc.) exists only in Doc 17 Provenance section. An LLM that reads only Doc 18 (v1.1 Baseline) will miss it.

- **What changed between versions?** Change scope tables exist in Docs 17 and 18, but not in a machine-diffable format. No changelog with identifiers.

- **What evidence exists for Architecture C retirement?** Requires traversing Doc 2 → Doc 3 → Doc 12 → Doc 9 → Doc 10 → Doc 13 → Doc 6 → Doc 4/5/14 → Doc 11. The chain is documented via controlling references, but only if the LLM follows them in order.

---

## 5. Authority and Provenance Risks

### 5.1 Potentially conflicting authority

| Conflict | Evidence | Risk Level |
|---|---|---|
| Doc 16 vs Doc 17 vs Doc 18 all claim "current working canonical" | Doc 16 Authority Statement: "current primary design authority" (dated earlier in lineage); Doc 17 and 18 also claim current | **High** — three concurrent canonical claims |
| Doc 6 vs Doc 7 same version number | Both v0.1, but Doc 6 says it supersedes earlier spec | **Medium** — simulation reproducibility fails if wrong spec used |
| HD-EIC-01 to 04 vs HD-EIC-05 to 08 | Same date, overlapping authority, but sequential | **Low** — correctly chained via controlling references, but requires HD identifier parsing |
| Authority review queue (WRLD-001, PLR-001, PLR-003, CRFT-001, SAFE-001) in Doc 0 | Marked "Authority review required" while still LOCKED in v1.1 | **Resolved** in Doc 17 which confirms them as HUMAN-LOCKED, but an LLM reading only Doc 18 will see unresolved queue |

### 5.2 Provenance risks

- **GR-001 violation risk via LLM consensus:** The corpus itself warns "LLM agreement does not constitute human approval" in multiple places. Yet without a manifest, a future LLM performing frequency analysis will count VIS-001 appearing in 6 docs as stronger than HD-EIC-05 appearing in 1 doc, inverting actual authority.

- **Prototype as authority:** Docs 16-18 contain PROTOTYPE references with note "No prototype was inspected". The Evidence Reconciliation Pass remains open. Without an explicit evidence manifest, a future LLM that finds prototype code may treat prototype behaviour as authority, violating PROT-001.

- **Historical SWG material:** Marked HISTORICAL in vocabulary, but appears in explanations throughout. GR-002 protects against automatic reintroduction, but only if the LLM reads GR-002.

### 5.3 Redundant material that should not be active

All EIC evidence documents (Docs 1-9, 12-14) are marked non-canonical and contain explicit disclaimers. Direct evidence: "This document creates no canonical design", "All content is PROPOSED / simulation assumption". They should not be in the active reading set for a state reconstruction audit. They belong in archive with evidence tags.

---

## 6. Cross-LLM Auditability Assessment

### A. State Reconstruction

**Current:** Can an independent LLM determine current project state without reading every historical document?

- **Verdict:** Marginally, but at high cost. The answer is spread across Doc 17 (GDD principles), Doc 0 (governance rules), Docs 10-11 (EIC human rulings), and Doc 18 lineage (if Doc 17 missing). No single document answers "What is the project currently?" Doc 17 Executive Summary does, but lacks EIC program state. Doc 11 contains post-falsification program state table but not GDD vision.

- **Gap:** No single "Project State" entry point. No manifest.

### B. Authority Reconstruction

**Current:** Can an independent LLM distinguish human decisions from LLM proposals, logical deductions, historical material, prototype evidence, assumptions, audit recommendations?

- **Verdict:** Yes, if it reads Doc 0 and Doc 17 carefully. The vocabulary is exemplary: HUMAN-LOCKED, DERIVED CONSTRAINT, PROPOSED, TBD, PROTOTYPE/EVIDENCE, DEFERRED, ASSUMPTION, HISTORICAL are defined in How to Read sections. Authority Classes A-G in Doc 0 distinguish explicit human ruling (A) from derived (C) from multi-source inference (E) from prototype (G).

- **Gap:** Vocabulary evolved between Doc 16 (no DERIVED CONSTRAINT) and Doc 17/18 (DERIVED CONSTRAINT added). An LLM reading only Doc 16 will lack the distinction. Also, DERIVED CONSTRAINT strictness rule GR-003 exists only in Doc 0, not repeated in GDD How to Read.

### C. Temporal Reconstruction

**Current:** Can an LLM determine what is current versus superseded?

- **Verdict:** Partially. Lineage fields exist: Doc 17 says Prior Version 1.1, Controlling Audit v1.0 Canonical Audit, Authority Matrix. Doc 6 says supersedes earlier spec. However, dates are all identical (2026-08-25), so temporal ordering depends on parsing natural-language lineage sentences, not machine-readable supersedes/superseded-by fields.

- **Gap:** No standardized version graph. No superseded-by field. No manifest with ordering.

### D. Provenance

**Current:** Can important current statements be traced to origin and status changes?

- **Verdict:** For GDD rules, yes via Rule IDs (VIS-001, etc.) and Provenance and Change Notes tables in Docs 17/18 that list reclassifications (LOCKED → DERIVED CONSTRAINT for CRFT-006, MFG-004, etc.) with parent principles. For EIC rulings, traceability table exists (HD-EIC-01 maps to Gate 1).

- **Gap:** Traceability is in prose tables, not machine-readable provenance chains. No per-statement provenance (e.g., VIS-001 derived from explicit project brief per Doc 0 table). No change identifiers.

### E. Cross-LLM Comparability

**Current:** Could two independent LLMs produce comparable audit results?

- **Verdict:** No, in current sprawl. Without a prescribed reading order, LLM A might read Doc 16 (v1.0 Consolidated, which has 53 PROPOSED items with more detail) and LLM B might read Doc 17 (v1.1.1 Status Patch, which demotes some items to PROPOSED). They will disagree on what is LOCKED. Similarly, LLM A might include Architecture C as leading hypothesis (Doc 2) while LLM B reads retirement (Doc 11).

- **Evidence of risk:** The corpus itself contains a failed assumption: Doc 16 calls itself primary authority, but is superseded. A frequency-based LLM will prefer Doc 16 because it is larger (45,749 chars vs 51,426 for Doc 17) and appears authoritative.

### F. Change Detection

**Current:** Can an auditor identify what changed since previous state?

- **Verdict:** Partially for GDD (status reclassifications table lists 5 items LOCKED → DERIVED CONSTRAINT). For EIC program, program state tables exist (Doc 10 §4, Doc 11 §3) showing stage transitions. However, no diff format, no change identifiers, no versioned manifest.

- **Gap:** No CHANGELOG.md or equivalent.

### G. Scope Control

**Current:** Does structure prevent historical or proposed material from silently becoming current canon?

- **Verdict:** Governance rules do (GR-002: Historical Proposal Boundary; Critical Rule: Never silently promote). Structural enforcement does not. All files are in same flat namespace with no archival directory, no file-level status header that is machine-readable, and no manifest that declares active vs archive.

- **Risk:** Future LLM with no governance context that ingests all files will merge PROPOSED mechanisms from Doc 2 (M1-M8) into canon because they look detailed and well-structured.

### H. Scalability

**Current:** Will architecture still work when project contains many detailed system specifications?

- **Verdict:** No. Current pattern is to add a new document per investigation cycle. At 19 docs for one subsystem (EIC), the project will have 100+ docs when Provenance & Reputation, Commerce & Services, Society, World & Logistics, and Combat are specified. Flat namespace with no module boundary will become un-auditable.

---

## 7. Recommended Documentation Architecture

**Design Principle:** Small active set (5-7 documents) + preserved archive + machine-readable manifest. Preserve existing status vocabulary. Do not rewrite Master GDD for style.

### Active Set Philosophy

- **Active = what an independent LLM must read to reconstruct current state.** No more than 7 documents.
- **Archive = everything else, preserved with containment markers.** Must be readable only when auditor explicitly needs provenance or evidence.
- **Manifest = single source of truth for reading order, current vs superseded, and authority graph.**

### Proposed Active Documents (Moderately Consolidated)

The proposed architecture consolidates 19 documents into 6 active documents plus 1 manifest. It does NOT merge unrelated content into one enormous document.

| # | Proposed Filename | Content Source | Rationale |
|---|---|---|---|
| 1 | `00_MANIFEST.md` | New — does not exist today | Machine-readable reading order, version graph, active/archive list |
| 2 | `01_GOVERNANCE.md` | Doc 0 + GR rules from Docs 17/18 How to Read + Critical Rule | Single governance root |
| 3 | `02_MASTER_GDD_v1.1.1.md` | Doc 17 (Status Patch) as base, with provenance appendix merged from Doc 18 if needed | Current working canonical reference — only one copy |
| 4 | `03_HUMAN_RULINGS.md` | Doc 10 + Doc 11 merged (HD-EIC-01 to 08) with traceability | All human authority in one place, cumulative |
| 5 | `04_EIC_PROGRAM_STATE.md` | New synthesis — not new design — summarizing program state table from Doc 11 §3 + gates from Doc 9 + what is forbidden | Replaces need to read 12 EIC docs to know program status |
| 6 | `05_OPEN_QUESTIONS.md` | Consolidated from Doc 17 Remaining open decisions + OQ registry + historical non-canonical boundary list | Single source for unresolved material |
| 7 | `06_CHANGELOG.md` | From Doc 17 and Doc 18 change scope tables + human rulings applied tables | Machine-readable change history |

**What is intentionally NOT in active set:**

- All simulation specifications, results, functional shapes, investigation prose, falsification passes — these become archive evidence.
- v1.0 Consolidated and v1.1 Baseline — these become archive provenance.
- v1.0 Canonical Audit — becomes archive audit record but is referenced by manifest.

---

## 8. Active Documentation Set

For each proposed active document:

### 00_MANIFEST.md

- **Filename:** `00_MANIFEST.md`
- **Purpose:** Deterministic entry point for any LLM auditor. Declares active set, archive set, version graph, reading order, and authority precedence.
- **Authority:** Governance — creates no game-design decisions. Authority Class D (Audit/Governance).
- **Audience:** LLMs first, humans second.
- **Belongs inside:** Document inventory with IDs, current vs superseded graph, controlling references chain, file hashes or version dates, active/archive classification, status vocabulary declaration, reading protocol reference.
- **Must not belong inside:** Game mechanics, proposals, simulation results.
- **Relationship:** Controls all other documents. Read first.
- **Authoritative:** Yes for documentation structure, no for game design.
- **Update frequency:** On every documentation change (add/remove/promote).
- **Should LLM read during audit:** Yes — always first.

### 01_GOVERNANCE.md

- **Filename:** `01_GOVERNANCE.md`
- **Purpose:** Single source for authority classes, status vocabulary, governance rules, and anti-contamination rules.
- **Authority:** Governance — defines how to interpret authority, not what game is. Authority Class D.
- **Audience:** LLMs and humans.
- **Belongs inside:** Authority Classes A-G table (from Doc 0), Status vocabulary with meanings (LOCKED, HUMAN-LOCKED, DERIVED CONSTRAINT, PROPOSED, TBD, PROTOTYPE/EVIDENCE, DEFERRED, ASSUMPTION, HISTORICAL), Governance Rules GR-001, GR-002, GR-003, Critical Rule (Never silently promote), Design Documentation Standard (Principle → Constraint → Invariant → Failure → Test).
- **Must not belong inside:** Specific game rules, EIC mechanisms, simulation data.
- **Relationship:** Referenced by all other active docs. Supersedes How to Read sections scattered across Master GDDs.
- **Authoritative:** Yes for governance interpretation.
- **Update frequency:** Rare — only when governance rules change via explicit human approval.
- **Should LLM read during audit:** Yes — second after manifest.

### 02_MASTER_GDD_v1.1.1.md

- **Filename:** `02_MASTER_GDD_v1.1.1.md` (retain version in filename for provenance, but manifest declares it as current)
- **Purpose:** Current working canonical design reference for TCIndustries game design and rules. Technical implementation belongs in future TDD.
- **Authority:** Contains HUMAN-LOCKED principles (A/B) and DERIVED CONSTRAINTS (C/D) as marked. All status tags authoritative per Doc 17.
- **Audience:** Humans and LLMs — primary game design source.
- **Belongs inside:** Vision, Pillars, Core Experience, World, Skills & Professions (principles only), Resources (principle only), Crafting (principle only), Manufacturing (principle only), Economy, Retail, Services, Combat (as one lifestyle), Society, Progression, Emergence, Safety, Glossary, Design Documentation Standard reference, Historical non-canonical boundary list (12 items), Remaining open design decisions list.
- **Must not belong inside:** EIC investigation prose, simulation specs, functional shapes, detailed taxonomies that are PROPOSED (those belong in archive or future system specs), numeric balance values.
- **Relationship:** Controlled by Governance doc, audited by Canonical Audit, supersedes v1.0 and v1.1. Referenced by all future system specs.
- **Authoritative:** Yes for game design principles — status tags define per-rule authority.
- **Update frequency:** Only via controlled correction with human approval. No LLM consolidation without human authorization (per Doc 17).
- **Should LLM read during audit:** Yes — third. This is the core state.

### 03_HUMAN_RULINGS.md

- **Filename:** `03_HUMAN_RULINGS.md`
- **Purpose:** Cumulative record of explicit human/project rulings with scope limits.
- **Authority:** Explicit human authority (Class A) for the decisions stated. Does not promote downstream mechanics to LOCKED.
- **Audience:** LLMs (for authority reconstruction) and humans (for decision tracking).
- **Belongs inside:** HD-EIC-01 to HD-EIC-08 with full Decision, Explicitly does not determine, Status, Implementation status, Established vs Not established table, Updated Programme State tables from Docs 10 and 11, Traceability tables, Governance note per GR-001.
- **Must not belong inside:** Proposed mechanisms, simulation assumptions, new design.
- **Relationship:** Controls EIC program state doc, references Governance doc for interpretation, referenced by Manifest as human authority source.
- **Authoritative:** Yes — for the rulings stated only.
- **Update frequency:** Only when new human rulings are issued.
- **Should LLM read during audit:** Yes — fourth, to understand EIC authority.

### 04_EIC_PROGRAM_STATE.md

- **Filename:** `04_EIC_PROGRAM_STATE.md`
- **Purpose:** Replace need to read 12 EIC docs by providing current program state, what is retired, what is forbidden, and what is next authorized.
- **Authority:** None for game mechanics — summary of evidence and rulings. Authority Class D (governance summary).
- **Audience:** LLMs auditing EIC, humans planning next work.
- **Belongs inside:** Investigation → Candidate → Falsification → Minimum Simulation → Human Decision Brief → Human Rulings → Provisional Shapes → Discriminating Simulation → Falsification Gate narrative in 1 page, Current stage table (Architecture C RETIRED, PIL-003 still HUMAN-LOCKED, HD-EIC-01 to 08 in force, numeric tuning NOT AUTHORISED, Provenance & Reputation DEFERRED, automatic Architecture D NOT AUTHORISED), List of what is explicitly forbidden without further human ruling, Next authorized operation (human must supply structural boundaries per HD-EIC-07).
- **Must not belong inside:** Detailed M1-M8 mechanisms (those remain archive), functional shape formulas, simulation parameters.
- **Relationship:** Summarizes archive evidence, controlled by Human Rulings doc and Governance doc.
- **Authoritative:** No for mechanics, yes for program status.
- **Update frequency:** When program stage changes.
- **Should LLM read during audit:** Yes — fifth, if audit includes EIC.

### 05_OPEN_QUESTIONS.md

- **Filename:** `05_OPEN_QUESTIONS.md`
- **Purpose:** Single authoritative registry of unresolved material, TBDs, and historical containment boundary.
- **Authority:** Governance — declares what is not decided, not what is decided. No design authority.
- **Audience:** LLMs and humans.
- **Belongs inside:** OQ-001 to OQ-017 registry with status (TBD/PROPOSED/DEFERRED), cross-reference to where each OQ appears, Remaining open design decisions from Master GDD (10 items), Historical non-canonical boundary list (12 items) with note "NOT CANONICAL unless explicitly promoted", Evidence Reconciliation Pass open requirement.
- **Must not belong inside:** Proposed solutions to OQs (those belong in archive proposals or future system specs).
- **Relationship:** Referenced by Master GDD and Program State doc. Updated when OQs are resolved via human ruling.
- **Authoritative:** Yes for "what is unresolved", no for design.
- **Update frequency:** When OQs are opened or closed.
- **Should LLM read during audit:** Yes — sixth, to identify unresolved.

### 06_CHANGELOG.md

- **Filename:** `06_CHANGELOG.md`
- **Purpose:** Machine-readable change history enabling change detection.
- **Authority:** Governance — records changes, does not create design.
- **Audience:** LLMs for change detection, humans for review.
- **Belongs inside:** Version graph: v1.0 Consolidated → v1.1 Baseline (status reclassifications: 5 items LOCKED → DERIVED CONSTRAINT) → v1.1.1 Status Patch (human rulings applied: CRFT-006 demoted to PROPOSED, etc.) → Human Rulings HD-EIC-01 to 08 applied, with date, controlling audit, change scope table, and supersedes/superseded-by fields per document.
- **Must not belong inside:** Game mechanics, rationale beyond change scope.
- **Relationship:** Referenced by Manifest, summarizes provenance notes from Master GDDs.
- **Authoritative:** Yes for change history.
- **Update frequency:** On every documentation change.
- **Should LLM read during audit:** Yes — seventh, for change detection.

---

## 9. Archive / Provenance Structure

### Directory layout proposal (proposal-only, no file moves executed)

```
/active/   (6 docs + manifest, as above)
  00_MANIFEST.md
  01_GOVERNANCE.md
  02_MASTER_GDD_v1.1.1.md
  03_HUMAN_RULINGS.md
  04_EIC_PROGRAM_STATE.md
  05_OPEN_QUESTIONS.md
  06_CHANGELOG.md

/archive/
  /gdd_history/
    Master_GDD_v1.0_Consolidated.md  (Doc 16) — status: SUPERSEDED / PROVENANCE
    Master_GDD_v1.1_Canonical_Baseline.md (Doc 18) — status: SUPERSEDED / PROVENANCE
  /audits/
    Master_GDD_v1.0_Canonical_Audit.md (Doc 15) — status: AUDIT RECORD
    Authority_Provenance_Reconciliation_Matrix.md (Doc 0) — status: GOVERNANCE / PROVENANCE (copy remains in active as 01_GOVERNANCE, original preserved here)
  /eic_investigation/
    Economic_Interdependence_Core_Design_Investigation.md (Doc 1) — status: INVESTIGATION / NON-CANONICAL
    EIC_Candidate_v0.1.md (Doc 2) — status: PROPOSED CANDIDATE / ARCHIVED
    EIC_Candidate_v0.1_Falsification_Pass.md (Doc 3) — status: EVIDENCE / ADVERSARIAL ANALYSIS
  /eic_evidence/
    EIC_Comparative_Simulation_Results_v0.1.1.md (Doc 4)
    EIC_Comparative_Simulation_Execution_Results_v0.1.md (Doc 5)
    EIC_Comparative_Simulation_Specification_v0.1_post_shapes.md (Doc 6) — SUPERSEDES pre-shapes spec
    EIC_Comparative_Simulation_Specification_v0.1_pre_shapes.md (Doc 7) — SUPERSEDED
    EIC_Functional_Shapes_Refinement_v0.2.md (Doc 8)
    EIC_Minimum_Comparative_Simulation_Results_v0.1.md (Doc 12)
    EIC_Provisional_Functional_Shapes_v0.1.md (Doc 13)
    EIC_Targeted_γ′_Adversarial_Retest_Results_v0.1.md (Doc 14)
  /eic_gates/
    EIC_Human_Decision_Brief_v0.1.md (Doc 9) — DECISION BRIEF / NON-CANONICAL
  /human_rulings_history/
    (originals of Docs 10 and 11 preserved, with cumulative merged into active 03_HUMAN_RULINGS.md)
```

### Containment rules for archive

- Every archived file must retain original status disclaimers ("Authority: None", "PROPOSED / Non-Canonical").
- Every archived file must be referenced in Manifest with `status: ARCHIVE / PROVENANCE / EVIDENCE / SUPERSEDED` and `active: false`.
- Archive files must not be in LLM default reading set. Audit protocol must state "Consult archive only when provenance or evidence is required, and cite archive source explicitly".
- Historical non-canonical boundary list must be enforced by Governance doc, not just by archive location.

### Candidates for eventual retirement

- Doc 7 (pre-shapes simulation spec) — superseded by Doc 6, which explicitly supersedes it. Could be marked SUPERSEDED and eventually removed from active search, but per project owner requirement, historical material is preserved rather than deleted. Recommendation: keep in archive with SUPERSEDED tag, not delete.
- Duplicate Human Rulings originals after merge — preserve originals in `/human_rulings_history/` for provenance, but active doc is the merged cumulative record.

---

## 10. LLM Audit Reading Protocol

### Minimal deterministic protocol (proposed)

**Goal:** Two independent LLMs given same active docs and same protocol produce substantially comparable findings.

**Step 1 — Read Manifest first**
- File: `00_MANIFEST.md`
- Action: Identify active set (6 docs), version graph, and authority precedence: Governance > Master GDD status tags > Human Rulings > Program State > Open Questions > Changelog.
- Output: List of documents to read in order.

**Step 2 — Identify current canonical state**
- Files: `01_GOVERNANCE.md` + `02_MASTER_GDD_v1.1.1.md` + `03_HUMAN_RULINGS.md` + `04_EIC_PROGRAM_STATE.md`
- Rule: Current state = all statements in Master GDD with status tags as marked, plus human rulings HD-EIC-01 to 08, plus program state table. No other document may add to current state.
- Cite: Use Rule IDs (VIS-001, HD-EIC-01) and document version.

**Step 3 — Identify authority**
- File: `01_GOVERNANCE.md`
- Action: For each significant statement, map to Authority Class:
  - A = Explicit Human/Project Ruling (e.g., HD-EIC-01 Decision: APPROVE D — Layered)
  - B = Explicit Project Principle/Pillar (e.g., VIS-001 Living persistent player-driven virtual world)
  - C = Derived from Locked Principle (e.g., MFG-004 Anti-Factorio Drift)
  - D = Audit/Governance Reclassification (e.g., PROT-001 GDD over prototype)
  - E = Multi-Source Inference (consistent across sources but not human-approved)
  - F = Historical/External Reference (SWG inspiration)
  - G = Prototype/Evidence (behaviour demonstrated, not authority)
- Rule: Only A and B may be HUMAN-LOCKED. C-G may not be treated as HUMAN-LOCKED per GR-001.

**Step 4 — Distinguish current facts from proposals**
- Use status tags per Governance doc:
  - LOCKED / HUMAN-LOCKED = current fact (requires A or B authority)
  - DERIVED CONSTRAINT = binding consequence, not independent decision
  - PROPOSED = recommended direction, not approved
  - TBD = unresolved
  - PROTOTYPE/EVIDENCE = evidence only
  - DEFERRED = intentionally postponed
  - ASSUMPTION = working assumption
  - HISTORICAL = reference only
- Rule: Never silently promote. If status is missing, treat as PROPOSED and flag uncertainty.

**Step 5 — Identify unresolved questions**
- File: `05_OPEN_QUESTIONS.md`
- Action: List all OQs, TBDs, and DEFERRED items. Cross-check with Master GDD "Remaining open design decisions" and Human Rulings "Still forbidden without further human ruling" tables.
- Output: Explicit list of what remains unresolved.

**Step 6 — Inspect changes**
- File: `06_CHANGELOG.md` + Manifest version graph
- Action: Compare current active set against previous manifest (if available) or against lineage tables. Identify status reclassifications (e.g., CRFT-006 DERIVED CONSTRAINT → PROPOSED), new human rulings, and retired architectures (Architecture C RETIRED per HD-EIC-05).
- Output: What changed since previous state.

**Step 7 — Consult historical/provenance material only when necessary**
- Rule: Archive is not part of default reading set. Consult only if:
  - Provenance of a HUMAN-LOCKED statement is questioned (check archive/gdd_history)
  - Evidence for falsification or validation is needed (check archive/eic_evidence)
  - Audit record explains a reclassification (check archive/audits)
- When consulting archive, cite archive path and note "ARCHIVE / EVIDENCE / SUPERSEDED — not current canon per GR-002".

**Step 8 — Report uncertainty**
- Rule: If authority cannot be established, preserve uncertainty as TBD/PROPOSED/ASSUMPTION per Governance doc. Do not infer authority from frequency. Explicitly state "Information not established by corpus" when missing.
- Example: "No prototype was inspected during consolidation — Evidence Reconciliation Pass remains open per Doc 17 § Remaining open decisions #10".

**Step 9 — Cite evidence**
- Format: `[Document ID] Section / Rule ID — direct quote or paraphrase with status and authority class`.
- Example: `02_MASTER_GDD_v1.1.1.md VIS-003 — Status: LOCKED — Authority: Explicit project direction — "Economy primarily player-driven"`.
- For human rulings: `03_HUMAN_RULINGS.md HD-EIC-05 — Decision: ACCEPT FALSIFICATION — Architecture C retired`.

**Step 10 — Produce standardized audit result**
- Template:
  ```
  Executive Verdict: [PASS/FAIL/MARGINAL] per Test 1/2/3
  Current State: [Summary of active GDD + EIC program state]
  Authority Map: [Table of HUMAN-LOCKED vs DERIVED vs PROPOSED counts]
  Unresolved: [List of OQs and TBDs]
  Changes Since Previous: [List from changelog]
  Evidence Consulted: [Active docs + any archive docs consulted]
  Uncertainty: [Explicit gaps]
  Containment Check: [Historical material did/did not contaminate]
  ```

---

## 11. Recommended Standard Metadata

### Machine/LLM-readable envelope (proposed for every active and archived document)

Every document should begin with a YAML or markdown front-matter block that is machine-parseable without natural-language parsing:

```yaml
---
document_id: TCIndustries_Master_GDD_v1.1.1_Status_Patch
filename: 02_MASTER_GDD_v1.1.1.md
version: 1.1.1
date: 2026-08-25
document_type: Working Canonical Reference
authority_class: A/B/C/D mix (per-rule authority in body)
authority_created_by_this_artifact: None / Explicit human rulings / etc.
status: ACTIVE / ARCHIVE / SUPERSEDED / EVIDENCE
controlling_documents:
  - TCIndustries_Master_GDD_v1.1_Canonical_Baseline.md
  - TCIndustries_Master_GDD_v1.0_Canonical_Audit.md
supersedes:
  - TCIndustries_Master_GDD_v1.1_Canonical_Baseline.md
superseded_by: null (or filename if superseded)
provenance: |
  Source consolidation: v1.0 Consolidated
  Controlling audit: v1.0 Canonical Audit
change_id: CHG-2026-08-25-001
human_approval_state: HUMAN-LOCKED set confirmed where flagged / no new human approval
open_questions_referenced: [OQ-001, OQ-002, ...]
evidence_status: NO_PROTOTYPE_INSPECTED / SIMULATION_EVIDENCE / etc.
---
```

### Fields required for reliable reconstruction

| Field | Purpose | Required for Active? | Required for Archive? |
|---|---|---|---|
| document_id | Stable identifier | Yes | Yes |
| version | Version string | Yes | Yes |
| date | ISO date | Yes | Yes |
| document_type | Governance / Canonical / Human Authority / Investigation / Evidence | Yes | Yes |
| authority_created_by_this_artifact | What authority this doc creates, if any | Yes | Yes |
| status | ACTIVE / ARCHIVE / SUPERSEDED / PROVENANCE / EVIDENCE / AUDIT RECORD | Yes | Yes |
| controlling_documents | List of docs that control interpretation | Yes | Optional |
| supersedes | List of docs this supersedes | Yes if applicable | Yes if applicable |
| superseded_by | Doc that supersedes this | Yes if superseded | Yes if superseded |
| human_approval_state | Whether human approval is claimed and for what scope | Yes | Yes |
| open_questions_referenced | OQ IDs mentioned | Yes if applicable | Yes if applicable |
| evidence_status | Whether evidence is simulation, prototype, etc. | Yes if applicable | Yes |
| change_id | Unique change identifier for changelog | Yes | No |

**Why not unnecessary bureaucracy:** These fields replace the need to parse natural-language lineage sentences ("This document supersedes the earlier...") which is the current failure mode for temporal reconstruction. They enforce GR-001 to GR-003 structurally.

### Manifest fields

Manifest should contain:

- `manifest_version`
- `manifest_date`
- `active_set`: list of active docs with document_id, filename, version, status, authority_class, primary_consumer (LLM/Human)
- `archive_set`: list of archived docs with reason (SUPERSEDED, EVIDENCE, PROVENANCE, AUDIT)
- `reading_order`: ordered list for LLM audit
- `authority_precedence`: Governance > Human Rulings > Master GDD status tags > Program State
- `version_graph`: edges supersedes → superseded_by
- `open_questions_registry`: OQ-001 to OQ-017 with status
- `historical_containment_boundary`: list of 12 items that remain NOT CANONICAL
- `change_history`: list of change_ids with date and scope

---

## 12. Migration / Consolidation Plan

**Plan is proposal-only. No files to be changed without human authorization.**

### Stage 0 — Baseline (current state, no changes)

- Corpus is 19 docs in flat namespace, no manifest, overlapping canonical claims.
- Direct evidence: 3 Master GDD copies, 2 simulation spec copies, 12 EIC docs outnumbering canonical.

### Stage 1 — Establish Governance Root and Manifest (no content changes, only new files)

1. Create `00_MANIFEST.md` with inventory of all 19 docs, marking current active vs archive per this report's analysis, but not moving files.
2. Create `01_GOVERNANCE.md` by extracting Authority Classes A-G, GR-001 to GR-003, status vocabulary, Critical Rule, and Design Documentation Standard from Doc 0 and Docs 17/18 How to Read sections. No new governance rules.
3. Human authorization gate: Owner approves Manifest and Governance doc as accurate extraction.

**Effort:** Low — extraction only. **Risk:** Low — no design changes.

### Stage 2 — Declare Single Current Canonical and Single Human Rulings Record (consolidation by reference, not rewrite)

1. Declare `02_MASTER_GDD_v1.1.1.md` as the single current canonical (Doc 17). Mark Doc 16 and Doc 18 as SUPERSEDED / PROVENANCE in Manifest, but keep files.
2. Merge Docs 10 and 11 into `03_HUMAN_RULINGS.md` — cumulative record HD-EIC-01 to 08 with traceability tables preserved. Preserve originals in `/human_rulings_history/` archive.
3. Create `04_EIC_PROGRAM_STATE.md` as summary (no new design), citing evidence from archive.
4. Create `05_OPEN_QUESTIONS.md` consolidating Remaining open decisions (10 items from Doc 17) + OQ registry (17 items) + historical non-canonical boundary (12 items).
5. Create `06_CHANGELOG.md` from change scope tables in Docs 17/18 and human rulings applied tables.

**Effort:** Medium — synthesis of existing tables, no new mechanics. **Risk:** Medium — requires careful preservation of "Explicitly does not determine" scoping to avoid accidental promotion.

### Stage 3 — Archive Containment (file moves into archive directories, manifest updated)

1. Move Docs 1-9, 12-14 into `/archive/eic_*` directories per structure in Section 9, with status headers preserved.
2. Move Docs 16, 18 into `/archive/gdd_history/`.
3. Move Doc 15 into `/archive/audits/`.
4. Keep Doc 0 original in `/archive/audits/` while active governance is `01_GOVERNANCE.md`.
5. Update Manifest to reflect new paths and `active: false` for archive.

**Effort:** Medium — file operations only. **Risk:** Low if Manifest is authoritative — but requires human authorization before any moves.

### Stage 4 — Metadata Envelope Rollout (add machine-readable front-matter to active docs)

1. Add YAML front-matter to 6 active docs + manifest with fields per Section 11.
2. No content changes to body — only envelope.
3. Validate that envelope does not contradict body (e.g., authority_created_by_this_artifact must be "None" for evidence docs).

**Effort:** Low. **Risk:** Low — but must be validated that envelope parsing does not create new authority.

### Stage 5 — LLM Audit Protocol Validation

1. Run two independent LLM audits using protocol in Section 10 against active set only.
2. Compare findings for Test 2 — Audit Reproducibility. Adjust manifest reading order or governance doc if findings diverge.
3. Document validation results as audit record in `/archive/audits/`.

**Effort:** Low — but requires two LLMs. **Risk:** May reveal need to tighten Governance doc wording.

---

## 13. Risks and Trade-offs

| Decision | Benefit | Risk | Mitigation |
|---|---|---|---|
| 6 active docs + manifest vs one enormous document | Preserves modularity, allows human to update one doc without rewriting all | Still requires LLM to read 6 docs (approx 80k chars) | Manifest reading order ensures deterministic reconstruction; still small enough for context window |
| Merging two Human Rulings docs into one cumulative record | Single source for human authority, prevents duplicate-date confusion | Could obscure that HD-EIC-05 to 08 retired Architecture C after HD-EIC-01 to 04 approved it as leading hypothesis | Preserve originals in archive, include traceability table with chronology, retain "ACCEPT FALSIFICATION" decision wording verbatim |
| Creating EIC Program State summary (new doc) | Replaces need to read 12 EIC docs to know program status, improves Test 1 | Summary could be mistaken as new design authority if not carefully scoped | Explicit "Authority: None — summary only" and "Status: PROGRAM STATE / NOT CANONICAL" in envelope, plus "Explicitly does not determine" pattern |
| Preserving all 19 docs in archive vs deleting superseded | Satisfies project owner requirement to preserve historical provenance, prevents information loss | Archive remains large (461k chars) — future LLM may still ingest all if not respecting manifest | Manifest declares `active: false` and audit protocol says "Consult archive only when necessary" with explicit citation requirement; containment rule GR-002 remains in Governance doc |
| Standardized metadata envelope | Enables machine-readable temporal reconstruction and scope control without natural-language parsing | Adds bureaucracy — envelope could drift from body | Envelope fields are derived from body, not independent — validation step checks consistency; envelope does not create authority |
| Keeping Master GDD version in filename (02_MASTER_GDD_v1.1.1.md) | Preserves provenance — filename itself signals supersession chain | May encourage future proliferation of versioned filenames | Manifest declares single current canonical; future versions replace file, changelog records supersession, old versions moved to archive/gdd_history |

**Trade-off not taken:** One enormous document containing everything. Rejected because it would violate secondary objective to maintain small clearly identifiable active set and would make change detection impossible (diff of 460k chars). Also rejected: dozens of narrowly divided documents (e.g., one doc per OQ) — would increase sprawl and harm cross-LLM comparability.

---

## 14. Human Authorization Gates

Per constraints, all proposed restructuring requires later human authorization. No file operations are to be actioned by LLM.

---

## 15. Target State

After Stages 1-4 complete:

- **Active set:** 7 files (manifest + 6 docs) totaling ~80-100k chars, all with machine-readable envelopes, all explicitly declaring authority and status, with deterministic reading order.
- **Archive set:** 19 original files preserved plus 2 new originals (original human rulings preserved separately), organized into `/archive/gdd_history/`, `/archive/audits/`, `/archive/eic_investigation/`, `/archive/eic_evidence/`, `/archive/eic_gates/`, `/archive/human_rulings_history/`.
- **Governance:** Single governance root with GR-001, GR-002, GR-003, Authority Classes A-G, status vocabulary, Critical Rule, and Design Documentation Standard.
- **Canonical:** Single Master GDD v1.1.1 Status Patch as current working canonical, with supersession chain documented in Manifest and Changelog.
- **Human authority:** Single cumulative Human Rulings doc HD-EIC-01 to 08 with explicit "Explicitly does not determine" and "Still forbidden" scoping.
- **Program state:** Single EIC Program State doc declaring Architecture C RETIRED per HD-EIC-05, PIL-003 still HUMAN-LOCKED, numeric tuning NOT AUTHORISED, Provenance & Reputation DEFERRED, automatic Architecture D NOT AUTHORISED.
- **Open questions:** Single registry of 17 OQs + 10 remaining open decisions + 12 historical non-canonical boundary items.
- **Change detection:** Changelog with change IDs and version graph.
- **Auditability:** Any new LLM can reconstruct state by reading 7 files in order, distinguish human decisions from proposals via status tags and authority classes, identify unresolved questions via OQ registry, detect changes via changelog, and consult archive only when provenance or evidence is required, with explicit citation.

Success tests after target state:

- **Test 1 — State Reconstruction:** PASS — new LLM reads 7 files, not 19, and determines current project state.
- **Test 2 — Audit Reproducibility:** PASS — two LLMs given same manifest and reading protocol produce comparable findings (same HUMAN-LOCKED set, same retired architecture, same unresolved OQs).
- **Test 3 — Historical Containment:** PASS — old proposals, superseded docs, prototype behaviour, LLM recommendations cannot become canon because archive is marked `active: false` and governance rules GR-001 and GR-002 are enforced by manifest and protocol.

---

## 16. Open Questions / Information Not Established by the Corpus

Explicitly stating what corpus does not establish (evidence discipline):

1. **No prototype inspected:** Direct evidence from Docs 16,17,18: "No prototype was inspected during consolidation or status passes" and "Prototype references remain PROTOTYPE/ASSUMPTION". The Evidence Reconciliation Pass (GDD ↔ prototype ↔ decision matrix) remains open. Corpus does not establish whether any existing prototype implements specialisation budget, resource lifecycle, or crafting differentiation.

2. **Specialisation budget model unresolved:** OQ-001 remains TBD across 16 docs. Corpus contains multiple PROPOSED alternatives (formal concentration budget, soft capacity, attention load, knowledge maintenance cost, asset concentration, pure opportunity-cost gradients) but no HUMAN-LOCKED selection. Doc 2 and Doc 10 explicitly state implementation remains Open / PROPOSED.

3. **Resource lifecycle mathematics:** OQ-002 remains TBD. Corpus establishes LOCKED principle RES-002 "Resources exist to create discovery/scarcity/economic events" but not attribute set, lifecycle mathematics, or pool model (time-limited vs extraction-limited remains NOT CANONICAL per historical boundary list).

4. **Crafting experimentation model:** OQ-003 remains TBD. Corpus contains PROPOSED experimentation point allocation / success-critical-failure model but marks it NOT CANONICAL unless explicitly promoted.

5. **Manufacturing and automation constraints:** OQ-004 remains TBD. Corpus establishes LOCKED principle MFG-001 "Automation must not eliminate player relevance" and DERIVED CONSTRAINT MFG-004 "Anti-Factorio Drift", but exact factory quality rules, residual channels for specialist comparative value under industrial parity, and automation limits remain PROPOSED/TBD. Hard factory-quality ceiling explicitly rejected as non-canonical (Authority Matrix §5).

6. **Multi-accounting and exploit policies:** OQ-012 remains TBD. HD-EIC-03 establishes HUMAN-LOCKED validation stance M1 "Legitimate multi-accounting is an in-scope adversarial condition" but explicitly does not authorise account-level restrictions or efficiency parameters.

7. **Numeric balance values:** Explicitly NOT AUTHORISED per HD-EIC-08 and Governance. Corpus contains no numeric values by design (BAL-001). No formulas with numbers are canonical.

8. **Provenance & Reputation Engine:** Explicitly DEFERRED per HD-EIC-08 and design sequence Phase 2. OQ-017 remains unresolved. Corpus does not establish multi-contributor product or brand rules.

9. **City formation, commerce, society, world, combat:** All DEFERRED or TBD per design sequence. No canonical rules beyond high-level LOCKED principles (BLD-001, SOC-001, etc.).

10. **Authority review queue resolution:** Docs 0 and 15 flagged WRLD-001, PLR-001, PLR-003, CRFT-001, SAFE-001 as "Authority review required". Doc 17 confirms them as HUMAN-LOCKED (principle only). Corpus establishes confirmation occurred 2026-08-25, but final human confirmation of entire HUMAN-LOCKED vision/pillar set remains listed as open per Doc 18 § Remaining Human Decisions #10.

11. **Manifest existence:** No manifest currently exists in corpus. This assessment proposes one, but corpus does not establish manifest fields, reading order, or active/archive boundary. All manifest proposals are new governance artifacts requiring human authorization.

12. **Change identifiers:** No change_id exists in current docs. Change scope tables exist but are not machine-readable.

---

## Recommended Architecture Table

| Document | Purpose | Authority | Active/Archive | Primary Consumer |
|---|---|---|---|---|
| 00_MANIFEST.md | Deterministic entry point, version graph, active/archive list, reading order | Governance — no game design (Class D) | Active | LLM auditor |
| 01_GOVERNANCE.md | Authority classes A-G, status vocabulary, GR-001 to GR-003, Critical Rule, design standard | Governance — defines interpretation (Class D) | Active | LLM + Human |
| 02_MASTER_GDD_v1.1.1.md | Current working canonical design reference — vision, pillars, principles, constraints, glossary, non-canonical boundary, open decisions | Contains HUMAN-LOCKED (A/B) and DERIVED CONSTRAINT (C/D) per status tags | Active | Human + LLM |
| 03_HUMAN_RULINGS.md | Cumulative human rulings HD-EIC-01 to 08 with scope limits and traceability | Explicit human authority (Class A) for decisions stated only | Active | LLM + Human |
| 04_EIC_PROGRAM_STATE.md | Current EIC program stage, retired architecture, forbidden operations, next authorized operation | None for mechanics — summary of evidence and rulings (Class D) | Active | LLM + Human |
| 05_OPEN_QUESTIONS.md | Registry of OQ-001 to OQ-017, TBDs, DEFERRED, historical NOT CANONICAL boundary (12 items), evidence reconciliation open | Governance — declares unresolved (Class D) | Active | LLM + Human |
| 06_CHANGELOG.md | Machine-readable change history, version graph, supersession chain | Governance — records changes (Class D) | Active | LLM + Human |
| archive/gdd_history/Master_GDD_v1.0_Consolidated.md | Superseded working reference — provenance only | Superseded — originally claimed primary authority, now provenance | Archive / Provenance | Human (when provenance needed) |
| archive/gdd_history/Master_GDD_v1.1_Canonical_Baseline.md | Superseded canonical baseline — provenance | Superseded — provenance | Archive / Provenance | Human |
| archive/audits/Master_GDD_v1.0_Canonical_Audit.md | Audit that identified over-promoted LOCKED and recommended DERIVED CONSTRAINT | Audit record — no design authority | Archive / Audit Record | LLM (when explaining reclassifications) |
| archive/audits/Authority_Provenance_Reconciliation_Matrix.md | Original governance matrix defining Authority Classes and GR rules | Governance / Provenance — source for 01_GOVERNANCE.md | Archive / Governance Provenance | LLM (for authority provenance) |
| archive/eic_investigation/Economic_Interdependence_Core_Design_Investigation.md | Investigation developing three competing architectures, recommending Architecture C as leading hypothesis | None — investigation (Non-Canonical) | Archive / Evidence | Human / LLM when investigating EIC origins |
| archive/eic_investigation/EIC_Candidate_v0.1.md | Falsifiable structural candidate M1-M8 — leading candidate before falsification | None — PROPOSED CANDIDATE | Archive / Evidence | Human |
| archive/eic_investigation/EIC_Candidate_v0.1_Falsification_Pass.md | Adversarial attack on Architecture C — identifies attack surfaces | None — evidence | Archive / Evidence | Human |
| archive/eic_evidence/EIC_Comparative_Simulation_Results_v0.1.1.md | Qualitative adversarial structural evaluation — packages α and β eliminated, Architecture C underdetermined | None — evidence | Archive / Evidence | Human |
| archive/eic_evidence/EIC_Comparative_Simulation_Execution_Results_v0.1.md | Minimum simulation execution — INCONCLUSIVE / UNDERDETERMINED finding | None — evidence | Archive / Evidence | Human |
| archive/eic_evidence/EIC_Comparative_Simulation_Specification_v0.1_post_shapes.md | Controlled experiment spec after human rulings and provisional shapes — supersedes pre-shapes spec | None — specification | Archive / Evidence (current spec) | Human |
| archive/eic_evidence/EIC_Comparative_Simulation_Specification_v0.1_pre_shapes.md | Earlier spec before functional shapes — superseded | None — SUPERSEDED spec | Archive / Evidence / Superseded | Human (provenance) |
| archive/eic_evidence/EIC_Functional_Shapes_Refinement_v0.2.md | Targeted γ closure refinement — narrow provisional refinement | None — provisional | Archive / Evidence | Human |
| archive/eic_evidence/EIC_Minimum_Comparative_Simulation_Results_v0.1.md | Minimum simulation — under-specification finding | None — evidence | Archive / Evidence | Human |
| archive/eic_evidence/EIC_Provisional_Functional_Shapes_v0.1.md | Provisional non-numeric functional shapes for M1 and M4 | None — simulation assumptions only | Archive / Evidence | Human |
| archive/eic_evidence/EIC_Targeted_γ′_Adversarial_Retest_Results_v0.1.md | Adversarial re-test — evidence for falsification | None — evidence | Archive / Evidence | Human |
| archive/eic_gates/EIC_Human_Decision_Brief_v0.1.md | Minimal human decision set — 4 gates required before discriminating simulation | None — decision brief | Archive / Gate Provenance | Human |

---

## Human Authorization Required

**This assessment creates no new design and performs no file operations. All structural changes below require explicit project owner approval before implementation.**

1. **Creation of 00_MANIFEST.md** — new governance artifact declaring active vs archive, reading order, version graph, and authority precedence. Requires approval because manifest becomes controlling reference for all future LLMs.

2. **Creation of 01_GOVERNANCE.md** — extraction and consolidation of Authority Classes A-G, status vocabulary, GR-001 to GR-003, Critical Rule, and Design Documentation Standard from Docs 0, 17, 18. Requires approval because it becomes the single governance root and supersedes scattered How to Read sections.

3. **Declaration of 02_MASTER_GDD_v1.1.1.md as single current canonical** — marking Docs 16 (v1.0 Consolidated) and 18 (v1.1 Baseline) as SUPERSEDED / PROVENANCE in manifest. Requires approval because it retires two documents that currently claim to be current working canonical reference. Per Doc 17 lineage, this is the intended state, but explicit human confirmation is required per Remaining Human Decisions #10.

4. **Merging Human Rulings Docs 10 and 11 into 03_HUMAN_RULINGS.md** — cumulative record HD-EIC-01 to 08. Requires approval because it merges two human authority records with same date but sequential decisions (HD-EIC-05 ACCEPT FALSIFICATION retires Architecture C). Must preserve verbatim decision wording and "Explicitly does not determine" scoping.

5. **Creation of 04_EIC_PROGRAM_STATE.md** — new summary doc with no design authority, summarizing program state table and forbidden operations from Doc 11 §3. Requires approval because it interprets evidence and rulings into a program status statement.

6. **Creation of 05_OPEN_QUESTIONS.md** — consolidating OQ registry (17 items), remaining open decisions (10 items), and historical non-canonical boundary (12 items). Requires approval because it becomes authoritative for "what is unresolved" and must not accidentally promote proposals.

7. **Creation of 06_CHANGELOG.md** — machine-readable change history from change scope tables in Docs 17/18. Requires approval because it becomes authoritative for change detection.

8. **Archive containment moves** — moving 12 EIC docs, 2 superseded GDDs, 1 audit, and 1 governance matrix original into `/archive/` subdirectories. Requires approval because it changes file locations and active set definition, even though content is preserved per owner requirement to preserve historical material.

9. **Adding machine-readable metadata envelope (YAML front-matter)** to all active documents with fields: document_id, version, date, document_type, authority_created_by_this_artifact, status, controlling_documents, supersedes, superseded_by, human_approval_state, open_questions_referenced, evidence_status, change_id. Requires approval because envelope becomes machine-readable authority and must not contradict body or create new authority.

10. **Adopting LLM Audit Reading Protocol (Section 10)** as standard operating procedure for future audits. Requires approval because it defines how independent LLMs must reconstruct state and could affect comparability of future audits.

11. **Retiring Doc 7 (pre-shapes simulation spec) as SUPERSEDED** — Doc 6 explicitly supersedes it, but retirement marking requires human approval per change control.

No game-design decisions (specialisation budget model, resource lifecycle, crafting experimentation, manufacturing constraints, multi-accounting policies, currency, city formation, combat scope) are authorized by this assessment. Those remain TBD/PROPOSED and require separate explicit human rulings per HD-EIC-07 and HD-EIC-08.

---

**End of Assessment — Proposal Only, No Files Modified**
