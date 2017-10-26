package main

import (
	model "github.com/golang-bristol/beer-model"
)

const (
	// StorageTypeJSON is json type indicator
	StorageTypeJSON = "JSON" // storage_json.go

	// StorageTypeMemory will store data in memory
	StorageTypeMemory = "Memory" // storage_mem.go
)

// Storage represents all possible action available to deal with data
type Storage interface {
	SaveBeer(model.Beer) error
	SaveReview(model.Review) error
	SaveBunchOfBeers([]model.Beer) error
	SaveBunchOfReviews([]model.Review) error
	FindBeer(model.Beer) ([]*model.Beer, error)
	FindReview(model.Review) ([]*model.Review, error)
}

func newStorage(storageType string) (Storage, error) {
	var stg Storage
	var err error

	switch storageType {
	case StorageTypeMemory:
		stg, err = newStorageMemory()

	case StorageTypeJSON:
		// TODO: configuration setup - For the moment storage location for
		// JSON files is current working directory.
		stg, err = newStorageJSON("./")
	}

	return stg, err
}
