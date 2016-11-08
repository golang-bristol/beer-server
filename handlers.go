package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// GetBeers ... Returns the cellar.
func GetBeers(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	beers, err := ListBeers()
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(beers)
}

// AddBeer ... Add a new beer to the cellar.
func AddBeer(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")

	decoder := json.NewDecoder(r.Body)

	var newBeer Beer
	err := decoder.Decode(&newBeer)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(http.StatusBadRequest)
		fmt.Println("Bad beer - this will be a HTTP status code soon!")
		return
	}

	err = newBeer.Persist()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(http.StatusBadRequest)
		fmt.Println("Bad beer - this will be a HTTP status code soon!")
		return
	}

	json.NewEncoder(w).Encode("New beer added.")
}
