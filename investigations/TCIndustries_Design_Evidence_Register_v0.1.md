# TCIndustries — Design Evidence Register v0.1

**Status:** EVIDENCE (a finding aid). This register decides nothing, promotes
nothing, and answers no open question. Every mapping below states what evidence
**can** inform and — equally — what it **cannot** establish.

**Author:** Buffy — Codebuff agent (Freebuff), T-02 deliverable, 2026-09-17.
**Task authorization:** owner approval of T-02 under
`proposals/TCIndustries_Gap_Analysis_SWG_Reuse_and_Test_Program_Proposal_v0.1.md`
§9 (recorded there 2026-09-17). This register **executes the Master GDD §34
step 0 process requirement** — the Evidence Reconciliation Pass — against the
testbed evidence corpus, as scoped by the approved task.

**Method:** repository-wide inventory of evidence artifacts (every artifact below
was verified to exist at its cited location on 2026-09-17), then a mapping table
to the GDD open questions (OQ-001…017) with explicit "what this does NOT
establish" columns, then a gap list (evidence needed but absent).

**Standing rule (SWG-002 / HD-TST-01):** all testbed behavior is
PROTOTYPE/EVIDENCE. All predecessor-GDD values are HISTORICAL. Nothing here is
TCIndustries design authority. The companion Invariant Register
(`proposals/TCIndustries_Invariant_Register_v0.1.md`) holds the corresponding
candidate invariants; cross-references use its INV-xxx IDs.

---

## 1. Evidence artifact inventory (verified 2026-09-17)

### 1.1 Testbed verification artifacts (`testbed/swg-phase3-combat/swg-server/`)

| ID | Artifact | Location | Verified content |
|---|---|---|---|
| E-01 | Phase 0–2 predecessor suites | `sources/swg-pre-cu/code-phase-0..2/`, zip snapshots | Server foundation/HAM; movement/spatial/chat; skills/XP/training — predecessor lineage, pre-fork |
| E-02 | Phase 3–8 integration clients | `cmd/phase3test` … `cmd/phase8test/` | Combat core; resources/crafting; economy; services; civic; faction/PvP — all suite-green in the 2026-09-17 full chain |
| E-03 | Phase 9 suite (8 tests) | `cmd/phase9test/`, `docs/phase9_test_report.md`, `run_phase9.sh` | 28 professions/504 boxes; elite gates; lairs; missions incl. bounty escrow; commerce XP hooks; `PHASE9_EXIT:0` |
| E-04 | Phase 10 proxy harness | `cmd/phase10test/` | Five §1.3 proxies, three-verdict verdicts; `PHASE10_EXIT:0`; observed values: combat cohort 2/2 receipts, novice-vs-+tier1 ratio ≈0.93–1.08×, crafted share 1/1, mission mix 29 rows |
| E-05 | Balance validation report | `cmd/phase10report/`, `docs/phase10_balance_report.md` (43 table rows) | OBSERVED / CODE-CONSTANT / DERIVED / NOT-MEASURABLE rows; headline: fork unarmed 5–15 OUT-OF-BAND vs HISTORICAL §9.7 50–150; drop/reward constructions IN-BAND |
| E-06 | Service-XP characterization + options paper | `docs/phase10_service_xp_options.md` | All flat `[PROVISIONAL]` service-XP sites cited file:line; §7.2.3-shaped options, no rate changed |
| E-07 | Telemetry layer | `internal/database/phase10_db.go`, `internal/handlers/phase10_handlers.go` | Append-only `interdependence_events` (wound-heal / buff / crafted-purchase hooks; heal-kind code-path-verified only — buff+purchase observed live); `phase10_jobs` §29.8-shape proven running (7 completed records); read-only REST `/api/phase10/*` |
| E-08 | Anomaly review lists (B5(i)) | `phase10_handlers.go` | Wash-trade pairs, near-zero-net churn, placement bursts, price spread — structural invariants, no thresholds, human-review-only |
| E-09 | Ledger & market patterns | `internal/database/economy_db.go` (`AtomicPurchase`), `market_records` | Faucet/sink-tagged credit ledger; append-only sale-price history; provenance-checked crafted purchases |
| E-10 | Full-chain regression record | `run_chain.sh`; `investigations/SWG_Phase10_Implementation_Verification_Report_2026-09-17.md` | Phases 1–9 fresh-DB sequential green (`CHAIN:ALL-GREEN`) + Phase 10 (`PHASE10_EXIT:0`), commits `f9852ec`…`2730c9a` |

### 1.2 EIC programme artifacts (`investigations/eic/`)

| ID | Artifact | Verified content |
|---|---|---|
| E-11 | EIC Design Investigation → Candidate v0.1 → Falsification Pass | M1–M8 candidate mechanisms; falsification record |
| E-12 | Comparative Simulation Specs + Results v0.1/v0.1.1 (Tests A/B; α/β/γ packages) | α FAIL, β FAIL, γ INCONCLUSIVE under adversarial packages |
| E-13 | γ′-1/γ′-2 Adversarial Retest Results | Both FAIL — full adversarial closure achieved by combined neutralisation strategies |
| E-14 | HD-EIC-09 Reconciliation (2026-09-12) | Gate-record reconciliation of the EIC chain |
| — | **Programme state** (`proposals/TCIndustries_EIC_Program_State.md`) | Architecture C RETIRED; no leading hypothesis; HD-EIC-07 gate awaits human structural-boundary statement |

### 1.3 Predecessor craft demos (`sources/swg-pre-cu/`, HISTORICAL)

| ID | Artifact | Verified content |
|---|---|---|
| E-15 | `craft-seed2-1-pro-preview/` (PROT-002 identity CONFIRMED 2026-09-12) | Vite/React/TS crafting-core demo: assembly + experimentation engine, pure-TS core |
| E-16 | `craft-claude-opus-5-max/` | Crafting architecture demo: galaxy resource generation, spawn browser, FSM craft bench, self-test |
| E-17 | `java-swg-skillcalc/` | Pre-CU skill-rule calculator (250-pt pool, prereq chains, surrender guards) |
| E-18 | `track-a-b-code/` | Reference functions: combat resolution, resource-spawn generator, inflation sim, experimentation outcomes, XP curve |

### 1.4 Governance & analysis artifacts

| ID | Artifact | Content class |
|---|---|---|
| E-19 | Master GDD v1.1.1 + Open Questions Register + Authority Matrix + Human Rulings Register | Canonical design reference + status/authority record (not evidence of behavior — evidence of *decisions*) |
| E-20 | `investigations/OpenCode_Project_Documentation_Report_2026-09-17.md` | Reusability split (~30% direct / ~50% adaptable / ~20% never); Functioning-Prototype assessment |
| E-21 | `investigations/PreCU_SWG_Program_User_Guide_2026-09-17.md` | Program documentation: purpose, operation, history, SWG background |
| E-22 | `investigations/SWG_Code_Due_Diligence_Evidence_Log_2026-09-14.md` + evidence-reconciliation passes | Due-diligence record for the predecessor corpus |

---

## 2. Mapping: evidence → GDD open questions

Legend: **Informs** = can legitimately shape the human/LLM design discussion for
that OQ. **Does NOT establish** = the specific thing the evidence is often
mistaken for, which it cannot provide. INV-xxx = the companion invariant the
evidence bears on.

| OQ | Evidence | Informs | Does NOT establish |
|---|---|---|---|
| **OQ-001** budget model | E-04 (Proxy 2: bounded progression measured ≈0.93–1.08× with/without tier-1 under the fork's cap regime); E-17 (250-pt budget shape, HISTORICAL); E-02 phase-2 skills | What a working budget *regime* looks like mechanically; that budget-vs-output relations are measurable | Any TCIndustries cap model or value; INV-001 requires actor-complete adversarial analysis (T-03) — not yet run |
| **OQ-002** resource lifecycle | E-16/E-18 (spawn-generator patterns, HISTORICAL); E-02 phase-4 spawn behavior (fork implementation, PROTOTYPE) | That temporary/variable spawns are implementable and surveyable; shape vocabulary (cycles, windows, attribute ranges) | Lifecycle mathematics; whether discovery→event propagation holds at TCIndustries scale (INV-002; T-04 sim unrun) |
| **OQ-003** attribute→outcome mapping | E-05 (report rows on drop/reward constructions, IN-BAND observations); E-18 experimentation-outcome reference; E-16 | Attribute-recipe coupling is implementable; reporting taxonomy for the coverage matrix (INV-003) | Any attribute set; any mapping curve — zero TCIndustries attributes exist yet |
| **OQ-004** experimentation model | E-15/E-16 (two independent experimentation-engine demos, HISTORICAL); E-04 crafted-share 1/1; E-09 provenance-checked purchase | Experimentation is implementable in ≥3 independent shapes; differentiation + provenance attach is demonstrable | Which experimentation shape TCIndustries should choose; point-allocation remains non-canonical regardless of demo quality |
| **OQ-005** manufacturing/automation | E-09 (transactional purchase core); MFG-001/004 LOCKED constraints | Requirement vocabulary for INV-005a/b; what "closed-loop self-sufficiency" would look like as a test (T-06 brief) | Any automation rules; no factory system exists in any codebase — NOT-MEASURABLE today |
| **OQ-006** transportation | E-05 (NOT-MEASURABLE rows); E-20 (adaptation tier) | Nothing direct — honest gap | No transport evidence exists anywhere in the corpus (see §3 gaps) |
| **OQ-007** city governance | E-02 phase-7 civic suite (cities/guilds/mail — PROTOTYPE) | That civic scaffolding (membership, permissions, mail) is implementable and suite-testable | Governance/taxation design; INV-007a/b untested anywhere |
| **OQ-008** combat scope/risk | E-02 phase-3/8 suites (HAM combat, flagging, duels, death penalties); E-05 OUT-OF-BAND finding; E-06 service-XP link | Combat-economy coupling is implementable (combat consumes/creates service demand — E-07 buff receipts); scale-band observation | TCIndustries scope/risk decisions; the OUT-OF-BAND finding is a fork observation, not a TCIndustries band |
| **OQ-009** currency & sinks | E-09 (faucet/sink-tagged ledger — observed working); E-07/E-08 (observability surface) | Ledger-tagging + sink accounting is implementable and auditable (INV-009a's *observability half*); single-currency-as-default remains non-canonical | Any currency choice or sink rate; INV-009a's sufficiency threshold needs authorized numerics |
| **OQ-010** durability/repair | E-06 (repair-adjacent service XP sites); item-condition absence noted in E-05 | Requirement framing for INV-010's (a)–(d) | Whether decay exists at all; no durability evidence exists (gap, §3) |
| **OQ-011** acquisition/respecialisation | E-01/E-02 phase-2 (trainer/XP acquisition working); E-17 (drop-based respec, HISTORICAL); INV-011 derivation | Acquisition via trainers+pools is implementable; drop-cascade respec is implementable without identity loss | TCIndustries acquisition model; use-based SP remains non-canonical; cost boundedness (INV-011b) needs design first |
| **OQ-012** multi-account implementation | E-08 (anomaly lists detecting wash-trade/churn patterns — demonstrated); E-12/E-13 (multi-account adversarial packages that broke three architectures) | Coordination-detection is implementable; multi-account neutralisation power is empirically severe (EIC record) | Implementation policy (the human decision); that any new architecture survives it |
| **OQ-013** organisation property | E-02 phase-7 (org-adjacent guild scaffolding) | Membership/permission state machinery is implementable | Property-model rules; INV-013 state-machine audit has no model to audit yet |
| **OQ-014** species/setting | — | Nothing (DEFERRED; no evidence applicable) | Everything — correctly out of scope |
| **OQ-015** first-playable scope | E-10 (what a 10-phase build sequence costs/verifies like); E-20 | Effort-shape reference for phasing; INV-015's single-player-closure check has a worked example (the fork's own loop closure) | The scoping decision itself |
| **OQ-016** vendor mechanics | E-02 phase-5 suite + E-09 (vendor flows, sale history, purchase transactions) | Vendor/retail loops are implementable and measurable; identity attach demonstrable | Exact vendor mechanics; INV-016b (NPC-primacy guard) untested — no NPC-economy competition exists in any build |
| **OQ-017** provenance depth | E-09 (`crafted_items` provenance inside purchase tx); E-04 crafted-share; E-15/E-16 | Creator-granularity provenance is implementable and transaction-safe (INV-017 granularity 1) | Visibility rules; org/resource-origin granularities unimplemented anywhere (gap, §3) |

---

## 3. Evidence gap list (needed but absent — feeds the T-programme)

1. **Transportation (OQ-006):** zero evidence in any artifact. Any transport
   design proceeds blind without a simulation (candidate task, currently
   unapproved in the gap-analysis proposal).
2. **Durability (OQ-010):** no durability/decay system exists in any build.
   INV-010 is untestable until either a design or a sim host exists.
3. **Provenance granularities 2–3 (OQ-017):** organisation- and resource-origin
   provenance unimplemented; only creator granularity evidenced.
4. **Heal-kind interdependence event:** code-path-verified only; never observed
   live end-to-end (needs incap→revive scenario — flagged in the Phase 10
   verification report and the post-Phase-10 proposal).
5. **NPC-economy competition (INV-016b):** no build contains NPC supply
   competing with player supply; ECO-001's core boundary is therefore untested
   behaviorally.
6. **Long-run economy drift:** all telemetry observations are short-window
   (single session scale). Faucet/sink balance over extended horizons
   (INV-009a's domain) is NOT-MEASURABLE with current evidence.
7. **City governance dynamics (INV-007a/b):** civic scaffolding exists; no
   governance/taxation mechanics exist to test.
8. **Actor-complete adversarial analysis for non-EIC systems (INV-001 etc.):**
   the multi-account/organisation neutralisation methodology exists (E-12/E-13)
   but has been applied only to EIC architecture candidates, not to budget/
   automation/vendor shapes (T-03/T-06 would apply it; unapproved).

---

## 4. Reconciliation statement (GDD §34 step 0 closure, scoped)

The §34 step-0 requirement — *"Inspect any existing prototype; produce GDD ↔
prototype ↔ decision evidence matrix"* — is executed here **within the approved
task scope**: the inspected prototype corpus is the Pre-CU SWG testbed (the only
implementation in the repository), inventoried in §1, mapped in §2, gap-listed
in §3. Two scoping notes for the record:

1. **PROT-002's Seed-2.1 prototype** is inspected as E-15 (identity confirmed
   2026-09-12): it evidences crafting-core implementability only; its content is
   SWG-specific and non-transferable (SWG-002).
2. **The matrix is one-directional by design:** GDD → (informs via) evidence.
   No evidence entry feeds *back* into GDD status anywhere in this register;
   such promotion remains a reserved human act (PROT-001; AGENTS.md).

## 5. Acceptance self-audit (T-02 acceptance criteria)

| Criterion | Result |
|---|---|
| Every cited artifact verified to exist at cited location | ✅ all §1 entries located on the live tree 2026-09-17 (testbed `cmd/`, `docs/`, `internal/`; `investigations/eic/` 12 files; `sources/swg-pre-cu/` 3 demo trees; governance docs) |
| Every mapping states what the evidence does NOT establish | ✅ explicit "Does NOT establish" column for all 16 applicable OQ rows (OQ-014 recorded as out of scope) |
| OQ/§34 citations exact | ✅ OQ IDs match GDD §31 and the Open Questions Register v1.1; §34 step-0 language quoted from GDD line 898 region |
| Status discipline intact | ✅ EVIDENCE status only; no promotion; HISTORICAL/PROTOTYPE walls explicit; EIC chain cited without opening HD-EIC-07 |

*Verification note: auditable by spot-checking any §1 artifact path against the
repository and any §2 row against the cited GDD sections.*

---

*End of Design Evidence Register v0.1. This register is a maintained finding
aid: new evidence artifacts should be appended with IDs continuing the E-series.*
