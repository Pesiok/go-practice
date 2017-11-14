package main

import (
	"fmt"
)

func manyTimesOne() {
	n := 10
	c := make(chan int)
	done := make(chan bool)

	fmt.Println("ManyTimesToOne")

	for i := 0; i < n; i++ {
		// launched n goroutines
		go func() {
			var soFar int
			for i := 0; i < 10; i++ {
				soFar += i
				c <- soFar
			}
			done <- true
		}()
	}

	go func() {
		// check for done n times
		for i := 0; i < n; i++ {
			<-done
		}
		close(c)
	}()

	for n := range c {
		fmt.Println(n)
	}

}
