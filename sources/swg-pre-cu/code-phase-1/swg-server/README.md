# SWG Pre-CU Sandbox Recreation — Phase 1

A ground-up, systems-faithful recreation of Star Wars Galaxies (Pre-CU era), built from the GDD specification.

**Phase 0** (complete): Server foundation, account/character creation, world entry with verified HAM values.

**Phase 1** (complete): Server-validated movement, spatial interest management, spatial chat with radius filtering, multi-player proximity sync.

**Classification:** Fan-made, non-commercial, educational preservation project.

## Phase 1 Exit Criteria

> *"A player can walk around a persistent zone on one planet and see/chat with another connected player in real time."* — GDD Section 30, Phase 1

✅ Server-validated movement (7 m/s on-foot speed, teleport rejection)
✅ Spatial grid interest management (128m radius, cell-based partitioning)
✅ Proximity-based entity spawn/despawn (only sync nearby entities)
✅ Spatial chat with radius filtering (/say 20m, /shout 50m, /planet-wide)
✅ Position correction messages (anti-cheat)
✅ Position persistence (every 5s + on disconnect)
✅ Multi-player integration test (7 tests, all passing)
✅ Godot client with WASD movement + position correction handling

## Tech Stack

| Component | Technology |
|---|---|
| Game Server | Go 1.22 |
| Database | SQLite (dev) — schema designed to be portable to PostgreSQL |
| Real-time transport | WebSocket (gorilla/websocket) |
| Client | Godot 4 (GDScript) |
| Auth | Bearer token + bcrypt |

## Project Structure

```
swg-server/                 # Go game server
├── cmd/
│   ├── server/main.go      # Server entry point
│   ├── testclient/main.go  # Phase 0 integration test
│   └── phase1test/main.go  # Phase 1 integration test (movement, spatial, chat)
├── internal/
│   ├── database/db.go      # Database layer + migrations + position persistence (GDD 29.4)
│   ├── handlers/
│   │   ├── auth.go         # Register, login, character CRUD (GDD 27.3)
│   │   └── world.go        # WebSocket world handler: spatial grid, visibility diff, chat (GDD 29.2)
│   ├── models/character.go # Character + appearance data models (GDD 6)
│   ├── protocol/protocol.go# Client-server message protocol + position_correction (GDD 29.2.4)
│   ├── server/server.go    # HTTP server wiring
│   ├── species/species.go  # 9 species + HAM computation (GDD 6.2)
│   └── world/             # Spatial grid + movement validation (Phase 1)
│       ├── spatial.go     # Grid partitioning, entity tracking, Nearby() query (GDD 29.2.2)
│       └── movement.go    # Speed validation, position correction, chat channels (GDD 20.2.1, 18.2.1)
├── go.mod
└── swg.db                  # SQLite database (auto-created)

swg-godot/                  # Godot 4 client
├── project.godot           # Godot project file
├── scripts/
│   ├── SWGApiClient.gd     # HTTP API client
│   ├── WSClient.gd         # WebSocket client for real-time world
│   ├── Login.gd            # Login/register screen
│   ├── CharacterSelect.gd  # Character selection + creation
│   ├── World.gd            # 3D world scene + HAM display
│   └── GameState.gd        # Global state autoload
└── scenes/                 # Scene files (create in Godot editor)
```

## Setup

### Prerequisites

- **Go 1.22+** ([download](https://go.dev/dl/))
- **Godot 4.2+** ([download](https://godotengine.org/download/)) — for the client only

### Running the Server

```bash
cd swg-server
go mod tidy
go run ./cmd/server/
```

The server starts on `http://localhost:8080`. It auto-creates the SQLite database (`swg.db`) on first run.

### Running the Integration Test

```bash
cd swg-server
go run ./cmd/testclient/
```

This test:
1. Registers a new account
2. Logs in
3. Retrieves the species list (all 9)
4. Creates characters for all 9 species
5. Verifies HAM values match the GDD spec for each species
6. Lists all characters
7. Opens a WebSocket connection and enters the world
8. Verifies the world_enter response has correct character data + HAM

### Running the Phase 1 Integration Test

```bash
cd swg-server
go run ./cmd/phase1test/
```

This test verifies all Phase 1 exit criteria with two WebSocket clients:
1. Both players enter world at spawn and see each other (proximity entity spawn)
2. Movement is broadcast to nearby players
3. Spatial chat (/say) is delivered within 20m
4. Entity despawns when a player moves beyond 128m interest radius
5. Spatial chat is NOT delivered beyond 20m
6. Teleport (1000m jump) is rejected with position_correction
7. Entity respawns when a player moves back into range

### Running the Godot Client

1. Open Godot 4
2. Import the `swg-godot/` folder as a project
3. In Project Settings → Autoload, add `GameState.gd` as an autoload named `GameState`
4. Create the scene files (see below)
5. Press F5 to run

**Note:** The scene files (.tscn) need to be created in the Godot editor. The scripts are ready — you just need to set up the node hierarchy in each scene matching the `@onready` references in the scripts.

## API Endpoints

| Method | Path | Auth | Description |
|---|---|---|---|
| POST | `/api/register` | No | Create a new account |
| POST | `/api/login` | No | Login, returns session token |
| GET | `/api/species` | No | List all 9 playable species |
| POST | `/api/characters` | Yes | Create a character |
| GET | `/api/characters` | Yes | List account's characters |
| GET | `/api/characters/{id}` | Yes | Get character with HAM state |
| GET | `/ws` | Token | WebSocket world connection |
| GET | `/health` | No | Health check |

## HAM Values (GDD Section 6.2.2)

Baseline (Human) HAM values and species modifiers:

| Species | Health | Strength | Constitution | Action | Quickness | Stamina | Mind | Focus | Willpower |
|---|---|---|---|---|---|---|---|---|---|
| Human | 1000 | 500 | 500 | 1000 | 500 | 500 | 1000 | 500 | 500 |
| Bothan | 1000 | 490 | 500 | 1000 | 510 | 500 | 1000 | 500 | 500 |
| Rodian | 1000 | 500 | 490 | 1000 | 510 | 500 | 1000 | 500 | 500 |
| Trandoshan | 1000 | 520 | 500 | 1000 | 490 | 500 | 1000 | 500 | 500 |
| Twi'lek | 1000 | 500 | 500 | 1000 | 500 | 490 | 1000 | 500 | 510 |
| Wookiee | 1100 | 550 | 500 | 1000 | 450 | 500 | 1000 | 500 | 500 |
| Zabrak | 1000 | 500 | 500 | 1000 | 500 | 510 | 1000 | 500 | 490 |
| Mon Calamari | 1000 | 500 | 490 | 1000 | 500 | 500 | 1000 | 510 | 500 |
| Sullustan | 1000 | 500 | 500 | 1000 | 510 | 490 | 1000 | 500 | 500 |

## Phase 1: Spatial Interest Management

### Spatial Grid (GDD 29.2.2)

The world is partitioned per-planet into a grid of 64m cells. Each entity is tracked in the cell corresponding to its position. Proximity queries check only the cells within the query radius, then filter by exact 2D distance.

| Parameter | Value | GDD Reference |
|---|---|---|
| Cell size | 64m | 29.2.2 |
| Interest radius | 128m | 29.2.2 |
| /say radius | 20m | 18.2.1 |
| /shout radius | 50m | 18.2.1 |
| Planet bounds | ±8192m | — |

### Movement Validation (GDD 20.2.1, 29.2.1)

The server tracks each player's position and the timestamp of their last accepted move. When a move message arrives:

1. Compute the implied speed (distance / time delta)
2. If speed exceeds 7 m/s × 1.5 tolerance + 1m jitter allowance, reject the move
3. Send a `position_correction` message with the server-authoritative position
4. The client must snap to the corrected position

Movement messages with dt < 50ms are coalesced (ignored) to avoid false positives from jitter.

### Visibility Diff

Each client maintains a set of visible entity IDs. On every accepted move:
1. Query nearby entities within the 128m interest radius
2. Diff against the current visible set
3. Send `entity_spawn` for newly visible entities
4. Send `entity_despawn` for entities that left the radius
5. Symmetric: if A starts seeing B, B is also notified it can see A

### Position Persistence (GDD 29.4)

Positions are persisted to the database every 5 seconds while moving, and on disconnect. This prevents excessive database writes while ensuring positions survive server restarts.

## Architecture Notes

- **Authoritative Server:** All game state (position, HAM, inventory) is computed server-side. Clients send action requests; the server resolves and pushes results. Movement is validated server-side — clients predict locally but must accept server corrections. (GDD 29.2.1)
- **Spatial Interest Management:** Only entities within 128m of a player are synced to that client. This limits bandwidth and processing to relevant entities. (GDD 29.2.2)
- **ACID Transactions:** Character creation and HAM state are written in a single database transaction to prevent partial writes. (GDD 29.5)
- **SQLite Dev / PostgreSQL Production:** The schema is designed to be portable to PostgreSQL. Migration requires changing the driver (`lib/pq` or `jackc/pgx`), the connection string, and replacing SQLite `?` placeholders with PostgreSQL `$1, $2, ...` in queries. PRAGMA directives must be removed. (GDD 29.4)

## Next Phases (GDD Section 30)

| Phase | Goal |
|---|---|
| Phase 1 | World movement + spatial chat + multi-player sync ✅ |
| Phase 2 | Skills & professions (XP, skill points, trainers) |
| Phase 3 | Combat core (HAM-driven resolution, creatures, lairs) |
| Phase 4 | Resources & crafting (surveying, harvesters, schematics) |
| Phase 5 | Economy infrastructure (credits, vendors, bazaar, housing) |
| Phase 6 | Entertainer & Medic (Battle Fatigue, Wounds, buffs) |
| Phase 7 | Civic systems (cities, guilds, mail) |
| Phase 8 | Faction & PvP |
| Phase 9 | Remaining elite professions & missions |
| Phase 10 | Balance, telemetry & polish |
