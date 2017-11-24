package main

import (
	"fmt"
	"html/template"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func redirects() {
	http.HandleFunc("/", fooThree)
	http.HandleFunc("/bar", barOne)
	http.HandleFunc("/barred", barred)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func fooThree(res http.ResponseWriter, req *http.Request) {
	fmt.Print("Request method at foo: ", req.Method, "\n\n")
}

func barOne(res http.ResponseWriter, req *http.Request) {
	fmt.Print("Request method at bar: ", req.Method, "\n\n")
	// ~ process form

	// 303
	// redirect: See Other (POST => GET)
	// res.Header().Set("Location", "/")
	// res.WriteHeader(http.StatusSeeOther)

	// 307
	// redirect: temporary (POST => POST)
	// http.Redirect(res, req, "/", http.StatusTemporaryRedirect)

	// 302
	// moved permanently
	// broweser will automaticaly redirect to moved resource
	http.Redirect(res, req, "/", http.StatusMovedPermanently)

}

func barred(res http.ResponseWriter, req *http.Request) {
	fmt.Print("Your request method at barred: ", req.Method, "\n\n")
	tpl.ExecuteTemplate(res, "index.gohtml", nil)
}
