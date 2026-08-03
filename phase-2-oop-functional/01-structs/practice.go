package main

import (
	"encoding/json"
	"errors"
)

var _ = errors.New
var _ = json.Marshal

type Book struct {
	ID        int
	Title     string
	Author    string
	Available bool
}

type Library struct {
	Name  string
	Books []Book
}

type Address struct {
	City  string
	State string
}

type Person struct {
	Name    string
	Age     int
	Address Address
}

type Employee struct {
	Person
	ID       int
	Position string
}

type Product struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

// Exercise 1: Borrow a Book
// Set Available to false on the Book pointer and return a success message.
// If the book is already borrowed, return an error.
func BorrowBook(book *Book) (string, error) {
	// TODO: Implement
	return "", nil
}

// Exercise 2: New Employee Constructor
// Return a new Employee struct initialized with composition.
func NewEmployee(name string, age int, city, state string, id int, pos string) Employee {
	// TODO: Implement
	return Employee{}
}

// Exercise 3: Product to JSON
// Convert a Product struct instance to its JSON string representation.
func ToJSON(p Product) (string, error) {
	// TODO: Implement
	return "", nil
}
