Poniżej znajduje się przykładowy kod w języku Go (Golang), który generuje kody kreskowe. Użyjemy do tego popularnej biblioteki github.com/boombuler/barcode, która obsługuje różne typy kodów kreskowych.


```go
package main

import (
	"flag"
	"fmt"
	"image/png"
	"os"

	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/code128"
	"github.com/boombuler/barcode/code39"
	"github.com/boombuler/barcode/qr"
)

func main() {
	// Definiowanie flag wiersza poleceń
	data := flag.String("data", "1234567890", "Dane do zakodowania")
	barcodeType := flag.String("type", "code128", "Typ kodu kreskowego (code128, code39, qr)")
	outputFile := flag.String("output", "barcode.png", "Nazwa pliku wyjściowego")

	flag.Parse()

	var barcode barcode.Barcode
	var err error

	// Wybór odpowiedniego typu kodu kreskowego na podstawie argumentu
	switch *barcodeType {
	case "code128":
		barcode, err = code128.Encode(*data)
	case "code39":
		barcode, err = code39.Encode(*data, false, false)
	case "qr":
		barcode, err = qr.Encode(*data, qr.M, qr.Auto)
	default:
		fmt.Println("Nieznany typ kodu kreskowego. Dostępne typy: code128, code39, qr")
		return
	}

	if err != nil {
		panic(err)
	}

	// Skalowanie kodu kreskowego do odpowiednich rozmiarów
	barcode, err = barcode.Scale(300, 100)
	if err != nil {
		panic(err)
	}

	// Tworzenie pliku, do którego zapisany zostanie kod kreskowy
	file, err := os.Create(*outputFile)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Zapisanie kodu kreskowego jako PNG
	err = png.Encode(file, barcode)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Kod kreskowy został wygenerowany i zapisany jako %s\n", *outputFile)
}
```

### Jak działa zmodyfikowany kod:

1. **Argumenty wiersza poleceń**:
   - `-data`: Określa dane, które mają być zakodowane w kodzie kreskowym (domyślnie `"1234567890"`).
   - `-type`: Określa typ kodu kreskowego, który ma być wygenerowany (domyślnie `"code128"`). Dostępne opcje to `"code128"`, `"code39"` oraz `"qr"`.
   - `-output`: Nazwa pliku wyjściowego, w którym zostanie zapisany kod kreskowy (domyślnie `"barcode.png"`).

2. **Generowanie kodu kreskowego**:
   - Na podstawie wybranego typu kodu kreskowego, koduje dane za pomocą odpowiedniej funkcji (np. `code128.Encode`, `code39.Encode`, `qr.Encode`).

3. **Skalowanie**:
   - Kod kreskowy jest skalowany do rozmiaru 300x100 pikseli. Dla kodu QR możesz rozważyć użycie innych rozmiarów, ponieważ QR code zazwyczaj jest kwadratowy.

4. **Zapis**:
   - Wynikowy obraz jest zapisywany do pliku PNG o nazwie podanej w argumentach.

### Uruchomienie programu

Aby uruchomić program, użyj wiersza poleceń, na przykład:

```bash
go run main.go -data="HelloWorld" -type="qr" -output="qr_code.png"
```

Ten przykład wygeneruje kod QR dla tekstu `"HelloWorld"` i zapisze go jako `qr_code.png`. Możesz zmienić typ kodu kreskowego na `code128` lub `code39` według potrzeb.