#!/usr/bin/env bash
set -euo pipefail

echo "Installing dependencies..."
go get -u github.com/gin-gonic/gin
go get -u gorm.io/gorm
go get -u gorm.io/driver/postgres
go mod tidy