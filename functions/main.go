package main 

import (
	"fmt"
)

func main() {
	// 1. function signature 
	fmt.Println(absSub(5,10))

	// 2. multiple paramaters 
	multipleParameters(30, 10, "Goblin")

	// 3. complex signature 
	complexSignature(absSub, 20, 15)

	// 4. pass by value 
	passByValue()

	// 5. ignoring return values
	x, _ := ingnoringReturnValues()
	fmt.Println("Printing only first value:", x)

	// 6. named return
	sum, product := namedReturn()
	fmt.Printf("Named return - Sum: %d, Product: %d\n", sum, product)

	// 7. anonymous function 
	anonymousFunc()	

	// 8. defer keyword
	useDefer()

	//9. closure
	closure()

	// 10. closure examples
	ex1()
	ex2()
	ex3()
}