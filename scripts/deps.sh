#!/usr/bin/env bash
set -euo pipefail

echo "Installing dependencies..."
go get -u github.com/gin-gonic/gin
go mod tidy