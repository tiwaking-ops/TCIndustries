package world

import (
	"math"
	"sync"
)

// Position represents a 3D world position on a specific planet.
// Y is altitude (up). X and Z are the horizontal plane used for proximity.
type Position struct {
	Planet string
	X      float64
	Y      float64
	Z      float64
}

// Distance2D returns the horizontal distance between two positions on the same planet.
// Returns +Inf if they are on different planets. (GDD 29.2.2: per-planet partitioning)
func Distance2D(a, b Position) float64 {
	if a.Planet != b.Planet {
		return math.Inf(1)
	}
	dx := a.X - b.X
	dz := a.Z - b.Z
	return math.Sqrt(dx*dx + dz*dz)
}

// CellKey identifies a grid cell on a planet.
type CellKey struct {
	Planet string
	CX     int // cell X coordinate = floor(worldX / cellSize)
	CZ     int // cell Z coordinate = floor(worldZ / cellSize)
}

// Entity represents a tracked world entity (player character) in the spatial grid.
type Entity struct {
	CharacterID string
	Name        string
	Species     string
	Pos         Position
	Heading     float64
}

// Grid is a spatial partitioning grid for interest management.
// It partitions the world into fixed-size cells per planet, allowing O(1)
// lookup of which cell an entity is in, and O(n) query of nearby entities
// by checking only the neighboring cells within the query radius.
//
// GDD Section 29.2.2: "World space is partitioned per-planet, then into a
// grid of zones/cells for interest-management (only sync entities near a
// player to that player's client)."
type Grid struct {
	mu       sync.RWMutex
	cellSize float64
	cells    map[CellKey]map[string]*Entity // CellKey → characterID → Entity
}

// NewGrid creates a new spatial grid with the given cell size in meters.
func NewGrid(cellSize float64) *Grid {
	return &Grid{
		cellSize: cellSize,
		cells:    make(map[CellKey]map[string]*Entity),
	}
}

// cellKey converts a world position to a cell key.
func (g *Grid) cellKey(pos Position) CellKey {
	return CellKey{
		Planet: pos.Planet,
		CX:     int(math.Floor(pos.X / g.cellSize)),
		CZ:     int(math.Floor(pos.Z / g.cellSize)),
	}
}

// AddEntity adds or updates an entity in the grid. If the entity already exists,
// it is removed from its old cell and placed in the new one.
func (g *Grid) AddEntity(e *Entity) {
	g.mu.Lock()
	defer g.mu.Unlock()

	key := g.cellKey(e.Pos)
	cell, ok := g.cells[key]
	if !ok {
		cell = make(map[string]*Entity)
		g.cells[key] = cell
	}
	cell[e.CharacterID] = e
}

// RemoveEntity removes an entity from the grid.
func (g *Grid) RemoveEntity(characterID string, pos Position) {
	g.mu.Lock()
	defer g.mu.Unlock()

	key := g.cellKey(pos)
	if cell, ok := g.cells[key]; ok {
		delete(cell, characterID)
		if len(cell) == 0 {
			delete(g.cells, key)
		}
	}
}

// MoveEntity updates an entity's position. If the entity moved to a new cell,
// it is removed from the old cell and added to the new one.
// Returns the old position so callers can detect cell changes.
func (g *Grid) MoveEntity(characterID string, oldPos, newPos Position, heading float64) (oldCellKey, newCellKey CellKey, cellChanged bool) {
	g.mu.Lock()
	defer g.mu.Unlock()

	oldKey := g.cellKey(oldPos)
	newKey := g.cellKey(newPos)

	// Remove from old cell
	if oldCell, ok := g.cells[oldKey]; ok {
		if e, exists := oldCell[characterID]; exists {
			if oldKey == newKey {
				// Same cell, just update position and heading
				e.Pos = newPos
				e.Heading = heading
				return oldKey, newKey, false
			}
			// Different cell — remove from old
			delete(oldCell, characterID)
			if len(oldCell) == 0 {
				delete(g.cells, oldKey)
			}
		}
	}

	// Add to new cell, preserving existing entity data if available
	newCell, ok := g.cells[newKey]
	if !ok {
		newCell = make(map[string]*Entity)
		g.cells[newKey] = newCell
	}
	// Try to preserve Name/Species from the old entity
	name := ""
	species := ""
	if oldCell, ok := g.cells[oldKey]; ok {
		if oldE, exists := oldCell[characterID]; exists {
			name = oldE.Name
			species = oldE.Species
		}
	}
	newCell[characterID] = &Entity{
		CharacterID: characterID,
		Name:        name,
		Species:     species,
		Pos:         newPos,
		Heading:     heading,
	}

	return oldKey, newKey, true
}

// UpdateEntityInfo updates the name and species of an entity in-place
// without moving it between cells. Used when the entity first enters the world.
func (g *Grid) UpdateEntityInfo(characterID string, name, species string, pos Position) {
	g.mu.Lock()
	defer g.mu.Unlock()

	key := g.cellKey(pos)
	if cell, ok := g.cells[key]; ok {
		if e, exists := cell[characterID]; exists {
			e.Name = name
			e.Species = species
		}
	}
}

// Nearby returns all entities within the given radius of the query position.
// It checks the query cell and all neighboring cells, then filters by exact
// 2D distance. The calling entity is excluded from the results.
//
// GDD Section 29.2.2: "only sync entities near a player to that player's client"
func (g *Grid) Nearby(pos Position, radius float64, excludeID string) []*Entity {
	g.mu.RLock()
	defer g.mu.RUnlock()

	var result []*Entity

	// Calculate how many cells we need to check in each direction
	cellRadius := int(math.Ceil(radius / g.cellSize))
	centerKey := g.cellKey(pos)

	for cx := centerKey.CX - cellRadius; cx <= centerKey.CX+cellRadius; cx++ {
		for cz := centerKey.CZ - cellRadius; cz <= centerKey.CZ+cellRadius; cz++ {
			key := CellKey{Planet: pos.Planet, CX: cx, CZ: cz}
			cell, ok := g.cells[key]
			if !ok {
				continue
			}
			for id, e := range cell {
				if id == excludeID {
					continue
				}
				if Distance2D(pos, e.Pos) <= radius {
					result = append(result, e)
				}
			}
		}
	}

	return result
}

// EntityCount returns the total number of entities in the grid.
func (g *Grid) EntityCount() int {
	g.mu.RLock()
	defer g.mu.RUnlock()

	count := 0
	for _, cell := range g.cells {
		count += len(cell)
	}
	return count
}
