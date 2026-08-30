package main

import (
	"fmt"
	"time"
)

func main() {

	dayOfMonth := 20 //time.Now().Day()
	switch dayOfMonth {
	case 10:
		fmt.Println("zaplac ZUS")
		fallthrough
	case 20:
		fmt.Println("zaplac PIT")
		fmt.Println("Ide na koszykowke")
		fallthrough
	case 25:
		fmt.Println("zaplac VAT")
		fallthrough
	case 7, 14:
		fmt.Println("Ide na koszykowke")
		fallthrough
	default:
		fmt.Println("Dzis nie place")

	}

	fmt.Println("===================================")
	switch today := time.Now(); {
	case today.Day() < 10:
		fmt.Println("Przed zaplata ZUSu")
	case today.Day() == 10:
		fmt.Println("Zaplac ZUS")
	case today.Day() > 10:
		fmt.Println("Chyba zaplaciles ZUS???")
	}

}
