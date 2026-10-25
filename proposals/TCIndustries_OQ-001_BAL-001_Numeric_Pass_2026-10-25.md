# TCIndustries — OQ-001 BAL-001 Numeric Authorisation Pass

- **Author:** Claude (Anthropic)
- **Version:** 1.0
- **Date:** 2026-10-25
- **Status:** **BAL-001-AUTHORISED (owner authorisation 2026-10-25).** This pass supplies the numeric magnitudes that `proposals/TCIndustries_OQ-001_Specialisation_Budget_Proposed_Ruling_2026-10-25.md` (enacted as HD-PROF-04) deferred to a BAL-001 pass, and the PROG-003 + BAL-001 magnitudes the Phase 9 replacement proposal (`TCIndustries_Phase9_Elite_and_TaskLayer_Proposal_v0.1.md`, §4 item #2) deferred.
- **Grounded in:** HD-PROF-04 (two-layer model); GDD §6.3.1 (corrected Basic ladder, ~40 SP/profession), §6.3.2 (typed-XP / credit cost formulas), §8.2 (six basic professions); the live fork's `internal/skills/professions.go` cost-derivation (`xpCostForPoints`, `creditCostForPoints`); HD-EIC-10 C2/C4 commitments; HD-ECO-01 (faucet-exceeds-sink).
- **Companion:** implements the deferred magnitudes of HD-PROF-04 §5 and the Phase 9 proposal §4 item #2. Does **not** authorise fork code — see §6.

---

## 0. Authority posture (read first)

Per BAL-001 (LOCKED), numeric values may be set only once the structural design is stable and the values can be tested. OQ-001's *structure* is stable (RULED by HD-PROF-04, two-layer model). This pass therefore **authorises the magnitudes** the structure needs, at the owner's explicit 2026-10-25 authorisation.

BAL-001's own caveat — *"false precision is worse than explicit unresolved questions"* — is honoured two ways:
1. Every value below is explicitly **revisable** after testbed observation; this pass is a *starting authorised set*, not a final balance.
2. Magnitudes whose structural partners do not yet exist are **deferred, not invented** (see §5) — specifically task-layer reward/sink values, which require a matched sink under HD-ECO-01 and a built task layer.

**Out of scope for this pass:** any value not enumerated in §1–§4. The fork's `MaxSkillPoints = 250` was PROTOTYPE/EVIDENCE; it is here *authorised as-is* (§1), so the fork's inherited baseline becomes a TCIndustries value rather than merely evidence.

---

## 1. Layer (i) — character-layer skill-point hygiene cap

| Parameter | Authorised value | Basis |
|---|---|---|
| `MaxSkillPoints` (per character) | **250** | Fork-inherited baseline (SWG historical 250-point economy; GDD §6.3.1 references it). At 42 SP/profession (§6.3.1 corrected ladder), 250 permits ~5–6 professions mastered — the historical "broad but not total" coverage, consistent with Layer (ii) being the real mastery constraint. **Revisable** in a future BAL-001 re-tune. |

This cap is **subordinate** (HD-PROF-04 (i)): it keeps the character sheet legible and enforces the free-drop-on-overflow hygiene (HD-PROF-02 / PROF-004), but it does **not** by itself prevent single-character mastery — that is Layer (ii)'s job.

---

## 2. Layer (ii) — participant-layer interdependence budget (C2/C4 anchored)

### 2.1 Decay / maintenance cadence (Shape 3 element)

| Parameter | Authorised value | Basis |
|---|---|---|
| Decay scope | **Master-tier (Tier 5) boxes only** | Basic tiers (I–IV) never decay — preserves PROF-004 casual participation / hybrid-role hygiene. Only the capstone high-tier box is activity-bound. |
| Decay window | **30 in-game days** of no re-affirming activity in the role | Shape 3 (maintenance/decay). Re-engagement is cheap and free-drop exists, so lapse is recoverable, not punitive. Long enough to tolerate casual play; short enough that a dormant alt-army "mastery" cannot be parked indefinitely. |
| Lapse effect | Box reverts to Tier-IV state; Master must be re-trained (Master XP/credit re-paid) | No asset loss beyond the box; identity survives (PROF-004). |

### 2.2 Complementarity binding strength (C4 associational gating)

| Parameter | Authorised value | Basis |
|---|---|---|
| High-tier (Master) access prerequisite | **Active membership in a player association with ≥5 distinct participant-owners** (not all alts of one participant) **AND a role-provenance credential** | C4 (non-economic institutional/social constraint). The "≥5 distinct participants" floor blocks single-participant alt-internalisation of the institutional bond (HD-EIC-01 validation unit = participant). The provenance credential enforces C2 (persistent information/discovery asymmetry) — high-tier status depends on knowledge/provenance not trivially copied across an alt army. |
| Minimum association size `K` | **5** | Small-guild floor; low enough to be reachable, high enough that a lone participant cannot self-satisfy the bond. **Revisable.** |

This delivers PIL-003's LOCKED high-tier non-self-sufficiency as **soft comparative advantage** (HD-PROF-04 (ii)): a high-tier actor *can* perform the role but must sustain institutional + provenance bonds that cross the G2 "literal effortlessness" threshold, so broad self-coverage imposes sustained organisational effort, not mere capital.

---

## 3. Acquisition cost curve (PROG-003 + BAL-001) — authorised formulas & per-tier values

The fork derives costs from skill-point cost per GDD §6.3.2. This pass **authorises those formulas and their outputs** as TCIndustries values:

```
XP_cost(SP)     = 1000 × SP²
Credit_cost(SP) = round( XP_cost(SP) / 20 ) to nearest 10   (integer arithmetic)
```

| Tier | SP | XP cost | Credit cost |
|---|---|---|---|
| Novice | 0 | 0 | 0 |
| I | 1 | 1,000 | 50 |
| II | 2 | 4,000 | 200 |
| III | 2 | 4,000 | 200 |
| IV | 3 | 9,000 | 450 |
| Master | 10 | 100,000 | 5,000 |
| **Per profession total** | **42** | **118,000** | **5,900** |

- XP is **typed** (GDD §6.3.2): combat XP cannot buy crafting skills. The six `XPType` pools in the fork (combat, crafting, scouting, medical, entertainer, merchant) are authorised as the typed pools.
- The acquisition *gate* remains trainer-gated (HD-PROF-03): the credit cost above is the NPC-trainer path; a player who already holds the class may teach for free (no credit, no XP-at-source change).

**XP earn rate (the rate of progression, not the cost of a box)** is **deferred** — see §5, because it is entangled with combat/Phase-8 BAL-001 tuning and the unbuilt task layer.

---

## 4. What is now authorised vs. still deferred

| Magnitude | State |
|---|---|
| Character-layer cap (`MaxSkillPoints`) | **AUTHORISED = 250** (§1) |
| High-tier decay window | **AUTHORISED = 30 in-game days, Master tier only** (§2.1) |
| C4 association floor `K` | **AUTHORISED = 5 distinct participants** (§2.2) |
| Provenance-credential requirement for high tier | **AUTHORISED as a required binding** (§2.2) |
| Acquisition cost curve (XP/Credit = f(SP)) | **AUTHORISED formulas + per-tier table** (§3) |
| Typed XP pools | **AUTHORISED** (§3) |
| XP earn / progression rate (PROG-003 pace) | **DEFERRED — §5; re-examined 2026-10-25, still deferred (see OQ-001 Deferred BAL-001 Pass Disposition)** |
| Task-layer reward magnitudes (Q8) | **DEFERRED — §5; re-examined 2026-10-25, still deferred (no task layer exists)** |
| Task-layer sink magnitudes (Q8 / HD-ECO-01) | **DEFERRED — §5; re-examined 2026-10-25, still deferred (no matched sink)** |

---

## 5. Deliberately deferred (not invented)

These are **not** set here, despite being "numeric magnitudes," because their structural partners are absent — inventing them would violate BAL-001's false-precision caution and/or HD-ECO-01:

- **XP earn rate / PROG-003 progression pace.** The *cost* of a box is set (§3); the *rate at which XP is earned* is not, because PROG-003 (GDD §655) is still `TBD` and the fork's XP earn→spend pool loop is internally inconsistent (earn channels write pools that are not spendable). Re-examined 2026-10-25: still deferred — see `TCIndustries_OQ-001_Deferred_BAL-001_Pass_Disposition_2026-10-25.md` §1. A PROG-003 ruling + pool-taxonomy reconciliation must precede any rate authorisation.
- **Task-layer reward magnitudes (Q8).** No task layer exists in the fork (HISTORICAL only per HD-TST-01), so there is no hook and no magnitude to set. Re-examined 2026-10-25: still deferred — Disposition §2.
- **Task-layer sink magnitudes.** Companion to the above; deferred with it (and HD-ECO-01 requires a matched sink before any reward faucet is authorised).

---

## 6. Relationship to implementation (no code authorised here)

This pass **authorises magnitudes only**. It does **not**:
- modify the live fork's `professions.go` or any other code;
- authorise an elite-tier or task-layer *implementation* — that requires a separate implementation proposal built against these numbers and verified end-to-end per the testbed "verify, don't claim" discipline;
- change any GDD status beyond recording that the deferred magnitudes now exist (§7).

The phase-8 combat balance values (HD-PVP-01) remain independently BAL-001-gated; this pass does not touch them.

---

## 7. Status-line updates enacted by this pass

- GDD §31 OQ-001 row: `TBD → RULED (HD-PROF-04)`; its "magnitudes deferred to a BAL-001-authorised pass" note is now resolved by this document.
- Human Rulings Register HD-PROF-04: superseded its own "all magnitudes deferred" posture with the authorised values above (recorded as a BAL-001 numeric-authorisation subsection).
- Phase 9 replacement proposal §4 item #2 (PROG-003 + BAL-001 numeric pass): **RESOLVED by this pass**; the implementation gate it represented is now cleared for an implementation *proposal* (not yet for code — see §6).

*End of BAL-001 numeric pass (AUTHORISED 2026-10-25). Supplies the OQ-001 magnitudes HD-PROF-04 deferred; defers only the magnitudes whose structural partners do not yet exist.*
