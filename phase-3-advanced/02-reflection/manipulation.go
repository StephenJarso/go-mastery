package reflection

import (
	"errors"
	"fmt"
	"reflect"
)

// In Go, values are passed by value. To modify a value using reflection,
// we must pass a POINTER to the value so that it is addressable and settable.
// If we pass a value, v.CanSet() will return false.

// SetStructStringField attempts to dynamically update a string field on a struct.
func SetStructStringField(obj interface{}, fieldName string, newVal string) error {
	v := reflect.ValueOf(obj)

	// To modify the original object, it MUST be a pointer.
	if v.Kind() != reflect.Ptr {
		return errors.New("must pass a pointer to a struct")
	}

	// Get the struct value the pointer points to.
	elem := v.Elem()
	if elem.Kind() != reflect.Struct {
		return errors.New("pointer must point to a struct")
	}

	// Lookup the field by name.
	fieldVal := elem.FieldByName(fieldName)
	if !fieldVal.IsValid() {
		return fmt.Errorf("field %s does not exist on struct", fieldName)
	}

	// Verify that the field is settable and of type string.
	if !fieldVal.CanSet() {
		return fmt.Errorf("field %s is not settable (unexported?)", fieldName)
	}

	if fieldVal.Kind() != reflect.String {
		return fmt.Errorf("field %s is not of type string", fieldName)
	}

	// Update the field!
	fieldVal.SetString(newVal)
	return nil
}

// Greeter is a type with methods that we can invoke dynamically.
type Greeter struct {
	Prefix string
}

func (g Greeter) Greet(name string) string {
	return fmt.Sprintf("%s, %s!", g.Prefix, name)
}

func (g Greeter) SayGoodbye() string {
	return "Goodbye!"
}

// InvokeMethodByName invokes a method on obj by its string name.
func InvokeMethodByName(obj interface{}, methodName string, args ...interface{}) ([]interface{}, error) {
	v := reflect.ValueOf(obj)
	method := v.MethodByName(methodName)

	if !method.IsValid() {
		return nil, fmt.Errorf("method %s not found on object", methodName)
	}

	// Convert arguments to reflect.Value slices
	reflectArgs := make([]reflect.Value, len(args))
	for i, arg := range args {
		reflectArgs[i] = reflect.ValueOf(arg)
	}

	// Invoke the method!
	// Call returns a slice of reflect.Value representing the return values.
	results := method.Call(reflectArgs)

	// Convert results back to interface{} slice
	out := make([]interface{}, len(results))
	for i, res := range results {
		out[i] = res.Interface()
	}

	return out, nil
}
