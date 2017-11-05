package memory

import "fmt"

// Mem - show memory addresses
func Mem() {
	a := 34
	fmt.Println("a: ", a)
	fmt.Println("a's memory address: ", &a)
	fmt.Printf("%d \n", &a)
}

func MetersToYards() {
	const constant float64 = 1.09361
	var meters float64

	fmt.Print("Enters meters: ")
	fmt.Scan(&meters)
	fmt.Println(meters, " meters is ", meters*constant, " yards.")
}

func Pointers() {
	// everything in go is passed by value:
	// every time new copy with different memory adress is created
	// when passing pointer new copy of pointer to the same memory address is created

	// we can pass around memory adresses instead of values
	// to make program more performant

	a := 43

	// value -> 43
	fmt.Println(a)

	// address of that value in memory -> 0xc04204e298
	fmt.Println(&a)

	// pointer
	// points to &a memory address
	b := &a

	// -> 0xc04204e298
	fmt.Println(b)

	// get a value from a pointer, dereferencing
	// -> 43
	fmt.Println(*b)

	// change value at which b is pointing at
	*b = 42

	fmt.Println(a)

}

// standard pass bya value

// func Usingpointers() {
// 	x := 5
// 	zero(x)
// 	fmt.Println(x)
// }

// func zero(x int) {
// 	x = 0
// }

// 'pass by reference' with pointers

func Usingpointers() {
	x := 5
	fmt.Println(&x)
	zero(&x)
	// now x is changed
	fmt.Println(x)
}

func zero(x *int) {
	// changing value at x pointer
	*x = 0
}
