# TCIndustries Technical Architecture & State Specification: Phase 8 (Faction & PvP) Preparation

- **Author:** Claude (Lead Agentic Coding Assistant, Antigravity/Google DeepMind Advanced Agentic Coding Team)
- **Version:** 1.0.0
- **Date:** October 24, 2026
- **Target Audience:** AI Engineering Agents (Claude / Custom Agents) preparing to design and implement Phase 8 (Faction & PvP) without repository access.
- **Project Base:** Tiwakings Craftworld Industries (TCIndustries) — Go 1.25 testbed game server with SQLite (pure-Go `modernc.org/sqlite`, no CGO), Gorilla WebSocket real-time synchronization, and standard Go HTTP API handlers.

---

## 1. Executive Summary & Design Governance

### 1.1 Project Context
TCIndustries is an original sandbox MMORPG design project focused on player-driven society, economy, identity, reputation, interdependence, and emergent gameplay. The testbed server fork implements mechanics phase-by-phase against the Master GDD. Phases 0–7 (Movement, Skills, Combat, Resources & Crafting, Economy, Social Support, and Civic Systems) are fully implemented, seeded, and verified end-to-end via automated integration suites.

### 1.2 Governance & Authority Model
Per project instructions (`AGENTS.md`), AI agents act as documentarians, stewards, and implementers. 
- **Critical Rule:** The Master GDD explicitly marks **Phase 8 (Faction & PvP scope, risk model, and progression integration)** as **TBD** (`OQ-008`). 
- **No Unsolicited Game Design:** This specification provides the exact technical architecture, database schemas, message protocols, and architectural extension points of the existing server so that when human-supplied design boundaries for Phase 8 are established, they can be implemented cleanly, idiomatically, and without breaking existing systems.

---

## 2. Existing System Architecture

### 2.1 Three-Layer Architectural Split
The server strictly follows the GDD Section 29 three-layer architecture:
1. **Real-Time Tick Loop:** Managed by `internal/handlers/combat.go` (`StartTickLoop`) and `internal/world` (spatial partitioning interest-management grid with 64m cells). Handles position updates, interpolation, proximity checks, and AI combat ticks.
2. **Transactional Request/Response API:** REST-style JSON endpoints under `/api/*` (auth, characters, skills, crafting, services, economy, civic) guarded by `AuthMiddleware` in `internal/handlers/auth.go`.
3. **Persistence Layer:** Owned entirely by `internal/database/*_db.go` modules using SQLite in WAL mode with pure-Go drivers. Schema migration is idempotent and executes at startup.

### 2.2 Core Package Layout (`tc-server/internal/`)
- `database/`: Owns all SQLite tables, connections (`*sql.DB`), raw queries, and transactions (`db.go`, `combat_db.go`, `resource_db.go`, `economy_db.go`, `civic_db.go`, `civic_social_db.go`, `service_db.go`).
- `protocol/`: Defines wire types (`protocol.go`), including client-to-server WS message constants (`MsgCombatAction`, `MsgMove`, etc.) and server-to-client receipts (`MsgCombatResult`, `MsgHAMUpdate`, etc.).
- `handlers/`: Orchestrates requests, validating business rules before calling database persistence or broadcasting state over WebSockets (`world.go`, `combat.go`, `economy.go`, `civic.go`, `civic_http2.go`).
- `world/`: Spatial grid math, bounding box checks, distance calculations (`movement.go`, `spatial.go`).
- `combat/`: Pure mathematical combat calculations (HAM pool damage rolls, hit chance, mitigation, regeneration).

---

## 3. Data Models & Persistence Reference (Pre-Phase 8 Baseline)

### 3.1 Database Connection & Row ID Generation
- Driver: `modernc.org/sqlite` (registered as `"sqlite"`).
- Connection string: `dbPath + "?_foreign_keys=on"`.
- Unique ID generation: Collision-proof millisecond-timestamp + sequence counters (`db.NewRowID(prefix)`).

### 3.2 Existing Core Tables
- `accounts`: `id` (TEXT PK), `username` (UNIQUE), `password_hash`, `created_at`.
- `characters`: `id` (TEXT PK), `account_id`, `name`, `species`, `planet`, `credits`, position (`pos_x`, `pos_y`, `pos_z`, `heading`), equip slots (`equipped_weapon_id`, `equipped_armor_id`, `clone_bind_*`).
- `ham_pool_states`: Current and max values for Health, Action, Mind, wounds, battle fatigue, posture, stance, incapacitation flag.
- `structures`: `id` (TEXT PK), `owner_character_id`, `zone`, `pos_x`, `pos_z`, `kind` (`house`, `vendor`, `city_hall`), `tier`, `maintenance_pool`, `upkeep_weekly`, `status`, `entry_perm`.
- `cities`: `id` (TEXT PK), `name`, `zone`, `center_x`, `center_z`, `radius_m`, `rank`, `treasury`, `mayor_character_id`, `status` (`forming`, `active`, `dissolved`), `zoning_open`.
- `guilds` & `guild_members`: Leadership roles (`leader`, `officer`, `member`), dues diversion, treasury.

---

## 4. Phase 8 Extension Points & Integration Hooks

When human-supplied design boundaries for Faction & PvP are authorized, the following extension points are already embedded in the architecture and ready for wiring:

### 4.1 Combat Gating (`internal/handlers/combat.go` & `protocol.go`)
- **Current State:** `CombatActionMsg` processes attacks against creatures or (if targeted) other players, resolving via server-seeded randomness (`combat.ResolveAttack`).
- **Faction/PvP Hook:** Player-vs-player combat currently lacks a faction alignment or PvP flagging gate. Phase 8 will introduce:
  - Faction membership (`neutral`, `rebels`, `imperials`, or custom project factions).
  - PvP flagging state (overt vs. covert / on-the-spot combat tracking).
  - Safety checks in `handleCombatAction`: rejecting attacks against allies or unflagged neutral players unless mutual consent / war declaration exists.

### 4.2 Spatial Interest Management (`internal/world/spatial.go`)
- **Current State:** `InterestRadius` (128.0m) broadcasts player and creature movements and spawns to nearby clients.
- **Faction/PvP Hook:** Overhead nameplates and radar indicators (`EntitySpawnMsg`) will need to carry faction affiliation and PvP combat state (`overt` indicator) so clients can render enemy vs. ally highlights.

### 4.3 Database Schema Migration (`internal/database/faction_db.go` — to be created)
- Following the established pattern of `EnsureEconomySchema()` and `EnsureCivicSchema()`, Phase 8 will introduce `EnsureFactionSchema()` to create:
  - `character_factions`: mapping `character_id` to faction ID and standing/points (PvP rating / faction rank).
  - `pvp_flags`: tracking combat timer, last aggressor, and overt/covert status.

### 4.4 Protocol Extensions (`internal/protocol/protocol.go`)
- New Client → Server messages:
  - `MsgTogglePvP` / `MsgSetFaction`: client requests to toggle overt status or declare allegiance.
- New Server → Client pushes:
  - `MsgFactionUpdate` / `MsgPvPStateChange`: updates faction standing or combat-flag countdown timers.

---

## 5. Development & Verification Standard

Any future implementation of Phase 8 must adhere strictly to the project's non-negotiable verification standard:
1. **Clean Compilation & Vet:** `go build ./...` and `go vet ./...` must return zero errors.
2. **End-to-End Integration Tests:** A dedicated test binary (`cmd/phase8test/main.go`) must be written and executed against a live server instance using `TESTBED_FAST_CYCLE=1` and a fresh SQLite database.
3. **Regression Safety:** All prior test suites (`phase1test`, `phase2test`, `phase3test`, `phase6test`, `phase7test`) must continue to pass without regression.
4. **No Speculative Design:** Do not invent lore, canon factions, or unapproved risk models. Implement strictly to project rulings and bounded parameters.