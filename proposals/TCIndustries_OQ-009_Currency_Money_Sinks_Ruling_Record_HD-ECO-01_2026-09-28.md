# TCIndustries — OQ-009 Currency & Money Sinks — Human Ruling Record (HD-ECO-01)

**Status:** HUMAN RULED (PARTIAL) 2026-09-28 (owner decisions recorded this
session). Structure and INV-009a are ruled; INV-009b and the scope confirmation
are **NOT** ruled and are recorded as open. This document preserves the owner's
decisions verbatim for filing.
**Filing state:** decision record filed in `proposals/` by the recorder
(Freebuff/OpenCode, per owner's instruction). Ready-to-file register entry in §3.
**Author:** OpenCode (space-bunny-free) — ruling-session recorder, 2026-09-28.
**Assessor:** none (self-recorded; no independent assessment).
**Gives effect to:** `proposals/TCIndustries_OQ-009_Currency_Money_Sinks_Proposal_Drafter_v1_2026-09-19.md` (retained unaltered as provenance).
**Ruling ID:** **HD-ECO-01** — register's HD-<Area>-<NN> pattern (area: ECO,
economy). ID confirmed free at filing.

---

## 1. The ruling (as decided by the owner)

| # | Question | Owner decision |
|---|---|---|
| 1 | OQ-009 primary structural model | **Structure A — Single unified currency.** Owner's stated reason, verbatim: *"Historicity. SWG only had credits as a currency."* |
| 2 | Candidate invariant INV-009a | **APPROVED**, with the **faucet-exceeds-sink mandate** reading. Owner's stated reason, verbatim: *"Historicity. SWG always had a problem with not enough money sinks. Players became money sinks by hoarding money."* |
| 3 | Scope confirmation | **NOT RULED — OPEN.** The three sub-confirmations (sink parameters to BAL-001; sink instrument classes named-open; backlog 6c early-game on-ramp remains separate) were not answered this session. |
| 4 | Candidate invariant INV-009b | **NOT RULED.** Note: INV-009b binds only *if multiple currencies are adopted*. Under the ruled single-currency structure it is **moot on the current design**; it remains available as a standing constraint on any future multi-currency decision. No approval or decline is recorded either way. |

### 1a — Non-canonical-list promotion (required by Structure A, recorded here)

Structure A requires an **explicit promotion recorded in the ruling**. The Master
GDD's Historical non-canonical boundary list contains the entry **"Single unified
currency as the default"** (`canonical/gdd/TCIndustries_Master_GDD_v1.1.1_Status_Patch.md:1001`).

**PROMOTED to canonical design by HD-ECO-01, 2026-09-28, on the owner's explicit
ruling:** a single unified currency is the design's medium. The non-canonical
entry is struck from the historical boundary list and replaced with a pointer to
this ruling. This promotion is a **direct consequence of decision 1** and is
recorded as part of it, not as a separate decision. No other item on the
non-canonical list is promoted.

### 1b — INV-009a reading, recorded precisely

The candidate invariant was approved **in the faucet-exceeds-sink mandate
reading**, not the detectability reading. The two readings were distinguished in
the proposal because the single sentence admits both. As ruled, the commitment is
the **stronger** one:

> Total faucet flow exceeds sink flow in every bounded observation window **by
> design**, by a measurable and monitorable margin that the observability system
> can detect and attribute — **inflation cannot proceed silently.**

Recorded consequence: this is a standing economic **policy**, not only a
monitoring requirement. It commits the design to deliberate net money growth
relative to sinks. The owner's stated basis is that Pre-CU SWG persistently
suffered from insufficient money sinks, with hoarding players acting as money
sinks.

**Interaction with prior rulings, recorded not resolved:** under HD-RET-01
(player-only vendors; NPC commerce limited to baseline services) money faucets
are predominantly player creation — crafting, services, harvesting. This ruling
and HD-RET-01 are complementary rather than in tension. **No numeric faucet/sink
values are set** (BAL-001).

**Interaction with HD-PROF-03, recorded:** the ruled trainer-gated acquisition
model charges credits for class training. The faucet-exceeds-sink mandate is the
condition under which that training gate remains affordable. The two rulings
compose deliberately; neither implies the other.

## 2. What this ruling explicitly does NOT authorise

- Any numeric value: faucet rates, sink rates, margins, thresholds, credit
  quantities, or the 1000-credit figure cited as SWG precedent (BAL-001).
- Any specific sink instrument class. Sink instruments remain named-open; each
  requires its own future ruling.
- **Backlog 6c** — the early-game economic on-ramp question named open by
  HD-SCOPE-01 §5a. This ruling does not answer it and does not reduce it.
- Any medium other than the single unified currency. Multi-currency was not
  chosen; INV-009b is therefore moot on the current design and unanswered.
- Any Provenance & Reputation work (DEFERRED, HD-EIC-08).
- Any EIC shaping (HD-EIC-05/07 in force — the HD-EIC-07 gate is untouched and
  remains open).
- Any implementation, simulation, or testbed work.
- Approval of the scope-confirmation sub-items (decision 3 above).

## 3. Ready-to-file register entry (verbatim, for the Human Rulings Register)

```markdown
## HD-ECO-01 — OQ-009 Currency & Money Sinks (PARTIAL: structure + INV-009a; scope and INV-009b open)

**Source:** Owner ruling session, 2026-09-28, on
`proposals/TCIndustries_OQ-009_Currency_Money_Sinks_Proposal_Drafter_v1_2026-09-19.md`;
decision record `proposals/TCIndustries_OQ-009_Currency_Money_Sinks_Ruling_Record_HD-ECO-01_2026-09-28.md`.

**Status:** HUMAN-LOCKED (currency structure + the faucet-exceeds-sink policy
only). Complements HD-EIC-01–09, HD-GDD-01, HD-TST-01, HD-RET-01, HD-BLD-01,
HD-SCOPE-01, HD-ITM-01, HD-BLD-02, HD-PROF-02, HD-ITM-02, HD-BLD-03, which
remain in force. Supersedes nothing.

### Decision

1. **PRIMARY — Structure A (single unified currency).** Owner's stated reason,
   verbatim: *"Historicity. SWG only had credits as a currency."* Answers OQ-009
   (GDD §31) at the structural layer.
2. **NON-CANONICAL PROMOTION.** The Master GDD historical non-canonical list
   entry *"Single unified currency as the default"* is **PROMOTED to canonical
   design** by this ruling and struck from that list, replaced by a pointer here.
   This is a direct consequence of decision 1, recorded as part of it. No other
   non-canonical item is promoted.
3. **INVARIANT — INV-009a APPROVED**, in the **faucet-exceeds-sink mandate**
   reading: total faucet flow exceeds sink flow in every bounded observation
   window **by design**, by a measurable and monitorable margin the observability
   system can detect and attribute; inflation cannot proceed silently. Owner's
   stated reason, verbatim: *"Historicity. SWG always had a problem with not
   enough money sinks. Players became money sinks by hoarding money."* This is a
   standing economic policy, not only a monitoring requirement. Sets no numeric
   value.
4. **SCOPE — NOT CONFIRMED, REMAINS OPEN.** The three sub-confirmations (sink
   parameters to BAL-001; sink instrument classes named-open, each by its own
   future ruling; backlog 6c separate and unanswered) were not ruled this
   session. Default boundaries nonetheless remain in force: numerics are
   BAL-001-gated, and no sink instrument class is authorised by this ruling.
5. **INVARIANT — INV-009b NOT RULED, MOOT ON THE CURRENT DESIGN.** It binds only
   if multiple currencies are adopted; a single currency was ruled. It remains
   available as a standing constraint on any future multi-currency decision.
   Neither approved nor declined.

**Explicitly does not authorise:** any numeric value (BAL-001, including the
1000-credit SWG precedent figure); any specific sink instrument class; backlog 6c
(remains open per HD-SCOPE-01); any other currency medium; provenance or
reputation work (DEFERRED, HD-EIC-08); any EIC shaping (HD-EIC-07 gate untouched
and open); any implementation, simulation, or testbed work.

### Traceability
| Item | Source |
|---|---|
| Proposal given effect | `proposals/TCIndustries_OQ-009_Currency_Money_Sinks_Proposal_Drafter_v1_2026-09-19.md` (retained as provenance) |
| Decision record | `proposals/TCIndustries_OQ-009_Currency_Money_Sinks_Ruling_Record_HD-ECO-01_2026-09-28.md` |
| Invariant approved | INV-009a, faucet-exceeds-sink mandate reading (Invariant Register v0.1, OQ-009 entry) |
| Non-canonical entry promoted | "Single unified currency as the default" — Master GDD §36 historical boundary list |
| Composing rulings relied upon | HD-RET-01 (player-only vendors, NPC baseline only — the faucet base); HD-PROF-03 (trainer-gated acquisition — a credit-consuming sink this ruling keeps affordable) |
| Prior rulings affected | None |
```

## 4. Filing-time acts (reserved; owner/coordinator executes at register filing)

1. Paste §3 into `governance/TCIndustries_Human_Rulings_Register.md` as **Section E**, after Section D.
2. Update `proposals/TCIndustries_Open_Questions_Register.md` §2 OQ-009 status cell → "RULED (Structure A single unified currency; INV-009a approved, faucet-exceeds-sink mandate) — see HD-ECO-01".
3. Annotate `proposals/TCIndustries_Invariant_Register_v0.1.md` OQ-009 entry → INV-009a "APPROVED as design invariant (HD-ECO-01, 2026-09-28; faucet-exceeds-sink mandate reading)"; INV-009b recorded as not-ruled and moot under a single currency.
4. **Master GDD status line** for OQ-009 (canonical edit — owner's act), **and** strike "Single unified currency as the default" from the §36 historical non-canonical boundary list, replacing it with a pointer to HD-ECO-01.
5. Update `proposals/TCIndustries_Continuous_Queue_Tracker.md` — OQ-009 exits the live queue; add a RULED line.
6. `governance/project_memory.md` event line.
7. Run the **GR-006.5** mechanical check: confirm HD-ECO-01 is findable in this register and in every dependent document named in GR-006.2.

## 5. Immediate consequences recorded by the recorder

- **Queue:** OQ-009 exits the live queue. The queue is reduced to skill
  acquisition (now ruled as HD-PROF-03) and multi-account/exploit, i.e. **0 live
  items** against the tracker's 3–5 band. Per the tracker's own method this
  requires new backlog items to be drafted to restore the band.
- **6c interaction:** the early-game on-ramp question is *more* tractable under
  this ruling, because a deliberate faucet-exceeds-sink policy guarantees money
  supply growth in the early game without requiring NPC goods supply (which
  HD-RET-01 forbids). The question is **not** answered by this ruling; it is
  annotated as no longer blocked on money-supply growth.
- **6d interaction:** the mandatory faucet reading strengthens the case for
  consumption-side sinks. It does not select any sink instrument class, and it
  does not reopen the ruled no-salvage position (HD-ITM-02 ruling 6, HD-BLD-03
  ruling 1).
- **HD-RET-01 interaction:** no conflict. Player-only vendors remain the
  structural model; the mandate governs aggregate flow, not who may trade.
- **Still open from this proposal:** scope confirmation (3 sub-items) and
  INV-009b. Both recorded as unruled.

---
*End of decision record. Authority = the owner's 2026-09-28 rulings as recorded in §1.*
