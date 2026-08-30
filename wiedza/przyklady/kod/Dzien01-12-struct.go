package main

import (
	"encoding/json"
	"fmt"
)

type Rectangle struct {
	a     float32
	b     float32
	color string
}

type Employee struct {
	Firstname string `json:firstname`
	Lastname  string `json:lastname`
	Location  string `json:location`
}

func main() {
	var rect1 = new(Rectangle)
	rect1.a = 5
	rect1.b = 10
	rect1.color = "green"
	fmt.Printf("%+v\n", rect1)

	var rect2 = Rectangle{10, 5, "red"}
	fmt.Printf("%+v\n", rect2)

	var rect3 = Rectangle{a: 3, b: 10}
	fmt.Printf("%+v\n", rect3)
	rect3.color = "yellow"
	fmt.Printf("%+v\n", rect3)

	jsonString := `

	{
		"firstname" : "Jan",
		"lastname" : "Kowalski",
		"location" : "Warszawa"
	}

	`
	emp := new(Employee)
	json.Unmarshal([]byte(jsonString), emp)
	fmt.Printf("%+v\n", emp)

	emp.Firstname = "Zenek"
	emp.Lastname = "Martyniuk"
	str, _ := json.Marshal(emp)
	fmt.Printf("%s\n", str)

}
