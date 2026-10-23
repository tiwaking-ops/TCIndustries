# SWG Testbed — Post-Phase-10 Program Proposal (Decision Menu + Roadmap) v0.1

**Status:** PROPOSED plan. No authority. No design decision. No code until the §9
decisions are answered by the project owner.

**Author / Assessor:** Buffy — Codebuff agent (`codebuff/freebuff`) via Freebuff,

**Provenance note (GR-007.3, added 2026-10-18):** This author field names no separate model
name or version, and on that basis does not satisfy GR-007.3 on its face. The attribution
stands as originally authored and has been assessed as correct by the human owner
(ruling 2026-10-18). This note does not alter the attribution, does not re-attribute the
document, and confers no authority. No reason for the absence of model disclosure is
recorded here.
**Assessor (llm):** HY4 Preview (`orcarouter/tencent/hy4-preview-free`) — 2026-10-18, per
explicit owner authorization.
**Proposal v0.1**, 2026-09-17 — research, planning, and decision structuring only;
all authority decisions remain human. Prepared from an owner interview (decision
menu + numeric asks included + multi-phase roadmap with first phase detailed).

**Responds to:** completion of Phase 10 (Balance, Telemetry & Polish) — verified
2026-09-17 (`CHAIN:ALL-GREEN`, `PHASE10_EXIT:0`, commits `f9852ec`…`2730c9a`;
proposal v0.2 §10 EXECUTED record).

**Controlling references:** `AGENTS.md`; `proposals/TCIndustries_Manifest.md`;
`canonical/gdd/TCIndustries_Master_GDD_v1.1.1_Status_Patch.md`;
`proposals/TCIndustries_SWG_Testbed_Reuse_Proposed_Ruling.md` (HD-TST-01);
`proposals/SWG_Phase10_Balance_Telemetry_Polish_Proposal_v0.2.md` (§9/§10);
`investigations/SWG_Phase10_Implementation_Verification_Report_2026-09-17.md`;
`investigations/OpenCode_Project_Documentation_Report_2026-09-17.md` (reusability
split: ~30% direct / ~50% adaptable / ~20% not usable).

---

## 1. Status discipline (read first)

- This proposal is **PROPOSED**. Nothing here is canon, and nothing here changes
  any status of any other document.
- **No code, tests, or commits are authorized by this document.** Everything in
  §3–§8 is contingent on explicit owner answers in §9.
- The Phase 10 result this proposal builds on is **EVIDENCE**
  (PROTOTYPE/EVIDENCE per `AGENTS.md`), not design authority for TCIndustries.
- All numeric values referenced are either already-committed fork constants
  (PROTOTYPE) or HISTORICAL predecessor-GDD bands; this proposal sets no values.
- Per the no-automatic-successor-architecture boundary: this proposal does not
  *assume* a next phase exists — it structures the choice of whether one does.

## 2. Situation after Phase 10 (verified, this repository)

1. **Authorized scope fully closed.** Ten testbed phase scopes (0–10) implemented,
   each with a green end-to-end suite on fresh DBs; single verified code state at
   commit `2730c9a`. No open work remains inside any authorization to date.
2. **No successor exists by design.** The predecessor GDD's ladder ends at
   Phase 10 (HISTORICAL). No proposal, ruling, or authorization defines anything
   after it. Under the governance boundaries, initiation is an owner decision.
3. **The working tree is not fully provenance-complete.** ~65 pre-existing paths
   (Phase 3–9 sources incl. the `db.go` HAM repair, both Phase 10 proposals,
   investigations, Godot client) remain uncommitted — the Phase 10 commits
   deliberately touched only Phase 10 work, per the owner's commit-scope
   instruction. Future verify-don't-claim diffs are incomplete until adopted.
4. **Three flagged evidence gaps** (Phase 10 verification report §5):
   heal-event end-to-end (needs an incap→revive scenario), port-8080 hardcoding
   in phase 1–9 clients, mission-lifecycle counters (needs row-retention decision).
5. **Two pending numeric/decision items** (Phase 10 proposal §6 + verification
   report): the B4 service-XP options paper (adopt / decline / defer — default:
   flat `[PROVISIONAL]` rates stand), and the OUT-OF-BAND combat-scale finding
   (fork unarmed 5–15 vs HISTORICAL §9.7 band 50–150 — recorded, no fix sought).
6. **Human-pass items:** Proxy 4 artifact review — **DEFERRED** (2026-10-18); Historical Accuracy proxy — **formal questionnaire authored** (`proposals/SWG_Historical_Accuracy-v1.md` v1.0), pending human answers
   (§6.4 human-pass-only); HD-TST-01 draft traceability table still reads
   "Owner approval: PENDING."
7. **Reusability assessment exists** (OpenCode report 2026-09-17): testbed is a
   Functioning Prototype; TCIndustries-direct share ~30%, adaptable ~50%,
   unusable ~20% (SWG IP surface, HISTORICAL bands as targets, provisionals).

## 3. The decision menu (what comes after step 10)

Four options. They are not mutually exclusive in sequence — the recommendation
in §4 orders them — but each answers the question differently if chosen alone.
Costs are suite-relative (the Phase 9 suite ≈ 20 min; full chain ≈ 60–90 min on
the project's Windows host) and in proposals+verifications, matching house style.

### Option A — Consolidate (make the record complete)

Freeze feature work; complete provenance and evidence hygiene.

- **IN:** owner-pass adoption of the ~65 pre-existing paths into git (per-file
  provenance review, honest commit messages, per
  `investigations/…Verification_Report…` §4 inventory); close the heal-evidence
  gap (new incap→revive interdependence scenario asserting a real `heal` event);
  port overrides for the nine hardcoded test clients; mission-row retention once
  decided (§6 item 3); record the pending human-pass items; file HD-TST-01's
  final approval text in `governance/`.
- **Cost:** one implementation session + one verification session; no new systems;
  all changes are process/testing/persistence, zero gameplay-behavior change.
- **Defers:** any new capability; any TCIndustries-side work.
- **Evidence for:** every future verification diff becomes complete and
  auditable; the flagged gaps close at their cheapest point (now, while context
  is fresh); A is a precondition that benefits every other option.

### Option B — Testbed Phase 11 (extend the phase ladder)

Invent a successor phase in the testbed's own style: propose → authorize →
implement → verify. Candidate content families (each would need its own
proposal and §9 decision list before any code):

- **Economy depth:** bazaar/NPC-mediated flows, inflation/deflation simulation
  harness extending the Phase 10 telemetry; long-run faucet/sink drift studies.
- **Ecology & persistence scale:** resource-spawn lifecycle depth, creature
  population dynamics beyond lairs, multi-zone persistence.
- **Service interdependence depth:** group-wide service effects, heal-event
  closure, buff economy under load.

- **Cost:** at minimum one full proposal cycle per content family (like Phase 9
  or 10), each with fresh research, tensions, and numeric asks; suites grow
  longer; maintenance of an already-wide surface.
- **Defers:** the TCIndustries pivot (C) until later.
- **Evidence for:** the testbed is the project's only working evidence engine;
  deeper evidence would strengthen any future design decision. **Evidence
  against:** the reusability report suggests the *existing* evidence is already
  sufficient to inform the next design decisions; more depth risks polishing
  evidence nobody has consumed yet.

### Option C — TCIndustries pivot (consume the evidence)

Shift the program's center of gravity from producing evidence to using it,
on the TCIndustries side of the boundary. Candidate content (each item its own
PROPOSED artifact, human-approved before adoption):

- **Design-evidence register:** index every testbed evidence artifact (suites,
  balance report, telemetry, options paper) and map each to the TCIndustries
  design questions it can and cannot inform — with explicit
  `HISTORICAL`/`PROTOTYPE` status walls.
- **Reusability extraction plan:** operationalize the ~30/50/20 split — what
  lifts as-is (telemetry layer, report taxonomy, harness conventions,
  `newRowID`, runners), what adapts (mechanics needing re-derivation +
  authorization), what never crosses (SWG IP, HISTORICAL bands as targets,
  provisionals as finals).
- **Design-input memos:** short per-topic documents carrying specific evidence
  (e.g. interdependence receipts, OUT-OF-BAND finding) into TCIndustries design
  discussions as EVIDENCE — never as design.

- **Cost:** documentation/planning sessions, no testbed code; EIC-facing items
  remain gated by HD-EIC-05/07 (no shaping without human-supplied boundaries —
  this proposal supplies none).
- **Defers:** testbed growth (B) until the evidence deficit is demonstrated.
- **Evidence for:** the manifest's stated purpose — TCIndustries is the project;
  the testbed is its instrument. Instrument upkeep beyond demonstrated need is
  scope drift.

### Option D — Halt and archive

Record the program's exit state (Phase 10 complete, all suites green), freeze
the testbed maintenance-only, file a completion note in `investigations/`, and
stop. Reversible at any time by later owner decision.

- **Cost:** one documentation session. **Defers:** everything above.
- **Evidence for:** if TCIndustries work is pausing anyway, a clean recorded
  halt beats an idle live surface.

## 4. Recommendation (owner judgment overrides; not a decision)

**A → C, with B held until C demonstrates an evidence deficit.** Reasoning:
A is cheap, benefits all options, and closes known gaps while context is fresh
(its absence is the largest standing threat to future verify-don't-claim
evidence); C is where the project's actual purpose lies, and the reusability
report shows the evidence base is already actionable; B's marginal evidence has
no demonstrated consumer yet. If C's register work reveals real gaps (e.g. no
long-run economy drift evidence exists), the specific B-family content can be
proposed then, in the ordinary way.

## 5. Tensions and honest limits (human resolves, proposal does not)

1. **Provenance completeness vs mixed provenance.** Adopting the tree (A)
   commits files containing pre-existing, partially unattributed content
   (Group C sources, untracked Phase 3–9 work). Honest commit messages mitigate
   but do not eliminate the ambiguity; the alternative (never adopting) leaves
   verification diffs permanently incomplete.
2. **Menu framing vs invented successor.** Option B requires inventing a phase
   the governing GDD ladder does not define (HISTORICAL ladder ends at 10).
   This proposal presents B as a *choice*, not a default — the
   no-automatic-successor boundary forbids assuming it.
3. **Pivot eagerness vs evidence sufficiency.** Option C consumes evidence that
   is strong but partial (heal event unexercised end-to-end; single-host
   persistence). If TCIndustries decisions need the missing evidence, B-first
   ordering is correct.
4. **Numeric asks below are decision *channels*, not decisions.** Voting to
   "address" them here does nothing; each needs its specific §9 answer.

## 6. Numeric authorization requests (values NOT set by this proposal)

1. **B4 service-XP options paper** (`docs/phase10_service_xp_options.md`):
   adopt an option (requires specifying every rate), decline (provisionals
   stand permanently), or defer (provisionals stand for now). *Default if
   unanswered: defer — the current operative state.*
2. **OUT-OF-BAND combat-scale finding** (fork unarmed 5–15 vs HISTORICAL 50–150):
   (a) record as accepted evidence, no change (recommended — it is a fork
   constant, not a TCIndustries decision); (b) authorize a scoped investigation
   proposal (patterns-only, no tuning); (c) defer.
3. **Mission-row retention** (telemetry prerequisite, structural not numeric):
   authorize retention of completed mission rows for per-terminal completion
   telemetry, or accept row-level mix reporting permanently.

## 7. Acceptance criteria (verify, don't claim)

Whichever option(s) the owner selects, the resulting work is accepted only with:

1. Per-package verification with real pasted output (static gates, fresh-DB
   suites, exit codes), one commit per verified step, per house methodology.
2. For A specifically: a provenance-honest commit inventory (per-file review
   notes), a heal-event receipt observed live in telemetry, and chain runs
   passing with no port-8080 dependency.
3. For C specifically: the register/memos filed as PROPOSED with explicit
   status walls, and zero design-authority claims anywhere.

## 8. Sequencing (non-binding, assumes A→C)

1. §9 answered (all items; nothing moves before this).
2. If A: consolidation session → verification session → report.
3. If C: evidence register → reusability extraction plan → design-input memos,
   each separately proposed/approved where it creates new artifacts.
4. If B is ever selected: its content family gets its own full proposal cycle
   (research → tensions → numeric asks → decisions) before any code.
5. D at any point: single documentation session, recorded halt.

## 9. Decisions requested (no code until answered)

| # | Decision | Options | Default if unanswered |
|---|---|---|---|
| 1 | **Post-Phase-10 direction** | A Consolidate / B Testbed Phase 11 / C TCIndustries pivot / D Halt / an ordered combination | None — no default exists; work stays closed |
| 2 | **Adopt ~65 pre-existing paths into git** (owner pass, provenance-honest) | Yes / No / Partial (specify) | No (status quo; diffs stay incomplete) |
| 3 | **Close heal-evidence gap** (new incap→revive scenario) | Yes / No | No |
| 4 | **Port overrides for phase 1–9 clients** | Yes / No | No |
| 5 | **Mission-row retention** | Authorize / accept row-level mix | Accept mix |
| 6 | **B4 options paper** | Adopt / Decline / Defer | Defer |
| 7 | **OUT-OF-BAND finding** | Record-only / investigate / defer | Record-only |
| 8 | **Proxy 4 artifact review + Historical Accuracy proxy** | Owner review (split decision recorded 2026-10-18) | Proxy 4 = DEFERRED; Historical Accuracy = formal questionnaire authored (`SWG_Historical_Accuracy-v1.md` v1.0), pending human answers |
| 9 | **File HD-TST-01 final approval text** in `governance/` | Yes (provide date/text) / defer | Defer |
| 10 | **If C:** authorize evidence-register drafting as the first artifact | Yes / No | No |

## 9.1 Recorded owner answers (2026-10-18)

Owner answers recorded against §9 above (per the proposal's own enactment rule:
answers are recorded here and in the relevant documents before any work begins).
Recorded by OpenCode (agent: Beam Free, `opencode/beam-free`) at the owner's direction;
these are the owner's decisions, not agent recommendations. The original author
byline of this file (line 6, "Buffy — Codebuff agent") predates these 2026-10-18
additions and is not altered by this recording act.

| # | Decision | Chosen option | Note |
|---|---|---|---|
| 1 | Post-Phase-10 direction | **A — Consolidate** | Ordered combination not selected; B/C/D not selected. Recommendation (A→C) partially adopted (A only). |
| 2 | Adopt ~65 pre-existing paths into git | **Yes** | Provenance-honest, per-file review; honest commit messages. |
| 3 | Close heal-evidence gap | **Yes** | New incap→revive scenario asserting a real `heal` event. |
| 4 | Port overrides for phase 1–9 clients | **Yes** | Remove port-8080 dependency. |
| 5 | Mission-row retention | **Accept row-level mix** | No retention authorization; aggregate mix reporting stands. |
| 6 | B4 service-XP options paper | **Defer** | Provisional rates stand for now. |
| 7 | OUT-OF-BAND combat-scale finding | **Record-only** | Fork unarmed 5–15 vs HISTORICAL 50–150 recorded as accepted evidence; no change. |
| 8 | Proxy 4 artifact review + Historical Accuracy proxy | **Split decision (2026-10-18)** | **Proxy 4 — DEFERRED** (emergent player-organized play requires players; no live cohort). **Historical Accuracy — Human-approved formal proposal authored:** `proposals/SWG_Historical_Accuracy-v1.md` (v1.0) created as the comprehensive human-recollection questionnaire; owner to answer from recollection; answers are EVIDENCE only. |
| 9 | File HD-TST-01 final approval text in `governance/` | **Yes** | Filed 2026-10-18 in `governance/TCIndustries_Testbed_Reuse_Human_Ruling_HD-TST-01_2026-09-14.md` (filing note appended; mechanical act only; preconditions remain Pending). |
| 10 | If C: authorize evidence-register drafting | **Yes (moot)** | Direction is A, not C, so no C work is authorized by this answer; recorded as given. |

**Consequences (A-only):** consolidation work (Option A) is now authorized to begin once
the §7 acceptance criteria are met: per-package verification with real pasted output,
one commit per verified step, provenance-honest commit inventory, heal-event receipt
observed live in telemetry, and chain runs passing with no port-8080 dependency.
B, C and D remain unselected and are not authorized. No code, tests, or commits are
performed by this recording act.

## 10. Implementation status

**IN PROGRESS (Option A authorized 2026-10-18).** Owner answers recorded in §9.1.
Option A (Consolidate) is authorized to begin; B/C/D remain unselected and una-
uthorized. Phase 10 remains the last executed phase until A's verification gates
pass. Items 5 (defer) and 7 (record-only) carry no code work. Item 8 split: Proxy 4 deferred (no code work); Historical Accuracy resolved by authoring `proposals/SWG_Historical_Accuracy-v1.md` (v1.0) — documentation act only, no code, no design authority.
Item 9 filed 2026-10-18 (date/text supplied by owner).

---

*End of proposal v0.1. Owner answers recorded 2026-10-18 in §9.1; Option A authorized
to begin. B/C/D remain unselected. Item 8 split decision recorded: Proxy 4 deferred;
Historical Accuracy questionnaire authored as `proposals/SWG_Historical_Accuracy-v1.md`
(v1.0) — pending human answers. Item 9 filed 2026-10-18. Implementation proceeds
per §7 acceptance criteria.*
