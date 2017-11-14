package main

import "fmt"

func main() {
	// one()
	// two()
	// fmt.Println(<-factorial(3))
}

func one() {
	// deadlock
	c := make(chan int)
	// launch goroutine
	go func() {
		c <- 1
	}()
	// wait unitl it produce value
	fmt.Println(<-c)
}

func two() {
	c := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		close(c)
	}()

	for value := range c {
		fmt.Println(value)
	}
}

func factorial(n int) <-chan int {
	output := make(chan int)
	go func() {
		if n == 1 {
			output <- 1
		} else {
			output <- n * <-factorial(n-1)
		}
	}()
	return output
}
