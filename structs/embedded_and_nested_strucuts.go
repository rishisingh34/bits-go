package main

import "fmt"

func embeddedVsNested() {
	type Address struct {
		Street, City, State string
	}

	// --- Embedded struct ---
	type PersonEmbedded struct {
		Name    string
		Age     int
		Address // ✅ Embedded (anonymous field)
	}

	// --- Nested struct ---
	type PersonNested struct {
		Name    string
		Age     int
		Address Address // ❌ Nested (named field)
	}

	// Example person with embedded Address
	p1 := PersonEmbedded{
		Name: "Alice",
		Age:  25,
		Address: Address{
			Street: "111 Embedded St",
			City:   "GoTown",
			State:  "CA",
		},
	}

	// Example person with nested Address
	p2 := PersonNested{
		Name: "Bob",
		Age:  30,
		Address: Address{
			Street: "222 Nested Ave",
			City:   "StructCity",
			State:  "NY",
		},
	}

	// --- Accessing fields ---
	fmt.Println("=== Embedded Example ===")
	fmt.Println("Name:", p1.Name)          // direct
	fmt.Println("Street:", p1.Street)      // ✅ promoted (shortcut)
	fmt.Println("Full Address:", p1.Address.Street, p1.Address.City, p1.Address.State) // also valid 

	fmt.Println("\n=== Nested Example ===")
	fmt.Println("Name:", p2.Name)              // direct
	// fmt.Println(p2.Street)                  // ❌ does NOT work, not promoted
	fmt.Println("Street:", p2.Address.Street)  // ✅ must go through Address
}
