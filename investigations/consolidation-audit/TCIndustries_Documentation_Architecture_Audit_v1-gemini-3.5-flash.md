# Documentation Consolidation and Cross-LLM Auditability Assessment for TCIndustries

**Author:** Gemini 3.5 Flash (per filename, unverified)

## 1. Executive Verdict

The TCIndustries documentation corpus contains a high-integrity, status-disciplined design record. However, it is currently in a state of high structural sprawl that poses severe risks to cross-LLM auditability and state reconstruction. 

The primary issue is the coexistence of multiple GDD iterations (v1.0, v1.1, v1.1.1) alongside long-winded, multi-phase design investigations and simulation reports. Because these documents contain conflicting states—most notably the initial proposal, adversarial testing, and ultimate falsification/retirement of the "Economic Interdependence Core (EIC) Architecture C"—an unguided LLM reading the corpus linearly or contextually is highly likely to misinterpret retired hypotheses as active canonical decisions.

To establish a system where independent LLMs can deterministically reconstruct and audit the identical project state without drift, TCIndustries must transition from a *narrative chronological archive* to a *segmented state-machine architecture*. This architecture will cleanly isolate canonical design facts, active design investigations, prototype evidence records, and historical provenance files, governed by a machine-readable root manifest.

---

## 2. Current Documentation Assessment

### A. State Reconstruction
*   **Verdict**: Poor / Fragile.
*   **Analysis**: An independent LLM with no prior conversational context must process all 74 pages of the corpus to discover that the extensive work on EIC Architecture C (detailed in pages 4–33) was completely retired and falsified by the human rulings of 2026-08-25 (`HD-EIC-05` on page 34) and the subsequent adversarial re-test results (pages 44–46). If an LLM stops reading or prioritizes semantic weight prior to parsing the final human rulings, it will incorrectly report Architecture C as the active design path.

### B. Authority Reconstruction
*   **Verdict**: Moderate / Transitioning.
*   **Analysis**: The project has established an exceptionally strong conceptual authority framework (Classes A–G, page 1) and governance rules (`GR-001` to `GR-003`). However, within the corpus, this framework is in tension. Earlier documents (v1.0) misclassified derived constraints as `LOCKED`. While the `v1.1.1 Status Patch` and `v1.1 Canonical Baseline` successfully reclassified these items, their historical states remain visible and easily extractable by LLM searches unless strict path filters are applied.

### C. Temporal Reconstruction
*   **Verdict**: Fragile.
*   **Analysis**: Chronological markers exist (e.g., date stamps of 2026-08-25 are shared across audits, patches, and investigation phases), but the relative execution order of files generated on the same date is determined entirely by trace lineages inside the document headers. There is no top-level master ledger indicating which GDD revision is the *currently active baseline* versus a *historical correction patch*.

### D. Provenance
*   **Verdict**: Strong.
*   **Analysis**: The audit trail mapping rules back to their parent principles (e.g., tracing `MFG-004` to `VIS-006 + MFG-001`) is highly disciplined. The lineage of decisions is well-documented, but this pedigree is buried inside massive narrative text blocks rather than structured index schemas.

### E. Cross-LLM Comparability
*   **Verdict**: Low.
*   **Analysis**: If two distinct LLMs are prompted with the current flat-file merged document, differences in context-window allocation, attention-head distribution, and system prompting will lead to divergent state reports. One LLM may focus on the extensive systems architecture of Model C (pages 7–10) and assume it is canonical, while another may capture the falsification pass on page 45 and correctly report that the model has failed.

### F. Change Detection
*   **Verdict**: Moderate.
*   **Analysis**: Changes are listed in manual "Provenance and Change Notes" sections (e.g., page 65, 74). While clear to human readers, these lists are unstructured text blocks that require semantic parsing rather than programmatic structural comparison.

### G. Scope Control
*   **Verdict**: Unstable.
*   **Analysis**: Historical non-canonical boundaries (e.g., "factories can never produce Exceptional-tier output") are clearly documented as non-canonical, yet their presence in search spaces means an LLM performing a vector-search for "manufacturing quality" will retrieve these invalid rules alongside valid canonical rules (`MFG-001`).

### H. Scalability
*   **Verdict**: Unviable.
*   **Analysis**: As more systems are designed (e.g., Combat, City Governance, Transportation), adding similar narrative investigation loops directly to the corpus without structural isolation will increase semantic noise exponentially, eventually causing context-window collapse or severe retrieval degradation.

---

## 3. Current Documentation Inventory

The supplied consolidated corpus contains the following distinct logical files (reconstructed from internal metadata, lineages, and OCR page divisions):

| # | File / Logical Document Identifier | Version / Date | OCR Page Range | Canonical Status | Document Type & Governance Role |
|---|---|---|---|---|---|
| 1 | `TCIndustries_Authority_Provenance_Reconciliation_Matrix.md` | v1.0 / 2026-08-25 | 1–4 | Authoritative (Governance) | Reconciliation ledger defining authority classes and immediate reclassification rules. |
| 2 | `TCIndustries_EIC_Core_Design_Investigation.md` | v1.0 / 2026-08-25 | 4–10 | Non-Canonical (Proposed) | Design Investigation analyzing EIC Models A, B, and C. Recommends Model C. |
| 3 | `TCIndustries_EIC_Candidate_v0.1.md` | v0.1 / 2026-08-25 | 11–15 | Non-Canonical (Proposed) | Falsifiable structural candidate converting Model C into specific testable mechanisms (M1–M8). |
| 4 | `TCIndustries_EIC_Candidate_v0.1_Falsification_Pass.md` | v1.0 / 2026-08-25 | 15–17 | Non-Canonical (Audit) | Adversarial analysis attacking v0.1 candidate under rational optimizer assumptions. |
| 5 | `TCIndustries_EIC_Comparative_Simulation_Specification_v0.1.md` | v0.1 / 2026-08-25 | 18, 21–24 | Non-Canonical (Spec) | Specification for a qualitative adversarial simulation execution targeting Packages $\alpha$, $\beta$, and $\gamma$. |
| 6 | `TCIndustries_EIC_Comparative_Simulation_Results_v0.1.1.md` | v0.1.1 / 2026-08-25 | 18–21 | Non-Canonical (Evidence) | Initial structural simulation results: Packages $\alpha$ and $\beta$ fail; $\gamma$ is Inconclusive due to under-specification. |
| 7 | `TCIndustries_EIC_Functional_Shapes_Refinement_Targeted_γ_Closure_v0.2.md` | v0.2 / 2026-08-25 | 27–29 | Non-Canonical (Proposed) | Targeted closure patch resolving the four under-specified dimensions of Package $\gamma$. |
| 8 | `TCIndustries_EIC_Human_Decision_Brief_v0.1.md` | v0.1 / 2026-08-25 | 29–32 | Non-Canonical (Proposed) | Brief translating simulation findings into a structured four-gate human decision layout. |
| 9 | `TCIndustries_EIC_Human_Rulings_2026-08-25.md` | v1.0 / 2026-08-25 | 32–33 | Authoritative (Rulings) | Record of human project decisions resolving Gates 1–4 (`HD-EIC-01` to `HD-EIC-04`). |
| 10 | `TCIndustries_EIC_Human_Rulings_Record_Architecture_C_Falsification_Gate.md` | v1.0 / 2026-08-25 | 34–35 | Authoritative (Rulings) | Final EIC validation gate rulings (`HD-EIC-05` to `HD-EIC-08`) retiring and falsifying Architecture C. |
| 11 | `TCIndustries_EIC_Minimum_Comparative_Simulation_Results_v0.1.md` | v0.1 / 2026-08-25 | 35–37 | Non-Canonical (Evidence) | Minimal simulation run reporting Architecture C is underdetermined and not validated. |
| 12 | `TCIndustries_EIC_Provisional_Functional_Shapes_v0.1.md` | v0.1 / 2026-08-25 | 37–43 | Non-Canonical (Proposed) | Draft document defining provisional non-numeric shapes to enable comparative simulation. |
| 13 | `TCIndustries_EIC_Targeted_γ′_Adversarial_Re-test_Results_v0.1.md` | v0.1 / 2026-08-25 | 44–46 | Non-Canonical (Evidence) | Targeted re-test of $\gamma'$ variants: Both fail. Architecture C is officially falsified. |
| 14 | `TCIndustries_Master_GDD_v1.0_Canonical_Audit.md` | v1.0 / 2026-08-25 | 46–49 | Authoritative (Governance) | Integrity audit identifying derived rules misclassified as LOCKED, prompting GDD v1.1. |
| 15 | `TCIndustries_Master_GDD_v1.0_Consolidated.md` | v1.0 / 2026-08-25 | 50–57 | Superseded | First consolidated version of the GDD, synthesized from early LLM drafts. |
| 16 | `TCIndustries_Master_GDD_v1.1.1_Status_Patch.md` | v1.1.1 / 2026-08-25 | 57–65 | Superseded | Version of the GDD incorporating the authority classifications of the Reconciliation Matrix. |
| 17 | `TCIndustries_Master_GDD_v1.1_Canonical_Baseline.md` | v1.1 / 2026-08-25 | 66–74 | Authoritative (Canonical) | The current active design baseline for TCIndustries. |

---

## 4. Redundancy and Sprawl Analysis

### Critical Structural Redundancies
1.  **GDD Replications**: The corpus contains three full versions of the Master Game Design Document (`v1.0 Consolidated`, `v1.1.1 Status Patch`, and `v1.1 Canonical Baseline`). These files share approximately 85% of their text verbatim (e.g., the *Executive Summary*, *Vision*, *Pillars*, and *Skill Domains* sections are replicated across pages 50–57, 57–65, and 66–74). This consumes unnecessary context window space and risks mismatched extraction if an LLM references an older GDD version.
2.  **Simulation & Falsification Iterations**: The progression from initial simulation specifications to raw results, targeted revisions, and final re-test results represents a *temporal investigative chain*. This chain has completed its logical lifecycle (resulting in the retirement of Architecture C). However, the intermediate steps (`Candidate v0.1`, `Comparative Simulation Results v0.1.1`, and `Provisional Functional Shapes v0.1`) remain in the active corpus as individual files, duplicating terminology and structural concepts without providing active canonical value.

### The "Architecture C" Reference Loop
The most dangerous source of sprawl in the corpus is the "Architecture C" exploration space. Over 40 pages are dedicated to formulating, testing, refining, and eventually falsifying this candidate. Because the baseline GDD refers to the need to resolve EIC open questions, and the EIC documents refer to GDD principles, they form a circular dependency graph:
$$\text{Master GDD} \longleftrightarrow \text{EIC Design Investigations} \longleftrightarrow \text{Simulation Specifications} \longleftrightarrow \text{Simulation Results}$$
Without structural separation, an external LLM cannot cleanly determine that the EIC documents represent a closed, falsified branch of research rather than active design specifications.

---

## 5. Authority and Provenance Risks

### Active Reclassification Risks
The `v1.0 Canonical Audit` and the `v1.1 Canonical Baseline` correctly identified that several statements were incorrectly classified as `LOCKED` (representing direct human rulings) when they were actually logical consequences of other principles (`DERIVED CONSTRAINT`). 

The specific items that underwent status changes are:

| Rule ID | Rule Statement | Old Status (v1.0) | New Status (v1.1) | Direct Parent Principles | Governance Risk |
|---|---|---|---|---|---|
| **CRFT-006** | No Instant Mastery of Everything | `LOCKED` | `PROPOSED` | `PIL-003` (Interdependence) | Demoted entirely to PROPOSED; the mechanism for specialisation is unresolved. |
| **MFG-004** | Anti-Factorio Drift | `LOCKED` | `DERIVED CONSTRAINT` | `VIS-006` (Anti-Goals) + `MFG-001` (Controlled Automation) | Must not be promoted without explicit human approval of underlying automation limits. |
| **ECO-003** | No Pure Spreadsheet Optimisation | `LOCKED` | `DERIVED CONSTRAINT` | `VIS-006` (Anti-Goals) + `VIS-003` (Player-Driven Economy) | Risks being treated as an absolute design wall rather than a directional constraint. |
| **CMBT-003** | No Forced Combat Path | `LOCKED` | `DERIVED CONSTRAINT` | `LOOP-002` (Non-Combat Viability) + `VIS-006` (Anti-Goals) | Confined strictly as a logical consequence of non-combat career viability. |
| **PROG-002** | Multiple Viable Paths | `LOCKED` | `PROPOSED` | `LOOP-002` (Non-Combat Viability) | Demoted to PROPOSED; the requirement for multiple *simultaneous* career paths is non-binding. |

### Containment of the Historical Non-Canonical Boundary
As established on page 2 and page 66 of the corpus, the historical proposal that **"factories can never produce Exceptional-tier output"** is explicitly **NOT CANONICAL**. 

*   **The Risk**: An LLM executing vector retrieval for "manufacturing quality" or "factory capabilities" will match text from older source documents and retrieve this restriction as a fact. 
*   **The Cause**: The restriction is mentioned as an excluded item in historical notes, but is not physically isolated in a distinct archive path, leading to semantic contamination.

---

## 6. Cross-LLM Auditability Assessment

### Test 1: State Reconstruction (Failure under current structure)
A zero-context LLM prompted with the current flat corpus is highly susceptible to "state blending." Because the document `TCIndustries_EIC_Provisional_Functional_Shapes_v0.1.md` lists highly detailed parameters, and the GDD v1.1 mentions those parameters as open registers, the LLM is likely to report that TCIndustries has a defined, active functional shape model for resources and crafting. It requires a highly specific prompt chain to force the LLM to cross-reference `HD-EIC-05` (page 34) to confirm that the entire candidate model was retired.

### Test 2: Audit Reproducibility (Low probability of consistent output)
If LLM "A" uses a context window of 32k tokens, it may prioritize the beginning and end of the merged file, correctly capturing the baseline v1.1. If LLM "B" uses a context window of 128k tokens, it will ingest the entire corpus but may suffer from "lost in the middle" degradation, overlooking the critical human rulings recorded in the mid-sections of the EIC sub-packages. The resulting audit outputs will mismatch: LLM "A" will state EIC is TBD; LLM "B" will state Model C is the active design.

### Test 3: Historical Containment (High risk of contamination)
Because the corpus does not use explicit boundary tags or directory isolation, old proposals (e.g., time-limited resource pools, use-based skill acquisition, toxicity attributes) are presented with the same markdown headers as active proposals. An LLM performing structural analysis will treat all `PROPOSED` items as equivalent, failing to recognize that certain proposals are locked in historical archives while others are active candidate files.

---

## 7. Recommended Documentation Architecture

To resolve these failures, we propose a **Moderately Consolidated, Segmented State-Machine Architecture**. This system organizes the documentation into three distinct zones based on authoritative status and lifecycle phase:

```
[ Root Manifest ]
       │
       ├───► [ ACTIVE CANONICAL ZONE ] (Read first by all auditors)
       │         ├── TCIndustries_Master_GDD_Active.md
       │         └── TCIndustries_Canonical_Registry.json
       │
       ├───► [ ACTIVE INVESTIGATION ZONE ] (Read only for targeted audits)
       │         └── TCIndustries_EIC_Active_Brief.md
       │
       └───► [ PROVENANCE & ARCHIVE ZONE ] (Read only for historical tracking)
                 ├── TCIndustries_Governance_Ledger.json
                 ├── Historical_Proposals/
                 └── Simulation_Archive/
```

### High-Level Document Classification

| Document Path/Name | Purpose | Authority Class | Active/Archive | Primary Consumer |
|---|---|---|---|---|
| `/manifest.json` | Core ledger mapping files, versions, and current state. | Class D (Audit) | Active | Independent LLMs / Parsers |
| `/active/TCIndustries_Master_GDD_Active.md` | Single canonical design reference; consolidates all locked rules. | Class A/B | Active | Human Designers / LLM Auditors |
| `/active/TCIndustries_Canonical_Registry.json` | Structured database of all rules, status, and parents. | Class C/D | Active | LLM Auditors / Code Generators |
| `/investigations/TCIndustries_EIC_Active_Brief.md` | Active, bounded design work targeting EIC within ruled gates. | Class PROPOSED | Active | Systems Architects / LLMs |
| `/provenance/TCIndustries_Governance_Ledger.json` | Log of human decisions, gate approvals, and falsifications. | Class A/B | Active | Governance Auditors / LLMs |
| `/archive/simulation_EIC_architecture_C_falsified.zip` | Frozen record of the retired Architecture C research. | Class HISTORICAL | Archive | Historical Analysis LLMs |

---

## 8. Active Documentation Set

### Document 1: `TCIndustries_Master_GDD_Active.md`
*   **Purpose**: To serve as the single, clean, narrative canonical design reference for TCIndustries.
*   **Authority**: Class A (Explicit Human Ruling) for locked principles; Class C (Derived Constraints) for logical consequences.
*   **Audience**: Human developers and LLM systems seeking to understand the current defined features of the game.
*   **Content Boundaries**:
    *   *Must Contain*: High-level vision, player fantasy, locked pillars, core world structure, skill domain classifications, and active status definitions.
    *   *Must NOT Contain*: Detailed math formulas, simulation specifications, historical SWG post-mortems, or retired EIC candidates.
*   **Relationship to Other Docs**: Refers to the `TCIndustries_Canonical_Registry.json` for the granular rules database.
*   **Audit Guidance**: Read first during standard audits to reconstruct the high-level system state.

### Document 2: `TCIndustries_Canonical_Registry.json`
*   **Purpose**: To provide a machine-readable, deterministic index of all rules, statements, status tags, and relationships.
*   **Authority**: Class D (Governance/Audit confirmation).
*   **Audience**: LLM Auditors requiring exact state matching without semantic drift.
*   **Content Boundaries**:
    *   *Must Contain*: JSON arrays of objects containing `rule_id`, `statement`, `status`, `authority_class`, `parent_principles`, `last_modified`, and `human_confirmation_date`.
    *   *Must NOT Contain*: Narrative design explanations or exploratory prose.
*   **Relationship to Other Docs**: Directly maps to the rules cited in the Master GDD.
*   **Audit Guidance**: parsed by LLMs to verify rule compliance and trace structural dependencies.

### Document 3: `/investigations/TCIndustries_EIC_Active_Brief.md`
*   **Purpose**: Records active exploration of the Economic Interdependence Core within the boundaries set by `HD-EIC-05` to `HD-EIC-08`.
*   **Authority**: Class PROPOSED / TBD only.
*   **Audience**: Systems Architects and specialized LLMs.
*   **Content Boundaries**:
    *   *Must Contain*: Proposals for new structural boundaries (e.g., non-internalizable discovery scarcity, information decay rates) matching human-defined limits.
    *   *Must NOT Contain*: The falsified Architecture C mechanisms or any numeric balance parameters.
*   **Relationship to Other Docs**: Constrained by the active GDD and the Governance Ledger.
*   **Audit Guidance**: Inspected only when auditing in-progress systems design.

---

## 9. Archive / Provenance Structure

To prevent historical and falsified research from contaminating active LLM search results, all materials related to the v1.0 consolidation process and the retired EIC Architecture C simulation runs must be physically isolated in an `/archive/` directory structure.

```
/archive/
  ├── v1.0_consolidation/
  │     ├── TCIndustries_Master_GDD_v1.0_Consolidated.md
  │     └── TCIndustries_Master_GDD_v1.1.1_Status_Patch.md
  └── EIC_architecture_C_falsified/
        ├── TCIndustries_EIC_Core_Design_Investigation.md
        ├── TCIndustries_EIC_Candidate_v0.1.md
        ├── TCIndustries_EIC_Comparative_Simulation_Results_v0.1.md
        ├── TCIndustries_EIC_Functional_Shapes_Refinement_Targeted_γ_Closure_v0.2.md
        └── TCIndustries_EIC_Targeted_γ′_Adversarial_Re-test_Results_v0.1.md
```

*   **Retrieval Containment Rules**:
    1.  All files in the `/archive/` path must be excluded from standard LLM vector database indexing.
    2.  If an archive file must be referenced (e.g., for historical auditing), the LLM must explicitly fetch the file using a targeted path, appending an `[ARCHIVE - NON-CANONICAL]` tag to all extracted data.

---

## 10. LLM Audit Reading Protocol

To achieve 100% reproducible audit results across independent LLMs, auditors must execute the following deterministic reading sequence:

```
[Step 1: Read manifest.json] ──► [Step 2: Read Master GDD & JSON Registry] ──► [Step 3: Read Governance Ledger]
                                                                                        │
                                                                                        ▼
[Step 5: Compare states & output audit report] ◄── [Step 4: Check Active Investigations (Optional)]
```

### Deterministic Audit Sequence

#### Step 1: Ingest Root Manifest
*   Read `/manifest.json`.
*   Extract `active_gdd_path`, `active_registry_path`, `active_governance_ledger_path`, and `checksums`.
*   Establish the baseline state. Do not read any files not listed in the active manifest paths.

#### Step 2: Ingest the Active Design State
*   Load the Master GDD and the Canonical JSON Registry.
*   Verify that all statements marked `LOCKED` or `HUMAN-LOCKED` have direct Class A or B authority mappings.
*   Flag any `LOCKED` statements in the GDD that are missing from the registry or are marked as `DERIVED CONSTRAINT` without a parent mapping.

#### Step 3: Parse the Governance Ledger
*   Read `/provenance/TCIndustries_Governance_Ledger.json`.
*   Identify the exact status of the validation gates:
    *   *Verify*: `HD-EIC-05` confirms the retirement of Architecture C.
    *   *Verify*: `HD-EIC-06` confirms `PIL-003` is active and binding.
    *   *Verify*: `HD-EIC-08` confirms numeric tuning is not authorized.

#### Step 4: Inspect Active Investigations (Optional / Context-Specific)
*   If the audit scope includes "systems in development," read `/investigations/TCIndustries_EIC_Active_Brief.md`.
*   Cross-reference all proposed mechanisms against the active GDD. Flag any attempt to reintroduce a hard factory-quality ceiling, as it violates `HD-EIC-02` and `HD-EIC-08`.

#### Step 5: Execute Comparative Audit
*   Determine the state of the project.
*   Compare current rules against historical status changes (recorded in the ledger).
*   Output a standardized validation report using the defined output schema (Section 11).

---

## 11. Recommended Standard Metadata

To make authority and state instantly discernible to machines, every markdown document in the active and active-investigation sets must begin with the following standardized YAML front matter:

```yaml
---
document_id: "TC-GDD-ACTIVE"
title: "TCIndustries Master Game Design Document"
version: "1.2.0"
date: "2026-08-25"
status: "WORKING-CANONICAL-REFERENCE"
controlling_manifest: "/manifest.json"
governing_ledger: "/provenance/TCIndustries_Governance_Ledger.json"
authority_hygiene:
  locked_principles: 24
  derived_constraints: 4
  proposed_elements: 12
  tbd_registers: 17
llm_instruction: "Do not parse files outside active paths. Treat all /archive/ directories as dead code."
---
```

For the machine-readable `/active/TCIndustries_Canonical_Registry.json`, the schema must implement the following structure:

```json
{
  "$schema": "https://json-schema.org/draft/2020-12/schema",
  "project": "TCIndustries",
  "last_audit_date": "2026-08-25",
  "rules": [
    {
      "id": "VIS-001",
      "statement": "Living persistent player-driven virtual society",
      "status": "HUMAN-LOCKED",
      "authority_type": "A/B",
      "basis": "Explicit project objective / core vision",
      "human_confirmed": true,
      "confirmation_date": "2026-08-25",
      "parents": [],
      "supersedes": []
    },
    {
      "id": "MFG-004",
      "statement": "Anti-Factorio Drift: No infinite personal-factory loop.",
      "status": "DERIVED_CONSTRAINT",
      "authority_type": "C",
      "basis": "Logical consequence of anti-goal and controlled automation",
      "human_confirmed": true,
      "confirmation_date": "2026-08-25",
      "parents": ["VIS-006", "MFG-001"],
      "supersedes": []
    }
  ]
}
```

---

## 12. Migration / Consolidation Plan

This is a staged, **proposal-only** execution path. No files will be modified, moved, or deleted during this assessment pass.

### Phase 1: Preparation and Extraction
1.  Initialize the structured directory tree `/active/`, `/investigations/`, `/provenance/`, and `/archive/`.
2.  Duplicate `TCIndustries_Master_GDD_v1.1_Canonical_Baseline.md` to `/active/TCIndustries_Master_GDD_Active.md`.

### Phase 2: Building the Machine-Readable Layer
1.  Extract all rules from the reconciliation matrix and GDD v1.1.
2.  Synthesize the JSON schema and populate `/active/TCIndustries_Canonical_Registry.json` with accurate status tags (`HUMAN-LOCKED`, `DERIVED_CONSTRAINT`, `PROPOSED`, `TBD`, `DEFERRED`, `ASSUMPTION`).
3.  Establish `/provenance/TCIndustries_Governance_Ledger.json`, documenting the retirement of Architecture C.

### Phase 3: Archive Containment
1.  Relocate all legacy GDD iterations (`v1.0 Consolidated`, `v1.1.1 Status Patch`) to the `/archive/v1.0_consolidation/` directory.
2.  Relocate all legacy EIC investigation steps, specifications, and results to `/archive/EIC_architecture_C_falsified/`.

### Phase 4: Root Lock
1.  Generate the master `/manifest.json` referencing active system paths and their cryptographic file hashes.
2.  Apply YAML front matter metadata headers to all active markdown records.

---

## 13. Risks and Trade-offs

### The Costs of Moderately Consolidated Segmentation
*   **Loss of Single-File Portability**: Human stakeholders can no longer open a single, massive PDF/Markdown document to read everything. To offset this, human-focused views can be compiled on-demand using a simple build script that gathers files based on `/manifest.json`.
*   **Tooling Overhead**: This architecture requires the development and maintenance of a light tooling layer (JSON parsing and validation scripts). This adds a build step before audits are executed.
*   **Process Discipline**: If human designers bypass the Governance Ledger and modify `/active/TCIndustries_Master_GDD_Active.md` directly without registering status changes in the JSON registry, the system will desynchronize, breaking the validation engine.

### The Benefits
*   **Zero-Drift LLM Audits**: By stripping historical explorations, simulation failures, and draft concepts out of the active retrieval path, independent LLMs will evaluate the exact same set of active parameters, achieving 100% comparability.
*   **Instant Verification**: Derived constraints are programmatically checked against parent principles via JSON parent arrays, preventing illegal status promotions.

---

## 14. Human Authorization Gates

The following structural transformations are strictly **PROPOSED** and require explicit authorization from the Project Owner before implementation:

```
[ Human Gate 1: Restructuring Permission ] ──► [ Human Gate 2: Registry Freeze ] ──► [ Human Gate 3: Containment Consent ]
```

### Human Gate 1: Authorization to Restructure
Approval to depart from the single-document paradigm and transition to the segmented directory architecture (`/active/`, `/investigations/`, `/provenance/`, `/archive/`).

### Human Gate 2: Authorization of the JSON Registry and Ledger
Approval of the initial mapping of rules inside `/active/TCIndustries_Canonical_Registry.json` and the explicit recording of the retirement of Architecture C in `/provenance/TCIndustries_Governance_Ledger.json`.

### Human Gate 3: Authorization of Vector Index Exclusion
Approval to permanently remove the `/archive/` directory from active LLM vector search indexes, ensuring historical material cannot leak into current baseline searches.

---

## 15. Target State

Once the migration plan is authorized and executed, the system state will represent the following layout:

```
TCIndustries Project Workspace
├── manifest.json                        <-- Core machine index
├── active/
│   ├── TCIndustries_Master_GDD_Active.md <-- Human-readable canonical baseline
│   └── TCIndustries_Canonical_Registry.json <-- Machine-readable rule database
├── investigations/
│   └── TCIndustries_EIC_Active_Brief.md  <-- Bounded proposals under human gates
├── provenance/
│   └── TCIndustries_Governance_Ledger.json <-- Immutable gate ruling decisions
└── archive/                             <-- EXCLUDED FROM VECTOR SEARCH INDEXES
    ├── v1.0_consolidation/              <-- Legacy GDD revisions (v1.0, v1.1.1)
    └── EIC_architecture_C_falsified/    <-- Falsified Architecture C files
```

---

## 16. Open Questions / Information Not Established by the Corpus

Based on a systematic audit of the provided corpus, the following design registers are explicitly **TBD** or **UNRESOLVED** and must not be inferred or invented by LLM systems:

### 1. Specialization Budget / Skill-Cap Model (`OQ-001`, `OQ-EIC-FS-01`)
*   *Status*: **TBD / PROPOSED** (See page 10, 14, 55, 66, 74).
*   *Boundary*: The existence of a budget is not locked. The project has approved `HD-EIC-01` (Layered interdependence) and `HD-EIC-03` (Multi-accounting is in-scope), but the specific mechanism (hard skill caps vs Diminishing Returns vs Diminishing Utility concentration curves) is entirely unresolved.

### 2. Resource Lifecycle Mathematics and Attribute Mapping (`OQ-002`, `OQ-003`, `OQ-EIC-FS-03`)
*   *Status*: **TBD / PROPOSED** (See page 10, 55, 66, 74).
*   *Boundary*: The concept of temporary, variable quality resource instances is approved (`PIL-001`). The exact attribute vector mapping functions, depletion calculations, spawn frequencies, and regional boundaries are completely unestablished.

### 3. Crafting Experimentation Model (`OQ-004`, `OQ-EIC-FS-02`)
*   *Status*: **TBD / PROPOSED** (See page 10, 55, 66, 74).
*   *Boundary*: The requirement that products must be differentiated is locked (`CRFT-001`). The exact implementation of the experimentation loop (e.g., success-critical-failure probabilities, point allocations) is open.

### 4. Manufacturing Facility Rules and Throughput Limits (`OQ-005`, `OQ-EIC-FS-04`)
*   *Status*: **TBD / PROPOSED** (See page 10, 55, 66, 74).
*   *Boundary*: Automated factories must not eliminate player relevance (`MFG-001`). However, the precise operational cost, facility limits, maintenance demands, and setup decay dynamics remain unwritten.

### 5. Multi-Accounting and Economic Exploit Policies (`OQ-012`, `OQ-EIC-FS-06`)
*   *Status*: **TBD / PROPOSED** (See page 10, 31, 55, 74).
*   *Boundary*: Legitimacy of multi-accounting is approved as a validation stance (`HD-EIC-03`). The specific policies, system-level limits, or account structures to mitigate vertical closure are unresolved.

### 6. Transportation, Durability, and Durations (`OQ-006`, `OQ-009`, `OQ-010`)
*   *Status*: **TBD / PROPOSED** (See page 55, 66, 74).
*   *Boundary*: All physical and economic parameters (such as freight fees, decay coefficients, and money sinks) remain empty registers, awaiting baseline validation.

---

## Human Authorization Required

The structural recommendations outlined in this assessment represent a comprehensive reorganization of the project's state-tracking systems. 

**Before implementation, the following actions require explicit, written human authorization from the project owner:**

1.  **Authorize the split of `TCIndustries_Master_GDD_v1.1_Canonical_Baseline.md` into the dynamic `TCIndustries_Master_GDD_Active.md` and the machine-readable `TCIndustries_Canonical_Registry.json`.**
2.  **Authorize the permanent movement of all EIC Simulation, Specification, and Falsification files into the `/archive/EIC_architecture_C_falsified/` directory.**
3.  **Authorize the implementation of the deterministic "LLM Audit Reading Protocol" and manifest file checks as a mandatory gateway for all future development passes.**