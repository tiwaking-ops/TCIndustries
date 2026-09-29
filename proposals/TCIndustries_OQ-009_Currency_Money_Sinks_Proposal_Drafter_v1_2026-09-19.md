# TCIndustries — OQ-009 Currency & Money Sinks — Structural Proposal

**Filename:** `TCIndustries_OQ-009_Currency_Money_Sinks_Proposal_Drafter_v1_2026-09-19.md`
**Status:** PROPOSED. No authority. Ruling-ready candidate only. Nothing here is
LOCKED, and no element of this proposal changes any GDD status.
**Author:** Buffy — Codebuff agent (Freebuff), autonomous-loop drafter, Cycle 9.
**Assessor:** None — single-session authoring; no independent review pass (per GR-004, recorded as such rather than guessed).
**Date:** 2026-09-19
**Queue status:** **drafted-ahead — backlog rank 8, NOT a live queue item.** Enters the queue at the next freed slot (live queue is 5/5). See `proposals/TCIndustries_Continuous_Queue_Tracker.md`.
**Pre-flight:** OQ-009 appears in the GDD §31 table, the Open Questions Register §2, the Invariant Register (OQ-009 entry), and the Evidence Register (OQ-009 row + §3 gap 6). No prior proposal or investigation structures the currency/sink decision. No duplicate pass.

---

## Status wall

- Every option below is **PROPOSED** and requires an explicit HD ruling to advance.
- **Zero numeric values.** No sink rates, faucet values, price floors, or
  emission schedules (BAL-001). Every sink below is an instrument *class*
  whose parameters are later-defined under BAL-001 authorisation.
- **Single unified currency as the default is on the non-canonical list**
  (Open Questions Register §6; GDD §36). Per that register's own terms these
  items remain non-canonical "**unless explicitly promoted by a future human
  decision**" — therefore Option A below is lawful ONLY as an explicit,
  recorded promotion of that specific item, never as a silent default. This
  is stated in the option itself so a violating adoption is detectable.
- **ECO-004 is TBD** and names the possibility space: "Exact currency
  systems, money sinks, and pricing mechanics are unresolved. Multiple
  currencies or regional media of exchange remain possible design directions."
- **Player-driven economy binds (ECO-001 LOCKED):** NPC vendors must not make
  player production irrelevant — no option may route economic dominance
  through a currency mechanism (e.g. NPC-only emission points that gate
  player commerce).
- **ECO-003 (DERIVED)** binds: no pure spreadsheet optimisation — currency
  design must stay grounded in world interaction, location, and risk.
- **HD-RET-01 interaction (ruled):** all sell-side commerce is player-vendored;
  any currency medium flows through player vendors — pricing mechanics
  (RET-002) and this ruling compose but neither pre-selects the other.
- **Backlog 6c interaction (named-open by HD-SCOPE-01):** the early-game
  on-ramp question is recorded as interacting with OQ-009; this proposal
  structures the currency/sink layer the on-ramp ruling will compose with.
  It does not answer the on-ramp.
- Evidence honesty: faucet/sink-tagged ledger and snapshot accounting are
  PROTOTYPE-evidenced (E-09); **long-run faucet/sink balance is
  NOT-MEASURABLE with current evidence** (Evidence Register §3 gap 6 — all
  telemetry is single-session scale). No option claims long-run validation
  it cannot have.

## Anchors (controlling references)

| Anchor | Status | What it locks/constrains |
|---|---|---|
| OQ-009 (GDD §31) | TBD | "Currency system(s) and money sinks?" — economic stability; the open question structured here |
| ECO-004 (GDD §15) | TBD | Currency systems, money sinks, pricing unresolved; multiple/regional media remain possible directions |
| ECO-001 (GDD §15) | LOCKED | Player-driven economy; NPC vendors must not make player production irrelevant |
| ECO-003 (GDD §15) | DERIVED CONSTRAINT | No pure spreadsheet optimisation — world-grounded economy |
| BAL-001 (GDD §25) | LOCKED | No premature numeric tuning — structure only |
| INV-009a/b (Invariant Register v0.1, OQ-009 entry) | Candidates | Sink-sufficiency observability; currency contestability (quoted under Invariants) |
| SAFE-003 (GDD §26) | PROPOSED | Economic observability — detect and respond to systemic problems without relying solely on player reports |
| Non-canonical list (Open Questions Register §6) | Confirmed boundary | Single unified currency as *default* excluded; explicit promotion path stated in GDD §36 |
| HD-RET-01 (decision record `proposals/TCIndustries_OQ-016_Ruling_Record_HD-RET-01_2026-09-19.md`) | HUMAN-LOCKED | Player-only vendors — the commerce surface any medium flows through |
| E-09 (Design Evidence Register §1.1) | PROTOTYPE | Faucet/sink-tagged credit ledger; append-only sale history — observed working; pattern evidence only |
| Evidence Register §3 gap 6 | EVIDENCE gap | Long-run faucet/sink balance NOT-MEASURABLE (single-session telemetry only) |
| T-07 (Gap Analysis §7; NOT STARTED, not approved) | — | Economic observability requirements brief would derive TCIndustries requirements from E-07/E-08/E-09 patterns — cited as informing, not assumed |

## Principle → Constraint → Invariant → Failure Condition → Test

- **Principle (PROPOSED):** money in TCIndustries is a player-economy organ,
  not a game-score: whatever medium(s) exist, every faucet and sink must be
  tagged and observable from day one (SAFE-003), sinks must flow through
  player-provided services and world systems (repair, maintenance, transit,
  civic obligations) rather than arbitrary fees (ECO-003), and no medium may
  be mandated into acceptance that players would not choose (ECO-001's
  player-driven principle applied to money itself).
- **Constraint (derives from LOCKED layer):** any currency/sink design must
  (i) tag every faucet and sink at the ledger level (INV-009a's
  observability requirement), (ii) keep NPC emission/absorption points from
  dominating player production (ECO-001), (iii) carry no numeric parameters
  (BAL-001), (iv) if adopting a single unified currency, do so only by
  explicit promotion of the non-canonical-list item (recorded in the
  ruling text), (v) if adopting multiple media, satisfy contestability
  (INV-009b) rather than mandating zombie currencies.
- **Invariant (candidate INV-009a, quoted verbatim from the register):** "*For
  any currency design, total faucet flow exceeds sink flow in every bounded
  observation window by a measurable, monitorable margin that the
  observability system can detect and attribute — i.e., inflation cannot
  proceed silently.* (Structure, not numbers: no rate is proposed; the
  invariant is that the *imbalance is observable and attributable*.)"
  **Precision note for the ruling (recorded, not resolved):** the opening
  clause reads two ways — (a) mandating that faucets exceed sinks (which
  would be a standing inflation mandate), or (b) requiring that *if* faucets
  exceed sinks, the imbalance is measurable and attributable (the
  parenthetical's "imbalance is observable and attributable" supports this
  reading). The owner should gloss the intended reading at approval time —
  the same per-item precision the HD-EIC-09 reconciliation recorded for
  "trivially".
- **Invariant (candidate INV-009b, quoted verbatim from the register):** "*If
  multiple currencies/media are adopted, each retains at least one use case
  in which it is strictly preferred, or it disappears through player choice —
  no forced-acceptance zombie currency is design-mandated.* (Applies only if
  multi-currency is chosen; records the shape so the choice can be tested.)"
- **Failure condition:** unattributable aggregate money growth (no faucet/sink
  accounting — INV-009a failure); a mandated medium with no preferred use
  (INV-009b failure); or a currency design that lets NPC/system emission
  points out-compete player production (ECO-001 breach).
- **Test (shape-level):** ledger-audit walk-through of the ruled design:
  enumerate every faucet and sink the design creates, verify each is tagged
  and attributable at the ledger level (checklist-class, runnable at design
  time — the E-09 pattern is the reference shape); full behavioural testing
  requires an implementation or simulation host, and long-run balance
  validation additionally requires authorised numerics plus extended-horizon
  telemetry (currently NOT-MEASURABLE, Evidence Register §3 gap 6).

## Options (structural only; tradeoffs and failure modes stated)

### Option A — Single unified currency (explicit non-canonical-list promotion)
One universal medium, explicitly promoted from the non-canonical list by this
ruling (recorded as such in the register entry); faucet/sink-tagged ledger
accounting from day one; sinks as instrument classes (repair/maintenance,
travel, civic obligations, listing/service fees) with parameters deferred to
BAL-001 passes.
  **[Cycle-12 annotation, applied at the Cycle-13 consistency audit — this
  annotation was recorded in HD-ITM-01 §5c and HD-BLD-02 §5 and in the
  tracker's Cycle-12 log, but not applied to this file at Cycle 12.]**
  Two ruled interactions with this proposal's sink classes: (1) **HD-ITM-01
  (OQ-010, no-baseline-decay)** — decay-driven repair demand does not exist,
  so the "repair/maintenance" class shrinks to **destruction/loss-driven
  restoration**, and its adoption remains a future structural + BAL-001
  decision. (2) **HD-BLD-02 (BLD-004, maintenance-credit removal)** — the
  owner's ruled maintenance-credit instrument (structures consume ongoing
  maintenance credit; zero credit ⇒ removal) gives "civic obligations" a
  concrete ruled exemplar: maintenance credits as a settlement-maintenance
  sink candidate. Neither interaction alters this proposal's options A/B/C;
  both will be presented at the OQ-009 ruling session.
- *Serves:* simplest trade surface across the ruled player-vendor economy
  (HD-RET-01); cleanest ledger observability (one unit of account); regional
  price *differences* still expressible (prices vary; the medium does not).
- *Tradeoffs:* discards ECO-004's named "multiple currencies or regional
  media" direction in one ruling (irreversible-feeling — though a later
  ruling could add media); the promotion itself is a canonical act the
  register must record explicitly (status discipline cost, paid at filing);
  the spreadsheet-optimisation anti-goal pressure is highest with a single
  fungible medium (ECO-003 must be enforced through world-grounded sinks).
- *Failure modes:* silent promotion (adopting A without the explicit
  promotion record — the failure this proposal is structured to make
  impossible); single-medium monoculture smoothing away regional economic
  character (ECO-002 pressure; mitigated by regional price variation, not
  by the medium).
- *Non-canonical-list check:* requires explicit promotion — stated; absent
  that statement in the ruling, A is not a lawful outcome.

### Option B — Multiple / regional media (contestability-governed)
A primary medium plus optional regional or functional media (e.g. settlement
scrip under HD-BLD-01's civic surface, organisation credit under a future
OQ-013 ruling); INV-009b governs: each medium must retain a strictly
preferred use case or be allowed to die by player choice.
- *Serves:* ECO-004's named direction at full strength; regional economic
  character (ECO-002) gets a monetary expression, not just price variation;
  settlement/organisation media give PIL-004 institutions real economic
  functions; INV-009b is the binding integrity condition.
- *Tradeoffs:* the heaviest design and UX surface (exchange rates — a
  numeric layer, deferred; arbitrage surfaces; player confusion); interacts
  with OQ-007 basic depth (settlement media would need civic-depth
  instruments that are currently deferred — a flagged dependency, not
  resolved here); exchange-rate arbitrage must be surveilled by the
  observability surface (SAFE-002/INV-012-family patterns).
- *Failure modes:* zombie currency (mandated acceptance without preference —
  INV-009b failure); arbitrage exploits around media boundaries (SAFE-002
  surface); complexity suppressing commerce liquidity in the early economy
  (tension with the 6c on-ramp question — flagged, not resolved).
- *Non-canonical-list check:* clean (does not touch the single-currency
  item; B is ECO-004's own named direction).

### Option C — Medium-agnostic accounting skeleton (defer the medium, rule the plumbing)
Rule only what every future medium choice inherits: faucet/sink-tagged
ledger accounting, append-only transaction history, snapshot jobs,
attributability of imbalances (INV-009a's observability half regardless of
reading), and the named-open sink instrument classes. The medium choice —
single, multiple, or regional — is explicitly deferred to a later ruling.
- *Serves:* smallest ruling; unblocks the on-ramp question (6c) and all
  commerce-adjacent design without committing a medium; exactly the layer
  T-07's (unapproved) requirements brief would detail — the two compose;
  consistent with the "smallest ruling that unblocks the most downstream
  work" loop principle.
- *Tradeoffs:* INV-009b becomes vacuous-for-now (no media chosen — recorded
  as open rather than approved-with-design); vendors/cities/scope design
  must proceed with "a medium will exist" as the only assumption (mild
  under-specification — the recorded-assumption discipline applies); the
  medium ruling itself is not removed from the backlog, only deferred.
- *Failure modes:* medium-less drift — downstream proposals accumulating
  local medium assumptions (mitigated by recording them per proposal, as
  this loop does); repeated reopening of adjacent questions.
- *Non-canonical-list check:* clean — defers rather than assumes; the
  single-currency item stays non-canonical (no promotion, no default).

### Cross-option notes (recorded, not decided)
- **Sink instrument classes (any option):** repair/maintenance (OQ-010 —
  live queue item; whether repair demand exists depends on that ruling),
  settlement maintenance (HD-BLD-01 basic depth implies the obligation;
  its monetary instrument is a sink candidate), travel/transit (OQ-006,
  zero evidence, backlog), respecialisation fees (OQ-011 Option A — live
  queue item; flagged there). Each sink's *adoption* is a later structural
  + BAL-001 decision; this proposal only names the classes so the ledger
  tags exist from day one.
- **T-07 interface:** if the owner later approves T-07 (observability
  requirements brief), its output becomes the requirement detail for
  whichever option is ruled here; this proposal claims none of its content.

## What this proposal does NOT resolve

- The medium choice itself (that is the requested ruling; Option C defers it
  explicitly).
- Any numeric parameter: sink rates, emission values, exchange rates,
  thresholds (BAL-001; separate authorised passes).
- Pricing mechanics (RET-002 vendor pricing — open; composes with HD-RET-01).
- The 6c early-game on-ramp question (named-open by HD-SCOPE-01; this
  proposal structures the layer it will compose with, nothing more).
- Regional economic design (ECO-002, WRLD-002 — PROPOSED layers; OQ-006
  transport, zero evidence, backlog).
- Organisation/settlement media mechanics (OQ-013; civic-depth instruments —
  deferred in HD-BLD-01).
- Any implementation, simulation, or testbed work; E-09 is PROTOTYPE pattern
  evidence only (PROT-001); long-run balance evidence does not exist
  (Evidence Register §3 gap 6).

## Ruling requested (HD-<Area>-<NN> slot reserved in the Human Rulings Register)

**Question for the owner — rule yes / no / amend (when the item is formally
queued):**

1. **Primary ruling:** Select the currency/sink structure (OQ-009): **A**
   (single unified currency — requires explicit non-canonical-list promotion
   recorded in the ruling), **B** (multiple/regional media under
   contestability), **C** (medium-agnostic accounting skeleton; medium
   deferred), or **defer** (amendments welcome).
2. **Invariant ruling (independent):** Do you approve candidate invariant
   INV-009a as a design invariant (TEST-001 class) — and if so, with which
   reading of the precision note (detectability-of-imbalance vs.
   faucet-exceeds-sink mandate)? Do you approve INV-009b (approvable now to
   bind any future multi-medium choice, or with Option B)?
3. **Scope confirmation:** Confirm that (a) all sink parameters defer to
   BAL-001-authorised passes, (b) sink instrument classes are named-open
   (each adopted by its own future ruling), and (c) the 6c on-ramp question
   remains separate and unanswered here.

*Every claim above cites its file and anchor. On any ruling, the tracker
updates and the queue maintains 3–5 live items per the loop rules.*

---
*End of proposal. PROPOSED — awaiting HD ruling (upon formal queueing). No status changed anywhere by this document.*
