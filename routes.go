package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func setStaticFolder(r *mux.Router) {
	fs := http.FileServer(http.Dir("./public/"))
	r.PathPrefix("/public/").Handler(http.StripPrefix("/public/", fs))
}

func routes(r *mux.Router) {

	setStaticFolder(r)

	r.HandleFunc("/", renderHome)

	r.HandleFunc("/user", getUsers).Methods("GET")

	r.HandleFunc("/user", insertUser).Methods("POST")

	r.HandleFunc("/user/{id}", deleteUser).Methods("DELETE")

	r.HandleFunc("/user", updateUser).Methods("PUT")

	fmt.Println("Routes are Loded.")
}
