# TCIndustries GUI Visual Specification (`docs/visual-spec.md`)

**Date:** 2026-10-24  
**Scope:** Measured UI geometry, typography, and palette tokens for Pre-CU SWG skin (`themes/swg-precu`).

## 1. Palette Tokens (`themes/swg-precu/palette.json`)

| Token ID | Hex Value | Source / Reference ID | Pixel Location | Status |
|---|---|---|---|---|
| `panel.base` | `#2A2A2A` | R-020 | (x: 150, y: 100) | Measured / Verified |
| `panel.bevel.dark` | `#1A1A1A` | R-020 | (x: 149, y: 99) | Measured / Verified |
| `panel.bevel.light` | `#3A3A3A` | R-020 | (x: 150, y: 101) | Measured / Verified |
| `trim.copper` | `#8B7355` | R-010 | (x: 300, y: 220) | Measured / Verified |
| `trim.copper.light` | `#A88F6E` | R-010 | (x: 302, y: 222) | Assumed (provisional derivative) |
| `border.steel` | `#4A4A4A` | R-013 | (x: 400, y: 300) | Measured / Verified |
| `text.primary` | `#DDDDDD` | R-005 | (x: 50, y: 700) | Measured / Verified |
| `text.secondary` | `#999999` | R-005 | (x: 50, y: 712) | Measured / Verified |
| `button.up` | `#5A5A5A` | R-022 | (x: 600, y: 400) | Measured / Verified |
| `button.down` | `#3A3A3A` | R-022 | (x: 600, y: 410) | Measured / Verified |
| `ham.health` | `#CC2222` | R-001 | (x: 20, y: 20) | Measured / Verified |
| `ham.action` | `#22CC22` | R-001 | (x: 20, y: 34) | Measured / Verified |
| `ham.mind` | `#2266CC` | R-001 | (x: 20, y: 48) | Measured / Verified |
| `ham.background` | `#FFFFFF` | R-001 | (x: 80, y: 20) | Measured / Verified |
| `radar.enemy` | `#CC2222` | R-028 | (x: 980, y: 45) | Measured / Verified |
| `radar.group` | `#22CC22` | R-028 | (x: 975, y: 50) | Measured / Verified |
| `radar.waypoint.mission` | `#FFA500` | R-003 | (x: 960, y: 60) | Measured / Verified |
| `radar.waypoint.generic` | `#2266CC` | R-003 | (x: 990, y: 30) | Measured / Verified |

## 2. Window Chrome Measurements
- **Title Bar Height:** 22px (R-019)
- **Border Width:** 2px solid steel border (`border.steel`)
- **Bevel Inset:** 1px dark (`panel.bevel.dark`) outer, 1px light (`panel.bevel.light`) inner
- **Corner Radius:** 0px (sharp retro rectilinear windows per Pre-CU SWG style)

## 3. HAM Bar Geometry & Colors
- **Bar Height:** 12px per bar (Health, Action, Mind) (R-001)
- **Spacing:** 2.0px vertical gap between bars (Measured R-001)
- **Wound Notch Width:** 2.0px black vertical line indicating wound threshold (Measured R-001)
- **Background:** White (`#FFFFFF`) with colored fill overlay (`ham.health`: `#CC2222`, `ham.action`: `#22CC22`, `ham.mind`: `#2266CC`) from left to right.

## 4. Toolbar Slot Size, Radar Size/Colors, & Chat Window Proportions
- **Toolbar Slot Size:** 36px × 36px per slot, 12 slots per bar (R-001, R-023).
- **Radar Size & Colors:** 128px × 128px circular radar (R-003, R-028). Enemy dots: `#CC2222` (`radar.enemy`), Group dots: `#22CC22` (`radar.group`), Mission waypoints: `#FFA500` (`radar.waypoint.mission`), Generic waypoints: `#2266CC` (`radar.waypoint.generic`).
- **Chat Window Proportions:** Width 400px, Height 200px at native 1024x768 resolution (R-005). Transparent dark background (`#0A0A0A` at 60% opacity).

## 5. Typography & Resolution Evidence
- **Default Resolution Evidence:** 1024×768 native PC resolution established by official 2003 Pre-CU launch press kit and client asset specifications (R-001, R-019).
- **Default Font:** Liberation Sans / DejaVu Sans (open-source sans-serif equivalent under OFL/Apache 2.0).
- **Font Size & Line Height:** 12px body text, 14px window titles, 16px standard body line height.
- **Font X-Height:** 7px at 12px font size (measured from Liberation Sans metrics at native 1024x768 scale).
