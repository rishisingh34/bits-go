package main 

import "fmt"


func typeSwitches() {
	var num any = 42 
	switch v := num.(type) {
	case int:
		fmt.Println("Integer:", v)
	case string:
		fmt.Println("String:", v)
	default:
		fmt.Println("Unknown type")
	}

	type EmptyStruct struct{}
	var es any = EmptyStruct{}
	switch v := es.(type) {
	case int:
		fmt.Println("Integer:", v)
	case string:
		fmt.Println("String:", v)
	default:
		fmt.Println("Unknown type")
	}
}