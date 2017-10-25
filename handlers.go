package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/golang-bristol/beer-model"
	"github.com/julienschmidt/httprouter"
)

// GetBeers returns the cellar
func GetBeers(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Cellar)
}

// GetBeer returns a beer from the cellar
func GetBeer(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// TODO
}

// GetBeerReviews returns all reviews for a beer
func GetBeerReviews(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ID, err := strconv.Atoi(ps.ByName("id"))
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	// TODO: Consider checking if a beer matching the ID actually exists, and
	// 404 if that is not the case.

	results := []model.Review{}
	for _, v := range Reviews {
		if v.BeerID == ID {
			results = append(results, v)
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(results)
}

// AddBeer adds a new beer to the cellar
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

// AddBeerReview adds a new review for a beer
func AddBeerReview(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// TODO
}
