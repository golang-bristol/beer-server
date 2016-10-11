package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

var (
	// Cellar ... The collection of beer.
	Cellar Beers
)

func GetBeers(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Cellar)
}

func populateBeers() {
	Cellar = []Beer{
		Beer{Id: 1, Name: "Beer 1"},
		Beer{Id: 2, Name: "Beer 2"},
	}
}

func main() {
	populateBeers()

	router := httprouter.New()
	router.GET("/beers", GetBeers)
	//router.GET("/hello/:name", Hello)

	fmt.Println("The Beer server is on tap now.")
	log.Fatal(http.ListenAndServe(":8080", router))
}
