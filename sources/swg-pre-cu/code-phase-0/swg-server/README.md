# SWG Pre-CU Sandbox Recreation — Phase 0

A ground-up, systems-faithful recreation of Star Wars Galaxies (Pre-CU era), built from the GDD specification. This is Phase 0: server foundation, account/character creation, and world entry with verified HAM values.

**Classification:** Fan-made, non-commercial, educational preservation project.

## Phase 0 Exit Criteria

> *"A player can create an account, create a character with species/appearance, and see it standing in a bare world with correct starting HAM values."* — GDD Section 30, Phase 0

✅ All 9 species implemented with correct HAM modifiers
✅ Account registration and login (bcrypt password hashing)
✅ Character creation with species selection and appearance data
✅ WebSocket world entry with authoritative HAM state
✅ Entity spawn/despawn for other connected players
✅ Spatial chat relay

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
│   └── testclient/main.go  # Phase 0 integration test
├── internal/
│   ├── database/db.go      # Database layer + migrations (GDD 29.4)
│   ├── handlers/
│   │   ├── auth.go         # Register, login, character CRUD (GDD 27.3)
│   │   └── world.go        # WebSocket world handler (GDD 29.2)
│   ├── models/character.go # Character + appearance data models (GDD 6)
│   ├── protocol/protocol.go# Client-server message protocol (GDD 29.2.4)
│   ├── server/server.go    # HTTP server wiring
│   └── species/species.go  # 9 species + HAM computation (GDD 6.2)
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

## Architecture Notes

- **Authoritative Server:** All game state (position, HAM, inventory) is computed server-side. Clients send action requests; the server resolves and pushes results. (GDD 29.2.1)
- **ACID Transactions:** Character creation and HAM state are written in a single database transaction to prevent partial writes. (GDD 29.5)
- **SQLite Dev / PostgreSQL Production:** The schema is designed to be portable to PostgreSQL. Migration requires changing the driver (`lib/pq` or `jackc/pgx`), the connection string, and replacing SQLite `?` placeholders with PostgreSQL `$1, $2, ...` in queries. PRAGMA directives must be removed. (GDD 29.4)

## Next Phases (GDD Section 30)

| Phase | Goal |
|---|---|
| Phase 1 | World movement + spatial chat + multi-player sync |
| Phase 2 | Skills & professions (XP, skill points, trainers) |
| Phase 3 | Combat core (HAM-driven resolution, creatures, lairs) |
| Phase 4 | Resources & crafting (surveying, harvesters, schematics) |
| Phase 5 | Economy infrastructure (credits, vendors, bazaar, housing) |
| Phase 6 | Entertainer & Medic (Battle Fatigue, Wounds, buffs) |
| Phase 7 | Civic systems (cities, guilds, mail) |
| Phase 8 | Faction & PvP |
| Phase 9 | Remaining elite professions & missions |
| Phase 10 | Balance, telemetry & polish |
