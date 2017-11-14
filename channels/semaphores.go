package main

import (
	"fmt"
)

func semaphores() {
	c := make(chan int)
	done := make(chan bool)

	fmt.Println("Semaphores")

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		done <- true
	}()

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		done <- true
	}()

	go func() {
		// takes done from channel if it is accessible & continues
		<-done
		<-done
		close(c)
	}()

	// without /\ goroutine we would block our thread
	// because our program never reaches `range c` and we need receiver
	// if there is no receiver then our goroutines will never finish the loops
	// and won't put true to done channel
	for n := range c {
		fmt.Println(n)
	}

}
