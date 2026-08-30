## Parsowanie danych CSV

W Go (Golang) parsowanie danych z plików CSV jest prostym zadaniem dzięki standardowej bibliotece `encoding/csv`. W przypadku dużych plików CSV (powyżej 50 MB), ważne jest, aby przetwarzać dane w sposób wydajny, minimalizując zużycie pamięci. Można to osiągnąć, korzystając ze strumieniowego przetwarzania danych.

Poniżej znajduje się przykładowy kod, który pokazuje, jak parsować duży plik CSV w sposób wydajny. Każdy fragment kodu jest szczegółowo opisany.

### 1. Podstawowe parsowanie CSV

Zacznijmy od podstawowego przykładu, który pokazuje, jak odczytać dane z pliku CSV i wyświetlić je na konsoli.

```go
package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func main() {
	// Otwieranie pliku CSV
	file, err := os.Open("large_file.csv")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Tworzenie nowego czytnika CSV
	reader := csv.NewReader(file)

	// Odczytanie wszystkich rekordów (niezalecane dla dużych plików)
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error reading CSV:", err)
		return
	}

	// Wyświetlenie odczytanych rekordów
	for _, record := range records {
		fmt.Println(record)
	}
}
```

#### Wyjaśnienie:
- **os.Open**: Otwiera plik CSV bez ładowania go w całości do pamięci.
- **csv.NewReader**: Tworzy nowy czytnik CSV, który będzie przetwarzał dane linia po linii.
- **reader.ReadAll**: Odczytuje wszystkie rekordy na raz. Ta metoda nie jest zalecana dla dużych plików, ponieważ może zużyć dużo pamięci.

### 2. Wydajne przetwarzanie dużych plików CSV

W przypadku dużych plików CSV, zamiast używać `ReadAll`, który ładuje cały plik do pamięci, lepiej jest przetwarzać plik linia po linii, używając metody `Read`.

```go
package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func main() {
	// Otwieranie pliku CSV
	file, err := os.Open("large_file.csv")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Tworzenie nowego czytnika CSV
	reader := csv.NewReader(file)

	// Odczytywanie linii po linii
	for {
		record, err := reader.Read()
		if err != nil {
			// Jeśli dojdziemy do końca pliku
			if err.Error() == "EOF" {
				break
			}
			fmt.Println("Error reading line:", err)
			return
		}

		// Przetwarzanie każdego rekordu
		processRecord(record)
	}
}

func processRecord(record []string) {
	// Tutaj można dodać logikę przetwarzania rekordu, np. zapisanie do bazy danych
	fmt.Println("Processing record:", record)
}
```

#### Wyjaśnienie:
- **reader.Read**: Odczytuje jedną linię z pliku CSV naraz. Zwraca `[]string`, które reprezentuje rekord CSV (każda wartość w osobnej komórce tabeli).
- **EOF**: Sprawdzamy, czy doszliśmy do końca pliku, aby zakończyć pętlę przetwarzania.

### 3. Optymalizacje i wskazówki dotyczące wydajności

Przy pracy z bardzo dużymi plikami CSV (np. powyżej 100 MB), warto rozważyć dodatkowe optymalizacje:

- **Buforowane odczytywanie**: Można zwiększyć wydajność odczytu pliku, używając buforowanego odczytywania przy pomocy `bufio.Reader`.

- **Przetwarzanie równoległe**: Jeśli przetwarzanie każdego rekordu jest niezależne od innych, można rozważyć wykorzystanie gorutyn do równoległego przetwarzania danych.

- **Zmniejszenie zużycia pamięci**: Jeśli w trakcie przetwarzania dane są zapisywane do bazy danych lub innego magazynu, można ograniczyć przechowywanie danych w pamięci.

### 4. Przykład z buforowanym odczytem i równoległym przetwarzaniem

Poniższy przykład pokazuje, jak użyć `bufio.Reader` dla buforowanego odczytu i równoległego przetwarzania danych przy użyciu gorutyn.

```go
package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"sync"
)

func main() {
	// Otwieranie pliku CSV
	file, err := os.Open("large_file.csv")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Tworzenie buforowanego czytnika
	bufferedReader := bufio.NewReader(file)
	reader := csv.NewReader(bufferedReader)

	// Używanie grupy WaitGroup do synchronizacji gorutyn
	var wg sync.WaitGroup

	// Odczytywanie linii po linii
	for {
		record, err := reader.Read()
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			fmt.Println("Error reading line:", err)
			return
		}

		// Inkrementacja licznika WaitGroup
		wg.Add(1)

		// Przetwarzanie rekordu w osobnej gorutynie
		go func(record []string) {
			defer wg.Done()
			processRecord(record)
		}(record)
	}

	// Czekamy na zakończenie wszystkich gorutyn
	wg.Wait()
}

func processRecord(record []string) {
	// Tutaj można dodać logikę przetwarzania rekordu
	fmt.Println("Processing record:", record)
}
```

#### Wyjaśnienie:
- **bufio.NewReader**: Tworzy buforowany czytnik, co może zwiększyć wydajność odczytu danych z pliku.
- **sync.WaitGroup**: Używana do synchronizacji gorutyn, aby program poczekał na zakończenie przetwarzania wszystkich rekordów przed zakończeniem działania.
- **gorutyny**: Umożliwiają równoległe przetwarzanie rekordów, co może znacząco przyspieszyć przetwarzanie dużych plików CSV.


## Alterrnatywne biblioteki

W Go, oprócz standardowej biblioteki `encoding/csv`, istnieje kilka alternatywnych bibliotek do pracy z plikami CSV. Te biblioteki oferują dodatkowe funkcje, lepszą wydajność lub specyficzne możliwości, które mogą być przydatne w zależności od wymagań projektu. Poniżej przedstawiam najpopularniejsze alternatywne biblioteki do pracy z CSV w Go.

### 1. **`github.com/gocarina/gocsv`**

- **Opis**: `gocsv` to biblioteka, która upraszcza pracę z plikami CSV poprzez automatyczne mapowanie danych CSV do struktur Go i odwrotnie. Jest szczególnie przydatna, gdy chcemy pracować z plikami CSV w sposób bardziej strukturalny.

- **Zalety**:
  - Automatyczne mapowanie danych CSV do struktur Go i odwrotnie.
  - Obsługa nagłówków w plikach CSV, co ułatwia pracę z danymi o zdefiniowanym formacie.
  - Łatwa w użyciu składnia, która redukuje ilość kodu potrzebnego do obsługi CSV.

- **Instalacja**:
  ```bash
  go get -u github.com/gocarina/gocsv
  ```

- **Przykład użycia**:
  ```go
  package main

  import (
      "fmt"
      "github.com/gocarina/gocsv"
      "os"
  )

  type Person struct {
      FirstName string `csv:"first_name"`
      LastName  string `csv:"last_name"`
      Age       int    `csv:"age"`
      City      string `csv:"city"`
  }

  func main() {
      file, err := os.OpenFile("people.csv", os.O_RDWR|os.O_CREATE, os.ModePerm)
      if err != nil {
          fmt.Println("Error opening file:", err)
          return
      }
      defer file.Close()

      // Wczytywanie danych z CSV do slice struktury Person
      var people []Person
      if err := gocsv.UnmarshalFile(file, &people); err != nil {
          fmt.Println("Error unmarshalling CSV:", err)
          return
      }

      // Wyświetlanie danych
      for _, person := range people {
          fmt.Printf("%+v\n", person)
      }

      // Serializacja danych z powrotem do CSV
      fileOut, err := os.Create("output.csv")
      if err != nil {
          fmt.Println("Error creating file:", err)
          return
      }
      defer fileOut.Close()

      if err := gocsv.MarshalFile(&people, fileOut); err != nil {
          fmt.Println("Error marshalling CSV:", err)
      }
  }
  ```

- **Podsumowanie**: `gocsv` to świetny wybór, gdy potrzebujesz wygodnego sposobu na mapowanie danych CSV do struktur Go i odwrotnie, szczególnie gdy pliki CSV zawierają nagłówki.

### 2. **`github.com/jszwec/csvutil`**

- **Opis**: `csvutil` to biblioteka, która koncentruje się na wysokiej wydajności oraz elastycznym mapowaniu danych CSV do struktur Go. Jest bardziej niskopoziomowa niż `gocsv`, co daje większą kontrolę nad procesem przetwarzania CSV.

- **Zalety**:
  - Wysoka wydajność dzięki optymalizacji parsowania CSV.
  - Obsługa bardziej zaawansowanych przypadków użycia, takich jak niestandardowe typy, aliasy kolumn, itp.
  - Elastyczne mapowanie CSV do struktur Go z możliwością użycia tagów strukturalnych.

- **Instalacja**:
  ```bash
  go get -u github.com/jszwec/csvutil
  ```

- **Przykład użycia**:
  ```go
  package main

  import (
      "encoding/csv"
      "fmt"
      "github.com/jszwec/csvutil"
      "os"
  )

  type Person struct {
      FirstName string `csv:"first_name"`
      LastName  string `csv:"last_name"`
      Age       int    `csv:"age"`
      City      string `csv:"city"`
  }

  func main() {
      file, err := os.Open("people.csv")
      if err != nil {
          fmt.Println("Error opening file:", err)
          return
      }
      defer file.Close()

      // Tworzenie czytnika CSV
      reader := csv.NewReader(file)

      // Dekodowanie CSV do slice struktur Go
      var people []Person
      dec, err := csvutil.NewDecoder(reader)
      if err != nil {
          fmt.Println("Error creating decoder:", err)
          return
      }

      for {
          var p Person
          if err := dec.Decode(&p); err != nil {
              if err.Error() == "EOF" {
                  break
              }
              fmt.Println("Error decoding CSV:", err)
              return
          }
          people = append(people, p)
      }

      // Wyświetlanie danych
      for _, person := range people {
          fmt.Printf("%+v\n", person)
      }
  }
  ```

- **Podsumowanie**: `csvutil` jest doskonałym wyborem, gdy potrzebujesz wysokiej wydajności i elastyczności przy mapowaniu CSV do struktur Go, szczególnie w bardziej skomplikowanych przypadkach.

### 3. **`github.com/tealeg/xlsx` (dla plików Excel)**

- **Opis**: Choć `xlsx` nie jest biblioteką bezpośrednio do obsługi CSV, jest ona przydatna, jeśli masz do czynienia z plikami Excel, które chcesz przetworzyć na dane CSV. Biblioteka ta pozwala na odczytywanie i zapisywanie plików Excel (XLSX), a także konwersję arkuszy Excel do formatu CSV.

- **Zalety**:
  - Obsługa plików Excel, które mogą być konwertowane na CSV.
  - Umożliwia przetwarzanie danych złożonych w formatach, takich jak XLSX.

- **Instalacja**:
  ```bash
  go get -u github.com/tealeg/xlsx
  ```

- **Przykład użycia**:
  ```go
  package main

  import (
      "fmt"
      "github.com/tealeg/xlsx"
  )

  func main() {
      file, err := xlsx.OpenFile("example.xlsx")
      if err != nil {
          fmt.Println("Error opening file:", err)
          return
      }

      for _, sheet := range file.Sheets {
          for _, row := range sheet.Rows {
              var rowData []string
              for _, cell := range row.Cells {
                  text := cell.String()
                  rowData = append(rowData, text)
              }
              fmt.Println(rowData)
          }
      }
  }
  ```

- **Podsumowanie**: `xlsx` jest świetnym wyborem, jeśli pracujesz z plikami Excel, ale potrzebujesz konwertować te dane na CSV lub przetwarzać dane z tych plików w inny sposób.

### 4. **`github.com/bsm/csvutil` (nowsza, minimalna alternatywa)**

- **Opis**: Jest to lekka biblioteka, która koncentruje się na minimalizmie i wydajności przy pracy z plikami CSV. `bsm/csvutil` oferuje podstawowe funkcje do przetwarzania plików CSV, będąc jednocześnie bardzo wydajną.

- **Zalety**:
  - Bardzo lekka i szybka biblioteka.
  - Prosta w użyciu, z minimalnym interfejsem.

- **Instalacja**:
  ```bash
  go get -u github.com/bsm/csvutil
  ```

- **Podsumowanie**: Ta biblioteka jest idealna, jeśli potrzebujesz minimalnej, ale wydajnej alternatywy do pracy z plikami CSV, bez zbędnych funkcji.

### Podsumowanie

Każda z wymienionych bibliotek do obsługi CSV w Go ma swoje unikalne zalety i jest odpowiednia do różnych przypadków użycia:

- **`gocsv`**: Idealna do prostego i szybkiego mapowania danych CSV do struktur Go i odwrotnie, zwłaszcza przy pracy z plikami zawierającymi nagłówki.
- **`csvutil`**: Skierowana do zaawansowanych użytkowników potrzebujących wysokiej wydajności i elastyczności przy przetwarzaniu CSV.
- **`xlsx`**: Przydatna, jeśli musisz pracować z plikami Excel i konwertować je do formatu CSV.
- **`bsm/csvutil`**: Minimalistyczna, ale wydajna alternatywa, idealna dla prostych, wydajnych operacji na plikach CSV.

Wybór odpowiedniej biblioteki zależy od specyfiki projektu, wymagań dotyczących wydajności oraz tego, jak skomplikowane są dane, które chcesz przetwarzać.