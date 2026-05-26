package solutions

import (
	"io"
	"errors"
	"fmt"
)


type StringReader struct {
	content string
	offset  int
}

func NewStringReader(content string) *StringReader {
	return &StringReader{content: content}
}

func (sr *StringReader) Read(p []byte) (n int, err error) {
	if sr.offset >= len(sr.content) {
		return 0, io.EOF
	}
	n = copy(p, sr.content[sr.offset:])
	sr.offset += n
	return n, nil
}

func InspectType(val interface{}) string {
	switch v := val.(type) {
	case int:
		return fmt.Sprintf("Integer: %v", v)
	case float64:
		return fmt.Sprintf("Float: %v", v)
	case string:
		return fmt.Sprintf("String: %v", v)
	case bool:
		return fmt.Sprintf("Boolean: %v", v)
	default:
		return "Unknown Type"
	}
}

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
	if amount <= 0 {
		return "", errors.New("invalid amount")
	}
	return fmt.Sprintf("Processed Credit Card payment of $%.2f", amount), nil
}

func (pp PayPal) Process(amount float64) (string, error) {
	if amount <= 0 {
		return "", errors.New("invalid amount")
	}
	return fmt.Sprintf("Processed PayPal payment of $%.2f", amount), nil
}

func ExecutePayment(p PaymentProcessor, amount float64) (string, error) {
	return p.Process(amount)
}
