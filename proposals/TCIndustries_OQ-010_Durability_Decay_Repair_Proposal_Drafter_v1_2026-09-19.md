# TCIndustries — OQ-010 Durability, Decay & Repair — Structural Proposal

**Filename:** `TCIndustries_OQ-010_Durability_Decay_Repair_Proposal_Drafter_v1_2026-09-19.md`
**Status:** PROPOSED. No authority. Ruling-ready candidate only. Nothing here is
LOCKED, and no element of this proposal changes any GDD status.
**Author:** Buffy — Codebuff agent (Freebuff), autonomous-loop drafter, Cycle 1.
**Assessor:** None — single-session authoring; no independent review pass (per GR-004, recorded as such rather than guessed).
**Date:** 2026-09-19
**Queue position:** 1 of 5 (see `proposals/TCIndustries_Continuous_Queue_Tracker.md`)
**Pre-flight:** no prior proposal or investigation addresses OQ-010 durability design; the gap ID appears only as an invariant row and an evidence-gap entry (both cited below). No duplicate pass.

---

## Status wall

- Every option below is **PROPOSED** and requires an explicit HD ruling to advance.
- **Zero numeric values.** No decay rates, repair costs, durability ceilings, or
  time constants appear anywhere in this document (BAL-001, GDD §25:
  "Do not invent skill caps, spawn rates, prices, decay curves…"). Any future
  numeric layer requires separate BAL-001 authorisation.
- **SAFE-001 is HUMAN-LOCKED principle-only** (HD-GDD-01): ownership/transfer
  safeguards are required, mechanism open. Option C touches safeguards directly
  and proposes nothing — it only identifies the interaction surface.
- **PLR-003 is HUMAN-LOCKED principle-only** (HD-GDD-01): meaningful persistent
  ownership is locked; any decay model must not make ownership meaningless.
- The Provenance & Reputation **Engine remains DEFERRED** (HD-EIC-08; GDD §34
  phase 2); durability interacts with provenance only at the requirements level
  (INV-010b) — no engine design is proposed.
- Item durability/decay/repair is **TBD** (OQ-010, GDD §31; ITM-002, GDD §14:
  "Whether and how items decay, require repair, or have limited lifespan is TBD").
  Evidence register gap 2 confirms: "no durability/decay system exists in any
  build. INV-010 is untestable until either a design or a sim host exists"
  (`investigations/TCIndustries_Design_Evidence_Register_v0.1.md` §3).

## Anchors (controlling references)

| Anchor | What it locks/constrains |
|---|---|
| OQ-010 (GDD §31) | The open question this proposal structures |
| ITM-002 (GDD §14, PROPOSED/TBD) | Any durability system "must create demand for repair services and materials without becoming pure friction" |
| ITM-001 (GDD §14, PROPOSED) | Items support identity: quality, provenance, condition, creator association |
| INV-010 (Invariant Register v0.1, OQ-010 entry) | Candidate four-part invariant (a)–(d) — reproduced under Invariants below |
| SERV-001 / SERV-002 (GDD §17) | Repair is a LOCKED viable service career; services should reinforce interdependence |
| PIL-002 (GDD §6, LOCKED) | Products/services preserve meaningful association with creator |
| CRFT-001 (GDD §12, HUMAN-LOCKED principle-only) | Differentiated products — decay/repair must not erase differentiation |
| SAFE-001 (GDD §26, HUMAN-LOCKED principle-only) | Safeguards against fraud, abandoned assets, exploitative transfer |
| LOOP-002 (GDD §7, LOCKED) | Non-combat viability — repair demand must not be combat-gated |
| Evidence E-06 (Design Evidence Register §1.1) | Repair-adjacent service-XP sites observed in fork — pattern evidence only, no TCIndustries rule |

## Principle → Constraint → Invariant → Failure Condition → Test

- **Principle (PROPOSED):** where item permanence and wear exist, they generate
  ongoing, non-punitive demand for player repair services and materials
  (ITM-002), reinforcing interdependence (VIS-004/PIL-003) without confiscating
  ownership (PLR-003) or erasing identity (PIL-002).
- **Constraint (derives from LOCKED layer):** decay must never (i) make
  ownership meaningless, (ii) erase provenance, (iii) gate non-combat careers
  behind combat item loss, (iv) become pure friction with no service demand
  flowing to other players.
- **Invariant (candidate, from Invariant Register INV-010, quoted):** "*If decay
  exists, it is decoupled-enough from activity that all four hold: (a) decay
  creates repair demand flowing to other players (services/crafters), (b) item
  identity/provenance survives the repair cycle (PIL-002), (c) ownership
  remains meaningful across the item's life (no pre-ordained total loss), and
  (d) decay never destroys the only copy of an irreplaceable
  provenance-bearing asset without player-meaningful recourse."*
- **Failure condition:** decay that is pure attrition (no service demand);
  provenance erasure; de-facto ownership confiscation; repair demand that only
  combat generates.
- **Test (shape-level, stated now):** lifecycle walk-through of the ruled model
  against INV-010 (a)–(d) — checklist-class, runnable at design time; full
  behavioural testing requires an implementation or simulation host (none
  exists today; Evidence Register §3 item 2).

## Options (structural only; tradeoffs and failure modes stated)

### Option A — Universal decay with player repair (decay-for-demand)
All equipment-class items carry condition; use lowers it; player repairers and
crafted repair materials restore it.
- *Serves:* strongest SERV-002 interdependence loop; continuous crafter demand
  (RES-002 support: consumption keeps resource demand alive); ITM-002's
  "demand for repair services and materials" read at full strength.
- *Tradeoffs:* highest friction risk — must be tuned (later, under BAL-001) to
  avoid feeling like a tax; largest UI/telemetry surface; touches every item
  class.
- *Failure modes:* if repair is too cheap/absent → pure attrition (INV-010a
  fails); if too punishing → ownership feels confiscated (INV-010c fails);
  if repair erases history → provenance loss (INV-010b fails).
- *Non-canonical-list check:* none of the §6 exclusions are assumed by this option.

### Option B — Decay only for crafted/quality-tier items (differentiation-weighted)
Generic commodity items do not decay; crafted items bearing quality tiers,
experimentation outcomes, or provenance carry condition and repair cycles.
- *Serves:* targets decay exactly where PIL-002/CRFT-001 identity lives;
  protects commodity logistics from friction; smaller blast radius for a first
  ruling; keeps repair as a specialist-adjacent service (master-crafted goods
  deserve specialist care — PROPOSED framing, not canon).
- *Tradeoffs:* creates a two-class item economy that must be explained clearly
  to players; repair demand is narrower, so the SERV-002 loop is weaker than A;
  boundary rule ("what counts as identity-bearing") becomes a design object
  itself.
- *Failure modes:* if the boundary is gamed (players prefer decaying-free
  generics), crafted differentiation is undercut (CRFT-001 pressure); if the
  boundary is too broad, A's friction failures return.
- *Note:* this option proposes no tier structure — quality tiers themselves
  remain open under CRFT-001 (HD-GDD-01 explicitly leaves "whether Exceptional
  exists… how many tiers exist" undecided).

### Option C — No decay at baseline (permanence default)
Items do not decay; demand for repair/replacement comes from destruction,
loss, obsolescence, or voluntary upgrade instead.
- *Serves:* strongest PLR-003 ownership meaning; lowest friction; INV-010
  becomes vacuously satisfied on (c)/(d) and (a) must be satisfied by other
  demand systems.
- *Tradeoffs:* the ITM-002 text contemplates decay ("Whether and how items
  decay… is TBD" — it does not require it, but the repair-service economy
  named in SERV-001/SERV-002 must then find demand elsewhere); risk that
  crafter demand flattens over time as the item stock saturates — this is an
  ASSUMPTION-class risk (cf. AS-005's shape: demand generation may need
  redesign), not a demonstrated failure.
- *Failure modes:* saturated item stock erodes craft demand long-term (would
  surface as an INV-001-family/ECO pressure, testable only with a simulation
  host); salvage/destruction rules become the load-bearing safeguard surface
  (SAFE-001 interaction, unresolved here).
- *Status caution:* choosing C is a legitimate "none of the above" ruling on
  OQ-010's decay half; it does not close OQ-010's repair-economy half.

### Cross-option note (recorded, not decided)
Any option combining salvage, trade, or destruction must eventually route
through SAFE-001's open mechanism surface (HUMAN-LOCKED principle-only) —
e.g. exploit-safe transfer and abandonment handling. That mechanism decision
is separate from this proposal and is NOT resolved here.

## What this proposal does NOT resolve

- Whether decay exists (that is exactly the ruling requested) and any numeric
  parameter of it (BAL-001-gated; requires separate authorisation).
- The quality-tier/experimentation model decay would attach to (OQ-003/OQ-004,
  both open; OQ-003/OQ-004 EIC-coupled per the Invariant Register).
- SAFE-001 safeguard mechanisms (soul-binding, transfer restrictions — all
  explicitly non-canonical-by-default per Open Questions Register §6 and open
  under HD-GDD-01).
- Repair *profession* structure (SERV-002 service lists are TBD).
- Any implementation, simulation, or testbed work (no code proposed; no T-task
  executed or assumed approved).
- Provenance engine design (DEFERRED, HD-EIC-08).

## Ruling requested (HD-<Area>-<NN> slot reserved in the Human Rulings Register)

**Question for the owner — rule yes / no / amend:**

1. **Primary ruling:** Select the structural model for OQ-010 item durability:
   **A** (universal decay with player repair), **B** (decay only for
   identity-bearing crafted items), **C** (no baseline decay; demand via
   destruction/obsolescence), or **defer** OQ-010 (any of the three may also be
   amended).
2. **Invariant ruling (independent):** Do you approve candidate invariant
   INV-010 (a)–(d) as stated in the Invariant Register v0.1 as a design
   invariant (TEST-001 class) for whichever model is chosen? (Approval records
   it as a design invariant; it sets no numeric value.)
3. **Scope confirmation:** Confirm that any numeric decay/repair layer is
   deferred to a future BAL-001-authorised pass, so this ruling settles
   structure only.

*Every claim above cites its file and anchor. If the owner rules "defer", the
tracker's next-ranked gap (OQ-016 already queued) is unaffected; the queue
pulls forward per the loop rules.*

---
*End of proposal. PROPOSED — awaiting HD ruling. No status changed anywhere by this document.*
