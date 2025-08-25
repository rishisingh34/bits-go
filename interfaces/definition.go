package main

import (
	"fmt"
	"math"
)

//  ** for a type to implement an interface it has 
// to implement all methods of that interface ** 

// all structs implements empty interface 
// interface{} - implemented by all types

// ** a type can implement multiple interfaces 

type shape interface {
  area() float64
  perimeter() float64
}

type rect struct {
    width, height float64
}
// type rect implements shape
func (r rect) area() float64 {
    return r.width * r.height
}
func (r rect) perimeter() float64 {
    return 2*r.width + 2*r.height
}

type circle struct {
    radius float64
}
// type circle implements shape
func (c circle) area() float64 {
    return math.Pi * c.radius * c.radius
}
func (c circle) perimeter() float64 {
    return 2 * math.Pi * c.radius
}

func printShapeData(s shape) {
	fmt.Printf("Area: %v - Perimeter: %v\n", s.area(), s.perimeter())
}

func definition() {
	printShapeData(rect{width: 3, height: 4})
	printShapeData(circle{radius: 5})
}