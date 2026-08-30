### 1. Wstęp

Poniższy program w Go to kompletny przykład aplikacji do automatycznego przetwarzania obrazów, która monitoruje katalog, wykrywa nowe pliki graficzne, zmniejsza ich rozmiar o 50% i zapisuje wynik w innym katalogu z odpowiednią nazwą. Kluczowym elementem tego rozwiązania jest wykorzystanie biblioteki `fsnotify` do monitorowania zmian w systemie plików, co zapewnia wydajność i responsywność aplikacji.

### 2. Użyte technologie

- **Go (Golang)**: Wydajny, statycznie typowany język programowania, który doskonale nadaje się do tworzenia aplikacji serwerowych i narzędzi systemowych.
- **fsnotify**: Biblioteka Go, która umożliwia monitorowanie systemu plików pod kątem zmian, takich jak tworzenie, modyfikowanie, usuwanie lub przenoszenie plików. Jest to wydajniejsze i bardziej profesjonalne rozwiązanie, niż czasowe sprawdzanie (co kilka/kilkanaście sekund) czy w katalogu nie pojawił się nowy plik.
- **image/jpeg i image/png**: Standardowe biblioteki Go do pracy z formatami obrazów JPEG i PNG.
- **nfnt/resize**: Biblioteka Go do skalowania obrazów. Umożliwia zmianę rozmiaru obrazów przy użyciu różnych algorytmów.

### 3. Importowanie bibliotek

Na początku programu, importujemy potrzebne biblioteki:

```go
import (
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/fsnotify/fsnotify"
	"github.com/nfnt/resize"
)
```

#### Wyjaśnienie importowanych bibliotek:

- **"fmt"**: Standardowa biblioteka Go do formatowania wejścia i wyjścia. W tym przypadku nie jest bezpośrednio używana w programie, ale często jest pomocna do debugowania i wyświetlania informacji na konsoli.
- **"image", "image/jpeg", "image/png"**: Biblioteki Go, które umożliwiają manipulowanie obrazami oraz dekodowanie i kodowanie formatów JPEG i PNG.
- **"log"**: Biblioteka do logowania błędów i informacji. W naszym przypadku używana do rejestrowania błędów podczas przetwarzania obrazów i zdarzeń systemowych.
- **"os"**: Biblioteka do pracy z funkcjami systemu operacyjnego, takimi jak otwieranie plików, tworzenie katalogów itp.
- **"path/filepath"**: Biblioteka do pracy ze ścieżkami plików. Umożliwia manipulowanie ścieżkami w sposób niezależny od systemu operacyjnego.
- **"strings"**: Biblioteka do manipulowania ciągami znaków. W naszym przypadku używana do przetwarzania rozszerzeń plików.
- **"github.com/fsnotify/fsnotify"**: Zewnętrzna biblioteka, która umożliwia monitorowanie systemu plików pod kątem zmian.
- **"github.com/nfnt/resize"**: Zewnętrzna biblioteka do zmniejszania rozmiarów obrazów.

### 4. Funkcja main

```go
func main() {
	// Katalog do nasłuchiwania
	watchDir := "./watch"
	outDir := "./out"

	// Upewnij się, że katalog wyjściowy istnieje
	if _, err := os.Stat(outDir); os.IsNotExist(err) {
		os.Mkdir(outDir, 0755)
	}

	// Utwórz nowego obserwatora fsnotify
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	// Uruchom nasłuchiwanie w osobnym wątku
	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				// Obsługa zdarzenia stworzenia nowego pliku
				if event.Op&fsnotify.Create == fsnotify.Create {
					if isImage(event.Name) {
						processImage(event.Name, outDir)
					}
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Println("Error:", err)
			}
		}
	}()

	// Dodaj katalog do obserwacji
	err = watcher.Add(watchDir)
	if err != nil {
		log.Fatal(err)
	}

	// Zapobiega zakończeniu programu
	<-make(chan struct{})
}
```

#### Wyjaśnienie:

1. **watchDir i outDir**: 
   - `watchDir` to ścieżka katalogu, który będzie monitorowany. Każdy nowy plik dodany do tego katalogu zostanie przetworzony.
   - `outDir` to ścieżka katalogu, do którego będą zapisywane przetworzone pliki. Jeśli katalog nie istnieje, program go tworzy.

2. **Tworzenie katalogu wyjściowego**:
   - Program sprawdza, czy katalog `outDir` istnieje za pomocą funkcji `os.Stat`. Jeśli katalog nie istnieje, zostaje utworzony za pomocą `os.Mkdir`.

3. **fsnotify.NewWatcher**:
   - Tworzymy nowego "watchera" (obserwatora), który monitoruje zmiany w systemie plików. Jeśli nie uda się go utworzyć, program zakończy działanie z komunikatem o błędzie.

4. **Uruchamianie nasłuchiwania w osobnym goroutine**:
   - Tworzona jest nowa gorutyna (wątek), która zajmuje się przetwarzaniem zdarzeń związanych z systemem plików. W pętli `select` obsługiwane są dwa kanały:
     - `watcher.Events`: Kanał, który odbiera zdarzenia z systemu plików, takie jak tworzenie, modyfikowanie, usuwanie plików. W tym przypadku interesuje nas tylko zdarzenie `fsnotify.Create`, które oznacza, że nowy plik został utworzony.
     - `watcher.Errors`: Kanał, który odbiera błędy związane z monitorowaniem systemu plików. Jeśli wystąpi błąd, zostanie on zalogowany.

5. **Dodanie katalogu do obserwacji**:
   - Funkcja `watcher.Add` dodaje katalog `watchDir` do monitorowania. Jeśli operacja zakończy się błędem, program zakończy działanie.

6. **Zapobieganie zakończeniu programu**:
   - Aby program nie zakończył się po dodaniu katalogu do obserwacji, tworzymy pusty kanał i oczekujemy na dane, co zapobiega zakończeniu działania głównego wątku programu.

### 5. Funkcja isImage

```go
func isImage(filename string) bool {
	ext := strings.ToLower(filepath.Ext(filename))
	// Obsługiwane rozszerzenia plików graficznych
	return ext == ".jpg" || ext == ".jpeg" || ext == ".png"
}
```

#### Wyjaśnienie:

1. **Sprawdzanie rozszerzenia pliku**:
   - Funkcja `isImage` sprawdza, czy plik jest obrazem na podstawie jego rozszerzenia. Najpierw za pomocą `filepath.Ext(filename)` pobieramy rozszerzenie pliku, a następnie zmieniamy je na małe litery przy pomocy `strings.ToLower`.

2. **Obsługiwane formaty**:
   - Funkcja zwraca `true` tylko wtedy, gdy plik ma rozszerzenie `.jpg`, `.jpeg` lub `.png`. W przeciwnym razie zwraca `false`.

### 6. Funkcja processImage

```go
func processImage(filePath, outDir string) {
	// Otwórz plik obrazka
	file, err := os.Open(filePath)
	if err != nil {
		log.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Dekoduj obraz z pliku
	img, format, err := image.Decode(file)
	if err != nil {
		log.Println("Error decoding image:", err)
		return
	}

	// Wylicz nową szerokość i wysokość obrazka (50% oryginalnej)
	newWidth := img.Bounds().Dx() / 2
	newHeight := img.Bounds().Dy() / 2

	// Zmniejsz obrazek używając funkcji Resize z biblioteki resize
	thumbnail := resize.Resize(uint(newWidth), uint(newHeight), img, resize.Lanczos3)

	// Wygeneruj ścieżkę pliku wyjściowego z odpowiednim sufiksem
	outPath := generateOutPath(filePath, outDir, format)

	// Utwórz nowy plik w katalogu wyjściowym
	outFile, err := os.Create(outPath)
	if err != nil {
		log.Println("Error creating output file:", err)
		return
	}
	defer outFile.Close()

	// Zapisz obrazek do pliku w odpowiednim formacie
	switch format {
	case "jpeg":
		// Zapisz w formacie JPEG
		err = jpeg.Encode(outFile, thumbnail, nil)
	case "png":
		// Zapisz w formacie PNG
		err = png.Encode(outFile, thumbnail)
	default:
		// Format nieobsługiwany
		log.Println("Unsupported format:", format)
		return
	}

	if err != nil {
		log.Println("Error encoding image:", err)
	}
}
```

#### Wyjaśnienie:

1. **Otwieranie pliku**:
   - Funkcja `processImage` rozpoczyna od otwarcia pliku obrazu przy użyciu `os.Open(filePath)`. Jeśli operacja zakończy się niepowodzeniem, błąd zostanie zalogowany, a funkcja zakończy działanie.

2. **Dekodowanie obrazu**:
   - Po otwarciu pliku, obraz jest dekodowany przy użyciu funkcji `image.Decode`. Funkcja ta rozpoznaje format pliku (JPEG lub PNG) i dekoduje obraz do formatu zrozumiałego dla Go.

3. **Skalowanie obrazu**:
   - Nowe wymiary obrazu są obliczane jako 50% oryginalnych wymiarów (`img.Bounds().Dx() / 2` i `img.Bounds().Dy() / 2`). Obraz jest następnie skalowany przy użyciu funkcji `resize.Resize` z algorytmem `Lanczos3`, który jest znany z wysokiej jakości wyników przy zmniejszaniu rozmiarów obrazów.

4. **Generowanie ścieżki wyjściowej**:
   - Funkcja `generateOutPath` tworzy ścieżkę dla nowego pliku z dodanym sufiksem `-thumbnail` przed rozszerzeniem.

5. **Tworzenie pliku wyjściowego**:
   - Nowy plik zostaje utworzony w katalogu `outDir`. Jeśli operacja zakończy się błędem, błąd zostanie zalogowany, a funkcja zakończy działanie.

6. **Zapis obrazu w odpowiednim formacie**:
   - Zależnie od formatu obrazu (`jpeg` lub `png`), obraz zostaje zapisany do pliku wyjściowego przy użyciu odpowiednich funkcji kodujących (`jpeg.Encode` lub `png.Encode`).

### 7. Funkcja generateOutPath

```go
func generateOutPath(inputPath, outDir, format string) string {
	// Pobierz nazwę pliku z ścieżki
	fileName := filepath.Base(inputPath)
	// Rozdziel nazwę pliku na nazwę i rozszerzenie
	ext := filepath.Ext(fileName)
	nameWithoutExt := strings.TrimSuffix(fileName, ext)
	// Dodaj sufiks "-thumbnail" przed rozszerzeniem
	newFileName := nameWithoutExt + "-thumbnail" + ext
	// Połącz ścieżkę katalogu wyjściowego z nową nazwą pliku
	return filepath.Join(outDir, newFileName)
}
```

#### Wyjaśnienie:

1. **Pobieranie nazwy pliku**:
   - Funkcja `generateOutPath` rozpoczyna od pobrania nazwy pliku z pełnej ścieżki przy użyciu `filepath.Base(inputPath)`.

2. **Rozdzielanie nazwy pliku na nazwę i rozszerzenie**:
   - Funkcja `filepath.Ext(fileName)` pobiera rozszerzenie pliku, np. `.jpg`. Następnie `strings.TrimSuffix(fileName, ext)` usuwa to rozszerzenie z nazwy pliku, pozostawiając tylko nazwę bez rozszerzenia.

3. **Dodanie sufiksu "-thumbnail"**:
   - Do nazwy pliku zostaje dodany sufiks `-thumbnail`, po czym rozszerzenie pliku jest ponownie dołączane.

4. **Tworzenie pełnej ścieżki wyjściowej**:
   - Funkcja `filepath.Join(outDir, newFileName)` łączy katalog wyjściowy `outDir` z nową nazwą pliku, tworząc pełną ścieżkę do pliku wyjściowego.

