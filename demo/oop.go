package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type human struct {
	age    int
	height int
	Race   string `json:"race"`
}

type person struct {
	human
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

func (p *person) live() {
	p.age++
}

func (p person) MakeSound() string {
	return "Hello"
}

func newPerson(age int, height int, firstName string, lastName string) *person {
	return &person{
		human:     human{age, height, "homo sapiens"},
		FirstName: firstName,
		LastName:  lastName,
	}
}

func sendPerson(w http.ResponseWriter, req *http.Request) {
	p := newPerson(34, 134, "Krzysztof", "Krawczyk")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	err := json.NewEncoder(w).Encode(&p)
	if err != nil {
		fmt.Println(err)
	}

}

func oop() {
	http.HandleFunc("/", sendPerson)
	http.ListenAndServe(":8080", nil)
}
