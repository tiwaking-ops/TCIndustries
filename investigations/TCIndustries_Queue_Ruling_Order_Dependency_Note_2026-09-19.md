# TCIndustries — Queue Ruling-Order & Dependency Note (Cycle 3)

**Status:** ADVISORY analysis. No authority. This note recommends a ruling
order; the owner may rule in any order and any sequence is lawful. It adds no
design content, opens no gate, and changes no status.

**Author:** Buffy — Codebuff agent (Freebuff), autonomous-loop drafter, Cycle 3.
**Assessor:** None — single-session authoring; no independent review pass (per GR-004, recorded as such rather than guessed).
**Date:** 2026-09-19
**Pre-flight:** grep of `proposals/` + `investigations/` for ruling-order/dependency analysis returned zero prior artifacts (verified 2026-09-19). Gap Analysis §7.1 sequences the T-01–T-09 *tasks*; the Multi-LLM Plan §5 scopes *gaps* — neither maps proposal-to-proposal ruling dependencies. No duplicate pass.
**Subject:** the five Cycle-1 queued proposals (queue 5/5, cap in force —
`proposals/TCIndustries_Continuous_Queue_Tracker.md`):
OQ-010 durability · OQ-016 vendor/retail · OQ-011 respecialisation ·
OQ-007 city formation · OQ-015 first-playable scope.

---

## 1. Why ruling order matters

Ruling a proposal whose options *embed a sub-decision of another proposal* can
force a silent assumption unless the embedded sub-decision is explicitly
recorded. Each Cycle-1 proposal already flags its embedded sub-decisions
("flagged" notes in Options sections) — this note maps them so the owner can
rule in the order that avoids back-door assumptions entirely, the same
discipline `AGENTS.md` imposes on all documentation work.

## 2. Dependency matrix (who embeds whose sub-decisions)

| Proposal | Embeds a sub-decision of | Where (exact location in the proposal) | Effect if ruled first |
|---|---|---|---|
| **OQ-015** (first-playable scope) | OQ-016 | Option S1: "the NPC-supply floor assumed for early-game shop viability is exactly the OQ-016 Option B shape — flagged: if OQ-016 is ruled Option A (player-only vendors), S1's baseline assumption needs its own ruling" | Choosing S1 before OQ-016 is ruled imports an OQ-016 assumption; must then be recorded as flagged-TBD, not silently adopted |
| **OQ-015** (first-playable scope) | OQ-007 | Option S2: "embeds a sub-decision of OQ-007 (formation model) — flagged: ruling S2 before OQ-007 risks a city-formation assumption entering by the back door; recommended sequencing if chosen is OQ-007 first" | Choosing S2 before OQ-007 imports a formation-model assumption; same recording obligation |
| **OQ-016** (vendor/retail) | Nothing among the five | — | Fully independent |
| **OQ-007** (city formation) | Nothing among the five (its "civic depth" axis *feeds* backlog OQ-013 org rights, but embeds no queued proposal's decision) | Cross-option note: civic depth "will bind to OQ-013/OQ-009 follow-ups" | Fully independent of the five |
| **OQ-011** (respecialisation) | Nothing among the five | — | Fully independent |
| **OQ-010** (durability) | Nothing among the five | — | Fully independent |

**Read of the matrix:** OQ-010, OQ-011, OQ-016, OQ-007 are mutually
independent — ruleable in any order, any number per session. OQ-015 is the
only *composing* ruling: it depends on two of the others.

## 3. Recommended order (advisory; deviations are lawful)

1. **OQ-016 and/or OQ-007 first** (or any time before OQ-015) — independent
   rulings that OQ-015's S1/S2 options reference.
2. **OQ-010, OQ-011 any time** — fully independent; zero ordering risk.
3. **OQ-015 last** — rule it once OQ-016 and OQ-007 are decided (or rule it
   first *with the embedded assumptions explicitly recorded*, per §4).

Rationale: this order means every OQ-015 option can be evaluated against
*ruled* vendor and city structure rather than against flagged placeholders —
the smallest change that removes all silent-assumption risk from the scoping
ruling (the owner's "smallest ruling that unblocks the most downstream work"
principle, applied to ruling order itself).

## 4. If ruled out of order (deviation protocol, for the record)

If OQ-015 is ruled before OQ-016/OQ-007, the ruling text should state which
embedded assumption is being adopted *for the slice only* (e.g., "S1's NPC
floor is an OQ-016-Option-B-shaped placeholder, not an OQ-016 ruling") — the
OQ-015 proposal already words this as a flagged condition; adopting it via an
explicit sentence in the ruling keeps it recorded rather than silent. This
matches the repository's standing rule: assumptions preserved as ASSUMPTION,
never promoted by use (AGENTS.md authority model #6).

## 5. Backlog couplings (context only; no action needed)

Recorded so future cycles pull backlog items in the same discipline:
- OQ-011 **Option A** (economic cost) → interacts with OQ-009 currency/sinks
  (backlog rank 8): a respec fee needs a currency to be paid in.
- OQ-007 **civic depth** axis → feeds OQ-013 organisation rights (backlog
  rank 16).
- OQ-015 **any slice** → placeholder acquisition shape must respect OQ-011's
  acquisition half (backlog rank 6) — flagged PROTOTYPE in the proposal.
None of these block the five rulings; they only shape which backlog item gets
drafted after each ruling lands.

## 6. What this note does NOT resolve

- Any of the five rulings themselves (all remain PROPOSED, ruling-requested).
- Any backlog drafting decision (queue at 5/5 cap; exceeding it requires
  explicit owner authorisation per the loop rules).
- Any design content, invariant approval, or status change — advisory mapping only.

## 7. Cycle-4 update (2026-09-19)

OQ-016 **RULED** (HD-RET-01, Option A — player-only vendors; decision record
`proposals/TCIndustries_OQ-016_Ruling_Record_HD-RET-01_2026-09-19.md`;
register filing = owner act). Consequences for this note:
- OQ-015's S1 sub-decision (NPC floor) is **dissolved** — no NPC goods supply
  exists under the ruled model; no deviation-protocol recording needed.
- Remaining recommended order: **OQ-007 → OQ-015 last**; OQ-010 / OQ-011
  (respecialisation) anytime.
- New queued item (OQ-011a skill acquisition, drafted Cycle 4 from backlog
  rank 6) is **independent** of all five queued rulings.

---

## 8. Cycle-6 update (2026-09-19)

OQ-007 **RULED** (HD-BLD-01: Option A free-claim formation, basic governance
depth, INV-007a approved / INV-007b remains candidate; decision record
`proposals/TCIndustries_OQ-007_Ruling_Record_HD-BLD-01_2026-09-19.md`;
register filing = owner act). Consequences for this note:
- OQ-015's second embedded sub-decision (S2/formation) is **resolved** —
  OQ-015 is now fully unblocked and ruleable in any order with zero imports.
- The ordering problem this note was written for is dissolved: all remaining
  live items (OQ-010, OQ-011 respec, OQ-015, skill acquisition,
  multi-account/exploit) are mutually independent.
- Recorded for future re-ranks: BLD-004 abandonment handling surfaced as the
  priority adjacent gap by HD-BLD-01 §2 (free-claim + INV-007a makes
  abandonment the load-bearing mechanism).

---

## 9. Cycle-7 update (2026-09-19)

Pre-flight re-check: HD-RET-01 and HD-BLD-01 remain **unfiled in
governance/** (register grep clean; OQR §2 pointers still read TBD;
project_memory last entry 2026-09-18). Rulings stand as recorded in their
proposals/ decision records; queue unchanged at 5/5.

**Per-invariant approval recorded by the owner this programme (precision
note for filing time):** HD-RET-01 approved INV-016a + INV-016b; HD-BLD-01
approved INV-007a only (INV-007b explicitly remains candidate). Filing-time
invariant-register annotations must reflect the per-invariant granularity —
no batch approval is to be inferred from either record.

Drafted-ahead status: the BLD-004 abandonment & placement proposal now
exists as `proposals/TCIndustries_BLD-004_Abandonment_Placement_Proposal_Drafter_v1_2026-09-19.md`
(backlog priority 6b, drafted-ahead at Cycle 7; formally queued Cycle 8 —
see §10). It composes with the ruled free-claim + basic-depth structure but
embeds no other queued proposal's sub-decision.

## 10. Cycle-8 update (2026-09-19)

OQ-015 **RULED** (HD-SCOPE-01: Option S1 economic spine; early-game on-ramp
recorded as named-open; INV-015 approved; deferred-TBD confirmed; decision
record `proposals/TCIndustries_OQ-015_Ruling_Record_HD-SCOPE-01_2026-09-19.md`).
Consequences for this note:
- All three composing/ruling-order threads this note tracked are now closed:
  every originally-queued item that had embedded sub-decisions is ruled, and
  the five current live items (OQ-010, OQ-011 respec, skill acquisition,
  multi-account/exploit, BLD-004) are mutually independent.
- New named-open item created (backlog 6c: early-game economic on-ramp) —
  small, targeted, dependent on nothing currently queued.
- Approved-invariant count for filing-time annotations: INV-016a, INV-016b,
  INV-007a, INV-015 (per-invariant granularity; INV-007b remains candidate).

---

## 11. Cycle-10 update (2026-09-19)

OQ-010 **RULED** (HD-ITM-01: Option C no-baseline-decay; INV-010 approved in
conditional form; structure-only; decision record
`proposals/TCIndustries_OQ-010_Ruling_Record_HD-ITM-01_2026-09-19.md`).
Consequences for this note:
- Live queue (all mutually independent): OQ-009 currency/sinks, OQ-011
  respec, skill acquisition, multi-account/exploit, BLD-004 abandonment.
- New priority backlog item 6d (salvage & destruction handling) surfaced by
  the ruling — it composes with OQ-009 (restoration sink class) and the
  non-canonical list; embeds no live item's sub-decision.
- OQ-009's INV-009a precision note (two readings) remains the one open gloss
  request pending a ruling.

---

## 12. Cycle-12 closing update (2026-09-19; recorded at the Cycle-13 audit)

Cycles 11–12 ruled without changing any ordering relationship this note
tracks: OQ-011 respecialisation → **HD-PROF-01 (PARTIAL)** (INV-011a
approved; cost instrument DEFERRED — item stays queued, defer frees no slot)
and BLD-004 abandonment → **HD-BLD-02** (owner's own maintenance-credit
removal model, verbatim; INV-BLD-004b approved). **Current live queue (4,
within the 3–5 band; §11's five-item list is superseded):** OQ-009
currency/sinks · OQ-011 respecialisation (partial) · Skill acquisition ·
Multi-account/exploit — all still mutually independent; the
advisory-ordering purpose of this note remains fully served. Backlog added
since §11: 6d (salvage & destruction — designated next draft) and 6e
(HD-BLD-02 follow-ups: contents handling on removal; settlement-institution
dissolution).

---

*End of note. ADVISORY — no authority, no status changes.*
