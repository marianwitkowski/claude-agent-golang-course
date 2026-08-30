package main

import "fmt"

func removeIndex(s []string, index int) []string {
	return append(s[:index], s[index+1:]...)
}

func main() {

	a := make([]string, 3, 5)
	a[0], a[1], a[2] = "A", "B", "C"
	fmt.Println(a, len(a), cap(a))

	var countries = []string{"Polska", "USA", "UK"}
	fmt.Println(countries, len(countries), cap(countries))

	a = append(a, "D", "E", "F")
	fmt.Println(a, len(a), cap(a))

	// a klonujemy do c
	c := make([]string, len(a))
	copy(c, a)
	fmt.Println(c)

	c[0] = "Z"
	fmt.Println(c)

	c = removeIndex(c, 2)
	fmt.Println(c)

	for index, value := range c {
		fmt.Println(index, "\t", value)
	}
}
