package main

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// JSONError writes a HTTP status code and shows an error message as JSON.
func JSONError(w http.ResponseWriter, status int, err error) {
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(map[string]string{
		"err": err.Error(),
	})
}

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
		JSONError(w, http.StatusBadRequest, err)
		return
	}

	err = newBeer.Persist()
	if err != nil {
		JSONError(w, http.StatusBadRequest, err)
		return
	}

	json.NewEncoder(w).Encode("New beer added.")
}
