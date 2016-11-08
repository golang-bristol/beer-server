package main

import (
	"github.com/golang-bristol/beer-model"
)

func populateBeers() {
	Cellar = []model.Beer{
		model.Beer{Id: 1, Name: "Beer 1"},
		model.Beer{Id: 2, Name: "Beer 2"},
	}
}
