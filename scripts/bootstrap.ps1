# kwiki bootstrap for Windows
# Usage: irm https://raw.githubusercontent.com/you/kwiki/main/scripts/bootstrap.ps1 | iex

$ErrorActionPreference = "Stop"

Write-Host "🚀 kwiki bootstrap starting..." -ForegroundColor Magenta

# Check if Go is installed, install if not
if (-not (Get-Command go -ErrorAction SilentlyContinue)) {
    Write-Host "📦 Installing Go..." -ForegroundColor Cyan
    winget install GoLang.Go --silent --accept-package-agreements --accept-source-agreements
    $env:PATH += ";C:\Program Files\Go\bin"
}

# Clone or download kwiki
$kwikiDir = "$env:USERPROFILE\.kwiki"
if (-not (Test-Path $kwikiDir)) {
    if (Get-Command git -ErrorAction SilentlyContinue) {
        git clone https://github.com/you/kwiki.git $kwikiDir
    } else {
        Write-Host "📦 Installing Git first..." -ForegroundColor Cyan
        winget install Git.Git --silent --accept-package-agreements --accept-source-agreements
        $env:PATH += ";C:\Program Files\Git\bin"
        git clone https://github.com/you/kwiki.git $kwikiDir
    }
}

# Build kwiki
Set-Location $kwikiDir
go build -o kwiki.exe .

# Add to PATH
$userPath = [Environment]::GetEnvironmentVariable("PATH", "User")
if ($userPath -notlike "*$kwikiDir*") {
    [Environment]::SetEnvironmentVariable("PATH", "$userPath;$kwikiDir", "User")
    $env:PATH += ";$kwikiDir"
}

Write-Host "✅ kwiki installed! Run: kwiki install" -ForegroundColor Green
kwiki install
