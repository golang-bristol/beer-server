package main

import (
	"fmt"
	"log"
	"net/http"

	model "github.com/golang-bristol/beer-model"
	"github.com/julienschmidt/httprouter"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var (
	// Cellar is the collection of beers
	Cellar []model.Beer

	// Reviews is the colletion of beer reviews
	Reviews []model.Review
)

func dbInitialize(db *gorm.DB) {
	if !db.HasTable(&model.Beer{}) {
		db.CreateTable(model.Beer{})
	}
	if !db.HasTable(&model.Review{}) {
		db.CreateTable(&model.Review{})
	}
}

func main() {
	db, _ := gorm.Open("sqlite3", "main.db")
	defer db.Close()

	dbInitialize(db)

	router := httprouter.New()

	router.GET("/beers", GetBeers(db))
	router.GET("/beers/:id", GetBeer)
	router.GET("/beers/:id/reviews", GetBeerReviews)

	router.POST("/beers", AddBeer(db))
	router.POST("/beers/:id/reviews", AddBeerReview)

	fmt.Println("The beer server is on tap now.")
	log.Fatal(http.ListenAndServe(":8080", router))
}
