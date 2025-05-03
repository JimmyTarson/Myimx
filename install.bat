@echo off
setlocal enabledelayedexpansion

echo Installing myimx ASCII Art CLI Tool...

:: Set download URL
set DOWNLOAD_URL=https://github.com/JimmyTarson12/Myimx/releases/download/v1.0.0/myimx.exe

:: Create installation directory
set INSTALL_DIR=%USERPROFILE%\AppData\Local\Programs\myimx
if not exist "%INSTALL_DIR%" mkdir "%INSTALL_DIR%"

:: Download the executable directly
echo Downloading myimx...
echo From: %DOWNLOAD_URL%

powershell -Command "[Net.ServicePointManager]::SecurityProtocol = [Net.SecurityProtocolType]::Tls12; Invoke-WebRequest -Uri '%DOWNLOAD_URL%' -OutFile '%INSTALL_DIR%\myimx.exe'"
if %ERRORLEVEL% neq 0 (
  echo Failed to download the file. Please check your internet connection and try again.
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