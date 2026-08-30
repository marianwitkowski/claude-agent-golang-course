#!/usr/bin/env pwsh
# Uruchamia narzędzie postep (Go) niezależnie od bieżącego katalogu.
# Odpowiednik postep.sh dla Windows PowerShell.
#
#   powershell -ExecutionPolicy Bypass -File .claude\skills\postep\postep.ps1 read
#   powershell -ExecutionPolicy Bypass -File .claude\skills\postep\postep.ps1 add-lekcja --id 4.1 --trudnosc 3
#
# Katalog główny projektu wyliczamy ze ścieżki tego skryptu, więc wołanie
# działa tak samo z korzenia repozytorium, jak i z kurs\zadania.
#
# Binarkę budujemy raz do .bin\ i przebudowujemy tylko po zmianie źródeł.

[Console]::OutputEncoding = [System.Text.Encoding]::UTF8

if (-not (Get-Command go -ErrorAction SilentlyContinue)) {
    [Console]::Error.WriteLine('BŁĄD: nie znalazłem komendy `go` w PATH')
    exit 2
}

$KatalogSkilla = $PSScriptRoot
$KatalogGlowny = (Resolve-Path (Join-Path $KatalogSkilla '..\..\..')).Path
$KatalogBin    = Join-Path $KatalogSkilla '.bin'
$Binarka       = Join-Path $KatalogBin 'postep.exe'

$zrodla = @('postep.go', 'go.mod') | ForEach-Object { Join-Path $KatalogSkilla $_ }

$trzebaBudowac = -not (Test-Path $Binarka)
if (-not $trzebaBudowac) {
    $czasBinarki = (Get-Item $Binarka).LastWriteTimeUtc
    foreach ($zrodlo in $zrodla) {
        if ((Get-Item $zrodlo).LastWriteTimeUtc -gt $czasBinarki) {
            $trzebaBudowac = $true
            break
        }
    }
}

if ($trzebaBudowac) {
    New-Item -ItemType Directory -Force -Path $KatalogBin | Out-Null
    Push-Location $KatalogSkilla
    try {
        & go build -o $Binarka .
        if ($LASTEXITCODE -ne 0) {
            [Console]::Error.WriteLine('BŁĄD: nie udało się zbudować narzędzia postep')
            exit 1
        }
    }
    finally {
        Pop-Location
    }
}

& $Binarka -root $KatalogGlowny @args
exit $LASTEXITCODE
