# TCIndustries – Project Snapshot Manifest

**Date of Snapshot:** 2026-08-26  
**Author:** Author LLM Unknown
**Purpose:** Provide a clean handover for LLMs receiving the TCIndustries project. This manifest defines the **minimum active authority set**, the archive, the authority hierarchy, and the rules governing this snapshot.

---

## 1. Active Documentation Set – Minimum Authority

The following four files are the **active authority-bearing documents**. They constitute the **minimum operational set for routine determination of current authority and state**, subject to the provenance and precedence rules stated in this manifest. Where the consolidated human‑rulings record conflicts with an original human ruling, the original human ruling controls.

Authority applies only according to each document's stated **status and provenance** (e.g., LOCKED/HUMAN‑LOCKED, DERIVED CONSTRAINT, PROPOSED, TBD, etc.).

| File | Role | Authority Level |
|------|------|----------------|
| `00_SNAPSHOT_MANIFEST.md` | This manifest – entry point, handover instructions, source mapping | Informational (not authoritative design) |
| `TCIndustries_Master_GDD_v1.1.1_Status_Patch.md` | **Current canonical design reference** – contains all LOCKED principles, derived constraints, PROPOSED directions, and open questions. | HUMAN‑LOCKED (for LOCKED principles) / DERIVED / PROPOSED / TBD |
| `TCIndustries_Authority_Provenance_Reconciliation_Matrix.md` | **Governance and authority framework** – defines authority classes (A–G), status rules, and the distinction between human approval and LLM inference. | HUMAN‑LOCKED (governance) |
| `TCIndustries_EIC_Human_Rulings_Consolidated.md` | **Binding human rulings for the Economic Interdependence Core (EIC)** – consolidates HD‑EIC‑01 through HD‑EIC‑08. No new authority is created; original rulings control in case of discrepancy. | HUMAN‑LOCKED (programme‑specific rulings) |

**All other files are archived and are not part of the active design baseline.** They may be consulted for provenance/evidence but do not override the active documents.

---

## 2. Authority Hierarchy (Explicit)

Following the Authority & Provenance Reconciliation Matrix:

1. **Explicit Human / Project Ruling (Class A)** – Highest authority.
2. **Explicit Project Principle / Pillar (Class B)** – Binding principles.
3. **Current Canonical GDD** – Representation of approved design (with its own internal statuses).
4. **Derived Constraints (Class C)** – Logical consequences, not independent decisions.
5. **Proposed / TBD / Assumption** – Non‑authoritative design material.
6. **Prototype / Evidence (Class G)** – Evidence only.
7. **Historical / Archived Material** – Provenance only; cannot override current authority.

---

## 3. Archived Material – Classification and Purpose

The archive contains three distinct types of material:

- **Historical proposals** – Superseded design directions (e.g., older GDDs, discarded mechanics).
- **Decision provenance** – Original source records of human rulings (retained for traceability).
- **Evidence** – Investigations, simulations, and falsification results that informed current rulings.

None of these grant design authority, but evidence and provenance may be consulted to understand **why** a current ruling or constraint exists.

The following files are moved to `/archive/` when the snapshot is activated:

- `TCIndustries_Master_GDD_v1.0_Consolidated.md`
- `TCIndustries_Master_GDD_v1.1_Canonical_Baseline.md`
- `TCIndustries_Master_GDD_v1.0_Canonical_Audit.md`
- `TCIndustries_Economic_Interdependence_Core_Design_Investigation.md`
- `TCIndustries_EIC_Candidate_v0.1.md`
- `TCIndustries_EIC_Candidate_v0.1_Falsification_Pass.md`
- `TCIndustries_EIC_Comparative_Simulation_Results_v0.1.1.md`
- `TCIndustries_EIC_Comparative_Simulation_Results_v0.1.md`
- `TCIndustries_EIC_Comparative_Simulation_Specification_v0.1.1.md`
- `TCIndustries_EIC_Comparative_Simulation_Specification_v0.1.md`
- `TCIndustries_EIC_Functional_Shapes_Refinement_Targeted_γ_Closure_v0.2.md`
- `TCIndustries_EIC_Human_Decision_Brief_v0.1.md` *(source provenance for HD‑EIC‑01–04)*
- `TCIndustries_EIC_Human_Rulings_2026-08-25.md` *(source provenance; consolidated into active record)*
- `TCIndustries_EIC_Human_Rulings_Architecture_C_Falsification_2026-08-25.md` *(source provenance; consolidated into active record)*
- `TCIndustries_EIC_Minimum_Simulation_Results_v0.1.md`
- `TCIndustries_EIC_Provisional_Functional_Shapes_v0.1.md`
- `TCIndustries_EIC_Targeted_γ_Prime_Adversarial_Retest_Results_v0.1.md`

---

## 4. How to Use This Snapshot

An LLM receiving this snapshot should:

1. **Read this manifest first** to understand the active set, authority hierarchy, and archive classification.
2. **Read the GDD** (`TCIndustries_Master_GDD_v1.1.1_Status_Patch.md`) as the primary design reference.
3. **Read the Authority Matrix** (`TCIndustries_Authority_Provenance_Reconciliation_Matrix.md`) to understand the status system and authority classes.
4. **Read the Consolidated Human Rulings** (`TCIndustries_EIC_Human_Rulings_Consolidated.md`) to know the current EIC state and binding rulings.
5. **Consult the archive only if** you need detailed provenance/evidence (e.g., why a ruling was made). Archived documents are **not** design authority.

**Critical rules:**
- No status may be upgraded (e.g., PROPOSED → LOCKED) without explicit human ruling.
- Historical proposals (e.g., skill‑box architecture, specific resource models) are **not** part of the current baseline.
- Prototype behaviour is evidence, not design authority (see GDD PROT‑001).
- The EIC is **not yet designed**; the GDD states that the Economic Interdependence Core is the next substantive design phase after prototype evidence reconciliation.
- **This snapshot is a documentation consolidation/handover operation only. It does not authorise further consolidation, rewriting, expansion, or reinterpretation of the Master GDD.**

---

## 5. Source‑to‑Destination Mapping (Document Level)

| Original File | Disposition | Notes |
|---------------|-------------|-------|
| `TCIndustries_Master_GDD_v1.1.1_Status_Patch.md` | **Active** | Current GDD |
| `TCIndustries_Authority_Provenance_Reconciliation_Matrix.md` | **Active** | Governance framework |
| `TCIndustries_EIC_Human_Rulings_2026-08-25.md` | **Archived – source provenance** | Content consolidated into active rulings record |
| `TCIndustries_EIC_Human_Rulings_Architecture_C_Falsification_2026-08-25.md` | **Archived – source provenance** | Content consolidated into active rulings record |
| `TCIndustries_EIC_Human_Decision_Brief_v0.1.md` | **Archived – source provenance** | Source for HD‑EIC‑01–04 gate questions |
| `TCIndustries_EIC_Candidate_v0.1.md` | **Archived** | Proposed candidate architecture; falsified |
| `TCIndustries_Master_GDD_v1.1_Canonical_Baseline.md` | **Archived** | Superseded by v1.1.1 |
| `TCIndustries_Master_GDD_v1.0_Canonical_Audit.md` | **Archived** | Historical audit; corrections already applied |
| `TCIndustries_Master_GDD_v1.0_Consolidated.md` | **Archived** | Superseded consolidation source |
| All other EIC investigation, simulation, falsification, and functional‑shapes files | **Archived** | Historical evidence; conclusions captured in consolidated rulings |

---

## 6. Current EIC State (Authoritative State from Active Human Rulings)

The following state is reproduced from the active consolidated human rulings, which derive their authority from explicit human decisions. Archived investigations and simulation evidence provide provenance and evidentiary support but do not themselves grant design authority.

- **Architecture C** is **falsified / retired in its tested form** (HD‑EIC‑05). This does **not** prohibit reuse of every individual concept; any reuse would require fresh evaluation under the current human‑authorised boundaries.
- **PIL‑003** (interdependence pillar) remains **HUMAN‑LOCKED** and unchanged (HD‑EIC‑06).
- **Layered validation unit** (character vs. independent economic participant) is locked (HD‑EIC‑01).
- **Multi‑accounting** is an in‑scope adversarial condition (HD‑EIC‑03).
- **No hard factory‑quality ceiling** is permitted; specialist advantage must be tested under quality parity (HD‑EIC‑02).
- **Bounded new shaping** is authorised **in principle**, but **execution is gated on explicit human‑defined structural boundaries** (HD‑EIC‑07).
- **Numeric tuning, Provenance & Reputation engine, and automatic Architecture D are not authorised** without explicit human boundaries (HD‑EIC‑08).
- The **next required action** is human definition of the structural boundaries for a new EIC shaping effort. Until that occurs, no new architectural work may begin.

---

## 7. Human‑Review / Authorisation Queue

The following actions are not yet taken and require explicit human authorisation:

1. [ ] Move the four active documents to an `/active/` directory (or keep at root, as preferred).
2. [ ] Move all archived documents to an `/archive/` directory.
3. [ ] Confirm that the consolidated human rulings document is the operational record and that the original rulings are retained in the archive as source provenance.
4. [ ] Confirm that this manifest and the accompanying audit artifacts are acceptable.

Once authorised, this snapshot becomes the official point‑in‑time representation of the TCIndustries project.

---

**End of Manifest**