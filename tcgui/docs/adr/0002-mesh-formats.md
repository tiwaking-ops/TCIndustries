# ADR 0002: Mesh Asset Format Policy (`docs/adr/0002-mesh-formats.md`)

**Status:** Accepted
**Date:** 2026-10-24
**Decision type:** Standalone (no preceding ADR)

## Context

Milestone M8 requires a documented position on mesh formats: import them, or
defer. SWG's native scene format is `.sat`/` .msh` inside `.tre` archives,
none of which exist in this repository (see `docs/importer/findings.md` — T-80
deliverable). The existing 3D plumbing in `tcgui-scene` (`Vertex3D`,
`RetroLighting`, `Camera3D`/`OrthoCamera3D`/`Transform3D`) is generic and
format-agnostic, so the renderer imposes no constraint here.

## Decision

1. **Defer mesh import.** No mesh-importer is implemented in M8. SWG's
   `.sat`/`.msh` are undocumented binary formats; the only references to them are
   scattered across community reverse-engineering wikis, none of which
   constitute an authoritative, license-clear source this project may encode
   (AGENTS.md rule 7: "Never invent facts about Star Wars Galaxies";
   rule 9: no SWG-owned code/assets).

2. **Keep the renderer mesh-agnostic.** `tcgui-scene` already accepts
   pre-transformed, vertex-lit `Vertex3D` slices — no `.sat`/`.msh` types are
   introduced, so the scene crate's public API (frozen in
   `docs/interfaces.md`) is unchanged.

3. **If stakeholder-supplied extracts are provided later**, M8-T-83 is lifted
   only after:
   - a format-authority check (T-08 / license review), and
   - a new ADR recording the exact format(s) found and the intended importer
     crate(s) — e.g. `gltf` (permissive MIT) for interchange `.glb` assets,
     or a hand-written reader gated behind the `swg-import` feature flag.

## Consequences

- `tcgui-scene` remains render-ready but is purely geometry-passthrough; it
  consumes vertices the application (or a future importer) supplies.
- The `swg-import` Cargo feature flag (T-81) is reserved and currently exposes
  only the texture-import path. Adding a mesh importer later is one
  `#[cfg(feature = "swg-import")]` block plus its dependency declaration — no
  API rewrite needed.
- No new dependencies are added by this decision record; the ADR allowlist in
  `0001-stack.md` is unchanged.
