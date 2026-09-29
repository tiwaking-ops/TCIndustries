# TCIndustries — OQ-010 Durability, Decay & Repair — Human Ruling Record (HD-ITM-01)

**Status:** HUMAN RULED 2026-09-19 (owner decision recorded this session). This
document is the decision record; authority attaches to the owner's explicit
ruling, and this record preserves it verbatim for filing.
**Filing state:** decision record filed in `proposals/` by the drafter
(hard boundary: no `governance/` writes). The ready-to-file register entry is
in §3; `governance/TCIndustries_Human_Rulings_Register.md` is updated only by
the owner/coordinator at filing time.
**Author:** Buffy — Codebuff agent (Freebuff), ruling-session recorder, Cycle 10.
**Date:** 2026-09-19
**Gives effect to:** `proposals/TCIndustries_OQ-010_Durability_Decay_Repair_Proposal_Drafter_v1_2026-09-19.md` (retained unaltered as provenance).
**Ruling ID:** **HD-ITM-01** — register's HD-<Area>-<NN> pattern (area: ITM,
items); register filing confirms or renumbers.

---

## 1. The ruling (three parts, as decided by the owner)

| # | Question | Owner decision |
|---|---|---|
| 1 | OQ-010 primary structural model | **Option C — No baseline decay (permanence default):** items do not decay; demand for repair/replacement comes from destruction, loss, obsolescence, or voluntary upgrade instead. Options A (universal decay) and B (identity-bearing-only decay) are not adopted in any form |
| 2 | Candidate invariant INV-010 (a)–(d) | **APPROVED** as a design invariant (TEST-001 class), in its conditional register form: "*If decay exists, it is decoupled-enough from activity that all four hold: (a) decay creates repair demand flowing to other players (services/crafters), (b) item identity/provenance survives the repair cycle (PIL-002), (c) ownership remains meaningful across the item's life (no pre-ordained total loss), and (d) decay never destroys the only copy of an irreplaceable provenance-bearing asset without player-meaningful recourse.*" Owner-approved reading recorded in §5a |
| 3 | Scope confirmation | **Confirmed — structure only:** no decay rate, repair cost, durability ceiling, or time constant is set by this ruling; any such numeric layer requires separate BAL-001 authorisation |

## 2. What this ruling explicitly does NOT authorise

- Any numeric parameter of durability, decay, repair, salvage, or destruction (BAL-001).
- Any decay mechanism (A and B are unruled alternatives retained in the proposal's provenance; a future ruling could introduce decay, and INV-010's conditional then binds it).
- SAFE-001 safeguard mechanisms: **under C, salvage/destruction handling becomes the load-bearing adjacent surface** (consequence recorded in §5b; the mechanism itself remains open under HD-GDD-01's principle-only lock).
- Quality-tier/experimentation structures decay would have attached to (OQ-003/OQ-004 open, EIC-coupled — untouched by this ruling).
- Repair *profession* structure (SERV-002 service lists remain TBD).
- Provenance-engine work (DEFERRED, HD-EIC-08 — INV-010(b) remains a requirements-level obligation only).
- Any implementation, simulation, or testbed work.
- GDD text edits and status-pointer updates (reserved controlled acts at register filing, §4).

## 3. Ready-to-file register entry (verbatim, for `governance/TCIndustries_Human_Rulings_Register.md`)

```markdown
## HD-ITM-01 — OQ-010 Durability/Decay/Repair Structural Model
**Source:** Owner ruling session, 2026-09-19, on `proposals/TCIndustries_OQ-010_Durability_Decay_Repair_Proposal_Drafter_v1_2026-09-19.md`; decision record `proposals/TCIndustries_OQ-010_Ruling_Record_HD-ITM-01_2026-09-19.md`.
**Status:** HUMAN-LOCKED (durability structure only). Complements HD-EIC-01–09, HD-GDD-01, HD-TST-01, HD-RET-01, HD-BLD-01, HD-SCOPE-01, which remain in force. Supersedes nothing.

### Decision
1. **PRIMARY — Option C (No baseline decay):** items do not decay; demand for repair/replacement comes from destruction, loss, obsolescence, or voluntary upgrade. Answers OQ-010 (GDD §31) at the structural layer. ITM-002's repair-demand obligation ("demand for repair services and materials without becoming pure friction") is carried by the destruction/obsolescence/salvage layer, not by a decay system.
2. **INVARIANT — INV-010 APPROVED** as a design invariant (TEST-001 class) per `proposals/TCIndustries_Invariant_Register_v0.1.md` (OQ-010 entry), in its conditional form ("If decay exists…"). Owner-approved reading: under the no-decay model the conditional is binding on any future decay introduction, and clause (a)'s repair-demand-to-other-players obligation attaches to whatever systems generate repair demand in the ruled model. Sets no numeric value.
3. **SCOPE — structure-only:** no durability/decay/repair numeric parameter is set; all such values require separate BAL-001 authorisation.

**Explicitly does not authorise:** any numeric parameter; any decay mechanism (Options A/B remain unruled alternatives); SAFE-001 salvage/destruction mechanisms (open — flagged as the load-bearing adjacent surface under this model); OQ-003/OQ-004 structures; repair profession lists (SERV-002); provenance-engine work (DEFERRED, HD-EIC-08); GDD text edits (reserved controlled acts at filing time).

### Traceability
| Item | Source |
|---|---|
| Proposal given effect | `proposals/TCIndustries_OQ-010_Durability_Decay_Repair_Proposal_Drafter_v1_2026-09-19.md` (retained as provenance) |
| Decision record | `proposals/TCIndustries_OQ-010_Ruling_Record_HD-ITM-01_2026-09-19.md` |
| Invariant approved | INV-010, conditional reading (Invariant Register v0.1, OQ-010 entry) |
| Prior rulings affected | None |
```

## 4. Filing-time acts (reserved; owner/coordinator executes at register filing)

1. Paste §3 into `governance/TCIndustries_Human_Rulings_Register.md` (chronological entry, after the pending HD-RET-01/HD-BLD-01/HD-SCOPE-01 entries).
2. Update `proposals/TCIndustries_Open_Questions_Register.md` §2 OQ-010 status cell → "RULED — see HD-ITM-01".
3. Annotate `proposals/TCIndustries_Invariant_Register_v0.1.md` OQ-010 entry → INV-010 "APPROVED as design invariant (HD-ITM-01, 2026-09-19; conditional reading per decision record §5a)".
4. Master GDD status-patch line for OQ-010/ITM-002 (canonical edit — owner's act).
5. `governance/project_memory.md` event line.

## 5. Immediate consequences recorded by the drafter (proposals/ side)

- **Queue:** OQ-010 exits (4 live) → the drafted-ahead OQ-009 currency/sinks
  proposal (backlog rank 8) is **formally queued**, restoring 5/5. Pipeline
  behind the queue is now empty until the next pre-draft.
- **5a — INV-010 owner-approved reading (recorded for filing-time
  annotation):** the invariant is approved *in its conditional form*. Under
  the ruled no-decay model: (c)/(d) bind unconditionally as ownership/
  recourse obligations on any destruction/loss system; (a) attaches to
  whatever systems generate repair demand in the ruled model (destruction,
  loss, obsolescence, voluntary upgrade); (b) binds the provenance
  requirements of any repair/restoration mechanism. If decay is ever
  introduced by a future ruling, the full conditional binds it as written.
- **5b — Priority adjacent gap surfaced (backlog 6d):** *salvage &
  destruction handling* — under C, destruction/loss is the sole recurring
  item-destruction source, and its rules are the load-bearing SAFE-001
  surface (the proposal's Option C failure-mode note, now operative). Must
  respect INV-010(c)/(d) and the non-canonical list (no soul-binding
  default). Ranked at the next re-rank.
- **5c — OQ-009 interaction annotated:** the drafted-ahead currency/sinks
  proposal's sink instrument class "repair/maintenance (whether repair
  demand exists depends on OQ-010)" is annotated resolved: decay-driven
  repair demand does not exist under HD-ITM-01; that sink class shrinks to
  destruction/loss-driven restoration, and its adoption remains a future
  structural + BAL-001 decision.
- **SERV-002 note:** repair careers remain LOCKED-viable (SERV-001); their
  demand sources are now destruction/loss/obsolescence/upgrade — service
  lists stay TBD.
- **Approved invariants this programme (per-invariant, for filing-time
  annotations):** INV-016a, INV-016b, INV-007a, INV-015, INV-010.
  INV-007b remains candidate.

---
*End of decision record. Authority = the owner's 2026-09-19 ruling as recorded in §1.*
