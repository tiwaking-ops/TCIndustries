#!/usr/bin/env bash
set -e
rustup default stable-x86_64-pc-windows-gnu
cargo check --workspace
cargo test -p tcgui-core
cargo clippy --workspace -- -D warnings
