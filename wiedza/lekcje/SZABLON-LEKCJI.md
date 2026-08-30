---
lekcja: X.Y
tytul: Krótki, konkretny tytuł
modul: NN-nazwa-modulu
zrodlo: NN-plik.md (konkretne sekcje)
aktualizacja: [moduł NN], [ogólne]
czas_min: 30-60
zalozenia: co uczeń musi już umieć (lekcje poprzedzające)
---

<!-- KONWENCJA KOMEND:
     Komendy pisane w wersji macOS/Linux. Różnice na Windows:
     - separator ścieżek: `kurs\zadania\...` zamiast `kurs/zadania/...`
     - binarka po `go build` ma rozszerzenie `.exe`
     - sama komenda `go` jest identyczna na wszystkich systemach
     Uczeń pracuje w `kurs/zadania/` — JEDEN moduł Go, każde zadanie w podkatalogu.
     Uruchamianie: `go run ./NN-temat` z katalogu `kurs/zadania`.
     AGENT NIGDY NIE URUCHAMIA KODU UCZNIA — uruchamia uczeń i wkleja wynik. -->

# Lekcja X.Y — Tytuł

## Cel
Po lekcji uczeń:
- konkret 1
- konkret 2
- konkret 3

## Krok 1 — Zakotwiczenie (3-5 min)
**Pytanie wejściowe:**
> "Coś z życia, co aktywuje intuicję — bez terminologii technicznej."

**Co chcesz usłyszeć:** krótkie naprowadzenie, co odpowiedź ma uchwycić.
**Czego NIE rób:** nie wprowadzaj jeszcze terminu technicznego.

## Krok 2 — Mostek (3-5 min)
**Mostek konceptualny:** "To co opisałeś — [intuicja] — w Go nazywamy [termin]."

**Minimalny przykład (2-3 linijki, zawsze kompletny plik jeśli to pierwsze lekcje):**
```go
package main

import "fmt"

func main() {
	// najprostsze co pokazuje koncept
}
```

**Pytania:**
- "Co tu jest [X]? Co [Y]?"
- "Co robi linia 2?"

## Krok 3 — Eksperyment (10-15 min)
Mini-zadania (uczeń sam pisze i uruchamia — `go run ./NN-temat`):
1. zadanie 1
2. zadanie 2
3. zadanie 3

**Po każdym pytaj:** "Co zobaczyłeś? Czy się spodziewałeś?"
**Gdy kompilator odmówi:** "Wklej dokładnie, co powiedział kompilator" → czytacie komunikat razem.

## Krok 4 — Pogłębienie (10-15 min)
Wariacje i "co jeśli":
- "Co się stanie, gdy...?"
- "A jeśli zamiast X dasz Y?"

**Typowe błędy do sprowokowania:**
- `declared and not used` / `imported and not used` — Go nie wybacza nieużywanych rzeczy
- konkretny błąd typów, np. `invalid operation: mismatched types int and string`

## Krok 5 — Ćwiczenie
→ Wywołaj skill `cwiczenie`. Sugerowany kontekst:
- temat: [konkretny temat]
- elementy do użycia: [...]

## Pułapki
- Pułapka 1 — opis krótko + jak ją rozpoznać
- Pułapka 2

## Aktualizacja 2026 (z `AKTUALIZACJE.md`)
Co dodać do treści ze źródła:
- punkt 1 (z sekcji `[moduł NN]`)
- punkt 2

**Kolejność:** najpierw klasyczny sposób (jak w `zrodlo/`), potem nowoczesny idiom.

## Notatki tutora
- Jeśli uczeń mówi X → reaguj Y
- Częsta dygresja, której unikać
- Co powiedzieć, gdy uczeń wydaje się znudzony / przytłoczony

## Po lekcji
- Notatka w `kurs/lekcje/X.Y-temat.md` (3-5 zdań o tym, co zostało zrozumiane + 1 pułapka)
- `postep`: dopisz lekcję `X.Y` do `ukonczone_lekcje` z subiektywną trudnością (1-5)
- `aktualna_lekcja` ustaw na kolejną
- Jeśli uczeń się potknął → wpis do `do_powtorki`
