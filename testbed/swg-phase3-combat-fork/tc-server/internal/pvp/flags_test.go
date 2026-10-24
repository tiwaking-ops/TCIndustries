package pvp

import (
	"testing"
	"time"
)

func TestFlagStateMachine_OvertAndCovert(t *testing.T) {
	now := time.Now()
	view := NewFlagView()

	// Initially covert, no aggro.
	if view.State != FlagCovert || view.HasPending() || view.AggroActive(now) {
		t.Fatal("bad initial flag view")
	}

	// Request overt.
	trans := view.RequestOvert(now)
	if trans.Pending != FlagOvert {
		t.Fatalf("pending = %q, want overt", trans.Pending)
	}

	// Apply view update.
	view.Pending = trans.Pending
	view.ReadyAt = trans.ReadyAt

	// Before delay passes, applying does nothing.
	view, changed := view.ApplyDueTransition(now.Add(1 * time.Second))
	if changed || view.State != FlagCovert {
		t.Fatal("transition applied too early")
	}

	// After delay passes, state becomes overt.
	future := now.Add(Cfg.OvertDelay + time.Millisecond)
	view, changed = view.ApplyDueTransition(future)
	if !changed || view.State != FlagOvert {
		t.Fatalf("transition failed to apply: changed=%v, state=%q", changed, view.State)
	}

	// Try to request covert while aggression is active -> blocked by D-02.
	view = view.RefreshAggro(future)
	if !view.AggroActive(future) {
		t.Fatal("aggression should be active")
	}
	_, ok := view.RequestCovert(future)
	if ok {
		t.Fatal("covert request should be blocked while aggression is active")
	}

	// After aggression expires, covert request succeeds.
	expired := future.Add(Cfg.AggressionTimer + time.Millisecond)
	view = view.MaybeExpireAggro(expired)
	if view.AggroActive(expired) {
		t.Fatal("aggression should have expired")
	}
	cTrans, ok := view.RequestCovert(expired)
	if !ok {
		t.Fatal("covert request should succeed after aggression expires")
	}
	if cTrans.Pending != FlagCovert {
		t.Fatalf("pending = %q, want covert", cTrans.Pending)
	}
}

func TestFactionChangePreconditions(t *testing.T) {
	now := time.Now()
	view := NewFlagView()

	// Covert, no aggro -> can change faction.
	ok, _ := view.CanChangeFaction(now)
	if !ok {
		t.Fatal("covert without aggro should be allowed to change faction")
	}

	// Overt -> blocked by D-04.
	view.State = FlagOvert
	ok, reason := view.CanChangeFaction(now)
	if ok || reason != "faction_change_blocked_overt" {
		t.Fatalf("overt should block faction change: ok=%v, reason=%q", ok, reason)
	}

	// Covert with active aggro -> blocked by D-04.
	view.State = FlagCovert
	view = view.RefreshAggro(now)
	ok, reason = view.CanChangeFaction(now)
	if ok || reason != "faction_change_blocked_aggression" {
		t.Fatalf("aggression should block faction change: ok=%v, reason=%q", ok, reason)
	}
}
