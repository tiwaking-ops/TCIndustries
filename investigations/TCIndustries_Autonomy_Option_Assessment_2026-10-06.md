---
doc_id: TCI-AUTONOMY-OPTION-ASSESSMENT-2026-10-06
title: "TCIndustries — improve-in-place vs new autonomous project: comparative assessment and recommendation"
version: "1.0"
date: 2026-10-06
status: ADVISORY (observation and recommendation)
authority_class: null
human_approved: false
author: "OrcaRouter Auto (orcarouter/auto)"
author_llm: {name: "OrcaRouter Auto", version: "orcarouter/auto"}
assessor_llm: []              # none yet — independent assessment required before any promotion
controlling_documents: [AGENTS.md, governance/TCIndustries_Human_Rulings_Register.md, canonical/gdd/TCIndustries_Master_GDD_v1.1.1_Status_Patch.md]
supersedes: []
superseded_by: null
---

# TCIndustries — Improve the existing project system, or create a new autonomous LLM-managed project?

**Comparative assessment against the `TiwasProjects` `.pm` system, with a recommendation.**

- **Author:** OrcaRouter Auto (`orcarouter/auto`)
- **Date:** 2026-10-06
- **Subject A:** `C:\temp\projects\tcindustries` (hereafter **TCI**)
- **Subject B:** `C:\temp\projects\TiwasProjects` (hereafter **TW**)
- **Authority:** ADVISORY. This document is an observation. It changes no status, promotes nothing,
  amends no governance file, and validates no design decision. Placed in `investigations/`
  per `AGENTS.md` §"Documentation layout" (research, audits, comparisons). Any recommendation
  below requires its own human ruling before it becomes binding.

---

## 0. Verdict (answer first)

> **Improve TCI in place. Do not create a new repository and do not move the corpus.**
> More precisely: adopt a **third option** — transplant TW's deterministic coordination kernel
> *underneath* the existing TCI repository, leaving every governing document byte-identical in
> place, and add a small TCI-specific layer of checks that neither system has today.

**One-sentence reason.** TCI's stall is real but it is **not** caused by the repository, its layout,
its history, or its authority model — all four are assets — it is caused by the fact that TCI has
**zero executable instruments**: every consistency guarantee is performed by hand by an LLM session
and therefore decays the moment the session ends, while TW has already built and self-tested the
deterministic substrate that performs exactly those guarantees.

**The load-bearing caution.** A new repository would not remove the human gate. There are roughly
**8 genuine design decisions** and one structural boundary statement (HD-EIC-07) that only you can
make. No amount of autonomy answers them. What autonomy does is: (a) stop the project from losing
authority that you have *already* given, and (b) make each remaining decision cheap, bounded and
pre-packaged instead of open-ended.

**Quantified basis.** Your own record shows a **~9-day propagation latency** between six rulings being
given and them becoming detectable in the authority record (ruled 2026-09-19/20, filed 2026-09-28).
That interval was filed as the precipitating incident for GR-006. It is not "you being slow" — it is a
propagation defect, and propagation is precisely what machines do.

---

## 1. Purpose, method, and reproducibility

**Purpose.** Compare the project-management infrastructure of TCI against TW, diagnose why TCI has
stalled, and recommend whether to (A) improve TCI's infrastructure in place, (B) create a new
autonomous project and migrate TCI into it, or (C) some third course.

**Method.** Static read-only survey of both repositories conducted 2026-10-06:

- directory and file inventories with sizes and last-write times;
- git history, dirty-tree state, ahead/behind versus upstream, tracked-file largest-size check;
- full read of TCI `AGENTS.md`, `README.md`, `governance/README.md`, `governance/project_memory.md`,
  `proposals/TCIndustries_Continuous_Queue_Tracker.md`,
  `proposals/TCIndustries_Open_Questions_Register.md`,
  `investigations/TCIndustries_Corpus_Verification_Audit_2026-09-28.md`;
- full read of TW `AGENTS.md`, `README.md`, `policy.yaml`, `.pm/decisions/D-0002.md`,
  `docs/roadmap.md`, all `.pm/events/*.jsonl`, both `views/`, the task card, incident and pattern files;
- inventory of TW kernel subcommands and source size;
- **re-verification today** of the specific defects from TCI's own 2026-09-28 corpus audit (see §5.2).

**No file in either repository was modified by this assessment** except this report. No TCI or TW
tool or kernel command was executed; all counts are single-snapshot filesystem/git measurements and
will drift. Commands used are reproducible and listed in **Appendix A**.

**Honest frame.** Both repositories belong to the same operator, on the same host, are pushed to the
same GitHub account, and both already contain comparison reports written about *other* projects
(a Babylon 5 comparison exists in TW; a mirrored TW comparison exists in the Babylon 5 repo). The
operator is evidently iterating on a reusable autonomy pattern. This report is one input to that.

---

## 2. Subject A — TCIndustries as built

### 2.1 Measured state (2026-10-06)

| Signal | Measured value |
|---|---|
| Git | 24 commits, single branch `main`, remote `github.com/tiwaking-ops/TCIndustries`, **0 ahead / 0 behind** |
| Last commit | `3ccc198`, 2026-09-29 20:05 (+1300) — *"Rule HD-ECO-01 … and HD-PROF-03 … → register Section E"* |
| Elapsed since last tracked change | **7.06 days** (no activity anywhere in the tracked tree) |
| Working tree | 5 dirty entries (see §2.3) |
| `.git` size | 10.4 MB (large artifacts are **not** committed — good) |
| Total on-disk size | 2,232 MB, of which ~2,212 MB is untracked local testbed data |
| Markdown corpus (excluding testbed) | **160 files / 89,996 lines / 4.58 MB** |
| Human rulings on record | **22 distinct `HD-*` IDs** across the record set |
| Rulings with a dedicated register entry | **20** — HD-EIC-09 has *zero* occurrences in the register; HD-TST-01 has a cumulative-state row only |
| Invariant candidates | 28 distinct `INV-*` IDs |
| Master GDD open questions (OQ-001–017) | 6 RULED, 1 DEFERRED, **10 TBD** |
| GDD §36 remaining human decisions | 10 listed, 2 ruled → **8 open** |
| Live ruling queue | **0 items** (own stated policy: maintain 3–5) |
| **Executable governance/coordination instruments** | **0** (no `.pm`, no tools/, no scripts/, no tests, no sentinel, no citation checker) |
| One-off analysis code | 1 file (`investigations/resource-sim/main.go`) |

Directory census (excluding `.git`):

| Directory | Files | .md | Size | Last write |
|---|---|---|---|---|
| `testbed/` | 94 | 5 | 2,212.38 MB | 2026-09-17 |
| `sources/` | 223 | 44 | 5.44 MB | 2026-09-14 |
| `investigations/` | 43 | 40 | 0.95 MB | 2026-09-29 |
| `proposals/` | 47 | 47 | 0.64 MB | 2026-09-29 |
| `archive/` | 9 | 9 | 0.41 MB | 2026-09-12 |
| `governance/` | 15 | 14 | 0.20 MB | 2026-09-29 |
| `canonical/` | 3 | 3 | 0.05 MB | 2026-09-29 |

### 2.2 What the system actually is

TCI is a **human-ruling-gated documentation republic**, and it is unusually well-designed for what it
is. Its distinctive properties, all of which are assets:

1. **Status tokens with teeth.** `LOCKED`, `DERIVED CONSTRAINT`, `PROPOSED`, `TBD`, `PROTOTYPE/EVIDENCE`,
   `DEFERRED`, `ASSUMPTION`, `HISTORICAL` (`AGENTS.md` §"Mandatory status discipline"). The rule
   "never silently promote any non-LOCKED status into LOCKED" is enforced by instruction, and a
   2026-09-28 status-discipline sweep found **no violations**.
2. **Evidence and history separation.** Prototype behaviour is evidence, never design authority;
   Pre-CU SWG material is `HISTORICAL` and never inherits authority (SWG-001/002).
3. **A real authority spine.** `governance/TCIndustries_Human_Rulings_Register.md` accumulates
   Sections A–E of verbatim human rulings, each with decision and traceability blocks, plus a full
   cumulative-state table. Versioned (v1.4), dated, `human_approved: true`.
4. **A deliberate anti-invention doctrine.** "Document the project that exists; do not invent the
   project that ought to exist." Several gate records explicitly refuse to invent a successor
   architecture after falsification (HD-EIC-05/07/08) — a discipline most LLM-driven projects
   fail at completely.
5. **Governance rules that emerged from observed defects.** GR-004 (attribution), GR-005 (filename
   convention), GR-006 (detective/corrective control) each exist because a specific failure was
   measured. GR-006 additionally states the right remedy rule: *a filing defect never voids a ruling.*
6. **Preserved failure history.** Duplicate passes, mis-targeted authorisations, invented values,
   partial-context sessions — all recorded rather than smoothed over. This is rare and valuable.

### 2.3 What the system lacks — the structural hole

There is **no machine anywhere in the system.** Every guarantee above is performed by an LLM session
reading files and promising consistency. Observable consequences today:

| Observation (measured 2026-10-06) | Class |
|---|---|
| Root contains untracked `swg.db`, `swg.db-shm`, `swg.db-wal` (~1.67 MB SQLite working state) | hygiene / root litter |
| A **0-byte file literally named `fs.statSync(x).isFile()))`** exists at `sources/swg-pre-cu/track-a-b-code/` and is untracked — the fossil of a mistyped shell command that created a file instead of branching | hygiene / no filename gate |
| Untracked `.freebuff/project-id` at root | hygiene |
| `testbed/swg-phase3-combat/swg-server/swg.db` is **2,168 MB** of local data, excluded only by a *nested* `.gitignore` inside the vendored tree — not by any project policy | repo policy |
| Duplicate detection is a **manual pre-flight grep** repeated every cycle (`Queue Tracker` §"Method"), a ritual standing in for a tool | automation gap |
| The authority record is hand-pasted into the register by an agent at explicit instruction; each ruling triggers a **five-act filing ritual** across register, OQR, Invariant Register, GDD status lines and `project_memory` | propagation gap |
| **All due-diligence checklist checkboxes are still unchecked**, although `investigations/SWG_Code_Due_Diligence_Evidence_Log_2026-09-14.md` records steps 1–4 **and** step 5 as executed | tracker/evidence divergence |
| `project_memory.md` is a **41-line append-only Markdown file** that is simultaneously narrative memory, change log and audit trail | mutable shared state |

None of these are failures of care. Every one is the predictable result of asking sessions to be
deterministic.

---

## 3. Subject B — TiwasProjects `.pm` as built

### 3.1 Measured state (2026-10-06)

| Signal | Measured value |
|---|---|
| Kernel | `.pm/kernel/pmk.py`, Python **stdlib only**, **1,574 lines**, version `k-0.2.0` |
| Commands (19) | `init brief admit claim renew heartbeat init-bundle run submit record-decision record-incidents note-pattern sweep sentinel views status digest check-citations check-root` |
| Event log | 8 events across 2 day-files, **hash-chained** (`seq` + `prev_hash` + `hash`, verified end-to-end from `GENESIS`) |
| Tasks | 1 (`T-0001`, state `READY`, **never claimed**; declared `scope.write`, oracles, premises, budget, stop-conditions) |
| Decisions | 2 (`D-0001` bootstrap; `D-0002` adopting R1–R15, hash `1a00f2fa…`, `prev_hash: GENESIS`) |
| Evidence bundles | **0** (`.pm/evidence/` empty) |
| Tests | `tests/test_pmk.py` (14 tests incl. **red-witness fixtures that must fail**) + `tests/test_repo_hygiene.py` (4) |
| Gate results claimed in README, dated 2026-10-06 | unittest exit 0; sentinel exit 0; `check-citations` exit 0; `check-root` exit 0 |
| Policy | `.pm/policy.yaml`, hash-pinned, **two-key change**: spend envelopes, dependency allow-list (`stdlib` only), autonomy ceiling R2, risk floors → R3/R4, lease TTL/renew/grace (30/10/15 min, max 3 attempts), V0–V4 verification ladder, human queue 72 h default-hold, **interrupt budget 10/week, 15 human-minutes per accepted unit** |
| Incident/pattern hygiene | `INC-0001` auto-filed from the single sentinel finding; `PAT-0001` generalises it, `review_by` date set, tier explicitly "advisory (never evidence for a claim)" |
| Honest self-reported gaps (D-0002 §"Residual gaps") | no merge queue (SUBMITTED→PROMOTED path absent); bundle narrative sections hand-written; V1/V2 tiers specified not executed; heartbeat correlation weak; interrupt budget is a proxy count |

### 3.2 What it actually provides

TW is **young and thin** (one never-claimed task, zero bundles) but it has the property TCI lacks:
**invariants are code.** Concretely available today:

- **Sentinel invariants** (I1 chain integrity, I3, I4, I5 sequence contiguity, I8 kernel-pin match)
  runnable as one command, with self-tests and red witnesses.
- **Citation existence checking** (`check-citations`) — every `evidence[]` path in a decision must resolve.
- **Root hygiene policy** (`ROOT_ALLOW` + `check-root`) enforced by test.
- **Verified writes** — write, re-read, compare; no snapshot-restore clobber class.
- **Harness-recorded evidence** (`pmk run -- <cmd>` records cmd/cwd/exit/stdout-hash into the bundle),
  so "harness-recorded, not typed" is mechanically true.
- **Generated views** (`views/status.md`, `views/queue.md`, `views/ledger.md`) stamped
  `_measured=<timestamp> kernel=<version> seq=<n>_` — projections, never hand-edited.
- **A human contract with numbers on it**: interrupt budget and default-hold, so the human's cost is
  bounded by policy rather than by the agent's enthusiasm.
- **Clock discipline**: tool-stamped timestamps with payload-vs-clock acceptance, and a three-signal
  liveness rule (claim + heartbeat + report; absent ⇒ UNKNOWN, never STALE).

### 3.3 What TW does **not** have — read this before adopting anything

TW is a *code-project* kernel. It has **no concept of a design-document domain**:

1. **No status-token semantics.** It knows task states, not `LOCKED`/`PROPOSED`/`TBD`/`DERIVED CONSTRAINT`.
2. **No cross-register equality checking.** Nothing compares "the same fact in six documents".
3. **No per-ID completeness or uniqueness enforcement.** Nothing would catch two rules sharing `GR-004`.
4. **No Markdown metadata lint** (nothing catches a document missing `Author`/`Version`/`Date`).
5. **No content-addressed document registry.** `.pm/registry/sources/` is **empty** and no source has a hash.
6. **No ruling-packet generator.** Its human queue lists decision classes, not prepared one-screen
   decisions with options and ready-to-file text.
7. **Merge path incomplete.** `SUBMITTED` has no route to `PROMOTED`, which means autonomous work has
   no terminal landing in TW today.

**Consequence:** transplanting the kernel into TCI is necessary but **not sufficient**. §7 specifies
the TCI-specific layer that must be built either way.

---

## 4. Dimension-by-dimension comparison

Legend: ✅ present and exercised · 🟡 present but thin/partial · ❌ absent.

| Dimension | TCI | TW | Note |
|---|---|---|---|
| Executable gate(s) | ❌ **zero** | ✅ sentinel + unittest + citation + root, self-tested | The headline gap |
| Tests / red witnesses | ❌ none | ✅ 18 tests, must-fail fixtures | TW's best property |
| Mutation-resistant history | ❌ mutable append-only Markdown (`project_memory.md`) | ✅ hash-chained JSONL, kidney-owned | TCI's highest-loss-risk artifact |
| Generated vs hand-maintained views | ❌ six registers maintain the same facts by hand (documented divergence, V-07) | ✅ projections carry `measured=` stamps | Directly explains V-07 |
| Citation / reference existence | ❌ manual; **live phantom document cited by 7 docs** | ✅ `check-citations` | Directly explains V-03/V-04 |
| Per-ID completeness & uniqueness | ❌ missing ruling entries (V-01/V-02), duplicate rule ID (V-05) | 🟡 IDs exist for decisions/incidents; nothing for document-domain IDs | Needs the new check layer |
| Metadata lint | ❌ **69 of 82 documents** lack ≥1 required GR-004 field | ❌ none either | Neither has it |
| Evidence standard | 🟡 strong doctrine (2.3/2.4), manual execution | ✅ harness-recorded receipts, V0 machinery | Doctrinal vs mechanical |
| Verification independence | ❌ self-verification is the terminal step | 🟡 V0–V4 ladder; **V1/V2 not yet executed** | Both incomplete; TW's design wins |
| Failure infrastructure | ✅ excellent narrative history; ❌ no incident objects | ✅ `INC-0001` + `PAT-0001` with review dates | Different maturity axes |
| Human surface / cost bound | ❌ ad hoc; one round-trip per ruling | ✅ policy-bounded (10/wk, 15 min/unit, 72 h hold), not yet populated | TW design wins, unproven in practice |
| Task locking / concurrency | ❌ none needed (single thread today) | ✅ leases + fencing + quarantine | Only matters if TCI parallelises |
| Provenance conventions | ✅ rich (author_llm/assessor, compaction) | ✅ same global rules, applied sparsely | Both strong |
| Domain content itself | ✅ 22 rulings, 28 invariants, 90k-line GDD corpus | 🟡 bootstrap | **TCI's asset — do not move it** |
| Maturity / real traffic | ✅ ~6 weeks of real ruling traffic | ❌ 1 unclaimed task, 0 bundles | TCI has the field data TW lacks |

**Summary.** TCI has the *content and doctrine*; TW has the *machinery*. They fail in almost
perfectly complementary ways, which is exactly why "import one into the other" is better than
either "keep going by hand" or "start over".

---

## 5. Diagnosis: why TCI actually stalled

The brief says the cause is "the human being slow to decide and implement changes". That is partly
true and partly a misdiagnosis. Separating the three mechanisms matters, because they have different
cures.

### 5.1 Mechanism 1 — the irreducible human gate (real, small, permanent)

Genuinely yours to decide, today:

- **HD-EIC-07**: what structural source of economic comparative advantage the project will durably
  commit to. **Open since 2026-08-25 — 42 days.** It blocks OQ-001, OQ-004, OQ-005 and pre-work T-03/T-06.
- **8 open GDD §36 decisions**: specialisation budget, resource lifecycle maths, crafting
  experimentation, manufacturing constraints, combat scope/risk, currency stabilisers,
  multi-account *implementation*, prototype reconciliation.
- **Unclosed slots on partial rulings**: HD-ECO-01 scope + INV-009b; HD-PROF-03 scope + INV-011c.
- **GR-006's four unanswered slots** (WINDOW / RETRO / REPORT / EDGE) — the rule is in force while its
  own operating parameters are undetermined, with "no default inferred".
- **GR-005** is owner-confirmed in the Manifest but **never merged into the Matrix** and therefore
  not in force there.

No tool reduces these. But notice their *shape*: this is roughly **ten decisions**, several of them
one-line confirmations, not an unbounded stream. The project is not far from the point where the
human gate stops being the throughput limiter.

### 5.2 Mechanism 2 — authority-decay from having no detector (bigger than it looks)

Re-verified **today**, from TCI's own 2026-09-28 corpus audit — these are *still live*:

| ID | Finding (re-verified 2026-10-06) | Still live? |
|---|---|---|
| V-01 | Register claims to be the cumulative record of **every** ruling; **HD-EIC-09 has 0 occurrences** in it — while **11** entries cite "HD-EIC-01–09 … remain in force" | ✅ **YES** |
| V-02 | HD-TST-01 cited 11 times; no dedicated entry, only a cumulative-state row | ✅ **YES** |
| V-03 | Register's provenance guarantee points to `/archive/human_rulings_source/` — **path does not exist** (verified `Test-Path` = False; never in git history) | ✅ **YES** |
| V-04 | `..._Comparative_Simulation_Specification_v0.1.1.md` cited as existing by **seven** documents; absent from repo **and** from all git history | ✅ **YES** |
| V-05 | Two unrelated governance rules share the ID `GR-004` (Matrix = attribution; Manifest = ruling-context disclosure) | ✅ **YES** |
| V-09 | 69 of 82 documents lack ≥1 GR-004 required metadata field; 20 lack ≥3 | ✅ presumed (no checker exists to refute it) |
| V-07 | Live-queue table disagreed with three other sources on which items were live | ⚠️ superseded by queue collapse (§5.3) |
| V-11 | Authority record existed only as uncommitted working-tree state; 48 untracked `.md` | ➖ **mitigated** (committed 2026-09-29; 0 untracked `.md` today) |

Every one of V-01…V-05 is a **machine-checkable property**:

- *every ID cited anywhere must have a dedicated record somewhere* → catches V-01, V-02;
- *every path asserted to exist must resolve* → catches V-03, V-04;
- *IDs are globally unique and monotone per namespace* → catches V-05;
- *every document carries required frontmatter fields* → catches V-09.

They are live because nothing checks them continuously. An audit found them once, brilliantly —
and then the audit aged and the defects stayed. A detector is what converts "found once" into
"cannot recur". This is the single highest-value thing the TW kernel brings.

### 5.3 Mechanism 3 — the loop has nothing to do while waiting (self-reinforcing)

- Live queue is **0 items** versus the project's own stated policy of 3–5. The last surviving item
  (OQ-012 multi-account *implementation*) is blocked because **you asked for the question to be
  re-explained before ruling** — and that re-explanation has not been produced for **8 days**.
  That is an open request from the owner sitting unfulfilled, not a decision waiting on you.
- The remaining backlog is systematically un-draftable under current rules: ranks 10/13/14 are
  EIC-coupled; rank 11 has "zero evidence anywhere in the corpus"; rank 15 is behind others; ranks
  16/17 await transport and organization rulings.
- **Cycle-14 implementation authorization** (2026-09-20: HD-ITM-02 / HD-BLD-03 / HD-PROF-02 against
  the SWG fork, scaffold only) has **no artifact in the repository after that date** evidencing
  execution; `testbed/` last write is 2026-09-17. Authorised work that was never done is the most
  expensive kind of stall, because it looks like progress in the record.
- **Due-diligence steps 1–5 are recorded as executed** in an evidence log, yet every checkbox in the
  checklist is still unchecked and `HD-TST-01`'s own text says its preconditions are "not yet
  executed". So the precondition state of a HUMAN-LOCKED ruling is **ambiguous on the record** —
  the evidence exists and the tracker denies it.

**The compounding loop:** no queue → nothing drafts itself → no ruling surfaces → nothing files →
nothing drafts. Each iteration further from machine-checkable state than the last.

### 5.4 Root cause statement

> TCI's stall is **not** primarily a decision-making deficit. It is the absence of a deterministic
> substrate that can (a) hold authority state without decay, (b) detect divergence continuously
> instead of audibly, and (c) keep the human queue populated with *cheap, pre-packaged, single-screen*
> decisions. Its own records measure the cost precisely: **~9 days** for given rulings to become
> detectable (GR-006 precipitating incident), **42 days** for one structural boundary statement, and
> **8 days** for one requested re-explanation.

---

## 6. Option analysis

| Criterion | (A) Improve TCI in place | (B) New autonomous repo, migrate | (C) **Transplant kernel into TCI** (recommended) |
|---|---|---|---|
| Preserves the authority chain verbatim (22 rulings, GDD lineage, provenance) | ✅ untouched | ❌ **high risk** — see below | ✅ untouched |
| Fixes the detector gap (V-01…V-05 class) | 🟡 only if you build it | 🟡 still must be built | ✅ kernel + small TCI layer |
| Bound the human's cost (interrupt budget, pre-packed decisions) | ❌ no mechanism | 🟡 same as TW today (unproven) | ✅ policy-bounded, then populated with *your real* 10 open decisions |
| Keeps corpus location stable for humans and tooling | ✅ | ❌ new paths/pointers; every cross-reference in 160 docs is a path-based citation | ✅ |
| Startup cost | 🟢 lowest | 🔴 highest (new repo + migration + re-baselining) | 🟡 moderate (copy kernel; write checks) |
| Reversibility | 🟢 high | 🔴 low once commits move | 🟢 **high** — `.pm/` is additive and deletable |
| Evidence of failure risk | n/a | 🔴 **TCI's own history documents this exact loss class** (below) | 🟢 none beyond added files |
| Removes the human gate | ❌ no option does | ❌ no option does | ❌ no option does — but shrinks it |

### 6.1 Why not (B) — evidence already in TCI's own record

This is the decisive argument, and it is not theoretical:

1. **TCI has already lost documents to a corpus move.** `Open Questions Register` §5a records that a
   prior consolidation gathered files by **prefix-matching**, so documents saved without the
   `TCIndustries_` prefix were **silently absent from the merge with no error raised** — four
   confirmed cases, one of which (the Architecture C Validation Gate) was "entirely absent from the
   merged upload" and had to be reconstructed from conversation context. Its own root-cause note says
   the loss "cannot be verified from a merged export that may itself have already dropped the evidence".
2. **A move reproduces that mechanism.** One hundred sixty documents cross-reference each other by
   path and filename. Any automated harvest during migration inherits the same silent-drop failure
   mode, and this time there is no parallel chat history to reconstruct from.
3. **The new repo would inherit the same human.** EIC would still be blocked on HD-EIC-07, the same
   ten decisions would be open, and the receiver would start at zero evidence bundles and zero
   field data — losing TCI's only asymmetric asset (22 rulings of real human traffic).
4. **Cost asymmetry.** Migration spends effort on logistics; Option C spends the same effort on the
   actual defect — detection — which migration does not deliver either.

> **Rule of thumb.** Move a corpus only when the *location* is the defect. TCIndustries' location is
> fine. Its *location-blind machinery* is missing. Fix the missing thing.

### 6.2 Why (A) alone is insufficient

"Improve the existing infrastructure" today means writing more instructions into more Markdown. But
GR-006 already demonstrates the ceiling of the instruction approach: it was written *because* six
rulings were undetectable in the authority record for nine days, and it has been in force since
2026-09-28 with four of its own slots unanswered — while the five V-class defects it was born from
remain open. **Instructions written after a failure do not execute themselves.** Adding a 42nd rule
to a system with zero instruments adds surface without adding holding power. This is precisely the
failure mode recorded for the *other* reference project and summarised there as
*"an assertion you read past is not a gate."*

### 6.3 Recommendation — Option C, stated precisely

> Keep the TCI repository, its paths, its seven-directory layout, its `AGENTS.md`, its 22 rulings and
> its full document corpus **exactly as they are**. Add a `.pm/` coordination tier beside them
> (same shape as TW's kernel), plus a TCI-specific **document-domain check layer** the kernel does
> not have. The check layer is pointed at the *existing* files; it never rewrites them.

Three properties make this safe:

- **Additive.** New directory, new tests. No existing governance document is edited by the migration.
- **Reversible.** Delete `.pm/` and the check layer; the project is byte-identical to today.
- **Non-authoritative by construction.** The machinery records and detects. It never invents a ruling,
  never promotes a status, never answers an open question. Its powers stop exactly where TCI's
  authority model already draws the line — and that boundary is machine-enforced, which is stronger
  than today's instruction-based boundary.

---

## 7. Recommended implementation (Option C)

Ordered so each phase is independently useful and each has an exit test. Phase gates are **R3**
(two-key: agent proposal + your approval) under TW's own policy convention.

### P0 — Stop-the-line preservation (before anything else)

| # | Action | Exit test |
|---|---|---|
| P0.1 | Content-address the corpus: record SHA-256 for every file under `canonical/`, `governance/`, `proposals/`, `investigations/`, `archive/` and assert stability | `check-corpus` exit 0; any unexplained change is a finding |
| P0.2 | Adopt **verified-write** for the two mutable authority artifacts (`Human_Rulings_Register.md`, `project_memory.md`): write → re-read → compare → refuse success otherwise | A seeded clobber attempt is refused |
| P0.3 | Ban new root files (`.gitignore` already minimal); quarantine the three untracked SQLite files and the 0-byte `fs.statSync(x).isFile()))` artifact — or delete the latter after confirming it is a shell fossil | `check-root` exit 0 |
| P0.4 | Record that the 2,168 MB testbed database is local-only, and move the exclusion from a vendored nested `.gitignore` into project policy | `check-root` documents the exclusion |

**Why first:** this is the cheapest possible insurance, and it is the one thing that makes every later
phase safe to attempt. P0 needs no design decision from you.

### P1 — Transplant the kernel (additive)

Copy TW's `pmk.py` (stdlib-only Python; the `py` launcher is the working form on this host) plus
`policy.yaml`, `views/`, `events/`, `memory/{facts,assumptions,patterns,incidents}/`,
`registry/`, with these TCI-specific adaptations:

- **Policy values set for a design project, not a code project**: dependencies `stdlib` only;
  `autonomy_ceiling: R2`; **any write to `canonical/` or `governance/` is R3 by risk floor**;
  status-token promotion is R3; **invention of any number** remains forbidden and is machine-checked
  rather than instructed (BAL-001).
- Adopt the **interrupt budget** (say 10/week, 15 min/unit) and **72 h default-hold** — then populate
  the queue with its first real entries: HD-EIC-07, the 8 open §36 decisions, GR-006's four slots,
  the GR-005 merge.

| Exit test | Expectation |
|---|---|
| `sentinel` green on a fresh clone | invariants hold after clone, not only locally |
| `check-root` green | after P0.3/P0.4 |
| `views/*` generated with `measured=` stamps | no hand-edited projection exists |
| one `INC-*` auto-filed per sentinel finding | the incident→pattern habit starts at first finding |

### P2 — Build the TCI check layer (the part that must be written new)

Six checks, each with a **must-fail fixture** drawn from a *real, already-recorded* TCI defect. This
is the unusual opportunity here: TCI already has labelled ground truth, so red witnesses do not need
inventing.

| Check | Fires on | Red witness available from history |
|---|---|---|
| `check-rulings` — every `HD-*` ID cited anywhere has a dedicated record | V-01, V-02 | yes (HD-EIC-09, HD-TST-01 today) |
| `check-refs` — every referenced path/anchor resolves | V-03, V-04 | yes (nonexistent archive path; phantom v0.1.1 spec) |
| `check-id-unique` — IDs unique and monotone per namespace | V-05 (GR-004 collision) | yes |
| `check-registers` — OQ statuses equal across GDD §31, OQR §2 and the Queue Tracker; invariant approval set equal across three registers | V-07 | yes (the composition disagreement) |
| `check-metadata` — required frontmatter present per GR-004 | V-09 (69/82 docs) | yes |
| `check-boundaries` — no numeric values introduced outside an authorising ruling; no non-canonical concept resurrected from the explicit list | BAL-001, HD-EIC-08 | partially |

**Expected effect, and the honest caveat.** Running this layer on today's tree **will go red** —
V-01…V-05 are live. That is the point and also the trap: a gate that is permanently red becomes
wallpaper. Adopt therefore the practice recorded from the other reference project as
*"declared known-red instead of implied green"*: each check carries its open findings inline, with
IDs, so the red set is visible, dated and shrinking rather than hidden. **Do not let the layer be
introduced as green.**

### P3 — Change the shape of the ask (this is the part that fixes "I am slow")

Filing five dependent documents per ruling by hand is why rulings sat nine days undetected, and it
is why ruling feels expensive. Replace it with a machine-generated **ruling packet**:

1. One screen. The single question, in plain language.
2. Exactly the options, each with its consequences already computed against the 28 invariants.
3. **A recommended default**, clearly marked as a proposal with no authority.
4. The ready-to-file register entry, pre-rendered — so approving means *paste*, not *compose*.
5. On approval: one command executes **all five dependent-document acts atomically** and verifies
   each landed (GR-006's core obligation, mechanical at last).

Target: **cost per ruling ≤ 15 minutes**, batched ≤ 10 interrupts/week. That converts your ten open
decisions from an open-ended obligation into roughly two hours of bounded work, spread over two weeks.

**Companion act, worth doing before P3 finishes:** fulfil the outstanding request. Produce the OQ-012
implementation-policy re-explanation you asked for on 2026-09-28. It is the item that emptied the
queue, and it is owed regardless of which option is chosen.

### What NOT to import from TW

1. **Its emptiness as a virtue.** Declarative-but-unread policy was already flagged there as a defect.
   In TCI every policy value should be read by at least one check, or deleted.
2. **Bundle-oriented completion as the primary workflow.** TCI is not a code-build project; forcing
   every act through `init-bundle → run → submit` would ritualise work that should be a one-line check.
3. **A premature merge queue.** Until TCI parallelises multiple workers, leases/fencing are needed for
   safety but a merge queue is ceremony.
4. **Any change to `AGENTS.md` authority semantics.** TCI's status model is coherent and, on
   2026-09-28, measured violation-free. Leave it alone; put mechanics underneath it.

---

## 8. Answering the question directly

**Should you improve TCI's existing infrastructure, or create a new autonomous project and move?**

> **Improve — in place — by adding machinery rather than rewriting documents. Do not move the corpus.**
> Create a new repository only if you later decide to *rename or re-home* the project for reasons
> unrelated to this diagnosis; even then, move by `git clone`/history-preserving copy, never by
> re-harvesting document sets.

**Conditions under which this recommendation should be revisited:**

| Signal | Meaning | Action |
|---|---|---|
| After P0–P1 the check layer finds **nothing** new beyond V-01…V-05 | The defects were one-off, not systemic | Stop; revert to lightweight improvement (A) |
| You find yourself wanting the GDD **rewritten** rather than extended | The complaint is about content, not infrastructure | Do not migrate — write a scoped v1.2 *status-patch*, which is already the house style |
| Two or more agents will genuinely work TCI **concurrently** | Concurrency is now the binding constraint | Then the full TW lease/fence/merge machinery earns its cost; adopt it wholesale |
| The EIC stays blocked another 60 days | The human gate really is the limiter | Then **de-scope deliberately**: rule that EIC is deferred and unblock ranks 10/13/14 explicitly. Recordable in one ruling; no migration required |

**The one thing worth saying plainly.** Moving to a new autonomous system will not make you decide
faster, because nothing can. What it will do is guarantee that the moment you *do* decide, every
dependent document updates in once command instead of drifting for nine days — and that no decision
you have already made has quietly disappeared from the record. Given that four relevant defects are
open in the record **right now** (V-01…V-04, all re-verified today), that difference is worth having
before the next ruling rather than after it.

---

## 9. Limitations and honest caveats

- Single-day static snapshot; both trees drift (TCI by elapsed time, TW by active commits — it was
  committed within hours of this survey).
- **No tool was executed.** Gate results for TW are quoted from its README (dated 2026-10-06) and
  its own decision record; they were not independently re-run here. Re-run them before relying on them.
- TW's own record states it has **zero evidence bundles and one never-claimed task**. Its machinery is
  self-tested but has carried almost no real traffic. Its advantage here is *design plus tests*, not
  demonstrated throughput — whereas TCI's advantage is demonstrated throughput with no machinery.
  Anyone reading this should discount TW's maturity accordingly.
- V-09 is presumed live because no mechanism exists to refute it; that is the finding, not certainty.
- Whether the phantom `..._Specification_v0.1.1.md` exists outside the repository is a question for
  you, deliberately not adjudicated here (the 2026-09-28 audit made the same call). Same for which
  `GR-004` is authoritative.
- No design opinion is expressed. This report does not assess whether any ruling was *correct*, and
  it does not propose any EIC boundary, number, or architecture.
- This document has no independent assessor yet (`assessor_llm: []`). Before it is relied upon for a
  structural decision, it should be assessed by a different model or by you.

---

## Appendix A — Reproducible measurements

Repository facts (either repository, PowerShell):

```text
git -C <repo> log --date=short --pretty='%h|%ad|%s'          # history and cadence
git -C <repo> log -1 --date=iso --pretty='%H%n%ad%n%s'       # last commit + timestamp
git -C <repo> rev-list --left-right --count origin/main...main   # ahead/behind
git -C <repo> status --porcelain                             # dirty tree
git -C <repo> ls-files <path>                                # is a large artifact tracked?
git -C <repo> log --all --oneline -- '<filename>'            # did it ever exist in history?
git -C <repo> check-ignore -v <path>                         # which ignore rule applies
```

Corpus and defect re-verification (this assessment):

```text
Get-ChildItem <repo> -Recurse -File -Filter *.md | Measure-Object      # corpus counts
Select-String -Path <file> -Pattern 'HD-[A-Z]+-\d+' -AllMatches |      # ID census
  ForEach-Object { $_.Matches } | ForEach-Object { $_.Value } | Group-Object
Test-Path <asserted path>                                             # V-03 style existence check
Select-String -Path <file> -Pattern '\[[ xX]\]'                       # checklist completion state
```

Encoding note (already recorded in the 2026-09-28 audit, re-confirmed here): PowerShell 5.1 can
mis-decode UTF-8; several TCI files contain ¶, α/β/γ and smart punctuation. Read via
`[IO.File]::ReadAllText($path, [Text.Encoding]::UTF8)` when byte accuracy matters.

## Appendix B — Mapping: recorded TCI defect → candidate machine check

| Defect (source) | Would have been caught by | Source of red witness |
|---|---|---|
| V-01 ruling cited but not recorded | `check-rulings` | today's HD-EIC-09 state |
| V-02 same class, second instance | `check-rulings` | today's HD-TST-01 state |
| V-03 provenance points to missing path | `check-refs` | register line 30 |
| V-04 phantom document cited 7× | `check-refs` | Changelog, OQR, Manifest, Consensus Plan |
| V-05 two rules share `GR-004` | `check-id-unique` | Matrix vs Manifest |
| V-07 registers disagree on live queue | `check-registers` (generated views replace hand tables) | Cycle-13 fix set D1–D4 |
| V-09 metadata missing in 69/82 docs | `check-metadata` | any doc lacking frontmatter |
| V-11 authority record only in working tree | `check-corpus` + P0.1 hashing | 2026-09-28 working-tree state |
| Six rulings undetectable for ~9 days (GR-006 precipitating incident) | atomic multi-target filing + `check-rulings` | Section C→D filing gap |
| Checklist unchecked though evidence log exists | `check-registers` | checklist v0.2 vs evidence log |
| Authorised Cycle-14 work with no artifact | lease/claim expiry + no terminal state without evidence | Cycle-14 authorization record |

---

*End of report. Author: OrcaRouter Auto (`orcarouter/auto`), 2026-10-06. ADVISORY — no authority;
no status changed; no file amended by this assessment other than this report.*
