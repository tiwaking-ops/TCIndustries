# Project Documentation Report — SWG Phase 10 Session (Balance, Telemetry & Polish)

| Field | Value |
|---|---|
| **Document ID** | TCIND-PDR-2026-09-17-01 |
| **Title** | Complete Results of the Phase 10 Implementation Session — Testbed Status, Completeness, and TCIndustries Reusability Assessment |
| **Author** | Buffy — Codebuff coding agent (Freebuff client). LLM: **GLM, version 4.6**, trained by Z.ai |
| **Assessor / Recipient** | OpenCode — project documentarian per `AGENTS.md` |
| **Date** | 2026-09-17 |
| **Status** | EVIDENCE (observation report per `AGENTS.md` status table; no design authority claimed, no status promotion performed) |
| **Distribution** | Project owner, OpenCode |
| **Authority basis** | Owner authorization of `proposals/SWG_Phase10_Balance_Telemetry_Polish_Proposal_v0.2.md` §9/§6, recorded 2026-09-17 in that file's §10 |
| **Related documents** | `proposals/SWG_Phase10_Balance_Telemetry_Polish_Proposal_v0.2.md` · `investigations/SWG_Phase10_Implementation_Verification_Report_2026-09-17.md` · `canonical/gdd/TCIndustries_Master_GDD_v1.1.1_Status_Patch.md` · predecessor `SWG_PreCU_GDD.md` (HISTORICAL) |

---

## 1. Executive summary

On 2026-09-17 a single working session (a) produced the research-verified Phase 10 proposal v0.2, (b) received and recorded owner authorization for its seven decision packages (B1–B7), (c) implemented all seven packages in the Phase 3 combat testbed (`testbed/swg-phase3-combat/swg-server/`), and (d) verified the result with static gates, a five-proxy success-criteria harness, and a full-chain regression across Phases 1–10 — all green. Seven commits (`f9852ec`…`2730c9a`) were made under a mixed-provenance discipline that committed only Phase 10 work and disclosed every shared file.

The testbed as a whole is assessed a **Functioning Prototype** (engineering-verified through Phase 10; not a product). For the TCIndustries project — an original-IP design effort — the session's output is reusable in three tiers: **directly usable** non-gameplay infrastructure (~30% of the new material), **adaptation-required** mechanics and evidence (~50%), and **not usable** SWG-derived IP, reference material, and unauthorized numerics (~20%). No design document status changed; BAL-001 and HD-TST-01 constraints were honored and are provable from the commit range.

---

## 2. Scope of the documented session

The session covered the following work items, in order:

1. **Phase 10 proposal v0.2** (`proposals/SWG_Phase10_Balance_Telemetry_Polish_Proposal_v0.2.md`): comprehensive research pass over the predecessor GDD (§1.3, §9.7, §12.4, §29.6, §29.8, §30.2, §33, §34), Master GDD v1.1.1 statuses, governance decisions (HD-TST-01, HD-EIC-05/07, HD-GDD-01), Phase 9 §10 decision 6c (service-XP residual routed to Phase 10), and a first-hand testbed source inventory. v0.1 retained unaltered as provenance. Status: PROPOSED.
2. **Owner authorization** recorded in proposal §10: B1–B3, B5(i), B6, B7 in full; B4 as characterization + options paper only; zero numeric changes anywhere.
3. **Implementation of packages B1–B7** (§3 below), committed per verified package.
4. **Verification** (§4 below), including repair of two pre-existing defects discovered by the regression chain.
5. **Reporting**: implementation verification report filed under `investigations/`; proposal §10 updated with an EXECUTED record; README phase table refreshed with evidence-dated checkmarks.

Out of scope, and deliberately untouched: numeric tuning of any gameplay value, new residual channels, the ~69 pre-existing dirty/untracked working-tree paths, and adoption of the remainder of the testbed into git (flagged as owner work).

---

## 3. Delivered work (per package, with artifact inventory)

| Package | Delivered | Primary artifacts |
|---|---|---|
| **B1 — Telemetry completion** | Append-only `interdependence_events` written best-effort at three existing hook points (wound-heal, buff receipt, crafted-item purchase with `crafted_items` provenance check inside the purchase transaction); read-only REST `GET /api/phase10/telemetry|anomalies|jobs`; §29.6 nightly economic-snapshot job as a §29.8-shaped `phase10_jobs` record (single-row claim, idempotent completion, failure recorded), cadence 24 h → 60 s under `TESTBED_FAST_CYCLE=1`, separate from the real-time world loop; queries for profession distribution (registry-joined), mission-outcome mix, wealth concentration (Gini, diagnostic only), and price volatility from the existing append-only `market_records`. | `internal/database/phase10_db.go`, `internal/handlers/phase10_handlers.go`, hook insertions in `internal/handlers/services.go`, `internal/database/economy_db.go`, route wiring in `internal/server/server.go` |
| **B2 — §1.3 proxy harness** | Five success-criteria proxies with three verdict forms (PASS / FAIL-with-telemetry / NOT-MEASURABLE), raw counts beside every percentage, real-resolution PvP progression proxy, env-overridable server URL (`P10_SERVER_URL`). | `cmd/phase10test/main.go` |
| **B3 — Balance validation report** | 13-band generator producing OBSERVED / CODE-CONSTANT-cited / DERIVED / NOT-MEASURABLE rows; committed report records the fork's unarmed damage 5–15 as **OUT-OF-BAND** against §9.7's 50–150 start band (evidence, not a fix) and drop/reward constructions as IN-BAND. | `cmd/phase10report/main.go`, `docs/phase10_balance_report.md` |
| **B4 — Service-XP residual** | All flat `[PROVISIONAL]` service-XP sites characterized with file:line citations (medical 50/50/25/100; entertaining 25; image designer 50×2; smuggler 50; stim 25; scouting 10/25/25; ranger 10/25; pets 25/100/50; mentoring 50+50; merchant price/100-shaped; structure 100×4); options paper holds §7.2.3 redesign shapes (per-point wounds, per-tick battle-fatigue, per-application buff, diminishing-returns hook) as explicit numeric asks. **No rate changed; provisionals stand.** | B4 section of `docs/phase10_balance_report.md`, `docs/phase10_service_xp_options.md` |
| **B5(i) — Anomaly-lite** | Ranked human-review lists (wash-trade pairs, near-zero-net churn, placement bursts, price spread) using structural invariants only — no thresholds, no sigma, never auto-actioned. | In `internal/handlers/phase10_handlers.go` |
| **B6 — Polish / stability** | COMBAT-DEBUG logs gated behind `TESTBED_COMBAT_DEBUG=1`; all 30 `UnixNano` ID-mint sites swept onto a collision-proof exported helper (`newRowID`, RNG seeds deliberately untouched); dead `MailTimestamp` removed; `.gitignore` source-dir negations; dedicated-port suite runner; README phase table refreshed. | `internal/handlers/combat.go` + 12 further files; `run_phase10.sh`; `.gitignore`; `README.md` |
| **B7 — Full-chain regression** | Chain runner: per-phase fresh DB, own server lifecycle, inter-phase listener reaping; exposed and fixed two pre-existing defects (below). | `run_chain.sh`; assertion refresh in `cmd/phase2test/main.go` |

**Commit inventory (7 commits, this session):** `f9852ec` (B1+B5), `55ec433` (B2), `12ca3db` (B3+B4), `ce09b35` (B6 runners/gitignore), `458c311` (B6 sweep + gate), `e6c1187` (B3/B4 committed evidence), `2730c9a` (B7 + phase2test refresh). Mixed-provenance cases (pre-existing untracked content inside files also edited by this session) are disclosed in the respective commit messages per the owner's commit-scope instruction.

---

## 4. Verified results (real output, this host)

Static gates: `go build ./...`, `go vet ./...`, `go test ./...` — all clean across every package.

Full-chain regression, phases 1–9 each on a fresh DB and own server lifecycle, then the Phase 10 suite:

```
=== ALL PHASE 1..9 TESTS PASSED ===   (phases 1 through 9, individually)
=== CHAIN RESULTS: p1=0 p2=0 p3=0 p4=0 p5=0 p6=0 p7=0 p8=0 p9=0 ===
CHAIN:ALL-GREEN
```

Phase 10 suite (final standalone run of the same code state):

```
[PASS] Proxy 1: Economic Interdependence — combat cohort 2/2 covered via receipts
       (buffF=true buffE=true craftedPurchase=true; healF=false, see note); all-world 3/4 (75%)
[PASS] Proxy 2: Horizontal Progression — novice-vs-+tier1 damage-per-swing ratio ≈ 0.93×
[PASS] Proxy 3: Player-Driven Economy — crafted_share = 1/1 (100%)
[NOT-MEASURABLE] Proxy 4: Sandbox Validation — qualitative by definition; artifacts staged for human review
[PASS] Proxy 5: Mission-economy telemetry — lifecycle mix observable (29 rows)
[PASS] Telemetry job completed with §29.8-shaped result payload
=== PHASE10_EXIT:0 ===
```

Honest-evidence notes (recorded, not smoothed over): the wound-heal leg is durably unexercisable without an incap→revive cycle (the server refuses with "target has no wounds"; surfaced, not swallowed), so the `heal` event kind is code-path-verified while buff and crafted-purchase receipts were observed live. Proxy 2 reports the observed ratio consistent with the fork's tier structure not yet feeding damage mods into resolution; the six-month horizon form remains NOT-MEASURABLE by construction.

Defects the chain exposed and this session fixed in pre-existing code: (1) a stale Phase 2 registry assertion ("6 basic professions", registry is 28 post-Phase 9); (2) MSYS/Windows `kill` non-reaping causing silent port-8080 bind failures between phases (addressed with listener reaping and a dedicated Phase 10 port).

**Compliance proof:** `git diff 19e714a..HEAD` across gameplay-constant packages (`internal/services`, `internal/skills`, `internal/combat`, `internal/economy`) contains zero changed numeric constants; gameplay files received only best-effort telemetry writes, ID-mint replacements, and the log gate; nothing reads telemetry to steer behavior; all new REST routes are GET-only.

---

## 5. Project completeness assessment

**Assessment: Functioning Prototype — engineering-verified through Phase 10; pre-MVP as a product.**

Rationale, positive:

- All ten phase scopes (0–9 plus the Phase 10 polish/telemetry layer) are implemented and pass end-to-end integration suites on a fresh database, on Windows, in one contiguous chain run (2026-09-17).
- Core loops are real, not stubbed: movement/spatial chat, skills/XP/training, resources/crafting, economy/vendors, services/heals, faction/PvP, elite professions, missions with escrow, lairs, and now telemetry/anomaly/report machinery.
- Exit criteria were run, not asserted: three proxies PASS with data, one honestly NOT-MEASURABLE with staged artifacts, one human-pass-only by rule; FAIL-with-telemetry paths are implemented and were demonstrated during shakeout.

Rationale, limiting (why not "Fully Functional" and not yet an MVP):

- Single-host SQLite persistence; no deployment, operations, or multi-server story.
- The Godot client is an early shell (login/character-select/world scenes); the integration evidence is driven by Go test clients.
- Breadth is intentionally narrow: it is a Phase 3 combat testbed that accreted Phases 4–10, not a content-complete world.
- Known open items: phases 1–9 test clients hardcode port 8080; `missions.go` lifecycle statuses do not feed B1 counters (row-level mix only); the heal-event end-to-end gap above; all service-XP rates remain `[PROVISIONAL]` pending a numeric authorization that has not been given; ~69 pre-existing working-tree paths (including the `db.go` HAM repair and most Phase 3–9 sources) remain uncommitted pending an owner "adopt the testbed into git" pass.

Per `AGENTS.md`, everything the testbed demonstrates carries status **PROTOTYPE / EVIDENCE** only.

---

## 6. Reusability assessment for TCIndustries

TCIndustries is an original-IP design project governed by `AGENTS.md` authority rules; the testbed is a SWG Pre-CU recreation fork whose purpose is to produce evidence for that design. The split below follows from that relationship.

### 6.1 Directly usable as-is (~30% of session output)

Non-gameplay, IP-neutral engineering that can be lifted wholesale:

- Telemetry infrastructure: `interdependence_events` schema and writers, `phase10_jobs` §29.8-shaped job record (single-row claim, idempotent completion), read-only telemetry/anomaly/job REST pattern, snapshot and volatility/Gini query set.
- `cmd/phase10report` report-generator pattern: OBSERVED / CODE-CONSTANT / DERIVED / NOT-MEASURABLE row taxonomy with citations — directly reusable for any TCIndustries balance-audit tooling.
- `cmd/phase10test` harness conventions: three-verdict proxy design, raw-counts-beside-percentages, fresh-DB lifecycle, env-overridable endpoints.
- `newRowID` collision-proof ID mint and the portability sweep pattern; `run_chain.sh` / `run_phase10.sh` orchestration (fresh DB per phase, listener reaping, dedicated ports).
- Documentation and governance artifacts: the proposal→authorization→execution→verification paper trail, the mixed-provenance commit discipline, and both 2026-09-17 reports as templates.
- `phase10_service_xp_options.md` as a *process* template for options papers (structure only — its contents are fork-specific).

### 6.2 Adaptable with design work and human authorization (~50%)

Mechanics whose structure transfers but whose values and semantics require TCIndustries-specific derivation:

- Combat resolution, HAM/wounds state, services/heals, skills-and-boxes progression model, elite-gate registry structure.
- Economy: ledger faucet/sink tagging, `market_records` sale history, vendor flows, crafting pipelines, resource spawns/surveying.
- Missions/contracts (including bounty escrow), lair population/regen/relocation, faction/PvP flagging.
- The interdependence-event evidence itself: the observed buff/crafted-purchase receipts and the B4 characterization are direct inputs to the deferred Economic Interdependence Core (EIC) work — subject to HD-EIC-05/07 gating (no new shaping without human-supplied structural boundaries).

Adaptation requirements, all binding: every numeric is SWG-derived and therefore needs re-derivation against the Master GDD plus explicit numeric authorization (HD-TST-01; no numeric tuning without it); structural reuse must pass through TCIndustries design authority, since prototype behavior is evidence, never canon.

### 6.3 Not usable for TCIndustries (~20%)

- **SWG-specific IP surface:** SWG branding, planet/NPC/profession-flavor naming, and any content identifying the fork as a Star Wars Galaxies recreation. TCIndustries is original IP; none of this may carry over.
- **`sources/swg-pre-cu/` reference material:** HISTORICAL inspirational/provenance material with licensing sensitivity; reference and evidence only, never transferable content.
- **Predecessor-GDD balance bands (§9.7, §33, §34):** HISTORICAL comparison baselines. They may contextualize observations (as B3 does) but must never become tuning targets or canonical TCIndustries values.
- **All `[PROVISIONAL]` rates and fork constants as final values:** explicitly unauthorized; adopting any of them would violate the governance boundary against numeric tuning without human authorization.
- **Placeholder mechanics:** the Phase 2 `earn-xp` placeholder action and other scaffolding whose only purpose was to exercise testbed plumbing.
- **Early Godot client scenes/scripts:** too incomplete to reuse as a client foundation for TCIndustries.

---

## 7. Constraints compliance and status discipline

- BAL-001 (observe-don't-steer): honored — telemetry is write-only from gameplay, read-only via REST, and nothing reads it to change behavior.
- HD-TST-01 (patterns-only fork): honored — zero numeric changes, proven by the cited diff scope.
- No new residual channels; no automatic successor architectures; provenance/reputation remain DEFERRED; EIC untouched (HD-EIC-05/07).
- Status discipline: proposal v0.2 remains PROPOSED with an EXECUTED record added in §10; no document was promoted to LOCKED; this report claims EVIDENCE status only.

## 8. Open items handed forward

1. Owner pass: adopt remaining pre-existing testbed sources into git so future verification diffs are complete (itemized in `investigations/SWG_Phase10_Implementation_Verification_Report_2026-09-17.md` §4).
2. Close the heal-event evidence gap with an incap→revive interdependence scenario.
3. Give phase 1–9 test clients env-overridable server URLs (the Phase 10 client already has `P10_SERVER_URL`).
4. Owner decision (B1-adjacent): whether mission rows should be retained rather than deleted to feed per-terminal completion telemetry.
5. B4 numeric decision: adopt, decline, or defer the options paper — default is that the flat provisionals stand.

---

*Prepared by Buffy (Codebuff/Freebuff coding agent; LLM: GLM 4.6, Z.ai), 2026-09-17. This report documents what exists and what was verified; it invents no decisions and changes no design authority.*
