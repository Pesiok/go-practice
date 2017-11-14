package main

import (
	"fmt"
)

func oneToMany() {
	n := 10
	c := make(chan int)
	done := make(chan bool)

	fmt.Println("OneToMany")

	go func() {
		for i := 0; i < 100000; i++ {
			// write to channel
			c <- i
		}
		close(c)
	}()

	for i := 0; i < n; i++ {
		go func() {
			// stop
			// listen to channel change
			for n := range c {
				fmt.Println(n)
			}
			done <- true
		}()
	}

	// expect n times done because n goroutines was launched
	for i := 0; i < n; i++ {
		// stop
		// listen to channel change
		<-done
	}

}
