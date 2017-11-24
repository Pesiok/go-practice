package main

import (
	"errors"
	"net/http"
)

// type user struct {
// 	UserName string
// 	First    string
// 	Last     string
// }

// var tpl *template.Template
// var dbUsers = map[string]user{}
// var dbSessions = map[string]string{}

// func init() {
// 	tpl = template.Must(template.ParseGlob("templates/*"))
// }

// func gosession() {
// 	http.HandleFunc("/", fooSes)
// 	http.HandleFunc("/bar", barSes)
// 	http.ListenAndServe(":8080", nil)
// }

// func fooSes(res http.ResponseWriter, req *http.Request) {

// 	// get cookie
// 	cookie, err := req.Cookie("session")
// 	if err != nil {
// 		sID := uuid.NewV4()
// 		cookie = &http.Cookie{
// 			Name:  "session",
// 			Value: sID.String(),
// 		}
// 		http.SetCookie(res, cookie)
// 	}

// 	// if user exists already, get user based on value in the cookie
// 	var u user
// 	if un, ok := dbSessions[cookie.Value]; ok {
// 		u = dbUsers[un]
// 	}

// 	// process form submission
// 	// set user data
// 	if req.Method == http.MethodPost {
// 		un := req.FormValue("username")
// 		f := req.FormValue("firstname")
// 		l := req.FormValue("lastname")
// 		u = user{un, f, l}
// 		dbSessions[cookie.Value] = un
// 		dbUsers[un] = u
// 	}

// 	tpl.ExecuteTemplate(res, "index.gohtml", u)
// }

// func barSes(res http.ResponseWriter, req *http.Request) {
// 	// get cookie
// 	cookie, err := req.Cookie("session")
// 	if err != nil {
// 		http.Redirect(res, req, "/", http.StatusSeeOther)
// 	}
// 	// get user based on value in cookie
// 	un, ok := dbSessions[cookie.Value]
// 	if !ok {
// 		http.Redirect(res, req, "/", http.StatusSeeOther)
// 	}
// 	u := dbUsers[un]
// 	tpl.ExecuteTemplate(res, "bar.gohtml", u)
// }

func getUser(w http.ResponseWriter, req *http.Request) (user, error) {
	var u user
	// get cookie
	c, err := req.Cookie("session")
	if err != nil {
		return u, errors.New("there is no cookie")
	}
	http.SetCookie(w, c)

	// if the user exists already, get user
	if s, ok := dbSessions[c.Value]; ok {
		u = dbUsers[s.un]
	}
	return u, nil
}

func alreadyLoggedIn(w http.ResponseWriter, req *http.Request) bool {
	c, err := req.Cookie("session")
	if err != nil {
		return false
	}
	s := dbSessions[c.Value]
	_, ok := dbUsers[s.un]
	// refresh session
	c.MaxAge = sessionLength
	http.SetCookie(w, c)
	return ok
}
