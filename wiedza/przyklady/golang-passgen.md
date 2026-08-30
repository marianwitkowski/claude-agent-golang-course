### 1. Generator haseł
Program pozwala na wygenerowanie hasła o żądanej długości i poziomie trudności. Dane wejściowe wprowadzane są ze standardowego wejścia.


### 2. Importowanie bibliotek

```go
import (
	"crypto/rand"
	"fmt"
	"math/big"
	"strconv"
	"strings"
)
```

#### Wyjaśnienie:

- **"crypto/rand"**: Używana do generowania kryptograficznie bezpiecznych losowych liczb, co jest istotne dla bezpieczeństwa wygenerowanych haseł.
- **"fmt"**: Standardowa biblioteka Go do formatowania tekstu i interakcji z konsolą. Używana do wyświetlania komunikatów i pobierania danych wejściowych.
- **"math/big"**: Biblioteka do obsługi dużych liczb całkowitych, używana w połączeniu z "crypto/rand" do generowania losowych indeksów dla zestawu znaków.
- **"strconv"**: Używana do konwersji tekstu (string) na liczby całkowite (int), co jest niezbędne przy przetwarzaniu danych wejściowych od użytkownika.
- **"strings"**: Biblioteka do manipulacji ciągami znaków, używana do przycinania spacji z danych wejściowych oraz porównywania tekstów.

### 3. Definicja zestawów znaków

```go
const (
	lowerChars    = "abcdefghijklmnopqrstuvwxyz"
	upperChars    = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	digitChars    = "0123456789"
	specialChars  = "!@#$%^&*()-_+="
)
```

#### Wyjaśnienie:

- **lowerChars**: Zawiera wszystkie małe litery alfabetu angielskiego.
- **upperChars**: Zawiera wszystkie duże litery alfabetu angielskiego.
- **digitChars**: Zawiera cyfry od 0 do 9.
- **specialChars**: Zawiera popularne znaki specjalne używane w hasłach, takie jak `!`, `@`, `#`, itd.

Zestawy te będą używane do generowania haseł w zależności od wybranego poziomu złożoności.

### 4. Funkcja `GeneratePassword`

```go
func GeneratePassword(length int, complexity int) (string, error) {
	var charSet string

	switch complexity {
	case 1:
		charSet = lowerChars
	case 2:
		charSet = lowerChars + upperChars
	case 3:
		charSet = lowerChars + upperChars + digitChars
	case 4:
		charSet = lowerChars + upperChars + digitChars + specialChars
	default:
		return "", fmt.Errorf("invalid complexity level")
	}

	password := make([]byte, length)
	for i := range password {
		randomIndex, err := cryptoRandInt(len(charSet))
		if err != nil {
			return "", err
		}
		password[i] = charSet[randomIndex]
	}

	return string(password), nil
}
```

#### Wyjaśnienie:

1. **Wybór zestawu znaków**:
   - Funkcja przyjmuje dwa parametry: `length` (długość hasła) i `complexity` (poziom złożoności hasła).
   - W zależności od poziomu złożoności, wybieramy odpowiedni zestaw znaków:
     - **Poziom 1**: Tylko małe litery.
     - **Poziom 2**: Małe i duże litery.
     - **Poziom 3**: Małe, duże litery oraz cyfry.
     - **Poziom 4**: Małe, duże litery, cyfry oraz znaki specjalne.

2. **Generowanie hasła**:
   - Tworzymy pustą tablicę bajtów o długości określonej przez użytkownika.
   - Dla każdego znaku w haśle losujemy indeks z zestawu znaków (`charSet`) i przypisujemy odpowiedni znak do hasła.

3. **Zwracanie hasła**:
   - Zwracamy wygenerowane hasło jako string.

4. **Obsługa błędów**:
   - Jeśli poziom złożoności jest nieprawidłowy, funkcja zwróci błąd.

### 5. Funkcja generująca losowe liczby (`cryptoRandInt`)

```go
func cryptoRandInt(max int) (int, error) {
	nBig, err := rand.Int(rand.Reader, big.NewInt(int64(max)))
	if err != nil {
		return 0, err
	}
	return int(nBig.Int64()), nil
}
```

#### Wyjaśnienie:

1. **Generowanie liczby losowej**:
   - Funkcja `cryptoRandInt` przyjmuje maksymalną wartość (czyli długość zestawu znaków) i zwraca losowy indeks z przedziału [0, max).
   - Używamy tutaj funkcji `rand.Int` z pakietu `crypto/rand`, która generuje kryptograficznie bezpieczną losową liczbę całkowitą.

2. **Obsługa błędów**:
   - Jeśli wystąpi błąd podczas generowania losowej liczby, funkcja zwraca błąd, który jest obsługiwany w głównej funkcji generatora haseł.

### 6. Funkcja pobierająca dane wejściowe (`getInput`)

```go
func getInput(prompt string) string {
	var input string
	fmt.Print(prompt)
	fmt.Scanln(&input)
	return strings.TrimSpace(input)
}
```

#### Wyjaśnienie:

- **Pobieranie danych wejściowych**:
  - Funkcja `getInput` wyświetla komunikat (`prompt`), odczytuje dane wejściowe od użytkownika za pomocą `fmt.Scanln`, usuwa ewentualne spacje z początku i końca ciągu znaków przy użyciu `strings.TrimSpace`, a następnie zwraca wynik.

### 7. Funkcja główna (`main`)

```go
func main() {
	for {
		// Pobierz długość hasła od użytkownika
		lengthInput := getInput("Enter the password length (minimum 4 characters): ")
		length, err := strconv.Atoi(lengthInput)
		if err != nil || length < 4 {
			fmt.Println("Invalid input. Please enter a valid number greater than or equal to 4.")
			continue
		}

		// Pobierz poziom skomplikowania od użytkownika
		complexityInput := getInput("Enter the complexity level (1 = lowercase, 2 = lowercase+uppercase, 3 = lowercase+uppercase+digits, 4 = lowercase+uppercase+digits+special characters): ")
		complexity, err := strconv.Atoi(complexityInput)
		if err != nil || complexity < 1 || complexity > 4 {
			fmt.Println("Invalid input. Please enter a number between 1 and 4.")
			continue
		}

		// Wygeneruj hasło
		password, err := GeneratePassword(length, complexity)
		if err != nil {
			fmt.Println("Error generating password:", err)
			return
		}

		fmt.Println("Generated password:", password)

		// Zapytaj, czy użytkownik chce wygenerować kolejne hasło
		again := getInput("Do you want to generate another password? (y/n): ")
		if strings.ToLower(again) != "y" {
			break
		}
	}
}
```

#### Wyjaśnienie:

1. **Pętla główna**:
   - Cały proces generowania hasła jest zawarty w pętli `for`, dzięki czemu użytkownik może generować kolejne hasła bez ponownego uruchamiania programu.

2. **Pobieranie długości hasła**:
   - Program pyta użytkownika o długość hasła, a następnie konwertuje dane wejściowe z tekstu na liczbę całkowitą za pomocą `strconv.Atoi`.
   - Jeśli użytkownik wprowadzi nieprawidłową wartość (np. tekst zamiast liczby lub liczbę mniejszą niż 4), program wyświetli komunikat o błędzie i poprosi o ponowne wprowadzenie danych.

3. **Pobieranie poziomu skomplikowania**:
   - Program pyta użytkownika o poziom skomplikowania hasła. Użytkownik musi wprowadzić liczbę od 1 do 4, co określa, jakie znaki będą używane do generowania hasła.
   - Ponownie, jeśli użytkownik wprowadzi nieprawidłową wartość, program wyświetli komunikat o błędzie i poprosi o ponowne wprowadzenie danych.

4. **Generowanie i wyświetlanie hasła**:
   - Po uzyskaniu poprawnych danych wejściowych, program generuje hasło przy użyciu funkcji `GeneratePassword` i wyświetla je na konsoli.

5. **Pytanie o kolejne hasło**:
   - Po wygenerowaniu hasła program pyta, czy użytkownik chce wygenerować kolejne hasło. Jeśli użytkownik wprowadzi `y`, proces rozpoczyna się od nowa. W przeciwnym razie program kończy działanie.

