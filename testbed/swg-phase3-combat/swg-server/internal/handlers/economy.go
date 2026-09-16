// Economy world handlers for Phase 5: structures (houses/vendors), vendor
// purchase/till, bazaar search, structure maintenance ticks, and seeding.
// Testbed fork; generic content only. GDD-silent numbers [PROVISIONAL] per the
// owner-approved convention; GDD-given values cited.
package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"swg-server/internal/database"
	"swg-server/internal/economy"
	"swg-server/internal/protocol"
)

// SeedEconomyWorld prepares Phase 5 persistence (ledger, structures, listings,
// snapshots, terminals, market history, vendor-till column). Idempotent.
func (h *WorldHandler) SeedEconomyWorld() error {
	if err := h.db.EnsureEconomySchema(); err != nil {
		return err
	}
	if err := h.db.EnsureMarketTable(); err != nil {
		return err
	}
	if err := h.db.EnsureVendorTillColumn(); err != nil {
		return err
	}
	log.Printf("Economy world seeded: ledger, structures, listings, terminal")
	return nil
}

// tickStructures runs structure maintenance each harvest interval (= one game
// hour by construction): deduct pro-rated weekly upkeep (sink-tagged); lapsed
// vendors close (stock safe); lapsed houses enter condemned grace, then destroy
// past the grace window (fast-cycle mapping test-config-only, defaults GDD-given).
func (h *WorldHandler) tickStructures() {
	all, err := h.db.AllStructures()
	if err != nil {
		log.Printf("structure tick: list failed: %v", err)
		return
	}
	now := time.Now()
	for _, st := range all {
		if st.Status == economy.StatusDestroyed {
			continue
		}
		if st.Kind == "city_hall" {
			continue // Phase 7: hall upkeep is city-treasury-level, never pooled
		}
		if st.Kind == "vendor" {
			h.tickVendor(&st)
			continue
		}
		// Houses: upkeep or grace progression.
		fee := (st.UpkeepWeekly + 167) / 168 // pro-rated ceil (flagged discretization)
		if st.MaintenancePool >= fee && fee > 0 {
			st.MaintenancePool -= fee
			_ = h.db.RecordLedger(st.OwnerCharacterID, -fee, "maintenance_fee", "sink", st.ID, st.Zone)
			if st.Status == economy.StatusCondemned {
				st.Status = economy.StatusActive // caught up: grace lifted (flagged leniency)
			}
			_ = h.db.UpdateStructure(&st)
			continue
		}
		if st.Status == economy.StatusActive {
			st.Status = economy.StatusCondemned
			_ = h.db.UpdateStructure(&st) // lapsed_at stamped by UpdateStructure
			continue
		}
		if st.Status == economy.StatusCondemned {
			lapsed, err := h.db.StructureLapsedAt(st.ID)
			if err == nil && now.Sub(lapsed) >= economy.HouseGracePeriod {
				st.Status = economy.StatusDestroyed
				_ = h.db.UpdateStructure(&st)
			}
		}
	}
}

// tickVendor applies vendor upkeep (base + per active listing); lapse closes.
// Phase 9: Merchant Vendor Management tiers cut upkeep to −50% at tier IV
// (GDD 8.3.5 "up to 50% reduction" — per-tier curve provisional).
func (h *WorldHandler) tickVendor(st *database.StructureRow) {
	listings, err := h.db.ActiveListings(st.ID)
	if err != nil {
		return
	}
	weekly := economy.VendorUpkeepBase + economy.VendorUpkeepPerListing*len(listings)
	weekly = merchantDiscountedUpkeep(weekly,
		h.skillTier(st.OwnerCharacterID, "merchant_vendor_management_"))
	fee := (weekly + 167) / 168
	if st.MaintenancePool >= fee && fee > 0 {
		st.MaintenancePool -= fee
		_ = h.db.RecordLedger(st.OwnerCharacterID, -fee, "maintenance_fee", "sink", st.ID, st.Zone)
		if st.Status == economy.StatusClosed {
			st.Status = economy.StatusActive // ref ud: reopen on funding
		}
		_ = h.db.UpdateStructure(st)
		return
	}
	if st.Status == economy.StatusActive {
		st.Status = economy.StatusClosed // stock safe (GDD 21.2.2)
		_ = h.db.UpdateStructure(st)
	}
}

// --- Client message handlers (positional WS actions) ---

func (h *WorldHandler) handlePlaceHouse(client *Client, raw []byte) {
	if client.CharacterID == "" {
		h.sendError(client, "not in world")
		return
	}
	var msg protocol.WSMessage
	json.Unmarshal(raw, &msg)
	dataBytes, _ := json.Marshal(msg.Data)
	var pm protocol.PlaceStructureMsg
	if err := json.Unmarshal(dataBytes, &pm); err != nil {
		h.sendError(client, "invalid place_house data")
		return
	}
	if pm.Tier != "small" && pm.Tier != "medium" && pm.Tier != "large" {
		h.sendError(client, "unknown house tier (small/medium/large)")
		return
	}
	deed, err := h.db.GetItem(client.CharacterID, pm.DeedItemID)
	if err != nil || deed.Schematic != "structure_deed" {
		h.sendError(client, "valid structure deed required")
		return
	}
	upkeep, ok := economy.HouseUpkeepWeekly[pm.Tier]
	if !ok {
		h.sendError(client, "unknown house tier")
		return
	}
	if err := h.checkNoBuild(client.Pos.Planet, pm.X, pm.Z); err != nil {
		h.sendError(client, err.Error())
		return
	}
	if err := h.checkZoning(client, pm.X, pm.Z); err != nil {
		h.sendError(client, err.Error())
		return
	}
	// Phase 7: unique IDs (one character may own many structures — the fixed
	// per-character ID was a simplification; corrected, recorded in hygiene).
	id := h.db.NewRowID("st")
	if err := h.db.PlaceStructure(id, client.CharacterID, client.Pos.Planet,
		pm.X, pm.Z, "house", pm.Tier, upkeep); err != nil {
		h.sendError(client, "house placement failed")
		return
	}
	_ = h.db.DeleteItem(client.CharacterID, pm.DeedItemID)
	// Phase 9: structure_crafting XP for Architect progression (gate-add,
	// flagged — deed schematics stay artisan-gated with generic crafting XP).
	_ = h.db.AddCharacterXP(client.CharacterID, "structure_crafting", 100)
	h.send(client, protocol.MsgStructurePlaced, protocol.StructurePlacedMsg{
		StructureID: id, Kind: "house",
	})
}

func (h *WorldHandler) handlePlaceVendor(client *Client, raw []byte) {
	if client.CharacterID == "" {
		h.sendError(client, "not in world")
		return
	}
	var msg protocol.WSMessage
	json.Unmarshal(raw, &msg)
	dataBytes, _ := json.Marshal(msg.Data)
	var pm protocol.PlaceStructureMsg
	if err := json.Unmarshal(dataBytes, &pm); err != nil {
		h.sendError(client, "invalid place_vendor data")
		return
	}
	deed, err := h.db.GetItem(client.CharacterID, pm.DeedItemID)
	if err != nil || deed.Schematic != "vendor_deed" {
		h.sendError(client, "valid vendor deed required")
		return
	}
	if err := h.checkNoBuild(client.Pos.Planet, pm.X, pm.Z); err != nil {
		h.sendError(client, err.Error())
		return
	}
	if err := h.checkZoning(client, pm.X, pm.Z); err != nil {
		h.sendError(client, err.Error())
		return
	}
	id := h.db.NewRowID("st")
	if err := h.db.PlaceStructure(id, client.CharacterID, client.Pos.Planet,
		pm.X, pm.Z, "vendor", "", 0); err != nil {
		h.sendError(client, "vendor placement failed")
		return
	}
	_ = h.db.DeleteItem(client.CharacterID, pm.DeedItemID)
	_ = h.db.AddCharacterXP(client.CharacterID, "structure_crafting", 100)
	h.send(client, protocol.MsgStructurePlaced, protocol.StructurePlacedMsg{
		StructureID: id, Kind: "vendor",
	})
}

// merchantDiscountedUpkeep applies the Vendor Management fee curve: −12.5%
// per tier to a −50% floor at tier IV (GDD 8.3.5 "up to 50%"; curve
// provisional). Tiers clamp to [0,4].
func merchantDiscountedUpkeep(weekly, tier int) int {
	if tier < 0 {
		tier = 0
	}
	if tier > 4 {
		tier = 4
	}
	return weekly * (8 - tier) / 8
}

// checkZoning enforces the mayor's placement gate: inside an active city with
// closed zoning, only the mayor may place (GDD 14.2.3 approve/deny shape as a
// binary gate — flagged simplification, no per-request queue).
func (h *WorldHandler) checkZoning(client *Client, x, z float64) error {
	c, err := h.db.CityContaining(client.Pos.Planet, x, z)
	if err != nil || c.ZoningOpen {
		return nil
	}
	if c.MayorID.Valid && c.MayorID.String == client.CharacterID {
		return nil
	}
	return fmt.Errorf("city zoning closed by mayor")
}

// checkNoBuild enforces separation between structures [PROVISIONAL 20 m radius;
// GDD gives 100 m for harvesters only].
func (h *WorldHandler) checkNoBuild(zone string, x, z float64) error {
	all, err := h.db.AllStructures()
	if err != nil {
		return err
	}
	for _, s := range all {
		if s.Zone != zone || s.Status == economy.StatusDestroyed {
			continue
		}
		dx, dz := x-s.PosX, z-s.PosZ
		if dx*dx+dz*dz < economy.NoBuildRadiusM*economy.NoBuildRadiusM {
			return fmt.Errorf("too close to another structure")
		}
	}
	return nil
}

func (h *WorldHandler) handlePurchase(client *Client, raw []byte) {
	if client.CharacterID == "" {
		h.sendError(client, "not in world")
		return
	}
	var msg protocol.WSMessage
	json.Unmarshal(raw, &msg)
	dataBytes, _ := json.Marshal(msg.Data)
	var pm protocol.PurchaseMsg
	if err := json.Unmarshal(dataBytes, &pm); err != nil {
		h.sendError(client, "invalid purchase data")
		return
	}
	listing, err := h.db.GetListing(pm.ListingID)
	if err != nil {
		h.sendError(client, "listing unknown")
		return
	}
	vendor, err := h.db.GetStructure(listing.VendorID)
	if err != nil {
		h.sendError(client, "vendor unknown")
		return
	}
	if vendor.Zone != client.Pos.Planet {
		h.sendError(client, "vendor not in this zone")
		return
	}
	// Phase 7: banned-list entry denial (GDD 13.2.4; without interior cells
	// the enforced analog is refusing the banned buyer's purchase — flagged).
	if h.db.IsPermListed(vendor.ID, "banned", client.CharacterID) {
		h.sendError(client, "entry denied by structure owner")
		return
	}
	dx := client.Pos.X - vendor.PosX
	dz := client.Pos.Z - vendor.PosZ
	if dx*dx+dz*dz > economy.PurchaseRangeM*economy.PurchaseRangeM {
		h.sendError(client, "too far from vendor (walk up to buy)")
		return
	}
	res, err := h.db.AtomicPurchase(client.CharacterID, pm.ListingID, client.Pos.Planet)
	if err != nil {
		h.sendError(client, err.Error()) // sold out / closed / insufficient — never partial
		return
	}
	_ = res
	// Phase 9 Merchant XP (GDD 7.2.5: 1 XP per 100 credits; unique buyers
	// full, repeats diminished to 25% — provisional): post-tx best-effort
	// (tip-XP precedent).
	xp := listing.Price / 100
	if xp < 1 {
		xp = 1
	}
	if fresh, _ := h.db.RecordCustomer(listing.VendorID, client.CharacterID); !fresh {
		xp /= 4
	}
	_ = h.db.AddCharacterXP(vendor.OwnerCharacterID, "merchant", xp)
	h.send(client, protocol.MsgPurchaseReceipt, protocol.PurchaseReceiptMsg{
		ListingID: pm.ListingID, ItemID: listing.ItemID,
	})
}

func (h *WorldHandler) handleCollectTill(client *Client, raw []byte) {
	if client.CharacterID == "" {
		h.sendError(client, "not in world")
		return
	}
	var msg protocol.WSMessage
	json.Unmarshal(raw, &msg)
	dataBytes, _ := json.Marshal(msg.Data)
	var tm protocol.TillMsg
	if err := json.Unmarshal(dataBytes, &tm); err != nil {
		h.sendError(client, "invalid collect_till data")
		return
	}
	vendor, err := h.db.GetStructure(tm.VendorID)
	if err != nil || vendor.Kind != "vendor" {
		h.sendError(client, "vendor unknown")
		return
	}
	if vendor.OwnerCharacterID != client.CharacterID {
		h.sendError(client, "not vendor owner")
		return
	}
	dx := client.Pos.X - vendor.PosX
	dz := client.Pos.Z - vendor.PosZ
	if dx*dx+dz*dz > economy.CollectRangeM*economy.CollectRangeM {
		h.sendError(client, "too far from vendor (collect in person)")
		return
	}
	amount, err := h.db.CollectTill(tm.VendorID, client.CharacterID)
	if err != nil {
		h.sendError(client, err.Error())
		return
	}
	h.send(client, protocol.MsgTillCollected, protocol.TillCollectedMsg{
		VendorID: tm.VendorID, Amount: amount,
	})
}
