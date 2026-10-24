#!/usr/bin/env bash
set -e
rustup default stable-x86_64-pc-windows-gnu
cargo check --workspace
cargo test -p tcgui-core
cargo clippy --workspace -- -D warnings

# Optional importer surface (M8 / task T-81). Off by default, so it must be
# checked explicitly here to avoid silently rotting behind the feature flag.
cargo check -p tcgui-render --features swg-import
cargo test -p tcgui-render --features swg-import
cargo clippy -p tcgui-render --features swg-import --all-targets -- -D warnings
