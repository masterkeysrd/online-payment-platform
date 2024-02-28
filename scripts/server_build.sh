#!/usr/bin/env bash
set -euo pipefail

# Build the simulator
echo "Building the simulator"
go build -o bin/server cmd/server/main.go