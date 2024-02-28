#!/usr/bin/env bash
set -euo pipefail

# Download dependencies
echo "Executing ./scripts/deps.sh"
bash ./scripts/deps.sh

# Build the simulator
echo "Building the simulator"
go build -o bin/simulator cmd/simulator/main.go
