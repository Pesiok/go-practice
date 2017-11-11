package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func main() {
	//////// 1. arrays //////////

	// ordered
	// static
	// does not change size
	// elements identified by an index
	// single type
	// one-dimensional, but could have nested arrays
	var arr [58]string
	// array length
	fmt.Println(len(arr))
	fmt.Println(arr[0])
	// ASCI (numbers) to string
	for i := 65; i <= 122; i++ {
		arr[i-65] = string(i)
	}
	fmt.Println(arr[0])

	//////// 2. slices /////////

	// reference type that points to the underlying array
	// have: pointer, lenght, capacity
	// uninitialized -> nil
	slice := []int{1, 4, 5, 6}
	// slicing a slice
	// returns everything from position 2 to up to 4 (without it)
	fmt.Println(slice[2:4])
	// string -> slice of runes (made of bytes) -> slice of bytes
	fmt.Println("string"[0])
	// -----
	// slice := make([]T, length - starting size of slice, capacity - length of array )
	// when capacity is exceeded a new underlying array is created with 2x capacity
	// and old one deleted
	// -----
	// make([]int, 50, 100)
	// new([100]int)[0:50]
	// or []int{1, 4, 6} -> capacity 3
	appendSlice := func() {
		slice := make([]int, 0, 5)

		fmt.Println("----------")
		fmt.Println(slice)
		fmt.Println(len(slice))
		fmt.Println(cap(slice))
		fmt.Println("----------")

		for i := 0; i < 80; i++ {
			// adding new elements to slice
			slice = append(slice, i)
			fmt.Println("Len: ", len(slice), "Capacity: ", cap(slice), "value: ", slice[i])
			// merging slices
			// slice = append(slice, otherSlice...)
		}
		// deleting from slice
		// from 2 (without 2) to 3(including 3)
		slice = append(slice[:2], slice[3:]...)
	}
	appendSlice()
	// increment
	//slice[0]++

	////// 3. Maps //////

	// key - value pairs
	// dictionary
	// unordered
	// reference type
	// uninitialized -> nil
	// a bit like JS's {}
	// built upon hashtable, which are built on arrays
	// ------
	// empty (initialized) map with string key and int value
	myMapThree := map[string]string{}
	myMap := make(map[string]int)
	// ------
	// non-empty map
	myMapTwo := map[string]int{"foo": 1, "bar": 2}
	myMap["kek"] = 1
	myMap["kak"] = 2
	myMap["kik"] = 3
	// delete
	delete(myMap, "kak")
	// get
	if value, ok := myMap["kek"]; ok {
		fmt.Println(value, ok)
	}
	// ------
	// loop
	for key, value := range myMap {
		fmt.Println(key, " - ", value)
	}
	fmt.Println("mapOne: ", myMap)
	fmt.Println("mapTwo", myMapTwo)
	fmt.Println("mapThree", myMapThree)

	structs()
	jsonThings()
}

////// 4. Structs //////

// OOP in (Go)
// 1. encapsulation: state (struct fields), behavior (methods), public/private (exports)
// 2. reuseability: inheritence (embedded types)
// 3. polymorhpism: interfaces
// 4. overriding: (promotion)
// aggregate type
// a bit like class/JS's constructor function

// used declared types
type Person struct {
	// fields & tags
	first string
	last  string `datastore:"a,noindex"`
	age   int
}

type SuperPerson struct {
	Person
	// overriding
	first   string
	license bool
}

func structs() {
	// object like
	p1 := Person{"James", "Bond", 34}
	p2 := Person{"Milla", "Kek", 333}
	p3 := SuperPerson{p2, "Ayy", true}
	p4 := SuperPerson{
		Person: Person{
			first: "Ayy",
			last:  "Lmao",
			age:   43,
		},
		license: true,
	}
	fmt.Println(p1.first, p2.first)
	fmt.Println(p3)
	fmt.Println(p4)
	// methods
	fmt.Println(p3.fullName())
	fmt.Println(p3.fullNameSuperPerson())

	// pointers to struct
	p5 := &Person{"Ayy", "Lmao", 0}
	fmt.Println(p5)
	// fmt.Printf("%T\n", p5)
	// getting values straight out of the pointer
	fmt.Println(p5.first)
}

// 'free' function, can work with on any Person type
func (p Person) fullName() string {
	// p is like `this`
	// gettin values only from initial Person, not overidden by new SuperPerson object
	// if method will be modifying the reciever then it should ba a pointer
	return p.first + " " + p.last
}

func (p SuperPerson) fullNameSuperPerson() string {
	// p is like `this`
	return p.first + " " + p.last
}

func jsonThings() {
	// big letters are exported
	type myStruct struct {
		First string
		// do not include this
		Last string `json:"-"`
		// alt
		Age         int `json:"yearsLived"`
		notExported int
	}
	myPerson := myStruct{"Ayy", "Lmao", 4, 5}

	// Marshal/Unmarshal -> convert static data

	bs, _ := json.Marshal(myPerson)
	fmt.Println(bs)
	fmt.Printf("%T \n", bs)
	fmt.Println(string(bs))

	var mySecondPerson myStruct
	bs2 := []byte(`{"First": "James", "Last": "Bond", "yearsLived": 23}`)
	json.Unmarshal(bs2, &mySecondPerson)
	fmt.Println(mySecondPerson)

	// Decode/Encode -> convert data from streams
	json.NewEncoder(os.Stdout).Encode(myPerson)
	var myThirdPerson myStruct
	reader := strings.NewReader(`{"First": "James", "Last": "Bond", "yearsLived": 23}`)
	json.NewDecoder(reader).Decode(&myThirdPerson)
	fmt.Println(myThirdPerson)

}
