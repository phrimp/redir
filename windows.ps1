# .\windows.ps1 -ProjectDir "C:\Path\To\Your\Project"
# Function to print messages
function Print-Message {
    param (
        [string]$Message
    )
    Write-Host $Message -ForegroundColor Green
}

Print-Message "Starting the Go project setup script for Windows..."

# Check if Chocolatey is installed
if (-not (Get-Command choco -ErrorAction SilentlyContinue)) {
    Print-Message "Chocolatey not found. Installing Chocolatey..."
    Set-ExecutionPolicy Bypass -Scope Process -Force
    [System.Net.ServicePointManager]::SecurityProtocol = [System.Net.SecurityProtocolType]::Tls12
    Invoke-Expression ((New-Object System.Net.WebClient).DownloadString('https://chocolatey.org/install.ps1'))
} else {
    Print-Message "Chocolatey is already installed."
}

# Install Golang
if (-not (Get-Command go -ErrorAction SilentlyContinue)) {
    Print-Message "Installing Golang..."
    choco install golang -y
} else {
    Print-Message "Golang is already installed. Skipping installation."
}

# Verify Golang installation
if (-not (Get-Command go -ErrorAction SilentlyContinue)) {
    Write-Host "Failed to install Golang. Exiting." -ForegroundColor Red
    exit 1
}

$GoVersion = go version
Print-Message "Golang installed successfully. Version: $GoVersion"

# Check for an existing Go project
param (
    [string]$ProjectDir
)

if (-not $ProjectDir) {
    Write-Host "No project directory specified. Usage: .\windows.ps1 -ProjectDir 'C:\Path\To\Your\Project'" -ForegroundColor Yellow
    exit 1
}

if (-not (Test-Path "$ProjectDir\go.mod")) {
    Write-Host "No 'go.mod' file found in $ProjectDir. Ensure this is a valid Go project." -ForegroundColor Red
    exit 1
}

Print-Message "Navigating to project directory: $ProjectDir"
Set-Location -Path $ProjectDir

# Download all project dependencies
Print-Message "Downloading all dependencies for the Go project..."
go mod download

if ($LASTEXITCODE -eq 0) {
    Print-Message "Dependencies downloaded successfully!"
} else {
    Write-Host "Failed to download dependencies. Please check for errors above." -ForegroundColor Red
    exit 1
}

# Verify the project dependencies
Print-Message "Verifying dependencies..."
go mod tidy

Print-Message "Go project dependencies setup complete. You're ready to code!"
