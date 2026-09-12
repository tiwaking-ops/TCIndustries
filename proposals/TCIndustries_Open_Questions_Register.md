---
doc_id: OPEN-QUESTIONS-REGISTER
title: TCIndustries — Open Questions & Gates Register
version: "1.1"
date: 2026-08-25
status: active
authority_class: null   # index only; answers questions, decides nothing
human_approved: false
supersedes: []
superseded_by: null
controlling_documents: [MASTER-GDD-v1.1.1, HUMAN-RULINGS-REGISTER]
reading_priority: mandatory
archived_author_note: "Author LLM Unknown (filed 2026-09-12 as unadjudicated proposal; see proposals/README.md)"
change_note: "v1.1 — Item 5 (authority-review queue) upgraded from passive flag to an active gate (GDD-AUTH-REVIEW-GATE) per project-owner direction during the consolidation audit. WRLD-001, PLR-001, PLR-003, CRFT-001, SAFE-001 downgraded from HUMAN-LOCKED back to authority-review-queue status pending a re-review that captures decision rationale, not just outcome."
---

# TCIndustries — Open Questions & Gates Register

**This document lists what is currently unresolved. It resolves nothing.** When a question below is answered by a future human ruling or GDD status patch, this register is updated to point at that ruling — the ruling itself remains the authority, never this register.

---

## 1. Highest-Priority Open Gates

| ID | Question | Blocks | Status |
|---|---|---|---|
| **HD-EIC-07-GATE** | What structural source(s) of economic comparative advantage is the project willing to make a durable design commitment to? | **All further EIC architectural work** (no new architecture, no residual-channel invention, no Architecture D) | **OPEN — awaiting human statement.** See `TCIndustries_Human_Rulings_Register.md`, HD-EIC-07. |
| **GDD-AUTH-REVIEW-GATE** | Should WRLD-001, PLR-001, PLR-003, CRFT-001, and SAFE-001 remain HUMAN-LOCKED? | ~~Confidence in the GDD's LOCKED status tags for these 5 items~~ | **CLOSED — reconfirmed HUMAN-LOCKED.** See `TCIndustries_Human_Rulings_Register.md`, HD-GDD-01. |
| **FILENAME-CONVENTION-GATE** | Are there more unprefixed duplicate documents beyond the 4 already found? | Confidence that the archive is complete; risk of further silent document loss in future merges | **CLOSED for the 4 confirmed pairs — OPEN as a standing check.** See §5a below. Project owner should re-check their source folders for any other unprefixed `.md` files before the next merge/export. |

---

## 2. Master GDD Open Question Register (OQ-001–017)

Restated from `Master_GDD_v1.1.1_Status_Patch.md` §31 — not re-derived, not re-scoped.

| ID | Question | Affected Systems | Status |
|---|---|---|---|
| OQ-001 | Exact specialisation budget / skill-cap model? | Skills, Professions, Interdependence | TBD |
| OQ-002 | Resource spawn rates, durations, depletion mathematics? | Resources, Economy, Discovery | TBD |
| OQ-003 | Exact attribute sets for resources and mapping to product outcomes? | Resources, Crafting | TBD |
| OQ-004 | Crafting experimentation model details? | Crafting, Identity | TBD |
| OQ-005 | Manufacturing/factory throughput, maintenance, ownership rules? | Manufacturing, Automation | TBD |
| OQ-006 | Transportation model and fast-travel constraints? | World, Economy, Logistics | TBD |
| OQ-007 | City formation, governance, taxation/maintenance models? | Cities, Society, Buildings | TBD |
| OQ-008 | Combat scope (PvE/PvP), risk model, progression integration? | Combat, Economy, Services | TBD |
| OQ-009 | Currency system(s) and money sinks? | Economy | TBD |
| OQ-010 | Item durability, decay, repair model? | Items, Services, Economy | TBD |
| OQ-011 | Respecialisation exact costs and rules? | Progression, Identity | TBD |
| OQ-012 | Multi-accounting and economic exploit policies (implementation)? | Economy, Safety | TBD — *validation stance* (in-scope, no account restriction) is settled per HD-EIC-03; *implementation* is not |
| OQ-013 | Organisation rights, hierarchy, and property model? | Social Systems | TBD |
| OQ-014 | Species/ancestry/starting location design? | Character | DEFERRED |
| OQ-015 | Which systems are required for first playable prototype vs. persistent alpha? | Production planning | TBD |
| OQ-016 | Exact vendor and retail mechanics? | Retail, Economy | TBD |
| OQ-017 | Provenance depth and visibility rules? | Identity, Crafting, Items | TBD |

## 3. Assumption Register (AS-001–005)

| ID | Assumption | Status | Consequence if Wrong |
|---|---|---|---|
| AS-001 | A functional prototype (Seed-2.1 or similar) exists or will exist for testing | ASSUMPTION | Testing approach and reconciliation process must be redesigned |
| AS-002 | Data-driven content expansion is preferred over hard-coded systems | ASSUMPTION / PROPOSED | Content pipeline and technical architecture change |
| AS-003 | Regional economic differentiation is desirable and feasible | ASSUMPTION | World design and transport systems simplify or change |
| AS-004 | Meaningful provenance tracking is technically and design-feasible at scale | ASSUMPTION | Reputation systems may need alternative approaches |
| AS-005 | Non-combat careers can generate sufficient player demand without combat gating | ASSUMPTION | Demand generation systems may need redesign |

## 4. Remaining Human Decisions (GDD §36)

Not status issues — genuine unresolved design decisions requiring human/project approval before becoming LOCKED:

1. Final specialisation budget / skill-cap model.
2. Resource lifecycle mathematics and attribute set.
3. Crafting experimentation model.
4. Manufacturing and automation constraint details.
5. City formation and governance rules.
6. Combat scope and risk model.
7. Currency and economic stabiliser design.
8. Exact respecialisation costs and rules.
9. Multi-accounting and exploit policies (implementation detail; stance already settled — see OQ-012).
10. Reconciliation of any existing prototype against this GDD (Evidence Reconciliation Pass) — process requirement, not a design question.

## 5. GDD-AUTH-REVIEW-GATE — Resolved (Closed 2026-08-25)

**Status: CLOSED.** This gate is retained here, in resolved form, as the audit trail for how it was closed — it is no longer open.

**What happened:** `Authority_Provenance_Reconciliation_Matrix.md` flagged five GDD items as resting on "required by pillars/vision" language rather than a citable human decision. `Master_GDD_v1.1.1_Status_Patch.md` tagged all five `HUMAN-LOCKED (confirmed 2026-08-25)`, but no decision record existed to back that tag. During this consolidation audit, the five items were temporarily reverted to authority-review-queue status pending re-confirmation with rationale captured.

**Resolution:** The project owner located the original decision exchange (conducted across separate ChatGPT and Grok chat sessions prior to this consolidation effort) and supplied it directly. It is now recorded in full as **`HD-GDD-01`** in `TCIndustries_Human_Rulings_Register.md`, Section B. All five items are reconfirmed HUMAN-LOCKED — WRLD-001 and PLR-001 unqualified; PLR-003, CRFT-001, and SAFE-001 explicitly principle-only, each with a named open implementation surface.

**Root-cause finding (important beyond this one gate):** The project owner identified that the ChatGPT session which contributed to the original exchange **did not have access to the full project file set**. This is the specific mechanism that produced this instance of documentation drift — a partial-context LLM session made a real decision, but the record of that decision never made it into a document with awareness of the rest of the corpus, so the *outcome* got copied forward while the *process* was lost. See `TCIndustries_Changelog.md` for this recorded as a standing process-risk note, independent of this specific gate.

| ID | Principle (verbatim from GDD) | Status |
|---|---|---|
| WRLD-001 | Persistent Shared World | **HUMAN-LOCKED** (HD-GDD-01) |
| PLR-001 | Persistent Character Identity | **HUMAN-LOCKED** (HD-GDD-01) |
| PLR-003 | Ownership (meaningful persistent ownership) | **HUMAN-LOCKED, principle only** — implementation open (HD-GDD-01) |
| CRFT-001 | Differentiated Products | **HUMAN-LOCKED, principle only** — quality/experimentation/tiers open (HD-GDD-01) |
| SAFE-001 | Ownership and Transfer Safeguards | **HUMAN-LOCKED, principle only** — mechanism open (HD-GDD-01) |

## 5a. FILENAME-CONVENTION-GATE — Partially Resolved (Standing Check)

**Status: 4 confirmed pairs resolved; general risk remains open.**

**What happened:** At least one prior LLM session saved project documents without the required `TCIndustries_` prefix. Because the merged corpus used in this consolidation (`TCIndustries-project-files-v1-merged.txt`) was apparently gathered by prefix-matching, this caused **silent loss**, not just confusion — a properly-named document could be entirely absent from a merge with no error raised. Confirmed instances:

| Canonical (correctly prefixed) | Found without prefix | Resolution |
|---|---|---|
| `TCIndustries_EIC_Human_Validation_Gate_Architecture_C_Falsification_v0.1.md` | `EIC_Human_Validation_Gate_Architecture_C_Falsification_v0.1.md` | Reconstructed from original conversation context (this document was entirely absent from the merged upload); archived under canonical name. |
| `TCIndustries_EIC_Comparative_Simulation_Specification_v0.1.1.md` | `EIC_Comparative_Simulation_Specification_v0.1.1.md` | Extracted from the merged upload; restored to canonical filename. |
| `TCIndustries_EIC_Comparative_Simulation_Results_v0.1.md` | `EIC_Comparative_Simulation_Results_v0.1.md` | Extracted from the merged upload; already present under canonical filename in that source. |
| `TCIndustries_EIC_Comparative_Simulation_Results_v0.1.1.md` | `EIC_Comparative_Simulation_Results_v0.1.1.md` | Extracted from the merged upload; restored to canonical filename. |

**Correction to prior finding:** The earlier "naming collision" reported between the Test A/B and α/β/γ Comparative Simulation Specification documents (§C7 of `TCIndustries_Consolidated_Consensus_Plan_v1.md`) is **retracted**. Once correctly prefixed, `TCIndustries_EIC_Comparative_Simulation_Specification_v0.1.md` and `TCIndustries_EIC_Comparative_Simulation_Specification_v0.1.1.md` are genuinely distinct, non-colliding filenames. The apparent collision was an artifact of (a) the missing prefix hiding the true filename and (b) both documents' internal H1 headings imprecisely rendering as "v0.1" regardless of actual version. See the Erratum appended to the Consolidated Consensus Plan.

**Governing rule going forward:** GR-005 (Filename Convention Compliance) — see `TCIndustries_Manifest.md` §2.

**Remaining open item:** whether any *other* unprefixed duplicates exist beyond these 4 cannot be ruled out from this corpus alone. The project owner should check their own source folders directly; this cannot be verified from a merged export that may itself have already dropped the evidence.

## 6. Explicitly Non-Canonical (Do Not Resurrect Without New Human Decision)

Restated from `Authority_Provenance_Reconciliation_Matrix.md` §5 — historical proposals that remain non-canonical regardless of how many times they recur in older source material:

- Factories can never produce Exceptional-tier output
- Vitality / Stamina / Focus pools
- Skill Discipline / Skill Box architecture
- Use-based Skill Point acquisition
- Two-to-three concurrent profession target
- Toxicity attribute (as defined resource attribute)
- Time-limited vs. extraction-limited resource pool model (as a chosen model)
- Harvesters as installed structures (specific implementation)
- Experimentation point allocation / success-critical-failure model
- Single unified currency as the default
- Commerce Directory
- Fully tradeable by default / no account or character soul-binding
- Services required to be recurring/consumable by design

*End of Open Questions Register.*
