package main

import "fmt"

func main() {
	// fmt.Println(greetMultiple("KEK ", "KAK"))
	// fmt.Println(average(1, 2, 1))
	// rest()

	// func expression ->
	// only way you can declare function in another one
	greeting := func() {
		fmt.Println("Hello")
	}

	greeting()
	fmt.Printf("%T\n", greeting)

	// IIFE
	func() {
		fmt.Println("World")
	}()

	greet := makeGreeter()
	fmt.Println(greet())

}

// returns
func greet(fname, lname string) string {
	return fmt.Sprint(fname, lname)
}

// named returns
func greetNamed(fname, lname string) (s string) {
	s = fmt.Sprint(fname, lname)
	return
}

// return multiple
func greetMultiple(fname, lname string) (string, string) {
	return fmt.Sprint(fname, lname), "ok"
}

// variadic (params) - could be invoked with 0 or more arguments
func average(sf ...float64) float64 {
	var total float64
	// looping over a collection, (sf Slice)
	// it return index and value
	// index is not used so it's blank
	for _, v := range sf {
		total += v
	}
	return total / float64(len(sf))
}

func rest() {
	data := []float64{43, 2, 34, 35, 65}
	// varaidic (arguments)
	n := average(data...)
	fmt.Println(n)
}

// function maker
func makeGreeter() func() string {
	return func() string {
		return "greetings"
	}
}
