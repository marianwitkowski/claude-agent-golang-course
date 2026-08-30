---
name: cwiczenie
description: Generuje 1-3 ćwiczenia w Go do samodzielnego rozwiązania przez ucznia, dopasowane do bieżącej lekcji i poziomu trudności. Trzy poziomy: rozgrzewka, główne, gwiazdka. Użyj na końcu lekcji lub gdy uczeń prosi o "więcej zadań".
---

# Cel

Dać uczniowi **konkretne, mierzalne** zadanie do napisania w Go — samodzielnie, bez podpowiedzi w kodzie.

# Trzy poziomy

Dla każdej lekcji wygeneruj zestaw 3 ćwiczeń:

| Poziom        | Cel                                              | Czas       | Wskazówka                |
| ------------- | ------------------------------------------------ | ---------- | ------------------------ |
| Rozgrzewka 🔥 | Sprawdzenie, czy uczeń rozumie składnię          | 5-10 min   | "Powinno być łatwe"      |
| Główne ⭐     | Czy umie złożyć z poznanych klocków              | 15-20 min  | "Pomyśl, zanim napiszesz" |
| Gwiazdka ⚡   | Wyzwanie — łączy bieżącą lekcję z poprzednimi    | 20-30 min  | "Może być trudne, to OK" |

Uczeń wybiera, ile robi. Minimum: rozgrzewka + główne.

**Sekcja "Krok 5 — Ćwiczenie" w pliku lekcji zwykle podaje już trzy propozycje.** Zacznij od nich — są dopasowane do materiału. Generuj nowe tylko, gdy uczeń prosi o więcej albo tamte okazały się źle wycelowane.

# Struktura katalogu

Wszystkie ćwiczenia mieszkają w **jednym module Go**:

```
kurs/zadania/
├── go.mod                   ← jeden moduł na cały kurs, NIE twórz nowych
├── 01-hello/main.go
├── 12-slices/main.go
└── 12-slices/ZADANIA.md
```

Nazwa katalogu: `NN-temat`, gdzie `NN` to kolejny numer ćwiczenia (nie numer lekcji — jedna lekcja może dać kilka katalogów).

**Uruchamianie — z katalogu `kurs/zadania/`:**
```sh
go run ./12-slices
```

Zwróć uwagę: `./12-slices` to **katalog**, nie plik. Uczeń przyzwyczajony do "uruchamiam plik" musi to raz usłyszeć wprost: w Go uruchamia się pakiet.

**Jeden `package main` na katalog.** Trzy rozwiązania (🔥/⭐/⚡) w jednym katalogu **nie skompilują się** — trzy razy `func main`. Rozwiązania:
- osobne katalogi: `12-slices-a/`, `12-slices-b/`, `12-slices-c/` — **domyślnie tak rób**
- albo jeden `main.go` z trzema funkcjami (`rozgrzewka()`, `glowne()`, `gwiazdka()`) wywoływanymi kolejno z `main` — dobre od modułu 5 wzwyż, gdy uczeń zna już funkcje

# Format ćwiczenia

Zapisz w `kurs/zadania/NN-temat/ZADANIA.md`:

```markdown
# Lekcja N.M: [temat] — ćwiczenia

## 🔥 Rozgrzewka
**Cel:** [jednolinijkowo, co ćwiczy]
**Zadanie:** [opis w 1-2 zdaniach]
**Oczekiwany wynik:**

    $ go run ./NN-temat-a
    Cześć, Anno!
    Masz 32 lata.

## ⭐ Główne
[jw.]

## ⚡ Gwiazdka
[jw.]
```

# Zasady dobrego ćwiczenia

- **Konkretny, sprawdzalny wynik.** Nie "napisz program o kotach", tylko "wypisz 5 razy 'miau', każde w nowej linii".
- **Pokaż oczekiwane wyjście dosłownie** — jako blok tekstu z `$ go run ...` i wypisanymi liniami. Uczeń sam porówna.
- **Realistyczny kontekst.** Nie `x := 5; y := 7; policz z`. Lepiej: "w koszyku jest 5 jabłek i 7 gruszek — ile owoców razem?"
- **Wymaga myślenia, nie kopiowania.** Jeśli rozwiązanie to przykład z lekcji z podmienionymi liczbami — ćwiczenie za łatwe.
- **Tylko biblioteka standardowa.** Żadnego `go get` do modułu 10 włącznie; w projekcie (moduł 12) uczeń może sięgnąć po zewnętrzną bibliotekę, jeśli sam uzasadni, po co.
- **Nie wprowadzaj konstrukcji spoza dotychczasowych lekcji.** Zajrzyj do `wiedza/INDEX.md`, żeby sprawdzić, co uczeń już miał. Typowe wpadki: `slices.Sort` przed 4.2, wskaźnik przed 7.2, `defer` przed 5.3, gorutyna gdziekolwiek przed 11.1.
- **Od modułu 6 każde ćwiczenie dotykające plików albo konwersji musi wymuszać obsługę błędu.** Ignorowanie `err` w ćwiczeniu to nauka złego nawyku.
- **Od modułu 10 dołączaj do gwiazdki wymóg testu** — jeden `_test.go` z dwoma przypadkami.

# Przykłady (lekcja 2.1: zmienne i typy)

🔥 **Rozgrzewka:** Zadeklaruj `imie` (twoje imię), `wiek` (twój wiek) i `wzrost` (w metrach, np. 1.78). Wypisz każdą w osobnej linii.

⭐ **Główne:** Napisz program-wizytówkę. Zmienne `imie`, `nazwisko`, `wiek`, `miasto`. Wypisz jedną linią przez `fmt.Printf`:
```
Cześć, jestem Anna Kowalska, mam 32 lata i mieszkam w Krakowie.
```

⚡ **Gwiazdka:** Ojciec ma 45 lat, syn 12. Zadeklaruj obie zmienne. Wypisz różnicę wieku. Potem wypisz, ile lat będzie miał syn, gdy ojciec skończy 60. (Bez `if`, bez wczytywania — same zmienne i arytmetyka.) Zadeklaruj też zmienną, której **nie użyjesz** — zobacz, co powie kompilator, i dopiero potem ją usuń.

# Co po wygenerowaniu

- Pokaż uczniowi tylko 🔥 i ⭐ (gwiazdkę odsłoń, gdy skończy oba)
- Powiedz dokładnie:
  - **gdzie** ma zapisać kod: `kurs/zadania/NN-temat-a/main.go`
  - **jak** uruchomić — z komendą wg `srodowisko.go_cmd`:
    ```sh
    cd kurs/zadania
    go run ./NN-temat-a
    ```
  - **że przed pokazaniem kodu robi `gofmt`** (od lekcji 1.3): `gofmt -w NN-temat-a/main.go`, albo po prostu zapisuje plik w VS Code z włączonym "Format on Save"
- Jeśli uczeń wygląda na zagubionego co do workflow — przypomnij: "Zajrzyj do `kurs/JAK-PISAC-KOD.md`, sekcja 4 — cały cykl krok po kroku."
- **Uczeń sam uruchamia kod i wkleja wynik.** Ty potem robisz review (skill: review-kodu)

# Reguła komend — ZAWSZE z `student.json`

Przed wypisaniem jakiejkolwiek komendy:
```bash
bash .claude/skills/postep/postep.sh read --field srodowisko
```

> **Windows bez Git Basha:** zamień `bash <ścieżka>.sh` na `powershell -ExecutionPolicy Bypass -File <ścieżka>.ps1`. Argumenty bez zmian. Reguła wyboru — w `go-tutor.md`, sekcja „Narzędzia kursu".
Użyj `go_cmd` i konwencji z `system` (Windows: `.\program.exe`, `type` zamiast `cat`). Jeśli puste → zapytaj o system i zaktualizuj.

# Twarde zasady

- **Nigdy nie pisz rozwiązania** w `ZADANIA.md`. Tylko opis i oczekiwane wyjście.
- **Nie twórz nawet szkieletu** w `main.go` — to plik ucznia. Wyjątek: obudowa `package main` / `import` / `func main` na pierwszych dwóch lekcjach, zanim uczeń ją zapamięta; od lekcji 1.3 pisze ją sam.
- **Nie twórz nowych plików `go.mod`.** Ćwiczenia żyją w jednym module: `kurs/zadania/go.mod`. Uczeń wywołuje `go mod init` tylko dwa razy w całym kursie: w 10.2 na module próbnym poza repozytorium (`~/moj-projekt`, kasowanym po lekcji) i w 12.1 dla własnego projektu. **W `kurs/zadania/` nigdy** — zagnieżdżony moduł rozsypie uruchamianie pozostałych ćwiczeń.
- **Nie uruchamiaj rozwiązania ucznia**, żeby "sprawdzić, czy wychodzi". Poproś o wklejenie wyniku.
- Jeśli uczeń prosi "daj mi szablon" — odmów miękko: "Wypisz mi w czacie, jakie zmienne będą ci potrzebne i jakiego typu. Kod napiszesz, jak będziesz miał to na papierze."
