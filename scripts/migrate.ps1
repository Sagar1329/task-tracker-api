param(
    [Parameter(Position = 0, Mandatory = $true)]
    [ValidateSet("up", "down", "version", "force", "create")]
    [string]$Command,

    [Parameter(Position = 1)]
    [string]$Argument
)

# -----------------------------
# Load .env
# -----------------------------
Get-Content ".env" | ForEach-Object {

    if ($_ -match '^\s*#') { return }
    if ($_ -match '^\s*$') { return }

    $parts = $_ -split '=', 2

    if ($parts.Count -eq 2) {
        [Environment]::SetEnvironmentVariable(
            $parts[0].Trim(),
            $parts[1].Trim(),
            "Process"
        )
    }
}

# -----------------------------
# URL Encode Password
# -----------------------------
Add-Type -AssemblyName System.Web

$encodedPassword = [System.Web.HttpUtility]::UrlEncode($env:DB_PASSWORD)

$databaseUrl = "postgres://$($env:DB_USER):$encodedPassword@$($env:DB_HOST):$($env:DB_PORT)/$($env:DB_NAME)?sslmode=$($env:DB_SSLMODE)"

# -----------------------------
# Commands
# -----------------------------
switch ($Command) {

    "up" {
        migrate `
            -path migrations `
            -database $databaseUrl `
            up
    }

    "down" {
        migrate `
            -path migrations `
            -database $databaseUrl `
            down 1
    }

    "version" {
        migrate `
            -path migrations `
            -database $databaseUrl `
            version
    }

    "force" {

        if (-not $Argument) {
            Write-Host ""
            Write-Host "Usage:"
            Write-Host ".\scripts\migrate.ps1 force <version>"
            exit 1
        }

        migrate `
            -path migrations `
            -database $databaseUrl `
            force $Argument
    }

    "create" {

        if (-not $Argument) {
            Write-Host ""
            Write-Host "Usage:"
            Write-Host ".\scripts\migrate.ps1 create <migration_name>"
            exit 1
        }

        migrate create `
            -ext sql `
            -dir migrations `
            -seq `
            $Argument
    }

}