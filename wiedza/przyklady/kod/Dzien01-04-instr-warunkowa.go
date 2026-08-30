package main

import (
	"fmt"
	"strconv"
)

func main() {
	var age int = 20
	if age >= 18 {
		fmt.Println("jesteś pełnoletni")
	} else {
		fmt.Println("nie jesteś pełnoletni")
	}

	var x = 16
	if (x%5 == 0) && (x%3 == 0) {
		fmt.Println("podzielna przez 3 i przez 5 jednoczesnie")
	} else if x%5 == 0 {
		fmt.Println("podzielna przez 5 tylko")
	} else if x%3 == 0 {
		fmt.Println("podzielna przez 3 tylko")
	} else {
		fmt.Println("ani to, ani tamto")
	}

	if x = 30; x == 30 {
		fmt.Println("Liczba 30")
	}

	v := "123"
	if v_int, err := strconv.Atoi(v); err == nil {
		fmt.Printf("%T, %v\n", v_int, v_int)
	} else {
		fmt.Println("konwersja nie powiodła się")
	}

}
