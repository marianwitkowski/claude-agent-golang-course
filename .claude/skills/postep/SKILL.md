---
name: postep
description: Czyta i atomowo aktualizuje plik postep/student.json przez narzędzie postep napisane w Go. Każda modyfikacja przechodzi przez atomowy protokół (backup + walidacja + atomic rename) wykonywany przez narzędzie — agent NIE składa JSON-a samodzielnie. Użyj na początku sesji (odczyt) i po każdej istotnej zmianie (zapis).
---

# Cel

Trzymać **jeden** plik z pełnym stanem ucznia (`postep/student.json`), z gwarancją, że żadna operacja go nie uszkodzi — wszystkie zapisy idą przez deterministyczne narzędzie, **nie przez ręczne składanie JSON-a przez agenta**.

# Zasada twarda — kluczowa

**Agent NIGDY nie wykonuje `Write` ani `Edit` na `postep/student.json` bezpośrednio.**

Wszystkie operacje przez:
```bash
bash .claude/skills/postep/postep.sh <komenda> [argumenty]
```

Narzędzie wykonuje 7-krokowy protokół: odczyt → migracja → backup → modyfikacja → zapis `.tmp` → walidacja → atomowy `rename`. Agent tylko **woła je z argumentami**.

## Co to jest i dlaczego przez wrapper

`postep` to program w Go (`postep.go`, biblioteka standardowa, bez zależności) w osobnym module `.claude/skills/postep/`. Wrapper:
- wylicza katalog główny projektu ze swojej własnej ścieżki, więc **wołasz go z dowolnego miejsca** — z korzenia repozytorium tak samo jak z `kurs/zadania`
- buduje binarkę do `.claude/skills/postep/.bin/` przy pierwszym użyciu i przebudowuje ją tylko po zmianie źródeł (kolejne wywołania są natychmiastowe)

Są dwa równoważne wrappery:

| Powłoka | Wywołanie |
| --- | --- |
| bash / zsh (macOS, Linux, WSL, Git Bash) | `bash .claude/skills/postep/postep.sh <cmd>` |
| Windows PowerShell | `powershell -ExecutionPolicy Bypass -File .claude\skills\postep\postep.ps1 <cmd>` |

**Argumenty są identyczne.** W tym pliku wszystkie przykłady są w wersji `bash` — na Windows bez Git Basha podmień sam prefiks. Wyboru dokonujesz raz na sesję, zgodnie z sekcją „Narzędzia kursu" w `go-tutor.md`.

To **jedyne** miejsce, w którym wolno ci wywołać `go build`, i dotyczy narzędzia kursu, nie kodu ucznia. Zakaz uruchamiania kodu ucznia obowiązuje bez zmian.

> **Uwaga:** `postep` to narzędzie kursu, nie materiał do nauki. Uczeń nigdy go nie uruchamia ani nie czyta — robisz to wyłącznie ty. Nie omawiaj go na lekcji, nawet gdy jesteście przy module 9 albo 10 i wygląda na dobry przykład.

# Schemat student.json (schema_version 2)

```json
{
  "schema_version": 2,
  "imie": "Anna",
  "cel": "praca",
  "tempo_godz_tydz": "2-5",
  "rozpoczeto": "2026-08-30",
  "ostatnia_sesja": "2026-08-30",
  "liczba_sesji": 3,
  "aktualna_lekcja": "4.1",
  "srodowisko": {
    "system": "macOS",
    "go_cmd": "go",
    "go_version": "1.25.3",
    "shell": "zsh",
    "edytor": "VS Code"
  },
  "ukonczone_lekcje": [
    {"id": "1.1", "data": "2026-08-30", "trudnosc_subiektywna": 2}
  ],
  "ukonczone_cwiczenia": [
    {"lekcja": "1.1", "poziom": "warmup", "data": "2026-08-30"}
  ],
  "mocne_strony": ["czytanie komunikatów kompilatora"],
  "do_powtorki": [
    {"temat": "slice a tablica", "lekcja": "4.1", "data_zauwazenia": "2026-08-31"}
  ],
  "notatki_tutora": ["Anna lubi konkretne przykłady z życia"]
}
```

**Migracja schematu** jest automatyczna:
- v0 (bez `schema_version`) → v1 (dopisuje pole)
- v1 → v2 (dopisuje puste `srodowisko`)

**Pola z nowszego schematu są zachowywane.** Jeśli plik zawiera klucze, których to narzędzie nie zna, wracają do zapisu nietknięte — starsza wersja narzędzia nie skasuje stanu, którego nie rozumie. Plik z `schema_version` wyższą niż obsługiwana jest odrzucany, a nie nadpisywany.

**`go_version` zapisuj bez prefiksu `go`** — sam numer, np. `1.25.3`. To pole jest sprawdzane przed lekcjami 5.2 i 11.1 (wymagają ≥1.22).

# Komendy skryptu

Wszystkie wykonują pełen protokół atomowy. **Zawsze sprawdź kod wyjścia** — niezerowy = stary plik nietknięty.

## Inicjalizacja (po onboardingu)

```bash
bash .claude/skills/postep/postep.sh init \
  --imie "Anna" \
  --cel "hobby" \
  --tempo "2-5" \
  --system "macOS" \
  --go-cmd "go" \
  --go-version "1.25.3" \
  --shell "zsh" \
  --edytor "VS Code"
```

Tworzy plik z danymi z onboardingu + **pełen snapshot środowiska**. Listy puste, `liczba_sesji=1`, `aktualna_lekcja="1.1"`.

Błąd, jeśli plik już istnieje (ochrona przed nadpisaniem).

## Odczyt

```bash
bash .claude/skills/postep/postep.sh read
bash .claude/skills/postep/postep.sh read --field aktualna_lekcja
bash .claude/skills/postep/postep.sh read --field srodowisko
bash .claude/skills/postep/postep.sh read --field srodowisko.go_cmd
bash .claude/skills/postep/postep.sh read --field do_powtorki
```

## Ustawienie pola

```bash
bash .claude/skills/postep/postep.sh set --field aktualna_lekcja --value "4.2"
bash .claude/skills/postep/postep.sh set --field srodowisko.edytor --value "GoLand"
```

## Dopisanie ukończonej lekcji

```bash
bash .claude/skills/postep/postep.sh add-lekcja --id "4.1" --trudnosc 3
```

`--trudnosc` to 1-5 (subiektywna ocena ucznia: "1 = banalne, 5 = bardzo trudne"). Pytaj po lekcji. Skrypt sam aktualizuje `ostatnia_sesja`.

## Dopisanie ukończonego ćwiczenia

```bash
bash .claude/skills/postep/postep.sh add-cwiczenie --lekcja "4.1" --poziom warmup
# --poziom: warmup | main | star   (odpowiada 🔥 / ⭐ / ⚡)
```

## Mocne strony / do powtórki

```bash
bash .claude/skills/postep/postep.sh add-mocna-strona "samodzielne czytanie komunikatów kompilatora"
bash .claude/skills/postep/postep.sh add-do-powtorki --temat "odbiornik wskaźnikowy" --lekcja "7.2"
bash .claude/skills/postep/postep.sh remove-do-powtorki --temat "odbiornik wskaźnikowy"
```

`add-mocna-strona` trzyma max 7 najnowszych, duplikaty pomija.
`add-do-powtorki` nie dubluje tego samego tematu.

## Środowisko

```bash
bash .claude/skills/postep/postep.sh update-srodowisko \
  --system "Windows" \
  --go-cmd "go" \
  --go-version "1.25.3" \
  --shell "PowerShell"
```

Można podać dowolny podzbiór pól — tylko one się zmienią.

**Aktualizuj `go_version` po każdej aktualizacji Go u ucznia.** Nieaktualna wartość w pliku sprawi, że będziesz go ostrzegał przed nieistniejącym problemem albo przeoczysz prawdziwy.

## Notatki tutora (prywatne dla agenta)

```bash
bash .claude/skills/postep/postep.sh add-notatka "Anna chce kiedyś napisać narzędzie do porządkowania zdjęć"
```

Max 20 najnowszych. **Nie pokazuj uczniowi**, jeśli sam nie zapyta.

**Dobre kandydatki na notatkę:** pomysł ucznia na własny program (wraca w lekcji 12.1), co go zniechęca, co go wciąga, jak reaguje na utknięcie.

## Zakończenie sesji

```bash
bash .claude/skills/postep/postep.sh end-session
```

Ustawia `ostatnia_sesja=dziś`, `liczba_sesji+=1`. Wywołuj **raz** na koniec każdej sesji rozmowy.

## Recovery (gdy student.json uszkodzony)

```bash
bash .claude/skills/postep/postep.sh recovery
```

Skrypt szuka najnowszego **działającego** backupu w `postep/backups/`, przenosi uszkodzony do `postep/student.broken.<TS>.json` (NIE kasuje), kopiuje backup na miejsce, wypisuje podsumowanie przywróconego stanu.

# Procedura sesji

## Start sesji

1. **Odczyt:**
   ```bash
   bash .claude/skills/postep/postep.sh read
   ```
2. Błąd "plik nie istnieje" → uczeń nowy, uruchom onboarding
3. Błąd "JSON nie parsuje się" → `recovery` (zapytaj ucznia najpierw)
4. W normalnym przypadku zwróć agentowi:
   - `imie`, `aktualna_lekcja`
   - 2-3 ostatnie wpisy z `ukonczone_lekcje`
   - `do_powtorki` (jeśli niepusta)
   - liczbę dni od `ostatnia_sesja` (>7 → quiz odświeżający)
   - **`srodowisko.go_cmd`, `srodowisko.go_version`, `srodowisko.system`** — potrzebne w każdej komendzie pokazywanej uczniowi
5. **Sprawdź `go_version`.** Jeśli < 1.22 → zatrzymaj i wywołaj skill `setup-go` przed lekcją.

## Onboarding (pierwsza sesja)

Po wywiadzie + skill `setup-go` (który zna system, komendę i wersję):

```bash
bash .claude/skills/postep/postep.sh init \
  --imie <imię_z_wywiadu> \
  --cel <cel> \
  --tempo <tempo> \
  --system <z_setup-go> \
  --go-cmd <z_setup-go> \
  --go-version <z_setup-go>
```

## Po każdej ukończonej lekcji

1. Zapytaj: "Od 1 do 5, jak trudna była ta lekcja?"
2. `add-lekcja --id <X.Y> --trudnosc <N>`
3. `set --field aktualna_lekcja --value <następna_z_INDEX.md>`
4. Opcjonalnie `add-mocna-strona` / `add-do-powtorki`

Sekcja **Po lekcji** w pliku lekcji podaje dokładnie, jaka jest następna lekcja.

## Po każdym ukończonym ćwiczeniu

```bash
bash .claude/skills/postep/postep.sh add-cwiczenie --lekcja <X.Y> --poziom <warmup|main|star>
```

## Moduł 12 — projekt

Lekcja 12.2 rozciąga się na kilka sesji. Nie czekaj z zapisem do jej końca:
```bash
bash .claude/skills/postep/postep.sh add-notatka "projekt: dodany zapis do JSON, działa; następnie argumenty CLI"
bash .claude/skills/postep/postep.sh end-session
```
`add-lekcja --id 12.2` dopisz dopiero, gdy etap 1 projektu jest skończony.

## Koniec sesji rozmowy

```bash
bash .claude/skills/postep/postep.sh end-session
```

# Backupy

Skrypt tworzy backup do `postep/backups/student.{ISO_timestamp}.json` przed każdą modyfikacją.

**Nie kasuj backupów automatycznie.** Jeśli `postep/backups/` rośnie (>50 plików), powiadom ucznia:
> "Masz 53 backupy `student.json`, najstarszy z 2026-01-15. Chcesz przenieść starsze niż 30 dni do `postep/backups/_old/`?"

Po `tak`:
```bash
mkdir -p postep/backups/_old
# wypisz listę, przenieś przez mv — NIGDY find -delete
```

# Przy >7 dniach przerwy

```bash
LAST=$(bash .claude/skills/postep/postep.sh read --field ostatnia_sesja | tr -d '"')
```

Powiedz uczniowi:
> "Cześć [imię]! Ostatnio rozmawialiśmy [N] dni temu. Chcesz najpierw szybką powtórkę, czy lecimy dalej z lekcją [aktualna_lekcja]?"

# Twarde zasady

- **NIGDY** bezpośredni `Write` / `Edit` na `student.json`. ZAWSZE przez skrypt.
- **NIGDY** nie buduj nowego JSON-a "z pamięci" — skrypt czyta, modyfikuje wskazane pola, zapisuje. To chroni przed utratą pól z przyszłych wersji schematu.
- **Nie wymyślaj danych.** Nie znasz wartości → pytaj ucznia.
- **`notatki_tutora` są prywatne** — nie pokazuj bez prośby.
- **Daty** zawsze ISO `YYYY-MM-DD` — skrypt robi to sam.
- **`student.json` jest w `.gitignore`** — to stan konkretnego ucznia, nie część kursu.

# Test poprawności (dla autora)

Narzędzie ma własny moduł, więc sprawdzasz je jak każdy kod Go:

```bash
cd .claude/skills/postep
gofmt -l .          # cisza = sformatowane
go vet ./...        # cisza = brak zastrzeżeń
go build -o /dev/null .
```

Test na żywym stanie — **tylko gdy `postep/student.json` nie istnieje**, inaczej `init` odmówi:

```bash
cd /Users/marian/Projects/ITMOBILE-agent-nauka-golang
P() { bash .claude/skills/postep/postep.sh "$@"; }

P init --imie Test --cel hobby --tempo "2-5" --system macOS --go-cmd go --go-version 1.25.3
P add-lekcja --id 1.1 --trudnosc 2
P read --field srodowisko.go_version
P read
rm -f postep/student.json && rm -rf postep/backups   # sprzątanie po teście
```
