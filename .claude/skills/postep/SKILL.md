---
name: postep
description: Czyta i atomowo aktualizuje plik postep/student.json przez helper-skrypt postep.py. Każda modyfikacja przechodzi przez atomowy protokół (backup + walidacja + atomic mv) wykonywany przez skrypt — agent NIE składa JSON-a samodzielnie. Użyj na początku sesji (odczyt) i po każdej istotnej zmianie (zapis).
---

# Cel

Trzymać **jeden** plik z pełnym stanem ucznia (`postep/student.json`), z gwarancją, że żadna operacja go nie uszkodzi — wszystkie zapisy idą przez deterministyczny skrypt, **nie przez ręczne składanie JSON-a przez agenta**.

# Zasada twarda — kluczowa

**Agent NIGDY nie wykonuje `Write` ani `Edit` na `postep/student.json` bezpośrednio.**

Wszystkie operacje przez:
```bash
python3 .claude/skills/postep/postep.py <komenda> [argumenty]
```

Skrypt wykonuje 7-krokowy protokół: read → migrate → backup → modify → write tmp → validate → atomic mv. Agent tylko **woła go z argumentami**.

> **Uwaga:** `postep.py` to narzędzie kursu, nie materiał do nauki. Uczeń nigdy go nie uruchamia ani nie czyta — robisz to wyłącznie ty. Nie omawiaj go na lekcji.

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

**`go_version` zapisuj bez prefiksu `go`** — sam numer, np. `1.25.3`. To pole jest sprawdzane przed lekcjami 5.2 i 11.1 (wymagają ≥1.22).

# Komendy skryptu

Wszystkie wykonują pełen protokół atomowy. **Zawsze sprawdź kod wyjścia** — niezerowy = stary plik nietknięty.

## Inicjalizacja (po onboardingu)

```bash
python3 .claude/skills/postep/postep.py init \
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
python3 .claude/skills/postep/postep.py read
python3 .claude/skills/postep/postep.py read --field aktualna_lekcja
python3 .claude/skills/postep/postep.py read --field srodowisko
python3 .claude/skills/postep/postep.py read --field srodowisko.go_cmd
python3 .claude/skills/postep/postep.py read --field do_powtorki
```

## Ustawienie pola

```bash
python3 .claude/skills/postep/postep.py set --field aktualna_lekcja --value "4.2"
python3 .claude/skills/postep/postep.py set --field srodowisko.edytor --value "GoLand"
```

## Dopisanie ukończonej lekcji

```bash
python3 .claude/skills/postep/postep.py add-lekcja --id "4.1" --trudnosc 3
```

`--trudnosc` to 1-5 (subiektywna ocena ucznia: "1 = banalne, 5 = bardzo trudne"). Pytaj po lekcji. Skrypt sam aktualizuje `ostatnia_sesja`.

## Dopisanie ukończonego ćwiczenia

```bash
python3 .claude/skills/postep/postep.py add-cwiczenie --lekcja "4.1" --poziom warmup
# --poziom: warmup | main | star   (odpowiada 🔥 / ⭐ / ⚡)
```

## Mocne strony / do powtórki

```bash
python3 .claude/skills/postep/postep.py add-mocna-strona "samodzielne czytanie komunikatów kompilatora"
python3 .claude/skills/postep/postep.py add-do-powtorki --temat "odbiornik wskaźnikowy" --lekcja "7.2"
python3 .claude/skills/postep/postep.py remove-do-powtorki --temat "odbiornik wskaźnikowy"
```

`add-mocna-strona` trzyma max 7 najnowszych, duplikaty pomija.
`add-do-powtorki` nie dubluje tego samego tematu.

## Środowisko

```bash
python3 .claude/skills/postep/postep.py update-srodowisko \
  --system "Windows" \
  --go-cmd "go" \
  --go-version "1.25.3" \
  --shell "PowerShell"
```

Można podać dowolny podzbiór pól — tylko one się zmienią.

**Aktualizuj `go_version` po każdej aktualizacji Go u ucznia.** Nieaktualna wartość w pliku sprawi, że będziesz go ostrzegał przed nieistniejącym problemem albo przeoczysz prawdziwy.

## Notatki tutora (prywatne dla agenta)

```bash
python3 .claude/skills/postep/postep.py add-notatka "Anna chce kiedyś napisać narzędzie do porządkowania zdjęć"
```

Max 20 najnowszych. **Nie pokazuj uczniowi**, jeśli sam nie zapyta.

**Dobre kandydatki na notatkę:** pomysł ucznia na własny program (wraca w lekcji 12.1), co go zniechęca, co go wciąga, jak reaguje na utknięcie.

## Zakończenie sesji

```bash
python3 .claude/skills/postep/postep.py end-session
```

Ustawia `ostatnia_sesja=dziś`, `liczba_sesji+=1`. Wywołuj **raz** na koniec każdej sesji rozmowy.

## Recovery (gdy student.json uszkodzony)

```bash
python3 .claude/skills/postep/postep.py recovery
```

Skrypt szuka najnowszego **działającego** backupu w `postep/backups/`, przenosi uszkodzony do `postep/student.broken.<TS>.json` (NIE kasuje), kopiuje backup na miejsce, wypisuje podsumowanie przywróconego stanu.

# Procedura sesji

## Start sesji

1. **Odczyt:**
   ```bash
   python3 .claude/skills/postep/postep.py read
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
python3 .claude/skills/postep/postep.py init \
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
python3 .claude/skills/postep/postep.py add-cwiczenie --lekcja <X.Y> --poziom <warmup|main|star>
```

## Moduł 12 — projekt

Lekcja 12.2 rozciąga się na kilka sesji. Nie czekaj z zapisem do jej końca:
```bash
python3 .claude/skills/postep/postep.py add-notatka "projekt: dodany zapis do JSON, działa; następnie argumenty CLI"
python3 .claude/skills/postep/postep.py end-session
```
`add-lekcja --id 12.2` dopisz dopiero, gdy etap 1 projektu jest skończony.

## Koniec sesji rozmowy

```bash
python3 .claude/skills/postep/postep.py end-session
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
LAST=$(python3 .claude/skills/postep/postep.py read --field ostatnia_sesja | tr -d '"')
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

```bash
cd /tmp && rm -rf test-postep && mkdir -p test-postep && cd test-postep
cp /Users/marian/Projects/ITMOBILE-agent-nauka-golang/.claude/skills/postep/postep.py .
python3 postep.py init --imie Test --cel hobby --tempo "2-5" --system macOS --go-cmd go --go-version 1.25.3
python3 postep.py add-lekcja --id 1.1 --trudnosc 2
python3 postep.py read --field srodowisko.go_version
python3 postep.py read
cd /tmp && rm -rf test-postep
```
