@echo off
echo Building Auto-Unterhaltsrechner...

REM Set CGO environment
set CGO_ENABLED=1

REM Try to build with go build first
echo Attempting go build...
go build -ldflags="-H windowsgui" -o auto-unterhaltsrechner.exe ./cmd/app
if %ERRORLEVEL% EQU 0 (
    echo Build successful! Executable: auto-unterhaltsrechner.exe
    goto :end
)

echo Go build failed, trying with fyne package...

REM Check if fyne command is available
fyne version >nul 2>&1
if %ERRORLEVEL% NEQ 0 (
    echo Installing fyne command...
    go install fyne.io/fyne/v2/cmd/fyne@latest
)

REM Try fyne package
echo Building with fyne package...
fyne package -o auto-unterhaltsrechner.exe ./cmd/app
if %ERRORLEVEL% EQU 0 (
    echo Build successful with fyne! Executable: auto-unterhaltsrechner.exe
) else (
    echo Build failed. Please check the README.md for build requirements.
    echo You may need to install TDM-GCC or Visual Studio.
)

:end
pause