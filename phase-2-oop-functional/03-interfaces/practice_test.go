package phase2oopfunctional

import (
	"io"
	"strings"
	"testing"
)

func TestPracticeStringReader(t *testing.T) {
	// Skip tests if not implemented
	sr := NewStringReader("hello world")
	buf := make([]byte, 5)
	n, err := sr.Read(buf)
	if err != nil && strings.Contains(err.Error(), "not implemented") {
		t.Skip("StringReader not implemented yet")
	}

	if n != 5 || string(buf) != "hello" {
		t.Errorf("Expected to read 'hello', read %d bytes: %q", n, string(buf[:n]))
	}

	n, err = sr.Read(buf)
	if n != 5 || string(buf) != " worl" {
		t.Errorf("Expected to read ' worl', read %d bytes: %q", n, string(buf[:n]))
	}

	// Should EOF at the end
	smallBuf := make([]byte, 10)
	n, err = sr.Read(smallBuf)
	if n != 1 || string(smallBuf[:n]) != "d" {
		t.Errorf("Expected to read remaining 'd', got %d bytes: %q", n, string(smallBuf[:n]))
	}

	n, err = sr.Read(smallBuf)
	if err != io.EOF {
		t.Errorf("Expected EOF at the end, got error: %v", err)
	}
}

func TestPracticeCountingWriter(t *testing.T) {
	cw := &CountingWriter{}
	n, err := cw.Write([]byte("hello"))
	if err != nil && strings.Contains(err.Error(), "not implemented") {
		t.Skip("CountingWriter not implemented yet")
	}

	if n != 5 || err != nil {
		t.Errorf("Expected to write 5 bytes, got %d, err: %v", n, err)
	}

	if cw.BytesWritten() != 5 {
		t.Errorf("Expected 5 bytes written, got %d", cw.BytesWritten())
	}

	cw.Write([]byte(" world"))
	if cw.BytesWritten() != 11 {
		t.Errorf("Expected 11 bytes written, got %d", cw.BytesWritten())
	}

	if cw.String() != "hello world" {
		t.Errorf("Expected buffer content to be 'hello world', got %q", cw.String())
	}
}

func TestPracticeInspectType(t *testing.T) {
	if InspectType(0) == "not implemented" {
		t.Skip("InspectType not implemented yet")
	}

	tests := []struct {
		val      interface{}
		expected string
	}{
		{42, "Integer: 42"},
		{3.14, "Float: 3.14"},
		{"hello", "String: hello"},
		{true, "Boolean: true"},
		{[]int{1}, "Unknown Type"},
	}

	for _, tt := range tests {
		res := InspectType(tt.val)
		if !strings.Contains(res, tt.expected) {
			t.Errorf("InspectType(%v) = %q, want %q", tt.val, res, tt.expected)
		}
	}
}

func TestPracticeAuthenticator(t *testing.T) {
	ba := BasicAuth{Username: "admin", Password: "secret"}
	ok, err := CheckCredentials(ba)
	if err != nil && strings.Contains(err.Error(), "not implemented") {
		t.Skip("BasicAuth/TokenAuth not implemented yet")
	}

	if !ok || err != nil {
		t.Errorf("Expected basic auth to pass, got %v, err: %v", ok, err)
	}

	baBad := BasicAuth{Username: "admin", Password: "wrong-password"}
	ok, _ = CheckCredentials(baBad)
	if ok {
		t.Error("Expected bad basic auth to fail")
	}

	ta := TokenAuth{Token: "valid-token-xyz"}
	ok, err = CheckCredentials(ta)
	if !ok || err != nil {
		t.Errorf("Expected token auth to pass, got %v, err: %v", ok, err)
	}

	taBad := TokenAuth{Token: "bad-token"}
	ok, _ = CheckCredentials(taBad)
	if ok {
		t.Error("Expected bad token auth to fail")
	}
}
