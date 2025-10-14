package main

import (
	"fmt"
)

func main() {
	var option int

	fmt.Println("Select an option:")
	fmt.Scanln(&option)

	switch option {
	case 1:
		// 1. arrays in go
		var arr [5]int // --> all initlized to 0
		fmt.Println(arr)
	case 2:
		// 2. declare and initialize an array in one line
		arr := [5]int{1, 2, 3, 4, 5}
		fmt.Println(arr)
	case 3:
		// 3. slice ( dynamic array )

		// an array of 5 integers
		arr := [5]int{1, 2, 4, 3, 5}

		// arrayname[startindex : highindex]
		// arrayname[startindex:]
		// arrayname[:highindex]
		// arrayname[:]

		slice := arr[1:4] // 2,4,3
		fmt.Println(slice)

	case 4:
		// 4. declare and initialize a slice in one line
		slice := []int{1, 2, 3, 4, 5} // --> not specifying size creates a slice
		fmt.Println(slice)

	case 5:
		// 5. append to a slice
		slice := []int{1, 2, 3}

		slice = append(slice, 4)          // append returns a new slice with new underlying array if needed
		slice = append(slice, 5, 6, 7, 8) // append multiple values

		fmt.Println(slice)

	case 6:
		// Create a slice literal with 3 elements
		slice1 := []int{1, 2, 3} // len = 3, cap = 3

		// Create another slice of length 2 (initialized with zeros)
		slice2 := make([]int, 2) // [0, 0], len = 2, cap = 2

		// Copy elements from slice1 → slice2
		// copy(dst, src) copies min(len(dst), len(src)) elements
		n := copy(slice2, slice1) // copies first 2 elements: 1, 2

		// slice1 remains unchanged
		// slice2 now has the first two elements of slice1
		// n is the number of elements actually copied
		fmt.Println(slice1, slice2, n) // Output: [1 2 3] [1 2] 2

	case 7:
		// 7. length and capacity of a slice
		slice := []int{1, 2, 3, 4, 5}

		fmt.Println("Length:", len(slice))   // length
		fmt.Println("Capacity:", cap(slice)) // capacity

	case 8:
		// 8. spread operator to pass a slice to a variadic function
		names := []string{"Alice", "Bob", "Charlie", "Diana", "Eve"}
		newNames := []string{"Frank", "Grace"}

		// Append newNames to names -- using spread operator to pass slice to variadic function
		names = append(names, newNames...)
		fmt.Println(names)

	case 9:
		// 9. variadic function
		sum := func(nums ...int) int { // nums is a slice of int
			total := 0
			for _, num := range nums {
				total += num
			}
			return total
		}

		result := sum(1, 2, 3, 4, 5)
		fmt.Println("Sum:", result)

		result = sum(10, 20)
		fmt.Println("Sum:", result)

		// another example of designing variadic function
		print(1, "hello", 3.14, true)

	case 10: 
		// 10. Range 
		slice := []int{10, 20, 30, 40, 50}
		
		for index, value := range slice {
			pf("Index: %d, Value: %d\n", index, value)
		}

	case 11: 
		// 11. Multidimensional slices
		matrix := [][]int{
			{1, 2, 3},
			{4, 5, 6},
			{7, 8, 9},
		}

		for i, row := range matrix {
			for j, value := range row {
				pf("matrix[%d][%d] = %d\n", i, j, value)
			}
		}
	case 12:
		// 12. Slicing a slice
		slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

		subSlice1 := slice[2:5] // elements from index 2 to 4
		subSlice2 := slice[:4]  // first four elements
		subSlice3 := slice[5:]  // elements from index 5 to end

		fmt.Println("Original Slice:", slice)
		fmt.Println("Sub-slice 1 (2:5):", subSlice1)
		fmt.Println("Sub-slice 2 (:4):", subSlice2)
		fmt.Println("Sub-slice 3 (5:):", subSlice3)

	case 13:
		// 13. Deleting an element from a slice
		slice := []int{1, 2, 3, 4, 5}

		// Delete element at index 2 (value 3)
		indexToDelete := 2
		slice = append(slice[:indexToDelete], slice[indexToDelete+1:]...)

		fmt.Println("After deletion:", slice) // Output: [1 2 4 5]

	default:
		fmt.Println("Invalid option")
	}
}

func print(a ...any) {
	for i, v := range a {
		if i != len(a)-1 {
			fmt.Print(v, " ")
			continue
		}
		fmt.Print(v)
	}
}

func pf(format string, a ...any) {
	fmt.Printf(format, a...)
}