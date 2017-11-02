package main

import (
	model "github.com/golang-bristol/beer-model"
)

// StorageType defines available storage types
type StorageType int

const (
	// JSON will store data in JSON files saved on disk
	JSON StorageType = iota
	// Memory will store data in memory
	Memory
)

// Storage represents all possible actions available to deal with data
type Storage interface {
	SaveBeer(...model.Beer) error
	SaveReview(...model.Review) error
	FindBeer(model.Beer) ([]*model.Beer, error)
	FindReview(model.Review) ([]*model.Review, error)
	FindBeers() []*model.Beer
	FindReviews() []*model.Review
}

func newStorage(storageType StorageType) (Storage, error) {
	var stg Storage
	var err error

	switch storageType {
	case Memory:
		stg = new(StorageMemory)

	case JSON:
		// TODO: configuration setup - For the moment storage location for
		// JSON files is current working directory.
		stg, err = newStorageJSON("./data/")
	}

	return stg, err
}
