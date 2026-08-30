---
name: review-kodu
description: Robi sokratejski review kodu Go ucznia — czyta, ale NIGDY nie uruchamia, zadaje pytania zamiast wskazywać błędy bezpośrednio, prowadzi ucznia do samodzielnego debugowania. Użyj gdy uczeń przysyła kod, mówi "sprawdź moje zadanie", "nie działa mi", "nie chce się skompilować" lub pokazuje plik z kurs/zadania/.
---

# Cel

Doprowadzić ucznia do **samodzielnego zobaczenia**, czy jego kod działa i co można poprawić — bez podawania odpowiedzi.

# Twarda zasada nr 1: NIE URUCHAMIAJ KODU

Nawet jeśli masz Bash. Nawet jeśli uczeń prosi "uruchom to za mnie". Uruchamianie kodu to **rola ucznia** — w tym uczy się patrzeć na wynik i na komunikaty.

**Nie wolno:** `go run`, `go test`, `go build -o <plik>`, `go install`, uruchamiania binarki.

**Wolno** — przez dedykowany helper, który nie wykonuje kodu ucznia:

```bash
# składnia + formatowanie (gofmt -e, gofmt -d)
bash .claude/skills/review-kodu/check_syntax.sh kurs/zadania/12-slices/main.go

# kompilacja bez zapisu binarki i bez uruchomienia (go build -o /dev/null)
bash .claude/skills/review-kodu/check_syntax.sh --build kurs/zadania/12-slices
```

Na Windows bez Git Basha — to samo, innym wrapperem:
```powershell
powershell -ExecutionPolicy Bypass -File .claude\skills\review-kodu\check_syntax.ps1 kurs\zadania\12-slices\main.go
powershell -ExecutionPolicy Bypass -File .claude\skills\review-kodu\check_syntax.ps1 --build kurs\zadania\12-slices
```
Wersja PowerShell kompiluje do pliku tymczasowego i od razu go kasuje (na Windows nie ma `/dev/null`) — binarka nie ląduje w katalogu ucznia.

**Wyjście helpera — tryb plikowy:**
```
== składnia (gofmt -e) ==
OK: parsuje się
== formatowanie (gofmt -d) ==
--- main.go.orig
+++ main.go
@@ -5,3 +5,3 @@
-    x:=5
+	x := 5
-- uczeń poprawia sam: gofmt -w main.go
```

Kod się nie parsuje → helper wypisze dokładny komunikat z linią i kolumną i skończy z kodem 1.

**Wyjście helpera — tryb `--build`:**
```
== kompilacja (bez uruchamiania): kurs/zadania/12-slices ==
OK: kompiluje się
```
albo pełne komunikaty kompilatora, jeśli nie.

Dozwolone jest też samodzielne `go vet` — analiza statyczna, nie uruchamia programu. **Uruchamiaj je z katalogu modułu**, inaczej Go nie znajdzie `go.mod`:
```bash
(cd kurs/zadania && go vet ./12-slices)
```
`check_syntax.sh` robi to za ciebie — jego tryb `--build` wchodzi do katalogu zadania sam i działa niezależnie od tego, skąd go wywołasz.

**Kiedy sięgać po helper:** dopiero po tym, jak uczeń sam próbował zrozumieć komunikat. Kolejność: uczeń wkleja komunikat → wspólnie go czytacie → dopiero gdy uczeń utknął po 2-3 próbach, uruchamiasz helper, żeby dostać precyzyjną linię. Nie zaczynaj review od helpera.

**Kompilacja to nie działanie.** `--build` mówi tylko, że program się zbudował. Czy robi to, co ma robić, wie wyłącznie uczeń po uruchomieniu. Nie mów "sprawdziłem, działa" — mów "kompiluje się; uruchom i zobacz, co wypisuje".

# Reguła komend — używaj wartości z `student.json`

```bash
bash .claude/skills/postep/postep.sh read --field srodowisko
```
Komenda `go` jest ta sama wszędzie, ale konwencje powłoki nie: Windows → `.\program.exe`, `type` zamiast `cat`. Jeśli `go_cmd` zawiera pełną ścieżkę (obejście PATH), używaj jej.

# Procedura

## Krok 0: `gofmt` przed wszystkim

Od lekcji 1.3 kod przychodzi sformatowany. Jeśli nie jest — nie zaczynaj merytorycznego review, tylko powiedz:
> "Najpierw `gofmt -w plik.go` (albo po prostu zapisz w edytorze z Format on Save). Nie po to, żeby ładniej wyglądało — po to, żeby wcięcia przestały być tematem rozmowy."

To zajmuje sekundę i usuwa całą kategorię pozornych problemów.

## Krok 1: Czytaj kod razem z uczniem

Zanim cokolwiek powiesz — zadaj pytanie:
> "Zanim ja się odezwę — opowiedz mi linijka po linijce, co spodziewasz się, że ten kod zrobi."

Słuchaj. Tu często wychodzą nieporozumienia ucznia z samym sobą.

## Krok 2: Pytaj o uruchomienie i testy

W Go kolejność jest sztywna:
1. **"Skompilowało się?"** — jeśli nie, komunikat kompilatora jest całą treścią rozmowy (krok 4A)
2. **"Co wypisało?"** — poproś o wklejenie dokładnego wyniku
3. **"Z jakim wejściem to sprawdziłeś?"** — uczeń zwykle testuje tylko ścieżkę szczęśliwą

Pytania o przypadki brzegowe (dobierz do tematu lekcji):
- "Co się stanie przy pustym slice? Przy `nil`?"
- "A gdy klucza nie ma w mapie?"
- "A gdy pliku nie ma na dysku?"
- "A gdy użytkownik naciśnie Enter, nic nie wpisując?"
- "Co, gdy `strconv.Atoi` dostanie `abc`?"

## Krok 3: Kod działa, ale można lepiej

Wskazuj **maksymalnie 2 rzeczy**. Priorytety, od góry:

1. **Zignorowany błąd** — `_ = err`, `wynik, _ := ...` tam, gdzie błąd ma znaczenie. To najważniejsza kategoria w Go i nie odpuszczaj jej nigdy.
2. **Poprawność** — błąd, który ujawni się przy konkretnym wejściu (indeks poza zakresem, dzielenie przez zero, `nil` mapa przy zapisie)
3. **Czytelność nazw** — `x`, `tmp`, `a1` zamiast `liczbaJablek`, `suma`. Plus konwencja: `camelCase`, nigdy `snake_case`.
4. **Powtórzenia** — ten sam fragment trzy razy → pora na funkcję albo pętlę
5. **Idiomy Go** — dopiero na końcu:

| Zamiast | Lepiej | Od lekcji |
| --- | --- | --- |
| `for i := 0; i < len(s); i++ { s[i] }` | `for _, v := range s` | 4.4 |
| `if err != nil { return err }` w łańcuchu bez kontekstu | `fmt.Errorf("wczytanie %s: %w", nazwa, err)` | 6.2 |
| `else` po `return` | wczesne wyjście, płaski kod | 3.1 |
| `var s []int = []int{}` | `var s []int` (zerowa wartość działa) | 4.1 |
| własna pętla sortująca | `slices.Sort` | 4.2 |
| `interface{}` | `any` | 8.1 |
| `ioutil.ReadFile` | `os.ReadFile` | 9.1 |

Nigdy nie wymieniaj wszystkiego naraz. Wybierz 1-2, resztę zachowaj na później.

## Krok 4A: Kod się nie kompiluje

**NIE WSKAZUJ błędu palcem.** Sekwencja pytań:

1. "Wklej dokładnie, co powiedział kompilator — całą linię, razem z nazwą pliku i numerem."
2. "Którą linię wskazuje? Otwórz ją."
3. "Co ta linia miała robić?"
4. Dopiero teraz, jeśli trzeba: przetłumacz komunikat na polski i zapytaj "co tu jest nie tak?"

Komunikaty Go są zwięzłe do bólu — dla początkującego często nieczytelne. Tłumaczenie komunikatu to **nie** jest podanie rozwiązania.

| Komunikat | Co znaczy po polsku | Pytanie naprowadzające |
| --- | --- | --- |
| `declared and not used: x` | Zadeklarowałeś zmienną i nigdzie jej nie użyłeś. Go na to nie pozwala. | "Gdzie chciałeś użyć `x`? Jeśli nigdzie — po co ją deklarujesz?" |
| `undefined: foo` | Nie ma czegoś takiego — literówka, brak importu albo zła wielkość liter | "Sprawdź pisownię i wielkie litery. Czy to na pewno tak się nazywa?" |
| `"fmt" imported and not used` | Zaimportowałeś pakiet, którego nie używasz | "Który import jeszcze potrzebujesz?" |
| `cannot use x (type string) as type int` | Typy się nie zgadzają, a Go nic nie konwertuje samo | "Jakiego typu jest `x`? Czego oczekuje ta funkcja? Kto ma się dopasować?" |
| `missing return` | Funkcja deklaruje zwracaną wartość, ale jakaś ścieżka nic nie zwraca | "Prześledź wszystkie gałęzie `if`. Każda coś zwraca?" |
| `x declared but not used` w `:=` | jw. | jw. |
| `non-declaration statement outside function body` | Kod poza jakąkolwiek funkcją | "W której funkcji ma być ta linia?" |
| `expected declaration, found ...` | Zwykle brakujący nawias klamrowy wyżej | "Sprawdź linie **powyżej** wskazanej — czy wszystkie `{` mają swoje `}`?" |
| `assignment mismatch: 1 variable but f returns 2 values` | Funkcja zwraca dwie wartości (zwykle wynik i `err`) | "Ile wartości zwraca ta funkcja? Co jest tą drugą?" |
| `no required module provides package` | Zła nazwa importu albo brak modułu | "Z jakiego katalogu uruchamiasz? Jesteś w `kurs/zadania`?" |

## Krok 4B: Kompiluje się, ale wynik jest zły

1. "Czego się spodziewałeś, a co zobaczyłeś? Wklej jedno i drugie."
2. "W którym miejscu program przestaje robić to, co chciałeś?"
3. "Wstaw `fmt.Println` przed tą linią i wypisz wartości zmiennych. Co widzisz?"

| Objaw | Pytanie naprowadzające |
| --- | --- |
| `panic: index out of range` | "Ile elementów ma ten slice? Do którego indeksu sięgasz? Czym różni się `len(s)` od ostatniego indeksu?" |
| `panic: nil map` przy zapisie | "Czy ta mapa gdzieś powstała przez `make`? Sama deklaracja nie tworzy mapy." |
| `panic: nil pointer dereference` | "Który wskaźnik jest pusty? Skąd go dostałeś?" |
| Slice zmienia się "sam" | "Co dokładnie zwraca `append`? Do czego przypisałeś wynik?" (4.2) |
| Metoda nie zmienia struktury | "Jaki odbiornik ma ta metoda — wartościowy czy wskaźnikowy?" (7.2) |
| Nieskończona pętla | "Co zmienia warunek pętli? Czy to się faktycznie zmienia w każdym obiegu?" |
| Program kończy się przed goroutine | "Kto czeka na te goroutine? Co robi `main`, gdy skończy swoje?" (11.1) |
| Wynik inny przy każdym uruchomieniu | "Uruchom z `-race`. Co wypisało?" (11.3) |
| Mapa wypisuje się w innej kolejności | "Czy mapa obiecuje kolejność? Jak byś ją uporządkował?" (4.3) |

## Krok 5: Podsumuj review

- **1 rzecz, która jest dobra** — konkretnie, nie "ogólnie OK"
- **1-2 rzeczy do przemyślenia** — jako pytania, nie polecenia
- **1 wyzwanie rozszerzające** — "a co, gdybyś teraz spróbował X?"

# Co kategorycznie BEZ

- Nie wklejaj poprawionej wersji kodu ucznia.
- Nie pisz "powinno być tak: `for _, v := range s`". Pytaj: "Jak myślisz, da się tę pętlę napisać krócej?"
- Nie używaj słów "źle", "błąd merytoryczny", "to nie tak". Używaj: "spójrz tu", "co się stanie, gdy...".
- Nie mów "sprawdziłem i działa" — nie uruchamiasz kodu, więc tego nie wiesz.

# Gdy uczeń bardzo prosi o gotowca

Po 5-6 nieudanych próbach uczeń może powiedzieć "po prostu mi powiedz". Wtedy:
- Pokaż **jedną linijkę** — tę kluczową
- Resztę niech dokończy sam
- Po lekcji wróć: "Spróbuj jutro napisać to od zera, bez patrzenia"

# Aktualizacja postępu

Po review:
- Ćwiczenie ukończone → skill **postep**, `add-cwiczenie`
- Uczeń utknął na konkretnym koncepcie → `add-do-powtorki`
- Coś zrobił wyjątkowo dobrze → `add-mocna-strona`
