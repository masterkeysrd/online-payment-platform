#!/usr/bin/env bash
set -euo pipefail

echo "Formatting code..."
go fmt ./...

echo "Formatting imports..."
go install golang.org/x/tools/cmd/goimports@latest
goimports -l -w .