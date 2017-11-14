package main

import (
	"fmt"
	"sync"
)

func runfact() {
	in := gen2()

	cs := fanOut2(in, 10)

	for n := range merge2(cs...) {
		fmt.Println(n)
	}
}

// generate values passed to factorial
func gen2() <-chan int {
	out := make(chan int)
	go func() {
		for i := 0; i < 1000; i++ {
			for j := 3; j < 13; j++ {
				out <- j
			}
		}
		close(out)
	}()
	return out
}

func factorial2(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- func(n int) int {
				total := 1
				for i := n; i > 0; i-- {
					total *= i
				}
				return total
			}(n)
		}
		close(out)
	}()
	return out
}

// FAN OUT
// multiple funcs (factorial) reading from channel until it's closed
func fanOut2(in <-chan int, n int) []<-chan int {
	// cs := make([]<-chan int, 0, 10)
	// cs := make([]<-chan int, n)
	var cs []<-chan int
	for i := 0; i < n; i++ {
		cs = append(cs, factorial2(in))
	}
	return cs
}

// FAN IN
// multiple channels writing to the same channel (out)
func merge2(cs ...<-chan int) <-chan int {
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
