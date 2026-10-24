//! Theme, palette, and style resolution with JSON loading and schema generation.

use schemars::JsonSchema;
use serde::{Deserialize, Serialize};
use std::collections::HashMap;
use std::path::Path;

#[derive(Debug, Clone, Serialize, Deserialize, JsonSchema)]
pub struct PaletteToken {
    pub hex: String,
    pub source: Option<String>,
}

#[derive(Debug, Clone, Serialize, Deserialize, JsonSchema)]
pub struct PaletteFile {
    pub schema_version: u32,
    pub tokens: HashMap<String, PaletteToken>,
}

#[derive(Debug, Clone, Serialize, Deserialize, JsonSchema)]
pub struct FontConfig {
    pub file: String,
    pub size_px: f32,
}

#[derive(Debug, Clone, Serialize, Deserialize, JsonSchema)]
pub struct NineSliceConfig {
    pub texture: String,
    pub margins: [u32; 4],
}

#[derive(Debug, Clone, Serialize, Deserialize, JsonSchema)]
pub struct StyleConfig {
    pub background: Option<String>,
    pub background_opacity: Option<f32>,
    pub tint: Option<String>,
    pub opacity: Option<f32>,
    pub nine_slice: Option<String>,
    pub text_color: Option<String>,
    pub text: Option<String>,
}

#[derive(Debug, Clone, Serialize, Deserialize, JsonSchema)]
pub struct ThemeFile {
    pub schema_version: u32,
    pub fonts: Option<HashMap<String, FontConfig>>,
    pub nine_slices: Option<HashMap<String, NineSliceConfig>>,
    pub styles: Option<HashMap<String, HashMap<String, StyleConfig>>>,
}

#[derive(Debug, thiserror::Error)]
pub enum ThemeError {
    #[error("IO error loading {path}: {source}")]
    Io { path: String, source: std::io::Error },
    #[error("JSON parse error in {path} at {json_path}: {message}")]
    Json { path: String, json_path: String, message: String },
    #[error("Validation error in {path}: {message}")]
    Validation { path: String, message: String },
}

pub fn load_palette_from_str(path: &str, json_str: &str) -> Result<PaletteFile, ThemeError> {
    let palette: PaletteFile = serde_json::from_str(json_str).map_err(|e| {
        let (line, col) = (e.line(), e.column());
        ThemeError::Json {
            path: path.to_string(),
            json_path: format!("line {}, col {}", line, col),
            message: e.to_string(),
        }
    })?;
    Ok(palette)
}

pub fn load_theme_from_str(path: &str, json_str: &str) -> Result<ThemeFile, ThemeError> {
    let theme: ThemeFile = serde_json::from_str(json_str).map_err(|e| {
        let (line, col) = (e.line(), e.column());
        ThemeError::Json {
            path: path.to_string(),
            json_path: format!("line {}, col {}", line, col),
            message: e.to_string(),
        }
    })?;
    Ok(theme)
}

pub fn generate_schemas(schemas_dir: &Path) -> Result<(), std::io::Error> {
    std::fs::create_dir_all(schemas_dir)?;

    let palette_schema = schemars::schema_for!(PaletteFile);
    let palette_json = serde_json::to_string_pretty(&palette_schema)?;
    std::fs::write(schemas_dir.join("palette.schema.json"), palette_json)?;

    let theme_schema = schemars::schema_for!(ThemeFile);
    let theme_json = serde_json::to_string_pretty(&theme_schema)?;
    std::fs::write(schemas_dir.join("theme.schema.json"), theme_json)?;

    Ok(())
}

#[cfg(test)]
mod tests {
    use super::*;

    const APPENDIX_D_PALETTE: &str = r##"{
      "schema_version": 1,
      "tokens": {
        "panel.base":         { "hex": "#2A2A2A", "source": "T-08, assumed" },
        "panel.bevel.dark":   { "hex": "#1A1A1A", "source": "T-08, assumed" },
        "panel.bevel.light":  { "hex": "#3A3A3A", "source": "T-08, assumed" },
        "trim.copper":        { "hex": "#8B7355", "source": "T-08, assumed" },
        "border.steel":       { "hex": "#4A4A4A", "source": "T-08, assumed" },
        "text.primary":       { "hex": "#DDDDDD", "source": "T-08, assumed" },
        "ham.background":     { "hex": "#FFFFFF", "source": "assumed; update from visual-spec row" }
      }
    }"##;

    const APPENDIX_D_THEME: &str = r##"{
      "schema_version": 1,
      "fonts": { "default": { "file": "fonts/default.ttf", "size_px": 12 } },
      "styles": {
        "window": {
          "normal": { "nine_slice": "window_frame", "background": "panel.base", "background_opacity": 0.9 }
        },
        "button": {
          "normal":   { "nine_slice": "button_up",   "text": "text.primary" },
          "hover":    { "nine_slice": "button_up",   "tint": "trim.copper.light", "text": "text.primary" },
          "pressed":  { "nine_slice": "button_down", "text": "text.primary" },
          "disabled": { "nine_slice": "button_up",   "text": "text.secondary", "opacity": 0.6 }
        }
      }
    }"##;

    #[test]
    fn test_load_appendix_d_palette() {
        let palette = load_palette_from_str("palette.json", APPENDIX_D_PALETTE).unwrap();
        assert_eq!(palette.schema_version, 1);
        assert!(palette.tokens.contains_key("panel.base"));
        assert_eq!(palette.tokens.get("panel.base").unwrap().hex, "#2A2A2A");
    }

    #[test]
    fn test_load_appendix_d_theme() {
        let theme = load_theme_from_str("theme.json", APPENDIX_D_THEME).unwrap();
        assert_eq!(theme.schema_version, 1);
        assert!(theme.styles.is_some());
        let styles = theme.styles.unwrap();
        assert!(styles.contains_key("button"));
        assert!(styles.get("button").unwrap().contains_key("hover"));
    }

    #[test]
    fn test_generate_schemas() {
        let manifest_dir = env!("CARGO_MANIFEST_DIR");
        let schemas_path = std::path::Path::new(manifest_dir).join("../../schemas");
        generate_schemas(&schemas_path).unwrap();
        assert!(schemas_path.join("palette.schema.json").exists());
        assert!(schemas_path.join("theme.schema.json").exists());
    }
}
