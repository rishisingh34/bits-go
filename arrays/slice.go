package main

import "fmt"

func sliceInGo() {
    arr := []int{1, 2, 3, 4}
    slice := arr[1:3]

    fmt.Println("Slice:", slice)

    // Print address of first element in slice
    fmt.Println("Address of slice[0]:", &slice[0])

    // Print address of arr[1], which is backing slice[0]
    fmt.Println("Address of arr[1]:", &arr[1])

    // Print address of arr[0] just for comparison
    fmt.Println("Address of arr[0]:", &arr[0])


	// make 
	// func make([]T, len, cap) []T -- cap can be omitted 
	mySlice := make([]int, 10)

	fmt.Println(len(mySlice))
	fmt.Println(mySlice)

}