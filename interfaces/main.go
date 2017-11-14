package main

import (
	"fmt"
	"strconv"

	"github.com/Pesiok/go-practice/interfaces/exercises"
)

// a contract
// an _abstract_ type that declares methods
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
	exc.SortExc()
	conversionAndAssertion()
}

// interface{} <- every type implements it
// example: fmt.Prinln(a ...interface{})

// method sets:

// # value receiver -> value & pointer
// func (c Circle) area() float64 { ... }

// # pointer receiver -> only pointer
// (given value address might not exist)
// func (c *Circle) area() float64 { ... }
// used when:
// 1. method need to modify the receiver (mutation)
// 2. for efficiency when passed data is large (html page?)

// conversion vs assertion

func conversionAndAssertion() {
	// conversion -> from one basic type to other
	x := 12
	y := 12.13123123
	ż := "12"
	// widening the value
	fmt.Println("Float conversion: ", y+float64(x))
	// asci to int
	z, _ := strconv.Atoi(ż)
	fmt.Println("ASCI/INT conversion: ", z+x)

	// assertions -> only for interfaces
	var name interface{} = "Namae"
	var val interface{} = 7
	str, ok := name.(string)
	if ok {
		fmt.Printf("%T\n", str)
	} else {
		fmt.Println("Value is not a string")
	}

	// k - assertion
	fmt.Println(val.(int) + 7)
	// err - conversion
	// fmt.Println(int(val) + 7)
}
