package main

import "fmt"

func useDefer() error {
	// Close the connection *anywhere* the useDefer function returns 
	defer fmt.Println("Deferred call")


	fmt.Println("Regular call")

	return nil
}