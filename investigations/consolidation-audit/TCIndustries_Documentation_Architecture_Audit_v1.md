# TCIndustries — Documentation Consolidation & Cross-LLM Auditability Assessment

**Document Type:** Independent Documentation Architecture / Governance Audit (Proposal Only)
**Auditor Role:** Documentation Architect & Evidence/Governance Auditor
**Date of Assessment:** 2026-08-25 (corpus as supplied)
**Authority of This Document:** None. This is an external assessment. It makes no game-design decisions, resolves no open questions, and promotes no proposal to canon. All recommendations require human authorization before any file is created, renamed, merged, or retired.
**Corpus Reviewed:** 19 supplied TCIndustries documents (Master GDD lineage, Authority & Provenance Reconciliation Matrix, Economic Interdependence Core investigation/candidate/falsification/simulation/rulings chain) plus the merged upload of the same corpus.
**Author:** Claude (per Consolidated Consensus Plan inputs list — this instance's own prior audit; unverified)

---

## 1. Executive Verdict

**The project's governance *vocabulary* (LOCKED, DERIVED CONSTRAINT, PROPOSED, TBD, PROTOTYPE/EVIDENCE, DEFERRED, ASSUMPTION, HISTORICAL, Authority Classes A–G) is unusually rigorous and worth preserving unchanged.** The discipline of GR-001 ("canonical presence is not human approval") is genuinely followed in most documents — status tags are applied conservatively, and several documents actively self-limit their own authority ("Creates no design authority," "Authority: None").

**However, the *file-level architecture* has not kept pace with the governance vocabulary.** State is currently reconstructable only by reading a long prose lineage chain across ~13 EIC documents plus ~4 GDD documents, with no manifest, no changelog, and no single "what is true right now" document. Two independent LLMs given this corpus and asked "what is the current state of Architecture C?" would very plausibly diverge, because the authoritative answer (Architecture C is **RETIRED**, per `TCIndustries_EIC_Human_Rulings_Architecture_C_Falsification_2026-08-25.md`, HD-EIC-05) sits at the end of a 13-document evidentiary chain, while a document that *looks* current and canonical (`TCIndustries_Master_GDD_v1.1.1_Status_Patch.md`, §34) still names "Economic Interdependence Core (cross-system package)" as the *next* design phase without any acknowledgment that its leading candidate has since been falsified and retired.

This is not a hypothetical risk — it is a **demonstrated inconsistency already present in the current corpus** (see §5). It is the exact failure mode the project's own governance rules exist to prevent, and it has occurred not through LLM misbehavior but through **normal, correct, incremental document production without a synchronized state layer.**

**Recommendation in one sentence:** Keep every existing document (nothing should be deleted), but introduce a small number of new **pointer/state documents** — a Project State Manifest, a consolidated EIC current-state summary, a consolidated Open Questions/Gates register, and a chronological change log — so that "what is true right now" is answered by reading 3–4 short files rather than reconstructing it from prose lineage notes scattered across 19 documents.

---

## 2. Current Documentation Assessment

Scored against the eight assessment criteria requested:

| Criterion | Assessment | Evidence |
|---|---|---|
| **A. State Reconstruction** | **Weak.** No single document states current project state end-to-end. An LLM must read the full GDD lineage (4 docs) *and* the full EIC lineage (13 docs) and correctly order them by cross-referenced "Controlling References" prose to determine current state. | No manifest or index document exists in the corpus. |
| **B. Authority Reconstruction** | **Strong in vocabulary, moderate in practice.** Authority Classes A–G (`Authority_Provenance_Reconciliation_Matrix.md` §1) are well defined and consistently referenced. But several "HUMAN-LOCKED" tags (e.g., WRLD-001, PLR-001, PLR-003, CRFT-001, SAFE-001 in `Master_GDD_v1.1.1_Status_Patch.md`) assert human confirmation occurred without the corpus containing independent evidence of that confirmation beyond the patch document's own assertion — see §5. |
| **C. Temporal Reconstruction** | **Weak.** No document carries an explicit `supersedes` / `superseded-by` field. Supersession is only inferable from prose ("This patch applies... to `Master_GDD_v1.1_Canonical_Baseline.md`"). Two "Comparative Simulation Specification" documents share the label "v0.1" internally while differing in filename version (v0.1 vs v0.1.1) — see §4. |
| **D. Provenance** | **Good.** Nearly every document has a "Controlling References" and "Lineage" block naming its direct predecessor(s). This is the strongest part of the corpus and should be preserved as-is; it just needs to be indexed rather than only embedded in prose. |
| **E. Cross-LLM Comparability** | **Weak.** Two independent LLMs would likely disagree on (1) whether Architecture C is currently viable, (2) whether the EIC "next authorized operation" is a new simulation or waiting on a human structural-boundary statement, and (3) which of the two "Comparative Simulation Specification v0.1" documents is controlling for which package set. |
| **F. Change Detection** | **Weak.** No changelog exists. Each document's own "Document Control" footer records what *it* changed, but there is no cumulative record, so "what changed since last time" requires diffing dates across all 19 files. |
| **G. Scope Control** | **Strong.** This is the corpus's best-executed property. Documents are scrupulous about phrases like "Creates no design authority," "Does not promote Architecture C," "No mechanism in this document is LOCKED... unless already established." HISTORICAL/non-canonical boundaries (Authority Matrix §5) are explicit and specific. |
| **H. Scalability** | **At risk.** The EIC track alone required 13 documents to go from investigation to retirement of a single architecture candidate. If this per-topic document velocity continues project-wide (skills, resources, crafting, manufacturing, cities, combat...), the corpus will grow to hundreds of files within a few project phases unless a consolidation discipline is introduced now. |

---

## 3. Current Documentation Inventory

Classified by genuine function (not by the label the document gives itself), with explicit note of authority basis (Authority Class per the Reconciliation Matrix framework, applied here for consistency, not because the source documents used that exact class letter for themselves in all cases).

| # | Document | Function | Authority Class | Current / Superseded |
|---|---|---|---|---|
| 1 | `Master_GDD_v1.0_Consolidated.md` | Original consolidation of first-pass GDDs | E (multi-source inference) for most content; B for vision/pillars | **Superseded** by v1.1 |
| 2 | `Master_GDD_v1.0_Canonical_Audit.md` | Audit of doc 1 | D (audit/governance) | Historical audit record (its recommendations were applied, not itself canon) |
| 3 | `Master_GDD_v1.1_Canonical_Baseline.md` | Applies audit corrections to doc 1 | Mixed A/B/C per rule | **Superseded** by v1.1.1 |
| 4 | `Authority_Provenance_Reconciliation_Matrix.md` | Governance checkpoint reviewing v1.1; defines Authority Classes A–G, GR-001–003 | D (governance rule) | **Current** — foundational governance reference |
| 5 | `Master_GDD_v1.1.1_Status_Patch.md` | Applies human rulings from doc 4 to doc 3 | A/B/C per rule (explicit) | **Current** canonical GDD |
| 6 | `Economic_Interdependence_Core_Design_Investigation.md` | Non-canonical investigation proposing Architectures A/B/C | E (inference) / F (external evidence, SWG/EVE) | Historical — informed doc 7, superseded in relevance by doc 19 |
| 7 | `EIC_Candidate_v0.1.md` | Converts Architecture C into M1–M8 mechanisms | PROPOSED, no authority | **Superseded/retired** (see doc 19) |
| 8 | `EIC_Candidate_v0.1_Falsification_Pass.md` | Adversarial attack on doc 7 | Evidence/analysis, no authority | Historical evidence record |
| 9 | `EIC_Comparative_Simulation_Specification_v0.1.md` | Defines Test A / Test B (factory parity, formal budget) | Specification, no authority | Historical — superseded in scope by doc 13's α/β/γ specification |
| 10 | `EIC_Minimum_Simulation_Results_v0.1.md` | Executes doc 9; result INCONCLUSIVE / under-specified | Evidence (Class G-equivalent) | Historical evidence record |
| 11 | `EIC_Human_Decision_Brief_v0.1.md` | Presents Gates 1–4 arising from doc 10's under-specification | Decision brief, no authority until ruled | **Resolved** by doc 12 |
| 12 | `EIC_Human_Rulings_2026-08-25.md` | Rules Gates 1–4 → **HD-EIC-01–04** | **A — explicit human ruling** | **Current** — HUMAN-LOCKED (principles only) |
| 13 | `EIC_Provisional_Functional_Shapes_v0.1.md` | Non-numeric functional shapes issued under HD-EIC-04 | PROPOSED / simulation assumption | Historical — superseded in scope by doc 16 (γ-only refinement) |
| 14 | `EIC_Comparative_Simulation_Specification_v0.1.1.md` *(filename v0.1.1; internal header says "Version: 0.1, post-Functional-Shapes revision")* | Defines α/β/γ package testing | Specification, no authority | Historical — controlling spec for doc 15's results |
| 15 | `EIC_Comparative_Simulation_Results_v0.1.md` | Executes doc 14; α FAIL, β FAIL, γ INCONCLUSIVE | Evidence | Historical evidence record |
| 16 | `EIC_Functional_Shapes_Refinement_Targeted_γ_Closure_v0.2.md` | Narrows the 4 relationships blocking γ | PROPOSED / simulation assumption | Historical — superseded in relevance by doc 17 |
| 17 | `EIC_Targeted_γ_Prime_Adversarial_Retest_Results_v0.1.md` | Tests γ′-1/γ′-2 per doc 16; both FAIL | Evidence | Historical evidence record |
| 18 | `EIC_Human_Validation_Gate_Architecture_C_Falsification_v0.1.md` | Presents completed evidence, requests HD-EIC-05–08 | Decision gate, no authority until ruled | **Resolved** by doc 19 |
| 19 | `EIC_Human_Rulings_Architecture_C_Falsification_2026-08-25.md` | Rules HD-EIC-05–08: **Architecture C RETIRED**; bounded new shaping authorized pending human structural boundaries | **A — explicit human ruling** | **Current — most recent EIC-track state** |

**Net current state, as established by the corpus itself (not by this auditor's inference):**
- **Canonical design reference:** `Master_GDD_v1.1.1_Status_Patch.md` (item 5).
- **Canonical EIC state:** Architecture C is **retired** (item 19, HD-EIC-05). No architecture is currently the leading hypothesis. No further EIC architectural work is authorized until the human supplies structural boundaries (HD-EIC-07). This is the single most important fact in the entire corpus for any downstream LLM, and it currently has no dedicated, easily located home.
- **Governing rulings currently HUMAN-LOCKED:** HD-EIC-01–04 (item 12) and HD-EIC-05–08 (item 19), plus whatever subset of the GDD's LOCKED/HUMAN-LOCKED tags is genuinely class A/B per item 4's authority-review queue.

---

## 4. Redundancy and Sprawl Analysis

**4.1 Genuine duplication (same information, multiple homes).** Low. The corpus is largely non-duplicative in *content* — each document adds a distinct evidentiary or decision step. The GDD lineage (items 1, 3, 5) does restate the full pillar/vision text three times nearly verbatim across versions, which is expected for a "baseline" pattern but means an LLM naively reading all three will process ~3x the pillar text without new information.

**4.2 Naming collisions (different information, confusable homes) — this is the corpus's real sprawl risk:**

- **Two "Comparative Simulation Specification" documents claim overlapping version identity.** `EIC_Comparative_Simulation_Specification_v0.1.md` (item 9) and `EIC_Comparative_Simulation_Specification_v0.1.1.md` (item 14) are materially different specifications (Test A/B vs. α/β/γ packages) governing different result documents, yet item 14's *internal* header states `**Version:** 0.1 (post-Functional-Shapes revision)` — not 0.1.1. A downstream LLM sorting by internal version metadata (rather than filename) would treat these as the same version and could cite the wrong controlling spec for a given results document.
- **Two "Results v0.1" documents are one keyword-search away from being conflated:** `EIC_Minimum_Simulation_Results_v0.1.md` (item 10, controlled by item 9) and `EIC_Comparative_Simulation_Results_v0.1.md` (item 15, controlled by item 14). Both are plausible hits for a search on "EIC simulation results."
- **Two "Human Rulings" documents share a date and name prefix but cover disjoint ruling ID ranges:** `EIC_Human_Rulings_2026-08-25.md` (HD-EIC-01–04) and `EIC_Human_Rulings_Architecture_C_Falsification_2026-08-25.md` (HD-EIC-05–08). This is workable because the second filename is more specific, but nothing in either *filename* signals the ID range it covers, so an LLM citing "HD-EIC-06" has to open both files to find it.

**4.3 Structural sprawl driver.** The EIC track shows a clear pattern: **investigate → candidate → falsify → specify → simulate → results → refine → re-specify → re-simulate → results → gate → rule**, each step its own document (13 documents for one architecture's full lifecycle, ending in retirement). This pattern is methodologically sound (it is exactly the "candidate → attack → simulation → evidence → human ruling" discipline the documents themselves repeatedly assert), but at this document-per-step rate, a second architecture attempt (post-HD-EIC-07) will add another 10+ documents, and the project has ~30 more systems to design (per GDD §30/§31). **Without a consolidation discipline, this single pattern alone could produce 200+ documents before the game's core mechanical design is even drafted.**

**4.4 What should NOT be reduced.** The step-by-step evidentiary chain (investigate/candidate/falsify/specify/simulate/results/rule) should **not** be collapsed into fewer documents going forward — this is exactly the audit trail that makes the project's negative result (Architecture C retirement) credible and falsifiable. The sprawl problem is not that these documents exist; it is that **nothing points at them as a set** once their evidentiary job is done.

---

## 5. Authority and Provenance Risks

**5.1 Confirmed inconsistency — stale "next steps" pointer in current canon.**
`Master_GDD_v1.1.1_Status_Patch.md` §34 ("Recommended Design Sequence") states the next substantive design phase is "**Economic Interdependence Core (cross-system package)**" and its closing line repeats "the next operations are... then the Economic Interdependence Core package." This document is dated 2026-08-25, the same date as `EIC_Human_Rulings_Architecture_C_Falsification_2026-08-25.md`, which records that Architecture C — the EIC's leading hypothesis — has been **retired**, and that no further EIC architectural work is authorized until the human supplies structural boundaries. The GDD's pointer is not *false* (an Economic Interdependence Core package genuinely is still the design objective), but taken at face value by an LLM that reads only the GDD, it implies EIC work can proceed directly to system design, when the actual controlling state (per the more specific and more recent EIC ruling) is that **EIC architectural work is currently blocked pending a human structural-boundary statement.** This is a direct instance of the exact failure mode the project's own governance discipline warns against (a document being read as more current/authoritative than it is), and it exists in the corpus **today**, not hypothetically.

**5.2 "HUMAN-LOCKED (confirmed 2026-08-25)" tags rest on the patch document's own assertion.**
`Master_GDD_v1.1.1_Status_Patch.md` marks WRLD-001, PLR-001, PLR-003, CRFT-001, and SAFE-001 as `LOCKED — HUMAN-LOCKED (confirmed 2026-08-25)`. The `Authority_Provenance_Reconciliation_Matrix.md` (item 4) explicitly flagged these five as an **"authority-review queue"** — items resting on "required by pillars/vision" language rather than an independently documented human ruling, requiring "explicit human confirmation before being treated as frozen." The v1.1.1 patch (item 5) asserts this confirmation happened and applies it. The corpus, as supplied, contains **no separate artifact recording the confirmation act itself** (e.g., no HD-EIC-style ruling document analogous to items 12/19 for these five GDD items). This may simply mean the confirmation record exists outside the supplied corpus — the corpus does not establish that it does *not* exist — but as supplied, an LLM cannot independently verify this class-A claim; it can only note that the patch document asserts it. Flagged per **GR-001** (canonical presence is not itself human approval) — the patch document is, structurally, the kind of statement GR-001 is designed to make an LLM skeptical of, even though its authors clearly intended good-faith compliance.

**5.3 Authority Matrix vs. GDD patch have a subtle sequencing dependency that is easy to invert.**
Read in isolation, `Authority_Provenance_Reconciliation_Matrix.md` (item 4) reads as if it is *itself* the terminal governance state (it proposes demotions and flags a review queue but "creates no new game-design decisions"). It is only `Master_GDD_v1.1.1_Status_Patch.md` (item 5) that reports these proposals as *applied*. An LLM given only the Matrix, without also being told the Patch exists and supersedes it in effect, would correctly conclude CRFT-006 and PROG-002 are still LOCKED (their pre-Matrix status) — the opposite of current truth.

**5.4 No conflicting *decisions* were found.** To be clear: the audit found no case of two HUMAN-LOCKED rulings contradicting each other. The risks above are all **staleness/pointer risks and unverifiable-confirmation risks**, not decision conflicts. This is a meaningfully better finding than it could have been, and reflects well on the discipline applied while producing each individual document.

---

## 6. Cross-LLM Auditability Assessment

Applying the project's own three success tests conceptually to the *current* corpus (before any restructuring):

- **Test 1 (State Reconstruction) — currently FAILS.** A new LLM cannot determine "is Architecture C viable" without reading, at minimum, items 12, 14, 15, 16, 17, 18, 19 in the correct order (7 of the 19 documents), because no single document states the terminal answer plus its full authority chain.
- **Test 2 (Audit Reproducibility) — currently AT RISK.** Two LLMs are likely to agree on high-level pillars (well-labeled, low ambiguity) but are likely to *disagree* on: (a) which Comparative Simulation Specification governs which results document (§4.2), (b) whether the GDD's §34 "next phase" pointer is still operative (§5.1), and (c) whether the five "authority-review queue" GDD items should be treated as confirmed or still-pending (§5.2).
- **Test 3 (Historical Containment) — currently PASSES.** This is the corpus's strength. Every historical/superseded/non-canonical document is explicit about its own non-authority ("Creates no design authority," "Authority: None," explicit NOT CANONICAL tables in item 4 §5). No document in the corpus silently treats a proposal as decided. The discipline works; it is the *indexing* of that discipline across files that is missing.

---

## 7. Recommended Documentation Architecture

**Principle: add a thin, explicit "state layer" on top of the existing evidentiary/design corpus; do not restructure the evidentiary corpus itself.**

The architecture has three tiers:

1. **State Layer (new, small, high-update-frequency, always read first).** A manifest plus a handful of consolidated "current state" summaries. These documents contain *no new design decisions* — they only point to and summarize decisions already made elsewhere, with citations back to the authoritative document. This is the layer that is currently missing.
2. **Active Canonical Layer (existing documents, kept as-is, low update frequency).** The current Master GDD and the Authority & Provenance Reconciliation Matrix. These remain the actual authority; the manifest points to them rather than duplicating their content.
3. **Archive / Evidentiary Layer (existing documents, kept as-is, effectively frozen).** Every investigation, candidate, falsification pass, specification, results report, and superseded GDD version. These are never edited, only added to. Their existing "Controlling References"/"Lineage" sections are preserved unchanged — they are good provenance records and should not be rewritten.

This directly avoids the two extremes the brief warns against: it does not merge the corpus into one document (which would destroy the falsifiability/evidentiary value of the step-by-step EIC chain), and it does not add dozens of new narrow documents (it adds four).

---

## 8. Active Documentation Set

For each proposed **new or actively-maintained** document:

### 8.1 `PROJECT_STATE_MANIFEST.md` *(new)*
- **Purpose:** Single entry point. Machine-readable summary of current canonical status, current EIC status, currently HUMAN-LOCKED ruling IDs, and pointers (not copies) to the documents that establish each.
- **Authority:** None — it is a pointer/index, not a decision-making document. Every claim in it must cite its source document.
- **Audience:** Every LLM performing any TCIndustries task, read first, always.
- **Belongs inside:** Current GDD version + filename; current EIC architecture status (currently: none, retired, pending structural boundaries); list of currently HUMAN-LOCKED ruling IDs (HD-EIC-01–08, and GDD LOCKED items) with source-document citation; "authority-review queue" items still open; last-updated date; pointer to the change log.
- **Must NOT contain:** Any new design content, any restatement of pillar text, any resolution of open questions.
- **Relationship to other documents:** Points to §8.2 and §8.4 for detail; never duplicates their content.
- **Update frequency:** On every human ruling or GDD status change (i.e., whenever a new HD-EIC-xx or GDD status patch is issued).
- **Read during audit?** Yes — always, first.

### 8.2 `TCIndustries_Master_GDD_[current-version].md` *(existing — currently `v1.1.1_Status_Patch.md`)*
- **Purpose:** Canonical design reference (vision, pillars, LOCKED/DERIVED CONSTRAINT/PROPOSED/TBD register).
- **Authority:** Authoritative — highest for game-design principles (Class A/B per item, subject to §5.2's caveat on the five queued items).
- **Audience:** Any LLM doing design work; read second, after the manifest.
- **Belongs inside:** Exactly what it currently contains.
- **Must NOT contain:** EIC architectural detail beyond the principle level (this already correctly holds — EIC mechanics live in the EIC track, not the GDD).
- **Relationship to other documents:** Superseded in place going forward — future revisions should overwrite/replace this file's *role* (with the old version moved to archive) rather than producing a fourth parallel "vX.Y Baseline"/"Patch" file, unless a genuine audit is being performed and needs its own document (as v1.0 Canonical Audit correctly did).
- **Update frequency:** Low — only on explicit human status ruling.
- **Read during audit?** Yes.

### 8.3 `TCIndustries_Authority_Provenance_Reconciliation_Matrix.md` *(existing, unchanged)*
- **Purpose:** Defines the authority-class vocabulary (A–G) and governance rules (GR-001–003) that every other document's status tags depend on.
- **Authority:** Governance-level (Class D) — authoritative for *how to interpret* status tags, not for game design itself.
- **Audience:** Any LLM needing to adjudicate a status/authority question.
- **Belongs inside:** Exactly what it currently contains.
- **Must NOT contain:** New game-design content.
- **Relationship to other documents:** Read whenever an LLM needs to classify a statement it finds elsewhere; not needed for routine design-reading.
- **Update frequency:** Very low.
- **Read during audit?** Yes, but only as a reference/rules document, not for content.

### 8.4 `TCIndustries_EIC_Current_State.md` *(new — consolidates, does not replace, items 12/19)*
- **Purpose:** A short, current-state-only summary of the EIC track: what architecture (if any) is the leading hypothesis, what is retired, what HD-EIC-01–08 established, and what the explicitly next-authorized operation is.
- **Authority:** None of its own — every statement must cite the HD-EIC ruling that establishes it (items 12 and 19 remain the actual authority).
- **Audience:** Any LLM asked to continue, audit, or reference EIC work.
- **Belongs inside:** A restatement (not reinterpretation) of the "Programme State" tables already present in items 12 and 19 (e.g., item 19 §3), merged into one place, with the retirement of Architecture C stated in the first sentence.
- **Must NOT contain:** Any new EIC design content, new residual channels, or numeric values (same prohibitions the underlying documents already carry).
- **Relationship to other documents:** Points into the archive (items 6–18) for full evidentiary detail; those documents remain the record of *how* the state was reached.
- **Update frequency:** On every new HD-EIC ruling.
- **Read during audit?** Yes, whenever the audit touches EIC/economic design.

### 8.5 `TCIndustries_Open_Questions_and_Gates_Register.md` *(new — consolidates existing OQ tables)*
- **Purpose:** One place listing every currently open design question (OQ-001…OQ-017 from the GDD, plus any future open EIC "HD-" gates such as the structural-boundary statement required by HD-EIC-07) with status and the document that should be updated when each is resolved.
- **Authority:** None — index only.
- **Audience:** Human project owner deciding what to rule on next; any LLM checking "is X already decided."
- **Belongs inside:** OQ ID, question text, affected systems, current status, resolving-document-once-resolved.
- **Must NOT contain:** Answers to the questions (that would be silently resolving them).
- **Relationship to other documents:** Sources from GDD §31 and any EIC-track open-decision tables (e.g., item 19 §6, item 11 §5–6). When a question is resolved by a human ruling, this register is updated to point at the ruling document; the ruling document itself remains the authority.
- **Update frequency:** Whenever a new open question is raised or resolved.
- **Read during audit?** Yes, when identifying unresolved items.

### 8.6 `TCIndustries_Change_Log.md` *(new)*
- **Purpose:** Chronological, append-only log of state-changing events: every HD-EIC ruling, every GDD status patch, every architecture retirement/promotion.
- **Authority:** None — factual log only, each entry cites its source document.
- **Audience:** Any LLM asked "what changed since I last looked" or "what happened in order."
- **Belongs inside:** Date, event, one-line description, source document.
- **Must NOT contain:** Analysis, interpretation, or new decisions.
- **Relationship to other documents:** Derived entirely from Document Control footers already present in every existing document — this file adds no new information, only order and a single reading location.
- **Update frequency:** Every state change.
- **Read during audit?** Yes, for "what's new since—" queries.

---

## 9. Archive / Provenance Structure

**Recommendation: a single flat `archive/` (or equivalently tagged) collection, not a deep taxonomy**, because the corpus is not yet large enough to need sub-folders, and the existing "Controlling References"/"Lineage" fields already provide the navigational structure a taxonomy would otherwise need to reproduce.

Proposed archive membership (no files moved — this is a classification proposal only):

| Category | Documents |
|---|---|
| **Superseded canonical GDD versions** | `Master_GDD_v1.0_Consolidated.md`, `Master_GDD_v1.1_Canonical_Baseline.md` |
| **Audit records** | `Master_GDD_v1.0_Canonical_Audit.md` |
| **EIC investigation / candidate history** | `Economic_Interdependence_Core_Design_Investigation.md`, `EIC_Candidate_v0.1.md` |
| **EIC adversarial/evidence records** | `EIC_Candidate_v0.1_Falsification_Pass.md`, `EIC_Comparative_Simulation_Specification_v0.1.md`, `EIC_Minimum_Simulation_Results_v0.1.md`, `EIC_Comparative_Simulation_Specification_v0.1.1.md`, `EIC_Comparative_Simulation_Results_v0.1.md`, `EIC_Provisional_Functional_Shapes_v0.1.md`, `EIC_Functional_Shapes_Refinement_Targeted_γ_Closure_v0.2.md`, `EIC_Targeted_γ_Prime_Adversarial_Retest_Results_v0.1.md` |
| **EIC decision-gate records (resolved)** | `EIC_Human_Decision_Brief_v0.1.md`, `EIC_Human_Validation_Gate_Architecture_C_Falsification_v0.1.md` |
| **EIC ruling records (terminal authority — kept in archive but cited by the state layer, not duplicated)** | `EIC_Human_Rulings_2026-08-25.md`, `EIC_Human_Rulings_Architecture_C_Falsification_2026-08-25.md` |

Note: the two EIC "Human Rulings" documents are **archival in the sense of "not re-edited,"** but they remain the **actual authority** — `TCIndustries_EIC_Current_State.md` (§8.4) must cite them directly rather than treating archive status as reduced authority. Archive = "frozen and not read by default," not "demoted in authority."

---

## 10. LLM Audit Reading Protocol

A deterministic minimal reading order for an independent LLM performing a standard TCIndustries audit:

1. **Read `PROJECT_STATE_MANIFEST.md` first, in full.** This establishes current GDD version, current EIC status, and pointers to everything else. Do not read anything else until this is done.
2. **Identify the current canonical state** by following the manifest's pointer to the current Master GDD file and to `TCIndustries_EIC_Current_State.md`. Treat only the document(s) the manifest names as "current" as canonical; treat everything else as evidentiary/historical by default.
3. **Identify authority** for any specific claim by checking its status tag (LOCKED/HUMAN-LOCKED/DERIVED CONSTRAINT/PROPOSED/TBD/etc.) in its source document, then cross-checking that tag's meaning against `Authority_Provenance_Reconciliation_Matrix.md` §1 (Authority Classes). A claim is only Class A/B if it traces to an explicit ruling document (e.g., an HD-EIC-xx item or a dated human-approval note) — not merely to a LOCKED tag inside a consolidation document.
4. **Distinguish current facts from proposals** by treating anything outside the manifest's "current" pointer set as PROPOSED/EVIDENCE/HISTORICAL unless it is itself an explicit human ruling document.
5. **Identify unresolved questions** via `TCIndustries_Open_Questions_and_Gates_Register.md`. Do not answer them; report them as open.
6. **Inspect changes** via `TCIndustries_Change_Log.md` for anything since the audit's reference date.
7. **Consult historical/provenance (archive) material only when the manifest or state-layer documents cite it**, or when the audit task explicitly requires evidentiary detail (e.g., "why was Architecture C retired" requires reading the archived falsification/simulation chain, not just the one-line retirement statement).
8. **Report uncertainty explicitly** whenever a claim's authority class cannot be verified from the supplied corpus alone (e.g., the five-item authority-review queue in §5.2) — state "asserted by document X, independent confirmation not present in corpus" rather than treating the assertion as settled.
9. **Cite evidence** by document filename plus section/ID (e.g., "HD-EIC-05, `EIC_Human_Rulings_Architecture_C_Falsification_2026-08-25.md`"), never by paraphrase alone, so a second LLM can locate the same source.
10. **Produce a standardized audit result** using the metadata fields in §11 as the reporting schema, explicitly separating (a) current canonical state, (b) open questions, (c) authority risks/inconsistencies found, and (d) evidence citations.

---

## 11. Recommended Standard Metadata

A lightweight front-matter block (not a new bureaucracy — six fields, optional extras) for every **active-layer** document going forward. Archive documents are not retroactively edited.

```yaml
doc_id: EIC-RULE-05-08          # short stable identifier, independent of filename
title: EIC Human Rulings — Architecture C Falsification
version: 1.0
date: 2026-08-25
status: current                 # current | superseded | archived | proposed
authority_class: A              # per Reconciliation Matrix §1 (A–G)
supersedes: []                  # doc_id(s) this replaces, if any
superseded_by: null             # doc_id, once replaced
controlling_documents: [EIC-GATE-04]   # doc_id(s) this depends on
human_approved: true            # explicit boolean, not inferred from prose
open_questions: []              # OQ/HD IDs this leaves unresolved
```

This schema reuses fields already implicitly present in every document's existing "Document Control" prose block — it does not ask authors to track anything new, only to make explicit as structured fields what is currently only stated in prose (this directly closes the Temporal Reconstruction gap in §2.C and the naming-collision risk in §4.2, since `doc_id` + `version` would immediately disambiguate the two "Comparative Simulation Specification" documents).

---

## 12. Migration / Consolidation Plan

**Staged, proposal-only — no files are to be created, moved, renamed, merged, or deleted by this document.**

**Stage 0 (human decision):** Approve or reject the state-layer concept in principle (§7–§8). No content changes yet.

**Stage 1 (create state layer, zero risk to existing corpus):**
- Draft `PROJECT_STATE_MANIFEST.md`, `TCIndustries_EIC_Current_State.md`, `TCIndustries_Open_Questions_and_Gates_Register.md`, `TCIndustries_Change_Log.md`.
- Every claim in these four documents must cite an existing document; no new decisions are made.
- Existing 19 documents are untouched.

**Stage 2 (light-touch tagging of existing documents):**
- Add the six-field metadata block (§11) to the *current* documents only (the GDD's current version, the Authority Matrix, and the two EIC ruling documents) — four files. This is additive (a front-matter block), not a rewrite.
- Archive documents are left completely unedited.

**Stage 3 (naming-collision remediation, optional, human-authorized):**
- Consider a filename clarification for the two "Comparative Simulation Specification" documents and the two "Human Rulings" documents (e.g., appending the ruling-ID range or package scope to the filename) to reduce the §4.2 confusability risk. This would be a rename only — content unchanged — and still requires explicit human sign-off since it touches existing, already-referenced filenames that other documents' prose cites by name.

**Stage 4 (ongoing discipline):**
- Every future HD-EIC ruling or GDD status patch triggers an update to the manifest and change log as part of the same authoring step that produces the ruling — not a separate cleanup pass.
- Every future EIC "investigate → candidate → falsify → specify → simulate → results" cycle continues to produce its own step-by-step documents (this pattern should **not** be consolidated — see §4.4), but each cycle's *terminal* ruling document must update `TCIndustries_EIC_Current_State.md` as part of closing the cycle.

---

## 13. Risks and Trade-offs

- **Risk: the state layer itself becomes stale.** A manifest is only as good as its update discipline. Mitigation: tie manifest updates to the same authoring step that produces a ruling (Stage 4), not a separate maintenance task that can be skipped.
- **Risk: state-layer documents are mistaken for authority.** Because §8's new documents are explicitly designed to be read first, there is a risk an LLM treats them as the decision-making authority rather than as pointers. Mitigation: every state-layer document must carry `authority_class: none / pointer-only` and cite its source for every claim (already specified in §8.1, §8.4–8.6).
- **Trade-off: four new documents vs. reduced reconstruction cost.** This adds documents, nominally working against "reduce sprawl." The trade-off is deliberate: four small, high-value, always-current pointer documents in exchange for eliminating the need to read 13 documents in the correct order just to answer "is Architecture C still viable." This is judged to satisfy the brief's explicit priority ("cross-LLM auditability... more important than achieving the smallest possible number of files").
- **Risk: renaming existing files (Stage 3) breaks internal cross-references.** Many documents cite each other by exact filename in prose ("Controlling References: `TCIndustries_EIC_Candidate_v0.1.md`"). Any rename requires updating every citing document's prose, which is a larger and riskier operation than Stages 1–2. This is why Stage 3 is marked optional and lowest priority.
- **Risk: the "authority-review queue" (§5.2) cannot be resolved by documentation architecture alone.** Whether WRLD-001/PLR-001/PLR-003/CRFT-001/SAFE-001 were genuinely human-confirmed is a factual question about project history that this audit cannot answer from the corpus alone (see §16). The state layer can only surface the uncertainty, not resolve it.

---

## 14. Human Authorization Gates

(Restated and consolidated in §17 as the final required section — see below. Listed here as they arise structurally.)

- Approval of the state-layer concept (Stage 0).
- Approval to create the four new documents (Stage 1).
- Approval to add metadata front-matter to the four current documents (Stage 2).
- Approval of any filename change to existing documents (Stage 3) — **higher scrutiny required**, since this is the only stage that touches existing, already-cross-referenced files.
- Independent confirmation (or correction) of the five "authority-review queue" GDD items' human-approval status (§5.2), since this audit cannot resolve it from the corpus alone.

---

## 15. Target State

Six months from now, under this architecture, a new independent LLM should be able to:
1. Read `PROJECT_STATE_MANIFEST.md` (≈1 page) and correctly state the current GDD version and current EIC status without reading any other document.
2. Read `TCIndustries_EIC_Current_State.md` (≈1–2 pages) and correctly state whether any EIC architecture is currently authorized, and what the next-required human action is, citing the exact ruling ID.
3. Read `TCIndustries_Open_Questions_and_Gates_Register.md` and correctly enumerate every currently unresolved decision without missing any (currently, these are scattered across at least four documents: GDD §31, item 19 §6, item 11 §5–6, item 16 §1).
4. Read `TCIndustries_Change_Log.md` and correctly state what has changed since any given prior date.
5. Only then, if the task requires it, descend into the archive to verify *why* a given state holds — and find that the archive's existing lineage/provenance discipline (already strong) makes that descent reliable once it is correctly entered.

Two independent LLMs following this protocol should converge on the same answers to "what is the current state" and "what is still open," even if they differ in prose style — satisfying the brief's Test 1 and Test 2. Test 3 (historical containment) is already satisfied by the existing corpus and is preserved unchanged by this proposal, since no archive document is edited.

---

## 16. Open Questions / Information Not Established by the Corpus

- **Whether the five "authority-review queue" GDD items (WRLD-001, PLR-001, PLR-003, CRFT-001, SAFE-001) were actually independently human-confirmed**, as `Master_GDD_v1.1.1_Status_Patch.md` asserts, or whether that assertion is itself an unverified carry-forward. The corpus does not contain a dedicated confirmation record for these five items analogous to the HD-EIC ruling documents. **Not established by the corpus; flagged, not resolved, per instructions.**
- **Whether a structural-boundary statement (required by HD-EIC-07 before any new EIC architecture work) has since been supplied.** The most recent EIC-track document (`EIC_Human_Rulings_Architecture_C_Falsification_2026-08-25.md`) explicitly states this has not yet occurred as of 2026-08-25 and that "the programme remains at the post-falsification gate" until it does. Nothing in the supplied corpus post-dates this. **Current as of the supplied corpus; may be outdated by the time this is read.**
- **Whether any file-system, repository, or tooling constraints exist** (e.g., a wiki, a version-control system, a specific document-management tool already in use) that would affect how the recommended manifest/state-layer documents should actually be implemented. **Not established by the corpus** — this assessment is format-agnostic by necessity.
- **Whether the project owner intends the EIC per-cycle document pattern (§4.3/§4.4) to continue at its current granularity for future systems** (resources, crafting, manufacturing, cities, combat, etc.), or whether a lighter-weight process is intended for less contentious systems. **Not established by the corpus; a scope decision for the project owner, not this audit.**
- **The actual mechanism by which "human ruling" documents are produced** (i.e., what evidence, outside this text corpus, records that a human — as opposed to an LLM summarizing an intended human position — issued HD-EIC-01–08). This audit takes the documents' self-description ("Explicit human project rulings") at face value, consistent with the instruction not to invent assumptions, but notes it cannot independently verify authorship from the corpus alone.

---

## Proposed Active Documentation Set (Summary Table)

| Document | Purpose | Authority | Active/Archive | Primary Consumer |
|---|---|---|---|---|
| `PROJECT_STATE_MANIFEST.md` *(new)* | Single entry point; pointers to current GDD, current EIC state, HUMAN-LOCKED ruling list | None (pointer only) | Active | Any LLM, read first |
| `Master_GDD_v1.1.1_Status_Patch.md` *(existing)* | Canonical design reference | A/B (per item, with §5.2 caveat) | Active | Design-task LLMs |
| `Authority_Provenance_Reconciliation_Matrix.md` *(existing)* | Authority-class vocabulary and governance rules | D (governance) | Active | Any LLM adjudicating status/authority |
| `TCIndustries_EIC_Current_State.md` *(new)* | Current EIC architecture status and next-authorized-operation | None (cites HD-EIC rulings) | Active | EIC-task or audit LLMs |
| `TCIndustries_Open_Questions_and_Gates_Register.md` *(new)* | Consolidated open-question/gate index | None (index only) | Active | Project owner; audit LLMs |
| `TCIndustries_Change_Log.md` *(new)* | Chronological state-change record | None (log only) | Active | Audit LLMs, "what's new" queries |
| `EIC_Human_Rulings_2026-08-25.md` *(existing)* | HD-EIC-01–04 | **A — human ruling** | Active authority / archived form | Cited by state layer |
| `EIC_Human_Rulings_Architecture_C_Falsification_2026-08-25.md` *(existing)* | HD-EIC-05–08 | **A — human ruling** | Active authority / archived form | Cited by state layer |
| All remaining 15 documents (investigation, candidate, falsification, specs, results, refinements, superseded GDD versions, audit) | Evidentiary/provenance record of how current state was reached | Evidence / historical / superseded | Archive | Deep-audit or "why" queries only |

---

## Human Authorization Required

The following require explicit project-owner approval before any implementation occurs. **Nothing in this assessment has been implemented; no file has been created, moved, renamed, merged, or deleted.**

1. **Approval to create four new state-layer documents**: `PROJECT_STATE_MANIFEST.md`, `TCIndustries_EIC_Current_State.md`, `TCIndustries_Open_Questions_and_Gates_Register.md`, `TCIndustries_Change_Log.md` (§8.1, §8.4–8.6, Stage 1).
2. **Approval to add a metadata front-matter block** (§11) to the four current active documents (GDD current version, Authority Matrix, both EIC ruling documents) (Stage 2).
3. **Approval (or rejection) of any filename disambiguation** for the two "Comparative Simulation Specification" documents and the two "Human Rulings" documents (§4.2, Stage 3) — explicitly flagged as higher-scrutiny since it touches existing, cross-referenced filenames.
4. **Independent confirmation or correction of the human-approval status** of the five "authority-review queue" GDD items (WRLD-001, PLR-001, PLR-003, CRFT-001, SAFE-001) — this audit can surface the gap (§5.2) but cannot resolve it.
5. **Confirmation of the corrected understanding that EIC architectural work is currently blocked** pending the HD-EIC-07 structural-boundary statement, so the GDD's §34 "next design phase" pointer (§5.1) can be either updated or explicitly annotated as conditional, at the project owner's discretion — this audit does not recommend which wording change to make, only that the current inconsistency be addressed.

No other action is proposed or should be inferred from this document.

*End of Assessment.*
