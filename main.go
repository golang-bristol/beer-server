package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

var (
	// Cellar ... The collection of beer.
	Cellar Beers
)

func main() {
	populateBeers()

	router := httprouter.New()
	router.GET("/beers", GetBeers)
	router.POST("/beers", AddBeer)
	//router.GET("/hello/:name", Hello)

	fmt.Println("The Beer server is on tap now.")
	log.Fatal(http.ListenAndServe(":8080", router))
}
