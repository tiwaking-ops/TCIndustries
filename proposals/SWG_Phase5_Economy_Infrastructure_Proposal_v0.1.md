# SWG Pre-CU Phase 5 — Economy Infrastructure: Proposal Plan v0.1

**Status:** PROPOSED plan. No authority. No design decision. Human approval required
before any implementation, and separately before any numeric value is treated as tuned.
**Author / Assessor:** Muse Spark (`opencode/muse-spark-1.3-contributor-free`) via OpenCode —
research and planning only; all authority decisions remain human.
**Date drafted:** 2026-09-14.
**Controlling references:** `AGENTS.md`; `proposals/SWG_Phase4_Resources_Crafting_Proposal_v0.1.md`
(executed — review gate pending); `testbed/swg-phase3-combat/HYGIENE_NOTE.md`;
`investigations/SWG_Code_Due_Diligence_Evidence_Log_2026-09-14.md` (§6 row 8);
`governance/TCIndustries_Testbed_Reuse_Human_Ruling_HD-TST-01_2026-09-14.md`;
predecessor GDD `sources/swg-pre-cu/SWG_PreCU_GDD.md` (§§12, 13, 21, 22, 29.5, 30 —
HISTORICAL, never TCIndustries canon). TCIndustries fence: ECO-001/VIS-003 (player-driven
economy, LOCKED) are the direction; OQ-009/OQ-012/OQ-016 stay TBD; EIC untouched
(HD-EIC-07) — this plan builds transfer infrastructure + ledger observability ONLY, never
faucet/sink rate design beyond GDD-given values.

---

## 1. Objective

Implement **Phase 5 — Economy Infrastructure** on the testbed fork (Phases 0–4 complete
there, subject to the pending Phase 4 review gate), satisfying the predecessor-GDD §30
exit criteria verbatim: *"A player can place a house, place a vendor, list a crafted item,
and have a second player find and buy it via Bazaar search — with correct atomic
credit/item transfer (Section 29.5)."*

Entry criteria (§30: Phase 4 complete — "needs crafted items to sell") are MET in the
fork (sidearm/plate/deed craft paths verified), conditional on the review gate — see §10.

## 2. Research basis

### 2.1 Predecessor-GDD map (HISTORICAL input, not authority)

| GDD section | Specifies | Phase-5 relevance |
|---|---|---|
| 12.2 faucets | Mission 500–10k (PRIMARY), NPC junk buy-back 10–100 (destroyed, never resold), faction payouts, quest one-offs; bounty/treasure/city-salary secondary rules | Creature credit drops implementable now (10–200 CL-scaled, GDD band); missions/NPC-buy-back need unbuilt systems → OUT |
| 12.3 sinks | Maintenance (largest sustained), training costs (10–50 early / 5k–20k Master), travel, bazaar/vendor fees, repair, deed costs, faction bases; recurring-not-one-time; wealth-tier scaling | Existing sinks ledger-tagged unchanged; vendor/house maintenance + listing fees new |
| 12.4 inflation | 5–10% annual band, 30-day rolling window; nightly telemetry; lever = raise sinks not cut faucets; P2P excluded from faucet/sink accounting | Snapshot table + recording (MVP); no dashboard, no lever-pulling |
| 12.6 entities | CreditLedgerEntry (category faucet/sink/transfer + counterparty + location); EconomicSnapshot (circulation, 30d sums, inflation %, gini, top prices) | Ledger + snapshot shapes (gini/dashboard EXPANSION) |
| 12.8 edge cases | Atomicity mandate; 180-day [ASSUMPTION] reclamation; starter bootstrapping; cartels monitored-not-prevented | Atomic primitive; reclamation states (untestable durations flagged) |
| 12.9 priorities | MVP: ledger, primary faucets/sinks, basic telemetry, structure maintenance | Scope backbone |
| 13 housing | Deed→place→60–120 s construction [ASSUMPTION]→owner; tiers S/M/L (+guild hall EXPANSION); interiors/cells; permissions; maintenance + 14-day [ASSUMPTION] condemned grace → destroy; edge cases (placement races, redeed-blocking) | House MVP minus interiors/decoration/cells (no client geometry phase) |
| 13.3 entities | Structure base (location, condition, pools, costs, status, lists) + House extension + PlacedItem | Persistence shapes |
| 21 vendors | Deed→place (house-adjacent typical) → stock/price → offline sales → till (in-person collect); maintenance scales with listings; 50-slot cap [ASSUMPTION]; Merchant −50% (no Merchant profession exists — recorded gap); closed-not-destroyed on lapse; row-lock + atomicity edge cases | Vendor MVP (remote mgmt/appearance EXPANSION) |
| 22 bazaar | Terminals index ALL vendor listings galaxy-wide; filters (category, resource stats, price, planet/city, seller); stalls = escrow consignment w/ fee + 7-day expiry + MAIL delivery; 4–6% commission [ASSUMPTION]; last-N-sales MVP | Search MVP over vendor listings; stalls default OUT (needs mail — Phase 7) |
| 29.5 integrity | Single-atomic-DB-transaction mandate for every credit/item move; server-side rolls; server-side rate limiting | E2 transfer primitive; fork already has transaction precedent (training) |

### 2.2 Existing implementations (evidence, not authority)

- **Track A/B**: `economy_ledger.py` / `vendor_system.py` one-line stubs; `economy_inflation_sim`
  linear toy (evidence §3.2 — NOT an economy basis); TDR-002 one-line "Accepted"; schema
  comment names Ledger/Structure/City without defining them. Nothing transfers except the
  toy's warning value.
- **Fork wallet** (the entire current economy): 5000-start balances, insufficiency-guarded
  `DeductCredits`, in-transaction training-cost deduction (ACID precedent for §29.5),
  wallet→harvester-pool funding. To be ledger-tagged, not re-tuned.
- **Evidence gap row 8**: balance fields + train-cost sink + linear toy "do not approach
  ECO-001 needs" — this phase is the first approach.

## 3. Base and working location (RECOMMENDED — human confirms)

Build in place on `testbed/swg-phase3-combat/` (new `internal/economy/` + DB extensions +
handlers + `cmd/phase5test/`); no new fork, `sources/` untouched. Rationale: every Phase 5
flow terminates in Phase ≤4 systems (crafted items, stacks, HAM-adjacent wallet, kill path
for drop faucet).

## 4. Scope IN (work packages, each verify-by-running before the next)

- **E1 — Credit ledger** (§12.6 shape): append-only entries on EVERY credit mutation
  (existing training/harvester-fee paths retrofitted with categories — rates UNCHANGED,
  no tuning); categories faucet/sink/transfer + counterparty + zone coords.
- **E2 — Atomic transfer primitive** (§§29.5/12.8/21.5): single-transaction
  credit↔item moves with row-level serialization (BEGIN IMMEDIATE); the one way to move
  value, used by all packages below. Includes the mandated double-purchase test
  (two buyers, one unit → exactly one succeeds, loser never charged).
- **E3 — Houses MVP**: structure-deed schematic → place (valid terrain = in-zone,
  outside no-build radius, provisional 20 m — GDD silent) → owner + permission tiers
  (owner/admin/friends/banned, data only — no interior cells, no decoration, no geometry);
  maintenance pool + weekly upkeep; condemned-grace → destroy state machine; redeed
  blocked while vendor stocked. Tiers S/M/L (guild hall EXPANSION). Storage: DEFERRED
  (exit criteria needs none; §9 confirms).
- **E4 — Vendors MVP**: vendor-deed schematic → place (standalone allowed; house-adjacent
  typical, no requirement) → stock/price/description → in-person purchase by any player
  → till (in-person collect, GDD-conformant) → maintenance scaling with listings →
  lapse = closed (stock safe) per 21.2.2. Slot cap 50 (GDD [ASSUMPTION]). Merchant
  hooks ABSENT (no Merchant profession — recorded, not invented).
- **E5a — Bazaar search MVP**: one seeded terminal at zone spawn (documented stand-in for
  "terminals in every city"); galaxy-wide (single-zone) index over live vendor listings;
  filters category/price/seller/resource-stat-range; last-N-sales view (MVP per 22.2.4).
  This PLUS E4 satisfies the exit criteria without stalls or mail.
- **E6 — Creature-drop faucet**: 10–200 credits per kill scaled within the GDD band by
  CLMax (provisional mapping inside GDD-given bounds, flagged); ledger-tagged faucet.
  NPC buy-back, missions, bounties, salaries: OUT (need unbuilt systems).
- **E7 — Telemetry MVP**: snapshot table (circulation, faucet/sink sums, net % over the
  literal 30-day window — correct semantics, sparse in tests) + per-tick aggregation
  hook; test asserts recording + snapshot existence, never 30-day math. No dashboard,
  no lever-pulling, no gini (EXPANSION).

## 5. Scope OUT (explicitly deferred, with homes)

Stalls/escrow (needs Phase 7 mail — E5b default OUT); direct trade (not in exit criteria,
follow-up candidate reusing E2); house storage/interiors/decoration/cells; guild halls;
cities/tax/mayors (Phase 7); power/factory economics (R8 stayed OUT); travel/shuttle
sinks; repair/reload; NPC buy-back; missions/bounties; price-setting of any kind
(11.6/22 never hardcoded — emergent only); anomaly detection (29.6, tooling EXPANSION);
wash-trade monitoring. **GDD-internal tension recorded (human resolves):** §12.9 MVP
demands "all primary faucets" while missions/NPC-sales need unbuilt phases — the plan
implements the faucet SUBSET reachable now (E6 + existing training/harvester sinks as
sinks, not faucets) and records the remainder as blocked, not silently dropped.

## 6. Numeric authorization requests (values NOT set by this proposal)

GDD-GIVEN (implement as written): faucet bands (500–10k missions [unbuilt], 10–100 junk
[unbuilt], 10–200 kill drops [E6], 200–1000 bounty floor [unbuilt]); sink categories +
training bands (existing fork values LEFT UNTOUCHED — predecessor content, not re-tuned);
weekly maintenance SHAPE + wealth-tier scaling rule; 4–6% commission [ASSUMPTION, stalls
only — dormant]; 50-slot cap [ASSUMPTION]; house tiers/storage figures (tiers used,
storage deferred); 14-day grace + 180-day reclaim [ASSUMPTION]s; 30-day window + 5–10%
band (observed, never steered); 60–120 s construction [ASSUMPTION].
GDD-SILENT / INDEPENDENT (each needs approval OR interim-[ASSUMPTION] sign-off):
(1) structure-deed + vendor-deed schematic costs/stats; (2) house upkeep values by tier;
(3) vendor upkeep base + per-listing curve; (4) stall-style listing fee — only if E5b
flipped IN (default OUT); (5) creature-drop→CL mapping inside the 10–200 band;
(6) no-build radius (recommended 20 m); (7) seeded-terminal placement/model;
(8) snapshot cadence + retention. PROPOSED interim convention (approval requested):
Group-C-style NON-CANONICAL flagging on every non-GDD number (code comments +
HYGIENE_NOTE.md), reusing fork-local magnitudes (training-cost scale, harvester-fee
scale) where a magnitude must exist to avoid inventing fresh tuning.

## 7. Acceptance criteria (verify, don't claim)

`go build ./...` + `go vet ./...` clean; fresh DB; phase0–4 tests green PLUS new
`phase5test` end-to-end with TWO live clients (A: place house → place vendor → list
crafted sidearm; B: bazaar search finds it → buys → atomicity asserts: A till +=
price, B wallet −= price, ownership moved, ledger shows faucet/sink/transfer triple,
snapshot row exists; double-purchase race → exactly one winner; lapse vendor →
closed-not-destroyed; house lapse → condemned flag under fast-cycle mapping) with real
pasted output; fast-cycle mapping test-config-only, defaults GDD-given; hygiene note
extended; commit per verified package (E1→E7).

## 8. Risks and dependencies

- Maintenance-window durations (14 d grace, 180 d reclaim, weekly fees) are untestable
  live: fast-cycle mapping covers grace (→ minutes) the same way Phase 4 covered
  lifespans; reclaim stays code-only (assert states, not the 180-day wait).
- Vendor-without-Merchant: the GDD's distribution-role economics assume Merchant
  discounts that cannot exist yet — vendors will be mechanically complete but
  economically unrepresentative; recorded, not patched over.
- No-mail constraint shapes E5 (search-only bazaar) — any stall/mail follow-up is a
  Phase 7 dependency, not a Phase 5 stretch goal.
- Ordering: Phase 4 review gate (now requested) MUST pass before E1 starts; EIC,
  provenance, reputation untouched throughout.

## 9. Decisions requested of the human (the plan asks; it does not decide)

1. Base/location (§3) — confirm building on `testbed/swg-phase3-combat/`.
2. House storage (§4/E3) — deferred (recommended) vs minimal chest MVP.
3. Stalls E5b — OUT (recommended) vs IN with mail-substitute design.
4. Deed schematics + upkeep/fee values (§6) — approve magnitudes or supply.
5. Creature-drop mapping (§6.5) — approve in-band mapping or supply.
6. Snapshot/terminal/radius conventions (§6.6–6.8) — approve or supply.
7. Faucet-subset honesty (§5 tension) — confirm recording remainder as blocked.
8. Authorisation to implement on approval of 1–7 (cf. §10).

## 10. Implementation status

APPROVED AND EXECUTED 2026-09-14/15. Owner decisions (§9) received: base confirmed,
storage deferred, stalls OUT, magnitudes approved, drop mapping approved, conventions
approved, faucet-subset recorded as blocked, implementation AUTHORIZED (also covering
the Phase 4 review-gate precondition for this line — owner may still call the gate
separately).
Results on `testbed/swg-phase3-combat/`: `go build` + `go vet` clean; fresh-DB full
suite testclient 9/9 + phase1test 7/7 + phase2test 11/11 + phase3test 6/6 +
phase4test 9/9 + phase5test 9/9, ALL PASS (real output; OVERALL=0).
Observed: buyer −500/seller +500 atomic with receipt; double-buy race exactly one
winner, loser never charged; closed vendor refuses stocking; unfunded house runs
condemned → destroyed under fast-cycle grace; snapshot with live faucet/sink sums
(e.g. net −0.13%); drop faucet +10 with ledger entry. Bugs found by running:
stale-receipt test bug (vendor read house receipt), DB-persist cadence vs terminal
rule (sync-nudge helper), snapshot key case, phase4 trio-count assertion (now
presence-based), harvester-deed quota corruption by a reconstructive edit
(caught by suite, reverted to verified 8/4), phase3 posture/incap race (posture
moved before lair exposure + fast kill-poll). HYGIENE_NOTE.md extended (§Phase 5
additions + correction log); species-ID + Godot known-remainings unchanged.
Next: human review gate before any Phase-6-equivalent work.

*End of proposal v0.1. To enact: owner approves (or amends) §§3–6/9–10; approval and date
recorded before any code is written.*
