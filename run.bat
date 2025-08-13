@echo off
echo Starting Auto-Unterhaltsrechner...

REM Check if executable exists
if exist "auto-unterhaltsrechner.exe" (
    echo Running auto-unterhaltsrechner.exe...
    start auto-unterhaltsrechner.exe
) else (
    echo Executable not found. Trying to run with go run...
    set CGO_ENABLED=1
    go run ./cmd/app
)

pause