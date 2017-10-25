package main

import (
	"time"

	"github.com/golang-bristol/beer-model"
)

// PopulateBeers populates the Cellar variable with Beers
func PopulateBeers() {
	Cellar = []model.Beer{
		model.Beer{
			ID:      1,
			Name:    "Pliny the Elder",
			Brewery: "Russian River Brewing Company",
			Abv:     8,
			ShortDesc: "Pliny the Elder is brewed with Amarillo, " +
				"Centennial, CTZ, and Simcoe hops. It is well-balanced with " +
				"malt, hops, and alcohol, slightly bitter with a fresh hop " +
				"aroma of floral, citrus, and pine.",
			Created: time.Date(2017, time.October, 24, 22, 6, 0, 0, time.UTC),
		},
		model.Beer{
			ID:      2,
			Name:    "Oatmeal Stout",
			Brewery: "Samuel Smith",
			Abv:     5,
			ShortDesc: "Brewed with well water (the original well at the " +
				"Old Brewery, sunk in 1758, is still in use, with the hard well " +
				"water being drawn from 85 feet underground); fermented in " +
				"‘stone Yorkshire squares’ to create an almost opaque, " +
				"wonderfully silky and smooth textured ale with a complex " +
				"medium dry palate and bittersweet finish.",
			Created: time.Date(2017, time.October, 24, 22, 12, 0, 0, time.UTC),
		},
		model.Beer{
			ID:      3,
			Name:    "Märzen",
			Brewery: "Schlenkerla",
			Abv:     5,
			ShortDesc: "Bamberg's speciality, a dark, bottom fermented " +
				"smokebeer, brewed with Original Schlenkerla Smokemalt from " +
				"the Schlenkerla maltings and tapped according to old tradition " +
				"directly from the gravity-fed oakwood cask in the historical " +
				"brewery tavern.",
			Created: time.Date(2017, time.October, 24, 22, 17, 0, 0, time.UTC),
		},
		model.Beer{
			ID:      4,
			Name:    "Duvel",
			Brewery: "Duvel Moortgat",
			Abv:     9,
			ShortDesc: "A Duvel is still seen as the reference among strong " +
				"golden ales. Its bouquet is lively and tickles the nose with an " +
				"element of citrus which even tends towards grapefruit thanks to " +
				"the use of only the highest-quality hop varieties.",
			Created: time.Date(2017, time.October, 24, 22, 24, 0, 0, time.UTC),
		},
		model.Beer{
			ID:      5,
			Name:    "Negra",
			Brewery: "Modelo",
			Abv:     5,
			ShortDesc: "Brewed longer to enhance the flavors, this Munich " +
				"Dunkel-style Lager gives way to a rich flavor and remarkably " +
				"smooth taste.",
			Created: time.Date(2017, time.October, 24, 22, 27, 0, 0, time.UTC),
		},
	}
}

// PopulateReviews populates the Reviews variable with Reviews
func PopulateReviews() {
	Reviews = []model.Review{
		model.Review{ID: 1, BeerID: 1, FirstName: "Joe", LastName: "Tribiani", Score: 5, Text: "This is good but this is not pizza!", Created: time.Date(2017, time.November, 10, 12, 36, 0, 0, time.UTC)},
	}
}
