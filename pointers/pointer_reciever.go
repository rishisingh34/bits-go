package main

import (
	"fmt"
)

type circle struct {
	x int
	y int
    radius int
}

func (c *circle) grow() {
    c.radius *= 2
}

func pointer_receiver() {
    c := circle{
        x: 1,
        y: 2,
        radius: 4,
    }

    // notice c is not a pointer in the calling function
    // but the method still gains access to a pointer to c
    c.grow()
    fmt.Println(c.radius)
    // prints 8
}