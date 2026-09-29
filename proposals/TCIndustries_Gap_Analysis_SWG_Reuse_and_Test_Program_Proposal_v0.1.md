# TCIndustries — Gap Analysis, Pre-CU SWG Reuse Mapping, and Test Program Proposal v0.1

**Status:** PROPOSED plan. No authority. No design decision. No code until the §9
task list is approved item-by-item by the project owner.

**Author / Assessor:** Buffy — Codebuff agent (`codebuff/freebuff`) via Freebuff,
**Proposal v0.1**, 2026-09-17 — analysis, mapping, and test planning only; all
authority decisions remain human.

**Responds to:** owner request of 2026-09-17: *"analyze the current TCIndustries
game development document. Identify what needs to be completed. Identify what can
be used from the Pre-CU SWG. Write a formal document proposal to test current
missing parts of TCIndustries. Make this an actionable document which gives tasks
I can approve for an LLM to complete."*

**Controlling references:** `AGENTS.md`;
`canonical/gdd/TCIndustries_Master_GDD_v1.1.1_Status_Patch.md` (the "current game
development document" analyzed here);
`proposals/TCIndustries_Open_Questions_Register.md`;
`proposals/TCIndustries_EIC_Program_State.md`;
`proposals/TCIndustries_SWG_Testbed_Reuse_Proposed_Ruling.md` (HD-TST-01);
`investigations/OpenCode_Project_Documentation_Report_2026-09-17.md`;
`investigations/PreCU_SWG_Program_User_Guide_2026-09-17.md`.

---

## 1. Status discipline (read first)

- This proposal is **PROPOSED**. It changes no GDD status, opens no gate, and
  answers no open question.
- **Every task in §7 requires explicit owner approval before an LLM executes
  it.** Approval is per-task, not all-or-nothing.
- The Pre-CU SWG material referenced is **HISTORICAL** (SWG-002): it supplies
  *test patterns and evidence*, never design content. Nothing in this proposal
  promotes any SWG mechanic, value, or name into TCIndustries.
- All testing here is **shape-testing** (does a structure hold up?), never
  numeric tuning (BAL-001; numerics stay TBD and out of scope).
- EIC boundaries: **HD-EIC-05** (Architecture C retired) and **HD-EIC-07**
  (no new shaping without human-supplied structural boundaries) remain in force.
  Where a task would touch EIC, it **informs** the gate; it does not open it.
- Each task is a self-contained LLM assignment with an acceptance gate —
  designed so approval of one task neither requires nor triggers any other.

## 2. Part I — What the current TCIndustries development document contains

Analyzed document: **Master GDD v1.1.1 Status Patch** (2026-08-25, 1,019 lines,
35 rule sections + registers), the project's working canonical design reference.

**What is already settled (LOCKED / HUMAN-LOCKED):** the full vision layer —
living persistent world (VIS-001), citizen-not-hero fantasy (VIS-002),
player-driven economy (VIS-003), interdependence (VIS-004), emergence (VIS-005),
anti-goals (VIS-006); the four design pillars (discovery/gold rush, identity &
reputation, interdependence, player-created society); non-combat viability
(LOOP-002) and option-expanding progression (LOOP-003/PROG-001); flexible
skill-based professions (PROF-001) with respecialisation (PROF-004); resource
design goal (RES-002); differentiated crafting (CRFT-001, principle-only);
controlled automation (MFG-001); player retail (RET-001), services (SERV-001),
structures (BLD-001), organisations (SOC-001), emergence (EMRG-001); ownership/
transfer safeguards (SAFE-001, principle-only); SWG inspiration-only boundary
(SWG-001/002); prototype-subordination (PROT-001); no-premature-numerics (BAL-001).

**What remains open — the completion backlog (all verified against the document
and its registers):**

| Backlog area | GDD anchors | Current status |
|---|---|---|
| Specialisation budget / skill-cap model (what a character *cannot* be at once) | OQ-001, PROF-005, §36.1 | TBD |
| Skill acquisition model (points / XP / training / discovery) | OQ-011, PROG-003, PROF-004 rules | TBD |
| Resource lifecycle mathematics & attribute→outcome mapping | OQ-002, OQ-003, RES-005/006 | TBD (numbers explicitly forbidden for now) |
| Crafting experimentation & schematic structure | OQ-004, CRFT-002/003 | PROPOSED / TBD |
| Manufacturing facility rules & automation constraints | OQ-005, MFG-002/003 | PROPOSED / TBD |
| Transportation & logistics | OQ-006, WRLD-005 | PROPOSED / TBD |
| City formation, governance, taxation/maintenance | OQ-007, BLD-003, WRLD-003 | PROPOSED / TBD |
| Combat scope, risk model, progression integration | OQ-008, CMBT-002 | PROPOSED / TBD |
| Currency, money sinks, economic stabilisers | OQ-009, ECO-004 | TBD |
| Item durability, decay, repair | OQ-010, ITM-002 | PROPOSED / TBD |
| Respecialisation exact costs/rules | OQ-011, PROF-004 | LOCKED principle, TBD rules |
| Multi-accounting & exploit *implementation* | OQ-012, SAFE-002 | stance settled (HD-EIC-03), implementation TBD |
| Organisation rights, hierarchy, property | OQ-013, SOC-002 | PROPOSED / TBD |
| Vendor & retail mechanics | OQ-016, RET-002 | PROPOSED / TBD |
| Provenance depth & visibility | OQ-017, CRFT-004, ITM-001 | PROPOSED (engine DEFERRED) |
| First-playable vs persistent-alpha scoping | OQ-015 | TBD |
| Species/ancestry/setting | OQ-014 | DEFERRED (out of scope here) |

**Process state of the GDD's own §34 sequence:** step 0 (Evidence Reconciliation
Pass — GDD ↔ prototype ↔ decision matrix) is described there as the required next
work and remains **open**; the EIC programme sits at the **post-falsification
gate** (HD-EIC-07: awaiting a human statement of durable structural boundaries;
no Architecture D, no residual-channel invention, no tuning). Tasks T-03 and T-04
below are the only ones that touch these, and both produce *analysis artifacts*,
not design decisions.

## 3. Part II — What can be used from Pre-CU SWG (reuse map)

Per HD-TST-01 (patterns-only), the OpenCode reusability split (~30% direct /
~50% adaptable / ~20% never), and SWG-002, the usable material falls in three
tiers. This table is the bridge between Part I's gaps and Part III's tests.

### 3.1 Directly reusable (IP-neutral engineering — lifts as-is)

| Asset | Serves GDD gap |
|---|---|
| Telemetry layer: `interdependence_events`, `phase10_jobs` (§29.8-shaped claim/complete job record), read-only REST pattern | Economic observability (SAFE-003); EIC evidence collection |
| Report-generator taxonomy: OBSERVED / CODE-CONSTANT / DERIVED / NOT-MEASURABLE rows with citations | TEST-001 invariants; every future validation report |
| Test-harness conventions: three-verdict proxies (PASS / FAIL-with-telemetry / NOT-MEASURABLE), raw counts beside percentages, fresh-DB lifecycles, env-overridable endpoints | TEST-002 simulation preference; all test tasks |
| `newRowID` ID mint; suite runners (`run_chain.sh` pattern: fresh DB per phase, listener reaping, dedicated ports) | Any TCIndustries implementation's infrastructure |

### 3.2 Adaptable (mechanical patterns — structure informative, values never)

| SWG-precedent pattern | Informs GDD gap (shape only) |
|---|---|
| Skill-and-box progression with 250-pt specialisation budget; respecialisation via skill-dropping | OQ-001 (budget model), OQ-011 (respecialisation costs) — shapes, not values |
| Typed XP pools + trainer-gated acquisition | OQ-011 (acquisition model) |
| Resource spawns: temporary, variable-attribute, survey-gated information | OQ-002/003 (lifecycle + attribute *shapes*; no numbers) |
| Crafting with experimentation and resource-quality-sensitive outcomes; `crafted_items` provenance | OQ-004 (experimentation model), OQ-017 (provenance depth) |
| Faucet/sink-tagged credit ledger; append-only sale history (`market_records`); volatility/Gini queries | OQ-009 (currency/sink *structure*), SAFE-003 |
| Service professions: medical wounds/heals, entertainer buffs — the interdependence loop | SERV-002 (service interdependence), EIC evidence |
| Vendor/retail flows, structure placement, city-like civic scaffolding | OQ-016 (vendor mechanics), BLD-003 shapes |
| Mission/contract lifecycle incl. escrow | Demand-generation shapes (EIC-adjacent) |

### 3.3 Never usable (hard exclusions — restated for the record)

SWG branding, lore, names, species, places, or any protected content; HISTORICAL
balance bands (§9.7/§33/§34) as tuning targets; all `[PROVISIONAL]`/fork numeric
constants as TCIndustries values; the non-canonical list (Open Questions Register
§6 — e.g. skill-discipline/box architecture as a *chosen* model, single unified
currency as default, soul-binding rules, use-based skill-point acquisition);
predecessor LLM outputs as authority of any kind.

## 4. Part III rationale — why "test the missing parts" means shape-testing

The GDD's own methods (TEST-001, TEST-002, §34) define what testing can
legitimately happen now: **design invariants** made testable, and **simulation
before implementation**. BAL-001 forbids inventing numbers. Therefore the missing
parts of TCIndustries cannot be "tested" by tuning — they can be tested by:

1. **Invariant formulation** — turning each open design area into explicit
   Principle → Constraint → Invariant → Failure Condition → Test forms (the
   §34 EIC expression), so that future numeric work has something to test.
2. **Simulation** — checking structural properties (scarcity, interdependence,
   anti-monopoly, anti-automation-drift) hold *under adversarial pressure*,
   using patterns from the Pre-CU fork as reference behavior, with all numbers
   clearly labeled NON-CANONICAL test scaffolding.

Both are exactly the LLM-executable, human-approvable task shapes the owner
requested. Where a gap is purely a human decision (currency choice, combat
scope), the correct "test" is a decision-support brief, not a simulation.

## 5. Scope

**IN:** analysis artifacts, invariant registers, simulation specifications and
sandboxed runs, decision-support briefs, evidence registers — all as PROPOSED
documents under `proposals/` and `investigations/`.

**OUT:** numeric balance values; new EIC architecture or residual channels
(HD-EIC-05/07); GDD edits or status changes; provenance/reputation engine design
(DEFERRED); setting/species content (OQ-014 DEFERRED); any commit, test run, or
prototype modification before per-task approval.

## 6. Assumptions and honest limits

1. **No TCIndustries implementation exists.** All simulations are sandboxed
   studies (likely extending the EIC comparative-simulation toolchain in
   `archive/`/`investigations/eic/`), not features of a game build.
2. **Non-canonical shapes are still informative.** Simulating a *disallowed*
   shape (e.g. skill-box architecture) to demonstrate *why* it fails a locked
   principle is legitimate testing — it does not promote the shape.
3. **Task effort is unquantified.** Proposals here deliberately give no
   calendar estimates (consistent with prior proposals).
4. **Task ordering is recommendation, not requirement.** T-01/T-02 unblock the
   most; everything else stands alone.
5. **The EIC gate bounds everything.** If simulations appear to "choose" an
   interdependence architecture, they have overstepped — findings route to the
   owner as input to an HD-EIC-07 boundary statement, never as selection.

## 7. Part III — Actionable task list (approve per task; each is one LLM assignment)

**Task T-01 — TCIndustries Invariant Register (TEST-001 operationalized).**
*Objective:* convert every open backlog area (§2) into candidate design
invariants in the §34 form (Principle → Constraint → Invariant → Failure
Condition → Test), each traced to its GDD anchor (OQ-xxx / rule ID) and each
marked PROPOSED. *Method:* documentarian analysis of the GDD + registers; no
numbers; where an area cannot yield an invariant (e.g. species), record why.
*Output:* `proposals/TCIndustries_Invariant_Register_v0.1.md`.
*Acceptance:* every OQ-001…017 addressed (addressed = invariant proposed, or
documented as not-invariant-shaped); zero numeric values; every entry carries a
GDD citation; status discipline intact. *Authority consumed:* none.

**Task T-02 — Design-Evidence Register (Evidence Reconciliation Pass, step 0).**
*Objective:* index every testbed evidence artifact (Phase 0–10 suites, balance
report, telemetry, service-XP options paper, verification reports) and map each
to the GDD questions it can and cannot inform, with explicit
HISTORICAL/PROTOTYPE status walls. This *executes the GDD §34 step 0 process
requirement* against the testbed evidence corpus. *Method:* repository-wide
inventory; mapping table; gap list (evidence needed but absent).
*Output:* `investigations/TCIndustries_Design_Evidence_Register_v0.1.md`.
*Acceptance:* every cited artifact verified to exist at cited location; every
mapping states what the evidence does NOT establish; OQ/§34 citations exact.
*Authority consumed:* none (registers evidence; decides nothing).

**Task T-03 — Specialisation-Budget Shape Study (OQ-001, PROF-005, CRFT-006).**
*Objective:* compare candidate specialisation-budget *shapes* (e.g. single
point-pool, per-domain pools, activity-based decay of unused mastery, org-level
complementarity) against the LOCKED constraints: PIL-003, PROF-004
(respecialisation), HD-EIC-01 (validation at independent-participant layer),
HD-EIC-02 (no hard factory-quality ceiling), anti-goals. *Method:* structured
analysis + adversarial walk-throughs per shape (how does a multi-account or
organisational actor neutralise each shape?); explicitly **no shape selection**
— comparison table + failure modes + what evidence would discriminate.
*Output:* `investigations/TCIndustries_Specialisation_Shape_Study_v0.1.md`.
*Acceptance:* each shape expressed without numbers; adversarial analysis covers
the full HD-EIC-01 actor layer; recommendations framed as "candidate shapes for
human evaluation," not winners. *Authority consumed:* none (EIC-adjacent but
shapes analysis, not architecture).

**Task T-04 — Resource Lifecycle Sandbox Simulation Spec + Run (OQ-002/003,
TEST-002, PIL-001).**
*Objective:* specify and run a small simulation testing structural properties of
temporary variable-quality resource spawns against PIL-001/RES-002 invariants:
do discovery→depletion→rediscovery cycles generate *events* (regional traffic,
price movement, information trade) rather than inventory churn; does information
asymmetry (RES-007) survive bot/simulated-actor pressure; does anti-monopoly
behavior hold. *Method:* parameters generated as labeled NON-CANONICAL test
scaffolding (BAL-001 respected — they test the *shape*, they propose nothing);
reuse the fork's resource-spawn generator *pattern* where useful; all results
reported with the three-verdict taxonomy.
*Output:* `investigations/TCIndustries_Resource_Sim_Spec_v0.1.md` +
`investigations/TCIndustries_Resource_Sim_Results_v0.1.md`.
*Acceptance:* every parameter marked NON-CANONICAL with rationale; results
state which invariants held/failed/NOT-MEASURABLE; no tuning recommendations.
*Authority consumed:* none (TEST-002 simulation, not implementation).

**Task T-05 — Crafting-Differentiation Invariant Test Brief (OQ-003/004,
CRFT-001).**
*Objective:* from the LOCKED principle (differentiated products) derive testable
invariants (e.g. "equivalent inputs + equivalent skill ⇒ equivalent outputs;
differing inputs/skill ⇒ measurable, visible differentiation") and specify the
simulation/experiment that would test candidate experimentation *shapes*
(point-allocation is on the non-canonical list; alternatives must be explored)
without implementing any. *Method:* invariant derivation + test specification;
the testbed's crafted-share/provenance evidence (T-02) informs what is already
evidenced. *Output:*
`investigations/TCIndustries_Crafting_Invariants_Brief_v0.1.md`.
*Acceptance:* respects the non-canonical exclusion of point-allocation/
success-critical-failure models as *chosen* designs (analysis of why they were
excluded is in scope); zero numbers; tests specified to be runnable if approved.
*Authority consumed:* none.

**Task T-06 — Automation-Drift Structural Test Brief (OQ-005, MFG-001, MFG-004).**
*Objective:* define the adversarial test for the DERIVED CONSTRAINT: "no
configuration of factories/automation makes other players optional." Specify the
actor model (single specialist, generalist, multi-accounter, organisation),
the automation configurations to pit against them, and the measurable
interdependence signals (drawn from the fork's interdependence-event *pattern*).
*Method:* test design only; no implementation. *Output:*
`investigations/TCIndustries_Automation_Drift_Test_Brief_v0.1.md`.
*Acceptance:* test is falsifiable and actor-complete per HD-EIC-01; explicitly
states it cannot be run until an implementation or simulation host exists or a
host task is approved. *Authority consumed:* none.

**Task T-07 — Economic Observability Requirements Brief (OQ-009, SAFE-003,
RET-002).**
*Objective:* convert the fork's proven observability pattern (faucet/sink-tagged
ledger, append-only sale history, snapshot jobs, volatility/concentration
queries, anomaly review lists) into *TCIndustries requirements* — what must be
observable, at what event boundaries, with what privacy/fairness constraints —
so that whatever currency/stabiliser design the owner later chooses is
measurable from day one. *Method:* requirements derivation from evidence; no
currency choice. *Output:*
`proposals/TCIndustries_Economic_Observability_Requirements_v0.1.md`.
*Acceptance:* every requirement traced to observed evidence or a stated GDD
principle; no currency model selected or implied. *Authority consumed:* none.

**Task T-08 — Decision-Support Briefs for the Human-Only Gaps.**
*Objective:* for the areas that are pure human decisions, produce one-page
structured briefs each stating: the decision, the LOCKED constraints bounding
it, the candidate answer-space (from GDD + HISTORICAL precedent + testbed
evidence), what evidence exists, what T-01–T-07 would add, and the minimal
experiment that would discriminate. Topics: currency & sinks (OQ-009);
combat scope & risk (OQ-008); transportation (OQ-006); city governance (OQ-007);
item durability (OQ-010); respecialisation costs (OQ-011); multi-account
implementation (OQ-012); vendor mechanics (OQ-016); first-playable scoping
(OQ-015). *Method:* research + structuring; each brief ends in a decision
question for the owner, never an answer. *Output:*
`investigations/TCIndustries_Decision_Briefs_v0.1.md` (one section per topic).
*Acceptance:* no brief recommends; every brief cites its GDD anchors; the
"minimal discriminating experiment" is specified for each. *Authority consumed:*
none.

**Task T-09 — Post-Approval Assembly Pass (only after T-01–T-08 are approved
and delivered).**
*Objective:* consolidate delivered artifacts into the GDD's §34 sequence state —
a single "where the design programme stands" summary registering which invariants
exist, which simulations ran, which evidence gaps remain, and which §36 human
decisions are now ready for ruling — filed as an EVIDENCE report. *Output:*
`investigations/TCIndustries_Design_Programme_State_<date>.md`. *Acceptance:*
pure assembly of delivered artifacts; no new analysis; zero status changes.
*Authority consumed:* none.

### 7.1 Recommended sequencing and dependencies

| Order | Task | Depends on | Unblocks |
|---|---|---|---|
| 1 | T-01 Invariant Register | — | T-04/T-05/T-06 (invariants to test) |
| 2 | T-02 Evidence Register | — | T-03–T-08 (what's already evidenced) |
| 3 | T-03, T-04, T-05, T-06, T-07 | T-01 + T-02 | human decisions; future sims |
| 4 | T-08 Decision Briefs | T-02 (T-01 helps) | owner rulings on §36 items |
| 5 | T-09 Assembly | all above | §34 programme-state picture |

### 7.2 Effort class and approval mechanics

Each task is one focused LLM engagement (analysis/documentation class, no
server builds except the T-04 sandbox which creates no game code). Approval is
by task ID ("Approve T-01, T-02"), optionally with amendments. Every delivered
artifact returns as PROPOSED/EVIDENCE with its own acceptance section the owner
can audit against §7's criteria. No task consumes design authority; collectively
they prepare the material the owner needs for the §36 decisions and the
HD-EIC-07 boundary statement — which remain reserved human acts.

## 8. Tensions and honest limits (human resolves)

1. **Analysis vs paralysis.** Nine tasks before any "real" design work is a
   process-heavy path; the counterweight is that the GDD's own §34 sequence
   mandates step 0 and invariant-first design. The owner may approve a subset
   (e.g. T-01, T-02, T-08 only).
2. **NON-CANONICAL scaffolding risk.** Simulated numbers can acquire de-facto
   authority by repetition. Mitigation: every artifact carries status walls;
   the Open Questions Register §6 non-canonical list is restated in relevant
   tasks.
3. **Testbed coupling.** Reusing fork *patterns* is authorized (HD-TST-01);
   reusing fork *values* is not. The boundary is restated per task.
4. **EIC gravity.** T-03/T-06 sit near the EIC programme. They analyze shapes
   and specify tests; if any output reads as architecture selection, it
   violates HD-EIC-07 and must be corrected before filing.

## 9. Decisions requested (no task starts until answered)

**RECORDED 2026-09-17:** the owner approved **T-01** and **T-02** exactly as
specified ("Approve and execute T-01 and T-02 … then report both for audit").
Both tasks are EXECUTED; delivery and audit status are recorded in §10. All
other rows below remain unanswered with their defaults in force.

| # | Decision | Options | Default if unanswered |
|---|---|---|---|
| 1 | Approve T-01 Invariant Register | **APPROVED 2026-09-17 — EXECUTED** | — |
| 2 | Approve T-02 Evidence Register (GDD §34 step 0 execution) | **APPROVED 2026-09-17 — EXECUTED** | — |
| 3 | Approve T-03 Specialisation Shape Study | Approve / Amend / Decline | Not started |
| 4 | Approve T-04 Resource Lifecycle Sim (spec + run) | **APPROVED 2026-09-17 — EXECUTED** | — |
| 5 | Approve T-05 Crafting Invariants Brief | Approve / Amend / Decline | Not started |
| 6 | Approve T-06 Automation-Drift Test Brief | Approve / Amend / Decline | Not started |
| 7 | Approve T-07 Economic Observability Requirements | Approve / Amend / Decline | Not started |
| 8 | Approve T-08 Decision-Support Briefs (9 topics) | Approve all / select subset / Decline | Not started |
| 9 | Approve T-09 Assembly Pass | Approve / Decline | Not started |
| 10 | Overall sequencing (recommendation: T-01+T-02 first) | Confirm / reorder | Recommendation stands, no work |

## 10. Implementation status

| Task | Status | Artifact |
|---|---|---|
| T-01 Invariant Register | **EXECUTED 2026-09-17** (owner-approved) | `proposals/TCIndustries_Invariant_Register_v0.1.md` |
| T-02 Evidence Register | **EXECUTED 2026-09-17** (owner-approved; executes GDD §34 step 0 against the testbed evidence corpus) | `investigations/TCIndustries_Design_Evidence_Register_v0.1.md` |
| T-04 Resource Lifecycle Sim | **EXECUTED 2026-09-17** (owner-approved; spec + standalone simulator + envelope run + invariant verdicts) | `investigations/TCIndustries_Resource_Sim_Spec_v0.1.md`, `investigations/TCIndustries_Resource_Sim_Results_v0.1.md`, `investigations/resource-sim/` |
| T-03, T-05 … T-09 | NOT STARTED (not approved; §9 defaults in force) | — |

All artifacts delivered as PROPOSED/EVIDENCE with auditable acceptance sections.
Phase-10/post-Phase-10 proposals remain separate documents with their own
pending decisions; nothing here supersedes them.

---

*End of proposal v0.1. To enact: reply with task approvals by ID (amendments
welcome). Each approved task will be executed, verified against its acceptance
criteria, and returned for audit before the next begins.*
