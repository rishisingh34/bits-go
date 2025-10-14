package main 

import (
	"fmt"
)

func main() {
	var option int

	fmt.Println("Choose an option:")
	fmt.Scanln(&option)

	switch option {
	case 1: 
		// 1. creating a map 
		ages := make(map[string]int)

		// zero value of map is nil
		ages["john"] = 25
		ages["doe"] = 30
		ages["alice"] = 28

		print_mp(ages)

	case 2: 
		// 2. creating and initlizing a map 
		ages := map[string]string{
			"john": "25",
			"doe": "30",
			"alice": "28",
		}
		print_mp(ages)

	case 3: 
		// 3. operations on map

		ages := map[string]int{
			"john": 25,
			"doe": 30,
			"alice": 28,
		}

		// insert or update an element
		ages["john"] = 26

		// get an element 
		print("John's age is ", ages["john"])

		// delete an element 
		delete(ages, "doe")
		print_mp(ages)

		// check if a key exists
		if age, ok := ages["doe"]; ok {
			print("Doe's age is ", age)
		} else {
			print("Doe not found")
		}	

		// len of map
		print("Length of map is ", len(ages))
		
	case 4: 
		// 5. nested maps
		people := map[string]map[string]int{
			"john": {
				"age": 25,
				"height": 175,
			},
			"doe": {
				"age": 30,
				"height": 180,
			},
		}

		fmt.Println(people)
	default: 
		print("Invalid option")
	}
}

func print(a ...any) {
	for i, v := range a {
		fmt.Print(v)
		if i != len(a)-1 {
			fmt.Print(" ")
		}
	}
	fmt.Println()
}

func print_mp(m any) {
	mm, ok := m.(map[string]int)
	if !ok {
		fmt.Println("Not a map")
		return
	}
	for k, v := range mm {
		fmt.Println("Key: ", k, ", Value: ", v)

	}
}