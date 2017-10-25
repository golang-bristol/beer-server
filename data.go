package main

import (
	"time"

	"github.com/golang-bristol/beer-model"
)

// PopulateBeers populates the Cellar variable with Beers
func PopulateBeers() {
	Cellar = []model.Beer{
		model.Beer{ID: 1, Name: "Beer 1"},
		model.Beer{ID: 2, Name: "Beer 2"},
	}
}

// PopulateReviews populates the Reviews variable with Reviews
func PopulateReviews() {
	Reviews = []model.Review{
		model.Review{ID: 1, BeerID: 1, FirstName: "Joe", LastName: "Tribiani", Score: 5, Text: "This is good but this is not pizza!", Created: time.Date(2017, time.November, 10, 12, 36, 0, 0, time.UTC)},
	}
}
