# Kurs języka Go z tutorem Claude

Interaktywny kurs podstaw **Go** dla **kompletnie początkujących**, prowadzony przez agenta Claude Code metodą **sokratejską** — uczeń sam dochodzi do rozwiązań przez pytania naprowadzające.

Od `fmt.Println("Cześć")` do własnego narzędzia wiersza poleceń: **38 lekcji w 12 modułach**.

## Dla kogo

- Osoby, które **nigdy nie programowały** — kurs nie zakłada znajomości żadnego innego języka
- Chcące uczyć się w swoim tempie, z prowadzącym, który nie podaje gotowców
- Mające zainstalowane lub gotowe zainstalować Claude Code

## Jak zacząć

> 🚀 **Szybki start:** zobacz **[QUICKSTART.md](QUICKSTART.md)** — przewodnik krok po kroku z przykładami.

W katalogu kursu uruchom Claude Code i napisz:

```
ucz mnie Go
```

albo

```
zacznij lekcję
```

Agent `go-tutor` przeprowadzi Cię przez:
1. Sprawdzenie środowiska (Go 1.22+, edytor)
2. Krótki wywiad (cel, dostępny czas)
3. Wygenerowanie programu kursu dopasowanego do Ciebie
4. Pierwszą lekcję

## Struktura projektu

```
.
├── .claude/
│   ├── agents/go-tutor.md              # główny agent
│   └── skills/                         # specjalistyczne umiejętności
│       ├── setup-go/                   # sprawdzenie środowiska, instalacja Go
│       ├── program-kursu/              # generator programu
│       ├── lekcja/                     # prowadzenie lekcji
│       ├── cwiczenie/                  # generator ćwiczeń (3 poziomy)
│       ├── review-kodu/                # sokratejski review + check_syntax.sh/.ps1
│       ├── quiz/                       # quizy powtórkowe między lekcjami
│       ├── reset-kursu/                # reset miękki/pełny z backupem
│       ├── pomoc/                      # lista komend w czacie
│       ├── baza-wiedzy/                # odświeżanie z repozytoriów źródłowych
│       └── postep/                     # śledzenie postępu (narzędzie w Go + wrappery)
├── wiedza/                             # lokalna baza wiedzy
│   ├── zrodlo/                         # mirror golang-for-python-developers (32 pliki)
│   ├── przyklady/                      # mirror golang-examples (13) + kod/ z golang20230314 (24 .go)
│   ├── AKTUALIZACJE.md                 # delta: Go 1.22 → 1.27
│   ├── INDEX.md                        # mapowanie źródeł → 38 lekcji sokratejskich
│   └── lekcje/                         # 38 gotowych lekcji + SZABLON-LEKCJI.md
├── kurs/
│   ├── JAK-PISAC-KOD.md                # ⬅ przeczytaj na początku: workflow ćwiczeń
│   ├── program.md                      # Twój program kursu (powstanie po onboardingu)
│   ├── lekcje/                         # notatki z każdej lekcji
│   ├── zadania/                        # Twój kod — jeden moduł Go, katalog na ćwiczenie
│   │   └── go.mod
│   └── projekt/                        # Twój program z modułu 12 (osobny moduł)
└── postep/
    ├── student.json                    # Twój stan: lekcje, mocne strony, do powtórki
    ├── backups/                        # automatyczne kopie przed każdym zapisem
    └── archiwum/                       # backupy po resetach (nigdy nie kasowane automatycznie)
```

## Program kursu — 12 modułów, 38 lekcji

| Moduł | Temat | Lekcje |
| --- | --- | --- |
| 1 | Wprowadzenie i środowisko | 3 |
| 2 | Podstawy języka — zmienne, typy, wejście/wyjście | 4 |
| 3 | Decyzje i powtórzenia — `if`, `switch`, `for` | 3 |
| 4 | Kolekcje — tablice, slices, mapy, `range` | 4 |
| 5 | Funkcje — wiele wartości, domknięcia, `defer`, pakiety | 3 |
| 6 | Błędy jako wartości — `error`, opakowywanie, `panic` | 3 |
| 7 | Struktury i metody — `struct`, wskaźniki, embedding | 3 |
| 8 | Interfejsy — kontrakt niejawny, `Stringer`, `io.Writer` | 2 |
| 9 | Pliki i dane — pliki, JSON, argumenty CLI | 3 |
| 10 | Testy i narzędzia — `go test`, moduły, `go vet` | 3 |
| 11 | Współbieżność — goroutine, kanały, `select`, `-race` | 3 |
| 12 | Projekt i dalsze kroki — własne narzędzie CLI | 4 |

Źródłem prawdy dla struktury jest [`wiedza/INDEX.md`](wiedza/INDEX.md).

**Czego w kursie nie ma:** generyków, aplikacji webowych, baz danych, mikroserwisów, `context.Context`. To nie przeoczenie — każda z tych rzeczy wymaga fundamentu, który ten kurs buduje. Mapa dalszych kroków czeka w lekcji 12.4.

## Zanim zaczniesz pierwszą lekcję

Przeczytaj **[`kurs/JAK-PISAC-KOD.md`](kurs/JAK-PISAC-KOD.md)** — 5 minut, ale wyjaśnia:
- gdzie zapisywać kod (`kurs/zadania/NN-temat/main.go`)
- jak uruchamiać programy (`go run ./NN-temat` — z katalogu `kurs/zadania`)
- jak czytać komunikaty kompilatora i paniki
- czym jest `gofmt` i czemu nie ma o czym dyskutować
- cały workflow ćwiczenia od początku do końca
- najczęstsze pułapki początkujących

## Filozofia

- **Sokratejsko, nie wykładowo.** Agent zadaje pytania zamiast wyjaśniać. To wolniejsze, ale głębsze.
- **Nie uruchamiamy kodu za Ciebie.** Agent nigdy nie robi `go run` ani `go test` na Twoim kodzie — może najwyżej sprawdzić składnię (`gofmt -e`) i to, czy się kompiluje (`go build -o /dev/null`). Sam piszesz, sam uruchamiasz, sam czytasz wynik. To część nauki, nie ograniczenie narzędzia.
- **Bez porównań do innych języków.** Kurs zakłada, że to Twój pierwszy język. Nic nie jest wyjaśniane przez „to jak w...".
- **Postęp jest Twój.** Wszystko w `postep/student.json` — możesz przeglądać, edytować, eksportować.
- **Twoje tempo.** Czas trwania lekcji to wskazówka, nie termin. Możesz wrócić za tydzień, agent będzie wiedział, gdzie skończyliście.

## Lista komend

Komendy wpisujesz w Claude Code — to **frazy w języku naturalnym**, nie formalne komendy. Agent rozpozna intencję, nawet jeśli sformułujesz to inaczej. Poniżej wersje kanoniczne.

### 🚀 Start i kontynuacja

| Komenda | Co zrobi agent |
| --- | --- |
| `ucz mnie Go` | Start kursu — onboarding albo powitanie i kontynuacja |
| `zacznij lekcję` | To samo, bardziej formalnie |
| `kontynuujemy` | Kolejna lekcja z programu (`kurs/program.md`) |
| `pokaż program kursu` | Wyświetla zawartość `kurs/program.md` |
| `zmień program kursu` | Edycja programu (np. zmiana celu, tempa) |

### 📚 W trakcie lekcji

| Komenda | Co zrobi agent |
| --- | --- |
| `nie rozumiem [konceptu]` | Wraca do podstaw konceptu nowym kątem |
| `daj mi przykład` | Pokazuje minimalny przykład kodu (nie rozwiązanie) |
| `co to znaczy [termin]?` | Wyjaśnia termin przez pytania naprowadzające |
| `powtórzmy tę lekcję` | Wraca do bieżącej lekcji od początku |

### ✏️ Ćwiczenia i review

| Komenda | Co zrobi agent |
| --- | --- |
| `daj mi zadanie` | Generuje ćwiczenie z bieżącej lekcji (3 poziomy) |
| `daj mi więcej zadań` | Dodatkowe ćwiczenia na opanowany koncept |
| `sprawdź moje zadanie` | Sokratejski review kodu z `kurs/zadania/` |
| `skończyłem [rozgrzewkę/główne/gwiazdkę]` | Review konkretnego rozwiązania |
| `nie chce się skompilować` | Wspólne czytanie komunikatu kompilatora |
| `nie działa mi` | Pomoc w debugowaniu — agent pyta o wynik i oczekiwania |
| `pokaż gwiazdkę` | Odsłania zadanie ⚡ (po ukończeniu pozostałych) |

### 🎯 Quizy i powtórki

| Komenda | Co zrobi agent |
| --- | --- |
| `quiz` | Szybki quiz (3 pytania) z ostatnich lekcji |
| `quiz pełny` | Pełny quiz (5-7 pytań) z całości materiału |
| `quiz słabe` | Quiz z tematów oznaczonych w `do_powtorki` |
| `powtórzmy [temat]` | Krótka powtórka konkretnego konceptu |

### 📊 Postęp

| Komenda | Co zrobi agent |
| --- | --- |
| `pokaż postępy` | Podsumowanie ze `student.json` |
| `gdzie skończyliśmy?` | Przypomnienie aktualnej lekcji i ostatniej sesji |
| `co mam do powtórki?` | Lista tematów z pola `do_powtorki` |
| `co umiem najlepiej?` | Lista z pola `mocne_strony` |

### 🔄 Reset i backup

| Komenda | Co zrobi agent |
| --- | --- |
| `zresetuj kurs` | Reset **miękki** — czyści postęp i program (z backupem) |
| `pełny reset kursu` | Reset **pełny** — czyści wszystko, też Twój kod (z backupem) |
| `cofnij reset` | Przywrócenie ostatniego stanu z `postep/archiwum/` |
| `pokaż backupy` | Wypisuje katalogi w `postep/archiwum/` |

### 🛠️ Środowisko i pomoc

| Komenda | Co zrobi agent |
| --- | --- |
| `sprawdź Go` | Weryfikacja `go version` (skill `setup-go`) |
| `jak uruchomić kod?` | Odsyła do `kurs/JAK-PISAC-KOD.md`, sekcja 3-4 |
| `sformatuj mój kod` | Przypomnienie o `gofmt -w` |
| `lista komend` | **Wyświetla tę listę w czacie** (skill `pomoc`) |
| `pomoc` / `help` / `co mogę zrobić?` | To samo co `lista komend` |
| `krótka pomoc` | Skrócona wersja — tylko najważniejsze |

### 📖 Baza wiedzy

| Komenda | Co zrobi agent |
| --- | --- |
| `odśwież bazę wiedzy` | Pobiera najnowsze pliki z repozytoriów (z backupem) |
| `pokaż stan bazy` | Statystyki: ile plików, data pobrania, SHA commita |
| `sprawdź czy baza aktualna` | Porównuje lokalne pliki ze zdalnym repo |
| `przywróć poprzednią bazę` | Rollback z `wiedza/zrodlo.backup-*` |

### ⚙️ Tryb pracy (dla autora kursu)

| Komenda | Co zrobi agent |
| --- | --- |
| `tryb autora` | Włącza tryb modyfikacji lekcji i skilli (wymaga potwierdzenia pełną frazą) |
| `tryb student` | Powrót do trybu nauki (domyślny) |

> 💡 **Wskazówka:** Nie musisz pamiętać dokładnych fraz. „Zrób mi quiz", „wyczyść wszystko", „co robiłam ostatnio" — agent dopyta o szczegóły.

## Wymagania

- **System operacyjny:** macOS, Linux lub Windows — skill `setup-go` ma dedykowane gałęzie dla każdego. Na Windows agent używa Git Basha, jeśli jest; w przeciwnym razie sięga po wersje `.ps1` swoich narzędzi (`-ExecutionPolicy Bypass`, bez zmiany ustawień systemu)
- **Go 1.22 lub nowsze** — to twarda granica. Od 1.22 zmienna pętli `for` powstaje na nowo w każdej iteracji; na starszych wersjach lekcje 5.2 i 11.1 dadzą **cichy, błędny wynik**. Sprawdzenie i instalacja przez skill `setup-go`.
- **Claude Code** (https://claude.com/code)
- **Edytor tekstu** — rekomendacja: VS Code + rozszerzenie Go (działa identycznie na każdym systemie)

> 💡 **Notatka o komendach:** wszystkie komendy `go ...` są identyczne na każdym systemie. Różnią się tylko komendy powłoki (`./program` kontra `.\program.exe`, `cat` kontra `type`). Agent tłumaczy je automatycznie; pełna mapa różnic w `kurs/JAK-PISAC-KOD.md` (sekcja 3) i w skillu `setup-go`.

## Materiały źródłowe

Baza wiedzy pochodzi z trzech repozytoriów:

- [`marianwitkowski/golang-for-python-developers`](https://github.com/marianwitkowski/golang-for-python-developers) — kanon merytoryczny (32 pliki)
- [`marianwitkowski/golang-examples`](https://github.com/marianwitkowski/golang-examples) — materiał na ćwiczenia i projekt (13 plików)
- [`marianwitkowski/golang20230314`](https://github.com/marianwitkowski/golang20230314) — minimalne przykłady kodu (24 pliki `.go`)

Materiały zakładają czytelnika, który zna już inny język programowania; ten kurs bierze z nich wyłącznie treść o Go, a wszystkie porównania pomija — bo uczeń nie ma z czym porównywać. Aneks [`wiedza/AKTUALIZACJE.md`](wiedza/AKTUALIZACJE.md) prostuje to, co zdezaktualizowało się między Go 1.22 a 1.27.
