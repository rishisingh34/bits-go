package main

import (
	"fmt"
)

func main() {
	lets_start := 5

	// 1. "if" don't use paranthese here in go. 
	if lets_start > 5 {
		fmt.Println("Let's start now!")
	} else if lets_start <= 5 && lets_start > 0  {
		fmt.Println("Let's start later!")
	} else {
		fmt.Println("Invalid option!")
	}

	// 2. initial statement of if block
	if length := getLength("🙂🙂🙂"); length < 3 {
		fmt.Println("Not possible")
	} else if length > 3 {
		fmt.Println("it is not returning the correct number of emojis.")
		fmt.Println("Expected: 3, Got:", length) // --> Expected: 3, Got: 12
	} else {
		fmt.Println("it is returning the correct number of emojis.")
		fmt.Println("Expected: 3, Got:", length) // --> Expected: 3, Got: 3
	}

	//3 . switch 
	switchInGo("1")
}

func getLength(s string) int {

	// return len(s)

	// conversion from string to []runes is important because it will len(s) will return the number of bytes not runes
	// and each emoji is represented by multiple bytes 
	s_to_runes := []rune(s)
	return len(s_to_runes)
	
}