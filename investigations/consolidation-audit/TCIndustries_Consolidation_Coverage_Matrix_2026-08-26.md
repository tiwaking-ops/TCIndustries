# TCIndustries – Consolidation Coverage Matrix (Audit Artifact)

**Date:** 2026-08-26  
**Purpose:** Demonstrate that no substantive information has been discarded from the supplied corpus during consolidation. This matrix is for human review and is **not** part of the active authority set.

**Method:** For each source file, list the key substantive items and their disposition.
**Author:** Author LLM Unknown
**Filing note (2026-09-12):** Renamed from `TCIndustries_Authority_Provenance_Reconciliation_Matrix.md` (duplicate filename, distinct document) to prevent collision with the governance Matrix; original name recorded here.

---

## Legend

| Column | Meaning |
|--------|---------|
| Source | Original file name |
| Key Content | Substantive items (rulings, principles, constraints, findings) |
| Active Representation | Where content appears in the active set (if any) |
| Archive Status | Whether original file is retained in full |
| Transformation | How content is represented (verbatim, summarised, etc.) |

---

## Matrix

| Source | Key Content | Active Representation | Archive Status | Transformation |
|--------|-------------|-----------------------|----------------|----------------|
| `TCIndustries_Master_GDD_v1.1.1_Status_Patch.md` | All LOCKED principles, derived constraints, PROPOSED, TBD, open questions | **Active** – the file itself | N/A | None |
| `TCIndustries_Authority_Provenance_Reconciliation_Matrix.md` | Authority classes A–G, GR‑001–003, status model, non‑canonical boundary | **Active** – the file itself | N/A | None |
| `TCIndustries_EIC_Human_Rulings_2026-08-25.md` | HD‑EIC‑01–04 | Consolidated in `TCIndustries_EIC_Human_Rulings_Consolidated.md` | Retained as source provenance | Summarised/verbatim |
| `TCIndustries_EIC_Human_Rulings_Architecture_C_Falsification_2026-08-25.md` | HD‑EIC‑05–08 | Consolidated in `TCIndustries_EIC_Human_Rulings_Consolidated.md` | Retained as source provenance | Summarised/verbatim |
| `TCIndustries_EIC_Human_Decision_Brief_v0.1.md` | Gate questions 1–4 | Represented in HD‑EIC‑01–04 | Retained as source provenance | Summarised/verbatim |
| `TCIndustries_EIC_Candidate_v0.1.md` | Proposed candidate architecture | Only the falsification outcome is active | Retained | Summarised in rulings |
| `TCIndustries_Master_GDD_v1.1_Canonical_Baseline.md` | Earlier GDD statuses | Not active; superseded by v1.1.1 | Retained | Historical |
| `TCIndustries_Master_GDD_v1.0_Canonical_Audit.md` | Audit findings leading to v1.1 corrections | Corrections already applied | Retained | Historical |
| `TCIndustries_Master_GDD_v1.0_Consolidated.md` | Original consolidation source | Superseded | Retained | Historical |
| `TCIndustries_Economic_Interdependence_Core_Design_Investigation.md` | Initial investigation, proposed Architecture C | Only the falsification outcome is active | Retained | Summarised in rulings |
| `TCIndustries_EIC_Candidate_v0.1_Falsification_Pass.md` | Adversarial attack on Architecture C | Not active; evidence | Retained | Summarised in rulings |
| `TCIndustries_EIC_Comparative_Simulation_Results_v0.1.1.md` | Simulation evidence | Not active; evidence | Retained | Summarised in rulings |
| `TCIndustries_EIC_Comparative_Simulation_Results_v0.1.md` | Simulation evidence | Not active; evidence | Retained | Summarised in rulings |
| `TCIndustries_EIC_Comparative_Simulation_Specification_v0.1.1.md` | Simulation spec | Not active; historical spec | Retained | Historical |
| `TCIndustries_EIC_Comparative_Simulation_Specification_v0.1.md` | Simulation spec | Not active; historical spec | Retained | Historical |
| `TCIndustries_EIC_Functional_Shapes_Refinement_Targeted_γ_Closure_v0.2.md` | Provisional refinements | Not active; used only in tested Architecture C | Retained | Historical |
| `TCIndustries_EIC_Minimum_Simulation_Results_v0.1.md` | Earlier simulation results | Not active; evidence | Retained | Historical |
| `TCIndustries_EIC_Provisional_Functional_Shapes_v0.1.md` | Functional shapes | Not active; used only in tested Architecture C | Retained | Historical |
| `TCIndustries_EIC_Targeted_γ_Prime_Adversarial_Retest_Results_v0.1.md` | Final retest evidence | Not active; evidence of falsification | Retained | Summarised in rulings |

---

## Coverage Summary

| Category | Active Representation | Archive Retention |
|----------|------------------------|-------------------|
| Current GDD | Yes (v1.1.1) | N/A |
| Governance | Yes (Matrix) | N/A |
| Human rulings (all) | Yes (Consolidated) | Originals retained |
| Human decision brief | Yes (represented in rulings) | Retained as source provenance |
| Candidate architecture | No (only conclusion) | Retained |
| Earlier GDD versions | No | Retained |
| EIC investigation & evidence | No (only conclusion) | Retained |
| Simulation evidence | No (only conclusion) | Retained |
| Functional shapes | No (only conclusion) | Retained |
| Falsification evidence | No (only conclusion) | Retained |

**Conclusion:** Within the source corpus inventoried by this matrix, no substantive information was intentionally discarded. Material not required for active authority is retained as historical, provenance, or evidence material according to its classification.

---

*End of Coverage Matrix*