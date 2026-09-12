# SWG Pre-CU Code — Due-Diligence Checklist v0.2

**Status:** PROPOSED procedure. Not executed. Creates no design authority.
**Authority:** None. Human approval required before any reuse decision.
**Revision:** v0.2 extends v0.1 scope to all later arrivals. No step executed yet.
**Material under review:** `sources/swg-pre-cu/` — everything in it:
- Group A (Perplexity — GPT-5.2): Phase 0–2 zips + `code-phase-0/1/2/` + READMEs + GDD v1.
- Group B (arena thread): two Vite/React/TS crafting demos + zips (`craft-seed2-1-pro-preview/` — CONFIRMED PROT-002 prototype — and `craft-claude-opus-5-max/`).
- Group C (authors unknown): phases 0–3 evolved tree + zip (`code-phase0-3-claudecode/`, incl. Phase 3 Combat, GDD v2), Track A/B scaffold + zip (`track-a-b-code/`), Java skill-calc + zip (`java-swg-skillcalc/`).
Status HISTORICAL throughout (PROT-002 confirmation changes custody, not status: prototype evidence only).

## What this is (plain language)

Before the SWG code can serve as a basis for anything in TCIndustries, four questions must be answered with evidence: (1) does it actually work as claimed, (2) is it legally safe to reuse, (3) does the technology fit TCIndustries needs, and (4) how much of TCIndustries does it actually cover. This checklist defines how to answer each. Executing it is technical verification work, not a design decision — but the final reuse decision stays human.

## Steps

### 1. Integrity check (no build needed)
- [ ] Confirm each zip extracts cleanly and matches its extracted tree: `code-phase-0/1/2/` (37 / 43 / 49 entries), `code-phase0-3-claudecode/` (62 zip entries; runtime `swg.db-shm`/`swg.db-wal` excluded at extraction — confirm absent), `track-a-b-code/` (24 entries incl. GDD v2 copy), `java-swg-skillcalc/` (11 files), `craft-seed2-1-pro-preview/` (22 entries), `craft-claude-opus-5-max/` (24 entries).
- [ ] Record SHA-256 hashes of all eight zips in the results note.
- Evidence: hash list + extraction log.

### 2. Build verification (per tree, toolchain per tree)
- [ ] Go servers (`code-phase-0/1/2/`, `code-phase0-3-claudecode/`): `go build ./...` with the toolchain in each tree's `go.mod` (Go 1.22 per phase READMEs; phase0-3 `go.mod` differs — record actual).
- [ ] Godot clients: import each `project.godot` in Godot 4, confirm scenes open without errors.
- [ ] Java skill-calc: `javac` + launch `Main` per its README (Java 8+, no dependencies).
- [ ] Crafting demos (both): `npm install` + `vite build` (Node version recorded; React 19 / Vite 7 per `package.json`).
- [ ] Record exact tool versions, pass/fail, and any error output verbatim.
- Evidence: build log per tree. Do not fix failures yet — record them.

### 3. Test verification (per tree)
- [ ] Go phases: run `testclient`, `phase1test`, `phase2test`, `phase3test`; record claimed-vs-observed (Phase 0 exit criteria; 7 tests Phase 1; 11 tests Phase 2; "all 4 phases pass together" Phase 3 bundle claim).
- [ ] Track A/B: smoke-run the non-stub reference functions (`combat_resolution`, `resource_spawn_generator`, `economy_inflation_sim`, `crafting_experimentation` outcomes, `xp_progression_curve`); record stubs as stubs (`class X: pass` files are scaffolding, not implementation).
- [ ] Java skill-calc: launch smoke test (UI opens; 250-point pool behaves per README rules).
- [ ] Crafting demos: `vite preview` click-through (assembly → experimentation → finalize path); run the claude demo's built-in SelfTest section and record results.
- Evidence: test output logs. All authored claims (READMEs, CLAUDE.md "fully verified" statements) count as unverified until reproduced here.

### 4. IP / hygiene review (all trees)
- [ ] Confirm no Lucasfilm-owned assets (art, audio, text, names beyond generic mechanics) exist in code or scenes.
- [ ] Confirm no Star Wars lore, names, factions, or characters that would violate SWG-002 strict separation if reused. Known SWG-specific content already spotted (record more as found): Phase 3 creatures (Worrt, Womp Rat), planet reference (Tatooine), species/HAM content across Go trees, SWG schematic/crate vocabulary in demos. Any reuse must strip or replace these.
- [ ] Note each artifact's own classification (fan-made, non-commercial, educational) and any reuse implications.
- Evidence: short hygiene note with file references for any finding. If in doubt, flag — do not clear silently.

### 5. Technology fit assessment (all trees)
- [ ] Stack fit per tree: Go + SQLite + WebSocket servers; Godot 4 clients; Java 8 Swing skill-calc; Vite/React/TS demos (one explicitly written for UE5/C++ port). Compare against TCIndustries needs (persistent world, server-authoritative economy, future scale path).
- [ ] Identify what transfers as-is (e.g. auth/world-entry patterns, skill-box/prereq logic, craft-session FSM), what needs rework (e.g. SWG-specific systems, demo UI shells), and what is missing (EIC, TCIndustries economy, provenance).
- Evidence: fit table, one row per major subsystem per tree. No numeric tuning, no new mechanics — assessment only.

### 6. Coverage gap map (all trees)
- [ ] Map each tree's implemented systems against v1.1.1 §30 (Systems Requiring Deeper Design) and mark covered / partial / absent. Include GDD v2 as a documentation input (supersedes v1 for the phase0-3 tree only where its patch notes say so).
- Evidence: gap table. This shows how far the code goes toward a TCIndustries build — it does not promote any of it.

### 7. Reuse recommendation (for human decision)
- [ ] Present exactly one of: reuse as-is / reuse patterns only / reference only / reject — with evidence citations for the recommendation.
- [ ] The recommendation is PROPOSED. The decision is human. Record the decision and date when made.

## Out of scope for this checklist

- No TCIndustries design decisions. No EIC shaping (HD-EIC-07 gate unaffected). No numeric tuning. No provenance/reputation work. No GDD edits. PROT-002 evidence reconciliation (GDD §34 step 0) is a separate track: this checklist verifies the seed demo *runs as claimed*; reconciliation judges what it *proves* — different questions, different artifacts.
