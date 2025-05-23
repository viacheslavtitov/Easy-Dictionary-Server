@echo off
echo Converting...
go tool cover -html='coverage.out'
pause