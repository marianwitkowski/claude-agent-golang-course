# Indeks bazy wiedzy — mapowanie źródeł na lekcje

> **Po co ten plik?** Materiały w `wiedza/zrodlo/` to 31 plików wykładowych napisanych dla czytelnika, który zna już inny język programowania. Kurs prowadzony przez agenta jest dla **kompletnie początkujących** — potrzebuje mniejszych jednostek (30–60 min) i innej kolejności. Ten plik mówi: który fragment którego źródła odpowiada której lekcji.

> **⚠️ Uwaga o źródłach:** pliki w `zrodlo/` miejscami zestawiają Go z innym językiem („tam napisałbyś..."). W lekcjach dla początkującego **te porównania są wycięte** — bierzemy tylko treść o Go. Nie cytuj ich uczniowi.

> **⚠️ Komendy w lekcjach** są w wersji macOS/Linux (ścieżki z `/`). Na Windows separator to `\`, a sama komenda `go` jest identyczna na każdym systemie. Agent tłumaczy ścieżki w trakcie sesji.

> **Jak agent z tego korzysta:**
> 1. Przy generowaniu programu (`program-kursu`) — bierze listę lekcji stąd
> 2. Przy prowadzeniu lekcji (`lekcja`) — **najpierw** gotowy plik `wiedza/lekcje/NN.MM-*.md`, **potem** `zrodlo/` + `AKTUALIZACJE.md`
> 3. Gdy uczeń pyta o temat — szuka tu pasującej lekcji

> **Status gotowych lekcji sokratejskich:** ✅ **38/38** w `wiedza/lekcje/` (moduły 1-12 kompletnie pokryte)

---

## Katalogi bazy

| Katalog | Zawartość | Rola |
| --- | --- | --- |
| `wiedza/zrodlo/` | 31 md z repozytorium źródłowego | kanon merytoryczny (składnia, semantyka) |
| `wiedza/przyklady/` | 13 md z `marianwitkowski/golang-examples` | materiał na ćwiczenia i projekt końcowy |
| `wiedza/przyklady/kod/` | 24 pliki `.go` z `marianwitkowski/golang20230314` | minimalne działające przykłady do eksperymentów |
| `wiedza/lekcje/` | 38 lekcji sokratejskich | **scenariusze prowadzenia** — to czytasz w pierwszej kolejności |
| `wiedza/AKTUALIZACJE.md` | delta Go 1.22 → 1.27 | źródła są z 2024, to je aktualizuje |

---

## Mapowanie źródeł → lekcje

### Moduł 1 — Wprowadzenie i środowisko

| Lekcja | Temat | Źródło | Aktualizacja |
| --- | --- | --- | --- |
| 1.1 | Czym jest Go i do czego służy | `01-dlaczego-go.md`, `02-podstawy-go.md` (bez porównań do innych języków) | `[ogólne]` |
| 1.2 | Pierwszy program — `go run`, `package main`, `func main` | `11-pierwszy-program.md` | `[moduł 1]` |
| 1.3 | Edytor, terminal, `gofmt` — workflow ucznia | `11-pierwszy-program.md`, `71-skladnia.md` (tylko część o formatowaniu) | `[moduł 1]` |

### Moduł 2 — Podstawy języka

| Lekcja | Temat | Źródło | Aktualizacja |
| --- | --- | --- | --- |
| 2.1 | Zmienne i typy (`var`, `:=`, int/float64/string/bool) | `12-typy-danych.md` | `[moduł 2]` |
| 2.2 | Operatory i wyrażenia; brak konwersji niejawnej | `13-operatory-wyrazenia.md` | — |
| 2.3 | Wypisywanie: `fmt.Println`, `fmt.Printf`, czasowniki formatu | `14-stdin-stdout.md` | `[moduł 2]` |
| 2.4 | Wejście od użytkownika — `bufio.Scanner` | `14-stdin-stdout.md` | `[moduł 2]` |

### Moduł 3 — Decyzje i powtórzenia

| Lekcja | Temat | Źródło | Aktualizacja |
| --- | --- | --- | --- |
| 3.1 | `if` / `else if` / `else` + instrukcja inicjująca | `21-instrukcje-warunkowe.md` | — |
| 3.2 | `switch` — bez `break`, z `case` wielokrotnym | `21-instrukcje-warunkowe.md`, `kod/Dzien01-06-instr-switch.go` | — |
| 3.3 | `for` — jedyna pętla Go, trzy formy | `22-petle.md`, `kod/Dzien01-05-petla-for.go` | `[moduł 3]` — `range` po liczbie |

### Moduł 4 — Kolekcje

| Lekcja | Temat | Źródło | Aktualizacja |
| --- | --- | --- | --- |
| 4.1 | Tablice i slices — różnica | `12-typy-danych.md`, `kod/Dzien01-09-tablice.go`, `kod/Dzien01-10-slices.go` | — |
| 4.2 | `append`, `len`, `cap`, kopiowanie i pułapka współdzielenia | `12-typy-danych.md` | `[moduł 4]` — pakiet `slices` |
| 4.3 | Mapy — klucz→wartość, idiom `v, ok :=` | `kod/Dzien01-11-maps.go` | `[moduł 4]` — pakiet `maps` |
| 4.4 | `range` po slice i mapie; losowa kolejność map | `22-petle.md` | `[moduł 4]` |

### Moduł 5 — Funkcje

| Lekcja | Temat | Źródło | Aktualizacja |
| --- | --- | --- | --- |
| 5.1 | `func`, parametry, wiele wartości zwracanych | `23-funkcje.md`, `kod/Dzien01-07-funkcje.go` | — |
| 5.2 | Wariadyczne, funkcje jako wartości, domknięcia | `23-funkcje.md`, `kod/Dzien02-01-funkcje-parametry-nazwane.go` | `[moduł 5]` — zmienna pętli 1.22 |
| 5.3 | `defer`, zasięg, własny pakiet w projekcie | `kod/Dzien01-08-defer-func.go`, `63-moduly.md` | — |

### Moduł 6 — Błędy zamiast wyjątków

| Lekcja | Temat | Źródło | Aktualizacja |
| --- | --- | --- | --- |
| 6.1 | `error` to zwykła wartość — `if err != nil` | `64-wyjatki.md` | — |
| 6.2 | `errors.New`, `fmt.Errorf`, `%w`, `errors.Is`/`As` | `64-wyjatki.md` | `[moduł 6]` — `errors.Join`, `AsType` |
| 6.3 | `panic` / `recover`, pakiet `log` | `64-wyjatki.md`, `15-logowanie.md`, `kod/Dzien02-05-log.go` | `[moduł 6]` — `log/slog` |

### Moduł 7 — Struktury i metody

| Lekcja | Temat | Źródło | Aktualizacja |
| --- | --- | --- | --- |
| 7.1 | `struct` — własny typ danych | `31-struktury.md`, `kod/Dzien01-12-struct.go` | — |
| 7.2 | Metody i wskaźniki — kiedy `*T`, kiedy `T` | `31-struktury.md` | — |
| 7.3 | Osadzanie (embedding), wartości zerowe, konstruktory `NewX` | `31-struktury.md` | — |

### Moduł 8 — Interfejsy

| Lekcja | Temat | Źródło | Aktualizacja |
| --- | --- | --- | --- |
| 8.1 | Interfejs jako kontrakt — implementacja niejawna | `32-interfejsy.md`, `kod/Dzien02-02-interfejs-przyklad.go`, `03`, `04` | — |
| 8.2 | `Stringer`, `io.Writer`, `any`, type switch | `32-interfejsy.md` | `[moduł 8]` — `any` zamiast `interface{}` |

### Moduł 9 — Pliki i dane

| Lekcja | Temat | Źródło | Aktualizacja |
| --- | --- | --- | --- |
| 9.1 | Czytanie i pisanie plików (`os`, `bufio`) | `51-obsluga-plikow.md` | `[moduł 9]` — `os.ReadFile`, `os.Root` |
| 9.2 | JSON — `Marshal`, `Unmarshal`, tagi pól | `przyklady/golang-json.md` | `[moduł 9]` — json/v2 |
| 9.3 | Argumenty CLI — `os.Args`, pakiet `flag` | `81-aplikacja-cli.md`, `przyklady/golang-cli.md` | — |

### Moduł 10 — Testy i narzędzia

| Lekcja | Temat | Źródło | Aktualizacja |
| --- | --- | --- | --- |
| 10.1 | `go test`, funkcje `TestXxx`, testy tabelkowe | `61-testowanie.md`, `kod/Dzien02-09-http-test.go` | `[moduł 10]` — `t.Context`, `b.Loop` |
| 10.2 | `go mod` — moduły i zależności zewnętrzne | `63-moduly.md`, `73-zaleznosci.md` (bez porównań do pip) | `[moduł 10]` — `go mod tidy`, `tool` |
| 10.3 | `go vet`, `gofmt`, debugowanie przez wypisywanie i `delve` | `62-debugowanie.md` | `[moduł 10]` — `go fix` |

### Moduł 11 — Współbieżność (łagodnie)

| Lekcja | Temat | Źródło | Aktualizacja |
| --- | --- | --- | --- |
| 11.1 | Goroutine — "rób to obok" + `sync.WaitGroup` | `33-channels-goroutine.md`, `41-prog-wspolbiezne1.md`, `kod/Dzien02-06-goroutune.go` | `[moduł 11]` — `wg.Go()` |
| 11.2 | Kanały — komunikacja między goroutine | `33-channels-goroutine.md`, `42-prog-wspolbiezne2.md` | — |
| 11.3 | `select`, `sync.Mutex`, wyścigi i `-race` | `42-prog-wspolbiezne2.md` | `[moduł 11]` |

### Moduł 12 — Projekt i dalsze kroki

| Lekcja | Temat | Źródło | Aktualizacja |
| --- | --- | --- | --- |
| 12.1 | Wybór projektu CLI i rozpisanie go na kroki | `81-aplikacja-cli.md`, `93-cwiczenia.md` | — |
| 12.2 | Implementacja — od szkieletu do działania | `81-aplikacja-cli.md`, `przyklady/golang-passgen.md`, `golang-shortlinks.md` | — |
| 12.3 | Testy, README, `go build` i dystrybucja binarki | `61-testowanie.md`, `63-moduly.md` | `[moduł 12]` — cross-compile |
| 12.4 | Mapa ekosystemu — co dalej (web, bazy danych, mikroserwisy) | `82-aplikacja-web.md`, `83-aplikacja-bazadanych.md`, `84-mikroserwisy.md`, `91-zasoby.md`, `92-biblioteki.md` | `[moduł 12]` |

---

## Podsumowanie liczbowe

| Moduł | Lekcje |
| --- | --- |
| 1. Wprowadzenie i środowisko | 3 |
| 2. Podstawy języka | 4 |
| 3. Decyzje i powtórzenia | 3 |
| 4. Kolekcje | 4 |
| 5. Funkcje | 3 |
| 6. Błędy zamiast wyjątków | 3 |
| 7. Struktury i metody | 3 |
| 8. Interfejsy | 2 |
| 9. Pliki i dane | 3 |
| 10. Testy i narzędzia | 3 |
| 11. Współbieżność | 3 |
| 12. Projekt i dalsze kroki | 4 |
| **Razem** | **38** |

**Ten plik jest źródłem prawdy dla liczby 38.** Jeśli inny plik podaje inną liczbę — to błąd dokumentacji.

---

## Czego w kursie nie ma (świadome decyzje)

| Temat | Dlaczego pominięty | Gdzie wspomniany |
| --- | --- | --- |
| Generyki (`[T any]`) | Początkujący nie ma jeszcze problemu, który generyki rozwiązują | 12.4 — wzmianka |
| Aplikacja webowa (`net/http`, Gin, Fiber) | Wymaga rozumienia HTTP; osobny kurs | 12.4 |
| Bazy danych (`database/sql`, GORM) | Wymaga SQL | 12.4 |
| Mikroserwisy, gRPC, Kubernetes | Poziom architektoniczny | 12.4 |
| `context.Context` | Sensowny dopiero przy sieci i timeoutach | 11.3 — jedno zdanie |
| Refleksja, `unsafe`, cgo | Nigdy dla początkującego | — |
| Zestawienia Go z innym językiem (`71`, `72`, `73`) | Uczeń nie zna żadnego innego języka | — |
