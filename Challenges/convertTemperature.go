// Challenge 1

package Challenges

import (
	"fmt"
)

func convertToCelsius(kelvin int64) int64{
	var celsius int64
	celsius = kelvin - 273
	return celsius
}

func CalcTemperature() {
	var boilingPoint int64 = 373
	c := convertToCelsius(boilingPoint)
	fmt.Printf("The boiling point of water in kelvin is: %d and the conversion to celsius is: %d", boilingPoint, c)
}
	