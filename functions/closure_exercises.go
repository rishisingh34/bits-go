package main

import "fmt"

// Write a closure that counts how many times it’s called.
// -------- EX1 --------
func newCounter() func() int {
	cnt := 0
	return func() int {
		cnt++
		return cnt
	}
}

func ex1() {
	counter := newCounter()
	fmt.Println(counter())
	fmt.Println(counter())
	fmt.Println((counter()))
}

// -------- EX1 --------

// Write a closure that remembers a person’s name and always greets them.
// -------- EX2 --------

func greeter(name string) func() {
	return func() {
		fmt.Printf("Hello, %s\n", name)
	}
}

func ex2() {
	greetRishi := greeter("Rishi")
	greetRishi()
	greetRishi()
	greetRishi()
}

// -------- EX2 --------

// Write a closure that multiplies numbers by a fixed factor.
// -------- EX3 --------

func multiplier(num int) func(int) {
	return func(num2 int) {
		fmt.Println(num2 * num)
	}
}

func ex3() {
	double := multiplier(2)
	triple := multiplier(3)

	double(10) // 20
	triple(10) // 30
}

// -------- EX3 --------

//Create a closure that allows only N calls, then blocks further calls.
// -------- EX4 --------

// Define a User
type User struct {
	name   string
	email  string
	number int
}

// rateLimiter wraps a handler and allows only N calls
func rateLimiter(handler func(User) User, N int) func(User) (User, error) {
	cnt := 0
	return func(u User) (User, error) {
		cnt++
		if cnt > N {
			return User{}, fmt.Errorf("❌ rate limit exceeded, try later")
		}
		return handler(u), nil
	}
}

// Example usage
func ex4() {
	var users []User

	addUser := func(u User) User {
		fmt.Println("Processing user:", u.name)
		users = append(users, u)
		return u
	}

	limitedHandler := rateLimiter(addUser, 2)

	u1 := User{name: "Rishi", email: "singh34rishi@gmail.com", number: 11232}
	u2 := User{name: "Bob", email: "bob@example.com", number: 5678}
	u3 := User{name: "Eve", email: "eve@example.com", number: 9999}

	if res, err := limitedHandler(u1); err == nil {
		fmt.Println("OK:", res)
	} else {
		fmt.Println(err)
	}

	if res, err := limitedHandler(u2); err == nil {
		fmt.Println("OK:", res)
	} else {
		fmt.Println(err)
	}

	if res, err := limitedHandler(u3); err == nil {
		fmt.Println("OK:", res)
	} else {
		fmt.Println(err)
	}

	// Finally, print all users in global slice
	fmt.Println("\n📦 Users stored globally:")
	for _, u := range users {
		fmt.Printf("- %s (%s)\n", u.name, u.email)
	}
}

// -------- EX4 --------
