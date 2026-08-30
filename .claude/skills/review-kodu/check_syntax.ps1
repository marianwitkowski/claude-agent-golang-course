#!/usr/bin/env pwsh
# Sprawdza kod ucznia BEZ URUCHAMIANIA go. Odpowiednik check_syntax.sh dla Windows.
#
#   powershell -ExecutionPolicy Bypass -File .claude\skills\review-kodu\check_syntax.ps1 <plik.go>
#   powershell -ExecutionPolicy Bypass -File .claude\skills\review-kodu\check_syntax.ps1 --build <katalog>
#
# NIGDY nie używa `go run`. Tryb --build kompiluje do pliku tymczasowego
# i od razu go kasuje — binarka nie ląduje w katalogu ucznia.

[Console]::OutputEncoding = [System.Text.Encoding]::UTF8

function Uzycie {
    [Console]::Error.WriteLine('użycie: check_syntax.ps1 <plik.go> | check_syntax.ps1 --build <katalog>')
    exit 2
}

foreach ($narzedzie in @('go', 'gofmt')) {
    if (-not (Get-Command $narzedzie -ErrorAction SilentlyContinue)) {
        [Console]::Error.WriteLine("BŁĄD: nie znalazłem komendy ``$narzedzie`` w PATH")
        exit 2
    }
}

$argumenty = @($args)
if ($argumenty.Count -lt 1) { Uzycie }

# --- tryb kompilacji ---
if ($argumenty[0] -eq '--build' -or $argumenty[0] -eq '-build') {
    if ($argumenty.Count -ne 2) { Uzycie }
    $katalog = $argumenty[1]
    if (-not (Test-Path $katalog -PathType Container)) {
        [Console]::Error.WriteLine("BŁĄD: brak katalogu $katalog")
        exit 2
    }

    Write-Host "== kompilacja (bez uruchamiania): $katalog =="
    # cd do katalogu zadania: go szuka go.mod idąc w górę, więc działa
    # niezależnie od tego, skąd wywołano skrypt.
    $kod = 1
    $tymczasowa = Join-Path ([System.IO.Path]::GetTempPath()) ("check-syntax-" + [System.Guid]::NewGuid().ToString('N') + '.exe')
    Push-Location $katalog
    try {
        & go build -o $tymczasowa .
        $kod = $LASTEXITCODE
    }
    finally {
        Pop-Location
        if (Test-Path $tymczasowa) { Remove-Item $tymczasowa -Force }
    }

    if ($kod -eq 0) { Write-Host 'OK: kompiluje się' }
    exit $kod
}

# --- tryb pliku ---
$plik = $argumenty[0]
if (-not (Test-Path $plik -PathType Leaf)) {
    [Console]::Error.WriteLine("BŁĄD: brak pliku $plik")
    exit 2
}

Write-Host '== składnia (gofmt -e) =='
$wynik = & gofmt -e $plik 2>&1
if ($LASTEXITCODE -ne 0) {
    # Przy błędzie parsowania gofmt nie wypisuje kodu na stdout — zostają
    # same komunikaty o błędach, z numerem linii i kolumny.
    $wynik | ForEach-Object { Write-Host $_ }
    exit 1
}
Write-Host 'OK: parsuje się'

Write-Host '== formatowanie (gofmt -d) =='
$roznica = & gofmt -d $plik 2>$null
if (-not $roznica) {
    Write-Host 'OK: sformatowany kanonicznie'
}
else {
    $roznica | ForEach-Object { Write-Host $_ }
    Write-Host "-- uczeń poprawia sam: gofmt -w $plik"
}
