#!/usr/bin/env bash
set -euo pipefail

# Stop Docker-Compose App
echo "Stopping Docker-Compose App"
docker-compose -f ./deployments/docker/app.yml down --remove-orphans