package main

import (
	"time"

	"github.com/golang-bristol/beer-model"
)

// PopulateBeers populates the Cellar variable with Beers
func PopulateBeers() {
	Cellar = []model.Beer{
		model.Beer{ID: 1, Name: "Coors Light"},
		model.Beer{ID: 2, Name: "Brahma"},
		model.Beer{ID: 3, Name: "Heineken"},
		model.Beer{ID: 4, Name: "Carlsberg"},
		model.Beer{ID: 5, Name: "Skol"},
		model.Beer{ID: 6, Name: "Tsingtao"},
		model.Beer{ID: 7, Name: "Bud Light"},
		model.Beer{ID: 8, Name: "Budweiser"},
		model.Beer{ID: 9, Name: "Corona"},
		model.Beer{ID: 10, Name: "Yanjing"},
	}
}

// PopulateReviews populates the Reviews variable with Reviews
func PopulateReviews() {
	Reviews = []model.Review{
		model.Review{ID: 1, BeerID: 1, FirstName: "Joe", LastName: "Tribiani", Score: 5, Text: "This is good but this is not pizza!", Created: time.Date(2017, time.November, 10, 12, 36, 0, 0, time.UTC)},
	}
}
