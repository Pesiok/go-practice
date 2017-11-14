package main

import (
	"fmt"
	"sync"
)

func manyToOne() {
	c := make(chan int)

	var wg sync.WaitGroup
	wg.Add(2)

	fmt.Println("Many to one")

	go func() {
		// wg.Add(1) creates race condition here
		for i := 0; i < 10; i++ {
			c <- i
		}
		wg.Done()
	}()

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		wg.Done()
	}()

	go func() {
		wg.Wait()
		close(c)
	}()

	for n := range c {
		fmt.Println(n)
	}

}
