# T-80 — Extracted Asset Findings (`docs/importer/findings.md`)

**Status:** BLOCKED (no extracted files provided). This document records what was
actually found. No assumptions about file formats were carried over from T-08.

**Date:** 2026-10-24
**Task:** T-80 — establish whether extracted Pre-CU SWG files exist before
building any importer.

## Summary

**No extracted SWG asset files exist in this repository or anywhere reachable
from it.** Every importer work item in M8 that depends on real SWG data is
therefore blocked, and nothing downstream can be built against observed files.

This is recorded rather than worked around, per `AGENTS.md` rule 8
("If blocked: stop and report") and rule 9 ("Never invent facts about
Star Wars Galaxies").

## What was searched

| Location | Result |
|---|---|
| `sources/swg-pre-cu/` (197 files) | Documentation and source code only — Markdown, Go, GDScript, `.tscn`/`.project.godot`. Zero asset files. |
| All 8 `.zip` archives in `sources/swg-pre-cu/` | 271 files total across all archives; **zero** files with a game-asset extension. |
| Whole repository (excluding `target/`) | No files matching `.dds`, `.tga`, `.apt`, `.ifs`, `.stf`, `.png`, `.jpg` under any asset path. |
| Gitignored / ignored paths | `.freebuff/`, `swg.db*`, `.choco/`, `.tmp_gcc/`, `testbed/**` — none contain extracted game assets. |

The search covered these extensions, which are the formats M8 was scoped to
import: `.tre` (archive), `.dds` and `.tga` (textures), `.apt` (declarative
UI), `.ifs`, `.stf` (string tables), plus generic image formats.

### Archive contents breakdown

| Archive | Files | Asset-like |
|---|---|---|
| `SWG Pre-CU Phase 0 — Server + Client.zip` | 37 | 0 |
| `SWG Pre-CU Phase 1 — Server + Client.zip` | 43 | 0 |
| `SWG Pre-CU Phase 2 — Server + Client.zip` | 49 | 0 |
| `java_swg_files.zip` | 11 | 0 |
| `swg-crafting-system-architecture-claude-opus-5-max.zip` | 24 | 0 |
| `swg-crafting-system-architecture-llm-seed2-1-pro-preview.zip` | 22 | 0 |
| `swg-phase0-3-claudecode-ready.zip` | 62 | 0 |
| `swg_track_a_b_code.zip` | 23 | 0 |
| **Total** | **271** | **0** |

`java_swg_files.zip` despite its name contains Java sources, not `.tre` archive
extracts.

## What is present at `sources/swg-pre-cu/`

Reference material and reimplementation attempts, not extracted originals:

- `SWG_PreCU_GDD.md` (221 KB) and five phase READMEs — design documentation.
- `code-phase-0/`, `code-phase-1/`, `code-phase-2/` — a Go server plus Godot
  client per phase.
- `code-phase0-3-claudecode/`, `craft-claude-opus-5-max/`,
  `craft-seed2-1-pro-preview/`, `java-swg-skillcalc/` — additional
  reimplementations.

None of these are the original client data files, and none are evidence of the
on-disk formats (see rule 7: no SWG-owned assets in this repository).

## Impact on M8

| Task | Scoped acceptance | Status |
|---|---|---|
| T-81 | Texture import via the `image` crate, tested on a **locally generated** DDS/TGA fixture | **Deliverable.** Acceptance explicitly permits a locally generated fixture, not SWG data. |
| T-82 | Convert **at least one real** SWG declarative UI file | **BLOCKED.** No real file exists. A converter against invented fixtures would be fiction dressed as evidence. Deferred. |
| T-83 | Mesh formats: decision record only | **Deliverable.** An ADR is written. |

T-82 is the task the importer milestone exists to serve, so M8 delivers partial
value: the generic texture-import half (T-81) and the decision record (T-83).

## What is needed to unblock T-82

1. A extracted-file location supplied by the stakeholder — the card's T-80
   instruction is "stop and ask the stakeholder for the extracted-file
   location," and that request has not been answered with a path.
2. Specifically: at least one `.apt` (declarative UI) file, and ideally the
   `.tre` archive it came from so sibling dependencies can be resolved.
3. Legal clearance recorded in governance before any extracted original is
   placed here. `AGENTS.md` rule 7 forbids committing SWG-owned assets, so
   extracted files must live in the gitignored `extracted-swg/` directory and
   must not enter this repository.

Until (1) and (2) are satisfied, T-82 stays deferred rather than being
approximated. Writing a converter against guessed schemas would produce code
that appears to work and silently mis-converts real files later — worse than
absence.
