package pvp

import (
	"database/sql"
	"os"
	"testing"
	"time"

	"tc-server/internal/database"
	"tc-server/internal/models"
)

func TestFlagPersistenceAcrossRestart(t *testing.T) {
	dbPath := "/tmp/test_pvp_restart.db"
	_ = os.Remove(dbPath)
	_ = os.Remove(dbPath + "-shm")
	_ = os.Remove(dbPath + "-wal")
	defer func() {
		_ = os.Remove(dbPath)
		_ = os.Remove(dbPath + "-shm")
		_ = os.Remove(dbPath + "-wal")
	}()

	// Open first database instance.
	db1, err := database.New(dbPath)
	if err != nil {
		t.Fatalf("db1 open failed: %v", err)
	}
	if err := db1.EnsureFactionSchema(); err != nil {
		t.Fatalf("ensure schema failed: %v", err)
	}

	// Character row is required by the pvp_flags foreign key.
	if err := db1.CreateAccount("acc-1", "u1", "h"); err != nil {
		t.Fatalf("account insert failed: %v", err)
	}
	if err := db1.CreateCharacter(&models.Character{
		ID: "char-1", AccountID: "acc-1", Name: "TestChar", Species: "human",
	}); err != nil {
		t.Fatalf("character insert failed: %v", err)
	}

	// Write flag state with specific timers.
	now := time.Now().Unix()
	row := &database.PvPFlagRow{
		CharacterID:         "char-1",
		State:               string(FlagOvert),
		PendingState:        sql.NullString{String: string(FlagCovert), Valid: true},
		TransitionAt:        sql.NullInt64{Int64: now + 30, Valid: true},
		AggressionExpiresAt: now + 120,
		UpdatedAt:           now,
	}
	if err := db1.UpsertPvPFlag(row); err != nil {
		t.Fatalf("upsert flag failed: %v", err)
	}
	db1.Close()

	// Open second database instance on the same file (simulating server restart).
	db2, err := database.New(dbPath)
	if err != nil {
		t.Fatalf("db2 open failed: %v", err)
	}
	defer db2.Close()

	loaded, err := db2.GetPvPFlag("char-1")
	if err != nil {
		t.Fatalf("get pvp flag after restart failed: %v", err)
	}

	if loaded.State != string(FlagOvert) {
		t.Fatalf("state = %q, want overt", loaded.State)
	}
	if !loaded.PendingState.Valid || loaded.PendingState.String != string(FlagCovert) {
		t.Fatalf("pending = %v, want covert", loaded.PendingState)
	}
	if loaded.TransitionAt.Int64 != now+30 {
		t.Fatalf("transition_at = %d, want %d", loaded.TransitionAt.Int64, now+30)
	}
	if loaded.AggressionExpiresAt != now+120 {
		t.Fatalf("aggression_expires_at = %d, want %d", loaded.AggressionExpiresAt, now+120)
	}
}

// TestPvPEventLogging proves the D-06 path: a defeat is logged with its basis,
// and never awards anything (there is no award channel to test — the point is
// that the row exists and carries the basis the engine returned).
func TestPvPEventLogging(t *testing.T) {
	dbPath := "/tmp/test_pvp_events.db"
	for _, s := range []string{"", "-shm", "-wal"} {
		_ = os.Remove(dbPath + s)
	}
	defer func() {
		for _, s := range []string{"", "-shm", "-wal"} {
			_ = os.Remove(dbPath + s)
		}
	}()

	db, err := database.New(dbPath)
	if err != nil {
		t.Fatalf("db open failed: %v", err)
	}
	defer db.Close()
	if err := db.EnsureFactionSchema(); err != nil {
		t.Fatalf("ensure schema failed: %v", err)
	}
	if err := db.CreateAccount("acc-1", "u1", "h"); err != nil {
		t.Fatalf("account insert failed: %v", err)
	}
	for _, c := range []struct{ id, name string }{
		{"char-a", "Alpha"}, {"char-b", "Bravo"},
	} {
		if err := db.CreateCharacter(&models.Character{
			ID: c.id, AccountID: "acc-1", Name: c.name, Species: "human",
		}); err != nil {
			t.Fatalf("character insert failed: %v", err)
		}
	}

	zoneID := sql.NullString{String: "zone-circle-01", Valid: true}
	if err := db.RecordPvPEvent("ev-1", "defeat", string(BasisBattlefield), "char-a", "char-b", zoneID); err != nil {
		t.Fatalf("record event failed: %v", err)
	}

	events, err := db.PvPEventsFor("char-a")
	if err != nil {
		t.Fatalf("read events failed: %v", err)
	}
	if len(events) != 1 {
		t.Fatalf("got %d events, want 1", len(events))
	}
	e := events[0]
	if e.Basis != string(BasisBattlefield) {
		t.Fatalf("basis = %q, want battlefield", e.Basis)
	}
	if e.AttackerID != "char-a" || e.TargetID != "char-b" {
		t.Fatalf("attacker/target = %s/%s, want char-a/char-b", e.AttackerID, e.TargetID)
	}
	if !e.ZoneID.Valid || e.ZoneID.String != "zone-circle-01" {
		t.Fatalf("zone = %v, want zone-circle-01", e.ZoneID)
	}

	// The target sees it too (query by the other side alone).
	events, err = db.PvPEventsFor("char-b")
	if err != nil {
		t.Fatalf("read events for target failed: %v", err)
	}
	if len(events) != 1 {
		t.Fatalf("target got %d events, want 1", len(events))
	}
}
