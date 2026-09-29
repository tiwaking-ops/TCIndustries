# TCIndustries — BLD-004 Abandonment Handling — Human Ruling Record (HD-BLD-02)

**Status:** HUMAN RULED 2026-09-19 (owner decision recorded this session). This
document is the decision record; authority attaches to the owner's explicit
ruling, and this record preserves it verbatim for filing.
**Filing state:** decision record filed in `proposals/` by the drafter (hard
boundary: no `governance/` writes). Ready-to-file register entry in §3.
**Author:** Buffy — Codebuff agent (Freebuff), ruling-session recorder, Cycle 12.
**Date:** 2026-09-19
**Gives effect to:** `proposals/TCIndustries_BLD-004_Abandonment_Placement_Proposal_Drafter_v1_2026-09-19.md` (retained as provenance; the owner's ruling **amends** it — see §1.1).
**Ruling ID:** **HD-BLD-02** — register's HD-<Area>-<NN> pattern (area: BLD; second of area).

---

## 1. The ruling (three parts, as decided by the owner)

### 1.1 Primary ruling — OWNER AMENDMENT (owner's own model, verbatim)

The owner did not select any offered option (A/B/C/hybrid). The owner ruled, in
their own words:

> **"Once the credit paid to the house becomes 0 then the house is removed from the world."**

**Owner's sentence controls.** The following structural restatement is
**drafter gloss** (recorded to make the ruling implementable; it carries no
authority beyond the owner's sentence and may be corrected at filing review):

- (a) Placed structures require ongoing **maintenance credit** contribution —
  the concrete form of the shared-maintenance responsibility ruled at basic
  depth in HD-BLD-01.
- (b) When a structure's maintenance credit **reaches zero**, the structure is
  **removed from the world** — no recourse window after zero, no title
  transfer, no reclamation-to-claimant: removal is the terminal state.
- (c) The recourse path is the **credit balance itself**: a visible, declining
  state the owner can top up at any time before zero. (This is why
  INV-BLD-004a — which mandates a post-degradation recourse window — was not
  needed and was not approved; see §1.2.)
- (d) **Open structural questions created by the owner's model** (backlog 6e,
  small targeted future rulings): what happens to a structure's
  **contents/items** on zero-credit removal; and the **settlement-institution
  dissolution rule** (a settlement whose last structure is removed has no
  structures left — INV-007a is satisfied at structure scale, but the civic
  institution's end-state needs a rule).
- (e) All credit parameters (rates, costs, depletion schedules) are
  **numeric-gated** (BAL-001) — this ruling settles the instrument ("credit vs
  zero threshold"), no values.

### 1.2 Invariant ruling

| Candidate | Owner decision |
|---|---|
| **INV-BLD-004b** (no ghost equilibrium: zero-sustenance settlements converge to non-functional state; no placed structure persists at full function with zero involvement) | **APPROVED** as a design invariant (TEST-001 class). The owner's model satisfies it by construction: credit → 0 ⇒ removal ⇒ no full-function zero-involvement structures persist |
| **INV-BLD-004a** (function-before-title graduation with mandated recourse window) | **NOT approved** — remains a candidate. Consistent with the owner's model: the credit balance is the recourse path; no post-zero window is mandated |

### 1.3 Scope confirmation

**Confirmed:** abandonment handling only — placement rules, lot systems, and
density limits remain TBD (BLD-004 remainder); title transitions only to the
world structure or via player-authored succession (moot under the owner's
model — removal is terminal); all parameters deferred to BAL-001-authorised
passes.

## 2. What this ruling explicitly does NOT authorise

- Any numeric value: credit rates, costs, depletion schedules, thresholds (BAL-001).
- Placement/lot/density rules (BLD-004 remainder — stays TBD).
- **Contents-handling rule** on zero-credit removal (open — backlog 6e; this
  is also the natural composition point with backlog 6d salvage & destruction).
- **Settlement-institution dissolution rule** (open — backlog 6e).
- Any intermediate functional-decay staging (the ruled model is binary at
  structure scale: credit > 0 = present; 0 = removed; decay *staging* would
  need a follow-up ruling).
- SAFE-001 mechanism design beyond what removal implies; OQ-013; any EIC
  shaping (HD-EIC-05/07 in force).
- GDD text edits and status-pointer updates (reserved controlled acts, §4).
- Options A, B, C of the proposal are not adopted in any form; the owner's
  amendment replaces them (they remain in the proposal's provenance).

## 3. Ready-to-file register entry (verbatim, for `governance/TCIndustries_Human_Rulings_Register.md`)

```markdown
## HD-BLD-02 — BLD-004 Abandonment Handling (Maintenance-Credit Removal)
**Source:** Owner ruling session, 2026-09-19, on `proposals/TCIndustries_BLD-004_Abandonment_Placement_Proposal_Drafter_v1_2026-09-19.md`; decision record `proposals/TCIndustries_BLD-004_Ruling_Record_HD-BLD-02_2026-09-19.md`.
**Status:** HUMAN-LOCKED (abandonment-handling structure only). Complements HD-EIC-01–09, HD-GDD-01, HD-TST-01, HD-RET-01, HD-BLD-01, HD-SCOPE-01, HD-ITM-01, HD-PROF-01, which remain in force. Supersedes nothing; amends the proposal's option set (owner's own model adopted in place of Options A/B/C).

### Decision
1. **PRIMARY — Owner's model, verbatim:** "Once the credit paid to the house becomes 0 then the house is removed from the world." Structures require ongoing maintenance credit (the concrete form of HD-BLD-01's shared-maintenance responsibility); at zero credit the structure is removed from the world. The credit balance is the recourse path (visible, declining, top-up before zero); no post-zero recourse window exists. Recorded interpretation (drafter gloss, no authority beyond the owner's sentence) in the decision record §1.1, including two open structural questions routed to the backlog: contents handling on removal, and settlement-institution dissolution.
2. **INVARIANT — INV-BLD-004b APPROVED** as a design invariant (TEST-001 class): no ghost equilibrium — zero-sustenance settlements converge to non-functional state; no placed structure persists at full function with zero ongoing involvement. (Satisfied by construction under the ruled model.) **INV-BLD-004a NOT approved** and remains a candidate.
3. **SCOPE — confirmed:** abandonment handling only; placement/lots/density remain TBD; no numeric parameter authorised (credit rates/costs/schedules deferred to BAL-001-authorised passes).

**Explicitly does not authorise:** any numeric value; placement/lot/density rules; contents-handling or settlement-dissolution rules (open, backlog); intermediate decay staging; SAFE-001 mechanisms beyond removal's implication; GDD text edits (reserved controlled acts at filing time).

### Traceability
| Item | Source |
|---|---|
| Proposal given effect (as amended) | `proposals/TCIndustries_BLD-004_Abandonment_Placement_Proposal_Drafter_v1_2026-09-19.md` |
| Decision record | `proposals/TCIndustries_BLD-004_Ruling_Record_HD-BLD-02_2026-09-19.md` |
| Invariant approved | INV-BLD-004b (new candidate authored in the proposal; approved this ruling); INV-BLD-004a remains candidate |
| Composing rulings relied upon | HD-BLD-01 (free-claim formation + basic maintenance depth — the credit instrument's ruled context); INV-007a (approved — enforced at structure scale by this model) |
| Prior rulings affected | None |
```

## 4. Filing-time acts (reserved; owner/coordinator executes at register filing)

1. Paste §3 into `governance/TCIndustries_Human_Rulings_Register.md` (after the pending HD-RET-01/HD-BLD-01/HD-SCOPE-01/HD-ITM-01/HD-PROF-01 entries).
2. Annotate `proposals/TCIndustries_Invariant_Register_v0.1.md` → INV-BLD-004b "APPROVED as design invariant (HD-BLD-02, 2026-09-19)" (new register entry authored in the proposal); INV-BLD-004a recorded as candidate.
3. Master GDD status-patch line for BLD-004 (abandonment surface ruled; placement/lots/density remain TBD) — canonical edit, owner's act.
4. `governance/project_memory.md` event line. (No Open Questions Register row exists for BLD-004 — it is a GDD TBD, not an OQ; the tracker is the pointer record.)

## 5. Immediate consequences recorded by the drafter (proposals/ side)

- **Queue:** BLD-004 (abandonment) exits — 4 live items (within the 3–5 band). **Backlog 6d (salvage & destruction handling) is now the designated next draft**: the owner's removal model makes it load-bearing immediately (what zero-credit removal means for a structure's materials/contents is exactly the salvage surface; contents-handling is also open question 6e(a)). Drafts at the next cycle's start to restore 5/5.
- **Backlog 6e created:** zero-credit-removal follow-ups — (a) contents handling on removal; (b) settlement-institution dissolution rule. Small targeted rulings; 6e(a) may fold into the 6d draft as one composition.
- **OQ-009 interaction annotated:** the owner's maintenance-credit instrument gives OQ-009's "settlement maintenance" sink class a concrete ruled shape (maintenance credits as a sink candidate) — the live OQ-009 proposal's cross-option note stands, now with a ruled exemplar.
- **INV-007b relevance:** zero-credit removal recycles location capacity — the condition the (still-candidate) INV-007b would need if the owner ever revisits its approval.
- **Approved invariants this programme (per-invariant, filing-time annotations):** INV-016a, INV-016b, INV-007a, INV-015, INV-010, INV-011a, INV-BLD-004b. Candidates explicitly not approved: INV-007b, INV-011b, INV-011c, INV-BLD-004a.
- **Six ruled items now await governance filing** (HD-RET-01, HD-BLD-01, HD-SCOPE-01, HD-ITM-01, HD-PROF-01, HD-BLD-02).

---
*End of decision record. Authority = the owner's 2026-09-19 ruling as recorded in §1; the owner's verbatim sentence controls over all gloss.*
