package main

import "fmt"

func constants() {
	// constants can be primitive types
	// strings, integers, booleans, floats
	const Pi = 3.14

	// constants cannot be assigned from variable values
	radius := 5.0
	// const Area = Pi * radius // This will cause a compilation error

	// constants can be made from other constants or constant expressions
	const TwoPi = 2 * Pi

	fmt.Println("Value of Pi:", Pi)
	fmt.Println("Value of TwoPi:", TwoPi)

	// this next line will throw error cause radius is not const 
	// const area = Pi * radius * radius

	// fmt.Println("Value of Area:", area)

	fmt.Println("Value of Radius:", radius)
}