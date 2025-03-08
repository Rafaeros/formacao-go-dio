// Challenge 4
/* A student (you) of the Go Language course wrote a code that represents the basic functions of a /).
After writing this code you should run tests to make sure that it returns the expected values. */

package Challenges

import "fmt"


func Sum(x int, y int,) int {
	fmt.Println(x + y)
	result := x + y
	return result
}

func Subtraction(x int, y int,) int {
	fmt.Println(x - y)
	result := x - y
	return result
}

func Multiplication(x int, y int,) int {
	fmt.Println(x * y)
	result := x * y
	return result
}

func Division(x int, y int,) int {
	fmt.Println(x / y)
	result := x / y
	return result
}


func Calculator(x int, y int,) {
	Sum(x, y)
	Subtraction(x, y)
	Multiplication(x, y)
	Division(x, y)
}
