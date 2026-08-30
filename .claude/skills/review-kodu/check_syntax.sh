#!/usr/bin/env bash
# Sprawdza kod ucznia BEZ URUCHAMIANIA go.
#   ./check_syntax.sh <plik.go>        -> parsowanie (gofmt -e) + różnica formatowania
#   ./check_syntax.sh --build <katalog> -> kompilacja do /dev/null (NIE uruchamia binarki)
# NIGDY nie używa `go run`. `go build -o /dev/null` kompiluje i wyrzuca wynik.
set -uo pipefail

usage() { echo "użycie: $0 <plik.go> | $0 --build <katalog>"; exit 2; }
[ $# -ge 1 ] || usage

if [ "$1" = "--build" ]; then
  [ $# -eq 2 ] || usage
  dir="$2"
  [ -d "$dir" ] || { echo "BŁĄD: brak katalogu $dir"; exit 2; }
  echo "== kompilacja (bez uruchamiania): $dir =="
  go build -o /dev/null "./$dir" && echo "OK: kompiluje się"
  exit $?
fi

plik="$1"
[ -f "$plik" ] || { echo "BŁĄD: brak pliku $plik"; exit 2; }

echo "== składnia (gofmt -e) =="
if err=$(gofmt -e "$plik" 2>&1 >/dev/null) && [ -z "$err" ]; then
  echo "OK: parsuje się"
else
  echo "$err"
  exit 1
fi

echo "== formatowanie (gofmt -d) =="
diff=$(gofmt -d "$plik" 2>/dev/null)
if [ -z "$diff" ]; then
  echo "OK: sformatowany kanonicznie"
else
  echo "$diff"
  echo "-- uczeń poprawia sam: gofmt -w $plik"
fi
