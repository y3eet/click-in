#!/usr/bin/env bash
set -euo pipefail

(
    cd backend
    air -c .air.toml
)
air_pid=$!

(
    cd web
    bun dev
)
web_pid=$!

wait "$air_pid" "$web_pid"