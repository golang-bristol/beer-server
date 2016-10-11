package main

import "time"

type Review struct {
	Id         int       `json:"id"`
	Beer_id    int       `json:"beer_id"`
	First_name string    `json:"first_name"`
	Last_name  string    `json:"last_name"`
	Score      int       `json:"score"`
	Text       string    `json:"text"`
	Created    time.Time `json:"created"`
}

type Beer struct {
	Id         int      `json:"id"`
	Name       string   `json:"name"`
	Brewery    string   `json:"brewery"`
	Abv        int      `json:"abv"`
	Short_desc string   `json:"short_description"`
	Created    tim.Time `json:"created"`
}

type Beers []Beer
