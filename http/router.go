package main

import (
	"io"
	"net/http"
)

type one int

func (h one) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "one")
}

type two int

func (h two) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "two")
}

func startRouterOne() {
	// var o one
	// var t two

	// custom serve mux
	// router := http.NewServeMux()
	// router.Handle("/one/", o)
	// router.Handle("/two", t)
	// http.ListenAndServe(":8080", router)

	// default serve mux
	// http.Handle("/one/", o)
	// http.Handle("/two", t)
	// http.ListenAndServe(":8080", nil)

	// or handle func

	http.HandleFunc("/one", func(res http.ResponseWriter, req *http.Request) {
		io.WriteString(res, "one")
	})

	http.HandleFunc("/two", func(res http.ResponseWriter, req *http.Request) {
		io.WriteString(res, "two")
	})

	http.ListenAndServe(":8080", nil)
}
