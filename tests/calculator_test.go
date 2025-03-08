// Challenge 4
/* A student (you) of the Go Language course wrote a code that represents the basic functions of a /).
After writing this code you should run tests to make sure that it returns the expected values. */

package Challenges

import (
	"formacao-go-dio/Challenges"
	"testing"
)

func ShouldSumCorrect(t *testing.T) {
	test := Challenges.Sum(1, 2)
	result := 3
	if test != result {
		t.Errorf("Expected %d, got %d", result, test)
	}
}

func ShouldSumIncorrect(t *testing.T) {
	test := Challenges.Sum(1, 2)
	result := 4
	if test != result {
		t.Errorf("Expected %d, got %d", result, test)
	}
}

func ShouldSubtractionCorrect(t *testing.T) {
	test := Challenges.Subtraction(1, 2)
	result := -1
	if test != result {
		t.Errorf("Expected %d, got %d", result, test)
	}
}

func ShouldSubtractionIncorrect(t *testing.T) {
	test := Challenges.Subtraction(1, 2)
	result := -2
	if test != result {
		t.Errorf("Expected %d, got %d", result, test)
	}
}

func ShouldMultiplicationCorrect(t *testing.T) {
	test := Challenges.Multiplication(1, 2)
	result := 2
	if test != result {
		t.Errorf("Expected %d, got %d", result, test)
	}
}

func ShouldMultiplicationIncorrect(t *testing.T) {
	test := Challenges.Multiplication(1, 2)
	result := 3
	if test != result {
		t.Errorf("Expected %d, got %d", result, test)
	}
}

func ShouldDivisionCorrect(t *testing.T) {
	test := Challenges.Division(1, 2)
	result := 0
	if test != result {
		t.Errorf("Expected %d, got %d", result, test)
	}
}

func ShouldDivisionIncorrect(t *testing.T) {
	test := Challenges.Division(1, 2)
	result := 1
	if test != result {
		t.Errorf("Expected %d, got %d", result, test)
	}
}
