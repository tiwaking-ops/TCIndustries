package world

import (
	"fmt"
	"time"
)

// Movement validation constants.
// GDD Section 20.2.1: On Foot speed is 5-7 m/s.
// GDD Section 29.2.1: Server is authoritative — all game state computed server-side.

const (
	// MaxOnFootSpeed is the maximum on-foot movement speed in m/s (GDD 20.2.1).
	MaxOnFootSpeed = 7.0

	// SpeedTolerance is the multiplier applied to the max speed to allow for
	// minor client-server clock differences, jitter, and rounding.
	// 1.5x gives us a practical ceiling of 10.5 m/s before we flag it.
	SpeedTolerance = 1.5

	// MinMoveInterval is the minimum time between move messages before we
	// coalesce them (ignore tiny dt to avoid divide-by-zero / jitter).
	MinMoveInterval = 50 * time.Millisecond

	// JitterAllowance is a flat distance in meters that is always allowed
	// regardless of dt, to accommodate floating-point jitter and minor
	// client-side prediction errors. Kept small to prevent speed bypass.
	JitterAllowance = 0.3 // meters

	// PlanetBounds is the world boundary for each planet in meters.
	// SWG planets are roughly 16km x 16km (-8192 to +8192).
	PlanetBounds = 8192.0

	// InterestRadius is the distance at which entities become visible to
	// each other. Within this radius, spawn/move/despawn messages are sent.
	// GDD Section 29.2.2: interest management grid.
	InterestRadius = 128.0

	// SayRadius is the range of /say spatial chat in meters (GDD 18.2.1).
	SayRadius = 20.0

	// ShoutRadius is the range of /shout spatial chat in meters (GDD 18.2.1).
	ShoutRadius = 50.0
)

// MoveValidationResult indicates whether a movement was accepted or rejected.
type MoveValidationResult struct {
	Accepted       bool
	Coalesced      bool // true if the move was accepted but should NOT be applied (dt too small)
	Reason         string
	CorrectedPos   Position // The position the client should be at (if rejected)
	MaxAllowedDist float64  // The maximum distance that was allowed
	ActualDist     float64  // The distance the client tried to move
}

// ValidateMovement checks whether a client-reported position change is valid.
//
// The server tracks each player's last known position and the timestamp of
// their last accepted move. When a new move message arrives, the server
// computes the implied speed (distance / time delta) and rejects the move
// if it exceeds the maximum allowed speed.
//
// GDD Section 29.2.1: "every combat roll, crafting result, and economic
// transaction is resolved server-side and pushed to clients — never
// client-computed and merely reported." Movement follows the same principle.
//
// GDD Section 29.5: "All combat, crafting, and experimentation rolls are
// computed server-side using server-seeded randomness; clients never submit
// a roll result, only an action request." Movement is an action request.
func ValidateMovement(oldPos, newPos Position, dt time.Duration) MoveValidationResult {
	// Check planet boundary
	if err := CheckPlanetBounds(newPos); err != "" {
		return MoveValidationResult{
			Accepted:     false,
			Reason:       err,
			CorrectedPos: oldPos,
		}
	}

	// Different planet — this should never happen via a move message
	if oldPos.Planet != newPos.Planet {
		return MoveValidationResult{
			Accepted:     false,
			Reason:       "cannot move between planets via movement",
			CorrectedPos: oldPos,
		}
	}

	// Ignore tiny time deltas to avoid false positives from jitter.
	// Coalesced moves are accepted but NOT applied — the client's position
	// stays where it was. This prevents a bypass where a client sends a huge
	// move within 50ms.
	if dt < MinMoveInterval {
		return MoveValidationResult{
			Accepted:   true,
			Coalesced:  true,
			Reason:     "coalesced (dt too small)",
			ActualDist: Distance2D(oldPos, newPos),
		}
	}

	dist := Distance2D(oldPos, newPos)
	seconds := dt.Seconds()
	maxAllowed := MaxOnFootSpeed*SpeedTolerance*seconds + JitterAllowance

	if dist > maxAllowed {
		return MoveValidationResult{
			Accepted:       false,
			Reason:         fmt.Sprintf("speed exceeded: %.1f m in %.3fs (%.1f m/s, max %.1f m/s)", dist, seconds, dist/seconds, MaxOnFootSpeed*SpeedTolerance),
			CorrectedPos:   oldPos,
			MaxAllowedDist: maxAllowed,
			ActualDist:     dist,
		}
	}

	return MoveValidationResult{
		Accepted:       true,
		ActualDist:     dist,
		MaxAllowedDist: maxAllowed,
	}
}

// CheckPlanetBounds validates that a position is within the planet's boundaries.
func CheckPlanetBounds(pos Position) string {
	if pos.X < -PlanetBounds || pos.X > PlanetBounds {
		return fmt.Sprintf("X coordinate %.1f out of bounds (%.0f to %.0f)", pos.X, -PlanetBounds, PlanetBounds)
	}
	if pos.Z < -PlanetBounds || pos.Z > PlanetBounds {
		return fmt.Sprintf("Z coordinate %.1f out of bounds (%.0f to %.0f)", pos.Z, -PlanetBounds, PlanetBounds)
	}
	// Y (altitude) is not bounded for Phase 1 — terrain validation comes later
	return ""
}

// ChatChannel represents the scope of a chat message.
// GDD Section 18.2.1.
type ChatChannel string

const (
	ChannelSpatial ChatChannel = "spatial" // /say, ~20m radius
	ChannelShout   ChatChannel = "shout"   // /shout, ~50m radius
	ChannelPlanet  ChatChannel = "planet"  // planet-wide
	ChannelSystem  ChatChannel = "system"  // server messages
)

// ChatRadius returns the delivery radius for a chat channel.
// Returns 0 for channels that are not range-limited (planet, system).
func ChatRadius(channel ChatChannel) float64 {
	switch channel {
	case ChannelSpatial:
		return SayRadius
	case ChannelShout:
		return ShoutRadius
	default:
		return 0 // planet-wide or system-wide
	}
}
