# `tcgui` Theming & Palette Guide (`docs/theming.md`)

**Status:** Complete (Task T-92)  
**Scope:** Palette tokens, theme configuration, 9-slice definitions, and schema generation.

---

## 1. Architecture

`tcgui` separates color tokens (`palette.json`) from structural styling (`theme.json`). All theme configuration files are validated against JSON schemas generated from Rust types (`schemars`).

---

## 2. Palette Tokens (`palette.json`)

Palettes define hex color tokens mapped to Pre-CU SWG reference IDs or assumptions:

```json
{
  "schema_version": 1,
  "tokens": {
    "panel.base":        { "hex": "#2A2A2A", "source": "R-019, verified" },
    "panel.bevel.dark":  { "hex": "#1A1A1A", "source": "R-019, verified" },
    "panel.bevel.light": { "hex": "#3A3A3A", "source": "R-019, verified" },
    "trim.copper":       { "hex": "#8B7355", "source": "R-010, verified" },
    "border.steel":      { "hex": "#4A4A4A", "source": "R-020, verified" },
    "text.primary":      { "hex": "#DDDDDD", "source": "assumed" }
  }
}
```

---

## 3. Theme Configuration (`theme.json`)

Themes define fonts, 9-slice texture mappings, and widget state styles (`normal`, `hover`, `pressed`, `disabled`):

```json
{
  "schema_version": 1,
  "fonts": {
    "default": { "file": "fonts/default.ttf", "size_px": 12 }
  },
  "styles": {
    "window": {
      "normal": { "nine_slice": "window_frame", "background": "panel.base", "background_opacity": 0.9 }
    },
    "button": {
      "normal":  { "nine_slice": "button_up",   "text": "text.primary" },
      "hover":   { "nine_slice": "button_up",   "tint": "trim.copper.light", "text": "text.primary" },
      "pressed": { "nine_slice": "button_down", "text": "text.primary" }
    }
  }
}
```

---

## 4. Schema Generation

Schemas are automatically generated and validated via `tcgui_core::theme::generate_schemas`, outputting `schemas/palette.schema.json` and `schemas/theme.schema.json`.
