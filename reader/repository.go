package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	host     = "db"
	port     = 5432
	user     = "root"
	password = "root"
	dbname   = "paack"
)

// connection against the DB
func connect() (*sql.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	return sql.Open("postgres", psqlInfo)
}

//inserts a new customer into DB
func insert(customer Customer, db *sql.DB) error {
	statement := "INSERT INTO paack.customers (id, firstname, lastname, email, phone) VALUES ($1, $2, $3, $4, $5) RETURNING id"
	var cid []uint8
	return db.QueryRow(statement, customer.Id, customer.Firstname, customer.Lastname, customer.Email, customer.Phone).Scan(&cid)
}

//Decorator which wraps statements against DB
//The only porpouse is to keep common code aside in a single place
func action(customer Customer, op Action) (int, error) {
	db, err := connect()
	defer db.Close()
	if err != nil {
		log.Fatal(err)
		return -1, err
	}
	//invoke the method to perorm
	err = op(customer, db)
	if err != nil {
		log.Fatal(err)
		return -1, err
	}
	return 0, nil
}
