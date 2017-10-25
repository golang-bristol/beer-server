package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	model "github.com/golang-bristol/beer-model"
)

func TestGetBeers(t *testing.T) {
	var cellar []model.Beer

	w := httptest.NewRecorder()
	r, _ := http.NewRequest("GET", "/beers", nil)

	router.ServeHTTP(w, r)

	json.Unmarshal(w.Body.Bytes(), &cellar)

	if w.Code != http.StatusOK {
		t.Errorf("Expected route GET /beers to be valid.")
		t.FailNow()
	}

	if len(cellar) != len(Cellar) {
		t.Error("Expected number of beers from request /beers to be the same as exported Cellar variable")
		t.FailNow()
	}

	var mapCellar = make(map[model.Beer]int, len(Cellar))
	for _, beer := range Cellar {
		mapCellar[beer] = 1
	}

	for _, beerResp := range cellar {
		if _, ok := mapCellar[beerResp]; !ok {
			t.Errorf("Expected all results to match existing records")
			t.FailNow()
			break
		}
	}
}

func TestAddBeer(t *testing.T) {
	newBeer := model.Beer{
		ID:      555,
		Name:    "Testing beer",
		Abv:     333,
		Brewery: "Testing Beer Inc",
	}

	newBeerJSON, err := json.Marshal(newBeer)
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()
	r, _ := http.NewRequest("POST", "/beers", bytes.NewBuffer(newBeerJSON))

	router.ServeHTTP(w, r)

	if w.Code != http.StatusOK {
		t.Errorf("Expected route POST /beers to be valid.")
		t.FailNow()
	}

	var mapCellar = make(map[model.Beer]int, len(Cellar))
	for _, beer := range Cellar {
		mapCellar[beer] = 1
	}

	if _, ok := mapCellar[newBeer]; !ok {
		t.Errorf("Expected to find new entry in the exported variable `Cellar`")
		t.FailNow()
	}

}
