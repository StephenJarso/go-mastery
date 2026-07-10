package phase2oopfunctional

import (
	"testing"
)

// ===== BASICS TESTS =====

func TestSpeakerInterface(t *testing.T) {
	tests := []struct {
		name     string
		speaker  Speaker
		expected string
	}{
		{
			name:     "Dog speaks",
			speaker:  Dog{Name: "Rex"},
			expected: "Rex says woof!",
		},
		{
			name:     "Cat speaks",
			speaker:  Cat{Name: "Whiskers"},
			expected: "Whiskers says meow!",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.speaker.speak()
			if result != tt.expected {
				t.Errorf("got %q, want %q", result, tt.expected)
			}
		})
	}
}

// ===== DUCK-TYPING TESTS =====

func TestWriterInterface(t *testing.T) {
	writers := []Writer{
		ConsoleWriter{Name: "console"},
		FileWriter{Filename: "file.txt"},
		NetworkWriter{Address: "localhost"},
	}

	for _, w := range writers {
		err := w.Write("test data")
		if err != nil {
			t.Errorf("Write failed: %v", err)
		}
	}
}

func TestSaverAndLoader(t *testing.T) {
	tests := []struct {
		name     string
		storage  SaverAndLoader
		shouldOK bool
	}{
		{
			name:     "Database",
			storage:  Database{Name: "TestDB"},
			shouldOK: true,
		},
		{
			name:     "FileStorage",
			storage:  FileStorage{Path: "/tmp/test"},
			shouldOK: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.storage.Load()
			if err != nil && tt.shouldOK {
				t.Errorf("Load failed: %v", err)
			}

			err = tt.storage.Save()
			if err != nil && tt.shouldOK {
				t.Errorf("Save failed: %v", err)
			}
		})
	}
}

// ===== EMPTY INTERFACE TESTS =====

func TestLogValue(t *testing.T) {
	// Test that LogValue doesn't panic with various types
	values := []interface{}{
		"string",
		42,
		3.14,
		true,
		[]int{1, 2, 3},
	}

	for _, v := range values {
		LogValue("test", v) // Should not panic
	}
}

func TestSumNumbers(t *testing.T) {
	tests := []struct {
		name      string
		numbers   []interface{}
		expected  int64
		shouldErr bool
	}{
		{
			name:     "valid sum",
			numbers:  []interface{}{10, 20, 30},
			expected: 60,
		},
		{
			name:      "with non-int",
			numbers:   []interface{}{10, "invalid"},
			shouldErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := SumNumbers(tt.numbers...)
			if tt.shouldErr && err == nil {
				t.Errorf("expected error, got nil")
			}
			if !tt.shouldErr && result != tt.expected {
				t.Errorf("got %d, want %d", result, tt.expected)
			}
		})
	}
}

func TestProcessValue(t *testing.T) {
	tests := []struct {
		value    interface{}
		contains string
	}{
		{nil, "nil"},
		{true, "boolean"},
		{42, "integer"},
		{3.14, "float"},
		{"hello", "string"},
	}

	for _, tt := range tests {
		result := ProcessValue(tt.value)
		if !contains(result, tt.contains) {
			t.Errorf("ProcessValue(%v) = %q, should contain %q", tt.value, result, tt.contains)
		}
	}
}

// ===== TYPE ASSERTIONS TESTS =====

func TestBasicTypeAssertion(t *testing.T) {
	var i interface{} = "hello"

	s, ok := i.(string)
	if !ok || s != "hello" {
		t.Errorf("expected 'hello', got %q", s)
	}

	num, ok := i.(int)
	if ok {
		t.Errorf("should not assert string as int, got %d", num)
	}
}

func TestDescribeFunction(t *testing.T) {
	tests := []struct {
		value    interface{}
		contains string
	}{
		{nil, "nil"},
		{true, "bool"},
		{42, "int"},
		{"test", "string"},
	}

	for _, tt := range tests {
		result := Describe(tt.value)
		if !contains(result, tt.contains) {
			t.Errorf("Describe(%v) = %q, should contain %q", tt.value, result, tt.contains)
		}
	}
}

// ===== IO INTERFACES TESTS =====

func TestCustomReader(t *testing.T) {
	reader := NewCustomReader("test data")
	buffer := make([]byte, 4)

	n, err := reader.Read(buffer)
	if n != 4 || string(buffer) != "test" {
		t.Errorf("expected 'test', got %q", string(buffer[:n]))
	}

	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestCustomWriter(t *testing.T) {
	writer := &CustomWriter{}

	writer.Write([]byte("line1"))
	writer.Write([]byte("line2"))

	lines := writer.GetLines()
	if len(lines) != 2 {
		t.Errorf("expected 2 lines, got %d", len(lines))
	}
}

func TestPersonStringer(t *testing.T) {
	p := Person{"Alice", 30}
	expected := "Alice (age 30)"

	result := p.String()
	if result != expected {
		t.Errorf("got %q, want %q", result, expected)
	}
}

func TestBookStringer(t *testing.T) {
	b := Book{"Go Programming", "John Doe"}
	expected := "'Go Programming' by John Doe"

	result := b.String()
	if result != expected {
		t.Errorf("got %q, want %q", result, expected)
	}
}

func TestResourceCloser(t *testing.T) {
	r := &Resource{name: "test"}

	if r.IsClosed() {
		t.Errorf("expected resource to be open")
	}

	err := r.Close()
	if err != nil {
		t.Errorf("unexpected error on first close: %v", err)
	}

	if !r.IsClosed() {
		t.Errorf("expected resource to be closed")
	}

	err = r.Close()
	if err == nil {
		t.Errorf("expected error on second close")
	}
}

// Helper function
func contains(s, substr string) bool {
	return len(s) > 0 && len(substr) > 0 && (s == substr || len(s) >= len(substr))
}

// ===== TABLE-DRIVEN TESTS =====

func TestInterfacePolymorphism(t *testing.T) {
	speakers := []Speaker{
		Dog{Name: "Buddy"},
		Cat{Name: "Mittens"},
		Dog{Name: "Max"},
	}

	expectedCount := 3
	if len(speakers) != expectedCount {
		t.Errorf("expected %d speakers, got %d", expectedCount, len(speakers))
	}

	for _, s := range speakers {
		result := s.speak()
		if result == "" {
			t.Errorf("speaker.speak() returned empty string")
		}
	}
}

func TestCalculatorInterface(t *testing.T) {
	calcs := []Calculator{
		SimpleCalc{},
		ScientificCalc{},
	}

	for _, calc := range calcs {
		if calc.Add(2, 3) != 5 {
			t.Errorf("Add(2, 3) should be 5")
		}
		if calc.Subtract(5, 3) != 2 {
			t.Errorf("Subtract(5, 3) should be 2")
		}
		if calc.Multiply(3, 4) != 12 {
			t.Errorf("Multiply(3, 4) should be 12")
		}
	}
}

// ===== INTERFACE COMPOSITION TESTS =====

func TestInterfaceComposition(t *testing.T) {
	var composed SaverAndLoader

	// Test with Database
	composed = Database{Name: "TestDB"}
	if err := composed.Save(); err != nil {
		t.Errorf("Save failed: %v", err)
	}
	if err := composed.Load(); err != nil {
		t.Errorf("Load failed: %v", err)
	}

	// Test with FileStorage
	composed = FileStorage{Path: "/tmp/test"}
	if err := composed.Save(); err != nil {
		t.Errorf("Save failed: %v", err)
	}
	if err := composed.Load(); err != nil {
		t.Errorf("Load failed: %v", err)
	}
}

// ===== BENCHMARK TESTS =====

func BenchmarkTypeAssertion(b *testing.B) {
	var i interface{} = "benchmark"

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_, _ = i.(string)
	}
}

func BenchmarkTypeSwitch(b *testing.B) {
	values := []interface{}{"string", 42, 3.14, true}

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		for _, v := range values {
			_ = Describe(v)
		}
	}
}

func BenchmarkInterfaceCall(b *testing.B) {
	dog := Dog{Name: "Buddy"}
	var speaker Speaker = dog

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_ = speaker.speak()
	}
}
