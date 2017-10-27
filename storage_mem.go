package main

import (
	"fmt"

	model "github.com/golang-bristol/beer-model"
)

// StorageMemory data storage type save only in memory
type StorageMemory struct {
	cellar  []model.Beer
	reviews []model.Review
}

func newStorageMemory() (*StorageMemory, error) {
	var err error

	stg := new(StorageMemory)

	return stg, err
}

// SaveBeer insert or update beers
func (s *StorageMemory) SaveBeer(beers ...model.Beer) error {
	for _, beer := range beers {
		var err error

		beersFound, err := s.FindBeer(model.Beer{ID: beer.ID})
		if err != nil {
			return err
		}

		if len(beersFound) == 1 {
			*beersFound[0] = beer

		} else if len(beersFound) > 1 {
			return fmt.Errorf("can't update because found multiple entries with ID %d", beer.ID)
		}

		s.cellar = append(s.cellar, beer)
	}

	return nil
}

// SaveReview insert or update reviews
func (s *StorageMemory) SaveReview(reviews ...model.Review) error {
	for _, review := range reviews {
		var err error

		reviewsFound, err := s.FindReview(model.Review{ID: review.ID})
		if err != nil {
			return err
		}

		if len(reviewsFound) == 1 {
			*reviewsFound[0] = review

		} else if len(reviewsFound) > 1 {
			return fmt.Errorf("can't update because found multiple entries with ID %d", review.ID)
		}

		s.reviews = append(s.reviews, review)
	}

	return nil
}

// FindBeer locate full data set based on given criteria
func (s *StorageMemory) FindBeer(criteria model.Beer) ([]*model.Beer, error) {
	var err error
	var beers []*model.Beer
	var blankBeer model.Beer

	for _, b := range s.cellar {
		beer := b

		if criteria == blankBeer { // if blank return all
			beers = append(beers, &beer)

		} else if beer.ID == criteria.ID {
			beers = append(beers, &beer)
		}
	}

	return beers, err
}

// FindReview locate full data set based on given criteria
func (s *StorageMemory) FindReview(criteria model.Review) ([]*model.Review, error) {
	var err error
	var reviews []*model.Review
	var blankReview model.Review

	for _, r := range s.reviews {
		review := r

		if criteria == blankReview {
			reviews = append(reviews, &review)

		} else if review.ID == criteria.ID {
			reviews = append(reviews, &review)
		}
	}

	return reviews, err
}
