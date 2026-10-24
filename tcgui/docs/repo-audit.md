# TCIndustries GUI Repository Audit (`docs/repo-audit.md`)

**Date:** 2026-10-24  
**Project:** `tcgui` (TCIndustries GUI System)

## 1. Verified T-08 Repository Paths

Per Task T-00 requirements, the repository paths cited in T-08 and the project proposal have been audited:

- `sources/swg-pre-cu/` — **EXISTS**. Contains predecessor Godot projects (`swg-godot`), Go servers (`swg-server`), and documentation/skill calculators.
- `testbed/swg-phase3-combat/` — **EXISTS**. Contains Phase 3 combat fork implementation (`swg-server` and `swg-godot`).
- `testbed/swg-phase3-combat-fork/` — **EXISTS**. Contains extended testbed code.

## 2. Workspace Layout

The `tcgui/` workspace has been established with the following modular crate structure:
- `tcgui-core`: GPU-free retained-mode UI toolkit (widget tree, layout, input routing, theming, data binding).
- `tcgui-render`: WGPU 2D batch renderer, bitmap-font text, 9-slice panels.
- `tcgui-platform`: winit window/input layer for Windows and Linux.
- `tcgui-scene`: Retro profile 3D renderer for login/character backdrops.
- `tcgui-demo`: Runnable demo with Login, Character Select, HUD, Inventory, and Chat windows.
