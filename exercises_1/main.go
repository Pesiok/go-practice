package main

import (
	"fmt"
)

func main() {
	// three()
	// four()
	// five(100)
	// six(100)
	seven(1000)
}

func one() {
	fmt.Println("Hello world")
}

func two() {
	name := "my name"
	fmt.Println("Hello" + name)
}

func three() {
	var name string
	fmt.Print("What's your name?")
	// write value to &name memory address
	fmt.Scan(&name)
	fmt.Println("Hello " + name)
}

func four() {
	var (
		small int
		big   int
	)

	fmt.Print("Small number: ")
	fmt.Scan(&small)
	fmt.Print("Big number: ")
	fmt.Scan(&big)

	fmt.Println(big % small)

}

func five(n int) {
	for i := 0; i <= n; i++ {
		if i%2 != 0 {
			continue
		}

		fmt.Println(i)
	}
}

func six(n int) {
	for i := 0; i <= n; i++ {
		var output string
		if i%3 == 0 {
			output += "Fizz"
		}
		if i%5 == 0 {
			output += "Buzz"
		}

		if output != "" && i != 0 {
			fmt.Println(output)
			continue
		}

		fmt.Println(i)
	}
}

func seven(n int) {
	var sum int
	for i := 0; i < n; i++ {
		if i%3 == 0 || i%5 == 0 {
			sum += i
		}
	}
	fmt.Println(sum)
}
