# TCIndustries — OQ-011 Respecialisation Costs & Rules — Structural Proposal

**Filename:** `TCIndustries_OQ-011_Respecialisation_Cost_Rules_Proposal_Drafter_v1_2026-09-19.md`
**Status:** PROPOSED. No authority. Ruling-ready candidate only. Nothing here is
LOCKED, and no element of this proposal changes any GDD status.
**Author:** Buffy — Codebuff agent (Freebuff), autonomous-loop drafter, Cycle 1.
**Assessor:** None — single-session authoring; no independent review pass (per GR-004, recorded as such rather than guessed).
**Date:** 2026-09-19
**Queue position:** 3 of 5 (see `proposals/TCIndustries_Continuous_Queue_Tracker.md`)
**Pre-flight:** no prior proposal or investigation addresses OQ-011 respecialisation design; the gap ID appears only in invariant and evidence rows (cited below). No duplicate pass.

---

## Status wall

- Every option below is **PROPOSED** and requires an explicit HD ruling to advance.
- **Zero numeric values.** No cooldown lengths, retraining times, fees, or decay
  rates appear anywhere (BAL-001). Cost magnitudes require BAL-001 authorisation
  later; this proposal rules the *instrument class* only.
- **Use-based skill-point acquisition is non-canonical** (Open Questions
  Register §6; GDD §36) — no option assumes it, and no acquisition model
  (OQ-011 acquisition half / PROG-003) is selected here.
- **Non-canonical architecture guard:** Skill Discipline / Skill Box
  architecture is on the non-canonical list — Option C **analyzes** the
  HISTORICAL drop-cascade pattern (E-17, HISTORICAL per Evidence Register) and
  does not promote it.
- **PROF-004 is LOCKED** (respecialisation allowed) with the GDD's own PROPOSED
  implementation principle recorded in it — this proposal structures exactly
  that open implementation surface (OQ-011).
- **PROF-005/OQ-001 specialisation budget is open**; nothing here pre-selects
  a budget model (EIC-coupled per the Invariant Register; HD-EIC-07 untouched).

## Anchors (controlling references)

| Anchor | Status | What it locks/constrains |
|---|---|---|
| PROF-004 (GDD §10) | LOCKED | Players must be able to respecialise; no permanent traps |
| PROF-004 implementation principle (GDD §10) | PROPOSED | Meaningful opportunity cost, cooldown, retraining time, loss of active capacity, or economic cost — without erasing identity, provenance, ownership, reputation history |
| OQ-011 (GDD §31) | TBD | The open question this proposal structures |
| PROG-003 (GDD §21) | TBD | Acquisition model — must support budgets + respecialisation; NOT selected here |
| INV-011a/b/c (Invariant Register v0.1, OQ-011 entry) | Candidate | No identity loss; cost real-but-bounded; acquisition-route equivalence |
| SAFE-001 (GDD §26) | HUMAN-LOCKED principle-only | Ownership/transfer safeguards — bounds protection of identity-bearing assets |
| E-17 (Design Evidence Register §1.3) | HISTORICAL | Skill-surrender guards + 250-pt budget *pattern* — evidence of implementability of drop-based respec, not a TCIndustries rule (SWG-002) |
| E-01/E-02 phase-2 (Design Evidence Register §2, OQ-011 row) | PROTOTYPE | Trainer/XP acquisition working in fork — acquisition model implementability only |

## Principle → Constraint → Invariant → Failure Condition → Test

- **Principle (PROPOSED):** respecialisation exists for identity development,
  not power re-rolling: players must be able to redirect their character's
  development (PROF-004 LOCKED) while specialist capacity remains economically
  meaningful at the participant layer (PIL-003; HD-EIC-01) and identity
  investments stay intact.
- **Constraint (derives from LOCKED layer):** any cost model must (i) never
  erase identity, provenance, business ownership, or reputation history
  (PROF-004 implementation principle + INV-011a), (ii) carry a cost players
  cannot trivially nullify (INV-011b), (iii) keep all legitimate careers
  reachable (INV-011c), (iv) not function as a de-facto permanent trap
  (PROF-004) and (v) not pre-suppose any acquisition model or budget model
  (PROG-003/OQ-001 open).
- **Invariant (candidate INV-011a, quoted):** "*Under any respecialisation
  mechanism, name, appearance, item provenance, business ownership,
  organisation membership, and reputation history survive intact.*"
- **Invariant (candidate INV-011b):** "*Respecialisation carries a cost the
  player cannot trivially nullify (time, economic, or capability opportunity
  cost) while remaining achievable — respecialisation is neither free nor
  punitive.*"
- **Invariant (candidate INV-011c):** "*Whatever the acquisition model,
  alternative specialisations remain reachable from any starting state — no
  acquisition path locks a character out of any legitimate career
  permanently.*"
- **Failure condition:** identity erasure under respecialisation; zero-cost
  respec ping-ponging as a dominant strategy; an unreachable-specialisation
  dead end.
- **Test (shape-level):** state-machine walk-through of the ruled model against
  INV-011a/b/c (checklist-class, runnable at design time); behavioural testing
  requires an implementation or simulation host (none exists today).

## Options (structural only; tradeoffs and failure modes analyzed)

### Option A — Economic-cost respecialisation (fee/materials)
Respecialisation consumes a meaningful economic cost — fee and/or special
retraining materials — with no capability loss beyond the switch itself.
- *Serves:* identity-safe (nothing lost but currency/materials); clean
  interaction with the player economy (demand flows to other players — money
  sinks connect to OQ-009, open, flagged); simplest audit surface (one
  transaction to walk through).
- *Tradeoffs:* wealthier players can redirect faster — a participant-layer
  concern (PIL-003/HD-EIC-01: does wealth substitution for time threaten
  independent-participant comparative advantage? flagged for the actor-complete
  walk-through, not decided here); sinks via respec fees require the OQ-009
  currency decision first (flagged dependency).
- *Failure modes:* if cost is trivially affordable by established participants
  → INV-011b "cannot trivially nullify" fails; pure-currency costs with weak
  sinks risk inflation interaction (OQ-009 open; INV-009a observable-imbalance
  shape cited, not designed).
- *Non-canonical-list check:* clean.

### Option B — Retraining-time respecialisation (time cost)
Switching starts a retraining process consuming real (in-game) time during
which the outgoing capacity winds down and the incoming capacity builds.
- *Serves:* cost is identity-safe and wealth-independent — a structural answer
  to A's wealth-substitution concern; matches the GDD's own PROPOSED
  implementation principle word-for-word ("retraining time").
- *Tradeoffs:* friction concentrates in real-world terms (player returns to
  find capacity changed); asynchronous players (limited play windows) bear
  the same wall-clock cost as daily players — an interdependence-opportunity
  ("interdependence must create opportunity rather than coercive
  inconvenience", PIL-003) tension the walk-through must check; requires some
  state machinery for winding-down/building capacity.
- *Failure modes:* if retraining is fast, INV-011b fails (trivially nullifiable
  by patience); if slow, respecialisation becomes de-facto punished and
  PROF-004's "must be able to respecialise" weakens — the boundedness judgment
  is exactly the later BAL-001-class call, not made here.
- *Non-canonical-list check:* clean.

### Option C — Capability-decay respecialisation (surrender-based)
Respecialisation = surrendering outgoing capacities; incoming capacity builds
per whatever acquisition model is later ruled (not selected here); the
HISTORICAL drop-cascade *pattern* (E-17) evidences that surrender-based
respecialisation is implementable without identity loss.
- *Serves:* strongest interdependence alignment — the cost is capability
  itself, the most direct guard against respec-as-power-re-roll; no currency
  or wall-clock dependency; the GDD's PROPOSED principle explicitly names
  "loss of active specialisation capacity" as a candidate instrument.
- *Tradeoffs:* interlocks with the OQ-001 budget model (whatever a character
  "gives up" must exist as a thing to give up) — a flagged dependency, not a
  pre-selection; "cost" is paid in-kind, so INV-011b's "cannot trivially
  nullify" reads as "capacity gone is gone" — audit surface is the surrender
  state machine.
- *Failure modes:* if surrender is instant and free of transition cost, it can
  approach respec-as-free-re-roll (INV-011b pressure); if surrender erases
  history/reputation, INV-011a fails (explicitly forbidden by PROF-004's
  implementation principle — the instrument must attack *capacity*, never
  identity).
- *Non-canonical-list check:* the HISTORICAL drop-cascade *pattern* is cited as
  implementability evidence only (SWG-002); the non-canonical
  Skill Discipline/Skill Box *architecture* is not promoted — surrender-based
  switching is not unique to that architecture (recorded reasoning, not
  promotion).

### Cross-option note (recorded, not decided)
A+B, A+C, or B+C combinations are lawful outcomes of this ruling ("cost" may
stack instruments) — the ruling may pick one, pick several, or amend.

## What this proposal does NOT resolve

- The specialisation budget model (OQ-001/PROF-005 — EIC-coupled, gated behind
  HD-EIC-07 boundaries; explicitly not pre-selected by any option).
- The skill acquisition model (acquisition half of OQ-011 / PROG-003 — tracked
  separately in the queue backlog; use-based SP stays non-canonical).
- Any cost *magnitude* (BAL-001-gated; requires separate numeric authorisation).
- SAFE-001 safeguard mechanisms; currency design (OQ-009) that Option A's sink
  interaction would eventually touch; multi-account implementation (OQ-012).
- Any implementation, simulation, or testbed work; E-17 is cited as
  implementability evidence only (PROT-001).

## Ruling requested (HD-<Area>-<NN> slot reserved in the Human Rulings Register)

**Question for the owner — rule yes / no / amend:**

1. **Primary ruling:** Select the cost instrument class for OQ-011
   respecialisation: **A** (economic cost), **B** (retraining-time), **C**
   (capability-decay/surrender), any lawful combination, or **defer**.
2. **Invariant ruling (independent):** Do you approve candidate invariants
   INV-011a/b/c as design invariants (TEST-001 class) for whichever instrument
   is chosen? (Approval records them as design invariants; they set no numeric
   value.)
3. **Scope confirmation:** Confirm that cost magnitudes are deferred to a
   future BAL-001-authorised pass, and that the acquisition model (PROG-003)
   and budget model (OQ-001) remain separately open.

*Every claim above cites its file and anchor. On any ruling, the tracker updates
and the queue maintains 3–5 live items per the loop rules.*

---
*End of proposal. PROPOSED — awaiting HD ruling. No status changed anywhere by this document.*
