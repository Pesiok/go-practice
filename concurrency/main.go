package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

var wg sync.WaitGroup
var counter int64
var mutex sync.Mutex

// special function
// runs before main
func init() {
	// use all cores
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	// no concurrency

	// foo()
	// bar()

	// concurrency

	// composition of independently executing processes
	// _dealing_ with lots of things at once
	// way to structure program
	// many things but one at time
	// can have concurrency without parallelism
	// here switching between 3 threads: main, foo, bar
	// wg.Add(2)
	// go foo()
	// go bar()
	// wg.Wait()

	// parallelism

	// simultaneous execution of (possibly related) computations
	// running things in parallel
	// can't have parallelism without concurrency
	// _doing_ a lots of things at once

	// race conditions

	// check for race: go run -race file.go
	wg.Add(2)
	go incrementor("Foo:")
	go incrementor("Bar:")
	wg.Wait()
	fmt.Println("Final Counter:", counter)

	// mutex
	// mutual exclusion object

	// atomics
}

func foo() {
	fmt.Println("-----")
	for i := 0; i < 5; i++ {
		fmt.Println("Foo:", i)
		time.Sleep(time.Duration(3 * time.Millisecond))
	}
	wg.Done()
}

func bar() {
	fmt.Println("-----")
	for i := 0; i < 5; i++ {
		fmt.Println("Bar:", i)
		time.Sleep(time.Duration(30 * time.Millisecond))
	}
	wg.Done()
}

// race conditions - many processes are manipulating the same value
// one can overwrite another

func incrementor(s string) {
	for i := 0; i < 20; i++ {
		// delay
		time.Sleep(time.Duration(rand.Intn(3)) * time.Millisecond)

		// no race, stops other threads from accessing
		{
			mutex.Lock()
			counter++
			fmt.Println(s, i, "Counter:", counter)
			mutex.Unlock()
		}

		// also no race
		// atomic.AddInt64(&counter, 1)
		// but addin print gives race condition, but atomic doesnt
		// fmt.Println(s, i, "Counter:", counter)
	}
	wg.Done()
}
