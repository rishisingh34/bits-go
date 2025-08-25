package main

import "fmt"

// clouse demonstrates the concept of closures in Go.
//
// --- What is a Closure? ---
// A closure is an *anonymous function* that "remembers" the variables
// from the scope in which it was created, even after that scope has ended.
//
// In simple words:
// ➝ A closure is a function + the surrounding state it carries along.
// ➝ This allows the function to "capture" and "use" variables defined outside it.
//
// Closures are extremely useful when you want functions that maintain state,
// create custom function generators, or encapsulate logic with its own environment.
func closure() {
	// Example 1: Closure capturing an external variable
	message := "Hello"

	// This function "remembers" the variable `message`
	printMessage := func() {
		fmt.Println("Message from closure:", message)
	}
	printMessage() // Output: Message from closure: Hello

	// Example 2: Closure that maintains its own state
	// We define a function that returns another function
	counter := func() func() int {
		count := 0
		return func() int {
			count++
			return count
		}
	}

	// Each call to `counter()` creates a new independent closure
	firstCounter := counter()
	secondCounter := counter()

	fmt.Println("First Counter:", firstCounter())   // 1
	fmt.Println("First Counter:", firstCounter())   // 2
	fmt.Println("Second Counter:", secondCounter()) // 1
	fmt.Println("First Counter:", firstCounter())   // 3

	// Example 3: Closure as a function generator
	// Let's make a multiplier generator
	multiplier := func(factor int) func(int) int {
		return func(n int) int {
			return n * factor
		}
	}

	double := multiplier(2)
	triple := multiplier(3)

	fmt.Println("Double 5:", double(5)) // 10
	fmt.Println("Triple 5:", triple(5)) // 15
}
