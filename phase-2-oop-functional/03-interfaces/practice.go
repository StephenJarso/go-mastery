package phase2oopfunctional

import (
	"errors"
	"fmt"
	"io"
)

// ==========================================
// EXERCISE 1: Custom String Reader
// ==========================================
// StringReader reads bytes from an underlying string.
// Implement the io.Reader interface.
type StringReader struct {
	content string
	offset  int
}

func NewStringReader(content string) *StringReader {
	return &StringReader{content: content}
}

// Read implements io.Reader. It reads up to len(p) bytes from the content.
func (sr *StringReader) Read(p []byte) (n int, err error) {
	if sr.offset >= len(sr.content) {
		return 0, io.EOF
	}
	n = copy(p, sr.content[sr.offset:])
	sr.offset += n
	return n, nil
}

// ==========================================
// EXERCISE 2: Counting Writer
// ==========================================
// CountingWriter writes data to an underlying slice and counts total bytes written.
// Implement the io.Writer interface.
type CountingWriter struct {
	buf   []byte
	count int
}

func (cw *CountingWriter) Write(p []byte) (n int, err error) {
	cw.buf = append(cw.buf, p...)
	cw.count += len(p)
	return len(p), nil
}

func (cw *CountingWriter) BytesWritten() int {
	return cw.count
}

func (cw *CountingWriter) String() string {
	return string(cw.buf)
}

// ==========================================
// EXERCISE 3: Type Switch Value Inspector
// ==========================================
// InspectType accepts any value and returns a description of its type and value.
// It should support: int, float64, string, bool, and any other type (as default).
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

// ==========================================
// EXERCISE 4: Interface Composition
// ==========================================
// ReadWriteStub combines both Reader and Writer stubs from above.
// Implement io.ReadWriter.
type ComposedReadWriter struct {
	reader io.Reader
	writer io.Writer
}

func NewComposedReadWriter(r io.Reader, w io.Writer) *ComposedReadWriter {
	return &ComposedReadWriter{reader: r, writer: w}
}

func (crw *ComposedReadWriter) Read(p []byte) (n int, err error) {
	if crw.reader == nil {
		return 0, io.EOF
	}
	return crw.reader.Read(p)
}

func (crw *ComposedReadWriter) Write(p []byte) (n int, err error) {
	if crw.writer == nil {
		return 0, io.ErrClosedPipe
	}
	return crw.writer.Write(p)
}

// ==========================================
// EXERCISE 5: Authenticator System
// ==========================================
// Authenticator defines the interface for verifying credentials.
type Authenticator interface {
	Authenticate() (bool, error)
}

// BasicAuth authenticates via username and password.
type BasicAuth struct {
	Username string
	Password string
}

// TODO: Implement Authenticator for BasicAuth
// Authenticate should return true if Username == "admin" and Password == "secret",
// otherwise return false or error.
func (ba BasicAuth) Authenticate() (bool, error) {
	if ba.Username == "admin" && ba.Password == "secret" {
		return true, nil
	}
	return false, errors.New("invalid credentials")
}

// TokenAuth authenticates via token string.
type TokenAuth struct {
	Token string
}

// TODO: Implement Authenticator for TokenAuth
// Authenticate should return true if Token is exactly "valid-token-xyz",
// otherwise return false or error.
func (ta TokenAuth) Authenticate() (bool, error) {
	if ta.Token == "valid-token-xyz" {
		return true, nil
	}
	return false, errors.New("invalid token")
}

// CheckCredentials runs Authenticate on any Authenticator and returns result.
func CheckCredentials(auth Authenticator) (bool, error) {
	return auth.Authenticate()
}
