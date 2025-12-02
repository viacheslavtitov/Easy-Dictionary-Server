#!/usr/bin/env bash
set -euo pipefail

echo "Setting up environment..."

export APP_ENV=local
export SERVER_CONFIG_JWT_EXP_TIME_MINUTES=60
export SERVER_CONFIG_REFRESH_JWT_EXP_TIME_MINUTES=180
export SERVER_CONFIG_JWT_SECRET="Local"
export SERVER_CONFIG_ADDRESS="localhost"
export DB_HOST="127.0.0.1"
export DB_NAME="easy_dictionary_dev"
export DB_PORT=5432
export DB_USER="postgres_local"
export DB_PASSWORD="postgres123"
export SERVER_CONFIG_PORT=8080
export SERVER_CONFIG_TIMEOUT=60

export NO_PROXY="localhost,127.0.0.1,::1"
export no_proxy="$NO_PROXY"

echo "Starting server..."
go run main.go Local