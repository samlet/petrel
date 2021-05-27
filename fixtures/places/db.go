package main

import (
	"database/sql"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var schema = `
CREATE TABLE person (
    first_name text,
    last_name text,
    email text
);

CREATE TABLE place (
    country text,
    city text NULL,
    telcode integer
)`

type Person struct {
	FirstName string `db:"first_name"`
	LastName  string `db:"last_name"`
	Email     string
}

type Place struct {
	Country string
	City    sql.NullString
	TelCode int
}

type PlacesDb struct {
	db *sqlx.DB
}

func NewPlacesDb() (*PlacesDb, error) {
	// this Pings the database trying to connect
	// use sqlx.Open() for sql.Open() semantics
	db, err := sqlx.Connect("postgres", "user=xiaofeiwu dbname=places sslmode=disable")
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	return &PlacesDb{db: db}, nil
}

func (c PlacesDb) Seed() {
	// exec the schema or fail; multi-statement Exec behavior varies between
	// database drivers;  pq will exec them all, sqlite3 won't, ymmv
	c.db.MustExec(schema)

	tx := c.db.MustBegin()
	tx.MustExec("INSERT INTO person (first_name, last_name, email) VALUES ($1, $2, $3)", "Jason", "Moiron", "jmoiron@jmoiron.net")
	tx.MustExec("INSERT INTO person (first_name, last_name, email) VALUES ($1, $2, $3)", "John", "Doe", "johndoeDNE@gmail.net")
	tx.MustExec("INSERT INTO place (country, city, telcode) VALUES ($1, $2, $3)", "United States", "New York", "1")
	tx.MustExec("INSERT INTO place (country, telcode) VALUES ($1, $2)", "Hong Kong", "852")
	tx.MustExec("INSERT INTO place (country, telcode) VALUES ($1, $2)", "Singapore", "65")
	// Named queries can use structs, so if you have an existing struct (i.e. person := &Person{}) that you have populated, you can pass it in as &person
	tx.NamedExec("INSERT INTO person (first_name, last_name, email) VALUES (:first_name, :last_name, :email)", &Person{"Jane", "Citizen", "jane.citzen@example.com"})
	tx.Commit()
}

func (c PlacesDb) AllPerson() ([]Person, error) {
	// Query the database, storing results in a []Person (wrapped in []interface{})
	people := []Person{}
	err := c.db.Select(&people, "SELECT * FROM person ORDER BY first_name ASC")
	return people, err
}

func (c PlacesDb) AllPlaces() ([]Place, error) {
	places := []Place{}
	err := c.db.Select(&places, "SELECT * FROM place ORDER BY telcode ASC")
	return places, err
}

// CreatePerson
// 	{
//	"first": "Bin",
//	"last": "Smuth",
//	"email": "bensmith@allblacks.nz",
//    }
func (c PlacesDb) CreatePerson(data map[string]interface{}) (int64, error) {
	r, err := c.db.NamedExec(`INSERT INTO person (first_name,last_name,email) VALUES (:first,:last,:email)`,
		data)
	if err != nil {
		return 0, err
	} else {
		return r.RowsAffected()
	}
}

// CreatePersons
// ```personStructs := []Person{
//        {FirstName: "Ardie", LastName: "Savea", Email: "asavea@ab.co.nz"},
//        {FirstName: "Sonny Bill", LastName: "Williams", Email: "sbw@ab.co.nz"},
//        {FirstName: "Ngani", LastName: "Laumape", Email: "nlaumape@ab.co.nz"},
//    }
//    c.CreatePersons(personStructs)
// ```
func (c PlacesDb) CreatePersons(personStructs []Person) (int64, error) {
	r, err := c.db.NamedExec(`INSERT INTO person (first_name, last_name, email)
        VALUES (:first_name, :last_name, :email)`, personStructs)
	if err != nil {
		return 0, err
	} else {
		return r.RowsAffected()
	}
}

/**
refs:
	- https://github.com/jmoiron/sqlx
*/
