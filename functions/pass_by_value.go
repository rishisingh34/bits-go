package main

import "fmt"

func passByValue() {
	fmt.Println("\n--- Pass by Value Demonstration ---")

	// 1. Basic types (int, float, bool) → fully copied
	x := 10
	changeInt(x)
	fmt.Println("int after function call:", x) // still 10

	// 2. Array → fully copied
	arr := [3]int{1, 2, 3}
	changeArray(arr)
	fmt.Println("array after function call:", arr) // unchanged [1 2 3]

	// 3. Struct → fully copied
	type Point struct{ X, Y int }
	p := Point{1, 2}
	changeStruct(p)
	fmt.Println("struct after function call:", p) // unchanged {1 2}

	// 4. Slice → header copied, but underlying array shared
	slice := []int{1, 2, 3}
	changeSlice(slice)
	fmt.Println("slice after function call:", slice) // modified [99 2 3]

	// 5. Map → header copied, but underlying hash table shared
	m := map[string]int{"a": 1, "b": 2}
	changeMap(m)
	fmt.Println("map after function call:", m) // modified map[a:100 b:2]

	// 6. Channel → reference to same channel
	ch := make(chan int, 1)
	changeChannel(ch)
	fmt.Println("channel after function call has:", <-ch) // 42

	// 7. String → descriptor copied, but immutable so no effect
	str := "hello"
	changeString(str)
	fmt.Println("string after function call:", str) // still "hello"
}

// ------------------ Helper functions ---------------------

func changeInt(n int) {
	fmt.Println("int inside function before change:", n) // 10
	n = 99
	fmt.Println("int inside function:", n) // 99, but does not affect caller
}

func changeArray(a [3]int) {
	fmt.Println("array inside function before change:", a) // [1 2 3]
	a[0] = 99
	fmt.Println("array inside function:", a) // [99 2 3], but does not affect caller
}

func changeStruct(p struct{ X, Y int }) {
	fmt.Println("struct inside function before change:", p) // {1 2}
	p.X = 99
	fmt.Println("struct inside function:", p) // {99 2}, but does not affect caller
}

func changeSlice(s []int) {
	s[0] = 99
}

func changeMap(m map[string]int) {
	m["a"] = 100
}

func changeChannel(ch chan int) {
	ch <- 42
}

func changeString(s string) {
	// strings are immutable; cannot modify contents
	// assigning new value does not affect caller
	// s = "changed"
	fmt.Println("string is immutable:", s)
}
