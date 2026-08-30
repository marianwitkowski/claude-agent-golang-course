package main

import "fmt"

type Car struct{}

func (c Car) Drive() {
	fmt.Println("Samochód jedzie")
}

func (c Car) Break() {
	fmt.Println("Samochód zatrzymuje się")
}

type Bicycle struct{}

func (b Bicycle) Drive() {
	fmt.Println("Rower jedzie")
}

func (b Bicycle) Break() {
	fmt.Println("Rower zatrzymuje się")
}

func main() {
	c := Car{}
	b := Bicycle{}

	c.Drive()
	c.Break()

	b.Drive()
	b.Break()
}
