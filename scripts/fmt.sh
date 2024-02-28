#!/usr/bin/env bash
set -euo pipefail

echo "Formatting code..."
go fmt ./...
