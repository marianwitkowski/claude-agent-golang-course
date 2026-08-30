package main

import (
	"fmt"
	"reflect"
)

var (
	x = 10
	y = -10
	z = false
)

func main() {
	fmt.Println("Hello world!")

	// deklarowanie zmiennych
	var i int
	var s string
	i = 10
	s = "ABCDEFG"
	fmt.Println(i, s)

	// deklarowanie zmiennych bez podania typu
	var j = 123456
	var t = "XYZABC"
	fmt.Println(reflect.TypeOf(j), reflect.TypeOf(t))

	// deklarowanie zmiennych "wiele na raz"
	var fname, lname string = "Jan", "Kowalski"
	a, b, c := 10, 20, 30
	d := 40
	fmt.Println(fname + " " + lname)
	fmt.Println(a + b + c + d)

	// deklarowanie stalych
	const PI float64 = 3.14159
	fmt.Println(PI)

}
