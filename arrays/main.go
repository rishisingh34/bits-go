package main 

import "fmt"

func main() {
	var option int 

	fmt.Print("Choose an option: ")
	fmt.Scan(&option)

	switch option {
		case 1:
			// 1. declaration 
			var myInts [10]int

			fmt.Println(myInts)
			// initialization
			for i := range 10 {
				myInts[i] = i * 2
			}
			// usage
			for i, v := range myInts {
				fmt.Printf("myInts[%d] = %d\n", i, v)
			}
		case 2:
			// 2. Another way to declare
			myInts := [10]int{0, 2, 4, 6, 8, 10, 12, 14, 16, 18}
			fmt.Println(myInts)
		case 3:
			// making a slice 
			sliceInGo()
		default:
			fmt.Println("Invalid option")
	}
}