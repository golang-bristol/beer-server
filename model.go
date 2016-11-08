package main

import "database/sql"

// Review represents a single review
type Review struct {
	ID        int            `json:"id" db:"ID"`
	Beer      int            `json:"beer_id" db:"Beer"`
	Name      string         `json:"name" db:"Name"`
	Score     int            `json:"score" db:"Score"`
	Text      sql.NullString `json:"text" db:"Text"`
	CreatedAt []uint8        `json:"created" db:"CreatedAt"`
	UpdatedAt []uint8        `json:"created" db:"UpdatedAt"`
	DeletedAt sql.NullInt64  `json:"created" db:"DeletedAt"`
}

// Brewery represents a single brewery
type Brewery struct {
	ID        int            `json:"id" db:"ID"`
	Name      string         `json:"name" db:"Name"`
	Location  sql.NullString `json:"location" db:"Location"`
	CreatedAt []uint8        `json:"created" db:"CreatedAt"`
	UpdatedAt []uint8        `json:"created" db:"UpdatedAt"`
	DeletedAt sql.NullInt64  `json:"created" db:"DeletedAt"`
}

// Beer represents a single beer
type Beer struct {
	ID        int            `json:"id" db:"ID"`
	Name      string         `json:"name" db:"Name"`
	Brewery   sql.NullInt64  `json:"brewery" db:"Brewery"`
	ABV       sql.NullInt64  `json:"abv" db:"ABV"`
	ShortDesc sql.NullString `json:"short_description" db:"ShortDesc"`
	CreatedAt []uint8        `json:"created" db:"CreatedAt"`
	UpdatedAt []uint8        `json:"created" db:"UpdatedAt"`
	DeletedAt sql.NullInt64  `json:"created" db:"DeletedAt"`
}

// Beers is 0 or more beers
type Beers []Beer

// Persist to the database.
func (b Beer) Persist() error {
	// New
	if b.ID == 0 {
		_, err := DB.Exec("insert into beers (name) values(?)",
			b.ID)
		return err
	}

	// Existing
	_, err := DB.Exec("update beers set Name=? where ID=?",
		b.Name, b.ID)

	return err
}

// ListBeers gets all beers
func ListBeers() (beers Beers, err error) {
	err = DB.Select(&beers, "select * from beers where DeletedAt is null")
	if err != nil {
		return nil, err
	}
	return beers, nil
}
