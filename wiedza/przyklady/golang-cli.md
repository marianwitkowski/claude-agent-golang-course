### 1. Instalacja biblioteki Cobra

Aby skorzystać z biblioteki Cobra, musisz ją najpierw zainstalować:

```bash
go get -u github.com/spf13/cobra@latest
```

### 2. Struktura programu

Program będzie miał główną komendę `config` (lub podobną nazwę), która przyjmie różne typy argumentów:
- **Wartość numeryczna**: Liczba całkowita z określonego zakresu.
- **Wartość tekstowa**: Dowolny ciąg znaków.
- **Wartość logiczna**: Flaga true/false.
- **Lista wartości**: Lista ciągów znaków (stringów).

### 3. Kod programu

```go
package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "app",
	Short: "A CLI tool for demonstrating input parsing with Cobra",
	Long: `This tool demonstrates how to use Cobra to parse different types of input from the command line, 
including numerical values, strings, booleans, and lists.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Main logic goes here
		fmt.Println("Please specify options using the flags provided.")
	},
}

// Definiowanie zmiennych globalnych do przechowywania wartości flag
var (
	numValue   int
	textValue  string
	boolValue  bool
	listValues []string
)

func init() {
	// Definiowanie flag komendy root
	rootCmd.PersistentFlags().IntVarP(&numValue, "num", "n", 0, "A numeric value (must be between 1 and 100)")
	rootCmd.PersistentFlags().StringVarP(&textValue, "text", "t", "", "A text value")
	rootCmd.PersistentFlags().BoolVarP(&boolValue, "bool", "b", false, "A boolean value")
	rootCmd.PersistentFlags().StringSliceVarP(&listValues, "list", "l", []string{}, "A list of values separated by commas")

	// Dodanie walidacji flagi numValue
	rootCmd.PersistentPreRunE = func(cmd *cobra.Command, args []string) error {
		if numValue < 1 || numValue > 100 {
			return errors.New("num value must be between 1 and 100")
		}
		return nil
	}

	// Dodanie komendy root do uruchomienia
	rootCmd.Run = func(cmd *cobra.Command, args []string) {
		fmt.Printf("Numeric Value: %d\n", numValue)
		fmt.Printf("Text Value: %s\n", textValue)
		fmt.Printf("Boolean Value: %v\n", boolValue)
		fmt.Printf("List Values: %s\n", strings.Join(listValues, ", "))
	}
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}
```

### 4. Szczegółowy opis kodu

#### 4.1. Importowanie bibliotek

```go
import (
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)
```

- **errors**: Standardowa biblioteka do tworzenia błędów.
- **fmt**: Biblioteka do formatowania tekstu i interakcji z konsolą.
- **log**: Biblioteka do logowania informacji i błędów.
- **os**: Biblioteka do interakcji z systemem operacyjnym, np. do uzyskiwania dostępu do zmiennych środowiskowych.
- **strconv**: Używana do konwersji wartości tekstowych na liczby, jeśli zajdzie taka potrzeba.
- **strings**: Biblioteka do manipulacji ciągami znaków.
- **cobra**: Główna biblioteka do tworzenia aplikacji CLI.

#### 4.2. Definiowanie głównej komendy (root command)

```go
var rootCmd = &cobra.Command{
	Use:   "app",
	Short: "A CLI tool for demonstrating input parsing with Cobra",
	Long: `This tool demonstrates how to use Cobra to parse different types of input from the command line, 
including numerical values, strings, booleans, and lists.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Main logic goes here
		fmt.Println("Please specify options using the flags provided.")
	},
}
```

- **Use**: Nazwa aplikacji, którą użytkownik wpisuje w wierszu poleceń.
- **Short**: Krótkie podsumowanie, które pojawia się, gdy użytkownik wywołuje pomoc (`--help`).
- **Long**: Dłuższy opis, który pojawia się w szczegółowej dokumentacji.
- **Run**: Funkcja, która jest wykonywana, gdy nie zostanie podana żadna subkomenda. W tym przypadku wyświetla informację, aby użytkownik użył dostępnych flag.

#### 4.3. Definiowanie zmiennych globalnych do przechowywania wartości flag

```go
var (
	numValue   int
	textValue  string
	boolValue  bool
	listValues []string
)
```

- **numValue**: Zmienna do przechowywania wartości numerycznej.
- **textValue**: Zmienna do przechowywania wartości tekstowej.
- **boolValue**: Zmienna do przechowywania wartości logicznej (true/false).
- **listValues**: Zmienna do przechowywania listy wartości (ciągów znaków).

#### 4.4. Inicjalizacja flag i walidacja wartości numerycznej

```go
func init() {
	// Definiowanie flag komendy root
	rootCmd.PersistentFlags().IntVarP(&numValue, "num", "n", 0, "A numeric value (must be between 1 and 100)")
	rootCmd.PersistentFlags().StringVarP(&textValue, "text", "t", "", "A text value")
	rootCmd.PersistentFlags().BoolVarP(&boolValue, "bool", "b", false, "A boolean value")
	rootCmd.PersistentFlags().StringSliceVarP(&listValues, "list", "l", []string{}, "A list of values separated by commas")

	// Dodanie walidacji flagi numValue
	rootCmd.PersistentPreRunE = func(cmd *cobra.Command, args []string) error {
		if numValue < 1 || numValue > 100 {
			return errors.New("num value must be between 1 and 100")
		}
		return nil
	}

	// Dodanie komendy root do uruchomienia
	rootCmd.Run = func(cmd *cobra.Command, args []string) {
		fmt.Printf("Numeric Value: %d\n", numValue)
		fmt.Printf("Text Value: %s\n", textValue)
		fmt.Printf("Boolean Value: %v\n", boolValue)
		fmt.Printf("List Values: %s\n", strings.Join(listValues, ", "))
	}
}
```

- **PersistentFlags**: Flagi globalne, które są dostępne dla wszystkich subkomend (choć w tym przykładzie mamy tylko jedną komendę).
- **IntVarP**: Ustawia flagę numeryczną `--num` (skróconą do `-n`), z wartością domyślną 0.
- **StringVarP**: Ustawia flagę tekstową `--text` (skróconą do `-t`), z wartością domyślną pusty ciąg znaków.
- **BoolVarP**: Ustawia flagę logiczną `--bool` (skróconą do `-b`), z wartością domyślną `false`.
- **StringSliceVarP**: Ustawia flagę, która akceptuje listę wartości `--list` (skróconą do `-l`), z wartością domyślną pustą listą.
- **PersistentPreRunE**: Funkcja, która jest wykonywana przed uruchomieniem głównej komendy. W tym przypadku służy do walidacji wartości numerycznej `numValue`. Jeśli wartość `numValue` nie mieści się w przedziale 1-100, zwracany jest błąd.

#### 4.5. Funkcja `main`

```go
func main() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}
```

- **rootCmd.Execute**: Funkcja, która wywołuje przetwarzanie komend. Cobra automatycznie zajmie się rozpoznaniem, którą komendę użytkownik chciał uruchomić, oraz przekaże odpowiednie argumenty i flagi.

- **Error handling**: W przypadku błędu (np. niepoprawne argumenty), program zakończy działanie

 i wyświetli błąd.

### 5. Testowanie programu

Możesz przetestować program, uruchamiając go z różnymi argumentami w wierszu poleceń:

- **Przykład uruchomienia**:
  ```bash
  go run main.go --num 42 --text "Hello, World!" --bool --list item1,item2,item3
  ```

  Wynik:
  ```
  Numeric Value: 42
  Text Value: Hello, World!
  Boolean Value: true
  List Values: item1, item2, item3
  ```

- **Przykład walidacji**:
  ```bash
  go run main.go --num 150 --text "Hello, World!" --bool --list item1,item2,item3
  ```

  Wynik:
  ```
  Error: num value must be between 1 and 100
  ```

### 6. Podsumowanie

Ten program pokazuje, jak używać biblioteki Cobra do parsowania różnych typów danych przekazywanych z linii poleceń w aplikacji Go. Program przyjmuje wartość numeryczną, tekstową, logiczną oraz listę wartości, a także weryfikuje, czy wartość numeryczna mieści się w określonym przedziale. Dzięki zastosowaniu Cobra, zarządzanie tymi parametrami jest intuicyjne i łatwe do rozbudowy o kolejne funkcjonalności.