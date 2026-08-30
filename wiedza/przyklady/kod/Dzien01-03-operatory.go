package main

import "fmt"

func main() {
	x, y := 12, 10
	println(x + y)
	println(x - y)
	println(x * y)
	println(x / y)
	println(x % y)
	x++
	println(x)
	x--
	println(x)

	x = 100
	x += 2 //x + 2
	x -= 2 //x - 2
	x *= 2 //x * 2
	x /= 3 //x / 3
	x %= 3 //x % 3

	/*
		operatory komparacji
		== != < <= > >=
	*/
	z := x > y
	println(z)

	// operatory logiczne
	z = (x > y) && x < 10 //AND
	z = (x < y) || x > 10 //OR
	z = !(x >= y)         //NOT

	// operatory binarne
	var m uint = 0b1010 //10
	var n uint = 0b0110
	var p uint

	p = m & n
	fmt.Printf("AND: %b\n", p)
	p = m | n
	fmt.Printf("OR: %b\n", p)
	p = m ^ n
	fmt.Printf("XOR: %b\n", p)
	p = m << 1
	fmt.Printf("SHIFT-L: %b\n", p)
	p = m >> 1
	fmt.Printf("SHIFT-R: %b\n", p)

}
