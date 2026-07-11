package reflection

import (
	"errors"
	"fmt"
	"reflect"
)

// PRACTICE EXERCISE #1: Extract tags from a struct
// Implement a function that accepts any struct (or pointer to a struct)
// and returns a map where the key is the field name, and the value is the tag value
// for the given tag name. If the field does not have the specified tag, omit it from the map.
// Return an error if the input is not a struct or pointer to a struct.

func ExtractTags(obj interface{}, tagName string) (map[string]string, error) {
	v := reflect.ValueOf(obj)
	t := reflect.TypeOf(obj)

	if t.Kind() == reflect.Ptr {
		t = t.Elem()
		v = v.Elem()
	}

	if t.Kind() != reflect.Struct {
		return nil, errors.New("input is not a struct or a pointer to struct")
	}

	results := make(map[string]string)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		tagVal := field.Tag.Get(tagName)
		if tagVal != "" {
			results[field.Name] = tagVal
		}
	}

	return results, nil
}

// PRACTICE EXERCISE #2: Invoke no-argument method returning string
// Implement a function that accepts any object and a method name.
// It should invoke the method (which must accept no arguments and return a string)
// and return the string value. If the method does not exist or has a different
// signature, return an error.

func CallMethod(obj interface{}, methodName string) (string, error) {
	v := reflect.ValueOf(obj)
	method := v.MethodByName(methodName)

	if !method.IsValid() {
		return "", fmt.Errorf("method %s not found", methodName)
	}

	mType := method.Type()
	// Validate signature: 0 inputs, 1 output which is string
	if mType.NumIn() != 0 {
		return "", errors.New("method must accept 0 arguments")
	}
	if mType.NumOut() != 1 {
		return "", errors.New("method must return exactly 1 value")
	}
	if mType.Out(0).Kind() != reflect.String {
		return "", errors.New("method must return a string")
	}

	// Call the method
	results := method.Call(nil)
	return results[0].String(), nil
}
