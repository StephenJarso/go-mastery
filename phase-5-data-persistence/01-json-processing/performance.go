package jsonprocessing

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
)

// JSON Performance & Best Practices:
// 1. json.Marshal vs json.Encoder:
//    - Use json.Marshal when you already have the data in memory and need a byte slice.
//    - Use json.NewEncoder(writer).Encode(v) when writing directly to an HTTP response,
//      a file, or any io.Writer. It streams the output directly, avoiding large buffer allocations.
// 2. json.Unmarshal vs json.Decoder:
//    - Use json.Decoder(reader).Decode(&v) when reading from a network connection, file, or io.Reader.
// 3. Precision:
//    - Use decoder.UseNumber() when parsing large numbers to avoid precision loss from conversion to float64.

// StreamEncodeUser writes a user struct directly to an io.Writer.
func StreamEncodeUser(w io.Writer, u User) error {
	encoder := json.NewEncoder(w)
	// Optionally disable HTML escaping if encoding URL values or strings containing <, >, &
	encoder.SetEscapeHTML(false)
	
	err := encoder.Encode(u)
	if err != nil {
		return fmt.Errorf("failed to stream encode user: %w", err)
	}
	return nil
}

// StreamDecodeUser reads a user struct directly from an io.Reader.
func StreamDecodeUser(r io.Reader) (User, error) {
	var u User
	decoder := json.NewDecoder(r)
	
	err := decoder.Decode(&u)
	if err != nil {
		return User{}, fmt.Errorf("failed to stream decode user: %w", err)
	}
	return u, nil
}

// DecodeWithNumericPrecision decodes JSON numbers as json.Number (string wrapper)
// rather than converting them to float64, preserving exact precision for large IDs/amounts.
func DecodeWithNumericPrecision(jsonData string) (map[string]interface{}, error) {
	reader := bytes.NewReader([]byte(jsonData))
	decoder := json.NewDecoder(reader)
	
	// CRITICAL: UseNumber causes the decoder to unmarshal numbers into the interface{}
	// as a json.Number type instead of a float64.
	decoder.UseNumber()
	
	var result map[string]interface{}
	err := decoder.Decode(&result)
	if err != nil {
		return nil, fmt.Errorf("failed to decode with precision: %w", err)
	}
	
	return result, nil
}
