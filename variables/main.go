package main

import (
	"fmt"
)

// primitive types --->  
// bool 
// string 
// int 
// float64 
// byte: 8 bits of data 

func main() {
	// 1. declaring variable 
	var str string = "A string" 
	fmt.Println(str)

	// 2. shorthand declaration 
	str2 := "another string" // type inference
	fmt.Println(str2)

	// 3. comments
	comments()

	// 4. type sizes
	type_sizes()

	// 5. same line declaration 
	name, email, phone_number := "Rishi", "singh34rishi@gmail.com", 9999999999
	fmt.Println("Name:", name, "Email:", email, "Phone Number:", phone_number)

	// 6. constants
	constants()

	// 7. formatting
	formatting()

	// 8. runes and strings
	runesAndStrings()
}