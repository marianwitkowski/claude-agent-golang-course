### Opis biblioteki `go-funk` na przykładzie zaawansowanego kodu

`go-funk` to wszechstronna biblioteka do pracy z kolekcjami w Go, która umożliwia wykonywanie zaawansowanych operacji na slice'ach i innych strukturach danych. Poniżej znajdziesz zaawansowany przykład kodu, który pokazuje, jak można wykorzystać różne funkcje tej biblioteki do pracy ze slice'ami. Kod został podzielony na fragmenty, a każdy fragment jest szczegółowo opisany.

### 1. Instalacja biblioteki `go-funk`

Zanim zaczniesz, musisz zainstalować bibliotekę `go-funk`:

```bash
go get -u github.com/thoas/go-funk
```

### 2. Definicja struktury i danych wejściowych

Najpierw zdefiniujemy strukturę `Person` oraz slice zawierający kilka przykładowych danych.

```go
package main

import (
	"fmt"
	"github.com/thoas/go-funk"
)

type Person struct {
	Fname string
	Lname string
	Age   int
	City  string
}

func main() {
	// Definiowanie slice osób (Person)
	people := []Person{
		{"John", "Doe", 16, "New York"},
		{"Jane", "Smith", 20, "Los Angeles"},
		{"Bob", "Johnson", 25, "New York"},
		{"Alice", "Brown", 17, "Chicago"},
		{"Tom", "Clark", 22, "New York"},
		{"Emma", "Davis", 30, "Los Angeles"},
	}
```

#### Wyjaśnienie:
- **Person**: Struktura `Person` zawiera cztery pola: `Fname` (imię), `Lname` (nazwisko), `Age` (wiek) i `City` (miasto zamieszkania).
- **people**: Slice zawierający kilka przykładowych danych typu `Person`.

### 3. Filtrowanie danych

Załóżmy, że chcemy filtrować osoby, które są pełnoletnie (wiek > 18) i mieszkają w Nowym Jorku.

```go
	// Filtracja osób pełnoletnich (Age > 18) z Nowego Jorku
	adultsInNY := funk.Filter(people, func(p Person) bool {
		return p.Age > 18 && p.City == "New York"
	}).([]Person)

	fmt.Println("Dorośli z Nowego Jorku:")
	for _, person := range adultsInNY {
		fmt.Printf("%s %s, Age: %d, City: %s\n", person.Fname, person.Lname, person.Age, person.City)
	}
```

#### Wyjaśnienie:
- **funk.Filter**: Funkcja `Filter` przetwarza slice `people` i zwraca nowy slice zawierający tylko te osoby, które spełniają warunki określone w funkcji anonimowej.
- **Funkcja anonimowa**: Filtruje osoby na podstawie dwóch warunków: `Age > 18` i `City == "New York"`.
- **Rzutowanie typu**: Wynik z `funk.Filter` jest rzutowany na `[]Person`, ponieważ `funk.Filter` zwraca `interface{}`.

### 4. Mapowanie danych

Następnie możemy stworzyć nowy slice zawierający tylko pełne imiona (czyli połączenie `Fname` i `Lname`) osób w slice.

```go
	// Mapowanie na pełne imiona (Fname + Lname)
	fullNames := funk.Map(people, func(p Person) string {
		return p.Fname + " " + p.Lname
	}).([]string)

	fmt.Println("\nPełne imiona wszystkich osób:")
	for _, name := range fullNames {
		fmt.Println(name)
	}
```

#### Wyjaśnienie:
- **funk.Map**: Funkcja `Map` przetwarza każdy element slice `people`, tworząc nowy slice zawierający tylko pełne imiona (`Fname + Lname`).
- **Rzutowanie typu**: Wynik z `funk.Map` jest rzutowany na `[]string`, ponieważ `funk.Map` zwraca `interface{}`.

### 5. Grupowanie danych

Możemy również pogrupować osoby według miasta, w którym mieszkają.

```go
	// Grupowanie osób według miasta
	groupedByCity := funk.GroupBy(people, func(p Person) string {
		return p.City
	}).(map[string][]Person)

	fmt.Println("\nGrupowanie osób według miasta:")
	for city, persons := range groupedByCity {
		fmt.Printf("City: %s\n", city)
		for _, person := range persons {
			fmt.Printf("  %s %s, Age: %d\n", person.Fname, person.Lname, person.Age)
		}
	}
```

#### Wyjaśnienie:
- **funk.GroupBy**: Funkcja `GroupBy` grupuje elementy slice na podstawie wartości zwracanej przez funkcję anonimową. W tym przypadku osoby są grupowane według miasta (`City`).
- **Rzutowanie typu**: Wynik z `funk.GroupBy` jest rzutowany na `map[string][]Person`, gdzie kluczem jest nazwa miasta, a wartością jest slice osób z tego miasta.

### 6. Wyszukiwanie elementu

Możemy wyszukać pierwszą osobę w slice, która spełnia określony warunek, np. osoba o imieniu "Tom".

```go
	// Wyszukiwanie pierwszej osoby o imieniu "Tom"
	tom, found := funk.Find(people, func(p Person) bool {
		return p.Fname == "Tom"
	}).(Person)

	if found {
		fmt.Printf("\nZnaleziono Toma: %s %s, Age: %d, City: %s\n", tom.Fname, tom.Lname, tom.Age, tom.City)
	} else {
		fmt.Println("\nNie znaleziono osoby o imieniu Tom.")
	}
```

#### Wyjaśnienie:
- **funk.Find**: Funkcja `Find` zwraca pierwszy element slice, który spełnia określony warunek. Jeśli element zostanie znaleziony, wynik jest rzutowany na typ `Person`.
- **Rzutowanie typu**: Wynik z `funk.Find` jest rzutowany na `Person`, ponieważ `funk.Find` zwraca `interface{}`.
- **`found`**: Wartość logiczna, która wskazuje, czy element został znaleziony.

### 7. Unikalne wartości

Możemy również uzyskać listę unikalnych miast, w których mieszkają osoby z naszego slice.

```go
	// Uzyskanie unikalnych miast
	uniqueCities := funk.Uniq(funk.Map(people, func(p Person) string {
		return p.City
	})).([]string)

	fmt.Println("\nUnikalne miasta:")
	for _, city := range uniqueCities {
		fmt.Println(city)
	}
```

#### Wyjaśnienie:
- **funk.Uniq**: Funkcja `Uniq` zwraca slice z unikalnymi wartościami.
- **Zagnieżdżone wywołanie `funk.Map`**: Najpierw mapujemy `people`, aby uzyskać slice miast (`City`), a następnie uzyskujemy unikalne miasta za pomocą `funk.Uniq`.

### 8. Redukcja danych

Na koniec możemy zsumować wiek wszystkich osób z slice.

```go
	// Sumowanie wieku wszystkich osób
	totalAge := funk.Reduce(people, func(acc int, p Person) int {
		return acc + p.Age
	}, 0).(int)

	fmt.Printf("\nSuma wieku wszystkich osób: %d\n", totalAge)
}
```

#### Wyjaśnienie:
- **funk.Reduce**: Funkcja `Reduce` przetwarza każdy element slice, aby zredukować go do jednej wartości (w tym przypadku suma wieku).
- **`acc`**: Akumulator, który przechowuje bieżącą sumę wieku.
- **Wartość początkowa `0`**: Wartość początkowa dla akumulatora, czyli 0.

Dodanie sortowania oraz znajdowania elementów wspólnych w dwóch slice'ach za pomocą biblioteki `go-funk` jest możliwe i rozszerza funkcjonalność naszego przykładowego programu. Poniżej znajdziesz zaktualizowany kod, który zawiera te dodatkowe operacje, a także szczegółowe wyjaśnienie każdej z nich.

### 9. Sortowanie elementów

Sortowanie slice osób według wieku i nazwiska w porządku rosnącym.

```go
	// Sortowanie osób według wieku, a następnie nazwiska
	sortedPeople := funk.Sort(people, func(a, b Person) bool {
		if a.Age == b.Age {
			return a.Lname < b.Lname
		}
		return a.Age < b.Age
	}).([]Person)

	fmt.Println("Osoby posortowane według wieku i nazwiska:")
	for _, person := range sortedPeople {
		fmt.Printf("%s %s, Age: %d, City: %s\n", person.Fname, person.Lname, person.Age, person.City)
	}
```

#### Wyjaśnienie:
- **funk.Sort**: Funkcja `Sort` sortuje slice na podstawie funkcji porównawczej, która przyjmuje dwa elementy i zwraca `true`, jeśli pierwszy element powinien być przed drugim.
- **Funkcja porównawcza**: Sortuje najpierw według `Age` w porządku rosnącym. Jeśli dwie osoby mają ten sam wiek, sortuje według `Lname` (nazwiska).

### 10. Znajdowanie elementów wspólnych w dwóch slice'ach

Następnie możemy znaleźć osoby, które mają imiona wspólne z drugim slice'em `names`.

```go
	// Znajdowanie osób, których imiona są wspólne z slice 'names'
	commonPeople := funk.Filter(people, func(p Person) bool {
		return funk.Contains(names, p.Fname)
	}).([]Person)

	fmt.Println("\nOsoby, których imiona są wspólne z podanym slice:")
	for _, person := range commonPeople {
		fmt.Printf("%s %s, Age: %d, City: %s\n", person.Fname, person.Lname, person.Age, person.City)
	}
```

#### Wyjaśnienie:
- **funk.Filter**: Filtruje slice `people`, zachowując tylko te osoby, których `Fname` znajduje się w slice `names`.
- **funk.Contains**: Sprawdza, czy `names` zawiera daną wartość `Fname`.

