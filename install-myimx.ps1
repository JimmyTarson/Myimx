# install-myimx.ps1
$installDir = Join-Path $env:USERPROFILE 'AppData\Local\Programs\myimx'
$exeUrl = 'https://github.com/JimmyTarson12/Myimx/releases/download/v1.1.2/myimx.exe'

# Create installation directory if it doesn't exist
if(!(Test-Path $installDir)) {
    New-Item -Path $installDir -ItemType Directory -Force
    Write-Host "Created installation directory: $installDir"
}

# Download the executable
Write-Host "Downloading myimx..."
[Net.ServicePointManager]::SecurityProtocol = [Net.SecurityProtocolType]::Tls12
Invoke-WebRequest -Uri $exeUrl -OutFile (Join-Path $installDir 'myimx.exe')
Write-Host "Download complete!"

# Add to PATH if not already there
$currentPath = [Environment]::GetEnvironmentVariable('Path', 'User')
if($currentPath -notlike "*$installDir*") {
    [Environment]::SetEnvironmentVariable('Path', "$currentPath;$installDir", 'User')
    Write-Host "Added to PATH"
}

Write-Host "`n✅ Installation complete!" -ForegroundColor Green
Write-Host "You can now use myimx from your command line."
Write-Host "NOTE: You may need to restart your command prompt for the PATH changes to take effect."
Write-Host "Try it now: myimx list`n"