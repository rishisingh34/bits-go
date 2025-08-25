package main

import "fmt"

// A normal function that adds two numbers
func add(a, b int) int {
	return a + b
}

// A curried version of add
func curryAdd(a int) func(int) int {
	return func(b int) int {
		return a + b
	}
}

func currying() {
	fmt.Println(add(2, 3))
	add2 := curryAdd(2)
	fmt.Println(add2(3))  // 5
	fmt.Println(add2(10)) // 12
}
