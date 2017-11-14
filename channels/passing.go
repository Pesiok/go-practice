package main

import (
	"fmt"
)

func passing() {
	// c := incrementor()
	// cSum := puller(c)

	fmt.Println(<-puller(incrementor()))
}

// or func incrementor() <-chan int {...}
// if we only want to receive values from that channel
func incrementor() chan int {
	out := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			out <- i
		}
		close(out)
	}()
	return out
}

func puller(c chan int) chan int {
	out := make(chan int)
	go func() {
		var sum int
		for n := range c {
			sum += n
		}
		out <- sum
		close(out)
	}()
	return out
}
