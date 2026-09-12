I'll start by mapping the full merged corpus—document boundaries, titles, and versions—before assessing authority, redundancy, or recommending any architecture.The corpus is a long chain of EIC investigation, candidate, and simulation artifacts. I'll continue mapping later pages so I can inventory every source document.The later pages contain human rulings and simulation results. I'll keep inventorying remaining documents so the architecture recommendation is based on the full set.The EIC chain is now mapped through Architecture C’s retirement. Next I’ll locate the Master GDD, audits, and remaining source documents.The Master GDD and its audit are now in view. I'll continue through the remaining pages to complete the document inventory.The GDD lineage is clearer now. I’ll finish the remaining pages and then search for any documents that still aren’t accounted for.Several referenced first-pass GDDs and later EIC artifacts may be missing from the merge. I'll search titles and filenames to confirm what's actually present.# TCIndustries Documentation Consolidation and Cross-LLM Auditability Assessment

**Auditor role:** Independent documentation architect / evidence–governance auditor  
**Scope:** The supplied merged corpus only  
**Date of assessment:** based solely on documents dated 2026-08-25 in the merge  
**Authority of this assessment:** None. This is a documentation-architecture proposal. It creates no game-design decisions, does not promote or demote design status, and authorises no file operations.
**Author:** Grok 4.6 xhigh (per filename, unverified)

**Evidence rule used below:**  
- **Direct** = a named document states it.  
- **Inference** = my reconstruction from multiple statements.  
- **Not established** = the merged corpus does not contain the fact.

---

## 1. Executive Verdict

The project already has unusually strong **status vocabulary and authority hygiene**. GR-001–GR-003, HUMAN-LOCKED vs DERIVED CONSTRAINT vs PROPOSED, the non-canonical boundary, and the EIC “creates no design authority” discipline are real assets. They must be preserved.

The corpus **fails the three success tests today**.

| Test | Result | Why |
|---|---|---|
| **1 — State reconstruction** | **FAIL** | There is no single current-state pointer. Three near-complete Master GDD bodies sit in one merge. EIC documents assert contradictory “current” programme states. Almost every artifact is dated 2026-08-25. |
| **2 — Audit reproducibility** | **FAIL** | Two independent LLMs can honestly pick different “current” documents (v1.1 vs v1.1.1; Investigation vs HD-EIC-05) and produce non-comparable audits. |
| **3 — Historical containment** | **PARTIAL** | Governance *rules* are good. Physical containment is weak: superseded GDD text, retired Architecture C, and historical proposals remain in the live reading stream without a machine-obvious “not current” gate. |

**Root cause (inference):** the project solved *authority language* before it solved *document identity, succession, and intake order*. New work is being recorded as a new full document per epistemic step. That is scientifically honest and operationally unsustainable.

**What is actually current (direct, if an auditor follows the latest human-authority records rather than merge order):**

- Working canonical design reference: **Master GDD v1.1.1 Status Patch**.
- Governance rules: **Authority & Provenance Reconciliation Matrix** (GR-001, GR-002, GR-003).
- EIC programme: **Architecture C RETIRED** (HD-EIC-05). PIL-003 unchanged (HD-EIC-06). Next authorised EIC act is **human structural boundaries** (HD-EIC-07). No Architecture D, no GDD rewrite, no numeric tuning (HD-EIC-08).
- Prototype: **unreconciled**. Seed-2.1 is referenced, not inspected (PROT-002 / AS-001).
- Canonical EIC mechanics: **not established**.

**Recommended direction:** a **small active set** (about seven living documents) plus a **strict archive**, keyed by a **project-state manifest**. Do **not** collapse everything into one mega-document. Do **not** rewrite the Master GDD. Do **not** delete history. All of this requires later human authorisation (and HD-EIC-08 already forbids a GDD rewrite).

---

## 2. Current Documentation Assessment

### What the system does well (direct)

- Explicit status tags on design statements (LOCKED, DERIVED CONSTRAINT, PROPOSED, TBD, PROTOTYPE/EVIDENCE, DEFERRED, ASSUMPTION, HISTORICAL).
- Explicit rule that **canonical presence ≠ human approval** (GR-001).
- Explicit rule that **historical proposals do not re-enter canon** (GR-002; Matrix §5 list).
- Explicit rule that **derived ≠ independently decided** (GR-003).
- EIC documents repeatedly refuse to promote themselves.
- Human rulings isolate *what was decided* from *what remains open*.
- Open-question and assumption registers exist inside the GDD.
- Non-canonical mechanical list is concrete and repeated in v1.1.1.

### What the system fails (direct + inference)

| Criterion | Verdict | Evidence |
|---|---|---|
| **A. State reconstruction** | Fail | No manifest. Merge contains v1.0, v1.1, and v1.1.1 GDD bodies. EIC “programme state” tables disagree across documents. |
| **B. Authority reconstruction** | Partial | Vocabulary is clear; *location* of the latest human ruling is not. An LLM that stops at the Investigation or at GDD v1.1 will mis-state CRFT-006, PROG-002, and Architecture C. |
| **C. Temporal reconstruction** | Fail | Nearly every document is dated 2026-08-25. Two different specs share version **v0.1**. v1.1 appears *after* v1.1.1 in the merge. |
| **D. Provenance** | Partial | Rule IDs exist. First-pass GDD sources and the Validation Gate document are **cited but not in the merge**. Prototype not inspectable. |
| **E. Cross-LLM comparability** | Fail | No published reading order, no audit schema, no single status snapshot. |
| **F. Change detection** | Partial | Change notes exist *inside* v1.1 and v1.1.1. No standalone changelog. EIC state changes are scattered. |
| **G. Scope control** | Partial at rule layer; weak at corpus layer | GR-001–003 are strong. Full retired/proposed prose is still in the default intake. |
| **H. Scalability** | Fail if continued | One new long document per investigation/spec/result/gate. Three full GDD copies already. |

### Important inference

The documentation problem is **not** “too little governance prose.” It is **too many complete snapshots with no front door**. Adding another investigation-length document would make auditability worse.

---

## 3. Current Documentation Inventory

Documents **actually present** in the merged PDF, in merge order (not process order).

### 3.1 Present artifacts

| # | Document (as titled in corpus) | Type | Stated authority | Role in lineage | Suggested disposition |
|---|---|---|---|---|---|
| 1 | Authority & Provenance Reconciliation Matrix v1.0 | Governance / provenance | None for game design; adopts GR-001–003 | Defines authority classes; flags review queue; lists non-canon | **Active governance** (or freeze as current governance baseline) |
| 2 | Economic Interdependence Core: Rigorous Design Investigation v1.0 | Design investigation | None; all PROPOSED | Competes A/B/C; recommends C | **Archive / investigation** |
| 3 | EIC-Candidate v0.1 | Structural candidate | None | M1–M8; Architecture C as leading candidate | **Archive / candidate** |
| 4 | EIC-Candidate v0.1 Falsification Pass | Adversarial analysis | None | Not destroyed, not validated | **Archive / audit record** |
| 5 | EIC Comparative Simulation Results v0.1.1 | Qualitative evidence | None | α FAIL, β FAIL, γ INCONCLUSIVE | **Archive / evidence** |
| 6 | EIC Comparative Simulation Execution & Results v0.1 | Adversarial evidence | None | Same packages; γ INCONCLUSIVE leans FAIL | **Archive / evidence** (near-duplicate of 5) |
| 7 | EIC Comparative Simulation Specification v0.1 *(post-Functional-Shapes)* | Experiment spec | None | Packages α/β/γ; 7-phase sequence; says it supersedes earlier spec | **Archive / evidence spec** |
| 8 | EIC Comparative Simulation Specification v0.1 *(earlier Test A/B)* | Experiment spec | None | Factory-parity vs formal/soft budget | **Archive / superseded spec** — **version collision** |
| 9 | EIC Functional Shapes Refinement — Targeted γ Closure v0.2 | Provisional simulation assumptions | None | Closes four γ gaps only | **Archive / simulation assumption** |
| 10 | EIC Human Decision Brief v0.1 | Decision brief | None | Recommends Gates 1–4 | **Archive / brief** (not a ruling) |
| 11 | EIC Human Rulings 2026-08-25 (HD-EIC-01–04) | **Human authority record** | Explicit human rulings | Layered unit; comparative value; multi-account in-scope; shapes permission | **Active rulings source** (extract into register) |
| 12 | EIC Human Rulings — Architecture C Falsification Gate (HD-EIC-05–08) | **Human authority record** | Explicit human rulings | C retired; requirement unchanged; bounded shaping only after human boundaries; gov boundaries confirmed | **Active rulings source** |
| 13 | EIC Minimum Comparative Simulation Results v0.1 | Evidence | None | Test A/B INCONCLUSIVE / under-specified | **Archive / evidence** (earlier experiment) |
| 14 | EIC Provisional Functional Shapes v0.1 | Provisional model | None (HD-EIC-04 permission only) | Non-numeric shape catalogue | **Archive / simulation assumption** |
| 15 | EIC Targeted γ′ Adversarial Re-test Results v0.1 | Evidence | None | γ′-1 FAIL, γ′-2 FAIL; C falsified as shaped | **Archive / evidence** (input to HD-EIC-05) |
| 16 | Master GDD v1.0 Canonical Audit | Audit | Audit recommendations only | Reclassify five LOCKED items; produce v1.1 | **Archive / audit record** |
| 17 | Master GDD v1.0 Consolidated | Design reference | Then “primary design authority”; later superseded | First consolidated GDD | **Superseded / archive** |
| 18 | **Master GDD v1.1.1 Status Patch** | Working canonical GDD | Working canonical reference; human rulings applied | Current design text | **Sole active GDD** |
| 19 | Master GDD v1.1 Canonical Baseline | Working canonical (then) | Supersedes v1.0; superseded by v1.1.1 | Status pass before human confirmations | **Superseded / archive** |

**Direct observation:** merge order is not chronological or authoritative order. v1.1.1 appears *before* v1.1. Several EIC results appear before the specs and rulings that supposedly control them.

### 3.2 Cited but **not present** in the merged corpus

| Cited artifact | Cited by | Consequence |
|---|---|---|
| `TCIndustries Master Game Design Document v0.1 - gpt-5.6-terra-medium.md` | GDD v1.0 § sources | First-pass provenance not inspectable |
| `TCIndustries_Master_GDD-chatgptv1-1.md` | GDD v1.0 | Same |
| `TCIndustries_Master_GDD_v1.0-claude-opus-4-6-thiking.md` | GDD v1.0 | Same |
| `TCIndustries_Master_GDD-claude1.md` (said identical to chatgptv1-1) | GDD v1.0 | Same |
| `Tiwakings Craftworld Industries — Master GDD v0.1-chatgpt1.md` | GDD v1.0 | Same |
| `TCIndustries_EIC_Human_Validation_Gate_Architecture_C_Falsification_v0.1.md` | HD-EIC-05–08 controlling priors | Human-gate *brief* not inspectable; rulings themselves are present |
| `TCIndustries_EIC_Comparative_Simulation_Specification_v0.1.1.md` | Results v0.1.1 | Controlling spec for one results file is missing |
| Seed-2.1 TypeScript/Vite/React prototype | PROT-002, AS-001 | Evidence track empty |
| Project-state manifest / changelog | — | **Does not exist** (search found no “manifest”) |
| HD-EIC-07 structural-boundary statement | HD-EIC-07 itself | **Not established**; programme waits on it |

**Not established:** whether those missing files exist outside this merge. For any new LLM given *only this merge*, they are absent.

### 3.3 Which documents are genuinely authoritative

| Class | Documents | Notes |
|---|---|---|
| **Human design/governance authority** | HD-EIC-01–04 record; HD-EIC-05–08 record; human confirmations recorded *inside* GDD v1.1.1 (WRLD-001, PLR-001, PLR-003, CRFT-001, SAFE-001; CRFT-006/PROG-002 demotions) | Authority is the *ruling*, not surrounding prose |
| **Working canonical design text** | **GDD v1.1.1 only** | v1.1 and v1.0 are prior snapshots |
| **Binding governance rules** | Authority Matrix GR-001–003; status vocabulary in v1.1.1 §3–4 | Matrix creates no game design |
| **Evidence only** | All simulation specs/results, falsification pass, γ′ retest, functional shapes | Do not grant design authority (their own statements) |
| **Recommendation only** | Investigation; Candidate v0.1; Human Decision Brief | Must not be treated as rulings |
| **Audit record** | Canonical Audit of v1.0 | Recommendations; several were later applied |

**Do not infer authority from self-description alone.** GDD v1.0 calls itself “current primary design authority.” That is historically true and **currently false** relative to v1.1.1 (direct succession statements in v1.1 and v1.1.1).

---

## 4. Redundancy and Sprawl Analysis

### 4.1 Triple GDD body (highest-cost duplication)

v1.0, v1.1, and v1.1.1 each contain essentially the same ~36-section GDD. Differences that matter are **status tags and a few authority lines**, not new design.

**Direct material differences an auditor must not miss:**

| Item | v1.0 | v1.1 | v1.1.1 |
|---|---|---|---|
| CRFT-006 | LOCKED (derived) | DERIVED CONSTRAINT | **PROPOSED** (human demotion) |
| PROG-002 | LOCKED (derived) | DERIVED CONSTRAINT | **PROPOSED** (human demotion) |
| MFG-004, ECO-003, CMBT-003 | LOCKED (anti-goal/derived) | DERIVED CONSTRAINT | DERIVED CONSTRAINT |
| WRLD-001, PLR-001, PLR-003, CRFT-001, SAFE-001 | LOCKED, often “required by vision/pillars” | LOCKED, still thin authority | **HUMAN-LOCKED confirmed 2026-08-25, principle only** |
| DERIVED CONSTRAINT vocabulary | Absent | Present | Present |
| Document status | Primary Design Reference | Working Canonical Reference | Working Canonical Reference (human rulings applied) |

An LLM that reads v1.1 (last in the merge) will **incorrectly** treat CRFT-006 and PROG-002 as binding derived constraints.

### 4.2 EIC document explosion

On a single calendar date the project produced, at least:

- 1 investigation  
- 1 candidate  
- 1 falsification pass  
- 2 differently designed simulation specs, both labelled v0.1  
- 3 results reports (minimum Test A/B; package α/β/γ execution; package α/β/γ “v0.1.1”)  
- 1 functional-shapes catalogue  
- 1 γ-closure refinement  
- 1 γ′ retest  
- 1 decision brief  
- 2 human-rulings records  

This is epistemically disciplined and **documentationally chaotic**. Several files restate the same M1–M8 / αβγ / ETA–ICR machinery.

**Near-duplicates (inference):**

- Results v0.1.1 vs Execution & Results v0.1 — same packages, same FAIL/FAIL/INCONCLUSIVE pattern, slightly different γ wording.
- Investigation § architectures A/B/C vs Candidate M1–M8 vs Functional Shapes S/R/C/Mfg catalogues — three encodings of one design space.
- Decision Brief Gates 1–4 vs HD-EIC-01–04 — brief is the question; rulings are the answer. Both remain in the live stream.

### 4.3 Version-identity collision (direct, serious)

Two documents are both titled **EIC Comparative Simulation Specification v0.1** and ask **different questions**:

1. Post-Functional-Shapes spec: packages α / β / γ, Phases 0–6, ETA/ICR.  
2. Earlier spec: Test A (factory residual collapse A0–A8) and Test B (formal vs soft budget).

An LLM citing “Spec v0.1” is ambiguous. This alone can desynchronise two audits.

### 4.4 Information that exists in many places

| Information | Where it is repeated | Risk |
|---|---|---|
| Vision/pillars VIS/PIL | All three GDDs; Investigation §2.1; Matrix §3 | Harmless if statuses stay synced; they did not |
| Non-canonical list | Matrix §5; GDD v1.1.1 §36; Investigation §2.2 | Good, if treated as one list |
| OQ-001–017 | All three GDD §31 copies | Drift if one copy is edited |
| EIC “programme state” tables | Decision Brief, both rulings records, several results docs | **Contradictory “current” rows** |
| “No further LLM consolidation of the Master GDD” | Audit, Matrix, v1.1, v1.1.1, HD-EIC-08 | Consistent and binding as a process rule |

### 4.5 Information that is important but hard for a new LLM to find

1. **Architecture C is retired** — only reliable in HD-EIC-05, late in a long EIC chain.  
2. **CRFT-006 / PROG-002 are PROPOSED** — only reliable in v1.1.1, which is *not* the last GDD in the merge.  
3. **Next legal EIC action** — human structural boundaries (HD-EIC-07), not “write Architecture D.”  
4. **Which simulation is which** — Test A/B vs α/β/γ vs γ′.  
5. **What LOCKED actually means** — Matrix authority classes A–G; many GDD LOCKED rows were later only principle-confirmed.  
6. **Missing controlling files** — Validation Gate brief; Spec v0.1.1; first-pass GDDs.

---

## 5. Authority and Provenance Risks

### 5.1 Documented rules (direct)

From the Authority Matrix:

- **A/B** may remain LOCKED. **C–E are not substitutes for human approval.**  
- **GR-001:** inclusion in the Master GDD ≠ human approval.  
- **GR-002:** superseded/first-pass proposals stay non-canonical unless promoted.  
- **GR-003:** DERIVED CONSTRAINT only for genuine logical consequence.

These rules are correct and should not be weakened.

### 5.2 Risks the current corpus still creates

**Risk 1 — Stale GDD snapshot wins.**  
v1.1 is last in the file and still calls itself the working canonical reference. **Inference:** a naive LLM will treat it as current.

**Risk 2 — Retired leading hypothesis remains rhetorically current.**  
Investigation, Candidate, and several mid-chain docs still say Architecture C is the leading hypothesis. HD-EIC-05 retires it. Without a manifest, both statements look “live.”

**Risk 3 — Human rulings vs GDD scope confusion.**  
HD-EIC-01–08 are HUMAN-LOCKED *for the EIC programme* and **explicitly do not** promote Architecture C or mechanisms into the GDD. An LLM may either (a) ignore them as “not GDD” or (b) smuggle them into GDD LOCKED. Both errors are available.

**Risk 4 — Authority-review queue already resolved, but Matrix still shows “review required.”**  
Matrix flags WRLD-001, PLR-001, PLR-003, CRFT-001, SAFE-001. v1.1.1 records them as confirmed HUMAN-LOCKED, principle only. The Matrix is historically true and **not the current register**.

**Risk 5 — First-pass content is cited, not inspectable.**  
GDD v1.0 claims LOCKED items come from “explicit project brief / consistent across sources,” but those sources are not in the merge. A new LLM cannot independently verify that claim. That is a provenance hole, not a reason to demote those items (demotion would itself be a design act).

**Risk 6 — Prototype-shaped design by rumour.**  
PROT-002 / AS-001: prototype referenced, not inspected. PROT-001 says GDD wins over prototype. The risk is the opposite of promotion: *unreconciled prototype behaviour silently informing later EIC intuition*. Corpus does not show that happening; the gap remains.

**Risk 7 — ID collision.**  
Candidate v0.1 uses **HD-1–HD-9**. Later rulings use **HD-EIC-01–08**. These are different objects. Easy to conflate.

**Risk 8 — Simulation evidence over-read.**  
Minimum Results v0.1 prints numeric tables and immediately classifies them as illustrative / specification-limited. Results v0.1.1 states it is **not** a coded ABM. An LLM can still quote 0.22 specialist share as if it were empirical validation. The documents warn against this; the architecture must make the evidentiary tier unavoidable.

### 5.3 Conflicts (report, do not “fix” design)

| Conflict | Parties | What the project’s own rules imply |
|---|---|---|
| Architecture C leading vs retired | Investigation/Candidate vs HD-EIC-05 | Later **human ruling** wins. C is retired as leading hypothesis. |
| CRFT-006 / PROG-002 status | v1.1 DERIVED vs v1.1.1 PROPOSED | v1.1.1 human ruling wins. |
| Authority-review items | Matrix “review required” vs v1.1.1 confirmed | v1.1.1 confirmation wins for those five principles. |
| γ status | Results v0.1.1 INCONCLUSIVE vs Execution v0.1 leans FAIL vs γ′ retest FAIL | Different experiments. HD-EIC-05 accepts falsification of C *as expressible through the tested shapes*, including γ′. |
| Two Spec v0.1 | Test A/B vs α/β/γ | Identity failure. Neither is design authority. Must be distinguished by experiment ID, not version string. |
| GDD v1.0 “primary authority” vs later working canonical | v1.0 vs v1.1 / v1.1.1 | Succession statements in later GDDs win. |

I am **not** choosing among remaining EIC design options. HD-EIC-07 forbids a new architecture until the human supplies boundaries.

---

## 6. Cross-LLM Auditability Assessment

### If six months from now two LLMs receive this same merge

They can diverge on all of the following without being “careless”:

1. Which GDD version is current.  
2. Whether Architecture C is the working EIC candidate.  
3. Whether CRFT-006 is a constraint or a proposal.  
4. Whether a formal concentration budget is “in” the design (Investigation A/C vs Candidate HD-1 open vs HD-EIC-01 “implementation open”).  
5. Which simulation “result” to cite.  
6. Whether human rulings are GDD-canonical.  
7. What the next authorised action is.

That is a **comparability failure**.

### What would make audits comparable

A new LLM must be able to answer ten questions from a **short, ordered pack**, not from 74 pages of interleaved snapshots:

1. What the project currently is.  
2. What is approved (HUMAN-LOCKED / LOCKED).  
3. What is derived.  
4. What is only proposed.  
5. What is unresolved.  
6. What evidence exists, and its tier.  
7. What is superseded.  
8. Where each important statement came from.  
9. What changed since last state.  
10. What the auditor must independently verify.

Today, (1), (2), (7), and (9) require a full-corpus read plus judgement about succession. That is not reproducible.

---

## 7. Recommended Documentation Architecture

### Design principles (documentation, not game design)

1. **One current pointer.** A manifest is the only document an LLM must trust for “what is current.”  
2. **One current GDD body.** Older GDD versions become immutable archive snapshots.  
3. **Rulings live in a register, not only inside narrative docs.**  
4. **Evidence never lives in the same file as canon.**  
5. **History is preserved, not mixed into the default read.**  
6. **Identity is unique.** No two live or archived controlling docs share the same `{title, version}` meaning different content.  
7. **Active set stays small** even when system specs later appear.  
8. **Do not regenerate the Master GDD** for structure, style, or consolidation (already forbidden by the GDD itself and HD-EIC-08).

### Shape: small active set + typed archive

```
ACTIVE (LLM default pack)
├── TCI-MANIFEST          project state / reading order / succession
├── TCI-GDD-CURRENT       Master GDD v1.1.1 (unchanged body)
├── TCI-GOV-AUTHORITY     authority classes, GR-001–003, non-canon list
├── TCI-GOV-RULINGS       human rulings register (extract, not rewrite)
├── TCI-GOV-CHANGELOG     change identifiers
├── TCI-REG-OQ            machine view of OQ/AS (generated from GDD, or pointer)
└── TCI-EIC-STATUS        thin programme status only

ARCHIVE (consult on demand)
├── gdd/snapshots         v1.0, v1.1, first-pass sources
├── gdd/audits            Canonical Audit
├── eic/investigation
├── eic/candidates
├── eic/specs
├── eic/evidence
├── eic/briefs
└── evidence/prototype    when it exists
```

This is **moderately consolidated**: not one blob, not dozens of live files.

### Concise architecture table

| Document | Purpose | Authority | Active/Archive | Primary Consumer |
|---|---|---|---|---|
| `TCI_Project_State_Manifest` | Declare current files, statuses, reading order, succession | Index only; no design authority | **Active** | Every LLM, every audit |
| `TCIndustries_Master_GDD_v1.1.1_Status_Patch` | Working canonical design text | Working canonical reference | **Active** | Humans + LLMs for *what the game currently is* |
| `TCI_Authority_Governance` | Status vocabulary, GR-001–003, authority classes, non-canon boundary | Governance rules | **Active** | Auditors reconstructing authority |
| `TCI_Human_Rulings_Register` | All human rulings with scope and non-decisions | Human authority (existing rulings only) | **Active** | Auditors reconstructing approval |
| `TCI_Changelog` | What changed, IDs, from→to | Record only | **Active** | Delta audits |
| `TCI_Open_Questions_Index` | Pointers/extract of OQ/AS + EIC open items | None; mirrors sources | **Active** | Auditors listing unresolved items |
| `TCI_EIC_Programme_Status` | Current EIC state in one page | None for mechanics; cites HD-EIC-* | **Active** | Auditors; blocks Architecture D by default |
| GDD v1.0, v1.1 | Prior snapshots | None now | **Archive / superseded** | Provenance only |
| Canonical Audit v1.0 | Why v1.1 happened | Audit record | **Archive / audit** | Provenance |
| Authority Matrix v1.0 | Historical governance checkpoint | Governance (if not folded) | **Active or freeze-as-source** of TCI_Authority_Governance | Provenance |
| All EIC investigation/candidate/spec/result/brief docs | History of a retired search | None | **Archive** | Evidence / lineage |
| HD-EIC source records | Original ruling text | Human authority sources | **Archive + cited by register** | Verification |
| First-pass GDDs (when found) | Pre-consolidation sources | None | **Archive / historical** | Provenance |
| Prototype Seed-2.1 (when available) | Behavioural evidence | Evidence only | **Archive / evidence** | Reconciliation pass |
| Validation Gate brief / Spec v0.1.1 (when found) | Missing controlling files | As stamped on those files | **Archive** | Completeness |

---

## 8. Active Documentation Set

For each proposed **active** document:

### 8.1 `TCI_Project_State_Manifest.yaml` (+ short `.md` wrapper)

- **Purpose:** The front door. Answers “what is current?” without design prose.  
- **Authority:** None for game design. Binding only as an **index** after human authorisation.  
- **Audience:** Any new LLM; human auditors.  
- **Belongs:** current file pointers; versions; statuses; supersedes/superseded-by; ruling-set ID; EIC programme one-liner; prototype status; forbidden operations currently in force; deterministic reading order; last change ID.  
- **Must not belong:** vision essays, mechanisms, recommendations, simulation narrative.  
- **Relationships:** points at every other active file; lists archive roots.  
- **Authoritative?** For *document identity and currency only*.  
- **Update frequency:** every time a document is added, superseded, or a human ruling is recorded.  
- **LLM audit read?** **Always first.**

### 8.2 `TCIndustries_Master_GDD_v1.1.1_Status_Patch.md` (existing; do not regenerate)

- **Purpose:** Sole working canonical design reference.  
- **Authority:** Working canonical. Status tags inside it are authoritative *as of v1.1.1*.  
- **Audience:** Humans and LLMs needing design content.  
- **Belongs:** existing v1.1.1 body.  
- **Must not belong:** new mechanics; EIC Architecture C; simulation numbers; “improvements.”  
- **Relationships:** controlled by governance + rulings register; superseded GDD versions archived.  
- **Authoritative?** Yes, for tagged design statements at their stated status.  
- **Update frequency:** only on explicit human-authorised status patches (same discipline as v1.1.1).  
- **LLM audit read?** Yes, after manifest + governance + rulings. First pass: Document Control, §3–4, §31–33, §36. Full body only if the audit question needs it.

### 8.3 `TCI_Authority_Governance.md`

- **Purpose:** Durable rules for what LOCKED means; GR-001–003; authority classes A–G; non-canonical list; recommended status model.  
- **Authority:** Governance. No new game design.  
- **Audience:** LLMs before they interpret any GDD sentence.  
- **Belongs:** Matrix substance (classes, GR rules, §5 non-canon list, status model). Current confirmation that §3 reconciliation table is **historical** relative to v1.1.1.  
- **Must not belong:** live LOCKED registers (those live in GDD + rulings). Do not keep the Matrix §3 table as if it were the current status board.  
- **Relationships:** interprets GDD tags; does not override human rulings.  
- **Authoritative?** For governance rules only.  
- **Update frequency:** rare (only if humans change governance).  
- **LLM audit read?** **Always, second.**

Implementation note (proposal): this may be the existing Matrix **plus a banner** “reconciliation tables are historical; current statuses are GDD v1.1.1 + Rulings Register.” Prefer a banner over a rewrite.

### 8.4 `TCI_Human_Rulings_Register.md`

- **Purpose:** One scannable list of every explicit human ruling in force, with scope and explicit non-decisions.  
- **Authority:** Restates existing human rulings. Creates none.  
- **Audience:** LLMs reconstructing approval.  
- **Belongs:**  
  - GDD 2026-08-25 confirmations and demotions (from v1.1.1 §36).  
  - HD-EIC-01 through HD-EIC-08, copied as *decision + non-decision + status*, not as essays.  
- **Must not belong:** Decision Brief recommendations; Investigation recommendations; “the LLM thinks…”  
- **Relationships:** cites original ruling records in archive; GDD remains the home of design-statement text.  
- **Authoritative?** Yes, as an index of already-issued human rulings. On conflict, original signed ruling record wins until a human says otherwise.  
- **Update frequency:** whenever a human issues a ruling.  
- **LLM audit read?** **Always, third.**

### 8.5 `TCI_Changelog.md`

- **Purpose:** Change detection without rereading everything.  
- **Authority:** None.  
- **Audience:** Delta audits.  
- **Belongs:** `change_id`, date, actor (human/LLM/process), files affected, from→to status, ruling ID, note “no design change” where true.  
- **Must not belong:** rationale essays, new proposals.  
- **Relationships:** last_change_id duplicated in the manifest.  
- **Authoritative?** For history of documentation state only.  
- **Update frequency:** every authorised documentation event.  
- **LLM audit read?** Yes for “what changed”; skippable for a first full-state audit if manifest `as_of` is enough.

### 8.6 `TCI_Open_Questions_Index.md` (thin)

- **Purpose:** Make unresolved items machine-discernible without forking them from the GDD.  
- **Authority:** None.  
- **Audience:** Auditors listing TBD/OQ/AS.  
- **Belongs:** table of OQ-001–017 and AS-001–005 **pointing at GDD v1.1.1 §31–32**; EIC-only opens (OQ-EIC-FS-*, HD-EIC-07 boundaries not supplied) pointing at EIC status; explicit “not in GDD.”  
- **Must not belong:** answers, preferred options, new questions invented by an LLM.  
- **Relationships:** GDD remains controlling text for GDD OQs. Index must not drift. Prefer generated extract over a second edited list.  
- **Authoritative?** No.  
- **Update frequency:** when GDD OQ section or EIC status changes.  
- **LLM audit read?** Yes, short.

If generation is unwanted, **do not create this file**. Put OQ pointers in the manifest instead. Fewer files beats a second editable register.

### 8.7 `TCI_EIC_Programme_Status.md` (one to two pages)

- **Purpose:** Prevent Architecture C / M1–M8 / Packages αβγ from being read as current design.  
- **Authority:** None for mechanics. Cites HD-EIC-01–08.  
- **Audience:** Any LLM that might otherwise enter the EIC archive first.  
- **Belongs:** programme state table (Investigation complete; C retired; requirement unchanged; next gate = human boundaries; numeric tuning forbidden; Provenance engine deferred; automatic D forbidden); list of archive folders; evidentiary-tier reminder (qualitative ≠ ABM).  
- **Must not belong:** new architectures, residual channels, numbers, “how we might save C.”  
- **Relationships:** HD-EIC-07/08 are the constraints; archive holds the science history.  
- **Authoritative?** Only as a **status pointer**.  
- **Update frequency:** only after a new human EIC ruling.  
- **LLM audit read?** Yes, always, before any EIC archive file.

### Active set size

**Core living set: 5 documents** if OQ index is folded into the manifest.  
**Maximum recommended living set: 7.**  
System-specific design specs should be added later as `TCI-SYS-*` **only after** humans authorise that design phase. Do not pre-create empty system files.

---

## 9. Archive / Provenance Structure

Logical structure (proposal only; no moves now):

```
/archive
  /gdd
    /first-pass                 # the five source GDDs when located
    /snapshots
      Master_GDD_v1.0_Consolidated.md
      Master_GDD_v1.1_Canonical_Baseline.md
    /audits
      Master_GDD_v1.0_Canonical_Audit.md
  /governance
    Authority_Provenance_Reconciliation_Matrix_v1.0.md  # if superseded by TCI_Authority_Governance
    Human_Decision_Brief_EIC_v0.1.md
    Human_Validation_Gate_Architecture_C_v0.1.md        # if located
    Human_Rulings_HD-EIC-01-04_source.md
    Human_Rulings_HD-EIC-05-08_source.md
  /eic
    /investigation
      Economic_Interdependence_Core_Design_Investigation_v1.0.md
    /candidates
      EIC_Candidate_v0.1.md
    /functional_shapes
      EIC_Provisional_Functional_Shapes_v0.1.md
      EIC_Functional_Shapes_Refinement_gamma_closure_v0.2.md
    /specs
      EIC_SimSpec_TestAB_v0.1.md          # rename to kill collision
      EIC_SimSpec_PackagesABG_v0.1.md
      EIC_SimSpec_v0.1.1.md               # if located
    /evidence
      EIC_Falsification_Pass_Candidate_v0.1.md
      EIC_Minimum_Simulation_Results_TestAB_v0.1.md
      EIC_Comparative_Results_Packages_v0.1.md
      EIC_Comparative_Results_v0.1.1.md
      EIC_gamma_prime_Retest_Results_v0.1.md
  /historical_proposals
    non_canonical_boundary.md             # copy of Matrix §5, labelled HISTORICAL
  /prototype
    seed-2.1/                             # when available; evidence only
    reconciliation_matrix.md              # not yet created; do not invent
```

**Rules for archive (proposal):**

- Immutable once archived. Corrections happen by **new** documents, not silent edits.  
- Every archive file header: `status: SUPERSEDED | EVIDENCE | HISTORICAL | AUDIT_RECORD` and `superseded_by` / `do_not_cite_as_current`.  
- Historical proposals stay readable. GR-002 remains the contamination control, not deletion.  
- Prototype folder is an **evidence track**. PROT-001 still applies.

**Candidates for eventual retirement (not deletion):**

- Near-duplicate results (keep one as “record of publication,” the other marked `DUPLICATE_OF`). Retirement = “not in any default pack,” not erase.  
- Decision Brief after its rulings are in the register (keep as provenance of the question set).

**Not candidates for retirement:** human ruling source records; GDD snapshots; Canonical Audit; Matrix; γ′ retest (basis of HD-EIC-05).

---

## 10. LLM Audit Reading Protocol

Publish this protocol **in the manifest**. Same instructions → comparable audits.

### Deterministic reading order

| Step | Read | Stop condition |
|---|---|---|
| 0 | This protocol + manifest | If manifest is missing, **report UNCERTAIN / INTAKE FAILURE** and do not invent current state |
| 1 | `TCI_Authority_Governance` | Learn status vocabulary and GR-001–003 |
| 2 | `TCI_Human_Rulings_Register` | Build the approved-decision set |
| 3 | GDD v1.1.1 Document Control + §3–4 + §31–33 + §36 | Current design status board |
| 4 | `TCI_EIC_Programme_Status` | Current EIC programme; do not open archive yet |
| 5 | GDD body sections **only as needed** for the audit question | Do not “re-summarise the whole GDD” by default |
| 6 | `TCI_Changelog` if the task is a delta audit | Compare `from` state if provided |
| 7 | Archive **only** to verify a cited claim, reconstruct provenance, or inspect evidence tier | Never treat archive prose as current |

### How to identify current canonical state

1. Manifest `current_gdd.version` must match the GDD header.  
2. If they disagree: **report conflict; do not choose**.  
3. GDD status tags are the design-state. Rulings register may constrain EIC *programme* without changing a GDD tag.

### How to identify authority

| If the statement is… | Treat as… |
|---|---|
| In rulings register as HUMAN-LOCKED, scoped | Human decision, within stated scope only |
| GDD tag LOCKED / HUMAN-LOCKED | Working canon, still subject to GR-001 (presence is not extra approval) |
| GDD tag DERIVED CONSTRAINT | Binding consequence, not independent human decision |
| GDD or EIC PROPOSED / TBD / ASSUMPTION | Not approved |
| Prototype / simulation | Evidence, not design |
| First-pass / Matrix §5 list | HISTORICAL / NOT CANONICAL |
| Repeated in many LLM docs | Still PROPOSED unless a human ruling says otherwise |

### How to distinguish current facts from proposals

- Ignore verbs like “should,” “recommended,” “leading hypothesis” unless the **status tag** or **ruling ID** agrees.  
- Architecture C recommendation is **historical investigation**, not current fact.  
- Numeric tables in simulation reports are **illustrative evidence**, not balance values (BAL-001; reports’ own disclaimers).

### How to identify unresolved questions

- GDD §31 OQ-*, §32 AS-*.  
- EIC status: HD-EIC-07 boundaries not supplied; OQ-EIC-FS-* remain unresolved; no canonical EIC.  
- Do not close an OQ because a simulation “suggested” an answer.

### How to inspect changes

- Changelog `change_id` sequence.  
- If no changelog, compare manifest `as_of` + GDD §36 + rulings dates — and **report low change-detection confidence**.

### How to consult historical material

Only if:

- verifying a provenance citation,  
- checking whether a proposal is on the non-canonical list,  
- reconstructing why a ruling was made,  
- assessing evidence tier.

When consulting it, prefix every extracted claim: `ARCHIVE / NON-CURRENT`.

### How to report uncertainty

Required phrases:

- `NOT ESTABLISHED BY CORPUS`  
- `CONFLICTING DOCUMENTS: A vs B`  
- `EVIDENTIARY TIER: qualitative | illustrative-numeric | human-ruling | gdd-tag`  
- `INTAKE FAILURE` if the active pack is incomplete

Never fill a gap with a preferred design.

### How to cite evidence

Citation minimum:

`[doc_id | version | section/ID | status-tag-or-ruling | evidentiary-tier]`

Example: `[GDD-1.1.1 | §12 CRFT-006 | PROPOSED | gdd-tag]`  
Example: `[HD-EIC-05 | Architecture C retired | HUMAN-LOCKED | human-ruling]`  
Example: `[γ′-retest v0.1 | §6 | evidence only | qualitative-adversarial]`

### Standardized audit result (required output shape)

```text
AUDIT_RESULT
  auditor: <model/id>
  protocol_version: <from manifest>
  manifest_as_of: <date/id>
  intake: COMPLETE | INTAKE_FAILURE
  current_gdd: <version or CONFLICT>
  human_rulings_in_force: [IDs]
  eic_programme_state: <one line from EIC status>
  locked_sample_verified: [IDs checked]
  derived_constraints: [IDs]
  proposed_not_to_be_treated_as_locked: [IDs found at risk]
  unresolved: [OQ/AS/HD-EIC-07…]
  superseded_seen_and_ignored: [docs]
  evidence_consulted: [docs + tier]
  conflicts: [list]
  not_established: [list]
  independent_verification_needed: [list]
  design_decisions_made_by_this_audit: NONE
```

Two LLMs using this shape on the same active pack should match on `current_gdd`, ruling IDs, EIC one-liner, and conflict list even if their prose differs.

---

## 11. Recommended Standard Metadata

Apply as YAML front matter on **every** document going forward. Do not retrofit by rewriting GDD prose.

```yaml
doc_id: TCI-GDD                 # stable ID
title: Master Game Design Document
version: "1.1.1"
date: 2026-08-25
doc_class: GDD | GOVERNANCE | HUMAN_RULING | EVIDENCE | SPEC | BRIEF | AUDIT | MANIFEST | HISTORICAL
lifecycle: ACTIVE | SUPERSEDED | ARCHIVE | MISSING_CITED
authority: working_canonical | human_ruling | governance | none
human_approval_state: not_applicable | recorded | none
supersedes: [TCI-GDD@1.1]
superseded_by: null
controlling_document: TCI-MANIFEST
provenance_sources: [TCI-GDD@1.0, TCI-AUDIT-GDD-1.0, TCI-MATRIX-1.0]
evidentiary_tier: n/a | qualitative | illustrative_numeric | prototype | human_ruling
creates_design_authority: false
llm_default_read: true | false
change_id: CHG-0000
```

**Keep bureaucracy minimal.** The mandatory fields are: `doc_id`, `version`, `lifecycle`, `authority`, `supersedes`/`superseded_by`, `creates_design_authority`, `llm_default_read`.

**Also assign unique experiment IDs** (proposal):

| Experiment ID | Current colliding name |
|---|---|
| `EIC-EXP-TESTAB-0.1` | Comparative Simulation Specification v0.1 (Test A/B) |
| `EIC-EXP-ABG-0.1` | Comparative Simulation Specification v0.1 (packages) |
| `EIC-EXP-ABG-0.1.1` | Spec v0.1.1 (missing from merge) |
| `EIC-EXP-GAMMA-PRIME-0.1` | γ′ retest |

Do not leave two files as “Spec v0.1.”

**Manifest fields (minimum):**

```yaml
project: TCIndustries
manifest_version: 1
as_of: 2026-08-25
current_gdd: { doc_id: TCI-GDD, version: "1.1.1" }
rulings_in_force: [GDD-2026-08-25-CONFIRMS, HD-EIC-01, HD-EIC-02, HD-EIC-03, HD-EIC-04, HD-EIC-05, HD-EIC-06, HD-EIC-07, HD-EIC-08]
eic:
  architecture_c: RETIRED
  canonical_eic: NOT_ESTABLISHED
  next_authorized_operation: human_structural_boundaries
  automatic_architecture_d: NOT_AUTHORISED
prototype_reconciliation: NOT_DONE
gdd_rewrite: NOT_AUTHORISED
numeric_tuning: NOT_AUTHORISED
audit_protocol_version: 1
audit_reading_order: [MANIFEST, AUTHORITY, RULINGS, GDD_STATUS_SECTIONS, EIC_STATUS, CHANGELOG, GDD_BODY_AS_NEEDED, ARCHIVE_ON_DEMAND]
last_change_id: CHG-xxxx
conflicts_known: [SPEC-V0.1-COLLISION, GDD-V1.1-AFTER-V1.1.1-IN-MERGE]
missing_cited: [first-pass GDDs, Validation Gate brief, SimSpec v0.1.1, Seed-2.1]
```

---

## 12. Migration / Consolidation Plan

**Proposal only. No files are to be changed until the project owner authorises a stage.**

### Stage 0 — Authorise the architecture (human)

Approve or amend this assessment. Confirm that creating **index/register** documents is **not** a Master GDD rewrite (HD-EIC-08). If the owner considers any new file a “GDD consolidation,” **stop**.

### Stage 1 — Manifest + changelog only

- Create empty/populated manifest pointing at **existing filenames**.  
- Create changelog entries describing *documentation state*, not design.  
- Do not move, rename, or edit GDD/EIC bodies.  
- **Exit:** a new LLM can name the current GDD and HD-EIC set from the manifest alone.

### Stage 2 — Logical succession, no deletion

- Manifest marks v1.0 and v1.1 `SUPERSEDED`, v1.1.1 `ACTIVE`.  
- Optional one-line banner at top of archived GDD copies: “SUPERSEDED — see manifest.” Only if humans allow a one-line stamp.  
- **Exit:** Test 1 begins to pass.

### Stage 3 — Rulings register (extract, do not rewrite)

- Copy HD-EIC-01–08 and v1.1.1 §36 rulings into the register.  
- Leave source records untouched.  
- **Exit:** Test 2 on authority questions becomes possible.

### Stage 4 — EIC status page

- One-page pointer: C retired; do not read archive as current.  
- **Exit:** Test 3 (historical containment) becomes possible for Architecture C.

### Stage 5 — Identity repair (labels only)

- Assign experiment IDs to the two Spec v0.1 files in the manifest.  
- Do not merge or delete either file.  
- Mark Results v0.1.1 and Execution Results v0.1 as related/overlapping evidence.

### Stage 6 — Locate missing cited files

- Place first-pass GDDs, Validation Gate brief, Spec v0.1.1, and prototype into archive **if they exist**.  
- If they do not exist, manifest `missing_cited` stays populated. Do not recreate first-pass GDDs from memory.

### Stage 7 — Optional OQ index

- Only if humans want a machine table. Prefer generation from GDD §31–32.

### Stage 8 — Metadata on *new* documents only

- All future artifacts carry front matter.  
- Ban new documents that lack `lifecycle`, `authority`, `creates_design_authority`.

### Stage 9 — Prototype evidence reconciliation (already required by GDD)

- Separate work item. Produce GDD ↔ prototype ↔ decision matrix.  
- Do not start it by regenerating the GDD.

### Explicitly out of this plan

- Regenerating or restyling the Master GDD.  
- Resolving OQ-001–017.  
- Inventing Architecture D or HD-EIC-07 boundaries.  
- Deleting any historical file.  
- Promoting CRFT-006, PROG-002, or any EIC mechanism.  
- Numeric tuning.

---

## 13. Risks and Trade-offs

| Choice | Benefit | Cost / risk |
|---|---|---|
| Add a manifest | State reconstruction becomes possible | Another file; if not maintained, it becomes a new lie |
| Keep GDD v1.1.1 untouched | Obeys HD-EIC-08 and “no further consolidation” | Active GDD remains long; auditors must use sectional first pass |
| Extract rulings register | Approval set is scannable | Dual-source risk if register drifts from signed records |
| Archive all EIC long-form | Stops C from re-entering canon | Scientists must know to open archive for evidence |
| Fold OQ into manifest vs separate index | Fewer files vs nicer tables | Second editable OQ list will drift — **avoid dual edit** |
| One mega-document | Fewer files | Destroys containment; already failed conceptually in this 74-page merge |
| Dozens of system stubs now | Looks organised | Empty files become fake canon; mechanical design is early (Matrix §8) |
| Stamping SUPERSEDED on old GDDs | Stops last-file-wins | Requires human permission to touch those files at all |
| Qualitative sims sitting next to GDD | Looks like one project | LLMs quote illustrative numbers as balance |

**Largest residual risk after adoption:** an LLM is given the **old merge** instead of the **active pack**. The manifest must be the intake contract.

---

## 14. Human Authorization Gates

Nothing in this assessment is self-executing.

| Gate | What the owner must decide | If declined |
|---|---|---|
| G0 | Accept this *documentation* architecture as a proposal | Status quo remains; Tests 1–3 still fail |
| G1 | Allow creation of manifest + changelog (index only) | Cannot fix state reconstruction |
| G2 | Confirm v1.1.1 is the sole active GDD pointer | Succession remains ambiguous |
| G3 | Allow a rulings register that *copies* existing rulings | Authority stays buried |
| G4 | Allow a one-page EIC status pointer | Architecture C containment stays weak |
| G5 | Allow SUPERSEDED banners / logical archive labels | Optional; stronger containment |
| G6 | Allow experiment-ID aliases for colliding Spec v0.1 | Optional; reduces citation errors |
| G7 | Confirm new index files are **not** a GDD rewrite under HD-EIC-08 | If the owner says they *are* a rewrite, **do not create them** |
| G8 | Later: prototype reconciliation; locating missing cited files | Completeness, not structure |

**Not requested and not authorised here:** HD-EIC-07 structural boundaries, any EIC redesign, any GDD status change.

---

## 15. Target State

Six months from now, a cold-start LLM should receive:

1. The **active pack** (5–7 files), not a 74-page unordered merge.  
2. A manifest whose `as_of` and `last_change_id` define the audit baseline.  
3. One GDD body, one rulings list, one governance ruleset, one EIC status page.  
4. An archive that is complete, labelled, and off the default path.  
5. Unique IDs so “Spec v0.1” cannot mean two experiments.  
6. An audit protocol that forces `design_decisions_made_by_this_audit: NONE`.

**Target test outcomes:**

- **Test 1:** Pass — current GDD version, LOCKED/DERIVED/PROPOSED split, EIC “C retired / no D,” and OQ list recoverable from the active pack alone.  
- **Test 2:** Pass — two LLMs emit the same `AUDIT_RESULT` core fields.  
- **Test 3:** Pass — encountering Candidate v0.1 or GDD v1.0 cannot change current canon because those files are `lifecycle: SUPERSEDED` / `ARCHIVE` and not in `llm_default_read`.

**What remains intentionally early (direct from corpus, not a documentation defect):**

- Mechanical design maturity: early.  
- Economic interdependence: not canonically demonstrated; Architecture C retired as currently shaped.  
- Prototype: unreconciled.  
- Ready for canonical freeze: **No** (v1.1.1 and Matrix already say this).  
- Ready for major system design: after human confirmation (done for the listed principles) **and** optional prototype reconciliation; EIC still waiting on HD-EIC-07 boundaries.

---

## 16. Open Questions / Information Not Established by the Corpus

These are **documentation/completeness** gaps, not invitations to invent design.

1. Whether the five first-pass Master GDDs still exist and match v1.0’s source list.  
2. Full text of `TCIndustries_EIC_Human_Validation_Gate_Architecture_C_Falsification_v0.1.md`.  
3. Full text of `TCIndustries_EIC_Comparative_Simulation_Specification_v0.1.1.md`.  
4. Whether Seed-2.1 exists, what it implements, and how it differs from GDD (AS-001 / PROT-002).  
5. The **human structural boundaries** required by HD-EIC-07 — **not supplied**.  
6. Whether Results v0.1.1 and Execution Results v0.1 were meant as two publications or a duplicate.  
7. Whether the owner considers the Authority Matrix still the live governance text or a completed checkpoint.  
8. Whether any human rulings exist *outside* this merge (not established).  
9. Intended filename/storage conventions in the real repo (the merge has titles, not a file tree).  
10. Whether “identical sources” chatgptv1-1 and claude1 are truly byte-identical (asserted by v1.0; not verifiable here).  
11. Process owner for maintaining a manifest (not established).  
12. Whether qualitative simulation reports should be one evidence series or retired as non-discriminating — a **records** question, not a design question.

I do not fill these gaps.

---

# Human Authorization Required

The following structural actions require the project owner’s approval **before anyone implements them**. This assessment does not perform them.

1. **Create** a project-state manifest and changelog.  
2. **Designate** Master GDD v1.1.1 as the only active GDD in that manifest (leave file contents unchanged).  
3. **Create** a Human Rulings Register by extracting already-issued rulings (HD-EIC-01–08 and v1.1.1 §36).  
4. **Create** a one-page EIC Programme Status pointer stating Architecture C is retired and Architecture D is not authorised.  
5. **Mark** GDD v1.0 and v1.1 as superseded in the index (optional one-line SUPERSEDED banner if you allow any edit).  
6. **Assign** distinct experiment IDs to the two “Spec v0.1” documents.  
7. **Adopt** the LLM audit protocol and metadata fields as intake contract.  
8. **Confirm** that items 1–7 are **documentation indexing**, not a Master GDD rewrite under HD-EIC-08.  
9. **Do not** move, rename, merge, delete, or regenerate existing design files until you explicitly authorise a later stage.  
10. **Do not** treat this assessment as a game-design ruling, an EIC architecture, or a promotion/demotion of any VIS/PIL/OQ/HD item beyond what the cited human records already say.

**If no authorisation is given, the correct documentation state remains exactly the corpus as supplied.**