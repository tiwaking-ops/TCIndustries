# TCIndustries Phase 9 — Elite Tier & Task Layer: Scoped Replacement Proposal

- **Author:** Claude (Anthropic)
- **Version:** 0.1 (PROPOSED — not approved; supersedes the build-authorisation posture of `SWG_Phase9_Elite_Missions_Proposal_v0.1.md`, whose addendum §9 Q10 directs this replacement)
- **Date:** 2026-10-25
- **Status:** **PROPOSED.** This document changes no GDD status, ruling, or authority line. It inherits the ten owner decisions recorded in the addendum to `SWG_Phase9_Elite_Missions_Proposal_v0.1.md` (2026-10-25) and re-scopes the work to what those decisions actually authorise. **No code may be written from this document until the owner separately authorises it and the gating open questions below are resolved.**
- **Baseline:** `TCIndustries_Master_GDD_v1.1.1_Status_Patch.md` (current canonical design reference). Testbed target: `testbed/swg-phase3-combat-fork/tc-server` (live fork, per addendum Q7). The other fork (`swg-phase3-combat/`) is EVIDENCE/PROTOTYPE only per HD-TST-01.

---

## 1. Why this document exists

The predecessor (`SWG_Phase9_Elite_Missions_Proposal_v0.1.md`) was a Muse Spark PROPOSED plan carrying SWG Pre-CU content (7–14 named elites, missions, Ranger, surveying) that its own header disclaimed as *"No authority."* Its 2026-10-25 addendum ran a four-agent authority pass and recorded ten owner decisions (Q1–Q10). **Q10** rules that the proposal must not be enacted by its §10 Item 7 as written (which would green-light that SWG content), and that the path forward is a **replacement proposal** scoped to the decisions actually taken.

This document is that replacement. It:

- (a) adopts the **TCIndustries-native** task layer definition from **Q1** (no SWG names, mechanics, or structures — all SWG content stays HISTORICAL per HD-TST-01);
- (b) carries **only** the elite/specialisation question set that survives the open-question set, bound to its **LOCKED** constraints (PROF-001, PROF-004, PIL-003, PROG-001, BAL-001, and the §36 non-canonical boundary);
- (c) is itself **PROPOSED** and authorises nothing until fresh owner authorisation.

---

## 2. Scope that this proposal resolves (structural only)

### 2.1 Task Layer — TCIndustries-native, procedurally repeatable

**Decision (Q1):** The "mission" content is redefined as a **TCIndustries-native, procedurally-repeatable task layer**: a recurrent, system-generated work offering that players may accept, perform, and complete for a defined outcome. It is **structural only** — this proposal specifies the *shape* of the layer (repeatable task contracts, acceptance, completion, reward hook) and deliberately **not** any SWG-specific mechanic, name, or value. SWG mission behaviour remains HISTORICAL input, never promoted to canon (HD-TST-01).

**Compatibility (Q2):** The task layer is compatible with **VIS-006** and **§24 guard #1** (anti theme-park / quest-driven drift) **provided it never becomes the primary narrative or progression driver.** It is a secondary demand/activity source, consistent with **LOOP-002** parity (a meaningful career must not require combat-as-primary; equally, a task layer must not become the only path to progression).

**Rewards (Q8):** Rewards are **PERMITTED only once a matched sink exists** (HD-ECO-01, faucet-exceeds-sink). The task layer includes a *reward hook* but no reward magnitude — magnitudes are deferred to a **BAL-001** numeric-authorisation pass, which this proposal does not perform.

### 2.2 Elite / Specialisation Tier — gated, not authorised

**Decision (Q3):** Any elite or high-tier profession **requires an explicit OQ-001 ruling before it ships.** OQ-001 (the specialisation budget / skill-cap model) is **TBD** in the Master GDD and remains the live prerequisite. The OQ-011 register row's *"blocked behind the open HD-EIC-07 gate"* is **stale** — HD-EIC-10 closed that gate (2026-10-23) — but closure lifts a *prohibition on starting*, not an *authorisation to skip OQ-001*. PROF-003 and PROF-005 stay **PROPOSED**.

**Interdependence instrument (Q4):** PIL-003's LOCKED high-tier non-self-sufficiency is delivered by **soft comparative advantage** (HD-EIC-02 channel) — high-tier status confers specialist advantage without autarky — **not** a hard numeric budget. This avoids importing the §36-**NOT-CANONICAL** "Skill Discipline / Skill Box" cap model and avoids colliding with BAL-001 at the magnitude layer. Degrees of advantage are deferred to a BAL-001 pass.

**Identity tension (Q5):** HD-PROF-03's trainer-gated **class acquisition** does **not** supersede PROF-001's LOCKED flexible-skill identity. They are a **held tension** on distinct axes — acquisition *access* (trainer-gated) vs identity *fixity* (not immutable) — kept coherent by HD-PROF-02's **free-drop respec**. Any elite design must keep acquisition trainer-gated **and** keep identities freely respec-able.

**Progression rates (Q9):** XP / second-tier progression rates are **deferred to PROG-003 via an explicit BAL-001 pass.** None are invented here. The GDD §36 *"Use-based Skill Point acquisition"* line stays **NOT CANONICAL** and is not adopted as the rate source.

### 2.3 Phase framing and target (Q6, Q7)

"Phase 9" denotes the **fork's next build-sequence phase** on the live testbed, **not** a TCIndustries design unit; no canon is conferred by the label. Build target, when authorised, is the **live fork** (`swg-phase3-combat-fork/tc-server`, v2 42-point ladder). The old fork's 21 elite `ProfessionDef` entries and `Standard80`/`Elite120` cost tables are uncommitted EVIDENCE only.

---

## 3. The LOCKED constraints any built tier/task layer must satisfy

| Constraint | Status | Binding requirement on this work |
|---|---|---|
| PROF-001 flexible skill-based (not rigid classes) | LOCKED | Elite/identity design must stay emergent-from-skills, not immutable-class |
| PROF-004 respecialisation | LOCKED | Free-drop respec must be preserved (HD-PROF-02) — gated ≠ immutable |
| PIL-003 high-tier non-self-sufficiency | LOCKED | Delivered via soft comparative advantage (Q4), not autarky |
| PROG-001 option expansion, not linear power | LOCKED | Tier expands choices/specialisation, not a single power ladder |
| BAL-001 no premature numeric tuning | LOCKED | No caps, rates, reward values, or sink magnitudes set here |
| §24 guard #1 / VIS-006 | LOCKED | Task layer must not become primary progression or theme-park driver |
| GDD §36 non-canonical boundary | NOT CANONICAL list | "Skill Discipline/Skill Box", "Use-based SP acquisition", "2–3 profession target" are excluded as model sources |
| HD-ECO-01 faucet-exceeds-sink | RULED | Task-layer rewards require a matched sink |

PROF-003 (specialisation investment) and PROF-005 (progression caps / specialisation budget) remain **PROPOSED** and are not relied upon as authority.

---

## 4. Open decisions still required before any build

This proposal deliberately **does not** resolve these. They gate implementation:

1. **OQ-001 ruling pass** (cap shape + budget) — per Q3, the hard prerequisite for any elite tier. Recommended as a standalone owner ruling (cap shape + numeric budget), independent of the closed HD-EIC-07 gate.
2. **PROG-003 + BAL-001 numeric pass** (XP / second-tier progression rates and any reward/sink magnitudes) — per Q9 and Q8.
3. **Owner authorisation of this document itself** — it is PROPOSED; writing code from it requires explicit owner approval, recorded with date.

Until all three exist, the implementation status below is deliberately empty.

---

## 5. Implementation status

| Item | Status |
|---|---|
| Task layer definition (structural) | Defined in §2.1 — PROPOSED, structural only, no mechanics/values |
| Task layer rewards | Reward *hook* only; magnitude deferred (Q8, BAL-001) |
| Elite / specialisation tier | **Not authorised** — gated behind OQ-001 ruling (Q3) |
| High-tier non-self-sufficiency instrument | Soft comparative advantage specified (Q4); degrees deferred |
| Progression rates | Deferred to PROG-003 + BAL-001 (Q9) |
| Code written | **None.** No `cmd/phase9test`, no mission/elite/pet handlers exist in the live fork, and none are added under this PROPOSED document |

---

## 6. What this proposal does NOT do

- It does **not** authorise SWG Pre-CU content (named elites, missions, Ranger, surveying). That material is HISTORICAL (HD-TST-01) and out of scope.
- It does **not** change any GDD status, ruling, or authority line (per `AGENTS.md`).
- It does **not** set any numeric value (caps, rates, rewards, sinks) — BAL-001 (LOCKED) forbids it.
- It does **not** resolve OQ-001, PROG-003, or the PROF-003/005 PROPOSED status.
- It does **not** promote the old fork's elite `ProfessionDef` / cost-table code into canon (EVIDENCE only).

---

## 7. Recommended next decision

1. Owner authorises this replacement proposal (or amends it).
2. A dedicated **OQ-001 ruling pass** defines the specialisation budget/cap model.
3. A **BAL-001 numeric pass** supplies PROG-003 progression rates and any task-layer reward/sink magnitudes.
4. Only then is an implementation proposal (with code, handlers, `cmd/phase9test`) warranted — built against the live fork and verified end-to-end per the testbed's "verify, don't claim" discipline.

*End of proposal v0.1 (PROPOSED). Supersedes the build-authorisation posture of `SWG_Phase9_Elite_Missions_Proposal_v0.1.md` per its addendum §9 Q10; inherits all ten owner decisions recorded there on 2026-10-25.*
