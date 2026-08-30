package main

import "fmt"

func main() {
	var intArray [3]int
	intArray[0] = 1
	intArray[1] = 10
	intArray[2] = 100

	a := [5]int{10, 20, 30, 40, 50}
	var b [5]int = [5]int{10, 20, 30, 40, 50}
	c := [...]int{10, 20, 30, 40, 50}
	d := [10]int{0: 1, 2: 1} //zmiana dla indeksu 0 i 2 na wartosc 1
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)

	for i := 0; i < len(c); i++ {
		fmt.Print(c[i], "|")
	}
	fmt.Println()

	for index, value := range c {
		fmt.Println(index, " -> ", value)
	}

	copyA := a
	fmt.Println(copyA)
	copyA[0] = -10
	fmt.Println(a, "|", copyA)

	copyAPointer := &a //przekazanie referencji do tablicy a
	copyAPointer[0] = -100
	fmt.Println(a, "|", copyAPointer)

	// wybieranie wielu elementow tablicy
	// 10, 20, 30, 40, 50
	fmt.Println("2 pierwsze elem. B:", b[:2])
	fmt.Println("2-gi, 3 i 4-ty element", b[1:4])
	fmt.Println("Od wart. 30 do konca", b[2:])
	fmt.Println("2 ostatnie elementy", b[len(b)-2:])

	s := "Ala ma kota"
	fmt.Println(s[2:8])
	fmt.Println(s[0:1])

}
