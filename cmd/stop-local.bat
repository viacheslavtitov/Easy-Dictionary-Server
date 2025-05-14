@echo off
FOR /F "tokens=5" %%A IN ('netstat -aon ^| find ":8080" ^| find "LISTENING"') DO (
    echo Killing process on port 8080 with PID %%A
    taskkill /PID %%A /F
)
pause