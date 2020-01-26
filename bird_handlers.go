package main

import (
	"net/http"
	"encoding/json"
	"fmt"
)

type Bird struct {
	Species	string `json:"species"`
	Description string `json:"description"`
}

var birds []Bird

func getBirdHandler(w http.ResponseWriter, r *http.Request) {
	
	// Convert "birds" variable into JSON
	birdListBytes, err := json.Marshal(birds)

	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	// If all goes well, write the JSON list of birds to the response
	w.Write(birdListBytes)

}

func createBirdHandler(w http.ResponseWriter, r *http.Request) {
	
	// instantiate a new bird object
	bird := Bird{}

	// All data comes from HTML form data
	// 'ParseForm' method of the request parses the
	// form values
	err := r.ParseForm()
	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
	}

	// Get bird info from form
	bird.Species 	= r.Form.Get("species")
	bird.Description  = r.Form.Get("description")

	// Append bird to existing birds list
	birds = append(birds, bird)

	// Finally redirect user to original HTML page
	http.Redirect(w, r, "/assets/", http.StatusFound)
}