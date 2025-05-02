@echo off
setlocal enabledelayedexpansion

echo Installing myimx ASCII Art CLI Tool...

:: Set version
set VERSION=1.0.0

:: Detect architecture
set ARCH=amd64
if "%PROCESSOR_ARCHITECTURE%"=="x86" set ARCH=386
if "%PROCESSOR_ARCHITECTURE%"=="ARM64" set ARCH=arm64

:: Set download URL
set DOWNLOAD_URL=https://github.com/yourusername/myimx/releases/download/v%VERSION%/myimx_%VERSION%_windows_%ARCH%.zip

:: Create temp directory
set TEMP_DIR=%TEMP%\myimx_install
if exist "%TEMP_DIR%" rd /s /q "%TEMP_DIR%"
mkdir "%TEMP_DIR%"

:: Download the zip file
echo Downloading myimx v%VERSION% for Windows/%ARCH%...
echo From: %DOWNLOAD_URL%

powershell -Command "[Net.ServicePointManager]::SecurityProtocol = [Net.SecurityProtocolType]::Tls12; Invoke-WebRequest -Uri '%DOWNLOAD_URL%' -OutFile '%TEMP_DIR%\myimx.zip'"
if %ERRORLEVEL% neq 0 (
  echo Failed to download the file. Please check your internet connection and try again.
  exit /b 1
)

:: Extract the zip
echo Extracting...
powershell -Command "Expand-Archive -Path '%TEMP_DIR%\myimx.zip' -DestinationPath '%TEMP_DIR%\extract'"

:: Create installation directory
set INSTALL_DIR=%USERPROFILE%\AppData\Local\Programs\myimx
if not exist "%INSTALL_DIR%" mkdir "%INSTALL_DIR%"

:: Copy the binary
echo Installing to %INSTALL_DIR%...
copy "%TEMP_DIR%\extract\myimx.exe" "%INSTALL_DIR%"
if %ERRORLEVEL% neq 0 (
  echo Failed to copy the file. Please ensure you have permission to write to %INSTALL_DIR%.
  exit /b 1
)

:: Add to PATH if not already there
echo Checking PATH...
set "PATH_UPDATED="
echo %PATH% | findstr /C:"%INSTALL_DIR%" > nul
if %ERRORLEVEL% neq 0 (
  echo Adding to PATH...
  setx PATH "%PATH%;%INSTALL_DIR%"
  set "PATH_UPDATED=yes"
)

:: Clean up
rd /s /q "%TEMP_DIR%"

echo.
echo ✅ Installation complete!
echo You can now use myimx from your command line.
echo Try it now: myimx list
echo.

if defined PATH_UPDATED (
  echo NOTE: You need to restart any open Command Prompt or PowerShell windows 
  echo for the PATH changes to take effect.
)

pause