package phase2oopfunctional

import (
	"io"
	"errors"
	"fmt"
)


var _ io.Reader
var _ = errors.New
var _ = fmt.Println

// Exercise 1: Custom String Reader
// StringReader reads bytes from content. Implement io.Reader.
type StringReader struct {
	content string
	offset  int
}

func NewStringReader(content string) *StringReader {
	// TODO: Implement
	return nil
}

func (sr *StringReader) Read(p []byte) (n int, err error) {
	// TODO: Implement
	return 0, nil
}

// Exercise 2: Type Switch Value Inspector
// Return string description of type and value.
func InspectType(val interface{}) string {
	// TODO: Implement
	return ""
}

// Exercise 3: Payment Processor System
type PaymentProcessor interface {
	Process(amount float64) (string, error)
}

type CreditCard struct {
	CardNumber string
}

type PayPal struct {
	Email string
}

func (cc CreditCard) Process(amount float64) (string, error) {
	// TODO: Implement
	return "", nil
}

func (pp PayPal) Process(amount float64) (string, error) {
	// TODO: Implement
	return "", nil
}

func ExecutePayment(p PaymentProcessor, amount float64) (string, error) {
	// TODO: Implement
	return "", nil
}
