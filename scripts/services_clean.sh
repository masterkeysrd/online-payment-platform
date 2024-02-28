#!/usr/bin/env bash
set -euo pipefail

# Stop Docker-Compose Services
echo "Stopping Docker-Compose Services..."
docker-compose -f ./deployments/docker/services.yml down

# Remove Docker-Compose Services
echo "Removing Docker-Compose Services..."
docker-compose -f ./deployments/docker/services.yml rm -f

# Remove Docker-Compose Services Images
echo "Removing Docker-Compose Services Images..."
docker-compose -f ./deployments/docker/services.yml down --rmi all

# Remove Docker-Compose Services Volumes
echo "Removing Docker-Compose Services Volumes..."
docker-compose -f ./deployments/docker/services.yml down -v