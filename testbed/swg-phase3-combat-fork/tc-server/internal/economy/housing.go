package economy

import (
	"errors"
	"math"
	"time"
)

// HouseInstance represents a placed player structure (GDD 13).
type HouseInstance struct {
	ID          string    `json:"id"`
	OwnerID     string    `json:"owner_id"`
	TypeName    string    `json:"type_name"` // e.g. "small_house"
	X           float64   `json:"x"`
	Z           float64   `json:"z"`
	Maintenance int       `json:"maintenance_pool"` // credits remaining
	LastUpkeep  time.Time `json:"last_upkeep"`
	Condemned   bool      `json:"condemned"`
}

// VendorInstance represents a placed player vendor (GDD 21).
type VendorInstance struct {
	ID       string             `json:"id"`
	OwnerID  string             `json:"owner_id"`
	Name     string             `json:"name"`
	X        float64            `json:"x"`
	Z        float64            `json:"z"`
	Till     int                `json:"till"` // uncollected sales credits
	Stock    map[string]VendorItem // itemID -> item details
}

// VendorItem represents an item stocked on a player vendor.
type VendorItem struct {
	ItemID      string  `json:"item_id"`
	ItemName    string  `json:"item_name"`
	Price       int     `json:"price"`
	Description string  `json:"description"`
}

// EconomyManager manages houses and vendors.
type EconomyManager struct {
	houses  map[string]*HouseInstance
	vendors map[string]*VendorInstance
}

func NewEconomyManager() *EconomyManager {
	return &EconomyManager{
		houses:  make(map[string]*HouseInstance),
		vendors: make(map[string]*VendorInstance),
	}
}

// PlaceHouse validates distance and places a player house.
func (em *EconomyManager) PlaceHouse(id, ownerID, typeName string, x, z float64) (*HouseInstance, error) {
	// Check minimum separation from other houses (20m provisional per GDD proposal)
	for _, h := range em.houses {
		dx := h.X - x
		dz := h.Z - z
		if math.Sqrt(dx*dx+dz*dz) < 20.0 {
			return nil, errors.New("too close to another structure (20m separation required)")
		}
	}

	h := &HouseInstance{
		ID:          id,
		OwnerID:     ownerID,
		TypeName:    typeName,
		X:           x,
		Z:           z,
		Maintenance: 1000, // starting maintenance pool
		LastUpkeep:  time.Now(),
		Condemned:   false,
	}

	em.houses[id] = h
	return h, nil
}

// PlaceVendor places a player vendor.
func (em *EconomyManager) PlaceVendor(id, ownerID, name string, x, z float64) (*VendorInstance, error) {
	v := &VendorInstance{
		ID:      id,
		OwnerID: ownerID,
		Name:    name,
		X:       x,
		Z:       z,
		Till:    0,
		Stock:   make(map[string]VendorItem),
	}

	em.vendors[id] = v
	return v, nil
}

// StockVendor adds an item to a vendor (max 50 slots per GDD [ASSUMPTION]).
func (v *VendorInstance) StockVendor(itemID, itemName, description string, price int) error {
	if len(v.Stock) >= 50 {
		return errors.New("vendor stock is full (50 slot limit)")
	}
	if price <= 0 {
		return errors.New("price must be positive")
	}

	v.Stock[itemID] = VendorItem{
		ItemID:      itemID,
		ItemName:    itemName,
		Price:       price,
		Description: description,
	}
	return nil
}
