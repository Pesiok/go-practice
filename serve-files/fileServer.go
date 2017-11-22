package main

import (
	"io"
	"net/http"
	"os"
)

func fileServer() {
	http.Handle("/", http.FileServer(http.Dir(".")))
	http.HandleFunc("/fjords", fjords)
	http.ListenAndServe(":8080", nil)

	// setting static folder to assets
	// http.Handle("/resources/", http.StripPrefix("/resources", http.FileServer(http.Dir("./assets"))))

	// listenAndServe returns an error
	// log.Fatal(http.ListenAndServe(":8080", nil))

	// not found handler
	// http.Handle("/favicon.ico", http.NotFoundHandler())
}

func fjords(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "text/html; charset=utf-8")

	io.WriteString(res, `<img src="/fjords.jpg">`)
}

///

func manualFileServer() {
	http.HandleFunc("/fjords", fjords)
	http.HandleFunc("/fjords.jpg", fjordsPicThree)
	http.ListenAndServe(":8080", nil)
}

// 1

func fjordsPic(res http.ResponseWriter, req *http.Request) {
	file, err := os.Open("fjords.jpg")
	if err != nil {
		http.Error(res, "file not found", 404)
		return
	}
	defer file.Close()

	io.Copy(res, file)
}

// 2

func fjordsPicTwo(res http.ResponseWriter, req *http.Request) {
	file, err := os.Open("fjords.jpg")
	if err != nil {
		http.Error(res, "file not found", 404)
		return
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		http.Error(res, "file not found", 404)
		return
	}

	http.ServeContent(res, req, file.Name(), fileInfo.ModTime(), file)
}

// 3

func fjordsPicThree(res http.ResponseWriter, req *http.Request) {
	http.ServeFile(res, req, "fjords.jpg")
}
