package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/golang-bristol/beer-model"
	"github.com/jinzhu/gorm"
	"github.com/julienschmidt/httprouter"
)

// GetBeers returns the cellar
func GetBeers(db *gorm.DB) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		w.Header().Set("Content-Type", "application/json")
		var beersFromDB []model.Beer
		db.Find(&beersFromDB)
		json.NewEncoder(w).Encode(beersFromDB)
	}
}

// GetBeer returns a beer from the cellar
func GetBeer(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// TODO
}

// GetBeerReviews returns all reviews for a beer
func GetBeerReviews(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// TODO
}

// AddBeer adds a new beer to the cellar
func AddBeer(db *gorm.DB) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
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
		db.Create(&newBeer)
	}
}

// AddBeerReview adds a new review for a beer
func AddBeerReview(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// TODO
}
