package jsonprocessing

import (
	"encoding/json"
	"fmt"
)

// Real-world JSON often contains nested structures, arrays, and objects.
// Additionally, when schema is dynamic or unknown beforehand, we use map[string]interface{}
// or map[string]any (introduced in Go 1.18).

type Employee struct {
	Name string `json:"name"`
	Role string `json:"role"`
}

type Department struct {
	DeptName  string     `json:"dept_name"`
	Manager   Employee   `json:"manager"`   // Nested object
	Employees []Employee `json:"employees"` // Nested array of objects
}

type Company struct {
	CompanyName string       `json:"company_name"`
	Departments []Department `json:"departments"`
}

// UnmarshalDynamicJSON unmarshals a JSON string into a map, allowing access
// to fields without declaring a struct.
func UnmarshalDynamicJSON(jsonData string) (map[string]interface{}, error) {
	var result map[string]interface{}
	err := json.Unmarshal([]byte(jsonData), &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal dynamic JSON: %w", err)
	}
	return result, nil
}

// ExtractFloatFromMap retrieves a float64 from a dynamic map.
// Note: Go's json package decodes all JSON numbers into float64 inside map[string]interface{}.
func ExtractFloatFromMap(m map[string]interface{}, key string) (float64, error) {
	val, ok := m[key]
	if !ok {
		return 0, fmt.Errorf("key %q not found", key)
	}

	floatVal, ok := val.(float64)
	if !ok {
		return 0, fmt.Errorf("value for key %q is not a number", key)
	}

	return floatVal, nil
}
