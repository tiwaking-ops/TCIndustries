// Economy HTTP handlers for Phase 5: vendor listing management, bazaar search,
// structure funding, ledger queries, and telemetry snapshots. Testbed fork;
// generic content only.
package handlers

import (
	"net/http"
	"strconv"
	"strings"

	"swg-server/internal/crafting"
	"swg-server/internal/database"
	"swg-server/internal/economy"
)

// EconomyHandler serves /api/economy/* (AuthMiddleware-wrapped, except nothing;
// all routes require auth like the craft routes).
type EconomyHandler struct {
	db *database.DB
}

// NewEconomyHandler creates the handler.
func NewEconomyHandler(db *database.DB) *EconomyHandler {
	return &EconomyHandler{db: db}
}

// HandleEconomyRoute dispatches /api/economy/* sub-paths.
func (eh *EconomyHandler) HandleEconomyRoute(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimPrefix(r.URL.Path, "/api/economy/")
	charID := r.URL.Query().Get("character_id")
	switch {
	case path == "vendor/list" && r.Method == "POST":
		eh.stock(w, r)
	case path == "vendor/reprice" && r.Method == "POST":
		eh.reprice(w, r)
	case path == "vendor/pull" && r.Method == "POST":
		eh.pull(w, r)
	case path == "vendor/listings" && r.Method == "GET":
		eh.vendorListings(w, r, r.URL.Query().Get("vendor_id"))
	case path == "bazaar/search" && r.Method == "GET":
		eh.search(w, r, charID)
	case path == "structure/fund" && r.Method == "POST":
		eh.fund(w, r)
	case path == "ledger" && r.Method == "GET":
		eh.ledger(w, r, charID)
	case path == "structures" && r.Method == "GET":
		eh.structures(w, r, charID)
	case path == "snapshot" && r.Method == "GET":
		eh.snapshot(w, r)
	default:
		writeCraftJSON(w, http.StatusNotFound, map[string]string{"error": "unknown economy route"})
	}
}

func (eh *EconomyHandler) ownVendor(w http.ResponseWriter, charID, vendorID string) (*database.StructureRow, bool) {
	v, err := eh.db.GetStructure(vendorID)
	if err != nil || v.Kind != "vendor" {
		writeCraftJSON(w, http.StatusNotFound, map[string]string{"error": "vendor unknown"})
		return nil, false
	}
	if v.OwnerCharacterID != charID {
		writeCraftJSON(w, http.StatusForbidden, map[string]string{"error": "not vendor owner"})
		return nil, false
	}
	return v, true
}

// stock lists an owned, unequipped crafted item on the owner's vendor.
func (eh *EconomyHandler) stock(w http.ResponseWriter, r *http.Request) {
	body, ok := decodeBody(w, r)
	if !ok {
		return
	}
	charID := strOf(body, "character_id")
	v, ok := eh.ownVendor(w, charID, strOf(body, "vendor_id"))
	if !ok {
		return
	}
	if v.Status != economy.StatusActive {
		writeCraftJSON(w, http.StatusBadRequest, map[string]string{"error": "vendor not open"})
		return
	}
	item, err := eh.db.GetItem(charID, strOf(body, "item_id"))
	if err != nil {
		writeCraftJSON(w, http.StatusNotFound, map[string]string{"error": "item not found"})
		return
	}
	if item.Equipped {
		writeCraftJSON(w, http.StatusBadRequest, map[string]string{"error": "cannot list equipped item"})
		return
	}
	live, err := eh.db.ActiveListings(v.ID)
	if err != nil {
		writeCraftJSON(w, http.StatusInternalServerError, map[string]string{"error": "listing lookup failed"})
		return
	}
	if len(live) >= economy.VendorSlotCap {
		writeCraftJSON(w, http.StatusBadRequest, map[string]string{"error": "vendor inventory full"})
		return
	}
	price := 0
	if f, ok := body["price"].(float64); ok {
		price = int(f)
	}
	if price <= 0 {
		writeCraftJSON(w, http.StatusBadRequest, map[string]string{"error": "price must be positive"})
		return
	}
	id, err := eh.db.CreateListing(v.ID, item.ID, price, strOf(body, "description"))
	if err != nil {
		writeCraftJSON(w, http.StatusInternalServerError, map[string]string{"error": "listing failed"})
		return
	}
	writeCraftJSON(w, http.StatusCreated, map[string]string{"listing_id": id})
}

// reprice changes a live listing's price (owner only).
func (eh *EconomyHandler) reprice(w http.ResponseWriter, r *http.Request) {
	body, ok := decodeBody(w, r)
	if !ok {
		return
	}
	charID := strOf(body, "character_id")
	l, err := eh.db.GetListing(strOf(body, "listing_id"))
	if err != nil {
		writeCraftJSON(w, http.StatusNotFound, map[string]string{"error": "listing unknown"})
		return
	}
	if _, ok := eh.ownVendor(w, charID, l.VendorID); !ok {
		return
	}
	price := 0
	if f, ok := body["price"].(float64); ok {
		price = int(f)
	}
	if price <= 0 {
		writeCraftJSON(w, http.StatusBadRequest, map[string]string{"error": "price must be positive"})
		return
	}
	if err := eh.db.RepriceListing(l.ID, price); err != nil {
		writeCraftJSON(w, http.StatusInternalServerError, map[string]string{"error": "reprice failed"})
		return
	}
	writeCraftJSON(w, http.StatusOK, map[string]string{"status": "repriced"})
}

// pull delists a live listing (item stays with the owner).
func (eh *EconomyHandler) pull(w http.ResponseWriter, r *http.Request) {
	body, ok := decodeBody(w, r)
	if !ok {
		return
	}
	charID := strOf(body, "character_id")
	l, err := eh.db.GetListing(strOf(body, "listing_id"))
	if err != nil {
		writeCraftJSON(w, http.StatusNotFound, map[string]string{"error": "listing unknown"})
		return
	}
	if _, ok := eh.ownVendor(w, charID, l.VendorID); !ok {
		return
	}
	if err := eh.db.SetListingActive(l.ID, false); err != nil {
		writeCraftJSON(w, http.StatusInternalServerError, map[string]string{"error": "pull failed"})
		return
	}
	writeCraftJSON(w, http.StatusOK, map[string]string{"status": "pulled"})
}

// vendorListings returns live listings for one vendor (public browse shape).
func (eh *EconomyHandler) vendorListings(w http.ResponseWriter, r *http.Request, vendorID string) {
	if vendorID == "" {
		writeCraftJSON(w, http.StatusBadRequest, map[string]string{"error": "vendor_id required"})
		return
	}
	list, err := eh.db.ActiveListings(vendorID)
	if err != nil {
		writeCraftJSON(w, http.StatusInternalServerError, map[string]string{"error": "lookup failed"})
		return
	}
	writeCraftJSON(w, http.StatusOK, map[string]interface{}{"listings": list})
}

// search implements bazaar search over live vendor listings. The searcher must
// be near a terminal (provisional 30 m rule); filters: slot (weapon/armor),
// min/max price, seller name, and schematic id. Returns live listings with
// vendor location plus last-5 sales for matched schematics (MVP market data).
func (eh *EconomyHandler) search(w http.ResponseWriter, r *http.Request, charID string) {
	if charID == "" {
		writeCraftJSON(w, http.StatusBadRequest, map[string]string{"error": "character_id required"})
		return
	}
	char, err := eh.db.GetCharacterByID(charID)
	if err != nil {
		writeCraftJSON(w, http.StatusNotFound, map[string]string{"error": "character unknown"})
		return
	}
	terms, err := eh.db.Terminals(char.Planet)
	if err != nil || len(terms) == 0 {
		writeCraftJSON(w, http.StatusBadRequest, map[string]string{"error": "no terminal in zone"})
		return
	}
	near := false
	for _, t := range terms {
		dx, dz := char.PosX-t.X, char.PosZ-t.Z
		if dx*dx+dz*dz <= economy.TerminalSearchRadiusM*economy.TerminalSearchRadiusM {
			near = true
		}
	}
	if !near {
		writeCraftJSON(w, http.StatusBadRequest, map[string]string{"error": "no bazaar terminal in range"})
		return
	}
	q := r.URL.Query()
	slot, schem := q.Get("slot"), q.Get("schematic")
	minP, _ := strconv.Atoi(q.Get("min_price"))
	maxP, _ := strconv.Atoi(q.Get("max_price"))
	seller := q.Get("seller")
	list, err := eh.db.ActiveListings("")
	if err != nil {
		writeCraftJSON(w, http.StatusInternalServerError, map[string]string{"error": "search failed"})
		return
	}
	type hit struct {
		ListingID string  `json:"listing_id"`
		VendorID  string  `json:"vendor_id"`
		ItemID    string  `json:"item_id"`
		Schematic string  `json:"schematic"`
		Price     int     `json:"price"`
		Seller    string  `json:"seller"`
		Zone      string  `json:"zone"`
		VendorX   float64 `json:"vendor_x"`
		VendorZ   float64 `json:"vendor_z"`
	}
	hits := []hit{}
	seenSchem := map[string]bool{}
	for _, l := range list {
		v, err := eh.db.GetStructure(l.VendorID)
		if err != nil || v.Status != economy.StatusActive {
			continue
		}
		item, err := eh.db.GetItem(v.OwnerCharacterID, l.ItemID)
		if err != nil {
			continue
		}
		sc := crafting.SchematicByID(item.Schematic)
		if sc == nil {
			continue
		}
		if slot != "" && sc.EquipSlot != slot {
			continue
		}
		if schem != "" && item.Schematic != schem {
			continue
		}
		if minP > 0 && l.Price < minP {
			continue
		}
		if maxP > 0 && l.Price > maxP {
			continue
		}
		sellerName := eh.sellerName(v.OwnerCharacterID)
		if seller != "" && sellerName != seller {
			continue
		}
		hits = append(hits, hit{ListingID: l.ID, VendorID: v.ID, ItemID: item.ID,
			Schematic: item.Schematic, Price: l.Price, Seller: sellerName,
			Zone: v.Zone, VendorX: v.PosX, VendorZ: v.PosZ})
		seenSchem[item.Schematic] = true
	}
	history := []database.MarketRecord{}
	for s := range seenSchem {
		recs, err := eh.db.RecentSales(s, 5)
		if err == nil {
			history = append(history, recs...)
		}
	}
	writeCraftJSON(w, http.StatusOK, map[string]interface{}{
		"results": historyKeeper(hits), "recent_sales": history,
	})
}

func historyKeeper[T any](in []T) []T {
	if in == nil {
		return []T{}
	}
	return in
}

func (eh *EconomyHandler) sellerName(charID string) string {
	c, err := eh.db.GetCharacterByID(charID)
	if err != nil {
		return ""
	}
	return c.Name
}

// fund moves wallet credits into a structure pool (transfer, ledger-tagged).
func (eh *EconomyHandler) fund(w http.ResponseWriter, r *http.Request) {
	body, ok := decodeBody(w, r)
	if !ok {
		return
	}
	charID, stID := strOf(body, "character_id"), strOf(body, "structure_id")
	st, err := eh.db.GetStructure(stID)
	if err != nil {
		writeCraftJSON(w, http.StatusNotFound, map[string]string{"error": "structure unknown"})
		return
	}
	// Phase 7: owner-or-admin may fund (GDD 13.2.4 admin manages; redeed and
	// ownership transfer stay owner-only elsewhere).
	if !eh.db.CanManageStructure(stID, charID) {
		writeCraftJSON(w, http.StatusForbidden, map[string]string{"error": "not structure owner or admin"})
		return
	}
	credits := 0
	if f, ok := body["credits"].(float64); ok {
		credits = int(f)
	}
	if credits <= 0 {
		writeCraftJSON(w, http.StatusBadRequest, map[string]string{"error": "credits must be positive"})
		return
	}
	if err := eh.db.DeductCredits(charID, credits); err != nil {
		writeCraftJSON(w, http.StatusBadRequest, map[string]string{"error": "wallet deduction failed: "+err.Error()})
		return
	}
	if err := eh.db.FundStructure(stID, credits); err != nil {
		writeCraftJSON(w, http.StatusInternalServerError, map[string]string{"error": "funding failed"})
		return
	}
	_ = eh.db.RecordLedger(charID, -credits, "structure_fund", "transfer", stID, st.Zone)
	writeCraftJSON(w, http.StatusOK, map[string]string{"status": "funded"})
}

// ledger returns a character's ledger entries.
func (eh *EconomyHandler) ledger(w http.ResponseWriter, r *http.Request, charID string) {
	if charID == "" {
		writeCraftJSON(w, http.StatusBadRequest, map[string]string{"error": "character_id required"})
		return
	}
	entries, err := eh.db.LedgerFor(charID)
	if err != nil {
		writeCraftJSON(w, http.StatusInternalServerError, map[string]string{"error": "ledger lookup failed"})
		return
	}
	if entries == nil {
		entries = []database.LedgerEntry{}
	}
	writeCraftJSON(w, http.StatusOK, map[string]interface{}{"entries": entries})
}

// snapshot writes and returns a fresh telemetry snapshot (GDD 12.4/12.6 shape).
func (eh *EconomyHandler) snapshot(w http.ResponseWriter, r *http.Request) {
	snap, err := eh.db.WriteSnapshot()
	if err != nil {
		writeCraftJSON(w, http.StatusInternalServerError, map[string]string{"error": "snapshot failed"})
		return
	}
	writeCraftJSON(w, http.StatusOK, map[string]interface{}{"snapshot": snap})
}

// structures returns a character's owned structures with pool/status (test +
// management introspection aid).
func (eh *EconomyHandler) structures(w http.ResponseWriter, r *http.Request, charID string) {
	if charID == "" {
		writeCraftJSON(w, http.StatusBadRequest, map[string]string{"error": "character_id required"})
		return
	}
	all, err := eh.db.AllStructures()
	if err != nil {
		writeCraftJSON(w, http.StatusInternalServerError, map[string]string{"error": "lookup failed"})
		return
	}
	type view struct {
		ID     string `json:"id"`
		Kind   string `json:"kind"`
		Status string `json:"status"`
		Pool   int    `json:"pool"`
		Till   int    `json:"till"`
	}
	out := []view{}
	for _, s := range all {
		if s.OwnerCharacterID != charID {
			continue
		}
		till, _ := eh.db.VendorTill(s.ID)
		out = append(out, view{ID: s.ID, Kind: s.Kind, Status: s.Status, Pool: s.MaintenancePool, Till: till})
	}
	writeCraftJSON(w, http.StatusOK, map[string]interface{}{"structures": out})
}
