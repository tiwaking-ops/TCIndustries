# SWG Pre-CU — Predecessor Project Material

**Status:** HISTORICAL — reference/inspirational material. Never automatically a TCIndustries rule.
**Provenance:** Two groups, recorded separately below. Group A: Perplexity thread (author: Perplexity — GPT-5.2). Group B: arena.ai thread (authors per artifact, as stated by the project owner).

## Purpose

This directory preserves the predecessor project that informed TCIndustries:

- `SWG_PreCU_GDD.md` — a fan-made, non-commercial, educational preservation design document for a Star Wars Galaxies (Pre-CU era) sandbox recreation, written as a build specification for an autonomous AI agent.
- `SWG Pre-CU Phase N — README.md` — milestone snapshots of a working build (Go 1.22 server + SQLite + WebSocket; Godot 4 client).
- `SWG Pre-CU Phase N — Server + Client.zip` — byte-for-byte provenance snapshots of the source trees at each phase boundary.
- `code-phase-N/` — the same source trees extracted for direct reading/searching. Evidence only; not design authority.
- `perplexity_computer-SWG Game Development Questions_1789200439254.md` — the original chat thread export (source of Group A material).
- `swg-crafting-system-architecture-llm-seed2-1-pro-preview.zip` (Group B) — Vite + React + TypeScript interactive tech demo of the SWG Pre-CU crafting core (assembly + experimentation engine, pure-TS core written for 1:1 port to UE5 C++). Author: seed-2.1-pro-preview LLM.
- `swg-crafting-system-architecture-claude-opus-5-max.zip` (Group B) — Vite + React + TypeScript architectural recommendation + live tech demo (galaxy resource generation, spawn browser, FSM craft bench, self-test). Author: claude-opus-5-max LLM.
- `craft-seed2-1-pro-preview/`, `craft-claude-opus-5-max/` — the same two trees extracted for direct reading/searching. Evidence only; not design authority.
- Group B provenance: both files from https://arena.ai/c/01a01b4e-5bf1-74b2-9f39-582ec2cefc23 (shared chat record).
- `swg-phase0-3-claudecode-ready.zip` (Group C) — later evolved server tree, phases 0–3 (adds Phase 3 Combat: HAM resolution, postures, wounds/fatigue, incapacitation/revive/clone, creatures + lairs, retaliation AI; plus `migrations/`, agent instructions `CLAUDE.md`). Built against embedded **GDD v2** (skill-point ladder / XP-formula corrections, patch notes near §6.3.1). `CLAUDE.md` records a verify-by-running methodology with real bugs found and fixed. Author: Author LLM Unknown (embedded GDD v2 carries `**Author:** Perplexity — GPT-5.2`).
- `swg_track_a_b_code.zip` (Group C) — early Track A/B prototyping scaffold: mostly class stubs plus small reference functions (combat resolution, resource-spawn generator, inflation sim, experimentation outcomes, XP curve), two one-line Accepted TDRs (server-authoritative architecture; ledger economy), phase backlog to Phase 10, a GDD copy (pre-attribution original, duplicate of Group A GDD), and a second copy of **GDD v2** (byte-identical to the claudecode bundle's copy — corroborated in two independent bundles). Author: Author LLM Unknown.
- `code-phase0-3-claudecode/`, `track-a-b-code/` — the same two trees extracted for direct reading/searching. Evidence only; not design authority. Runtime SQLite junk (`swg.db-shm`, `swg.db-wal`) excluded at extraction.
- Group C provenance: files received at repo root from the project owner; prior history unrecorded.
- `java_swg_files.zip` / `java-swg-skillcalc/` (Group C) — small Java Swing skill-calculator prototype implementing pre-CU rules (250-point pool, prerequisite chains, surrender guards, 3 professions × 18 boxes). Author: Author LLM Unknown.

## Open provenance question (unconfirmed)

The TCIndustries GDD (PROT-002) references a "Seed-2.1 TypeScript/Vite/React prototype" never inspected in-repo. The `craft-seed2-1-pro-preview/` tree is a TypeScript/Vite/React prototype by an authoring model named seed-2.1-pro-preview — a strong nominal match and therefore a **candidate** for the referenced prototype. Candidacy is **unconfirmed**: "Seed-2.1" in the GDD may be a version label rather than this model. No reconciliation claim is made here; confirmation awaits the project owner.

## Governance notes

- **No status is inherited.** Nothing here is LOCKED, DERIVED CONSTRAINT, PROPOSED, or canonical for TCIndustries from being stored in this repository.
- Using this code as a basis for a future TCIndustries implementation is a human-gated decision, not established by storage here.
- `SWG_PreCU_GDD.md` concerns Star Wars Galaxies and is written in the SWG emulation/preservation tradition. It is not TCIndustries design authority.
- Star Wars and Star Wars Galaxies are trademarks of Lucasfilm Ltd. TCIndustries is an original intellectual property and is not affiliated with Lucasfilm.

## Contents

| Artifact | Phase | Nature |
|---|---|---|
| `SWG_PreCU_GDD.md` | — | 34-section systems-design GDD for autonomous AI implementation |
| `SWG Pre-CU Phase 0 – README.md` | 0 | Server foundation, account/character creation, HAM |
| `SWG Pre-CU Phase 1 – README.md` | 1 | Movement, spatial interest management, chat |
| `SWG Pre-CU Phase 2 – README.md` | 2 | Skills, professions, XP, training persistence |
| `SWG Pre-CU Phase 0 – Server + Client.zip` | 0 | Source tree snapshot (37 entries) |
| `SWG Pre-CU Phase 1 – Server + Client.zip` | 1 | Source tree snapshot (43 entries) |
| `SWG Pre-CU Phase 2 – Server + Client.zip` | 2 | Source tree snapshot (49 entries) |
| `code-phase-0/` | 0 | Extracted source tree |
| `code-phase-1/` | 1 | Extracted source tree |
| `code-phase-2/` | 2 | Extracted source tree |
| `perplexity_computer-SWG Game Development Questions_1789200439254.md` | — | Original chat thread export (Group A) |
| `swg-crafting-system-architecture-llm-seed2-1-pro-preview.zip` | — | Crafting-core tech demo snapshot, seed-2.1-pro-preview LLM (Group B) |
| `swg-crafting-system-architecture-claude-opus-5-max.zip` | — | Architecture + live-loop demo snapshot, claude-opus-5-max LLM (Group B) |
| `craft-seed2-1-pro-preview/` | — | Extracted crafting-core tree (Group B) |
| `craft-claude-opus-5-max/` | — | Extracted architecture-demo tree (Group B) |
| `swg-phase0-3-claudecode-ready.zip` | — | Phases 0–3 server snapshot incl. Phase 3 Combat, GDD v2 (Group C, author unknown) |
| `swg_track_a_b_code.zip` | — | Track A/B scaffold + reference functions (Group C, author unknown) |
| `code-phase0-3-claudecode/` | — | Extracted phases 0–3 tree (Group C) |
| `track-a-b-code/` | — | Extracted scaffold tree (Group C) |