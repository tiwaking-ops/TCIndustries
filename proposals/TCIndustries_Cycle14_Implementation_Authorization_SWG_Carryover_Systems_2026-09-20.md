# TCIndustries — Cycle 14 Implementation Authorization: SWG Carryover Systems

**Status:** HUMAN AUTHORIZATION (implementation gate opened by owner ruling, 2026-09-20). This document records that the owner has authorised implementation of three design surfaces as SWG carryover systems, within the bounds stated below. It is **not** a design ruling (the design rulings are the HD-ITM-02, HD-BLD-03, HD-PROF-02 records this session, plus the already-filed HD-ITM-01 and HD-BLD-02); it is the implementation gate that the project's standing rules require before any code work begins.
**Author:** Buffy — Codebuff agent (Freebuff), Cycle 14.
**Date:** 2026-09-20
**Assessor:** None — single-session authoring; no independent review pass (per GR-004, recorded as such rather than guessed).
**Controlling references:** owner's in-session rulings of 2026-09-20 (re-stated in the HD-ITM-02, HD-BLD-03, and HD-PROF-02 decision records this session); `governance/TCIndustries_Human_Rulings_Register.md`; `canonical/gdd/TCIndustries_Master_GDD_v1.1.1_Status_Patch.md`; `proposals/TCIndustries_Open_Questions_Register.md`; `proposals/TCIndustries_Invariant_Register_v0.1.md`; `investigations/TCIndustries_Design_Evidence_Register_v0.1.md`; `governance/eic/...HD-EIC-07...` (in force); HD-TST-01 (testbed reuse, patterns-only); AGENTS.md.
**Scope of this record:** authorizes implementation of the three ruled systems against the SWG Pre-CU fork (`sources/swg-pre-cu/craft-seed2-1-pro-preview/`) as the implementation base, within the explicit bounds below. Does **not** authorize: new EIC architecture, residual-channel invention, numeric tuning beyond the owner-authorized values, provenance-engine work, any GDD status change, any `canonical/` or `governance/` edit, or any change to ruled design.

---

## 1. What is authorized for implementation

The owner has authorized implementation of the following three ruled systems, scoped to the SWG Pre-CU fork as the base:

1. **Respec free-drop (HD-PROF-02):** skill tiers gained by spending skill points; dropping a class-skill tier returns the points spent on that tier. No fee, no time cost, no capability penalty beyond the tier drop. No 1/1 immortal crit-fail bug. (Implementation surface: skill-tier/point account + tier-drop point-return. No respec cost instrument beyond free-drop.)

2. **Item damage / repair / destruction (HD-ITM-02):**
   - No usage/wear decay (deferred — "item decay later").
   - Death while holding an item → 1% of max durability damage on that item (owner-authorized value, historically grounded).
   - Items are repairable (repair mechanics, cost, stations, profession structure open — not specified this session beyond repairability).
   - Repair critical-fail → no effect (the Pre-CU SWG 1/1 immortal crit-fail bug is explicitly ruled out).
   - Items may be destroyed only by the player "destroy item" menu action; death-damage does not destroy items.
   - Menu-destroy → straight removal from the world; no salvage now (a salvage system may be added later).

3. **Building contents on zero-credit removal (HD-BLD-03):**
   - Building maintenance credit reaches 0 → after a 24-hour period (owner-authorized value, historically grounded), the building is removed from the world, and **everything in the building is also removed** — items, vendors, everything.
   - No salvage, no recovery window, no restitution.
   - This does not alter HD-BLD-02's removal instrument (zero-credit → terminal structure removal; credit balance is recourse path); it only specifies the contents outcome.
   - Settlement-institution dissolution (backlog 6e(b)) is **not** included in this authorization — it remains backlog.

---

## 2. Owner-authorized numeric values (with stated basis)

Two numeric values are authorized for implementation by this record, each with an explicit historical basis stated by the owner:

| Value | Where used | Basis (as stated by owner) | Status |
|---|---|---|---|
| 1% of max durability damage on death while holding the item | HD-ITM-02 ruling 2 | Pre-CU SWG: items damaged ~1% on death while holding them | Owner-authorized, historically grounded — not an independently derived balance value. Implementation may use this figure; no other durability tuning is authorized. |
| 24-hour upkeep lapse period before building+contents removal | HD-BLD-03 ruling 2 | Pre-CU SWG building-maintenance behaviour | Owner-authorized, historically grounded — not an independently derived balance value. Implementation may use this figure; no other building-maintenance timing is authorized. |

**Boundary (per BAL-001 and HD-EIC-08):** these two values are authorized *because the owner authorized them with a stated basis in this session*. They are not invented by the drafter; they are not a balance-tuning pass; and no other numeric is authorized by this record. Any further durability timing, repair cost, repair rate, crafting-durability interaction, or building-maintenance cost/schedule is **not authorized** here — it remains either open (to be ruled later) or BAL-001-gated.

---

## 3. Implementation base and scope limits

**Implementation base:** the SWG Pre-CU fork at `sources/swg-pre-cu/craft-seed2-1-pro-preview/`. That tree is a Vite/React/TS crafting-core demo (types, crafting engine, resource gen, RNG, schematics, data/resourceTree, UI sections, StatBar). It does **not** currently contain: a durable inventory, a death system, an item-durability field, a repair system, a building-maintenance system, or building contents. Those are new surface to add against the existing types/engine, not ports of existing code in that tree.

**De-SWG discipline (HD-TST-01, in force):** the implementation must be a de-SWG carry-over — keep the *mechanical behaviour* the owner ruled (1% on-death damage, free-drop point return, 24h upkeep→remove-everything, menu-destroy→removal, crit-fail→no effect), but do **not** carry SWG lore, names, species, planets, faction identifiers, or protected content into TCIndustries. Concretely, the fork's `PlanetId` union of Star Wars planet names and any SWG-named resources/schematics are **out of scope** for this implementation; the carry-over is the ruled mechanics, not the SWG content.

**Out of scope for this authorization (even though related):**
- Usage/wear item decay (deferred per owner's ruling — "item decay later").
- The 1/1 immortal repair crit-fail bug (explicitly ruled out).
- Any salvage system on item destruction (none now; may come later — "we can implement one later").
- Repair cost, repair materials, repair stations, repair profession structure, or any SERV-002 service-list decision (open — not specified this session).
- Any durability ceiling, durability-to-stat mapping, or durability range beyond the authorized 1% death-damage parameter (BAL-001).
- Any building-maintenance cost, credit rate, or credit-funding mechanic beyond the ruled "credit balance is the recourse path; 0 after 24h → removal with contents" shape (BAL-001 / HD-BLD-01/02).
- Settlement-institution dissolution (backlog 6e(b) — not authorized here).
- Any respec point-return arithmetic beyond "drop tier → return points spent on that tier" (the acquisition model and any magnitude are deferred; this authorization covers only the free-drop behaviour ruled in HD-PROF-02).
- Any EIC architecture, residual channel, or interdependence mechanism (HD-EIC-05/07 in force).
- Any provenance-engine work (DEFERRED, HD-EIC-08).
- Any combat system changes beyond the death-damage hook specified (combat scope/risk is OQ-008, still TBD).

---

## 4. Sequencing note (recorded, not a commitment)

The three systems are authorized together as a carry-over package because they are interdependent in the ruled model:
- Death-damage, repair, and menu-destroy together define the **item lifecycle** surface (HD-ITM-02).
- Building upkeep and contents removal define the **structure-contents** surface (HD-BLD-03).
- Respec free-drop defines the **skill-tier account** surface (HD-PROF-02).

They do **not** depend on each other for implementation order, and the owner did not direct a sequence. If implementation is begun, the order is the implementer's choice subject to the bounds above; this record does not schedule or estimate.

---

## 5. Open implementation-detail questions (recorded for the implementer; not resolved by this record)

These were confirmed with the owner in-session and are closed as stated; any further detail not covered below is left to the implementer within the bounds:

1. **Repair critical-fail behaviour:** confirmed — no effect (fail, nothing happens). The 1/1 immortal bug is ruled out.
2. **Item menu-destroy result:** confirmed — straight removal, no salvage (simplicity; Pre-CU SWG had no salvage system for items; salvage may be added later).
3. **Implementation base:** confirmed — SWG Pre-CU fork at `sources/swg-pre-cu/craft-seed2-1-pro-preview/`.

**Left open (implementer should not invent; flag if implementation needs an answer):**
- Repair mechanic shape (cost, materials, station, profession gating) — open; SERV-002 / crafting / facility surfaces.
- Whether "holding an item" at death is a single equipped/carried item, an inventory item, or a set of items — implementer's scope decision within the ruled behaviour (1% max-durability damage to the held item(s)), bounded by the owner's intent (death-damage is a hit, not a wipe).
- How max durability is established per item (the ruling assumes a "max durability" property exists; the durability range/ceiling is not authorized here — flag if the implementation needs a durability model before the 1% rule can be expressed).
- Any UI for the "destroy item" menu action — implementer's scope, bounded by the ruled behaviour (menu action → straight removal).

---

## 6. Explicit non-authorizations (standing)

- **No numeric tuning beyond the two owner-authorized values** (1%, 24h). Any other number is not authorized by this record.
- **No new EIC architecture or residual channel** (HD-EIC-05/07 in force).
- **No provenance-engine design** (DEFERRED, HD-EIC-08; depth/visibility requirements only where ruled).
- **No `canonical/` or `governance/` edits** from this record or from the implementation scaffold (reserved controlled acts; owner/coordinator only).
- **No GDD status changes** — the design rulings are in the HD-ITM-02, HD-BLD-03, HD-PROF-02 records (and the earlier HD-ITM-01, HD-BLD-02); this record is the implementation gate only.
- **No promotion of SWG content** — carry-over is the ruled mechanics; SWG lore/names/species/planets/factions are out of scope (HD-TST-01; SWG-001/002).
- **No prototype evidence treated as design authority** (PROT-001 — the fork is evidence of implementability, not design authority; the design here is the owner's in-session rulings, not the fork's prior behaviour).

---

## 7. Filing-time acts (reserved; owner/coordinator executes at register filing)

This record itself is filed in `proposals/` by the drafter (hard boundary: no `governance/` writes). If implementation work is begun, the owner/coordinator updates:
- `governance/project_memory.md` — event line for the authorization.
- `governance/TCIndustries_Human_Rulings_Register.md` — this record is referenced from the ruling records (HD-ITM-02, HD-BLD-03, HD-PROF-02) as the implementation authorization; it does not itself paste into the register unless the owner directs (the register records rulings; this record is the authorization attachment).

---

*End of implementation authorization record. Authority = the owner's 2026-09-20 implementation authorization as recorded in §1; the design authority for the three systems remains in the HD-ITM-02, HD-BLD-03, and HD-PROF-02 decision records.*
