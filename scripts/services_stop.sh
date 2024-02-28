#!/usr/bin/env bash
set -euo pipefail

# Stop Docker-Compose Services
echo "Stopping Docker-Compose Services..."
docker-compose -f ./deployments/docker/services.yml down --remove-orphans