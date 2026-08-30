Tworzenie aplikacji sieciowej w Go jest stosunkowo proste dzięki wbudowanym bibliotekom, takim jak `net`, które oferują wsparcie dla tworzenia serwerów i klientów TCP oraz UDP. Poniżej przedstawię, jak stworzyć prosty, tekstowy czat peer-to-peer (P2P) w Go. W tym przykładzie każdy węzeł może działać zarówno jako serwer, jak i klient, co umożliwia bezpośrednią komunikację między uczestnikami.

### Krok 1: Tworzenie Serwera TCP

Każdy uczestnik czatu będzie nasłuchiwał na porcie TCP i odbierał wiadomości od innych użytkowników.

```go
package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func handleConnection(conn net.Conn) {
	defer conn.Close()
	for {
		message, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Print("Message received: ", string(message))
	}
}

func startServer(port string) {
	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		fmt.Println("Error starting server:", err)
		return
	}
	defer listener.Close()

	fmt.Println("Server started on port", port)
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}
		go handleConnection(conn)
	}
}
```

### Krok 2: Tworzenie Klienta TCP

Każdy węzeł będzie mógł wysyłać wiadomości do innych węzłów poprzez połączenie TCP.

```go
func connectToPeer(address string) net.Conn {
	conn, err := net.Dial("tcp", address)
	if err != nil {
		fmt.Println("Error connecting to peer:", err)
		return nil
	}
	return conn
}

func sendMessage(conn net.Conn, message string) {
	_, err := fmt.Fprintf(conn, message+"\n")
	if err != nil {
		fmt.Println("Error sending message:", err)
	}
}
```

### Krok 3: Uruchamianie Węzła P2P

Główna funkcja pozwoli użytkownikowi na uruchomienie serwera i podłączenie się do innych węzłów.

```go
func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter port to listen on: ")
	port, _ := reader.ReadString('\n')
	port = strings.TrimSpace(port)

	go startServer(port)

	for {
		fmt.Print("Enter peer address (ip:port) or 'exit' to quit: ")
		address, _ := reader.ReadString('\n')
		address = strings.TrimSpace(address)

		if address == "exit" {
			break
		}

		conn := connectToPeer(address)
		if conn == nil {
			continue
		}

		for {
			fmt.Print("Enter message: ")
			message, _ := reader.ReadString('\n')
			message = strings.TrimSpace(message)

			if message == "exit" {
				break
			}

			sendMessage(conn, message)
		}

		conn.Close()
	}
}
```

### Krok 4: Uruchomienie Aplikacji

- **Serwer:** Uruchom aplikację na pierwszym węźle, wpisując port, na którym będzie nasłuchiwać, np. `3000`.
- **Klient:** Na drugim węźle uruchom aplikację i podaj adres IP oraz port serwera, np. `127.0.0.1:3000`.

### Wyjaśnienie Kodów

1. **Serwer TCP:** Każdy węzeł uruchamia serwer TCP, który nasłuchuje na określonym porcie. Gdy nadejdzie połączenie, serwer uruchamia osobną gorutynę, aby obsłużyć to połączenie, co pozwala na równoległe przetwarzanie wielu połączeń.
   
2. **Klient TCP:** Węzeł może połączyć się z innym węzłem przez adres IP i port. Po połączeniu może wysyłać wiadomości do tego węzła, które są przesyłane poprzez połączenie TCP.

3. **Komunikacja:** Użytkownicy mogą wysyłać wiadomości do podłączonych węzłów, a także odbierać wiadomości od innych. Komunikacja odbywa się przez TCP, co zapewnia niezawodność przesyłu danych.

### Podsumowanie

Ten prosty czat P2P pokazuje podstawy programowania sieciowego w Go. Używając `net.Listen` i `net.Dial`, można łatwo zbudować aplikacje komunikujące się przez sieć. Możliwości Go w zakresie współbieżności, dzięki gorutynom, sprawiają, że obsługa wielu połączeń jest efektywna i prosta. W dalszym rozwoju można dodać funkcje takie jak obsługa wielu jednoczesnych połączeń, szyfrowanie komunikacji czy zaawansowane zarządzanie siecią.