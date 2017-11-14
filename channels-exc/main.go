package main

import "fmt"

func main() {
	// one()
	// two()
	// fmt.Println(<-factorial(3))
	// pipeline()
	// three()
	// runFanIn()
	// runFanInOut()
	// runfact()
	runIncr()
}

func one() {
	// deadlock
	c := make(chan int)
	// launch goroutine
	go func() {
		c <- 1
	}()
	// wait unitl it produce value
	fmt.Println(<-c)
}

func two() {
	c := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		close(c)
	}()

	for value := range c {
		fmt.Println(value)
	}
}

func three() {
	var counter int
	for number := range genFactorials(100, 5) {
		counter++
		fmt.Println(number, counter)
	}
}

func genFactorials(times, number int) <-chan int {
	output := make(chan int)
	done := make(chan bool)
	go func() {
		for i := 0; i < times; i++ {
			go func() {
				output <- <-factorial(number)
				done <- true
			}()
		}
		for i := 0; i < times; i++ {
			<-done
		}
		close(done)
		close(output)
	}()

	return output
}

func factorial(n int) <-chan int {
	output := make(chan int)
	go func() {
		if n == 1 {
			output <- 1
		} else {
			output <- n * <-factorial(n-1)
		}
	}()
	return output
}
