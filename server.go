package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	fmt.Println("Server will start at http://localhost:8000/")

	connectDatabse()

	r := mux.NewRouter()

	routes(r)

	log.Fatal(http.ListenAndServe(":8000", r))
}
