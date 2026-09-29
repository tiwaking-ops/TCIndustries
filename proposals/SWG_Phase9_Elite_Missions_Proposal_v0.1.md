# SWG Pre-CU Phase 9 — Remaining Elite Professions & Missions: Proposal Plan v0.1

**Status:** PROPOSED plan. No authority. No design decision. Human approval required
before any implementation, and separately before any numeric value is treated as tuned.
**Author / Assessor:** Muse Spark (`opencode/muse-spark-1.3-contributor-free`) via OpenCode —
research and planning only; all authority decisions remain human.
**Date drafted:** 2026-09-16.
**Controlling references:** `AGENTS.md`; `proposals/SWG_Phase8_Faction_PvP_Proposal_v0.1.md`
(executed — review gate pending); `testbed/swg-phase3-combat/HYGIENE_NOTE.md`;
`governance/TCIndustries_Testbed_Reuse_Human_Ruling_HD-TST-01_2026-09-14.md`;
predecessor GDD `sources/swg-pre-cu/SWG_PreCU_GDD.md` (§§6.3, 7.2–7.3, 8.1–8.4,
9.2.4, 9.3.1, 9.5.3, 10, 11.3–11.4, 12.2.1, 14.6, 15, 16, 17, 18.2.3, 19.2.3,
23, 24, 25, 26, 27.2.5, 28.3, 29.5, 30, 31, 32, 33.3 — HISTORICAL, never
TCIndustries canon).
TCIndustries fence: PROF-001 (flexible skill architecture, LOCKED) and PROF-004
(respecialisation allowed, LOCKED) are the direction — the existing
drop-cascade is reused unchanged. PROF-003/005 stay PROPOSED/TBD and OQ-001
stays TBD: the fork's 250-cap/66/80/120 cost economy is testbed mechanics,
never the TCIndustries specialisation answer. OQ-011 (respec costs) untouched —
no new costs invented. CRFT-006 stays PROPOSED. PIL-003/VIS-004
(interdependence): §8.4's dependency chains are directional support only, not
canon. LOOP-002: non-combat elites (Merchant, Architect, Image Designer,
Creature Handler, Bio-Engineer-roster) must be viable without combat. EIC
untouched (HD-EIC-07) — Merchant fees/XP and mission credit rewards are
mechanics and faucet *rates* (approval via §6/§9 under BAL-001), not EIC
architecture. MFG/factories stay OUT (MFG-001/004, HD-EIC-05/07). SAFE-001
governs contract-escrow atomicity; OQ-012 stays TBD (poster-target collusion
allowed but observed — no invented policy). OQ-007/013 untouched. OQ-010:
condition work stays at the Phase 8 minimal surface (tissue OUT with DNA).
Provenance stays deferred (points/ranks are system standing, not reputation).

**Change discipline (per AGENTS.md):** candidate plan in `proposals/` (no authority
change). HD-TST-01 patterns-only rules apply: generic names/content only, all
fork behaviour PROTOTYPE/EVIDENCE, no TCIndustries canon touched.

---

## 1. Objective

Implement **Phase 9 — Remaining Elite Professions & Missions** on the testbed fork
(Phases 0–8 complete there, subject to pending review gates), satisfying the
predecessor-GDD §30 exit criteria verbatim: *"Full Section 8 profession roster is
playable; mission terminals generate and reward correctly across all MVP-tagged
mission types."*

Entry criteria (§30: Phases 3–8 complete — "these professions build on combat,
crafting, and faction systems already live") are MET in the fork (PvE/PvP
resolution, schematic/experimentation pipeline, overt roster + kill hooks,
guilds, chat, ledger), conditional on the review gates — see §10.

**Scope warning (the central §9 decision):** the GDD contradicts itself three ways
on what Phase 9 contains — §30 names 7 elites (Bounty Hunter, Commando, Smuggler,
Creature Handler, Bio-Engineer, Image Designer, TKA) plus missions/Ranger/
surveying; §31.4 demands 14 *different* MVP elites (Armorsmith, Weaponsmith,
Architect, Merchant, Pistoleer, Rifleman, Carbineer, Fencer, Swordsman, Doctor,
Combat Medic, Musician, Dancer, Creature Handler) with no §30 home; §31.5 defers
Droid Engineer, Bio-Engineer, Ranger, Pikeman, TKA, Image Designer, Bounty Hunter,
Commando, Smuggler (+ un-spec'd Chef/Tailor, Politician, Squad Leader) as
EXPANSION while §30 simultaneously schedules six of them. §4 recommends the
reconciliation (option b, with a split fallback); §9 asks the human to settle it.
Nothing here resolves the contradiction — it is surfaced for decision.

## 2. Research basis

### 2.1 Predecessor-GDD map (HISTORICAL input, not authority)

**Elite roster (§8.3 — all 23 entries read; prereq / trees / XP / cost):**

| Elite | Prereq | XP pool (new pools marked *) | Cost | Phase-9 role |
|---|---|---|---|---|
| Armorsmith | Master Artisan | armor_crafting * | ~80 | Trees + elite armor schematics + XP hook |
| Weaponsmith | Master Artisan | weapon_crafting * | ~80 | Trees + elite weapon schematics + XP hook |
| Architect | Master Artisan | structure_crafting * | ~80 | Trees + civic-deed schematics (cantina/hospital/shuttleport kinds complete Phase 7 flags) + XP hook |
| Droid Engineer | Master Artisan | droid_crafting * | ~80 | OUT — EXPANSION per 31.5 (needs pet-AI-equivalent + programming system) |
| Merchant | Master Artisan | merchant (pool exists, 7.2.5) | ~80 | Trees + sale-XP hook (1 XP/100 cr + unique-buyer weighting) + fee reduction to 50% + slot bonus (fills Phase 5 recorded gap) |
| Bio-Engineer | Master Scout + Novice Medic | bio_engineering * | ~80 | Roster data + FULLER DNA mechanics IN (owner decision 2026-09-16 via §9.3 — explicitly overrides the 17.6 EXPANSION deferral for the testbed): DNA sampling + combining into enhanced-stat pet variants + tissue sampling (§5.d amended) |
| Creature Handler | Master Scout | creature_handling * | ~80 | Trees + tame/own/command/release loop (CL cap 20 Master; max 3 pets) |
| Ranger | Master Scout | ranger (7.2.4) | ~80 | Trees + camps (MVP) + creature tracking (MVP-lite) + survey bonuses |
| Pistoleer/Rifleman/Carbineer | Master Marksman | pistol/rifle/carbine_combat * | ~80 ea | Trees + style-tagged schematics; specials OUT (no ability system) |
| Fencer/Swordsman/Pikeman | Master Brawler | fencing/sword/polearm_combat * | ~80 ea | Trees + style schematics; specials OUT (Pikeman IN despite 31.5 — spec'd + mechanics-free, §5) |
| TKA | Master Brawler + 3 melee masters | teras_kasi * | ~80 (§8.3.15) vs 120 (§33.3 multi-prereq) — tension §5, recommend 120 | Trees + unarmed scaling + meditation regen; combos OUT |
| Doctor | Master Medic | medical (exists) | ~80 | Trees + buff-pack crafting (extends stim pattern) + heal scaling; cures BLOCKED (no disease/poison states) |
| Combat Medic | Master Medic | medical (exists) | ~80 | Trees + in-combat healing (stance-break exemption); AoE/shields OUT; revive gate unchanged (no relitigation of Phase 6) |
| Musician/Dancer | Master Entertainer | entertaining (exists) | ~80 ea | Trees + elite BF-heal scaling + performance group buffs + instrument/dance content (data) |
| Image Designer | Master Entertainer | image_designer * | ~80 | Trees + appearance-edit handshake + holoemotes |
| Bounty Hunter | Master Marksman/Scout + 4 masters | bounty_hunter * | ~120 | Trees + player tracking (roster exists) + BH-terminal visibility + unique-weapon schematics; specials OUT |
| Commando | Master Marksman + 3 masters | commando * | ~120 | Trees + heavy-weapon schematics (single-target + HAM-cost representation; AoE option for human) + ammo consumables |
| Smuggler | Master Pistoleer + Master Scout | smuggler * | ~120 | Trees + slicing (handshake + one-time flag) + spice consumables + smuggling missions; evasion/contacts OUT (unscoped) |
| Chef / Tailor / Miner | — (no 8.3 entries; named only in §8.4 chains + §17) | — | — | OUT — unspecifiable from source (no trees/prereqs/costs) |
| Politician (full) / Squad Leader | HISTORICAL §§14/19 + 31.5 deferred | — | — | OUT — EXPANSION, no §30 home |

**Supporting systems:**

| GDD section | Specifies | Phase-9 relevance |
|---|---|---|
| 7.2.1 combat XP | Kill scaling, lair bonus, mission XP, group bonus 1.0–1.5x, diminishing, boss 5–10x; **PvP kills grant factional + combat XP** | Style-pool kill awards (new); group bonus UNBUILT (missions accepter-only, §5); **prior-phase residual: Phase 8 awarded points but no combat XP on PvP kills — correction candidate, human-directed (not silently patched here)** |
| 7.2.2 crafting XP | Per-item + complexity bands (10–50 simple, 500–2000 complex) + repeat diminishing (50% after 10th) + crit 2x | Complexity mapping + diminishing + crit bonus (GDD-given mechanics; fork flat-500s untouched for existing schematics) |
| 7.2.3 social XP | Per-tick BF / per-point wounds + tier scaling + diminishing | Residual: fork flat-50s stand (Phases 5/6 provisionals, not relitigated) |
| 7.2.4 scout/ranger XP | Survey, tracking, harvest, camping use | Tracking/camping earn hooks (new); survey/sample stay scouting |
| 7.2.5 merchant XP | Per-transaction + 1 XP/100 cr + unique-buyer weighting; brokering | Sale hook + diminishing (brokering OUT — no direct trade) |
| 8.4 chains | Crafter/combat/support dependency maps (incl. Chef/Tailor/Miner mentions) | Directional support only; Miner unspec'd (flagged observation) |
| 9.2.4 abilities | Special-attack system (unbuilt since Phase 3) | **Major residual: most combat-elite "specials" wait on it — trees trainable, specials OUT; future Ability-System phase is human-directed** |
| 9.5.3 + 17.2.3 lairs | 3–10 spawns; **population regenerates to max without lair destruction**; destroyed lairs relocate in 12–24 h (5.4.2); lair HP | **E2: regen + relocation + lair damage hook — GDD-given mechanics the fork never built (root-cause closure for the Phase 6 depletion flake, in-scope via 17.6 MVP + mission dependency, not scope creep)** |
| 17.2.2 harvest | Fixed yield/template, per-corpse quality, 10-min window [ASSUMPTION], group-loot interplay | Reuse (harvest-gating by Hunting tiers is new: rare-harvest gate, flagged) |
| 17.2.4/17.3/17.5/17.6 taming | Tamable flag + Command-tier CL caps; aggro-at-third-party block; Pet entity; loyalty + DNA + pack depth + boss uniques EXPANSION; **basic taming MVP** | E5 pet loop (loyalty/DNA/packs/bosses OUT) |
| 12.2.1 mission faucet | Mission credit rewards 500–10k band | In-band provisional mapping (BAL-001 approval) |
| 16.2–16.6 missions | Terminal types; 5–10 min refresh/rate-limit; 8–15 listings; 2 h expiry; difficulty scaling; lifecycle (log cap 3–5 [ASSUMPTION], waypoints, turn-in); 50% other-destroyed rule [ASSUMPTION]; group rules; delivery-to-NPC-only anti-griefing | E4 (MVP set: Destroy Lair, Recon, Delivery + player bounties; Assault + NPC-audience OUT — no NPCs, blocked with reason) |
| 18.2.3 holoemotes | Custom player-designed, Image-Designer-dependent premium goods | E6 (dynamic per-player customs) |
| 23/24 service specs | Musician/Dancer demand, Doctor packs, CM field healing (already mapped in Phase 6 research) | E6 elite extensions reuse Phase 6 machinery |
| 25 ranger | Tracking (creature + PvP player), camps (30–60 min [ASSUMPTION], wilderness-only, Scout-reduced, non-stacking, 24.5 persistence), traps, camo; **MVP: camps only** | E7 (camps + creature tracking + Scout bonuses; player-tracking/traps/camo OUT per 25.6) |
| 26 surveying | Tiers basic (1/sample, ±500 m) / crafted (2–3, ±250 m) / master+Scout (±100 m, faster); Scout Hunting cooldown/detection; Ranger yield + lair-detect; triangulation refine; **MVP: basic + crafted + Scout bonuses** | E3 (imprecision mechanic + refinement rule are new, flagged) |
| 27.2.5 logout | 10–20 s interruptible camping channel [ASSUMPTION] | OUT — distinct from Ranger buff camps; do not conflate (flagged) |
| 28.3 schema | Profession/SkillTree/SkillBox entity shapes | Fork mirrors it; elite defs plug into `generateProfession` + cost-table parameterization + master-gate validation |
| 31.4/31.5/33.3 | MVP-14 vs deferred lists; 40–60 h basic / 80–120 h elite / 150–250 h multi-prereq / 200–300 h cap | Time targets (not inputs); TKA 80-vs-120 tension (§5) |
| 8.2.4 Scout (re-read) | Trees: Exploration/Hunting/Trapping/Survival; ~40 to Master; Hunting = tracking + harvest bonuses; Survival = camps; elite paths listed | **Fork deviation noted: fork basics cost 66 (verified ladder), GDD says ~40 — fork numbers govern the testbed (NON-CANONICAL placeholders either way); Hunting-tree home EXISTS in fork (README table saying "Harvesting" is stale doc, code governs)** |

### 2.2 Existing implementations (evidence, not authority)

- **Fork skills**: 6 basics as data (66-pt ladder, 250 cap, prereq chains, novice gates); `CategoryElite/Hybrid` constants exist with ZERO elite defs; XP pools combat/crafting/scouting/medical/entertaining/merchant(+mentoring) as free-form strings; trainer registry 6 basics (INSERT OR IGNORE); drop-cascade respec.
- **Fork gathering/combat**: survey (fixed 500 m radius, exact-center waypoints, basic yield, 10 s cooldown), harvesters, corpse harvest (ungated, fixed-type + quality roll, 10-min window); lairs (HP columns exist, NO damage path, NO regen, NO relocation); creatures retaliation-only AI (no passive/flee/pack/humanoid classes); no camps/traps/tracking/pets/DNA/tools-tiers/triangulation.
- **Fork economy/social hooks to reuse**: schematic registry + sessions + XPReward; vendor listings + E2 atomic purchase (sale hook site for Merchant XP); ledger faucet/sink/transfer; buff machinery (pool/amount/expiry/persistence/non-stacking — camp + performance + spice buffs ride it); mail/waypoint/chat/presence; overt roster + pending-kill hook (bounty completion rides it); guild/treasury.
- **Prior-phase OUT-lists landing here**: Combat-Medic-gated revive (stays base-Medic — no relitigation), DoTs/conditions (stay OUT), factories/R8 (stay OUT — EIC/MFG), stalls/escrow (UNBLOCKED by Phase 7 mail but unassigned — §30 doesn't list it; human direction needed, not Phase 9 scope), E5b stalls (same).
- **Group C / Track A/B / demos**: zero elite/mission/pet/camp content. Nothing transfers.

## 3. Base and working location (RECOMMENDED — human confirms)

Build in place on `testbed/swg-phase3-combat/` (new `internal/elites/` constants? — rather: extend `internal/skills/` defs + `internal/resources/` tool tiers + new `internal/missions/`, `internal/pets/`, `internal/ranger/` pure packages + DB extensions + handlers + `cmd/phase9test/`); no new fork, `sources/` untouched. Rationale: every Phase 9 flow terminates in Phase ≤8 systems (trees/XP/trainers, combat resolution, schematic sessions, overt roster, kill hooks, buffs, ledger, waypoints). Single-zone simplification carries over.

## 4. Scope IN (work packages, each verify-by-running before the next)

- **E1 — Elite roster data**: ~21 ProfessionDefs (all §8.3-spec'd elites except Droid Engineer; §5) with GDD-literal trees/names/prereqs/XP pools; cost tables parameterized (standard 80-class: tiers 3/4/5/6 + Master 8; 120-class: 4/6/8/10 + Master 8 — provisional distributions summing right, flagged); master-gate validation in `PrerequisiteCheck` (cross-profession Master boxes); new XP pool constants; trainer coverage mapping (elites train at matching base trainers — flagged stand-in, registrar precedent; no NPC authoring); existing 500-XP artisan schematics UNTOUCHED (new elite schematics carry complexity-scaled XP per E6-craft rule below).
- **E2 — Lair cycle (GDD-given, fork-missing)**: `combat_action` vs lair IDs (base-attack precedent: skill + range + conscious, no faction/window); lair HP damage; population respawn up to max while the lair stands (owner-supplied Pre-CU recollection of a 30-minute tick, 2026-09-16 — recorded as HISTORICAL-USER input, unverified against the corpus; AMENDED 2026-09-16 per §10 amendment: the 30-minute rate is NOT applied in the testbed — implementation uses the provisional +1 per harvest tick; the production datum is retained as a HISTORICAL-USER note only; exact refill shape, incremental vs full, remains TBD); destroyed-lair relocation outside active cities after 12 h default / 10 min fast-cycle (provisional mapping of the GDD 12–24 h window); per-kill XP + destroy bonus XP (provisional 150). **Side effect, recorded as evidence: this closes the root cause of the Phase 6 shared-DB depletion flake (regen + relocation); the flake note stands as history.** Until E2 executes, the flake persists (no interim fix — that would be an unapproved design change).
- **E3 — Surveying depth (26.2/26.6 MVP)**: crafted tool schematics per category (artisan-gated, 15M+8P precedent); crafted tier = 2–3 yield + ±250 m waypoint imprecision (NEW: offsets within precision — flagged); triangulation refinement (successive same-spawn surveys halve error, 60 s window — provisional rule, flagged); Scout bonuses keyed on Exploration/Survival tiers (cooldown −10%/tier, radius +100 m/tier — provisional; Hunting-tree citation adapted — fork HAS Hunting (tracking/harvest), recommendation: Hunting tiers gate rare-corpse-harvest + tracking-assist instead, flagged); Ranger Wilderness Survival yield bonus (provisional); master tier + Ranger lair-detect OUT (26.6 EXPANSION).
- **E4 — Mission terminals**: seeded combat + crafting + player-bounty terminals (bazaar-terminal precedent; no city terminals — no geometry phase); lazy regeneration (no new tick: list/accept/turn-in sweep expiry + refill to 10 listings — flagged simplification); mission log cap 4 ([ASSUMPTION] midpoint); Destroy Lair (lair-damage attribution via E2 log pattern; 50% other-destroyed rule [ASSUMPTION]), Recon (radius-presence check), Delivery (craft + terminal turn-in, goods destroyed as ledger sink — registrar/terminal stand-in precedent), Sample (units consumed as sink); rewards in the 500–10k band by difficulty (provisional in-band mapping) + typed XP mapping (flagged) + zero faction points in MVP ("if applicable" never applies yet — flagged); accepter-only completion (group sharing OUT — needs group-bonus design, flagged); Bounty-Player contracts (target overt at posting; upfront escrow deducted to contract row; BH-gated visibility (Novice BH); kill hook pays ALL stacked contracts + 100 BH XP [PROVISIONAL]; expiry refunds poster (anti-grief, flagged); poster=killer allowed (zero-sum, flagged)).
- **E5 — Pets + DNA (17.2.4; 17.6 partially overridden by §9.3)**: pets table (17.3 Pet entity minus loyalty); tame interaction (8 m, target idle — strict reading of the 17.5 third-party block, flagged; CL caps [3,8,12,16,20] provisional; success 50% + 10%/Command-tier provisional; fail aggros tamer — CONFIRMED risk rule); tameable flags (0001 yes / 0002 no — flagged test content); own (cap [1,1,2,2,3] provisional) / follow (tick moves toward owner outside 5 m — flagged) / stay / attack-owner's-target (retaliation-AI-shaped ticks — flagged) / release; training flags (data only); loyalty/packs/bosses OUT. **DNA (amended per §9.3 fuller-DNA decision):** Bio-gated DNA sampling from in-window wild corpses + owned pets anytime (flagged reading of "harvestable window only"); DNA items tradeable with per-sample quality roll (corpse-quality precedent); combining endpoint (2 DNA → enhanced pet bound to engineer; outcome quality = max(parent qualities) + 5×Engineering-tier capped at template max — PROVISIONAL formula, confirmation requested with item 7); engineered pets use uniform ownership (no CH-command synergy — flagged future); tissue sampling (corpse window, Bio-gated) as tradeable resource + optional tissue slot in new elite armor schematics (marginal durability/protection, magnitudes in §6.9); Bio XP per sample (provisional 50 — §6.4).
- **E6 — Elite services**: Doctor buff-pack schematics (pool/magnitude/duration consumables — stim pattern + buff legs); Musician/Dancer elite BF scaling + performance group buffs (buff machinery, new source); Image Designer appearance-edit handshake (invite-pattern reuse; per-field tree mapping provisional) + holoemote customs (dynamic per-player emotes, flagged); Smuggler slicing handshake (one-time flag + fixed boost provisional) + spice schematics (buff consumables); Commando heavy schematics (**single-target + HAM-cost representation per 9.7.2 bands — CONFIRMED conservative option, no splash invention**) + ammo consumables (charged auto-consume hook — flagged); BH unique-weapon schematics (generic names — DL-44 is SW vocabulary, never enters); TKA unarmed scaling + meditation regen (flagged; **TKA cost 120 CONFIRMED** — §33.3 multi-prereq rule over §8.3.15's 80); style-tagged weapon schematics per combat elite (new `weapon_style` field — old sidearm untagged, non-breaking); cures OUT (no states); AoE heals/shields/combos/flourishes OUT (ability system residual, §5).
- **E7 — Ranger + camps + Merchant hooks**: camps (timed non-structures, wilderness-only incl. outside cities, HAM-regen group buff radius, 30–60 min [ASSUMPTION] / 5 min fast-cycle, Scout-reduced, non-stacking, 24.5 persistence; camping-XP per member tick provisional); creature tracking (directional waypoint to nearest matching creature — survey pattern reuse; player-tracking OUT per 25.6); Merchant sale-XP (1 XP/100 cr + unique-buyer log diminishing), fee curve to −50% (applies in upkeep formula — flagged), slot bonus (+10/tier provisional); Architect civic-deed schematics (cantina/hospital/shuttleport kinds — completes Phase 7 rank flags); Armorsmith/Weaponsmith elite schematics (better bands, gated; existing artisan schematics untouched).
- **E8 — `phase9test`** (3–4 clients): elite train (prereq gates incl. refusal paths; 250-cap accounting) → lair damage/destroy + regen/relocation observed → crafted-tool survey (imprecision + refine) → mission accept/complete across 4 MVP types + bounty contract post/kill/pay → tame/command/release + CL-cap/tamable refusals → camp deploy/buff/XP → ID appearance edit + holoemote → slice/spice/heavy-ammo spot checks → Merchant XP/fee/slots → ranks/leaning untouched; plus re-run of ALL Phase 0–8 tests per verify-don't-claim.

## 5. Scope OUT (explicitly deferred, with homes and reasons)

Chef / Tailor / Miner (no 8.3 entries — unspecifiable; chains-only mentions); Droid Engineer (31.5 EXPANSION — needs pet-AI-equivalent + programming); Politician-full + Squad Leader (31.5 EXPANSION, no §30 home); loyalty/packs/boss uniques (DNA/enhanced variants + tissue are IN per amended §9.3 — no longer in this list); player-tracking/traps/camo (25.6 EXPANSION); master-tier tools + Ranger lair-detect (26.6 EXPANSION); organic-yield expansion (no consumers yet — avoids dead content); special attacks + combos + flourishes + AoE (ability system 9.2.4 unbuilt — **the major residual; future Ability-System phase is human-directed**); cures (no states); shields; Assault + NPC-audience missions (no NPCs — blocked with reason); group mission sharing (needs group-bonus design); 27.2.5 logout-camping channel (distinct system — do not conflate with Ranger camps); free-roam spawns; factories/R8 (EIC/MFG — the factory gap persists post-Phase 9, recorded again); stalls/escrow (unblocked but unassigned — human direction); TEF/duels observations (Phase 8 scope); Jedi/space/quests/instancing (§32, never).
**GDD-internal tensions recorded (human resolves):** (a) §30-seven vs §31.4-fourteen vs §31.5-deferred vs "full roster" exit — §4 option (b) recommended; (b) TKA 80 (§8.3.15) vs 120 (§33.3 multi-prereq) — recommend 120; (c) fork 66-pt basics vs GDD ~40 — fork governs the testbed (both NON-CANONICAL); (d) Bio-Engineer named by §30, deferred by §17.6 — AMENDED per §9.3 decision: fuller DNA mechanics IN (explicit override of the 17.6 deferral for the testbed); loyalty/packs/bosses remain OUT; (e) BH terminal EXPANSION per 16.6 but §15-depth now met — recommend IN player-bounty-only; (f) Sample Collection unmentioned in 16.6-MVP but §31.4 names only "Delivery" — recommend IN (trivially reachable, human confirms); (g) mission faction-point rewards never applicable in MVP ("if applicable" vacuous — flagged zero, not dropped).

## 6. Numeric authorization requests (values NOT set by this proposal)

GDD-GIVEN (implement as written): tool precisions ±500/±250/±100 m ([ASSUMPTION]-tagged where noted); yields 1 / 2–3; cooldown 10 s; corpse window 10 min [ASSUMPTION]; camp 30–60 min [ASSUMPTION]; mission refresh 5–10 min / expiry 2 h / 8–15 listings / log cap 3–5 [ASSUMPTION] / 50% other-destroyed [ASSUMPTION]; lair 3–10 pop + 12–24 h relocation window; mission credit band 500–10k (12.2.1); tag/tame CL-20-at-Master + 3-pet cap; 250 cap (fork-verified); XP tier costs (exist). GDD-SILENT / INDEPENDENT (each needs approval OR interim-[ASSUMPTION] sign-off): (1) elite SP distributions (80-class 3/4/5/6+M8; 120-class 4/6/8/10+M8); TKA 120; (2) crafting XP complexity mapping + repeat-50% + crit-2x windows; (3) style-pool kill awards; (4) tracking/camping/slicing/taming/sample XP awards + DNA sample XP 50 + combine formula (max(parent Q) + 5×tier, capped); (5) Merchant diminishing windows + fee curve + slot bonus; (6) mission credit/XP in-band mappings + slot cap 4 + refresh 7 min + listings 10; (7) lair respawn tick 30 min (HISTORICAL-USER per owner 2026-09-16; refill shape TBD) + destroy bonus 150 + relocation 12 h / 10 min fast; (8) tame chance/caps/pet caps/camp radius + buff magnitudes; (9) slicing boost + spice/buff-pack magnitudes + heavy HAM costs + ammo stocking + tissue-slot armor bonus; (10) contract rules (escrow/refund/pay-all/self-claim); (11) ID field mapping + holoemote pricing shape; (12) triangulation halving rule + Scout/Ranger bonus magnitudes; (13) fast-cycle camp 5 min. PROPOSED interim convention (approval requested): NON-CANONICAL flagging on every non-GDD number (code + HYGIENE_NOTE.md), magnitudes anchored to fork-local scales (deed 15M+8P, kill-drop 10×CL, upkeep/fee bands, XP award bands).

## 7. Acceptance criteria (verify, don't claim)

`go build ./...` + `go vet ./...` clean; fresh DB; phase0–8 tests green PLUS new
`phase9test` end-to-end (elite trains with prereq refusals + cap accounting;
lair damaged→destroyed→relocated with regen observed; crafted-tool survey shows
imprecision then refinement; 4+ mission types accepted→completed→rewarded with
ledger pairs; bounty contract posted→BH-killed→paid from escrow; creature tamed
(CL-cap/tamable refusals sampled) → commanded → released; DNA sampled →
combined → enhanced pet + tissue slot spot check; camp deployed
(wilderness-only refusal inside city sampled) → buff + XP observed; ID edit +
holoemote performed; slice/spice/heavy spot checks; Merchant XP + fee + slots
observed) with real pasted output; fast-cycle mapping test-config-only, defaults
GDD-given; hygiene note extended; commit per verified package (E1→E8).

## 8. Risks and dependencies

- **Size**: the largest phase (roster + lairs + surveying + missions + pets +
  services + camps). Mitigation: verify-per-package commits (E1→E8); fallback —
  split into 9a (E1–E3 + E7-Merchant/Scout hooks) / 9b (E4–E7) at the gate (human
  decides, §9).
- Lair regen interacts with Phase 3/6 dynamics (more creatures over time):
  full-suite regression is the detector (verify-don't-claim); new-instance IDs
  must not collide with test enumerations.
- Pet AI + camps + contracts are the three biggest new-state surfaces; each
  rides existing machinery (retaliation ticks, buffs, invite/escrow patterns) —
  no new tick loops proposed (mission sweep is lazy).
- Arranged content (bounty collusion, mission sharing by alts): observed via
  events/ledger, unenforced (OQ-012 TBD — no invented throttles).
- EIC/BAL boundaries: every §6 number needs approval or interim sign-off before
  code; EIC architecture, GDD edits, and status changes are untouched.
- Ordering: Phase 5/6/7/8 review gates MUST pass before E1 starts (E4–E7
  terminate in their systems).
- Scope-creep magnets: ability system, factories, DNA depth, TEF/duels
  leftovers, Chef/Tailor invention — all §5-excluded; enforce at review.

## 9. Decisions requested of the human (the plan asks; it does not decide)

1. Base/location (§3) — confirm building on `testbed/swg-phase3-combat/`.
2. Scope reconciliation (§§1/4/5) — confirm option (b) (recommended), (a)
   §30-literal, or (c) split 9a/9b.
3. Bio-Engineer shape (§5.d) — confirm roster-only vs fuller DNA options.
4. Commando AoE + TKA cost + fail-aggro taming (§§4–6) — confirm recommendations
   or supply.
5. Numeric set (§6, 13 groups) — approve magnitudes/mappings or supply.
6. Unassigned leftovers (§§2.2/5: stalls/escrow follow-up, PvP-combat-XP
   residual, service-XP scaling residual) — direct where each goes.
7. Authorisation to implement on approval of 1–6 (cf. §10).

## 10. Implementation status

DECISIONS 1–6 RECORDED 2026-09-16 (no code written; item 7 NOT authorized —
implementation remains explicitly pending):
1. Base/location CONFIRMED (`testbed/swg-phase3-combat/`).
2. Scope option (b) CONFIRMED (full spec'd roster + reachable mechanics; §4/§5
   stand as written except the §9.3 amendment below).
3. Bio-Engineer FULLER DNA CONFIRMED — §§4 (E5), 5(d), 6(4/9), 7 amended
   accordingly (DNA sampling + combining + enhanced variants + tissue sampling
   and tissue slots; loyalty/packs/bosses stay OUT). This explicitly overrides
   the 17.6 EXPANSION deferral for the testbed. The DNA combine formula and DNA
   XP award in §6.4 are provisional and ride with the item-7 authorization
   request (no separate numbers supplied).
4. CONFIRMED as recommended: Commando conservative single-target heavies (no
   splash invention); TKA cost 120 (§33.3 over §8.3.15); fail-aggro taming.
5. Numeric set APPROVED as proposed (all 13 §6 groups, magnitudes/mappings
   unchanged, plus the §6.4/§6.9 DNA additions above).
6. Leftovers EXPLAINED (no implementation directed — placements recorded here):
   (a) stalls/escrow follow-up → standalone commerce package AFTER Phase 9
   (its mail blocker cleared in Phase 7; §30 never listed it; needs its own
   mini-proposal for commission/fee/expiry numerics); (b) PvP-combat-XP residual
   (GDD 7.2.1 grants combat XP on PvP kills; Phase 8 gave points only) →
   Phase 8 amendment micro-package (one hook + amount decision), sequenced after
   Phase 9 or beside (a), separately authorized; (c) service-XP scaling residual
   (flat-50s vs GDD per-point/per-tick) → Phase 10 balance pass (rate redesign
   belongs with Balance/Telemetry/Polish + §33 targets); provisionals stand.
AMENDMENT RECORDED 2026-09-16 (respawn datum — APPROVED): the §4/E2 30-minute
   production tick is set aside for testbed purposes (owner direction: correct for
   production, unusable for the test). Testbed E2 regen = provisional +1 instance
   per harvest tick while living < max (as built in `tickLairs`); exact refill
   shape (incremental vs full repopulation) stays TBD. The 30-minute datum stands
   as HISTORICAL-USER input, not as an implemented value.
7. PENDING — no authorization given. Next step is an explicit item-7
    authorization (covering §§1–6 as amended); until then no code is written.
   On authorization, execution follows the verify-don't-claim methodology
   (`go build` + `go vet` clean, fresh DB, all phase tests re-run old+new with
   real pasted output, hygiene note extended per package, commit per verified
   package E1→E8), with the human review gate before any Phase-10-equivalent
   work (HD-TST-01 pattern).

*End of proposal v0.1. To enact: owner approves (or amends) §§3–6/9–10; approval and date
recorded before any code is written.*
