# TCIndustries — Deferred BAL-001 Numeric Pass: Disposition (2026-10-25)

- **Author:** Claude (Anthropic)
- **Version:** 1.0
- **Date:** 2026-10-25
- **Status:** **DISPOSITION — no magnitudes set.** This document answers the request to "do the deferred BAL-001 pass" (the magnitudes the OQ-001 BAL-001 numeric pass of 2026-10-25 deliberately deferred: XP earn/progression rate, task-layer reward, task-layer sink). It records that, on inspection, **none of the three can be responsibly authorised yet**, and states the concrete reason for each. It authorises *no* numeric value and *no* code.
- **Companion:** `proposals/TCIndustries_OQ-001_BAL-001_Numeric_Pass_2026-10-25.md` (the pass that set the OQ-001/HD-PROF-04 magnitudes and listed these three as deferred).

---

## 0. Why this is a disposition, not a set of numbers

The original deferral was grounded on "structural partners do not yet exist." Re-investigation on 2026-10-25 shows that premise is now **partially** true and **partially** false:

| Deferred magnitude | Original deferral reason | State now (2026-10-25) | Set now? |
|---|---|---|---|
| XP earn / progression rate (PROG-003) | model not stable/testable | Earning channels exist & are Phase-8-verified, **but** PROG-003 is still `TBD` and the earn→spend pool loop is internally inconsistent | **No — see §1** |
| Task-layer reward (Q8) | no task layer / matched sink | Still no task layer in the fork (HISTORICAL only) | **No — see §2** |
| Task-layer sink (Q8 / HD-ECO-01) | no matched sink exists | Still no task layer ⇒ no sink | **No — see §2** |

BAL-001 forbids inventing magnitudes where the base is undefined or inconsistent. Setting values here would be false precision, so this pass **disposes** rather than **authorises**.

---

## 1. XP earn / progression rate — DEFERRED (new reasons)

### 1.1 PROG-003 is still `TBD`
The XP earn rate is a component of the **Skill Acquisition Model**, which GDD §655 holds at **Status: TBD** ("Exact skill acquisition … remains unresolved"). The OQ-001 numeric pass deferred this magnitude explicitly "until the model exists." The model still does not exist. Authorising an isolated earn rate without a ruled acquisition model contradicts the deferral's own condition.

### 1.2 The earn → spend pool loop is internally inconsistent (must be resolved first)
The fork's XP plumbing accepts **any** pool string (`AddCharacterXP` is untyped: it `INSERT … ON CONFLICT … DO UPDATE` on `(character_id, xp_type)`), while the **train** path deducts from a *fixed* set of six pools derived from `ProfessionDef.XPType` (`combat, crafting, scouting, medical, entertainer, merchant` — the six the OQ-001 pass authorised).

Concrete mismatches found in the live fork:
- `internal/crafting/crafting.go` recipes award `XPPool` values `weapon_crafting`, `armor_crafting`, `smuggler`, and `medical` — sub-pools that are **not** spendable, because no box deducts from `weapon_crafting`/`armor_crafting`/`smuggler`.
- `internal/handlers/civic_http2.go` awards a `"mentoring"` pool (via `civic.MentorshipXP = 50`) — also not among the spendable six.
- The generic test endpoint `POST /api/characters/{id}/earn-xp` likewise accepts any of the six by validation but the real channels above write pools outside that set.

**Consequence:** XP earned through crafting/mentorship currently accumulates in pools that **cannot be spent to train any box**. Authorising an earn *rate* on top of this would enshrine a loop where players grind XP they can never convert — exactly the "false precision" BAL-001 warns against. The pool taxonomy must be reconciled (earn pools == spendable pools, or a documented mapping) before any rate is set.

### 1.3 Combat-rate scale (informational only, NOT a ruling)
The one channel that does feed a spendable pool is combat: `AddCharacterXP("combat", 50*tmpl.CLMax)`. The live bestiary (`internal/creatures/creatures.go`) spans CL 1–3 only, so a kill currently yields 50 / 100 / 150 combat XP against a per-profession training cost of **118,000 XP**. This is presented so the eventual PROG-003 ruling has the real ratio in hand; it is **not** an authorised rate.

**Resolution required before this magnitude can be set:** (a) PROG-003 moved from `TBD` to RULED; (b) XP pool taxonomy reconciled so every earn pool is spendable. Then a PROG-003 + BAL-001 pass sets the earn rates.

---

## 2. Task-layer reward & sink — DEFERRED (original reason still holds)

### 2.1 No task layer exists
The Phase 9 replacement proposal (`proposals/TCIndustries_Phase9_Elite_and_TaskLayer_Proposal_v0.1.md`, §2.1) defines the task layer **structurally only** — repeatable system-generated contracts with a *reward hook* and no magnitude. There is **no task layer in the fork**: a tree-wide grep for `taskLayer`/`task layer`/`phase9`/`mission`/`quest` returns matches only inside `SWG_PreCU_GDD.md`, which is HISTORICAL reference material explicitly out-of-scope per HD-TST-01. There is no hook to attach a reward magnitude to and no magnitude to set.

### 2.2 HD-ECO-01 blocks a reward without a matched sink
Even if a reward *hook* existed, HD-ECO-01 (faucet-exceeds-sink) requires a **matched sink** before any reward faucet magnitude is authorised. No task-layer sink is designed. Authorising a reward magnitude here would violate HD-ECO-01.

**Resolution required before these magnitudes can be set:** (a) the task layer built (against the Phase 9 structural definition); (b) a matched sink designed; then a Q8 + BAL-001 pass sets reward and sink magnitudes together.

---

## 3. Status of the three deferred magnitudes

| Magnitude | State (this disposition) |
|---|---|
| XP earn / progression rate (PROG-003) | **DEFERRED** — PROG-003 still `TBD`; pool taxonomy inconsistent (§1) |
| Task-layer reward (Q8) | **DEFERRED** — no task layer exists (§2) |
| Task-layer sink (Q8 / HD-ECO-01) | **DEFERRED** — no task layer, no matched sink (§2) |

This disposition does **not** change any GDD status, ruling, or authority line. It clarifies the reason for the continued deferral so the project does not carry a stale "outstanding BAL-001 pass" line that is in fact unblocked-and-broken rather than merely pending.

---

## 4. Recommended next decisions (owner-gated)

1. **Reconcile the XP pool taxonomy** in the fork so every earn pool is spendable (or document a mapping). *Code/design change — not a numeric tuning act, but a prerequisite for any rate authorisation.*
2. **Rule PROG-003** (the Skill Acquisition Model) — move GDD §655 from `TBD` to a ruled model. *Owner ruling.*
3. **Only then**, a PROG-003 + BAL-001 pass sets the XP earn rates, grounded in the real cost-curve ratios (§1.3).
4. **Build the task layer** (against Phase 9 §2.1) **and** design its matched sink, then a Q8 + BAL-001 pass sets reward/sink magnitudes together.

*End of disposition (no magnitudes authorised, 2026-10-25).*
