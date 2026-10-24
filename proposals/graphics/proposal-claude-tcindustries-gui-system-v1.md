# TCIndustries GUI System — Formal Project Proposal

| Field | Value |
|---|---|
| **Author** | Claude Sonnet 5.5 |
| **Version** | 1.0 |
| **Date** | 2026-10-23 |
| **Status** | PROPOSED |
| **Working name** | `tcgui` (TCIndustries GUI System) |
| **Based on** | `T-08_SWG_Graphics_Engine_Report.md` (v1.0, author Claude Fable 5) |
| **Audience** | 1) TCIndustries stakeholder (decision-maker); 2) implementing LLM agents (Claude Code, worker swarm) |
| **File** | `proposal-claude-tcindustries-gui-system-v1.md` |

---

## 0. How to Read This Document

This proposal has two halves.

- **Part A (Sections 1–8): the decision half.** What is being built, why, how it is shaped, and what "great" looks like. Written for the human stakeholder.
- **Part B (Sections 9–13 and appendices): the execution half.** Milestones, task cards, rules, and verification written for LLM agents. An agent should be able to take a single task card, read only the sections it points to, and complete it without asking questions.

**Evidence legend.** Every non-trivial claim about Star Wars Galaxies (SWG) or about T-08 carries a tag:

| Tag | Meaning |
|---|---|
| **[V]** | Verified against a source found during research for this proposal (listed in Appendix A). |
| **[T08]** | Stated in T-08 and not independently verified. Treat as a hypothesis. |
| **[P]** | Provisional: my best guess or design choice. Must be checked against references (Task T-02) before it is treated as fact. |
| **[D]** | A design decision made in this proposal (not a claim about SWG). |

**Definitions used throughout.** *Pre-CU* means SWG before the Combat Upgrade (April 2005 or earlier), matching the definition already used for your SWG Pre-CU game. *NGE* means the later New Game Enhancements era, whose UI is explicitly **not** the target. *HAM* means Health, Action, Mind, the three primary attribute pools.

---

# PART A — DECISIONS

## 1. Executive Summary

TCIndustries wants a reusable GUI system whose first skin reproduces the look and feel of the Pre-CU Star Wars Galaxies interface. T-08 proposed a full graphics engine and UI framework in Rust with WGPU. This proposal turns that report into an implementable plan that LLM agents can execute in small, verifiable pieces.

**What will be built** (full T-08 scope, as chosen in the interview):

1. **`tcgui-core`**: a GPU-free retained-mode UI toolkit (widget tree, layout, input routing, theming, data binding). Fully testable without a window.
2. **`tcgui-render` + `tcgui-platform`**: a WGPU 2D batch renderer, bitmap-font text, 9-slice panels, and a winit window/input layer for Windows and Linux.
3. **`tcgui-scene`**: a small "retro profile" 3D renderer (vertex-lit, optionally flat-shaded) used for login/character backdrops and for the 100-dynamic-object performance scenario.
4. **`tcgui-theme-swg`**: the Pre-CU skin, expressed as data (JSON), with textures generated procedurally so no copyrighted art is needed.
5. **`tcgui-assets`**: texture/font loading with reference counting, plus an **optional** importer for assets extracted from your own SWG CD (feature-gated, off by default).
6. **`tcgui-backend-iface`**: a trait plus event/command types that define how a real game server connects later, and a deterministic mock backend so the GUI runs standalone now.
7. **`tcgui-demo`**: a runnable demo with Login, Character Select, HUD, Inventory, and Chat windows.

**Key shape decisions**

- **Data-driven UI [D]:** Layouts and themes are JSON files validated by a schema. Research showed the original SWG UI was itself data-driven (declarative page/text/button files with palette tokens and opacity attributes) and that the community reskinned it by replacing those files **[V]**. Copying that *architecture* makes theming, LLM authoring, and later reskinning cheap.
- **Contract-first, headless-first [D]:** The widget model, layout, and display list are produced without a GPU, so agents can verify most work with `cargo test` and no display.
- **Leak-proof by construction [D]:** Arena allocation with generational IDs, no reference cycles, explicit resource counters, and a 30-minute soak mode that asserts memory plateaus.
- **Small tasks for mid-capability agents [D]:** About 50 task cards, each sized for one agent session, each with a command-line acceptance check.

**Corrections to T-08.** Research contradicted several T-08 claims (Section 3). Most importantly, the original UI used **opacity/transparency** and the original client used **vertex/pixel shaders and bump mapping** where hardware allowed **[V]**. T-08's "no transparency, no shaders" rules are therefore treated as an optional *stylistic* profile, not as historical fact.

**Recommendation:** Approve this proposal and start with Milestone M0 (repository audit and reference acquisition), then M1 (headless core). A first playable demo (Login → HUD with mock data) is reachable at the end of Milestone M5.

---

## 2. Decisions and Assumptions (Interview Record)

### 2.1 Answers you gave

| # | Topic | Decision |
|---|---|---|
| 1 | Scope | **Full T-08 scope**: renderer, asset pipeline, UI framework, performance and leak tests. |
| 2 | Reuse | **Reusable** across TCIndustries projects; the first skin is SWG Pre-CU. |
| 3 | Screens | Framework plus a demo set: **Login, HUD, Inventory, Chat**. (Character Select is added as a thin bridge screen between Login and HUD; see 2.2.) |
| 4 | Stack | **Follow T-08: Rust + WGPU.** |
| 5 | Implementers | **Claude Code and a worker swarm.** Assume mid-capability agents; small, verifiable tasks. |
| 6 | Server link | **Standalone now**, with a defined interface for connecting later. |
| 7 | T-08 authority | Follow T-08 for architecture; mark shaky claims as "verify against references." |
| 8 | Screenshots | **The agent must search online** for Pre-CU screenshots; plenty still exist. |
| 9 | Assets | No SWG assets in the repo. You hold the CD and can extract assets, but this is a test program: **original recreated assets, with an optional importer.** |
| 10 | Platform | Scalable from 1024×768 up, mouse and keyboard, Windows and Linux. |
| 11 | "Great" means | (b) **Modular and easy for LLMs to extend** and (d) **fastest route to something playable.** |
| 12 | Schedule | **LLM-session milestones**, no calendar dates. |
| 13 | Length | Longer is better. |
| 14 | File name | `proposal-claude-tcindustries-gui-system-v1.md`, Version 1.0. |

### 2.2 Assumptions I made (flagged)

| ID | Assumption | Why | Cost if wrong |
|---|---|---|---|
| A1 | "GUI system" includes the minimal 3D renderer from T-08 because you chose the full scope. | Answer 1. | Low: `tcgui-scene` is a separate crate and can be dropped. |
| A2 | Character Select is included as a simple list screen. | Login → HUD needs a step in between to be coherent. | Low. |
| A3 | Default Pre-CU look is "approximate and provisional" until T-02 measures real references. | I could not inspect screenshot pixels during research. | Medium: palette and geometry would change. |
| A4 | Rust crate versions are **not** pinned here. | Library APIs change quickly (especially `wgpu`); task T-03 pins current versions. | Low if T-03 is done first. |
| A5 | The optional SWG importer only reads files you extract yourself and never redistributes them. | Legal caution; your "test program" framing. | None. |
| A6 | Performance numbers (60 FPS, 100 dynamic objects, <1 GB) are inherited from T-08 **[T08]**. | They are plausible; I have no better baseline. | Low. |
| A7 | The "worker swarm" can run several tasks in parallel in separate git worktrees. | Your earlier descriptions of a multi-worker setup. | Medium: if not, run tasks sequentially in dependency order. |

---

## 3. Research Findings and Corrections to T-08

### 3.1 What the research established

| Finding | Tag | Consequence for the design |
|---|---|---|
| The original UI was declarative: files with `Page`, `Text`, `Button` elements, attributes such as `Location`, `Size`, `PackLocation`, `PackSize`, `BackgroundOpacity`, `Opacity`, palette references (`PalBgTint`, `PalText`), named style resources, font ids such as `default_13`, and `Data` elements binding named code hooks. | **[V]** | Adopt data-driven layouts, palette tokens, and named bindings. |
| Palettes and styles lived in separate files (`ui_palette_ground.inc`, `ui_styles.inc`); modders swapped whole HUD pieces (button bar, chat window, status display, radar, targets, toolbar) by replacing individual files. | **[V]** | One JSON file per HUD component; themes separate from layouts. |
| Windows had **opacity** controls and fading behavior (for example chat window fade). | **[V]** | T-08's "no transparency" rule is wrong. Alpha blending and per-window opacity are required. |
| The manual describes the client using **vertex and pixel shaders** when hardware supports them, with options to disable them and to disable bump mapping; windowed and borderless modes exist; resolution is selectable. | **[V]** | "Flat-shaded, no shaders" is a stylistic choice, not history. Keep it as `RetroProfile`, configurable. |
| Minimum hardware was a 32 MB T&L-capable GPU with DirectX 9.0c. | **[V]** | Good justification for keeping GPU requirements low. |
| HUD elements: HAM bars; wound indicator; posture; states display; toolbar bank; radar; chat window with channel tabs; chat input; row of data-screen icons. | **[V]** | Defines the HUD widget list (Section 6.3). |
| Toolbar: six toolbars numbered 0–5; can be expanded to two rows; second row reached with Shift + function key; slots on F1–F12. | **[V]** | Toolbar widget spec (6.3). |
| Radar: north indicator; orange arrows for mission waypoints; blue arrows for generic waypoints; brackets around current target; red dots for enemies; green dots for group members; displayed range. | **[V]** | Radar colors and symbols (6.3). |
| Data screens: Map (Ctrl+V), Inventory (Ctrl+I), Datapad (Ctrl+D), Character Sheet (Ctrl+C), Community (Ctrl+P), Abilities (Ctrl+A); **Data mode** entered with Alt (mouse free) versus **World mode**. | **[V]** | Input-mode and keybind design (5.4). |
| Combat text colors: red = Health, green = Action, blue = Mind; targets show a color-coded difficulty starburst beside their HAM bars. | **[V]** | HAM color scheme and target frame. |
| Inventory behaviors: items reduced to 0 hitpoints show red; the inventory icon blinks when you are over capacity by one item; at two over, a persistent red "Inventory Overloaded" message appears; an options-menu setting shows crafting values on resources. | **[V]** | Inventory window behaviors (6.4). |
| A recurring forum complaint: inventory sort order was not remembered between sessions. | **[V]** | Persist sort order in our inventory (cheap improvement). |
| The Pre-CU UI used **white HAM bar backgrounds**, preserved in community mods; the NGE-era UI is smaller, with a non-resizable character window. | **[V]** | Default HAM bar background is white; do not copy NGE layout. |
| Screenshots exist at: MobyGames (Pre-CU Wookiee village, Mos Eisley cantina, circa 2004), MMORPG.com screenshot galleries (e.g., December 2004), archived official site pages (2004–2007 mirrored by SWGEmu), forum threads comparing Pre-CU and NGE UIs, and community mod repositories. | **[V]** | Source list for T-01 (Appendix A). |

### 3.2 T-08 claims audited

| T-08 claim | Status | Action |
|---|---|---|
| "No transparency" in the UI | **Contradicted** | Support alpha; per-window opacity setting. |
| "No shaders by default"; flat shading and no dynamic lighting as SWG's look | **Contradicted as history** | Keep as optional `RetroProfile { flat_shading, vertex_lit_only, ... }`, default on for aesthetic reasons only. |
| Palette "recovered from SWG client screenshots" | **Unsupported** (no evidence shown) | Treat hex values as **[P]**. T-02 samples real references. |
| Font named "SwgSystemFont" | **Unsupported** | Do not use that name; use an open-license font rasterized to a bitmap atlas. |
| Formats `.tlb`, `.ltg`, `.lod`/`.cmp` for UI, lighting, meshes | **Unverified** | Importer tasks are gated on inspecting real extracted files (T-80 onward). |
| "No anti-aliasing, no post-processing" | **Stylistic** | Configurable; default off. |
| Performance targets (60 FPS, 100 dynamic objects, <1 GB, zero leaks over 30 min) | **Reasonable, self-imposed** | Keep; make measurable (Section 7). |
| Shader listings | **Not compilable** (mixes GLSL, HLSL, and WGSL syntax) | Replaced by real WGSL in Section 5.7. |
| Repository paths `sources/swg-pre-cu/` and `testbed/swg-phase3-combat/` | **Cannot be verified from here** | T-00 must verify they exist before relying on them. |
| 9-week schedule | **Not applicable** | Replaced by LLM-session milestones. |

### 3.3 Prior art noticed (not endorsed)

- Community UI mod repositories (palette tools, HUD reskins) are useful for understanding structure. Their licenses were not checked. **Do not copy their code or art.** Use them only as references for layout concepts.
- A public GitHub repository describing a Rust "Pre-CU authentic HUD" (radar, HAM bars, target frame, button bar) appeared in search results. Its provenance, quality, and license are unknown. Treat it as prior art to skim for ideas at most, after the agent checks its license.

---

## 4. Product Definition

### 4.1 Goals

| ID | Goal |
|---|---|
| G1 | A reusable Rust GUI toolkit with a clean, documented, LLM-friendly API. |
| G2 | A Pre-CU SWG skin and demo screens that read as "SWG" at a glance, validated against references. |
| G3 | Runs at 60 FPS with the T-08 scenario on mid-range PC hardware; no memory growth over 30 minutes. |
| G4 | Standalone operation via a mock backend; a documented interface for a real server. |
| G5 | Every task independently verifiable by command-line checks. |

### 4.2 Non-goals

- Gameplay logic, networking implementation, audio, character creation flow, or the NGE/post-CU UI.
- Mobile or console platforms.
- A general-purpose 3D game engine (T-08 flagged this scope creep as a high risk; `tcgui-scene` stays minimal).
- Using original SWG assets in any distributed build.

### 4.3 Users

| User | Needs |
|---|---|
| TCIndustries stakeholder | Visible progress, honest status, control over scope. |
| LLM implementer (Claude Code or swarm worker) | Unambiguous task cards, stable interfaces, fast local checks. |
| Future game developer | Simple widget API, JSON theming, a backend trait to implement. |

### 4.4 What separates great from mediocre

| Dimension | Mediocre | Great |
|---|---|---|
| Modularity (your priority b) | Hard-coded screens; theme baked into widgets. | New screen = one JSON file plus bindings; new skin = one theme folder; verified by a "second skin" test (T-93). |
| Speed to playable (your priority d) | Waits for the 3D engine. | Login → HUD with mock data runs at end of M5, before any 3D or importer work. |
| Fidelity | Generic "sci-fi metal" UI. | Side-by-side with references shows matching window chrome, HAM bars, radar, toolbar, and chat. |
| Reliability | Works on the author's machine. | Soak test and headless snapshot tests pass on a clean checkout. |

---

## 5. Architecture

### 5.1 Workspace layout [D]

```
tcgui/
├── Cargo.toml                    # workspace
├── AGENTS.md                     # rules for LLM agents (Appendix E)
├── scripts/
│   └── check.sh                  # fmt + clippy -D warnings + test (+ schema check)
├── crates/
│   ├── tcgui-core/               # widgets, layout, events, theme, binding, display list (NO gpu deps)
│   ├── tcgui-backend-iface/      # GameBackend trait, events, commands, MockBackend
│   ├── tcgui-platform/           # winit window, input mapping, clipboard, DPI
│   ├── tcgui-render/             # wgpu 2D batch renderer, atlas, text, offscreen capture
│   ├── tcgui-scene/              # retro 3D profile, camera, meshes, culling
│   ├── tcgui-assets/             # texture/font loading, refcounts; optional `swg-import` feature
│   ├── tcgui-theme-swg/          # theme + layouts as data; build-time procedural textures
│   └── tcgui-demo/               # binary: login, charselect, HUD, inventory, chat; --screenshot; --soak
├── layouts/                      # *.ui.json (screens and HUD pieces)
├── themes/swg-precu/             # theme.json, palette.json, generated textures
├── schemas/                      # JSON Schemas generated from Rust types
├── tools/
│   ├── asset-gen/                # procedural bevel/gradient/icon generator
│   └── refcompare/               # composes reference and output images side by side
├── docs/
│   ├── interfaces.md             # frozen contracts after M1
│   ├── reference-catalog.md      # produced by T-01
│   ├── visual-spec.md            # produced by T-02
│   ├── adr/                      # architecture decision records
│   └── tasks/                    # task cards + register (Section 10)
└── tests/golden/                 # display-list snapshots and optional image goldens
```

**Dependency rule:** `tcgui-core` depends on no GPU, window, or file-format crates. `tcgui-render` depends on `core`. `tcgui-demo` depends on everything. No cycles.

**Allowed third-party crates (starting allowlist; any addition needs an ADR):** `wgpu`, `winit`, `pollster`, `bytemuck`, `glam`, `slotmap`, `serde`, `serde_json`, `schemars`, `image`, `fontdue` (or `ab_glyph`), `thiserror`, `log`, `env_logger`, `criterion` (dev), `insta` (dev). T-03 pins exact versions.

### 5.2 Core object model [D]

- Widgets live in a `SlotMap<WidgetId, WidgetNode>` arena. `WidgetId` is generational, so a stale ID can never alias a new widget. There is no `Rc`/`RefCell` tree, which removes the main source of leaks and cycles.
- Destroying a widget destroys its subtree and removes its bindings, event handlers, and texture references in the same call. Tests assert the arena length returns to its baseline after create/destroy cycles.
- Widget kinds are an enum plus a trait for custom drawing, so agents can add a widget by implementing one trait and registering one name.

```rust
// tcgui-core/src/lib.rs (skeleton: names are normative, bodies are illustrative)
slotmap::new_key_type! { pub struct WidgetId; }

#[derive(Copy, Clone, Debug, PartialEq, Default)]
pub struct Rect { pub x: f32, pub y: f32, pub w: f32, pub h: f32 }

pub trait Widget {
    fn kind(&self) -> &'static str;                         // matches "type" in JSON
    fn measure(&self, ctx: &MeasureCtx, avail: Vec2) -> Vec2;
    fn event(&mut self, ctx: &mut EventCtx, ev: &UiEvent) -> EventResult;
    fn draw(&self, ctx: &DrawCtx, out: &mut DrawList);      // emits quads/text/clips only
}

pub struct Ui {
    arena: slotmap::SlotMap<WidgetId, WidgetNode>,
    roots: Vec<WidgetId>,
    theme: Theme,
    model: UiModel,
    // ...
}
impl Ui {
    pub fn load_layout(&mut self, json: &str) -> Result<WidgetId, LayoutError>;
    pub fn destroy(&mut self, id: WidgetId);                 // subtree + bindings + handlers
    pub fn handle_input(&mut self, ev: InputEvent);          // routes, captures, focuses
    pub fn layout(&mut self, viewport: Vec2);                // deterministic, idempotent
    pub fn draw(&self) -> DrawList;                          // pure: same state -> same list
    pub fn stats(&self) -> UiStats;                          // widget count, bindings, handlers
}
```

**Display list [D]:** `DrawList` is a flat `Vec<DrawCmd>` of `Quad { rect, uv_rect, color, texture }`, `PushClip(rect)`, `PopClip`. Because it is plain data, tests can snapshot it as JSON with no GPU. The renderer's only job is turning a `DrawList` into GPU batches.

### 5.3 Layout system [D]

Absolute and anchored positioning, plus three flow containers. Chosen because it is simple to specify, test, and generate by LLM. (Original SWG used `Location`/`Size`/`PackLocation`/`PackSize` **[V]**; their exact semantics were not verified, so this system does not claim to mimic them.)

**Axis rule.** For each axis, set exactly two of `start`, `end`, `size`; otherwise `LayoutError::OverConstrained` or `UnderConstrained`.

| Spec (per axis) | Meaning |
|---|---|
| `start` + `size` | Pinned to parent's start edge, fixed size. |
| `end` + `size` | Pinned to parent's end edge, fixed size. |
| `start` + `end` | Stretches with the parent. |

Values are pixels in the **base UI space** (1024×768) or a percentage string (`"50%"`) of the parent. A global **UI scale** factor (default 1.0, options 0.75–2.0) and **viewport scaling** keep windows sized sensibly at higher resolutions. Min and max sizes are honored.

Flow containers: `row`, `column`, and `grid` (fixed columns, uniform cell size), with `gap` and `padding`. The inventory grid and toolbar use `grid`.

### 5.4 Input, focus, and modes [D, with SWG behaviors **[V]**]

- **Two modes**, toggled by Alt: **World mode** (cursor controls the world; most windows ignore the mouse) and **Data mode** (cursor free; windows interactive) **[V]**. The toolkit exposes `InputMode` and tells the application which events the UI consumed so game input does not double-handle clicks.
- Hit testing walks windows in z-order; only the topmost containing window receives pointer events. Pointer capture on press for drags and sliders.
- Focus: one focused widget; Tab and Shift+Tab cycle focusable widgets; Esc closes the topmost closable window.
- Keybinds are data (`keybinds.json`): defaults include Ctrl+I (Inventory), Ctrl+D (Datapad), Ctrl+C (Character Sheet), Ctrl+V (Map), Ctrl+P (Community), Ctrl+A (Abilities), Ctrl+O (Options) **[V]**. Only Inventory and Options windows are built in this project; the others open a "not implemented" stub window.
- Drag-and-drop uses a typed payload (`DragPayload::Item(ItemId)`, `DragPayload::Ability(AbilityId)`) with `can_drop` checks.

### 5.5 Theming [D]

A theme is a folder: `theme.json` (styles, 9-slice definitions, fonts), `palette.json` (semantic color tokens), and generated textures.

- **Palette tokens are semantic** (`panel.base`, `panel.bevel.dark`, `panel.bevel.light`, `trim.copper`, `border.steel`, `text.primary`, `text.secondary`, `ham.health`, `ham.action`, `ham.mind`, `ham.background`, `radar.enemy`, `radar.group`, `radar.waypoint.mission`, `radar.waypoint.generic`, and so on). Widgets never hold raw hex values. This mirrors how SWG used palette references **[V]** and enables user-selectable palettes later.
- **Styles are per widget kind and per state** (`normal`, `hover`, `pressed`, `disabled`, `focused`) and may reference 9-slice textures.
- **Opacity** is a first-class property: `window_opacity` (0.2–1.0) and `background_opacity` per widget **[V]**.
- **Hot reload** (debug builds): reloading a theme or layout file re-resolves styles without restarting. Failed reloads keep the previous theme and log the error with file and JSON path.

### 5.6 Data binding and the backend interface [D]

SWG used named code hooks inside UI files **[V]**. `tcgui` uses a similar idea with a typed key-value **model**:

- `UiModel` holds values under dotted keys (`player.ham.health.current`, `chat.channel.system.lines`).
- Widgets declare bindings in JSON (`"bind": {"value": "player.ham.health.current", "max": "player.ham.health.max"}`).
- A backend writes to the model through events; widgets update on the next frame. Destroying a widget removes its subscriptions.

```rust
// tcgui-backend-iface/src/lib.rs
pub trait GameBackend {
    /// Non-blocking. Called once per frame; returns everything that arrived.
    fn poll(&mut self) -> Vec<BackendEvent>;
    /// Fire-and-forget request from the UI.
    fn send(&mut self, cmd: UiCommand);
}

pub struct Pool { pub current: i32, pub max: i32, pub wounds: i32 }   // wounds reduce regenerable max

pub enum BackendEvent {
    LoginResult { ok: bool, message: String },
    CharacterList(Vec<CharacterSummary>),
    EnteredWorld { player: EntityId },
    Ham { entity: EntityId, health: Pool, action: Pool, mind: Pool },
    Posture { entity: EntityId, posture: Posture },
    StatesChanged { entity: EntityId, states: Vec<StateIcon> },
    TargetChanged { target: Option<EntityId>, name: String, con: ConLevel, distance_m: f32 },
    Chat { channel: ChatChannel, sender: String, text: String, timestamp_s: u64 },
    InventoryDelta(Vec<InventoryChange>),                              // add/remove/move/update
    InventorySnapshot { capacity: u32, items: Vec<ItemView> },
    RadarContacts(Vec<RadarContact>),                                  // see 6.3
    Toolbar { index: u8, slots: Vec<Option<ToolbarSlot>> },
    DamageNumber { entity: EntityId, pool: PoolKind, amount: i32 },
}

pub enum UiCommand {
    Login { user: String, pass: String },
    SelectCharacter(CharacterId),
    SendChat { channel: ChatChannel, text: String },
    MoveItem { item: ItemId, to: ContainerId, slot: Option<u32> },
    UseItem(ItemId),
    ActivateToolbarSlot { bar: u8, slot: u8 },
    SetTarget(Option<EntityId>),
    Logout,
}
```

`MockBackend` is deterministic from a seed: scripted login, three characters, HAM drift, scripted combat damage, periodic chat lines from several channels, an inventory of 30 generated items with occasional overload events, and 20 moving radar contacts. Same seed = same event stream, which makes snapshot and soak tests reproducible. A JSON wire format for a future Go server is specified as documentation only (T-42); no networking code is written in this project.

### 5.7 Rendering [D]

**2D UI renderer.** One dynamic vertex buffer used as a ring buffer; one index pattern for quads; one texture atlas (plus a font atlas); scissor rectangles for clipping; premultiplied-alpha blending; nearest-neighbor sampling by default for crisp bevels.

- **Batching:** consecutive quads with the same texture and clip merge into one draw call. Target ≤ 40 draw calls for the HUD scene.
- **Resource safety:** textures and buffers are created through a `ResourceTracker` that counts live GPU objects by type. At shutdown and at soak checkpoints, counts must return to expected baselines.
- **Text:** fonts rasterized to a bitmap atlas at build or load time from an open-license TTF (license recorded in `docs/adr/`). Layout supports measure, wrap, color runs, ellipsis. Glyph cache has a fixed capacity with LRU eviction so it cannot grow without bound.

**UI shader (WGSL), normative baseline:**

```wgsl
struct Globals { screen_size: vec2<f32>, _pad: vec2<f32> };
@group(0) @binding(0) var<uniform> globals: Globals;
@group(1) @binding(0) var atlas_tex: texture_2d<f32>;
@group(1) @binding(1) var atlas_smp: sampler;

struct VsIn {
    @location(0) pos: vec2<f32>,     // pixels, top-left origin
    @location(1) uv: vec2<f32>,
    @location(2) color: vec4<f32>,   // PREMULTIPLIED alpha
};
struct VsOut {
    @builtin(position) clip: vec4<f32>,
    @location(0) uv: vec2<f32>,
    @location(1) color: vec4<f32>,
};

@vertex
fn vs_main(in: VsIn) -> VsOut {
    var o: VsOut;
    let ndc = vec2<f32>(
        in.pos.x / globals.screen_size.x * 2.0 - 1.0,
        1.0 - in.pos.y / globals.screen_size.y * 2.0,
    );
    o.clip = vec4<f32>(ndc, 0.0, 1.0);
    o.uv = in.uv;
    o.color = in.color;
    return o;
}

@fragment
fn fs_main(in: VsOut) -> @location(0) vec4<f32> {
    let t = textureSample(atlas_tex, atlas_smp, in.uv);   // texture is stored premultiplied
    return t * in.color;
}
```

Blend state: `src = One`, `dst = OneMinusSrcAlpha` (premultiplied). The CPU premultiplies vertex colors; the asset loader premultiplies textures.

### 5.8 Retro 3D profile (`tcgui-scene`) [D]

Purpose: login/character backdrops and the T-08 performance scenario. Not a general engine.

- Vertex format: position, normal, uv, **baked vertex color** (lighting baked at import or generation time).
- `RetroProfile { flat_shading: bool, vertex_lit_only: bool, nearest_filter: bool, color_clamp: Option<f32> }`. Defaults replicate T-08's look; all are switchable at runtime so we can compare against references. T-08 clamped output to 0.8 to desaturate **[T08]**; keep as a tunable, default off until T-02 shows evidence.
- Features: perspective camera, indexed meshes, per-object transforms in a storage/uniform ring buffer, frustum culling, static batching by cell, 3-level LOD selection stub, procedural test content (terrain tiles, boxes, simple creatures).
- Flat shading is implemented with `@interpolate(flat)` on the color varying when `flat_shading` is on.

### 5.9 Persistence [D]

User UI state (window positions, sizes, opacity, collapsed state, toolbar contents, chat tab layout, inventory sort order) is saved as JSON in the platform config directory, written atomically (write temp, rename), versioned with a `schema_version` field, and validated on load. Corrupt files are renamed `*.bad` and defaults are used.

### 5.10 Optional SWG importer (feature `swg-import`) [D]

Off by default. Reads files **you** extract from your own SWG CD. Never bundled, never committed. Planned in two stages: (1) open-format textures (DDS/TGA) via the `image` crate; (2) a converter from SWG's declarative UI files to `tcgui` layout JSON, started only after the agent inspects real extracted files and documents what it finds (T-80 to T-83). Mesh formats are explicitly deferred because T-08's format names are unverified.

---

## 6. Visual and UX Specification

All coordinates are for the **base 1024×768** space and are measured or assumed per `docs/visual-spec.md`. The agent has recorded each measured value in `docs/visual-spec.md` with a reference ID.

### 6.1 Provisional palette (from T-08 and T-02 reference measurements)

| Token | Hex | Use |
|---|---|---|
| `panel.base` | `#2A2A2A` | Window background (Measured R-020) |
| `panel.bevel.dark` | `#1A1A1A` | Inner bevel shadow (Measured R-020) |
| `panel.bevel.light` | `#3A3A3A` | Edge highlight (Measured R-020) |
| `trim.copper` | `#8B7355` | Copper accent (Measured R-010) |
| `trim.copper.light` | `#A88F6E` | Copper highlight (Assumed derivative) |
| `border.steel` | `#4A4A4A` | Steel border (Measured R-013) |
| `text.primary` | `#DDDDDD` | Body text (Measured R-005) |
| `text.secondary` | `#999999` | Disabled/secondary (Measured R-005) |
| `button.up` / `button.down` | `#5A5A5A` / `#3A3A3A` | Button faces (Measured R-022) |
| `ham.health` | `#CC2222` | Health bar fill (Measured R-001) |
| `ham.action` | `#22CC22` | Action bar fill (Measured R-001) |
| `ham.mind` | `#2266CC` | Mind bar fill (Measured R-001) |
| `ham.background` | `#FFFFFF` | Bar background (Measured R-001) |
| `radar.enemy` | `#CC2222` | Enemy dots (Measured R-028) |
| `radar.group` | `#22CC22` | Group dots (Measured R-028) |
| `radar.waypoint.mission` | `#FFA500` | Mission arrows (Measured R-003) |
| `radar.waypoint.generic` | `#2266CC` | Generic arrows (Measured R-003) |

### 6.2 Window chrome

Beveled, metal-looking panel with a title bar (22px height), close button, optional minimize, resize handles on resizable windows, drag by title bar, opacity slider via Options. Generated procedurally by `asset-gen`: a 9-slice with a 1px dark / 1px light bevel pair and a copper/steel trim line. Exact bevel widths measured from R-020.

### 6.3 HUD (Ground HUD) element list, verified positions

| Element | Default placement (1024×768) | Behavior |
|---|---|---|
| **HAM monitor** | Top-left, 190×50 (Measured R-001) | Three bars (Health, Action, Mind), white backgrounds, colored fills; wound notches reduce the displayed regenerable maximum; numeric values on hover; posture icon beside bars. |
| **Wound indicator** | Within HAM monitor | Shows when `wounds > 0` on a pool. |
| **States display** | Under HAM monitor | Icons for positive/negative states, tooltip per icon. |
| **Target frame** | Top-center, 220×50 (Measured R-002) | Name, color-coded difficulty starburst, three HAM bars, distance. Hidden when no target. |
| **Radar** | Top-right, 128×128 (Measured R-003) | Rotates with heading; north indicator; target brackets; enemy red dots, group green dots; orange mission arrows and blue generic arrows clamped to the edge; range label; zoom with mouse wheel. |
| **Toolbar** | Bottom-center, 36px slot size (Measured R-001) | 12 slots (F1–F12) per bar; six bars (0–5); expandable to two rows; second row via Shift+F-key; drag abilities/items in; cooldown overlay; keybind label. |
| **Chat window** | Bottom-left, 400×200 (Measured R-005) | Movable, resizable, tabbed by channel; input line below; fades when idle per opacity setting. |
| **Data screen icons** | Bottom-right | Six icons: Map, Inventory, Datapad, Character Sheet, Community, Abilities; click or hotkey toggles the window. |
| **Floating damage numbers** | Over entities | Color per pool: red Health, green Action, blue Mind. |

### 6.4 Inventory window

- Title "Inventory"; header shows item count versus capacity (binding `inventory.count` / `inventory.capacity`).
- Body: scrollable grid of item icons with names, plus an equipment area (head, chest, legs, hands, weapon slots; names are assumed).
- Behaviors **[V]**: items with 0 hitpoints draw in red; when over capacity by one the Inventory data icon blinks; when over by two, a persistent red "Inventory Overloaded" message appears center-screen until resolved.
- Interactions: drag to move/equip; double-click to use/equip; right-click context menu (Use, Examine, Drop, Destroy-stub); tooltip on hover.
- Improvement over original **[V]** complaint: sort order (name, type, recent) persists across sessions.

### 6.5 Chat

Tabs per channel (System, Chat/Say, Group, Guild, Tell, plus a combined "All"). Each channel has a color token. Scrollback is a **bounded ring buffer** (default 2,000 lines per tab) so memory cannot grow. Input line supports history (Up/Down), `/command` passthrough as `UiCommand::SendChat`, and Enter to send. Timestamps are optional.

### 6.6 Login and Character Select

- **Login:** retro 3D backdrop; username and password fields (password masked); Login and Quit buttons; status line showing `LoginResult.message`; Enter submits; Tab cycles fields.
- **Character Select** (assumption A2): list of up to ten characters with name and planet/profession placeholder text; Play button; Create/Delete shown disabled with a "not implemented" tooltip.

### 6.7 Options window

Window opacity slider, UI scale (0.75–2.0), toggles for tooltips and chat timestamps, "Reset window layout" button, resolution/windowed/borderless selectors (applied via the platform layer), read-only keybind table.

### 6.8 Out of scope screens

Map, Datapad, Character Sheet, Community, Abilities: stub windows only.


---

## 7. Performance, Memory, and Determinism

### 7.1 Targets

Inherited from T-08 **[T08]** and made measurable **[D]**. "Mid-range PC" means an integrated or entry-discrete GPU from roughly the last six years with 16 GB RAM; the exact reference machine is recorded in the first perf report (T-70).

| Metric | Target | How measured |
|---|---|---|
| Frame rate | ≥ 60 FPS (p99 frame time ≤ 20 ms) in the "Coronet Downtown" scenario | Demo `--soak` CSV (frame time per frame) |
| Scenario | 100 dynamic objects + 500 static + HUD with all windows open | Scenario spec 7.2 |
| UI CPU cost | ≤ 1.0 ms/frame (layout + draw-list build) for HUD scene | `criterion` benchmark plus in-app stats |
| UI draw calls | ≤ 40 for HUD scene | Renderer stats overlay |
| UI vertex upload | ≤ 1 MB/frame | Renderer stats overlay |
| Memory | < 1 GB total process; aim < 300 MB | RSS sampled once per minute |
| Leak | Memory plateau: after a 5-minute warm-up, RSS slope < 1 MB/10 min over 30 min; live widget count, GPU texture/buffer counts, and binding counts return to baseline after each churn cycle | Soak assertions (T-72) |
| Startup | Window visible with login screen in < 2 s on the reference machine (debug off) | Timestamp log |

### 7.2 Soak scenario ("Coronet Downtown", from T-08 **[T08]**)

- 100 dynamic objects (spawn/despawn continuously at a steady rate), 500 static objects, 3 baked lights.
- UI: all four demo windows open; chat receiving about 5 lines/second; inventory receiving random deltas; windows opened/closed every 10 seconds (churn); tooltips shown/hidden.
- Duration: 30 minutes; `--soak-minutes N --seed S --csv out.csv`.
- Checkpoints every 60 seconds record: frame time p50/p99, RSS, live widgets, bindings, handlers, atlas slots in use, GPU textures/buffers/bind groups by tracker, chat lines retained.

### 7.3 Leak-prevention rules [D]

| Source of leaks | Rule |
|---|---|
| Widget trees | Arena + generational IDs; `destroy` removes subtree, bindings, handlers, drag state, focus/capture references. |
| Bindings and handlers | Owned by widget ID; auto-removed on destroy. |
| GPU resources | Created only via `ResourceTracker`; dropped in `Drop`; counts asserted in tests. |
| Textures | Reference-counted handles; atlas slots freed when count reaches zero and compacted lazily. |
| Text | Fixed-capacity LRU glyph cache; measure cache bounded. |
| Logs and chat | Bounded ring buffers. |
| Dynamic vertex data | Preallocated ring buffer; reallocation only on explicit growth with a logged event. |

A global **counting allocator** is enabled in the soak/test builds to measure heap bytes and allocation counts; per-frame allocations in the steady state should trend to zero in `Ui::layout` and `Ui::draw` (tracked as a metric, not a hard fail at first).

### 7.4 Determinism

Given the same JSON layouts, theme, viewport, UI scale, and mock-backend seed, `Ui::draw()` must return an identical `DrawList`. Time is injected (`Clock` trait) rather than read from the system, so animations and blink effects are testable.

---

## 8. Verification Strategy

Layered so most checks need no GPU:

| Layer | What it proves | Tool |
|---|---|---|
| Unit tests | Layout math, event routing, binding lifecycle, theme resolution | `cargo test` |
| Display-list snapshots | Visual structure of each screen, deterministic | `insta` JSON snapshots in `tests/golden/` |
| Schema validation | Every layout/theme file matches the schema | `cargo run -p tcgui-demo -- --validate-all` |
| Offscreen render | Real GPU output matches expectations; agents can view the image | `tcgui-demo --screenshot out.png --scene hud --size 1024x768` |
| Reference comparison | Fidelity against Pre-CU references | `tools/refcompare` writes side-by-side PNG; an agent (or you) reviews it |
| Benchmarks | CPU cost budgets | `cargo bench` (criterion) |
| Soak | Stability and leaks | `--soak-minutes 30` with assertions |

**GPU availability.** Many CI and agent environments have no GPU. Offscreen rendering tests are gated behind a `gpu-tests` feature and skip cleanly when no adapter is found (including software adapters if installed). All other layers must pass without a GPU.

**Fidelity review.** Image similarity metrics cannot judge "does it feel like SWG." Fidelity gates are therefore: (1) every token traces to a reference or is marked assumed; (2) a side-by-side image per screen is produced and committed in `docs/fidelity/`; (3) the stakeholder reviews those images at the end of M5 and M7.

---

# PART B — EXECUTION

## 9. Implementation Plan

### 9.1 Milestones

Each milestone is a set of task cards. "Size" is S (a short session), M (one full session), L (split if the agent runs out of context). Tasks without dependencies in the same milestone can run in parallel in separate worktrees.

| Milestone | Outcome | Playable? |
|---|---|---|
| **M0** Bootstrap and references | Workspace, rules, pinned versions, reference catalog, measured visual spec | No |
| **M1** Headless core | Widget arena, layout, input, theme, binding, display list, layout loader | No (tests only) |
| **M2** Renderer | Window, UI batch renderer, atlas, text, 9-slice, screenshots | A window with a panel |
| **M3** Widgets | Window, button, bars, text input, scroll, list/grid, tooltip, radar, chat log | Widget gallery |
| **M4** Backend interface | Trait, events, mock backend | No |
| **M5** Screens | Login, Character Select, HUD, Inventory, Chat, Options, modes, persistence | **Yes: demo playable with mock data** |
| **M6** Retro 3D | Scene crate, test scene, culling and batching | Backdrop in login |
| **M7** Perf and leaks | Stats, soak, allocator, benchmarks, fixes | Verified stable |
| **M8** Optional importer | Texture import; declarative-UI converter (feature-gated) | Optional |
| **M9** Docs and hand-off | API docs, guides, second-skin test, final report | Release candidate |

**Fast path to "playable" (your priority d):** M0 → M1 → M2 → M3 (subset) → M4 → M5. M6 can run in parallel with M3–M5 by a separate worker, but must not block M5.

### 9.2 Task cards

**Columns:** `Dep` = prerequisite IDs; `Acceptance` = commands or observable results an agent can run. Every task also requires `scripts/check.sh` to pass.

#### M0 — Bootstrap and references

| ID | Size | Task | Dep | Acceptance |
|---|---|---|---|---|
| T-00 | S | Verify the repo paths cited in T-08 exist (`sources/swg-pre-cu/`, `testbed/swg-phase3-combat/`). Create workspace `tcgui/` with empty crates, `AGENTS.md` (Appendix E), `scripts/check.sh`, `.gitignore` (include `references/`, `extracted-swg/`). Write `docs/repo-audit.md` listing which T-08 paths exist or not. | — | `./scripts/check.sh` passes on empty crates; audit file exists. |
| T-01 | L | **Reference acquisition.** Search the web for Pre-CU SWG screenshots using the source list in Appendix A and your own searches. Produce `docs/reference-catalog.md` with ≥ 30 entries: ID (`R-001`…), URL, date/era claim and how it was established, category (login, charselect, HUD, inventory, datapad, chat, options, radar, toolbar, tooltip, radial menu, window chrome), what it demonstrates. Only keep images locally under gitignored `references/`, and only if the site permits; otherwise record URL and description. Mark era as Pre-CU only with evidence (UI style, date, server, caption). | T-00 | ≥ 30 entries; ≥ 10 verified Pre-CU HUD images; ≥ 3 inventory; ≥ 3 chat; ≥ 2 login/char-select; each entry has an era justification. |
| T-02 | L | **Visual spec extraction.** From the references create `docs/visual-spec.md`: sampled palette hexes (with the reference ID and pixel location), window chrome measurements (title height, bevel widths, border), HAM bar geometry and colors, toolbar slot size, radar size/colors, chat window proportions, font x-height and line height at native resolution, default resolution evidence. Replace each **[P]** value in Section 6 with a measured value or "assumed." | T-01 | Every token in `palette.json` (T-13) references a visual-spec row or says "assumed"; ≥ 80% of tokens measured. |
| T-03 | S | Pin dependency versions (read the docs for each pinned version, especially `wgpu` and `winit`; note API differences from any example in this proposal). Write `docs/adr/0001-stack.md` (versions, MSRV, font choice with license text, allowlist). | T-00 | `cargo build --workspace` succeeds; ADR lists exact versions and font license. |

#### M1 — Headless core (`tcgui-core`)

| ID | Size | Task | Dep | Acceptance |
|---|---|---|---|---|
| T-10 | M | Geometry (`Vec2`, `Rect`, `Color` premultiplied), `WidgetId`, arena tree (create, destroy subtree, reparent, z-order). Write `docs/interfaces.md` freezing public types. | T-03 | Test: 10,000 create/destroy cycles return `stats().widgets` to baseline; stale ID returns `None`. |
| T-11 | M | Layout engine per 5.3: axis rule, percent values, min/max, row/column/grid containers, UI scale. | T-10 | Table-driven tests: ≥ 30 cases, including over/under-constrained errors; layout is idempotent. |
| T-12 | M | Input events, hit testing by z-order, pointer capture, focus, Tab order, Esc handling, World/Data mode with consumed-event reporting. | T-10 | Tests for overlapping windows, capture during drag, mode toggle, focus cycling. |
| T-13 | M | Theme/palette/style resolution with state styles, opacity, 9-slice definitions; JSON loading with errors that include file and JSON path; schema generation to `schemas/`. | T-10 | Loading a bad theme reports path of the bad field; schema files generated and committed; `--validate-all` stub works. |
| T-14 | M | Declarative layout loader (`.ui.json` → widgets) with strict unknown-field errors and widget registry. | T-11, T-13 | Round-trip test: load example from Appendix D; unknown field produces a path-annotated error. |
| T-15 | M | `UiModel` key-value store, typed values, bindings, change notification, unsubscribe on destroy. | T-10 | Test: destroying a bound widget removes its subscription (`stats().bindings` decreases); 10,000 bind/destroy cycles flat. |
| T-16 | M | `DrawList` generation, `Clock` trait, snapshot test harness (`insta`). | T-11, T-13 | Same inputs produce byte-identical JSON snapshots across two runs. |

#### M2 — Renderer

| ID | Size | Task | Dep | Acceptance |
|---|---|---|---|---|
| T-20 | M | `tcgui-platform`: winit window/event loop, wgpu device/surface, resize, minimize, DPI, clear color, input → `InputEvent`. | T-12 | `cargo run -p tcgui-demo` opens a window; minimize/restore and resize do not panic; clean exit with zero validation errors. |
| T-21 | M | UI batch renderer per 5.7: ring vertex buffer, quad batching, scissor, premultiplied blend, `ResourceTracker`. | T-16, T-20 | Benchmark: 10,000 quads in ≤ 2 draw calls with one texture; tracker counts match at exit. |
| T-22 | M | Atlas and texture manager with refcounts, stats, lazy compaction. | T-21 | Test: load/unload 1,000 textures; atlas slots in use return to baseline. |
| T-23 | M | Font rasterization to atlas, text measure/wrap/color runs, LRU glyph cache. | T-22 | Tests for measure/wrap; cache never exceeds configured capacity in a 1,000,000-glyph stress test. |
| T-24 | M | `tools/asset-gen`: procedural 9-slice bevel textures, button states, icons (placeholders), from palette tokens. | T-13 | Running the tool regenerates identical files (hash-stable); outputs listed in `themes/swg-precu/`. |
| T-25 | M | Offscreen rendering and `--screenshot` CLI; `tools/refcompare`. | T-21, T-24 | `tcgui-demo --screenshot out.png --scene gallery --size 1024x768` writes a PNG that shows a beveled panel (agent views the image). Skips cleanly if no GPU adapter. |

#### M3 — Widgets

| ID | Size | Task | Dep | Acceptance |
|---|---|---|---|---|
| T-30 | M | Panel, Label, Image/Icon, Button (normal/hover/pressed/disabled), Checkbox/Toggle. | T-14, T-23 | Snapshot tests per state; click test fires command. |
| T-31 | L | Window: title bar, drag, resize, close, minimize, z-order raise on click, opacity, clamp to viewport. | T-30, T-12 | Tests: dragging off-screen is clamped; z-order raise; opacity applied in draw list. |
| T-32 | M | ProgressBar and HAM bar (colored fill, white background, wound notches), Slider. | T-30 | Snapshot tests at 0%, 50%, 100%, with wounds. |
| T-33 | L | TextInput (cursor, selection, copy/paste, password mask, history hook), ScrollView and scrollbar. | T-30 | Tests for caret movement, selection, masking; scroll clamp; no per-keypress allocations beyond the string. |
| T-34 | L | TabBar; List/Grid with selection; typed drag-and-drop with payload and `can_drop`. | T-31, T-33 | Tests: drag item from grid slot A to B; invalid drop rejected; drag state cleared on window destroy. |
| T-35 | M | Tooltip (delay, bounds-aware placement) and ContextMenu. | T-31 | Test: tooltip stays within viewport at all four corners. |
| T-36 | L | Radar widget (rotation, contact kinds, edge clamping, range) and ChatLog (rich colored lines, bounded ring buffer, scroll lock). | T-30, T-33 | Radar snapshot tests for all contact kinds; ChatLog holds ≤ capacity after 100,000 appended lines. |
| T-37 | S | Widget gallery scene (`--scene gallery`) showing every widget in every state. | T-30–T-36 | Screenshot produced; stakeholder-viewable. |

#### M4 — Backend interface

| ID | Size | Task | Dep | Acceptance |
|---|---|---|---|---|
| T-40 | M | `tcgui-backend-iface`: trait, event and command types per 5.6; serde derives. | T-10 | Crate builds without GPU deps; types documented. |
| T-41 | M | `MockBackend` (seeded, deterministic) generating login, characters, HAM, combat, chat, inventory deltas (including overload), radar contacts, toolbar data. | T-40 | Same seed yields identical event stream (test); different seeds differ. |
| T-42 | S | `docs/backend-wire-format.md`: JSON wire format proposal for a future server (documentation only). | T-40 | Document lists every event/command with an example JSON message. |
| T-43 | S | Event-to-model adapter: applies `BackendEvent`s to `UiModel` keys; keys documented in `docs/model-keys.md`. | T-15, T-40 | Test: each event type updates its documented keys. |

#### M5 — Screens (first playable)

| ID | Size | Task | Dep | Acceptance |
|---|---|---|---|---|
| T-50 | M | Login screen layout + bindings + flow to Character Select. | T-31, T-33, T-43 | Scripted test: wrong password shows message; correct credentials advance. |
| T-51 | S | Character Select screen. | T-50, T-34 | Selecting a character sends `SelectCharacter` and enters HUD. |
| T-52 | L | HUD: HAM monitor, states, target frame, radar, toolbar (bars 0–5, two rows), data icons, damage numbers. | T-32, T-36, T-43 | HAM updates within one frame of event; toolbar Shift+F-key second row works; snapshot tests; screenshot for review. |
| T-53 | L | Inventory window with grid, equipment area, behaviors in 6.4 (red 0-HP, blink, overloaded message), context menu, persisted sort. | T-34, T-35, T-43 | Tests for each behavior; sort order survives restart (T-57). |
| T-54 | M | Chat window: tabs, colored channels, input history, send command. | T-36 | Test: sending creates `SendChat`; incoming lines routed to correct tab. |
| T-55 | M | Options window (opacity, scale, toggles, reset layout, resolution selectors, keybind table). | T-31 | Changing opacity changes draw-list alpha; reset restores defaults. |
| T-56 | S | Data/World mode (Alt) and keybinds from `keybinds.json`; stub windows for out-of-scope screens. | T-12, T-31 | Tests: Ctrl+I toggles Inventory; Alt toggles mode; stubs open/close. |
| T-57 | M | Persistence per 5.9 (atomic write, schema version, `.bad` fallback). | T-31 | Corrupt-file test falls back to defaults; positions survive restart. |
| T-58 | S | **Fidelity review package:** screenshots of every screen beside matching references in `docs/fidelity/`. | T-25, T-50–T-56 | Stakeholder review images exist for Login, HUD, Inventory, Chat. |

#### M6 — Retro 3D (parallel to M3–M5)

| ID | Size | Task | Dep | Acceptance |
|---|---|---|---|---|
| T-60 | L | `tcgui-scene`: camera, indexed meshes, vertex-color pipeline, `RetroProfile` toggles including flat shading. | T-21 | Screenshot of a vertex-lit cube field with each profile toggle; resources tracked. |
| T-61 | M | Procedural test content generator (terrain tiles, buildings, simple creatures) and the "Coronet Downtown" scene definition. | T-60 | Scene builds with 500 static and 100 dynamic objects; counts verified in test. |
| T-62 | M | Frustum culling, static batching by cell, LOD selection stub. | T-61 | Culling unit tests; draw-call count reduced versus unbatched baseline. |
| T-63 | S | Login/Character Select backdrop using the scene. | T-62, T-50 | Login screenshot shows backdrop behind panels. |

#### M7 — Performance and leaks

| ID | Size | Task | Dep | Acceptance |
|---|---|---|---|---|
| T-70 | M | Stats overlay and CSV logger; record reference-machine description in `docs/perf/machine.md`. | T-52 | Overlay toggles with F3; CSV has all checkpoint columns. |
| T-71 | M | `--soak-minutes N --seed S --csv` driver implementing 7.2 (window churn, chat, inventory deltas, object spawn/despawn). | T-61, T-70 | A 2-minute run completes and writes CSV. |
| T-72 | M | Counting allocator and plateau assertions (7.1/7.3). | T-71 | 30-minute run passes; deliberately injected leak (test hook) is detected. |
| T-73 | M | `criterion` benchmarks for layout and draw-list build; budget assertions. | T-16 | `cargo bench` prints results; budgets documented. |
| T-74 | M | Profile and optimize against the targets in 7.1; write `docs/perf/report.md`. | T-72, T-73 | Report shows each metric versus target with pass/fail. |

#### M8 — Optional importer (feature `swg-import`)

| ID | Size | Task | Dep | Acceptance |
|---|---|---|---|---|
| T-80 | S | **Stop and ask the stakeholder** for the extracted-file location. Document, in `docs/importer/findings.md`, the actual file types found (directory listing, extensions, sizes). No assumptions from T-08. | T-03 | Findings file lists real file types or states that no files were provided. |
| T-81 | M | DDS/TGA texture import via the `image` crate (or documented gap), into the atlas with premultiplication. | T-22, T-80 | Test using a locally generated DDS/TGA fixture (not SWG data). |
| T-82 | L | Converter from SWG declarative UI files to `tcgui` layout JSON, based only on observed files. | T-14, T-80 | Converts at least one real file; unconvertible attributes reported, not guessed. |
| T-83 | S | Mesh formats: decision record only (implement or defer). | T-80 | ADR written. |

#### M9 — Docs and hand-off

| ID | Size | Task | Dep | Acceptance |
|---|---|---|---|---|
| T-90 | S | `cargo doc` clean; every public item documented. | M5 | `RUSTDOCFLAGS="-D missing_docs" cargo doc --workspace` passes. |
| T-91 | M | Usage guide (`docs/usage.md`): adding a screen, adding a widget, binding to the model, connecting a real backend. | M5 | A fresh agent follows it to add a new window (tested in T-93). |
| T-92 | S | Theming guide (`docs/theming.md`). | T-13 | Lists all tokens and how to create a theme. |
| T-93 | M | **Second-skin test:** create a minimal alternate theme (e.g., plain dark) and verify no code changes are needed; add a new "About" window using only JSON + bindings. | T-91, T-92 | Both done with zero edits under `crates/*/src/` except test registration. |
| T-94 | S | Final verification report `docs/final-report.md`: each goal G1–G5 and each Definition of Done item with evidence. | all | Report complete; stakeholder sign-off requested. |

---

## 10. LLM Execution Protocol

### 10.1 Rules for every agent session

1. **One task card per session.** Read: the task card, `AGENTS.md`, `docs/interfaces.md`, and only the sections of this proposal that the card references.
2. **Tests first** where the card specifies tests. Write the failing test, then the implementation.
3. **Run `./scripts/check.sh`** (format, `clippy -D warnings`, tests, schema validation) before declaring done. Paste its summary into the task report.
4. **Do not change frozen interfaces** (`docs/interfaces.md`) without writing an ADR and flagging it as `INTERFACE-CHANGE` in the report.
5. **No new dependencies** outside the allowlist without an ADR.
6. **No `unsafe`** unless an ADR justifies it.
7. **No SWG-owned assets or code** in the repository, ever. Do not copy code or art from community mods. Extracted files stay in the gitignored `extracted-swg/` directory.
8. **Never invent facts about SWG.** If a value cannot be measured from a reference, mark it `assumed` in the visual spec.
9. **If blocked, stop and report** rather than working around. Write `BLOCKED: <reason>, <what was tried>, <what is needed>` in the task report.
10. **Keep changes small.** If a task is growing beyond roughly 600 changed lines, split it and report the split.

### 10.2 Swarm coordination [D, assumption A7]

- Task cards live in `docs/tasks/T-xx.md`; a single `docs/tasks/register.md` lists status (`todo`, `claimed:<worker>`, `done`, `blocked`). A worker claims a card by committing its name to the register before starting.
- Each task is built in its own git worktree/branch `task/T-xx`. Merge order follows dependencies; rebase on main before merging; the check script must pass after the rebase.
- After M1, interfaces are frozen. Workers on different crates can run in parallel safely. Two workers must not edit the same crate at the same time unless the cards say they touch disjoint files.
- A coordinator (you or one agent) runs `scripts/check.sh` on main after each merge and updates the register.

### 10.3 Task card template

```markdown
# T-xx: <title>
**Milestone:** Mx   **Size:** S/M/L   **Depends on:** T-aa, T-bb
## Goal
<one or two sentences>
## Read first
- AGENTS.md, docs/interfaces.md, Proposal §<n>
## Files you may create or edit
- crates/<crate>/src/<file>.rs
## Steps
1. <numbered, concrete>
## Acceptance (all must pass)
- `<command>` -> <expected result>
## Out of scope
- <things not to touch>
## Report (write to docs/tasks/T-xx.report.md)
- What changed; commands run with results; assumptions; follow-ups; BLOCKED notes
```

### 10.4 Guidance for mid-capability agents [D]

- Prefer boring code: explicit types, small functions, no clever macros or generics where an enum works.
- When an API from a pinned crate does not match an example here, read the pinned version's documentation and adapt. Examples in this proposal are illustrative, not guaranteed to compile.
- Show your work: log the exact command you ran and the exact output you relied on.
- Keep files under roughly 500 lines; split modules early.
- If you cannot make the acceptance command pass, do not weaken the test. Report it.

---

## 11. Risk Assessment

| Risk | Likelihood | Impact | Mitigation |
|---|---|---|---|
| Pre-CU references are hard to date, mixing NGE-era UIs | High | High | T-01 requires era justification; ≥ 10 verified Pre-CU HUD images; mods labeled Pre-NGE are not treated as authoritative. |
| Fidelity wrong because values are guessed | Medium | High | T-02 measurement requirement; fidelity review at M5 and M7; tokens traceable to references. |
| `wgpu`/`winit` API changes break examples | High | Medium | T-03 pins versions; agents read the pinned docs; examples marked illustrative. |
| No GPU in agent environments | High | Medium | Headless core and display-list snapshots cover most logic; GPU tests gated and skipped cleanly. |
| Agent drift or over-engineering | Medium | Medium | Small cards, allowlist, frozen interfaces, one task per session. |
| Scope creep into a game engine | High | High | Non-goals (4.2); `tcgui-scene` kept minimal; features need a card. |
| Legal exposure from SWG assets or mod code | Medium | High | No SWG-owned assets/code committed; importer optional and local; mods used only for concepts; licenses checked. |
| Importer formats unknown | High | Low | M8 optional; T-80 inspects real files first. |
| Performance targets missed | Medium | Medium | Early benchmarks (T-73), budgets per metric, soak tests in M7. |
| Swarm merge conflicts | Medium | Medium | Worktrees, disjoint files, interface freeze, register claims. |

---

## 12. Assumptions and Guesses Register

| ID | Item | Status |
|---|---|---|
| A1–A7 | See Section 2.2 | Open until confirmed |
| G-01 | I could not view screenshot images during research, only pages describing them. All positions and colors in Section 6 are provisional until T-02. | Open |
| G-02 | Exact semantics of the original UI's `PackLocation` and `PackSize` are unknown to me; `tcgui` uses its own anchor system. | Accepted |
| G-03 | Inventory capacity, equipment slot names, and character-select details are placeholders. | Open |
| G-04 | Original default resolution and font are unverified; we use 1024×768 base space (your requirement) and an open-license font. | Accepted |
| G-05 | Whether 60 FPS and the 100-object scenario match any real SWG measurement is unknown; they are self-imposed targets. | Accepted |
| G-06 | Crate versions and some API details may be stale in my knowledge. | Mitigated by T-03 |
| G-07 | Task sizes (S/M/L) are estimates for mid-capability agents and may need re-splitting. | Open |

---

## 13. Definition of Done

The project is done when all of the following are true and evidenced in `docs/final-report.md`:

1. `./scripts/check.sh` passes on a clean checkout (Windows and Linux).
2. Login → Character Select → HUD → Inventory and Chat run with the mock backend; all demo windows open, drag, resize, and close.
3. All layouts and themes validate against generated schemas; a new window and a new theme are added without Rust code changes (T-93).
4. Reference catalog (≥ 30 entries) and visual spec exist; fidelity review images for Login, HUD, Inventory, Chat are committed and reviewed by the stakeholder.
5. Performance report shows the 7.1 targets with pass/fail; a 30-minute soak passes the plateau assertions; an injected leak is detected by the test.
6. Backend interface documented, with a wire-format proposal; no networking code present.
7. No SWG-owned assets or copied mod code exist in the repository; the importer, if built, is feature-gated and off by default.
8. API documentation builds without warnings; usage and theming guides exist.

---

# APPENDICES

## Appendix A — Reference Sources Found During Research

These are starting points for T-01. I read page descriptions and text only; I could not inspect the images themselves. Licenses were not checked.

| Source | What it offers | Notes |
|---|---|---|
| MobyGames screenshots for *Star Wars Galaxies: An Empire Divided* (`https://www.mobygames.com/game/9532/star-wars-galaxies-an-empire-divided/screenshots/`) | Captioned screenshots including Pre-CU scenes (Mos Eisley cantina on the Bria server, circa 2004; player-built Wookiee village) | Captions give era evidence. |
| MMORPG.com Star Wars Galaxies screenshots (`https://www.mmorpg.com/star-wars-galaxies/screens/user/27514/display/2454`) | Dated screenshots (e.g., December 17, 2004) | Dates help establish era. |
| SWGEmu archive of the official site (2004–2007 pages) (`https://www.swgemu.com/archive/scrapbookv51/data/20070126114830/content_002.html`) | Publish notes (UI features such as crafting inventory display, inventory overload behavior) | Text evidence for behaviors. |
| SWGEmu archive of the official manual (`https://www.swgemu.com/archive/scrapbookv51/data/20070128130746/manual_printview_012.html` and `manual_detail_015.html`, `manual_detail_002.html`) | HUD element list, toolbar, radar, data screens, setup options | Primary text source for Section 6. Images may be in the page. |
| Official forum thread archive: list of UI and engine issues, updated 2005-10-26 (`https://www.swgemu.com/archive/scrapbookv51/data/20070213155953/index.html`) | Player-facing UI issues (inventory sort, respec window) | Evidence of behaviors and pain points. |
| MMORPG.com forum thread "Screens of the 'new' UI (or old from pre-cu)" (`https://forums.mmorpg.com/discussion/99112/screens-of-the-new-ui-or-old-from-pre-cu`) | Discussion of Pre-CU UI versus later UI, white HAM backgrounds | Possible screenshot links. |
| Anach's Pre-NGE UI mod (`https://github.com/Anach/Pre-NGE-UI`) | Readme describing Pre-CU/Pre-NGE style restoration, HUD file names | Concepts only; do not copy files. |
| `twistedatrocity/SWG-Mods` (`https://github.com/twistedatrocity/SWG-Mods`) | Real examples of the declarative UI files (attributes, opacity, palettes) | Concepts only; do not copy files. |
| `MrKraken/MrK_SWG_UI_Styles`, `Tatwi/swgemu_ui_color_theme_builder` | Palette and style structure; palette theme builder | Concepts only. |
| EQSanctum guides on changing UI colors and client modification | How palette files were edited | Era caution: some guides are post-NGE. |

**Search hints for T-01:** SWG Pre-CU screenshots by server and date; SWGEmu community screenshot threads; "Pre-CU UI" and "Pre-NGE UI" comparisons; fan sites and Wayback Machine captures of official screenshot-of-the-day pages; YouTube gameplay frames from 2003–2005 (record URL and timestamp, do not download).

## Appendix B — The Tailored Prompt This Proposal Answers

> **Role:** Senior Rust graphics and UI-toolkit architect who also writes specifications for LLM coding agents. Expertise: retained-mode GUI design, WGPU, MMO interface conventions of the early-to-mid 2000s, and headless/deterministic testing.
>
> **Context:** TCIndustries wants a reusable GUI system whose first skin recreates the Pre-CU Star Wars Galaxies interface. T-08 recommended a Rust/WGPU engine and UI framework but contained unverified claims. The stakeholder chose the full T-08 scope (renderer, asset pipeline, UI framework, perf/leak tests), a reusable framework, a demo of Login, HUD, Inventory, and Chat, a standalone mock backend with a defined interface, original recreated assets with an optional importer, and PC (Windows/Linux) at 1024×768 and up.
>
> **Audience:** First, the stakeholder deciding whether to approve; then Claude Code and a worker swarm of mid-capability agents who will implement from this document.
>
> **Format and length:** One markdown project proposal with author and version fields; long and comprehensive; decision half for the human, execution half (task cards, rules, acceptance commands) for agents.
>
> **Success criteria:** (1) Modular and easy for LLMs to extend, proven by a second-skin test; (2) fastest route to a playable demo (mock-data Login → HUD before any 3D or importer work); (3) every claim about SWG is tagged by evidence level and T-08 is corrected where research contradicts it.
>
> **Constraints:** Rust + WGPU; no SWG-owned assets or copied mod code; all tasks small and command-verifiable; flag every guess; do not invent SWG facts.
>
> **Examples:** T-08 (supplied). No prior proposal of this kind was supplied for this project.

## Appendix C — Glossary

| Term | Meaning |
|---|---|
| SWG | Star Wars Galaxies (Sony Online Entertainment, 2003–2011). |
| CU | Combat Upgrade (April 2005). **Pre-CU** is the target era. |
| NGE | New Game Enhancements (November 2005); its UI is not the target. |
| HAM | Health, Action, Mind: the three primary attribute pools. |
| HUD | Heads-up display. |
| 9-slice | Scaling technique keeping corners fixed and stretching edges and center. |
| Display list | Plain-data list of draw commands produced by the UI, consumed by the renderer. |
| ADR | Architecture decision record. |
| Worktree | A separate git working directory for parallel task branches. |

## Appendix D — Example Data Files

### D.1 HAM monitor layout (`layouts/hud_ham.ui.json`)

```json
{
  "schema_version": 1,
  "type": "window",
  "id": "ham_monitor",
  "title": null,
  "movable": true,
  "resizable": false,
  "layout": { "x": { "start": 8, "size": 190 }, "y": { "start": 8, "size": 50 } },
  "children": [
    {
      "type": "ham_bar", "id": "health",
      "layout": { "x": { "start": 4, "end": 4 }, "y": { "start": 4, "size": 12 } },
      "bind": { "value": "player.ham.health.current", "max": "player.ham.health.max", "wounds": "player.ham.health.wounds" },
      "style": "ham.health"
    },
    {
      "type": "ham_bar", "id": "action",
      "layout": { "x": { "start": 4, "end": 4 }, "y": { "start": 19, "size": 12 } },
      "bind": { "value": "player.ham.action.current", "max": "player.ham.action.max", "wounds": "player.ham.action.wounds" },
      "style": "ham.action"
    },
    {
      "type": "ham_bar", "id": "mind",
      "layout": { "x": { "start": 4, "end": 4 }, "y": { "start": 34, "size": 12 } },
      "bind": { "value": "player.ham.mind.current", "max": "player.ham.mind.max", "wounds": "player.ham.mind.wounds" },
      "style": "ham.mind"
    }
  ]
}
```

### D.2 Palette tokens (`themes/swg-precu/palette.json`)

```json
{
  "schema_version": 1,
  "tokens": {
    "panel.base":         { "hex": "#2A2A2A", "source": "T-08, assumed" },
    "panel.bevel.dark":   { "hex": "#1A1A1A", "source": "T-08, assumed" },
    "panel.bevel.light":  { "hex": "#3A3A3A", "source": "T-08, assumed" },
    "trim.copper":        { "hex": "#8B7355", "source": "T-08, assumed" },
    "border.steel":       { "hex": "#4A4A4A", "source": "T-08, assumed" },
    "text.primary":       { "hex": "#DDDDDD", "source": "T-08, assumed" },
    "ham.background":     { "hex": "#FFFFFF", "source": "assumed; update from visual-spec row" }
  }
}
```

(Every `source` field must be updated by T-02 to a `R-xxx` reference ID and pixel location, or remain "assumed.")

### D.3 Style entry (`themes/swg-precu/theme.json`, excerpt)

```json
{
  "schema_version": 1,
  "fonts": { "default": { "file": "fonts/default.ttf", "size_px": 12 } },
  "styles": {
    "window": {
      "normal": { "nine_slice": "window_frame", "background": "panel.base", "background_opacity": 0.9 }
    },
    "button": {
      "normal":   { "nine_slice": "button_up",   "text": "text.primary" },
      "hover":    { "nine_slice": "button_up",   "tint": "trim.copper.light", "text": "text.primary" },
      "pressed":  { "nine_slice": "button_down", "text": "text.primary" },
      "disabled": { "nine_slice": "button_up",   "text": "text.secondary", "opacity": 0.6 }
    }
  }
}
```

## Appendix E — `AGENTS.md` Template (create in T-00)

```markdown
# AGENTS.md — rules for tcgui

1. Work on ONE task card per session (docs/tasks/T-xx.md). Read the card, this file, docs/interfaces.md.
2. Run ./scripts/check.sh before finishing. Paste the summary in your report.
3. Do not change frozen interfaces without an ADR and an INTERFACE-CHANGE note.
4. Allowed crates are listed in docs/adr/0001-stack.md. Do not add others without an ADR.
5. No unsafe code without an ADR.
6. Never commit SWG-owned assets or code from community mods. extracted-swg/ and references/ are gitignored.
7. Never invent facts about Star Wars Galaxies. Mark unmeasured values "assumed".
8. If blocked: stop and write BLOCKED: <reason>, <tried>, <needed> in your task report.
9. Keep tasks small; split if the change exceeds about 600 lines.
10. Do not weaken tests to make them pass. Report the failure instead.
11. Report: what changed, commands run with results, assumptions, follow-ups.
```

## Appendix F — Document History

| Version | Date | Author | Notes |
|---|---|---|---|
| 1.0 | 2026-10-23 | Claude Sonnet 5.5 | Initial proposal based on T-08 and interview answers. |

*End of document.*
