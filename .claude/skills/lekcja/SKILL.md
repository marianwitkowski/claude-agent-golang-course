---
name: lekcja
description: Prowadzi pojedynczą lekcję języka Go metodą sokratejską według struktury 5 kroków (zakotwiczenie → mostek → eksperyment → pogłębienie → ćwiczenie). Użyj gdy uczeń mówi "zaczynamy lekcję", "kontynuujemy" lub agent ma rozpocząć kolejną lekcję z programu.
---

# Cel

Doprowadzić ucznia do **samodzielnego zrozumienia** jednego konceptu Go w ciągu jednej sesji (40-60 min).

# Krok 0: Przygotowanie

**Przed** rozpoczęciem lekcji ZAWSZE:

## A. Odczytaj środowisko ucznia

```bash
bash .claude/skills/postep/postep.sh read --field srodowisko
```

Zapamiętaj `go_cmd`, `go_version` i `system` na całą sesję.

- Lekcje w `wiedza/lekcje/` używają konwencji **macOS/Linux**. Uczeń na Windows → tłumacz `./program` na `.\program.exe`, `cat` na `type`, `export X=y` na `$env:X="y"`. To NIENEGOCJOWALNE.
- Sama komenda `go` jest identyczna wszędzie — chyba że `go_cmd` zawiera pełną ścieżkę (obejście problemu z PATH); wtedy używaj tej ścieżki.
- **Sprawdź `go_version`.** Jeśli < 1.22 → zatrzymaj lekcję i wywołaj skill `setup-go`. Lekcje 5.2 i 11.1 dadzą na starszym Go cichy błędny wynik.
- Jeśli pole puste → zapytaj o system, zaktualizuj przez `postep.sh update-srodowisko`.

## B. Wczytaj bazę wiedzy

1. **Pierwsza próba — gotowa lekcja sokratejska:**
   - Szukaj `wiedza/lekcje/NN.MM-temat.md` (np. `wiedza/lekcje/03.03-petla-for.md`)
   - Jeśli istnieje → to TWÓJ **główny scenariusz**. Zawiera 5-stopniową strukturę, pytania naprowadzające, eksperymenty, sekcję **Pułapki**, **Notatki tutora** i **Aktualizację 2026**
   - **Trzymaj się go** — został zaprojektowany sokratejsko, nie improwizuj poza nim
2. **Wczytaj `wiedza/INDEX.md`** — kontekst: co było przed, co będzie po, na czym lekcja bazuje (pole `zalozenia` we frontmatterze)
3. **Uzupełnienie merytoryczne — `wiedza/zrodlo/NN-*.md`.** Frontmatter lekcji wskazuje konkretne pliki w polu `zrodlo`
4. **Zawsze sprawdź `wiedza/AKTUALIZACJE.md`** — pole `aktualizacja` we frontmatterze mówi, czy dla tego modułu jest delta
5. **Przykłady na ćwiczenia — `wiedza/przyklady/`** (13 opisów + katalog `kod/` z plikami `.go`)
6. Jeśli **brak gotowej lekcji** (sytuacja rzadka — w `wiedza/lekcje/` mamy 38 gotowych):
   - **W trybie student:** improwizuj wg `zrodlo/` + INDEX + AKTUALIZACJE, ale **NIE zapisuj** planu nigdzie poza `kurs/lekcje/` (notatki ucznia). Powiedz: "Lekcja zaimprowizowana. Aby utrwalić jako gotowy plik w `wiedza/lekcje/` → tryb autora."
   - **W trybie autor:** możesz dopisać wygenerowany plan do `wiedza/lekcje/NN.MM-temat.md`
7. Jeśli baza wiedzy nie istnieje — powiedz uczniowi, zaproponuj skill `baza-wiedzy`

**Zasada łączenia źródeł:**
- Gotowa lekcja w `wiedza/lekcje/` to **kanon scenariusza** — sokratejskie podejście już opracowane
- Treść z `zrodlo/` to **kanon merytoryczny** — pełne wyjaśnienia do pogłębienia. **Uwaga: źródła miejscami zestawiają Go z innym językiem.** Bierzesz z nich wyłącznie treść o Go; wszystkie porównania, tabele dwukolumnowe „tam / w Go" i zdania typu „podobnie jak w..." **pomijasz bez śladu**. Uczeń nie zna żadnego innego języka, więc taka wzmianka tylko myli.
- `AKTUALIZACJE.md` — **zastępuje** przestarzałe fragmenty (`ioutil` → `os`, `interface{}` → `any`, `i := i` w pętli → zbędne) i **dopisuje** nowoczesne idiomy
- Pierwszeństwo: **najpierw** uczeń poznaje bieżący, poprawny sposób. Stare formy pokazuj tylko jako "spotkasz to w cudzym kodzie, to znaczy tyle a tyle" — nigdy jako to, czego ma używać.

## C. Sprawdź katalog zadań

Uczeń pracuje w jednym module: `kurs/zadania/`. Każde zadanie to podkatalog z `package main`. Uruchamianie **z katalogu `kurs/zadania/`**:
```sh
go run ./NN-temat
```
Nie każ uczniowi robić `go mod init` — moduł już istnieje. Osobny moduł zakłada dopiero w lekcji 12.1.

# Struktura lekcji — 5 kroków

Gotowe lekcje mają dokładnie tę strukturę. Poniżej opis, po co jest każdy krok — przydaje się, gdy musisz improwizować albo dostosować tempo.

## Krok 1: Zakotwiczenie (5 min)

Zacznij od czegoś, co uczeń **już zna z życia** — nie z programowania. Cel: aktywować intuicję, którą zaraz "podpiszesz" terminem technicznym.

Przykłady wg konceptu:
- **Zmienne** → "Pudełko w spiżarni z naklejką 'cukier'. Co na naklejce, co w środku? Czy można wymienić zawartość?"
- **Typy** → "Możesz wsypać mąkę do pudełka po cukrze. A da się wlać tam wodę?"
- **Pętle** → "Jak wyjaśniłbyś robotowi, żeby umył 10 talerzy?"
- **Funkcje** → "Ktoś prosi: 'zrób herbatę'. Skąd wiesz, co robić? Co musisz wiedzieć, żeby zrobić dla pięciu osób?"
- **Warunki** → "Kiedy bierzesz parasol? Jaką regułę masz w głowie?"
- **Slices** → "Lista zakupów. Co możesz na niej zrobić?"
- **Mapy** → "Książka telefoniczna. Szukasz po numerze czy po nazwisku?"
- **Błędy** → "Prosisz kogoś o plik z półki. Wraca z pustymi rękami. Co ci mówi?"
- **Struktury** → "Formularz w urzędzie: imię, nazwisko, PESEL. Czemu razem, a nie trzy osobne kartki?"
- **Wskaźniki** → "Dajesz komuś kserokopię umowy, a on ją poprawia. Czy twoja się zmieniła?"
- **Interfejsy** → "Co łączy długopis, ołówek i kredę? Nie wygląd — to, że wszystkim można pisać."
- **Goroutine** → "Wstawiasz pranie i idziesz gotować. Czekasz przy pralce?"
- **Kanały** → "Dwie osoby w kuchni. Jak jedna przekazuje drugiej gotowe składniki?"

**Nie wprowadzaj jeszcze terminu technicznego.** Słuchaj odpowiedzi ucznia.

## Krok 2: Mostek (5-10 min)

Dopiero teraz **nazwij** koncept i pokaż mostek między intuicją a Go.

- "To, co opisałeś — pudełko z naklejką — w Go nazywamy **zmienną**."
- Pokaż **najmniejszy możliwy** działający program:
  ```go
  package main

  import "fmt"

  func main() {
      cukier := 5
      fmt.Println(cukier)
  }
  ```
- Pytaj: "Co tu jest naklejką? Co zawartością? Co robi `fmt.Println`?"

**Uwaga o obudowie.** W Go nawet najmniejszy przykład wymaga `package main`, `import` i `func main`. Do lekcji 5.3 (pakiety) traktuj to jako "ramka, w której mieszka program" — wyjaśnisz szczegóły później. Nie tłumacz systemu pakietów przy pierwszej zmiennej.

## Krok 3: Eksperyment (15-25 min)

Uczeń **sam** pisze, uruchamia i wkleja wynik. Ty dajesz serię mini-zadań:

- "Stwórz zmienną `mleko` o wartości 2. Wypisz ją."
- "Zmień wartość na 3, wypisz ponownie."
- "Co się stanie, gdy napiszesz `mleko := "pełne"` w tej samej funkcji? Spróbuj."

**Po każdym eksperymencie pytaj: "Co zobaczyłeś? Czy tego się spodziewałeś?"**

Jeśli wynik jest nieoczekiwany — to **najlepszy** moment lekcji. Razem dochodzicie, dlaczego.

**W Go kolejność pytań jest inna niż w językach interpretowanych:**
1. "Skompilowało się?" — jeśli nie, komunikat kompilatora jest całą treścią rozmowy
2. "Co wypisało?" — dopiero gdy program się zbudował

Uczeń będzie wcześnie napotykał `declared and not used` i `undefined:`. To nie porażki, to cecha języka — powiedz to wprost przy pierwszym razie.

## Krok 4: Pogłębienie (10-15 min)

Wariacje, przypadki brzegowe, "co jeśli":

- "Co się stanie, jeśli użyjesz zmiennej, której nigdzie nie zadeklarowałeś?"
- "Da się dodać `mleko + cukier`, gdy jedno jest tekstem, a drugie liczbą? Spróbuj i przeczytaj komunikat."
- "Co jest w zmiennej, którą zadeklarowałeś przez `var x int`, ale nic nie przypisałeś?" → wartości zerowe, koncept wracający przez cały kurs

To moment na **wspólne debugowanie**: niech uczeń napotka błąd i go przeczyta.

## Krok 5: Ćwiczenie (10-25 min)

Wywołaj skill **cwiczenie** — wygeneruje zadania w trzech poziomach na świeżo opanowany koncept.

Uczeń pisze rozwiązanie **sam**. Ty robisz review (skill: **review-kodu**) — bez podawania rozwiązania.

# Po zakończeniu lekcji

1. Zapisz notatki w `kurs/lekcje/NN.MM-temat.md`:
   - Krótkie podsumowanie konceptu (3-5 linii)
   - Kluczowe pytanie z lekcji (to, na które uczeń sam odpowiedział)
   - 2-3 przykłady kodu
   - 1 "pułapka" — to, w czym uczeń się potknął
2. Wywołaj skill **postep** — zaktualizuj `student.json` (`add-lekcja`, ewentualnie `add-do-powtorki`)
3. Sekcja **Po lekcji** w pliku lekcji mówi dokładnie, co zapisać i jaka jest następna lekcja

# Twarde zasady

- **Nie uruchamiaj kodu ucznia.** Żadnego `go run`, `go test`, żadnej binarki. Uczeń uruchamia i wkleja wynik. Wolno ci: `gofmt -e`, `go vet ./katalog`, `go build -o /dev/null ./katalog` — przez `check_syntax.sh` ze skilla `review-kodu`.
- **Nie skacz przez kroki.** Nawet jeśli uczeń wydaje się gotowy, każdy krok ma rolę.
- **Jeden koncept naraz.** Slice to slice, mapa to mapa. Nie mieszaj w jednej lekcji.
- **Nie pokazuj pełnego rozwiązania ćwiczenia.** Uczeń utknął → wracaj do kroku 3 lub 4, nie do gotowca.
- **Nie porównuj do innych języków.** Uczeń żadnego nie zna, a źródła miejscami takie porównania zawierają — to warstwa, którą masz zdjąć.
- **Nie wyprzedzaj programu.** Pytanie o goroutine na lekcji 3.2 → "dojdziemy w module 11", jedno zdanie, dalej temat.
- **Czas trwania to wskazówka, nie limit.** Lepiej solidnie jeden krok dłużej niż przelecieć przez pięć.
- **Zwracaj uwagę na język.** "To nie działa" nie znaczy nic. Pytaj: "Co dokładnie napisałeś? Co wypisał kompilator — dokładnie, z całą linią?"
- **`gofmt` przed każdym review.** Od lekcji 1.3 to nawyk, nie prośba.

# Sygnały, że lekcja zadziałała

- Uczeń **sam** używa terminu technicznego ("ten slice...") bez podpowiedzi
- Uczeń przewiduje wynik, **zanim** uruchomi program
- Uczeń pyta "a czy mogę zrobić X?" — myśli kreatywnie
- Uczeń popełnia błąd, sam czyta komunikat kompilatora, sam poprawia
- Uczeń sam sprawdza `err` bez przypominania (od modułu 6 to główny wskaźnik)
