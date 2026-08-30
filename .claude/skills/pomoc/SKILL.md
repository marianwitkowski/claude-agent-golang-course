---
name: pomoc
description: Wyświetla uczniowi pogrupowaną listę dostępnych komend kursu Go — frazy, którymi może sterować agentem. Pokazuje kategorie start, lekcja, ćwiczenia, quizy, postęp, reset, środowisko, baza wiedzy. Użyj gdy uczeń mówi "lista komend", "pomoc", "help", "co mogę zrobić?", "jakie są komendy?", "nie pamiętam co mam wpisać".
---

# Cel

Pokazać uczniowi **szybką ściągawkę** dostępnych komend bez zmuszania go do otwierania README.

# Co wyświetlić

Wypisz w czacie poniższą listę. **Nie modyfikuj** kategorii ani ikon — uczeń może ją znać z pamięci.

```
📋 LISTA KOMEND KURSU GO

🚀 Start i kontynuacja
  • ucz mnie Go               → start kursu lub powitanie
  • kontynuujemy              → następna lekcja
  • pokaż program kursu       → zawartość kurs/program.md
  • zmień program kursu       → edycja celu/tempa

📚 W trakcie lekcji
  • nie rozumiem [konceptu]   → wraca do podstaw nowym kątem
  • daj mi przykład           → minimalny przykład kodu
  • co to znaczy [termin]?    → wyjaśnienie sokratejskie
  • powtórzmy tę lekcję       → od początku
  • powtórzmy [temat]         → krótka powtórka jednego konceptu

✏️  Ćwiczenia i review
  • daj mi zadanie            → ćwiczenie z bieżącej lekcji
  • daj mi więcej zadań       → dodatkowe ćwiczenia
  • sprawdź moje zadanie      → review kodu
  • skończyłem [rozgrzewkę/główne/gwiazdkę] → review konkretnego rozwiązania
  • nie działa mi             → pomoc w debugowaniu
  • nie chce się skompilować  → czytamy komunikat kompilatora
  • pokaż gwiazdkę            → odsłania zadanie ⚡

🎯 Quizy i powtórki
  • quiz                      → szybki (3 pytania)
  • quiz pełny                → pełny (5-7 pytań)
  • quiz słabe                → z tematów do powtórki

📊 Postęp
  • pokaż postępy             → podsumowanie student.json
  • gdzie skończyliśmy?       → bieżąca lekcja + ostatnia sesja
  • co mam do powtórki?       → lista do_powtorki
  • co umiem najlepiej?       → lista mocne_strony

🔄 Reset i backup
  • zresetuj kurs             → reset miękki (z backupem)
  • pełny reset kursu         → reset pełny (z backupem)
  • cofnij reset              → przywrócenie z archiwum
  • pokaż backupy             → lista postep/archiwum/

🛠️  Środowisko i pomoc
  • sprawdź Go                → weryfikacja `go version` (min. 1.22)
  • jak uruchomić kod?        → odsyła do JAK-PISAC-KOD.md
  • sformatuj mój kod         → przypomnienie o gofmt
  • lista komend / pomoc      → ten widok

📖 Baza wiedzy
  • odśwież bazę wiedzy       → pobierz najnowsze pliki z repozytoriów
  • pokaż stan bazy           → statystyki bazy
  • sprawdź czy baza aktualna → porównanie ze zdalnym repo
  • przywróć poprzednią bazę  → rollback z backupu

⚙️  Tryb pracy (dla autora kursu)
  • tryb autora               → włącz tryb modyfikacji curriculum
  • tryb student              → wróć do trybu nauki (domyślny)

💡 Nie musisz pamiętać dokładnych fraz — agent zrozumie też "wyczyść kurs",
   "co robiłam ostatnio", "zrób mi test" itp.

⚠️  Czego agent NIE zrobi: nie uruchomi twojego programu. `go run` należy do
   ciebie — wynik wklejasz do czatu i wtedy rozmawiamy.
```

# Wariant skrócony

Jeśli uczeń poprosi o "krótką pomoc" / "tylko najważniejsze":

```
🎯 NAJWAŻNIEJSZE KOMENDY

  • kontynuujemy              → następna lekcja
  • daj mi zadanie            → nowe ćwiczenie
  • sprawdź moje zadanie      → review kodu
  • nie działa mi             → debugowanie
  • quiz                      → szybka powtórka
  • pokaż postępy             → twój stan
  • lista komend              → pełna lista
```

# Ściąga komend `go` — na życzenie

Jeśli uczeń pyta "jakie są komendy Go?" (a nie komendy kursu), pokaż to:

```
🔧 KOMENDY GO, KTÓRE POZNASZ W TYM KURSIE

  go run ./NN-temat      uruchom program            (lekcja 1.2)
  gofmt -w plik.go       sformatuj kod              (lekcja 1.3)
  go build -o nazwa .    zbuduj samodzielną binarkę (lekcja 12.3)
  go test ./...          uruchom testy              (lekcja 10.1)
  go vet ./...           poszukaj podejrzanych miejsc (lekcja 10.3)
  go mod init nazwa      załóż nowy moduł           (lekcja 10.2)
  go doc fmt.Println     dokumentacja bez internetu (lekcja 10.2)
  go run -race ./...     wykryj wyścigi danych      (lekcja 11.3)

Zadania uruchamiasz z katalogu kurs/zadania/. Uwaga: `./NN-temat` to katalog,
nie plik — w Go uruchamia się pakiet.
```

# Twarde zasady

- **Wypisuj listę w jednym bloku** — nie dziel na wiele wiadomości.
- **Nie wymyślaj nowych komend** poza listą. Uczeń pyta o coś, czego nie ma → "tego nie ma, ale możesz [...]".
- **Po pokazaniu listy** zadaj jedno pytanie: "Z czego dziś korzystamy?" — żeby nie zostać w trybie "wyświetlam pomoc i czekam".
- **Nie pokazuj listy** w środku trwającej lekcji bez wyraźnej prośby — wybija z rytmu.
- **Nie pokazuj komend Go z lekcji, których uczeń jeszcze nie miał**, jako rzeczy do użycia teraz. Ściąga wyżej podaje numery lekcji właśnie po to.
