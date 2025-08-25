package main 

import "fmt"

func anonymStructs() {
	myCar := struct{
		brand string
		model string 
	} {
		brand: "Toyota", 
		model: "camry",
	}
	fmt.Println(myCar)
}