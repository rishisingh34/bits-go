package main

import "fmt"

func main() {
	var option int 

	fmt.Println("Choose an option: ")
	fmt.Scanln(&option)

	switch option {
	case 1:
		// 1. Definition of interfaces
		definition()
	case 2:
		// 2. Type assertions
		typeAssertions()
		typeAssertions2()
	case 3: 
		// 3. type switches 
		typeSwitches()
	default:
		fmt.Println("Invalid option")
	}
}