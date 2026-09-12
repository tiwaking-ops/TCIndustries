```markdown
# TCIndustries — GitHub Repository Architecture & Import Instructions
## Structured Archive for Multi-Version Design & Prototype Consolidation

**Document Type:** Formal Project Documentation Report (OpenCode / Repository Architecture)  
**Version:** 0.1  
**Date:** 2026-09-12  
**Status:** PROPOSED — Awaiting human review and execution  
**Authority:** Non-canonical. This document provides repository structure and operational instructions only. It does not alter any design status, promote any mechanism, or create new game design content.  
**Author:** Grok (xAI) — Grok 4.5  

**Document Purpose:**  
Define a durable, public GitHub repository architecture and the exact steps required to import and organise all existing TCIndustries design documents and competing multi-language prototype versions into a single structured archive. The repository must preserve exact decision status, testing history, and human-monitored decision process so that both the human owner and future LLMs can answer “What is the current status of decision X and what evidence led to it?” without tribal knowledge.

**Authority Statement:**  
This report is advisory. It does not supersede the Master GDD, any Human Rulings, or the Authority & Provenance Reconciliation Matrix. All design status tags (LOCKED, PROPOSED, TBD, DERIVED CONSTRAINT, etc.) remain under exclusive human control.

**Lineage:**  
- Master GDD v1.1.1 Status Patch  
- Authority & Provenance Reconciliation Matrix  
- EIC Human Rulings (HD-EIC-01 through HD-EIC-08)  
- All EIC Candidate, Falsification, Simulation, and Functional Shapes documents present in the working artifacts set  
- Project instructions governing status discipline and non-silent promotion of design material

---

# 1. Document Control

| Field              | Value                                                                 |
|--------------------|-----------------------------------------------------------------------|
| Project            | TCIndustries (Tiwakings Craftworld Industries)                        |
| Document           | GitHub Repository Architecture & Import Instructions                  |
| Version            | 0.1                                                                   |
| Status             | PROPOSED                                                              |
| Scope              | Repository layout, branch/tag strategy, import procedure, ongoing conventions |
| Change Control     | Structural changes require human approval before execution            |
| Author             | Grok (xAI) — Grok 4.5                                                 |
| Audience           | Human owner + LLMs (public repository)                                |
| Controlling Pillars| Exact status visibility, recoverable history, human-monitored decisions |

---

# 2. Recommended Repository Architecture

## 2.1 Top-Level Layout (Monorepo Structured Archive)

```
tcindustries/
├── README.md                          # Entry point: current baseline, how to read status, decision process
├── LICENSE                            # Choose and declare (recommended: proprietary or CC-BY-NC for design docs)
├── .gitignore
│
├── docs/                              # Design authority layer
│   ├── current/                       # Single source of truth for active design
│   │   ├── Master_GDD.md              # Current working canonical (symlink or copy of latest Status Patch)
│   │   ├── Authority_Matrix.md
│   │   └── Human_Rulings/             # All HD-EIC and future human authority records
│   │
│   ├── archive/                       # Immutable historical versions
│   │   ├── gdd/
│   │   │   ├── v1.0_Consolidated/
│   │   │   ├── v1.0_Canonical_Audit/
│   │   │   ├── v1.1_Canonical_Baseline/
│   │   │   └── v1.1.1_Status_Patch/
│   │   ├── eic/
│   │   │   ├── Candidate_v0.1/
│   │   │   ├── Falsification_Pass/
│   │   │   ├── Comparative_Simulation_*/
│   │   │   ├── Functional_Shapes_*/
│   │   │   └── Human_Rulings_*/
│   │   └── investigations/
│   │
│   └── process/                       # How decisions are made and tested
│       ├── Decision_Log.md            # Chronological human-monitored decisions
│       ├── Test_Evidence_Index.md     # Pointers to all simulation/falsification results
│       └── Status_Discipline.md       # Explicit restatement of LOCKED / PROPOSED / etc. rules
│
├── prototypes/                        # Competing implementations — never overwritten
│   ├── _index.md                      # Table of all prototype versions, language, status, relation to design
│   ├── lang-a/                        # e.g. TypeScript / Node prototype set
│   ├── lang-b/                        # e.g. Python / simulation harness
│   ├── lang-c/                        # additional languages as they exist
│   └── experiments/                   # Short-lived or discarded experiments
│
├── evidence/                          # Raw and summarised test results
│   ├── simulations/
│   ├── falsification-passes/
│   └── adversarial-retests/
│
└── meta/                              # Repository governance
    ├── CONTRIBUTING.md                # How LLMs and future humans should interact with the repo
    ├── STATUS.md                      # Machine-readable snapshot of current design baseline
    └── import-log.md                  # Record of every bulk import action
```

## 2.2 Branch & Tag Strategy

| Branch / Tag Pattern          | Purpose                                      | Lifetime          |
|-------------------------------|----------------------------------------------|-------------------|
| `main`                        | Current baseline only (docs/current + latest evidence) | Permanent        |
| `archive/gdd-vX.Y`            | Frozen historical design documents           | Permanent tags   |
| `archive/prototype-<lang>-vN` | Frozen language-specific prototype snapshots | Permanent tags   |
| `work/<topic>`                | Temporary human or LLM work branches         | Short-lived      |
| `evidence/<test-id>`          | Isolated simulation or falsification runs    | Merge then delete|

Rules:
- `main` never contains superseded design documents in active paths.
- Every major design version and every prototype language version receives an annotated Git tag.
- No force-push to `main` or any `archive/*` tag after the initial import is complete.

## 2.3 Multi-Language Prototype Handling

Each language-specific prototype lives in its own top-level folder under `prototypes/`.  
The folder name is the language or runtime (e.g. `typescript`, `python`, `csharp`).  
Inside each language folder:

```
prototypes/<lang>/
├── README.md                 # What this prototype implements, which design version it targets, known deviations
├── src/                      # Source as-is
├── docs/                     # Any notes that accompanied the prototype
└── STATUS.md                 # Explicit mapping: which GDD rules / EIC mechanisms this code claims to exercise
```

No attempt is made to unify languages. Side-by-side archival is the explicit design.

---

# 3. Step-by-Step Import Instructions

## 3.1 Create the Repository

1. Create a new public GitHub repository named `tcindustries` (or preferred public name).
2. Initialise with a README only. Do not add a .gitignore or license yet via the web UI.
3. Clone locally.

## 3.2 Establish Skeleton

```bash
mkdir -p docs/{current,archive/{gdd,eic,investigations},process}
mkdir -p prototypes evidence/{simulations,falsification-passes,adversarial-retests} meta
touch docs/process/{Decision_Log.md,Test_Evidence_Index.md,Status_Discipline.md}
touch prototypes/_index.md meta/{CONTRIBUTING.md,STATUS.md,import-log.md}
```

Commit the empty skeleton with message:  
`chore: establish structured archive skeleton`

## 3.3 Import Design Documents (Ordered)

Import in chronological / authority order so Git history itself becomes useful:

1. Place every existing Master GDD version into the corresponding `docs/archive/gdd/` subfolder.
2. Place every EIC document into the matching `docs/archive/eic/` subfolder.
3. Copy the current authoritative set into `docs/current/`:
   - Latest Master GDD Status Patch → `Master_GDD.md`
   - Authority Matrix
   - All Human Rulings documents
4. After each logical group, commit with a clear message, e.g.:  
   `docs: archive Master GDD v1.0 Consolidated + Canonical Audit`  
   `docs: archive full EIC Candidate → Falsification → Simulation sequence`  
   `docs: establish current baseline from v1.1.1 Status Patch + HD-EIC-05–08`

## 3.4 Import Competing Prototypes

For each language-specific prototype set:

1. Create `prototypes/<lang>/`.
2. Copy the entire working tree of that language version into the folder.
3. Write a short `STATUS.md` that states:
   - Design version it was written against
   - Which mechanisms it attempted to implement
   - Known deviations from the design documents of that era
4. Commit per language:  
   `prototypes: archive <lang> prototype set (pre-Architecture-C era)`

Update `prototypes/_index.md` with a table linking language, date range, design target, and Git tag.

## 3.5 Tag the Archive

After the complete import:

```bash
git tag -a archive/gdd-v1.0 -m "Master GDD v1.0 Consolidated + Audit"
git tag -a archive/gdd-v1.1.1 -m "Status Patch + Human Authority Rulings"
git tag -a archive/eic-architecture-c-falsified -m "Complete EIC Architecture C sequence through HD-EIC-08"
# one tag per language prototype set
git push origin main --tags
```

## 3.6 Finalise Entry Points

- Write `README.md` that points first to `docs/current/` and second to the Decision Log and Test Evidence Index.
- Populate `docs/process/Status_Discipline.md` with the exact status vocabulary from the Master GDD.
- Record the entire import sequence in `meta/import-log.md`.

---

# 4. Ongoing Conventions (Human-Monitored Decision Process)

1. **No silent status promotion**  
   Any change that would move a rule from PROPOSED → LOCKED (or equivalent) must be recorded as a new Human Ruling document under `docs/current/Human_Rulings/` and referenced in the Decision Log.

2. **Evidence before claim**  
   Every simulation, falsification pass, or adversarial retest is committed under `evidence/` and indexed in `Test_Evidence_Index.md` before any design document is allowed to cite it as support.

3. **LLM interaction rule**  
   LLMs may propose branches under `work/` and may update PROPOSED or TBD material only. They may never edit files under `docs/current/` or any `archive/` path without an explicit human commit that accepts the change.

4. **Decision Log format**  
   Each entry: Date | Ruling ID | Decision summary | Controlling documents | Status impact | Human confirmation marker.

5. **Prototype status**  
   A prototype never becomes design authority. Its `STATUS.md` must always declare the design version it targets and list deviations.

---

# 5. Success Criteria Check

| Criterion                                                                 | How the architecture satisfies it                                      |
|---------------------------------------------------------------------------|------------------------------------------------------------------------|
| Exact status of any decision + its history recoverable in < 2 minutes    | `docs/current/` + Decision Log + tagged archive versions               |
| Every competing version (docs + languages) remains recoverable            | Immutable `archive/` folders + permanent tags + per-language prototype trees |
| Layout itself encodes status discipline and human-gate process            | Explicit `Status_Discipline.md`, Human_Rulings folder, LLM contribution rules in CONTRIBUTING.md |

---

# 6. Explicit Non-Goals

- This document does not invent new design content.
- This document does not authorise numeric tuning, Provenance engine work, or a new EIC architecture.
- This document does not collapse multi-language prototypes into a single implementation.
- This document does not change the status of any existing rule or ruling.

---

**End of Report**

**Next authorised human action:** Review this architecture. If accepted, execute the import sequence and mark this document HISTORICAL or SUPERSEDED by a short “Repository Established” note under `meta/`.
```