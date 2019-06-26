package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// Users represents a collection of users
type Users struct {
	Users []User `json:"users"`
}

// User is a person who uses our services
type User struct {
	Name   string `json:"name"`
	Type   string `json:"type"`
	Age    int64  `json:"age"`
	Social Social `json:"social"`
}

// Social is the soclia media info of our users
type Social struct {
	Facebook string `json:"facebook"`
	Twitter  string `json:"twitter"`
}

func main() {
	jsonFile, err := os.Open("./test.json")
	if err != nil {
		fmt.Println(err)
	}

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var users Users

	json.Unmarshal(byteValue, &users)

	for _, user := range users.Users {
		fmt.Printf("This is %v\n", user.Name)
	}
}
