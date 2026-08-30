---
name: quiz
description: Przeprowadza krótki quiz powtórkowy (3-7 pytań) z ukończonych przez ucznia lekcji Go. Trzy tryby — szybki (3 pyt), pełny (5-7 pyt), słabe punkty (z do_powtorki). Pyta jedno pytanie naraz, czeka na odpowiedź, daje sokratejski feedback. Użyj gdy uczeń mówi "quiz", "powtórka", "sprawdź mnie", między lekcjami lub gdy przerwa była >7 dni.
---

# Cel

Utrwalić wiedzę z **ukończonych** lekcji przez krótkie, interaktywne pytania. Quiz to **diagnoza**, nie egzamin — chodzi o wyłapanie luk, nie o ocenianie.

# Trzy tryby

| Tryb            | Liczba pytań | Z jakich lekcji                         | Kiedy                                 |
| --------------- | ------------ | --------------------------------------- | ------------------------------------- |
| Szybki ⚡       | 3            | 1-2 ostatnio ukończone lekcje           | Rozgrzewka na początku sesji          |
| Pełny 📋        | 5-7          | Wszystkie ukończone lekcje (losowy mix) | Co kilka lekcji, na życzenie ucznia   |
| Słabe punkty 🎯 | 3-5          | Tematy z `do_powtorki` w `student.json` | Gdy `do_powtorki` ma ≥3 pozycje       |

Domyślny tryb przy "quiz" bez doprecyzowania: **szybki**.

# Procedura

## Krok 1: wybór zakresu

1. Odczytaj `postep/student.json` (skill: **postep**)
2. Sprawdź `ukonczone_lekcje` i `do_powtorki`
3. Jeśli `ukonczone_lekcje` ma <2 pozycje → powiedz, że za wcześnie na quiz, zaproponuj lekcję
4. Wybierz tryb (z prośby ucznia lub z kontekstu)

## Krok 2: dobór pytań

Mieszaj **3 rodzaje** pytań:

### A. Przewidywanie wyniku ("co wypisze ten kod?")
```go
s := []int{1, 2, 3}
t := s[:2]
t = append(t, 99)
fmt.Println(s)
```
Sprawdza rozumienie, nie pamięć.

### B. Konceptualne ("dlaczego / kiedy")
> Kiedy metoda potrzebuje odbiornika wskaźnikowego, a kiedy wystarczy wartościowy?

### C. "Znajdź błąd" / "popraw"
```go
// Kompiluje się bez zastrzeżeń. Co się stanie po uruchomieniu?
func main() {
    var m map[string]int
    m["a"] = 1
    fmt.Println(m["a"])
}
```
**Poprawna odpowiedź:** program panikuje — `panic: assignment to entry in nil map`. Sama deklaracja `var m map[string]int` daje mapę `nil`: czytać z niej wolno (dostaniesz wartość zerową), ale zapisać już nie. Mapa powstaje dopiero przez `make` albo literał.

> **Uwaga dla agenta:** nie uruchamiasz kodu, więc odpowiedzi z tego banku znasz wyłącznie stąd. Jeśli uczeń twierdzi co innego niż napisano — poproś o wklejenie wyniku, zanim go poprawisz. Gdy wynik ucznia przeczy temu plikowi, **rację ma wynik**; zgłoś rozbieżność użytkownikowi zamiast upierać się przy banku.

**Proporcja w pełnym quizie:** ~40% A, ~30% B, ~30% C.

**Pytania typu A pisz tak, żeby dało się je rozstrzygnąć w głowie.** Jeśli uczeń musi uruchomić kod, żeby odpowiedzieć — to nie jest pytanie quizowe, tylko ćwiczenie.

## Krok 3: prowadzenie quizu

**Jedno pytanie naraz.** Nie wrzucaj pięciu w jednej wiadomości.

Schemat dla każdego pytania:
1. Podaj pytanie z numerem: "Pytanie 2/5"
2. **Poczekaj** na odpowiedź
3. Po odpowiedzi:
   - **Poprawna** → krótkie potwierdzenie + pytanie pogłębiające ("A gdyby `t` powstało przez `append` zamiast `s[:2]`?")
   - **Częściowo** → naprowadzenie ("Blisko. Czy `t` i `s` dzielą tę samą tablicę pod spodem?")
   - **Błędna** → NIE podawaj odpowiedzi, naprowadź pytaniem. Po 2 nieudanych próbach pokaż odpowiedź i dopisz temat do `do_powtorki`
4. Następne pytanie

**Bez punktacji po każdym pytaniu** — to nie test.

## Krok 4: podsumowanie

- Ile było **na pewno OK**, ile **z pomocą**, ile **do powtórki**
- Wymień konkretnie 1-2 tematy do utrwalenia
- Uczeń zaliczył temat z `do_powtorki` → skill **postep**, `remove-do-powtorki`
- Nowe luki → `add-do-powtorki`
- Zaproponuj następny krok: "Wracamy do lekcji X" albo "Krótkie ćwiczenie na [temat]?"

# Bank pytań — przykłady wg modułów

Pytania **generuj na żywo** pod to, co uczeń przerobił. Poniżej szablony jako inspiracja.

## Moduł 1-2 (podstawy, zmienne, typy)
- A: Co wypisze `fmt.Println(7 / 2)`? A `fmt.Println(7.0 / 2)`?
- B: Czym różni się `var x int` od `x := 0`? Kiedy której formy użyć?
- C: Czemu to się nie kompiluje? `wiek := "30"; fmt.Println(wiek + 5)`

## Moduł 3 (decyzje i pętle)
- A: Ile razy wykona się `for i := 0; i < 10; i += 3 { }`?
- B: Czemu w Go jest tylko jedna pętla, skoro inne języki mają trzy?
- C: Co jest nie tak? `for i := 0; i < 10; i++ { fmt.Println(i) } fmt.Println(i)`

## Moduł 4 (kolekcje)
- A: Co wypisze? `s := make([]int, 3); s = append(s, 1); fmt.Println(len(s))`
- B: Czym różni się `len` od `cap`?
- C: Czemu `var m map[string]int; m["a"] = 1` wywala program, a `var s []int; s = append(s, 1)` nie?

## Moduł 5 (funkcje)
- A: Co się stanie przy `x := f()`, gdy `func f() (int, error)`? (odpowiedź: nie skompiluje się — `assignment mismatch: 1 variable but f returns 2 values`)
- B: Po co funkcji zwracać dwie wartości zamiast jednej?
- C: Kiedy wykona się `defer plik.Close()` — od razu czy na końcu funkcji?

## Moduł 6 (błędy)
- A: Co wypisze `errors.Is(fmt.Errorf("kontekst: %w", os.ErrNotExist), os.ErrNotExist)`?
- B: Czemu `error` jest zwykłą wartością, a nie wyjątkiem przerywającym program?
- C: Co jest nie tak? `dane, _ := os.ReadFile("plik.txt"); fmt.Println(len(dane))`

## Moduł 7 (struktury, metody, wskaźniki)
- A: Czy ta metoda zmieni strukturę? `func (l Licznik) Zwieksz() { l.n++ }`
- B: Kiedy odbiornik wskaźnikowy, a kiedy wartościowy?
- C: Co jest wartością zerową struktury `type P struct { Imie string; Wiek int }`?

## Moduł 8 (interfejsy)
- A: Co wypisze `fmt.Println(p)`, jeśli `p` ma metodę `String() string`?
- B: Czym interfejs w Go różni się od tego, że typ "deklaruje, że go implementuje"?
- C: Czemu `var w io.Writer = os.Stdout` działa, choć `os.File` nigdzie nie mówi o `io.Writer`?

## Moduł 9-10 (pliki, JSON, testy)
- A: Co się stanie, gdy `json.Unmarshal` dostanie pole, którego nie ma w strukturze?
- B: Po co pola struktury muszą zaczynać się wielką literą, żeby trafiły do JSON-a?
- C: Czemu ten test przechodzi zawsze? `func TestX(t *testing.T) { if 1 != 1 { t.Error("nie") } }`

## Moduł 11 (współbieżność)
- A: Co wypisze program, który uruchamia 3 goroutine i od razu kończy `main`?
- B: Kiedy kanał, a kiedy mutex?
- C: Czemu ten licznik daje różne wyniki przy każdym uruchomieniu?

# Twarde zasady

- **Tylko ukończone lekcje.** Nie pytaj o materiał, którego uczeń nie miał.
- **Jedno pytanie naraz.** Czekanie na odpowiedź to część quizu.
- **Nie uruchamiaj kodu z pytań**, żeby sprawdzić własną odpowiedź. Jeśli nie jesteś pewien wyniku — nie dawaj tego pytania.
- **Sokratejskie naprowadzanie**, nie podpowiedzi typu "to chyba wskaźnik".
- **Bez ocen liczbowych** ("5/7", "70%"). Mów jakościowo.
- **Aktualizuj `student.json`** przez skill **postep** po każdym quizie.
- **Quiz to nie lekcja.** Duża luka → zaproponuj powrót do lekcji, ale nie tłumacz materiału w trakcie quizu.

# Gdy uczeń wszystko wie

- Pochwal konkretnie
- Zaproponuj **gwiazdkowe** ćwiczenie z bieżącej lekcji (skill: **cwiczenie**)
- Albo przyspiesz przejście do następnej lekcji

# Gdy uczeń się "rozsypuje"

Jeśli >50% pytań idzie źle:
- Przerwij po 3-4 pytaniach (nie męcz)
- Powiedz wprost: "Widzę, że temat X wymaga powtórki — wróćmy do lekcji N"
- Dopisz tematy do `do_powtorki`
- NIE rób z tego porażki — to diagnoza, którą zrobiliście razem
