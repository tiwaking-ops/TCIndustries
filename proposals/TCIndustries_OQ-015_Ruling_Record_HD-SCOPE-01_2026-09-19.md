# TCIndustries — OQ-015 First-Playable Scope — Human Ruling Record (HD-SCOPE-01)

**Status:** HUMAN RULED 2026-09-19 (owner decision recorded this session). This
document is the decision record; authority attaches to the owner's explicit
ruling, and this record preserves it verbatim for filing.
**Filing state:** decision record filed in `proposals/` by the drafter
(hard boundary: no `governance/` writes). The ready-to-file register entry is
in §3; `governance/TCIndustries_Human_Rulings_Register.md` is updated only by
the owner/coordinator at filing time.
**FILED 2026-09-19:** §3 pasted verbatim into the governance register
(**Section C**) at the owner's explicit instruction; OQ Register OQ-015
pointer updated to "RULED — see HD-SCOPE-01"; Invariant Register INV-015
annotated APPROVED (per-invariant). Execution record:
`investigations/TCIndustries_Governance_Filing_Execution_HD-RET-01_HD-BLD-01_HD-SCOPE-01_2026-09-19.md`.
GDD status line and `project_memory` event line remain reserved owner acts
(not executed by the drafter).
**Author:** Buffy — Codebuff agent (Freebuff), ruling-session recorder, Cycle 8.
**Date:** 2026-09-19
**Gives effect to:** `proposals/TCIndustries_OQ-015_First_Playable_Scope_Proposal_Drafter_v1_2026-09-19.md` (retained unaltered as provenance).
**Ruling ID:** **HD-SCOPE-01** — register's HD-<Area>-<NN> pattern (area:
SCOPE, production planning); register filing confirms or renumbers.

---

## 1. The ruling (four parts, as decided by the owner)

| # | Question | Owner decision |
|---|---|---|
| 1 | OQ-015 primary scope | **Option S1 — Economic spine:** first playable = discover → acquire → process → craft → sell via player vendors, plus one player service (e.g. medical) generating recurring demand. Out of scope (deferred-TBD): combat, cities, organisations, transport, factories |
| 2 | S1 early-game economic on-ramp (shop viability before player commercial density exists) | **Recorded as a named-open question** at scope level (see §5a); resolved by a later targeted ruling — NOT decided inline by this ruling |
| 3 | Candidate invariant INV-015 | **APPROVED** as a design invariant (TEST-001 class), binding all future scope proposals: any "first playable" must include enough economy for at least one interdependence relation to be real |
| 4 | Scope confirmation | **Confirmed:** systems named out-of-scope are recorded as deferred-TBD (not assumed absent from the design); no schedule or numeric commitment is implied by this ruling |

## 2. What this ruling explicitly does NOT authorise

- Any schedule, effort estimate, or numeric commitment (none exists in the
  proposal; none is created by this ruling).
- Factories/manufacturing (MFG-001 LOCKED guard held; OQ-005 open,
  EIC-coupled), combat (OQ-008), city governance instruments (ruled
  formation/depth exist at structure level — HD-BLD-01 — but cities are not
  in the first-playable slice), organisations (OQ-013), transport (OQ-006),
  currency choice (OQ-009) — all deferred-TBD, per the confirmation.
- The skill-acquisition model: the slice's acquisition machinery ships as a
  **PROTOTYPE-labelled placeholder** until the live skill-acquisition queue
  item is ruled; nothing about the placeholder pre-selects that model.
- Any interdependence mechanism (INV-015 approved as invariant — a scope
  test, not an architecture; HD-EIC-07 untouched, HD-EIC-05 respected).
- Any implementation work — scope structure only; build tasks are separate
  future approvals (Gap Analysis §7.2 mechanics).
- S2/S3 are not adopted in any form; they remain unruled alternatives in the
  proposal's provenance.

## 3. Ready-to-file register entry (verbatim, for `governance/TCIndustries_Human_Rulings_Register.md`)

```markdown
## HD-SCOPE-01 — OQ-015 First-Playable Scope
**Source:** Owner ruling session, 2026-09-19, on `proposals/TCIndustries_OQ-015_First_Playable_Scope_Proposal_Drafter_v1_2026-09-19.md`; decision record `proposals/TCIndustries_OQ-015_Ruling_Record_HD-SCOPE-01_2026-09-19.md`.
**Status:** HUMAN-LOCKED (first-playable scope structure only). Complements HD-EIC-01–09, HD-GDD-01, HD-TST-01, HD-RET-01, HD-BLD-01, which remain in force. Supersedes nothing.

### Decision
1. **PRIMARY — Option S1 (Economic spine):** the first playable comprises resource discovery → acquisition → processing → crafting → sale via player vendors, plus one player service (e.g. medical) generating recurring demand. Systems out of scope are recorded as deferred-TBD: combat, cities, organisations, transport, factories. Answers OQ-015 (GDD §31) at the structural layer. Satisfies INV-015: craft→vendor→buyer and service demand are real, unavoidable between-player relations in normal play.
2. **NAMED-OPEN — Early-game economic on-ramp:** how vendor-based shop viability is established before player commercial density exists (no NPC goods supply exists under HD-RET-01's player-only-vendor structure) is recorded as an explicit open question at scope level, for a later targeted ruling. Not decided by this ruling.
3. **INVARIANT — INV-015 APPROVED** as a design invariant (TEST-001 class) per `proposals/TCIndustries_Invariant_Register_v0.1.md` (OQ-015 entry), binding all future scope proposals.
4. **SCOPE — deferred-TBD and no commitments:** excluded systems are not assumed absent from the design; no schedule, effort estimate, or numeric value is implied or authorised. The slice's skill-acquisition machinery ships as a PROTOTYPE-labelled placeholder pending the skill-acquisition ruling. No interdependence mechanism is selected (HD-EIC-05/07 in force).

**Explicitly does not authorise:** any schedule/effort/numeric commitment; factories (MFG-001; OQ-005 open), combat (OQ-008), organisations (OQ-013), transport (OQ-006), currency (OQ-009) — all deferred-TBD; any interdependence architecture selection; any implementation work (separate future approvals).

### Traceability
| Item | Source |
|---|---|
| Proposal given effect | `proposals/TCIndustries_OQ-015_First_Playable_Scope_Proposal_Drafter_v1_2026-09-19.md` (retained as provenance) |
| Decision record | `proposals/TCIndustries_OQ-015_Ruling_Record_HD-SCOPE-01_2026-09-19.md` |
| Invariant approved | INV-015 (`proposals/TCIndustries_Invariant_Register_v0.1.md`, OQ-015 entry) |
| Composing rulings relied upon | HD-RET-01 (OQ-016); HD-BLD-01 (OQ-007) — neither amended |
| Prior rulings affected | None |
```

## 4. Filing-time acts (reserved; owner/coordinator executes at register filing)

1. Paste §3 into `governance/TCIndustries_Human_Rulings_Register.md` (chronological entry, after the pending HD-RET-01/HD-BLD-01 entries).
2. Update `proposals/TCIndustries_Open_Questions_Register.md` §2 OQ-015 status cell → "RULED — see HD-SCOPE-01".
3. Annotate `proposals/TCIndustries_Invariant_Register_v0.1.md` OQ-015 entry → INV-015 "APPROVED as design invariant (HD-SCOPE-01, 2026-09-19)".
4. Master GDD status-patch line for OQ-015 (canonical edit — owner's act).
5. `governance/project_memory.md` event line.

## 5. Immediate consequences recorded by the drafter (proposals/ side)

- **Queue:** OQ-015 exits (4 live) → backlog priority 6b (BLD-004 abandonment,
  pre-drafted Cycle 7) is **formally queued**, restoring 5/5. No new drafting
  required — the pre-draft becomes a live item.
- **5a — Named-open question created by this ruling:** *"Early-game economic
  on-ramp: how is vendor-based shop viability established before player
  commercial density exists, given no NPC goods supply (HD-RET-01)?"* —
  recorded in the tracker backlog (6c); interacts with OQ-009 (currency) and
  the observability surface; small, targeted, future one-item ruling.
- **Placeholder machinery:** the slice's acquisition placeholder stays
  PROTOTYPE-labelled until the skill-acquisition item is ruled — that queue
  item's ruling now has a concrete downstream consumer.
- **INV-015 is now the fourth per-invariant approval of this programme**
  (INV-016a, INV-016b, INV-007a, then INV-015, in session order)
  *[count corrected at the Cycle-13 consistency audit — originally misstated
  as "third … four, counting RET's pair"; per-invariant counting gives four]*.
  Filing-time annotations must remain per-invariant (dependency note §9).

---
*End of decision record. Authority = the owner's 2026-09-19 ruling as recorded in §1.*
