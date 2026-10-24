// Package economy implements Phase 5 constants: upkeep magnitudes, ranges,
// drop mapping, and ledger category vocabulary. Testbed fork; generic only.
//
// GDD-given values are cited inline. Everything else is [PROVISIONAL] under the
// owner-approved convention (proposal §6): magnitudes reuse fork-local scales
// (training-cost and harvester-fee bands) to avoid inventing fresh tuning, and
// every number is flagged NON-CANONICAL here and in HYGIENE_NOTE.md.
package economy

import (
	"os"
	"time"
)

// Ledger flows (GDD 12.6 shape).
const (
	FlowFaucet   = "faucet"
	FlowSink     = "sink"
	FlowTransfer = "transfer"
)

// Ledger categories (GDD 12.6 vocabulary, subset reachable in this build).
const (
	CatTrainingCost    = "training_cost"    // sink (existing, retro-tagged)
	CatMaintenanceFee  = "maintenance_fee"  // sink (harvester/vendor/house)
	CatHarvesterFund   = "harvester_fund"   // transfer (wallet → pool)
	CatStructureFund   = "structure_fund"   // transfer (wallet → pool)
	CatVendorSale      = "vendor_sale"      // transfer pair (buyer → till)
	CatCreatureDrop    = "creature_drop"    // faucet (GDD 12.2.2 band)
	CatStartingCredits = "starting_credits" // faucet (character creation; tagged going forward)
)

// House tiers (GDD 13.2.2 S/M/L; guild hall EXPANSION, not built).
// Upkeep values [PROVISIONAL]: small relative to harvester-fee scale per GDD
// 13.2.5 ("housing should never be the dominant sink").
var HouseUpkeepWeekly = map[string]int{
	"small":  200,
	"medium": 500,
	"large":  1000,
}

// Vendor upkeep: base + per active listing (GDD 21.2.2 "scales with number of
// active listings" — the curve itself is GDD-silent). [PROVISIONAL]
const (
	VendorUpkeepBase       = 300
	VendorUpkeepPerListing = 50
)

// VendorSlotCap: GDD [ASSUMPTION] 50 for a Novice-placed vendor.
const VendorSlotCap = 50

// NoBuildRadiusM: minimum separation between structures [PROVISIONAL — GDD gives
// 100 m for harvesters only, nothing for houses/vendors].
const NoBuildRadiusM = 20.0

// PurchaseRangeM / CollectRangeM: in-person distances [PROVISIONAL — GDD says
// "walk up" / "in person" without metres].
const (
	PurchaseRangeM = 15.0
	CollectRangeM  = 15.0
)

// TerminalSearchRadiusM: must be near a terminal to search [PROVISIONAL].
const TerminalSearchRadiusM = 30.0

// DropCreditsPerCL: creature-drop faucet mapping inside the GDD 10–200 band
// (0001 CL1 → 10, 0002 CL3 → 30). [PROVISIONAL mapping, GDD-given band.]
const DropCreditsPerCL = 10

// Maintenance windows: GDD-given defaults; fast-cycle test mapping (same
// convention as Phase 4 — defaults GDD-given, compression test-config-only).
var (
	HouseGracePeriod = 14 * 24 * time.Hour // condemned grace (GDD [ASSUMPTION])
	ReclaimAfter     = 180 * 24 * time.Hour // inactive-structure reclaim (GDD [ASSUMPTION])
	fastCycle        = os.Getenv("TESTBED_FAST_CYCLE") == "1"
)

func init() {
	if fastCycle {
		HouseGracePeriod = 3 * time.Minute
		ReclaimAfter = 30 * time.Minute
	}
}

// Structure statuses.
const (
	StatusActive    = "active"
	StatusClosed    = "closed"    // vendor lapsed: no sales, stock safe (GDD 21.2.2)
	StatusCondemned = "condemned" // house lapsed past grace: destroy on expiry
	StatusDestroyed = "destroyed"
)
