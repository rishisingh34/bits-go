package main

import "fmt"

// Define a struct
type Person struct {
	Name string
	Age  int
}

// Define a method on Person
// Greet() has a receiver of (p Person)
// p is the placeholder
func (p Person) Greet() string {
	return fmt.Sprintf("Hello, my name is %s and I am %d years old.", p.Name, p.Age)
}

func structMethods() {
	// Create an instance of Person
	alice := Person{Name: "Alice", Age: 30}

	// Call the Greet method
	fmt.Println(alice.Greet())
}
