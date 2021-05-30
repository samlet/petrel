package main

import (
	"encoding/json"
	"fmt"
	"github.com/jmoiron/sqlx"
	"testing"
)

type PersonData struct {
	FirstName *string `db:"first_name"`
	LastName  *string `db:"last_name"`
	//Email     string
}

func allPerson(db *sqlx.DB) ([]PersonData, error) {
	// Query the database, storing results in a []Person (wrapped in []interface{})
	people := []PersonData{}
	err := db.Select(&people, `SELECT first_name, last_name 
		FROM person 
		ORDER BY first_name ASC
		LIMIT 5
		`)
	return people, err
}

func TestPerson(t *testing.T) {
	db, err := sqlx.Connect("postgres", "user=xiaofeiwu dbname=ofbiz sslmode=disable")
	if err != nil {
		t.Fatal(err)
	}
	rs, err := allPerson(db)
	if err != nil {
		panic(err)
	}
	for n, r := range rs {
		rowstr, _ := json.MarshalIndent(r, "", "  ")
		fmt.Printf("%d. %s\n", n, rowstr)
	}
}
