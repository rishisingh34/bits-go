package main 

import (
	"fmt"
)

func pass_by_reference(){
	x := "a string"
	fmt.Println("Before calling increment, x =", x)
	increment(&x) // passing the address of x
	fmt.Println("After calling increment, x =", x)
} 
func increment(ptr *string) {
	*ptr = *ptr + " plus something"
	fmt.Println("Inside increment, value at ptr =", *ptr)
}