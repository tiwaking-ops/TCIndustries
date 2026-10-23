# SWG Historical Accuracy — Core-Loop Comparison Checklist v0.2

**Status:** PROPOSED candidate material. **No design authority.** This document is a *procedure*, not a
decision. It closes no TBD, settles no open question, creates no invariant, and authorises no tuning.
Nothing in it promotes any value, mechanism, or number to TCIndustries canon. Executing it produces
**EVIDENCE** about how the fork's core loops compare to Pre-CU accounts — it does not produce a ruling.

**Instrument author (llm):** MiniMax M3.1 Flash (free) — model `openrouter/minimax/minimax-m3.1-flash:free`,
via OpenCode, 2026-10-22.
**Provenance note (GR-007.1 / GR-007.3).** This is the same identity that produced the assessment half of
`investigations/TCIndustries_SWG_Historical_Recollection_Report_Tiwa_Pene_2026-10-21.md` (v1.6). The
instrument is a **downstream consumer** of that filed report, not a new assessment of it, and the report's
findings are cited as filed rather than re-derived here. Independent assessment of *this instrument* is
available on request and is **recommended from a different model**, since this identity authored both the
underlying analysis and the instrument built on it.
**Subject-matter source (human):** Tiwa Pene — Human Operator — respondent and author of the recollection
record on which the Pre-CU side of every row below rests.
**Assessor (llm):** — not yet assessed.
**Version:** v0.2
**Date authored:** 2026-10-22 (v0.1); 2026-10-22 (v0.2 — status corrections, §3.6; see §11)

**Responds to:** Item 8, Historical Accuracy limb — the 5th of the five §1.3 success-criteria tests
(`proposals/SWG_Phase10_Balance_Telemetry_Polish_Proposal_v0.2.md` §2.2, §5.3, §6.4; §9 item 5). Per
`proposals/SWG_Post_Phase10_Program_Proposal_v0.1.md` §9.1 item 8, the split decision of 2026-10-18
resolved the Historical Accuracy limb by *authoring* a formal instrument (documentation act only, no code,
no design authority) and leaving it to the human. This checklist is the **next** artefact that limb calls
for: the instrument that turns a filed account into a runnable comparison. **Proxy 4 (Sandbox Validation) is
handled separately and remains DEFERRED** (no live cohort); it is not touched by this document.

**Authority basis:** `AGENTS.md` (global + project); `governance/TCIndustries_Authority_Provenance_Reconciliation_Matrix.md`
GR-001, GR-002, GR-003, GR-004, GR-006.8, GR-007.1–GR-007.6; `governance/TCIndustries_Testbed_Reuse_Human_Ruling_HD-TST-01_2026-09-14.md`;
`proposals/SWG_Phase10_Balance_Telemetry_Polish_Proposal_v0.2.md` §§1.2, 2.2, 2.3, 5.1, 5.2, 5.3, 6.4, 9;
`proposals/SWG_Post_Phase10_Program_Proposal_v0.1.md` §§9, 9.1, 10;
`investigations/TCIndustries_SWG_Historical_Recollection_Report_Tiwa_Pene_2026-10-21.md` (v1.6);
`canonical/gdd/TCIndustries_Master_GDD_v1.1.1_Status_Patch.md` §2 (central loop, anti-goals);
`sources/swg-pre-cu/SWG_PreCU_GDD.md` (HISTORICAL reference text only).

---

## 0. What this document is, and why it exists now

### 0.1 The obligation

The §1.3 success-criteria set has five tests. Two of them are qualitative by definition and are routed to
a human:

| §1.3 test | Method | Where it lives |
|---|---|---|
| Sandbox Validation (Proxy 4) | observation-log template (event, date, witnesses) | Phase 10 v0.2 §5.3 — **DEFERRED** |
| **Historical Accuracy** | **human checklist** | Phase 10 v0.2 §2.2, §5.3 — **this document is the missing instrument** |

Phase 10 v0.2 §5.3 states the constraint exactly: *"Sandbox Validation = observation-log template (event,
date, witnesses); Historical Accuracy = human checklist — the harness reports their existence/absence, never
a fabricated score."*

### 0.2 Why the checklist can be written now, and why it could not be written before

A historical-accuracy comparison needs **two sides**: the Pre-CU account, and an observation of the fork's
own behaviour. Until 2026-10-22 the Pre-CU side had only `sources/swg-pre-cu/SWG_PreCU_GDD.md` — a
*design document* describing what Pre-CU SWG was specified to do. It did not describe how the loops
actually played or felt, which is the thing a "core loops vs documented Pre-CU accounts" comparison
actually turns on.

`investigations/TCIndustries_SWG_Historical_Recollection_Report_Tiwa_Pene_2026-10-21.md` (v1.6) supplies
that side. It is a 279-item recollection record (223 questions in the first pass, 56 in a targeted second
pass) from a respondent who held **every class mastery** in Pre-CU SWG, organised into 24 instrument
categories, with a design-significant synthesis at §5.1–§5.6, a project-surface map at §5.2, and a
mechanism-level negative-result register at §6.

### 0.3 What this document is

A **fill-in comparison sheet**: for each core loop, the attested Pre-CU account (cited to question number),
the corresponding observation point in the fork, and an empty verdict slot for a human to fill.

### 0.4 What this document is not

- **Not a scoring system.** No rows are weighted, summed, averaged, or given a numeric verdict. A checklist
  with a composite score would be exactly the "fabricated score" Phase 10 v0.2 §5.3 forbids.
- **Not a pass/fail gate.** §1.3's exit criterion is explicitly disjunctive — *"the five success-criteria
  tests can be run and pass (**or** produce actionable telemetry showing where they don't)"*
  (Phase 10 v0.2 §2.1). An honest "DIVERGES" or "INSUFFICIENT EVIDENCE" is a *successful* run.
- **Not an authority record.** It cannot close OQ-009, OQ-010, OQ-011, OQ-012/SAFE-002, OQ-016,
  SERV-001/SERV-002, OQ-007 or BLD-004. Note that the report's characterisation of these surfaces as
  *"exactly as open"* is **stale** — six of the eight had already been ruled before the report was written.
  **§3.6** carries the corrected live status table; the report's §5.2 column is not relied on here.
- **Not a source of tuning values.** Every magnitude in the Pre-CU column is a **recollected magnitude**,
  never a target, floor, ceiling, or reference point (report §1.1, §7 item 7, §9). BAL-001 and HD-TST-01
  are unaffected by this document.
- **Not executable by an agent.** §1.1 below sets the execution rules; the verdict column is a human's
  to write.

---

## 1. The comparison model

### 1.1 Row shape

Every row has the same six columns:

| Column | Contents |
|---|---|
| **ID** | `HA-xx` row identifier. Stable. Citable in an observation log. |
| **Core loop** | The TCIndustries core loop under test, named in the GDD's own vocabulary. |
| **Pre-CU account** | What the recollection record attests, with a question-number citation and the report finding that carries it. |
| **Account strength** | `ATTESTED` / `THIN` / `ABSENT` — see §6. |
| **Fork observation** | Where the owner looks in the build, and what is looked at. Telemetry surfaces are named by their real REST endpoints. |
| **Verdict** | Empty until a human fills it. Vocabulary in §1.2. |

### 1.2 Verdict vocabulary

Five values, and the distinction between the middle three is the point of the whole instrument.

| Verdict | Meaning |
|---|---|
| **MATCH** | The fork's observed behaviour on this row is consistent with the attested Pre-CU account. |
| **DIVERGES** | The fork's observed behaviour is measurably or visibly different from the attested Pre-CU account. |
| **INSUFFICIENT EVIDENCE** | The Pre-CU account is too thin, or the observation too weak, to decide. **This is not MATCH.** |
| **NOT OBSERVABLE HERE** | The row cannot be exercised in the current environment (no live cohort, mechanic not built, etc.). Distinct from INSUFFICIENT EVIDENCE: the *account* is fine, the *arena* is not. |
| **NOT APPLICABLE** | The loop is out of scope for the current build or milestone. |

**Why INSUFFICIENT EVIDENCE is not a soft MATCH.** Report §7 item 6: *"It does not evidence any SWG system
as 'working as designed'. An affirmative one-word answer is weak evidence of satisfaction and near-zero
evidence of mechanism."* The report's densest categories (A, D, E, F) and its thinnest (L, M, N, O, R) are
separated by exactly this distinction. A checklist that collapsed them would manufacture agreement the
evidence does not support.

**Why NOT OBSERVABLE HERE is separate.** Report §3.3 / F-8 established that answer density tracks
**instrument fidelity to Pre-CU**, not respondent expertise — the same respondent produced 91.1% substantive
prose against 47.1% once the stems were written against the actual game. A row that could not be answered
because the question was not asked in Pre-CU terms is an instrument artefact, not a data point. Both
INSUFFICIENT EVIDENCE and NOT OBSERVABLE HERE are honest non-answers; they are logged differently so that
a later reader can tell a *thin account* from a *wrong question*.

### 1.3 Evidence class

Each row's Pre-CU citation carries an evidence class, so that a verdict is never read as carrying more
weight than its source:

| Class | Meaning |
|---|---|
| **RECALLED** | Human recollection, recalled after the fact. Magnitudes are recalled, not measured. |
| **DOCUMENTED** | Attested in `sources/swg-pre-cu/SWG_PreCU_GDD.md` or another written reference text. |
| **MEASURED** | Observed in the fork with a real count. (Only ever applies to the *fork* column — the Pre-CU side has no measured data at all: report §7 item 5, item 19.) |
| **INFERRED** | A reading placed on the record by the assessor's analysis. |

### 1.4 Observation modes

A row may be exercisable in one or more of three modes. Mark the mode on the execution log.

| Mode | What it means | What it needs | What it can establish |
|---|---|---|---|
| **A — DESIGN** | Compare the fork's *specified* behaviour (GDD, ruling records, phase proposals) to the Pre-CU account. | Nothing but documents. | Whether the *design* is faithful. Cannot establish that the build behaves so. |
| **B — FORK** | Observe the running fork's behaviour and telemetry. | The built systems, a cohort, a run. | What the build *does*. Limited by cohort size and by what is instrumented. |
| **C — COHORT** | Observe real players doing the loop, and judge whether the loop *plays* as the account describes. | A live player cohort and a human observer. | Whether the loop *feels* right. This is the only mode that can answer "is this historically accurate" in the sense the test intends. |

**Mode C is where Part 2 actually lives**, and it is also the mode that is currently unavailable: Proxy 4
was deferred for exactly the reason that no live cohort exists. Modes A and B are executable now against
the testbed, and are worth running — but they are *provisional* against a ~3-character cohort
(Phase 10 v0.2 §5.6), and a Mode B DIVERGES verdict on a cohort that size may be an artefact of the cohort.

**Mode A and Mode B results do not clear Mode C.** A completed Mode A/B run says: *the design is
consistent with the account, and the build is consistent with the design.* It does not say the loop plays
like Pre-CU SWG. That is the honest limit of what this checklist can currently deliver.

---

## 2. Standing rules for execution

These bind whoever fills the sheet.

1. **Observe, never steer.** Phase 10 v0.2 §1.2: validate = observe + report. A checklist row that
   reveals a divergence produces a *finding*, never a fix. Any change requires a separate authorised
   proposal (BAL-001).
2. **No fabricated verdicts.** Every unfilled row stays unfilled. `INSUFFICIENT EVIDENCE` and
   `NOT OBSERVABLE HERE` are correct, valuable answers. Phase 10 v0.2 §5.2 forbids simulated-player
   fabrication to manufacture a PASS.
3. **No numeric changes, no mechanism changes.** BAL-001 (LOCKED), HD-TST-01, and Phase 10 v0.2 §10
   ("zero numeric changes … zero gameplay-behavior changes"). This checklist is a read-only activity.
4. **Recollected magnitudes are never targets.** Any figure appearing in the Pre-CU column is a recalled
   magnitude. It may be quoted as a recalled magnitude and nothing else (report §1.1, §7 item 7).
5. **The historical record outranks the recollection where they conflict.** Report §7 item 8: *"Where it
   conflicts with the historical record of SWG Pre-CU, the historical record stands and this record is simply
   a recollection."*
6. **A row never closes an open question, and a row on a ruled surface never re-opens it.** The checklist
   produces evidence for a ruling the owner may later make. It is not that ruling (GR-001). Where a row
   bears on a surface that is **already ruled** (§3.6), a DIVERGES verdict is recorded as a **re-opening
   question for the owner** — the ruling stands until the owner moves it, and the divergence is logged
   beside it, never over it.
7. **Raw counts beside every percentage.** Phase 10 v0.2 §5.6: a 40/42 cohort is 95%, not a verdict.
8. **Record the cohort.** Every Mode B/C run logs cohort size, composition, and duration. A verdict
   without a cohort is not comparable to another verdict.
9. **Cite the report, not the report's findings-at-large.** Pre-CU claims are cited by question number
   (Q-n, C-Q-n) so a reader can land on the verbatim answer at §11 or §10C.
10. **Do not edit the Pre-CU column.** It is a quotation of a filed record. If it is wrong, that is a
    finding about the record, filed separately — not an edit.

---

## 3. Pre-checklist notes — three Phase 10 §2.3 bands with new provenance

Phase 10 v0.2 §2.3 lists HISTORICAL bands for validate-against comparison. Three of them now carry
provenance questions that the recollection report surfaces. **These notes are questions for the owner, not
corrections.** The bands are validate-against-only, none is a tuning target, and BAL-001 is unaffected.
Nothing in this section changes a band. Per §2 rule 5 and report §7 item 8, the owner decides.

### 3.1 `stall commission 4–6%` — uncorroborated, and possibly the wrong shape

Phase 10 v0.2 §2.3 (Economic): *"stall commission 4–6% (stalls dormant in fork — recorded, no value)"*.

The report finds this figure is the **weakest of the three** figure-bearing stems in the first instrument.
Report F-2, at the Q191 row: *"**UNCHANGED — still the weakest of the three.** The author never paid
listing fees (Q191, Q89) and so has no independent figure to offer. The `4–6%` value therefore remains
**uncorroborated and unsupported** by this record."* Report §7 item 5's residual anchoring concern is
**Q191 alone** (§7, item 5 closing note).

There is a second, distinct question. The report's mechanism-level account of the operative fee is
**not a percentage**:

- F-17 (C-Q15, C-Q18, Q18): a **20-credit listing fee rising to 100 credits** for premium (Merchant-only,
  Business II) — a flat credit charge, not a percentage of sale value. Merchants received fee reductions.
- F-30 (D-Q15): the fee was **economically inert and bypassable** — listing at 0 credits made the fee
  free, so a vendor became free storage, an XP faucet and a sales channel simultaneously. Report §5.2's
  OQ-016 row records *"Vendor fees were negligible and bypassable"* and that *"the operative sink was
  vendor **maintenance**, not the fee."*

**So the question for the owner is two-part:** (a) does the 4–6% band have any documentary support
independent of the respondent, given the respondent never paid it; and (b) is a *percentage* band the right
comparand at all, when the attested Pre-CU fee structure is a flat credit charge that was driven to zero?
Note the band is currently inert in the fork (*"stalls dormant in fork"*), so nothing is being compared
today — which is precisely why this is cheap to settle now and expensive to settle later.

### 3.2 `TTK 30–60 s PvP 1v1` — the figure belongs to group PvP

Phase 10 v0.2 §2.3 (Combat): *"TTK 20–40 s solo white-con, 60–120 s group-vs-elite, 30–60 s PvP 1v1"*.

The 30–60s figure entered the instrument as a **stem figure** at Q119 and was assessed as *figure-echoed*
(report F-2, Q119 row). Addendum A then re-scoped it (report F-5, **closed at v1.1**):

> *"PvP 1v1 will probably not be 30-60s TTK. In GROUPS TTK would be 30-60s as it would be multiple people
> against one person or focus fire or unfair matchups. Mathematically and practically speaking, 10 minutes
> would be a normal but long fight and that would not include healing."*

Report F-5's consequent model: **1v1 ≈ 10 minutes** excluding self-healing (Q6 allows 10–30 minutes with
self-healing); **group PvP / focus fire 30–60 seconds**; buffed-vs-unbuffed 10–20 min collapses to 1–5 min
or a one-hit kill (Q23).

The report's own note on this instance is the general one: *"This is a second instance of D-1: the
instrument's stems imported generic-MMO figures, and here the imported figure was wrong."*

**Question for the owner:** the band as written applies a figure the filed record re-scopes from 1v1 to
group PvP. Note the direction of authority carefully — the report is EVIDENCE and the historical record
outranks it where they conflict, so this is a flag for the owner, not a correction. Also note the band is
**validate-against-only**; it is not a target, and the fork's own combat numbers are reported beside it
whether or not it is amended.

### 3.3 Progression bands — one of four is corroborated, three are unaddressed

Phase 10 v0.2 §2.3 (Progression time): *"basic mastery 40–60 h; elite 80–120 h; multi-prereq elite
150–250 h; full 250-point cap 200–300 h cumulative."*

- **`150–250 h` (multi-prereq elite) — CORROBORATED.** Report F-2a: the figure was initially challenged
  and then **withdrawn as a contradiction**; Addendum A corroborated it at **240 h** for the *unassisted
  from-scratch* path, with **~24 h** for the *assisted* path. F-2a's consequence: *"The `150–250h / 240h`
  figure is now **corroborated** and can be carried forward as a recalled magnitude (still not a tuning
  target, and still conditional on the unassisted from-scratch path)."*
- **`40–60 h` (basic mastery), `80–120 h` (elite), `200–300 h` (full cap) — NOT ADDRESSED.** Q34 asked
  "How long did basic mastery feel? Elite mastery? Full cap?" and the record's answer survived on the
  elite paths only. **The report carries no counterpart for the other three bands.**
- **§2.3's own band comment remains accurate:** these are *projections*, not measurements (Phase 10 v0.2
  §2.3 — *"Unmeasurable live — projection method (§6.3)"*).

**Question for the owner:** three of the four progression bands have no documentary Pre-CU comparand in this
repository. When a projection is validated against a band with no comparand, the validation is vacuous and
should be logged as such rather than reported as agreement.

### 3.4 A unit-comparison caution (not a conflict)

Phase 10 v0.2 §2.3 (Combat): *"HAM 3000–6000 unbuffed, 6000–12000 buffed."* The record's HAM figures
(1,000 minimum per buff; 10,000–12,000 from top-tier doctors — Q-C5, F-7) describe **a single buff's
contribution to a pool**, not the pool total. A buffed pool of 10,000–12,000 is consistent with a
6,000–12,000 band and the 1,000 figure is a per-buff minimum, not a pool floor. These are different
quantities and must not be compared directly. Recording this prevents a false DIVERGES.

### 3.5 Provenance note on the bands themselves

The 30–60s, 150–250h and 4–6% figures all entered the corpus as **stem figures drawn from the predecessor
GDD** and were flagged as potentially anchoring at instrument filing
(`proposals/SWG_Historical_Accuracy-v1.md`, advisory finding). That pre-flag is what made the post-hoc
audit possible. It is a useful precedent: the Phase 10 §2.3 bands and the Pre-CU instrument's stems are
the **same class of object** — a figure asserted in advance and later either corroborated, re-scoped, or
left unsupported.

### 3.6 Live status of the surfaces the report names — the report's status column is stale

**This section was added at v0.2. It corrects a factual error the v0.1 draft inherited from the report.**

The report's §5.2 table carries a *"Status in the project"* column, and its closing assurance is that
*"Every entry is usable as precedent and as a warning; none is usable as a value"* — with the **"Closed?"**
column answering **No** for all eight named surfaces. The assurance that they *"remain exactly as open"*
(§7 item 11, §9) **does not match the authority record.** Six of the eight were **already ruled** when the
report was authored on 2026-10-21.

The report's authority basis cites the Authority & Provenance Reconciliation Matrix but **not**
`governance/TCIndustries_Human_Rulings_Register.md`, and the Open Questions Register v1.2 carries the
corrected statuses. Per the project `AGENTS.md` authority model, the **register is the authority**, not the
report. Corrected table:

| Surface | Report §5.2 says | **Live status** | Ruling | Filed |
|---|---|---|---|---|
| OQ-007 — city formation & governance | "TBD" | **RULED** — Option A, free-claim settlement formation; basic governance instruments (ownership + zoning presence + shared maintenance); **civic depth TBD** | HD-BLD-01 | 2026-09-19 |
| BLD-004 — abandonment handling | "TBD" | **RULED** — at zero maintenance credit the structure is removed from the world; contents removed with the structure (HD-BLD-03). **Placement / lots / density remain TBD** | HD-BLD-02 + HD-BLD-03 | 2026-09-19 / 2026-09-20 |
| OQ-009 — currency & money sinks | "TBD" | **RULED** — Structure A, single unified currency; INV-009a approved; **faucet-exceeds-sink mandate** | HD-ECO-01 | 2026-09-28 |
| OQ-010 — durability / decay / repair | "TBD; models A/B/C" | **RULED** — Option C, **no baseline decay**; repair demand carried by destruction/loss/obsolescence/upgrading; **numerics TBD** | HD-ITM-01 + HD-ITM-02 | 2026-09-19 / 2026-09-20 |
| OQ-011 — respecialisation cost | "TBD; cost classes A/B/C" | **RULED** — cost instrument **free-drop** (points returned, no fee/time/penalty); acquisition **trainer-gated**; **magnitudes TBD** (BAL-001). INV-011c and scope confirmation still open | HD-PROF-01 + HD-PROF-02 + HD-PROF-03 | 2026-09-19 / 09-20 / 09-28 |
| OQ-016 — vendor & retail structure | "TBD; models A/B/C" | **RULED** — Option A, identity-anchored player vendors (all sell-side capacity player-owned and player-stocked); NPC commerce limited to baseline services, never sells player-producible goods; INV-016a + INV-016b approved; **numerics TBD** (BAL-001) | HD-RET-01 | 2026-09-19 |
| SERV-001 / SERV-002 — services | SERV-001 LOCKED (principle only); SERV-002 PROPOSED | **Correct as stated** | — | — |
| OQ-012 / SAFE-002 — multi-account & exploit | "TBD" | **Correct in substance** — OQ-012 remains **TBD at the implementation layer**; the validation stance (in-scope, **no account-level restriction**) is settled per HD-EIC-03 | HD-EIC-03 | 2026-08-25 |

**Consequence for this checklist, and why it matters more than a bookkeeping note.** Four of the
corroborations below are **post-hoc**: the rulings predate the recollection report, so the report's
attestations were not available when the rulings were made. They are now available, and they point in
**both** directions:

| Surface | Ruling | What the record attests | Relationship |
|---|---|---|---|
| OQ-010 | **No baseline decay** (Option C) | Decay existed in the manual and the code **and did nothing in play**; there was **no repair economy** (F-15, F-16). The report states the record's effective historical answer was also **C** | **CORROBORATES** the ruling. The respondent's own stated preference is C-with-later-implementation (R§5.2 OQ-010 row) — but the respondent is not the ruling authority |
| OQ-011 | **Free-drop** respec | Respec was **free, unlimited, no cooldown, no cost per removal**; skill points returned, XP not refunded (F-14, C-Q10/C-Q11/C-Q12) | **CORROBORATES** the ruling, and adds the mechanism the ruling did not record: **XP is not refunded, skill points are** |
| OQ-016 | Player-vendor primacy (Option A) | Vendors were **granted by skill, not crafted**, bore their owner's identity, and NPC commerce did not displace the player economy (F-17, F-30, R§5.2 OQ-016 row) | **CORROBORATES** INV-016b. The record also attests the Pre-CU fee structure was **flat credits, not a percentage** (§3.1) — which bears on the still-deferred numerics |
| OQ-009 | **Faucet-exceeds-sink** mandate | Mission payouts dominated income; the operative sinks were housing maintenance, vehicle maintenance, player cities and insurance — and vendor fees were **inert and bypassable** (F-17, F-19, F-20) | **CORROBORATES** the mandate's shape. Also supplies two sink classes (**vehicle maintenance, insurance**) the record says are *new to it* and **not thereby adopted** (R§5.2 OQ-009 row) |
| BLD-004 | **Removal** at zero maintenance credit | Unpaid maintenance **damaged** structures; at 670 h the structure was affected, and destruction of the City Hall **disbanded the whole city** (F-23, F-34) | **DIVERGES on mechanism.** The ruling chose *removal*; the historical account records *damage-at-lapse* plus *dissolution-on-City-Hall-destruction*. **HA-29 is a live DIVERGES candidate, not a match** |

**This does not put any ruling in question.** A Pre-CU recollection is not authority, and the project
`AGENTS.md` is explicit that historical inspiration is never automatic canon. The rows are written so that
a DIVERGES on a **ruled** surface is recorded as a **re-opening question for the owner**, not as a defect
asserted by this document (§2 rule 6, §8.2). The divergence is the deliverable; the decision is not.

**Note on the report itself.** The stale status column is a dependency-consistency defect against the
authority record (GR-006.2). Per GR-006.8, **reporting a gap is never authorisation to correct it
unilaterally**, so the report has **not** been edited by this session. The remedy, if the owner wants it,
is a v1.7 amendment to the report carrying an explicit, dated *"Status column corrected"* note — not a
silent edit, and not a rewrite of the record's own findings. This is filed as a finding for the owner in
`investigations/OpenCode_Project_Documentation_Report_2026-10-22.md` §6.

---

## 4. The checklist rows

Pre-CU citations are to `investigations/TCIndustries_SWG_Historical_Recollection_Report_Tiwa_Pene_2026-10-21.md`.
`R§x` = that report's section. `F-n` = its assessment finding. `Q-n` = first-pass answer (§11);
`C-Q-n` = second-pass answer (§10C). Fork observation points name real surfaces:
`GET /api/phase10/telemetry`, `GET /api/phase10/anomalies`, `GET /api/phase10/jobs`.

**GDD central loop (the spine):** *Discover resources → acquire and process materials → craft goods or
provide services → specialise → establish reputation → participate in trade, society, and regional
economies.* (Master GDD v1.1.1 §2.)

---

### Domain A — Resource discovery & harvesting

| ID | Core loop | Pre-CU account | Strength | Fork observation | Verdict |
|---|---|---|---|---|---|
| **HA-01** | Discover resources — survey, sample, depletion | Surveying was exploration or a chore (Q69); resource spawns were dynamic and interesting vs predictable (Q70); "shift" changes in resource stats were meaningful gameplay (Q71); sampling was engaging or repetitive (Q72); quality mattered enough to seek specific spawns (Q73); depletion felt fair or frustrating (Q74); good resources were scarce and valuable (Q75). R§5.4, R§3.1 cat. E (8/9 substantive). | ATTESTED (RECALLED) | Design: resource tree, spawn tables, depletion rules. Fork: `profession_distribution` and `price_volatility` rows; per-resource spawn counts. | |
| **HA-02** | Harvesting as an economy, not a chore | Creature products and fishing were a **separate resource path** from gear loot — SWG had **no gear in creature loot at all**; loot was junk worth ~20 credits (Q-C6, F-6, R§5.4). Materials and gear were distinct economies. | ATTESTED (RECALLED) | Fork: does any gear enter inventories without a schematic? `crafted_share.with_schematic` vs `total` — the 95% crafted-share proxy is exactly this question, read as counts not percentage. | |
| **HA-03** | Creature ecology & spawn governance | Spawns were **mission-based or random**; apparent thinning is mission spawns appearing or disappearing; long-term thinning or migration **is not documented and did not happen** (Q30, F-33, R§7 item 15). Creature behaviour: hunt as a group, fight one at a time, or flee (Q26). | THIN (Q26 substantive; Q30 is a negative attestation) | Fork: `mission_outcomes` mix read against spawn visibility. Design: spawn lifecycle vs mission spawns. | |
| **HA-04** | Threat legibility before combat | Threat was legible **before** combat via a colour classification — and that legibility could mislead (Q130, F-21, R§5.1). What information was available about a creature, and what was not (Q27). | THIN (first pass, one-word heavy) | Design: con/colour system, any pre-combat threat indicator. | |

---

### Domain B — Acquisition & processing

| ID | Core loop | Pre-CU account | Strength | Fork observation | Verdict |
|---|---|---|---|---|---|
| **HA-05** | Material acquisition feeding production | Resources were hoarded and this affected markets (Q76); crafting inputs were specific and finite (C-Q32: 85 Iron, 40 Ferrous metal, 28 Inert petrochemical — RECALLED magnitudes, not targets). | ATTESTED (RECALLED) | Fork: `price_volatility` for resource rows; resource consumption vs availability. | |
| **HA-06** | Player-to-player trade of materials | Resource hoarding was a **market-shaping** behaviour, not just hoarding (Q76) — a player behaviour, not a designed sink. | THIN (Q76 is one-word; the market-shaping *reading* is R§3.1 cat. C inference) | Mode C only: did players hoard, and did it move prices? | |

---

### Domain C — Crafting & item creation

| ID | Core loop | Pre-CU account | Strength | Fork observation | Verdict |
|---|---|---|---|---|---|
| **HA-07** | Crafting as a *process* with partial determinism | **Materials set the weights, the experiment set the roll** (F-22, C-Q33/C-Q34) — partial, not full, determinism. Quality varied; a good roll mattered; players chased perfect rolls; generic taxonomy (frame/grip) did or did not make sense; assembly was a real process or a formality (Q55–Q64). | ATTESTED (RECALLED) | Design: the fork's RNG split between material and experimentation stages — the F-22 shape. Fork: distribution of crafted output quality across repeated runs of the same schematic. | |
| **HA-08** | Crafting as a *market* activity | Crafted items must be **meaningfully better than NPC/looted items** (Q65) — a design premise the GDD also carries. Factories were a meaningful endgame (Q66) or AFK (Q67); factory spam was a problem (Q68). Craft commissions had **no trust mechanism** (F-22). | ATTESTED (RECALLED) | Fork: `crafted_share` (proxy 3) as counts; `price_volatility` for crafted vs non-crafted. | |
| **HA-09** | Durability / decay / repair | **Decay existed in the manual and in the code and did nothing in play** — "intermittent and buggy and mainly didnt work or was too slow to do anything"; **there was no repair economy** (F-15, F-16, C-Q5/C-Q6, R§6.4). The two decay systems that did operate were the **death penalty** and the **civic maintenance obligation**. An item at zero condition became a **1/1 item that could not be destroyed**, by an early bug (F-16, F-31). | ATTESTED (RECALLED) | Design: the fork's durability model — **OQ-010 is RULED** (HD-ITM-01 Option C, no baseline decay; HD-ITM-02 death-damage/repair/destruction behaviour; numerics TBD under BAL-001). Fork: repair/consumption events per item if instrumented. | |
| **HA-10** | What is *not* reproduced from Pre-CU | Explicit negative result: *"For a new game starting without decay would be the correct choice and it can be implemented later."* (C-Q5 note, R§6.4) — the respondent's own design opinion. **Not authority**; recorded because it is a named negative result, not a recommendation for TCIndustries. | ATTESTED (RECALLED, opinion) | Design: **OQ-010 is RULED — no baseline decay.** The row therefore observes whether the build honours the ruling, and whether it reproduces the historical decay/damage pattern the record attests. **The respondent's preference is not authority and does not re-open the ruling.** | |

---

### Domain D — Services & the service economy

| ID | Core loop | Pre-CU account | Strength | Fork observation | Verdict |
|---|---|---|---|---|---|
| **HA-11** | Providing services as an economic role | **Service XP was act-based for every profession except the Merchant** (F-12, C-Q1). The Merchant's yield was a fixed **50 XP per active vendor-hour**, paid **regardless of sales**, with interaction XP capped once per 15 minutes (F-11, C-Q1/C-Q2). | ATTESTED (RECALLED) | Fork: `interdependence` covered/total (proxy 1) with raw counts; per-service-type event counts. | |
| **HA-12** | Whether services can be a *career* | Service professions were **weak earners**; players migrated to combat professions for money; **Doctor buffs were the one income exception** (F-12, C-Q3). Separate XP pools per activity meant a thin service economy **could not be rescued by mission XP** (F-13). | ATTESTED (RECALLED) | Fork: `profession_distribution` — is the service cohort populated or empty? Mode C: do players *choose* service professions? | |
| **HA-13** | Buffs & the social economy of the buff tier | The top tier of the buff economy was a **social/guild privilege, not a purchasable service** (F-7, Q-C5). Buff duration ~4 h (or ~2 h for normal buffs) — RECALLED. **Buff begging** was a social problem (Q92). | ATTESTED (RECALLED; buff-begging is one-word) | Design: buff HAM magnitudes vs the §2.3 band — read with §3.4's unit caution. Mode C: social tiering of buffs. | |
| **HA-14** | Medical / wound services | Wounds were a post-combat consequence (Q93); wound healing was accessible or scarce (Q94); wound rot / neglect was a real risk (Q95). Entertainers **healed the mind-wound and fatigue residue in cantinas** — a systemic penalty creating demand for a non-combat service profession (F-24). | THIN (cat. G mixed, 3/6 substantive) | Fork: heal events in `interdependence`; entertainer-service events specifically. | |

---

### Domain E — Specialisation & progression

| ID | Core loop | Pre-CU account | Strength | Fork observation | Verdict |
|---|---|---|---|---|---|
| **HA-15** | Progression pace by path | **240 h** unassisted from scratch vs **~24 h** assisted (F-2a, Q-C1) — RECALLED magnitudes for elite paths, **scope-limited**. Paths diverged for identifiable reasons: Bounty Hunter XP arithmetic against a 300,000–500,000 XP tree at ~30 min/mission on foot; Commando heavy weapons' HAM cost fighting its own XP loop; Master Carbineer/Pikeman poorly balanced; Smuggler slicing slow (F-2a). | ATTESTED (RECALLED) | Design: the fork's progression projection method (Phase 10 §2.3 / §6.3) against §3.3's band caveats. | |
| **HA-16** | Respecialisation and the retraining market | Respec was **free, unlimited, no per-day limit, no total limit, no cooldown, no cost per removal**; the **experience was not refunded** (the skill points were) (F-14, C-Q10/C-Q11/C-Q12). Because it was free and XP-destructive, a **paid retraining market formed** — 1,000 credits per skill on the author's server. | ATTESTED (RECALLED) | Design: **OQ-011 is RULED** (HD-PROF-01/02/03 — cost instrument free-drop, acquisition trainer-gated; magnitudes TBD under BAL-001; INV-011c and scope confirmation still open). Fork: respec events, XP-not-refunded behaviour, and whether a retraining service market emerges. | |
| **HA-17** | Separate vs shared progression pools | A service profession levelled on **its own XP type**, not on mission XP (F-13, C-Q4). | ATTESTED (RECALLED) | Design: XP pool separation in the fork. | |
| **HA-18** | An elite that taxes the group | **F-26 — the most generalisable finding in the record:** an elite that gates XP on group content **while its group skills do nothing** becomes **a tax on the group, not a benefit to it**. Squad Commander is the worked instance: required a group to use group skills and to earn XP, the skills never worked, and the class was never repaired (F-26, Q-C1, R§6.4). | ATTESTED (RECALLED) | Design: any group-gated progression path in the fork whose group skills are unbuilt. **This row is checkable against the build today** — Phase 10 v0.2 §2.3 notes *"specials unbuilt in fork — ability-system residual, Phase 9 §5/§10."* | |

---

### Domain F — Reputation & social standing

| ID | Core loop | Pre-CU account | Strength | Fork observation | Verdict |
|---|---|---|---|---|---|
| **HA-19** | Reputation earned through visible work | Social standing in Pre-CU was carried by **rank, titles and recognised work** rather than by a numeric reputation system: an item's worth to a buyer came from *whose* work it was (**C-Q31** — what made an item you made worth having to someone else rather than just being another item); the top tier of the buff economy was a **social/guild privilege** rather than a purchasable service (**F-7**); and social structure was **tooling, ranks and titles** (**F-27**). | ATTESTED (RECALLED) | Design: TCIndustries' reputation surface. **Note:** the project has *deferred* provenance and formal reputation systems (project `AGENTS.md`, governance boundaries) — this row therefore observes an **authored absence**, not a gap to fill. | |
| **HA-20** | Mastery vs grind as the reward | *"Mastery. Grind was also rewarded, but mastery made the game much faster."* (Q219, R§6.3) — RECALLED opinion. | ATTESTED (RECALLED, opinion) | Design: horizontal-progression band. | |
| **HA-21** | Guild structure and whether it enforced anything | Guilds were **tooling, ranks and titles**; whether they enforced anything is **unknown**; a Community Advisor programme existed (F-27, C-Q47/C-Q48, R§7 item 14). | **ABSENT** on the enforcement question (explicitly unknown) | Mode C only. Fork: guild objects exist in the build, but enforcement behaviour is not attributable to them. | |

---

### Domain G — Trade, market & the player economy

| ID | Core loop | Pre-CU account | Strength | Fork observation | Verdict |
|---|---|---|---|---|---|
| **HA-22** | Vendors as granted, not crafted | Vendors were **granted by skill, not crafted**: Terminal at Business III (3,000 XP), Robot at Business IV (6,000 XP), a second vendor via Novice Merchant (5,000 Merchant XP, 6,000 credits), 5 vendors at Master Merchant, **4,000 items per vendor** (F-17, §5.2 OQ-016, C-Q14/C-Q16). | ATTESTED (RECALLED) | Design: **OQ-016 is RULED** (HD-RET-01 Option A — identity-anchored player vendors; INV-016a + INV-016b approved; numerics TBD under BAL-001). | |
| **HA-23** | Five selling routes, capped per route and per skill tier | Direct trade, Bazaar instant sale, Bazaar auction, player vendors, Merchant-only premium auctions. Caps: **25-item Bazaar limit**; per-item credit cap **6,000 (possibly 4,000)**; listing fee **20 → 100 credits** (F-18, C-Q17/C-Q18, Q18). | ATTESTED (RECALLED) | Fork: `anomalies` — `wash_trade_pair`, `balance_churn`, `placement_burst`, `price_outlier`, ranked, **no thresholds** (Phase 10 v0.2 §6 item 2, option (i)). | |
| **HA-24** | What the market felt like | Bazaar felt like a **real market or a dumping ground** (Q52); undercutting and price wars were common and how sellers felt about it (Q53); cartels were **rare and short-lived on the author's server** (Q54). | THIN (cat. R returned 0/7 substantive in pass 1; the above is pass-2 + Addendum A only) | Fork: `price_volatility`; `wealth.gini` as **diagnostic only, never target-capped** (Phase 10 v0.2 §5.4). Mode C for the felt-quality question. | |
| **HA-25** | Money sinks that actually operated | The sink classes that operated for an ordinary player: **housing maintenance, vehicle maintenance, player cities, insurance**. Wealthy-only: **city maintenance and specialisations** (one specialisation added 20,000 credits/week) (F-20, C-Q20/C-Q21). Vendor fees were negligible and bypassable; the operative sink was vendor **maintenance** (F-17). | ATTESTED (RECALLED) | Design: **OQ-009 is RULED** (HD-ECO-01 — Structure A, single unified currency; INV-009a approved; faucet-exceeds-sink mandate; numerics TBD). Fork: sink events if instrumented. | |
| **HA-26** | Exploit-driven inflation and its remediation | A duplication glitch caused inflation; remediation was **mass bans, confiscation from holders and middlemen, and a deliberately run economy-wide deficit**, with **75% of duped credits removed** (F-19, C-Q22/C-Q23). The account is developer-level and general, **not server-local** — the author's server was less affected (R§7 item 17). | ATTESTED (RECALLED, general) | Design: **OQ-012 remains TBD at the implementation layer**; the validation stance (in-scope, no account-level restriction) is settled per HD-EIC-03. Mode C: does exploit behaviour appear in a live cohort? | |
| **HA-27** | Multi-character rules | **One character per server per account**, with the second slot unlocked by an in-game trait; **no other rules** — a character-level unlock, not an account-structure rule (F-25, C-Q41/C-Q42). The author **never played more than one character**, so the record carries **no self-report evidence on multi-account behaviour at all** (R§5.2 OQ-012 row; R§7 item 12). | **ABSENT** on multi-account behaviour ("No" is an absence of evidence, not evidence of absence) | Fork: the fork's account/character model as built. | |

---

### Domain H — Society, civic structure & cities

| ID | Core loop | Pre-CU account | Strength | Fork observation | Verdict |
|---|---|---|---|---|---|
| **HA-28** | City formation, placement & dissolution | A **City Hall deed**, placed once and **never re-deeded**; civic structures placed by the mayor, citizens placing their own; a **minimum distance from other cities**; **rank grew with citizens** and gated civic structures (cantina rank 2, hospital rank 3, theatre rank 4); **unpaid maintenance damaged structures**; destruction of the City Hall **disbanded the whole city** — dissolution, not abandonment; the **mayor could withdraw from the treasury** (F-23, C-Q35/C-Q36, §5.2 OQ-007/BLD-004). | ATTESTED (RECALLED) | Design: **OQ-007 is RULED** (HD-BLD-01 Option A, free-claim formation; basic governance instruments only, **civic depth TBD**) and **BLD-004 abandonment is RULED** (HD-BLD-02/03, zero-credit removal). Fork: city objects, maintenance events, dissolution path. | |
| **HA-29** | City abandonment vs dissolution — the distinction | The record distinguishes these: dissolution (City Hall destroyed) from abandonment. Maintenance lapse was **670 hours** (F-34), superseding the earlier "about four weeks" estimate. | ATTESTED (RECALLED) | Design: **BLD-004 abandonment is RULED** (HD-BLD-02: at zero maintenance credit the structure is removed from the world; HD-BLD-03: contents removed with the structure). Placement / lots / density remain **TBD**. **⚠ Live DIVERGES candidate: the ruling chose *removal*; the historical account records *damage-at-lapse* plus *dissolution-on-City-Hall-destruction*.** §3.6. | |
| **HA-30** | Whether cities felt like living communities or ghost towns | Q153. | **ABSENT** (one-word answer; cat. M returned 0/8 substantive in pass 1) | Mode C only. This is the **thinnest civic row** and it is the one a live cohort would most directly speak to. | |

---

### Domain I — Death, wounds & recovery

| ID | Core loop | Pre-CU account | Strength | Fork observation | Verdict |
|---|---|---|---|---|---|
| **HA-31** | The four-part cost of death | Death cost **cloning credits, wounds, battle fatigue, and item decay** (5% uninsured, 1% insured) (F-24, C-Q39/C-Q40). | ATTESTED (RECALLED) | Design: the fork's death-cost model. Fork: death events and their components, if instrumented. | |
| **HA-32** | How PvP death differed | **Not documented.** *"I have no retrieved source describing how PvP death differed for ordinary characters."* (Q39, R§7 item 13) — unanswered by **either** instrument, against the thin first-pass PvP category. | **ABSENT** | Nothing to compare against. A fork observation here has no account comparand. | |

---

### Domain J — World, travel & the client

| ID | Core loop | Pre-CU account | Strength | Fork observation | Verdict |
|---|---|---|---|---|---|
| **HA-33** | World persistence & aliveness | Did the world feel alive — dynamic spawns, events (Q209); did the galaxy feel vast and explorable (Q192); did different planets feel distinct (Q194). | THIN (cat. V mixed 1/4 substantive; cats. S and W thin) | Mode C only. | |
| **HA-34** | The client and information load | **A negative result in the other direction:** *"Nothing was hard to use as everything was customizable"*; the author used **only the in-game macro system** (F-28, C-Q50/C-Q51, R§6.4). **Nothing was hard to use** — and the author used only in-game tooling, so third-party-tool reliance did not occur. | ATTESTED (RECALLED) | Mode C only. Note the double-sided reading: the client was *not* a problem, and no third-party tool dependence arose. | |
| **HA-35** | Macros and third-party tooling | The author used only in-game macros; no add-ons (C-Q52). The game **required third-party information** (a source the author named as a failure: *"there also needed to be an official source of game information instead of all of the third party sites"*, Q223). | ATTESTED (RECALLED) | Design/fork: does the fork require external information to play? | |

---

### Domain K — Combat & threat

| ID | Core loop | Pre-CU account | Strength | Fork observation | Verdict |
|---|---|---|---|---|---|
| **HA-36** | Combat as a resource-managed loop | The three pools (health/Action/mental) were distinct and meaningful (Q20); **HAM was the gate** for high-level combat (Q15); specials were **availability-capped** and **HAM-cost-capped** — ~**12 casts per 1v1** because *"players would run out of HAM"* (Q14/Q-C4); buffs were a meaningful combat lever (Q23). Buffed vs unbuffed, TTK collapses from 10–20 min to 1–5 min or a one-hit kill (F-5). | ATTESTED (RECALLED) | Fork: the combat TTK spot-check (proxy 2) and the ability/HAM system. **Note:** specials are unbuilt in the fork (Phase 10 v0.2 §2.3) — HAM-cost gating cannot be observed until they are built. | |
| **HA-37** | Incapacitation, wounds & the bleed-out loop | Incapacitation, revive, deathblows, the incap→revive→death loop in group vs solo combat (Q25–Q28); **deliberate 3-incap griefing** was part of the game (Q26). | THIN (cat. A is 22/28 substantive, but these specific items are mixed) | Fork: the heal-evidence gap (post-phase10 item 3) exercises this surface. | |

---

### Domain L — Exploit & integrity surface

| ID | Core loop | Pre-CU account | Strength | Fork observation | Verdict |
|---|---|---|---|---|---|
| **HA-38** | The 0-credit vendor fee exploit | A vendor was at once **free storage, an XP faucet and a sales channel** because the fee could be driven to zero (F-30, D-Q15). | ATTESTED (RECALLED) | Fork: does the fork's vendor/fee structure admit a zero-cost path? Mode C: do players find it? | |
| **HA-39** | Broken-systems pattern (general) | **F-35 — a skill that claims an effect it does not deliver can only be a belief.** Several aggro-removal skills existed and **none functioned** (F-35, D-Q29). Contrast F-32: vendor interactions were **documented as functional and were not**; *"the numbers never made sense or added up."* | ATTESTED (RECALLED) | Fork: documented-vs-actual behaviour audit of the fork's own mechanics — a **documentation-integrity** check, not a balance check. | |

---

## 5. GDD anti-goal cross-check rows

The Master GDD v1.1.1 §2 states what TCIndustries must avoid. Each anti-goal is a testable claim about
core-loop shape, and each has a Pre-CU comparand in the record. These rows are unusual in that the
**fork** side is usually already decided by the GDD — so the observation is narrow: does the *build* drift
from the stated anti-goal?

| ID | GDD anti-goal (§2) | Pre-CU comparand | Strength | Fork observation | Verdict |
|---|---|---|---|---|---|
| **HA-40** | Not an automation-first industrial game | SWG's factories were a meaningful endgame (Q66) or AFK (Q67) with **factory spam** a real problem (Q68). | THIN | Design: is there a factory/automation loop, and is it the dominant progression? | |
| **HA-41** | Not a conventional class-based MMO | Pre-CU had **no levels** — gear-driven (Q179). Mastery and skill trees, not character levels, were the progression. | ATTESTED (RECALLED) | Design: does the fork have character levels? | |
| **HA-42** | Not a quest-driven theme park | Mission terminals were a **"go do something" entry point** (Q103) but **missions dominated the economy as a faucet** and that felt inflationary (Q111). Mission payouts were the **specific weak spot of the record** (R§7 item 5) — the Merchant's income came from selling, so mission payouts were not salient. | THIN, and the mission-payout figures are the **weakest numbers in the file** (R§7 item 5) | Fork: `mission_outcomes` mix vs `interdependence` — is the mission faucet dominant? | |
| **HA-43** | Not a single-player crafting game with multiplayer features | Crafting was a **social** activity: what made an item worth having to someone else (C-Q31); craft commissions had **no trust mechanism** (F-22); making something for another player raised whose materials and how payment worked (C-Q32). | ATTESTED (RECALLED) | Fork: `interdependence` covered/total (proxy 1) with raw counts. | |
| **HA-44** | Not a pure market spreadsheet simulator | Bazaar limits, per-item credit caps, and per-route fees restricted ordinary traders; the author's own complaint was *"the Bazaar nonsense… limited what could be bought by normal people without travelling"* (Q223, R§6.1). Market access was **skill-graded per route** (F-18). | ATTESTED (RECALLED) | Fork: market access rules; `anomalies` `price_outlier` and `wash_trade_pair` lists. | |
| **HA-45** | Not a combat-only MMO | Service professions were weak earners and players migrated to combat professions for money (F-12); **separate XP pools** meant a thin service economy could not be rescued by mission XP (F-13). | ATTESTED (RECALLED) | Fork: `profession_distribution` — is the non-combat cohort populated? | |
| **HA-46** | Not a world where NPCs make meaningful player economic roles irrelevant | Vendors were **granted by skill** and were the lynchpin of the small-scale economy — but their fees were inert and bypassable, and the *operative* trade was direct negotiation (F-17, F-30). NPCs did not displace the player economy; they **served** it, and served it badly. | ATTESTED (RECALLED) | Design: NPC vendor role vs player vendor role in the fork. | |
| **HA-47** | Not a system where one character can perform every economically important role | **Separate XP pools per activity** (F-13) and separate service professions with distinct acts (F-12) meant one character could not cover every economic role. | ATTESTED (RECALLED) | Design/fork: can one character perform every economic role? | |

---

## 6. Coverage and confidence map

**This section is the honest self-description of the instrument.** It says where the Pre-CU side is strong
enough to support a verdict and where it is not, so that an executor does not treat a thin row as a
failure of the fork.

| Pre-CU account strength | Rows | What it means for verdicts |
|---|---|---|
| **ATTESTED** (substantive prose, usually second pass) | HA-01, HA-02, HA-05, HA-07, HA-08, HA-09, HA-10, HA-11, HA-12, HA-13, HA-15, HA-16, HA-17, HA-18, HA-19, HA-20, HA-22, HA-23, HA-25, HA-26, HA-28, HA-29, HA-31, HA-34, HA-35, HA-36, HA-38, HA-39, HA-41, HA-43, HA-44, HA-45, HA-46, HA-47 | A MATCH or DIVERGES here is a real signal. |
| **THIN** (one-word affirmatives, or pass-1 density only) | HA-03, HA-04, HA-06, HA-14, HA-24, HA-33, HA-37, HA-40, HA-42 | A DIVERGES here is weak evidence. A MATCH here is **near-meaningless** — per R§7 item 6, an affirmative one-word answer is weak evidence of satisfaction and near-zero evidence of mechanism. |
| **ABSENT** (explicitly not documented) | HA-21 (guild enforcement), HA-27 (multi-account behaviour), HA-30 (cities as communities), HA-32 (PvP death) | **INSUFFICIENT EVIDENCE is the expected verdict.** These rows exist to record the gap, not to be won. |

### 6.1 The thin/ABSENT cluster is not a coincidence

HA-21, HA-27, HA-30, HA-32 are all **social, civic or behavioural** — the exact surfaces Proxy 4 (Sandbox
Validation) would observe. Two consequences follow, and both are worth stating plainly:

1. **The two qualitative §1.3 tests overlap on their weakest ground.** Part 1 (Sandbox Validation) and
   Part 2 (Historical Accuracy) both turn on player-organised social behaviour — and the Pre-CU account
   for that behaviour is the thinnest part of the record. Proxy 4 was deferred for want of a live cohort.
   If a cohort is ever convened, the Historical Accuracy rows for social/civic behaviour will be the ones
   with the least to compare against.
2. **A future combined observation pass should be scheduled as one activity, not two.** Running Part 1 and
   Part 2 against the same cohort would at least mean the observation cost is paid once. That is a
   scheduling observation, not a decision, and it is the owner's to act on.

### 6.2 Governing limits on the whole instrument

Carried from the report and applicable to every row:

| Limit | Citation |
|---|---|
| One respondent, one server, one point in time | R§7 item 4, item 20 |
| No measured data anywhere in the Pre-CU account | R§7 item 5, item 19 |
| A denser record is not a more authoritative record | R§7 item 19 |
| 22-year interval; official forum records purged | R§9 closing note (Q-C8) |
| Confidence field is the respondent's own summary judgement, not a weighting mechanism | R§3.3.1 |
| The record supports no design decision, invariant, or tuning value | R§7 item 7 |
| The record is not a substitute for the primary source | R§7 item 8 |
| The second pass closed **no** project surface | R§7 item 11, R§5.2, R§9 — **note:** this is the report's claim about the *record's* effect on the surfaces. It is **not** a statement of live surface status, which was already partly ruled before the record was written (§3.6) |

---

## 7. Execution log template

One block per run. Paste below or duplicate per run.

```
=== HISTORICAL ACCURACY CHECKLIST RUN ===
Run ID:                    HA-RUN-____
Date / time:               ____
Executor (human):          ____
Mode(s) exercised:         A (design) / B (fork) / C (cohort)
Cohort (size, composition, duration):
Fork build / commit:       ____
Observation surfaces used: ____
Telemetry taken_at:        ____
----------------------------------------------------------------------
Rows attempted:            ____ / 47
Verdict tally:             MATCH ____ | DIVERGES ____
                            INSUFFICIENT EVIDENCE ____ | NOT OBSERVABLE HERE ____
                            NOT APPLICABLE ____
Rows left unfilled:        ____ (reason: ____________________)
----------------------------------------------------------------------
FINDINGS (DIVERGES only, one line each — no fixes proposed here):
  HA-__ : ________________________________________________________
  HA-__ : ________________________________________________________
----------------------------------------------------------------------
QUESTIONS FOR THE OWNER (ruling required, not a recommendation):
  1.  ____________________________________________________________
----------------------------------------------------------------------
NOTE: This run produced EVIDENCE. It closed no open question, changed
no status, authorised no tuning, and set no target.
```

**Do not** add a composite score, a percentage-complete figure presented as a verdict, or any numeric
summary of agreement. The tally above is a **count of recorded outcomes**, not a score, and it must not be
converted into one. Phase 10 v0.2 §5.3 forbids a fabricated score; a composite would be a fabricated score
in a different costume.

---

## 8. Sign-off — what a completed checklist does and does not constitute

### 8.1 What it does

- Produce a **dated, cited, per-row comparison** of the fork's core loops against attested Pre-CU accounts.
- Produce a **list of specific divergences**, each with a row ID, a Pre-CU citation, and a fork observation
  — i.e. actionable telemetry in the §1.3 sense.
- Produce an **explicit map of what could not be decided and why**, which is itself a finding: it tells the
  owner which surfaces the evidence base cannot currently speak to.

### 8.2 What it does not

- **It does not close any OQ, TBD, or open question, and it does not re-open a ruled one.** Only a recorded
  human ruling does that, in `governance/TCIndustries_Human_Rulings_Register.md` (GR-006.1). A divergence
  found against an already-ruled surface is a **finding for the owner**, filed for their decision, not a
  status change made by this document.
- **It does not change any status.** A DIVERGES verdict is not a design decision. It is a finding.
- **It does not authorise a fix.** BAL-001 and HD-TST-01 are untouched. A divergence found by this
  checklist becomes a *new proposal*, separately authorised.
- **It does not make the fork historically accurate**, and a completed checklist is not a §1.3 PASS by
  itself. §1.3's exit criterion is disjunctive; the checklist is one input to it, and only a human records
  the test's outcome.
- **It does not speak for the owner.** A checklist filled in by an agent is not a valid run. §2's rules
  make the verdict a human's act; an agent may stage observations and log them, but may not write a verdict.

---

## 9. Open items requiring an owner ruling

These are recorded, not decided. Each is a question the owner must answer; none has an agent-recommended
answer, and this document proceeds on a stated default for each so it is usable if the owner defers.

| # | Item | Default if unanswered | Why it is not the author's to decide |
|---|---|---|---|
| 1 | **Adopt this checklist** as the §1.3 Historical Accuracy instrument, or specify a different shape. | Not adopted; Part 2 stays unrun. | Instrument adoption is a human act. The 2026-10-18 item-8 split decision authorised *authoring* a formal instrument; this file is a candidate continuation of that authorisation, and its own adoption is a further act. |
| 2 | **§3.1** — does the `4–6%` stall-commission band have independent documentary support, and is a percentage band the right comparand for a flat-credit, zero-bypassable fee? | Band stands as written, validate-against-only. | The band and the record are both filed material; reconciling them is a ruling. |
| 3 | **§3.2** — does the `30–60 s PvP 1v1` band stay, given the filed record re-scopes that figure to group PvP? | Band stands as written. | Same. Direction matters: the report is EVIDENCE and the historical record outranks it. |
| 4 | **§3.3** — are the three unaddressed progression bands to be recorded as having **no Pre-CU comparand**, so their validations are logged as vacuous? | Log as observed; no band change. | Same. |
| 5 | **Which mode to run first** — A (design, executable now), B (fork, executable now), or hold for C (cohort)? | Mode A only, as a documentation pass. | Mode B's results on a ~3-character cohort may be cohort artefacts; recommending a mode is an owner judgement about what the evidence can bear. |
| 6 | **Whether to co-schedule Part 1 and Part 2** against one cohort when a cohort exists (§6.1). | No scheduling change. | Scheduling is the owner's. |
| 7 | **Whether the row set should be extended** — 47 rows were derived from the report's §5.1–§5.6, §5.2, §6 and the GDD anti-goals. The report's 24 instrument categories were not walked exhaustively. | Row set as drafted. | Scope is a judgement call the owner can make against the checklist's purpose. |
| 8 | **§3.6 — the report's §5.2 "Status in the project" column is stale** (six of eight surfaces already ruled when the report was authored). Correct the report by a dated v1.7 amendment, or leave it and rely on the register as authority. | Leave the report unamended; the register governs. | GR-006.8: reporting a gap is never authorisation to correct it unilaterally, and the report is a filed, versioned document carrying a human author. The correct remedy, if any, is the owner's. |

---

## 10. Related documents

| Document | Relationship |
|---|---|
| `investigations/TCIndustries_SWG_Historical_Recollection_Report_Tiwa_Pene_2026-10-21.md` (v1.6) | **The Pre-CU account source for every row.** EVIDENCE + ANALYSIS. Not authority. |
| `investigations/tiwas-swg-recollections-21-10-2026.md` | Raw first-pass answer record (unaltered primary artefact). |
| `investigations/tiwas-swg-recollections-22-10-2026.md` | Raw second-pass answer record (unaltered primary artefact). |
| `proposals/SWG_Historical_Accuracy-v1.md` (v1.0) | The 223-question instrument; the Pre-CU corpus was gathered through it. |
| `proposals/SWG_Historical_Accuracy_v2_Targeted_ReAsk.md` (v1.0) | The 56-question second instrument; the 8 named-open surfaces. |
| `proposals/SWG_Phase10_Balance_Telemetry_Polish_Proposal_v0.2.md` | §§1.2, 2.2, 2.3, 5.2, 5.3, 6.4, 9 — the obligations this instrument serves. |
| `proposals/SWG_Post_Phase10_Program_Proposal_v0.1.md` | §9 item 8 and the 2026-10-18 split decision; §9 item 9. |
| `governance/TCIndustries_Authority_Provenance_Reconciliation_Matrix.md` | GR-001, GR-002, GR-003, GR-007; GR-006.1 (where a ruling would have to be filed). |
| `governance/TCIndustries_Testbed_Reuse_Human_Ruling_HD-TST-01_2026-09-14.md` | Patterns-only fork; fork behaviour is PROTOTYPE/EVIDENCE, never design authority. |
| `canonical/gdd/TCIndustries_Master_GDD_v1.1.1_Status_Patch.md` | §2 central loop and anti-goals; §33/§34 band and KPI provenance. |
| `sources/swg-pre-cu/SWG_PreCU_GDD.md` | HISTORICAL reference text — the other Pre-CU source, and the higher authority where the two conflict (R§7 item 8). |

---

## 11. Version history

| Version | Date | Change |
|---|---|---|
| v0.1 | 2026-10-22 | Initial draft. 47 rows across 12 core-loop domains (§4) plus 8 GDD anti-goal cross-check rows (§5). Pre-checklist notes on three Phase 10 §2.3 bands (§3). Verdict vocabulary, evidence classes, three observation modes, execution log template, owner-ruling list. Not yet assessed; not adopted; not run. |
| v0.2 | 2026-10-22 | **Status corrections.** The v0.1 draft inherited the report's §5.2 *"Status in the project"* column, which is **stale** — six of the eight named surfaces (OQ-007, BLD-004, OQ-009, OQ-010, OQ-011, OQ-016) were already ruled before the report was authored. Added **§3.6** carrying the corrected live-status table and the post-hoc corroboration/divergence analysis; corrected the fork-observation cells of HA-09, HA-10, HA-16, HA-22, HA-25, HA-26, HA-28, HA-29; amended §0.4, §2 rule 6, §6.2 and §8.2 for ruled surfaces; added §9 item 8. **No row was added, removed or re-worded**; the 47-row set and the verdict vocabulary are unchanged. The report itself was **not** edited (GR-006.8) — the defect is filed for the owner. |

---

*PROPOSED candidate material. Filed in `proposals/` as unadjudicated candidate procedure. Nothing in this
document carries authority unless a documented human ruling promotes it, and that promotion must be
recorded, never silent. Executing it produces EVIDENCE, not decisions.*
