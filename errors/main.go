package main

import (
	"fmt"
	"log"
	"os"
)

func init() {
	nf, err := os.Create("log.txt")
	if err != nil {
		fmt.Println(err)
	}

	log.SetOutput(nf)
}

func main() {
	_, err := os.Open("no-file.txt")
	if err != nil {
		// to console
		// fmt.Println("err happend", err)

		// prints to stdout (logger created in init)
		log.Println("err happend", err)

		// calls os.Exit(1) after writing msg
		// log.Fatalln(err)

		// show error stack
		// panic(err)

		//////////////

		_, err := Sqrt(-10)
		if err != nil {
			log.Fatalln(err)
		}
	}

}

type FunException struct {
	lat, long string
	err       error
}

func (n *FunException) Error() string {
	return fmt.Sprintf("Fun exception error: %v %v $v", n.lat, n.long, n.err)
}

func Sqrt(f float64) (float64, error) {
	if f < 0 {
		// custom errors

		// 1
		// return 0, errors.New("you can't do that")

		// 2
		// var ErrType = errors.New("Bad err")
		// return 0, ErrType

		// 3
		// return 0, fmt.Errorf("bad err: %v", f)

		// 4
		// return address to that struct
		return 0, &FunException{"50 N", "99 W", fmt.Errorf("you can't do that: %v", f)}
	}
	// implementation...
	return 42, nil
}
