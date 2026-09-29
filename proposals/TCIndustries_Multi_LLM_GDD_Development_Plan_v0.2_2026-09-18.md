# TCIndustries — Multi-LLM GDD Development Plan (v0.2)

**Status:** PROPOSED. No authority. Nothing here promotes, amends, supersedes, or reclassifies any Ruled or Canonical material. It proposes a *workflow* for rapidly developing the Master GDD with multiple LLMs; it designs no game content and resolves no design decision. Becomes operative only on project-owner approval of the plan and of each subsequent step it schedules.
**Version:** 0.2 (revision of v0.1, dated 2026-09-18)
**Date:** 2026-09-18 (ISO)
**Author:** OpenCode (big-pickle)
**Assessor:** None — single-session authoring; no independent review pass has been performed (per GR-004, recorded as such rather than guessed).
**Controlling references:** `AGENTS.md`; `README.md`; `canonical/gdd/TCIndustries_Master_GDD_v1.1.1_Status_Patch.md`; `governance/TCIndustries_Human_Rulings_Register.md` (HD-EIC-01–08, HD-EIC-09, HD-GDD-01, HD-TST-01 in force); `governance/TCIndustries_Authority_Provenance_Reconciliation_Matrix.md` (GR-001–005); `proposals/TCIndustries_Open_Questions_Register.md` (OQ-001–017); `investigations/consolidation-audit/README.md` (multi-LLM fan-out precedent); `governance/project_memory.md` (2026-09-17 duplicate-pass incident); `proposals/TCIndustries_Multi_LLM_GDD_Development_Plan_v0.1_2026-09-18.md` (the document this revises — retained verbatim as the historical record).
**Revision delta from v0.1:** (1) §3 orchestration option replaces Orca with OpenChamber, the installed, running orchestration layer; (2) new OpenChamber feature→pipeline mapping (§3.2); (3) new model/role matrix grounded in verified provider credentials (§3.3); (4) §5 re-cast as a "Run This Week" demo-sprint module with explicit day-scope; (5) §6 tooling reference updated; (6) §7 adds orchestration-layer risks. All other sections (working principles, intake packet contract, gates, file conventions, reconciliation loop, rollout order) are carried forward substantively unchanged.

---

## 1. Working Principles

1. **LLMs produce evidence; the project owner rules.** A drafter's output is advisory. Agreement between models is evidence of convergence, never a decision.
2. **One record, one repository.** All draft material enters through a single intake contract and is stored by the coordinator. The Human Rulings Register is the only place decisions are recorded.
3. **Fan out for divergence, converge by audit.** Same task to multiple models → reconcile by surfacing conflicts, not by silent merging (the `consolidation-audit/` pattern: seven independent LLM audits of one corpus, then a consensus plan).
4. **Full-context rule.** No drafter works from a partial snapshot; drafter LLMs must be able to see the governing-status excerpt for the material they touch. If context is missing, the drafter must say so, not guess. This directly counters the HD-GDD-01 root cause (a partial-context LLM session whose decision *outcome* was copied forward while its *process* was lost). A richer orchestration context does not relax this rule.
5. **Zero-tolerance for duplicate passes and unauthorised claims** (the 2026-09-17 incident: a duplicate Evidence Reconciliation Pass that made unauthorised design claims and temporary GDD edits). Pre-flight check before any new pass.
6. **Structure moves; numbers wait.** Structural proposals may advance even where numeric detail is gated. Numeric tuning of economic or progression systems requires explicit human authorisation (BAL-001, HD-EIC-08). EIC architectural work stays gated (HD-EIC-07); Architecture C stays retired (HD-EIC-05).

## 2. Standing Pipeline

### Roles

| Role | Entity | Authority |
|---|---|---|
| Ruler | Project owner (Tiwa) | Sole design authority; decides on decision briefs; supplies structural boundaries where required |
| Coordinator | OpenCode (big-pickle) | Intake, governance check, storage, audit, decision-brief builder, ruling recorder; never a designer |
| Drafter LLMs | Any model/agent behind the drafter endpoints (§3) | Advisory only; interchangeable behind the intake contract |
| Reviewer LLMs | Any model other than the drafter of record | Produce assessment annotations only |

### Intake packet contract (required of every drafter output)

- Author model + version
- Date
- Session type (chat / agent / API)
- Self-declared status (must be PROPOSED / ASSUMPTION / ADVISORY — never LOCKED)
- Purpose and scope
- Sources consulted (file paths, versions, GDD item IDs such as RES-005, CRFT-001, ITM-002)
- Claims with references
- Explicit list of what the packet does *not* resolve
- GR-004 attribution block

A packet missing a required field is returned, not stored.

### Gates

1. **Intake gate** — coordinator verifies packet completeness + governance compliance before filing.
2. **Consistency gate** — coordinator cross-checks the packet against the rulings register, invariants, and statuses; conflicts are recorded, never silently fixed.
3. **Ruling gate** — a one-page decision brief is delivered to the owner; nothing advances without a ruling.
4. **Promotion gate** — status changes only via a recorded human ruling; the coordinator executes the strip and notes it in the changelog.

### File conventions (GR-005)

| Material | Destination | Filename pattern |
|---|---|---|
| Draft design material / candidates | `proposals/` | `TCIndustries_<Gap>_Options_<Model>_vX_<date>.md` |
| Assessments / audits / analysis | `investigations/` | `TCIndustries_<Topic>_<date>.md` |
| Storage / repo events | `governance/project_memory.md` | Bulleted event line (existing pattern) |
| Owner rulings | `governance/TCIndustries_Human_Rulings_Register.md` | `HD-<Area>-<NN>` |
| Gap status pointers | `proposals/TCIndustries_Open_Questions_Register.md` | Status cell updated to point at ruling |

Drafter LLMs never write to the repository directly. The coordinator writes on their behalf at intake.

## 3. Parallel Drafting

### 3.1 Fan-out procedure

1. **Slice**: take one open gap; split into sub-questions, or send the same gap to multiple models for divergence.
2. **Spawn**: launch drafter sessions — one per gap per model — with the intake contract, the full-context packet, the 2026-09-17 incident note, and a finish-line goal (Section Goal), each in its own git worktree/branch.
3. **Collect**: outputs are returned as packets; the coordinator runs the intake gate and files them.
4. **Select**: the owner (not the pipeline) picks the leading candidate(s), aided by the decision brief.

### 3.2 Orchestration — OpenChamber (primary)

**Change from v0.1:** v0.1 named Orca as the full-parallel orchestration option and OpenCode subagents as the minimum. OpenChamber (`openchamber.dev`, MIT; runs on the OpenCode SDK) is now installed and running in this environment and **replaces Orca** as the parallel-orchestration recommendation. It covers Orca's worktree fan-out and adds fusion, goal-driven persistence, scheduling, guided diff review, and cross-device supervision. OpenCode subagents and Ark chat-app handoff remain as fallback/adjunct endpoints.

OpenChamber feature → pipeline-step mapping:

| OpenChamber feature | Pipeline role | Notes |
|---|---|---|
| **Multi-run** (same task to up to 5 models, each in own session, optionally own git worktree) | Fan-out (§3.1 step 2): one gap → N drafter sessions on `feat/<gap>-draft/<model>` | Worktrees and branches isolate drafting from canonical `main`; no two drafters collide |
| **Session Goals** | Per-pass finish line: the drafter keeps iterating, even with the app closed, until it can return an intake-complete packet or declares itself blocked | Goal string is the intake contract plus the gap brief — not an open-ended instruction |
| **Fusion** (seed a new session with the strongest patches from several runs) | Convergence (§4): combine structurally-strongest sections across N runs into a consolidated packet | Drafting-layer merge only. The *design* winner is decided by the owner via the decision brief; fusion never confers design authority |
| **Changes Walkthrough** (AI-guided diff narrative) | Ruler-facing review of the coordinator's consolidation before the ruling gate | Turns a multi-file consolidation into ordered, explained steps |
| **Scheduled / cron passes** (+ Session Goals) | Nightly consistency, coverage-matrix, and numeric-touch screens (§4) | Audits, screens, and files only; never rules; never touches `canonical/` |
| **Branch / worktree review** | Keeper-draft review per branch; coordinator merges to `main` only after intake | Merge discipline rests with the coordinator alone |

### 3.3 Model/role matrix (grounded in verified credentials)

Credentials verified in this environment: OpenCode authenticates OpenCode Zen, Google, Groq, Nvidia, Cerebras, OpenRouter, and Kilo Gateway; the owner additionally holds BytePlus Ark, Anthropic, OpenAI, Gemini, and Grok access. **Every endpoint below must be catalog-verified at invocation time** (`opencode models` / `arkcli models`) — this matrix is the target assignment, not a promise of availability.

| Role | Endpoints (ordered) | Why |
|---|---|---|
| Drafter fan-out (gap X) | OpenRouter-hosted Claude · OpenAI · Gemini · Grok (up to 5 in parallel) | Max model diversity to provoke divergence, which the pipeline surfaces rather than silences; OpenRouter exposes all five major families under one key |
| Fast fallback drafter | OpenCode Zen / Groq / Cerebras / Nvidia | Cheap structural-only seconds when fan-out cost is a concern |
| Reviewer (annotates another's packet) | Any endpoint **≠ drafter of record** | Cross-model critique per §4 |
| Ark critique surface | `arkcli` chat / understand on BytePlus models | Bonus endpoint; output routed through the same intake contract |
| Coordinator | This OpenCode session | Intake, gates, consistency, decision briefs, ruling recording |

### 3.4 Drafter endpoints, in order of availability

1. **OpenChamber Multi-run** (works today; installed and running) — the full parallel configuration.
2. **OpenCode subagents** (works today, no orchestration dependency) — the minimum runnable fallback.
3. **Ark models via arkcli** and chat-app handoff (Claude/ChatGPT/Grok surfaces) — for brainstorming and critique passes behind the same intake contract.

Worktrees beyond OpenChamber's own are optional acceleration and are documented, not assumed; anything repo-changing is surfaced for owner approval first.

## 4. Reconciliation & Audit Loop

- **Cross-model review**: a model other than the drafter of record reviews each packet; annotations stored beside it.
- **Consistency checks**: statuses vs. register; contradiction surfacing (never silent fixes).
- **Numeric-touch screen**: grep for numbers entering gated areas (economy, progression tuning, EIC) and flag or strip them before they reach a brief. Automatable as a scheduled OpenChamber cron pass.
- **Convergence**: 2-of-3 structural agreement marks a leading candidate — evidence, not decision. Fusion is an allowed drafting-layer tool for assembling convergence material, never a decision mechanism.
- **Coverage matrix**: updated per pass, mirroring the consolidation-audit coverage matrix; updatable via scheduled pass.

## 5. Demo Sprint — First Run ("Run This Week")

Scope: carry 2–3 open gaps to ruling-ready within the week, using tools already installed and credentialed.

### Scoping rule

Pull open items from `TCIndustries_Open_Questions_Register.md` (OQ-001–017 and GDD §36). Exclude EIC-linked and gated items (OQ-001, OQ-002, OQ-009, OQ-012, OQ-017) and numeric-first items. Candidate first targets (structural, single-section-anchored, non-EIC):

| Gap | Anchor | Fit |
|---|---|---|
| **OQ-010** — Durability, decay, repair model | ITM-002 | Well-bounded; Items/Services/Economy; lowest conflict surface |
| **OQ-007** — City formation, governance, taxation/maintenance | BLD-003 | Society-core; structural model needs no numbers |
| **OQ-016** — Vendor & retail mechanics | RET-002 / RET-003 | Structural framing; percentage fees deferred to the numeric gate |
| *(Alternate)* **OQ-008** — Combat scope & risk model | CMBT-002 | Prefer only if combat/testbed momentum is wanted |

### Execution steps

1. **Pre-flight** — coordinator greps `proposals/` + `investigations/` + `project_memory.md` for each candidate gap ID; any hit stops the pass (2026-09-17 rule).
2. **Scoping pass** — coordinator reads the anchor sections + register, re-verifies the gaps are genuinely open against the live GDD, produces the 3-candidate shortlist with a one-page decision brief.
3. **Owner confirms** 2–3 gaps.
4. **Multi-run fan-out** — one OpenChamber session folder per gap; dispatch the same task (intake contract + full-context excerpt + incident note + finish-line goal) to 3–5 models in parallel worktrees.
5. **Intake gate** — coordinator verifies packet completeness; incomplete packets are returned, not stored; complete packets are filed under `proposals/`.
6. **Consistency + numeric-touch screen** — statuses vs. register; numbers flagged in gated areas (surfaced, never silently removed).
7. **Decision brief per gap** — options, convergence evidence (2-of-3 structural agreement counts only as evidence), conflicts, and exactly what the owner must rule — never a dressed-up recommendation.
8. **Owner rules** — coordinator records a new `HD-<Area>-<NN>` entry, updates OQR status pointers, marks the GDD status-patch line, files provenance.
9. **Sprint close** — `project_memory.md` event line + coverage-matrix update.

Target: 2–3 gaps ruling-ready by week's end, starting with ~3 drafters per gap and scaling to 5 where OpenRouter coverage and cost permit.

## 6. Tooling Reference

- **OpenChamber** (primary orchestration): installed and running in this environment. Multi-run for fan-out, Fusion for convergence material, Session Goals for per-pass finish lines, cron for nightly screens, Changes Walkthrough for consolidation review. Requires Node.js 22+ for CLI use; the desktop app bundles OpenCode.
- **Subagent fan-out** (fallback): OpenCode `task` tool with the intake contract embedded in the agent prompt.
- **Orca**: retired from this plan. Retained only as a historical note — v0.1 named it; v0.2 replaces it with OpenChamber. Any future use requires a fresh proposal.
- **Ark drafter** (optional): verify catalog (`arkcli models`) before invocation; route drafter output into the intake contract.
- **Worktrees** (native to OpenChamber Multi-run; manual form if not using it): `git worktree add ../tcindustries-<name> -b feat/<gap>-draft`; merge back on a clean checkout; remove with `git worktree remove`.
- **Pre-flight duplicate check**: grep `proposals/` + `investigations/` + `project_memory.md` for the gap ID before authoring any pass.

## 7. Risks & Safeguards

- **Partial-context decision loss** → full-context packets mandatory even with large orchestration contexts; drafter declares context gaps; only the register records decisions.
- **Unauthorised claims / duplicate passes (2026-09-17 pattern)** → pre-flight check; drafter agents never touch `canonical/`; the incident note is included in every intake prompt.
- **Fusion mistaken for decision** → Fusion is a drafting-layer tool. The design winner is always the owner's, via the decision brief. Stated in the Session-Goal text and the fusion brief.
- **Worktree/session sprawl** → one branch per gap per model; merge discipline rests solely with the coordinator; pre-flight check prevents collisions; short-lived sessions.
- **Scheduled passes exceeding their mandate** → cron sessions audit, screen, and file only; they never rule and never write to `canonical/`; goals are scoped to auditable outcomes.
- **Catalog drift (provider/model availability)** → every invocation re-verifies the live catalog before fan-out; unverified endpoints are marked unavailable, not guessed.
- **Numeric creep into gated areas** → automated screen; structural-only default.
- **Status drift** → storage ≠ promotion; promotion only via recorded human ruling.
- **Attribution gaps** → GR-004 enforced at intake; missing attribution recorded as UNVERIFIED, not guessed.

## 8. Rollout Order

1. Owner approves this plan (and thereby the demo-sprint scope method).
2. Coordinator produces the 3-candidate shortlist brief (re-verifying gaps against the live GDD).
3. Owner confirms 2–3 gaps.
4. Multi-run drafting pass runs (OpenChamber first; subagents/Ark added on owner's signal).
5. Review + consistency passes.
6. Decision briefs → owner rulings → coordinator records and files.

---

## Explicit Non-Findings

- No DEC / HD ruling is assigned or inferred by this document.
- No rule is invented, amended, promoted, or superseded. In particular, Orca is retired as a *plan recommendation* only — no prior ruling about Orca existed to rescind.
- No GDD status is changed; all TBDs remain TBD; CRFT-006 stays PROPOSED.
- No EIC architectural work, residual-channel invention, or numeric tuning is proposed.
- This is a proposed *workflow*; the design content it will carry is whatever the owner subsequently approves.

*End of draft. To enact: owner reviews and approves this plan and its sequenced steps (in-session or as amended); each further step is then executed only with the owner's sign-off at its gate.*