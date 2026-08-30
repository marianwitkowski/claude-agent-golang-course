#!/usr/bin/env bash
# Uruchamia narzędzie postep (Go) niezależnie od bieżącego katalogu.
#
#   bash .claude/skills/postep/postep.sh read
#   bash .claude/skills/postep/postep.sh add-lekcja --id 4.1 --trudnosc 3
#
# Katalog główny projektu wyliczamy ze ścieżki tego skryptu, więc wołanie
# działa tak samo z korzenia repozytorium, jak i z kurs/zadania.
#
# Binarkę budujemy raz do .bin/ i przebudowujemy tylko po zmianie źródeł.
# `go run` byłby prostszy, ale dokleja własne "exit status N" do stderr,
# co zaciemnia komunikaty błędów narzędzia.
set -uo pipefail

KATALOG_SKILLA="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
KATALOG_GLOWNY="$(cd "${KATALOG_SKILLA}/../../.." && pwd)"
BINARKA="${KATALOG_SKILLA}/.bin/postep"

if [ ! -x "${BINARKA}" ] \
   || [ "${KATALOG_SKILLA}/postep.go" -nt "${BINARKA}" ] \
   || [ "${KATALOG_SKILLA}/go.mod" -nt "${BINARKA}" ]; then
  mkdir -p "${KATALOG_SKILLA}/.bin" || exit 1
  ( cd "${KATALOG_SKILLA}" && go build -o "${BINARKA}" . ) || {
    echo "BŁĄD: nie udało się zbudować narzędzia postep" >&2
    exit 1
  }
fi

exec "${BINARKA}" -root "${KATALOG_GLOWNY}" "$@"
