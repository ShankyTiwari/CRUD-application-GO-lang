package main

import (
	"net/http"
)

func main() {
	connectDatabse()
	routes()
	http.ListenAndServe(":1323", nil)
}
