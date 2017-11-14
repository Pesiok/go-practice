package main

import (
	"fmt"
	"strconv"
)

var counter int64

func runIncr() {
	c := incr(2)
	for n := range c {
		counter++
		fmt.Println(n)
	}
	fmt.Println("Final: ", counter)
}

func incr(n int) chan string {
	out := make(chan string)
	done := make(chan bool)

	for i := 0; i < n; i++ {
		// launch n goroutines that will spin the /\ range for loop
		go func(i int) {
			for k := 0; k < 20; k++ {
				out <- fmt.Sprint("Process: "+strconv.Itoa(i)+" printing:", k)
			}
			done <- true
		}(i)
	}

	// allows to exit main
	go func() {
		for i := 0; i < n; i++ {
			// if n done's were received then close out channel
			<-done
		}
		close(out)
		close(done)
	}()

	return out
}
