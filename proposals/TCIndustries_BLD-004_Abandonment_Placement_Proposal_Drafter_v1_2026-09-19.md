# TCIndustries — BLD-004 Abandonment & Placement Handling — Structural Proposal

**Filename:** `TCIndustries_BLD-004_Abandonment_Placement_Proposal_Drafter_v1_2026-09-19.md`
**Status:** PROPOSED. No authority. Ruling-ready candidate only. Nothing here is
LOCKED, and no element of this proposal changes any GDD status.
**Author:** Buffy — Codebuff agent (Freebuff), autonomous-loop drafter, Cycle 7.
**Assessor:** None — single-session authoring; no independent review pass (per GR-004, recorded as such rather than guessed).
**Date:** 2026-09-19
**Queue status:** **live queue item** (pre-drafted Cycle 7 as backlog priority 6b; formally queued Cycle 8 into OQ-015's freed slot). See `proposals/TCIndustries_Continuous_Queue_Tracker.md`.
**Pre-flight:** BLD-004 appears in the GDD (§19, TBD: "Structure placement rules, lot systems, density limits, and abandonment handling are unresolved"), in the Gap Analysis backlog table (city-formation row), and in HD-BLD-01 §2's consequence note (priority 6b). No prior proposal or investigation structures BLD-004 as a decision. No duplicate pass.

---

## Status wall

- Every option below is **PROPOSED** and requires an explicit HD ruling to advance.
- **Zero numeric values.** No vacancy timers, decay periods, reclaim fees, or
  density caps (BAL-001). Every option is stated so its parameters are
  later-defined instrument *classes*, not values.
- **Composes with ruled structure.** HD-BLD-01 (2026-09-19) ruled free-claim
  settlement formation (Option A) with basic governance depth (ownership +
  zoning presence + shared maintenance responsibility) and approved INV-007a
  ("City governance mechanisms require ongoing player participation to
  persist — no city remains fully functional indefinitely with zero player
  involvement"). This proposal implements the abandonment *mechanism* that
  keeps free-claim formation compatible with that approved invariant.
- **PLR-003 is HUMAN-LOCKED principle-only** (HD-GDD-01): meaningful
  persistent ownership is locked; any abandonment mechanism must degrade
  *function* without confiscating *title* except where the owner is offered a
  player-meaningful recourse path first (SAFE-001's safeguard obligation
  binds throughout).
- **Abandonment ≠ punishment.** Nothing here targets inactive players;
  every option governs *asset state under vacancy*, identity-neutral by
  construction (the same discipline the Cycle-6 multi-account proposal
  applies: pattern/state, not actor).
- **BLD-001 (LOCKED)** requires player-owned structures with persistent
  presence; **BLD-004 itself remains TBD** — this proposal structures it and
  does not promote any option by composition with the ruled formation model.
- **WRLD-003's wilderness/NPC-settlement structure (PROPOSED)** is respected:
  reclaim flows return land/structures toward the ruled world structure
  (wilderness or NPC-settlement baseline per ECO-001), never into
  system-owned economy participation (ECO-001 boundary held).

## Anchors (controlling references)

| Anchor | Status | What it locks/constrains |
|---|---|---|
| BLD-004 (GDD §19) | TBD | "Structure placement rules, lot systems, density limits, and abandonment handling are unresolved" — the surface structured here |
| HD-BLD-01 (decision record `proposals/TCIndustries_OQ-007_Ruling_Record_HD-BLD-01_2026-09-19.md`) | HUMAN-LOCKED (formation + depth) | Free-claim formation; basic depth; INV-007a approved — the ruled context this proposal composes with |
| INV-007a (Invariant Register v0.1; approved HD-BLD-01) | APPROVED design invariant | Ongoing-participation requirement — abandonment handling is its enforcement mechanism under free-claim |
| PLR-003 (GDD §9) | HUMAN-LOCKED principle-only | Meaningful persistent ownership — title is not confiscated by vacancy without recourse |
| SAFE-001 (GDD §26) | HUMAN-LOCKED principle-only | Safeguards against fraud, abandoned assets, exploitative transfer — abandonment is exactly the abandoned-assets safeguard surface |
| BLD-001 (GDD §19) | LOCKED | Player-owned structures with persistent presence |
| WRLD-001 (GDD §8) | HUMAN-LOCKED principle-only | Persistent shared world — player actions have durable consequences "where practical" (the "where practical" hinge is what abandonment calibration turns) |
| ECO-001 (GDD §15) | LOCKED | Reclaimed assets must not make NPC/system economy dominant (no system-owned production) |
| E-02 phase-7 (Design Evidence Register §2, OQ-007 row) | PROTOTYPE | Civic scaffolding implementable — placement/permission machinery exists as pattern; no abandonment mechanics exist anywhere (Evidence Register §3 gap 7) |

## Principle → Constraint → Invariant → Failure Condition → Test

- **Principle (PROPOSED):** settlements are leases on attention, not deeds
  held against the world: structures and civic roles persist through *use and
  sustenance*, and when sustenance stops, function decays through graduated,
  recourse-before-confiscation stages — keeping the map alive (INV-007a) while
  every ownership stake stays meaningful (PLR-003) and recoverable (SAFE-001).
- **Constraint (derives from LOCKED/ruled layer):** any abandonment mechanism
  must (i) never degrade title without a player-meaningful recourse path
  (PLR-003; SAFE-001), (ii) degrade function before it ever touches title
  (graduation by construction), (iii) return reclaimed capacity to the world
  structure — wilderness or NPC-settlement baseline — never into
  system-owned economic activity (ECO-001), (iv) treat identical vacancy
  states identically regardless of owner identity or account configuration
  (HD-EIC-03/09 pattern-not-structure discipline), (v) carry no numeric
  parameters (BAL-001).
- **Invariant (candidate INV-BLD-004a — function-before-title graduation):**
  *For any vacancy state, functional degradation precedes any title
  transition, and every title transition requires a recourse window in which
  the owner (or their explicitly-designated successor) can restore the asset
  to full function by performing the ruled sustenance obligations.*
- **Invariant (candidate INV-BLD-004b — no ghost equilibrium):** *Under any
  parameterisation, a settlement left with zero sustenance converges to
  non-functional state (INV-007a satisfied at the settlement scale); no
  configuration of placed structures persists indefinitely at full function
  with zero ongoing player involvement.*
- **Failure condition:** ghost-town equilibrium (vacant, fully-functional
  settlements persisting — INV-007a breach at scale); or confiscation drift
  (mechanisms that make ownership meaningfully temporary — PLR-003 breach).
- **Test (shape-level):** lifecycle walk-through per option: from active
  settlement through each vacancy stage, verify (a) function degrades before
  title, (b) a recourse path exists at every title-touching stage, (c) the
  terminal state is world-structure-consistent (wilderness/NPC baseline, not
  system economy). Checklist-class, runnable at design time; behavioural
  testing requires an implementation or simulation host (no abandonment
  mechanics exist in any build — Evidence Register §3 gap 7).

## Options (structural only; tradeoffs and failure modes stated)

### Option A — Pure functional decay (title never moves)
Vacancy degrades *function only*: unmaintained structures lose services,
zoning effect, and civic contribution (a city building's stage value
decays; a vendor's stall stops stocking; a workshop stops producing) but
title remains with the owner indefinitely; physical dereliction is cosmetic
plus functional, never a transfer.
- *Serves:* maximum PLR-003 ownership meaning (title is permanent);
  simplest recourse design (restoration = resume maintenance); zero
  confiscation-fraud surface; fully composable with any future placement/
  density ruling.
- *Tradeoffs:* land/lot scarcity is never recycled — under free-claim
  formation (HD-BLD-01), prime locations can be permanently withheld by
  absent owners (a *speculative lock* failure shape); INV-007b (unapproved,
  candidate) would be hardest to satisfy under this option since locational
  advantage can be frozen.
- *Failure modes:* speculative lock (the map fills with derelict-but-owned
  prime plots — economic dead weight under the ruled dense-polity world);
  cosmetic dereliction without functional meaning decays into noise.
- *Non-canonical-list check:* clean.

### Option B — Graduated decay with eventual reclamation (recourse-before-confiscation)
Functional decay as in A; below a structural threshold (instrument class,
not a number), the asset enters a *recourse window*: the owner and any
explicitly-designated successor are notified through in-world channels; if
the window passes with no restoration, title transitions to the ruled world
structure (wilderness state; or NPC-settlement custodianship for structures
inside NPC-adjacent zones per ECO-001) — never to another player
directly, never into system economy.
- *Serves:* recycles location capacity (kills the speculative-lock failure);
  INV-BLD-004a satisfied by construction (function → window → title);
  reclamation *to the world*, not to claimants — no land-grab economy, no
  griefing-by-neglect vector (a rival cannot weaponise the mechanism).
  Custodianship branch keeps NPC settlements as pure baseline (ECO-001).
- *Tradeoffs:* the recourse-window machinery is the design surface (who is
  notified, through what channels, and how successors are designated —
  identity-neutral state, but real design); title transition is a genuine
  PLR-003 pressure point requiring the SAFE-001 safeguard obligation to be
  taken seriously; "custodianship" needs a definition that keeps NPC
  custodians non-economic (storage-and-state only, no production).
- *Failure modes:* notification failure converting due process into
  confiscation (PLR-003/SAFE-001 breach — the window must be real, which is
  why channel design is named as design surface, not left implicit);
  custodianship drift if NPC custodians slowly gain economic behavior.
- *Non-canonical-list check:* clean — the non-canonical list excludes
  specific SWG housing-auto-decay *values*, not the concept; no
  tradeability/soul-binding rule is assumed (title moves only to the world
  structure, not player-to-player).

### Option C — Social-reclamation layer (successor continuity without state transfer)
Functional decay as in A; long-vacant assets become *transferable by
social mechanism*: the owner may pre-register a successor (craft-guild,
organisation, or named player) whose claim activates through a
player-performed ceremony (the ruled basic-depth governance surface — a
civic act); absent any pre-registration, the asset derelicts permanently
(Option A behaviour) but never changes hands automatically.
- *Serves:* keeps every title transition player-authored (no automatic
  confiscation at all — the strongest PLR-003 reading); builds succession
  into the player-created society pillar (PIL-004: organisations gain a
  real civic function); identity-neutral (anyone may be designated).
- *Tradeoffs:* cannot recycle unclaimed locations (speculative lock
  persists for owners who never designate — same residual failure as A);
  pre-registration UX and the ceremony's legal shape (what a successor
  may/may not do before activation) are new design surfaces; organisations
  as successors binds to OQ-013 (open — flagged, not resolved).
- *Failure modes:* succession-cartel behavior (organisations pressuring
  designations — a SAFE-001 "exploitative transfer" surface); dead
  pre-registrations (successor entities that themselves die — requiring a
  reclamation layer anyway, at which point B's machinery returns).
- *Non-canonical-list check:* clean.

### Cross-option notes (recorded, not decided)
- **Hybrids are lawful:** B + C is the maximal-recourse composition
  (recourse window AND succession designation; the window only triggers
  when no succession activates).
- **Placement/lot/density:** deliberately OUT of this proposal's ruling even
  though BLD-004's text names them — under free-claim formation, placement
  rules interact with OQ-006 transport (zero evidence, backlog) and OQ-015
  scope; bundling them here would front-load unruled dependencies. The
  proposal requests a ruling on **abandonment handling only**; placement/
  lots/density stay TBD (recorded explicitly so "ruling BLD-004" is never
  read as ruling all of it).
- **INV-007b interface:** Option A makes the (unapproved) INV-007b hardest;
  B recycles location capacity and thus *enables* locational competition —
  recorded per option because the owner explicitly declined INV-007b approval
  in HD-BLD-01 and may weigh this in the choice.

## What this proposal does NOT resolve

- Placement rules, lot systems, density limits (BLD-004 remainder —
  deliberately excluded; see cross-option note; OQ-006/OQ-015 interactions).
- Any numeric parameter (vacancy thresholds, window lengths — BAL-001).
- Organisation succession mechanics (OQ-013, open — Option C flags it).
- SAFE-001 transfer-mechanism design beyond the abandonment surface
  (player-to-player transfer rules stay open; soul-binding default excluded
  by the non-canonical list).
- What NPC custodianship may lawfully include (if B is chosen, a one-line
  boundary definition is needed at filing time — flagged in Option B).
- Any implementation, simulation, or testbed work; E-02 phase-7 evidences
  scaffolding only (PROT-001).

## Ruling requested (HD-<Area>-<NN> slot reserved in the Human Rulings Register)

**Question for the owner — rule yes / no / amend (when the item is formally
queued):**

1. **Primary ruling:** Select the abandonment-handling model (BLD-004
   abandonment surface): **A** (pure functional decay; title never moves),
   **B** (graduated decay with recourse window and reclamation to the world
   structure), **C** (social-reclamation succession layer), a lawful hybrid
   (e.g. B+C), or **defer**.
2. **Invariant ruling (independent):** Do you approve candidate invariants
   INV-BLD-004a (function-before-title graduation) and INV-BLD-004b (no
   ghost equilibrium) as design invariants (TEST-001 class)?
3. **Scope confirmation:** Confirm that (a) this ruling covers abandonment
   handling only — placement/lot/density remain TBD (BLD-004 remainder),
   (b) all parameters defer to BAL-001-authorised passes, and (c) title
   transitions, if any, run only to the world structure or through
   player-authored succession — never automatic player-to-player.

*Every claim above cites its file and anchor. On any ruling, the tracker
updates and the queue maintains 3–5 live items per the loop rules.*

---
*End of proposal. PROPOSED — awaiting HD ruling (upon formal queueing). No status changed anywhere by this document.*
