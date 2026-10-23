package resources

import (
	"errors"
)

// ResourceStack represents a stack of harvested resources keyed by spawn ID (GDD 11.5).
type ResourceStack struct {
	SpawnID      string `json:"spawn_id"`
	ResourceName string `json:"resource_name"`
	Category     string `json:"category"`
	Quantity     int    `json:"quantity"`
	Quality      int    `json:"quality"`
}

// CharacterInventory manages resource stacks for a player.
type CharacterInventory struct {
	CharacterID string                   `json:"character_id"`
	Stacks      map[string]*ResourceStack // spawnID -> stack
}

func NewCharacterInventory(characterID string) *CharacterInventory {
	return &CharacterInventory{
		CharacterID: characterID,
		Stacks:      make(map[string]*ResourceStack),
	}
}

// AddStack adds harvested resources to inventory, stacking by spawn ID.
func (inv *CharacterInventory) AddStack(spawnID, name, category string, qty, quality int) {
	if qty <= 0 {
		return
	}
	stack, ok := inv.Stacks[spawnID]
	if !ok {
		inv.Stacks[spawnID] = &ResourceStack{
			SpawnID:      spawnID,
			ResourceName: name,
			Category:     category,
			Quantity:     qty,
			Quality:      quality,
		}
		return
	}

	// Weighted mean quality when merging stacks (or keep existing)
	totalQty := stack.Quantity + qty
	stack.Quality = ((stack.Quality * stack.Quantity) + (quality * qty)) / totalQty
	stack.Quantity = totalQty
}

// RemoveStack removes quantity from a specific spawn ID stack.
func (inv *CharacterInventory) RemoveStack(spawnID string, qty int) error {
	stack, ok := inv.Stacks[spawnID]
	if !ok {
		return errors.New("resource stack not found")
	}
	if stack.Quantity < qty {
		return errors.New("insufficient resource quantity")
	}
	stack.Quantity -= qty
	if stack.Quantity == 0 {
		delete(inv.Stacks, spawnID)
	}
	return nil
}
