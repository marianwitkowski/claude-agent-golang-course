package main

import (
	"fmt"
	"reflect"
)

func dummyFunc() {
	fmt.Println("XYZ")
}

// deklaracja funkcji z parametrami i zwracającej wartośc
func add(x int, y int) int {
	return x + y
}

func cube(x int) (result int) {
	result = x * x * x
	return
}

func rect(a float32, b float32) (float32, float32) {
	area := a * b
	cir := 2*a + 2*b
	return area, cir
}

// deklaracja funkcji z parametrami przekazywanymi przez referencję
func updateByPointer(a *int, s *string) {
	*a += 20
	*s += "...chyba Ty"
}

// funkcja o zmiennej l. parametrów jednego typu
func anyParams(s ...string) {
	fmt.Println(s)
}

func anyParams2(s string, numbers ...int) {
	fmt.Println("========================")
	fmt.Println(s)
	fmt.Println(numbers)
}

func anyParam3(params ...interface{}) {
	fmt.Println("========================")
	for ix, v := range params {
		fmt.Println(ix, v, reflect.TypeOf(v))
	}
}

// domknięcia w Golang
func intGenerator(start int, step int) func() int {
	counter := start
	return func() int {
		//counter++
		counter += step
		return counter
	}

}

var (
	area = func(a int, b int) int {
		return a * b
	}
)

func main() {
	sum := add(20, 30)
	fmt.Println(sum)
	fmt.Println(cube(5))
	a, c := rect(2, 3)
	fmt.Printf("Pole=%v, obwod=%v\n", a, c)

	var x = 20
	var s = "Jestem głodny"
	fmt.Println("Przed:", x, s)
	updateByPointer(&x, &s)
	fmt.Println("Po:", x, s)

	fmt.Println("Wywołanie f. anonimowej ", area(4, 5))

	anyParams("Hello")
	anyParams("Hello", "world", "!")

	anyParams2("ABC")
	anyParams2("ABC", 1, 3, 5, 7, 11)

	anyParam3(1, "Hello world!", true, 12.34, []string{"A", "B", "C"})

	fmt.Println("========================")
	genCounter := intGenerator(10, 5)
	fmt.Println(genCounter())
	fmt.Println(genCounter())
	fmt.Println(genCounter())

}
