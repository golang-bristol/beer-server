package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/golang-bristol/beer-model"
)

// GetBeers ... Returns the cellar.
func GetBeers(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Cellar)
}

// AddBeer ... Add a new beer to the cellar.
func AddBeer(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")

	decoder := json.NewDecoder(r.Body)

	var newBeer model.Beer
	err := decoder.Decode(&newBeer)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(http.StatusBadRequest)
		fmt.Println("Bad beer - this will be a HTTP status code soon!")
	} else {
		json.NewEncoder(w).Encode("New beer added.")
	}
}
