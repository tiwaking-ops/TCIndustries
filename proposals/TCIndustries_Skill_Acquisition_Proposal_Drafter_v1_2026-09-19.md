# TCIndustries — Skill Acquisition Model — Structural Proposal

**Filename:** `TCIndustries_Skill_Acquisition_Proposal_Drafter_v1_2026-09-19.md`
**Status:** PROPOSED. No authority. Ruling-ready candidate only. Nothing here is
LOCKED, and no element of this proposal changes any GDD status.
**Author:** Buffy — Codebuff agent (Freebuff), autonomous-loop drafter, Cycle 4.
**Assessor:** None — single-session authoring; no independent review pass (per GR-004, recorded as such rather than guessed).
**Date:** 2026-09-19
**Queue position:** one of 5 live items — tracker row 2 (inherited the vacated OQ-016 slot; see `proposals/TCIndustries_Continuous_Queue_Tracker.md`; drafted from backlog rank 6 after OQ-016 was ruled HD-RET-01)
**Pre-flight:** the skill-acquisition gap is cited as the OQ-011 *acquisition half* / PROG-003 in the Gap Analysis backlog, the Invariant Register (OQ-011 entry), and the Evidence Register (OQ-011 row) — no prior proposal or investigation structures it as a decision. The respecialisation *cost* half is a separate live queue item (`TCIndustries_OQ-011_Respecialisation_Cost_Rules_Proposal_Drafter_v1_2026-09-19.md`); the two are deliberately split for minimal rulings and this proposal does not pre-empt it. No duplicate pass.

---

## Status wall

- Every option below is **PROPOSED** and requires an explicit HD ruling to advance.
- **Zero numeric values.** No XP rates, training times, costs, or curve shapes
  appear anywhere (BAL-001). Magnitudes require BAL-001 authorisation later.
- **Use-based skill-point acquisition is on the non-canonical list** (Open
  Questions Register §6; GDD §36) and may not be adopted. Options below are
  worded so that none assumes it; where a HISTORICAL pattern resembles it, the
  distinction is argued explicitly (Option B note).
- **Skill Discipline / Skill Box architecture is non-canonical** as a chosen
  model — no option adopts it; HISTORICAL evidence is cited as implementability
  only (SWG-002; E-17).
- **The specialisation budget / cap model (OQ-001, PROF-005) is NOT selected
  here** — it is EIC-coupled (Invariant Register [EIC-COUPLED] flag) and gated
  behind HD-EIC-07 boundaries. Every option is budget-agnostic: each must
  remain compatible with any future budget ruling (PROG-003's own text
  requires this).
- **HD-RET-01 interaction (ruled 2026-09-19):** under the ruled player-only
  vendor model, NPC commerce is limited to baseline services. Any option that
  uses NPC trainers must treat training as a *baseline service* under that
  ruling — and player-provided training becomes a service-economy opportunity
  (SERV-001/002) rather than a vendor question. Flagged per option.
- **Respecialisation cost instrument is not pre-selected** (live queue item):
  acquisition and respecialisation interlock, but each option below states its
  interface to the respecialisation ruling without assuming an instrument.

## Anchors (controlling references)

| Anchor | Status | What it locks/constrains |
|---|---|---|
| OQ-011 acquisition half (GDD §31; §36 item 8 adjacent) | TBD | "Exact skill acquisition (points, experience, training, discovery, etc.) remains unresolved" — the open question structured here |
| PROG-003 (GDD §21) | TBD | Acquisition model must support specialisation budgets and respecialisation principles |
| PROF-001 (GDD §10) | LOCKED | Flexible skill-based profession architecture — professions are emergent identities, not rigid classes |
| PROF-004 (GDD §10) | LOCKED | Respecialisation must be possible; no permanent traps |
| PROF-005 / OQ-001 (GDD §10/§31) | PROPOSED/TBD | Specialisation budget — open, EIC-coupled; NOT selected here, must be compatible |
| LOOP-003 / PROG-001 (GDD §7/§21) | LOCKED | Progression expands choices, capability, specialisation, reputation, access — not one linear power ladder |
| INV-011c (Invariant Register v0.1, OQ-011 entry) | Candidate | Acquisition-route equivalence (quoted under Invariants) |
| HD-RET-01 (2026-09-19, decision record `proposals/TCIndustries_OQ-016_Ruling_Record_HD-RET-01_2026-09-19.md`) | HUMAN-LOCKED (structural model only) | Player-only vendors; NPC baseline services only — bounds where NPC trainers may appear |
| SERV-001 / SERV-002 (GDD §17) | LOCKED / PROPOSED | Player-provided services viable careers; services reinforce interdependence — training can be a service |
| E-01/E-02 phase-2 (Design Evidence Register §2, OQ-011 row) | PROTOTYPE | Trainer/XP acquisition working in fork — implementability evidence only |
| E-17 (Design Evidence Register §1.3) | HISTORICAL | Skill calculator with prereq chains and surrender guards — pattern evidence only (SWG-002) |

## Principle → Constraint → Invariant → Failure Condition → Test

- **Principle (PROPOSED):** skills are acquired through means that generate
  meaningful choices and, where possible, player interaction — acquisition
  should expand what a character can *become* (LOOP-003/PROG-001) rather than
  measure time-served, and should leave room for teaching, study, discovery,
  and practice as distinct, composable means (PROF-001's emergent-profession
  architecture).
- **Constraint (derives from LOCKED layer):** any acquisition model must
  (i) leave every legitimate career reachable from any starting state
  (PROF-004; INV-011c), (ii) support a future specialisation budget without
  redesign (PROG-003; OQ-001 open), (iii) not adopt any non-canonical
  mechanism (use-based SP; skill-box architecture as chosen model), (iv)
  respect HD-RET-01's player-only-vendor structure for any commerce-adjacent
  acquisition step, (v) carry no numeric parameters.
- **Invariant (candidate INV-011c, quoted verbatim from the register):** "*Whatever
  the acquisition model, alternative specialisations remain reachable from any
  starting state — no acquisition path locks a character out of any legitimate
  career permanently.*"
- **Failure condition:** an unreachable-specialisation dead end; an acquisition
  model that silently hard-codes a budget (forcing OQ-001's hand); or one that
  converts training into an NPC-run commodity that displaces a player service
  career (ECO-001/SERV-001 pressure under HD-RET-01's structure).
- **Test (shape-level):** state-machine walk-through of the ruled model against
  INV-011c from at least three distinct starting states (checklist-class,
  runnable at design time); compatibility check against a placeholder budget
  interface (a slot where the future OQ-001 ruling plugs in — no budget model
  implied); behavioural testing requires an implementation or simulation host.

## Options (structural only; tradeoffs and failure modes stated)

### Option A — Trainer-gated study acquisition
Skills are acquired by study: a character seeks a trainer (player-provided
where the economy supports it; NPC baseline trainers where it does not, under
HD-RET-01's baseline-service boundary), pays whatever the respecialisation/
economy rulings later define, and trains the skill subject to prerequisite
chains (E-17 pattern: prereq structure without SWG values).
- *Serves:* strongest SERV-001/002 alignment — expert players can *teach*,
  making knowledge itself a player-economy good (information asymmetry with
  economic value, RES-007's spirit); cleanest prerequisite structure for
  professions-as-emergent-identity (PROF-001); implementability evidenced
  (E-01/E-02 phase-2, PROTOTYPE).
- *Tradeoffs:* trainer availability shapes progression pace — NPC baseline
  trainers are the safe floor, but if they are too complete, player teaching
  never becomes economically real (ECO-001 pressure inside the service band);
  if they are too thin, progression bottlenecks on player availability.
  Where that floor/ceiling line sits is a later structural question flagged,
  not decided, by this option.
- *Failure modes:* NPC trainers crowding out player teaching (the exact
  HD-RET-01 baseline-service boundary pressure); prerequisite chains
  degenerating into a fixed ladder (LOOP-003 breach) if prereqs are designed
  as a single spine rather than a graph.
- *Non-canonical-list check:* clean — no use-based SP; prereq chains are not
  skill boxes.

### Option B — Activity-earned pools + gate acquisition
Activity generates experiential pools (by domain, shape TBD); acquisition
still requires passing a gate (trainer, certification, or demonstrable
milestone) to convert pools into skills — activity earns the *capacity* to
learn, the gate confers the *skill*.
- *Serves:* the HISTORICAL typed-XP-pool pattern (Gap Analysis §3.2 reuse map)
  with the gate preserved; activity-feel progression without forcing all
  progression through trainers; pools-by-domain composes naturally with any
  future budget model (domains are the natural budget axis — flagged, not
  assumed).
- *Tradeoffs:* the most pattern-similar option to the non-canonical
  "use-based skill point acquisition" — the load-bearing distinction is
  recorded explicitly: use-based SP grants the *skill itself* by use, whereas
  this option grants only a convertible capacity and always requires a
  separate conferral step. If the owner judges the distinction too thin, this
  option should be declined on that ground alone (recorded here so the
  rejection is easy and explicit).
- *Failure modes:* without a meaningful gate, B collapses into the
  non-canonical model by behaviour (grinding = skills) — the gate is the
  option's integrity condition; pools-per-domain may silently pre-structure
  the OQ-001 budget (mitigated by the placeholder-interface check in Test).
- *Non-canonical-list check:* argued explicitly above; adoption requires the
  owner to accept the stated distinction.

### Option C — Mentorship / discovery acquisition
Skills transfer between players through teaching relationships and are
discovered through world engagement (experimentation, rare knowledge objects,
observation); no NPC acquisition path for advanced tiers at all.
- *Serves:* maximum interdependence and information-asymmetry value (RES-007);
  knowledge becomes the scarcest good — strongest PIL-003 alignment of the
  three; makes expertise genuinely social (VIS-001/002).
- *Tradeoffs:* heaviest design surface (teaching mechanics, knowledge
  representation); zero direct implementability evidence in the corpus
  (closest HISTORICAL analogies only — UNVERIFIED as to whether mentor-based
  acquisition existed in any referenced predecessor; stated honestly rather
  than claimed); population-dependent: thin servers starve progression.
- *Failure modes:* gatekeeping by incumbent experts (access cartelisation —
  must be checked against INV-011c's reachability requirement at walk-through);
  new-player experience collapses without an NPC baseline for fundamentals
  (HD-RET-01 permits baseline services — but "advanced tiers NPC-free" is a
  deliberate design stance this option takes, flagged for the owner's
  explicit acceptance).
- *Non-canonical-list check:* clean.

### Cross-option notes (recorded, not decided)
- **Hybrids are lawful ruling outcomes** — e.g. A-for-fundamentals + C-for-
  advanced, or B-pools + A-gates. The ruling may pick one, combine, or amend.
- **Respecialisation interface:** any option must state what respecialisation
  "consumes" without pre-selecting the instrument (live queue item OQ-011
  cost ruling): A consumes nothing at acquisition time; B consumes accrued
  pools; C consumes the teaching relationship's availability. Recorded per
  option; the interlock is resolved when both rulings exist.

## What this proposal does NOT resolve

- The specialisation budget / cap model (OQ-001/PROF-005 — EIC-coupled,
  HD-EIC-07-gated; placeholder interface only, no model implied).
- The respecialisation cost instrument (live queue item; not pre-selected).
- Any numeric value (XP rates, training times, costs, curves — BAL-001).
- Where the NPC-baseline vs player-taught service line sits (HD-RET-01 bounds
  it as "baseline services only"; the exact line is a later structural
  question, flagged in Option A).
- Prerequisite-chain content (which skills prereq which — content design,
  after a model is ruled).
- Any implementation, simulation, or testbed work; E-01/E-02/E-17 are
  implementability/pattern evidence only (PROT-001, SWG-002).

## Ruling requested (HD-<Area>-<NN> slot reserved in the Human Rulings Register)

**Question for the owner — rule yes / no / amend:**

1. **Primary ruling:** Select the skill acquisition model (OQ-011 acquisition
   half / PROG-003): **A** (trainer-gated study), **B** (activity-earned
   pools + conferral gate, with the non-canonical distinction as stated),
   **C** (mentorship/discovery; advanced tiers NPC-free by design), any
   lawful hybrid, or **defer**.
2. **Invariant ruling (independent):** Do you approve candidate invariant
   INV-011c (acquisition-route equivalence) as a design invariant
   (TEST-001 class) binding whichever model is chosen?
3. **Scope confirmation:** Confirm that (a) no budget model (OQ-001) is
   selected or implied, (b) all magnitudes defer to BAL-001, and (c) the
   respecialisation cost instrument remains with the live OQ-011 queue item.

*Every claim above cites its file and anchor. On any ruling, the tracker
updates and the queue maintains 3–5 live items per the loop rules.*

---
*End of proposal. PROPOSED — awaiting HD ruling. No status changed anywhere by this document.*
