package main

import (
	"fmt"
	"net/http"
)

var chttp = http.NewServeMux()

func routes() {
	chttp.Handle("/", http.FileServer(http.Dir("./")))

	http.HandleFunc("/", renderHome)

	http.HandleFunc("/getUsers", getUsers)

	http.HandleFunc("/insertUsers", inertUsers)

	fmt.Println("Routes are Loded.")
}
