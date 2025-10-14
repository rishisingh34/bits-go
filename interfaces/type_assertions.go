package main

import (
	"fmt"
)

func typeAssertions() {
	var x any = "hello"
	s, ok := x.(string)
	if ok {
		fmt.Println(s)
	} else {
		fmt.Println("x is not a string")
	}
}


func typeAssertions2() {
  var circ shape  = circle{radius: 5}

   c, ok := circ.(circle)
   if ok {
       fmt.Println(c.radius)
   } else {
       fmt.Println("shape is not a circle")
   }
}