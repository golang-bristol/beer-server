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

// SaveBeer insert or update beers
func (s *StorageJSON) SaveBeer(beers ...model.Beer) error {
	for _, beer := range beers {
		var err error
		var resource = fmt.Sprintf("%d", beer.ID)

		beersFound, err := s.FindBeer(model.Beer{ID: beer.ID})
		if err != nil {
			return err
		}

		if len(beersFound) > 0 {
			err = s.db.Delete(collectionBeer, resource)
			if err != nil {
				return err
			}
		}

		err = s.db.Write(collectionBeer, resource, beer)
		if err != nil {
			return err
		}
	}
	return nil
}

// SaveReview insert or update reviews
func (s *StorageJSON) SaveReview(reviews ...model.Review) error {
	for _, review := range reviews {
		var err error
		var resource = fmt.Sprintf("%d", review.ID)

		reviewsFound, err := s.FindReview(model.Review{ID: review.ID})
		if err != nil {
			return err
		}

		if len(reviewsFound) > 0 {
			err = s.db.Delete(collectionReview, resource)
			if err != nil {
				return err
			}
		}

		err = s.db.Write(collectionReview, resource, review)
		if err != nil {
			return err
		}
	}
	return nil
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
