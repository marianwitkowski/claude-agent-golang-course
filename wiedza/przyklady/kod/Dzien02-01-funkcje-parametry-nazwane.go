package main

import (
	"fmt"
)

type CalculateParams struct {
	Width  float64
	Height float64
}

func CalculateArea(params CalculateParams) float64 {
	return params.Height * params.Width
}

func main() {
	params := CalculateParams{
		Width:  10.0,
		Height: 15.0,
	}
	area := CalculateArea(params)
	fmt.Println(area)
}
