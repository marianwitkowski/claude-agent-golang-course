package main

import "fmt"

type Vehicle interface {
	Drive()
	Break()
}

func DriveVehicle(v Vehicle) {
	v.Drive()
	v.Break()
}

type Car struct{}

func (c Car) Drive() {
	fmt.Println("Samochod jedzie")
}

func (c Car) Break() {
	fmt.Println("Samochod zatrzymuje sie")
}

type Bicycle struct{}

func (b Bicycle) Drive() {
	fmt.Println("Rower jedzie")
}

func (b Bicycle) Break() {
	fmt.Println("Rower zatrzymuje sie")
}

func main() {
	c := Car{}
	b := Bicycle{}

	DriveVehicle(c)
	DriveVehicle(b)

}
