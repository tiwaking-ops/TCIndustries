#!/usr/bin/env bash
set -e
cargo check --workspace
cargo test --workspace
cargo clippy --workspace -- -D warnings
