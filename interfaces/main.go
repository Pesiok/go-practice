package main

import "fmt"

// a contract
// an abstract type that declare
// methods are implemented by concrete types
type Shape interface {
	area() float64
}

// this type implements Shape interface
type Square struct {
	side float64
}

// like implementing method on a class
func (z Square) area() float64 {
	return z.side + z.side
}

// this type also implements Shape interface
type Circle struct {
	radius float64
}

func (c Circle) area() float64 {
	return c.radius * c.radius
}

func info(z Shape) {
	// different implementation of the same functionality
	fmt.Println(z)
	fmt.Println(z.area())
}

func main() {
	// c and s implements the Shape
	s := Square{10}
	c := Circle{5}
	info(s)
	info(c)
}
