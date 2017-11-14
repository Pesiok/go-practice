package main

import "fmt"

func main() {
	// Do not communicate by sharing memory (do not mutex.lock, dont use atomic)
	// Share memory by communicating (use channels to communiciate between goroutines)
	// channel over which we will be communicating int

	// when data is pushed to an (unbuffered) channel
	// the code blocks until the data is taken of by another goroutine from that channel

	// unbuffered channel
	c := make(chan int)

	// buffered channel
	// c : make(chan int, 10)

	go func() {
		for i := 0; i < 10; i++ {
			// putting number to channel
			// everything stops until something takes value from the channel
			c <- i
		}
		// putting values on channel is no longer possible
		close(c)
	}()

	// go func() {
	// 	for {
	// 		// taking number of the channel
	// 		fmt.Println(<-c)
	// 	}
	// }()

	// block until range has finished looping
	for n := range c {
		fmt.Println(n)
	}

	// sleep no longer needed
	// 	time.Sleep(time.Second)

	// manyToOne()
	// semaphores()
	// manyTimesOne()
	// oneToMany()
	passing()
}
