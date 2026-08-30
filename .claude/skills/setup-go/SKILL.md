---
name: setup-go
description: Sprawdza, czy w systemie zainstalowany jest Go 1.22 lub nowszy (macOS / Linux / Windows), prowadzi ucznia przez instalację lub aktualizację, weryfikuje że kompilator działa, ustawia edytor i zapisuje środowisko do student.json. Użyj na początku pierwszej sesji lub gdy uczeń zgłasza, że komenda `go` nie działa.
---

# Cel

Doprowadzić ucznia do stanu, w którym **w jego terminalu** działa `go version` i pokazuje **Go 1.22 lub nowszy**.

**Dlaczego 1.22 to twarda granica:** od tej wersji zmienna pętli `for` powstaje na nowo w każdej iteracji i działa `for i := range 10`. Na starszym Go przykłady z lekcji 5.2 i 11.1 dadzą **cichy, błędny wynik** — nie błąd kompilacji. To najgorszy rodzaj problemu dla początkującego.

# Krok 0: rozpoznaj system operacyjny

**Zawsze najpierw** ustal, na jakim systemie pracuje uczeń — to determinuje instalator i konwencję ścieżek.

```bash
uname -s 2>/dev/null || echo "Windows (lub PowerShell)"
```

- `Darwin` → macOS
- `Linux` → Linux
- `MINGW*` / `MSYS*` / brak `uname` → Windows

Jeśli niejasne — zapytaj wprost: "Na jakim systemie pracujesz: macOS, Linux czy Windows?"

Dodatkowo na macOS warto znać architekturę (wpływa na wybór instalatora):
```bash
uname -m    # arm64 = Apple Silicon (M1-M4), x86_64 = Intel
```

# Krok 1: sprawdź obecność Go

Na **wszystkich** systemach komenda jest ta sama:

```bash
go version
```

**Interpretacja:**
- `go version go1.22.x` lub nowszy → **gotowe**, przejdź do kroku 4
- `go version go1.21.x` lub starszy → **konieczna aktualizacja**, przejdź do kroku 2 (traktuj jak instalację — nowy Go nadpisuje stary)
- `command not found` / `nie jest rozpoznawany` → krok 2

**Częsty przypadek:** Go jest zainstalowane, ale nie ma go w `PATH`. Zanim uznasz, że brakuje — sprawdź typowe lokalizacje:
```bash
ls /usr/local/go/bin/go 2>/dev/null          # macOS/Linux, instalator oficjalny
ls /opt/homebrew/bin/go 2>/dev/null           # macOS, Homebrew na Apple Silicon
ls "$HOME/go/bin" 2>/dev/null                 # katalog binarek użytkownika
```
Jeśli plik istnieje, a `go version` nie działa → problem z `PATH`, patrz krok 3B.

# Krok 2: instalacja — gałąź wg systemu

**Nie instaluj nic sam.** Daj instrukcję, uczeń wykonuje, potem wracacie do kroku 1 w **nowym terminalu** (PATH musi się odświeżyć).

## 2A — macOS

**Opcja 1 — instalator oficjalny (najprościej dla początkującego):**
1. https://go.dev/dl/
2. Pobierz `.pkg` — `darwin-arm64` dla Apple Silicon, `darwin-amd64` dla Intel (sprawdziłeś przez `uname -m`)
3. Kliknij, zainstaluj — trafi do `/usr/local/go`, a `PATH` ustawi instalator
4. **Zamknij i otwórz nowy terminal**

**Opcja 2 — Homebrew (jeśli uczeń już go ma):**
```bash
brew install go
```
Aktualizacja istniejącego:
```bash
brew upgrade go
```

## 2B — Linux

**Instalator oficjalny — zalecany**, bo wersje w repozytoriach dystrybucji bywają stare (a nam potrzeba ≥1.22):

```bash
# 1. Pobierz (podmień numer na aktualny z https://go.dev/dl/)
curl -LO https://go.dev/dl/go1.25.3.linux-amd64.tar.gz

# 2. Usuń starą instalację i rozpakuj nową
sudo rm -rf /usr/local/go
sudo tar -C /usr/local -xzf go1.25.3.linux-amd64.tar.gz

# 3. Dodaj do PATH (bash)
echo 'export PATH=$PATH:/usr/local/go/bin' >> ~/.profile
source ~/.profile
```

To **jedyny wyjątek** od zasady "nie modyfikuj plików konfiguracyjnych ucznia" — i tak wykonuje to uczeń, nie ty. Wyjaśnij mu, co robi każda linia, zanim ją wklei.

Przez menedżer pakietów (**sprawdź wersję po instalacji** — często za stara):
```bash
sudo apt install golang-go      # Ubuntu/Debian
sudo dnf install golang         # Fedora
sudo pacman -S go               # Arch
```

## 2C — Windows

**Rekomendacja: instalator MSI z go.dev**

1. https://go.dev/dl/ → "Microsoft Windows" `.msi` (`windows-amd64`, albo `windows-arm64` na maszynach ARM)
2. Uruchom, klikaj dalej — instalator **sam dodaje Go do PATH**
3. **Zamknij i otwórz NOWY** PowerShell
4. Sprawdź: `go version`

W przeciwieństwie do wielu innych środowisk, na Windows **nie ma tu haczyka z nazwą komendy** — `go` to `go` wszędzie.

Alternatywy: `winget install GoLang.Go` albo `choco install golang` (jeśli uczeń już używa tych menedżerów).

# Krok 3: weryfikacja

## 3A — normalna ścieżka

1. Uczeń **otwiera NOWY terminal** (stary nie zna nowego PATH)
2. Powtarza krok 1: `go version`

## 3B — Go zainstalowane, ale komenda nie działa

Problem z `PATH`. Diagnoza:

```bash
echo $PATH                    # macOS/Linux
$env:Path                     # Windows PowerShell
```

Naprawa (macOS/Linux, uczeń wykonuje sam):
```bash
echo 'export PATH=$PATH:/usr/local/go/bin' >> ~/.zshrc   # zsh (domyślny na macOS)
echo 'export PATH=$PATH:/usr/local/go/bin' >> ~/.bashrc  # bash
```
Potem **nowy terminal**.

**Obejście awaryjne:** jeśli `PATH` nie chce współpracować, zapisz w `student.json` pełną ścieżkę jako `go_cmd` (np. `/usr/local/go/bin/go`) i używaj jej we wszystkich komendach pokazywanych uczniowi. Działa, choć jest niewygodne — wróćcie do naprawy `PATH` przy okazji.

# Krok 4: test "działa"

Utwórz razem z uczniem pierwszy moduł zadań (jeśli jeszcze nie istnieje):

```sh
cd kurs/zadania
cat go.mod
```

Powinno pokazać:
```
module kurs/zadania

go 1.22
```

Test kompilacji **bez uruchamiania czegokolwiek**:
```sh
go build ./...
```

Cisza = wszystko działa. To jednocześnie pierwsza lekcja kultury Go: **narzędzia milczą, gdy jest dobrze**.

> **Uwaga dla agenta:** to jedyny moment, gdy sam wywołujesz `go build` — na pustym module bez kodu ucznia. Później: nigdy.

# Krok 5: sprawdzenie narzędzi towarzyszących

`gofmt` i `go vet` przychodzą razem z Go — nie trzeba nic instalować:

```sh
gofmt --help    # wypisuje użycie
go vet --help   # wypisuje użycie
```

Powiedz uczniowi jednym zdaniem, po co są (rozwiniesz to w lekcji 1.3):
> "`gofmt` formatuje kod — w Go jest jeden słuszny format i nie ma o czym dyskutować. `go vet` szuka podejrzanych miejsc, których kompilator nie uzna za błąd."

# Krok 6: edytor

Rekomendacja dla początkujących (działa na wszystkich systemach):
- **VS Code** (https://code.visualstudio.com/) + rozszerzenie **"Go"** od Google

Przy pierwszym otwarciu pliku `.go` rozszerzenie zapyta o doinstalowanie narzędzi pomocniczych (`gopls` — serwer języka) → **zgódź się**. Da to podpowiedzi, podświetlanie błędów w locie i formatowanie przy zapisie.

Alternatywy:
- **GoLand** (JetBrains, płatny, 30 dni próbnych)
- **Zed**, **Neovim** z `gopls` — dla osób, które już wiedzą, po co
- Jakikolwiek edytor tekstu — **NIE Word, NIE TextEdit w trybie sformatowanym na macOS** (wstawia znaki, których kompilator nie zrozumie)

**Ustawienie warte pięciu sekund:** "Format on Save" w VS Code. Kod formatuje się sam przy każdym zapisie i uczeń nigdy nie walczy z wcięciami.

# Krok 7: instrukcja workflow

**Zawsze** na koniec setupu wskaż uczniowi plik `kurs/JAK-PISAC-KOD.md`:

> "Zanim zaczniemy pierwszą lekcję — otwórz w edytorze plik `kurs/JAK-PISAC-KOD.md` i przeczytaj go. To 5 minut, a wyjaśnia: gdzie zapisywać kod, jak uruchamiać programy, jak czytać komunikaty kompilatora, plus różnice komend między macOS/Linux/Windows. Daj znać, gdy przeczytasz."

Nie idź dalej, dopóki uczeń nie potwierdzi.

# Mapa komend wg systemu — ściąga dla agenta

| Co                        | macOS / Linux                  | Windows PowerShell               |
| ------------------------- | ------------------------------ | -------------------------------- |
| Wersja Go                 | `go version`                   | `go version`                     |
| Uruchom program           | `go run ./01-hello`            | `go run ./01-hello`              |
| Zbuduj binarkę            | `go build -o hello .`          | `go build -o hello.exe .`        |
| Uruchom binarkę           | `./hello`                      | `.\hello.exe`                    |
| Sprawdź ścieżkę Go        | `which go`                     | `Get-Command go`                 |
| Katalog binarek           | `go env GOPATH`                | `go env GOPATH`                  |
| Zmienna środowiskowa      | `export GOOS=linux`            | `$env:GOOS="linux"`              |
| Podgląd pliku             | `cat plik.go`                  | `type plik.go`                   |
| Lista plików              | `ls -l`                        | `dir`                            |
| Bieżący katalog           | `pwd`                          | `pwd`                            |

**Dobra wiadomość:** wszystkie komendy `go ...` są identyczne na każdym systemie. Różnice dotyczą tylko powłoki i binarek (`.exe`, `.\`).

# Twarde zasady

- **Nie uruchamiaj instalatorów** za ucznia. To jego maszyna.
- **Nie modyfikuj** `~/.zshrc`, `~/.bashrc`, profilu PowerShell itp. — pokaż komendę, uczeń ją wykonuje.
- **Po instalacji ZAWSZE nowy terminal** — oszczędzi długich poszukiwań "czemu nie działa".
- **Go < 1.22 → nie zaczynaj kursu.** Cicho błędne wyniki w lekcjach o domknięciach i goroutine to gorsze niż brak Go.
- **`GOPATH` to relikt.** Jeśli uczeń trafi na stary poradnik każący układać kod w `~/go/src/...` — powiedz, że od Go 1.16 obowiązują moduły i kod może leżeć gdziekolwiek.
- **Nie uruchamiaj kodu ucznia** — także tutaj. Jedyne dozwolone wywołania w tym skillu to `go version`, `go env`, `gofmt --help`, `go vet --help` i `go build ./...` na pustym module z kroku 4.

# Zapis środowiska do `student.json`

Po zakończonym setupie ZAWSZE zapisz środowisko (podczas onboardingu może być od razu w `init`, później przez `update-srodowisko`):

```bash
python3 .claude/skills/postep/postep.py update-srodowisko \
  --system "macOS" \
  --go-cmd "go" \
  --go-version "1.25.3" \
  --shell "zsh" \
  --edytor "VS Code"
```

**Mapowanie systemu → wartości:**

| System  | `go_cmd`                    | `shell`      | Uwagi                            |
| ------- | --------------------------- | ------------ | -------------------------------- |
| macOS   | `go`                        | `zsh`        | pełna ścieżka tylko przy kłopocie z PATH |
| Linux   | `go`                        | `bash`/`zsh` | jw.                              |
| Windows | `go`                        | `PowerShell` | binarki z `.exe`, ścieżki z `\`  |
| WSL     | `go`                        | `bash`       | traktuj jak Linux                |

`go_version` zapisuj **bez prefiksu `go`** — sam numer, np. `1.25.3`. Wyciągnięcie z `go version`:
```bash
go version | awk '{print $3}' | sed 's/^go//'
```

Zapis przez `postep.py` jest atomowy — nie ma ryzyka uszkodzenia `student.json`.

# Zwrotka do agenta-rodzica

Po zakończeniu zwróć krótko:
- `OK: Go 1.X.Y na <system>, komenda: <go_cmd>, edytor: <nazwa> — środowisko zapisane`
- `STARY GO: 1.21.x — wymaga aktualizacji przed lekcją 1.2`
- `BLOCKED: <co nie działa>`

**Od tego momentu** WSZYSTKIE skille edukacyjne (lekcja, cwiczenie, review-kodu, quiz) muszą używać `srodowisko.go_cmd` z `student.json` i konwencji ścieżek z `srodowisko.system`.
