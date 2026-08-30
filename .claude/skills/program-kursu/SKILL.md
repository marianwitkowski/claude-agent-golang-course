---
name: program-kursu
description: Generuje plik kurs/program.md — spersonalizowany program 12 modułów / 38 lekcji podstaw języka Go, na podstawie bazy wiedzy w wiedza/INDEX.md. Dostosowuje akcenty do celu ucznia (praca/hobby/narzędzia/szkoła) i deklarowanego tempa. Użyj raz, podczas onboardingu, po krótkim wywiadzie z uczniem.
---

# Cel

Stworzyć `kurs/program.md` — plan kursu, do którego uczeń i tutor będą wracać. To **kompas**, nie sztywne tory.

# Źródło prawdy

**Zawsze** opieraj plan na pliku `wiedza/INDEX.md`. Nie wymyślaj modułów ani lekcji — tabela z INDEX.md to kanon: **12 modułów, 38 lekcji**, zbudowanych na repozytoriach `marianwitkowski/golang-for-python-developers`, `golang-examples` i `golang20230314`.

Jeśli `wiedza/INDEX.md` nie istnieje → coś jest nie tak z bazą wiedzy. Powiedz uczniowi i zaproponuj skill `baza-wiedzy`.

# Wejście

Wymagane od ucznia (przed wywołaniem skill):
- **Cel:** praca / narzędzia i automatyzacja / hobby / szkoła / inne
- **Czas tygodniowo:** <2h / 2-5h / 5-10h / 10+h
- **Doświadczenie z programowania:** brak (domyślnie) / coś dotykał / inny język

# Procedura

1. **Wczytaj** `wiedza/INDEX.md` — źródło struktury kursu
2. **Skopiuj kanon** (12 modułów, 38 lekcji — moduły mają po 2-4 lekcje)
3. **Personalizuj** akcenty wg celu (patrz niżej)
4. **Dostosuj tempo** wg dostępnego czasu
5. **Zapisz** do `kurs/program.md`

# Personalizacja wg celu

Personalizacja dotyczy **akcentów i projektu końcowego**, nie struktury. Wszystkie 38 lekcji zostaje w tej samej kolejności.

- **Cel: praca (backend, DevOps)** → mocniej moduł 10 (testy, moduły, `go vet`) i 11 (współbieżność — to jest to, po co firmy sięgają po Go); projekt: narzędzie CLI przetwarzające dane albo raportujące stan czegoś
- **Cel: narzędzia i automatyzacja** → mocniej moduł 9 (pliki, JSON, argumenty) i 12.3 (`go build`, kompilacja skrośna — jedna binarka bez zależności to główny powód, dla którego Go nadaje się do skryptów); projekt: narzędzie zastępujące ręczną czynność, którą uczeń faktycznie wykonuje
- **Cel: hobby / gry tekstowe** → mocniej moduł 3 (pętle, warunki) i 4 (kolekcje); projekt: gra tekstowa z tabelą wyników w pliku
- **Cel: szkoła / algorytmy** → mocniej moduł 5 (funkcje) i 4 (slices, mapy); projekt: solver albo kalkulator z testami tablicowymi (10.1)
- **Cel: inny** → dopytaj o konkret, dobierz akcent

**Współbieżność (moduł 11) zostaje zawsze**, niezależnie od celu. To wizytówka Go — uczeń, który jej nie dotknie, nie zobaczy, czym ten język się różni.

# Tempo

Lekcja trwa 40-60 minut plus ćwiczenie.

| Czas/tydz | Lekcji/tydz | Czas trwania kursu |
| --------- | ----------- | ------------------ |
| <2h       | 1           | ~38 tygodni        |
| 2-5h      | 2-3         | ~13-19 tygodni     |
| 5-10h     | 3-5         | ~8-13 tygodni      |
| 10+h      | 5-7         | ~6-8 tygodni       |

**Uwaga o module 12:** projekt rozciąga się na kilka sesji (lekcja 12.2 jest prowadzona wielokrotnie). Do szacunku doliczaj 2-4 dodatkowe sesje.

# Format pliku `kurs/program.md`

```markdown
# Program kursu Go — [imię]

**Cel:** [praca / narzędzia / hobby / szkoła]
**Tempo:** [X lekcji / tydzień]
**Rozpoczęto:** YYYY-MM-DD
**Wersja Go:** [z srodowisko.go_version] (minimum kursu: 1.22)
**Bazujemy na:** wiedza/INDEX.md

## Jak działa kurs

Uczysz się przez pytania, nie przez wykłady. Kod piszesz i uruchamiasz sam —
tutor czyta, pyta i podpowiada, ale nigdy nie uruchamia twojego programu za ciebie.
Postęp zapisuje się w `postep/student.json`, więc każdą sesję zaczynasz tam,
gdzie skończyłeś.

## Moduły i lekcje

### Moduł 1: Wprowadzenie i środowisko
- Lekcja 1.1: Czym jest Go i do czego służy
- Lekcja 1.2: Pierwszy program — `package main`, `func main`, `go run`
- Lekcja 1.3: Edytor, terminal, `gofmt`

### Moduł 2: Podstawy języka
- Lekcja 2.1: Zmienne i typy (`var`, `:=`)
- Lekcja 2.2: Operatory i brak niejawnych konwersji
- Lekcja 2.3: Wypisywanie — `Println`, `Printf`
- Lekcja 2.4: Wejście od użytkownika

[...kontynuuj wg INDEX.md, wszystkie 12 modułów, 38 lekcji...]

## Projekt końcowy (Moduł 12)

[Spersonalizowany pod cel ucznia — 2-3 propozycje do wyboru,
 wszystkie w formie narzędzia wiersza poleceń]

## Czego w tym kursie nie ma

Generyki, aplikacje webowe, bazy danych, mikroserwisy. To nie przeoczenie —
każda z tych rzeczy wymaga fundamentu, który ten kurs buduje.
Mapa dalszych kroków czeka w lekcji 12.4.
```

# Twarde zasady

- **Źródłem prawdy jest `wiedza/INDEX.md`.** Nie wymyślaj lekcji, nie pomijaj modułów bez zgody ucznia.
- **Trzymaj się 12 modułów i kolejności.** Kolejność nie jest przypadkowa: wskaźniki wymagają struktur (7.1 → 7.2), interfejsy wymagają metod (7.2 → 8.1), współbieżność wymaga funkcji i domknięć (5.2 → 11.1).
- **Nie wymyślaj modułów typu "Gin", "Kubernetes", "gRPC"** — te tematy są wzmiankowane w lekcji 12.4 (mapa ekosystemu), nigdy jako osobne lekcje.
- **Nie skracaj kursu przez wycięcie modułu 6 (błędy)** — nawet jeśli uczeń chce szybciej. W Go obsługa błędów jest wpleciona w każdą operację na plikach i danych; bez modułu 6 lekcje 9-12 nie mają sensu.
- Jeśli uczeń chce krótszej wersji → zaproponuj zatrzymanie się po module 10 i wrócenie do 11-12 później. Nie wycinaj środka.
- Plik nadpisujesz **tylko jeśli** uczeń świadomie chce zmienić program (np. zmienił się cel).

# Po wygenerowaniu

Pokaż uczniowi **spis modułów** (nie cały plik) i zapytaj, czy chce coś zmienić, zanim ruszycie z lekcją 1.1. Dodaj jedno zdanie kotwiczące:

> "Trzydzieści osiem lekcji brzmi dużo, ale pierwsze siedem to podstawy, które w innych językach wyglądają podobnie. Po module 6 zaczyna się to, co w Go jest naprawdę własne."
