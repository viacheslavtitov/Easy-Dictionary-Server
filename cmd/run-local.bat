@echo off
echo Setting up environment...

set APP_ENV=local
set SERVER_CONFIG_JWT_EXP_TIME_MINUTES=60
set SERVER_CONFIG_REFRESH_JWT_EXP_TIME_MINUTES=180
set SERVER_CONFIG_JWT_SECRET=Local
set SERVER_CONFIG_ADDRESS=localhost
set DB_HOST=127.0.0.1
set DB_NAME=local_database
set DB_PORT=5430
set DB_USER=admin
set DB_PASSWORD=qwerty123
set SERVER_CONFIG_PORT=8080
set SERVER_CONFIG_TIMEOUT=60
set NO_PROXY=localhost,127.0.0.1

echo Starting server...
go run main.go Local
pause
