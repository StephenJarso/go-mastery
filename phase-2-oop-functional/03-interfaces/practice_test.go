package phase2oopfunctional

import (
	"io"
	"testing"
)

func TestStringReader(t *testing.T) {
	sr := NewStringReader("test")
	if sr == nil {
		t.Skip("Exercise StringReader not implemented yet")
	}
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
	res := InspectType(42)
	if res == "" {
		t.Skip("Exercise InspectType not implemented yet")
	}
	if res != "Integer: 42" {
		t.Error("InspectType failed for int")
	}
}

func TestPaymentProcessor(t *testing.T) {
	cc := CreditCard{CardNumber: "1234"}
	res, err := ExecutePayment(cc, 100.50)
	if res == "" && err == nil {
		t.Skip("Exercise PaymentProcessor not implemented yet")
	}
	if err != nil || res != "Processed Credit Card payment of $100.50" {
		t.Errorf("credit card payment failed: %q, %v", res, err)
	}
}
