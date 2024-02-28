#!/usr/bin/env bash
set -euo pipefail

# Start Docker-Compose Services
docker-compose -f ./deployments/docker/services.yml up -d