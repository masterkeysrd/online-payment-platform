#!/usr/bin/env bash
set -euo pipefail

# Start Docker-Compose App
echo "Starting Docker-Compose App"
docker-compose -f ./deployments/docker/app.yml up -d --build