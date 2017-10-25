package main

import (
	"fmt"
	"log"
	"net/http"

	model "github.com/golang-bristol/beer-model"
	"github.com/julienschmidt/httprouter"
)

var (
	// Cellar is the collection of beers
	Cellar []model.Beer

	// Reviews is the colletion of beer reviews
	Reviews []model.Review
)

func main() {
	PopulateBeers()
	PopulateReviews()

	router := httprouter.New()

	router.GET("/beers", GetBeers)
	router.GET("/beers/:id", GetBeer)
	router.GET("/beers/:id/reviews", GetBeerReviews)

	router.POST("/beers", AddBeer)
	router.POST("/beers/:id/reviews", AddBeerReview)

	fmt.Println("The beer server is on tap now.")
	log.Fatal(http.ListenAndServe(":8080", router))
}
