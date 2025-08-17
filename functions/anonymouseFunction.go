package main

import "fmt"

func anonymousFunc() {
	// using a named function
	newX, newY, newZ := conversions(double, 1, 2, 3)
	// newX is 2, newY is 4, newZ is 6

    // using an anonymous function
	newX1, newY1, newZ1 := conversions(func(a int) int {
	    return a + a
	}, 1, 2, 3)
	// newX is 2, newY is 4, newZ is 6
	fmt.Println("Converted values using named function:", newX, newY, newZ)
	fmt.Println("Converted values using anonymous function:", newX1, newY1, newZ1)
}

func conversions(converter func(int) int, x, y, z int) (int, int, int) {
	convertedX := converter(x)
	convertedY := converter(y)
	convertedZ := converter(z)
	return convertedX, convertedY, convertedZ
}
func double(a int) int {
	return 2*a
}