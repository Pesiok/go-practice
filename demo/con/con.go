package con

import (
	"fmt"
)

// Animal is a living thing that makes sounds
type Animal interface {
	MakeSound() string
}

// Iterate is an async iterator
func Iterate(animals ...Animal) {
	pipe := sink(animals...)
	for sound := range sounds(pipe) {
		fmt.Println(sound)
	}
}

func sink(animals ...Animal) <-chan Animal {
	out := make(chan Animal)
	go func() {
		for _, animal := range animals {
			out <- animal
		}
		close(out)
	}()
	return out
}

func sounds(in <-chan Animal) <-chan string {
	out := make(chan string)
	go func() {
		for animal := range in {
			out <- animal.MakeSound()
		}
		close(out)
	}()
	return out
}
