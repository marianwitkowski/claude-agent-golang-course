package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {
	x := "100"
	// konwersja str to int
	x_int, err := strconv.Atoi(x)
	fmt.Println(x_int, err, reflect.TypeOf(x_int))

	y := "3.14159"
	y_float, err := strconv.ParseFloat(y, 64)
	fmt.Println(y_float, err, reflect.TypeOf(y_float))

	var i int = 123456
	fmt.Println(i, reflect.TypeOf(i))

	var j int64
	j = int64(i)
	fmt.Println(j, reflect.TypeOf(j))

	var k int8
	k = int8(i)
	fmt.Println(k, reflect.TypeOf(k))

	var a float32 = 123.456
	var b float64
	b = float64(a)
	fmt.Println(b, reflect.TypeOf(b))

	fmt.Printf("wartosc a=%v, b=%v, type zmiennej b to %v\n", a, b, reflect.TypeOf(b))

	num1 := 123
	strNum1 := fmt.Sprintf("|%d|", num1)
	fmt.Println(strNum1)

	strNum1 = fmt.Sprintf("|%10d|", num1)
	fmt.Println(strNum1)

	strNum1 = fmt.Sprintf("|%-10d|", num1)
	fmt.Println(strNum1)

	strNum1 = fmt.Sprintf("|%010d|", num1)
	fmt.Println(strNum1)

	strNum1 = fmt.Sprintf("|%b|", num1)
	fmt.Println(strNum1)

}
