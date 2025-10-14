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

	// consumed, err :=doSomething(false)
	// if err != nil {
	// 	fmt.Println("Error occurred:", err)
	// }
	// fmt.Println("Consumed:", consumed)
}


// ---- Custom Error Example ------
// type MyError struct {
// 	Code string
// 	Msg  string
// }

// func (e *MyError) Error() string {
// 	return fmt.Sprintf("Error - Code: %s, Msg: %s", e.Code, e.Msg)
// }

// func doSomething(flag bool) (string, error) {
// 	if !flag {
// 		return "", &MyError{Code: "400", Msg: "Bad Request"}
// 	}
// 	return "Success", nil
// }
