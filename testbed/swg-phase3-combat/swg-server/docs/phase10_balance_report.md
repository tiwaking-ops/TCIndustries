# Phase 10 Balance Validation Report (B3)

Generated: 2026-09-16T20:26:47Z

Author: Buffy — Codebuff agent, Phase 10 implementation (proposal v0.2 §3/B3).

> **Status discipline:** every fork value below is OBSERVED from the live DB,
> read from fork source with a file:line citation, DERIVED with the formula
> shown, or honestly NOT-MEASURABLE. The HISTORICAL §33 bands are targets to
> validate against, never tune-to (BAL-001; Master GDD "Inherited from prior
> versions" item 3). A gap recorded here is EVIDENCE, not a work item: any fix
> is a separate numeric-authorization decision (proposal §5.1).

| Area | Item | Fork value | HISTORICAL target | Verdict | Note |
|---|---|---|---|---|---|
| Combat §9.7 | unarmed/sidearm damage (start scale) | 5–15 per swing (unarmedWeapon, handlers/combat.go:29 [PROVISIONAL]) | 50–150 start band (predecessor GDD §9.7, HISTORICAL) | OUT-OF-BAND | fork scale sits BELOW the start band by 10× — a DPS-parity gap is not a failure (BAL-002 lens); recorded as evidence |
| Economic §12.2.2 | creature-drop cap | 10 × CLMax, capped 200 (handlers/combat.go:312-314 [PROVISIONAL]) | 10–200 band (§12.2.2, HISTORICAL) | IN-BAND | fork construction lands inside the band by design |
| Economic §12.2.1 | mission reward band (destroy-lair, CL1) | 2000 credits (= 1500 + 500×CL1, missions.go:183 [PROVISIONAL]) | 500–10,000 band (§12.2.1, HISTORICAL) | IN-BAND | low-CL floor sits inside the band |
| Economic §12.2.1 | mission reward band (destroy-lair, CL10) | 6500 credits (= 1500 + 500×CL10, derived) | 500–10,000 band (§12.2.1, HISTORICAL) | IN-BAND | upper-CL construction stays inside the band |
| Economic §12.2.1 | mission reward band (delivery, qty 2–4) | 1000–2000 credits (= 500×qty, missions.go:193 [PROVISIONAL]) | 500–10,000 band (§12.2.1, HISTORICAL) | IN-BAND | inside the band |
| Economic §12.2.1 | mission reward band (sample / recon) | 750 / 500 credits (missions.go:201,214 [PROVISIONAL]) | 500–10,000 band (§12.2.1, HISTORICAL) | IN-BAND | inside the band |
| Faction §15.2.3 | ranked faction thresholds | NOT-BUILT: fork has alignment declaration + overt flagging only; no ranked point thresholds exist | 2,500 / 10,000 / 30,000 / 75,000 (§15.2.3, HISTORICAL) | NOT-MEASURABLE | recorded: nothing to compare — an unbuilt mechanic, not a gap (ability-system-class residual) |
| Economic §12.4 | 30-day net inflation window | OBSERVED: circulation=18450 faucet30d=0 sink30d=57 net=-0.31% (snapshot at 2026-09-16T20:26:35Z) | ~5–10% annualized (§12.4, HISTORICAL) | OBSERVED | testbed windows are hours old, not 30 days — read as pipeline proof, not as an inflation measurement |
| Wealth §34.3 | credit distribution (diagnostic Gini) | OBSERVED: players=4 total_credits=18450 gini=0.047 | no §33 target — Gini is diagnostic only (§34.3) | OBSERVED | tracked, never target-capped, never a lever |
| Economic §33.4 | maintenance share of weekly income/sink | OBSERVED: maintenance_fee=7 of 57 total sink over trailing 7d (12.3%) | 5–15% of weekly income (§33.4, HISTORICAL) | OBSERVED | fresh-world windows are dominated by training sinks; the share is recorded, not judged |
| Economy §34 | crafted provenance share | OBSERVED: 1/1 items carry schematic provenance (100%) | ≥95% crafted (§1.3/§34, HISTORICAL) | OBSERVED | raw counts beside the percentage (small-cohort rule) |
| Progression §33.3 | basic mastery 40–60 h / elite 80–120 h / full cap 200–300 h | NOT-MEASURABLE live: XP rates are [PROVISIONAL] flat awards (service 50/25/100, combat 50×CLMax) and fast-cycle testbed cadences distort any wall-clock projection | 40–60 h basic / 80–120 h elite / 200–300 h cap (§33.3, HISTORICAL) | NOT-MEASURABLE | honest projection requires the B4 rate redesign (proposal §6); characterized in the B4 options paper, not estimated here |
| Combat §9.7 | PvP 1v1 TTK 30–60 s; solo white-con 20–40 s | NOT-MEASURABLE from this report: TTK is behavior, observable only from live-resolved sparring data (phase10test Proxy 2 records per-swing outcomes at the observed 5–15 dmg scale) | 30–60 s PvP / 20–40 s solo (§9.7, HISTORICAL) | OUT-OF-BAND (derived) | at 5–15 dmg vs 1000 health the arithmetic TTK is ~2–3 minutes, far outside 30–60 s — recorded as DERIVED observation, no fix |

## Reading this report

- **BAL-002 lens:** a DPS-parity gap is not a failure; a collapsed profession
  category is. Out-of-band rows are recorded observations.
- **NOT-MEASURABLE rows are honest exits** (§2.1 disjunction): instrumentation
  plus an honest failure signal satisfies the Phase 10 exit criteria.
- **No numeric changes** were made by or proposed inside this report.

## B4 — Service-XP characterization (Phase 9 decision 6c residual)

All service-XP awards in the fork are flat [PROVISIONAL] constants — the
same award regardless of target state, magnitude, or repetition. Sites
(verified 2026-09-17):

| Pool | Award | Shape | Site |
|---|---|---|---|
| medical | 50 | per wound-heal action | services.go:136 (services.HealXP=50) |
| medical | 50 | per buff application | services.go:190 (services.BuffXP=50) |
| medical | 25 | per stim use | services.go:427 [PROVISIONAL] |
| medical | 100 | per successful revive | combat.go:458 (services.ReviveXP=100) |
| entertaining | 25 | per tip received | services.go:350 (services.TipXP=25) |
| image_designer | 50 | per image-design application | elite_http.go:327 [PROVISIONAL] |
| image_designer | 50 | per holoemote created | elite_http.go:367 [PROVISIONAL] |
| smuggler | 50 | per slice accepted | elite_http.go:512 [PROVISIONAL] |
| (stim pool varies) | 25 | per stim applied | elite_http.go:613 [PROVISIONAL] |
| scouting | 10 | per survey | resources.go:265 [PROVISIONAL] |
| scouting | 25 | per sample | resources.go:347 [PROVISIONAL] |
| scouting | 25 | per corpse harvest | resources.go:503 [PROVISIONAL] |
| ranger | 10 | per occupied camp per harvest tick | ranger.go:101 (CampXPPerMember=10) |
| (camp action pool) | 25 | per camp-related action | ranger.go:176 [PROVISIONAL] |
| combat | 50 | per pet kill assist (flat) | ranger.go:266 [PROVISIONAL flat] |
| creature_handling | 25/100 | per tame attempt / success | pets.go:129,153 (constants pets.go:29-30) |
| bio_engineering | 50 | per DNA sample | pets.go:313 (DNASampleXP=50) |
| mentoring | 50+50 | per mentorship session, both sides | civic_http2.go:599-600 (civic.MentorshipXP=50) |
| merchant | price/100 | per sale — unique buyer full, repeat 25% | economy.go:297-303 (rate-shaped, not flat) |
| structure_crafting | 100 | per house/vendor/base placement | economy.go:157,196; civic.go:96; faction.go:367 |
| combat | 50×CLMax | per kill (reference rate) | combat.go:297 [PROVISIONAL] |

Observed per-pool totals (character_xp, live DB):

| XP pool | characters holding XP | total XP |
|---|---|---|
| combat | 2 | 9000 |
| crafting | 1 | 1000 |
| medical | 1 | 5100 |
| merchant | 1 | 5 |
| scouting | 1 | 1145 |
| structure_crafting | 1 | 100 |

Characterization conclusion: every award above is action-count-driven,
not outcome-driven — repeating the action repeats the XP (the merchant
sale hook is the only rate-shaped exception). This is the flat-rate shape
the Phase 9 decision 6c residual refers to. Any redesign is an option in
docs/phase10_service_xp_options.md and requires explicit numeric
authorization (proposal §6); the default if declined is that these
provisionals stand, characterized by this section.
