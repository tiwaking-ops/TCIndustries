# ADR 0001: Technology Stack and Pinned Versions (`docs/adr/0001-stack.md`)

**Status:** Accepted
**Date:** 2026-10-24

## Context
`tcgui` requires a robust, performant, GPU-accelerated UI framework and rendering backend in Rust that can execute headless tests without a display server and render faithfully on Windows and Linux.

## Decision
1. **Language & Edition:** Rust (Edition 2021), stable toolchain. **MSRV (Minimum Supported Rust Version):** 1.99.0. The workspace `rust-version` field enforces this. Newer toolchains are accepted; older toolchains are not supported.
2. **Graphics & Windowing:**
   - **WGPU** pinned to `24.0.1` — cross-platform, safe, pure-Rust graphics API abstracting Vulkan, DirectX 12, Metal, and OpenGL/WebGPU. Version 24 is used for stable `InstanceDescriptor` and surface API; API differences from wgpu 23 and earlier are noted below.
   - **Winit** pinned to `0.30.8` — native window management and event handling. Version 0.30 (not 0.31-beta) is stable and compatible with wgpu 24's `Surface` API. Note: winit 0.31 introduces breaking changes to the event loop API; the workspace remains on 0.30 until a coordinated upgrade is evaluated.
3. **Core Dependencies (all pinned via `workspace.dependencies`):**
   | Crate | Pinned Version | Purpose |
   |---|---|---|
   | `serde` | `1.0.217` | Declarative JSON themes, palettes, UI layouts |
   | `serde_json` | `1.0.135` | JSON serialization/deserialization |
   | `thiserror` | `2.0.11` | Structured error handling |
   | `wgpu` | `24.0.1` | GPU rendering abstraction |
   | `bytemuck` | `1.21.0` | Zero-copy vertex data casting in WGPU |
   | `ab_glyph` | `0.2.28` | Font rasterization and layout |
   | `winit` | `0.30.8` | Window management and event loop |
   | `pollster` | `0.4.0` | Synchronous future blocking for async wgpu calls |
   | `nalgebra` | `0.33.2` | 3D math transforms (`tcgui-scene`) |
4. **WGPU/Winit API Notes:**
   - `wgpu` 24 exposes `InstanceDescriptor` with `backend` set via `BackendFlags`. Surface creation uses `SurfaceTarget::Surface` (winit 0.30 integration).
   - `winit` 0.30's `EventLoop` API is stable; `Event::WindowEvent` and `ControlFlow` patterns are in use.
   - If upgrading to `wgpu` 25+ or `winit` 0.31+, an interface-change review is required (see AGENTS.md rule 3).
5. **Font Licensing:** Open-source fonts **DejaVu Sans** and **Liberation Sans** are referenced for UI text rendering. Both are licensed under the **SIL Open Font License (OFL) Version 1.1** and **Apache License Version 2.0** (Liberation Sans). No proprietary fonts are committed to the repository. Font asset files are downloaded at build time or referenced from system paths; see `docs/assets/fonts.md`.

   **SIL Open Font License (OFL) Summary:**
   - Fonts may be freely used, studied, copied, merged, embedded, modified, and redistributed.
   - Fonts must not be sold standalone.
   - Embedded components must remain under OFL.
   - Attribution is appreciated but not required.

   **Apache 2.0 (Liberation Sans) Summary:**
   - Free for commercial and non-commercial use.
   - Modification and distribution permitted.
   - Must include original copyright and license notice.

6. **Allowed Crate Allowlist:**
   Only the crates listed in this ADR and the workspace `Cargo.toml` files are permitted. Adding new external dependencies requires a formal ADR update. The full allowlist:
   - `serde`, `serde_json`, `thiserror` — core utilities
   - `wgpu`, `bytemuck`, `ab_glyph` — rendering and text
   - `winit`, `pollster` — platform/windowing
   - `nalgebra` — 3D scene math
   - `image` (`0.25.5`) — texture import for the optional `swg-import` feature (see ADR 0002, task T-81). Optional dependency; not compiled unless `--features swg-import` is enabled.

## Amendment 1 — `image` for T-81 texture import

**Date:** 2026-10-24
**Reason:** Task T-81 (M8) specifies DDS/TGA texture import via the `image` crate.

- **Double-licensed MIT / Apache-2.0** (permissive; MPL alternative may apply to certain optional sub-dependencies — not enabled by the `dds`/`tga` features used here).
- **Optional dependency.** `image` is only pulled in when the `swg-import` feature is enabled; the default build graph is unchanged.
- **Feature selection:** Only the `dds` and `tga` features are enabled. This avoids pulling in codecs the project does not need (JPEG/PNG/WebP/etc.), keeping the dependency tree minimal.
- **No unsafe directly in our code:** all decode operations go through `image`'s safe public API; no swg-owned data is processed (see ADR 0002, T-80 findings: no extracted assets exist).
- **Alternatives considered:** writing hand-rolled DDS/TGA parsers would require ~400+ lines of unsafe-free but unchecked bit manipulation and no spec authority; using `ddsfile`+`targa` (two deps instead of one) was rejected in favor of the consolidated `image` crate already widely used in the Rust ecosystem.

**Impact:** none on default (`check.sh`) builds; `image` is compiled only under `--features swg-import`.

## Consequences
- Pinning versions ensures reproducible builds across CI, development, and deployment environments.
- The `resolver = "2"` setting in the workspace `Cargo.toml` ensures correct feature unification across the dependency graph.
- The `rust-version = "1.99"` field prevents compilation on older toolchains that may lack required language features.
- Version pins are declared in `[workspace.dependencies]` and inherited via `.workspace = true` in member crates, simplifying updates.
- API differences between wgpu/winit versions are documented to guide future upgrades.
