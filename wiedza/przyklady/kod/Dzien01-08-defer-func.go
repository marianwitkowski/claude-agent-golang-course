package main

import (
	"fmt"
	"os"
)

// zwraca deskryptor pliku
func createFile(fn string) *os.File {
	fmt.Println("otwieram plik....")
	fd, err := os.Create(fn)
	if err != nil {
		panic(err)
	}
	return fd
}

func closeFile(fd *os.File) {
	fmt.Println("zamykam plik...")
	err := fd.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "wystapil błąd: %v\n", err.Error())
		os.Exit(1)
	}
}

func writeFile(fd *os.File, txt string) {
	fmt.Println("zapisuję coś do pliku....")
	fmt.Fprint(fd, txt)
}

func main() {
	fd := createFile("test.txt")
	defer closeFile(fd)
	writeFile(fd, "Linia tekstu....")

}
