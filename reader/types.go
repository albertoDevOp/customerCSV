package main

import "database/sql"

//Representation of the customer
type Customer struct {
	Id        string `json:"id"`
	Firstname string `json:"first_name"`
	Lastname  string `json:"last_name"`
	Email     string `json:"Email"`
	Phone     string `json:"Phone"`
}

//Function type which represents to the decorator for the repository
type ActionDB func(customer Customer, f Action) (int, error)

//Function type which represents any method in callbacks file
type CrmPost func(customer []byte) (string, error)

//Function type which represents the decorated function
type Action = func(Customer, *sql.DB) error
