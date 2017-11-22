package main

import (
	"html/template"
	"log"
	"net/http"
)

// func ServeHttp(res http.ResponseWriter, req *http.Request){
// 	res.Header().Set("Some-Header", "junk")
// 	fmt.Fprintln(res, someHTMLtemplate)
// }

// type Handler interface {
// 	ServerHTTP(ResponseWriter, *Request)
// }

type handler struct{}

func (h handler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if err != nil {
		log.Fatalln(err)
	}

	// data := struct {
	// 	Method        string
	// 	URL           *url.URL
	// 	Submissions   map[string][]string
	// 	Header        http.Header
	// 	ContentLength int64
	// }{
	// 	req.Method,
	// 	req.URL,
	// 	req.Form,
	// 	req.Header,
	// 	req.ContentLength,
	// }

	// fmt.Println(data)

	tpl.ExecuteTemplate(w, "index.gohtml", req.Form)
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("index.gohtml"))
}

func startHandler() {
	var h handler
	http.ListenAndServe(":8080", h)
}
