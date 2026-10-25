# TCIndustries — OQ-001 Specialisation Budget / Skill-Cap Model: Proposed Ruling Pass

- **Author:** Claude (Anthropic)
- **Version:** 0.1
- **Date:** 2026-10-25
- **Status:** **PROPOSED — NOT RULED.** This document resolves the *model shape* of OQ-001 and is ready for the owner's authority act (which would assign the HD-PROF-0N id and flip it to HUMAN-LOCKED). It changes no GDD status and authorises **no** numeric value. All magnitudes are explicitly deferred to a BAL-001 pass.
- **Resolves:** GDD open question **OQ-001** (exact specialisation budget / skill-cap model), currently **TBD**.
- **Grounded in:** `investigations/TCIndustries_Specialisation_Shape_Study_v0.1.md` (PROPOSED analysis, 2026-10-23); HD-EIC-10 structural boundaries (2026-10-23, gate closed); addendum Q3/Q4 to `SWG_Phase9_Elite_Missions_Proposal_v0.1.md`; the live fork `testbed/swg-phase3-combat-fork/tc-server` (`internal/skills/professions.go`).
- **Companion:** implementation-gated by `proposals/TCIndustries_Phase9_Elite_and_TaskLayer_Proposal_v0.1.md` §4 (gating decision #1).

---

## 0. Authority posture (read first)

Per `AGENTS.md`, I (the agent) do **not** make the game-design decision myself and do **not** silently promote OQ-001 from **TBD** to **LOCKED/HUMAN-LOCKED**. This document is the *ruling pass the owner requested*: a fully-formed recommendation — model shape selected, rejected alternatives named, magnitudes scoped to a deferred pass — presented for the owner's explicit conferral of authority. Until then, OQ-001 stays **TBD** in the GDD.

The HD-EIC-07 gate that previously blocked OQ-001 search is **closed** (HD-EIC-10, 2026-10-23); the OQ-011 register row's *"blocked behind the open HD-EIC-07 gate"* is **stale** (addendum Q3). OQ-001 is now *workable* but *unresolved* — this pass resolves it structurally, subject to owner authorisation.

---

## 1. The ruling (recommended — for owner approval)

| # | Question | Recommended decision |
|---|---|---|
| 1 | Exact specialisation budget / skill-cap model (OQ-001) | **Two-layer budget: (i) a retained character-layer skill-point hygiene cap as a *subordinate* floor, plus (ii) a participant-layer interdependence budget anchored on HD-EIC-10 commitments C2 (persistent information/discovery asymmetry) and C4 (non-economic institutional/social constraint), delivered as soft comparative advantage (addendum Q4).** |
| 2 | Is the pure character-sheet point cap (SWG "Skill Discipline / Skill Box" model) the answer? | **No — rejected as the sole or primary instrument.** It may persist only as the subordinate character-layer hygiene floor in layer (i); it is explicitly **NOT CANONICAL** as a complete model (GDD §36). |
| 3 | Numeric values (cap size, decay rates, complementarity binding strength)? | **Deferred to a BAL-001-authorised numeric pass.** None set here. The live fork's `MaxSkillPoints = 250` is **PROTOTYPE/EVIDENCE** (SWG historical baseline), not an authorised TCIndustries value. |

---

## 2. Model shape — two layers

### Layer (i) — Character-layer skill-point hygiene cap (subordinate, retained)
A per-character skill-point ceiling remains as a *hygiene* rule: a character cannot hold an unbounded number of skill boxes, and exceeding the cap requires dropping an owned box (free-drop per HD-PROF-02, preserving PROF-004). This layer does **not** by itself prevent single-character mastery — the shape study (§3, Shape 1; §6 conclusion) shows a pure character-sheet cap is *internally bypassable* by multi-account pooling (the Package β failure mode, HD-EIC-05). It is retained because it is cheap, keeps the character sheet legible, and satisfies PROF-005's "preserve experimentation, casual participation, hybrid roles, respecialisation" hygiene intent — **but it is explicitly subordinate** to Layer (ii).

### Layer (ii) — Participant-layer interdependence budget (primary, HD-EIC-10 anchored)
The binding constraint on single-actor mastery is a **participant-layer** budget (the validation unit is the participant, per HD-EIC-01, not merely the character). It is expressed as **soft comparative advantage** (addendum Q4) and anchored on two HD-EIC-10 durable commitments:

- **C2 — persistent information/discovery asymmetry:** high-tier effectiveness depends on knowledge, discovery, or provenance that is *not* trivially internalised across an alt army. A single participant holding every role requires sustaining persistent asymmetries that cross the G2 "literal effortlessness" threshold (HD-EIC-09 gloss adopted by HD-EIC-10), so broad coverage imposes sustained organisational effort, not mere capital.
- **C4 — non-economic institutional/social constraint:** peak effectiveness at a high-tier role requires non-economic associational bonds (guild/association membership, reputation-provenance, civic standing) that cannot be bought or alt-internalised. This is the historical "top-tier doctor buff" pattern (F-7), instantiated as structural interdependence rather than a numeric ceiling.

This is the shape study's **Shape 4 (Organisational Complementarity) + Shape 3 (Maintenance/Decay)**, combined and anchored on C2/C4 — the study's stated viable combination (§6). It delivers PIL-003's LOCKED high-tier non-self-sufficiency through *advantage without autarky*, not through a hard numeric budget and not through coercive bifurcation (Shape 2, which the study rules as violating PIL-003).

---

## 3. Rejected alternatives (and why)

| Alternative | Verdict | Reason |
|---|---|---|
| **Shape 1 — pure global skill-point pool** ("Skill Discipline / Skill Box" cap) | Rejected as sole/primary | Bypassed by multi-account alt-pooling (Package β); §36 marks "Skill Discipline / Skill Box architecture" **NOT CANONICAL**. Retained only as Layer (i) subordinate hygiene. |
| **Shape 2 — per-domain mutually-exclusive bifurcation** | Rejected | Creates "coercive inconvenience" PIL-003 forbids; forces multi-boxing; shape study §3 verdict. |
| **Hard factory-quality ceiling / account-level restrictions** | Forbidden | Explicitly excluded by HD-EIC-01 cumulative effect (line 68) and HD-EIC-10 soft-residual exclusions. |
| **Character-sheet cap as the *complete* OQ-001 answer** | Rejected | Cannot solve the participant-layer closure problem (shape study §6). |

---

## 4. Relationship to existing rulings and constraints

- **HD-PROF-02 (free-drop respec) / PROF-004 (LOCKED):** Layer (i) requires free-drop on cap overflow; respec identity survives. Preserved.
- **HD-PROF-03 (trainer-gated acquisition):** acquisition gating is orthogonal to the budget; both layers sit *above* acquisition. Held tension with PROF-001 (addendum Q5) is unaffected — gated acquisition ≠ immutable class.
- **PIL-003 (LOCKED):** satisfied by Layer (ii)'s comparative-advantage-without-autarky, not by a numeric ceiling.
- **PROG-001 (LOCKED option expansion):** the budget expands *choices/specialisation depth* (Layer ii's role access), not a linear power ladder.
- **BAL-001 (LOCKED):** no caps, rates, decay curves, or binding strengths set in this pass. Numeric authorisation is a separate, future act.
- **GDD §36 non-canonical boundary:** "Skill Discipline / Skill Box architecture" and "Two-to-three concurrent profession target" are **excluded** as model sources; "Use-based Skill Point acquisition" is **not** adopted as the rate source (addendum Q9).
- **HD-ECO-01:** Layer (ii) interdependence creates the *demand for others' services* that is the sink-side counterpart to any task-layer reward (addendum Q8).

---

## 5. What remains deferred (BAL-001 pass — not performed here)

| Open magnitude | State |
|---|---|
| Character-layer cap size (currently `250` in fork code) | **EVIDENCE only** — SWG historical baseline; not authorised. BAL-001 pass to set/replace. |
| Layer (ii) decay/maintenance rates (if Shape 3 elements used) | TBD — BAL-001 |
| Complementarity binding strength (C4 associational gating) | TBD — BAL-001 |
| XP / second-tier progression rates feeding the cap | TBD — PROG-003 + BAL-001 (addendum Q9) |

---

## 6. Recommended owner action

1. **Approve or amend** the two-layer model in §1/#1 (the structural shape). On approval, this becomes a HUMAN-LOCKED ruling (HD-PROF-0N, id assigned by filing), OQ-001 moves **TBD → RULED**, and the GDD §31 / register entries are updated.
2. **Authorise a BAL-001 numeric pass** for the magnitudes in §5.
3. Only then does `TCIndustries_Phase9_Elite_and_TaskLayer_Proposal_v0.1.md`'s gating decision #1 clear, unlocking an elite-tier implementation proposal for the live fork.

*End of proposed ruling pass (PROPOSED). Resolves OQ-001's model shape; authorises no numbers; requires a distinct owner authority act to take effect.*
