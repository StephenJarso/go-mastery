package serialization

import (
	"bytes"
	"encoding/gob"
)


var _ = bytes.NewBuffer
var _ = gob.NewEncoder

type CacheItem struct {
	Key   string
	Value []byte
}

// Exercise 1: Gob Serialize
// Serialize the given struct or value to byte slice using Gob encoder.
func GobSerialize(v interface{}) ([]byte, error) {
	// TODO: Implement
	return nil, nil
}

// Exercise 2: Gob Deserialize
// Deserialize Gob-encoded bytes into the target pointer.
func GobDeserialize(b []byte, target interface{}) error {
	// TODO: Implement
	return nil
}
