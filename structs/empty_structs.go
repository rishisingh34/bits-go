package main 

import "fmt"

func emptyStructs() {
	// anonymous empty struct
	a := struct{}{}
	fmt.Println(a)

	// Define an empty struct
	type EmptyStruct struct{}

	// Create an instance of the empty struct
	var e EmptyStruct

	// Print the empty struct
	fmt.Println(e)
}