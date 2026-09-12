# SWG Pre-CU Sandbox Recreation — Phase 3

A ground-up, systems-faithful recreation of Star Wars Galaxies (Pre-CU era), built from the GDD specification.

> **Patch note:** Phases 0-2 were originally built against the pre-correction GDD, then patched against **GDD v2** (skill-point ladder, XP/credit formula). **Phase 3 (Combat) is new in this build**, implemented directly against v2, and independently verified end-to-end — all 4 phases pass together against the same running server, confirmed by actually building, running, and testing the code, not by re-reading it:
> 1. **Combat core** (`internal/combat`): HAM-driven attack resolution, postures/stances, wounds/battle-fatigue-adjusted effective max, incapacitation/revive/clone state machine (GDD 9.2, 9.6) — a direct port of the standalone reference implementation already verified earlier in this project.
> 2. **Creatures** (`internal/creatures`, `internal/handlers/combat.go`): two starter Tatooine creatures (Worrt — named in GDD 5.2.1; Womp Rat — in-universe-consistent addition), lair-based spawning, MVP retaliation AI, all persisted via new `lairs`/`creature_instances` tables.
> 3. **Found and fixed while building this**, not before: `GetCharacterByID` was scanning both `health_current` and `health_max` into the same Go field (silently discarding current health) and reading wounds/battle-fatigue/posture/stance into throwaway variables. Latent and untested since Phase 0-2 never damaged a character — Phase 3 is the first phase where current-vs-max actually matters. Fixed in `combat_db.go`.
> 4. **A real design gap, not just a bug**, surfaced during testing: the initial creature AI dropped a target based on a loose distance multiplier without checking range *before* attacking, and the two starter lairs were seeded 74m apart — far outside either creature's 8-10m aggro radius, so natural proximity aggro couldn't work from any single position. Both fixed: attack range is now checked consistently, and the starter lairs are clustered ~8m apart, a more sensible tutorial-area layout regardless of the bug it also happened to fix.
>
> Everything else below is prior-phase documentation, left as-is where still accurate.

**Phase 0** (complete): Server foundation, account/character creation, world entry with verified HAM values.

**Phase 1** (complete): Server-validated movement, spatial interest management, spatial chat with radius filtering, multi-player proximity sync.

**Phase 2** (complete): Skills and professions — XP earning, skill point spending at trainers, prerequisite chains, skill persistence across logout/login.

**Classification:** Fan-made, non-commercial, educational preservation project.

## Phase 2 Exit Criteria

> *"A player can earn XP from a placeholder action, spend skill points at a trainer NPC, and see the skill persist across logout/login."* — GDD Section 30, Phase 2

✅ 6 basic professions with full skill trees (Artisan, Brawler, Marksman, Scout, Medic, Entertainer)
✅ 108 skill boxes total (18 per profession: Novice + 4 trees × 4 tiers + Master)
✅ XP earning via placeholder action (typed XP pools: combat, crafting, etc.)
✅ Skill training at trainer NPCs (REST API, trainer registry in DB)
✅ Prerequisite validation (must train Novice first, must train tiers in sequence)
✅ Skill point system (250 max, computed from owned boxes — not mutable state)
✅ Credit costs for training (starting credits: 5000)
✅ Skill dropping with dependent cascade (dropping a box drops all dependents)
✅ XP not refunded on drop (GDD 7.4.1)
✅ Skills persist across logout/login (database-backed)
✅ Integration test (11 tests, all passing)

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
│   ├── phase1test/main.go  # Phase 1 integration test (movement, spatial, chat)
│   ├── phase2test/main.go  # Phase 2 integration test (skills, XP, training)
│   └── phase3test/main.go  # Phase 3 integration test (combat, incapacitation, clone)
├── internal/
│   ├── combat/
│   │   ├── combat.go        # Attack resolution: hit chance, damage, postures, stances (GDD 9.2.2-9.2.3)
│   │   └── ham_state.go     # HAM pool state machine: wounds, BF, incapacitation, revive, clone (GDD 9.2.1, 9.6)
│   ├── creatures/creatures.go # Creature templates, instances, lairs (GDD 17, 9.5)
│   ├── database/
│   │   ├── db.go           # Database layer + migrations + position persistence (GDD 29.4)
│   │   ├── skills_db.go    # Trainer seeding, skill box CRUD, XP operations, training transaction
│   │   └── combat_db.go    # Combat state, clone binding, creature/lair persistence (Phase 3)
│   ├── handlers/
│   │   ├── auth.go         # Register, login, character CRUD (GDD 27.3)
│   │   ├── world.go        # WebSocket world handler: spatial grid, visibility diff, chat (GDD 29.2)
│   │   ├── skills.go       # Skills API: professions, trainers, train/drop/earn-xp (GDD 7)
│   │   └── combat.go       # Combat actions, creature AI tick, HAM regen, world seeding (Phase 3)
│   ├── models/character.go # Character + appearance data models (GDD 6)
│   ├── protocol/protocol.go# Client-server message protocol, incl. Phase 3 combat messages (GDD 29.2.4)
│   ├── server/server.go    # HTTP server wiring
│   ├── skills/professions.go# 6 basic professions, 108 skill boxes, prereq/dependent logic (GDD 7, 8.2)
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

### Running the Phase 2 Integration Test

```bash
cd swg-server
go run ./cmd/phase2test/
```

This test verifies all Phase 2 exit criteria:
1. List all 6 basic professions with full skill tree data
2. List trainer NPCs seeded near Mos Eisley spawn
3. Verify character starts with 250 skill points, 5000 credits, no skills
4. Earn XP from a placeholder action (typed XP pool)
5. Train Novice Marksman (free — 0 skill points)
6. Train Ranged Accuracy I (costs 1 SP, 1000 XP, 50 credits — corrected v2 ladder)
7. Cannot train Tier III without Tier II (prerequisite enforcement)
8. Cannot train Artisan skills without Novice Artisan
9. Drop skill and verify skill points refunded
10. Skill persists across logout/login (database-backed)
11. Typed XP pools work correctly (combat + crafting simultaneously)

### Running the Phase 3 Integration Test

```bash
cd swg-server
go run ./cmd/phase3test/
```

This test verifies all Phase 3 exit criteria. **Takes 1-2 minutes to run** — Test 5 waits for genuine tick-driven creature retaliation to incapacitate the player in real time; nothing is mocked or fast-forwarded:
1. combat_action rejected with no combat skill trained (skill gate)
2. Unknown target rejected
3. Fight and defeat a spawned creature (real HAM-driven resolution, retried against RNG variance)
4. Posture change reflected in server state
5. Take sustained creature retaliation damage until incapacitated (real time, ~60-120s)
6. Clone at bound facility; verify Wounds and Battle Fatigue were applied (GDD 9.6.2)

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
| GET | `/api/professions` | No | List all professions with full skill tree data |
| GET | `/api/trainers` | No | List all trainer NPCs |
| GET | `/api/characters/{id}/skills` | Yes | Get character's owned skills, XP pools, skill points |
| POST | `/api/characters/{id}/train` | Yes | Train a skill box at a trainer |
| POST | `/api/characters/{id}/drop-skill` | Yes | Drop a skill box (cascades to dependents) |
| POST | `/api/characters/{id}/earn-xp` | Yes | Earn XP (placeholder action, Phase 2 only) |
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

## Phase 2: Skills & Professions

### Profession System (GDD Section 7, 8.2)

Six basic professions are available, each with 4 skill trees containing 4 tiers, plus Novice and Master boxes:

| Profession | XP Type | Skill Trees |
|---|---|---|
| Artisan | crafting | Engineering, Domestic Arts, Business, Harvesting |
| Brawler | combat | Unarmed, One-Handed, Two-Handed, Polearm |
| Marksman | combat | Ranged Accuracy, Ranged Support, Carbine, Rifle |
| Scout | scouting | Exploration, Trapping, Survival, Harvesting |
| Medic | medical | Organic Chemistry, First Aid, Diagnostics, Drugging |
| Entertainer | entertaining | Music, Dance, Image Design, Entertainment Healing |

### Skill Point System (GDD v2 Section 6.3.1, 7.2)

Every character has 250 skill points. Skill point costs (corrected Basic-profession ladder — an earlier version used a flat 2/3/4/5 pattern that summed to 66 points, which didn't reconcile against the GDD's own "~40 points to Master" target; fixed in v2):

| Box Type | SP Cost |
|---|---|
| Novice | 0 |
| Tier I | 1 |
| Tier II | 2 |
| Tier III | 2 |
| Tier IV | 3 |
| Master | 10 |

A full profession (all 18 boxes) costs 42 SP (≈40, matching GDD Section 8). Skill points are computed from owned boxes (250 − sum of costs), not stored as mutable state.

### XP System (GDD v2 Section 6.3.2, 7.3)

XP is stored in typed pools per character. Training a box deducts XP from the matching pool. Costs are *derived* from skill-point cost via `XP_cost = 1000 × skill_point_cost²`, `Credit_cost = round(XP_cost / 20, nearest 10)`, rather than hand-picked, so they can't drift out of sync with the ladder above:

| Tier | SP Cost | XP Cost | Credit Cost |
|---|---|---|---|
| Tier I | 1 | 1,000 | 50 |
| Tier II | 2 | 4,000 | 200 |
| Tier III | 2 | 4,000 | 200 |
| Tier IV | 3 | 9,000 | 450 |
| Master | 10 | 100,000 | 5,000 |

Tiers II and III share the same SP/XP/credit cost because they share the same skill-point cost under the corrected ladder — that's expected, not a bug. XP is NOT refunded when dropping skills (GDD 7.4.1).

### Training Rules (GDD 6.3.3, 7.4.1)

- Must train Novice before any tier boxes in a profession
- Must train boxes in sequence within a tree (I → II → III → IV)
- Must complete all 4 trees to unlock Master
- Dropping a skill box drops all dependent boxes (cascade)
- Training requires a trainer NPC matching the profession
- Training deducts credits (50–10,000 depending on tier)

### Trainers (GDD 7.3)

Six trainer NPCs are seeded near the Mos Eisley spawn point on Tatooine, one per basic profession. The trainer registry is stored in the `trainers` table.

## Phase 3: Combat, Creatures, Incapacitation & Clone

### Exit Criteria (GDD Section 30, Phase 3)

> *"A player can fight and defeat a spawned creature using a skill-gated ability, take damage against the correct HAM pool, and be incapacitated/revived-or-cloned correctly (Section 9.6)."*

✅ HAM-driven attack resolution (hit chance, damage, armor mitigation — GDD 9.2.2)
✅ Postures and stances affect combat (GDD 9.2.3)
✅ combat_action gated on owning Novice Marksman or Novice Brawler
✅ Two starter creatures (Womp Rat, Worrt) spawned from lairs near Mos Eisley
✅ MVP retaliation AI: creatures aggro on proximity, attack in range, disengage out of range
✅ Wounds and Battle Fatigue reduce effective max HAM (GDD 9.2.1)
✅ Incapacitation at 0 Health, 5-minute timer, auto-clone on expiry (GDD 9.6.1-9.6.2)
✅ Manual clone request while incapacitated
✅ Revive (any nearby player — Combat Medic gating deferred, see below)
✅ HAM regeneration out of combat (GDD 29.2.3)
✅ Combat XP reward on creature kill (GDD 7.2.1)
✅ Integration test (6 tests, all passing, including a real ~100s tick-driven incapacitation — not simulated or mocked)

### Combat Resolution (GDD 9.2.2)

```
Hit Chance = Base Weapon Accuracy + Attacker Accuracy Skill
             - Target Defense Skill - Target Posture Penalty (melee only)
             + Attacker Posture Bonus (ranged only)
Damage = (Weapon Min-Max Damage + Skill Damage Bonus) × Stance Modifier
         × (1 - Armor Mitigation %) × Resistance Multiplier
```

No crafted-weapon system exists yet (Phase 4), so every character currently fights unarmed (`internal/handlers/combat.go`'s `unarmedWeapon`).

### Creatures & Lairs (GDD 17, 9.5)

| Creature | AI Class | CL | Health | Damage | Aggro Radius |
|---|---|---|---|---|---|
| Womp Rat | Aggressive | 1 | 60 | 3-8 | 10m |
| Worrt | Scavenger | 2-3 | 140 | 6-15 | 8m |

Both lairs are seeded near Mos Eisley spawn, clustered within ~8m of each other. Lairs and creature instances persist in the `lairs` and `creature_instances` tables and reload on server restart (idempotent seeding — won't re-spawn if lairs already exist).

### Known Phase 3 Simplifications (scoped intentionally, not gaps found late)

- **Revive is not yet gated by Combat Medic.** GDD 9.6.1 specifies Combat Medic abilities; that profession doesn't exist until a later phase. Any nearby player can revive for now — add the skill-box gate in `handleRevive` once Combat Medic is built, the same way `handleCombatAction` already gates on Marksman/Brawler.
- **Creatures don't move.** Retaliation AI aggros and attacks in range but has no pursuit/patrol behavior. This is why lairs are clustered close together rather than spread across the map.
- **No corpse/harvest window.** GDD 17.2.2's 10-minute harvestable-corpse mechanic is Phase 4 (crafting/harvesting) scope; creatures currently despawn immediately on death.
- **Combat XP is a flat reward per creature CL tier**, not the full formula in GDD 7.2.1 (which needs a player "effective level" concept not yet built).

## Architecture Notes

- **Authoritative Server:** All game state (position, HAM, inventory) is computed server-side. Clients send action requests; the server resolves and pushes results. Movement is validated server-side — clients predict locally but must accept server corrections. (GDD 29.2.1)
- **Spatial Interest Management:** Only entities within 128m of a player are synced to that client. This limits bandwidth and processing to relevant entities. (GDD 29.2.2)
- **ACID Transactions:** Character creation and HAM state are written in a single database transaction to prevent partial writes. (GDD 29.5)
- **SQLite Dev / PostgreSQL Production:** The schema is designed to be portable to PostgreSQL. Migration requires changing the driver (`lib/pq` or `jackc/pgx`), the connection string, and replacing SQLite `?` placeholders with PostgreSQL `$1, $2, ...` in queries. PRAGMA directives must be removed. (GDD 29.4)

## Next Phases (GDD Section 30)

| Phase | Goal |
|---|---|
| Phase 1 | World movement + spatial chat + multi-player sync ✅ |
| Phase 2 | Skills & professions (XP, skill points, trainers) ✅ |
| Phase 3 | Combat core (HAM-driven resolution, creatures, lairs) ✅ |
| Phase 4 | Resources & crafting (surveying, harvesters, schematics) |
| Phase 5 | Economy infrastructure (credits, vendors, bazaar, housing) |
| Phase 6 | Entertainer & Medic (Battle Fatigue, Wounds, buffs) |
| Phase 7 | Civic systems (cities, guilds, mail) |
| Phase 8 | Faction & PvP |
| Phase 9 | Remaining elite professions & missions |
| Phase 10 | Balance, telemetry & polish |
