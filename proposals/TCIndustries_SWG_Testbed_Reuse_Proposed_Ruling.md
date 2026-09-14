# TCIndustries — SWG Engine Testbed Reuse Ruling (DRAFT for human approval)

**Status:** PROPOSED draft. No authority. Nothing here is HUMAN-LOCKED until the project owner approves the text (as-is or amended) and it is filed in `governance/` with a ruling ID and date.
**Date drafted:** 2026-09-14 (OpenCode documentarian draft, at owner request).
**Responds to:** owner direction of 2026-09-14 selecting path (b) — SWG Pre-CU engine as future mechanical testbed for TCIndustries original IP.
**Controlling references:** `AGENTS.md`; `README.md`; `sources/swg-pre-cu/README.md`; `proposals/SWG_Code_Due_Diligence_Checklist_v0.1.md`; `governance/TCIndustries_Human_Rulings_Register.md` (HD-EIC-01–08 in force); `canonical/gdd/TCIndustries_Master_GDD_v1.1.1_Status_Patch.md`.

---

## Proposed decision (adopts path b)

The project owner authorises reuse of the SWG Pre-CU server codebase (`sources/swg-pre-cu/`, Groups A–C) **as patterns-only mechanical testbed material** for TCIndustries investigation and prototyping.

**Reuse mode authorised:** patterns-only fork (checklist §7 option "reuse patterns only"), subject to the due-diligence checklist and the constraints below. Not "reuse as-is".

## Explicitly does NOT authorise

- No promotion of any SWG mechanic, value, name, or structure into TCIndustries canon. All SWG content remains HISTORICAL; all fork behaviour remains PROTOTYPE/EVIDENCE.
- No numeric tuning of economic or progression systems (standing governance boundary).
- No EIC shaping — HD-EIC-05 (Architecture C retired) and HD-EIC-07 (new shaping only on human-supplied structural boundaries) remain in force.
- No provenance / formal reputation systems (still deferred).
- No GDD edits or status changes by this ruling. CRFT-006 stays PROPOSED, MFG-004 stays DERIVED CONSTRAINT, and all TBDs stay TBD.
- No Star Wars / Lucasfilm names, lore, factions, species, or places in the fork. Known items to strip per checklist step 4: Tatooine, Worrt, Womp Rat, 9-species HAM table, SWG schematic/crate vocabulary (plus anything else the hygiene pass finds).

## Preconditions before new testbed code

1. Execute `SWG_Code_Due_Diligence_Checklist_v0.1.md` steps 1–4 (integrity, build, test, IP hygiene) against `sources/swg-pre-cu/`; record evidence logs. Do not fix failures — record them.
2. Fork `sources/swg-pre-cu/code-phase0-3-claudecode/` (GDD v2 tree, Phases 0–3) to a new location; leave the SWG original unmodified for provenance.
3. De-SWG pass with generic IDs; hygiene note with file refs for every renamed string; all Phase 0–3 tests still passing. Human review gate before Phase 4-equivalent work.

## Scope of testbed work authorised after preconditions

- Generic pipeline exercises only (e.g. survey → harvest → craft-with-experimentation → equip), with placeholder stats explicitly marked NON-CANONICAL.
- Verify-don't-claim methodology per the fork's `CLAUDE.md`: `go build` + `go vet` clean, fresh DB, all phase tests run (old and new), real output pasted, commit per verified step.

## Traceability (to be completed on approval)

| Item | Value |
|---|---|
| Owner approval | PENDING — date and exact approved text to be recorded |
| Ruling ID | To be assigned on filing in `governance/` (suggested: HD-TST-01) |
| Prior rulings affected | None — HD-EIC-01–08, HD-EIC-09, HD-GDD-01 unchanged |
| Checklist outcome | To be attached (reuse-patterns-only recommendation + evidence) |

*End of draft. To enact: owner replies "approve" (or amends), then this text — not a paraphrase — is filed in `governance/` under its ruling ID, and `governance/project_memory.md` gains one repo-event line.*
