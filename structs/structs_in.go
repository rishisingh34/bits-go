package main

import "fmt"

type car struct {
	brand string
	model string 
	doors int 
	mileage int 
}


func structsInGo() {
	newCar := car {
		brand : "Toyota",
		model: "Camry",
		doors : 4 ,
		mileage: 20,
	}
	fmt.Println(newCar)

	// empty struct 
	newCar2 := car{}
	// by default brand & model is empty string and doors & mileage are 0 
	fmt.Println(newCar2)
}