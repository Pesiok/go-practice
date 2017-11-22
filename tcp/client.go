package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
)

// won't work from the same package as server btw
func client() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	// read from connection
	bs, err := ioutil.ReadAll(conn)
	// write to connection
	// fmt.Fprintln(conn, "msg")
	if err != nil {
		log.Println(err)
	}

	fmt.Println(string(bs))
}
