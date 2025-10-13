package main

import "fmt"

func counter(start int) (func() int, func() int) {
	cnt := start
	return func() int {
			cnt++
			return cnt
		}, func() int {
			return cnt
		}
}

func test() {
	inc, peek := counter(0)
	fmt.Println(peek())
	fmt.Println(inc())
	fmt.Println(peek())
}
