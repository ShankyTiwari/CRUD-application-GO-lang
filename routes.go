package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func setStaticFolder(route *mux.Router) {
	fs := http.FileServer(http.Dir("./public/"))
	route.PathPrefix("/public/").Handler(http.StripPrefix("/public/", fs))
}

func addApproutes(route *mux.Router) {

	setStaticFolder(route)

	route.HandleFunc("/", renderHome)

	route.HandleFunc("/user", getUsers).Methods("GET")

	route.HandleFunc("/user", insertUser).Methods("POST")

	route.HandleFunc("/user/{id}", deleteUser).Methods("DELETE")

	route.HandleFunc("/user", updateUser).Methods("PUT")

	fmt.Println("Routes are Loded.")
}
