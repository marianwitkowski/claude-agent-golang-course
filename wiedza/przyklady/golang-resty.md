Oto przegląd możliwości biblioteki Resty w Go, z przykładami kodu demonstrującymi różne funkcje, takie jak wysyłanie żądań GET i POST, przekazywanie parametrów, wysyłanie JSON-a w body, obsługę błędów serwera, ustawianie timeoutu, nagłówków HTTP oraz upload plików.

### 1. **Wywoływanie metody GET i przekazywanie parametrów**

Resty umożliwia łatwe wysyłanie żądań GET i przekazywanie parametrów zapytania.

#### Przykład:
```go
import (
    "fmt"
    "github.com/go-resty/resty/v2"
)

func main() {
    client := resty.New()

    resp, err := client.R().
        SetQueryParam("param1", "value1").
        SetQueryParam("param2", "value2").
        Get("https://api.example.com/data")

    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    fmt.Println("Response Status:", resp.Status())
    fmt.Println("Response Body:", resp.String())
}
```

### 2. **Wywoływanie metody POST i przekazywanie parametrów**

Resty pozwala również na wysyłanie żądań POST z parametrami formularza lub JSON-em w body.

#### Przykład z parametrami formularza:
```go
import (
    "fmt"
    "github.com/go-resty/resty/v2"
)

func main() {
    client := resty.New()

    resp, err := client.R().
        SetFormData(map[string]string{
            "param1": "value1",
            "param2": "value2",
        }).
        Post("https://api.example.com/form-submit")

    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    fmt.Println("Response Status:", resp.Status())
    fmt.Println("Response Body:", resp.String())
}
```

### 3. **JSON w Body**

Resty automatycznie serializuje struktury Go do formatu JSON i ustawia odpowiednie nagłówki.

#### Przykład:
```go
import (
    "fmt"
    "github.com/go-resty/resty/v2"
)

func main() {
    client := resty.New()

    data := map[string]interface{}{
        "name":  "John Doe",
        "email": "john.doe@example.com",
    }

    resp, err := client.R().
        SetHeader("Content-Type", "application/json").
        SetBody(data).
        Post("https://api.example.com/users")

    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    fmt.Println("Response Status:", resp.Status())
    fmt.Println("Response Body:", resp.String())
}
```

### 4. **Obsługa błędów serwera**

Resty oferuje łatwe zarządzanie błędami, np. w przypadku odpowiedzi z kodem 4xx lub 5xx.

#### Przykład:
```go
import (
    "fmt"
    "github.com/go-resty/resty/v2"
)

func main() {
    client := resty.New()

    resp, err := client.R().
        Get("https://api.example.com/not-found")

    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    if resp.IsError() {
        fmt.Printf("Error: Status %d, Body: %s\n", resp.StatusCode(), resp.String())
    } else {
        fmt.Println("Response Status:", resp.Status())
        fmt.Println("Response Body:", resp.String())
    }
}
```

### 5. **Ustawianie timeoutu**

Resty umożliwia ustawienie timeoutu zarówno na poziomie klienta, jak i pojedynczego żądania.

#### Przykład:
```go
import (
    "github.com/go-resty/resty/v2"
    "time"
)

func main() {
    client := resty.New()
    client.SetTimeout(5 * time.Second) // Ustawienie timeoutu na poziomie klienta

    resp, err := client.R().
        SetTimeout(2 * time.Second). // Ustawienie timeoutu na poziomie żądania
        Get("https://api.example.com/timeout-test")

    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    fmt.Println("Response Status:", resp.Status())
    fmt.Println("Response Body:", resp.String())
}
```

### 6. **Ustawianie nagłówka HTTP**

Resty pozwala na ustawienie dowolnych nagłówków HTTP w żądaniu.

#### Przykład:
```go
import (
    "fmt"
    "github.com/go-resty/resty/v2"
)

func main() {
    client := resty.New()

    resp, err := client.R().
        SetHeader("Authorization", "Bearer token").
        SetHeader("Custom-Header", "value").
        Get("https://api.example.com/data")

    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    fmt.Println("Response Status:", resp.Status())
    fmt.Println("Response Body:", resp.String())
}
```

### 7. **Upload plików**

Resty obsługuje upload plików, zarówno pojedynczych, jak i wielu plików naraz.

#### Przykład uploadu pojedynczego pliku:
```go
import (
    "fmt"
    "github.com/go-resty/resty/v2"
)

func main() {
    client := resty.New()

    resp, err := client.R().
        SetFile("file_field_name", "/path/to/file.jpg").
        Post("https://api.example.com/upload")

    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    fmt.Println("Response Status:", resp.Status())
    fmt.Println("Response Body:", resp.String())
}
```

#### Przykład uploadu wielu plików:
```go
import (
    "fmt"
    "github.com/go-resty/resty/v2"
)

func main() {
    client := resty.New()

    resp, err := client.R().
        SetFiles(map[string]string{
            "file1": "/path/to/file1.jpg",
            "file2": "/path/to/file2.png",
        }).
        Post("https://api.example.com/multi-upload")

    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    fmt.Println("Response Status:", resp.Status())
    fmt.Println("Response Body:", resp.String())
}
```
