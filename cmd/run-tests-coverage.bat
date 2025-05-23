@echo off
echo Running tests with coverage...
go test -v ./../... -coverprofile='coverage.out'
pause