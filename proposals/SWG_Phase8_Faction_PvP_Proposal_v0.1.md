# SWG Pre-CU Phase 8 — Faction & Advanced PvP: Proposal Plan v0.1

**Status:** PROPOSED plan. No authority. No design decision. Human approval required
before any implementation, and separately before any numeric value is treated as tuned.
**Author / Assessor:** Muse Spark (`opencode/muse-spark-1.3-contributor-free`) via OpenCode —
research and planning only; all authority decisions remain human.
**Date drafted:** 2026-09-15.
**Controlling references:** `AGENTS.md`; `proposals/SWG_Phase7_Civic_Systems_Proposal_v0.1.md`
(executed — review gate pending); `testbed/swg-phase3-combat/HYGIENE_NOTE.md`;
`governance/TCIndustries_Testbed_Reuse_Human_Ruling_HD-TST-01_2026-09-14.md`;
predecessor GDD `sources/swg-pre-cu/SWG_PreCU_GDD.md` (§§7, 8.3.21, 9.2.2, 9.3, 9.4,
9.6, 12, 14, 15, 16, 18.2.1, 19.2.4, 28, 29.5, 30, 31 — HISTORICAL, never TCIndustries canon).
TCIndustries fence: CMBT-001 (combat as one lifestyle among many, LOCKED) is the
direction; CMBT-002 stays PROPOSED/TBD and OQ-008 stays TBD — testbed PvP proves
mechanics, never canon scope. CMBT-003/LOOP-002 (no forced combat path) are
satisfied by construction: neutral and covert characters are never valid targets.
EIC untouched (HD-EIC-07) — faction points are not credits (no faucet/sink impact;
death transfers are transfers). Provenance / formal reputation stays deferred —
faction rank is game-system standing, explicitly NOT social reputation (no scores,
no leaderboards; personal status + chat tag only). BAL-001 governs §6; OQ-012
(exploit implementation) stays TBD — arranged-kill farming is surfaced (§8) with
observation only, no invented policy.

**Change discipline (per AGENTS.md):** candidate plan in `proposals/` (no authority
change). HD-TST-01 patterns-only fork rules apply throughout: no SW names, lore,
or factions enter the fork — alignments are deliberately flavorless
(`alignment_a` / `alignment_b` + neutral) so no canon contamination is possible.
All fork behaviour remains PROTOTYPE/EVIDENCE.

---

## 1. Objective

Implement **Phase 8 — Faction & Advanced PvP** on the testbed fork (Phases 0–7
complete there, subject to pending review gates), satisfying the predecessor-GDD
§30 exit criteria verbatim: *"Two Overt, opposing-faction players can engage in
valid PvP combat with correct death penalties (Section 9.4.2), and faction
points/ranks accrue correctly."*

Entry criteria (§30: Phase 3 and Phase 7 complete — "needs combat resolution +
guild/city structures for bases") are MET in the fork (HAM resolution + incap/
revive/clone verified; guilds, cities,.distance-gated placement, ledger
transfers verified), conditional on the review gates — see §10.

## 2. Research basis

### 2.1 Predecessor-GDD map (HISTORICAL input, not authority)

| GDD section | Specifies | Phase-8 relevance |
|---|---|---|
| 9.4.1 flagging | Rebel/Imperial allegiance (SW names stay OUT per HD-TST-01); Overt = attackable, Covert = safe; `/pvp` toggle; Overt→Covert 5-min timer (anti-combat-log); attacking enemy auto-flags Overt | F1 core (generic alignments) |
| 9.4.2 open-world PvP | Both-Overt + opposing-faction gate; same-faction safe; no PvP in NPC cities (safe zones); wilderness + player cities "if city allows" | F2 gate + F6 city flag |
| 9.4.2 death penalties | 10% item-condition loss (double PvE); BF accumulation; small credit drop (PvP: 10% of carried) | F3 core (condition-minimal + transfer) |
| 9.4.3 bases | Wilderness placement; enemy attack/destroy; destroyer + defender points; HP-to-0 destruction; war windows set by owning guild; turrets (NPC automated defenses) | F5 (turrets OUT — needs NPC AI) |
| 9.6.1/9.6.2 death | Incap 5-min → revive (Combat Medic; base-Medic interim per Phase 6 precedent); death → clone + wounds/BF + credit drop | Reuse paths (clone-from-incap needs no timer wait) |
| 15.2.1 alignment | Rebel/Imperial/Neutral (genericized); switching allowed, resets progress, 30-day cooldown [ASSUMPTION]; neutrals never Overt, never valid targets, never restricted elsewhere (§15.1) | F1 (cooldown code-only) |
| 15.2.3 ranks/points | Recruit→Sergeant→Major→Colonel→General at 0/2500/10k/30k/75k+sponsorship (thresholds [HISTORICAL, ASSUMED]); points from enemy-player kills, enemy-NPC kills, base assaults/defenses; 60-day inactivity decay [ASSUMPTION] | F4 (General+sponsorship OUT; NPC-kill points blocked — no factional NPCs) |
| 15.2.4 base gates | Base deeds need founder Major+; owning guild single-faction at officer level [ASSUMPTION]; overlap/zoning validation like cities | F5 gates |
| 15.2.5 perks | Rank-scaled vendor discounts 10–30% (needs faction NPC vendors — BLOCKED, none exist); Major+ exclusive schematics (EXPANSION per 15.6 — OUT); faction chat | Discounts/schematics OUT; chat IN (F7) |
| 15.3 entities | FactionStanding / FactionBase shapes (overt flag + expiry, points, sponsor, HP, turrets, war windows) | Persistence shapes (turrets omitted) |
| 15.5 edges | Neutral never a valid target; no mid-siege faction switch; sponsorship expiry 14 d; mixed guilds barred at base placement | Rules to implement (sponsorship OUT with General) |
| 15.6/31.4 MVP | Alignment, flagging, points, ranks through Colonel, vendor discounts (blocked); EXPANSION: full siege, General+sponsorship, schematics, specialization influence | Scope backbone |
| 14.2.3 militia | Mayor appoints city-sanctioned PvP defenders; §14.6 marks militia tools EXPANSION | Column stays dormant (OUT) |
| 14.4 leaning | City Rebel/Imperial/Neutral lean from citizen composition [ASSUMPTION]; gates nearby base tolerance | Display-only % (F6; enforcement OUT) |
| 19.2.4 guilds | Membership independent of faction; alignment label; hall/base ownership by guild entity | F7 (members unrestricted; officers gated at base placement — flagged reading) |
| 18.2.1 faction chat | Same-faction galaxy-wide channel | Fills Phase 7 stub (membership-gated) |
| 9.2.2 resolution | 8-step attack flow reused vs players (HAM pools, posture/stance, armor mitigation — all exist) | F2 (no new combat math) |
| 8.3.21/16 bounties | Bounty Hunter terminal, player-target bounties from faction pools + player contracts | Phase 9 (overt roster is its future input — interface only) |
| 12 faucets/sinks | Points are not credits; death credit move is a transfer (victim→killer), never a faucet | Accounting rule (F3) |

### 2.2 Existing implementations (evidence, not authority)

- **Fork combat**: 9.2.2 resolution vs creatures; HAM/wounds/BF; incap 5-min; revive
  gated on `medic_novice` (base-Medic interim); clone-from-incap (no timer wait);
  armor mitigation via crafted plate; posture/stance; combat XP (PvE) + kill-drop
  faucet. `handleCombatAction` targets creature IDs only ("player-vs-player
  targeting is later-phase scope" — this phase is that scope).
- **Fork civic (Phase 7 hooks)**: faction chat channel exists but refused ("later
  phase"); guild `faction` column (all `'neutral'` — no SW strings anywhere in the
  fork, verified); `is_militia` column dormant; entity_spawn without overt/faction;
  E2 AtomicPurchase with city-tax/dues split precedent (tx-scoped TableReads +
  ledger pairs); zoning binary-gate precedent; registrar-as-HTTP-endpoint precedent
  (guild founding without registrar NPCs — same honesty applies to faction
  declaration without recruiter NPCs).
- **Gaps this phase fills**: flagging, alignment, PvP targeting validity, PvP
  death penalties, points/ranks, bases (MVP-subset), faction chat membership,
  city PvP permission + leaning display. No PvP/faction code to adapt — greenfield
  on fork-local patterns, reusing combat + treasury + ledger machinery.
- **Group C / Track A/B / demos**: zero faction/PvP code. Nothing transfers.

## 3. Base and working location (RECOMMENDED — human confirms)

Build in place on `testbed/swg-phase3-combat/` (new `internal/faction/` + DB
extensions + handlers + `cmd/phase8test/`); no new fork, `sources/` untouched.
Rationale: every Phase 8 flow terminates in Phase ≤7 systems (combat resolution,
incap/clone/revive, guilds, cities, wallet/ledger, chat routing). Single-zone
simplification carries over (alignments are zone-agnostic character flags).

## 4. Scope IN (work packages, each verify-by-running before the next)

- **F1 — Alignment + flagging**: declare/switch alignment via HTTP
  (`alignment_a` / `alignment_b` / `neutral`; flavorless by HD-TST-01 design —
  human may rename); switch resets points + rank progress; 30-day cooldown
  [ASSUMPTION] (code-only, untestable live — flagged); neutrals can never go
  overt. Overt/covert WS toggle; Overt→Covert 5-min delay (GDD-given; fast-cycle
  30 s); Covert→Overt instant; accepted attacks auto-flag the attacker overt
  (GDD-literal; rejected attempts do NOT flag — flagged reading). Overt + faction
  carried on `entity_spawn` (client-distinguishable, creature-template precedent).
- **F2 — PvP resolution**: `combat_action` accepts character IDs; validity =
  attacker aligned (non-neutral) + victim overt + opposing alignments + same
  zone + in-range (8 m, existing) + conscious + city `pvp_allowed` at the fight
  location + target not self; neutral-attacker / covert-victim / same-faction /
  cross-zone / out-of-range / city-disallowed rejected with explicit errors.
  Covert attackers are NOT rejected outright: a covert attacker with an
  otherwise-valid target is auto-flagged overt by the attempt (reconciles
  "cannot attack enemies" with "attacking flags Overt" — covert cannot attack
  *while remaining* covert; victims must always be overt). Reuses 9.2.2 flow, posture/stance, armor mitigation;
  `combat_result` broadcast unchanged. No combat XP for PvP kills (points only —
  flagged; 7.2.1 formula is PvE-scoped and effective-level is unbuilt).
- **F3 — PvP death penalties** (exit-criteria core): existing incap path
  (5-min timer, base-Medic revive per Phase 6 precedent — Combat Medic elite
  stays deferred); clone-from-incap applies all three 9.4.2/9.6.2 penalties
  atomically: (a) wounds/BF via existing ranges; (b) 10% of carried credits
  victim→killer as a ledger-tagged transfer in the death tx (wallet-to-wallet,
  offline-safe — "drop" resolved as transfer, §5.d); (c) −10% item condition on
  equipped items via a new `condition_pct` column (default 100, floor 0, no
  breakage, no repair — flagged minimal surface for OQ-010; repair/breakage are
  follow-ups, never silent).
- **F4 — Points + ranks**: kill award 100 [PROVISIONAL] to the killer on a PvP
  kill (incap credit; no farming policy — observed only, §8); rank thresholds
  GDD-given (0/2500/10k/30k; fast-cycle 100/200/300 so Sergeant→Colonel execute
  live — defaults stay GDD-given, same compression precedent as founding counts);
  General (75k + sponsorship) EXPANSION-OUT; 60-day decay code-only; personal
  status endpoint (+ chat sender tag). NPC-kill points recorded BLOCKED (no
  factional NPCs exist — Phase-5 faucet-subset honesty precedent).
- **F5 — Bases (MVP-subset)**: base deed schematic (Architect-crafted by analogy
  — GDD names no crafter, flagged; placement gated on founder Major+); WS
  positional placement (wilderness-equivalent: outside city radii + no-overlap,
  hall-placement precedent); guild ownership (entity-owned, hall precedent);
  officers single-alignment at placement [ASSUMPTION-literal]; HP 3000
  [PROVISIONAL ≈ 90 s solo siege at sidearm scale]; war windows set by owning
  guild (attacks outside rejected); damage log for attribution (new telemetry
  table); destroy → 250/participant [PROVISIONAL] + defender-kill points via F4;
  no upkeep, no repair (GDD-silent — flagged omissions); turrets OUT (needs NPC
  combat AI — EXPANSION).
- **F6 — City integration**: `pvp_allowed` flag (default true per "unrestricted";
  mayor-toggable, zoning-gate precedent; containment enforced on attack);
  leaning as display-only derived % in `city/get` (14.4 [ASSUMPTION]; enforcement
  OUT). Militia: no endpoint (EXPANSION per 14.6); column stays dormant.
- **F7 — Social**: faction chat membership-gated (fills Phase 7 stub; opposing
  side excluded — test asserts); guild alignment label (leader-set; members
  unrestricted — §5.g reading). Overt roster query (read-only; Phase 9 bounty
  input — interface only, no bounties).
- **F8 — `phase8test`** (3–4 clients, crafted sidearms for PvP TTK): gating
  rejects → overt duel to incap → points/events → clone penalties (transfer +
  condition) → covert-delay + re-flag → ranks to Colonel live (fast-cycle) →
  base place (Major-gated; refusal first) → in-window siege → destroy awards →
  out-of-window refusal → city-disallow refusal → faction chat isolation →
  leaning display; plus re-run of ALL Phase 0–7 tests per verify-don't-claim.

## 5. Scope OUT (explicitly deferred, with homes and tensions)

General rank + sponsorship; turrets; base upkeep/repair; vendor discounts (no
faction NPC vendors — blocked); faction-exclusive schematics (EXPANSION);
specialization influence (Rank 4 unreachable + Politician deferred — doubly
blocked); Bounty Hunter + player bounties + missions (Phase 9); duels (no GDD
source — overt/overt is the only gate); TEF/associative flagging (no GDD source —
aiding an overt player carries no flag in this build, flagged omission); NPC-city
safe zones (vacuous — no NPC cities exist; spawn-area exposure to overt attack is
a flagged consequence, not a patched-over rule); DoTs/conditions beyond existing
states; group XP/threat in PvP (no XP); Jedi/space/quests/instancing (§32, never).
**GDD-internal tensions recorded (human resolves):** (a) discounts need unbuilt
vendors — blocked, not dropped; (b) "enemy NPC kills" need unbuilt factional
NPCs — blocked; (c) base-deed crafter unspecified — Architect by analogy
(flagged); (d) 9.6.2 "drop" — transfer chosen over destroy-as-sink (keeps
faucet/sink accounting clean); (e) guild mixed membership vs officer-gated bases
— members unrestricted, officers gated (flagged reading of 15.5);
(f) aiding-overt/TEF absence — omitted for lack of source, not oversight.

## 6. Numeric authorization requests (values NOT set by this proposal)

GDD-GIVEN (implement as written): Overt→Covert 5-min delay; rank thresholds
0/2500/10k/30k/75k+sponsorship ([HISTORICAL, ASSUMED] — defaults verbatim);
10% credit drop; 10% condition loss; wound/BF death ranges (exist); war-window
SHAPE (guild-set); 30-day switch cooldown + 60-day decay ([ASSUMPTION]-tagged,
implement flagged, code-only paths); PvP TTK 30–60 s as acceptance TARGET, not
input. GDD-SILENT / INDEPENDENT (each needs approval OR interim-[ASSUMPTION]
sign-off): (1) kill award 100; (2) destroy award 250/participant; (3) base HP
3000; (4) fast-cycle rank thresholds 100/200/300 + covert delay 30 s; (5) base
deed schematic cost/stats; (6) war-window bounds (min/max duration, max windows);
(7) condition floor-0 no-breakage rule; (8) overt roster visibility (self-side
only vs galaxy-wide). PROPOSED interim convention (approval requested):
Group-C-style NON-CANONICAL flagging on every non-GDD number (code comments +
HYGIENE_NOTE.md), magnitudes anchored to fork-local scales (kill-drop 10×CL,
training-cost bands, structure upkeep scale) to avoid fresh tuning.

## 7. Acceptance criteria (verify, don't claim)

`go build ./...` + `go vet ./...` clean; fresh DB; phase0–7 tests green PLUS new
`phase8test` with FOUR live clients (A align_a + B align_b duelists with crafted
sidearms, C align_a guildmate, D neutral): A/B overt → A attacks B (auto-flag
observed) → rejections sampled (C-targeted, covert B-targeted, same-side,
neutral-safe, city-disallowed after mayor toggle, out-of-window base hit) →
duel to incap (real resolution) → B clones: wallet −10% to A with ledger pair,
equipped condition −10%, wounds/BF per existing ranges → A points + events
recorded → 3 kills under fast-cycle thresholds → A Colonel (+Major-gated base
deed placed by A in wilderness for A's guild; non-Major placement refused) →
war window declared → B sieges to 0 HP → destroy awards → faction chat A↔C
received, B excluded → leaning displayed on city/get → all with real pasted
output; fast-cycle mapping test-config-only, defaults GDD-given; hygiene note
extended; commit per verified package (F1→F8).

## 8. Risks and dependencies

- Duel-duration variance (hit RNG around regen): sidearm-scale TTK keeps duels
  to minutes, but walk-backs after clone add up; budget ~35 min like Phase 7.
- Arranged-kill point farming is INHERENT (two consenting players can trade
  kills): GDD provides no counter-rule, OQ-012 implementation is TBD — this plan
  instruments (per-kill events table, SAFE-003 observability) and enforces
  nothing. Do not invent throttles without human approval.
- Neutral/covert safety is load-bearing for CMBT-003: every new attack path must
  re-check validity (regression + pre-registered negative tests in §7).
- No-safe-zone consequence: overt characters are attackable even at the spawn
  area. No geography exists to fix this; the plan records it rather than
  inventing sanctuary rules.
- Ordering: Phase 5 AND 6 AND 7 review gates MUST pass before F1 starts
  (F2–F5 terminate in their systems); EIC, provenance, TCIndustries-GDD edits
  untouched throughout.
- Scope-creep magnets: turrets, repair, duels, TEF, bounties — all §5-excluded;
  enforce at review.

## 9. Decisions requested of the human (the plan asks; it does not decide)

1. Base/location (§3) — confirm building on `testbed/swg-phase3-combat/`.
2. Flavorless alignment IDs (§4/F1) — confirm `alignment_a/b` or supply names.
3. Fast-cycle threshold compression (§6.4) — confirm or supply.
4. Award/HP set (§6.1–6.3: kill 100, destroy 250/participant, base HP 3000) —
   approve magnitudes or supply.
5. Death economics (§5.d): credit-transfer-always + condition-minimal (recommended)
   vs alternatives.
6. Base-deed crafter (§5.c: Architect) + no-upkeep/no-repair omissions — confirm.
7. Authorisation to implement on approval of 1–6 (cf. §10).

## 10. Implementation status

APPROVED AND EXECUTED 2026-09-15/16. Owner decisions (§9) received with no
amendments and no supplied values (all recommendations confirmed): fork base,
flavorless `alignment_a/b`, fast-cycle threshold compression, award/HP set
(kill 100, destroy 250/participant, base HP 3000), credit-transfer-always +
condition-minimal death economics, Architect base deeds + no-upkeep/no-repair,
implementation AUTHORIZED.
Results on `testbed/swg-phase3-combat/`: `go build` + `go vet` clean; fresh-DB
phase8test 39/39 ALL PASS (real output) — alignment/flagging, gating rejects,
covert-attack auto-flag, three overt duels to incap with 10% transfers +
condition −10 + wounds/BF + points to Sergeant→Major→Colonel, city permission
gate (live, both directions), Major-gated wilderness base, out-of-window
refusal (precise reason), siege to destruction + destroy awards, destroyed-base
refusal, faction-chat isolation, neutral refusal, overt roster, leaning
display, covert delay (refusal + success), leave-reset, cooldown refusal,
neutral declare/leave. Full 9-suite regression (testclient + phase1–8, one
fresh DB, fast-cycle): ALL GREEN, exit 0.
Bugs found by running: quota single-stack starvation (new ensureStack top-up +
guarantee anchoring); quota arithmetic shortfall (five deeds before the
sidearm); single-connection self-deadlock in GuildOfficerAlignments (open rows
+ nested lookup under SetMaxOpenConns(1) — fixed two-phase + static audit of
all 38 rows-loops clean); test rank arithmetic (250 = Major — test fixed, code
right). HYGIENE_NOTE.md extended (§Phase 8 additions + correction log +
F2-reconciliation review-gate note). Pre-existing Phase 6 lair-depletion flake
recurred once more on a shared-DB run (5/6 dead again); rerun green — still
out of scope (respawn design needs a human decision).
REVIEW-GATE ITEM (§4/F2 refinement) — CONFIRMED 2026-09-16. The approved text
said "both overt" while also quoting "attacking flags Overt" — implemented so
a covert attacker with an otherwise-valid target is auto-flagged and proceeds
(rejected attempts never flag). Owner confirmed as-implemented at the gate; no
amendment, no code change.
Human review gate (per §8/HD-TST-01 pattern) is now the next step before any
Phase-9-equivalent work.

*End of proposal v0.1. To enact: owner approves (or amends) §§3–6/9–10; approval and date
recorded before any code is written.*
