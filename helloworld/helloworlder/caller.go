package helloworlder

import "fmt"

// Call -> exported due to capital letter
func Call() {
	fmt.Println(Version)
	fmt.Println(hello)
}
