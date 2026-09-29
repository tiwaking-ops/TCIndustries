# TCIndustries — Governance Rule GR-006 (Proposed)
## Detective and Corrective Control of Project Decisions

**Document Type:** Proposed Governance Rule (Non-Canonical) — draft for owner approval
**Version:** 0.1
**Date:** 2026-09-28
**Status:** APPROVED and AUTHORISED by the project owner 2026-09-28; **in force**
in `governance/TCIndustries_Authority_Provenance_Reconciliation_Matrix.md` §2
from that date, filed there as GR-006. The owner's four §4 implementation slots
(WINDOW, RETRO, REPORT, EDGE) were **not** answered and remain **undetermined**,
reserved to the owner; filed under the **partially-ruled precedent of HD-PROF-01** —
the rule is in force and the filing-window deadline is expressly unset rather than
invented. This draft's §5 paste block was **not** used verbatim: the placeholder
`[[WINDOW]]` was not carried into the matrix, because a governance rule
containing a literal placeholder is a defective rule. The filed text substitutes
an explicit UNDETERMINED marking. This document is retained unaltered below as
provenance.
**Authority created by this artifact:** None. This document creates no
game-design decisions and grants no design authority.
**Author:** OpenCode (space-bunny-free) — documentarian draft, per project
owner's instruction of 2026-09-28 ("authorise as a governance rule, not an EIC
boundary"). **Assessor:** none (self-drafted; no independent assessment).
**Owner-stated rationale (verbatim):** *"Allows greater detective and corrective
control of project decisions."*
**Placement:** destined for `governance/TCIndustries_Authority_Provenance_Reconciliation_Matrix.md`
§2, alongside GR-001–GR-005.

**Related but distinct — the EIC draft this was *not*:** on 2026-09-28 the owner
directed that this authorisation apply to a **governance rule**, not to the
HD-EIC-07 boundary statement
(`proposals/TCIndustries_EIC_Structural_Boundary_Statement_Draft_HD-EIC-07_2026-09-28.md`).
That draft remains **unfilled and unruled**; HD-EIC-07 remains **open**; EIC
architectural work remains **blocked**. This document does not clear that gate
and does not substitute for it. See §7.

---

# 1. The gap this rule addresses (observed, not hypothetical)

All five failures below were **actually observed in this repository** before
this rule was drafted. None is a projection.

| # | Observed failure | Evidence |
|---|---|---|
| **F1** | **Six owner rulings were undetectable in the authority record.** HD-ITM-01, HD-PROF-01, HD-BLD-02, HD-PROF-02, HD-ITM-02, HD-BLD-03 were ruled 2026-09-19/20 and existed as decision records in `proposals/`. A grep of `governance/TCIndustries_Human_Rulings_Register.md` for all six IDs returned **zero matches** — for approximately nine days. | Register Section D absent until filed 2026-09-28; queue-tracker line 5 ("Rulings land only in the register") was the only pointer, and it was not a check |
| **F2** | **Dependent documents actively contradicted the rulings.** With those six ruled, `canonical/gdd/..._Status_Patch.md` still showed OQ-007/010/011/015/016 as `TBD`, and `proposals/TCIndustries_Open_Questions_Register.md` showed OQ-010/011 as `TBD`. Three artifacts, three different stories about the same decisions. | GDD §31 lines 811–820 (pre-filing); OQR §2 rows OQ-010/OQ-011 (pre-filing) |
| **F3** | **The project memory had no trace of any of the six rulings.** `governance/project_memory.md` ended at 2026-09-18. | Pre-filing tail inspection, 2026-09-28 |
| **F4** | **The register's own summary table was stale relative to the register.** "Full Cumulative State" was dated 2026-08-25 while the register itself had reached v1.2 / 2026-09-19. A summary that silently disagrees with the body is worse than no summary. | Register line 226 (pre-filing) vs line 5 |
| **F5** | **Reserved filing acts were indistinguishable from forgotten ones.** Each decision record reserved "filing-time acts (owner/coordinator executes at register filing)". Nothing tracked whether they were *performed*. The reservation read as a plan and functioned as a deferral, so the gap was invisible in both directions. | HD-ITM-01 §4(2)–(5); HD-PROF-01 §4; HD-BLD-02 §4; HD-PROF-02 §4; HD-ITM-02 §4; HD-BLD-03 §4 |

**Common root cause:** the project has strong *classification* machinery
(authority classes, status vocabulary, GR-001–GR-005) and weak **detection and
correction** machinery. Nothing in the repository could tell an auditor — or the
owner — whether a decision that *should* be in the authority record actually was.
GR-004's own rationale already gestures at this ("provenance checks, audit
trails"); this rule makes it a checkable obligation.

---

# 2. The rule

## 2.1 Detective obligations

**GR-006.1 — Authority-record presence.** Every ruling the owner issues must
appear in `governance/TCIndustries_Human_Rulings_Register.md`, in full, within
`[[OWNER: WINDOW — the period within which a ruling must be filed]]` of the
ruling date. A ruling recorded anywhere else is a **pending filing**, not a
completed decision, until the register entry exists.

**GR-006.2 — Dependent-document consistency.** Where a filing changes a status
that appears in any other repository document, that document must be updated in
the same filing action. The set of dependent documents is not left to judgement:
at minimum `canonical/gdd/` (status tables and TBD entries), the Open Questions
Register, the Invariant Register, `governance/project_memory.md`, and the
Continuous Queue Tracker. **Divergence between the authority record and a
dependent document is a defect**, and is reported as one.

**GR-006.3 — Summary-table currency.** Where a document carries a summary,
index, or "current state" table that asserts its own currency, that table must
be brought current in the same action as any change it summarises. A dated
summary table that is knowingly behind the body must say so on its face.

**GR-006.4 — Filing acts are tracked, not merely reserved.** A decision record
may not reserve filing-time acts without recording them in a trackable state
(pending / done / waived-with-reason). An act that is neither done nor waived is
an open obligation, and an open obligation is reportable.

**GR-006.5 — The check must be mechanical.** Satisfying this rule requires a
**repeatable query**, not a judgement call. A filing is not closed until a
search for the ruling's ID returns a hit in the register **and** every dependent
document listed in GR-006.2. This is the property that would have caught F1
on the day it happened.

## 2.2 Corrective obligations

**GR-006.3-C — Correction is filing, not re-litigation.** When a gap is found,
the remedy is to **complete the record**, not to revisit or reverse the decision.
The ruling's authority is unaffected by how long it took to file it.

**GR-006.7 — Absence of disclosure or filing is a documentation defect, never a
voiding one.** This follows the precedent set by GR-004 and by the HD-GDD-01
provenance-note pattern. A late, missing, or incomplete filing record **never**
invalidates the underlying ruling, and never removes a LOCKED status. Nothing in
this rule may be used to argue that a decision the owner made is not a decision.

**GR-006.8 — Detection is owed to the owner, not only to agents.** Any agent may
and should surface a filing gap when it encounters one, as a **finding for the
owner**, with the specific IDs and file paths. Reporting a gap is never
permission to fix it unilaterally.

---

# 3. What this rule does not do

- **It does not change any design status.** No game-design decision, status
  tag, invariant, or numeric value is created, altered, or promoted.
- **It does not create a new authority class.** §1 classes A–G are unchanged.
  A governance rule is not a source of design authority.
- **It does not make an agent authoritative.** GR-006.8 explicitly withholds
  unilateral corrective power. The owner remains the sole rounder.
- **It does not retroactively void or require re-ruling anything.**
- **It does not alter GR-001, GR-002, GR-003, GR-004, or GR-005.**
- **It does not open, close, or substitute for HD-EIC-07.**
- **It does not require LLM consensus on anything.** One checkable fact, one
  owner ruling.

---

# 4. Open implementation choices (owner)

| Slot | Question | Note |
|---|---|---|
| **WINDOW** | What period counts as "filed in time"? | Left to the owner. No default is proposed — a default would be an invented threshold. |
| **RETRO** | Does this rule apply only forward, or does it impose a one-off reconciliation sweep of all prior rulings? | Forward-only is the minimal reading. A sweep is a separate authorisation. |
| **REPORT** | Who runs the mechanical check (GR-006.5), and at what cadence — per filing, or on request? | |
| **EDGE** | If the owner rules *inside* a session with an agent that drafts the register entry immediately, is that "filing," or does filing require a separate later act? | This materially changes the burden. Worth deciding explicitly. |

---

# 5. Ready-to-paste text for the Matrix §2

*Inert until approved by the owner.*

```markdown
### GR-006 — Detective and Corrective Control of Project Decisions
**Status: GOVERNANCE RULE.**

Every ruling the owner issues must be present in full in
`governance/TCIndustries_Human_Rulings_Register.md` within [[WINDOW]] of the
ruling date, and every repository document that reflects a status changed by
that filing must be updated in the same action — at minimum the canonical GDD
status tables, the Open Questions Register, the Invariant Register,
`governance/project_memory.md`, and the Continuous Queue Tracker. Divergence
between the authority record and a dependent document is a defect and is
reported as one. Documents carrying their own dated summary or "current state"
table must bring it current in the same action as any change it summarises.
Reserved filing-time acts must be tracked in a reportable state; an act neither
done nor waived is an open obligation. Closure requires a repeatable query, not
a judgement call: the ruling's ID must be found in the register and in every
dependent document.

When a gap is found, the remedy is to **complete the record**. Late, missing, or
incomplete filing never invalidates a ruling and never removes a LOCKED status;
this rule may not be used to argue that a decision the owner made is not a
decision. Any agent may surface a filing gap as a finding for the owner;
reporting a gap is never authorisation to correct it unilaterally.

*Owner rationale: "Allows greater detective and corrective control of project
decisions."*
```

---

# 6. Filing-time acts (reserved; owner/coordinator executes at filing)

1. Paste §5 into `governance/TCIndustries_Authority_Provenance_Reconciliation_Matrix.md` §2.
2. Bump the matrix version/date metadata and its change note.
3. Update `governance/README.md` to list GR-006 alongside GR-004.
4. **Decide GR-005's disposition while in the file** — GR-005 is recorded as
   "in effect now, project-owner confirmed 2026-08-25" in
   `proposals/TCIndustries_Manifest.md:55` but has **never been pasted into the
   matrix**. The matrix is currently one governance rule short. Merging GR-005
   and GR-006 together is a single edit; leaving GR-005 stranded keeps the
   manifest and the matrix disagreeing, which is the exact failure mode F2
   describes. **Flagged for the owner — not actioned here.**
5. Record in `proposals/TCIndustries_Manifest.md` that GR-004's text is now
   genuinely present in the matrix, retiring the stale line at
   `Manifest.md:53` ("Until that file is actually edited, this manifest is
   GR-004's operative home").
6. `governance/project_memory.md` event line recording the authorisation, the
   rationale, and the non-EIC disposition.
7. If **RETRO** is answered yes, authorise and scope the reconciliation sweep as
   a separate task.

---

# 7. Explicit non-effect on the EIC gate

Recorded here so this filing cannot later be misread as having cleared it:

| Item | State after this rule | Unchanged? |
|---|---|---|
| HD-EIC-07 structural boundaries | **OPEN** — human statement still outstanding | Yes |
| EIC architectural search | **BLOCKED** | Yes |
| OQ-001, OQ-004, OQ-005, T-03, T-06 | **Still coupled to the HD-EIC-07 gate** | Yes |
| The EIC boundary-statement draft | **Unfilled, unruled, PROPOSED** | Yes |
| Architecture C | **RETIRED** (HD-EIC-05) | Yes |
| Numerical boundary statement | **Not made** | Yes |

The owner's 2026-09-28 authorisation was, on the owner's own direction, a
**governance** authorisation. HD-EIC-07 remains the register's standing "Next
required human action."

---

**Document Control**

**Status:** PROPOSED — HUMAN APPROVAL REQUIRED. Not a governance rule yet.
**No design authority created.** **No status promoted.** **No filing-gap repair
authorised by this document.**
**No retroactive requirement imposed.** **No unilateral corrective power granted
to any agent.**

---

*End of GR-006 draft. To enact: the owner answers §4, then approves §5.*
