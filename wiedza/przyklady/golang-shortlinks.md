### 1. Wstęp

Skracacze URL są szeroko używane do tworzenia krótszych, bardziej zwięzłych wersji długich adresów URL, które są łatwiejsze do udostępniania i zarządzania. Skrócone URL-e działają poprzez przekierowanie użytkownika na oryginalny adres URL po odwiedzeniu skróconego linku. W tej sekcji stworzymy prosty skracacz URL jako usługę REST API w Go.

### 2. Użyte technologie

- **Go (Golang)**: Język programowania używany do implementacji logiki backendowej.
- **net/http**: Standardowa biblioteka Go do obsługi serwera HTTP, która posłuży do implementacji REST API.
- **gorilla/mux**: Popularny router w Go, używany do obsługi bardziej zaawansowanego routingu w aplikacjach webowych.

### 3. Importowanie bibliotek

Na początku programu importujemy potrzebne biblioteki:

```go
import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"sync"
	"time"

	"github.com/gorilla/mux"
)
```

#### Wyjaśnienie importowanych bibliotek:

- **"encoding/json"**: Używana do kodowania i dekodowania danych JSON.
- **"fmt"**: Używana do formatowania ciągów znaków.
- **"math/rand"**: Używana do generowania losowych wartości, co pomoże nam w tworzeniu unikalnych skrótów.
- **"net/http"**: Standardowa biblioteka do obsługi serwera HTTP i zapytań HTTP.
- **"sync"**: Używana do zarządzania współbieżnością (np. ochrona dostępu do mapy przechowującej URL-e).
- **"github.com/gorilla/mux"**: Zewnętrzna biblioteka, która ułatwia routing w aplikacjach webowych Go. Możesz zainstalować ją poleceniem:
  ```bash
  go get -u github.com/gorilla/mux
  ```

### 4. Struktura danych

Najpierw zdefiniujemy strukturę, która będzie przechowywać skrócone URL-e i odpowiadające im oryginalne adresy.

```go
type URLStore struct {
	sync.RWMutex
	urlMap map[string]string
}
```

#### Wyjaśnienie:

- **URLStore**: Jest to struktura danych, która będzie przechowywać mapę związków między skróconymi URL-ami a oryginalnymi adresami URL. Używamy tu mapy (`urlMap map[string]string`), gdzie klucz to skrócony URL, a wartość to oryginalny URL.
- **sync.RWMutex**: Używamy tego mechanizmu, aby zapewnić bezpieczny dostęp do mapy w środowisku wielowątkowym. `RWMutex` pozwala na równoczesne odczyty, ale blokuje dostęp na czas zapisu.

### 5. Tworzenie serwera HTTP

Teraz utworzymy serwer HTTP i skonfigurujemy podstawowe ścieżki oraz funkcje obsługujące zapytania.

```go
func main() {
	store := &URLStore{
		urlMap: make(map[string]string),
	}

	r := mux.NewRouter()

	// Obsługa tworzenia skrótu URL
	r.HandleFunc("/shorten", store.ShortenURL).Methods("POST")

	// Obsługa przekierowania ze skróconego URL na oryginalny
	r.HandleFunc("/{shortURL}", store.Redirect).Methods("GET")

	fmt.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", r)
}
```

#### Wyjaśnienie:

1. **Inicjalizacja URLStore**:
   - Tworzymy nową instancję `URLStore`, która będzie przechowywać mapę skróconych URL-i.

2. **Konfiguracja routera**:
   - Używamy `mux.NewRouter()` do utworzenia nowego routera, który będzie zarządzać routingiem naszych zapytań HTTP.

3. **Ścieżka `/shorten`**:
   - Ta ścieżka obsługuje żądania POST, które przekazują URL do skrócenia. Obsługiwana przez funkcję `ShortenURL`.

4. **Ścieżka `/{shortURL}`**:
   - Ta ścieżka obsługuje żądania GET, które przekierowują ze skróconego URL na oryginalny adres URL. Obsługiwana przez funkcję `Redirect`.

5. **Uruchomienie serwera**:
   - Serwer HTTP zostaje uruchomiony na porcie 8080 przy użyciu `http.ListenAndServe(":8080", r)`.

### 6. Obsługa skracania URL-i

Teraz zaimplementujemy funkcję `ShortenURL`, która będzie odpowiadać za generowanie skróconych URL-i.

```go
func (s *URLStore) ShortenURL(w http.ResponseWriter, r *http.Request) {
	var request struct {
		URL string `json:"url"`
	}

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil || request.URL == "" {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	shortURL := s.generateShortURL()
	s.Lock()
	s.urlMap[shortURL] = request.URL
	s.Unlock()

	response := map[string]string{
		"short_url": fmt.Sprintf("http://localhost:8080/%s", shortURL),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
```

#### Wyjaśnienie:

1. **Dekodowanie żądania**:
   - Oczekujemy, że klient wyśle w żądaniu JSON zawierający klucz `"url"` z wartością URL-a do skrócenia. Jeśli format żądania jest niepoprawny lub URL jest pusty, zwracamy błąd 400 (`http.StatusBadRequest`).

2. **Generowanie skróconego URL-a**:
   - `generateShortURL()` to funkcja, która generuje unikalny skrócony URL.

3. **Przechowywanie skróconego URL-a**:
   - Używamy mechanizmu blokady (`s.Lock()`), aby zapewnić, że dostęp do mapy jest bezpieczny w środowisku wielowątkowym, a następnie przechowujemy skrócony URL oraz oryginalny adres URL w mapie.

4. **Zwracanie odpowiedzi**:
   - Tworzymy odpowiedź JSON zawierającą skrócony URL w formacie `http://localhost:8080/{shortURL}` i wysyłamy ją do klienta.

### 7. Funkcja generująca skrócone URL-e

Teraz zaimplementujemy funkcję `generateShortURL`, która będzie generować losowe, unikalne skrócone URL-e.

```go
func (s *URLStore) generateShortURL() string {
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	rand.Seed(time.Now().UnixNano())
	b := make([]byte, 6)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
```

#### Wyjaśnienie:

1. **Zbiór znaków**:
   - Używamy zbioru znaków `letters`, który zawiera małe i wielkie litery oraz cyfry. Na jego podstawie będziemy generować losowe ciągi znaków.

2. **Generowanie losowego ciągu**:
   - Funkcja `rand.Intn(len(letters))` losuje indeks z przedziału długości zbioru `letters`, a następnie przypisuje odpowiedni znak do tablicy `b`.

3. **Długość skróconego URL-a**:
   - Skrócony URL będzie miał długość 6 znaków. Można to łatwo zmienić, modyfikując rozmiar tablicy `b`.

4. **Zwracanie skróconego URL-a**:
   - Po wygenerowaniu skróconego URL-a funkcja zwraca go jako string.

### 8. Obsługa przekierowania

Teraz zaimplementujemy funkcję `Redirect`, która będzie obsługiwać przekierowanie użytkowników ze skróconego URL-a na oryginalny adres URL.

```go
func (s *URLStore) Redirect(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	shortURL := vars["shortURL"]

	s.RLock()
	originalURL, exists := s.urlMap[shortURL]
	s.RUnlock()

	if !exists {
		http.NotFound(w, r)
		return
	}

	http.Redirect(w, r, originalURL, http.StatusFound)


}
```

#### Wyjaśnienie:

1. **Pobieranie skróconego URL-a**:
   - Używamy `mux.Vars(r)`, aby pobrać wartość skróconego URL-a z parametrów ścieżki.

2. **Wyszukiwanie oryginalnego URL-a**:
   - Blokujemy dostęp do mapy na czas odczytu (`s.RLock()`), aby bezpiecznie pobrać oryginalny URL odpowiadający skróconemu URL-owi. Następnie zwalniamy blokadę (`s.RUnlock()`).

3. **Obsługa nieznalezionego URL-a**:
   - Jeśli skrócony URL nie istnieje w mapie, zwracamy błąd 404 (`http.NotFound(w, r)`).

4. **Przekierowanie**:
   - Jeśli skrócony URL istnieje, przekierowujemy użytkownika na oryginalny URL przy użyciu `http.Redirect(w, r, originalURL, http.StatusFound)`.


