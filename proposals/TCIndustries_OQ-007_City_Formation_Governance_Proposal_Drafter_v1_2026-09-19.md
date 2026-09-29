# TCIndustries — OQ-007 City Formation & Governance — Structural Proposal

**Filename:** `TCIndustries_OQ-007_City_Formation_Governance_Proposal_Drafter_v1_2026-09-19.md`
**Status:** PROPOSED. No authority. Ruling-ready candidate only. Nothing here is
LOCKED, and no element of this proposal changes any GDD status.
**Author:** Buffy — Codebuff agent (Freebuff), autonomous-loop drafter, Cycle 1.
**Assessor:** None — single-session authoring; no independent review pass (per GR-004, recorded as such rather than guessed).
**Date:** 2026-09-19
**Queue position:** 4 of 5 (see `proposals/TCIndustries_Continuous_Queue_Tracker.md`)
**Pre-flight:** no prior proposal or investigation addresses OQ-007 city design; the gap ID appears only in invariant and evidence rows (cited below). No duplicate pass.

---

## Status wall

- Every option below is **PROPOSED** and requires an explicit HD ruling to advance.
- **Zero numeric values.** No tax rates, maintenance fees, population thresholds,
  or decay timers appear anywhere (BAL-001). Taxation/maintenance *numbers* are
  explicitly excluded — structure only.
- **Taxation appears here only as an open structural question** (GDD §30 lists
  "city formation, governance, and maintenance" TBD; BLD-003 names "taxation or
  maintenance models" as the open surface). No instrument, rate, or sink design
  is proposed — that layer is BAL-001-gated and also interacts with OQ-009
  (currency, open).
- **Exact city-formation rules are TBD** (WRLD-003, GDD §8). SOC-003/EMRG-001
  (LOCKED layer) forbid scripting every institution — all options below are
  *condition-creation* shapes, not scripted-institution shapes.
- **No automated Architecture D / residual-channel content.** Cities touch
  interdependence obliquely (commerce placement, INV-007b) but nothing here
  designs an interdependence mechanism; HD-EIC-07 untouched, HD-EIC-05
  (Architecture C retired) respected.
- Evidence status: civic *scaffolding* (membership, permissions, mail) is
  PROTOTYPE-evidenced (E-02 phase-7, Evidence Register §2 OQ-007 row);
  governance/taxation *design* is untested anywhere ("INV-007a/b untested
  anywhere"). Nothing below claims testbed support it does not have.

## Anchors (controlling references)

| Anchor | Status | What it locks/constrains |
|---|---|---|
| OQ-007 (GDD §31) | TBD | The open question this proposal structures |
| BLD-003 (GDD §19) | PROPOSED/TBD | Player cities should support governance, zoning, taxation or maintenance models, civic structures, community identity; exact rules TBD |
| WRLD-003 (GDD §8) | PROPOSED | NPC settlements / player settlements / wilderness / economic-corridors world structure; "exact city-formation rules are TBD" |
| WRLD-001 (GDD §8) | HUMAN-LOCKED principle-only (HD-GDD-01) | Persistent shared world — cities are durable player-created consequences |
| PIL-004 (GDD §6) | LOCKED | Players must form organisations, settlements, social hubs with meaningful world presence |
| SOC-003 (GDD §20) | PROPOSED | Support emergent social institutions without scripting every institution |
| VIS-005 / EMRG-001 (GDD §5/§22) | LOCKED | Conditions for emergence, not scripted experience |
| ECO-001 (GDD §15) | LOCKED | NPC settlements serve baseline needs without displacing player economy |
| BLD-004 (GDD §19) | TBD | Placement rules, lot systems, density limits, abandonment — adjacent open surface, not resolved here |
| INV-007a/b (Invariant Register v0.1, OQ-007 entry) | Candidate | Ongoing-participation requirement; differentiated decision-relevant locational advantage, no strictly dominant placement |
| E-02 phase-7 (Design Evidence Register §1.1/§2) | PROTOTYPE | Civic scaffolding (cities/guilds/mail) implementable and suite-testable — implementability only |

## Principle → Constraint → Invariant → Failure Condition → Test

- **Principle (PROPOSED):** cities are player-created institutions that persist
  because players sustain them: formation must be achievable by organised
  players (PIL-004), governance must remain emergent (SOC-003/EMRG-001), and
  civic structure must make location matter economically (BLD-003, ECO-002)
  without any single placement being strictly dominant.
- **Constraint (derives from LOCKED layer):** any formation/governance model
  must (i) never make cities set-and-forget automation (INV-007a; anti-Factorio
  spirit, MFG-004 shape), (ii) keep NPC settlements from displacing the player
  economy (ECO-001), (iii) preserve emergence — governance mechanisms create
  conditions, not scripted outcomes (VIS-005), (iv) respect BLD-004 as an open
  adjacent surface (no placement/density rules invented here), (v) carry no
  numeric parameters.
- **Invariant (candidate INV-007a, quoted):** "*City governance mechanisms
  require ongoing player participation to persist — no city remains fully
  functional indefinitely with zero player involvement*" (guards "living
  society" vs. set-and-forget automation — guard note from the register).
- **Invariant (candidate INV-007b, quoted):** "*Cities create differentiated
  economic effects: at least one locational advantage (zoning, tax,
  maintenance, civic structure access) is decision-relevant for commerce
  placement, and no placement is strictly dominant for all purposes."*
- **Failure condition:** INV-007a fails under an abandoned-city persistence
  strategy; INV-007b fails when all locations are equivalent (no civic
  meaning) or one placement dominates everything (degenerate monoculture).
- **Test (shape-level):** policy-shape walk-through of the ruled model against
  INV-007a/b (checklist-class, runnable at design time); full testing requires
  city simulation — no governance mechanics exist in any build today (Evidence
  Register §3 item 7).

## Options (structural only; tradeoffs and failure modes stated)

### Option A — Free-claim settlement formation
Any player/organisation meeting *structural* preconditions (declared site,
declared governance instrument, minimum participating-actor structure — shape
defined at ruling time, no numeric threshold proposed) may found a settlement;
the world is dense with small polities from early on.
- *Serves:* PIL-004 at full strength; fastest path to emergent city politics;
  strongest INV-007b differentiation pressure (many polities → competition
  makes locational choice meaningful).
- *Tradeoffs:* map sprawl and ghost towns are likely without abandonment
  handling (BLD-004 surface, open); civic meaning may dilute if settlements
  are trivially many; NPC-settlement role must be carefully bounded (ECO-001).
- *Failure modes:* INV-007a pressure — free-claimed, cheaply-held towns that
  nobody tends must decay structurally or the world fills with autonomous
  ghost infrastructure; over-fragmentation making governance itself
  meaningless (all polities too small to matter).

### Option B — Charter / petition formation
Settlement formation requires an explicit charter act — a founding instrument
assembling land claims, participants, and a declared governance model —
approved through a (to-be-defined) institutional channel; fewer, more
committed polities.
- *Serves:* concentration of civic investment; each city is an event (PIL-001
  event-shape analogy: formation itself becomes a player-driven occurrence);
  cleaner abandonment story (charter maintenance lapses → settlement status
  changes, honoring INV-007a by construction).
- *Tradeoffs:* the approving institution becomes a design object with its own
  politics — who charters? (player institution? system adjudication? hybrid?)
  — a genuinely open sub-question this option intentionally leaves open;
  slower to reach the emergent-city state; risk of gatekeeping by incumbent
  polities.
- *Failure modes:* if the charter channel is system-adjudicated with fixed
  criteria, formation drifts toward scripted quest-like behaviour (VIS-006
  theme-park pressure); if player-adjudicated, incumbent capture risk — both
  need the walk-through, neither is resolved here.

### Option C — Settlement-climber (stages of civic maturity)
Settlements begin as minimal outposts (a claimed location with presence) and
*may* climb through structural stages (outpost → village → town → city) by
sustained player activity and civic construction; governance *surface* widens
with maturity (higher stages support more governance instruments).
- *Serves:* smooth onboarding into civic life; each stage transition is an
  emergent event; naturally encodes INV-007a (stages are maintained by
  continued participation, not one-time achievement); gives RET-003 commercial
  districts a natural home (mature cities).
- *Tradeoffs:* the stage ladder is the most system-heavy option — a ladder
  designable in PLAY-FEEL terms but with obvious numeric-tuning gravity
  (thresholds), which must stay deferred (BAL-001); risk of checklist-feel
  ("city XP") if stage criteria are designed as grind metrics rather than
  structural capacities — recorded as the option's principal drift danger.
- *Failure modes:* if stage criteria are purely metric-thresholds, drift
  toward theme-park progression (anti-goals); if purely organic with no
  criteria, the ladder collapses back into Option A; the boundary is the
  design judgment this ruling does NOT need to make (structure now, criteria
  later under BAL-001 and a follow-up ruling).

### Cross-option note (recorded, not decided)
- All three options leave **governance instrument depth** as a second-axis
  ruling: (i) *basic* — ownership, zoning presence, shared maintenance
  responsibility (minimum for INV-007a/b to be testable) versus (ii) *civic* —
  instruments for policy, treasury, and civic offices (deeper; interacts with
  OQ-013 organisations and OQ-009 currency). The owner may rule formation
  model and depth independently (e.g. "Option B formation, basic depth now,
  civic depth later").
- Taxation/maintenance appears only as named-open structural surface; its
  instrument choice (and every number) is BAL-001-gated and OQ-009-dependent.

## What this proposal does NOT resolve

- The formation model itself (that is the requested ruling) and any numeric
  threshold, rate, or stage criterion (BAL-001-gated, separately authorised).
- Governance instrument details (policy/treasury/office mechanics — second-axis
  ruling above; OQ-013 org rights, queued in backlog, feeds it).
- Placement/lot/density/abandonment rules (BLD-004, open — flagged as the
  immediately adjacent next gap; abandonment handling in particular bounds
  Option A).
- Transportation/corridor design (WRLD-005/OQ-006 — Evidence Register gap 1:
  zero evidence anywhere; backlog item 11).
- Any implementation, simulation, or testbed work; E-02 phase-7 evidences
  scaffolding implementability only (PROT-001).

## Ruling requested (HD-<Area>-<NN> slot reserved in the Human Rulings Register)

**Question for the owner — rule yes / no / amend:**

1. **Primary ruling:** Select the structural formation model for OQ-007:
   **A** (free-claim), **B** (charter/petition), **C** (settlement-climber),
   or **defer** (amendments welcome).
2. **Depth ruling (independent second axis):** Rule governance-instrument depth
   for now: **basic** (ownership + zoning presence + maintenance
   responsibility — minimum viable civic surface) or **civic** (full
   instruments; will bind to OQ-013/OQ-009 follow-ups), or **defer depth**.
3. **Invariant ruling (independent):** Do you approve candidate invariants
   INV-007a and INV-007b as design invariants (TEST-001 class) for whichever
   model is chosen?
4. **Scope confirmation:** Confirm that taxation/maintenance instruments and
   all numeric thresholds are deferred to future BAL-001-authorised passes,
   keeping this ruling structure-only.

*Every claim above cites its file and anchor. On any ruling, the tracker updates
and the queue maintains 3–5 live items per the loop rules.*

---
*End of proposal. PROPOSED — awaiting HD ruling. No status changed anywhere by this document.*
