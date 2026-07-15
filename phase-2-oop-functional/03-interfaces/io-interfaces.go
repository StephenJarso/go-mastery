package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
)

// io-interfaces.go demonstrates common interfaces from the Go standard library
// These interfaces are fundamental to Go programming and appear throughout the stdlib

// ---- io.Reader Interface ----

// Reader is Go's fundamental interface for reading data
// type Reader interface {
//     Read(p []byte) (n int, err error)
// }
// Any type that implements Read() is a Reader

// StringAsReader shows how strings.Reader implements io.Reader
func StringAsReader() {
	fmt.Println("\n=== io.Reader: strings.Reader ===")

	// Create a Reader from a string
	reader := strings.NewReader("Hello, World!")

	// Read into a buffer
	buffer := make([]byte, 5)
	n, err := reader.Read(buffer)
	if err != nil && err != io.EOF {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("Read %d bytes: %s\n", n, string(buffer))

	// Read more
	n, err = reader.Read(buffer)
	fmt.Printf("Read %d bytes: %s\n", n, string(buffer[:n]))

	// Continue until EOF
	n, err = reader.Read(buffer)
	if err == io.EOF {
		fmt.Println("Reached end of reader")
	}
}

// BytesAsReader shows how bytes.Buffer implements io.Reader
func BytesAsReader() {
	fmt.Println("\n=== io.Reader: bytes.Buffer ===")

	// Create a buffer
	var buf bytes.Buffer
	buf.WriteString("Go is awesome")

	// bytes.Buffer implements io.Reader
	buffer := make([]byte, 4)
	for {
		n, err := buf.Read(buffer)
		if err == io.EOF {
			break
		}
		fmt.Printf("Read: %s\n", string(buffer[:n]))
	}
}

// CustomReader implements the io.Reader interface
type CustomReader struct {
	data string
	pos  int
}

func NewCustomReader(data string) *CustomReader {
	return &CustomReader{data: data}
}

func (cr *CustomReader) Read(p []byte) (int, error) {
	if cr.pos >= len(cr.data) {
		return 0, io.EOF
	}

	n := copy(p, cr.data[cr.pos:])
	cr.pos += n
	return n, nil
}

// CustomReaderExample shows implementing io.Reader
func CustomReaderExample() {
	fmt.Println("\n=== Custom io.Reader Implementation ===")

	reader := NewCustomReader("Custom data source")

	buffer := make([]byte, 6)
	for {
		n, err := reader.Read(buffer)
		if err == io.EOF {
			break
		}
		fmt.Printf("Read from custom: %s\n", string(buffer[:n]))
	}
}

// ---- io.Writer Interface ----

// Writer is Go's fundamental interface for writing data
// type Writer interface {
//     Write(p []byte) (n int, err error)
// }

// BytesBufferAsWriter shows how bytes.Buffer implements io.Writer
func BytesBufferAsWriter() {
	fmt.Println("\n=== io.Writer: bytes.Buffer ===")

	var buf bytes.Buffer

	// bytes.Buffer implements io.Writer
	fmt.Fprintf(&buf, "Hello %s\n", "World")
	fmt.Fprintf(&buf, "Line 2\n")
	fmt.Fprintf(&buf, "Line 3\n")

	fmt.Println("Buffer contents:")
	fmt.Print(buf.String())
}

// CustomWriter implements io.Writer
type CustomWriter struct {
	lines []string
}

func (cw *CustomWriter) Write(p []byte) (int, error) {
	cw.lines = append(cw.lines, string(p))
	return len(p), nil
}

func (cw *CustomWriter) GetLines() []string {
	return cw.lines
}

// CustomWriterExample shows implementing io.Writer
func CustomWriterExample() {
	fmt.Println("\n=== Custom io.Writer Implementation ===")

	writer := &CustomWriter{}

	fmt.Fprintf(writer, "First line\n")
	fmt.Fprintf(writer, "Second line\n")
	fmt.Fprintf(writer, "Third line\n")

	fmt.Println("Custom writer collected:")
	for i, line := range writer.GetLines() {
		fmt.Printf("[%d] %s", i+1, line)
	}
}

// ---- io.ReadWriter Interface ----

// ReadWriter combines Reader and Writer
// type ReadWriter interface {
//     Reader
//     Writer
// }

// bytes.Buffer implements both Reader and Writer, so it implements ReadWriter
func ReadWriterExample() {
	fmt.Println("\n=== io.ReadWriter: bytes.Buffer ===")

	var buf bytes.Buffer

	// Write to buffer
	fmt.Fprintf(&buf, "Hello from bytes buffer\n")

	// Read from buffer
	data := make([]byte, 100)
	n, _ := buf.Read(data)

	fmt.Printf("Read from buffer: %s", string(data[:n]))
}

// ---- fmt.Stringer Interface ----

// Stringer interface requires only a String() method
// type Stringer interface {
//     String() string
// }

// Person implements fmt.Stringer
type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%s (age %d)", p.Name, p.Age)
}

// Book implements fmt.Stringer
type Book struct {
	Title  string
	Author string
}

func (b Book) String() string {
	return fmt.Sprintf("'%s' by %s", b.Title, b.Author)
}

// StringerExample shows fmt.Stringer usage
func StringerExample() {
	fmt.Println("\n=== fmt.Stringer Interface ===")

	people := []fmt.Stringer{
		Person{"Alice", 30},
		Person{"Bob", 25},
		Book{"Go Programming", "John Doe"},
	}

	fmt.Println("Using fmt.Stringer:")
	for i, s := range people {
		fmt.Printf("[%d] %v\n", i, s)
	}
}

// ---- io.Closer Interface ----

// Closer interface requires a Close() method
// type Closer interface {
//     Close() error
// }

// Resource is a custom resource that needs closing
type Resource struct {
	name   string
	closed bool
}

func (r *Resource) Close() error {
	if r.closed {
		return fmt.Errorf("resource already closed")
	}
	r.closed = true
	fmt.Printf("Closed resource: %s\n", r.name)
	return nil
}

func (r *Resource) IsClosed() bool {
	return r.closed
}

// CloserExample shows resource cleanup patterns
func CloserExample() {
	fmt.Println("\n=== io.Closer Interface ===")

	resource := &Resource{name: "Database Connection"}

	// Simulate defer pattern for cleanup
	defer func() {
		if err := resource.Close(); err != nil {
			fmt.Printf("Error closing: %v\n", err)
		}
	}()

	fmt.Printf("Using resource: %s\n", resource.name)
	fmt.Printf("Resource closed: %v\n", resource.IsClosed())
}

// ---- Multiple Interface Composition ----

// File implements Reader, Writer, and Closer
// (in reality, *os.File does this)

// LoggedFile wraps a file and logs operations
type LoggedFile struct {
	name     string
	contents string
	pos      int
}

func (lf *LoggedFile) Read(p []byte) (int, error) {
	fmt.Printf("[READ] Reading from %s\n", lf.name)
	if lf.pos >= len(lf.contents) {
		return 0, io.EOF
	}
	n := copy(p, lf.contents[lf.pos:])
	lf.pos += n
	return n, nil
}

func (lf *LoggedFile) Write(p []byte) (int, error) {
	fmt.Printf("[WRITE] Writing to %s: %s\n", lf.name, string(p))
	lf.contents += string(p)
	return len(p), nil
}

func (lf *LoggedFile) Close() error {
	fmt.Printf("[CLOSE] Closing %s\n", lf.name)
	return nil
}

// IOComposedInterfaceExample shows implementing multiple interfaces
func IOComposedInterfaceExample() {
	fmt.Println("\n=== Composed Interfaces (Reader + Writer + Closer) ===")

	file := &LoggedFile{name: "data.txt"}

	// Use as Writer
	file.Write([]byte("Hello World"))

	// Use as Reader
	buffer := make([]byte, 5)
	n, _ := file.Read(buffer)
	fmt.Printf("Read: %s\n", string(buffer[:n]))

	// Use as Closer
	file.Close()
}

// ---- io.Copy Pattern ----

// CopyExample demonstrates the power of Reader/Writer abstraction
func CopyExample() {
	fmt.Println("\n=== io.Copy: Reader + Writer Pattern ===")

	// Read from string
	source := strings.NewReader("This is content to copy")

	// Write to buffer
	var destination bytes.Buffer

	// io.Copy works with any Reader and Writer!
	n, err := io.Copy(&destination, source)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("Copied %d bytes\n", n)
	fmt.Printf("Destination: %s\n", destination.String())
}

// ---- Chaining Readers ----

// ChainedReaderExample shows composing multiple readers
func ChainedReaderExample() {
	fmt.Println("\n=== Chaining Readers ===")

	// Create multiple readers
	reader1 := strings.NewReader("First part. ")
	reader2 := strings.NewReader("Second part. ")
	reader3 := strings.NewReader("Third part.")

	// Chain them together
	chained := io.MultiReader(reader1, reader2, reader3)

	// Read from the chain
	buffer := make([]byte, 50)
	n, _ := chained.Read(buffer)
	fmt.Printf("Read: %s\n", string(buffer[:n]))

	// Continue reading
	n, _ = chained.Read(buffer)
	fmt.Printf("Read: %s\n", string(buffer[:n]))
}

// ---- Buffered I/O ----

// BufferedIOExample shows bufio usage
func BufferedIOExample() {
	fmt.Println("\n=== Buffered I/O ===")

	// Create a reader with buffer
	text := "Line 1\nLine 2\nLine 3\nLine 4"
	reader := strings.NewReader(text)
	bufferedReader := bufio.NewReader(reader)

	// Read line by line
	for i := 1; i <= 4; i++ {
		line, _ := bufferedReader.ReadString('\n')
		fmt.Printf("Line %d: %s", i, line)
	}

	// Buffered writer
	var output strings.Builder
	bufferedWriter := bufio.NewWriter(&output)

	fmt.Fprintf(bufferedWriter, "Writing ")
	fmt.Fprintf(bufferedWriter, "to ")
	fmt.Fprintf(bufferedWriter, "buffer\n")

	bufferedWriter.Flush() // Must flush to write to underlying writer
	fmt.Printf("Output: %s\n", output.String())
}

// ---- Practical Example: Custom Pipeline ----

// TransformReader wraps a reader and transforms data
type TransformReader struct {
	reader io.Reader
	transform func([]byte) []byte
}

func (tr *TransformReader) Read(p []byte) (int, error) {
	n, err := tr.reader.Read(p)
	if n > 0 {
		transformed := tr.transform(p[:n])
		copy(p, transformed)
		n = len(transformed)
	}
	return n, err
}

// PipelineExample shows building data transformation pipelines
func PipelineExample() {
	fmt.Println("\n=== Data Transformation Pipeline ===")

	// Original reader
	source := strings.NewReader("hello world")

	// Wrap with transformation (uppercase)
	transformed := &TransformReader{
		reader: source,
		transform: func(data []byte) []byte {
			return bytes.ToUpper(data)
		},
	}

	// Read through transformation
	buffer := make([]byte, 100)
	n, _ := transformed.Read(buffer)
	fmt.Printf("Transformed: %s\n", string(buffer[:n]))
}

// ---- Real World: File I/O ----

// FileIOExample shows real file operations
func FileIOExample() {
	fmt.Println("\n=== Real-World: File Operations ===")

	// Create a temporary file for demo
	content := []byte("Go is a great language!\n")
	fileName := "/tmp/go-example.txt"

	// Write to file (file implements io.Writer)
	err := os.WriteFile(fileName, content, 0644)
	if err != nil {
		fmt.Printf("Error writing: %v\n", err)
		return
	}

	// Read from file (file implements io.Reader)
	data, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Printf("Error reading: %v\n", err)
		return
	}

	fmt.Printf("File contents: %s", string(data))

	// Clean up
	os.Remove(fileName)
}

// IOInterfacesPlayground runs all io interface examples
func main() {
	fmt.Println("\n========== STANDARD LIBRARY I/O INTERFACES ==========")
	StringAsReader()
	BytesAsReader()
	CustomReaderExample()
	BytesBufferAsWriter()
	CustomWriterExample()
	ReadWriterExample()
	StringerExample()
	CloserExample()
	IOComposedInterfaceExample()
	CopyExample()
	ChainedReaderExample()
	BufferedIOExample()
	PipelineExample()
	FileIOExample()
}
