package main

import (
	"fmt"
)

func complexSignature(f func(int, int) int, a int, b int) {
	// this function takes a function and two integers as parameters
	// the function in the parameter has itself two integer parameters and returns an integer
	result := f(a, b)
	fmt.Println("Result:", result)
}
