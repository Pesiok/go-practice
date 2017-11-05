package main

// imported packages are in file block scope
import (
	"fmt"

	"github.com/Pesiok/go-practice/helloworld/helloworlder"
	"github.com/Pesiok/go-practice/helloworld/memory"
)

// constants
const (
	Pi             = 3.14
	constString    = "very important costant value"
	constStringTwo = "that's also a very important string"
)

const (
	hello      string = "hello"
	typedHello        = "hello"
)

// ioatas
const (
	a = iota
	b
	c
)

const (
	_ = iota
	d = iota * 10
	e = iota * 10
)

const (
	_ = iota
	// bitwise operations:
	// bitshift
	kb = 1 << (iota * 10)
	mb = 1 << (iota * 10)
)

func main() {
	// const
	fmt.Println("constString", constString)
	// iotas
	fmt.Println("iota A ", a)
	fmt.Println("iota B ", b)
	fmt.Println("iota C ", c)
	fmt.Println("iota D ", d)
	fmt.Println("iota E ", e)
	fmt.Println("iota KB ", kb)
	fmt.Println("iota MB ", mb)
	// text
	fmt.Println(helloworlder.Version)
	helloworlder.Call()
	// format (decimal, binary, hexidecimal, utf8)
	fmt.Printf("%d - %b - %#x - %q \n", 42, 42, 42, 42)
	fmt.Printf("%d - %b - %#X - %q \n", 42, 42, 42, 42)
	fmt.Printf("%d \t %b \t %x \t %q \n", 42, 42, 42, 42)
	// loop
	for i := 0; i < 15; i++ {
		fmt.Printf("%d - %b - %#X - %q \n", i, i, i, i)
	}

	// http & blank usaage
	// get.Google()

	// memory
	memory.Mem()
	// memory.MetersToYards()
	memory.Pointers()
	memory.Usingpointers()

	// default values -> 'zero' values

	// declare ...and assign
	// var b string
	// b = "kek"

	// variables - shorhand
	// only in func
	// decalre & assign
	a := 10
	b := "golang"
	c := 31.13
	d := false

	fmt.Printf("%T \n", a)
	fmt.Printf("%T \n", b)
	fmt.Printf("%T \n", c)
	fmt.Printf("%T \n", d)

	// scope: universe, package, file, block
	kek()
	kak()

	// closure and anonymous function expression
	times := 5
	firstInc := incrementer()
	secondInc := incrementer()

	for i := 0; i < times; i++ {
		fmt.Println(firstInc())
	}

	for i := 0; i < times; i++ {
		fmt.Println(secondInc())
		fmt.Println(firstInc())
	}

}

// package level scope
func incrementer() func() int {
	counter := 0
	return func() int {
		counter++
		return counter
	}
}

func kek() {
	y := "kek"
	fmt.Println(y)
}

func kak() {
	y := "kak"
	// block scope, shadowing
	{
		fmt.Println(y)
		y := "ajj"
		x := "test"
		fmt.Println(y, x)
	}
	fmt.Println(y, four)
	// fmt.Println(x) -> x out of scope
}

func multiply(x int, y int) int {
	// function body scope
	return x * y
}

// var and func are hoisted
var four = multiply(2, 2)

func remainder() {
	x := 13 % 3
	fmt.Println(x)

	for i := 1; i < 10; i++ {
		if i%2 == 1 {
			fmt.Println("Odd")
		} else {
			fmt.Println("Even")
		}
	}
}
