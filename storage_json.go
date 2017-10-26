package main

import (
	"encoding/json"
	"fmt"

	model "github.com/golang-bristol/beer-model"
	"github.com/nanobox-io/golang-scribble"
)

// StorageJSON is the data storage type using JSON file
type StorageJSON struct {
	db *scribble.Driver
}

var (
	collectionBeer   = "beers"
	collectionReview = "reviews"
)

func newStorageJSON(location string) (*StorageJSON, error) {
	var err error

	stg := new(StorageJSON)

	stg.db, err = scribble.New(location, nil)
	if err != nil {
		return nil, err
	}

	return stg, nil
}

// SaveBeer insert or update a single beer
func (s *StorageJSON) SaveBeer(beer model.Beer) error {
	var err error
	var resource = fmt.Sprintf("%d", beer.ID)

	beers, err := s.FindBeer(model.Beer{ID: beer.ID})

	if len(beers) > 0 {
		s.db.Delete(collectionBeer, resource)
	}

	s.db.Write(collectionBeer, resource, beer)

	return err
}

// SaveReview insert or update a single review
func (s *StorageJSON) SaveReview(review model.Review) error {
	var err error
	var resource = fmt.Sprintf("%d", review.ID)

	reviews, err := s.FindReview(model.Review{ID: review.ID})

	if len(reviews) > 0 {
		s.db.Delete(collectionReview, resource)
	}

	s.db.Write(collectionReview, resource, review)

	return err
}

// SaveBunchOfBeers insert or update a multiple beers
func (s *StorageJSON) SaveBunchOfBeers(beers []model.Beer) error {
	var err error

	for _, beer := range beers {
		s.SaveBeer(beer)
	}

	return err
}

// SaveBunchOfReviews insert or update a multiple reviews
func (s *StorageJSON) SaveBunchOfReviews(reviews []model.Review) error {
	var err error

	for _, review := range reviews {
		s.SaveReview(review)
	}

	return err
}

// FindBeer locate full data set based on given criteria
func (s *StorageJSON) FindBeer(criteria model.Beer) ([]*model.Beer, error) {
	var beers []*model.Beer
	var blankBeer model.Beer

	if criteria == blankBeer {
		records, err := s.db.ReadAll(collectionBeer)
		if err != nil {
			return beers, err
		}

		for _, b := range records {
			var beer model.Beer

			if err := json.Unmarshal([]byte(b), &beer); err != nil {
				return beers, err
			}

			beers = append(beers, &beer)
		}

	} else {
		var beer model.Beer
		resource := fmt.Sprintf("%d", criteria.ID)

		if err := s.db.Read(collectionBeer, resource, &beer); err != nil {
			return beers, err
		}

		beers = append(beers, &beer)
	}

	return beers, nil
}

// FindReview locate full data set based on given criteria
func (s *StorageJSON) FindReview(criteria model.Review) ([]*model.Review, error) {
	var reviews []*model.Review
	var blankReview model.Review

	if criteria == blankReview {
		records, err := s.db.ReadAll(collectionReview)
		if err != nil {
			return reviews, err
		}

		for _, r := range records {
			var review model.Review

			if err := json.Unmarshal([]byte(r), &review); err != nil {
				return reviews, err
			}

			reviews = append(reviews, &review)
		}

	} else {
		var review model.Review
		resource := fmt.Sprintf("%d", criteria.ID)

		if err := s.db.Read(collectionReview, resource, &review); err != nil {
			return reviews, err
		}

		reviews = append(reviews, &review)
	}

	return reviews, nil
}
