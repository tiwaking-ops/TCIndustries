package economy

import (
	"strings"
)

// BazaarListing represents a searchable vendor item entry in the galaxy-wide bazaar index (GDD 22).
type BazaarListing struct {
	VendorID    string `json:"vendor_id"`
	VendorName  string `json:"vendor_name"`
	ItemID      string `json:"item_id"`
	ItemName    string `json:"item_name"`
	Price       int    `json:"price"`
	Description string `json:"description"`
}

// BazaarSearchCriteria filters bazaar queries.
type BazaarSearchCriteria struct {
	Query    string // name filter
	MaxPrice int    // max price filter (0 if unconstrained)
	SellerID string // seller filter
}

// BazaarIndex aggregates live listings across all player vendors.
type BazaarIndex struct {
	SalesHistory []BazaarSale `json:"sales_history"` // last-N-sales view (GDD 22.2.4)
}

type BazaarSale struct {
	ItemName   string `json:"item_name"`
	Price      int    `json:"price"`
	BuyerID    string `json:"buyer_id"`
	SellerID   string `json:"seller_id"`
	Timestamp  string `json:"timestamp"`
}

func NewBazaarIndex() *BazaarIndex {
	return &BazaarIndex{
		SalesHistory: make([]BazaarSale, 0),
	}
}

// Search queries all active vendor stock against criteria.
func (bi *BazaarIndex) Search(vendors map[string]*VendorInstance, criteria BazaarSearchCriteria) []BazaarListing {
	var results []BazaarListing

	for _, v := range vendors {
		for _, item := range v.Stock {
			// Apply filters
			if criteria.Query != "" && !strings.Contains(strings.ToLower(item.ItemName), strings.ToLower(criteria.Query)) {
				continue
			}
			if criteria.MaxPrice > 0 && item.Price > criteria.MaxPrice {
				continue
			}
			if criteria.SellerID != "" && v.OwnerID != criteria.SellerID {
				continue
			}

			results = append(results, BazaarListing{
				VendorID:    v.ID,
				VendorName:  v.Name,
				ItemID:      item.ItemID,
				ItemName:    item.ItemName,
				Price:       item.Price,
				Description: item.Description,
			})
		}
	}

	return results
}
