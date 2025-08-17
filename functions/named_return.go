// name_return.go
package main


// namedReturn demonstrates the use of **named return values** in Go.
//
// In Go, you can declare return variable names in the function signature.
// These names act like regular local variables, initialized to their zero values.
// Using a bare `return` statement will return the current values of these variables.
//
// Benefits:
//  1. Improves readability when function has multiple return values.
//  2. Useful in longer functions where values are updated in multiple places.
//  3. Helps to avoid repeating variable names in the return statement.
//
// Downsides:
//  1. Can reduce clarity if overused in short functions.
//  2. Bare returns without context may be confusing.
//
// Example:
//  func divide(x, y int) (result int, err error) {
//      if y == 0 {
//          err = fmt.Errorf("cannot divide by zero")
//          return // returns (0, err)
//      }
//      result = x / y
//      return // returns (result, nil)
//  }
//
// In this example, both `result` and `err` are named return variables.
//
func namedReturn() (sum int, product int) {
	// Named returns: sum and product are already declared with type int, initialized to 0
	a, b := 5, 10

	// Assign values to the named return variables
	sum = a + b
	product = a * b

	// bare return → returns (sum, product)
	return
}