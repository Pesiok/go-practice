package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// fanOut: multiple funcs reading from channel until it's closed
// func one() { <-chan } func two() { <-chan}

// fanIn: a function that takes multiple channels  and writing to the one channel
// func one() { chan <- sth } func two() {chan <- sth }

func runFanIn() {
	c := fanIn(boring("Joe"), boring("Ann"))
	for i := 0; i < 10; i++ {
		fmt.Println(<-c)
	}
	fmt.Println("You're both boring; I'm leaving.")
}

func boring(msg string) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()
	return c
}

func fanIn(input1, input2 <-chan string) <-chan string {
	c := make(chan string)
	go func() {
		for {
			c <- <-input1
		}
	}()

	go func() {
		for {
			c <- <-input2
		}
	}()
	return c
}

///////

func runFanInOut() {
	in := gen(2, 3)
	// FAN OUT: multiple funcs reading from channel until it's closed
	// Distribute the sq work across two goroutines that both read from in
	c1 := sq(in)
	c2 := sq(in)

	// FAN IN: multiple channels writing to the same channel
	// Consume the merged output from multiple channels
	for n := range merge(c1, c2) {
		fmt.Println(n)
	}
}

func gen(numbers ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, number := range numbers {
			out <- number
		}
		close(out)
	}()
	return out
}
func sq(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for number := range in {
			out <- number * number
		}
		close(out)
	}()
	return out
}

func merge(cs ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	out := make(chan int)

	// Start an output goroutine for each input channel in cs
	// output copies values from c to out until c is close, then calls wg.Done
	wg.Add((len(cs)))
	for _, c := range cs {
		go func(c <-chan int) {
			for n := range c {
				out <- n
			}
			wg.Done()
		}(c)
	}

	// Start a goroutine to close out
	// once all the output goroutines are done (wg.Done())
	// This must start after thr wg.Add call.
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}
