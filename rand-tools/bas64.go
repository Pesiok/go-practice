package main

import (
	"encoding/base64"
	"fmt"
)

func runBase64() {
	s := "lslslsldfldsflsdlfsdlgsdlgsldglsd"
	s64 := base64.StdEncoding.EncodeToString([]byte(s))

	fmt.Println(len(s))
	fmt.Println(len(s64))
	fmt.Println(s)
	fmt.Println(s64)
}
