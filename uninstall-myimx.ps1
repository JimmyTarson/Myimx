Write-Host "Uninstalling myimx ASCII Art CLI tool..." -ForegroundColor Yellow

$installDir = Join-Path $env:USERPROFILE 'AppData\Local\Programs\myimx'

if (!(Test-Path $installDir)) {
    Write-Host "myimx doesn't appear to be installed at $installDir" -ForegroundColor Red
    exit 1
}

$currentPath = [Environment]::GetEnvironmentVariable('Path', 'User')
if ($currentPath -like "*$installDir*") {
    Write-Host "Removing myimx from PATH..." -ForegroundColor Yellow
    
    $newPath = ($currentPath -split ';' | Where-Object { $_ -ne $installDir }) -join ';'
    
    [Environment]::SetEnvironmentVariable('Path', $newPath, 'User')
    Write-Host "Removed from PATH" -ForegroundColor Green
}

try {
    Write-Host "Removing installation directory..." -ForegroundColor Yellow
    Remove-Item -Path $installDir -Recurse -Force
    Write-Host "Removed $installDir" -ForegroundColor Green
} catch {
    Write-Host "Error removing installation directory: $_" -ForegroundColor Red
    Write-Host "You may need to close any applications using myimx and try again." -ForegroundColor Yellow
    exit 1
}

Write-Host "`n✅ myimx has been successfully uninstalled!" -ForegroundColor Green
Write-Host "You may need to restart any open Command Prompt or PowerShell windows for the PATH changes to take effect."