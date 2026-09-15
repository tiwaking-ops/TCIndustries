# SWG Pre-CU Phase 7 — Civic Systems: Proposal Plan v0.1

**Status:** PROPOSED plan. No authority. No design decision. Human approval required
before any implementation, and separately before any numeric value is treated as tuned.
**Author / Assessor:** Muse Spark (`opencode/muse-spark-1.3-contributor-free`) via OpenCode —
research and planning only; all authority decisions remain human.
**Date drafted:** 2026-09-15.
**Controlling references:** `AGENTS.md`; `proposals/SWG_Phase5_Economy_Infrastructure_Proposal_v0.1.md`
(executed — review gate pending); `proposals/SWG_Phase6_Social_Support_Proposal_v0.1.md`
(executed — review gate pending); `testbed/swg-phase3-combat/HYGIENE_NOTE.md`;
`investigations/SWG_Code_Due_Diligence_Evidence_Log_2026-09-14.md`;
`governance/TCIndustries_Testbed_Reuse_Human_Ruling_HD-TST-01_2026-09-14.md`;
predecessor GDD `sources/swg-pre-cu/SWG_PreCU_GDD.md` (§§12, 13, 14, 18, 19, 20, 28, 29.5,
30, 31 — HISTORICAL, never TCIndustries canon).
TCIndustries fence: BLD-001 (player-owned structures, LOCKED) and SOC-001 (player
organisations, LOCKED) are the direction; BLD-003/SOC-002/SOC-003 (city governance,
organisation capabilities, social institutions) stay PROPOSED/TBD; OQ-007 (city
formation/governance/taxation) and OQ-013 (organisation rights/hierarchy/property) stay
TBD; EIC untouched (HD-EIC-05 retired, HD-EIC-07 gate — no shaping here, city salaries
are transfers not faucets); provenance / formal reputation stays deferred (conveniently,
GDD 18.2.5 wants NO formal reputation score either — agreement recorded, not authority);
BAL-001 (no numeric tuning without human authorisation) governs §6.

**Change discipline (per AGENTS.md):** this is a candidate plan filed in `proposals/`
(design-proposal change type, no authority change). Authority basis for every cited rule
is labelled inline (HISTORICAL vs LOCKED vs PROPOSED/TBD vs PROTOTYPE/EVIDENCE). Nothing
here promotes SWG mechanics into TCIndustries canon; all fork behaviour remains
PROTOTYPE/EVIDENCE under HD-TST-01.

---

## 1. Objective

Implement **Phase 7 — Civic Systems** on the testbed fork (Phases 0–6 complete there,
subject to pending review gates), satisfying the predecessor-GDD §30 exit criteria
verbatim: *"A group of players can found a city, elect a mayor, and see rank-gated
unlocks (shuttleport) appear."*

Entry criteria (§30: Phase 5 complete — "needs houses/vendors/structures to count and
tax") are MET in the fork (house/vendor/deed/ledger/telemetry paths verified 9/9 in
phase5test), conditional on the review gates — see §10. Phase 6 venue-adjacency notes
("venues means placed structures/anywhere, not cantinas — no cities exist") are exactly
what this phase resolves.

Note on §30 scope wording: Phase 7 "Implements: Section 14 (City
founding/governance), Section 19.2.4 (Guilds), Section 18.2.2 (Mail), full Section 13
permission/maintenance depth." Groups (19.2.1–19.2.3), guild/group chat routing,
waypoints/friends (18.2.4), and tell/planet channels (18.2.1) have no other phase home
in §30, but MVP §31.4 demands "Groups, loot rules, Guilds, Mentorship" and "All chat
channels, mail, waypoints, friends". §4 packages G1–G3 cover that gap explicitly as
RECOMMENDED IN with rationale; the human may sever them to a follow-up without
affecting the verbatim exit criteria (§9).

## 2. Research basis

### 2.1 Predecessor-GDD map (HISTORICAL input, not authority)

| GDD section | Specifies | Phase-7 relevance |
|---|---|---|
| 14.2.1 founding | Mayor-Founder buys City Hall deed (Architect source); ≥10 structures in radius [HISTORICAL]; Rank 1 Outpost on placement; auto-enrol + opt-out | V1 core |
| 14.2.2 ranks | Rank table (Outpost→Township→City→Metropolis; 20+/40+/75+ structure thresholds, treasury conditions) [ASSUMPTION — authoritative for this project]; unlocks: zoning → Cantina/Medical → Shuttleport+trainers → Clone+specialization | V1 ranks 1–3 + V4 flags; Rank 4 default OUT (§5) |
| 14.2.3 governance | Weekly elections, any-citizen candidacy, one-vote, 7-day term [ASSUMPTION]; mayor powers (tax flat/% of local vendor sales, placement approve/deny, militia appoint, specialization Rank 4+, evict for non-payment); Politician elite (Master Entertainer/Merchant prereq [ASSUMPTION]) with maintenance/zoning/specialization perks; specialization 5-way [EXPANSION] | V2 core; Politician/specialization OUT beyond lightweight perk (§5/§9) |
| 14.2.4 maintenance | Weekly upkeep from treasury (taxes+fees); downgrade one tier on shortfall; Rank-1 failure → dissolve, structures revert to standalone (owned); rank-gated structure access removed, citizen houses NOT destroyed | V3 core |
| 14.3 entities | City / CityCitizen / CityElectionBallot shapes (UUIDs, radius, rank, treasury, upkeep, tax_rate, specialization, mayor+term, joined/tax-paid, militia flag, ballot rows) | Persistence shapes |
| 14.4 interactions | Depends on §13 (structure count); unlocks §20 travel nodes Rank 3+, §9.6.3 clone binding Rank 4; city faction-leaning [ASSUMPTION]; Politician cross-link §8 | V4 flag-only unlocks; faction-lean OUT (Phase 8) |
| 14.5 edges | No-overlap radii; 14-day [ASSUMPTION] mayor-inactivity emergency election; post-founding structure-count drops do NOT dissolve (treasury failure only); tie → earliest candidacy [ASSUMPTION] | V1/V2 rules |
| 14.6 priorities | [MVP]: founding, Ranks 1–3, elections, tax, treasury/upkeep, shuttleport unlock; Politician lightweight perk only; [EXPANSION]: Rank 4, specialization, militia tools | Scope backbone |
| 19.2.4 guilds | Registrar founding + credit cost (sink) + ≥3 members [ASSUMPTION]; roles Leader/Officer/Member; independent of city/faction; hall owned by Guild entity (13.5); optional leader-configured dues % diverted from member bazaar sales into treasury [ASSUMPTION] | V5 core |
| 19.3/19.5 guild entities+edges | Guild / MentorshipBond shapes; leader-disconnect succession (group 2-min; guild → senior Officer, else 30-day dissolution countdown); group-cap-20 merge rejection | V5 (+G1 succession) |
| 19.2.1/19.2.2 groups+loot | /invite flow, transferable leader, cap 20 [ASSUMPTION]; Round-Robin / Master-Looter / Need-Greed | G1 (§30-gap, RECOMMENDED IN — §1) |
| 19.2.5 mentorship | Veteran→newcomer bond, 7-day XP boost both sides [ASSUMPTION, MVP-tagged] | G1 RECOMMENDED IN (cheap, MVP-tagged) |
| 19.2.3 squad leader | Elite (Master Brawler/Marksman [ASSUMPTION]), command-stance group buffs smaller than Doctor/Musician, [EXPANSION] | OUT (§5) |
| 18.2.1 chat | /say 20 m (exists), shout 50 m, planet, group, guild, faction, tell | G2 (faction channel stubbed, Phase 8 owns membership) |
| 18.2.2 mail | Async character mail, offline delivery, text + credit/item attachments (crafted-economy remote sales), 50-msg cap [ASSUMPTION], oldest-read cleanup flagged, never auto-delete with unclaimed attachments | V6 core; unlocks Phase 5 E5b stalls follow-up (OUT here) |
| 18.2.4 waypoints+friends | Personal waypoints (survey/mission/manual/shared) + chat-link sharing; Friends list with online/offline | G3 RECOMMENDED IN |
| 18.2.5 reputation | NO numeric score — emergent signals only [ASSUMPTION, deliberate] | Implemented by omission (aligns with deferred provenance) |
| 18.2.3 emotes | Standard library all-characters; holoemotes need Image Designer (§8.3.20, EXPANSION) | Standard emotes data-only; holoemotes OUT |
| 13.2.3–13.2.5/13.5 depth | Interiors/cells (29.2), furniture caps, public/private entry; Owner/Admin/Friends/Banned tiers; maintenance pool + 14-day [ASSUMPTION] condemned grace → destroy (harsh, UI-warned); placement-race transaction order; owner-deletion 30-day [ASSUMPTION] reclaim; redeed blocked while stocked; hall transfers to guild entity | V7 (interiors/cells/decoration-geometry OUT — no client geometry phase) |
| 13.3 entities | Structure base + House + PlacedItem shapes (pools, costs, status, lists) | Extend, don't rebuild (fork has structure/vendor tables) |
| 12.3/12.6/21/22 money hooks | Guild-registrar cost + city upkeep + dues as sinks/transfers; city salaries recirculate tax (transfer, NOT faucet — §12 p.2186); ledger categories faucet/sink/transfer; vendor-sale tax hook for mayor % | V3/V5 ledger tagging; salaries OUT (no salaried roles exist) |
| 28/29.5 entities+integrity | City/Citizen/Ballot/Mail/Waypoint/Friends/Group/Guild shapes; single-atomic-transaction mandate for every credit/item move | E2 reuse for dues/tax/till/mail-attachment moves |
| 31 MVP/EXPANSION | City founding→Rank 3 + elections + tax; groups/loot/guilds/mentorship; all chat + mail + waypoints + friends; Politician/Squad-Leader/Image-Designer beyond perk deferred | Scope arbiter where §30 is silent |

### 2.2 Existing implementations (evidence, not authority)

- **Fork (`testbed/swg-phase3-combat/`)**: structures/houses (S/M/L tiers, 20 m no-build
  radius provisional, upkeep 200/500/1000, condemned→destroyed under fast-cycle),
  vendors (50-slot [ASSUMPTION] cap, upkeep 300+50/listing, till in-person collect,
  closed-not-destroyed on lapse), E2 atomic transfer primitive (BEGIN IMMEDIATE +
  double-purchase test precedent), ledger (faucet/sink/transfer + counterparty +
  coords), bazaar search MVP (no stalls — "needs Phase 7 mail"), wallet, deed registry
  (5 schematics). **Gaps this phase fills**: zero city/guild/mail/group/chat-beyond-say,
  zero permission enforcement (tiers data-only), zero treasury/tax/dues/election flows,
  zero waypoint/friend/tell/planet channels. No civic code to adapt — greenfield build
  on fork-local patterns.
- **Group C `code-phase0-3-claudecode/`**: trainer `city` column is a locale string only;
  no guild/mail/city logic. Nothing transfers.
- **Track A/B**: schema comment names Ledger/Structure/City without defining them; zero
  civic code. Nothing transfers.
- **Chat export / java-swg / seed+claude demos**: no civic content. Nothing transfers.

## 3. Base and working location (RECOMMENDED — human confirms)

Build in place on `testbed/swg-phase3-combat/` (new `internal/civic/` + DB extensions +
handlers + `cmd/phase7test/`); no new fork, `sources/` untouched. Rationale: every
Phase 7 flow terminates in Phase ≤5 systems (structures for counting/zoning, vendors
for sale-tax, wallet+E2+ledger for treasury/dues/tax/mail-attachment moves, HAM-adjacent
identity for citizen/member rows). Single-zone simplification carries over:
planet = `zone-0001`; radii in metres; no NPC cities (seeded spawn terminal stands in
for "terminals in every city" precedent — city unlocks are server-side flags + API,
not client geometry).

## 4. Scope IN (work packages, each verify-by-running before the next)

- **V1 — City founding + ranks 1–3 + citizenship**: City Hall deed schematic (registry
  addition, Architect-gated like house/vendor deeds — first civic deed, flagged pattern
  continuation); place (no-overlap validation vs existing city radii, outside
  provisional no-build radius reuse); ≥10 structures-in-radius check (structure kinds:
  house/harvester/vendor/civic; harvesters count per 14.2.1 "any mix"); Rank 1 on
  placement; auto-enrol citizens (structure-in-radius) + opt-out/leave; rank evaluation
  (structure count + treasury-funded-days per 14.2.2 table, Ranks 1–3 only); zoning
  approve/deny hook point (mayor decision stored, enforced at placement validation).
- **V2 — Governance**: candidacy + one-vote-per-citizen elections + weekly period +
  7-day term [ASSUMPTION]; tie → earliest candidacy [ASSUMPTION]; mayor powers: tax rate
  set (flat and/or % of local vendor sales — % collected at vendor purchase via E2),
  placement approve/deny, evict-for-non-payment; inactivity (14-day [ASSUMPTION])
  emergency election; leadership-vacancy chain (mayor null → acting-mayor = earliest
  officer-equivalent citizen? NO — cities have no officers: vacancy triggers emergency
  election, flagged simplification).
- **V3 — Treasury / upkeep / downgrade / dissolve**: treasury funded by citizen taxes +
  structure fees + vendor-sale %; weekly upkeep deduction (values §6); shortfall →
  downgrade one tier (rank-gated flags revoked, citizen property untouched); Rank-1
  failure → dissolve (structures revert standalone-owned); every movement ledger-tagged
  (transfer; city salaries NOT built — no salaried roles exist, recorded not dropped).
- **V4 — Rank-gated unlock flags (flag-only, implementations OUT)**: Rank 2 →
  `cantina_allowed` + `medical_allowed` structure-kind flags; Rank 3 → `shuttleport`
  flag + `trainer_hosting` flag (satisfies exit criteria "shuttleport appear": API-visible
  flag + seeded shuttleport marker row, no §20 travel movement); Rank 4 shapes present
  as data (rank enum + specialization column) but unreachable (cap Rank 3, §5).
- **V5 — Guilds**: registrar founding (credit-cost sink, ≥3 members [ASSUMPTION]);
  roles Leader (disband/transfer)/Officer (invite/kick/hall-permissions)/Member
  (chat/hall-per-points); guild chat; treasury + optional dues % diverted from member
  bazaar sales (E2, ledger-tagged transfer); hall ownership by Guild entity (nullable
  `guild_hall_structure_id` FK + 13.5 transfer rule — hall BUILD itself EXPANSION, OUT);
  succession (leader → senior Officer; none → 30-day dissolution countdown, states only);
  independence: guild membership / city citizenship / faction alignment tracked
  separately, co-displayed.
- **V6 — Mail**: send/list/read/claim; offline delivery (recipient need never be
  online); credit attachments via E2 atomic + item attachments via inventory/stack
  transfer (ownership leaves sender at send time — 18.5 edge made structural);
  50-cap [ASSUMPTION] with oldest-read-cleanup flagging, never deleting unclaimed
  attachments; tell-to-unknown → explicit error. Mail is also the unblocker for the
  Phase 5 E5b stalls/escrow follow-up (that follow-up is OUT here, §5).
- **V7 — Section 13 permission/maintenance depth**: enforce Owner/Admin(entry+decorate+
  storage, no redeed/transfer)/Friends(entry+granted-storage)/Banned-despite-public +
  public/friends_only/private entry flags; placement-race transaction order (first
  validated wins); owner-deletion/ban → condemned fast path (states, durations §6);
  redeed-blocked-while-stocked (exists — extend to vendor-stocked civic structures).
- **G1 — Groups + loot-rule data + mentorship (§30-gap, RECOMMENDED IN)**: /invite →
  accept/decline → shared Group (cap 20 [ASSUMPTION], merge-over-cap rejected);
  transferable leader; 2-min auto-transfer on disconnect-timeout; loot_rule enum
  (round_robin/master_looter/need_greed, round-robin effective default — no loot
  pipeline exists yet to distribute beyond corpse/credit paths, recorded); MentorshipBond
  (7-day, threshold checked at bond time only). Justification: guild/group/planet chat
  routing + Phase 8 base-assault coordination need groups; MVP §31.4 lists them.
- **G2 — Chat routing**: shout 50 m + planet + group + guild + tell (+ existing /say
  20 m); faction channel EXISTS with empty membership (Phase 8 owns alignment);
  standard emote library as data-only broadcast (no Image Designer tooling).
- **G3 — Waypoints + friends**: waypoint CRUD (survey/mission/manual/shared sources) +
  chat-link share/import (cross-zone imports with "different planet" display rule);
  friends list + online/offline status.
- **V8 — `phase7test`** (mirrors the verified 6–9-test sets of prior phases): multi-client
  founding → election → tax/treasury → rank-flag → guild → mail-offline-claim →
  permissions → dissolve-path, plus re-run of ALL Phase 0–6 tests per verify-don't-claim.
  Durations compressed ONLY via the established TESTBED_FAST_CYCLE test-config mapping
  (defaults stay GDD-given).

## 5. Scope OUT (explicitly deferred, with homes)

Rank 4 Metropolis + 5-way specialization + Politician beyond a lightweight upkeep-discount
perk (EXPANSION per 14.6/31.5; full Politician/Squad-Leader/Image-Designer are deferred
elites); militia sanctioning tools + PvP-defender semantics (needs Phase 8 flagging);
city faction-leaning + base-toleration (Phase 8, §15); shuttle/starport TRAVEL movement
(§20 — V4 sets the flag only); Rank-4 clone binding (flag only — extends Phase 3 clone
path later); stalls/escrow consignment (Phase 5 E5b follow-up reusing V6+E2, NOT this
phase); house interiors/cells/furniture-geometry/decoration counts (no client geometry
phase); guild halls BUILT (FK + ownership rule only); guild alliances; salary
disbursement; Chef/food buffs; disease/poison-adjacent civic clinics beyond flags;
reputation/star-rating (omitted per 18.2.5 + deferred provenance); space/Jedi/quests/
instancing (GDD §32, never). **GDD-internal tensions recorded (human resolves):**
(a) 14.2.1 "≥10 [HISTORICAL]" vs 14.2.2 "[ASSUMPTION]-authoritative" thresholds — plan
implements the 14.2.2 table as written (per its own header) with the 10-minimum as the
Rank-1 floor; (b) Politician-prereq [ASSUMPTION] (Entertainer-vs-Merchant mastery)
untouched while the profession is deferred — lightweight perk gated on MAYOR STATUS
only, flagged simplification; (c) §12.9-style "all-primary" expectations vs unbuilt
sinks (travel/repair) — civic sinks implemented are registrar + upkeep + dues shapes
only, remainder recorded blocked.

## 6. Numeric authorization requests (values NOT set by this proposal)

GDD-GIVEN / [ASSUMPTION]-tagged-in-GDD (implement as written, flagged where tagged):
founding ≥10 structures [HISTORICAL]; rank table 20+/40+/75+ + treasury conditions
[AUTHORITATIVE-ASSUMPTION]; 7-day term, weekly elections, 14-day inactivity trigger,
earliest-candidacy tiebreak, 50-msg mail cap, group cap 20, guild ≥3 founders, 30-day
guild dissolution countdown, 2-min group succession, 14-day house grace + 30-day
owner-deletion reclaim (V7 fast path), 60–120 s construction [ASSUMPTION] (reuse Phase 5
mapping), upkeep SHAPE (weekly, wealth-scaled) + tax SHAPE (flat and/or vendor-%).
GDD-SILENT / INDEPENDENT (each needs approval OR interim-[ASSUMPTION] sign-off):
(1) City Hall deed schematic cost/stats + city radius metres; (2) upkeep values by rank
(anchor: house 200/500/1000 + vendor 300+50/listing scale); (3) tax bounds (flat cap +
vendor-% cap); (4) founding grace period for post-placement structures; (5) registrar
cost + dues-% cap; (6) upkeep/dissolve fast-cycle mapping; (7) waypoint caps + tell
rate-limit + mail body limit; (8) shuttleport/trainer-hosting marker model.
PROPOSED interim convention (approval requested): Group-C-style NON-CANONICAL flagging
on every non-GDD number (code comments + HYGIENE_NOTE.md), magnitudes anchored to
fork-local scales (deed/upkeep/XP bands) to avoid fresh tuning. No placeholder survives
contact with real TCIndustries design (all civic numerics stay TBD/PROTOTYPE; OQ-007 +
OQ-013 untouched).

## 7. Acceptance criteria (verify, don't claim)

`go build ./...` + `go vet ./...` clean; fresh DB; phase0–6 tests green PLUS new
`phase7test` with THREE live clients (A founder + B/C citizens): A places City Hall →
city Rank 1 → B/C auto-enrolled → election (A runs, B+C vote → A mayor; double-vote
rejected; tie-path unit-covered) → mayor sets vendor-% tax → C buys from A-vendor →
treasury += tax with ledger transfer rows → upkeep tick under fast-cycle → Rank 2 flags
(cantina/medical) appear → Rank 3 shuttleport+trainer flags appear → guild founded by
A+B+C via registrar (sink ledger row) → dues % divert observed on bazaar sale → C mails
B credits+item while B offline → B logs in, claims intact (atomicity asserts both
sides) → 51st mail capped with unclaimed-attachment protection → permission checks
(banned-despite-public, friends-only entry) → overlap placement rejected → Rank-1
treasury starvation → dissolve, structures standalone-owned; all with real pasted
output; fast-cycle mapping test-config-only, defaults GDD-given; hygiene note extended;
commit per verified package (V1→V8).

## 8. Risks and dependencies

- Timer testability (weekly upkeep, 7-day terms, 30-day countdowns, 14-day triggers)
  is untestable live — fast-cycle mapping covers elections/upkeep/dissolve the way
  Phases 4–5 covered lifespans/grace; reclaim/30-day paths stay code-only (assert
  states, not waits).
- Closing permission tiers changes Phase 5 behavior (vendors/houses were entry-open):
  full-suite regression covers ordering effects as usual.
- Mail-item atomicity rides E2 + stack-store paths from Phase 4 — cross-package
  transaction spans need the same BEGIN-IMMEDIATE discipline, plus the
  send-time-ownership-transfer invariant (or the 18.5 edge reappears).
- Triple-membership independence (guild × city × faction-stub) must never gate core
  loops — a guildless/cityless/neutral player keeps full Phase ≤6 access (Pillar 2 +
  §15.1 neutral-safety carried forward).
- Scope-creep magnets: Rank 4/specialization, militia PvP, travel movement, clone
  binding, stalls — all §5-excluded; enforce at review.
- Ordering: Phase 5 AND Phase 6 review gates MUST pass before V1 starts (V2–V6
  terminate in their systems); EIC, provenance, TCIndustries-GDD edits untouched
  throughout.

## 9. Decisions requested of the human (the plan asks; it does not decide)

1. Base/location (§3) — confirm building on `testbed/swg-phase3-combat/`.
2. §30-gap packages G1–G3 — confirm IN (recommended) or sever to a follow-up.
3. Politician lightweight perk (§5.b) — mayor-status-gated upkeep discount (recommended)
   vs full deferral.
4. Rank-4 reachability — capped at 3 (recommended) vs data-shapes-only vs IN.
5. Numeric set (§6) — approve magnitudes/mappings or supply replacements.
6. Stalls/escrow follow-up — confirm OUT of this phase (recommended) with V6 as its
   unblocker.
7. Authorisation to implement on approval of 1–6 (cf. §10).

## 10. Implementation status

APPROVED AND EXECUTED 2026-09-15. Owner decisions (§9) received: base confirmed,
G1–G3 IN, mayor-status-gated upkeep discount (20% — Politician-substitute),
Rank 4 capped at 3 (shapes present, unreachable), numeric set approved,
stalls/escrow OUT (V6 mail is its unblocker), implementation AUTHORIZED.
Results on `testbed/swg-phase3-combat/`: `go build` + `go vet` clean; fresh-DB
phase7test 50/50 ALL PASS (real output) — founding active (8/4 pre-existing),
forming grace, auto-enrol, unanimous election + double-vote rejection + 500→400
upkeep discount, vendor-% tax (buyer −500 / city +50 / till 450), township→city
with shuttleport+trainer flags (exit criteria), flat-tax ledger accrual,
evict-guard refusal, opt-out without re-enrol, guild + 10% dues (till 480),
groups + member-only routing, mentorship XP, offline mail claim (credits+deed),
delete guard, 50-cap, ban/unban purchase gate, admin funding, zoning gate,
overlap refusal, Rank-1 exhaustion dissolve, grace-expiry dissolve, waypoints/
friends/tell/emotes/planet, faction-channel refusal, offline-tell refusal.
Full 8-suite regression (testclient + phase1–7, one fresh DB, fast-cycle):
ALL GREEN, exit 0. Bugs found by running: UpdateCity arg-order corruption
(status/grace_ends swap — caught by Test 6 ledger assert); `structure_banned`
table-name generation; test-side placedIDs staleness; per-character fixed
structure IDs → unique IDs (recorded behavior change); SQLITE_BUSY under
rapid mail sends → `SetMaxOpenConns(1)` pool hardening (plumbing, no mechanic);
test-name collision (`BuyerChar` → `CityBuyer`) and +100 m geography shift for
shared-DB coexistence. HYGIENE_NOTE.md extended (§Phase 7 additions +
correction log + known pre-existing lair-depletion fragility affecting Phase 6
on shared DB — evidence, not this phase's scope). Human review gate
(per §8/HD-TST-01 pattern) is now the next step before any Phase-8-equivalent
work.

*End of proposal v0.1. To enact: owner approves (or amends) §§3–6/9–10; approval and date
recorded before any code is written.*
