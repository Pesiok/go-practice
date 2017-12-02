package main

import (
	"net/http"
)

func runHttps() {
	http.HandleFunc("/", fooTri)
	// generate cert and key:
	// go run <GOROOT>/src/crypto/tls/generate_cert.go --host=localhost
	http.ListenAndServeTLS(":10443", "cert.pem", "key.pem", nil)
}

func fooTri(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "text/plain")
	res.Write([]byte("halko"))
}
