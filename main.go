package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/julienschmidt/httprouter"
	_ "github.com/mattn/go-sqlite3" // Driver
)

var (
	// Cellar ... The collection of beer.
	Cellar Beers
)

var (
	// DB holds a connection to the database
	DB *sqlx.DB
)

func openDB() error {
	db, err := sqlx.Open("sqlite3", "db.sqlite3")
	if err != nil {
		if db != nil {
			db.Close()
		}
		return err
	}
	err = db.Ping()
	if err != nil {
		if db != nil {
			db.Close()
		}
		return err
	}

	// Disabled by default
	db.Exec("pragma foreign_keys = on;")

	DB = db
	return nil
}

// createDatabase initializes the database from schema.sql
func createDatabase() {
	err := openDB()
	if err != nil {
		panic(err)
	}
	defer DB.Close()

	schema, err := ioutil.ReadFile("schema.sql")
	if err != nil {
		panic(err)
	}

	_, err = DB.Exec(string(schema))
	if err != nil {
		panic(err)
	}
}

func main() {
	if _, err := os.Stat("db.sqlite3"); os.IsNotExist(err) {
		createDatabase()
	}

	err := openDB()
	if err != nil {
		panic(err)
	}
	defer DB.Close()

	router := httprouter.New()
	router.GET("/beers", GetBeers)
	router.POST("/beers", AddBeer)
	//router.GET("/hello/:name", Hello)

	fmt.Println("The Beer server is on tap at :8080 now.")
	log.Fatal(http.ListenAndServe(":8080", router))
}
