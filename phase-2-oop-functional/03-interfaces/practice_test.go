package phase2oopfunctional

import (
	"testing"
	"io"
)


func TestStringReader(t *testing.T) {
	sr := NewStringReader("test")
	buf := make([]byte, 2)
	n, err := sr.Read(buf)
	if n != 2 || string(buf) != "te" || err != nil {
		t.Errorf("read failed: n=%d, err=%v", n, err)
	}
	n, err = sr.Read(buf)
	if n != 2 || string(buf) != "st" || err != nil {
		t.Errorf("read failed: n=%d, err=%v", n, err)
	}
	_, err = sr.Read(buf)
	if err != io.EOF {
		t.Errorf("expected EOF, got %v", err)
	}
}

func TestInspectType(t *testing.T) {
	if InspectType(42) != "Integer: 42" {
		t.Error("InspectType failed for int")
	}
}

func TestPaymentProcessor(t *testing.T) {
	cc := CreditCard{CardNumber: "1234"}
	res, err := ExecutePayment(cc, 100.50)
	if err != nil || res != "Processed Credit Card payment of $100.50" {
		t.Errorf("credit card payment failed: %q, %v", res, err)
	}
}
