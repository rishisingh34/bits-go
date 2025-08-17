package main

import (
	"fmt"
)

func formatting() {
	// fmt.Printf() - prints a foramtted string to std out. 
	pi := 3.14
	fmt.Printf("Value of Pi: %.2f\n", pi)

	// fmt.Sprintf() - formats and returns a string
	formattedString := fmt.Sprintf("Value of Pi: %.2f", pi)
	fmt.Println(formattedString)


	// format verbs --> 
	// %f - float
	// %d - integer
	// %s - string
	// %t - boolean
	// %T - type
	// %p - pointer
	// %v - value
	// %w - error

	// %#x - prints the value in hexadecimal format with 0x prefix
	// %b - prints the value in binary format with 0b prefix
	// %x - prints the value in hexadecimal format without 0x prefix
	// %c - prints the value as a character
}