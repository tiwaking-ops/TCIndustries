# Project Documentation Report — OpenCode Session: Historical Accuracy Coverage & Checklist Instrument

| Field | Value |
|---|---|
| **Document ID** | TCIND-PDR-2026-10-22-01 |
| **Title** | Complete Results of the 2026-10-22 OpenCode Session — Part 2 Coverage Assessment, the SWG Historical Accuracy Core-Loop Comparison Checklist v0.2, and a Dependency-Consistency Finding Against the Recollection Record |
| **Author** | MiniMax M3.1 Flash (free) — model `openrouter/minimax/minimax-m3.1-flash:free`, via OpenCode, 2026-10-22 |
| **Assessor / Recipient** | Project owner (Tiwa Pene) — for review. **Independent LLM assessment: not yet performed.** No document in this repository records the same entity as both author and assessor of the same work, and this one does not. |
| **Date** | 2026-10-22 |
| **Status** | EVIDENCE (observation report per `AGENTS.md` status table; no design authority claimed, no status promotion performed) |
| **Distribution** | Project owner, OpenCode |
| **Authority basis** | Owner instruction of 2026-10-22: *"Write a formal project documentation report for OpenCode covering the complete results of this chat."* No design, tuning, or status change is authorised by this instruction. |
| **Related documents** | `proposals/SWG_Historical_Accuracy_Checklist_v0.2.md` (delivered, PROPOSED) · `proposals/SWG_Historical_Accuracy_Checklist.html` (delivered, local fill-and-export form) · `proposals/SWG_Historical_Accuracy_Checklist_build.ps1` (delivered, build script) · `investigations/TCIndustries_SWG_Historical_Recollection_Report_Tiwa_Pene_2026-10-21.md` (v1.6, EVIDENCE) · `proposals/SWG_Historical_Accuracy-v1.md` · `proposals/SWG_Historical_Accuracy_v2_Targeted_ReAsk.md` · `proposals/SWG_Phase10_Balance_Telemetry_Polish_Proposal_v0.2.md` · `proposals/SWG_Post_Phase10_Program_Proposal_v0.1.md` · `governance/TCIndustries_Human_Rulings_Register.md` · `proposals/TCIndustries_Open_Questions_Register.md` · `canonical/gdd/TCIndustries_Master_GDD_v1.1.1_Status_Patch.md` |

**Provenance note (GR-007.1 / GR-007.3).** The author identity named above is the same identity that produced
the **assessment half** of `investigations/TCIndustries_SWG_Historical_Recollection_Report_Tiwa_Pene_2026-10-21.md`
(v1.6). This report therefore both *cites* that record and *finds a defect* in one of its status
descriptions (§6.1). The defect is not self-assessed in the sense that would compromise it: it is a
mechanical mismatch between two filed documents, verified by direct comparison against the Human Rulings
Register and the Open Questions Register, and it is reproducible by any reader. The underlying record is
not edited (§6.2). Independent assessment of this report, and of the delivered checklist, is recommended
from a different model.

---

## 1. Executive summary

The session was opened with a request to explain what the Phase 10 harness's "artifact review" actually
entails, and closed with a request to build a playtesting form for the Historical Accuracy checklist.
Four work packages resulted, all documentary:

1. **Clarification of the two qualitative §1.3 tests.** The harness's Proxy 4 (Sandbox Validation) stages
   five telemetry surfaces and reports NOT-MEASURABLE; the fifth §1.3 test (Historical Accuracy) is
   human-pass-only and is **not asserted in the harness at all**. Both are qualitative by definition, so
   neither can be scored by an agent.
2. **An assessment of the 2026-10-21 recollection report against Part 2.** Conclusion: the report supplies
   the Pre-CU side of the comparison far better than the only prior source, and it supplies two of the
   three structural elements of a checklist — but it is a **reference corpus, not a checklist**. It also
   surfaced three live problems in the Phase 10 §2.3 HISTORICAL bands.
3. **A delivered instrument:** `proposals/SWG_Historical_Accuracy_Checklist_v0.2.md` — 47 comparison rows
   across 12 core-loop domains plus 8 GDD anti-goal cross-check rows, with a five-value verdict
   vocabulary, evidence classes, three observation modes, an execution log template, and an owner-ruling
   list. Filed PROPOSED, unadopted, unrun, unassessed, and uncommitted.
4. **A delivered playtesting form:** `proposals/SWG_Historical_Accuracy_Checklist.html` — a local
   fill-and-export page for the 47 rows, **generated from** the instrument by
   `proposals/SWG_Historical_Accuracy_Checklist_build.ps1` and verified against it character-for-character
   (§5.4, §8).

**The session's most consequential output is a finding, not the instrument.** Cross-checking the
recollection report's §5.2 *"Status in the project"* column against the Human Rulings Register showed that
**six of the eight surfaces the report lists as TBD had already been ruled** before the report was written
(§6.1). The report's closing assurance that those surfaces *"remain exactly as open"* is therefore false
against the authority record. This is a dependency-consistency defect (GR-006.2) in a filed document.

**The defect is reported, not corrected** (GR-006.8: reporting a gap is never authorisation to correct it
unilaterally). The recollection report was not edited. The delivered checklist was corrected, because it
is this session's own unfiled work and it had inherited the stale statuses; that correction is recorded as
a version bump to v0.2 with a full change note, not a silent edit.

**No design decision, status change, tuning value, commit, or edit to any pre-existing project document
was made by this session.**

---

## 2. Scope of the documented session

The session covered three owner instructions, in order:

1. **Explain the "artifact review."** The owner supplied a summary of the Proxy 4 surfaces, the
   "no agent action possible" reasoning, and the two-part structure (Item 8 observation-log review;
   Historical Accuracy checklist), and asked for verification.
2. **Assess whether a named file covers Part 2.** The owner asked whether
   `investigations/TCIndustries_SWG_Historical_Recollection_Report_Tiwa_Pene_2026-10-21.md` helps cover
   the Historical Accuracy checklist "during playtesting / observation", and quoted the Part 2 definition.
3. **Draft the checklist.** The owner answered *"yes draft"* to a proposed instrument, and the instrument
   was produced, verified, corrected, and filed.
4. **Build a playtesting form.** The owner asked for *"an html page for the swg historical accuracy
   checklist for playtesting similar to `proposals/SWG_Historical_Answers.html`"*, and the form was
   built, generated from the instrument, verified against it, and exercised in a live browser (§5.4).

Out of scope, and deliberately untouched: any numeric tuning; any change to a ruling's status; any edit to
the Master GDD, the Human Rulings Register, the Open Questions Register, the Invariant Register, or the
recollection report; any commit; any code change; any execution of the checklist.

---

## 3. What the "artifact review" actually entails — verified

The owner's opening summary was **checked against the source and confirmed**, with one added code-quality
observation (§3.4).

### 3.1 Proxy 4 — what the harness does

`testbed/swg-phase3-combat/swg-server/cmd/phase10test/main.go`, lines 633–656, implements *Test 4: Sandbox
Validation (qualitative — telemetry artifacts)*. The harness:

- **Stages five telemetry surfaces** from the run's activity: ranked anomaly candidates; mission outcomes
  (available / accepted / completed / expired); wealth distribution + Gini (diagnostic only, never
  target-capped); price volatility (from `market_records`); profession distribution (registry-joined).
- **Verifies the surfaces are populated** and, if so, reports **NM (NOT-MEASURABLE)** with the note that
  artifacts are *"ready for human review."* If the REST payloads are empty it reports **FAIL**.
- **Renders no qualitative verdict.** The harness can surface what a human reads; it cannot judge it.

### 3.2 The two obligations that remain with the owner

| Obligation | Nature | Where it lives |
|---|---|---|
| **Item 8 — Proxy 4 observation-log review** | Owner reviews the staged artifacts and annotates an **observation log** recording player-organized events, guilds and schemes *not authored by design*, keyed by **event, date, witnesses** | `SWG_Phase10_..._v0.2.md` §5.3; §9 item 5 |
| **Historical Accuracy checklist** | Owner runs a human checklist comparing core loops to documented Pre-CU accounts | §2.2, §5.3, §6.4 of the same proposal |

The **Historical Accuracy** limb is *explicitly not asserted in the harness*: lines 702–703 record
*"Historical Accuracy (5th §1.3 test) is HUMAN pass only — out of testbed scope by proposal §6.4. Not
asserted here."*

### 3.3 Why no agent action is possible

Both tests are qualitative by definition. BAL-001 (LOCKED), HD-TST-01, and §5.3 (*"Qualitative tests stay
qualitative"*) forbid fabricating a score or auto-resolving them. An agent can stage telemetry; it cannot
judge emergent player behaviour or historical fidelity of core loops. The recorded status is therefore
**owner review; no agent action** — with the default if never answered being **Open (stays open)**.

**Confirmed against the live record.** `proposals/SWG_Post_Phase10_Program_Proposal_v0.1.md` §9.1 item 8
records a **split decision of 2026-10-18**: *Proxy 4 — DEFERRED* (emergent player-organized play requires
players; no live cohort; fabricating one would violate the no-simulated-player-fabrication rule at §5.2);
*Historical Accuracy — human-approved formal proposal authored* as `proposals/SWG_Historical_Accuracy-v1.md`.
§10 confirms both limbs carry no code work, and that Option A (Consolidate) does not touch them.

**One clarification the owner's framing did not carry:** the item-8 split decision means the **Historical
Accuracy limb is no longer merely "pending"** — the instrument has been authored and answered, and the
remaining gap is the *comparison procedure*, not the Pre-CU account. That distinction is the reason the
delivered instrument exists.

### 3.4 Added observation — a vestigial condition in the harness (minor, non-blocking)

Line 647 reads `if len(aList) >= 0 && wealth["players"] != nil`. In Go, `len()` of a slice is never
negative, so `len(aList) >= 0` is **tautologically true** and the effective gate is
`wealth["players"] != nil` alone. This is a code-clarity observation, not a defect in behaviour: the
intended check (artifact list non-empty *and* wealth populated) is not what the code tests, since an empty
anomaly list still reports NM. Flagged for the owner; **no code change was made by this session**, and
the fork is under a zero-numeric-change discipline in any case.

---

## 4. Does the recollection report cover Part 2?

**Assessment: it covers the *source* side of Part 2 substantially, and two of the three structural
elements of a checklist. It does not cover Part 2 as a procedure, and it does not close Part 2.**

### 4.1 What the report supplies

| Part 2 need | Supplied by | Verdict |
|---|---|---|
| Pre-CU account per core loop | §5.1 combat/HAM · §5.2 progression/skills · §5.3 economy/market · §5.4 crafting/resources/world · §5.5 professions/service economy · §5.6 world/travel/client | **Yes** — organised on the GDD's central loop |
| Checklist row set | Report §5.2 — a table of 8 named-open surfaces × *what the record attests* × *findings* × *closed?* | **Yes, pre-built** |
| "Do not reproduce" column | Report §6.1 / §6.1a / §6.4 — mechanism-level failures, not complaints. F-26 generalises past its origin: *an elite that gates XP on group content while its group skills do nothing becomes a tax on the group* | **Yes** |
| The comparison procedure itself | — | **No** — no instrument existed on disk |
| The fork's own behaviour | — | **No** — the report is entirely the Pre-CU side |

Before this report, the only Pre-CU source was `sources/swg-pre-cu/SWG_PreCU_GDD.md` — a *design document*
describing what Pre-CU SWG was specified to do, not how its loops played. The report is the first record
of the latter.

### 4.2 What the report does not carry

- **It is EVIDENCE, not authority.** Report §7 item 7: it supports no design decision, invariant, or
  tuning value. §9: GR-001/002/003 apply; the findings are *"a report to the owner, not authorisation to
  alter the source record, the instrument, or any other document."*
- **Its coverage is uneven exactly where observation is richest.** Report §7 item 3: the v1.0 record
  carries no usable evidence on social structure, guilds, player cities, civic governance, housing,
  death/cloning, market mechanics, travel, or NPC/AI beyond a one-word affirmative. Addendum C repaired
  much of the *mechanism* gap, but F-27 leaves guild enforcement *"unknown"*, §7 items 13–14 leave
  PvP-death and guild enforcement unanswered, and §7 item 15 removes creature thinning/migration.
- **Governing limits:** one respondent, one server, 22 years, no measured data anywhere (report §7 items
  4, 5, 19, 20; §9 closing note).

### 4.3 Three live problems in the Phase 10 §2.3 HISTORICAL bands

Surfaced by the report and carried into the delivered instrument as §3.1–§3.3. These are **questions for
the owner, not corrections** — the bands are validate-against-only, and per report §7 item 8 the
historical record outranks the recollection where they conflict.

| # | Band (§2.3) | Finding |
|---|---|---|
| 1 | `stall commission 4–6%` | **Uncorroborated and possibly the wrong shape.** Report F-2 holds it as the one figure of three the record does not support (the author never paid listing fees, Q191/Q89). Separately, the attested Pre-CU fee was a **flat 20→100 credit** charge, and F-30 records it was **drivable to zero** — so a percentage band may not be the right comparand. Currently inert in the fork (*"stalls dormant"*), so nothing is compared today |
| 2 | `TTK 30–60 s PvP 1v1` | **Re-scoped.** Report F-5 (closed v1.1) establishes 30–60 s is **group PvP / focus fire**; 1v1 is **~10 minutes** excluding healing. The band carries a figure the record re-scopes |
| 3 | Progression bands | **One of four corroborated, three unaddressed.** `150–250 h` is corroborated at **240 h** (F-2a, unassisted-from-scratch path; ~24 h assisted). The `40–60 h` basic-mastery, `80–120 h` elite and `200–300 h` full-cap bands have **no Pre-CU comparand**, so validating a projection against them is vacuous and should be logged as such |

A fourth, minor caution is recorded at instrument §3.4: the §2.3 HAM band (3000–6000 unbuffed,
6000–12000 buffed) and the record's HAM figures (1,000 minimum per buff; 10,000–12,000 from top-tier
doctors) are **different quantities** and must not be compared directly — the report's 1,000 is a
per-buff contribution, not a pool floor. Recording this prevents a false DIVERGES.

---

## 5. Delivered artifact — the comparison checklist

**File:** `proposals/SWG_Historical_Accuracy_Checklist_v0.2.md` (509 lines, PROPOSED candidate material).

### 5.1 Structure

| Element | Content |
|---|---|
| **47 rows** | 39 across 12 core-loop domains (§4) + 8 GDD anti-goal cross-check rows (§5) — A resources · B acquisition · C crafting · D services · E specialisation · F reputation · G trade/market · H civic · I death · J world/client · K combat · L exploit surface |
| **Row shape** | `ID \| core loop \| Pre-CU account (cited to Q-number) \| account strength \| fork observation \| verdict` |
| **Verdict vocabulary** | MATCH · DIVERGES · **INSUFFICIENT EVIDENCE** (≠ MATCH) · **NOT OBSERVABLE HERE** (≠ INSUFFICIENT EVIDENCE) · NOT APPLICABLE |
| **Evidence classes** | RECALLED · DOCUMENTED · MEASURED (fork side only) · INFERRED |
| **Observation modes** | **A** design · **B** fork · **C** live cohort — with the explicit statement that A and B do not clear C |
| **Account strength** | ATTESTED (34) · THIN (9) · ABSENT (4), with per-strength guidance on what a verdict is worth |
| **Execution log** | Fill-in block; outcome counts only, **no composite score** |
| **Owner-ruling list** | 8 items (§9), each with a stated default |

### 5.2 Design decisions of substance

- **INSUFFICIENT EVIDENCE is not a soft MATCH.** Per report §7 item 6, an affirmative one-word answer is
  weak evidence of satisfaction and near-zero evidence of mechanism. For the 4 ABSENT rows, INSUFFICIENT
  EVIDENCE is the *expected* verdict; the rows exist to record the gap, not to be won.
- **NOT OBSERVABLE HERE is kept separate** from INSUFFICIENT EVIDENCE, so a later reader can distinguish a
  *thin account* from a *wrong question* — report F-8 established that answer density tracks instrument
  fidelity to Pre-CU, not respondent expertise.
- **Mode C is where Part 2 actually lives**, and it is the mode currently unavailable: Proxy 4 was
  deferred for the same reason — no live cohort.
- **No composite score, ever.** §5.3 of the Phase 10 proposal forbids a fabricated score; the instrument
  states the tally "must not be converted into one."

### 5.3 The §6.1 concentration finding (in the instrument)

The THIN/ABSENT cluster (HA-21 guild enforcement, HA-27 multi-account behaviour, HA-30 cities as
communities, HA-32 PvP death) is **entirely social, civic and behavioural** — the same surfaces Proxy 4
would observe. Both qualitative §1.3 tests therefore land on the ground where the Pre-CU account is
thinnest, and Proxy 4 is deferred for want of a cohort. The instrument records the observation that a
future combined observation pass should be **scheduled as one activity, not two**. That is a scheduling
observation for the owner, not a decision.

### 5.4 The playtesting form — `proposals/SWG_Historical_Accuracy_Checklist.html`

A local, single-file, fill-and-export page for the 47 rows, in the same category as the precedent forms
`proposals/SWG_Historical_Answers.html` and `..._v2.html`. **It carries no project authority of any kind.**
It is a transport mechanism, not a record and not a ruling.

**Build method — derived, not transcribed.** The form is **generated from** the instrument by
`proposals/SWG_Historical_Accuracy_Checklist_build.ps1`, and the build verifies its own output. This is
the direct answer to the D-3a failure class (rendered question text silently diverging from source):
the form's reference text cannot drift from the instrument, because it is not separately authored. The
build script is filed so the form is reproducible and auditable, and the run is byte-reproducible —
the same SHA-256 (`FA883959…`) when executed from a different working directory.

**Function set**

| Feature | Behaviour |
|---|---|
| Run header | 9 fields: Run ID, Date, Executor, Mode(s), Cohort, Fork build/commit, Observation surfaces, Telemetry `taken_at`, Reason for unfilled rows |
| Per row | Verdict (5-value dropdown), Evidence class, Mode checkboxes (A/B/C), Observation textarea, Observation-log ref, Notes |
| Read-only per row | Pre-CU account, account strength, and the fork observation point — all rendered from the instrument |
| Badges | `ATTESTED` / `THIN` / `ABSENT` per row; `RULED SURFACE` on the 7 rows touching a ruled surface; `KNOWN DIVERGENCE CANDIDATE` on HA-29 |
| Filters | Domain (14 options), Unfilled only, DIVERGES only, Ruled surfaces only, Thin/absent only |
| Export | Markdown run record matching the instrument's §7 template, plus a 47-entry per-row record |
| Autosave | `localStorage`; survives reload |
| Network | **Zero external resources.** No `src`/`href` to any remote origin |

**Governance guards built into the tool**

- **No verdict is pre-filled.** All 47 selects ship on *(not yet recorded)*. The tool cannot compute a
  verdict and offers no button that would.
- **Unfilled rows export as `UNFILLED`**, never as a match. Verified in the browser: every per-row
  verdict value is one of the five vocabulary terms or `UNFILLED`, and the tally reconciles exactly
  against the per-row entries.
- **The tally is labelled as not-a-score** in the live UI *and* in the exported record
  (*"This is a count of recorded outcomes, NOT a score. It must not be converted into one, and no
  composite figure may be derived from it."*).
- **Unrecorded metadata is stamped `UNRECORDED`**, never a plausible-looking placeholder — the D-2
  failure class, where a hardcoded date and a literal `<handle>` placeholder shipped in every exported
  record.
- **A DIVERGES on a ruled surface is framed as a re-opening question for the owner**, in the export
  footer, not as a status change.
- The export footer restates: *"This run produced EVIDENCE. It closed no open question, changed no
  status, authorised no tuning, and set no target."*

### 5.5 Defect found and fixed during the build — an encoding fault (F-4, §6.4)

The build script is a PowerShell file, and PowerShell 5.1 decodes a **BOM-less** `.ps1` using the ANSI
codepage. The script's own literal em dashes were therefore mis-decoded and written into the generated
HTML as mojibake (`â€”`, U+00E2 U+20AC U+201D) — **10 lines**, including the page `<title>`. This was
caught by verifying the *rendered* title in the browser, not by static reading, because the mojibake is
valid UTF-8 and therefore invisible to a file-level check.

Two fixes were applied and are now permanent: the build script is filed **with a UTF-8 BOM** (and
carries a header warning stating the requirement), and the build now **fails if any mojibake marker
reaches the output**. A second instance of the same class was found in the *verification* pass, where
`Get-Content` mis-decoded the instrument's `§` — fixed by reading the source explicitly as UTF-8.
**Both the form and the build now report 0 mojibake markers and 0 round-trip mismatches.**

---

## 6. Findings for the owner

### 6.1 F-1 (material) — the recollection report's §5.2 status column is stale

The report's §5.2 table carries a *"Status in the project"* column, and its closing assurances are that
*"none was closed"* and that the eight surfaces *"remain exactly as open"* (§7 item 11, §9). **That is
false against the authority record.** Six of the eight were already ruled when the report was authored on
2026-10-21:

| Surface | Report says | Live status | Ruling | Filed |
|---|---|---|---|---|
| OQ-007 — city formation & governance | "TBD" | **RULED** — Option A, free-claim formation; basic governance instruments; **civic depth TBD** | HD-BLD-01 | 2026-09-19 |
| BLD-004 — abandonment handling | "TBD" | **RULED** — zero-credit removal; contents removed with structure (HD-BLD-03). **Placement/lots/density TBD** | HD-BLD-02 + HD-BLD-03 | 2026-09-19 / 09-20 |
| OQ-009 — currency & money sinks | "TBD" | **RULED** — Structure A single unified currency; INV-009a approved; **faucet-exceeds-sink mandate** | HD-ECO-01 | 2026-09-28 |
| OQ-010 — durability/decay/repair | "TBD; models A/B/C" | **RULED** — Option C, **no baseline decay**; numerics TBD | HD-ITM-01 + HD-ITM-02 | 2026-09-19 / 09-20 |
| OQ-011 — respecialisation cost | "TBD; cost classes A/B/C" | **RULED** — **free-drop** cost; trainer-gated acquisition; magnitudes TBD (BAL-001) | HD-PROF-01 + HD-PROF-02 + HD-PROF-03 | 2026-09-19 / 09-20 / 09-28 |
| OQ-016 — vendor & retail structure | "TBD; models A/B/C" | **RULED** — Option A, identity-anchored player vendors; INV-016a/b approved; numerics TBD | HD-RET-01 | 2026-09-19 |
| SERV-001 / SERV-002 | SERV-001 LOCKED (principle only); SERV-002 PROPOSED | **Correct as stated** | — | — |
| OQ-012 / SAFE-002 | "TBD" | **Correct in substance** — TBD at the implementation layer; validation stance settled per HD-EIC-03 | HD-EIC-03 | 2026-08-25 |

**Verified against three independent authorities**, all agreeing: the Human Rulings Register (Sections C
and D), the Open Questions Register v1.2 (which carries a `change_note` recording the 2026-09-28 pointer
edits), and the Master GDD v1.1.1 Status Patch §"Applied 2026-09-28" status-line table.

**Probable cause (not asserted as fact).** The report's authority basis cites the Authority &
Provenance Reconciliation Matrix but **not** the Human Rulings Register. That is consistent with a
status column assembled from the GDD's original §31 TBD list rather than from the register. The register
is the authority; the GDD status-patch table is a *pointer* to it.

**Classification:** a dependency-consistency defect (GR-006.2) and, materially, a **false assurance**
inside a filed evidence document. The record's *own findings* are unaffected — F-1 to F-35, the addenda,
the answer records and the raw artefacts are untouched by this.

### 6.2 F-1 disposition — reported, not corrected

Per **GR-006.8** (*"Reporting a gap is never authorisation to correct it unilaterally"*), the recollection
report was **not edited** by this session. It is a filed, versioned document with a human author
(Tiwa Pene) and a multi-version assessment history that explicitly forbids silent amendment. The
instrument's §9 item 8 records the owner-ruling options: (a) issue a dated **v1.7 amendment** carrying an
explicit *"Status column corrected"* note, or (b) leave the record unamended and rely on the register as
authority. The default recorded is **(b)**.

The **delivered checklist was corrected**, because it is this session's own unfiled work and it had
inherited the stale statuses. The correction is a version bump to **v0.2** with a full change note
(§11 of the instrument): eight fork-observation cells corrected, §3.6 added, §0.4 / §2 rule 6 / §6.2 /
§8.2 amended, §9 item 8 added. **No row was added, removed or re-worded**; the 47-row set and the
verdict vocabulary are unchanged.

### 6.3 F-2 (material, for the owner) — post-hoc corroboration and one divergence

Because the rulings predate the record, the report's attestations were **not available when the rulings
were made**. They are now available, and they point in both directions:

| Surface | Ruling | What the record attests | Relationship |
|---|---|---|---|
| OQ-010 | No baseline decay (Option C) | Decay existed in the manual and the code **and did nothing in play**; there was **no repair economy** (F-15, F-16). The report states the record's effective historical answer was also **C** | **CORROBORATES** |
| OQ-011 | Free-drop respec | Respec was **free, unlimited, no cooldown, no cost per removal**; skill points returned, **XP not refunded** (F-14, C-Q10–C-Q12) | **CORROBORATES**, and adds the mechanism the ruling did not record — XP is not refunded, skill points are |
| OQ-016 | Player-vendor primacy (Option A) | Vendors were **granted by skill, not crafted**, bore their owner's identity; NPC commerce did not displace the player economy (F-17, F-30) | **CORROBORATES** INV-016b. Also attests the Pre-CU fee was **flat credits, not a percentage** — bearing on the still-deferred numerics |
| OQ-009 | Faucet-exceeds-sink mandate | Mission payouts dominated income; the operative sinks were housing maintenance, vehicle maintenance, player cities and insurance; vendor fees were **inert and bypassable** (F-17, F-19, F-20) | **CORROBORATES** the mandate's shape. Supplies two sink classes the record says are *new to it* and **not thereby adopted** |
| BLD-004 | **Removal** at zero maintenance credit | Unpaid maintenance **damaged** structures; at 670 h the structure was affected; destruction of the City Hall **disbanded the whole city** (F-23, F-34) | **DIVERGES on mechanism** — the ruling chose *removal*; the historical account records *damage-at-lapse* plus *dissolution-on-City-Hall-destruction* |

**This puts no ruling in question.** A Pre-CU recollection is not authority, and the project `AGENTS.md` is
explicit that historical inspiration is never automatic canon. The divergence is the deliverable; the
decision is the owner's. The delivered instrument encodes this: a DIVERGES on a ruled surface is recorded
as a **re-opening question for the owner**, never as a status change made by the document (instrument §2
rule 6, §8.2).

### 6.4 F-3 (minor, informational) — vestigial harness condition

See §3.4 above. No code change made; fork code untouched.

### 6.5 F-4 (minor, self-contained) — encoding fault in the form build, found and fixed

See §5.5 above. Summary: PowerShell 5.1's ANSI decoding of a BOM-less `.ps1` introduced mojibake into
10 lines of the generated form, including the page title. **Found by reading the rendered title in a
live browser** — a file-level check passes it, because the mojibake is valid UTF-8. Fixed by filing the
build script with a UTF-8 BOM and adding a build-time guard that fails on any mojibake marker.

**Bearing on the project's own precedent tools.** The same hazard class applies to
`proposals/SWG_Historical_Answers.html` and `..._v2.html`. This session did **not** audit those two
files for encoding damage and makes **no claim** about them; the observation is filed so the owner can
decide whether to check them. The v1 and v2 forms' own defect records (D-2, D-3) are concerned with
provenance and rendering fidelity, not character encoding.

### 6.6 Findings explicitly **not** made

- No finding that any ruling is wrong.
- No finding that the recollection record is unreliable. Its integrity metrics are strong (100% answered,
  0 placeholders, 91.1% substantive prose on the second pass) and its limits are already stated in its own
  §7 and §9.
- No finding that the checklist's row set is the correct one. Coverage is a judgement the owner can make
  (instrument §9 item 7).
- No finding that Proxy 4 should be re-opened. It stays deferred for the recorded reason (no live cohort).
- No finding about the encoding state of the two precedent forms. They were not audited; see §6.5 and
  §9 item 11.

---

## 7. Constraints compliance and status discipline

| Constraint | Compliance |
|---|---|
| **BAL-001** (no premature numeric tuning) | Honoured. No number was set, changed, proposed as a target, or validated into a constant. The report's magnitudes remain quoted magnitudes. |
| **HD-TST-01** (patterns-only fork) | Honoured. No code was written or modified. The vestigial condition at §3.4 was **reported, not fixed**. |
| **GR-001** (canonical presence ≠ approval) | Honoured. The delivered checklist is PROPOSED in `proposals/`, unadopted. |
| **GR-002 / GR-003** | Honoured. No historical material promoted; no derived constraint treated as locked. |
| **GR-004** (author/assessor attribution) | Honoured. Author, assessor state, and the same-identity caveat are all recorded; no self-assessment is claimed. |
| **GR-006.1** (rulings live in the register) | Honoured. No ruling was issued, so nothing was filed. The register was read, not written. |
| **GR-006.2** (dependent-document consistency) | **The session found a violation of this rule in a pre-existing document (§6.1) and reported it rather than fixing it.** The session's own new document does not assert a status contradicted by the register. |
| **GR-006.8** (detection is owed to the owner) | Honoured. The defect is surfaced with specific IDs, file paths and dates, for the owner's decision. |
| **GR-007.1 / GR-007.3** (true identity; model disclosure) | Honoured. The author field names the model, provider and version; the same-identity relationship to the recollection report is disclosed. |
| **GR-007.4** (no silent correction of another's byline) | Honoured. The recollection report was not edited. |
| **GR-007.6** (reporting is not authorisation) | Honoured. No rewrite, no re-attribution, no unilateral correction. |
| **Project `AGENTS.md` — "Document the project that exists"** | Honoured. No design decision was made; no contradiction was silently reconciled; every status assertion is traceable to a named authority. |
| **Project `AGENTS.md` — no silent promotion** | Honoured. The checklist is PROPOSED; the report is EVIDENCE; the form and its build script carry no authority of any kind. |
| **D-2 class** (hardcoded date / placeholder handle in exports) | Avoided. The form stamps the real date and the executor's real input; unrecorded fields export as the explicit token `UNRECORDED`, never a placeholder that could pass for data. |
| **D-3a class** (rendered text silently diverging from source) | Avoided by construction. The form is **generated from** the instrument and verified against it character-for-character — 0 mismatches across all 47 rows × 4 cells. |
| **D-3b class** (export path failing only at run time) | Avoided and **tested**. The browser script contains no regex literals and no template literals; the export path was exercised in a live browser by real button click, with zero runtime errors. |

---

## 8. Verification performed

All checks were run against the live files, not against this session's memory.

| Check | Method | Result |
|---|---|---|
| Row-ID integrity | Regex extraction + gap detection over `HA-01`…`HA-47` | **47 unique IDs, no gaps** |
| Coverage-map consistency | Parse every row's strength cell; compare against §6's three lists | **Exact match** — 34 ATTESTED / 9 THIN / 4 ABSENT, no row unlisted, no miscategorisation |
| `F-n` finding citations | Every cited finding resolved against the report's own heading | **29 distinct findings, all resolve** |
| `Q-n` answer citations | Every cited answer resolved against the report's `### Q-n` headings | **47 distinct, all resolve** |
| `C-Q-n` answer citations | Every cited second-pass answer resolved against `### C-Q-n` headings | **33 distinct, all resolve** |
| **Total citations verified** | | **109, zero unresolved** |
| Stale status strings in the instrument | Scan for `OQ-nn … (TBD)` | **0 remaining** after the v0.2 correction |
| Harness behaviour | Direct read of `cmd/phase10test/main.go` lines 633–711, plus `phase10_handlers.go` telemetry struct and anomaly-kind constants | Confirmed; §3.4 condition identified |
| Live surface statuses | Human Rulings Register, Open Questions Register v1.2, and Master GDD v1.1.1 §"Applied 2026-09-28" table read and cross-compared | **All three agree**; table in §6.1 |
| **Form ↔ instrument fidelity** | Build re-ran the instrument's own 47 rows through the form and compared every cell (core loop, Pre-CU account, strength, fork observation) | **0 mismatches across 188 cell comparisons**; 47 row blocks, 47 verdict selects, 14 domain options |
| **No pre-filled verdicts** | Enumerated all 47 verdict selects at run time | **0 pre-filled**; all ship on *(not yet recorded)* |
| **Verdict vocabulary parity** | Compared each select's option list against the instrument's §1.2 vocabulary | **47/47 selects identical**, 0 differences |
| **Export path (D-3b class)** | Live browser: real button click on *Export run record*, then read the produced DOM text | **No runtime error**; 13,484-char / 456-line record produced; header, tally, findings, owner-questions, 47 per-row entries all present |
| **Unfilled-row integrity** | Parsed every per-row `Verdict:` value in the produced export | **All are vocabulary terms or `UNFILLED`**; none written as a match; tally reconciles exactly with the per-row entries (5 recorded, 42 unfilled) |
| **Tally ≠ score** | Inspected the live UI and the exported record | Both carry the explicit *not-a-score* guard text; no composite figure is produced anywhere |
| **Filter behaviour** | Live browser, each filter applied in turn | DIVERGES→1, UNFILLED→43, RULED→7 (matches the 7 ruled-surface rows), Domain A→4, anti-goal→8, reset→47 |
| **Autosave / hydrate round-trip** | Filled, reloaded, re-read the DOM | 5 rows, the executor, the verdict values and the mode checkbox all restored; `localStorage` key `swg_ha_checklist_v0_2` |
| **Clear behaviour** | Clicked *Clear all* with `confirm` stubbed | Reset to 0/47, verdicts and executor cleared, `localStorage` reduced to the empty state |
| **Network isolation** | Scanned the generated file for remote `src`/`href` | **0 external resources**; no network calls |
| **Rendering** | Computed-style probe of the live page | 68 CSS rules parsed; dark theme, badges, inputs and the authority panel all styled; page title renders the em dash correctly |
| **Encoding (F-4)** | Byte- and codepoint-level scan of the generated file | **0 mojibake markers**; 55 correct em dashes, 37 section signs, 0 stray high codepoints |
| **Build reproducibility** | Re-ran the filed build script from a different working directory | **Byte-identical output** — SHA-256 `FA8839596B8C0E0009A4568E24D9F94F2B5E40358CF71016089EEBCD4614F7E1` both runs |
| Working-tree hygiene | `git status` | Only new files were added by this session (the checklist, the form, the build script, this report); **no commit was made**; the ~40 pre-existing modified and untracked paths are untouched and not adopted |

---

## 9. Unresolved items handed forward

1. **Owner ruling — adopt or amend the checklist** as the §1.3 Historical Accuracy instrument
   (instrument §9 item 1; default: not adopted, Part 2 stays unrun).
2. **Owner ruling — the three §2.3 band questions** (§4.3 / instrument §3.1–§3.3): the 4–6% stall
   commission, the 30–60 s PvP TTK band, and whether the three uncomparable progression bands should be
   logged as having no Pre-CU comparand. Defaults: all bands stand as written.
3. **Owner ruling — the report's stale status column** (§6.1 / instrument §9 item 8): issue a dated v1.7
   amendment, or rely on the register as authority. Default: leave the record unamended.
4. **Owner ruling — the BLD-004 mechanism divergence** (§6.3): the ruling chose *removal*; the historical
   account records *damage-at-lapse* plus *dissolution-on-City-Hall-destruction*. Recorded as a
   re-opening question, not a proposed change.
5. **Owner ruling — which observation mode to run first** (instrument §9 item 5). Default: Mode A
   (design) only, as a documentation pass.
6. **Owner ruling — whether to co-schedule Part 1 and Part 2** against one cohort when a cohort exists
   (instrument §6.1). Default: no scheduling change.
7. **Owner ruling — whether the row set should be extended** beyond the 47 derived rows (instrument §9
   item 7). Default: row set as drafted.
8. **Vestigial harness condition** (§3.4 / §6.4): owner decision on whether to clean up
   `len(aList) >= 0`. No code changed.
9. **Independent assessment** of the delivered checklist (and, separately, of the recall report's stale
   column) from a model other than this session's identity. Not yet performed; recommended before
   adoption.
10. **Carry-over unchanged from the 2026-10-18 programme state:** Proxy 4 remains DEFERRED pending a live
    cohort; the observation log and the Historical Accuracy verdict remain owner acts; no cohort exists to
    run Mode C against.
11. **Optional encoding audit of the two precedent forms** (§6.5 / F-4): `proposals/SWG_Historical_Answers.html`
    and `..._v2.html` were not audited for the same character-encoding hazard this session found in its
    own build. No claim is made about them either way. Default: not audited.

---

## 10. Related documents

| Document | Relationship |
|---|---|
| `proposals/SWG_Historical_Accuracy_Checklist_v0.2.md` | **Delivered by this session.** PROPOSED, unadopted, unrun, unassessed, uncommitted |
| `proposals/SWG_Historical_Accuracy_Checklist.html` | **Delivered by this session.** Local fill-and-export playtesting form. No project authority. Generated from the instrument and verified against it; zero external resources |
| `proposals/SWG_Historical_Accuracy_Checklist_build.ps1` | **Delivered by this session.** The form's build script, filed for reproducibility. Byte-reproducible; carries a UTF-8 BOM and a mojibake guard. No project authority |
| `proposals/SWG_Historical_Answers.html` · `SWG_Historical_Answers_v2.html` | The precedent fill-and-export forms the new form was modelled on. Same category: local transport tools carrying no project authority |
| `investigations/TCIndustries_SWG_Historical_Recollection_Report_Tiwa_Pene_2026-10-21.md` (v1.6) | The Pre-CU account source. EVIDENCE + ANALYSIS. **§6.1 finding filed against it; not edited** |
| `investigations/tiwas-swg-recollections-21-10-2026.md` · `tiwas-swg-recollections-22-10-2026.md` | Raw unaltered answer records backing the report |
| `proposals/SWG_Historical_Accuracy-v1.md` (v1.0) | The 223-question instrument that produced the corpus |
| `proposals/SWG_Historical_Accuracy_v2_Targeted_ReAsk.md` (v1.0) | The 56-question second instrument; the 8 named surfaces |
| `proposals/SWG_Phase10_Balance_Telemetry_Polish_Proposal_v0.2.md` | §§1.2, 2.1, 2.2, 2.3, 5.1–5.3, 6, 9, 10 — the obligations and bands the instrument serves |
| `proposals/SWG_Post_Phase10_Program_Proposal_v0.1.md` | §9 item 8 and the 2026-10-18 split decision; §9 item 9; §10 |
| `governance/TCIndustries_Human_Rulings_Register.md` | Sections C and D — the authority for the §6.1 status table |
| `proposals/TCIndustries_Open_Questions_Register.md` (v1.2) | Second, independent confirmation of the §6.1 statuses |
| `proposals/TCIndustries_Invariant_Register_v0.1.md` | INV-007a, INV-009a, INV-010, INV-011a, INV-016a/b, INV-BLD-004b — approved design invariants touched by the instrument's rows |
| `governance/TCIndustries_Authority_Provenance_Reconciliation_Matrix.md` | GR-001–GR-004, GR-006.1/2/8, GR-007.1–GR-007.6 |
| `canonical/gdd/TCIndustries_Master_GDD_v1.1.1_Status_Patch.md` | §2 central loop and anti-goals; the "Applied 2026-09-28" status-line table |
| `testbed/swg-phase3-combat/swg-server/cmd/phase10test/main.go` | Lines 633–711 — the Proxy 4 and Historical Accuracy behaviour documented in §3 |
| `sources/swg-pre-cu/SWG_PreCU_GDD.md` | HISTORICAL reference text — the other Pre-CU source, and the higher authority where the two conflict (report §7 item 8) |

---

*Prepared by MiniMax M3.1 Flash (free) — model `openrouter/minimax/minimax-m3.1-flash:free` — via OpenCode,
2026-10-22. This report documents what exists, what was verified, and what was found. It invents no
decisions, promotes no status, changes no authority, and alters no pre-existing project document. The one
defect it identifies in a pre-existing document (§6.1) is reported for the owner's decision, not corrected
unilaterally (GR-006.8). The three deliverables it produces are
`proposals/SWG_Historical_Accuracy_Checklist_v0.2.md` (PROPOSED, unadopted, unrun),
`proposals/SWG_Historical_Accuracy_Checklist.html` and
`proposals/SWG_Historical_Accuracy_Checklist_build.ps1` (local transport and build tools, no project
authority) — all uncommitted.*
