package main

import "fmt"

func main() {
	// For loop
	for i := 0; i < 5; i++ {
		fmt.Println("For Loop Iteration:", i)
	}

	// While loop (using for)
	j := 0
	for j < 5 {
		fmt.Println("While Loop Iteration:", j)
		j++
	}

	// Infinite loop
	// Uncommenting the following lines will cause an infinite loop
	/*
		for {
			fmt.Println("Infinite Loop")
		}
	*/
}
