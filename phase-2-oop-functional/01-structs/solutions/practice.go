package solutions

import (
	"errors"
	"encoding/json"
	"fmt"
)


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

func BorrowBook(book *Book) (string, error) {
	if book == nil {
		return "", errors.New("nil book pointer")
	}
	if !book.Available {
		return "", errors.New("book is already borrowed")
	}
	book.Available = false
	return fmt.Sprintf("You have successfully borrowed %q.", book.Title), nil
}

func NewEmployee(name string, age int, city, state string, id int, pos string) Employee {
	return Employee{
		Person: Person{
			Name: name,
			Age:  age,
			Address: Address{
				City:  city,
				State: state,
			},
		},
		ID:       id,
		Position: pos,
	}
}

func ToJSON(p Product) (string, error) {
	bytes, err := json.Marshal(p)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}
