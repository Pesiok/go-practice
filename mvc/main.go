package main

import (
	"net/http"

	"github.com/Pesiok/go-practice/mvc/controllers"
	"github.com/julienschmidt/httprouter"
	mgo "gopkg.in/mgo.v2"
)

const dbName = "go-practice"

func main() {
	r := httprouter.New()
	uc := controllers.NewUserController(getMongoSession(), dbName)
	r.GET("/users/:id", uc.GetUser)
	r.POST("/user", uc.CreateUser)
	r.DELETE("/user/:id", uc.DeleteUser)
	http.ListenAndServe("localhost:8080", r)
}

func getMongoSession() *mgo.Session {
	session, err := mgo.Dial("mongodb://localhost:27017")
	if err != nil {
		panic(err)
	}
	return session
}
