package main

import "fmt"

func main() {

	// loops
	for i := 4; i < 19; i++ {
		if i&2 == 0 {
			// stop this cycle and go to next one
			continue
		}
		fmt.Println(i)
	}

	for 1 > 5 {
		// never
	}

	// for {
	// 	// forever
	// 	if 5 == 1 {
	// 		// stop whole loop
	// 		break
	// 	}
	// }

	// runes - characters, int32 alias

	// conversion
	// string - collection of runes
	// convert to slice of bytes
	// []byte("Hello")

	foo := 'a'
	fmt.Println(foo)
	fmt.Printf("%T \n", foo)
	fmt.Printf("%T \n", rune(foo))

	// strings
	bar := "strings are made of int32 runes coded in utf8 (4 byte coding sheme)"
	literal := `literals
		maitain space
	` + string(foo) + ` even more`

	fmt.Println(bar, literal)

	// switches
	// if no value provided then run all true cases

	switch "value" {
	case "someone":
		// do sth
	case "medhi", "or this":
		{
			// do sth else
			// more
		}
		// and do all below
		fallthrough
	case "kek":
		{

		}
	default:
		// do nothing
	}
	b := false

	// initialization statement
	if food := "kek"; b {
		fmt.Println(food)
	}

}

type Contact struct {
	name  string
	email string
}
