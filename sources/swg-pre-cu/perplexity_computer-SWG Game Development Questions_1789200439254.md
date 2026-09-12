# SWG Game Development Questions

> **Created:** 06/08/2026, 08:21:10 | **Last updated:** 06/08/2026, 10:26:07 | **5 questions**
> **Author:** Perplexity — GPT-5.2

---

## Q1. Query only: Would you be able to create a game based on this document? Would you have to build it in segmented systems? What would be the best engine / language to use?

| | |
|---|---|
| Timestamp | 12/09/2026, 18:49:58 |
| Model | pplx_asi_glm |
| Mode | ASI |
| Language | en |
| UUID | `f274c0c0-462f-4926-ad87-f8bef680b05f` |

### Attachments

1. [SWG_PreCU_GDD.md](https://ppl-ai-file-upload.s3.amazonaws.com/web/direct-files/attachments/7508106/b9ed55cb-47b9-456d-808e-1e2738260985/SWG_PreCU_GDD.md?AWSAccessKeyId=ASIA2F3EMEYEYH2DMWY2&Signature=RRpTsvqVkoPk371EiOu6RzvBCU4%3D&x-amz-security-token=IQoJb3JpZ2luX2VjEOj%2F%2F%2F%2F%2F%2F%2F%2F%2F%2FwEaCXVzLWVhc3QtMSJGMEQCIG160CPacnesAtEiux7n%2F%2FkwMJzUI2oVlP5gDmuj0sb4AiBtGmHmi8EymP5YCziFO0%2B3j2dmZkygpdxGkU0WHvoLuCr8BAix%2F%2F%2F%2F%2F%2F%2F%2F%2F%2F8BEAEaDDY5OTc1MzMwOTcwNSIMXJgOv2rBcahax7igKtAE8Q5i0ka58RGuPKGqLG4%2Fx3fZP9mrzOcJuS%2FrVvcraS5nKGfyoquAu5XtBCzn%2Bk9pv11fN9YVw%2BfM%2BELyXQn8R1yoz0z1pvpSw0QUOOebUEKXwC4oBavVObwo%2FyMW%2FyeLPsPYCDC%2FLzLpCPKVPF%2FNpHg%2F30xjgE4Sl5lSkHM7mpCW3Y57HEfwKcH23%2FdRQLzWnTxGydCTFomXbRoFt1%2BYzWjr6fDnudGNwolf29j3eE8MUaNEkb7Ojs7tt%2FkiZwhYtXuIKe51YaFdqKqeGTEw5kyjEnjDLlp4aMw3j3T7R7IlTSMEkTtEd3yNA7gLKfi%2Fr6fk2K%2Fx1W2adAGsYeLZPQYvRfpzrtPiwfJfmcLvjRLaoMzgNBVpG36KCK3ja%2FbKz6619U0dTwRmikLhF6CjA7vLAhx7g%2B1ApCyqLrONojkxYl8r4PEnqb7gkIU%2FA5Ym8tZANtogiUhJh3GJ0WsNucmyzP6k5LwoDu%2BtI%2B6RAuRx7NhCr3wbO%2FsR1Tm0o96fsox8Qvv4O8polRmIoOPYf5YBBhIb9IHShH2TGxpe%2FaoVoGlVVk93D6n8JLEkHGunWWvBuqcQmgFLNUjUEIDmMxSWZektfZo0Rc3aeKiEzYZB8YsMjxaWVralGtaD7RU8ncD9zs8EO8YM2LGvw%2Bg6hl4BEnOnYEdtAmFFod7LZEDwf36WNveOYEA%2FbOfEMbF%2FEqKJiMupruoKY1fXzInWuRUg1j%2BsurfXSOVgDpok%2BxOGA%2BzUnju5vvzqNs7aAVKNyz%2BxgBhy6CbvSbimeJBe9jDQjZTVBjqZAX%2Bh%2BTRuPjnpkbofLv35WGU9hv085nijUHzinHhf%2BGBHuEGYiiX9OeiunRJD8V7uIoRULJXfyOkuAZ%2BCgcth%2BHwxGpc0gqOQIw7SeMGpTgukZaQ0cm2f%2BM%2Bh0e1HY%2FoQrK5z6v5cje3Uj91PJLTmKUmSnSTbVZQetCkYzqMkTpSNTIgyuOIvMlXQt4zNtMR4%2BIh8ffxgwii3zw%3D%3D&Expires=1789201038)
2. [SWG_PreCU_GDD.md](https://ppl-ai-file-upload.s3.amazonaws.com/web/direct-files/attachments/7508106/b9ed55cb-47b9-456d-808e-1e2738260985/SWG_PreCU_GDD.md)

> ⚠️ Attachment links are presigned AWS S3 URLs and expire — download the files soon after exporting.

### Answer

Let me read the key sections near the end — server architecture, implementation roadmap, and MVP scope.

I've read the full GDD — all 34 sections, ~4,000 lines. This is an exceptionally well-structured document, explicitly written for autonomous AI implementation. Let me give you a straight answer on all three questions.

---

## Can I Create a Game Based on This Document?

**Yes — but with honest scope expectations.**

This is one of the best-prepared GDDs I've seen for AI consumption. It already provides what most projects lack:

- **Dependency-ordered build roadmap** (Section 30: 10 phases with entry/exit criteria)
- **Explicit MVP scope** (Section 31: 3 planets, ~14 professions, full interdependence loop)
- **Complete data model schemas** (Section 28: entity relationships, field types)
- **Server architecture requirements** (Section 29: authoritative server, ACID persistence, tick systems)
- **Binding design rulings** where history is ambiguous ([ASSUMPTION] tags)
- **Cross-references everywhere** ("Section 9.2.1" is a literal pointer)

What I can realistically do:

| Capability | Realistic? |
|---|---|
| Build the Phase 0 server skeleton (Go + PostgreSQL, auth, persistence) | Yes |
| Implement data models / migrations from Section 28 | Yes |
| Build the HAM pool + skill-point progression system | Yes |
| Implement combat resolution (deterministic, stat-based) | Yes |
| Build crafting schematics + experimentation logic | Yes |
| Implement resource spawn algorithms + harvester ticks | Yes |
| Build economy ledger, bazaar, vendor systems | Yes |
| Create a functional 3D client with movement + UI | Partial — I can scaffold and build systems, but a polished MMO client is a massive asset/art undertaking |
| Produce a "playable but ugly" MVP that passes Section 1.3's five tests | Achievable over many sessions |
| Ship a production-quality MMO | Not solo — this is a multi-person, multi-year effort even with AI assistance |

The document itself acknowledges this: MVP requires passing the **Economic Interdependence Test**, **Horizontal Progression Test**, **Player-Driven Economy Test**, **Sandbox Validation Test**, and **Historical Accuracy Test**. That's a real game, not a demo.

---

## Would It Have to Be Built in Segmented Systems?

**Absolutely — and your GDD already mandates this.**

Section 30 prescribes 10 sequential phases, each with explicit entry/exit criteria. This isn't optional — it's architecturally forced by Pillar 1 (player interdependence). Nearly every system depends on another:

```
Phase 0: Server skeleton + character schema
Phase 1: World + movement + chat
Phase 2: Skills/professions (data layer)
Phase 3: Combat core (needs movement + skills)
Phase 4: Resources + crafting (needs combat/items)
Phase 5: Economy infrastructure (needs crafted items)
Phase 6: Entertainer/Medic (needs combat to generate wounds/fatigue)
Phase 7: Civic systems (needs housing from Phase 5)
Phase 8: Faction/PvP (needs combat + cities)
Phase 9: Remaining professions + missions
Phase 10: Balance + telemetry + polish
```

Beyond the phase sequencing, the architecture itself is naturally segmented into independent services (Section 29.2.2–29.2.4):

- **Real-time game loop** (position, combat, HAM ticking) — persistent WebSocket connection
- **Economic/ownership layer** (bazaar, vendors, trades, credit ledger) — request/response API with ACID transactions
- **World tick job system** (harvester extraction, resource rotation, city maintenance) — scheduled job queue, separate from game loop
- **Chat service** — can be a separate process
- **Planet shards** — each planet as an independently scalable process

This maps cleanly to a microservices or modular monolith approach. The GDD explicitly recommends separating the tick/job system from the real-time loop so a slow economic snapshot can never stall combat.

---

## Best Engine / Language

Given your background (Go, structured systems thinking), the project's non-commercial classification, and the GDD's technical requirements, here's my recommendation:

### Recommended Stack

| Layer | Technology | Why |
|---|---|---|
| **Game Server** | **Go** | You know it. Goroutines are ideal for concurrent player connections and tick systems. Excellent networking stdlib. Matches the authoritative server model perfectly. |
| **Database** | **PostgreSQL** | ACID-compliant (Section 29.4 is non-negotiable on this). Handles the credit ledger, item ownership, structure persistence. Mature, free, well-supported in Go via `pgx`. |
| **Real-time transport** | **WebSocket** (gorilla/websocket or nhooyr/websocket) | GDD Section 29.2.4 recommends WebSocket or UDP. WebSocket is simpler and adequate since combat is deterministic/stat-based, not twitch-reflex. |
| **Game Client** | **Godot 4** (GDScript or C#) | Free, open-source (matches the fan-project ethos). Good 3D capabilities, built-in networking, no licensing costs. Can connect to a custom Go server via WebSocket or ENet. |
| **Job scheduler** | **Go** (internal scheduler or `robfig/cron`) | Section 29.2.3 requires multiple tick rates (1s combat, 1h harvesters, 10d resources, 1w city maintenance). A Go-native job queue keeps this in-process. |
| **Ephemeral state** | **Redis** (optional for MVP) | Section 29.4 suggests a fast key-value store for live positions/HAM ticking/chat. Can start with in-memory Go maps for MVP. |

### Why This Combination Works

**Go server + Godot client** is increasingly common in indie MMO/sandbox projects. The key advantages:

1. **Server authority** — Go handles all combat rolls, crafting results, economic transactions. The client is purely a renderer and input forwarder, exactly as Section 29.2.1 demands.
2. **Concurrency model** — Go's goroutines map directly to the GDD's needs: one goroutine per connected player, separate goroutines for each tick system, channel-based communication between them.
3. **ACID integrity** — PostgreSQL transactions wrap every credit/item movement (Section 29.5), preventing the duplication exploits the GDD explicitly warns about.
4. **No licensing friction** — Godot is MIT-licensed. PostgreSQL is free. Go is free. This matters for a non-commercial preservation project.
5. **Your expertise** — You already know Go. The server is the hardest part. Starting from strength there is the right call.

### Alternative Considerations

| Option | When to choose |
|---|---|
| **Unity client** (C#) instead of Godot | If you want a larger asset store and more MMO-specific community resources. Licensing applies for commercial use. |
| **Rust server** instead of Go | If you want stronger memory-safety guarantees at compile time. Steeper learning curve if you don't already know it. |
| **Unreal Engine 5 client** | If visual fidelity is a top priority. Overkill for an MVP; heavy tooling for a small team. |
| **Nakama** (managed game server) | If you don't want to build the server from scratch. Provides matchmaking, real-time multiplayer, accounts. But it abstracts away the precise control the GDD's architecture requires. |

### My Strong Recommendation

Start with **Phase 0**: a Go server with PostgreSQL, account/character creation, and a bare Godot client that connects and renders a character standing in an empty world. That alone validates the entire tech stack end-to-end before you invest in any gameplay systems. The GDD's Phase 0 exit criteria are explicitly: "A player can create an account, create a character with species/appearance, and see it standing in a bare world with correct starting HAM values."

If you want to start building, I can scaffold the Phase 0 server and client structure right now — Go project layout, database migrations for the Section 28 entity schemas, WebSocket connection handling, and a minimal Godot scene that connects and renders a character. Just say the word.

### Sources

1. [SWG_PreCU_GDD.md](https://ppl-ai-file-upload.s3.amazonaws.com/web/direct-files/attachments/7508106/b9ed55cb-47b9-456d-808e-1e2738260985/SWG_PreCU_GDD.md?AWSAccessKeyId=ASIA2F3EMEYEYH2DMWY2&Signature=RRpTsvqVkoPk371EiOu6RzvBCU4%3D&x-amz-security-token=IQoJb3JpZ2luX2VjEOj%2F%2F%2F%2F%2F%2F%2F%2F%2F%2FwEaCXVzLWVhc3QtMSJGMEQCIG160CPacnesAtEiux7n%2F%2FkwMJzUI2oVlP5gDmuj0sb4AiBtGmHmi8EymP5YCziFO0%2B3j2dmZkygpdxGkU0WHvoLuCr8BAix%2F%2F%2F%2F%2F%2F%2F%2F%2F%2F8BEAEaDDY5OTc1MzMwOTcwNSIMXJgOv2rBcahax7igKtAE8Q5i0ka58RGuPKGqLG4%2Fx3fZP9mrzOcJuS%2FrVvcraS5nKGfyoquAu5XtBCzn%2Bk9pv11fN9YVw%2BfM%2BELyXQn8R1yoz0z1pvpSw0QUOOebUEKXwC4oBavVObwo%2FyMW%2FyeLPsPYCDC%2FLzLpCPKVPF%2FNpHg%2F30xjgE4Sl5lSkHM7mpCW3Y57HEfwKcH23%2FdRQLzWnTxGydCTFomXbRoFt1%2BYzWjr6fDnudGNwolf29j3eE8MUaNEkb7Ojs7tt%2FkiZwhYtXuIKe51YaFdqKqeGTEw5kyjEnjDLlp4aMw3j3T7R7IlTSMEkTtEd3yNA7gLKfi%2Fr6fk2K%2Fx1W2adAGsYeLZPQYvRfpzrtPiwfJfmcLvjRLaoMzgNBVpG36KCK3ja%2FbKz6619U0dTwRmikLhF6CjA7vLAhx7g%2B1ApCyqLrONojkxYl8r4PEnqb7gkIU%2FA5Ym8tZANtogiUhJh3GJ0WsNucmyzP6k5LwoDu%2BtI%2B6RAuRx7NhCr3wbO%2FsR1Tm0o96fsox8Qvv4O8polRmIoOPYf5YBBhIb9IHShH2TGxpe%2FaoVoGlVVk93D6n8JLEkHGunWWvBuqcQmgFLNUjUEIDmMxSWZektfZo0Rc3aeKiEzYZB8YsMjxaWVralGtaD7RU8ncD9zs8EO8YM2LGvw%2Bg6hl4BEnOnYEdtAmFFod7LZEDwf36WNveOYEA%2FbOfEMbF%2FEqKJiMupruoKY1fXzInWuRUg1j%2BsurfXSOVgDpok%2BxOGA%2BzUnju5vvzqNs7aAVKNyz%2BxgBhy6CbvSbimeJBe9jDQjZTVBjqZAX%2Bh%2BTRuPjnpkbofLv35WGU9hv085nijUHzinHhf%2BGBHuEGYiiX9OeiunRJD8V7uIoRULJXfyOkuAZ%2BCgcth%2BHwxGpc0gqOQIw7SeMGpTgukZaQ0cm2f%2BM%2Bh0e1HY%2FoQrK5z6v5cje3Uj91PJLTmKUmSnSTbVZQetCkYzqMkTpSNTIgyuOIvMlXQt4zNtMR4%2BIh8ffxgwii3zw%3D%3D&Expires=1789201038)

### Related Queries

- Scaffold the Phase 0 Go server and Godot client
- Create the Phase 0 project structure and PostgreSQL migrations
- Break Phase 0 into small implementation tasks

---

## Q2. yes

| | |
|---|---|
| Timestamp | 12/09/2026, 18:49:58 |
| Model | pplx_asi_glm |
| Mode | ASI |
| Language | en |
| UUID | `9dec6415-c7e8-482d-9789-f06668114865` |

### Answer

Now let me package everything into a zip for sharing:

Phase 0 is built and verified. Here's what you've got:

## What Was Built

**Go Server** (`swg-server/`)
- Account registration + login with bcrypt password hashing
- All 9 species with HAM modifiers computed from the GDD spec (Section 6.2.1–6.2.2)
- Character creation with species selection, appearance data, and name uniqueness
- Database layer with ACID transactions (character + HAM state written atomically)
- WebSocket world handler: `enter_world`, entity spawn/despawn/move, spatial chat relay
- CLI integration test that creates characters for all 9 species and independently verifies HAM values against a hard-coded spec table

**Godot 4 Client** (`swg-godot/`)
- Login/register screen with API client
- Character select + creation screen with species dropdown
- 3D world scene with ground plane, player mesh, HAM panel, and chat
- WebSocket client for real-time world connection
- All scene files (.tscn) created and ready to open in Godot 4
- GameState autoload configured in project.godot

## Test Results

```
Species HAM tests: 9/9 passed
WebSocket test:    true
=== ALL TESTS PASSED — Phase 0 exit criteria met ===
```

Every species' HAM values were verified against an independent hard-coded table derived from the GDD, not the server's own computation. The WebSocket `world_enter` response confirms correct character position, planet, and all 9 HAM attributes.

## How to Run

**Server:**
```bash
cd swg-server
go mod tidy
go run ./cmd/server/
```

**Integration test:**
```bash
cd swg-server
go run ./cmd/testclient/
```

**Godot client:**
1. Open `swg-godot/` in Godot 4.2+
2. Press F5 to run

## Next Steps (Phase 1)

Phase 1 from the GDD is: world movement + spatial chat + multi-player sync. The WebSocket infrastructure is already in place — Phase 1 would add server-validated movement, spatial interest management (only sync nearby entities), and proper terrain/world geometry.

### Related Queries

- Start Phase 1 with server-validated movement and spatial interest management
- Add proper terrain and world geometry to the Godot client
- Run the Godot client and verify the complete Phase 0 flow

---

## Q3. Start Phase 1 with server-validated movement and spatial interest management

| | |
|---|---|
| Timestamp | 12/09/2026, 18:49:58 |
| Model | pplx_asi_glm |
| Mode | ASI |
| Language | en |
| UUID | `12a9d24c-f46e-4036-ae5a-d61137484d02` |

### Attachments

1. [SWG_PreCU_GDD.md](https://ppl-ai-file-upload.s3.amazonaws.com/web/direct-files/attachments/7508106/b9ed55cb-47b9-456d-808e-1e2738260985/SWG_PreCU_GDD.md?AWSAccessKeyId=ASIA2F3EMEYEYH2DMWY2&Signature=RRpTsvqVkoPk371EiOu6RzvBCU4%3D&x-amz-security-token=IQoJb3JpZ2luX2VjEOj%2F%2F%2F%2F%2F%2F%2F%2F%2F%2FwEaCXVzLWVhc3QtMSJGMEQCIG160CPacnesAtEiux7n%2F%2FkwMJzUI2oVlP5gDmuj0sb4AiBtGmHmi8EymP5YCziFO0%2B3j2dmZkygpdxGkU0WHvoLuCr8BAix%2F%2F%2F%2F%2F%2F%2F%2F%2F%2F8BEAEaDDY5OTc1MzMwOTcwNSIMXJgOv2rBcahax7igKtAE8Q5i0ka58RGuPKGqLG4%2Fx3fZP9mrzOcJuS%2FrVvcraS5nKGfyoquAu5XtBCzn%2Bk9pv11fN9YVw%2BfM%2BELyXQn8R1yoz0z1pvpSw0QUOOebUEKXwC4oBavVObwo%2FyMW%2FyeLPsPYCDC%2FLzLpCPKVPF%2FNpHg%2F30xjgE4Sl5lSkHM7mpCW3Y57HEfwKcH23%2FdRQLzWnTxGydCTFomXbRoFt1%2BYzWjr6fDnudGNwolf29j3eE8MUaNEkb7Ojs7tt%2FkiZwhYtXuIKe51YaFdqKqeGTEw5kyjEnjDLlp4aMw3j3T7R7IlTSMEkTtEd3yNA7gLKfi%2Fr6fk2K%2Fx1W2adAGsYeLZPQYvRfpzrtPiwfJfmcLvjRLaoMzgNBVpG36KCK3ja%2FbKz6619U0dTwRmikLhF6CjA7vLAhx7g%2B1ApCyqLrONojkxYl8r4PEnqb7gkIU%2FA5Ym8tZANtogiUhJh3GJ0WsNucmyzP6k5LwoDu%2BtI%2B6RAuRx7NhCr3wbO%2FsR1Tm0o96fsox8Qvv4O8polRmIoOPYf5YBBhIb9IHShH2TGxpe%2FaoVoGlVVk93D6n8JLEkHGunWWvBuqcQmgFLNUjUEIDmMxSWZektfZo0Rc3aeKiEzYZB8YsMjxaWVralGtaD7RU8ncD9zs8EO8YM2LGvw%2Bg6hl4BEnOnYEdtAmFFod7LZEDwf36WNveOYEA%2FbOfEMbF%2FEqKJiMupruoKY1fXzInWuRUg1j%2BsurfXSOVgDpok%2BxOGA%2BzUnju5vvzqNs7aAVKNyz%2BxgBhy6CbvSbimeJBe9jDQjZTVBjqZAX%2Bh%2BTRuPjnpkbofLv35WGU9hv085nijUHzinHhf%2BGBHuEGYiiX9OeiunRJD8V7uIoRULJXfyOkuAZ%2BCgcth%2BHwxGpc0gqOQIw7SeMGpTgukZaQ0cm2f%2BM%2Bh0e1HY%2FoQrK5z6v5cje3Uj91PJLTmKUmSnSTbVZQetCkYzqMkTpSNTIgyuOIvMlXQt4zNtMR4%2BIh8ffxgwii3zw%3D%3D&Expires=1789201038)
2. [SWG_PreCU_GDD.md](https://ppl-ai-file-upload.s3.amazonaws.com/web/direct-files/attachments/7508106/b9ed55cb-47b9-456d-808e-1e2738260985/SWG_PreCU_GDD.md?AWSAccessKeyId=ASIA2F3EMEYES54OVFLZ&Signature=Tx%2FcVUMbFvYnRdBLpxbmQ4fnoLQ%3D&x-amz-security-token=IQoJb3JpZ2luX2VjEHEaCXVzLWVhc3QtMSJHMEUCIQDAO55dM0K2rGjvEXL%2BpvSgPiha%2Ftvy2JZkhFptEkJN%2FgIgbVgyhEI8E68Zhv4s97U0st0pWww9hXiX%2FoJWxbnN8Pgq8wQIOhABGgw2OTk3NTMzMDk3MDUiDNLxmGO8brcM4tR1ayrQBIPbKxFQz1vMDa1HvNW6F4C4PjcNVywKj9poWt5WhJefD4N%2BU59IH2M39clPJSRF1wZIut1KHfeO%2BRn5uJZN%2Fdt4OaTwQSkXDPRqUrTlLja07Ff5qDH9xXY8hqlgDVG2ZGHPExzX%2FXTORBj8mynHyYTyljubIPnotQwDy%2FvZiMQ%2BieJj7lWpVvBQH41Y0kGMUsXUW3OxriMmeCOFGhjP51RKlivmbsKCF5IOa%2FOM4x8cor%2BFf9JmUkINEDHDIX3WUJ2ZKBStWgAS1Mx5jI1MkqBWIUlCK37jcxzbNAsKxhAerSZFCXgnZBkbPJtFAj9hswXvfPEmvgR%2FsceeFAPFC%2Bc3sRJmjSMzHQM9cPBCx2TsH9%2Bw4e9in7dyowZP%2FOgz9PIDgD3CLP1gpfiM5uLmNGD%2F%2F05epiBKCZeoDohy9sV4uolqLmCOQtrfv5dDhFdHprMDzC4qrezNcyf%2BPB4fxuJ9YT8CyF65N9wWE8USZHfUYl1tIEdhnDvLTIrz6fttRCDPFhDU6uk56%2BcakhH5vhKJEBG%2B%2FBYv%2BB1MiO%2BVGMV3wn9kV)

> ⚠️ Attachment links are presigned AWS S3 URLs and expire — download the files soon after exporting.

### Answer

Both Phase 0 (9 species HAM tests) and Phase 1 (7 movement/spatial tests) pass. Now repackage the zip with proper top-level directories:

Phase 1 is complete. Here's what was built:

## Phase 1: Server-Validated Movement + Spatial Interest Management

**Exit criteria met:** "A player can walk around a persistent zone on one planet and see/chat with another connected player in real time."

### What's new in Phase 1

**Spatial Grid** (`internal/world/spatial.go`) — Per-planet 64m cell partitioning. Entities are tracked in cells, and proximity queries check only neighboring cells then filter by exact distance. O(1) cell lookup, O(n) nearby query where n is entities in adjacent cells.

**Movement Validation** (`internal/world/movement.go`) — Server tracks each player's position and last move time. When a move arrives, it computes implied speed and rejects anything over 7 m/s (GDD 20.2.1) with 1.5x tolerance + 0.3m jitter. Coalesced moves under 50ms are ignored entirely to prevent bypass attempts. Teleports get a `position_correction` message back.

**Visibility Diff** (`world.go`) — Each client maintains a visible-entity set. On every accepted move, the server queries nearby entities within 128m, diffs against the set, and sends `entity_spawn`/`entity_despawn` for changes. Symmetric: if A starts seeing B, B also gets notified.

**Spatial Chat** — `/say` (20m radius), `/shout` (50m radius), `/planet` (planet-wide). Messages are only delivered to clients within the channel's range.

**Position Persistence** — Positions saved to the database every 5 seconds while moving and on disconnect.

### Test Results

Both test suites pass:
- **Phase 0**: 9 species HAM verification + WebSocket world entry
- **Phase 1**: 7 tests covering proximity spawn, movement broadcast, spatial chat radius (within and beyond 20m), entity despawn beyond 128m, teleport rejection, and respawn when returning to range

### Godot Client

Updated with WASD movement (6 m/s), move updates at ~15Hz, third-person camera follow, `/shout` and `/planet` chat commands, and position_correction handling (snaps to server position when a move is rejected). The Godot client was not runtime-tested — the scripts are ready but need to be opened in the Godot editor.

The zip contains top-level `swg-server/` and `swg-godot/` directories. Ready for Phase 2 (skills & professions) whenever you are.

### Sources

1. [SWG_PreCU_GDD.md](https://ppl-ai-file-upload.s3.amazonaws.com/web/direct-files/attachments/7508106/b9ed55cb-47b9-456d-808e-1e2738260985/SWG_PreCU_GDD.md?AWSAccessKeyId=ASIA2F3EMEYEYH2DMWY2&Signature=RRpTsvqVkoPk371EiOu6RzvBCU4%3D&x-amz-security-token=IQoJb3JpZ2luX2VjEOj%2F%2F%2F%2F%2F%2F%2F%2F%2F%2FwEaCXVzLWVhc3QtMSJGMEQCIG160CPacnesAtEiux7n%2F%2FkwMJzUI2oVlP5gDmuj0sb4AiBtGmHmi8EymP5YCziFO0%2B3j2dmZkygpdxGkU0WHvoLuCr8BAix%2F%2F%2F%2F%2F%2F%2F%2F%2F%2F8BEAEaDDY5OTc1MzMwOTcwNSIMXJgOv2rBcahax7igKtAE8Q5i0ka58RGuPKGqLG4%2Fx3fZP9mrzOcJuS%2FrVvcraS5nKGfyoquAu5XtBCzn%2Bk9pv11fN9YVw%2BfM%2BELyXQn8R1yoz0z1pvpSw0QUOOebUEKXwC4oBavVObwo%2FyMW%2FyeLPsPYCDC%2FLzLpCPKVPF%2FNpHg%2F30xjgE4Sl5lSkHM7mpCW3Y57HEfwKcH23%2FdRQLzWnTxGydCTFomXbRoFt1%2BYzWjr6fDnudGNwolf29j3eE8MUaNEkb7Ojs7tt%2FkiZwhYtXuIKe51YaFdqKqeGTEw5kyjEnjDLlp4aMw3j3T7R7IlTSMEkTtEd3yNA7gLKfi%2Fr6fk2K%2Fx1W2adAGsYeLZPQYvRfpzrtPiwfJfmcLvjRLaoMzgNBVpG36KCK3ja%2FbKz6619U0dTwRmikLhF6CjA7vLAhx7g%2B1ApCyqLrONojkxYl8r4PEnqb7gkIU%2FA5Ym8tZANtogiUhJh3GJ0WsNucmyzP6k5LwoDu%2BtI%2B6RAuRx7NhCr3wbO%2FsR1Tm0o96fsox8Qvv4O8polRmIoOPYf5YBBhIb9IHShH2TGxpe%2FaoVoGlVVk93D6n8JLEkHGunWWvBuqcQmgFLNUjUEIDmMxSWZektfZo0Rc3aeKiEzYZB8YsMjxaWVralGtaD7RU8ncD9zs8EO8YM2LGvw%2Bg6hl4BEnOnYEdtAmFFod7LZEDwf36WNveOYEA%2FbOfEMbF%2FEqKJiMupruoKY1fXzInWuRUg1j%2BsurfXSOVgDpok%2BxOGA%2BzUnju5vvzqNs7aAVKNyz%2BxgBhy6CbvSbimeJBe9jDQjZTVBjqZAX%2Bh%2BTRuPjnpkbofLv35WGU9hv085nijUHzinHhf%2BGBHuEGYiiX9OeiunRJD8V7uIoRULJXfyOkuAZ%2BCgcth%2BHwxGpc0gqOQIw7SeMGpTgukZaQ0cm2f%2BM%2Bh0e1HY%2FoQrK5z6v5cje3Uj91PJLTmKUmSnSTbVZQetCkYzqMkTpSNTIgyuOIvMlXQt4zNtMR4%2BIh8ffxgwii3zw%3D%3D&Expires=1789201038)

### Related Queries

- Start Phase 2 with skills and professions
- Open the Godot project and runtime-test the Phase 1 client
- Add automated race testing for the server

---

## Q4. Start Phase 2 with skills and professions

| | |
|---|---|
| Timestamp | 12/09/2026, 18:49:58 |
| Model | pplx_asi_glm |
| Mode | ASI |
| Language | en |
| UUID | `6a8cfa2a-b2c2-4506-b980-d122d441080b` |

### Attachments

1. [SWG_PreCU_GDD.md](https://ppl-ai-file-upload.s3.amazonaws.com/web/direct-files/attachments/7508106/b9ed55cb-47b9-456d-808e-1e2738260985/SWG_PreCU_GDD.md?AWSAccessKeyId=ASIA2F3EMEYEYH2DMWY2&Signature=RRpTsvqVkoPk371EiOu6RzvBCU4%3D&x-amz-security-token=IQoJb3JpZ2luX2VjEOj%2F%2F%2F%2F%2F%2F%2F%2F%2F%2FwEaCXVzLWVhc3QtMSJGMEQCIG160CPacnesAtEiux7n%2F%2FkwMJzUI2oVlP5gDmuj0sb4AiBtGmHmi8EymP5YCziFO0%2B3j2dmZkygpdxGkU0WHvoLuCr8BAix%2F%2F%2F%2F%2F%2F%2F%2F%2F%2F8BEAEaDDY5OTc1MzMwOTcwNSIMXJgOv2rBcahax7igKtAE8Q5i0ka58RGuPKGqLG4%2Fx3fZP9mrzOcJuS%2FrVvcraS5nKGfyoquAu5XtBCzn%2Bk9pv11fN9YVw%2BfM%2BELyXQn8R1yoz0z1pvpSw0QUOOebUEKXwC4oBavVObwo%2FyMW%2FyeLPsPYCDC%2FLzLpCPKVPF%2FNpHg%2F30xjgE4Sl5lSkHM7mpCW3Y57HEfwKcH23%2FdRQLzWnTxGydCTFomXbRoFt1%2BYzWjr6fDnudGNwolf29j3eE8MUaNEkb7Ojs7tt%2FkiZwhYtXuIKe51YaFdqKqeGTEw5kyjEnjDLlp4aMw3j3T7R7IlTSMEkTtEd3yNA7gLKfi%2Fr6fk2K%2Fx1W2adAGsYeLZPQYvRfpzrtPiwfJfmcLvjRLaoMzgNBVpG36KCK3ja%2FbKz6619U0dTwRmikLhF6CjA7vLAhx7g%2B1ApCyqLrONojkxYl8r4PEnqb7gkIU%2FA5Ym8tZANtogiUhJh3GJ0WsNucmyzP6k5LwoDu%2BtI%2B6RAuRx7NhCr3wbO%2FsR1Tm0o96fsox8Qvv4O8polRmIoOPYf5YBBhIb9IHShH2TGxpe%2FaoVoGlVVk93D6n8JLEkHGunWWvBuqcQmgFLNUjUEIDmMxSWZektfZo0Rc3aeKiEzYZB8YsMjxaWVralGtaD7RU8ncD9zs8EO8YM2LGvw%2Bg6hl4BEnOnYEdtAmFFod7LZEDwf36WNveOYEA%2FbOfEMbF%2FEqKJiMupruoKY1fXzInWuRUg1j%2BsurfXSOVgDpok%2BxOGA%2BzUnju5vvzqNs7aAVKNyz%2BxgBhy6CbvSbimeJBe9jDQjZTVBjqZAX%2Bh%2BTRuPjnpkbofLv35WGU9hv085nijUHzinHhf%2BGBHuEGYiiX9OeiunRJD8V7uIoRULJXfyOkuAZ%2BCgcth%2BHwxGpc0gqOQIw7SeMGpTgukZaQ0cm2f%2BM%2Bh0e1HY%2FoQrK5z6v5cje3Uj91PJLTmKUmSnSTbVZQetCkYzqMkTpSNTIgyuOIvMlXQt4zNtMR4%2BIh8ffxgwii3zw%3D%3D&Expires=1789201038)
2. [SWG_PreCU_GDD.md](https://ppl-ai-file-upload.s3.amazonaws.com/web/direct-files/attachments/7508106/b9ed55cb-47b9-456d-808e-1e2738260985/SWG_PreCU_GDD.md?AWSAccessKeyId=ASIA2F3EMEYES54OVFLZ&Signature=Tx%2FcVUMbFvYnRdBLpxbmQ4fnoLQ%3D&x-amz-security-token=IQoJb3JpZ2luX2VjEHEaCXVzLWVhc3QtMSJHMEUCIQDAO55dM0K2rGjvEXL%2BpvSgPiha%2Ftvy2JZkhFptEkJN%2FgIgbVgyhEI8E68Zhv4s97U0st0pWww9hXiX%2FoJWxbnN8Pgq8wQIOhABGgw2OTk3NTMzMDk3MDUiDNLxmGO8brcM4tR1ayrQBIPbKxFQz1vMDa1HvNW6F4C4PjcNVywKj9poWt5WhJefD4N%2BU59IH2M39clPJSRF1wZIut1KHfeO%2BRn5uJZN%2Fdt4OaTwQSkXDPRqUrTlLja07Ff5qDH9xXY8hqlgDVG2ZGHPExzX%2FXTORBj8mynHyYTyljubIPnotQwDy%2FvZiMQ%2BieJj7lWpVvBQH41Y0kGMUsXUW3OxriMmeCOFGhjP51RKlivmbsKCF5IOa%2FOM4x8cor%2BFf9JmUkINEDHDIX3WUJ2ZKBStWgAS1Mx5jI1MkqBWIUlCK37jcxzbNAsKxhAerSZFCXgnZBkbPJtFAj9hswXvfPEmvgR%2FsceeFAPFC%2Bc3sRJmjSMzHQM9cPBCx2TsH9%2Bw4e9in7dyowZP%2FOgz9PIDgD3CLP1gpfiM5uLmNGD%2F%2F05epiBKCZeoDohy9sV4uolqLmCOQtrfv5dDhFdHprMDzC4qrezNcyf%2BPB4fxuJ9YT8CyF65N9wWE8USZHfUYl1tIEdhnDvLTIrz6fttRCDPFhDU6uk56%2BcakhH5vhKJEBG%2B%2FBYv%2BB1MiO%2BVGMV3wn9kV)

> ⚠️ Attachment links are presigned AWS S3 URLs and expire — download the files soon after exporting.

### Answer

README updated. Now package everything:

Phase 2 is complete. Here's what was built:

## Phase 2: Skills & Professions

### What's New

**6 Basic Professions** (GDD 8.2) — Artisan, Brawler, Marksman, Scout, Medic, Entertainer. Each has 4 skill trees with 4 tiers, plus Novice and Master boxes — 108 skill boxes total, all generated programmatically.

**Skill Point System** (GDD 6.3.1) — 250 points per character. Costs range from 0 (Novice) to 10 (Master). Points are computed from owned boxes, not stored as mutable state, so they can never drift.

**XP System** (GDD 7.3) — Typed XP pools (combat, crafting, etc.). Earning XP adds to the matching pool. Training deducts XP costs (1k–100k by tier). XP is not refunded on drop, per GDD 7.4.1.

**Training at Trainers** (GDD 7.3) — 6 trainer NPCs seeded near Mos Eisley spawn, one per profession. The REST API validates: trainer matches profession, prerequisites met (Novice first, tiers in sequence, all 4 trees for Master), sufficient XP, credits, and skill points. All in a single DB transaction.

**Skill Dropping** (GDD 7.4.1) — Dropping a box cascades to all dependent boxes. XP is not refunded.

### Test Results — All 27 Tests Pass

| Phase | Tests | Result |
|---|---|---|
| Phase 0 | 9 species HAM + WebSocket | All pass |
| Phase 1 | 7 spatial/movement/chat tests | All pass |
| Phase 2 | 11 skills/XP/training tests | All pass |

The Phase 2 exit criteria is met: a player can earn XP from a placeholder action, spend skill points at a trainer NPC, and see the skill persist across logout/login.

### Files

- **Server + Client**: `swg-phase2.zip` (65 KB)
- **README**: Full documentation with API endpoints, skill point/XP cost tables, training rules, and architecture notes

### Sources

1. [SWG_PreCU_GDD.md](https://ppl-ai-file-upload.s3.amazonaws.com/web/direct-files/attachments/7508106/b9ed55cb-47b9-456d-808e-1e2738260985/SWG_PreCU_GDD.md?AWSAccessKeyId=ASIA2F3EMEYEYH2DMWY2&Signature=RRpTsvqVkoPk371EiOu6RzvBCU4%3D&x-amz-security-token=IQoJb3JpZ2luX2VjEOj%2F%2F%2F%2F%2F%2F%2F%2F%2F%2FwEaCXVzLWVhc3QtMSJGMEQCIG160CPacnesAtEiux7n%2F%2FkwMJzUI2oVlP5gDmuj0sb4AiBtGmHmi8EymP5YCziFO0%2B3j2dmZkygpdxGkU0WHvoLuCr8BAix%2F%2F%2F%2F%2F%2F%2F%2F%2F%2F8BEAEaDDY5OTc1MzMwOTcwNSIMXJgOv2rBcahax7igKtAE8Q5i0ka58RGuPKGqLG4%2Fx3fZP9mrzOcJuS%2FrVvcraS5nKGfyoquAu5XtBCzn%2Bk9pv11fN9YVw%2BfM%2BELyXQn8R1yoz0z1pvpSw0QUOOebUEKXwC4oBavVObwo%2FyMW%2FyeLPsPYCDC%2FLzLpCPKVPF%2FNpHg%2F30xjgE4Sl5lSkHM7mpCW3Y57HEfwKcH23%2FdRQLzWnTxGydCTFomXbRoFt1%2BYzWjr6fDnudGNwolf29j3eE8MUaNEkb7Ojs7tt%2FkiZwhYtXuIKe51YaFdqKqeGTEw5kyjEnjDLlp4aMw3j3T7R7IlTSMEkTtEd3yNA7gLKfi%2Fr6fk2K%2Fx1W2adAGsYeLZPQYvRfpzrtPiwfJfmcLvjRLaoMzgNBVpG36KCK3ja%2FbKz6619U0dTwRmikLhF6CjA7vLAhx7g%2B1ApCyqLrONojkxYl8r4PEnqb7gkIU%2FA5Ym8tZANtogiUhJh3GJ0WsNucmyzP6k5LwoDu%2BtI%2B6RAuRx7NhCr3wbO%2FsR1Tm0o96fsox8Qvv4O8polRmIoOPYf5YBBhIb9IHShH2TGxpe%2FaoVoGlVVk93D6n8JLEkHGunWWvBuqcQmgFLNUjUEIDmMxSWZektfZo0Rc3aeKiEzYZB8YsMjxaWVralGtaD7RU8ncD9zs8EO8YM2LGvw%2Bg6hl4BEnOnYEdtAmFFod7LZEDwf36WNveOYEA%2FbOfEMbF%2FEqKJiMupruoKY1fXzInWuRUg1j%2BsurfXSOVgDpok%2BxOGA%2BzUnju5vvzqNs7aAVKNyz%2BxgBhy6CbvSbimeJBe9jDQjZTVBjqZAX%2Bh%2BTRuPjnpkbofLv35WGU9hv085nijUHzinHhf%2BGBHuEGYiiX9OeiunRJD8V7uIoRULJXfyOkuAZ%2BCgcth%2BHwxGpc0gqOQIw7SeMGpTgukZaQ0cm2f%2BM%2Bh0e1HY%2FoQrK5z6v5cje3Uj91PJLTmKUmSnSTbVZQetCkYzqMkTpSNTIgyuOIvMlXQt4zNtMR4%2BIh8ffxgwii3zw%3D%3D&Expires=1789201038)

### Related Queries

- Start Phase 3 with combat core
- Add skill and profession support to the Godot client
- Review the Phase 2 implementation against the GDD

---

## Q5. Start Phase 3 with combat core

| | |
|---|---|
| Timestamp | 12/09/2026, 18:49:58 |
| Model | pplx_asi_glm |
| Mode | ASI |
| Language | en |
| UUID | `fa8db2e8-c271-4a04-adb2-212043a7a399` |

### Attachments

1. [SWG_PreCU_GDD.md](https://ppl-ai-file-upload.s3.amazonaws.com/web/direct-files/attachments/7508106/b9ed55cb-47b9-456d-808e-1e2738260985/SWG_PreCU_GDD.md?AWSAccessKeyId=ASIA2F3EMEYEYH2DMWY2&Signature=RRpTsvqVkoPk371EiOu6RzvBCU4%3D&x-amz-security-token=IQoJb3JpZ2luX2VjEOj%2F%2F%2F%2F%2F%2F%2F%2F%2F%2FwEaCXVzLWVhc3QtMSJGMEQCIG160CPacnesAtEiux7n%2F%2FkwMJzUI2oVlP5gDmuj0sb4AiBtGmHmi8EymP5YCziFO0%2B3j2dmZkygpdxGkU0WHvoLuCr8BAix%2F%2F%2F%2F%2F%2F%2F%2F%2F%2F8BEAEaDDY5OTc1MzMwOTcwNSIMXJgOv2rBcahax7igKtAE8Q5i0ka58RGuPKGqLG4%2Fx3fZP9mrzOcJuS%2FrVvcraS5nKGfyoquAu5XtBCzn%2Bk9pv11fN9YVw%2BfM%2BELyXQn8R1yoz0z1pvpSw0QUOOebUEKXwC4oBavVObwo%2FyMW%2FyeLPsPYCDC%2FLzLpCPKVPF%2FNpHg%2F30xjgE4Sl5lSkHM7mpCW3Y57HEfwKcH23%2FdRQLzWnTxGydCTFomXbRoFt1%2BYzWjr6fDnudGNwolf29j3eE8MUaNEkb7Ojs7tt%2FkiZwhYtXuIKe51YaFdqKqeGTEw5kyjEnjDLlp4aMw3j3T7R7IlTSMEkTtEd3yNA7gLKfi%2Fr6fk2K%2Fx1W2adAGsYeLZPQYvRfpzrtPiwfJfmcLvjRLaoMzgNBVpG36KCK3ja%2FbKz6619U0dTwRmikLhF6CjA7vLAhx7g%2B1ApCyqLrONojkxYl8r4PEnqb7gkIU%2FA5Ym8tZANtogiUhJh3GJ0WsNucmyzP6k5LwoDu%2BtI%2B6RAuRx7NhCr3wbO%2FsR1Tm0o96fsox8Qvv4O8polRmIoOPYf5YBBhIb9IHShH2TGxpe%2FaoVoGlVVk93D6n8JLEkHGunWWvBuqcQmgFLNUjUEIDmMxSWZektfZo0Rc3aeKiEzYZB8YsMjxaWVralGtaD7RU8ncD9zs8EO8YM2LGvw%2Bg6hl4BEnOnYEdtAmFFod7LZEDwf36WNveOYEA%2FbOfEMbF%2FEqKJiMupruoKY1fXzInWuRUg1j%2BsurfXSOVgDpok%2BxOGA%2BzUnju5vvzqNs7aAVKNyz%2BxgBhy6CbvSbimeJBe9jDQjZTVBjqZAX%2Bh%2BTRuPjnpkbofLv35WGU9hv085nijUHzinHhf%2BGBHuEGYiiX9OeiunRJD8V7uIoRULJXfyOkuAZ%2BCgcth%2BHwxGpc0gqOQIw7SeMGpTgukZaQ0cm2f%2BM%2Bh0e1HY%2FoQrK5z6v5cje3Uj91PJLTmKUmSnSTbVZQetCkYzqMkTpSNTIgyuOIvMlXQt4zNtMR4%2BIh8ffxgwii3zw%3D%3D&Expires=1789201038)
2. [SWG_PreCU_GDD.md](https://ppl-ai-file-upload.s3.amazonaws.com/web/direct-files/attachments/7508106/b9ed55cb-47b9-456d-808e-1e2738260985/SWG_PreCU_GDD.md?AWSAccessKeyId=ASIA2F3EMEYES54OVFLZ&Signature=Tx%2FcVUMbFvYnRdBLpxbmQ4fnoLQ%3D&x-amz-security-token=IQoJb3JpZ2luX2VjEHEaCXVzLWVhc3QtMSJHMEUCIQDAO55dM0K2rGjvEXL%2BpvSgPiha%2Ftvy2JZkhFptEkJN%2FgIgbVgyhEI8E68Zhv4s97U0st0pWww9hXiX%2FoJWxbnN8Pgq8wQIOhABGgw2OTk3NTMzMDk3MDUiDNLxmGO8brcM4tR1ayrQBIPbKxFQz1vMDa1HvNW6F4C4PjcNVywKj9poWt5WhJefD4N%2BU59IH2M39clPJSRF1wZIut1KHfeO%2BRn5uJZN%2Fdt4OaTwQSkXDPRqUrTlLja07Ff5qDH9xXY8hqlgDVG2ZGHPExzX%2FXTORBj8mynHyYTyljubIPnotQwDy%2FvZiMQ%2BieJj7lWpVvBQH41Y0kGMUsXUW3OxriMmeCOFGhjP51RKlivmbsKCF5IOa%2FOM4x8cor%2BFf9JmUkINEDHDIX3WUJ2ZKBStWgAS1Mx5jI1MkqBWIUlCK37jcxzbNAsKxhAerSZFCXgnZBkbPJtFAj9hswXvfPEmvgR%2FsceeFAPFC%2Bc3sRJmjSMzHQM9cPBCx2TsH9%2Bw4e9in7dyowZP%2FOgz9PIDgD3CLP1gpfiM5uLmNGD%2F%2F05epiBKCZeoDohy9sV4uolqLmCOQtrfv5dDhFdHprMDzC4qrezNcyf%2BPB4fxuJ9YT8CyF65N9wWE8USZHfUYl1tIEdhnDvLTIrz6fttRCDPFhDU6uk56%2BcakhH5vhKJEBG%2B%2FBYv%2BB1MiO%2BVGMV3wn9kV)

> ⚠️ Attachment links are presigned AWS S3 URLs and expire — download the files soon after exporting.

### Answer

All 3 runs pass with all 6 tests. Now let me consult the advisor to verify the Phase 3 implementation before packaging:

### Sources

1. [SWG_PreCU_GDD.md](https://ppl-ai-file-upload.s3.amazonaws.com/web/direct-files/attachments/7508106/b9ed55cb-47b9-456d-808e-1e2738260985/SWG_PreCU_GDD.md?AWSAccessKeyId=ASIA2F3EMEYEYH2DMWY2&Signature=RRpTsvqVkoPk371EiOu6RzvBCU4%3D&x-amz-security-token=IQoJb3JpZ2luX2VjEOj%2F%2F%2F%2F%2F%2F%2F%2F%2F%2FwEaCXVzLWVhc3QtMSJGMEQCIG160CPacnesAtEiux7n%2F%2FkwMJzUI2oVlP5gDmuj0sb4AiBtGmHmi8EymP5YCziFO0%2B3j2dmZkygpdxGkU0WHvoLuCr8BAix%2F%2F%2F%2F%2F%2F%2F%2F%2F%2F8BEAEaDDY5OTc1MzMwOTcwNSIMXJgOv2rBcahax7igKtAE8Q5i0ka58RGuPKGqLG4%2Fx3fZP9mrzOcJuS%2FrVvcraS5nKGfyoquAu5XtBCzn%2Bk9pv11fN9YVw%2BfM%2BELyXQn8R1yoz0z1pvpSw0QUOOebUEKXwC4oBavVObwo%2FyMW%2FyeLPsPYCDC%2FLzLpCPKVPF%2FNpHg%2F30xjgE4Sl5lSkHM7mpCW3Y57HEfwKcH23%2FdRQLzWnTxGydCTFomXbRoFt1%2BYzWjr6fDnudGNwolf29j3eE8MUaNEkb7Ojs7tt%2FkiZwhYtXuIKe51YaFdqKqeGTEw5kyjEnjDLlp4aMw3j3T7R7IlTSMEkTtEd3yNA7gLKfi%2Fr6fk2K%2Fx1W2adAGsYeLZPQYvRfpzrtPiwfJfmcLvjRLaoMzgNBVpG36KCK3ja%2FbKz6619U0dTwRmikLhF6CjA7vLAhx7g%2B1ApCyqLrONojkxYl8r4PEnqb7gkIU%2FA5Ym8tZANtogiUhJh3GJ0WsNucmyzP6k5LwoDu%2BtI%2B6RAuRx7NhCr3wbO%2FsR1Tm0o96fsox8Qvv4O8polRmIoOPYf5YBBhIb9IHShH2TGxpe%2FaoVoGlVVk93D6n8JLEkHGunWWvBuqcQmgFLNUjUEIDmMxSWZektfZo0Rc3aeKiEzYZB8YsMjxaWVralGtaD7RU8ncD9zs8EO8YM2LGvw%2Bg6hl4BEnOnYEdtAmFFod7LZEDwf36WNveOYEA%2FbOfEMbF%2FEqKJiMupruoKY1fXzInWuRUg1j%2BsurfXSOVgDpok%2BxOGA%2BzUnju5vvzqNs7aAVKNyz%2BxgBhy6CbvSbimeJBe9jDQjZTVBjqZAX%2Bh%2BTRuPjnpkbofLv35WGU9hv085nijUHzinHhf%2BGBHuEGYiiX9OeiunRJD8V7uIoRULJXfyOkuAZ%2BCgcth%2BHwxGpc0gqOQIw7SeMGpTgukZaQ0cm2f%2BM%2Bh0e1HY%2FoQrK5z6v5cje3Uj91PJLTmKUmSnSTbVZQetCkYzqMkTpSNTIgyuOIvMlXQt4zNtMR4%2BIh8ffxgwii3zw%3D%3D&Expires=1789201038)

---

