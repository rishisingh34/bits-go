package main 

func switchInGo(option string) {
	switch option {
	case "1":
		println("Option 1 selected")
	case "2":
		println("Option 2 selected")

	// fallthrough if it arrives at "3" it will also execute "4" due to the fallthrough statement
	case "3":
		println("Option 3 selected")
		fallthrough
	case "4":
		println("Option 4 selected")
	default:
		println("Invalid option")
	}
}