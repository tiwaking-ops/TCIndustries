# TCIndustries — OQ-010 Item Damage, Repair & Destruction Elaboration — Human Ruling Record (HD-ITM-02)

**Status:** HUMAN RULED 2026-09-20 (owner decision recorded this session). This
document is the decision record; authority attaches to the owner's explicit
ruling, and this record preserves it verbatim for filing.
**Filing state:** decision record filed in `proposals/` by the drafter
(hard boundary: no `governance/` writes). The ready-to-file register entry is
in §3; `governance/TCIndustries_Human_Rulings_Register.md` is updated only by
the owner/coordinator at filing time.
**Author:** Buffy — Codebuff agent (Freebuff), ruling-session recorder, Cycle 14.
**Date:** 2026-09-20
**Gives effect to:** owner's in-session ruling (re-stated verbatim in §1). This
record elaborates the OQ-010 / HD-ITM-01 surface; it does **not** amend
HD-ITM-01's "no baseline decay" determination — it specifies the
death-damage/repair/destruction behaviour that lives inside that ruled surface.
**Ruling ID:** **HD-ITM-02** — register's HD-<Area>-<NN> pattern (area: ITM;
second of area).

---

## 1. The ruling (as decided by the owner)

|| # | Question | Owner decision |
||---|---|---|
|| 1 | Does item decay exist? | **No baseline usage/wear decay.** Consistent with HD-ITM-01 (no baseline decay). The owner's stated reason: in Pre-CU SWG the usage-decay system that was supposed to exist never worked; item decay was never solved. Item decay is deferred — "we can deal with item decay later." This ruling does **not** reinstate decay. |
|| 2 | Death-damage behaviour | **When a player dies while holding an item, the item takes 1% of its max durability damage.** The 1% and the death-trigger are owner-authorized values, stated to be historically grounded in Pre-CU SWG (items damaged ~1% on death while holding them). This is a **damage hit**, not a decay system: it is a discrete event tied to death, not a time-based wear process. |
|| 3 | Repair behaviour (nominal) | **Items are repairable.** (The owner did not specify a full repair system this session; repairability is ruled as a property of items, with repair mechanics left for implementation within the authorized scope — see §2 and the authorization record.) |
|| 4 | Repair critical-fail behaviour | **Do NOT implement the 1/1 immortal crit-fail bug.** The Pre-CU SWG critical-fail-to-1/1-immortal outcome is recorded as a bug, not a feature. The owner ruled it out. On a repair critical fail, the default behaviour is **no effect** (fail, nothing happens) unless a different behaviour is later ruled. |
|| 5 | Destruction source | **Items may be destroyed only by player action** — specifically the "destroy item" menu option. Death-damage does **not** destroy items (ruling 2 + ruling 5 together: death damage is a hit, not a destruction path). When an item is destroyed via the menu, it is **removed from the world** (straight removal — no salvage; see ruling 6). |
|| 6 | Salvage on destruction | **No salvage system for item destruction (this session).** Straight removal. Reason: Pre-CU SWG did not have a salvage system for items. A salvage system may be added later; for now, simplicity. |

**Owner's reasons (recorded for the file, not canonised beyond the ruling):**
- Authentic historical accuracy to Pre-CU SWG game rules where the precedent exists.
- Simplicity — the simplest system consistent with the precedent.
- The old Pre-CU SWG design made repairing items pointless because items could be destroyed directly from the menu; ruling 5 + ruling 6 keep destruction as a player action and keep repair meaningful (repair restores durability; the item is not auto-destroyed by death or wear).

---

## 2. Relationship to existing rulings and invariants

- **HD-ITM-01 (2026-09-19) — no baseline decay:** this record is consistent with, and does not amend, HD-ITM-01. It specifies the death-damage/repair/destruction behaviour that lives *inside* the ruled "no decay; attrition comes from destruction/loss/obsolescence/upgrade" surface. Death-damage is a discrete loss/trauma event; repair is a service/player-action surface; menu-destroy is the only destruction path.
- **INV-010 (approved conditional, HD-ITM-01):** under the ruled model:
  - clause (c) (ownership remains meaningful across the item's life) — addressed by the fact that items are not silently destroyed by decay or death; destruction is a player action.
  - clause (d) (no irreplaceable provenance-bearing asset destroyed without player-meaningful recourse) — under this ruling, destruction is a player action (menu), not an automatic event, so the owner's own action is the recourse; the dead-accumulation path does not reach destruction. Recorded for the file: the owner did not ask for a salvage/recourse path on menu-destroy; ruling 6 records that as a deliberate simplicity choice (no salvage now; may come later). If the owner wants INV-010(d) read more tightly later, that is a future ruling.
- **24-hour building upkeep period and 1% damage are owner-authorized numeric values** with a stated historical basis; they are **not** independently derived balance values. Recorded as such in the authorization record (§2 of `proposals/TCIndustries_Cycle14_Implementation_Authorization_SWG_Carryover_Systems_2026-09-20.md`).

**Explicitly does not authorise:**
- Any usage/wear decay system (deferred; "item decay later" is a deferral, not a ruling on decay mechanics).
- The 1/1 immortal repair crit-fail bug (explicitly ruled out).
- Any salvage system on item destruction (none now; simplicity; may come later).
- Any repair cost, repair material, repair station, or repair profession structure (those remain open — SERV-002 service lists, crafting/facility surfaces; no repair mechanic is specified this session beyond "items are repairable").
- Any durability ceiling, durability range, or durability-to-stat mapping (BAL-001; the 1% figure is an owner-authorized event parameter, not a full durability tuning pass).
- Any implementation beyond what is covered by the Cycle-14 implementation authorization record (scaffold only; no auto-modification to the fork).
- Any change to the OQ-010 structural determination (HD-ITM-01 stands); any change to HD-ITM-01; any provenance-engine work (DEFERRED, HD-EIC-08).

---

## 3. Ready-to-file register entry (verbatim, for `governance/TCIndustries_Human_Rulings_Register.md`)

```markdown
## HD-ITM-02 — OQ-010 Item Damage, Repair & Destruction Elaboration

**Source:** Owner ruling session, 2026-09-20, on owner's in-session ruling;
elaborates the OQ-010 / HD-ITM-01 surface without amending HD-ITM-01's "no
baseline decay" determination. Decision record:
`proposals/TCIndustries_OQ-010_Item_Damage_Repair_Destruction_Elaboration_Record_HD-ITM-02_2026-09-20.md`.

**Status:** HUMAN-LOCKED (item damage/repair/destruction behaviour only). Does
not reinstate decay. Complements HD-EIC-01–09, HD-GDD-01, HD-TST-01,
HD-RET-01, HD-BLD-01, HD-SCOPE-01, HD-ITM-01, HD-BLD-02, HD-PROF-01,
HD-PROF-02, which remain in force. Supersedes nothing; elaborates HD-ITM-01.

### Decision

1. **NO USAGE/WEAR DECAY.** Consistent with HD-ITM-01. The Pre-CU SWG
   usage-decay system that was supposed to exist never worked; item decay was
   never solved. Item decay is deferred — "we can deal with item decay later."
2. **DEATH DAMAGE:** when a player dies while holding an item, the item takes
   1% of its max durability damage. 1% and the death-trigger are
   owner-authorized values, historically grounded in Pre-CU SWG (items damaged
   ~1% on death while holding them). This is a discrete hit, not a decay
   system.
3. **REPAIR:** items are repairable (repair mechanics, cost, stations, and
   profession structure remain open — SERV-002 / crafting / facility surfaces;
   no repair mechanic specified this session beyond repairability).
4. **REPAIR CRIT-FAIL:** do NOT implement the Pre-CU SWG 1/1 immortal
   crit-fail bug — that outcome is a bug, not a feature. On a repair
   critical fail, default behaviour is no effect (fail, nothing happens) unless
   a different behaviour is later ruled.
5. **DESTRUCTION SOURCE:** items may be destroyed only by player action — the
   "destroy item" menu option. Death-damage does not destroy items. When an
   item is destroyed via the menu, it is removed from the world.
6. **SALVAGE:** no salvage system for item destruction this session — straight
   removal. Reason: Pre-CU SWG did not have a salvage system for items; a
   salvage system may be added later. For now, simplicity.

**Explicitly does not authorise:** any usage/wear decay system; the 1/1
immortal repair crit-fail bug; any salvage system now; any repair cost/material/
station/profession structure; any durability ceiling/range/mapping beyond the
authorized 1% death-damage parameter; any implementation beyond the
Cycle-14 implementation authorization record; any change to HD-ITM-01; any
provenance-engine work (DEFERRED, HD-EIC-08).

### Traceability

|| Item | Source |
||---|---|
|| Ruling that elaborates HD-ITM-01 | Owner in-session ruling, 2026-09-20 (re-stated verbatim in this record §1) |
|| Prior durability ruling | HD-ITM-01 (`proposals/TCIndustries_OQ-010_Ruling_Record_HD-ITM-01_2026-09-19.md`) — no baseline decay; INV-010 approved conditional; this record elaborates inside that surface |
|| Invariant register state | `proposals/TCIndustries_Invariant_Register_v0.1.md`, OQ-010 entry — INV-010 approved conditional (clauses a–d); this record's model is consistent with that conditional |
|| Prior rulings affected | None (elaborates HD-ITM-01; amends nothing) |
```

---

## 4. Filing-time acts (reserved; owner/coordinator executes at register filing)

1. Paste §3 into `governance/TCIndustries_Human_Rulings_Register.md` (chronological entry, after the pending HD-PROF-02 entry).
2. Update `proposals/TCIndustries_Open_Questions_Register.md` §2 OQ-010 status cell → "RULED (no decay; death-damage/repair/destruction behaviour per HD-ITM-01 + HD-ITM-02) — see HD-ITM-01, HD-ITM-02" (register convention: pointer only; ruling remains the authority).
3. Annotate `proposals/TCIndustries_Invariant_Register_v0.1.md` OQ-010 entry → INV-010 "remains APPROVED in conditional form (HD-ITM-01, 2026-09-19); elaboration per HD-ITM-02 is consistent with that conditional" (no new invariant approved; the candidate invariant in the 6d proposal is not approved this session).
4. Master GDD status-patch line for OQ-010/ITM-002 (canonical edit — owner's act, per GDD change-control).
5. `governance/project_memory.md` event line.

---

## 5. Immediate consequences recorded by the drafter (proposals/ side)

- **HD-ITM-01 relationship:** HD-ITM-01 remains the structural determination ("no baseline decay"). HD-ITM-02 adds the concrete death-damage/repair/destruction behaviour inside that surface. The two records are read together; neither amends the other.
- **INV-010(d) recourse question (flagged, not resolved):** the owner did not direct a salvage/recourse path for the menu-destroy case; ruling 6 records "no salvage now; may come later." The 1% death-damage path does not reach destruction, so INV-010(d)'s "irreplaceable provenance-bearing asset destroyed without recourse" condition is not triggered by death — that is the intended reading under this model. If the owner later wants a stronger INV-010(d) reading, that is a future ruling.
- **6d proposal surface (partial resolution):** the 6d proposal's "item destruction/loss" surface is partially answered: death damage is a hit (not destruction), repair exists (cost/station/profession open), and the only destruction path is player menu action with no salvage. The 6d proposal's "contents on zero-credit removal" sub-surface is resolved separately by HD-BLD-03 (this session). The 6d proposal itself is not withdrawn; it remains the framing document, with its options partially consumed by these rulings.
- **Pre-CU SWG basis (recorded, not canonised):** the 1% death-damage figure, the "destroy item" menu action, and the no-salvage choice are anchored in the owner's stated Pre-CU SWG precedent. The Evidence Register's E-17 / E-18 references are cited as historical pattern evidence only (SWG-002), not as TCIndustries design authority.

---

*End of decision record. Authority = the owner's 2026-09-20 ruling as recorded in §1.*
