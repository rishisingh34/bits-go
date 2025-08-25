package main

import (
	"fmt" 
)

func main() {
	var option int 
	fmt.Print("Enter the option: ")
	fmt.Scan(&option)

	switch option {
	case 1: 
		// 1. definition 
		structsInGo()
	case 2:
		// 2. anonymous structs
		anonymStructs()
	case 3: 
		// 3. embedded & nested structs 
		embeddedVsNested()
	case 4:
		// 4. struct methods
		structMethods()

	case 5: 
		// 5. memory layout in go for structs 
		// keep larger fields first in a struct 
	case 6:
		// 6. empty structs 
		emptyStructs()
	default:
		fmt.Println("Invalid Option")
	}
}