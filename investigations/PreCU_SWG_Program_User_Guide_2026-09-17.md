# Pre-CU SWG Program — User Guide and Program Documentation

| Field | Value |
|---|---|
| **Document ID** | TCIND-PDR-2026-09-17-02 |
| **Title** | Pre-CU SWG Program: Purpose, Operation, TCIndustries Integration, Development History, and SWG Background — User Guide |
| **Author** | Buffy — Codebuff coding agent (Freebuff client). LLM: **GLM, version 4.6**, trained by Z.ai |
| **Audience** | Project owner and any user who needs to build, run, or understand the program |
| **Date** | 2026-09-17 |
| **Status** | EVIDENCE (operational documentation; per `AGENTS.md` this document makes no design decisions and holds no design authority) |
| **Program location** | `testbed/swg-phase3-combat/swg-server/` (server) · `testbed/swg-phase3-combat/swg-godot/` (client) |
| **Related documents** | `sources/swg-pre-cu/README.md` (predecessor provenance) · `testbed/swg-phase3-combat/swg-server/README.md` (technical README) · `proposals/TCIndustries_SWG_Testbed_Reuse_Proposed_Ruling.md` (HD-TST-01 basis) · `AGENTS.md` · `proposals/TCIndustries_Manifest.md` |
| **Verification basis** | All run commands below were executed successfully on the project's Windows host during 2026-09-17 sessions (full-chain `CHAIN:ALL-GREEN`, `PHASE10_EXIT:0`); see `investigations/SWG_Phase10_Implementation_Verification_Report_2026-09-17.md` |

---

## 1. What the Pre-CU SWG program is, and its purpose

The **Pre-CU SWG program** is a ground-up, systems-faithful recreation of the sandbox systems of *Star Wars Galaxies* as they existed in the **Pre-CU era** (before the 2005 "Combat Upgrade"). It is implemented as a Go game server with a Godot 4 client, and it is classified — per its own README — as a **fan-made, non-commercial, educational preservation project**.

The program has two purposes, one inherited and one current:

1. **Original purpose (predecessor project):** to faithfully recreate the Pre-CU SWG sandbox — its character model, skill system, economy, services, and society systems — as a working, testable implementation of a written design specification (`SWG_PreCU_GDD.md`).
2. **Current purpose (within TCIndustries):** to serve as a **patterns-only mechanical testbed** for the TCIndustries project. It builds and exercises the *mechanical patterns* TCIndustries cares about — interdependence, player-driven economy, skill-based progression, service professions — so that TCIndustries design discussions can be grounded in working, measured evidence instead of speculation.

It is **not** a shippable game, not a TCIndustries build, and not design authority for anything. Everything it produces has the documentation status **PROTOTYPE / EVIDENCE**.

---

## 2. Program components

| Component | Technology | Location | State |
|---|---|---|---|
| Game server | Go 1.22, SQLite, WebSocket (`gorilla/websocket`), REST | `testbed/swg-phase3-combat/swg-server/` | Fully working through Phase 10; all integration suites green (2026-09-17) |
| Client | Godot 4 (GDScript) | `testbed/swg-phase3-combat/swg-godot/` | Early shell: login, character select, 3D world scene with HAM display; scenes/scripts present in the current fork |
| Integration test clients | Go (one per phase: `cmd/testclient`, `cmd/phase1test` … `cmd/phase9test`, `cmd/phase10test`) | `swg-server/cmd/` | Fully working; these, not the Godot client, are the primary way the program is operated and verified |
| Report generator | Go (`cmd/phase10report`) | `swg-server/cmd/` | Working; emits the balance validation report and B4 service-XP characterization |
| Suite runners | POSIX shell (`run_phase9.sh`, `run_phase10.sh`, `run_chain.sh`) | `swg-server/` | Working on Windows (MSYS/Git Bash) and POSIX hosts |

Architecture in one paragraph: an **authoritative server** — all game state (position, HAM, inventory, economy) is computed server-side; REST (`/api/*`) handles accounts, characters, and actions; a WebSocket (`/ws`) handles the real-time world (movement, spatial interest management, chat, combat updates); SQLite persists everything; per-phase integration test clients drive real end-to-end scenarios against a live server.

---

## 3. How to use and run the program

### 3.1 Prerequisites

- **Go 1.22+** — [go.dev/dl](https://go.dev/dl/)
- **A C toolchain (gcc)** on all hosts — the SQLite driver requires cgo. Linux/macOS usually have gcc preinstalled. **Windows:** install MinGW-w64 (e.g. `winget install BrechtSanders.WinLibs.POSIX.UCRT`) and place it at a **space-free path** such as `C:\mingw64` (spaces in the toolchain path break the build).
- **Godot 4.2+** — only if you want to run the client ([godotengine.org](https://godotengine.org/download/)).
- A POSIX-compatible shell for the suite runners (`sh`): Git Bash / MSYS2 on Windows.

### 3.2 Running the server (interactive play / inspection)

```bash
cd testbed/swg-phase3-combat/swg-server
go mod tidy
go run ./cmd/server/
```

- The server listens on `http://localhost:8080`.
- The SQLite database (`swg.db`) is **auto-created on first run**. Delete the file for a guaranteed-fresh world (suites assume clean state).
- Health check: `GET /health`.

### 3.3 Environment variables

| Variable | Effect |
|---|---|
| `TESTBED_FAST_CYCLE=1` | Remaps long timers at init to test-scale (e.g. the 24 h economic-snapshot job → 60 s; lair/mission windows similarly shortened). Required by the phase suites. |
| `TESTBED_COMBAT_DEBUG=1` | Enables `COMBAT-DEBUG` combat-log output (gated since Phase 10 B6; off by default). |
| `P10_SERVER_URL` | Overrides the server URL used by `cmd/phase10test` (lets the Phase 10 suite run against a server on any port). Phase 1–9 clients hardcode `:8080`. |

### 3.4 Verifying a build (always do this first)

```bash
export CGO_ENABLED=1                    # Windows Go defaults to CGO_ENABLED=0
export PATH="/c/mingw64/bin:$PATH"      # Windows: space-free MinGW-w64 location
go vet ./... && go test ./...
```

If a Windows binary exits at startup with *"Binary was compiled with 'CGO_ENABLED=0', go-sqlite3 requires cgo to work"*, the cgo/PATH setup above was skipped.

### 3.5 Running the per-phase integration tests

Each phase has an end-to-end client that registers accounts, creates characters, and exercises that phase's systems against a live server (start the server first in a second shell, on a fresh `swg.db`):

```bash
go run ./cmd/testclient/     # Phase 0: accounts, characters, 9-species HAM, world entry
go run ./cmd/phase1test/     # Movement, spatial interest management, spatial chat
go run ./cmd/phase2test/     # Skills: XP pools, training, prerequisites, persistence
go run ./cmd/phase3test/     # Combat core
go run ./cmd/phase4test/     # Resources & crafting
go run ./cmd/phase5test/     # Economy infrastructure
go run ./cmd/phase6test/     # Medic/Entertainer services (wounds, BF, buffs)
go run ./cmd/phase7test/     # Civic systems (cities, guilds, mail)
go run ./cmd/phase8test/     # Faction & PvP
go run ./cmd/phase9test/     # Elite professions & missions (full 8-test suite)
go run ./cmd/phase10test/    # Phase 10: §1.3 success-criteria proxies + telemetry
```

Each prints per-test results and ends with an exit code (`0` = all pass). The Phase 10 client additionally accepts `P10_SERVER_URL` and prints the five §1.3 proxy verdicts (PASS / FAIL-with-telemetry / NOT-MEASURABLE).

### 3.6 Running the full suites (recommended operating mode)

```bash
# Full Phase 9 suite (8 tests), ~20 min on Windows:
TESTBED_FAST_CYCLE=1 sh run_phase9.sh

# Phase 10 suite on its own dedicated port (does not require :8080):
TESTBED_FAST_CYCLE=1 sh run_phase10.sh

# Full-chain regression: Phases 1–9 sequentially, each on a fresh DB with its
# own server lifecycle, then Phase 10. Allow well over an hour on Windows;
# the phase 9 suite alone is ~20 minutes:
TESTBED_FAST_CYCLE=1 sh run_chain.sh
```

`run_chain.sh` is the program's master verification: a clean run ends with

```
=== CHAIN RESULTS: p1=0 p2=0 p3=0 p4=0 p5=0 p6=0 p7=0 p8=0 p9=0 ===
CHAIN:ALL-GREEN
PHASE10_EXIT:0
```

**Port hygiene:** phase 1–9 clients connect to `:8080` — nothing else may hold that port when their server boots (stale servers cause silent bind failures). `run_chain.sh` reaps listeners between phases; `run_phase10.sh` avoids the conflict entirely via a dedicated port.

### 3.7 Reports and telemetry (read-only)

- **Balance validation report** (fork values vs. historical §33 bands, every row OBSERVED / CODE-CONSTANT / DERIVED / NOT-MEASURABLE):
  ```bash
  go run ./cmd/phase10report/
  ```
  A committed copy lives at `testbed/swg-phase3-combat/swg-server/docs/phase10_balance_report.md`; the service-XP rate characterization and options paper at `docs/phase10_service_xp_options.md`.
- **Live telemetry REST** (all read-only, GET-only): `GET /api/phase10/telemetry`, `GET /api/phase10/anomalies`, `GET /api/phase10/jobs`.
- The REST surface also includes the documented account/character/skills endpoints (see the technical README's API table); Phases 3–9 added combat, resources, economy, services, civic, faction, and mission routes — consult `internal/server/server.go` for the complete live route list.

### 3.8 Running the Godot client (optional visual shell)

1. Open Godot 4 and import `testbed/swg-phase3-combat/swg-godot/`.
2. Add `scripts/GameState.gd` as an autoload named `GameState` (Project Settings → Autoload).
3. Press F5. Use the login/register screen, character select, then enter the world.

The client connects to `http://localhost:8080` and shows HAM state in the 3D world scene. **The Godot client is a shell, not the program's verification surface** — most scripted content (combat, crafting, missions) is exercised by the Go test clients, not by the GUI. The archived Phase 0–2 snapshots note that scene files were to be created in-editor; the current fork's `scenes/` directory contains them (Login, CharacterSelect, World).

### 3.9 Troubleshooting quick list

| Symptom | Cause / fix |
|---|---|
| Binary dies at startup mentioning `CGO_ENABLED=0` | Set up the C toolchain per §3.2/§3.4; on Windows use a space-free MinGW path and `CGO_ENABLED=1`. |
| Server starts then exits silently | Port 8080 is held by another process (often a stale test server). Free the port or use `run_phase10.sh`'s dedicated port for Phase 10 work. |
| Suite client can't connect | Start the server first, on a fresh `swg.db`; confirm `TESTBED_FAST_CYCLE=1` for any suite longer than Phase 0–2. |
| Test suite appears to hang on windows/timers | Timers are fast-cycled; the phase 9 and full-chain suites legitimately take 20–90 minutes. |
| Wrong species/skill data | Database is stale — delete `swg.db` and restart the server (world data is recreated by server bootstrap). |

---

## 4. How the program integrates with TCIndustries

### 4.1 The relationship in one sentence

TCIndustries is an **original-IP persistent-sandbox MMORPG design project**; the Pre-CU SWG program is its **mechanical testbed** — a separate, SWG-derived codebase used to generate working evidence about the systems TCIndustries is designing.

### 4.2 The governing ruling (HD-TST-01)

The reuse of this codebase by TCIndustries is governed by recorded human decisions, not by convenience:

- On **2026-09-14** the project owner selected **path (b)**: the SWG Pre-CU engine serves as a **patterns-only mechanical testbed** for TCIndustries investigation and prototyping. This is recorded as ruling **HD-TST-01** (`proposals/TCIndustries_SWG_Testbed_Reuse_Proposed_Ruling.md`, filed into governance).
- **Patterns-only** means: the testbed is authorized to *exercise structural patterns* (survey → harvest → craft-with-experimentation → equip; ledger economy; service interdependence), **not** to define TCIndustries design. It is explicitly **not** "reuse as-is".
- The fork was created from the predecessor tree `sources/swg-pre-cu/code-phase0-3-claudecode/` after a due-diligence pass, with a **de-SWG hygiene pass** (generic IDs, no Star Wars names/lore/species/places), and left the original in `sources/` unmodified for provenance.

### 4.3 What the testbed is allowed to mean for TCIndustries

Under `AGENTS.md` status discipline, every observable behavior of this program is **PROTOTYPE / EVIDENCE**:

| Testbed output | TCIndustries use |
|---|---|
| Phase-suite pass results | Proof that a mechanical pattern *can* be implemented and verified end-to-end |
| Balance report (`phase10report`) | OBSERVED fork values beside **HISTORICAL** Pre-CU bands (§9.7, §33) — comparison baselines, never tuning targets |
| Interdependence telemetry (`interdependence_events`) | Direct evidence input for the deferred **Economic Interdependence Core (EIC)** work — itself gated by rulings HD-EIC-05/07 |
| Service-XP characterization + options paper | Documents the residual from Phase 9 decision 6c; any adoption is a future explicit human numeric authorization |
| Anomaly review lists | A process pattern for future TCIndustries economy auditing (human review only, never auto-action) |

### 4.4 Hard boundaries (what integration forbids)

- **No promotion:** nothing in the testbed becomes TCIndustries canon by existing or passing tests. Status changes require explicit recorded human approval.
- **No numeric tuning:** economic/progression numbers may not be tuned without explicit human authorization (standing governance boundary; HD-TST-01).
- **No SWG IP leakage:** no Star Wars/Lucasfilm names, lore, factions, species, or places in the fork or in anything derived from it for TCIndustries.
- **No EIC shaping:** the testbed may inform, but never shape, EIC — new shaping requires human-supplied structural boundaries (HD-EIC-07).
- **Workflow:** testbed work follows *proposal → explicit owner authorization → implementation → verify-don't-claim evidence report* (the Phase 10 cycle of 2026-09-17 is the worked example).

---

## 5. History of the program

### 5.1 Predecessor origins (Groups A, B, C)

The program descends from a predecessor effort preserved byte-for-byte under `sources/swg-pre-cu/` (provenance recorded there; all of it **HISTORICAL**):

- **Group A** — a Perplexity chat thread (author: Perplexity — GPT-5.2) produced the **34-section systems-design GDD** (`SWG_PreCU_GDD.md`), written as a build specification for an autonomous AI agent, plus working **Phase 0–2 builds** (Go 1.22 server + SQLite + WebSocket; Godot 4 client), snapshotted as zips at each phase boundary. Content: server foundation/HAM (Phase 0), movement/spatial/chat (Phase 1), skills/professions (Phase 2).
- **Group B** — an arena.ai shared chat produced two **crafting-core tech demos** (Vite + React + TypeScript): one by the **seed-2.1-pro-preview** LLM (assembly + experimentation engine with a pure-TS core designed for 1:1 porting) and one by **claude-opus-5-max** (architecture recommendation + live demo with galaxy resource generation and an FSM craft bench). The seed-2.1 demo is the prototype referenced by TCIndustries GDD item PROT-002 (identity confirmed 2026-09-12).
- **Group C** — files received from the project owner with prior history unrecorded: a later evolved server tree carrying **Phases 0–3** (adds the combat core: HAM resolution, postures, wounds/battle fatigue, incapacitation/revive/clone, creatures + lairs), built against a revised **GDD v2** and embedding a `CLAUDE.md` **verify-by-running methodology**; an early **Track A/B prototyping scaffold** with reference functions (combat resolution, resource-spawn generator, inflation sim, experimentation outcomes, XP curve); and a small **Java Swing skill calculator** implementing Pre-CU skill rules.

### 5.2 Development method

The program was developed in a distinctive, fully documented style — **phase-gated autonomous-agent development**:

1. A written GDD defines each phase's **exit criteria** (GDD §30), e.g. Phase 2: *"A player can earn XP from a placeholder action, spend skill points at a trainer NPC, and see the skill persist across logout/login."*
2. An AI agent implements the phase, then writes an **integration test client** that drives the exit criteria against a live server over real network protocols.
3. Nothing is claimed without running: the methodology (inherited from the predecessor's `CLAUDE.md`) is **verify-don't-claim** — `go vet` clean, fresh database, real test output pasted into records, one commit per verified step.
4. Each phase boundary was snapshotted (zip + extracted tree) so provenance is byte-exact.

### 5.3 Adoption into TCIndustries (repository chronology)

| Date | Event |
|---|---|
| 2026-09-12 | TCIndustries documentation architecture established; corpus imported; **SWA predecessor material stored** under `sources/swg-pre-cu/` with a provenance README; Group B/PROT-002 identity confirmed; due-diligence checklist extended |
| 2026-09-14 | **HD-TST-01 recorded**: owner selects path (b) — SWG engine as patterns-only mechanical testbed; due diligence + de-SWG preconditions defined |
| (post-ruling) | Fork created at `testbed/swg-phase3-combat/swg-server/`; de-SWG pass (generic zone/IDs); Phases 3–6 (combat core, resources/crafting, economy, services) built and suite-verified |
| 2026-09-16 | **Phase 7** (civic: cities, guilds, groups, mail) and **Phase 8** (faction & PvP) executed and verified; Phase 9 proposal filed; owner numeric authorizations recorded |
| 2026-09-17 | **Phase 9** (22 elite professions, 28/504-box registry, missions incl. bounty escrow) verified (`PHASE9_EXIT:0`); **Phase 10** proposed (v0.2), owner-authorized, implemented (B1–B7), and verified: full chain `CHAIN:ALL-GREEN` + `PHASE10_EXIT:0` (commits `f9852ec`…`2730c9a`) |

The program is therefore at **Phase 10 complete**: ten phase scopes implemented, each with a green end-to-end suite, on a single verified code state.

---

## 6. Star Wars Galaxies background — and why TCIndustries uses it

### 6.1 What Star Wars Galaxies was

*Star Wars Galaxies* (SWG) was a massively multiplayer online role-playing game launched in **June 2003** by Sony Online Entertainment with LucasArts, set in the Star Wars universe. The **"Pre-CU"** designation refers to the period before the **Combat Upgrade** (March 2005) and the later **New Game Enhancements** (November 2005), two redesigns that replaced the original progression and combat models. The Pre-CU era is remembered — and studiously preserved by fan projects — for being one of the most ambitious **sandbox** MMO designs ever shipped: it de-emphasized theme-park questing in favor of a player-driven society.

### 6.2 The Pre-CU systems this program recreates

The fork recreates the *mechanical systems* of Pre-CU SWG (with a generic zone replacing Tatooine and generic IDs replacing SWG vocabulary, per the de-SWG pass):

| Pre-CU system | Recreation in this program |
|---|---|
| **HAM attribute model** | Health/Action/Mind bars, each with a strength pool — 3×3 attributes (Health, Strength, Constitution, Action, Quickness, Stamina, Mind, Focus, Willpower) with 9 species modifiers, verified against spec values |
| **Skill-based progression** | 250 skill points; skill boxes in trees (Novice → Tiers I–IV → Master); prerequisite chains; 6 basic professions grown to **28 professions / 504 boxes** including 22 elites; trainers; skill-dropping with cascades; XP never refunded on drop |
| **Resources & crafting** | Surveying, resource spawns with attribute ranges, harvesters, schematics, crafting with experimentation, item provenance (`crafted_items`) |
| **Player-driven economy** | Credits ledger with faucet/sink tagging, vendors, bazaar, housing/structure placement, append-only sale-price history (`market_records`), wealth/Gini telemetry |
| **Service interdependence** | Medic wound healing (typed wounds healed from HAM pools), Entertainer buffs (battle fatigue recovery), image designer, smuggler, stim distribution — the famous "players need other players" loop |
| **Civic systems** | Cities, guilds, groups, mail, permissions |
| **Faction & PvP** | Faction standing, PvP flagging, duels, death penalties, ranks, bases |
| **Missions & bounty hunting** | Mission terminals, lifecycles (available → accepted → completed/expired), bounty contracts with **escrow** and payments |
| **World ecology** | Creatures, lairs with damage/regen/destruction/relocation cycles |
| **Economic observation** | Nightly snapshot jobs, price volatility, interdependence event tracking (the Phase 10 telemetry layer) |

### 6.3 How and why TCIndustries uses SWG for this

**Why:** TCIndustries' design goals — player-driven society, economy, identity, reputation, **interdependence**, and emergent gameplay — are precisely the qualities Pre-CU SWG is historically famous for. It is the closest large-scale precedent in MMO history for the society TCIndustries wants to design. But TCIndustries is an **original IP**: it cannot copy SWG content, and its own governance forbids copying SWG *numbers* or treating SWG design as authority. What it can legitimately borrow is **evidence about mechanical patterns** — whether a system *works*, how players flow through it, where the economy leaks.

**How** (the pipeline, end to end):

1. **Patterns-only authorization** (HD-TST-01): the fork exists to test patterns, with SWG content stripped and values marked NON-CANONICAL / `[PROVISIONAL]`.
2. **Build the pattern faithfully** with a full integration suite, so the evidence is trustworthy (verify-don't-claim).
3. **Measure** the running system: telemetry jobs, interdependence events, balance reports comparing observed values against *HISTORICAL* Pre-CU bands (§9.7, §33) — as baselines for discussion, never as targets.
4. **Feed TCIndustries design work**: evidence reports inform canonical decisions, which remain exclusively with the human project owner. The Phase 10 cycle (proposal → authorization → implementation → verification report) is the template.
5. **Keep the boundaries**: no SWG names/lore/numbers cross into TCIndustries; the testbed remains a separate, fan-made, non-commercial educational recreation.

**Legal note:** *Star Wars* and *Star Wars Galaxies* are trademarks of Lucasfilm Ltd. This program is a fan-made, non-commercial, educational preservation effort and is not affiliated with or endorsed by Lucasfilm. TCIndustries is an original intellectual property and likewise unaffiliated.

---

## 7. Status and compliance summary

- **Program state:** Functioning Prototype, engineering-verified through Phase 10 (2026-09-17 full chain green). Not a product; no deployment/operations story beyond a dev server.
- **Documentation status of everything herein:** EVIDENCE. This guide instructs operation; it confers no design authority and promotes nothing into canon.
- **Canonical reading order for TCIndustries context:** `AGENTS.md` → `proposals/TCIndustries_Manifest.md` → Master GDD v1.1.1 → this guide for the testbed itself.

---

*Prepared by Buffy (Codebuff/Freebuff coding agent; LLM: GLM 4.6, Z.ai), 2026-09-17. All run commands in §3 were executed successfully on the project's Windows host during the 2026-09-17 sessions; historical claims in §5–6 cite `sources/swg-pre-cu/README.md`, the predecessor phase READMEs, `proposals/TCIndustries_SWG_Testbed_Reuse_Proposed_Ruling.md`, `proposals/TCIndustries_Manifest.md`, and the repository commit record.*
