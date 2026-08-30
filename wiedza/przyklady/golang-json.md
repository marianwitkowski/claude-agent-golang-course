## Obsługa JSON w Golang

W Go (Golang) parsowanie, serializowanie (kodowanie) i deserializowanie (dekodowanie) danych JSON jest łatwe dzięki standardowej bibliotece `encoding/json`. Poniżej znajdziesz przykłady kodu oraz wyjaśnienia dotyczące tego, jak te operacje działają, z naciskiem na wydajność przy pracy z dużymi plikami JSON.

### 1. Podstawowe operacje: Serializacja i deserializacja JSON

#### Serializacja: Konwersja struktury Go do JSON

```go
package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Age       int    `json:"age"`
	City      string `json:"city"`
}

func main() {
	person := Person{
		FirstName: "John",
		LastName:  "Doe",
		Age:       30,
		City:      "New York",
	}

	// Serializacja struktury Go do JSON
	personJSON, err := json.Marshal(person)
	if err != nil {
		fmt.Println("Error serializing to JSON:", err)
		return
	}

	fmt.Println("Serialized JSON:", string(personJSON))
}
```

#### Wyjaśnienie:
- **Struktura Person**: Definiujemy strukturę `Person`, która reprezentuje dane osoby. Każde pole jest oznaczone tagiem JSON, który wskazuje nazwę klucza w wynikowym JSON-ie.
- **json.Marshal**: Funkcja `Marshal` konwertuje strukturę Go do JSON, zwracając bajty (`[]byte`), które można przekonwertować na string i wyświetlić.

#### Deserializacja: Konwersja JSON do struktury Go

```go
package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	personJSON := `{"first_name":"John","last_name":"Doe","age":30,"city":"New York"}`

	var person Person

	// Deserializacja JSON do struktury Go
	err := json.Unmarshal([]byte(personJSON), &person)
	if err != nil {
		fmt.Println("Error deserializing JSON:", err)
		return
	}

	fmt.Printf("Deserialized Person: %+v\n", person)
}
```

#### Wyjaśnienie:
- **json.Unmarshal**: Funkcja `Unmarshal` konwertuje dane JSON (w formie bajtów) do struktury Go. Wynik jest przypisywany do zmiennej `person`.

### 2. Praca z dużymi plikami JSON

Przy pracy z dużymi plikami JSON (np. powyżej 50 MB) wydajność i zużycie pamięci stają się kluczowymi czynnikami. W Go, zamiast ładować cały plik JSON do pamięci na raz, można użyć strumieniowego przetwarzania JSON za pomocą dekodera (`json.Decoder`). Pozwala to na przetwarzanie pliku JSON w częściach, co jest bardziej wydajne pamięciowo.

#### Strumieniowe przetwarzanie dużych plików JSON

Załóżmy, że mamy duży plik JSON zawierający listę osób, i chcemy przetwarzać te dane w sposób wydajny.

```go
package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Person struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Age       int    `json:"age"`
	City      string `json:"city"`
}

func main() {
	// Otwieranie dużego pliku JSON
	file, err := os.Open("large_file.json")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Tworzenie dekodera JSON
	decoder := json.NewDecoder(file)

	// Oczekujemy, że plik JSON zawiera listę osób (array)
	_, err = decoder.Token() // Odczytanie początkowego tokena, który powinien być '['
	if err != nil {
		fmt.Println("Error reading start of JSON array:", err)
		return
	}

	// Iteracja przez elementy JSON array
	for decoder.More() {
		var person Person
		err := decoder.Decode(&person)
		if err != nil {
			fmt.Println("Error decoding JSON:", err)
			return
		}

		// Przetwarzanie danych osoby (możesz tu dodać logikę przetwarzania)
		fmt.Printf("Person: %+v\n", person)
	}

	// Odczytanie końcowego tokena, który powinien być ']'
	_, err = decoder.Token()
	if err != nil {
		fmt.Println("Error reading end of JSON array:", err)
		return
	}
}
```

#### Wyjaśnienie:
- **os.Open**: Otwiera plik JSON bez ładowania go w całości do pamięci.
- **json.NewDecoder**: Tworzy dekoder, który przetwarza JSON w sposób strumieniowy.
- **decoder.Token()**: Odczytuje pojedyncze tokeny JSON, co jest potrzebne na początku i na końcu przetwarzania array'a JSON.
- **decoder.More()**: Sprawdza, czy w strumieniu JSON są jeszcze elementy do przetworzenia.
- **decoder.Decode(&person)**: Dekoduje każdy element JSON (struktura `Person`) ze strumienia.

### 3. Serializacja dużych danych

Jeśli chcesz zapisać dużą ilość danych do pliku JSON, również możesz to zrobić w sposób strumieniowy, zamiast serializować wszystko na raz.

```go
package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	people := []Person{
		{"John", "Doe", 16, "New York"},
		{"Jane", "Smith", 20, "Los Angeles"},
		{"Bob", "Johnson", 25, "New York"},
		// ... więcej danych
	}

	// Tworzenie pliku JSON
	file, err := os.Create("large_output.json")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	// Tworzenie enkodera JSON
	encoder := json.NewEncoder(file)

	// Rozpoczynamy zapis od '['
	file.WriteString("[")

	for i, person := range people {
		err := encoder.Encode(person)
		if err != nil {
			fmt.Println("Error encoding JSON:", err)
			return
		}
		// Dodaj przecinek między elementami (ale nie po ostatnim elemencie)
		if i < len(people)-1 {
			file.WriteString(",")
		}
	}

	// Zakończenie zapisu ']'
	file.WriteString("]")
}
```

#### Wyjaśnienie:
- **json.NewEncoder**: Tworzy enkoder do strumieniowego zapisywania danych JSON do pliku.
- **encoder.Encode(person)**: Zapisuje pojedynczą strukturę `Person` jako JSON do pliku, co pozwala na strumieniowy zapis dużych ilości danych.

### 4. Wydajność

**Strumieniowe przetwarzanie JSON** jest kluczowe dla pracy z dużymi plikami. Oto kilka powodów:

- **Zużycie pamięci**: Strumieniowe przetwarzanie umożliwia przetwarzanie danych częściami, co znacznie zmniejsza zużycie pamięci, w przeciwieństwie do ładowania całego pliku JSON do pamięci.
- **Przepływ przetwarzania**: Dzięki strumieniowemu przetwarzaniu możesz natychmiast zacząć przetwarzanie danych, bez konieczności czekania na załadowanie całego pliku.
- **Elastyczność**: Możesz łatwo przetwarzać dane sekwencyjnie, co jest szczególnie przydatne w przypadku aplikacji przetwarzających strumienie danych w czasie rzeczywistym.


## Alternatywne biblioteki

W Go (Golang) standardowa biblioteka `encoding/json` jest najczęściej używana do obsługi JSON, ale istnieje również kilka alternatywnych bibliotek, które oferują dodatkowe funkcje, lepszą wydajność lub inne zalety. Poniżej opisano kilka z najbardziej popularnych alternatyw do pracy z JSON w Go:

### 1. **`jsoniter` (JSONIterator)**

- **Opis**: `jsoniter` (lub `json-iterator`) jest jedną z najbardziej popularnych alternatyw dla standardowej biblioteki JSON w Go. Została zaprojektowana z myślą o większej wydajności i elastyczności. `jsoniter` jest w pełni zgodny z `encoding/json`, więc można go łatwo zastąpić w istniejącym kodzie.

- **Zalety**:
  - **Wydajność**: `jsoniter` jest zwykle szybszy niż standardowa biblioteka JSON, szczególnie przy dużych plikach lub intensywnym przetwarzaniu JSON.
  - **Obsługa bardziej skomplikowanych typów danych**: `jsoniter` lepiej radzi sobie z niektórymi skomplikowanymi przypadkami, np. z dużymi liczbami całkowitymi.

- **Instalacja**:
  ```bash
  go get -u github.com/json-iterator/go
  ```

- **Przykład użycia**:
  ```go
  package main

  import (
      "fmt"
      jsoniter "github.com/json-iterator/go"
  )

  var json = jsoniter.ConfigCompatibleWithStandardLibrary

  type Person struct {
      FirstName string `json:"first_name"`
      LastName  string `json:"last_name"`
      Age       int    `json:"age"`
      City      string `json:"city"`
  }

  func main() {
      person := Person{
          FirstName: "John",
          LastName:  "Doe",
          Age:       30,
          City:      "New York",
      }

      // Serializacja do JSON
      personJSON, err := json.Marshal(person)
      if err != nil {
          fmt.Println("Error serializing to JSON:", err)
          return
      }
      fmt.Println("Serialized JSON:", string(personJSON))

      // Deserializacja JSON do struktury
      var personDeserialized Person
      err = json.Unmarshal(personJSON, &personDeserialized)
      if err != nil {
          fmt.Println("Error deserializing JSON:", err)
          return
      }
      fmt.Printf("Deserialized Person: %+v\n", personDeserialized)
  }
  ```

### 2. **`gjson`**

- **Opis**: `gjson` to biblioteka przeznaczona do szybkiego i łatwego odczytywania wartości z JSON-a bez konieczności deserializacji całej struktury. Umożliwia bezpośredni dostęp do poszczególnych pól JSON za pomocą ścieżek podobnych do XPath lub JPath.

- **Zalety**:
  - **Wydajność**: Bardzo szybkie przeszukiwanie i ekstrakcja wartości z dużych struktur JSON.
  - **Prosta składnia**: Łatwe wyszukiwanie danych bez konieczności definiowania struktur Go.

- **Instalacja**:
  ```bash
  go get -u github.com/tidwall/gjson
  ```

- **Przykład użycia**:
  ```go
  package main

  import (
      "fmt"
      "github.com/tidwall/gjson"
  )

  func main() {
      jsonData := `{"name":{"first":"John","last":"Doe"},"age":30,"city":"New York"}`

      // Ekstrakcja danych przy użyciu ścieżki JSON
      firstName := gjson.Get(jsonData, "name.first")
      age := gjson.Get(jsonData, "age")

      fmt.Println("First Name:", firstName.String())
      fmt.Println("Age:", age.Int())
  }
  ```

### 3. **`go-json`**

- **Opis**: `go-json` to kolejna alternatywna implementacja JSON w Go, która jest zoptymalizowana pod kątem wydajności. Celem projektu jest dostarczenie jeszcze szybszego parsowania JSON niż `jsoniter`.

- **Zalety**:
  - **Wydajność**: Jeszcze szybsza niż `jsoniter` w niektórych przypadkach, szczególnie przy dużych zbiorach danych.
  - **Zgodność**: Biblioteka stara się być w pełni zgodna z `encoding/json`.

- **Instalacja**:
  ```bash
  go get -u github.com/goccy/go-json
  ```

- **Przykład użycia**:
  ```go
  package main

  import (
      "fmt"
      "github.com/goccy/go-json"
  )

  type Person struct {
      FirstName string `json:"first_name"`
      LastName  string `json:"last_name"`
      Age       int    `json:"age"`
      City      string `json:"city"`
  }

  func main() {
      person := Person{
          FirstName: "John",
          LastName:  "Doe",
          Age:       30,
          City:      "New York",
      }

      // Serializacja do JSON
      personJSON, err := json.Marshal(person)
      if err != nil {
          fmt.Println("Error serializing to JSON:", err)
          return
      }
      fmt.Println("Serialized JSON:", string(personJSON))

      // Deserializacja JSON do struktury
      var personDeserialized Person
      err = json.Unmarshal(personJSON, &personDeserialized)
      if err != nil {
          fmt.Println("Error deserializing JSON:", err)
          return
      }
      fmt.Printf("Deserialized Person: %+v\n", personDeserialized)
  }
  ```

### 4. **`jettison`**

- **Opis**: `jettison` to biblioteka do szybkiej serializacji JSON w Go, zaprojektowana z myślą o wydajności. Skupia się głównie na szybkim kodowaniu JSON, z pominięciem niektórych funkcji dekodowania.

- **Zalety**:
  - **Wydajność**: Bardzo szybkie kodowanie JSON.
  - **Optymalizacja pamięci**: Zaprojektowana z myślą o minimalnym zużyciu pamięci.

- **Instalacja**:
  ```bash
  go get -u github.com/wI2L/jettison
  ```

- **Przykład użycia**:
  ```go
  package main

  import (
      "fmt"
      "github.com/wI2L/jettison"
  )

  type Person struct {
      FirstName string `json:"first_name"`
      LastName  string `json:"last_name"`
      Age       int    `json:"age"`
      City      string `json:"city"`
  }

  func main() {
      person := Person{
          FirstName: "John",
          LastName:  "Doe",
          Age:       30,
          City:      "New York",
      }

      // Serializacja do JSON za pomocą jettison
      personJSON, err := jettison.Marshal(person)
      if err != nil {
          fmt.Println("Error serializing to JSON:", err)
          return
      }
      fmt.Println("Serialized JSON:", string(personJSON))
  }
  ```

### 5. **`ffjson`**

- **Opis**: `ffjson` to biblioteka, która generuje specjalizowany kod do serializacji i deserializacji JSON, aby przyspieszyć te operacje. `ffjson` oferuje możliwość wygenerowania kodu, który jest szybszy niż standardowe `encoding/json`.

- **Zalety**:
  - **Wydajność**: Optymalizuje operacje JSON poprzez generowanie kodu dostosowanego do konkretnych struktur danych.
  - **Automatyzacja**: Narzędzie generujące kod eliminuje konieczność ręcznego pisania zoptymalizowanych funkcji JSON.

- **Instalacja**:
  ```bash
  go get -u github.com/pquerna/ffjson
  ```

- **Przykład użycia** (z generowaniem kodu):
  ```bash
  ffjson <file>.go
  ```

  - Następnie można użyć wygenerowanego kodu w programie Go, co przyspieszy operacje JSON.

### Podsumowanie

Każda z alternatywnych bibliotek ma swoje unikalne zalety, które mogą uczynić je bardziej odpowiednimi w określonych scenariuszach:

- **`jsoniter`**: Ogólnie uznawana za szybszą alternatywę dla `encoding/json`, szczególnie przy dużych i skomplikowanych strukturach JSON.
- **`gjson`**: Idealna do szybkiego przeszukiwania i ekstrakcji danych z JSON bez pełnej deserializacji.
- **`go-json`**: Jeszcze szybsza opcja niż `jsoniter`, szczególnie wydajna w pracy z dużymi plikami JSON.
- **`jettison`**: Skupia się na szybkim kodowaniu JSON z optymalizacją pamięci.


- **`ffjson`**: Narzędzie generujące kod, które pozwala na wygenerowanie zoptymalizowanych funkcji JSON dla specyficznych struktur danych.

Wybór odpowiedniej biblioteki zależy od wymagań projektu, takich jak wydajność, zużycie pamięci, oraz specyficzne operacje, które muszą być wykonane na danych JSON.