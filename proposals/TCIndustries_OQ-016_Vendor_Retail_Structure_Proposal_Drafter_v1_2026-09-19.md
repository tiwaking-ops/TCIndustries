# TCIndustries — OQ-016 Vendor & Retail Mechanics — Structural Proposal

**Filename:** `TCIndustries_OQ-016_Vendor_Retail_Structure_Proposal_Drafter_v1_2026-09-19.md`
**Status:** PROPOSED. No authority. Ruling-ready candidate only. Nothing here is
LOCKED, and no element of this proposal changes any GDD status.
**Author:** Buffy — Codebuff agent (Freebuff), autonomous-loop drafter, Cycle 1.
**Assessor:** None — single-session authoring; no independent review pass (per GR-004, recorded as such rather than guessed).
**Date:** 2026-09-19
**Queue position:** 2 of 5 (see `proposals/TCIndustries_Continuous_Queue_Tracker.md`)
**Pre-flight:** no prior proposal or investigation addresses OQ-016 vendor design; the gap ID appears only in invariant and evidence rows (cited below). No duplicate pass.

---

## Status wall

- Every option below is **PROPOSED** and requires an explicit HD ruling to advance.
- **Zero numeric values.** No listing fees, tax percentages, price bands, or
  transaction rates appear anywhere (BAL-001). Any fee/tax layer later requires
  BAL-001 authorisation and is explicitly out of scope here.
- **NPC substitution limits are EIC-adjacent.** ECO-001 ("NPC vendors, if
  present, must not make player production irrelevant", LOCKED) is cited as a
  bounding principle; the *NPC vendor design* itself is not EIC architecture,
  and no residual channel or interdependence mechanism is invented here
  (HD-EIC-07 untouched; Architecture C stays retired, HD-EIC-05).
- **Commerce Directory is on the non-canonical list** (Open Questions Register
  §6; GDD §36) — no option assumes it, and Option C explicitly re-opens the
  *generic* question of discovery surfaces without proposing that specific
  historical artifact.
- Single unified currency as default is non-canonical (same §6 list) — no
  option assumes any currency model (OQ-009 stays open).
- Vendor identity supports PIL-002 (LOCKED) but **no provenance-engine work is
  proposed** (DEFERRED, HD-EIC-08) — identity attachment here is structural
  bookkeeping, not a reputation engine.

## Anchors (controlling references)

| Anchor | What it locks/constrains |
|---|---|
| OQ-016 (GDD §31) | The open question this proposal structures |
| RET-001 (GDD §16, LOCKED) | Players must be able to operate shops, vendors, commercial spaces selling goods and services |
| RET-002 (GDD §16, PROPOSED) | Vendors support stock management, pricing, association with player/org identity; exact mechanics TBD |
| RET-003 (GDD §16, PROPOSED) | Commerce should support formation of commercial districts, recognised brands, trusted suppliers |
| ECO-001 (GDD §15, LOCKED) | Player-driven economy; NPC vendors must not make player production irrelevant |
| INV-016a/b (Invariant Register v0.1, OQ-016 entry) | Candidate invariants: identity persistence through the sale chain; player-supply not strictly dominated by NPC supply |
| PIL-002 (GDD §6, LOCKED) | Identity and reputation attach to products/services/providers |
| WRLD-001 / PLR-003 (HUMAN-LOCKED principle-only, HD-GDD-01) | Persistent world; meaningful persistent ownership — vendors are owned persistent assets |
| BLD-001 (GDD §19, LOCKED) | Player-owned structures — vendor placement ties to structures |
| E-02 phase-5 suite, E-09 (Design Evidence Register §1.1, §2) | Vendor flows, `AtomicPurchase`, append-only `market_records` observed working in fork — PROTOTYPE pattern evidence only, "identity attach demonstrable"; establishes implementability, not design |

## Principle → Constraint → Invariant → Failure Condition → Test

- **Principle (PROPOSED):** retail is the point where production meets
  identity: vendor mechanics must let player sellers be *known* (RET-002,
  PIL-002), keep player supply primary (ECO-001), and give location and
  aggregation social meaning (RET-003) without a mandatory central marketplace.
- **Constraint (derives from LOCKED layer):** no vendor design may (i) strip
  seller identity from goods in the sale chain (PIL-002 breach), (ii) let NPC
  supply dominate any player-producible good class (ECO-001 breach),
  (iii) require a single universal marketplace (anti-goals: pure market
  spreadsheet simulator, VIS-006; ECO-003 DERIVED), (iv) force all commerce
  through one payment medium (single-currency default is non-canonical).
- **Invariant (candidate INV-016a, quoted):** "*A vendor's goods remain
  attributable to their supplier/owner through the full sale chain, and vendor
  identity (who stocks it) is visible to buyers."*
- **Invariant (candidate INV-016b, quoted):** "*For every good class a player
  can produce, player-made supply is distinguishable from and not strictly
  dominated by NPC supply — NPC availability must not erase the market for the
  player version."*
- **Failure condition:** anonymous commodities (no brand/identity attach);
  NPC stock making any player production role pointless; a global
  one-price market with no local character (ECO-002 collapse).
- **Test (shape-level):** retail-flow audit of the ruled design against
  INV-016a/b (checklist-class); INV-016b is untested behaviourally anywhere —
  no build contains NPC supply competing with player supply (Evidence
  Register §3 item 5) — full testing requires a sim host or implementation.

## Options (structural only; tradeoffs and failure modes stated)

### Option A — Identity-anchored player vendors (player-only retail)
All sell-side vendor capacity is player-owned and player-stocked; every vendor
bears its owner's identity; NPC commerce limited to baseline services (e.g.
trait-style conveniences) that never sell player-producible goods.
- *Serves:* RET-001 at full strength; INV-016a trivially satisfiable; ECO-001
  inviolable by construction; strongest RET-003 brand/district incentives
  (players must gather commerce themselves).
- *Tradeoffs:* commerce liquidity depends entirely on player population
  distribution; new/returning players face discovery friction until commercial
  density exists; NPC settlements (WRLD-003) lose a convenient "baseline
  supply" role, which must be deliberately designed around.
- *Failure modes:* ghost-market problem (vendors placed but unstocked) needs
  abandonment handling (SAFE-001 surface); if density never emerges in a
  region, regional economies (ECO-002) starve — detectable via the
  observability surface (INV-012/T-07-informed, not designed here).

### Option B — Hybrid NPC convenience + player primacy
NPC vendors sell a *baseline convenience band* of generic goods; anything
crafted, differentiated, or quality-bearing comes only from player vendors;
NPC stock is explicitly the floor, never the ceiling, of the economy.
- *Serves:* softer early-game and low-population-hour experience; keeps RET-001
  meaningful (all interesting goods are player-made); INV-016b expressible as a
  hard design rule ("NPC stock never covers crafted/differentiated classes").
- *Tradeoffs:* the A/B boundary ("what is generic") becomes a recurring design
  object; risk of NPC convenience crowding out nascent player supply in the
  generic band — the exact ECO-001 pressure, here bounded but not eliminated;
  INV-016b needs the boundary rule to be enforceable, not aspirational.
- *Failure modes:* boundary creep (more goods reclassified "generic" over time
  under convenience pressure) silently hollows RET-001 — would need the
  INV-016b audit as a standing check; dual-economy arbitrage exploits around
  the boundary (SAFE-002 surface, unresolved).

### Option C — Minimal open-vendor framework (mechanism-light)
The GDD locks only the *capability* (RET-001); this option defers all vendor
*mechanism* structure: define the minimal vendor object (identity + stock +
pricing) and leave discovery, districts, NPC competition, and fees entirely
open for later rulings.
- *Serves:* smallest ruling possible; unblocks downstream work (cities,
  structures, economy plumbing) that only needs "player vendors exist with
  identity" to proceed; consistent with the "smallest ruling that unblocks
  the most downstream work" ranking principle.
- *Tradeoffs:* RET-003 (districts/brands) gets no structural support and may
  need a later dedicated pass; leaves ECO-001/INV-016b unanswered in practice
  (no NPC competition defined, so nothing to audit); risks repeated reopening
  of adjacent questions.
- *Failure modes:* downstream systems (OQ-007 cities, OQ-015 scope) inherit an
  under-specified retail layer and must make local assumptions — the exact
  silent-assumption drift the status discipline exists to prevent; mitigated
  only by recording the assumptions per system (as this loop's proposals do).

## What this proposal does NOT resolve

- Any fee, tax, price-band, or transaction-rate question (BAL-001-gated,
  separately authorised; also interacts with OQ-009 currency, open).
- Whether NPC vendors exist at all, and their exact stock rules (that is
  precisely the Option A/B/C ruling; Option C defers it entirely).
- Discovery surfaces (search, directories, bulletin boards) — the Commerce
  Directory is non-canonical by default; any successor surface would need its
  own proposal and ruling.
- Commercial-district mechanics (RET-003 support is option-dependent; zoning
  ties to OQ-007 cities, queued separately).
- Provenance *engine* work (DEFERRED, HD-EIC-08); vendor identity here is
  ownership/labelling structure only.
- Any implementation, simulation, or testbed work; E-09 patterns are cited as
  evidence of implementability only (PROT-001: prototype is evidence, not
  design authority).

## Ruling requested (HD-<Area>-<NN> slot reserved in the Human Rulings Register)

**Question for the owner — rule yes / no / amend:**

1. **Primary ruling:** Select the structural model for OQ-016 vendor/retail:
   **A** (player-only vendor capacity), **B** (hybrid NPC convenience band +
   player primacy), **C** (minimal vendor object; mechanism questions deferred),
   or **defer** (amendments welcome on any option).
2. **Invariant ruling (independent):** Do you approve candidate invariants
   INV-016a and INV-016b as design invariants (TEST-001 class) for whichever
   model is chosen? (If Option C is chosen, INV-016b is recorded as open
   rather than approved-with-model.)
3. **Scope confirmation:** Confirm that fees/taxation/price-band numerics are
   deferred to a future BAL-001-authorised pass, keeping this ruling
   structure-only.

*Every claim above cites its file and anchor. On any ruling, the tracker updates
and the queue maintains 3–5 live items per the loop rules.*

---
*End of proposal. PROPOSED — awaiting HD ruling. No status changed anywhere by this document.*
