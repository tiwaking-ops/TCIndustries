# TCIndustries — Multi-LLM GDD Development Plan (v0.1 draft)

**Status:** PROPOSED. No authority. Nothing here promotes, amends, supersedes, or reclassifies any Ruled or Canonical material. It proposes a *workflow* for rapidly developing the Master GDD with multiple LLMs; it designs no game content and resolves no design decision. Becomes operative only on project-owner approval of the plan and of each subsequent step it schedules.
**Version:** 0.1
**Date:** 2026-09-18 (ISO)
**Author:** OpenCode (big-pickle)
**Assessor:** None — single-session authoring; no independent review pass has been performed (per GR-004, recorded as such rather than assumed).
**Controlling references:** `AGENTS.md`; `README.md`; `canonical/gdd/TCIndustries_Master_GDD_v1.1.1_Status_Patch.md`; `governance/TCIndustries_Human_Rulings_Register.md` (HD-EIC-01–08, HD-EIC-09, HD-GDD-01, HD-TST-01 in force); `governance/TCIndustries_Authority_Provenance_Reconciliation_Matrix.md` (GR-001–005); `proposals/TCIndustries_Open_Questions_Register.md` (OQ-001–017); `investigations/consolidation-audit/README.md` (multi-LLM fan-out precedent); `governance/project_memory.md` (2026-09-17 duplicate-pass incident).

---

## 1. Working Principles

1. **LLMs produce evidence; the project owner rules.** A drafter's output is advisory. Agreement between models is evidence of convergence, never a decision.
2. **One record, one repository.** All draft material enters through a single intake contract and is stored by the coordinator. The Human Rulings Register is the only place decisions are recorded.
3. **Fan out for divergence, converge by audit.** Same task to multiple models → reconcile by surfacing conflicts, not by silent merging (the `consolidation-audit/` pattern: seven independent LLM audits of one corpus, then a consensus plan).
4. **Full-context rule.** No drafter works from a partial snapshot; drafter LLMs must be able to see the governing-status excerpt for the material they touch. If context is missing, the drafter must say so, not guess. This directly counters the HD-GDD-01 root cause (a partial-context LLM session whose decision *outcome* was copied forward while its *process* was lost).
5. **Zero-tolerance for duplicate passes and unauthorised claims** (the 2026-09-17 incident: a duplicate Evidence Reconciliation Pass that made unasuthorised design claims and temporary GDD edits). Pre-flight check before any new pass.
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

### Fan-out procedure

1. **Slice**: take one open gap; split into sub-questions, or send the same gap to multiple models for divergence.
2. **Spawn**: launch drafter sessions — one per gap per model — with the intake contract and the full-context packet.
3. **Collect**: outputs are returned as packets; the coordinator runs the intake gate and files them.
4. **Select**: the owner (not the pipeline) picks the leading candidate(s), aided by the decision brief.

### Orchestration option — Orca

Orca (`github.com/stablyai/orca`) is an agent development environment that runs multiple CLI agents side-by-side, each in its own isolated git worktree, with a "fan one prompt across N agents → compare → merge the winner" workflow. OpenCode is a supported agent. Using it makes the worktree fan-out mechanical:

- One worktree per gap per agent; each draft committed to its own branch (`feat/<gap>-draft/<agent>`).
- Owner-approved keeper drafts are merged to `main` by the coordinator only, after intake.
- **Consent caveat:** Orca's "merge the winner" is a git/drafting-layer decision. The *design* winner is decided by the owner via the decision brief. No agent output gains design authority from being merged.
- Parallel worktrees keep the working `main` (currently ahead of origin with uncommitted testbed changes) isolated from drafting activity.

Drafter endpoints, in order of availability:
1. **OpenCode subagents** (works today, no external dependencies) — the minimum runnable configuration.
2. **Orca** orchestrating OpenCode and/or other CLI agents (Claude Code, Codex, Grok, etc.) on the owner's own subscriptions — the full parallel configuration.
3. **Ark models via arkcli** and chat-app handoff (Claude/ChatGPT/Grok surfaces) — available for brainstorming and critique passes behind the same intake contract.

Worktrees beyond Orca's own are optional acceleration and are documented, not assumed; anything repo-changing is surfaced for owner approval first.

## 4. Reconciliation & Audit Loop

- **Cross-model review**: a model other than the drafter of record reviews each packet; annotations stored beside it.
- **Consistency checks**: statuses vs. register; contradiction surfacing (never silent fixes).
- **Numeric-touch screen**: grep for numbers entering gated areas (economy, progression tuning, EIC) and flag or strip them before they reach a brief.
- **Convergence**: 2-of-3 structural agreement marks a leading candidate — evidence, not decision.
- **Coverage matrix**: updated per pass, mirroring the consolidation-audit coverage matrix.

## 5. Demo Sprint — First Run (carry 2–3 open gaps to ruling-ready)

### Scoping rule

Pull open items from `TCIndustries_Open_Questions_Register.md` (OQ-001–017 and GDD §36). Exclude EIC-linked and gated items (OQ-001, OQ-002, OQ-009, OQ-012, OQ-017) and numeric-first items. Candidate first targets (structural, single-section-anchored, non-EIC):

| Gap | Anchor | Fit |
|---|---|---|
| **OQ-010** — Durability, decay, repair model | ITM-002 | Well-bounded; Items/Services/Economy; lowest conflict surface |
| **OQ-007** — City formation, governance, taxation/maintenance | BLD-003 | Society-core; structural model needs no numbers |
| **OQ-016** — Vendor & retail mechanics | RET-002 / RET-003 | Structural framing; percentage fees deferred to the numeric gate |
| *(Alternate)* **OQ-008** — Combat scope & risk model | CMBT-002 | Prefer only if combat/testbed momentum is wanted |

### Execution steps

1. **Scoping pass** — coordinator reads the anchor sections + register, verifies the gaps are genuinely open, produces a 3-candidate shortlist with a one-page decision brief.
2. **Owner confirms** 2–3 gaps.
3. **Drafting pass** — 2–3 endpoints per gap, structural-only, intake packets, filed under `proposals/`.
4. **Cross-review pass** — different-model annotations per draft.
5. **Consistency + numeric-touch screen.**
6. **Decision brief per gap** — options, convergence evidence, conflicts, and exactly what the owner must rule (never a dressed-up recommendation).
7. **Owner rules** — coordinator records a new `HD-<Area>-<NN>` entry, updates OQR status pointers, marks the GDD status-patch line, files provenance.
8. **Sprint close** — `project_memory.md` event line + coverage matrix update.

## 6. Tooling Reference

- **Subagent fan-out**: OpenCode `task` tool with the intake contract embedded in the agent prompt.
- **Orca** (optional accelerator): install from https://onorca.dev/download; run agents with the owner's own subscriptions; OpenCode is a supported agent; orphan CLI (`orca worktree create`, `snapshot`, etc.) is scriptable.
- **Ark drafter** (optional): verify catalog (`arkcli models`) before invocation; route drafter output into the intake contract.
- **Worktrees** (optional, if not using Orca): `git worktree add ../tcindustries-<name> -b feat/<gap>-draft`; merge back on a clean checkout; remove with `git worktree remove`.
- **Pre-flight duplicate check**: grep `proposals/` + `investigations/` + `project_memory.md` for the gap ID before authoring any pass.

## 7. Risks & Safeguards

- **Partial-context decision loss** → full-context packets mandatory; drafter declares context gaps; only the register records decisions.
- **Unauthorised claims / duplicate passes (2026-09-17 pattern)** → pre-flight check; drafter agents never touch `canonical/`; the incident note is included in every intake prompt.
- **Numeric creep into gated areas** → automated screen; structural-only default.
- **Status drift** → storage ≠ promotion; promotion only via recorded human ruling.
- **Attribution gaps** → GR-004 enforced at intake; missing attribution recorded as UNVERIFIED, not guessed.

## 8. Rollout Order

1. Owner approves this plan (and thereby the demo-sprint scope method).
2. Coordinator produces the 3-candidate shortlist brief (re-verifying gaps against the live GDD).
3. Owner confirms 2–3 gaps.
4. Drafting pass runs (OpenCode subagents first; Orca/Ark/chat added on owner's signal).
5. Review + consistency passes.
6. Decision briefs → owner rulings → coordinator records and files.

---

## Explicit Non-Findings

- No DEC / HD ruling is assigned or inferred by this document.
- No rule is invented, amended, promoted, or superseded.
- No GDD status is changed; all TBDs remain TBD; CRFT-006 stays PROPOSED.
- No EIC architectural work, residual-channel invention, or numeric tuning is proposed.
- This is a proposed *workflow*; the design content it will carry is whatever the owner subsequently approves.

*End of draft. To enact: owner reviews and approves this plan and its sequenced steps (in-session or as amended); each further step is then executed only with the owner's sign-off at its gate.*