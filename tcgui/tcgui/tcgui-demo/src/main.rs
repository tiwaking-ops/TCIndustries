//! Demo entry point for `tcgui`.
//!
//! Parses `--screenshot <path>` and `--size <WxH>` flags and prints a startup
//! banner. In headless mode a placeholder screenshot is written; full WGPU
//! rendering is wired up in `tcgui-platform` (see `docs/usage.md`).
#![allow(missing_docs)]

use std::env;

fn main() {
    let args: Vec<String> = env::args().collect();
    let mut screenshot_path = None;
    let mut size_str = String::from("1024x768");

    let mut i = 1;
    while i < args.len() {
        match args[i].as_str() {
            "--screenshot" if i + 1 < args.len() => {
                screenshot_path = Some(args[i + 1].clone());
                i += 1;
            }
            "--size" if i + 1 < args.len() => {
                size_str = args[i + 1].clone();
                i += 1;
            }
            _ => {}
        }
        i += 1;
    }

    println!("tcgui-demo starting. Screenshot: {:?}, Size: {}", screenshot_path, size_str);

    if let Some(path) = screenshot_path {
        println!("Writing mock headless screenshot to {}...", path);
        std::fs::write(&path, b"mock-screenshot-data").unwrap();
    }
}
