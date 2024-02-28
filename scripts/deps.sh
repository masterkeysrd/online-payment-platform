#!/usr/bin/env bash
set -euo pipefail

echo "Installing dependencies..."
go get -u github.com/gin-gonic/gin
go get -u gorm.io/gorm
go get -u gorm.io/driver/postgres
go get -u github.com/joho/godotenv
go get -u github.com/matoous/go-nanoid/v2
go mod tidy