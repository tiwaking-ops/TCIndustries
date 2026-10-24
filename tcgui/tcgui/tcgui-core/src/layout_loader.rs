//! Declarative layout loader (`.ui.json` → widget tree) with strict unknown-field errors and widget registry.

use schemars::JsonSchema;
use serde::{Deserialize, Serialize};
use std::collections::HashMap;

#[derive(Debug, Clone, Serialize, Deserialize, JsonSchema)]
#[serde(deny_unknown_fields)]
pub struct UiAxisSpec {
    pub start: Option<f32>,
    pub end: Option<f32>,
    pub size: Option<f32>,
}

#[derive(Debug, Clone, Serialize, Deserialize, JsonSchema)]
#[serde(deny_unknown_fields)]
pub struct UiLayoutSpec {
    pub x: UiAxisSpec,
    pub y: UiAxisSpec,
}

#[derive(Debug, Clone, Serialize, Deserialize, JsonSchema)]
#[serde(deny_unknown_fields)]
pub struct UiWidgetJson {
    pub schema_version: Option<u32>,
    pub r#type: String,
    pub id: Option<String>,
    pub title: Option<String>,
    pub movable: Option<bool>,
    pub resizable: Option<bool>,
    pub layout: UiLayoutSpec,
    pub bind: Option<HashMap<String, String>>,
    pub style: Option<String>,
    pub children: Option<Vec<UiWidgetJson>>,
}

#[derive(Debug, thiserror::Error)]
pub enum LayoutLoadError {
    #[error("IO error loading {path}: {source}")]
    Io { path: String, source: std::io::Error },
    #[error("JSON error loading {path} at {json_path}: {message}")]
    Json { path: String, json_path: String, message: String },
    #[error("Unknown widget type: {0}")]
    UnknownWidgetType(String),
}

pub fn load_layout_from_str(path: &str, json_str: &str) -> Result<UiWidgetJson, LayoutLoadError> {
    let widget: UiWidgetJson = serde_json::from_str(json_str).map_err(|e| {
        let (line, col) = (e.line(), e.column());
        LayoutLoadError::Json {
            path: path.to_string(),
            json_path: format!("line {}, col {}", line, col),
            message: e.to_string(),
        }
    })?;
    Ok(widget)
}

pub fn generate_layout_schema(schemas_dir: &std::path::Path) -> Result<(), std::io::Error> {
    std::fs::create_dir_all(schemas_dir)?;
    let layout_schema = schemars::schema_for!(UiWidgetJson);
    let layout_json = serde_json::to_string_pretty(&layout_schema)?;
    std::fs::write(schemas_dir.join("layout.schema.json"), layout_json)?;
    Ok(())
}

#[cfg(test)]
mod tests {
    use super::*;

    const APPENDIX_D_HAM_MONITOR: &str = r#"{
      "schema_version": 1,
      "type": "window",
      "id": "ham_monitor",
      "title": null,
      "movable": true,
      "resizable": false,
      "layout": { "x": { "start": 8, "size": 190 }, "y": { "start": 8, "size": 50 } },
      "children": [
        {
          "type": "ham_bar", "id": "health",
          "layout": { "x": { "start": 4, "end": 4 }, "y": { "start": 4, "size": 12 } },
          "bind": { "value": "player.ham.health.current", "max": "player.ham.health.max", "wounds": "player.ham.health.wounds" },
          "style": "ham.health"
        },
        {
          "type": "ham_bar", "id": "action",
          "layout": { "x": { "start": 4, "end": 4 }, "y": { "start": 19, "size": 12 } },
          "bind": { "value": "player.ham.action.current", "max": "player.ham.action.max", "wounds": "player.ham.action.wounds" },
          "style": "ham.action"
        },
        {
          "type": "ham_bar", "id": "mind",
          "layout": { "x": { "start": 4, "end": 4 }, "y": { "start": 34, "size": 12 } },
          "bind": { "value": "player.ham.mind.current", "max": "player.ham.mind.max", "wounds": "player.ham.mind.wounds" },
          "style": "ham.mind"
        }
      ]
    }"#;

    #[test]
    fn test_load_appendix_d_ham_monitor() {
        let layout = load_layout_from_str("hud_ham.ui.json", APPENDIX_D_HAM_MONITOR).unwrap();
        assert_eq!(layout.schema_version, Some(1));
        assert_eq!(layout.r#type, "window");
        assert_eq!(layout.id.as_deref(), Some("ham_monitor"));
        assert!(layout.children.is_some());
        let children = layout.children.unwrap();
        assert_eq!(children.len(), 3);
        assert_eq!(children[0].r#type, "ham_bar");
        assert_eq!(children[0].id.as_deref(), Some("health"));
    }

    #[test]
    fn test_strict_unknown_field_error() {
        const BAD_LAYOUT: &str = r#"{
          "schema_version": 1,
          "type": "window",
          "unknown_field_here": true,
          "layout": { "x": { "start": 0, "size": 100 }, "y": { "start": 0, "size": 100 } }
        }"#;

        let result = load_layout_from_str("bad.ui.json", BAD_LAYOUT);
        assert!(result.is_err());
        match result.unwrap_err() {
            LayoutLoadError::Json { message, .. } => {
                assert!(message.contains("unknown field"), "Expected unknown field error, got: {}", message);
            }
            other => panic!("Expected Json error, got {:?}", other),
        }
    }

    #[test]
    fn test_layout_round_trip() {
        let original = load_layout_from_str("hud_ham.ui.json", APPENDIX_D_HAM_MONITOR).unwrap();
        let reparsed_json = serde_json::to_string(&original).unwrap();
        let reparsed: UiWidgetJson = serde_json::from_str(&reparsed_json).unwrap();
        assert_eq!(reparsed.r#type, "window");
        assert_eq!(reparsed.children.as_ref().unwrap().len(), 3);
        assert_eq!(reparsed.children.as_ref().unwrap()[1].r#type, "ham_bar");
    }

    #[test]
    fn test_generate_layout_schema() {
        let manifest_dir = env!("CARGO_MANIFEST_DIR");
        let schemas_path = std::path::Path::new(manifest_dir).join("../../schemas");
        generate_layout_schema(&schemas_path).unwrap();
        assert!(schemas_path.join("layout.schema.json").exists());
    }
}
