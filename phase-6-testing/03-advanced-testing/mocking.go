package advanced_testing

import (
	"errors"
	"fmt"
	"strings"
)

// User represents our business entity.
type User struct {
	ID    int
	Name  string
	Email string
}

// UserStore defines the storage boundary contract.
type UserStore interface {
	Get(id int) (*User, error)
	Save(user *User) error
}

// UserService coordinates business logic for user management.
type UserService struct {
	Store UserStore
}

// RegisterUser validates user details and saves the user to storage.
func (us *UserService) RegisterUser(name, email string) (*User, error) {
	if name == "" {
		return nil, errors.New("name cannot be empty")
	}
	if !strings.Contains(email, "@") {
		return nil, errors.New("invalid email address")
	}

	newUser := &User{
		Name:  name,
		Email: email,
	}

	err := us.Store.Save(newUser)
	if err != nil {
		return nil, fmt.Errorf("failed to save user: %w", err)
	}

	return newUser, nil
}
