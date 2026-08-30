# Aktualizacje — Go 1.22 → 1.27

> **Po co ten plik?** Materiały w `wiedza/zrodlo/` powstały w 2024 (Go 1.21/1.22). Od tego czasu wyszło pięć wydań. Ten plik to **delta**: co się zmieniło, co jest dziś idiomem, czego już nie pisać.

> **Zasada dydaktyczna:** najpierw pokaż uczniowi sposób ze źródła (klasyczny, wszędzie działa), **potem** nowoczesny idiom. Uczeń ma rozumieć oba — kod w internecie jest w większości sprzed 2024.

> **Weryfikacja:** treść sprawdzona z `go.dev/doc/devel/release` oraz notatkami wydań `go1.22`–`go1.27` (stan: sierpień 2026).

---

## Kalendarz wydań

| Wersja | Data | Co wnosi dla początkującego |
| --- | --- | --- |
| 1.22 | 2024-02-06 | **zmiana semantyki zmiennej pętli**, `range` po `int` |
| 1.23 | 2024-08-13 | `range` po funkcji (iteratory), funkcje iteracyjne w `slices`/`maps` |
| 1.24 | 2025-02-11 | `os.Root`, `tool` w `go.mod`, `t.Context()`, `b.Loop()` |
| 1.25 | 2025-08-12 | `sync.WaitGroup.Go()`, `testing/synctest`, `go doc -http` |
| 1.26 | 2026-02-10 | `new(wyrażenie)`, `errors.AsType`, `go fix` z modernizatorami |
| 1.27 | 2026-08-19 | pakiet `uuid`, `encoding/json` przepisany na silnik v2 (bez zmiany zachowania), `strings.CutLast` |

**Wersja minimalna dla kursu: 1.22.** Wszystko poniżej ma inną semantykę pętli — to jest twarda granica, nie preferencja.

---

## `[ogólne]` — co powiedzieć na starcie kursu

- Aktualna wersja Go to **1.27** (sierpień 2026). Uczeń instaluje najnowszą.
- Go trzyma **obietnicę zgodności Go 1**: kod z 2015 kompiluje się dziś. Dlatego stare tutoriale są nadal poprawne — po prostu czasem nieidiomatyczne.
- Wersja w `go.mod` (`go 1.22`) to **nie** wersja kompilatora, tylko deklaracja "ten kod oczekuje semantyki 1.22". Kompilator 1.27 to uszanuje.
- `GOPATH` **nie jest już potrzebny** (od 1.16 moduły są domyślne). Jeśli tutorial każe ustawiać `GOPATH` i trzymać kod w `~/go/src` — jest przestarzały, ignoruj.

---

## `[moduł 1]` — środowisko

- Instalacja: `go.dev/dl` (oficjalny instalator), macOS `brew install go`, Linux — pakiet dystrybucji bywa stary, lepiej z go.dev.
- Edytor: **VS Code + oficjalne rozszerzenie Go** (formatuje przy zapisie, podpowiada, uruchamia testy). Źródło wspomina GoLand — nadal dobre, ale płatne.
- `go doc -http :8080` (od 1.25) — lokalna dokumentacja w przeglądarce, bez internetu. Warto pokazać w 1.3.
- `gofmt` vs `go fmt`: `gofmt -w plik.go` działa na pliku, `go fmt ./...` na całym pakiecie. W kursie używamy `gofmt`.

---

## `[moduł 2]` — podstawy

- `min(a, b)` i `max(a, b)` są **wbudowane** od 1.21 — nie trzeba pisać własnej funkcji ani sięgać po `math.Min` (która działa tylko na `float64`).
- `math/rand/v2` (od 1.22): `rand.IntN(10)` zamiast `rand.Intn(10)`, brak potrzeby `rand.Seed()` — ziarno jest losowe automatycznie. `rand.Seed()` w starym pakiecie od 1.24 **nic nie robi**.
- Jeśli uczeń zobaczy w tutorialu `rand.Seed(time.Now().UnixNano())` — to relikt, dziś zbędny.

---

## `[moduł 3]` — pętle ⚠️ NAJWAŻNIEJSZA ZMIANA

### `range` po liczbie (1.22)

```go
// nowe — od 1.22
for i := range 5 {
    fmt.Println(i) // 0 1 2 3 4
}

// klasyczne — działa zawsze
for i := 0; i < 5; i++ {
    fmt.Println(i)
}
```

Ucz **obu**. Klasyczna forma jest w każdym kodzie, jaki uczeń zobaczy.

### Zmienna pętli jest nowa w każdej iteracji (1.22)

To zmiana **semantyki**, nie składni — ten sam kod daje inny wynik na 1.21 i 1.22:

```go
funcs := []func(){}
for i := range 3 {
    funcs = append(funcs, func() { fmt.Println(i) })
}
for _, f := range funcs { f() }
// Go 1.21: 2 2 2
// Go 1.22+: 0 1 2
```

**Konsekwencja dla kursu:** stare tutoriale każą pisać `i := i` na początku ciała pętli ("shadowing"). Od 1.22 to zbędne. Jeśli uczeń to zobaczy — wyjaśnij, że to obejście nieistniejącego już problemu. Temat wraca w 5.2 (domknięcia) i 11.1 (goroutine w pętli).

### `range` po funkcji (1.23)

Iteratory (`iter.Seq`) — **poza kursem**. Uczeń zobaczy je w `slices.All()`. Wystarczy zdanie: "da się iterować po własnej funkcji, dojdziesz do tego później".

---

## `[moduł 4]` — kolekcje

Pakiety `slices` i `maps` są w bibliotece standardowej (od 1.21). Źródła z 2024 każą pisać pętle ręcznie — dziś jest krócej:

| Zadanie | Ręcznie (źródło) | Dziś |
| --- | --- | --- |
| sortowanie | `sort.Ints(s)` | `slices.Sort(s)` |
| szukanie | pętla `for` | `slices.Contains(s, x)` / `slices.Index(s, x)` |
| porównanie slice'ów | pętla | `slices.Equal(a, b)` |
| sklejenie | `append(a, b...)` | `slices.Concat(a, b, c)` (1.22) |
| max/min z slice'a | pętla | `slices.Max(s)` / `slices.Min(s)` |
| klucze mapy | pętla + `append` | `slices.Sorted(maps.Keys(m))` |
| domyślna wartość | `if x == "" { x = y }` | `cmp.Or(x, y)` (1.22) |

**Kolejność w lekcji:** najpierw pętla ręczna (uczeń ma zrozumieć, co się dzieje), na końcu "a teraz krótko".

---

## `[moduł 6]` — błędy

- `fmt.Errorf("...: %w", err)` + `errors.Is` / `errors.As` — standard od 1.13, w źródłach jest.
- **`errors.Join(err1, err2)`** (1.20) — łączenie wielu błędów w jeden. Przydatne przy walidacji formularza.
- **`errors.AsType[T](err)`** (1.26) — bezpieczniejsza wersja `errors.As` bez wskaźnika na zmienną:
  ```go
  if pe, ok := errors.AsType[*os.PathError](err); ok { ... }
  ```
- **`log/slog`** (1.21) — logowanie strukturalne, zastępuje surowy `log` w nowych projektach:
  ```go
  slog.Info("plik zapisany", "nazwa", f, "bajtow", n)
  ```
  W kursie: `log` w 6.3, `slog` jako wzmianka "gdy program rośnie".

---

## `[moduł 8]` — interfejsy

- **`any` zamiast `interface{}`** (alias od 1.18). Piszemy `any`. `interface{}` w kodzie ucznia nie jest błędem, ale to stary zapis.
- `errors.Is`/`As` (moduł 6) to praktyczne zastosowanie type switch — warto połączyć.

---

## `[moduł 9]` — pliki i dane

- **`os.ReadFile` / `os.WriteFile`** (od 1.16) zamiast `ioutil.ReadFile` / `ioutil.WriteFile`. Pakiet `ioutil` jest **przestarzały** — jeśli źródło go używa, tłumacz na `os`.
- **`os.Root`** (1.24, rozszerzony w 1.25) — operacje na plikach ograniczone do jednego katalogu, nie da się z niego "uciec" przez `../`. Wzmianka przy bezpieczeństwie, nie ćwiczenie.
- **`encoding/json`** stoi od 1.27 na nowym silniku (v2): ten sam interfejs, to samo zachowanie, wyraźnie lepsza wydajność. Surowsze reguły — odrzucanie zduplikowanych kluczy i niepoprawnego UTF-8 — to domena **osobnego pakietu `encoding/json/v2`**, po który trzeba sięgnąć świadomie. Kod ucznia i pliki, które dotąd działały, działają dalej.
- `strings.Lines(s)`, `strings.SplitSeq`, `strings.FieldsSeq` (1.24) — iteracja po liniach bez wczytywania całej tablicy.
- `strings.CutLast` (1.27) — obok istniejącego `strings.Cut`; przydatne do wycinania rozszerzenia pliku.

---

## `[moduł 10]` — testy i narzędzia

- **`t.Context()`** (1.24) — kontekst anulowany po zakończeniu testu.
- **`t.Chdir(dir)`** (1.24) — zmiana katalogu na czas testu, sprzątanie automatyczne.
- **`b.Loop()`** (1.24) — nowa forma benchmarku, zastępuje `for i := 0; i < b.N; i++`.
- **`testing/synctest`** (stabilne w 1.25) — testowanie kodu współbieżnego z wirtualnym zegarem, bez `time.Sleep` w testach. Wzmianka w 11.3.
- **`go fix ./...`** (przebudowane w 1.26) — automatyczna modernizacja kodu do nowych idiomów. Świetne narzędzie dydaktyczne: uczeń pisze po staremu, `go fix` pokazuje nowszy zapis.
- **`tool` w `go.mod`** (1.24) — deklarowanie narzędzi projektu; zastępuje hack z plikiem `tools.go`.
- `go mod init` w 1.26+ wpisuje niższą wersję minimalną niż zainstalowany kompilator — to celowe, dla zgodności.

---

## `[moduł 11]` — współbieżność

- **`wg.Go(func(){...})`** (1.25) — zastępuje trzyliniowy rytuał:
  ```go
  // klasycznie
  wg.Add(1)
  go func() { defer wg.Done(); praca() }()

  // od 1.25
  wg.Go(func() { praca() })
  ```
  **Ucz najpierw klasycznej formy** — uczeń musi rozumieć, co `Add`/`Done` liczą. `wg.Go` pokaż na końcu 11.1 jako nagrodę.
- Pętla + goroutine: od 1.22 `go func(){ fmt.Println(i) }()` w pętli działa "jak człowiek się spodziewa". Przed 1.22 wymagało przekazania `i` jako argumentu. Stare tutoriale są tu mylące.
- `GOMAXPROCS` (1.25) respektuje limity CPU kontenera — poza kursem, ale wyjaśnia różnice wydajności w Dockerze.

---

## `[moduł 12]` — projekt i ekosystem

- **Cross-compile bez dodatkowych narzędzi:**
  ```bash
  GOOS=windows GOARCH=amd64 go build -o narzedzie.exe
  GOOS=linux   GOARCH=arm64 go build -o narzedzie-linux-arm64
  ```
  To mocny moment dydaktyczny — uczeń robi binarkę dla innego systemu jedną komendą.
- **`uuid` w bibliotece standardowej** (1.27) — wcześniej wymagało `github.com/google/uuid`. Jeśli tutorial każe instalować tę zależność, na 1.27 już nie trzeba.
- Biblioteki wymienione w `zrodlo/92-biblioteki.md` żyją: Gin, Fiber, Cobra, GORM. Wersje sprawdzać na `pkg.go.dev`, nie w źródle.
- Do sprawdzenia aktualności zależności: `go list -m -u all`, aktualizacja `go get -u ./...` + `go mod tidy`.

---

## Czego już nie uczymy (przestarzałe w źródłach)

| Konstrukcja ze źródła | Zamiast tego | Od kiedy |
| --- | --- | --- |
| `ioutil.ReadFile` / `WriteFile` | `os.ReadFile` / `os.WriteFile` | 1.16 |
| `interface{}` | `any` | 1.18 |
| `rand.Seed(...)` | nic (ziarno automatyczne) | 1.20 / 1.24 |
| `sort.Ints`, `sort.Strings` | `slices.Sort` | 1.21 |
| `i := i` na starcie pętli | zbędne | 1.22 |
| `go func(i int){...}(i)` tylko dla kopii `i` | zbędne | 1.22 |
| `GOPATH`, `~/go/src` | moduły (`go mod init`) | 1.16 |
| plik `tools.go` | `tool` w `go.mod` | 1.24 |

---

## Jak to weryfikować później

Ten plik zestarzeje się jak każdy inny. Aktualizacja:

1. `go.dev/doc/devel/release` — lista wydań z datami
2. `go.dev/doc/go1.NN` — notatki konkretnego wydania
3. `go version` u ucznia — czy w ogóle ma wersję, o której mówimy
4. Skill `baza-wiedzy` odświeża `zrodlo/`, ale **nie ten plik** — ten pisze się ręcznie w trybie autora
