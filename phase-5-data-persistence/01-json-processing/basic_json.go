package jsonprocessing

import (
	"encoding/json"
	"fmt"
)

// In Go, the standard library package "encoding/json" handles JSON serialization.
// - "Marshaling" converts Go data structures (structs, maps, slices) into JSON bytes.
// - "Unmarshaling" parses JSON bytes into Go structures.
// Struct tags are used to customize key names, omit empty fields, or ignore fields entirely.

type User struct {
	ID       int    `json:"id"`       // Custom key name "id"
	Username string `json:"username"` // Custom key name "username"
	// omitempty: If Email is empty (""), this key will be omitted from the output JSON.
	Email string `json:"email,omitempty"`
	// json:"-": This field will be completely ignored during JSON marshal/unmarshal.
	SecretKey string `json:"-"`
	// Unexported fields (fields starting with a lowercase letter) are also ignored.
	password string
}

func NewUser(id int, username, email, secret, password string) User {
	return User{
		ID:        id,
		Username:  username,
		Email:     email,
		SecretKey: secret,
		password:  password,
	}
}

// MarshalUser converts a User struct to a JSON byte slice.
func MarshalUser(u User) ([]byte, error) {
	// json.Marshal converts the struct to compact JSON.
	data, err := json.Marshal(u)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal user: %w", err)
	}
	return data, nil
}

// MarshalUserIndent converts a User struct to a pretty-printed JSON byte slice.
func MarshalUserIndent(u User) ([]byte, error) {
	// json.MarshalIndent formats the JSON with prefixes and indentation (e.g. four spaces).
	data, err := json.MarshalIndent(u, "", "    ")
	if err != nil {
		return nil, fmt.Errorf("failed to marshal pretty user: %w", err)
	}
	return data, nil
}

// UnmarshalUser parses a JSON byte slice into a User struct.
func UnmarshalUser(data []byte) (User, error) {
	var u User
	// json.Unmarshal parses JSON-encoded data and stores the result in the value pointed to by &u.
	err := json.Unmarshal(data, &u)
	if err != nil {
		return User{}, fmt.Errorf("failed to unmarshal user: %w", err)
	}
	return u, nil
}
