# QUICKSTART — jak korzystać z agenta `go-tutor`

Krótki przewodnik dla osoby zaczynającej naukę Go z tym kursem.

> **⚠️ Użytkownicy Windows:** komendy `go ...` są takie same na każdym systemie, ale komendy powłoki już nie. W lekcjach i dokumentach są w wersji **macOS/Linux**. Na Windows zamień:
> - `./program` → `.\program.exe`
> - `cat plik.go` → `type plik.go`
> - `ls -l` → `dir`
> - `export GOOS=linux` → `$env:GOOS="linux"`
>
> Agent tłumaczy komendy automatycznie w trakcie sesji (jeśli w onboardingu zaznaczyłeś Windows). Ta uwaga dotyczy **ręcznego** czytania dokumentów.
>
> Własne narzędzia agenta (zapis postępu, sprawdzanie składni) mają wersje dla Git Basha i dla PowerShella — agent wybiera je sam podczas onboardingu. Ciebie to nie dotyczy: `go run`, `go build` i `gofmt` działają na Windows dokładnie tak samo jak wszędzie.

---

## 🚀 Pierwsze uruchomienie

### 1. Sprawdź, gdzie jesteś
```sh
cd ~/Projects/ITMOBILE-agent-nauka-golang
pwd
```
Musisz być **w tym katalogu** — agent i skille są lokalne (`.claude/`).

### 2. Uruchom Claude Code
```sh
claude
```

### 3. Napisz w czacie
```
ucz mnie Go
```

Agent automatycznie:
1. Wykryje brak `postep/student.json` → uruchomi **onboarding**
2. Sprawdzi, czy masz Go (`go version`) — wymagane **1.22 lub nowsze**
3. Każe Ci przeczytać `kurs/JAK-PISAC-KOD.md` (5 min)
4. Zapyta o imię, cel (praca / narzędzia / hobby / szkoła), tempo (h/tydzień)
5. Wygeneruje **Twój** `kurs/program.md` (38 lekcji dopasowanych do celu)
6. Utworzy `postep/student.json` ze stanem początkowym
7. Zaproponuje rozpoczęcie lekcji 1.1

---

## 📚 Typowa sesja nauki

### Układ dwóch okien obok siebie

```
┌──────────────────────────┬──────────────────────────┐
│                          │                          │
│      Claude Code         │     VS Code + terminal   │
│      (lewa połowa)       │     (prawa połowa)       │
│                          │                          │
│  Tutaj rozmawiasz        │  Tutaj piszesz kod       │
│  z agentem               │  i go uruchamiasz        │
│                          │                          │
└──────────────────────────┴──────────────────────────┘
```

### Przebieg lekcji

1. **W Claude Code:** napisz `kontynuujemy`
2. Agent wczytuje gotową lekcję z `wiedza/lekcje/03.03-petla-for.md` i prowadzi Cię **sokratejsko** (zadaje pytania, Ty odpowiadasz)
3. **Eksperyment** — agent każe Ci coś napisać; otwierasz VS Code, piszesz w `kurs/zadania/03-petle/main.go`
4. **W terminalu**, z katalogu `kurs/zadania`, uruchamiasz: `go run ./03-petle`
5. Wracasz do Claude Code, wklejasz wynik — agent dopytuje
6. **Ćwiczenie** — agent generuje 3 zadania (🔥/⭐/⚡)
7. Piszesz rozwiązania w osobnych katalogach: `03-petle-a/`, `03-petle-b/`
8. Mówisz `sprawdź moje zadanie` → agent robi **sokratejski review** (pyta, nie ocenia z góry)
9. Agent aktualizuje `postep/student.json`, ustawia następną lekcję

---

## 🎯 Najczęstsze komendy

| Co chcę zrobić | Komenda |
| --- | --- |
| Wrócić do nauki po przerwie | `kontynuujemy` |
| Coś przetestować | `daj mi zadanie` |
| Sprawdzić mój kod | `sprawdź moje zadanie` |
| Kod się nie kompiluje | `nie chce się skompilować` + wklej komunikat |
| Program działa źle | `nie działa mi` + wklej kod i wynik |
| Powtórka | `quiz` (3 pyt) / `quiz słabe` |
| Stan postępów | `pokaż postępy` |
| Zapomniałem, co mogę robić | `lista komend` |
| Zacząć od nowa | `zresetuj kurs` |

Pełna lista — napisz `lista komend` w czacie lub zobacz [README.md](README.md).

---

## ⏸ Co zrobić, gdy wracasz po tygodniach

Po prostu napisz `ucz mnie Go` albo `kontynuujemy`:
- Agent czyta `postep/student.json`
- Wita Cię po imieniu, pokazuje ostatnią lekcję
- Jeśli przerwa > 7 dni → **automatycznie** zaproponuje krótki quiz powtórkowy
- Potem ruszacie z bieżącą lekcją

---

## 🔧 Komendy techniczne (rzadziej)

| Sytuacja | Komenda |
| --- | --- |
| Pierwsze uruchomienie, brak Go | `sprawdź Go` |
| Odświeżenie materiałów z repozytoriów | `odśwież bazę wiedzy` |
| Sprawdzenie stanu bazy | `pokaż stan bazy` |
| Cofnięcie resetu | `cofnij reset` |

---

## 📂 Co jest gdzie

```
ITMOBILE-agent-nauka-golang/
│
├── .claude/                      ← agent + 10 skilli (NIE dotykaj)
├── wiedza/                       ← baza wiedzy (38 gotowych lekcji + źródła)
│
├── kurs/
│   ├── JAK-PISAC-KOD.md          ← PRZECZYTAJ NA POCZĄTKU (5 min)
│   ├── program.md                ← Twój plan kursu (powstanie po onboardingu)
│   ├── lekcje/                   ← Twoje notatki z lekcji (agent zapisuje)
│   ├── zadania/                  ← TU PISZESZ SWÓJ KOD
│   │   ├── go.mod                ← metryczka modułu, nie ruszaj
│   │   ├── 01-hello/main.go      ← tworzysz Ty
│   │   └── 03-petle/
│   │       ├── main.go
│   │       └── ZADANIA.md        ← zapisuje agent
│   └── projekt/                  ← Twój program z modułu 12 (osobny moduł Go)
│
└── postep/
    ├── student.json              ← Twój stan (możesz podglądać)
    ├── backups/                  ← automatyczne kopie stanu
    └── archiwum/                 ← backupy po resetach
```

---

## 💡 Mini-przykład pełnej sesji

```
Ty:    ucz mnie Go
Agent: Cześć! Widzę, że jesteś tu pierwszy raz. Uczę przez pytania — i od razu
       uprzedzam: kod uruchamiasz Ty, nie ja. Zacznijmy od sprawdzenia,
       czy masz Go. Otwórz terminal i wpisz: go version

Ty:    go version go1.25.3 darwin/arm64
Agent: Świetnie, 1.25.3 — z zapasem ponad wymagane minimum 1.22.
       Teraz przeczytaj kurs/JAK-PISAC-KOD.md — 5 minut, wyjaśnia gdzie
       piszesz kod i jak go uruchamiasz. Daj znać, gdy skończysz.

Ty:    przeczytałam
Agent: Super. Teraz 4 pytania, żeby dopasować kurs:
       1. Jak masz na imię?
       2. Czemu chcesz się uczyć Go? (praca / narzędzia / hobby / szkoła)
       3. Ile godzin tygodniowo możesz poświęcić?
       4. Programowałaś już kiedyś w innym języku?

Ty:    Anna, narzędzia, 3-5h tygodniowo, nigdy nie programowałam
Agent: [generuje program.md, student.json] Mam dla Ciebie 38 lekcji w 12 modułach,
       tempo ~3 lekcje/tydzień, czyli około 13 tygodni. Na końcu napiszesz
       własne narzędzie wiersza poleceń.
       Zaczynamy od lekcji 1.1 — "Czym jest Go"?

Ty:    tak, zaczynamy
Agent: Zanim powiem cokolwiek o Go — pomyśl o przepisie kulinarnym...
```

---

## ⚠️ Kilka pułapek na start

- **Plik niezapisany przed `go run`** — najczęstszy błąd. Zawsze Cmd+S / Ctrl+S.
- **Uruchamianie z niewłaściwego katalogu** — zadania odpalasz **z `kurs/zadania`**. Sprawdź `pwd`.
- **`go run 01-hello` bez `./`** — Go szuka wtedy pakietu w internecie. Zawsze `./01-hello`.
- **`declared and not used`** — w Go to błąd, nie ostrzeżenie. Usuń nieużywaną zmienną.
- **Nie kopiuj rozwiązań z internetu** — uczysz się przez próbowanie, nie przez wklejanie.
- **Go starsze niż 1.22** — część lekcji da wtedy **cichy, błędny wynik**. Zaktualizuj przed startem.

---

## 🆘 Co zrobić, gdy coś nie działa

| Problem | Co zrobić |
| --- | --- |
| Agent się nie aktywuje | Sprawdź `pwd` — musisz być w katalogu kursu |
| Nie pamiętam komend | `lista komend` |
| Utknąłem na zadaniu | `nie działa mi` + wklej kod + wklej komunikat |
| Kod się nie kompiluje | `nie chce się skompilować` + cały komunikat kompilatora |
| Chcę zacząć od nowa | `zresetuj kurs` (z backupem) |
| Zgubiłem postęp | `cofnij reset` (z `postep/archiwum/`) |
| Materiały wyglądają nieaktualnie | `odśwież bazę wiedzy` |
| Pytanie o Go, nie o kurs | Po prostu zapytaj agenta normalnie |

---

## 📤 Dla autora (publikacja na GitHubie)

Jeśli chcesz udostępnić kurs:

```sh
git init
git add .
git status                # sprawdź, czy nic prywatnego (student.json itd. jest ignorowany)
git commit -m "Initial: agent + baza wiedzy"
gh repo create claude-agent-go-course --public --source=. --push
```

`.gitignore` zadba, żeby Twój postęp i Twój kod nie trafiły do publicznego repozytorium — można bezpiecznie udostępnić **strukturę kursu**, a każdy uczeń sklonuje i ma własny `student.json`.

---

## 🎓 Filozofia kursu — w 3 zdaniach

1. **Sokratejsko, nie wykładowo** — agent zadaje pytania, Ty dochodzisz do odpowiedzi sam(a). Wolniej, ale głębiej.
2. **Nie uruchamiamy kodu za Ciebie** — sam piszesz, sam uruchamiasz, sam czytasz wynik. To część nauki, nie ograniczenie narzędzia.
3. **Twoje tempo** — czas trwania lekcji to wskazówka, nie termin. Możesz wrócić za tydzień, agent wie, gdzie skończyliście.

---

## 🔗 Powiązane dokumenty

- **[README.md](README.md)** — pełna dokumentacja, struktura, lista komend
- **[kurs/JAK-PISAC-KOD.md](kurs/JAK-PISAC-KOD.md)** — workflow pisania i uruchamiania kodu (przeczytaj raz na początku)
- **[wiedza/INDEX.md](wiedza/INDEX.md)** — mapa 38 lekcji
- **[wiedza/AKTUALIZACJE.md](wiedza/AKTUALIZACJE.md)** — co zmieniło się w Go od 1.22 do 1.27

---

**Powodzenia!** 🐹
