package main

import "fmt"


// Write a closure that counts how many times it’s called.
// -------- EX1 --------
func newCounter() func() int {
	cnt := 0
	return func() int {
		cnt++ 
		return cnt 
	}
}


func ex1() {
	counter := newCounter()
	fmt.Println(counter())
	fmt.Println(counter())
	fmt.Println((counter()))
}	
// -------- EX1 --------


// Write a closure that remembers a person’s name and always greets them.
// -------- EX2 --------

func greeter(name string) func () {
	return func () {
		fmt.Printf("Hello, %s\n", name)
	}
}

func ex2() {
	greetRishi := greeter("Rishi")
	greetRishi()
	greetRishi()
	greetRishi()
}

// -------- EX2 --------


// Write a closure that multiplies numbers by a fixed factor.
// -------- EX3 --------

func multiplier(num int) func(int)  {
	return func(num2 int) {
		fmt.Println(num2*num)
	}
} 

func ex3() {
	double := multiplier(2)
	triple := multiplier(3)

	double(10) // 20 
	triple(10) // 30  
}

// -------- EX3 --------