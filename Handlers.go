package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

var usr Person // made a global variable for user

func Index(w http.ResponseWriter, r *http.Request) {

	switch r.Method {

	case http.MethodGet:
		fmt.Fprintln(w, `<html><form action = "" method = "post">
		<button>Upvote</button>
	<form></html>`)

	case http.MethodPost:
		http.Redirect(w, r, "http://localhost:8000/user", 2)
	}

}

func User(w http.ResponseWriter, r *http.Request) {
	switch r.Method {

	case http.MethodGet:
		getUser(w, r)

	case http.MethodPost:
		createUser(w, r)

	case http.MethodPut:
		updateUser(w, r)

	case http.MethodDelete:
		deleteUser(w, r)
	}
}

func getUser(w http.ResponseWriter, r *http.Request) {
	userdoc, err := json.Marshal(usr)

	if err == nil {
		fmt.Fprintln(w, string(userdoc))
	}
}
func createUser(w http.ResponseWriter, r *http.Request) {
	fmt.Print("Creating....")

	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&usr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Fprintln(w, "User:", usr.Name, "created!")
}
func updateUser(w http.ResponseWriter, r *http.Request) {

	name := usr.Name

	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&usr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Fprintln(w, "User:", name, "updated to ", usr.Name)
}
func deleteUser(w http.ResponseWriter, r *http.Request) {
	name := usr.Name
	usr = Person{}
	fmt.Fprintln(w, "User", name, " deleted!")
}
