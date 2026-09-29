# SWG Pre-CU Phase 6 — Social Support Professions: Proposal Plan v0.1

**Status:** PROPOSED plan. No authority. No design decision. Human approval required
before any implementation, and separately before any numeric value is treated as tuned.
**Author / Assessor:** Muse Spark (`opencode/muse-spark-1.3-contributor-free`) via OpenCode —
research and planning only; all authority decisions remain human.
**Date drafted:** 2026-09-14.
**Controlling references:** `AGENTS.md`; `proposals/SWG_Phase5_Economy_Infrastructure_Proposal_v0.1.md`
(executed — review gate pending); `testbed/swg-phase3-combat/HYGIENE_NOTE.md`;
predecessor GDD `sources/swg-pre-cu/SWG_PreCU_GDD.md` (§§7, 8.2.5, 8.2.6, 8.3.16–8.3.20,
9.2.1, 9.2.4, 9.2.6, 9.6.1, 10.9, 18.2.5, 23, 24 — HISTORICAL, never TCIndustries canon).
TCIndustries fence: SERV-001 (meaningful non-combat services, LOCKED), VIS-004/PIL-003
interdependence, LOOP-002 non-combat viability, PROF-001/PROF-004; EIC untouched
(HD-EIC-07); provenance stays deferred (conveniently, GDD 23.2.3 wants NO formal
reputation score either — agreement recorded, not authority).

---

## 1. Objective

Implement **Phase 6 — Social Support Professions** on the testbed fork (Phases 0–5
complete there, subject to the pending Phase 5 review gate), satisfying the
predecessor-GDD §30 exit criteria verbatim: *"A combat player can accumulate Battle
Fatigue and Wounds through Phase 3's combat loop and have them healed by a second player
performing Section 23/24's mechanics — demonstrating Pillar 1's core interdependence
claim end-to-end."* This is the first phase whose exit criterion is irreducibly
two-player — the interdependence pillar made literal.

Entry criteria (§30: Phase 3 for the wound/BF source, Phase 5 for venue adjacency)
are MET in the fork (combat loop verified; structures exist), conditional on the review
gate — see §10. Note: "venues" in the fork means placed structures/anywhere, not
cantinas (no cities exist) — §5 covers the resulting simplifications.

## 2. Research basis

### 2.1 Predecessor-GDD map (HISTORICAL input, not authority)

| GDD section | Specifies | Phase-6 relevance |
|---|---|---|
| 23.2.1 performance | Routines from skill trees; venue-or-busk ([ASSUMPTION] reduced); stationary + vulnerable; watch-target + tick effects; Mind cost; interrupt on combat/incap | Perform/Watch session loop |
| 23.2.2 BF healing | Per-tick BF heal scaled by Music/Dance tier + quality roll; elite Mind-regen buffs stackable w/ Doctor, not w/ same-type Entertainer | BF tick heals + buffs |
| 23.2.3 tips | Direct credit tips any time; small XP bonus (§7.2.3); NO formal reputation (18.2.5) | Tip transfer + XP |
| 23.2.4 AFK watching | Historically accurate passive watching preserved; presence required, no rate difference — NOT an exploit to fix | Implemented by omission (no anti-AFK) |
| 23.3 entities | PerformanceSession, WatchSession, shared BuffInstance (source-profession enum incl. food — Chef-adjacent, OUT) | Persistence shapes |
| 23.5 edges | Incap ends performance cleanly; leaving radius ends watch; one-watch-target (no double-dip); PvP-zone performing allowed | Rules to implement |
| 24.2.1 wounds | Medic + tool → flat-plus-skill-scaled wound reduction; 2–5 s cast [ASSUMPTION] (Combat Medic mobile exception); cooldown/diminishing window | Wound-heal action |
| 24.2.2 revive | Combat Medic targets incapped player, cast + Mind cost → 10% pools + wounds/BF; revive must beat the 5-min timer, no grace | Revive gating (tension §5) |
| 24.2.3 buffs | Doctor Enhancement packs +500–3000 HAM (6.2.2 scale); out-of-combat, stand-still application; 30–60+ min [ASSUMPTION]; same-type non-stacking | Buffs (gating tension §5) |
| 24.2.4 cures | Disease/poison cures consume crafted components | OUT — no 9.2.6 states exist in fork |
| 24.3 entities | HealingAction (with per-pair cooldown), MedicalCenter (needs cities — OUT) | Action log shape |
| 8.2.5/8.2.6 basics | Medic trees (Healing/Injury/Medicine/Support); Entertainer trees (Music/Dance/Healing/Showmanship); ~40 SP; Medical + Entertainer XP types | Skill gates (trees exist as fork data) |
| 8.3.16–20 elites | Doctor/Combat Medic/Musician/Dancer/Image Designer (require Master bases) | DEFERRED with two tensions (§5) |
| 9.2.1/9.6.1 | Wound/BF mechanics; revive-at-10% + penalties (ALREADY in fork) | The demand side — no changes needed |
| 10.9 stim precedent | Seed-demo stimPackB only (healing/charges/potency) — pattern, not code | Stim schematic shape |

### 2.2 Existing implementations (evidence, not authority)

- **Fork**: Medic/Entertainer 18-box SKILL TREES as data only (gap row 12); wound/BF
  columns + EffectiveMax + revive/clone all live; revive UNGATED (any-nearby-player
  interim, flagged since Phase 3); NO heal actions, NO buffs, NO performance/watch,
  NO stim items or schematics, NO tips/XP. This phase fills exactly those gaps.
- **Track A/B**: zero service code (matches only inside GDD copies). **Chat export**:
  6 mentions, negligible. Nothing transfers.

## 3. Base and working location (RECOMMENDED — human confirms)

Build in place on `testbed/swg-phase3-combat/` (new `internal/services/` or extension of
`internal/combat/` + DB tables + handlers + `cmd/phase6test/`); no new fork, `sources/`
untouched. Rationale: every Phase 6 flow terminates in Phase ≤5 systems (HAM state,
revive path, crafted stim schematics, credit tips via wallet, XP pools).

## 4. Scope IN (work packages, each verify-by-running before the next)

- **S1 — Wound healing** (Medic-gated on `medic_novice`): targeted action, flat +
  Healing-tree-scaled wound reduction, per-healer-target cooldown/diminishing window,
  instant application (2–5 s cast choreography deferred — flagged simplification;
  Combat Medic mobility nuance deferred with elites). HealingAction log rows.
- **S2 — Revive gating**: `revive` (and only revive) now requires `medic_novice`,
  closing the Phase 3 any-player interim. Base-Medic — NOT Combat Medic — does the
  reviving: GDD 8.2.5's own benefits list grants base Medic revive, against 9.6.1/24.2.2
  naming Combat Medic (tension §5; base-Medic recommended, elite gate deferred).
- **S3 — Buffs**: Medic-gated basic HAM-pool buffs (magnitudes inside the GDD +500–3000
  band, low end; durations inside 30–60+ min) as BuffInstance rows with wall-clock
  expiry (persists logout/login per 24.5); same-type non-stacking incl. cross-source;
  out-of-combat + stand-still application checks. Doctor-elite buffs deferred.
- **S4 — Perform/Watch/BF-heal** (Entertainer-gated): session start/stop, watch-target
  pairing, per-tick BF reduction scaled by Healing-tree boxes with Music-vs-Dance tree
  differentiation (trees exist as data — captures 23.6's Musician/Dancer demand without
  elite professions); Mind drain per tick on performer; interrupt on combat/incap;
  one-watch-target rule. Venue = anywhere (busking fallback per the GDD's own
  [ASSUMPTION]) at the reduced rate; venue-bonus structures deferred.
  AFK-watching preserved by omission (23.2.4 — recorded, not an oversight).
- **S5 — Stim packs**: one generic stim schematic (registry addition, `medic_novice`
  gate — first non-Artisan schematic gate, flagged pattern extension); charges +
  potency props; use-action heals wounds per potency (consumes a charge). Cure actions
  OUT (no disease/poison states to cure — recorded, not dropped silently).
- **S6 — Tips + service XP**: direct credit tip transfer (wallet→wallet, ledger-tagged
  transfer) + small Entertainer/Medical XP awards (provisional). No reputation score
  (GDD 23.2.3 agrees with deferred provenance).

## 5. Scope OUT (explicitly deferred, with homes and tensions)

Disease/poison cures (no 9.2.6 states); flourishes (23.6 EXPANSION); Image Designer
(EXPANSION); MedicalCenter/clone-facility link (needs cities); Combat Medic mobility
nuance; cast-time choreography; venue structures/bonuses; Chef/food buffs; Militia/
Politician-adjacent content. **Tensions recorded (human resolves):** (a) revive gate —
base-Medic (recommended, 8.2.5's own list) vs Combat-Medic-literal (needs unbuilt
elites); (b) buff gate — base-Medic basic buffs (recommended) vs full deferral to
Doctor elites; (c) Musician/Dancer MVP demand (23.6) vs elite-deferral pattern —
resolved via tree-differentiation without elite professions (recommended).

## 6. Numeric authorization requests (values NOT set by this proposal)

GDD-GIVEN (implement as written): buff band +500–3000 and 30–60+ min durations
([ASSUMPTION]-tagged, implement flagged); revive-at-10% + wound/BF penalties (exist);
tip-XP existence (§7.2.3, magnitude silent → provisional); no-reputation rule.
GDD-SILENT / INDEPENDENT (each needs approval OR interim-[ASSUMPTION] sign-off):
(1) wound-heal flat + per-tier scaling + cooldown/diminishing windows;
(2) BF-heal per-tick rates + Music/Dance/tree multipliers + busking reduction;
(3) performer Mind drain per tick; (4) watch radius; (5) stim schematic costs +
heal-per-potency + charges; (6) buff magnitudes/durations within GDD bands;
(7) tip XP + service XP awards; (8) cast-time simplification (instant + cooldown).
PROPOSED interim convention (approval requested): Group-C-style NON-CANONICAL
flagging on every non-GDD number (code + HYGIENE_NOTE.md), magnitudes anchored to
fork-local scales (combat damage bands, fee bands, XP award bands) to avoid fresh tuning.

## 7. Acceptance criteria (verify, don't claim)

`go build ./...` + `go vet ./...` clean; fresh DB; phase0–5 tests green PLUS new
`phase6test` with TWO clients (A fighter + B service holding Medic/Entertainer
Novice): A takes wounds + BF through the combat loop → B wound-heals A (wounds drop,
cooldown observed on immediate re-heal) → B performs, A watches (BF drops over ticks,
Mind drains on B) → B tips flow A→B with ledger + XP → B buffs A (pools rise, expiry
persists a relog) → A incaps → B (and only B — ungated stranger C rejected) revives →
all with real pasted output; fast-cycle mapping test-config-only for any duration
compression; hygiene note extended; commit per verified package (S1→S6).

## 8. Risks and dependencies

- Tick-rate coupling: BF-heal and Mind-drain ride the 1 s world tick — balance between
  tick granularity and test duration (a 0.5%-per-tick heal needs hundreds of ticks to
  show; tests must run minutes, like Test 5's incap wait — accepted cost).
- Ungated-revive closure changes Phase 3 behavior: phase3test's revive-agnostic flow
  (it clones, never revives) is unaffected — but any shared-DB ordering effects get
  full-suite coverage as usual.
- Service-XP placeholder endpoints mirror the Phase 2 earn-xp pattern (provisional,
  flagged) — real XP sources (patients healed, audiences held) need design beyond MVP.
- Ordering: Phase 5 review gate MUST pass before S1 starts; EIC, provenance,
  reputation untouched throughout.

## 9. Decisions requested of the human (the plan asks; it does not decide)

1. Base/location (§3) — confirm building on `testbed/swg-phase3-combat/`.
2. Revive gate (§5.a) — base-Medic (recommended) vs Combat-Medic-literal.
3. Buff gate (§5.b) — base-Medic basic buffs (recommended) vs full deferral.
4. Busking-anywhere + tree-differentiation (§§4/S4, 5.c) — confirm or amend.
5. Stim schematic + medic gate (§4/S5) — confirm first non-Artisan gate.
6. Numeric set (§6) — approve magnitudes or supply.
7. Authorisation to implement on approval of 1–6 (cf. §10).

## 10. Implementation status

APPROVED AND EXECUTED 2026-09-14/15. Owner decisions (§9) received with no
amendments: base confirmed, base-Medic revive gate, base-Medic basic buffs,
busking-anywhere + tree-differentiation, stim schematic with medic gate, numeric
set approved, implementation AUTHORIZED (also covering the Phase 5 review-gate
precondition for this line — owner may still call the gate separately).
Results on `testbed/swg-phase3-combat/`: `go build` + `go vet` clean; fresh-DB
phase6test 7/7 ALL PASS (real output) — incap → stranger-rejected → medic revive
(wounds 125/BF 2.0) → heal 125→100 + cooldown rejection + medical XP → BF→0 by
watching + Mind drain → 100-credit tip + XP → +500 buff (max 1400) + no-stack
rejection + relog persistence → stim craft/use (wounds → 0).
Bugs found by running: shared-RNG tick-loop death (per-call instances now);
seed-time type gaps frozen by same-type relocation (per-type guarantee seeding
now); missing revive-XP award; timestamp parse shapes. HYGIENE_NOTE.md extended
(§Phase 6 additions + correction log); species-ID + Godot known-remainings
unchanged. Full 7-suite regression: see run transcripts.
Next: human review gate before any Phase-7-equivalent work.

*End of proposal v0.1. To enact: owner approves (or amends) §§3–6/9–10; approval and date
recorded before any code is written.*
