# SWG Pre-CU Code — Due-Diligence Checklist v0.1

**Status:** PROPOSED procedure. Not executed. Creates no design authority.
**Authority:** None. Human approval required before any reuse decision.
**Material under review:** `sources/swg-pre-cu/` — Phase 0–2 zips, extracted `code-phase-0/1/2/` trees, phase READMEs. Author: Perplexity — GPT-5.2. Status HISTORICAL.

## What this is (plain language)

Before the SWG code can serve as a basis for anything in TCIndustries, four questions must be answered with evidence: (1) does it actually work as claimed, (2) is it legally safe to reuse, (3) does the technology fit TCIndustries needs, and (4) how much of TCIndustries does it actually cover. This checklist defines how to answer each. Executing it is technical verification work, not a design decision — but the final reuse decision stays human.

## Steps

### 1. Integrity check (no build needed)
- [ ] Confirm each zip extracts cleanly and matches the extracted `code-phase-N/` tree (file count: 37 / 43 / 49 entries).
- [ ] Record SHA-256 hashes of the three zips in the results note.
- Evidence: hash list + extraction log.

### 2. Build verification (per phase 0, 1, 2)
- [ ] Go server: `go build ./...` with the Go toolchain noted in `go.mod` (Go 1.22 per READMEs).
- [ ] Godot client: import `project.godot` in Godot 4 and confirm scenes open without errors.
- [ ] Record exact tool versions, pass/fail, and any error output verbatim.
- Evidence: build log per phase. Do not fix failures yet — record them.

### 3. Test verification (per phase 0, 1, 2)
- [ ] Run the phase integration tests (`testclient`, `phase1test`, `phase2test`) and record claimed-vs-observed results (READMEs claim Phase 0 exit criteria, 7 passing tests Phase 1, 11 passing tests Phase 2).
- Evidence: test output logs. Claims in Perplexity-authored READMEs count as unverified until reproduced here.

### 4. IP / hygiene review
- [ ] Confirm no Lucasfilm-owned assets (art, audio, text, names beyond generic mechanics) exist in code or scenes.
- [ ] Confirm no Star Wars lore, names, factions, or characters that would violate SWG-002 strict separation if reused.
- [ ] Note the prototype's own classification (fan-made, non-commercial, educational) and any reuse implications.
- Evidence: short hygiene note with file references for any finding. If in doubt, flag — do not clear silently.

### 5. Technology fit assessment
- [ ] Stack fit: Go 1.22 + SQLite + WebSocket server; Godot 4 client. Compare against TCIndustries needs (persistent world, server-authoritative economy, future scale path).
- [ ] Identify what transfers as-is (e.g. auth/world-entry patterns), what needs rework (e.g. SWG-specific systems), and what is missing (EIC, TCIndustries economy, provenance).
- Evidence: fit table, one row per major subsystem. No numeric tuning, no new mechanics — assessment only.

### 6. Coverage gap map
- [ ] Map each phase's implemented systems against v1.1.1 §30 (Systems Requiring Deeper Design) and mark covered / partial / absent.
- Evidence: gap table. This shows how far the code goes toward a TCIndustries build — it does not promote any of it.

### 7. Reuse recommendation (for human decision)
- [ ] Present exactly one of: reuse as-is / reuse patterns only / reference only / reject — with evidence citations for the recommendation.
- [ ] The recommendation is PROPOSED. The decision is human. Record the decision and date when made.

## Out of scope for this checklist

- No TCIndustries design decisions. No EIC shaping (HD-EIC-07 gate unaffected). No numeric tuning. No provenance/reputation work. No GDD edits.
