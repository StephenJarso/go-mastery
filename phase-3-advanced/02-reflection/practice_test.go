package reflection

import (
	"testing"
)

type PracticeUser struct {
	Username string `validate:"required" db:"username_col"`
	Age      int    `validate:"min=18"`
	Role     string // no tags
}

func TestExtractTags(t *testing.T) {
	u := PracticeUser{
		Username: "sjarso",
		Age:      20,
		Role:     "Admin",
	}

	// 1. Success on struct
	tags, err := ExtractTags(u, "validate")
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if len(tags) != 2 {
		t.Errorf("Expected 2 tags, got %d", len(tags))
	}
	if tags["Username"] != "required" {
		t.Errorf("Expected Username tag 'required', got %s", tags["Username"])
	}
	if tags["Age"] != "min=18" {
		t.Errorf("Expected Age tag 'min=18', got %s", tags["Age"])
	}

	// 2. Success on pointer to struct
	tagsPtr, err := ExtractTags(&u, "db")
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if tagsPtr["Username"] != "username_col" {
		t.Errorf("Expected Username tag 'username_col', got %s", tagsPtr["Username"])
	}

	// 3. Error on non-struct
	_, err = ExtractTags(42, "validate")
	if err == nil {
		t.Error("Expected error for non-struct input, got nil")
	}
}

type GreeterPractice struct{}

func (g GreeterPractice) Hello() string {
	return "Hello World"
}

func (g GreeterPractice) BadMethod(val int) string {
	return "No"
}

func (g GreeterPractice) MultiReturn() (string, int) {
	return "No", 0
}

func TestCallMethod(t *testing.T) {
	gp := GreeterPractice{}

	// 1. Valid call
	val, err := CallMethod(gp, "Hello")
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if val != "Hello World" {
		t.Errorf("Expected 'Hello World', got %q", val)
	}

	// 2. Non-existent method
	_, err = CallMethod(gp, "NonExistent")
	if err == nil {
		t.Error("Expected error for non-existent method")
	}

	// 3. Method with args
	_, err = CallMethod(gp, "BadMethod")
	if err == nil {
		t.Error("Expected error for method with arguments")
	}

	// 4. Method returning multiple values
	_, err = CallMethod(gp, "MultiReturn")
	if err == nil {
		t.Error("Expected error for method returning multiple values")
	}
}
