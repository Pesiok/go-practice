package main

import (
	"net/http"
)

type animal interface {
	makeSound() string
}

type human struct {
	height int
	Race   string `json:"race"`
}

// Person is a type representing someone
type Person struct {
	human
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	age       int
}

type people []Person

func (p people) Len() int           { return len(p) }
func (p people) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p people) Less(i, j int) bool { return p[i].height < p[j].height }

http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
	// user := Person{}
	// jsonUser, _ := json.Marshal(user)
})

// http.ListenAndServe(":80", nil)
