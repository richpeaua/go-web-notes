package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func hello(w http.ResponseWriter, req *http.Request) {
	
	fmt.Fprint(w, "hello")
}

// newRouter returns an http request router
func newRouter() *mux.Router {
	
	// New router instantiated
	r := mux.NewRouter()

	// Handler for GET requests to /hello
	r.HandleFunc("/hello", hello).Methods("GET")

	// Declare static file directory
	staticFileDir := http.Dir("./assets/")

	// Strip prefix and create File Server handler
	staticFileHandler := http.StripPrefix("/assets/", http.FileServer(staticFileDir))

	r.PathPrefix("/assets/").Handler(staticFileHandler).Methods("GET")
	return r
}

func main() {

	// New router created calling the "newRouter" constructor function
	r := newRouter()

	// Server
	http.ListenAndServe(":8090", r)
}

