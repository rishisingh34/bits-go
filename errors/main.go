package main 

import "fmt"

// what is error-interface ---> 
// in Go, the built-in error type is an interface
// type error interface {
// 	Error() string
// }

func main() {
    var option int
    fmt.Print("Enter an option: ")
    fmt.Scan(&option)

    switch option {
    case 1:
        // 1. custom error
		_, err := divide(4, 0)
		if err != nil {
			fmt.Println("Error:", err)
		}
    case 2:
        // 2. errors package 
		_, err := errors_package(4, 0)
		if err != nil {
			fmt.Println("Error:", err)
		}
    case 3:
        // 3. panic and recover 
		// never use panic for control flow 

	case 4: 
		// 4. fmt.Errorf is GOATed: it wraps an error with additional context
		_, err := helper(4, 0)
		if err != nil {
			fmt.Println("Error:", err)
		}
    default:
        fmt.Println("Invalid option.")
    }
}


// ---- Helper Function ------ 
func helper(x, y float64) (float64, error) {
	if y == 0 {
		return 0, fmt.Errorf("no dividing by 0, %f", x)
	}
	return x / y, nil
}