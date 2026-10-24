# TCIndustries Phase 8 — Faction & PvP: Formal Implementation Proposal

- **Author:** Claude (Anthropic)
- **Version:** 1.0.0 (Draft for human confirmation)
- **Date:** October 24, 2026
- **Status:** **HUMAN-LOCKED (approved 2026-10-25 by project owner).** Section 2 rulings R-01..R-08 were human-supplied and final. Section 3 defaults D-01..D-16 were Claude's proposals and were confirmed by the owner's blanket approval ("I approve the proposal"), which confirms the defaults as written including their convention that silence on the Section 13 checklist means the proposed default stands. Both sections are now binding.
- **Target Audience:** An LLM coding agent that will implement Phase 8 of the TCIndustries testbed server.
- **Baseline Document:** `Phase_8_Formal_Specification_Report.md` v1.0.0 (current-state architecture, Phases 0–7 complete).
- **Resolves:** GDD open question `OQ-008` (Phase 8 scope, risk model, progression integration) — **resolved by this approval**.

---

## 1. Purpose and How to Use This Document

This proposal gives an implementing LLM the complete sketch of Phase 8: what to build, in what order, what is decided, what is only a proposed default, and how to prove it works.

**Reading order for the implementer:**
1. Read Section 2 (Rulings) and Section 3 (Pending defaults) first. They define every design decision.
2. Read Sections 4–9 for the exact behaviour, data, protocol and integration spec.
3. Follow Section 10 (work plan) in order, and satisfy Section 11 (verification) before declaring the phase complete.
4. Obey Section 12 (guardrails) throughout.

**Verify names against the code.** The baseline report names files, tables and message constants (e.g. `handleCombatAction`, `EnsureCivicSchema()`, `MsgCombatAction`). Confirm each against the repository before relying on it. If a name differs, follow the repository and note the discrepancy in your handoff.

---

## 2. Rulings (Human-Supplied, Final)

| ID | Ruling |
|----|--------|
| R-01 | Phase 8 PvP combines **three modes**: (a) duel / war declaration, (b) faction overt/covert flagging, (c) PvP designated zones (battlefields). |
| R-02 | Factions are **fixed and project-defined**: *neutral* plus **two** factions. |
| R-03 | Faction names are **placeholders in seed data** (`FACTION_A`, `FACTION_B`), to be renamed later by editing seed data only. No lore, no canon names, no code changes required to rename. |
| R-04 | Players **join, switch and leave** a faction only by interacting with a **faction NPC** (SWG-style). Switching sides is done at the NPC. Leaving is done only at the NPC by choosing "leave". There is no remote or UI-only faction change. |
| R-05 | Faction members **toggle** between covert and overt. An **aggression timer refreshes on combat**. |
| R-06 | A player defeated in PvP (any mode) goes through the **normal incapacitation / clone flow only**. No extra penalty: no item loss, no credit loss, no skill loss. |
| R-07 | Only **guilds** can declare war. Cities and factions cannot. |
| R-08 | Battlefield zones are **fixed rectangles or circles seeded in a database table**. |

---

## 3. Pending Defaults (Claude's Proposals — Need Confirmation)

These were not ruled on. Implement them as **configurable values in a single `PvPConfig` struct** (not scattered constants) so changing them never needs a refactor. Each is marked with the ID used in the confirmation checklist (Section 13).

| ID | Topic | Proposed default |
|----|-------|------------------|
| D-01 | Overt toggle delay (covert → overt) | 30 s |
| D-02 | Covert toggle delay (overt → covert) | 300 s, and only once the aggression timer has expired |
| D-03 | Aggression timer length | 120 s, refreshed to full on every PvP attack made or received |
| D-04 | Faction switch/leave restriction | Not allowed while overt or while the aggression timer is active |
| D-05 | Standing on switch/leave | Standing resets to 0 |
| D-06 | Standing awards | **None.** Phase 8 stores a `standing` column but awards no points. Kills are logged only |
| D-07 | Duel challenge expiry | 30 s |
| D-08 | Duel range | Both participants within 64 m to start; duel ends if separated beyond 128 m for more than 10 s |
| D-09 | Duel end | When either participant is incapacitated, forfeits, disconnects, or the duel times out (10 min) |
| D-10 | Who in a guild may declare/end war | Guild `leader` and `officer` roles |
| D-11 | War activation | Becomes `active` after a delay of 24 h (configurable, compressed under `TESTBED_FAST_CYCLE=1`) |
| D-12 | War end | Either guild's leader may end it (surrender); otherwise it never expires |
| D-13 | Where war combat is legal | Anywhere in the world, between members of the two warring guilds |
| D-14 | Battlefield rules | Only members of **opposing factions** may fight inside an active zone. Neutrals cannot attack or be attacked there by the zone rule. Entering forces **overt**. Leaving a zone applies the normal aggression timer before covert is allowed |
| D-15 | Faction NPC placement | Seeded placeholder positions, one recruiter NPC per faction; interaction radius 8 m |
| D-16 | Battlefield seed set | Two placeholder zones (one circle, one rectangle) at placeholder coordinates, marked active |

---

## 4. Domain Model

### 4.1 Factions
- Exactly three alignments exist: **neutral**, **FACTION_A**, **FACTION_B**.
- **Neutral is the absence of a `character_factions` row.** Do not create a neutral faction row.
- A character belongs to at most one faction.
- Faction rows live in a seeded `factions` table so names can change without code edits.

### 4.2 Faction membership (R-04)
Membership changes happen only through a **recruiter NPC interaction**:

| Action | Precondition | Result |
|--------|--------------|--------|
| `join` | Character is neutral; within NPC interaction radius | Creates `character_factions` row; initial PvP state `covert` |
| `switch` | Character is in the *other* faction; within radius; D-04 satisfied | Changes faction; applies D-05; state set to `covert` |
| `leave` | Character is in a faction; within radius; D-04 satisfied | Deletes row and flag state; character becomes neutral |

All three are validated server-side. The server never trusts a client-reported position or faction.

### 4.3 Flag states (R-05)
Per faction member, a PvP flag has a **state** and an **aggression timer**:

- `covert` — default. The member cannot attack or be attacked under the faction rule.
- `overt` — the member can attack and be attacked by **overt** members of the opposing faction.

**Toggle behaviour:**
- Requesting overt starts a pending transition that completes after D-01.
- Requesting covert starts a pending transition that completes after D-02, but is **rejected** (or deferred) while the aggression timer is active.
- Cancelling a pending transition is allowed by sending the opposite request.
- The aggression timer (D-03) refreshes to full on **every PvP attack made or received** (any mode), and also when leaving a battlefield zone.
- Neutral characters have no flag.

All transitions and timers are stored as absolute timestamps and processed in the tick loop, so they survive a server restart and work with `TESTBED_FAST_CYCLE=1`.

### 4.4 Duels (R-01a)
- A player challenges another player; the target accepts or declines (or the challenge expires, D-07).
- Only the two participants may attack each other under the duel rule; it grants no permissions against third parties.
- A duel ends per D-09. It never changes faction flags, but attacks made during a duel **do** refresh the aggression timer for faction members.
- Defeat is normal incapacitation (R-06). No penalty.

### 4.5 Guild wars (R-01a, R-07)
- Declared by a guild (D-10) against another guild. Status flow: `declared` → `active` (after D-11) → `ended`.
- While `active`, a member of one guild may attack a member of the other (D-13).
- Ended per D-12. History is kept.
- A guild cannot declare war on itself or on a guild it is already at war with (`declared` or `active`).
- Defeat is normal incapacitation (R-06).

### 4.6 Battlefields (R-01c, R-08)
- Seeded table of fixed zones; each is a circle or an axis-aligned rectangle in world coordinates (X/Z plane, matching existing structure and city coordinates).
- Entry and exit are detected server-side from authoritative position updates (not client claims).
- Rules per D-14.
- Defeat is normal incapacitation (R-06).

---

## 5. Attack Legality Engine

Create **one pure function** that decides whether player A may attack player T. All PvP gating goes through it so there is a single place to test and audit.

```go
// Sketch only: adapt names and types to the repository.
type PvPBasis string // "duel", "war", "faction", "battlefield", ""

type PvPDecision struct {
    Allowed bool
    Basis   PvPBasis // which rule permitted it
    Reason  string   // machine-readable rejection reason when !Allowed
}

func CanAttackPlayer(now time.Time, a, t CharacterPvPContext) PvPDecision
```

`CharacterPvPContext` carries what the decision needs: character ID, faction (or none), flag state, active duel ID (if any), guild ID (if any), and the battlefield zone ID containing the character (if any). The handler builds it from the database and spatial position; the function itself does no I/O.

**Evaluation order** (first match that allows wins; if none allow, reject):

1. **Self or same-account** attack → reject (`self`).
2. **Duel:** A and T are the two participants of the same `active` duel → allow (`duel`).
3. **War:** A and T are in guilds that are at `active` war with each other → allow (`war`).
4. **Battlefield:** A and T are both inside the **same active** zone, both have a faction, and the factions differ → allow (`battlefield`).
5. **Faction:** A and T have different factions, **both** are `overt` → allow (`faction`).
6. Otherwise → reject. Use specific reasons where useful: `neutral_target`, `same_faction`, `target_covert`, `attacker_covert`.

**Side effects of an allowed attack** (applied by the handler, not the pure function):
- Refresh the aggression timer for any participant that has a faction (D-03).
- Record the basis on the resulting `pvp_events` row when a defeat occurs.
- On defeat, run the **existing** incapacitation/clone flow unchanged (R-06).

**PvE is unchanged.** Attacks on creatures bypass this function entirely.

---

## 6. Database Schema

Create `internal/database/faction_db.go` with `EnsureFactionSchema()`, called at startup after the existing `Ensure*Schema()` calls. It must be **idempotent**, follow the existing `EnsureEconomySchema()` / `EnsureCivicSchema()` pattern, use `db.NewRowID(prefix)` for IDs, and respect `?_foreign_keys=on`. Timestamps follow whatever convention the existing schemas use (verify; absolute epoch milliseconds are recommended for timers).

```sql
CREATE TABLE IF NOT EXISTS factions (
    id           TEXT PRIMARY KEY,
    code         TEXT NOT NULL UNIQUE,      -- 'FACTION_A', 'FACTION_B'
    display_name TEXT NOT NULL,             -- renamed later via seed data
    created_at   INTEGER NOT NULL
);

CREATE TABLE IF NOT EXISTS character_factions (
    character_id TEXT PRIMARY KEY REFERENCES characters(id),
    faction_id   TEXT NOT NULL REFERENCES factions(id),
    standing     INTEGER NOT NULL DEFAULT 0,   -- stored, not awarded in Phase 8 (D-06)
    joined_at    INTEGER NOT NULL,
    changed_at   INTEGER NOT NULL
);

CREATE TABLE IF NOT EXISTS pvp_flags (
    character_id          TEXT PRIMARY KEY REFERENCES characters(id),
    state                 TEXT NOT NULL CHECK (state IN ('covert','overt')),
    pending_state         TEXT CHECK (pending_state IN ('covert','overt')),
    transition_at         INTEGER,              -- when pending_state applies
    aggression_expires_at INTEGER NOT NULL DEFAULT 0,
    last_aggressor_id     TEXT,
    updated_at            INTEGER NOT NULL
);

CREATE TABLE IF NOT EXISTS faction_recruiters (
    id            TEXT PRIMARY KEY,
    faction_id    TEXT NOT NULL REFERENCES factions(id),
    name          TEXT NOT NULL,
    zone          TEXT NOT NULL,
    pos_x         REAL NOT NULL,
    pos_z         REAL NOT NULL,
    interact_m    REAL NOT NULL DEFAULT 8.0
);

CREATE TABLE IF NOT EXISTS pvp_zones (
    id        TEXT PRIMARY KEY,
    name      TEXT NOT NULL,
    zone      TEXT NOT NULL,
    shape     TEXT NOT NULL CHECK (shape IN ('circle','rect')),
    center_x  REAL, center_z REAL, radius_m REAL,        -- circle
    min_x     REAL, min_z    REAL, max_x    REAL, max_z REAL,  -- rect
    active    INTEGER NOT NULL DEFAULT 1
);

CREATE TABLE IF NOT EXISTS pvp_duels (
    id            TEXT PRIMARY KEY,
    challenger_id TEXT NOT NULL REFERENCES characters(id),
    target_id     TEXT NOT NULL REFERENCES characters(id),
    status        TEXT NOT NULL CHECK (status IN ('pending','active','ended','declined','expired')),
    created_at    INTEGER NOT NULL,
    started_at    INTEGER,
    ended_at      INTEGER,
    winner_id     TEXT,
    end_reason    TEXT
);

CREATE TABLE IF NOT EXISTS guild_wars (
    id                      TEXT PRIMARY KEY,
    declaring_guild_id      TEXT NOT NULL REFERENCES guilds(id),
    target_guild_id         TEXT NOT NULL REFERENCES guilds(id),
    declared_by_character_id TEXT NOT NULL REFERENCES characters(id),
    status                  TEXT NOT NULL CHECK (status IN ('declared','active','ended')),
    declared_at             INTEGER NOT NULL,
    active_from             INTEGER NOT NULL,
    ended_at                INTEGER,
    ended_by_character_id   TEXT,
    end_reason              TEXT
);

CREATE TABLE IF NOT EXISTS pvp_events (
    id          TEXT PRIMARY KEY,
    kind        TEXT NOT NULL,             -- 'defeat'
    basis       TEXT NOT NULL,             -- duel|war|faction|battlefield
    attacker_id TEXT NOT NULL,
    target_id   TEXT NOT NULL,
    zone_id     TEXT,
    created_at  INTEGER NOT NULL
);
```

Add the supporting indexes you need (e.g. on `guild_wars(status)`, `pvp_duels(status)`, `pvp_events(attacker_id)`). Verify that `guilds(id)` and `characters(id)` are the real referenced key names.

**Seed data (idempotent, `INSERT OR IGNORE`):**
- Two `factions` rows: `FACTION_A`, `FACTION_B`, with placeholder display names.
- One `faction_recruiters` row per faction at placeholder positions (D-15).
- Two `pvp_zones` rows (D-16).

Provide DB-layer functions in the same file (get/set faction, upsert flag, list zones, create/advance duel and war, insert event). Keep all SQL inside `internal/database/`; handlers never run raw SQL.

---

## 7. Protocol

Add constants and payload types in `internal/protocol/protocol.go`, following existing naming and JSON conventions.

### 7.1 Client → Server (WebSocket)
| Message | Payload | Purpose |
|---------|---------|---------|
| `MsgTogglePvP` | `{ "state": "overt" \| "covert" }` | Request a flag transition (4.3). Also cancels a pending opposite transition |
| `MsgSetFaction` | `{ "npc_id": string, "action": "join" \| "switch" \| "leave", "faction_id"?: string }` | Faction NPC interaction (4.2). `faction_id` required for `join`/`switch` |
| `MsgDuelChallenge` | `{ "target_id": string }` | Challenge a player |
| `MsgDuelResponse` | `{ "duel_id": string, "accept": bool }` | Accept or decline |
| `MsgDuelForfeit` | `{ "duel_id": string }` | Forfeit an active duel |

### 7.2 Server → Client (WebSocket pushes)
| Message | Payload | Purpose |
|---------|---------|---------|
| `MsgFactionUpdate` | faction code/name, standing | Sent on join/switch/leave and on login |
| `MsgPvPStateChange` | `state`, `pending_state`, `transition_in_ms`, `aggression_remaining_ms`, `in_battlefield` | Sent on any flag, timer or zone change |
| `MsgDuelUpdate` | duel id, status, opponent, reason | Challenge received, started, ended |
| `MsgPvPError` (or the project's existing error pattern) | reason code | Rejected PvP request or attack |

Reuse the existing combat receipts (`MsgCombatResult`, `MsgHAMUpdate`) for PvP hits. Do not fork the combat result format.

### 7.3 Entity broadcast extension
Extend `EntitySpawnMsg` (and any entity update the interest grid sends) with optional, `omitempty` fields: `faction_id`, `pvp_state`. Existing clients that ignore unknown fields must keep working. Visibility rules remain governed by the existing 128 m `InterestRadius` grid.

### 7.4 HTTP API (guarded by `AuthMiddleware`)
| Endpoint | Purpose |
|----------|---------|
| `GET /api/faction/me` | Own faction, standing, flag state |
| `GET /api/pvp/zones` | List active battlefield zones |
| `GET /api/pvp/wars` | Wars involving the caller's guild |
| `POST /api/pvp/war/declare` | Body `{ "target_guild_id": string }` — D-10 role check |
| `POST /api/pvp/war/end` | Body `{ "war_id": string }` — D-12 |

Use the existing JSON error and response conventions in `internal/handlers/` (see `civic.go` for the pattern).

---

## 8. Server Integration Points

1. **`internal/handlers/combat.go`:** In `handleCombatAction`, when the target is a player, build both `CharacterPvPContext`s and call `CanAttackPlayer`. On reject, send the error and do **not** resolve the attack. On allow, resolve normally with `combat.ResolveAttack`, then apply the side effects in Section 5.
2. **Tick loop (`StartTickLoop`):** Each tick (or on a coarser sub-interval) process: pending flag transitions, aggression-timer expiry broadcasts, duel expiry/range checks/timeouts, and `declared → active` war activation.
3. **Movement/position handling (`internal/handlers/world.go`, `internal/world`):** After an accepted position update, test battlefield containment. On entry, force `overt` (D-14), mark `in_battlefield`, and push `MsgPvPStateChange`. On exit, refresh the aggression timer and push state.
4. **Incapacitation path:** On a defeat of a player by a player, insert a `pvp_events` row and end an active duel if the defeat was inside it. Do not alter the existing incapacitation/clone logic.
5. **Faction NPC interaction:** Validate NPC existence, interaction radius against authoritative position, and the preconditions in 4.2. Use whichever existing NPC/service interaction pattern the repo already has if one fits; otherwise add a minimal handler.
6. **Login/session start:** Send `MsgFactionUpdate` and `MsgPvPStateChange` so the client has current state.
7. **Guild integration:** Role checks use the existing `guild_members` roles (`leader`, `officer`, `member`) in `civic_social_db.go`. Do not change guild tables.
8. **Wiring order:** Register routes and WebSocket message handlers following the existing registration style; call `EnsureFactionSchema()` in the startup sequence.

---

## 9. Edge Cases the Implementation Must Handle

- Attacking while the target is already incapacitated (follow the existing combat rules; do not add PvP-specific behaviour).
- Disconnects: pending flag transitions and timers persist by timestamp. Active duels end on disconnect (D-09).
- A character leaving a guild or a guild disbanding while a war is `declared`/`active`: the war record persists; legality is evaluated from current guild membership at attack time.
- A character who changes faction while in a battlefield: rejected by D-04 (overt is forced there).
- Two overlapping zones: containment returns the first matching active zone; document the choice.
- Concurrent requests (two simultaneous toggles, two duel challenges): resolve with transactions so state never becomes inconsistent.
- A player in both a duel and a war with the same opponent: legality allows; the recorded basis follows the evaluation order in Section 5.

---

## 10. Work Plan (Implement in This Order)

| Step | Deliverable | Done when |
|------|-------------|-----------|
| 1 | `PvPConfig` struct with all D-values, plus a fast-cycle variant keyed to `TESTBED_FAST_CYCLE=1` | Compiles; defaults match Section 3 |
| 2 | `faction_db.go`: schema, seed data, DB functions | `EnsureFactionSchema()` is idempotent on a fresh and an existing DB |
| 3 | Protocol types and constants | Builds; JSON shapes match Section 7 |
| 4 | Faction membership via NPC (4.2) | join/switch/leave work and reject bad cases |
| 5 | Flag state, toggle, aggression timer, tick processing (4.3) | Transitions and timers behave per D-01..D-04 |
| 6 | `CanAttackPlayer` plus **unit tests** for every rule and rejection reason | Table-driven tests pass |
| 7 | Combat gating in `handleCombatAction` | PvP attacks gated; PvE unchanged |
| 8 | Duels (4.4) | Full challenge → fight → end flow works |
| 9 | Guild wars (4.5) | Declare, activate, attack legality, end |
| 10 | Battlefields (4.6) | Entry/exit, forced overt, legality |
| 11 | Broadcast extension and login push | Nearby clients see faction/state |
| 12 | `cmd/phase8test/main.go`, regression runs, docs handoff | Section 11 passes |

Do not start a step until the previous one builds and its tests pass.

---

## 11. Verification (Non-Negotiable)

1. `go build ./...` and `go vet ./...` return zero errors.
2. Add `cmd/phase8test/main.go`, run against a **live** server with `TESTBED_FAST_CYCLE=1` and a **fresh** SQLite database.
3. All prior suites (`phase1test`, `phase2test`, `phase3test`, `phase6test`, `phase7test`) still pass.
4. `phase8test` must cover at least:
   - Faction: join, switch, leave via NPC; rejected when out of radius, when flagged (D-04), when joining while already in a faction, when switching to the same faction.
   - Flags: covert → overt after delay; overt → covert blocked during the aggression timer; timer refresh on attack made and received; state persists across a server restart.
   - Legality matrix: neutral vs neutral rejected; same-faction rejected; covert vs overt rejected; overt vs overt (different factions) allowed.
   - Duel: challenge, accept, decline, expiry; only the two participants can fight; ends on incapacitation; no effect on third parties.
   - War: declare by leader/officer; rejected for plain members, self-war, duplicate war; activation after delay; cross-guild attacks allowed only when active; end by either leader.
   - Battlefield: forced overt on entry; opposing factions can fight; neutrals and same-faction cannot; exit applies the timer; circle and rectangle containment edges.
   - Defeat: normal incapacitation/clone flow, **no** item, credit or skill change (R-06), and a `pvp_events` row written.
   - Broadcast: nearby client receives faction and state fields.
   - PvE regression: creature combat unaffected.
5. Report results honestly. If a test fails, fix it or list it as a known issue. Never mark the phase complete on a partial pass.

---

## 12. Guardrails for the Implementer

- **No speculative design.** Implement Section 2 exactly. Implement Section 3 as configurable defaults. Anything not covered here is **not authorized**: stop and list it as a question rather than deciding it.
- **No lore.** Faction and zone names stay placeholders. Do not invent canon, backstory or NPC dialogue beyond minimal functional labels.
- **No extra penalties or rewards.** R-06 and D-06 stand. Do not add loot drops, standing awards, bounties or ranks.
- **Do not change** existing table columns, existing message formats' required fields, or Phase 0–7 behaviour. Additions only.
- **Server authority.** Positions, factions, guild roles and timers are always read server-side.
- **Style.** Follow existing package layout, naming, error handling and the `Ensure*Schema()` idempotent migration pattern.
- **Documentation.** Per project rules, record what was implemented and the rulings it rests on in the project documentation. Do not mark `OQ-008` resolved in the GDD unless the human owner has approved this proposal.

---

## 13. Confirmation Checklist — CONFIRMED 2026-10-25

The project owner approved the proposal as a whole on 2026-10-25 ("I approve the
proposal"). Per this document's own convention — *silence on the checklist means
the proposed default stands* — a blanket approval confirms every default below as
written. No item was corrected or withdrawn.

- [x] **D-01 / D-02 / D-03:** Toggle delays (30 s / 300 s) and aggression timer (120 s).
- [x] **D-04 / D-05:** Faction change blocked while overt or in combat; standing resets on switch/leave.
- [x] **D-06:** No standing awards in Phase 8 (kills logged only).
- [x] **D-07 – D-09:** Duel challenge expiry, range and end conditions.
- [x] **D-10:** Guild leader **and** officers may declare/end war (or leader only).
- [x] **D-11 / D-12:** 24 h war activation delay; either leader may end a war; no automatic expiry.
- [x] **D-13:** War combat is legal anywhere in the world.
- [x] **D-14:** Battlefields: opposing factions only, neutrals excluded, forced overt on entry.
- [x] **D-15 / D-16:** Placeholder NPC and battlefield seed positions are acceptable.
- [x] **Extra:** May a player have a faction **and** be in a duel/war at once (this proposal says yes).

Numeric values (durations, ranges) remain subject to a future balance pass
(BAL-001) if the owner chooses to run one — this approval fixes the *rules*, and
adopts the stated numbers as the current working set.

---

## 14. Change History

| Version | Date | Notes |
|---------|------|-------|
| 1.0.0 | 2026-10-24 | Initial draft built from owner rulings R-01..R-08 and proposed defaults D-01..D-16. |
