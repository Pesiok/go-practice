package main

// uncomment to use!!

func pipeline() {
	// setup pipeline
	// c := gen(2, 3)
	// out := sq(c)

	// out := sq(gen(2, 3, 4))
	// consume the output
	// for number := range sq(gen(2, 3, 4)) {
	// 	fmt.Println(number)
	// }
}

// func gen(numbers ...int) <-chan int {
// 	out := make(chan int)
// 	go func() {
// 		for _, number := range numbers {
// 			out <- number
// 		}
// 		close(out)
// 	}()
// 	return out
// }

// func sq(in <-chan int) <-chan int {
// 	out := make(chan int)
// 	go func() {
// 		for number := range in {
// 			out <- number * number
// 		}
// 		close(out)
// 	}()
// 	return out
// }
