package jsonprocessing

import (
	"bytes"
	"encoding/json"
	"strings"
	"testing"
	"time"
)

func TestBasicJSONOperations(t *testing.T) {
	u := NewUser(1, "StephenJarso", "stephen@example.com", "my-secret-key", "my-password")

	// 1. Marshal User
	data, err := MarshalUser(u)
	if err != nil {
		t.Fatalf("MarshalUser failed: %v", err)
	}

	// Verify excluded secret keys/password
	jsonStr := string(data)
	if strings.Contains(jsonStr, "my-secret-key") || strings.Contains(jsonStr, "my-password") {
		t.Errorf("secret data should have been omitted, JSON: %s", jsonStr)
	}

	// 2. Unmarshal User
	u2, err := UnmarshalUser(data)
	if err != nil {
		t.Fatalf("UnmarshalUser failed: %v", err)
	}

	if u2.ID != u.ID || u2.Username != u.Username {
		t.Errorf("unmarshaled values do not match original: got %+v, expected %+v", u2, u)
	}
}

func TestCustomJSONOperations(t *testing.T) {
	dateVal := time.Date(2026, 7, 13, 0, 0, 0, 0, time.UTC)
	evt := Event{
		Name: "Go Mastery Release",
		Date: CustomDate{Time: dateVal},
	}

	// Marshal
	data, err := json.Marshal(evt)
	if err != nil {
		t.Fatalf("failed to marshal CustomDate: %v", err)
	}

	expectedJSON := `{"name":"Go Mastery Release","date":"2026-07-13"}`
	if string(data) != expectedJSON {
		t.Errorf("expected JSON %q, got %q", expectedJSON, string(data))
	}

	// Unmarshal
	var evt2 Event
	err = json.Unmarshal(data, &evt2)
	if err != nil {
		t.Fatalf("failed to unmarshal CustomDate: %v", err)
	}

	if !evt2.Date.Equal(dateVal) {
		t.Errorf("expected date %v, got %v", dateVal, evt2.Date.Time)
	}
}

func TestNestedAndDynamicJSON(t *testing.T) {
	comp := Company{
		CompanyName: "Go Corp",
		Departments: []Department{
			{
				DeptName: "Engineering",
				Manager:  Employee{Name: "Stephen", Role: "Principal Engineer"},
				Employees: []Employee{
					{Name: "Jacob", Role: "Software Engineer"},
				},
			},
		},
	}

	// Marshal nested
	_, err := json.Marshal(comp)
	if err != nil {
		t.Fatalf("failed to marshal nested structure: %v", err)
	}

	// Dynamic JSON parsing
	dynamicJSON := `{"id": 101, "score": 98.5, "name": "Dynamic"}`
	m, err := UnmarshalDynamicJSON(dynamicJSON)
	if err != nil {
		t.Fatalf("failed to unmarshal dynamic: %v", err)
	}

	score, err := ExtractFloatFromMap(m, "score")
	if err != nil || score != 98.5 {
		t.Errorf("failed to extract score: %v", err)
	}
}

func TestPerformanceAndPrecision(t *testing.T) {
	u := NewUser(42, "StreamUser", "stream@example.com", "", "")

	// 1. Stream Encode
	var buf bytes.Buffer
	err := StreamEncodeUser(&buf, u)
	if err != nil {
		t.Fatalf("StreamEncodeUser failed: %v", err)
	}

	// 2. Stream Decode
	u2, err := StreamDecodeUser(&buf)
	if err != nil {
		t.Fatalf("StreamDecodeUser failed: %v", err)
	}

	if u2.Username != "StreamUser" {
		t.Errorf("expected 'StreamUser', got %q", u2.Username)
	}

	// 3. Precision check
	largeJSON := `{"big_id": 9007199254740991, "amount": 100.50}`
	m, err := DecodeWithNumericPrecision(largeJSON)
	if err != nil {
		t.Fatalf("DecodeWithNumericPrecision failed: %v", err)
	}

	bigIDVal := m["big_id"]
	jsonNum, ok := bigIDVal.(json.Number)
	if !ok {
		t.Fatalf("expected value to be of type json.Number, got %T", bigIDVal)
	}

	intVal, err := jsonNum.Int64()
	if err != nil || intVal != 9007199254740991 {
		t.Errorf("expected 9007199254740991, got %d (err: %v)", intVal, err)
	}
}
