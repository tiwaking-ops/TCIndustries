# Testbed Game Server

Patterns-only mechanical testbed server (forked from a Star Wars Galaxies Pre-CU recreation; all SWG names stripped per HD-TST-01). Go server + Godot 4 client. Built phase-by-phase against `SWG_PreCU_GDD.md` (v2 — see below), currently through **Phase 3 (Combat)**, fully verified.

## Source of Truth

`SWG_PreCU_GDD.md` in the repo root is the authoritative design spec (retained here only as HISTORICAL reference material for the patterns being exercised, not as testbed design authority). It is **v2** — already corrected for a skill-point ladder / XP-formula bug found and fixed mid-project. If you ever see a *different* copy of this document (an older upload, a stale reference, anything without the v2 patch notes near Section 6.3.1), treat this repo's copy as authoritative, not it.

The GDD uses an explicit tagging convention (Section 1.4) — respect it, don't flatten it:
- `[MVP]` — required now. `[EXPANSION]` — deferred, don't build yet unless asked.
- `[HISTORICAL]` — documented original-game behavior.
- `[ASSUMPTION]` — a gap in the historical record, already resolved with a binding ruling for this project. Not an open question — use the number/rule given, don't re-litigate it.
- `[IMPLEMENTATION RECOMMENDATION]` — a suggestion, not a requirement; free to substitute an equivalent approach.

Section 30 is the phase-by-phase build roadmap. Section 28 is the consolidated data model reference. Don't redesign systems already specified — implement what's there, and flag it clearly (in code comments, the same way the existing code does) if something is genuinely ambiguous or missing, rather than inventing details silently.

## Non-Negotiable Rule: Verify, Don't Claim

Every phase of this project so far was built by actually running the server and every integration test end-to-end, not by writing code and assuming it works. This caught real bugs that "it compiles" never would have (a `GetCharacterByID` bug that silently discarded current health, an AI range-check inconsistency, two creature lairs seeded 74m apart — none visible from reading the code, all found by running it).

**Do the same. Before saying a phase is done:**
1. `go build ./...` and `go vet ./...` must be clean.
2. Actually start the server and run every existing `cmd/*test` binary against it — old phases too, not just the new one. A change to a shared file (`world.go`, `protocol.go`, `db.go`) can silently break an earlier phase.
3. Show the real pass/fail output. Never report "should work" or "this implements X" as if it were "this ran and produced X."
4. If a test needs to wait on something real (e.g., Phase 3's incapacitation test takes ~1-2 minutes of genuine tick-driven combat), let it — don't mock or fast-forward the thing actually being tested.

## Build & Run

```bash
cd tc-server
go build ./...              # must be clean
go vet ./...                # must be clean
go build -o /tmp/tcserver ./cmd/server/
rm -f tc.db                # fresh DB for a clean test run
HTTP_ADDR=0.0.0.0:8080 /tmp/tcserver &   # start in background, same shell session as the tests below
# poll http://localhost:8080/health until it returns {"status":"ok"} before testing
go run ./cmd/testclient/     # Phase 0
go run ./cmd/phase1test/     # Phase 1
go run ./cmd/phase2test/     # Phase 2
go run ./cmd/phase3test/     # Phase 3 (~1-2 min, real combat)
go run ./cmd/phase6test/     # Phase 6
go run ./cmd/phase7test/     # Phase 7 (~5 min fast-cycle; TESTBED_FAST_CYCLE=1)
```

Server and test client(s) must run in the *same* shell session/background job — starting the server in one command and testing in a separate one can lose the background process depending on environment.

`go.mod` has a `replace golang.org/x/crypto => github.com/golang/crypto ...` directive — that was a workaround for a network-restricted sandbox during original development. On a normal connection it's harmless to keep, and safe to delete if it bothers you.

**SQLite driver: pure Go, no CGO.** Uses `modernc.org/sqlite` (registered as driver name `"sqlite"`, not `"sqlite3"`). This replaced `github.com/mattn/go-sqlite3`, which requires cgo plus a MinGW-w64 toolchain; the resulting binary would not load on this machine
("not a valid Win32 application"). If you see that error, the driver was swapped back — restore the pure-Go one. Builds are plain `go build` with **no** `CGO_ENABLED=1` and no toolchain PATH needed. Note the fast-cycle tests also want `HTTP_ADDR=0.0.0.0:8080`: without it the default `:8080` binds `::` (IPv6) only on Windows, so `localhost` (127.0.0.1) gets connection-refused.

## Architecture

Three-layer split per GDD Section 29.2: real-time tick loop (`internal/handlers/combat.go`'s `StartTickLoop`, `internal/world`'s spatial grid), transactional request/response (`internal/handlers/auth.go`, `skills.go`), and DB persistence (`internal/database`, SQLite now, schema written to be Postgres-portable). Combat/creature/skill packages are pure logic with no DB dependency; `internal/database/*_db.go` files own persistence; `internal/handlers/*.go` files orchestrate.

## Current Status

Phases 0-7 complete and independently verified (see `README.md`'s patch notes for the exact bugs found and fixed in each). Phases 4 (Resources & Crafting), 5 (Economy), 6 (Social Support Professions), and 7 (Civic Systems) were built and verified end-to-end after the Phase 3 documentation below was written; the per-phase README sections for them are still to be written, but the integration suites (`cmd/phase6test`, `cmd/phase7test`) exist and pass (Phase 4 crafting and Phase 5 economy are exercised end-to-end by the Phase 7 harness).

**Phase 8 (Faction & PvP) is next but has no design authority yet.** The Master GDD leaves combat scope, risk model, and progression integration explicitly **TBD** (open question OQ-008, and repeated in the combat-systems row of the status register). Do not invent faction/PvP mechanics to fill that gap — implement only against boundaries the project supplies. See `AGENTS.md` at the repo root for the authority rules.

Known, intentionally-scoped Phase 3 simplifications are documented in the README under "Known Phase 3 Simplifications" — don't treat those as bugs to silently fix; they're sequenced for later phases (e.g., Combat Medic gating on Revive belongs whenever Phase 9 builds that profession, not now).

## Workflow

- Commit after each verified phase, before starting the next.
- Run `/clear` before starting a new phase's session — don't let one phase's debugging context bleed into the next.
- If you find a real bug while building something else (like the three above), fix it, but say plainly that's what happened rather than folding it in silently.
