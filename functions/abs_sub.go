package main

// function signature -->
// func absSub(x int, y int) int

func absSub(x int, y int) int {
	if x < y {
		return y - x
	}
	return x - y
}
