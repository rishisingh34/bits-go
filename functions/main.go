package main

import (
	"fmt"
)

func main() {
	var option int
	fmt.Print("Enter the option: ")
	fmt.Scan(&option)

	switch option {
	case 1:

		// 1. function signature
		fmt.Println(absSub(5, 10))

	case 2:

		// 2. multiple paramaters
		multipleParameters(30, 10, "Goblin")

	case 3:

		// 3. complex signature
		complexSignature(absSub, 20, 15)

	case 4:

		// 4. pass by value
		passByValue()

	case 5:

		// 5. ignoring return values
		x, _ := ingnoringReturnValues()
		fmt.Println("Printing only first value:", x)

	case 6:

		// 6. named return
		sum, product := namedReturn()
		fmt.Printf("Named return - Sum: %d, Product: %d\n", sum, product)

	case 7:

		// 7. anonymous function
		anonymousFunc()

	case 8:

		// 8. defer keyword
		useDefer()

	case 9:

		//9. closure
		closure()

	case 10:

		// 10. closure examples
		ex1()
		ex2()
		ex3()
		ex4()

	case 11:

		// 11. Currying
		currying()

	case 12:

		// 12. currying examples
		curr_ex1()
		curr_ex2()
		curr_ex3()
		curr_ex4()

	default:
		fmt.Println("Invalid Option")
	}

}
