package pvp

// Battlefield zone geometry (R-08: fixed rectangles or circles seeded in the
// database). Shapes map directly to the database columns; the pure geometry
// lives here so containment tests need no DB.
type ZoneShape string

const (
	ZoneCircle ZoneShape = "circle"
	ZoneRect   ZoneShape = "rect"
)

// Zone is a battlefield zone in the X/Z plane, matching structure and city
// coordinates. For a circle only Center/Radius are meaningful; for a rect only
// Min/Max are.
type Zone struct {
	ID     string
	Shape  ZoneShape
	CenterX float64
	CenterZ float64
	Radius  float64
	MinX    float64
	MinZ    float64
	MaxX    float64
	MaxZ    float64
	Active  bool
}

// Contains reports whether the point (x, zPos) lies inside the zone.
func (z Zone) Contains(x, zPos float64) bool {
	if !z.Active {
		return false
	}
	switch z.Shape {
	case ZoneCircle:
		dx := x - z.CenterX
		dz := zPos - z.CenterZ
		return dx*dx+dz*dz <= z.Radius*z.Radius
	case ZoneRect:
		return x >= z.MinX && x <= z.MaxX && zPos >= z.MinZ && zPos <= z.MaxZ
	}
	return false
}

// ZoneContaining returns the first active zone containing (x, zPos), or "" when
// none apply. Overlapping zones resolve to the first match in the supplied
// order (proposal Section 9: documented choice).
func ZoneContaining(zones []Zone, x, zPos float64) string {
	for _, zone := range zones {
		if zone.Contains(x, zPos) {
			return zone.ID
		}
	}
	return ""
}
