package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"gopkg.in/mgo.v2/bson"

	"github.com/Pesiok/go-practice/mvc/models"
	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
)

// UserController for user models
type UserController struct {
	memory  map[string]models.User
	session *mgo.Session
	users   *mgo.Collection
}

// NewUserController creates a new user controller instance
func NewUserController(session *mgo.Session, dbName string) *UserController {
	return &UserController{
		session: session,
		users:   session.DB(dbName).C("users"),
		memory:  models.LoadUsers(),
	}
}

func (uc UserController) GetUser(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
	// get id from params
	id := params.ByName("id")

	// validate id
	if !bson.IsObjectIdHex(id) {
		res.WriteHeader(http.StatusNotFound)
		fmt.Println("not found")
		return
	}

	// get objectId
	objId := bson.ObjectIdHex(id)

	// get user from db
	user := models.User{}
	err := uc.users.FindId(objId).One(&user)
	if err != nil {
		fmt.Println(err)
		res.WriteHeader(http.StatusNotFound)
		return
	}
	// get user from memory
	// user := uc.memory[id]

	jsonUser, _ := json.Marshal(user)

	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)
	fmt.Fprintf(res, "%s\n", jsonUser)
}

func (uc UserController) CreateUser(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	user := models.User{}

	// decode json from stream
	json.NewDecoder(req.Body).Decode(&user)

	// new bson ID
	user.ID = bson.NewObjectId()

	// store user in mongodb
	uc.users.Insert(user)

	// store user in memory
	// uc.memory[string(user.ID)] = user
	// models.StoreUsers(uc.memory)

	// marshal do json
	jsonUser, _ := json.Marshal(user)

	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusCreated)
	fmt.Fprintf(res, "%s\n", jsonUser)
}

func (uc UserController) DeleteUser(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
	// get id from params
	id := params.ByName("id")

	// validate id
	if !bson.IsObjectIdHex(id) {
		res.WriteHeader(http.StatusNotFound)
		return
	}

	// get objectId
	objId := bson.ObjectIdHex(id)

	// delete from db
	err := uc.users.RemoveId(objId)
	if err != nil {
		res.WriteHeader(http.StatusNotFound)
		return
	}
	// delete from memory
	// delete(uc.memory, string(objId))
	// models.StoreUsers(uc.memory)

	res.WriteHeader(http.StatusOK)
	fmt.Fprint(res, "Deleted user", objId, "\n")
}
