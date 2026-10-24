//! Procedural 9-slice bevel texture generator for `themes/swg-precu/`.
//!
//! Generates hash-stable procedural textures (window frame, button states, placeholder
//! icons) from palette tokens. Reads `themes/swg-precu/palette.json` when present and
//! falls back to measured Pre-CU defaults from `docs/visual-spec.md` (R-019, R-020,
//! R-010) otherwise. Output is deterministic: a re-run produces byte-identical files.

use serde::Deserialize;
use std::collections::HashMap;
use std::fs;
use std::path::{Path, PathBuf};

/// Pixel type used when building textures.
type Rgb = (u8, u8, u8);

/// Where palette tokens and generated textures live.
const THEME_DIR: &str = "themes/swg-precu";

/// Parses an HTML/CSS-style hex color (`#rgb`, `#rrggbb`, or bare hex) into an RGB8 triple.
fn parse_hex_color(hex: &str) -> Rgb {
    let hex = hex.trim().trim_start_matches('#').trim();
    match hex.len() {
        3 => (
            hex_to_u8(&format!("{}{}", &hex[0..1], &hex[0..1])),
            hex_to_u8(&format!("{}{}", &hex[1..2], &hex[1..2])),
            hex_to_u8(&format!("{}{}", &hex[2..3], &hex[2..3])),
        ),
        6 => (
            hex_to_u8(&hex[0..2]),
            hex_to_u8(&hex[2..4]),
            hex_to_u8(&hex[4..6]),
        ),
        _ => (42, 42, 42),
    }
}

fn hex_to_u8(s: &str) -> u8 {
    u8::from_str_radix(s, 16).unwrap_or(0)
}

/// Layout describing a palette token in `palette.json`.
#[derive(Debug, Clone, Deserialize)]
#[allow(dead_code)]
struct PaletteToken {
    hex: String,
    source: Option<String>,
}

/// Top-level palette file shape (`themes/swg-precu/palette.json`).
#[derive(Debug, Deserialize)]
#[allow(dead_code)]
struct PaletteFile {
    schema_version: Option<u32>,
    tokens: HashMap<String, PaletteToken>,
}

/// Loads palette tokens from the theme directory. Falls back to measured Pre-CU
/// defaults from `docs/visual-spec.md` when the file is absent or unreadable.
fn load_palette(theme_dir: &Path) -> HashMap<String, Rgb> {
    let mut tokens: HashMap<String, Rgb> = HashMap::new();
    let palette_path = theme_dir.join("palette.json");

    if let Ok(text) = fs::read_to_string(&palette_path) {
        if let Ok(file) = serde_json::from_str::<PaletteFile>(&text) {
            for (name, token) in file.tokens {
                tokens.insert(name, parse_hex_color(&token.hex));
            }
        } else {
            eprintln!(
                "asset-gen: warning: failed to parse {} ; using default palette tokens",
                palette_path.display()
            );
        }
    } else {
        eprintln!(
            "asset-gen: palette.json not found at {}; using default Pre-CU palette tokens",
            palette_path.display()
        );
    }

    // Measured / verified Pre-CU defaults (visual-spec rows, R-010/R-019/R-020),
    // used as fallback for any token missing from palette.json.
    let defaults: &[(&str, Rgb)] = &[
        ("panel.base", parse_hex_color("#2A2A2A")),
        ("panel.bevel.dark", parse_hex_color("#1A1A1A")),
        ("panel.bevel.light", parse_hex_color("#3A3A3A")),
        ("trim.copper", parse_hex_color("#8B7355")),
        ("trim.copper.light", parse_hex_color("#A88F6E")),
        ("border.steel", parse_hex_color("#4A4A4A")),
        ("text.primary", parse_hex_color("#DDDDDD")),
        ("text.secondary", parse_hex_color("#999999")),
        ("ham.background", parse_hex_color("#FFFFFF")),
    ];
    for (name, rgb) in defaults {
        tokens.entry(name.to_string()).or_insert(*rgb);
    }

    tokens
}

/// Writes a flat RGB buffer to a `P3` ASCII PPM file (deterministic, no external deps).
fn write_ppm(path: &Path, width: u32, height: u32, pixels: &[Rgb]) {
    let mut data = format!("P3\n{} {}\n255\n", width, height);
    for (r, g, b) in pixels {
        data.push_str(&format!("{} {} {}\n", r, g, b));
    }
    fs::write(path, data).unwrap();
}

/// Bevel geometry: outer edge uses `light` for raised (up), inner edge uses `dark`.
/// When pressed, the roles invert for an inverted (depressed) look.
fn generate_beveled_panel(width: u32, height: u32, base: Rgb, dark: Rgb, light: Rgb, pressed: bool) -> Vec<Rgb> {
    let mut pixels = vec![base; (width * height) as usize];
    for y in 0..height {
        for x in 0..width {
            let idx = usize::try_from(y * width + x).unwrap();
            let on_outer = x == 0 || y == 0 || x == width - 1 || y == height - 1;
            let on_inner = x == 1 || y == 1 || x == width - 2 || y == height - 2;
            pixels[idx] = if on_outer {
                if pressed { dark } else { light }
            } else if on_inner {
                if pressed { light } else { dark }
            } else {
                base
            };
        }
    }
    pixels
}

/// Kind of procedural texture to generate from palette tokens.
enum TextureKind {
    WindowFrame { pressed: bool },
    ButtonUp,
    ButtonDown,
    PlaceholderIcon,
}

/// Resolves a texture kind into a pixel buffer using the loaded palette tokens.
fn build_texture(kind: &TextureKind, width: u32, height: u32, tokens: &HashMap<String, Rgb>) -> Vec<Rgb> {
    let panel_base = *tokens.get("panel.base").unwrap();
    let bevel_dark = *tokens.get("panel.bevel.dark").unwrap();
    let bevel_light = *tokens.get("panel.bevel.light").unwrap();
    let border = *tokens.get("border.steel").unwrap();
    let trim = *tokens.get("trim.copper").unwrap();

    match kind {
        TextureKind::WindowFrame { pressed } => generate_beveled_panel(width, height, panel_base, bevel_dark, bevel_light, *pressed),
        TextureKind::ButtonUp => generate_beveled_panel(width, height, trim, bevel_dark, bevel_light, false),
        TextureKind::ButtonDown => generate_beveled_panel(width, height, trim, bevel_dark, bevel_light, true),
        TextureKind::PlaceholderIcon => {
            let mut pixels = vec![border; (width * height) as usize];
            let inner_x0 = width / 8;
            let inner_y0 = height / 8;
            let inner_x1 = width - width / 8;
            let inner_y1 = height - height / 8;
            for y in inner_y0..inner_y1 {
                for x in inner_x0..inner_x1 {
                    pixels[usize::try_from(y * width + x).unwrap()] = panel_base;
                }
            }
            pixels
        }
    }
}

fn main() {
    println!("Running asset-gen for Pre-CU SWG theme...");

    let theme_dir = PathBuf::from(THEME_DIR);
    fs::create_dir_all(&theme_dir).unwrap();

    let tokens = load_palette(&theme_dir);

    let textures: Vec<(&str, TextureKind, u32, u32)> = vec![
        ("window_frame.ppm", TextureKind::WindowFrame { pressed: false }, 32, 32),
        ("button_up.ppm", TextureKind::ButtonUp, 32, 32),
        ("button_down.ppm", TextureKind::ButtonDown, 32, 32),
        ("placeholder_icon.ppm", TextureKind::PlaceholderIcon, 32, 32),
    ];

    let mut generated: Vec<&str> = Vec::new();
    for (file_name, kind, width, height) in &textures {
        let path = theme_dir.join(file_name);
        let pixels = build_texture(kind, *width, *height, &tokens);
        write_ppm(&path, *width, *height, &pixels);
        generated.push(file_name);
    }

    // Persist a manifest of generated outputs. Re-runs overwrite deterministically
    // (hash-stable): identical inputs always produce byte-identical files.
    let manifest = format!(
        "manifest_version = 1\noutputs = {}\n",
        generated.iter().map(|s| format!("\"{}\"", s)).collect::<Vec<_>>().join(", ")
    );
    fs::write(theme_dir.join("asset_gen_manifest.txt"), manifest).unwrap();

    println!(
        "Asset generation complete. Outputs written to {:?}:",
        theme_dir
    );
    for name in &generated {
        println!("  - {}", name);
    }
}
