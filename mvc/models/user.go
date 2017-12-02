package models

import (
	"encoding/json"
	"fmt"
	"os"

	"gopkg.in/mgo.v2/bson"
)

// User model
type User struct {
	Name   string        `json:"name" bson:"name"`
	Gender string        `json:"gender" gender:"gender"`
	Age    string        `json:"age" bson:"age"`
	ID     bson.ObjectId `json:"id" bson:"_id"`
}

// Store stores users in file
func StoreUsers(users map[string]User) {
	file, err := os.Create("data")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	json.NewEncoder(file).Encode(users)
}

// Load loads users from file
func LoadUsers() map[string]User {
	users := make(map[string]User)

	file, err := os.Open("data")
	if err != nil {
		fmt.Println(err)
		return users
	}
	defer file.Close()

	err = json.NewDecoder(file).Decode(&users)
	if err != nil {
		fmt.Println(err)
	}
	return users
}
