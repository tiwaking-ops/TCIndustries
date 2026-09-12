---
doc_id: MANIFEST
title: TCIndustries — Project State Manifest
version: "1.3"
date: 2026-08-25
status: active
authority_class: null    # this document is a pointer/index; it holds no authority of its own
human_approved: false    # the manifest itself was not separately human-ruled; its content is derived from documents that were
supersedes: []
superseded_by: null
controlling_documents: []
reading_priority: mandatory
archived_author_note: "Author LLM Unknown (filed 2026-09-12 as unadjudicated proposal; see proposals/README.md)"
change_note: "v1.3 — Added GR-005 (Filename Convention Compliance). Corrected/retracted the earlier 'naming collision' finding (C7 in the consolidation plan) — it was a missing-prefix artifact, not a real collision. Restored 4 documents to their correct canonical prefixed filenames. v1.2 — Added GR-004 (Ruling Context Disclosure), proposed for merge into the Authority Matrix and in effect now. Replaced the resolved GDD-AUTH-REVIEW-GATE warning with a permanent Process Risk Note. v1.1 — Added GDD-AUTH-REVIEW-GATE (§1a, since resolved)."
---

# TCIndustries — Project State Manifest

**Read this document first, always, before any other TCIndustries document.**

This manifest makes no design decisions and holds no authority of its own. Every claim below cites the document that actually establishes it. If this manifest and a cited source ever disagree, **the source document governs, not this manifest** — and that disagreement should be reported as a bug in this manifest.

---

## 1. Current Canonical Design Reference

**File:** `TCIndustries_Master_GDD_v1.1.1_Status_Patch.md`
**Status:** Working Canonical Reference (human authority rulings applied, 2026-08-25).
**Do not treat any other GDD file as current.** Three GDD bodies exist in archive (`v1.0_Consolidated`, `v1.0_Canonical_Audit`, `v1.1_Canonical_Baseline`) — all are superseded. `v1.1.1` explicitly names `v1.1` as its immediate predecessor.
**Note:** The GDD's own §34 states *"No further LLM consolidation of the Master GDD should occur."* Do not regenerate, rewrite, or produce a "v1.2."

### Process Risk Note (Confirmed, Not Hypothetical)

During this consolidation effort, one of the GDD's HUMAN-LOCKED tags (WRLD-001/PLR-001/PLR-003/CRFT-001/SAFE-001) was traced back to a real decision that had, in fact, been made — but made in a **ChatGPT session that did not have access to the full project file set.** The decision's *outcome* made it into the canonical GDD; the decision's *process and rationale* did not, until recovered and formally recorded here as `HD-GDD-01` (see `Human_Rulings_Register.md`, Section B). **This is the confirmed root cause of at least one instance of documentation drift in this project**, and is the concrete reason the reading-order discipline in this manifest matters: any LLM session (this one included) that operates without the full corpus can produce a real, human-endorsed decision whose paper trail then silently breaks.

## 2. Current Governance / Authority Rules

**File:** `TCIndustries_Authority_Provenance_Reconciliation_Matrix.md`
Defines Authority Classes A–G and governance rules GR-001–003 (most important: **GR-001 — canonical presence is not human approval**). Read this whenever adjudicating whether a statement elsewhere is genuinely LOCKED.

### GR-004 — Ruling Context Disclosure *(new — proposed for merge into the Authority Matrix; in effect now, project-owner confirmed 2026-08-25)*

**Status: GOVERNANCE RULE.**

Any document recording a human ruling (`HD-EIC-xx`, `HD-GDD-xx`, or any future prefix) must state which project documents the producing LLM session had loaded or available at the time the ruling was elicited and recorded — at minimum, whether the **full current corpus** (per the then-current Manifest) was available, or a **named subset / older snapshot**.

A ruling made without full corpus context is **not thereby invalid**. This rule does not retroactively void anything. It exists so a future reader can independently assess whether a ruling was made with awareness of all then-current constraints, other open questions, and potentially conflicting prior rulings — rather than discovering the gap by accident, as happened with `HD-GDD-01` (see `TCIndustries_Human_Rulings_Register.md`, Section B, and the Process Risk Note below).

**Absence of this disclosure in an existing ruling record is itself a documentation defect** to be corrected via a provenance note (the same pattern used for `HD-GDD-01`), not grounds to treat the ruling as void.

**Implementation:** the `ruling_context` field is now part of the standard metadata schema (see `Changelog.md` for the field addition) and has been retrofitted, where determinable, onto the existing `HD-EIC-01–08` and `HD-GDD-01` entries in `Human_Rulings_Register.md`.

**Human authorization required:** GR-004's text above is written ready to paste directly into `Authority_Provenance_Reconciliation_Matrix.md` §2, alongside GR-001–003, as a formal fourth governance rule. Until that file is actually edited, this manifest is GR-004's operative home.

### GR-005 — Filename Convention Compliance *(new — proposed for merge into the Authority Matrix; in effect now, project-owner confirmed 2026-08-25)*

**Status: GOVERNANCE RULE.**

All TCIndustries project documents must be saved with the `TCIndustries_` prefix, without exception.

Any process that gathers or merges project documents by filename pattern must explicitly report every candidate file it excludes for not matching the pattern, rather than silently omitting it. A file found without the prefix is a documentation defect to be corrected by renaming, not by treating its absence from a merge as evidence it doesn't exist.

**Why this rule exists — confirmed impact, not a hypothetical:** During this consolidation effort, 4 documents were found to exist under two names each — one correctly prefixed, one not:

| Correctly prefixed (canonical) | Found also existing without prefix |
|---|---|
| `TCIndustries_EIC_Human_Validation_Gate_Architecture_C_Falsification_v0.1.md` | `EIC_Human_Validation_Gate_Architecture_C_Falsification_v0.1.md` |
| `TCIndustries_EIC_Comparative_Simulation_Specification_v0.1.1.md` | `EIC_Comparative_Simulation_Specification_v0.1.1.md` |
| `TCIndustries_EIC_Comparative_Simulation_Results_v0.1.md` | `EIC_Comparative_Simulation_Results_v0.1.md` |
| `TCIndustries_EIC_Comparative_Simulation_Results_v0.1.1.md` | `EIC_Comparative_Simulation_Results_v0.1.1.md` |

The `TCIndustries-project-files-v1-merged.txt` upload used in this consolidation was almost certainly built by a prefix-matching gather process. It silently dropped the unprefixed Validation Gate document entirely (confirmed absent from that merge) and — because the α/β/γ Specification and both Results documents happened to have prefixed *and* unprefixed copies with subtly different internal headers — this same mechanism is the most likely explanation for why 4 of 8 independent LLM audits of this project (including this one, initially) misread the corpus and reported a false "naming collision" between the Test A/B and α/β/γ specifications. **That collision finding is retracted — see the Erratum in `TCIndustries_Consolidated_Consensus_Plan_v1.md`.** The two specification files, and the two results files, are genuinely distinct, non-colliding, correctly-named documents once the prefix convention is respected.

**Resolution applied:** all 4 documents restored to their canonical prefixed filenames in the archive built for this consolidation.

## 3. Current EIC (Economic Interdependence Core) State

**File:** `TCIndustries_EIC_Program_State.md`
**Headline fact:** Architecture C is **RETIRED** (falsified). No EIC architecture is currently the leading hypothesis. No new EIC architectural work is authorised until the human supplies the structural-boundary statement required by HD-EIC-07.
Full ruling detail: `TCIndustries_Human_Rulings_Register.md`.

## 4. Every Currently HUMAN-LOCKED Ruling

**File:** `TCIndustries_Human_Rulings_Register.md`
Contains HD-EIC-01 through HD-EIC-08 in full, chronologically, merged from their two original source documents (both preserved verbatim in archive).

## 5. Open Questions

**File:** `TCIndustries_Open_Questions_Register.md`
The single outstanding gate blocking further EIC work: **the HD-EIC-07 structural-boundary statement has not yet been supplied.** All GDD-level open questions (OQ-001–017) are also indexed there.

## 6. What Changed, and When

**File:** `TCIndustries_Changelog.md`

---

## Reading Order for a Standard Audit

1. This manifest.
2. `Authority_Provenance_Reconciliation_Matrix.md` (rules for interpreting status tags — skim, don't re-derive).
3. `Human_Rulings_Register.md` (what is actually decided).
4. `Master_GDD_v1.1.1_Status_Patch.md` (canonical design content — full read only if the task needs design detail; otherwise its Document Control + §3–4 + §31–36 suffice).
5. `EIC_Program_State.md` (current EIC status in one page).
6. `Open_Questions_Register.md` (what remains unresolved).
7. `Changelog.md`, if the task is "what changed since—".
8. **Only then**, descend into `/archive/` for evidentiary detail (e.g., "why was Architecture C retired" requires the archived falsification/simulation chain; "what did earlier GDD versions say" requires the archived snapshots).

## Authority Hierarchy (Summary)

```
Explicit human ruling (HD-EIC-xx, or a dated "HUMAN-LOCKED (confirmed ...)" GDD tag)
        │  highest authority
        ▼
Project principle / pillar (VIS-, PIL- items, explicitly attributed)
        │
        ▼
Derived constraint (logical consequence of a locked principle — never itself an independent decision)
        │
        ▼
Proposed / TBD (recommendation or open question — no authority)
        │
        ▼
Evidence / Prototype (informs design; grants no authority)
        │
        ▼
Historical (SWG or superseded material — inspiration/record only, never automatically current)
```

## Last Updated

2026-08-25 — GR-004 (Ruling Context Disclosure) adopted; retrofitted onto `Human_Rulings_Register.md`. Prior update: issuance of HD-EIC-05–08 (Architecture C retirement) and the `v1.1.1` GDD status patch. This manifest should be updated at the same time any future HD-EIC/HD-GDD ruling, GDD status patch, or governance rule is issued — not as a separate cleanup pass.

## Standard Ruling Metadata (per GR-004)

Every future `HD-*` ruling record must include, in addition to the schema from the original consolidation plan:

```yaml
ruling_context:
  full_corpus_available: true | false | unknown
  documents_available: []   # list, or "see Manifest as of <date>" if full corpus
  documents_missing: []     # named if known and partial
```

## Standard Document Save Convention (per GR-005)

Every project document, without exception: filename must begin with `TCIndustries_`. Any tool or process that gathers documents by pattern must log excluded non-matching files explicitly rather than silently dropping them.

*End of Manifest.*
