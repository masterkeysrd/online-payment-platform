#!/usr/bin/env bash
set -euo pipefail

# Stop Docker-Compose App
echo "Stopping Docker-Compose App..."
docker-compose -f ./deployments/docker/app.yml down

# Remove Docker-Compose App
echo "Removing Docker-Compose App..."
docker-compose -f ./deployments/docker/app.yml rm -f

# Remove Docker-Compose App Images
echo "Removing Docker-Compose App Images..."
docker-compose -f ./deployments/docker/app.yml down --rmi all

# Remove Docker-Compose App Volumes
echo "Removing Docker-Compose App Volumes..."
docker-compose -f ./deployments/docker/app.yml down -v