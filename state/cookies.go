package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
)

func cookies() {
	http.HandleFunc("/", setCookie)
	// http.HandleFunc("/read", readCookie)
	// http.HandleFunc("/del", deleteCookie)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func setCookie(res http.ResponseWriter, req *http.Request) {
	// http.SetCookie(res, &http.Cookie{
	// 	Name:  "tracker",
	// 	Value: "some value",
	// })
	// fmt.Fprintln(res, "cookie written")

	cookie, err := req.Cookie("tracker")
	if err == http.ErrNoCookie {
		cookie = &http.Cookie{
			Name:  "tracker",
			Value: "0",
		}
	}

	count, err := strconv.Atoi(cookie.Value)
	if err != nil {
		log.Fatalln(err)
	}
	count++
	cookie.Value = strconv.Itoa(count)

	http.SetCookie(res, cookie)

	io.WriteString(res, cookie.Value)
}

func readCookie(res http.ResponseWriter, req *http.Request) {
	cookie, err := req.Cookie("tracker")
	if err != nil {
		http.Error(res, err.Error(), http.StatusNoContent)
		return
	}

	fmt.Fprintln(res, "cookie: ", cookie)
}

func deleteCookie(res http.ResponseWriter, req *http.Request) {
	cookie, err := req.Cookie("tracker")
	if err != nil {
		http.Redirect(res, req, "/", http.StatusSeeOther)
		return
	}
	// delete
	cookie.MaxAge = -1
	http.SetCookie(res, cookie)
	http.Redirect(res, req, "/", http.StatusSeeOther)
}
