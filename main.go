package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type dummy struct {
	Name string
}

func GetBeers(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	result := dummy{"Here are some beers!"}
	json.NewEncoder(w).Encode(result)
}

func Hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "hello, %s!\n", ps.ByName("name"))
}

func main() {
	router := httprouter.New()
	router.GET("/beers", GetBeers)
	router.GET("/hello/:name", Hello)

	fmt.Println("The Beer server is on tap now.")
	log.Fatal(http.ListenAndServe(":8080", router))
}
