---
name: reset-kursu
description: Resetuje stan kursu Go (miękko lub w pełni), zawsze robi backup do postep/archiwum/. Miękki reset czyści tylko postęp i program — kod ucznia zostaje. Pełny reset usuwa wszystko poza skillami i instrukcją. Wymaga jawnego potwierdzenia od ucznia. Użyj gdy uczeń mówi "zresetuj kurs", "zacznij od nowa", "wyczyść postęp", "chcę nowy plan".
---

# Cel

Pozwolić uczniowi zacząć od nowa **bez utraty pracy** — backup robisz **zawsze**, nawet przy "pełnym" resecie. Dane nigdy nie giną, tylko schodzą z głównej ścieżki.

# Dwa tryby

| Tryb      | Co czyści                                        | Co zostawia                                                | Kiedy |
| --------- | ------------------------------------------------ | ---------------------------------------------------------- | ----- |
| Miękki 🧽 | `postep/student.json`, `kurs/program.md`         | Cały kod w `kurs/zadania/`, notatki w `kurs/lekcje/`, projekt w `kurs/projekt/` | Chcesz nowy program kursu, ale zachować dotychczasowe ćwiczenia |
| Pełny 🔥  | Wszystko: postęp, program, notatki, kod, projekt | Tylko `.claude/`, `wiedza/`, `README.md`, `QUICKSTART.md`, `kurs/JAK-PISAC-KOD.md` | Chcesz absolutnie świeży start |

**Domyślny tryb przy "zresetuj kurs" bez doprecyzowania: miękki.**

# Procedura — twardy protokół potwierdzenia

## Krok 1: rozpoznaj intencję

Gdy uczeń mówi "zresetuj kurs", **NIE rób nic od razu**. Najpierw dwa pytania:

1. **"Czemu chcesz zresetować?"** — czasem rozwiązanie jest inne
2. **"Miękki czy pełny?"** — pokaż tabelę powyżej

Alternatywy warte zaproponowania zamiast resetu:
- Zmiana celu/tempa → edytujemy `program.md`, bez resetu
- Powtórka tematu → quiz "słabe punkty", bez resetu
- Wyczyszczenie `do_powtorki` → operacja na `student.json`, bez resetu
- "Chcę zacząć zadania od nowa" → nowy katalog `NN-temat-v2/`, stary zostaje

## Krok 2: pokaż, co konkretnie zniknie

Przed wykonaniem **wypisz dokładnie** ścieżki, które trafią do backupu:

```
Zostaną przeniesione do postep/archiwum/2026-08-30-14-30/:
  - postep/student.json (12 ukończonych lekcji, 8 sesji)
  - kurs/program.md
  [pełny reset dodatkowo]:
  - kurs/lekcje/ (8 plików)
  - kurs/zadania/ (5 katalogów, 14 plików .go + go.mod)
  - kurs/projekt/ (jeśli istnieje)
```

**Ostrzeż wprost, jeśli reset dotyka `kurs/projekt/`** — to zwykle najwięcej pracy w całym kursie:
> "W `kurs/projekt/` masz swój program z modułu 12. Trafi do archiwum, ale i tak upewnijmy się: na pewno pełny reset?"

## Krok 3: poproś o jawne potwierdzenie

Uczeń musi napisać **literalnie**:
- miękki: `tak, miękki reset`
- pełny: `tak, pełny reset`

Inne potwierdzenia ("ok", "tak", "rób") — **odmawiasz** i prosisz o pełną frazę.

## Krok 4: wykonaj reset

### Katalog backupu

```bash
TIMESTAMP=$(date +%Y-%m-%d-%H-%M)
mkdir -p postep/archiwum/$TIMESTAMP
```

### Miękki reset

```bash
mv postep/student.json postep/archiwum/$TIMESTAMP/ 2>/dev/null
mv kurs/program.md postep/archiwum/$TIMESTAMP/ 2>/dev/null
```

### Pełny reset

```bash
mv postep/student.json postep/archiwum/$TIMESTAMP/ 2>/dev/null
mv kurs/program.md postep/archiwum/$TIMESTAMP/ 2>/dev/null
[ -d kurs/lekcje ]  && mv kurs/lekcje  postep/archiwum/$TIMESTAMP/lekcje
[ -d kurs/zadania ] && mv kurs/zadania postep/archiwum/$TIMESTAMP/zadania
[ -d kurs/projekt ] && mv kurs/projekt postep/archiwum/$TIMESTAMP/projekt

# Odtwórz strukturę
mkdir -p kurs/lekcje kurs/zadania
```

**Po pełnym resecie odtwórz `go.mod`** — bez niego uczeń nie uruchomi żadnego zadania:
```bash
cat > kurs/zadania/go.mod <<'EOF'
module kurs/zadania

go 1.22
EOF
```

**Nigdy nie używaj `rm -rf`.** Wyłącznie `mv` do `postep/archiwum/`. To gwarancja, że nawet po "pełnym" resecie dane fizycznie są na dysku.

**Binarki** (`kurs/zadania/*/nazwa`, `*.exe`) możesz pominąć w backupie — odbudowują się z kodu. Ale nie kasuj ich sam; zostaw i powiedz uczniowi, że może usunąć.

### Manifest backupu

W `postep/archiwum/$TIMESTAMP/MANIFEST.md`:

```markdown
# Backup z dnia 2026-08-30 14:30

**Tryb resetu:** miękki | pełny
**Powód podany przez ucznia:** [krótki cytat]
**Stan przed resetem:**
- Ukończonych lekcji: 12 z 38
- Aktualna lekcja: 5.2
- Liczba sesji: 8
- Wersja Go: 1.25.3

**Co jest w tym katalogu:**
- student.json
- program.md
- (przy pełnym: lekcje/, zadania/, projekt/)

Aby przywrócić: skopiuj pliki z tego katalogu z powrotem do kurs/ i postep/.
```

## Krok 5: potwierdź i zaproś do nowego startu

> "Gotowe. Stary stan jest w `postep/archiwum/2026-08-30-14-30/` — jakbyś chciał wrócić, daj znać. Zaczynamy od nowa? Napisz **'ucz mnie Go'**, przejdziemy przez onboarding."

# Przywracanie z backupu

Uczeń mówi "cofnij" / "wróć do poprzedniego stanu":
1. Pokaż listę katalogów w `postep/archiwum/` (najnowszy na górze)
2. Zapytaj, który przywrócić
3. Skopiuj pliki **z powrotem** do `kurs/` i `postep/` — **nie usuwaj** katalogu archiwum
4. Potwierdź odczytem `student.json` (skill: postep)

# Sprzątanie starych backupów

**NIE rób tego automatycznie.** Jeśli uczeń zauważy, że ma 30 katalogów w `archiwum/`, zaproponuj: "Chcesz, żebym usunął backupy starsze niż 30 dni?" — pokaż listę, poproś o literalne potwierdzenie z liczbą, dopiero wtedy usuń.

# Twarde zasady

- **Backup zawsze**, nawet przy pełnym resecie.
- **Tylko `mv`, nigdy `rm -rf`.**
- **Dwa pytania + jawne potwierdzenie** przed wykonaniem. Nie ma szybkiej ścieżki.
- **Nie resetuj `.claude/`** — to konfiguracja agenta.
- **Nie resetuj `wiedza/`** — to baza wiedzy i 38 lekcji, nie stan ucznia. Do odświeżenia bazy służy skill `baza-wiedzy`.
- **Nie resetuj `kurs/JAK-PISAC-KOD.md`, `README.md`, `QUICKSTART.md`** — to instrukcje, nie generowana treść.
- **Po resecie zawsze pokaż ścieżkę backupu.**
- **Po pełnym resecie odtwórz `kurs/zadania/go.mod`** — inaczej pierwsze zadanie nie ruszy.
