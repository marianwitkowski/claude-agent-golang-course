---
name: baza-wiedzy
description: Zarządza lokalną bazą wiedzy kursu Go (katalog wiedza/) — pobiera świeże wersje plików z trzech repozytoriów źródłowych, pokazuje stan bazy, pozwala porównać lokalną wersję ze zdalną, przywraca poprzednią wersję z backupu. Użyj gdy uczeń/autor mówi "odśwież bazę wiedzy", "pobierz najnowszą wersję materiałów", "sprawdź czy baza jest aktualna", "pokaż stan bazy wiedzy".
---

# Cel

Trzymać lokalną kopię materiałów źródłowych aktualną z repozytoriów na GitHubie, ale **nigdy nie aktualizować bez świadomej decyzji ucznia/autora** — żeby nie nadpisać lokalnych edycji.

# Architektura bazy wiedzy

```
wiedza/
├── zrodlo/                  # mirror golang-for-python-developers (32 pliki .md)
│   ├── 01-dlaczego-go.md
│   ├── ...
│   ├── 93-cwiczenia.md
│   ├── README.md
│   └── VERSION.json         # SHA commita, z którego pochodzi mirror
├── przyklady/               # mirror golang-examples (13 plików .md)
│   ├── golang-cli.md
│   ├── ...
│   └── kod/                 # pliki .go z golang20230314 (24 pliki)
│       ├── Dzien01-01-zmienne.go
│       └── ...
├── AKTUALIZACJE.md          # delta: Go 1.22 → 1.27, co zdezaktualizowało się w źródłach
├── INDEX.md                 # mapowanie źródeł → 38 lekcji sokratejskich
└── lekcje/                  # 38 gotowych lekcji + SZABLON-LEKCJI.md
```

**Reguła:** `zrodlo/` i `przyklady/` to czyste mirrory — nie edytujemy ich ręcznie. Zmiany merytoryczne lądują w `AKTUALIZACJE.md` (delta) albo w `lekcje/*.md` (już przerobione na sokratejskie).

# Trzy repozytoria źródłowe

| Repo | Cel w kursie | Katalog lokalny | Pliki |
| --- | --- | --- | --- |
| `marianwitkowski/golang-for-python-developers` | kanon merytoryczny | `wiedza/zrodlo/` | 32 `.md` |
| `marianwitkowski/golang-examples` | materiał na ćwiczenia i projekt | `wiedza/przyklady/` | 13 `.md` |
| `marianwitkowski/golang20230314` | minimalne przykłady kodu | `wiedza/przyklady/kod/` | 24 `.go` |

**Uwaga o materiale źródłowym:** repozytorium główne zakłada czytelnika, który zna już inny język programowania. Kurs jest dla osób, które nigdy nie programowały, więc bierzemy z niego wyłącznie treść o Go. Pliki `71-skladnia.md`, `72-przeksztalcenie.md` i `73-zaleznosci.md` to czyste zestawienia Go z innym językiem — **w lekcjach nieużywane**, trzymane wyłącznie jako referencja.

**Uwaga o `golang20230314`:** w oryginale pliki leżą w katalogach `Dzien01/` i `Dzien02/`. Lokalnie są spłaszczone do `Dzien01-01-zmienne.go` itd. (jeden katalog, prefiks w nazwie). Przy odświeżeniu zachowaj tę konwencję.

# Operacje

## 0. Krok wstępny — weryfikacja spójności (każda komenda)

**ZAWSZE** na początku **dowolnej** komendy bazy wiedzy wywołaj `verify_state`. Wykrywa ślady przerwanych aktualizacji (SIGKILL, padło zasilanie, agent zatrzymany w połowie).

```bash
verify_state() {
  # Sygnał 1: brak wiedza/zrodlo + istnieje wiedza/zrodlo.backup-*
  # → poprzednia aktualizacja padła między mv-A (zrodlo → backup) i mv-B (zrodlo.new → zrodlo)
  if [ ! -d wiedza/zrodlo ]; then
    NEWEST_BACKUP=$(ls -dt wiedza/zrodlo.backup-* 2>/dev/null | head -1)
    if [ -n "$NEWEST_BACKUP" ]; then
      echo "SIGNAL_1:$NEWEST_BACKUP"
      return 1
    fi
  fi

  # Sygnał 2: istnieje wiedza/zrodlo.new (pozostały po nieudanym buildzie)
  if [ -d wiedza/zrodlo.new ]; then
    echo "SIGNAL_2:wiedza/zrodlo.new"
    return 2
  fi

  # Sygnał 3: brak VERSION.json (baza sprzed wersjonowania — info, nie błąd)
  if [ -d wiedza/zrodlo ] && [ ! -f wiedza/zrodlo/VERSION.json ]; then
    echo "SIGNAL_3:no_version"
    return 3
  fi

  return 0
}

VERIFY_CODE=0
verify_state || VERIFY_CODE=$?
```

### Kanoniczny wzorzec wywołania

**ZAWSZE** `VERIFY_CODE=0; verify_state || VERIFY_CODE=$?` — niezależnie od tego, czy działasz pod `set -e`. Nigdy `verify_state; VERIFY_CODE=$?` (pod `set -e` shell przerwie, zanim odczytasz kod).

```bash
set -e
VERIFY_CODE=0
verify_state || VERIFY_CODE=$?

case $VERIFY_CODE in
  1) echo "Wykryto przerwaną aktualizację, pytam ucznia..." ;;
  2) echo "Zostawiony zrodlo.new — pytam ucznia..." ;;
  3) echo "ℹ️  Brak VERSION.json — kontynuuję bez wersjonowania" ;;
  0) ;;
esac
```

### Reakcja na każdy sygnał

**Sygnał 1 (kod 1):** `wiedza/zrodlo/` zniknął, jest backup. Agent **zatrzymuje wszystko** i pyta:

> "⚠️ Wykryłem ślady przerwanej aktualizacji bazy wiedzy:
> - `wiedza/zrodlo/` nie istnieje
> - Jest backup: `[NEWEST_BACKUP]` (SHA: [z VERSION.json], data: [...])
>
> Co robimy?
> - **`przywróć`** → odzyskaj `wiedza/zrodlo/` z backupu (zalecane)
> - **`pomiń`** → zostaw jak jest (bez materiałów kurs nie ruszy)
> - **`usuń backup`** → wyrzuć backup (NIE polecane)"

Po `przywróć`:
```bash
mv "$NEWEST_BACKUP" wiedza/zrodlo
echo "OK: przywrócono wiedza/zrodlo z $NEWEST_BACKUP"
```

**Sygnał 2 (kod 2):** `wiedza/zrodlo.new` z nieudanego buildu.
> "ℹ️ Wykryłem `wiedza/zrodlo.new/` z poprzedniej, nieudanej próby. Co robimy?
> - **`przenieś do failed`** → `mv` na `wiedza/zrodlo.new.failed-[TS]`
> - **`pomiń`** → zostawiam (kolejna aktualizacja może to nadpisać)"

```bash
TIMESTAMP=$(date +%Y-%m-%d-%H-%M-%S)
mv wiedza/zrodlo.new "wiedza/zrodlo.new.failed-${TIMESTAMP}"
```

**Sygnał 3 (kod 3):** brak `VERSION.json`. Kontynuuj operację, dopisz informacyjnie:
> "ℹ️ Baza nie ma `VERSION.json`. Zalecane: `odśwież bazę wiedzy`, by wpisać aktualny SHA."

**Kod 0:** stan spójny, kontynuuj bez powiadomień.

### Twarda reguła

**Żadna komenda bazy wiedzy nie wykonuje się bez wcześniejszego `verify_state`.**

## 1. Odśwież bazę (pobierz najnowszą wersję)

Protokół z walidacją, podglądem różnic i rollbackiem. Nigdy nie ruszamy `wiedza/zrodlo/`, dopóki **wszystko** w `/tmp/` nie przejdzie walidacji.

Poniższy protokół opisuje repo główne (`golang-for-python-developers` → `wiedza/zrodlo/`). Dla `golang-examples` i `golang20230314` przebiega analogicznie, z podmienioną nazwą repo, katalogiem docelowym i listą plików. **Odświeżaj po jednym repo naraz**, nie wszystkie w jednym bloku.

### Krok 1: pobierz SHA i metadane ze zdalnego repo

```bash
REPO="marianwitkowski/golang-for-python-developers"
SHA=$(curl -s "https://api.github.com/repos/${REPO}/commits/main" \
  | grep -m1 '"sha"' | cut -d'"' -f4)
DATE=$(curl -s "https://api.github.com/repos/${REPO}/commits/${SHA}" \
  | grep -m1 '"date"' | cut -d'"' -f4)
echo "Remote SHA: ${SHA}, date: ${DATE}"
```

Pusty `SHA` → STOP (brak sieci, API niedostępne, limit zapytań). Nie ruszaj nic.

Jeśli `gh` jest zalogowany, wygodniejsze i bez limitu anonimowego:
```bash
gh api "repos/${REPO}/commits/main" --jq .sha
gh api "repos/${REPO}/git/trees/HEAD?recursive=1" --jq '.tree[].path'
```
**Cudzysłowy wokół URL-a z `?` są obowiązkowe** — zsh inaczej próbuje go dopasować jako wzorzec i zgłasza `no matches found`.

### Krok 2: porównaj z lokalnym `VERSION.json`

```bash
if [ -f wiedza/zrodlo/VERSION.json ]; then
  LOCAL_SHA=$(grep -o '"sha"[[:space:]]*:[[:space:]]*"[^"]*"' wiedza/zrodlo/VERSION.json | cut -d'"' -f4)
  echo "Local SHA:  ${LOCAL_SHA}"
fi
```

- `SHA == LOCAL_SHA` → "Baza już aktualna (commit ${SHA:0:7} z ${DATE})." STOP.
- Różne lub brak `VERSION.json` → kontynuuj.

### Krok 3: pobierz pliki do katalogu tymczasowego

```bash
TMP="/tmp/go-kurs-update-${SHA:0:7}"
rm -rf "${TMP}" && mkdir -p "${TMP}"

FILES="01-dlaczego-go 02-podstawy-go \
       11-pierwszy-program 12-typy-danych 13-operatory-wyrazenia 14-stdin-stdout 15-logowanie \
       21-instrukcje-warunkowe 22-petle 23-funkcje \
       31-struktury 32-interfejsy 33-channels-goroutine \
       41-prog-wspolbiezne1 42-prog-wspolbiezne2 \
       51-obsluga-plikow 52-komunikacja-siec \
       61-testowanie 62-debugowanie 63-moduly 64-wyjatki \
       71-skladnia 72-przeksztalcenie 73-zaleznosci \
       81-aplikacja-cli 82-aplikacja-web 83-aplikacja-bazadanych 84-mikroserwisy \
       91-zasoby 92-biblioteki 93-cwiczenia README"

ALL_OK=true
for f in $FILES; do
  URL="https://raw.githubusercontent.com/${REPO}/${SHA}/${f}.md"
  HTTP=$(curl -sL -w "%{http_code}" -o "${TMP}/${f}.md" "$URL")
  if [ "$HTTP" != "200" ]; then
    echo "FAIL ${f}.md (HTTP ${HTTP})"
    ALL_OK=false
  fi
done

$ALL_OK || { echo "STOP: nie wszystkie pliki pobrały się"; rm -rf "${TMP}"; exit 1; }
```

Pobieramy z **konkretnego SHA**, nie z `main` — gwarancja spójnego snapshotu, nawet jeśli ktoś commitnie w trakcie pobierania.

**Jeśli lista plików w repo się zmieniła** (doszedł/zniknął plik) — nie zgaduj. Wypisz zdalne drzewo (`gh api ... git/trees`), pokaż różnicę uczniowi i zapytaj, zanim ruszysz dalej. Nowy plik źródłowy to zwykle sygnał, że `INDEX.md` wymaga ręcznej decyzji autora.

### Krok 4: walidacja każdego pliku

```bash
for f in $FILES; do
  P="${TMP}/${f}.md"
  SIZE=$(wc -c < "$P")
  FIRST=$(head -c 200 "$P")

  [ "$SIZE" -lt 500 ] && { echo "FAIL ${f}.md: za mały (${SIZE} B)"; ALL_OK=false; continue; }
  echo "$FIRST" | grep -qi "<!DOCTYPE\|<html" && { echo "FAIL ${f}.md: HTML zamiast .md"; ALL_OK=false; continue; }
  echo "$FIRST" | grep -qE "^#|^---" || { echo "FAIL ${f}.md: nie wygląda jak markdown"; ALL_OK=false; continue; }

  echo "OK   ${f}.md (${SIZE} B)"
done

$ALL_OK || { echo "STOP: walidacja nie przeszła"; rm -rf "${TMP}"; exit 1; }
```

Dla plików `.go` z `golang20230314` walidacja jest inna: rozmiar >50 B i obecność linii `package `:
```bash
head -5 "$P" | grep -q '^package ' || { echo "FAIL ${f}: brak deklaracji package"; ALL_OK=false; }
```
**Nie kompiluj ani nie uruchamiaj pobranych plików `.go`** — to kod z internetu.

### Krok 5: pokaż różnice uczniowi (PRZED nadpisaniem)

```bash
echo ""
echo "=== Zmiany ==="
for f in $FILES; do
  LOCAL="wiedza/zrodlo/${f}.md"
  REMOTE="${TMP}/${f}.md"
  if [ ! -f "$LOCAL" ]; then
    echo "NEW    ${f}.md ($(wc -c < $REMOTE) B)"
  elif ! cmp -s "$LOCAL" "$REMOTE"; then
    echo "MOD    ${f}.md ($(wc -c < $LOCAL) → $(wc -c < $REMOTE) B, $(diff "$LOCAL" "$REMOTE" | wc -l) linii diff)"
  fi
done
```

Pokaż też: lokalny SHA vs nowy, datę commita, liczbę zmienionych plików.

### Krok 6: poproś o potwierdzenie

> "Pobrałem 32 pliki z commita ${SHA:0:7} (${DATE}). Powyżej widzisz, które się zmieniły. Zaktualizować lokalną bazę? Wymagana jawna odpowiedź: **tak, zaktualizuj**."

Inna odpowiedź → STOP, `rm -rf ${TMP}`, koniec.

### Kroki 7-9: atomowy blok build + swap z automatycznym rollbackiem

Wszystkie operacje modyfikujące `wiedza/zrodlo/` idą w **jednym bloku bash** z `set -e; trap rollback ERR`. Jeśli cokolwiek padnie, rollback wykona się sam.

```bash
set -e
TIMESTAMP=$(date +%Y-%m-%d-%H-%M-%S)
NEW_DIR="wiedza/zrodlo.new"
BACKUP="wiedza/zrodlo.backup-${TIMESTAMP}"
EXPECTED=32

rollback() {
  local exit_code=$?
  echo ""
  echo "==================== BŁĄD — automatyczny rollback ===================="
  if [ ! -d wiedza/zrodlo ] && [ -d "$BACKUP" ]; then
    mv "$BACKUP" wiedza/zrodlo
    echo "ROLLBACK A: przywrócono wiedza/zrodlo z backupu"
  fi
  if [ -d "$NEW_DIR" ]; then
    mv "$NEW_DIR" "${NEW_DIR}.failed-${TIMESTAMP}"
    echo "ROLLBACK B: zrodlo.new → ${NEW_DIR}.failed-${TIMESTAMP}"
  fi
  echo "Stan po rollback:"
  ls -d wiedza/zrodlo* 2>/dev/null || echo "  (brak wiedza/zrodlo*)"
  echo "======================================================================"
  exit $exit_code
}
trap rollback ERR

# --- Krok 7: build zrodlo.new obok ---
if [ -d "$NEW_DIR" ]; then
  mv "$NEW_DIR" "${NEW_DIR}.stale-${TIMESTAMP}"
  echo "INFO: stary zrodlo.new → ${NEW_DIR}.stale-${TIMESTAMP}"
fi
mkdir -p "$NEW_DIR"
cp "${TMP}"/*.md "$NEW_DIR"/

# --- Krok 8: VERSION.json + walidacja ---
POBRANE=$(date -u +%Y-%m-%dT%H:%M:%SZ)
cat > "${NEW_DIR}/VERSION.json" <<EOF
{
  "sha": "${SHA}",
  "commit_date": "${DATE}",
  "pobrane": "${POBRANE}",
  "plikow": ${EXPECTED},
  "zrodlo": "https://github.com/${REPO}"
}
EOF

COUNT=$(ls "$NEW_DIR"/*.md 2>/dev/null | wc -l | tr -d ' ')
[ "$COUNT" = "$EXPECTED" ] || { echo "BŁĄD walidacji: oczekiwane ${EXPECTED}, znalezione $COUNT"; exit 1; }
# Walidacja VERSION.json: pięć kluczy obecnych i plik się domyka.
for klucz in sha commit_date pobrane plikow zrodlo; do
  grep -q "\"${klucz}\"" "${NEW_DIR}/VERSION.json" || { echo "BŁĄD: VERSION.json bez klucza ${klucz}"; exit 1; }
done
tail -c 2 "${NEW_DIR}/VERSION.json" | grep -q '}' || { echo "BŁĄD: VERSION.json ucięty"; exit 1; }

# --- Krok 9: atomowy swap przez 2x mv ---
if [ -d wiedza/zrodlo ]; then
  mv wiedza/zrodlo "$BACKUP"      # mv-A
fi
mv "$NEW_DIR" wiedza/zrodlo         # mv-B

trap - ERR
echo "OK: zaktualizowano wiedza/zrodlo/ do SHA ${SHA:0:7}"
echo "Backup poprzedniej wersji: ${BACKUP}"
```

**Co się dzieje przy awarii:**

| Co padło | Co robi `trap rollback` |
| --- | --- |
| `cp` z `/tmp` do `zrodlo.new` | `zrodlo` nietknięty, `zrodlo.new` → `failed-<TS>` |
| Walidacja (`COUNT ≠ 32`) | jw. |
| Zapis `VERSION.json` | jw. |
| mv-A (`zrodlo → backup`) | rzadkie; `zrodlo` nietknięty, `zrodlo.new` → failed |
| mv-B (`zrodlo.new → zrodlo`) | brak `zrodlo` + jest backup → `backup → zrodlo` |
| SIGKILL między mv-A i mv-B | trap nie odpali; następne uruchomienie wykryje to jako **Sygnał 1** w `verify_state` |

### Krok 10: sprzątanie `/tmp`

```bash
rm -rf "${TMP}"
```
To **jedyne** dozwolone `rm -rf` w całym protokole.

### Krok 11: powiadomienie

Powiedz uczniowi:
- Zaktualizowano do SHA `${SHA:0:7}` (commit z ${DATE})
- Backup: `wiedza/zrodlo.backup-${TIMESTAMP}/`
- Wyczyszczono `/tmp/...`
- Przypomnienie: "Jeśli zmiany były merytoryczne, przejrzyj `wiedza/AKTUALIZACJE.md` (czy delty nadal trafne) i `wiedza/INDEX.md` (czy mapowanie na 38 lekcji nadal się zgadza)."

**Odświeżenie źródeł nie zmienia lekcji.** `wiedza/lekcje/*.md` to osobny kanon — modyfikuje go wyłącznie autor, w trybie autora. Jeśli źródło zmieniło się merytorycznie, powiedz to wprost i zostaw decyzję.

### Awaryjny manualny rollback (bez `rm -rf`)

```bash
TIMESTAMP=$(date +%Y-%m-%d-%H-%M-%S)
mv wiedza/zrodlo "wiedza/zrodlo.failed-${TIMESTAMP}"
mv wiedza/zrodlo.backup-<TIMESTAMP_BACKUPU> wiedza/zrodlo
```

## 2. Stan bazy wiedzy

Gdy uczeń mówi "pokaż stan bazy":

```bash
[ -f wiedza/zrodlo/VERSION.json ] && cat wiedza/zrodlo/VERSION.json
ls wiedza/zrodlo/*.md | wc -l          # oczekiwane: 32
ls wiedza/przyklady/*.md | wc -l       # oczekiwane: 13
ls wiedza/przyklady/kod/*.go | wc -l   # oczekiwane: 24
ls wiedza/lekcje/*.md | wc -l          # oczekiwane: 39 (38 lekcji + SZABLON)
ls -d wiedza/zrodlo.backup-* 2>/dev/null
```

Wypisz uczniowi: lokalny SHA i datę commita, datę ostatniego pobrania, liczby plików, datę ostatniego commita w zdalnym repo (porównaj), listę backupów, czy `AKTUALIZACJE.md` i `INDEX.md` istnieją.

Brak `VERSION.json` → baza z pierwszego pobrania; zaproponuj odświeżenie.

## 3. Sprawdź różnice ze zdalnym repo (dry-run)

Gdy uczeń mówi "sprawdź czy baza aktualna": wykonaj **kroki 1-5**, **bez kroków 7-9**. Po pokazaniu różnic zapytaj: "Wykonać teraz aktualizację?". Nic nie modyfikujesz.

## 4. Przywróć poprzednią wersję bazy

1. Wylistuj backupy: `ls -dt wiedza/zrodlo.backup-*`
2. Pokaż listę z datami i SHA (z `VERSION.json` w każdym backupie)
3. Po wyborze:
   ```bash
   TIMESTAMP=$(date +%Y-%m-%d-%H-%M-%S)
   mv wiedza/zrodlo "wiedza/zrodlo.failed-${TIMESTAMP}"
   mv "wiedza/zrodlo.backup-${WYBRANY}" wiedza/zrodlo
   ```
4. Potwierdź: "Przywrócono z `zrodlo.backup-${WYBRANY}` (SHA: ...). Poprzednia wersja w `zrodlo.failed-${TIMESTAMP}/`."

## 5. Podaj konkretny plik źródłowy (dla skilla `lekcja`)

1. Sprawdź `wiedza/INDEX.md` — który plik i sekcja odpowiada lekcji (pole `zrodlo` we frontmatterze lekcji też to podaje)
2. Przeczytaj odpowiedni fragment `wiedza/zrodlo/NN-*.md`
3. **Sprawdź `wiedza/AKTUALIZACJE.md`** — czy temat ma deltę
4. **Zdejmij warstwę porównań do innych języków** — uczeń żadnego nie zna

# Lista plików źródłowych

## `wiedza/zrodlo/` — golang-for-python-developers

```
01-dlaczego-go            — czym jest Go, geneza, zastosowania
02-podstawy-go            — instalacja, moduły, struktura programu
11-pierwszy-program       — package main, func main, go run
12-typy-danych            — typy proste, zmienne, tablice, slices, mapy
13-operatory-wyrazenia    — operatory, konwersje, brak niejawnych
14-stdin-stdout           — fmt.Print*, bufio.Scanner, wejście
15-logowanie              — log, poziomy, slog
21-instrukcje-warunkowe   — if, switch, wartownicy
22-petle                  — for we wszystkich formach, range
23-funkcje                — funkcje, wiele zwracanych wartości, domknięcia, defer
31-struktury              — struct, metody, wskaźniki, embedding
32-interfejsy             — interfejsy niejawne, Stringer, type switch
33-channels-goroutine     — goroutine, kanały, podstawy
41-prog-wspolbiezne1      — WaitGroup, wzorce
42-prog-wspolbiezne2      — select, Mutex, wyścigi, -race
51-obsluga-plikow         — os, pliki, JSON, CSV
52-komunikacja-siec       — HTTP (poza kursem, referencja do 12.4)
61-testowanie             — go test, testy tablicowe, pokrycie
62-debugowanie            — go vet, delve, diagnostyka
63-moduly                 — go mod, zależności, wersjonowanie
64-wyjatki                — error jako wartość, panic/recover, opakowywanie
71-skladnia               — ZESTAWIENIE z innym językiem — nieużywane w lekcjach
72-przeksztalcenie        — ZESTAWIENIE z innym językiem — nieużywane w lekcjach
73-zaleznosci             — ZESTAWIENIE z innym językiem — nieużywane w lekcjach
81-aplikacja-cli          — narzędzie wiersza poleceń (moduł 12)
82-aplikacja-web          — web (tylko wzmianka w 12.4)
83-aplikacja-bazadanych   — bazy danych (tylko wzmianka w 12.4)
84-mikroserwisy           — mikroserwisy (tylko wzmianka w 12.4)
91-zasoby                 — gdzie szukać wiedzy (12.4)
92-biblioteki             — popularne biblioteki (12.4)
93-cwiczenia              — zadania do przerobienia
README                    — spis treści repo
```

## `wiedza/przyklady/` — golang-examples

`golang-barcode`, `golang-chat`, `golang-cli`, `golang-csv`, `golang-gofunk`, `golang-gofunk-top40`, `golang-goroutines`, `golang-gota`, `golang-json`, `golang-passgen`, `golang-resty`, `golang-shortlinks`, `golang-thumbnails`

Najużyteczniejsze w kursie: `golang-cli`, `golang-json`, `golang-csv`, `golang-passgen`, `golang-shortlinks` (moduły 9 i 12). Reszta używa bibliotek zewnętrznych — traktuj jako inspirację, nie materiał lekcyjny.

## `wiedza/przyklady/kod/` — golang20230314

24 pliki `.go`, spłaszczone z `Dzien01/` i `Dzien02/`. Minimalne, kompletne przykłady — dobre jako punkt wyjścia do eksperymentu w kroku 3 lekcji. `Dzien02-11-gin.go` i `Dzien02-12-scraping.go` wykraczają poza kurs.

# Repozytoria — szczegóły

```
https://github.com/marianwitkowski/golang-for-python-developers   (gałąź: main)
https://github.com/marianwitkowski/golang-examples                (gałąź: main)
https://github.com/marianwitkowski/golang20230314                 (gałąź: main)
```

Format raw: `https://raw.githubusercontent.com/<repo>/<SHA>/<plik>`

Ostatni commit:
```bash
curl -s "https://api.github.com/repos/<repo>/commits?per_page=1" | grep '"date"' | head -1
```

# Twarde zasady

- **`wiedza/zrodlo/` i `wiedza/przyklady/` to mirrory — nie edytuj ręcznie.** Zmiany merytoryczne idą do `AKTUALIZACJE.md`, dydaktyczne do `wiedza/lekcje/`.
- **Pobieranie tylko z konkretnego SHA**, nie z `main`.
- **Zawsze najpierw do `/tmp/`**, walidacja, dopiero potem dotykanie `wiedza/`.
- **Walidacja `.md`:** rozmiar >500 B, nie HTML, zaczyna się od `#` lub `---`. **Walidacja `.go`:** rozmiar >50 B, ma linię `package `.
- **Nigdy nie uruchamiaj ani nie kompiluj pobranego kodu `.go`.**
- **Buduj `zrodlo.new/` obok**, waliduj, dopiero swap. Nigdy nie modyfikuj `zrodlo/` "po fragmencie".
- **Atomowy swap przez `mv`.** Bez `rm -rf` na `wiedza/`.
- **Nieudane operacje archiwizuj** jako `*.failed-<TS>/`, nie kasuj.
- **`VERSION.json` commituj do gita** — pokazuje, z jakiego stanu pochodzi baza.
- **Nie usuwaj `AKTUALIZACJE.md` ani `INDEX.md`** przy odświeżeniu — to nasze aneksy.
- **Odświeżenie źródeł nie zmienia lekcji.** `wiedza/lekcje/` jest niezależnym kanonem; jego zmiana wymaga trybu autora.
- **Nie wykrywaj zmian automatycznie** — odświeżenie wymaga jawnej komendy.
- **Nie czyść backupów automatycznie.**
