package reflection

import (
	"testing"
)

func TestInspectVariables(t *testing.T) {
	// Call the inspecting print function to ensure it doesn't panic.
	InspectVariables()
}

func TestInspectStruct(t *testing.T) {
	u := User{
		ID:    1,
		Name:  "Stephen",
		Email: "stephen@example.com",
	}
	InspectStruct(u)
	InspectStruct(&u) // pointer
}

func TestSetStructStringField(t *testing.T) {
	u := User{
		ID:    1,
		Name:  "Stephen",
		Email: "stephen@example.com",
	}

	// 1. Trying to set on non-pointer must fail
	err := SetStructStringField(u, "Name", "Jacob")
	if err == nil {
		t.Error("Expected error when passing non-pointer struct, got nil")
	}

	// 2. Setting valid field on pointer must succeed
	err = SetStructStringField(&u, "Name", "Jacob")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if u.Name != "Jacob" {
		t.Errorf("Expected Name to be 'Jacob', got %q", u.Name)
	}

	// 3. Setting non-existent field must fail
	err = SetStructStringField(&u, "NonExistent", "Value")
	if err == nil {
		t.Error("Expected error when setting non-existent field, got nil")
	}

	// 4. Setting non-string field must fail
	err = SetStructStringField(&u, "ID", "Value")
	if err == nil {
		t.Error("Expected error when setting non-string field with string, got nil")
	}
}

func TestInvokeMethodByName(t *testing.T) {
	g := Greeter{Prefix: "Hello"}

	// 1. Invoke Greet with one argument
	res, err := InvokeMethodByName(g, "Greet", "Stephen")
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if len(res) != 1 {
		t.Fatalf("Expected 1 return value, got %d", len(res))
	}
	if res[0] != "Hello, Stephen!" {
		t.Errorf("Expected 'Hello, Stephen!', got %q", res[0])
	}

	// 2. Invoke SayGoodbye with no arguments
	res, err = InvokeMethodByName(g, "SayGoodbye")
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if len(res) != 1 {
		t.Fatalf("Expected 1 return value, got %d", len(res))
	}
	if res[0] != "Goodbye!" {
		t.Errorf("Expected 'Goodbye!', got %q", res[0])
	}

	// 3. Invoke non-existent method
	_, err = InvokeMethodByName(g, "NonExistent")
	if err == nil {
		t.Error("Expected error when calling non-existent method, got nil")
	}
}
