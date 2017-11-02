package main

import model "github.com/golang-bristol/beer-model"

// StorageMemory data storage type save only in memory
type StorageMemory struct {
	cellar  []model.Beer
	reviews []model.Review
}

// SaveBeer insert or update beers
func (s *StorageMemory) SaveBeer(beers ...model.Beer) error {
	for _, beer := range beers {
		var err error

		beersFound, err := s.FindBeer(beer)
		if err != nil {
			return err
		}

		if len(beersFound) == 1 {
			*beersFound[0] = beer
			return nil
		}

		beer.ID = len(s.cellar) + 1
		s.cellar = append(s.cellar, beer)
	}

	return nil
}

// SaveReview insert or update reviews
func (s *StorageMemory) SaveReview(reviews ...model.Review) error {
	for _, review := range reviews {
		var err error

		reviewsFound, err := s.FindReview(review)
		if err != nil {
			return err
		}

		if len(reviewsFound) == 1 {
			*reviewsFound[0] = review
			return nil
		}

		review.ID = len(s.reviews) + 1
		s.reviews = append(s.reviews, review)
	}

	return nil
}

// FindBeer locate full data set based on given criteria
func (s *StorageMemory) FindBeer(criteria model.Beer) ([]*model.Beer, error) {
	var beers []*model.Beer

	for idx := range s.cellar {

		if s.cellar[idx].ID == criteria.ID {
			beers = append(beers, &s.cellar[idx])
		}
	}

	return beers, nil
}

// FindReview locate full data set based on given criteria
func (s *StorageMemory) FindReview(criteria model.Review) ([]*model.Review, error) {
	var reviews []*model.Review

	for idx := range s.reviews {
		if s.reviews[idx].ID == criteria.ID || s.reviews[idx].BeerID == criteria.BeerID {
			reviews = append(reviews, &s.reviews[idx])
		}
	}

	return reviews, nil
}

// FindBeers return all beers
func (s *StorageMemory) FindBeers() []*model.Beer {
	var beers []*model.Beer

	for idx := range s.cellar {
		beers = append(beers, &s.cellar[idx])
	}

	return beers
}

// FindReviews return all reviews
func (s *StorageMemory) FindReviews() []*model.Review {
	var reviews []*model.Review

	for idx := range s.reviews {
		reviews = append(reviews, &s.reviews[idx])
	}

	return reviews
}
