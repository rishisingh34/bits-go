package main

import (
	"fmt"
)

func main() {

	var option int
	fmt.Println("Choose an option:")
	fmt.Scanln(&option)

	switch option {
	case 1:
		// 1. pointer is a variable that stores the memory address of another variable
		var a int = 10
		var p *int = &a
		fmt.Println("Value of a:", a)
		fmt.Println("Address of a:", &a)
		fmt.Println("Value of p (address of a):", p)
		fmt.Println("Value at address stored in p:", *p)

	case 2:
		// 2. dereferencing a pointer ( * operator )
		str := "a string"
		str_ptr := &str
		fmt.Printf("Value of str: %v\n", *str_ptr)

	case 3:
		// 3. pass by reference
		pass_by_reference()

	case 4:
		// 4. pointer arithmetic (not supported in Go)
		fmt.Println("Pointer arithmetic is not supported in Go")

	case 5:
		// 5. nil pointer
		// var p *int
		// if p == nil {
		// 	fmt.Println("p is a nil pointer")
		// } else {
		// 	fmt.Println("p is not a nil pointer")
		// }

	case 6:
		// 6. pointer to struct
		type Person struct {
			name string
			age  int
		}
		p := &Person{name: "Alice", age: 30}
		fmt.Println("Person name:", p.name)
		fmt.Println("Person age:", p.age)

		// 💡 Why use a pointer to a struct?
		/*
			1. Efficiency:
				Passing a large struct by value copies the entire struct.
				Using a pointer avoids that copy — only an 8-byte address is passed.

			2. Mutability:
				If you pass a struct by value, any changes inside a function are on the copy.
				Using a pointer allows you to modify the original struct’s fields.

			3. Method Receivers:
				Struct methods often use pointer receivers (*T) so that they can modify fields
				and avoid copying large data.

			4. Consistency:
				When you create objects dynamically (like &Person{}), they naturally return pointers.
				You can share them safely across functions without unnecessary copies.

			In short:
				- Passing struct → makes a copy (pass by value)
				- Passing *struct → references the original (pass by reference)
		*/
		fmt.Println("\n--- Why use pointer to struct? ---")
		fmt.Println("1. Efficiency: avoids copying large structs")
		fmt.Println("2. Mutability: allows modification of original struct")
		fmt.Println("3. Method receivers: enables modifying struct inside methods")
		fmt.Println("4. Consistency: easily share same instance across functions")

	case 7:
		// 7. pointer receiver
		pointer_receiver()
	default:
		fmt.Println("Invalid option")
	}
}