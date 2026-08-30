# Jak pisać i uruchamiać kod — instrukcja dla ucznia

Ten dokument odpowiada na trzy pytania:
1. **Gdzie** pisać kod?
2. **W czym** pisać kod (edytor)?
3. **Jak** uruchomić kod, żeby zobaczyć, co robi?

Przeczytaj raz, na początku kursu. Potem wracaj, gdy o czymś zapomnisz.

---

## 1. Gdzie pisać kod — struktura katalogów

Twój kod żyje w katalogu `kurs/zadania/`. Dla każdego ćwiczenia powstaje osobny podkatalog:

```
kurs/
└── zadania/
    ├── go.mod                       ← JEDEN plik na cały kurs, nie ruszaj go
    ├── 01-hello/
    │   └── main.go
    ├── 02-zmienne-a/
    │   └── main.go
    ├── 02-zmienne-b/
    │   └── main.go
    └── 03-petle/
        ├── main.go
        └── ZADANIA.md               ← zapisuje agent
```

### Reguły nazewnictwa

- Katalog: `NN-krotki-temat`, przy kilku zadaniach z jednej lekcji z literą: `-a`, `-b`, `-c`
- Plik: `main.go` (dla ćwiczeń zawsze tak; przy większych programach dojdą kolejne pliki)
- **Numerację i nazwę katalogu poda Ci agent** — nie musisz tego wymyślać sam

### Jeden `main.go` na katalog — czemu?

Każdy plik zaczyna się od `package main` i zawiera funkcję `func main()`. W jednym katalogu może być tylko **jedna** funkcja `main`. Dlatego trzy rozwiązania (🔥 ⭐ ⚡) trafiają do trzech katalogów, nie do trzech plików obok siebie.

### Czym jest `go.mod`

To metryczka projektu: nazwa i minimalna wersja Go. Dzięki niemu Go wie, że wszystkie katalogi obok są częścią jednej całości. **Masz go już gotowego** — nie twórz nowych, dopóki agent nie powie (stanie się to raz, przy projekcie w module 12).

### Eksperymenty na boku

Chcesz coś sprawdzić "na brudno" — poproś agenta o katalog roboczy albo utwórz `kurs/zadania/99-notatnik/main.go`. To plik do niszczenia, nie do oceny.

---

## 2. W czym pisać kod — edytor

### Rekomendacja: VS Code

1. Pobierz z https://code.visualstudio.com/
2. Zainstaluj rozszerzenie **Go** od Google (Extensions → wyszukaj "Go")
3. Przy pierwszym otwarciu pliku `.go` rozszerzenie zapyta o doinstalowanie narzędzi (`gopls`) → **zgódź się**
4. Otwórz cały katalog kursu: `File → Open Folder → claude-agent-golang-course`

### Jedno ustawienie warte pięciu sekund

Włącz **"Format on Save"** (Settings → wyszukaj "format on save"). Od tej chwili Twój kod formatuje się sam przy każdym zapisie i nigdy nie będziesz walczyć z wcięciami.

### Inne opcje

- **GoLand** (JetBrains) — pełne IDE, 30 dni za darmo
- **Zed**, **Neovim** z `gopls` — jeśli już wiesz, po co
- **NIE używaj:** Worda, TextEdit na macOS w trybie sformatowanym, Notatnika Windows w domyślnej konfiguracji — wstawiają znaki, których kompilator nie zrozumie

### Co musi umieć Twój edytor

- Zapisywać w **UTF-8** (dla polskich znaków `ą ę ć`)
- Kolorować składnię Go
- Najlepiej: uruchamiać `gofmt` przy zapisie

---

## 3. Jak uruchomić kod — terminal

### Co to terminal?

Okienko, w którym wpisujesz **komendy tekstowe** zamiast klikać.

| System       | Jak otworzyć                                     |
| ------------ | ------------------------------------------------ |
| **macOS**    | Cmd+Space → "Terminal" → Enter                   |
| **Linux**    | Ctrl+Alt+T (większość dystrybucji)               |
| **Windows**  | Win+X → "Terminal" / "Windows PowerShell"        |

**Dobra wiadomość:** komendy `go ...` są **identyczne na wszystkich systemach**. Różnice dotyczą tylko powłoki:

| W tej instrukcji (macOS/Linux) | Windows PowerShell |
| --- | --- |
| `./zadania` | `.\zadania.exe` |
| `cat plik.go` | `type plik.go` |
| `ls -l` | `dir` |
| `go build -o zadania .` | `go build -o zadania.exe .` |

Agent dopasuje komendy do Twojego systemu — powiedz mu na początku, jakiego używasz.

### Krok po kroku — uruchomienie programu

#### 1. Przejdź do katalogu z zadaniami

```sh
cd ~/Projects/claude-agent-golang-course/kurs/zadania
```

`cd` = "change directory". Sprawdź, gdzie jesteś:
```sh
pwd
```

**To ważne:** wszystkie zadania uruchamiasz **z katalogu `kurs/zadania`**, nie z katalogu głównego kursu i nie z wnętrza `01-hello`.

#### 2. Uruchom program

```sh
go run ./01-hello
```

Zwróć uwagę: `./01-hello` to **katalog**, nie plik. W Go uruchamia się pakiet — czyli wszystkie pliki `.go` w danym katalogu naraz. To jedna z pierwszych rzeczy, które trzeba przestawić w głowie.

To, co program wypisze przez `fmt.Println`, pojawi się **w terminalu**, pod komendą.

#### 3. Coś się nie zgadza? Czytaj komunikat

W Go są **dwa różne momenty**, w których coś może pójść nie tak.

**A. Program się nie skompilował** — nie uruchomił się w ogóle:

```
# kurs/zadania/01-hello
./main.go:6:2: declared and not used: imie
```

Czytaj tak:
- `./main.go:6:2` — plik `main.go`, linia 6, kolumna 2
- `declared and not used: imie` — zadeklarowałeś zmienną `imie` i nigdzie jej nie użyłeś. Go na to nie pozwala.

**B. Program się uruchomił i przerwał w trakcie** (panika):

```
panic: runtime error: index out of range [5] with length 3

goroutine 1 [running]:
main.main()
	/Users/anna/kurs/zadania/04-slices/main.go:9 +0x1d
```

Czytaj tak:
- pierwsza linia — **co** się stało: sięgnąłeś po element numer 5 w czymś, co ma 3 elementy
- linia z nazwą pliku — **gdzie**: `main.go`, linia 9
- reszta (`goroutine 1 [running]`, adresy `+0x1d`) — na razie ignoruj

Komunikaty kompilatora Go są krótkie — czasem zbyt krótkie. **To nie jest katastrofa, to informacja.** Twoją rolą jest najpierw **przeczytać samodzielnie**, a dopiero potem pokazać agentowi.

---

## 4. Cały workflow — od zadania do review

```
┌─────────────────────────────────────────────────────────────┐
│ 1. Agent prowadzi lekcję w czacie Claude Code               │
│    → odpowiadasz na pytania naprowadzające                  │
├─────────────────────────────────────────────────────────────┤
│ 2. Agent daje ćwiczenie ("napisz program, który...")        │
│    → mówi, gdzie zapisać: kurs/zadania/NN-temat/main.go     │
├─────────────────────────────────────────────────────────────┤
│ 3. Otwórz VS Code / terminal obok Claude Code               │
│    → utwórz plik, napisz kod, ZAPISZ (Cmd+S / Ctrl+S)       │
├─────────────────────────────────────────────────────────────┤
│ 4. W terminalu, z katalogu kurs/zadania:                    │
│    go run ./NN-temat                                        │
│    → patrzysz na wynik                                      │
├─────────────────────────────────────────────────────────────┤
│ 5. Sam(a) oceń: czy wynik jest taki, jakiego oczekiwałeś?   │
│    → nie? czytaj komunikat i poprawiaj                      │
├─────────────────────────────────────────────────────────────┤
│ 6. Gdy działa — wracasz do Claude Code i piszesz:           │
│    "skończyłem rozgrzewkę z lekcji 2.1"                     │
│    → wklejasz kod i wynik, agent robi review przez pytania  │
└─────────────────────────────────────────────────────────────┘
```

**Agent nie uruchomi Twojego programu.** To nie ograniczenie techniczne — to część metody. Patrzenie na wynik i czytanie komunikatów jest tym, czego się właśnie uczysz.

**Wskazówka praktyczna:** trzymaj **dwa okna obok siebie** — Claude Code po lewej, terminal/edytor po prawej.

---

## 5. `gofmt` — formatowanie, o którym się nie dyskutuje

W Go istnieje **jeden** poprawny format kodu i narzędzie, które go wymusza:

```sh
gofmt -w 01-hello/main.go     # -w = zapisz zmiany w pliku
gofmt -l .                    # wypisz pliki, które NIE są sformatowane
```

W innych językach ludzie kłócą się o spacje kontra tabulatory. W Go nie ma o czym — `gofmt` decyduje, a wszyscy programiści Go czytają kod w tym samym układzie.

**Praktycznie:** jeśli włączyłeś "Format on Save", nie musisz o tym pamiętać. Jeśli nie — uruchom `gofmt -w` przed pokazaniem kodu agentowi.

`gofmt -l .` milczy = wszystko sformatowane. **W Go cisza narzędzia znaczy „w porządku".**

---

## 6. Najczęstsze pułapki początkujących

### ❌ Uruchamiam z niewłaściwego katalogu

```
go: cannot find main module
```
Jesteś poza `kurs/zadania`. Sprawdź `pwd`, wróć na miejsce.

### ❌ `go run 01-hello` bez `./`

Go potraktuje to jak nazwę pakietu z internetu, nie katalog. Zawsze `./01-hello`.

### ❌ `declared and not used`

Zadeklarowałeś zmienną i nie użyłeś jej. **W Go to błąd, nie ostrzeżenie.** Usuń zmienną albo jej użyj. Wygląda to surowo, ale ratuje przed literówkami: zmienna, której nigdzie nie używasz, to zwykle zmienna, którą chciałeś gdzieś wpisać i wpisałeś inaczej.

### ❌ `"fmt" imported and not used`

To samo dla importów. Usuń niepotrzebny import (albo edytor zrobi to za Ciebie przy zapisie).

### ❌ Zapomniałem zapisać plik przed `go run`

Najczęstszy błąd na świecie. Cmd+S / Ctrl+S **zawsze** przed uruchomieniem.

### ❌ Polskie znaki wypisują się jako krzaczki

Sprawdź, czy edytor zapisuje w UTF-8 (w VS Code widać w prawym dolnym rogu). Same pliki `.go` Go zawsze czyta jako UTF-8.

### ❌ Program nie chce się zatrzymać

**Ctrl+C** przerywa działający program.

### ❌ Zbudowałem binarkę i mam śmieci w katalogu

`go build -o nazwa .` zostawia plik wykonywalny. Możesz go usunąć — odbuduje się w sekundę. `.gitignore` już go pomija.

---

## 7. Komendy terminala — minimum, które Ci wystarczy

| Komenda | Co robi |
| --- | --- |
| `pwd` | Pokaż, w którym katalogu jestem |
| `ls` | Pokaż pliki w bieżącym katalogu |
| `cd nazwa` | Wejdź do podkatalogu |
| `cd ..` | Wyjdź o katalog wyżej |
| `cd ~` | Wróć do katalogu domowego |
| **Ctrl+C** | Przerwij działający program |
| **strzałka w górę** | Powtórz ostatnią komendę |

---

## 8. Komendy Go, które poznasz w tym kursie

| Komenda | Co robi | Od lekcji |
| --- | --- | --- |
| `go version` | Pokaż wersję Go | 1.1 |
| `go run ./NN-temat` | Uruchom program | 1.2 |
| `gofmt -w plik.go` | Sformatuj kod | 1.3 |
| `gofmt -l .` | Wypisz niesformatowane pliki | 1.3 |
| `go build -o nazwa .` | Zbuduj samodzielny program | 12.3 |
| `go test ./...` | Uruchom testy | 10.1 |
| `go vet ./...` | Poszukaj podejrzanych miejsc | 10.3 |
| `go mod init nazwa` | Załóż nowy moduł | 10.2 |
| `go doc fmt.Println` | Dokumentacja, bez internetu | 10.2 |
| `go run -race ./NN` | Wykryj wyścigi danych | 11.3 |

Nie musisz ich pamiętać teraz. Agent poda właściwą, gdy przyjdzie na nią czas.

---

## 9. Nie ma trybu interaktywnego — i co z tego wynika

W niektórych językach można otworzyć konsolę i pisać kod linijka po linijce. **Go tego nie ma** — jest językiem kompilowanym, każdy program musi być kompletny: `package main`, `import`, `func main`.

Wygląda to na utrudnienie, ale ma zaletę: od pierwszej lekcji piszesz **prawdziwe programy**, nie fragmenty w konsoli.

Do szybkich eksperymentów masz dwie drogi:
1. Katalog roboczy (`99-notatnik`) — zmieniasz, uruchamiasz, kasujesz
2. **Go Playground** (https://go.dev/play) — piszesz w przeglądarce, klikasz Run. Wygodne do jednorazowego "co zwróci ta funkcja?". Nie ma dostępu do plików ani sieci.

---

## Pytania? Pisz do agenta

W Claude Code możesz w każdej chwili powiedzieć:

- *"jak mam to uruchomić?"*
- *"nie wiem, gdzie zapisać kod"*
- *"co znaczy ten komunikat?"* (wklej cały)
- *"nie chce się skompilować"*

Agent wróci do odpowiedniego fragmentu tej instrukcji albo poprowadzi Cię krok po kroku.
