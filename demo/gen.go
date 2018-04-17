package main

import (
	c "github.com/grzegorz-bielski/go-practice/demo/con"
)

type dog []int

func (d dog) MakeSound() string {
	return "woof"
}

func async() {
	p := newPerson(22, 123, "Ayy", "Lmao")
	d := dog{}
	c.Iterate(p, d)
}
