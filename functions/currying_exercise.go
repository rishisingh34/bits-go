package main

import "fmt"

// -------- EX1 --------
// Write a curried function in Go for adding two integers:

func addTwoIntegers(a int) func(int) int {
	return func(b int) int {
		return a + b
	}
}

func curr_ex1() {
	fmt.Println(addTwoIntegers(2)(3)) // successively calling nested closures until you reach the final return value

	// another way to call nested closures
	addAnotherIntegerWith2 := addTwoIntegers(2)
	fmt.Println(addAnotherIntegerWith2(3))
}

// -------- EX1 --------

// -------- EX2 --------
// string concatenation through currying

func concat(first string) func(string) string {
	return func(second string) string {
		return first + " " + second
	}
}

func curr_ex2() {
	first := "Hi"
	second := "rishi"
	fmt.Println(concat(first)(second))

}

// -------- EX2 --------

// -------- EX3 --------
// Filtering with Currying: Implement a curried function that creates filters:

func filterGreater(num int) func ([]int) []int {
	return func(arr []int) []int {
		newArr := []int{}  
		for _, val := range arr {
			if val > num {
				newArr = append(newArr, val)
			}
		}
		return newArr
	}
}  

func curr_ex3() {
	fmt.Println(filterGreater(10)([]int{1, 56, 20, 2, 99}))
}

// -------- EX3 --------


// -------- EX4 --------
// Currying with predicates
func filter(pred func(int) bool) func([]int) []int {
	return func(arr []int) []int {
		newArr := []int{}
		for _, val := range arr {
			if pred(val) {
				newArr = append(newArr, val)
			}
		}
		return newArr
	}
}

func curr_ex4() {
	// Example 1: filter even numbers
	isEven := func(n int) bool { return n%2 == 0 }
	evenFilter := filter(isEven)
	fmt.Println(evenFilter([]int{1, 2, 3, 4, 5, 6})) // [2 4 6]

	// Example 2: filter numbers greater than 10
	filteredArray := filter(func(num int) bool {
		return num > 10
	}) ([]int{5, 12, 8, 20})

	fmt.Println(filteredArray)
}

// -------- EX4 --------


