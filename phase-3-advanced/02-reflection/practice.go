package reflection

import (
	"errors"
	"reflect"
)


var _ = errors.New
var _ = reflect.TypeOf

// Exercise 1: Get Struct Fields
// Return slice of field names for the given struct using reflection.
func GetStructFields(s interface{}) []string {
	// TODO: Implement
	return nil
}

// Exercise 2: Invoke Method
// Call a method on obj by name that takes an int and returns int.
func InvokeByName(obj interface{}, methodName string, arg int) (int, error) {
	// TODO: Implement
	return 0, nil
}

// Exercise 3: Set Struct Field
// Find field by name and modify its value to newVal. Return error if not settable.
func SetStructField(s interface{}, fieldName string, newVal interface{}) error {
	// TODO: Implement
	return nil
}
