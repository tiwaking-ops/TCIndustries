---
doc_id: SWG-DUE-DILIGENCE-EVIDENCE-2026-09-14
title: SWG Pre-CU Code Due-Diligence — Evidence Log (Checklist Steps 1–4, Groups A–C)
version: "1.2"
date: 2026-09-14
status: EVIDENCE / ANALYSIS (no design authority)
authority_class: null
assessor: Muse Spark (opencode/muse-spark-1.3-contributor-free) via OpenCode
version_note: "1.2 adds §§6–7 (Steps 6–7 coverage map + reuse recommendation, 2026-09-14). §§1–5 unchanged."
human_approved: true
human_approval_note: Owner confirmed read + approved 2026-09-14 (chat); approval covers
  this evidence record only — grants no design authority, no status change, no reuse decision.
supersedes: []
superseded_by: null
controlling_documents: [MASTER-GDD-v1.1.1, HUMAN-RULINGS-REGISTER, SWG-PRE-CU-README]
reading_priority: reference
---

# SWG Pre-CU Code Due-Diligence — Evidence Log (Steps 1–4, Groups A–C)

**Assessor:** Muse Spark (`opencode/muse-spark-1.3-contributor-free`) via OpenCode —
evidence gathering and analysis only; all authority decisions remain human.

**Procedure:** `proposals/SWG_Code_Due_Diligence_Checklist_v0.1.md`, Steps 1–4 ONLY.
File-version note: the checklist file on disk identifies as **v0.2** ("extends v0.1 scope to
all later arrivals"); the material under `sources/swg-pre-cu/` (Groups A–C) is unchanged by
that revision, so Steps 1–4 below execute identically under either version.

**Role:** evidence gathering only. No code changed, nothing fixed, no fork, no reuse decision,
no status change. Steps 5–7 (fit, coverage, recommendation) NOT executed. EIC, numerics,
provenance, and GDD edits out of scope and untouched.

**Isolation:** every build/test/run below executed against copies under `C:\temp\swg-evidence\`
(temp working area, outside the repo). The repo's `sources/swg-pre-cu/` trees were only read.
`git status` after the work shows no modification to any tracked file (two pre-existing
untracked items only — see §1.5).

**Environment (recorded, not endorsed):** Windows / PowerShell 5.1; Go, Java, and Godot absent
natively at session start; Docker Desktop started during the session; Godot installed via
winget during the session (see §2). Native: node v26.7.0, npm 12.0.2 (both as observed),
Python 3.12.10, ripgrep 15.2.0, Google Chrome (for headless render check, §3.5).

---

## 1. Integrity (Step 1)

All eight zips extract cleanly (read fully with zipfile, zero errors).

### 1.1 SHA-256 of all eight zips

| Zip (exact name; Phase zips use U+2014 EM DASH) | SHA-256 |
|---|---|
| `SWG Pre-CU Phase 0 — Server + Client.zip` | `4960481aa26b63aaf28eb1fac619f63ffa216b4b5c14efa5fd17d3e82a51a85f` |
| `SWG Pre-CU Phase 1 — Server + Client.zip` | `dfa95f8bf2c6112f614156f994167e7659489428c1c9429ed0256cb7dbd1bff2` |
| `SWG Pre-CU Phase 2 — Server + Client.zip` | `3238b36cd5008db4bfcae28468d7f424b2a5b71551df322efdc16d1bf6666e30` |
| `java_swg_files.zip` | `380a8d9680de1e84fa365283583e1f3679e61394c59cda3737b7c78109bdb1ac` |
| `swg-crafting-system-architecture-claude-opus-5-max.zip` | `414431f1adb6d8b88e6f405da872280d5afe1c5b5ea095674b6375cbdd4aa5a9` |
| `swg-crafting-system-architecture-llm-seed2-1-pro-preview.zip` | `66c6069aed960c5c714e257c84f05816a68e56ab480221e8ba18526083ed83f5` |
| `swg-phase0-3-claudecode-ready.zip` | `6b9e601d79ed9656517506ddca404b9066e165adc288ad1d2f81138aa36e5374` |
| `swg_track_a_b_code.zip` | `0aabdf5d15dc21e0fc7536ca1c781cdb16260b0de8b97ea885caaaa04e2c1799` |

### 1.2 Zip-entry counts (checklist expectations confirmed)

| Zip | Entries (total) | Dirs | Files |
|---|---|---|---|
| Phase 0 | 37 (expect 37 ✓) | 16 | 21 |
| Phase 1 | 43 (expect 43 ✓) | 18 | 25 |
| Phase 2 | 49 (expect 49 ✓) | 20 | 29 |
| phase0-3-claudecode | 62 (expect 62 ✓) | 23 | 39 |
| track-a-b | 23 (expect 24 — see §1.4) | 0 | 23 |
| java | 11 (expect 11 ✓) | 0 | 11 |
| seed demo | 22 (expect 22 ✓) | 5 | 17 |
| claude demo | 24 (expect 24 ✓) | 5 | 19 |

### 1.3 Byte comparison zip-vs-tree (per-file SHA)

- `code-phase-0/`: 21/21 files byte-identical. No delta.
- `code-phase-1/`: 25/25 byte-identical. No delta.
- `code-phase-2/`: 29/29 byte-identical. No delta.
- `code-phase0-3-claudecode/`: 37/39 byte-identical. The 2 in-zip-but-not-in-tree files are
  `swg-server/swg.db-shm` (32768 bytes) and `swg-server/swg.db-wal` (1289592 bytes) —
  runtime SQLite junk, absent from the tree exactly as the checklist/README state
  ("excluded at extraction — confirm absent": CONFIRMED ABSENT).
- `craft-seed2-1-pro-preview/`: 17/17 byte-identical. No delta.
- `craft-claude-opus-5-max/`: 19/19 byte-identical. No delta.
- `java-swg-skillcalc/`: 11/11 byte-identical. No delta — BUT layout contradicts its README:
  the zip/tree are flat (11 files at root, no `src/` directory), while the README's compile
  and IDE instructions assume a `src/` source root (§2.4, §3.4).
- `track-a-b-code/`: 23/23 byte-identical, plus ONE extra 0-byte file in the tree NOT in the
  zip: `track-a-b-code/fs.statSync(x).isFile()))`. Name pattern indicates an extraction/paste
  artifact, not project content. Pre-existing; left untouched (see §1.5). Checklist "24
  entries" = 23 zip files + this artifact.

### 1.4 Checklist expectation deltas (Step 1)

1. track-a-b "24 entries": observed 23 in zip + 1 junk artifact in tree. No project file missing.
2. phase0-3 "runtime swg.db-shm/swg.db-wal excluded": confirmed absent from tree.
3. java "11 files": confirmed; flat layout undisclosed by README (finding, not fix).

### 1.5 Repo-cleanliness check

`git status --short` after all work: no tracked file modified. Only pre-existing untracked
items: `investigations/evidence-reconciliation/` and
`sources/swg-pre-cu/track-a-b-code/fs.statSync(x).isFile()))`. Neither created by this pass.

---

## 2. Build verification (Step 2)

### 2.1 Toolchains (actual, recorded)

| Tree claim | Actual toolchain used |
|---|---|
| Go 1.22 per phase READMEs; phase0-3 `go.mod` "differs" | All four `go.mod` declare `go 1.22` (identical requires: uuid v1.6.0, websocket v1.5.3, go-sqlite3 v1.14.22, x/crypto v0.24.0). The phase0-3 difference is ONLY an added `replace golang.org/x/crypto => github.com/golang/crypto v0.24.0` with an explanatory comment (network-restricted-sandbox workaround). Container: `golang:1.22` = **go1.22.12 linux/amd64**. |
| Godot 4.2+ | winget-installed **Godot v4.2-stable** (`4.2.stable.official.46dc27791`), satisfies "4.2+". |
| Java 8+ | `eclipse-temurin:8-jdk-jammy` = **OpenJDK 1.8.0_502 (Temurin)**. |
| Node + React 19 / Vite 7 | Native **node v26.7.0, npm 12.0.2** (as observed); both `package.json` pin **Vite 7.3.2, React 19.2.6, TS 5.9.3**. Build script is `vite build` ONLY — `tsc` is never run (type errors would not fail the build). |

### 2.2 Go servers — `go build ./...` (in temp copies, docker go1.22.12)

| Tree | Raw `go build ./...` | Per-README path | Result |
|---|---|---|---|
| code-phase-0 | FAILS (no `go.sum` in zip): `missing go.sum entry for module providing package github.com/mattn/go-sqlite3 (imported by swg-server/internal/database); to add: go get swg-server/internal/database` (+ identical lines for uuid, websocket, x/crypto/bcrypt) | `go mod tidy` (exit 0) then `go build ./...` (exit 0, no diagnostics) | PASS via documented path |
| code-phase-1 | PASS exit 0 (has `go.sum`) | n/a | PASS |
| code-phase-2 | PASS exit 0 (has `go.sum`) | n/a | PASS |
| code-phase0-3 | PASS exit 0 (`replace` resolves via github.com/golang/crypto download) | n/a | PASS |

No compile errors in any tree once dependencies resolve. `go vet` not run (not in checklist).

### 2.3 Godot clients — import check: INCONCLUSIVE (verbatim evidence)

- `Godot_v4.2-stable_win64_console.exe --headless --path <temp copy> --import`:
  attempt 1 hung with no output (killed after 300 s); attempt 2 (`--verbose`) hung after
  engine-init lines only (`TextServer…`, `CORE API HASH`, `EDITOR API HASH`,
  `Loaded system CA certificates` — then silence; killed after 150 s). No project, script,
  or scene error was emitted before either hang. No `.godot/` dir was created.
- Standalone `--check-only --script World.gd` fails ONLY on out-of-project scope
  (`Could not find type "WSClient"`, `Identifier "GameState" not declared`) — an artifact of
  checking one script outside its project (cross-script `class_name`/autoload), not project evidence.
- Static (read-only) findings: every `project.godot` targets 4.2 + GL Compatibility with
  `GameState` autoload pre-configured (README's manual "add autoload" step is unnecessary);
  all three `.tscn` scenes SHIP in the zips (README's "create the scene files in the editor"
  is outdated); `assets/` dirs are empty in all four trees; phase-1/2/phase03 Godot trees are
  identical (phase-0 has older `WSClient.gd`/`World.gd`).

### 2.4 Java skill-calc — FAILS (verbatim error, not fixed)

- Literal README path fails first: `ls: cannot access 'src': No such file or directory`
  (flat zip/tree vs `src/`-based instructions).
- Adapted flat compile (`javac -source 8 -target 8 -d out *.java`, Temurin JDK 8) fails:
  `ProfessionTreePanel.java:185: error: local variables referenced from a lambda expression must be final or effectively final` (×2, for `col` and `row` — `for`-loop variables
  captured by the `.filter(b -> …)` lambda at lines 164–185). This is a language-level rule
  unchanged since Java 8, so the file as shipped compiles on NO javac version. Launch (`Main`)
  not attempted — no classes produced.

### 2.5 Crafting demos — both PASS

- `npm install` (temp copies): 93 packages each, exit 0. Warning (both): esbuild postinstall
  blocked (`1 package had install scripts blocked… esbuild@0.27.7`) — build unaffected.
- seed demo `vite build`: `✓ 40 modules transformed`, `dist/index.html 277.35 kB (gzip 83.66 kB)`, exit 0.
- claude demo `vite build`: `✓ 42 modules transformed`, `dist/index.html 330.53 kB (gzip 99.57 kB)`, exit 0.
- Deprecation notice (both, non-fatal): `` `module.register()` is deprecated. Use `module.registerHooks()` `` (Node 26 vs Vite 7.3.2).

---

## 3. Test verification (Step 3)

All server tests ran in-container (fresh `swg.db` per tree): server binary + each `cmd/*test`
in one shell session per CLAUDE.md's own procedure. Full transcripts retained in the temp
working area.

### 3.1 Go phase tests — ALL PASS (claimed-vs-observed)

| Tree | testclient (Phase 0) | phase1test (7) | phase2test (11) | phase3test (6) |
|---|---|---|---|---|
| code-phase-0 | PASS 9/9 HAM + WS world_enter | n/a (absent) | n/a | n/a |
| code-phase-1 | PASS | PASS 7/7 | n/a | n/a |
| code-phase-2 | PASS | PASS 7/7 | PASS 11/11 (OLD ladder: 66 SP/profession, Tier I = 2 SP / 1000 XP) | n/a |
| code-phase0-3 | PASS | PASS 7/7 | PASS 11/11 (v2 ladder: 42 SP/profession, Tier I = 1 SP / 1000 XP / 50 cr) | PASS 6/6 |

Observed specifics (phase0-3 bundle — "all 4 phases pass together" claim REPRODUCED):
testclient 9/9 species HAM + world_enter at `(3528.0, 5.0, -4804.0) on tatooine`; phase1
teleport rejection verbatim `speed exceeded: 1000.0 m in 23.138s (43.2 m/s, max 10.5 m/s)`;
phase3 Test 3 creature kill with combat XP 50; Test 5 REAL tick-driven incapacitation
(`health_current is 0`); Test 6 clone with `wounds_health=125, battle_fatigue=2.0%`.
Claimed-vs-observed: Phase 0 exit criteria ✓, "7 tests" ✓, "11 tests" ✓ (both ladders),
"all 4 phases pass together" ✓, CLAUDE.md "fully verified" ✓ for the Go trees (reproduced
end-to-end here, not taken on faith).

### 3.2 Track A/B — 10 stubs recorded as stubs; 5 reference functions smoke-run OK

Stubs (1-line `class X: pass`, scaffolding not implementation): `city_system.py`
(`CitySystem`), `combat_system.py`, `crafting_system.py`, `creature_spawning.py`,
`economy_ledger.py`, `faction_system.py`, `guild_system.py`, `mission_generator.py`,
`resource_generation.py`, `vendor_system.py`. (Also 1-line `database_schema.sql` comment and
2-line TDR/contract/backlog/topology `.md` placeholders.)
Reference functions (imported from temp copies, executed, outputs verbatim):
`resolve_attack(70,20,10,0,10,20,5,10,1.0,roll=50)` → `{'hit_chance': 80, 'roll': 50,
'hit': True, 'damage': 18.9}`; `roll=99` → miss, damage 0; `gen('metal','tatooine')` →
`{'oq': 505.49, 'radius_km': 17, …}`; `simulate(5,1000,100,40)` → `[1060,1120,1180,1240,1300]`;
`resource_stat(100,800,0.5,0.05)` → `525.0`; `OUTCOMES` 5-row distribution table (data only —
`crafting_experimentation.py` defines NO roll function); `xp_for_tier(1..5)` →
`[1000, 4594, 11211, 21112, 34493]`.

### 3.3 Java skill-calc launch — NOT RUN

Blocked by §2.4 compile failure (no classes). The README's "250-point pool behaves per rules"
claim is UNVERIFIED (and unverifiable as shipped).

### 3.4 Crafting demos — serve + render + SelfTest PASS; manual click-through NOT performed

- `vite preview`: both serve HTTP 200 with full single-file payloads (seed 277,351 bytes;
  claude 330,528 bytes).
- Headless-Chrome render of served pages: seed demo renders "SWG Pre-CU Crafting Core —
  Tech Demo" with live spawn browser (84/84 spawns, Tatooine/Naboo/… rows); claude demo
  renders the full memo + live loop.
- Claude demo built-in SelfTest: IN-PAGE result captured from rendered DOM — badge
  `6/6 holding` (emerald) with all six invariants PASS (determinism `fnv 5cc43053==5cc43053`;
  cap containment 2548 lanes/0 violations; lane fidelity 410 spawns/0 mismatches; schematic
  reachability 3 schematics; ceiling invariant 400 sessions/54 destroyed/0 breaches; taxonomy
  46 classes/acyclic). Independently reproduced engine-level (esbuild-bundle + node, App
  defaults seed `ANCHORHEAD`/horizon 240): identical 6/6 PASS. (`FAIL`/`failing`/`CRITICAL
  FAILURE` strings elsewhere in the bundle are application-logic template text, not observed
  failures.)
- Interactive click-through (assembly → experimentation → finalize button path) NOT performed:
  no browser-automation harness exists in this environment (no Playwright/puppeteer; the
  installed Chrome attaches to the live user session unless given an isolated profile, and
  scripted DOM interaction was out of reach). Recorded as not-done, not as pass.

---

## 4. IP / hygiene review (Step 4)

Method: ripgrep sweep for 40 Lucasfilm/SWG-specific terms across `sources/swg-pre-cu/`
(file counts + match counts per term; full table retained in temp working area) + binary-asset
scan. "Flag, do not clear": this section lists findings with file refs; it grants no clearance.

### 4.1 No-go confirmations

- **No Lucasfilm-owned binary assets:** zero `.png/.jpg/.gif/.wav/.mp3/.ogg/.fbx/.obj/.blend/.ttf/.mp4/.bin/.dat/.db*` anywhere under `sources/swg-pre-cu/`. All four Godot `assets/` dirs are EMPTY (dir entries only).
- **No character names:** Boba / Skywalker / Vader / Padawan / Alliance: 0 matches.

### 4.2 SWG-specific strings embedded in CODE (any reuse must strip/replace these)

- **Species ×9** (Bothan, Rodian, Trandoshan, Twi'lek, Wookiee, Zabrak, Mon Calamari,
  Sullustan + Human baseline): `internal/species/species.go` in ALL FOUR Go trees (3–5 hits
  each) + all server READMEs + live test output (HAM tables).
- **Tatooine / Mos Eisley / `tatooine`**: `internal/models/character.go` (spawn
  `3528.0, 5.0, -4804.0` + planet), `internal/database/db.go`, `skills_db.go` (6 trainer-seed
  hits each in phase-2/phase03), phase03 `handlers/combat.go`, `creatures/creatures.go`,
  all `cmd/testclient` + `phase3test` mains, Phase 2 READMEs.
- **Worrt / Womp Rat**: phase03 `internal/creatures/creatures.go` (2+2) + phase03 README (3+3).
  The README's own patch note admits Womp Rat is an "in-universe-consistent addition" (i.e.
  NOT from the cited GDD §5.2.1 — author's admission, recorded).
- **Dewback**: phase03 `creatures.go` (1). Rancor / Bantha / Krayt: GDD docs only.
- **HAM** (Health/Action/Mind + 6 attributes): 59 files / 955 matches — pervasive across all
  Go handlers/models/protocol/db, all Godot scripts/scenes, all READMEs. Mechanic vocabulary,
  but SWG-instantiated throughout.
- **Battle Fatigue / `battle_fatigue`**: phase03 combat core (`ham_state.go` 7,
  `combat_db.go`, `handlers/combat.go`, `protocol.go`, `phase3test`) + `db.go` schema in every
  phase; **"Force Wave"**: `java-swg-skillcalc/ProfessionFactory.java:246`
  ("Mastery of polearms. Unlocks Force Wave.").
- **Schematic / survey / harvester vocabulary**: `schematic` 28 files/403 matches —
  both demo engines (`craft.ts` 18, `CraftBench.tsx` 35, `crafting.ts` 17, `schematics.ts`),
  `skills/professions.go` (phase-2 + phase03), `ProfessionFactory.java` (2); `survey` in both
  demos + java ×15 + professions.go; `harvester` in seed `schematics.ts` (6). `crate`: 0
  matches anywhere (checklist's "crate" expectation NOT found — recorded).
- **Planet names as live data**: BOTH demo engines ship them (claude `engine/resources.ts`,
  seed `engine/resourceGen.ts` + `App.tsx`): Naboo, Corellia, Dantooine, Endor, Dathomir,
  Yavin, Rori, Talus, Lok; Kashyyyk in seed `resourceGen.ts` only; `ANCHORHEAD` is the claude
  demo's DEFAULT SEED (`App.tsx`) — SWG geography executes at runtime, not just in docs.
- **Jedi / Force / lightsaber / Sith / Rebel / Imperial / Empire**: confined to the four GDD
  copies (Jedi 8, Rebel 9, Imperial 8 per copy) — NOT in Go/Godot/demo runtime code, EXCEPT the
  java "Force Wave" hit above.

### 4.3 Classification statements (authors' claims, recorded; not clearance)

- "Fan-made, non-commercial, educational preservation project" in all three Group A
  top-level READMEs, all four extracted server READMEs, and (short form) `CLAUDE.md`.
- All four GDD copies carry a trademark acknowledgement (~line 9):
  "Star Wars, Star Wars Galaxies, and all associated planet, species, creature, and vehicle
  names are trademarks of Lucasfilm Ltd. … for hobbyist, educational, and non-commercial use."
  This is the authors' assertion; it is reproduced here as provenance, NOT as a reuse clearance.
- Hygiene verdict (finding, NOT a reuse decision — Step 7 not executed): the material is
  NOT clean as-is. SWG-specific identity strings are embedded in executable code paths
  (species tables, spawn coordinates, trainer seeds, creature templates, demo engines'
  default data), not merely in docs. Any future reuse would require a strip/replace pass;
  that pass needs separate human approval and is NOT authorized by this log.

---

## 5. Technology fit assessment (Step 5 — executed 2026-09-14)

Assessment ONLY. No numeric tuning performed, no mechanics designed, no EIC shaping, no
status changes. "Transfers" below means transfer of pattern/logic into original-IP testbed
work — never as-is SWG content — per HD-TST-01 (patterns-only fork, HUMAN-LOCKED; generic
pipeline exercises only). CRFT-006 stays PROPOSED, MFG-004 stays DERIVED CONSTRAINT,
all TBDs stay TBD.

### 5.0 Comparison target (TCIndustries needs, sourced)

From the checklist plus Master GDD v1.1.1 (status patch): a **persistent world** (incl.
world clock, environmental variation — GDD §world line 279); a **server-authoritative,
player-driven economy** grounded in the world (VIS-003, ECO-001 LOCKED; VIS-004/PIL-003
interdependence LOCKED); **skill-box profession architecture** (GDD §30 line 960 area);
**player-owned structures** (BLD-001 LOCKED); and a **future scale path** (AS-001-adjacent
provenance-at-scale stays ASSUMPTION/TBD; AS-002 data-driven-over-hard-coded is
ASSUMPTION, not LOCKED — cited where relevant, never as authority).

### 5.1 Stack fit per tree (observed properties)

| Tree | Stack (observed) | Fit vs needs |
|---|---|---|
| Go lineage `code-phase-0/1/2/` + `code-phase0-3-claudecode/` (Group A + C head) | Go 1.22 single-process monolith; one SQLite file (`DB_PATH`, default `swg.db`); HTTP+WS same `:8080`; 1 s in-process tick (`combat.go:312`); in-memory spatial grid; Godot 4.2 scripts-only client, empty `assets/` | Closest structural match to "server-authoritative persistent world": real authoritative server, ACID persistence, validated movement. Scale path UNDEMONSTRATED: single process, single file DB, all world state in-memory; Postgres portability is CLAIMED in READMEs ("replace `?`, drop PRAGMA") but no Postgres path exists in-tree. Godot client is demo-grade; engine choice (Godot vs UE5 destination per §5.3 demos) is a human decision, not assessed here. |
| `track-a-b-code/` | Dependency-free Python scripts + paper `.md`s; no server, client, persistence, or network | NOT a runtime stack — arithmetic references only. TDR-001/002 "Accepted" one-liners and `api_contracts.md` ("Movement, Combat, Crafting, Bazaar") are paper, not implementations. |
| `java-swg-skillcalc/` | Swing desktop app; no network, server, or persistence; does not compile as shipped (§2.4) | Fits nothing at runtime: single-desktop calculator. Rules logic only (see table). 3 professions vs Go's 6; its "37 pts to Master" costs match NEITHER Go ladder (v1: 66, v2: 42) — recorded inconsistency, no reconciliation attempted. |
| `craft-seed2-1-pro-preview/` | Browser-local Vite/React/TS; engine explicitly written for 1:1 UE5 C++ port (`App.tsx:4` "entire src/engine/ maps 1:1 to C++ USTRUCTs"; `types.ts:3` "map 1:1 to future C++ structs"; `crafting.ts:3` "Port verbatim to C++"; `index.ts:2` "becomes a C++ module") | Portable LOGIC, non-fitting RUNTIME: single-player, offline, no server/persistence/multiplayer/economy. Port-fidelity is the authors' design claim (recorded, not verified — no C++ exists). |
| `craft-claude-opus-5-max/` | Same browser-local shape; memo explicitly targets UE5 DataTables (`Architecture.tsx:108` "tuning tables cross the boundary untouched as UE5 DataTable rows"; `:339` "must survive the port to UE5 unchanged"); self-describes as "single-player · offline · pre-engine" (`App.tsx`) and "No engine. No scene graph. No actors" (`Architecture.tsx:14`) | Same verdict as seed demo: portable logic + harness discipline, no runtime fit. Its own memo concedes the runtime gap — recorded as the author's assessment, not mine. |

The four Go extractions are cumulative snapshots of ONE lineage (P0 ⊂ P1 ⊂ P2 ⊂ P03 +
GDD-v2 recomputation), not independent architectures; subsystem rows below carry a Trees
column (P0/P1/P2/P03) instead of duplicating identical rows four times.

### 5.2 Subsystem fit table (one row per major subsystem per tree)

Legend: ✅ transfers as pattern/logic · 🔧 needs rework before any testbed use · ❌ missing.

| # | Tree(s) | Subsystem (evidence ref) | Transfers as-is/pattern ✅ | Needs rework 🔧 | Missing ❌ |
|---|---|---|---|---|---|
| G1 | P0–P03 | Auth: register/login, bcrypt, bearer token (`server.go:65-72`; token `id:expiry` observed §3.1) | ✅ request/response + middleware pattern | 🔧 token scheme, SWG account model; no hardening evidence | ❌ anything beyond single-server sessions |
| G2 | P0–P03 | World entry + spatial grid (64 m cells, 128 m interest) + movement validation + visibility diff (`world/`, §3.1 Tests 1–7) | ✅ server-validated movement, interest-management, correction-message patterns | 🔧 single-zone, in-memory grid; SWG planet/coords | ❌ multi-zone, sharding, world clock, environment |
| G3 | P1–P03 | Spatial chat (/say 20 m, /shout 50 m) | ✅ radius-filtered relay pattern | 🔧 channel set is SWG-specific | ❌ persistent mail/cities/guilds (paper/stubs only, Track B) |
| G4 | P2, P03 | Skills/professions/XP/trainers; SP computed from owned boxes; prereq chains; cascade drop; typed XP pools (`skills/`, `skills_db.go`, §3.1) | ✅ skill-box/prereq/cascade/typed-pool LOGIC pattern (aligns with GDD skill-box direction) | 🔧 professions, XP types, costs, 250-pool are SWG content; `earn-xp` is a placeholder endpoint | ❌ elite professions, missions, interdependence mechanics |
| G5 | P03 only | Combat: HAM state machine, wounds/BF, posture/stance, skill-gated `combat_action`, lair spawning, retaliation AI, incap/revive/clone, HAM regen tick (`combat/`, `creatures/`, `combat_db.go`, §3.1) | ✅ gated-action, pool-state-machine, tick-driven-AI patterns | 🔧 unarmed-only, no pursuit, flat CL XP, ungated revive (all README-admitted); SWG creatures/content | ❌ full combat economy linkage, corpse/harvest (Phase 4+) |
| G6 | P0–P03 | Persistence: SQLite ACID, migrations, 5 s position flush (`database/`) | ✅ transactional request/response + DB-ownership split (per CLAUDE.md three-layer note) | 🔧 Postgres path CLAIMED but absent; SWG schema content | ❌ ledger economy live (TDR-002 is Track-B paper); scale story |
| G7 | P0–P03 | Godot client (`project.godot`, 6 scripts, .tscn present, assets EMPTY) | ✅ thin-client-over-authoritative-server SHAPE | 🔧 entire client is demo-grade login/world-display; SWG strings in scenes/scripts | ❌ any production client; engine decision open |
| T1 | Track A/B | `combat_resolution.resolve_attack` (12 lines, §3.2 output) | ✅ hit/damage arithmetic REFERENCE (de-SWG + mark non-canon) | 🔧 SWG-flavoured params | ❌ armor/stance/posture integration (lives in Go, SWG-bound) |
| T2 | Track A/B | `resource_spawn_generator.gen` (gauss stats, OQ, radius) | ✅ spawn-statistic SHAPE reference | 🔧 distribution params unvalidated; SWG planets | ❌ depletion/regeneration lifecycle, harvesting |
| T3 | Track A/B | `economy_inflation_sim.simulate` (linear faucet−sink) | ✅ placeholder ledger-toy ONLY | 🔧 linear model is trivially gameable — not a TCIndustries economy basis | ❌ ledger, vendors, trade, sinks/faucets (the actual ECO-001 need) |
| T4 | Track A/B | `crafting_experimentation.resource_stat` + OUTCOMES table (NO roll fn) | ✅ stat-improvement formula + outcome-table SHAPE | 🔧 probabilities unvalidated; no execution path | ❌ session/points/risk loop (lives in demos, SWG-bound) |
| T5 | Track A/B | `xp_progression_curve.xp_for_tier` (`1000·tier^2.2`) | ✅ curve SHAPE reference | 🔧 exponent/costs unvalidated vs EITHER Go ladder | ❌ progression economy linkage |
| T6 | Track A/B | 10 `class X: pass` stubs + one-line TDR/backlog/contract `.md`s | — (scaffolding, transfers nothing) | — | ❌ everything they name (cities, factions, vendors, missions…) |
| J1 | Java | Skill-box rules: 250 pool, prereq chains, surrender guards, refunds, master lock (`SkillManager`, `ProfessionFactory`, §3.3 UNVERIFIED at runtime) | ✅ rules LOGIC pattern (matches G4 shape; independent implementation) | 🔧 does not compile as shipped; 3-profession SWG content; costs match no Go ladder | ❌ server, persistence, anything multiplayer |
| S1 | Seed demo | Craft session FSM: init→assign→assembly→experiment→finalize (`crafting.ts:173-286`); pure, seeded RNG (`rng.ts`) | ✅ session-FSM + pure-function discipline (directly exercises HD-TST-01's "craft-with-experimentation" pipeline SHAPE) | 🔧 SWG schematics/stats/planets; skill model (`NOVICE/MASTER_SKILL`) | ❌ server binding, persistence, item ownership/economy |
| S2 | Seed demo | Resource taxonomy + spawn tables (`resourceTree.ts`, `resourceGen.ts`) | ✅ taxonomy-with-inheritance + seeded-spawn-table pattern | 🔧 SWG classes/planets/stats (OQ/PE/UT/… vocabulary) | ❌ lifecycle, depletion, regional economy |
| C1 | Claude demo | Craft FSM + risk/reward experimentation points + ceilings (`craft.ts:54-280`) | ✅ points-budget + ceiling-invariant pattern (strongest "experimentation" logic of all trees) | 🔧 SWG schematics/content; tuning unvalidated | ❌ same as S1 |
| C2 | Claude demo | Galaxy spawn timeline: deterministic seeded `generateGalaxy`, cap windows, lane fidelity (`resources.ts`, SelfTest §3.4 6/6) | ✅ determinism + cap-containment + invariant-HARNESS discipline (the SelfTest tripwire pattern itself transfers) | 🔧 SWG taxonomy/planets | ❌ server-side galaxy service, persistence, multiplayer visibility |
| C3 | Claude demo | Architecture memo (verdict/port/demarcation sections) | ✅ port-discipline notes (DataTables, "five structures") as PROSE input | 🔧 author's opinions, not evidence; UE5 destination assumed | — (memo, not system) |

### 5.3 What transfers / needs rework / is missing (summary — no new claims)

- **Transfers (patterns/logic only):** authoritative server shape + validated movement +
  interest management (G1–G3); skill-box/prereq/cascade/typed-XP logic (G4, J1);
  gated-action/HAM-state-machine/tick-AI patterns (G5); transactional persistence layering
  (G6); craft-session FSMs + experimentation points/ceilings (S1, C1); deterministic
  seeded spawn timelines + cap containment + invariant-harness discipline (S2, C2);
  reference arithmetic (T1–T5). Closest to HD-TST-01's authorised pipeline
  (survey→harvest→craft-with-experimentation→equip): S1/C1 session logic + C2 spawn
  determinism + G4 skill gating — as PATTERNS over generic IDs, never as SWG systems.
- **Needs rework (before ANY testbed use):** all SWG content in code paths (§4.2 —
  species, planets, creatures, schematics, stats vocabulary); single-process/file/​in-memory
  scaling bounds; Postgres path (claimed, absent); placeholder endpoints (`earn-xp`,
  ungated revive); both demo UI shells; Java compile failure; T3's trivially linear model.
- **Missing (all trees, vs §5.0 needs):** EIC; live ledger economy (vendors, trade,
  faucets/sinks); TCIndustries resources/professions/structures; provenance/reputation
  (deferred); world clock/environment/multi-zone/sharding — i.e. everything that makes a
  persistent world and a server-authoritative economy at scale. No tree demonstrates load,
  telemetry, or a scale path (Phase 10-class work unbuilt everywhere).

---

## 6. Coverage gap map (Step 6 — executed 2026-09-14)

Maps each tree's IMPLEMENTED systems against v1.1.1 §30 (Systems Requiring Deeper Design,
12 items). Grading bar (applied uniformly): **Covered** = verified executable implementation
(§§2–3) mapping directly onto the §30 item's shape. **Partial** = executable code/data
bearing on the item, however fragmentary, with material gaps stated. **Absent** = nothing
executable (stubs, paper `.md`s, docs, or schematic DATA entries alone do not count).
Grades measure distance toward a TCIndustries build (PROT-001: GDD over prototype) — they
promote NOTHING; all TBDs stay TBD, DEFERRED stays DEFERRED.

GDD v2 input scope (checklist direction): v2's correction is confined to predecessor-GDD
§6.3.1 (skill-point ladder 1/2/2/3 + totals) and §6.3.2 (XP/credit derivation formula),
both internally tagged [ASSUMPTION] — predecessor-project assumptions, NOT TCIndustries
authority. V2 therefore affects ONLY row 1's P03 cell below; all other cells are v1/v2-
invariant. (Observed correlate: P02 implements the v1 flat 2/3/4/5 pattern — 66 SP,
Tier I 2 SP — while P03 implements v2 — 42 SP, Tier I 1 SP; §3.1.)

Columns P0/P1/P2/P03 = cumulative Go lineage snapshots. TAB = `track-a-b-code/`.
JAVA = `java-swg-skillcalc/`. SEED = `craft-seed2-1-pro-preview/`. CLD = `craft-claude-opus-5-max/`.

| # | v1.1.1 §30 item (status) | P0 | P1 | P2 | P03 | TAB | JAVA | SEED | CLD |
|---|---|---|---|---|---|---|---|---|---|
| 1 | Skill trees, costs, specialisation budgets (TBD) | Absent | Absent | Partial — 6 professions × 18 boxes, prereqs, cascade drop, typed XP, v1 ladder, all verified §3.1 | Partial — same + v2 ladder (v2-scope note above) | Partial-fragment — `xp_for_tier` curve formula ONLY, no trees | Partial — 3 professions, prereq/refund/master-lock logic; RUNTIME UNVERIFIED, doesn't compile §2.4 | Absent — `SkillProfile` modifiers only, no trees | Absent — `CrafterSkill` 0–100 modifiers only |
| 2 | Resource spawn, attribute, lifecycle math (TBD) | Absent | Absent | Absent | Absent | Partial-fragment — 9-stat gauss + OQ + radius; NO attributes, lifecycle, depletion | Absent | Partial — spawn tables + 9-stat attributes + taxonomy; NO lifecycle/depletion | Partial — deterministic galaxy timeline + caps + lanes; NO depletion/lifecycle |
| 3 | Crafting experimentation + schematic structure (TBD) | Absent | Absent | Absent | Absent | Partial-fragment — stat formula + outcome TABLE, no roll fn, no session | Absent | Partial — full session FSM (verified build+render); SWG content, non-canon numerics | Partial — FSM + points/ceilings + SelfTest 6/6; same caveat |
| 4 | Manufacturing facility rules + automation constraints (TBD) | Absent | Absent | Absent | Absent | Absent (stub) | Absent | Absent (`mineralHarvester` is a schematic DATUM, not rules) | Absent |
| 5 | Transportation and logistics (TBD) | Absent | Absent | Absent | Absent | Absent | Absent | Absent (`travelBiscuit` datum only) | Absent |
| 6 | City formation, governance, maintenance (TBD) | Absent | Absent | Absent | Absent | Absent (`city_system.py` stub + 2-line backlog) | Absent | Absent | Absent |
| 7 | Combat PvE/PvP scope, risk, progression (TBD) | Absent | Absent | Absent | Partial — HAM resolution, creatures/lairs, incap/clone verified; SWG content; NO PvP, factions, risk model | Partial-fragment — 12-line `resolve_attack` arithmetic ONLY | Absent | Absent | Absent |
| 8 | Currency, money sinks, stabilisers (TBD) | Partial-fragment — credits balance field only | Partial-fragment — same | Partial-fragment — 5000 start + train-cost sink, verified; NO sinks/stabilisers design | Partial-fragment — same | Partial-fragment — linear faucet−sink TOY (`simulate`), trivially gameable, not an economy basis | Absent | Absent | Absent |
| 9 | Organisation governance and rights (TBD) | Absent | Absent | Absent | Absent | Absent (stubs: guild/faction) | Absent | Absent | Absent |
| 10 | Item durability, repair, decay (TBD) | Absent | Absent | Absent | Absent | Absent | Absent | Absent | Absent (a `hitpoints` ITEM STAT exists — not a durability system) |
| 11 | Species/ancestry/setting content (DEFERRED) | Partial — 9-species HAM table verified; DEFERRED flag stands, NOT a license | Partial — same | Partial — same | Partial — same | Absent | Absent | Absent | Absent |
| 12 | Full service profession definitions (TBD) | Absent | Absent | Partial-fragment — Medic/Entertainer 18-box TREES as data only; NO service mechanics (Phase 6 unbuilt, README-admitted) | Partial-fragment — same | Absent | Absent (3 combat/craft professions only) | Absent | Absent |

Reading: no §30 item is Covered by any tree. The densest precedent clusters are
skill-box mechanics (row 1: P2/P03/JAVA), resource-spawn fragments (row 2: TAB/SEED/CLD),
experimentation sessions (row 3: SEED/CLD), and SWG combat mechanisms (row 7: P03).
Rows 4, 5, 6, 9, 10 are Absent across ALL trees — no executable precedent exists for
manufacturing, transport, cities, organisations, or durability/decay. Row 8's fragments
(balance fields, train-cost sink, linear toy) do not approach ECO-001 needs. Row 11's
Partial is fenced by DEFERRED: executable SWG precedent that TCIndustries must NOT draw on
without lifting deferral — recorded, not cleared.

---

## 7. Reuse recommendation (Step 7 — PROPOSED, not a decision)

**Recommendation: REUSE PATTERNS ONLY** — lift verified patterns/logic (§5.2 ✅ column,
§6 Partial cells) into original-IP testbed work over generic IDs, subject to the §4
strip/replace pass and HD-TST-01 constraints. NOT as-is, NOT reference-only, NOT reject.

Evidence citations (all in this log): FOR patterns — Go suites pass live incl. the
4-phase bundle (§3.1); both demos build, serve, render, and the claude SelfTest holds 6/6
in-page and headless (§§2.5, 3.4); reference functions execute (§3.2); invariant-harness,
deterministic-timeline, and session-FSM discipline are directly reusable shapes (§5.2
S1/C1/C2). AGAINST as-is — SWG strings embedded in code paths (§4.2); Java doesn't compile
(§2.4); Godot import inconclusive (§2.3); single-process/file/in-memory scale bounds (§5.1);
placeholders and paper-stubs throughout (§§3.2, 5.2 T6). AGAINST reference-only — the
material runs as claimed (mostly), so it is more than read-only inspiration; its value is
executable precedent (passing tests, SelfTest tripwires), not prose. AGAINST reject —
rejection would discard verified, pattern-grade evidence contradicting §§2–3.

Standing HUMAN decision (recorded here per the checklist's "record the decision and date";
made by the owner, not by this log): **HD-TST-01, 2026-09-14 — patterns-only fork of the
SWG Pre-CU server codebase as mechanical testbed material**, with preconditions (steps 1–4
recorded in §§1–4 above; fork + de-SWG + review gate still pending) and standing boundaries
(no numeric tuning, no EIC shaping per HD-EIC-07, no provenance systems, no GDD edits, no
SWG names in the fork). This PROPOSED recommendation coincides with that ruling; where the
two could ever diverge, the human ruling governs.

---

## 8. What was NOT done (scope boundary)

All 7 checklist steps executed (§§1–4, 5, 6, 7). Standing exclusions observed throughout:
no EIC shaping, no numeric tuning, no provenance/reputation work, no GDD edits, no
code fixes (including the two recorded build failures), no fork/de-SWG step — that step
needs separate human approval. This log remains the SOLE new file from this pass.

## 9. Raw-evidence custody

Full command transcripts (Go builds/tests ×4, javac, npm/vite builds, preview/Chrome/SelfTest
captures, SHA table, zip namelists, byte-comparison, IP sweep table) retained in the
out-of-repo temp working area (`C:\temp\swg-evidence\`). Re-run procedures are documented
inline in §§1–4 above (not as repo scripts — no new repo files beyond this log).
