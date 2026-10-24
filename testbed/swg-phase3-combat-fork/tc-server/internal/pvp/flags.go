package pvp

import "time"

// FlagTransition is the outcome of resolving a toggle request against current
// state. At most one pending transition exists at a time; the opposite request
// cancels it (proposal Section 4.3).
type FlagTransition struct {
	// State is the confirmed current state (unchanged by the request itself).
	State FlagState
	// Pending is the state that will apply when ReadyAt passes ("" = none).
	Pending    FlagState
	ReadyAt    time.Time
	// AggressionExpiresAt is the absolute aggression-timer expiry (0 = clear).
	AggressionExpiresAt time.Time
}

// FlagView is the current per-character PvP flag state.
type FlagView struct {
	State               FlagState
	Pending             FlagState
	ReadyAt             time.Time
	AggressionExpiresAt time.Time
}

// HasPending reports whether a transition is queued and not yet applied.
func (v FlagView) HasPending() bool { return v.Pending != "" }

// TransitionIn returns the time until the pending state applies (0 if none).
func (v FlagView) TransitionIn(now time.Time) time.Duration {
	if !v.HasPending() || v.ReadyAt.IsZero() {
		return 0
	}
	if d := time.Until(v.ReadyAt); d > 0 {
		return d
	}
	// Give the tick loop a moment to apply it rather than reporting negative.
	return 0
}

// AggressionIn returns the remaining aggression time (0 if clear/expired).
func (v FlagView) AggressionIn(now time.Time) time.Duration {
	if v.AggressionExpiresAt.IsZero() {
		return 0
	}
	if d := time.Until(v.AggressionExpiresAt); d > 0 {
		return d
	}
	return 0
}

// AggroActive reports whether the aggression timer is currently running.
func (v FlagView) AggroActive(now time.Time) bool { return v.AggressionIn(now) > 0 }

// NewFlagView returns the default flag for a faction member: covert, no
// pending transition, no timer.
func NewFlagView() FlagView {
	return FlagView{State: FlagCovert}
}

// RequestOvert applies D-01: requesting overt queues a transition that
// completes after OvertDelay. Re-requesting overt while an overt transition is
// already pending is idempotent.
func (v FlagView) RequestOvert(now time.Time) FlagTransition {
	if v.State == FlagOvert {
		return FlagTransition{
			State:               FlagOvert,
			AggressionExpiresAt: v.AggressionExpiresAt,
		}
	}
	return FlagTransition{
		State:               v.State,
		Pending:             FlagOvert,
		ReadyAt:             now.Add(Cfg.OvertDelay),
		AggressionExpiresAt: v.AggressionExpiresAt,
	}
}

// RequestCovert applies D-02: requesting covert queues a transition after
// CovertDelay, but is refused while the aggression timer is active.
func (v FlagView) RequestCovert(now time.Time) (FlagTransition, bool) {
	if v.AggroActive(now) {
		return FlagTransition{}, false // blocked: D-02
	}
	if v.State == FlagCovert {
		return FlagTransition{
			State:               FlagCovert,
			AggressionExpiresAt: v.AggressionExpiresAt,
		}, true
	}
	return FlagTransition{
		State:               v.State,
		Pending:             FlagCovert,
		ReadyAt:             now.Add(Cfg.CovertDelay),
		AggressionExpiresAt: v.AggressionExpiresAt,
	}, true
}

// CancelPending clears a queued transition (opposite request cancels).
func (v FlagView) CancelPending() FlagView {
	v.Pending = ""
	v.ReadyAt = time.Time{}
	return v
}

// ApplyDueTransition promotes a pending state to current once ReadyAt passes.
// Returns the new view plus whether a change occurred.
func (v FlagView) ApplyDueTransition(now time.Time) (FlagView, bool) {
	if !v.HasPending() || v.ReadyAt.IsZero() || now.Before(v.ReadyAt) {
		return v, false
	}
	v.State = v.Pending
	v.Pending = ""
	v.ReadyAt = time.Time{}
	return v, true
}

// RefreshAggro extends the aggression timer to now + AggressionTimer (D-03).
func (v FlagView) RefreshAggro(now time.Time) FlagView {
	v.AggressionExpiresAt = now.Add(Cfg.AggressionTimer)
	return v
}

// MaybeExpireAggro clears the timer once it has passed.
func (v FlagView) MaybeExpireAggro(now time.Time) FlagView {
	if !v.AggressionExpiresAt.IsZero() && !now.Before(v.AggressionExpiresAt) {
		v.AggressionExpiresAt = time.Time{}
	}
	return v
}

// CanChangeFaction applies D-04: switching or leaving a faction is blocked
// while overt or while the aggression timer is active.
func (v FlagView) CanChangeFaction(now time.Time) (bool, string) {
	if Cfg.BlockFactionChangeWhileOvert && v.State == FlagOvert {
		return false, "faction_change_blocked_overt"
	}
	if Cfg.BlockFactionChangeWhileAggro && v.AggroActive(now) {
		return false, "faction_change_blocked_aggression"
	}
	return true, ""
}
