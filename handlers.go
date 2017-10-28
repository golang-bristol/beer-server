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
func GetBeer(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ID, err := strconv.Atoi(ps.ByName("id"))
	if err != nil {
		http.Error(w, fmt.Sprintf("%s is not a valid Beer ID, it must be a number.", ps.ByName("id")), http.StatusBadRequest)
		return
	}

	for _, v := range Cellar {
		if v.ID == ID {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(v)
			return
		}
	}

	http.Error(w, "The beer you requested does not exist.", http.StatusNotFound)
}

// GetBeerReviews returns all reviews for a beer
func GetBeerReviews(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ID, err := strconv.Atoi(ps.ByName("id"))
	if err != nil {
		http.Error(w, fmt.Sprintf("%s is not a valid Beer ID, it must be a number.", ps.ByName("id")), http.StatusBadRequest)
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
	decoder := json.NewDecoder(r.Body)

	var newBeer model.Beer
	err := decoder.Decode(&newBeer)
	if err != nil {
		http.Error(w, "Bad beer - this will be a HTTP status code soon!", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("New beer added.")
}

// AddBeerReview adds a new review for a beer
func AddBeerReview(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")

	ID, err := strconv.Atoi(ps.ByName("id"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(fmt.Sprintf("%s is not a valid Beer ID, it must be a number.", ps.ByName("id")))
		return
	}

	for _, v := range Cellar {
		if v.ID == ID {
			decoder := json.NewDecoder(r.Body)

			var newReview model.Review
			err := decoder.Decode(&newReview)

			newReview.BeerID = ID

			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				json.NewEncoder(w).Encode(http.StatusBadRequest)
				fmt.Println("Bad beer - this will be a HTTP status code soon!")
			} else {
				Reviews = append(Reviews, newReview)
				json.NewEncoder(w).Encode("New beer review added.")
			}
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode("The beer selected for the review does not exist.")
}
