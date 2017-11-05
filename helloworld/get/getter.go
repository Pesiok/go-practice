package get

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// Google -> get google page
func Google() {
	// blank identifier -> used if we do not want to use given variable for now
	// res, _ := http.Get('www')
	res, err := http.Get("http://www.google.com/")
	if err != nil {
		log.Fatal(err)
	}
	page, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", page)
}
