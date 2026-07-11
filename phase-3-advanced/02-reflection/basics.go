package reflection

import (
	"fmt"
	"reflect"
)

// Reflection in Go allows inspecting and manipulating types and values at runtime.
// The two main entry points are reflect.Type and reflect.Value.
// - reflect.Type represents the Go type itself.
// - reflect.Value represents the actual runtime value.

type User struct {
	ID    int    `json:"id" db:"user_id"`
	Name  string `json:"name"`
	Email string `json:"email" validate:"required"`
}

// InspectVariables shows how to reflect on basic variables.
func InspectVariables() {
	fmt.Println("=== Inspecting Variables with Reflection ===")
	
	x := 42
	y := "hello"
	z := []int{1, 2, 3}

	// 1. reflect.TypeOf returns the Type interface representation of a variable.
	tX := reflect.TypeOf(x)
	tY := reflect.TypeOf(y)
	tZ := reflect.TypeOf(z)

	fmt.Printf("x: type = %v, kind = %v\n", tX, tX.Kind())
	fmt.Printf("y: type = %v, kind = %v\n", tY, tY.Kind())
	fmt.Printf("z: type = %v, kind = %v\n", tZ, tZ.Kind())

	// 2. reflect.ValueOf returns the Value struct representation of a variable.
	vX := reflect.ValueOf(x)
	vY := reflect.ValueOf(y)
	vZ := reflect.ValueOf(z)

	// We can convert a reflect.Value back to its original interface type.
	fmt.Printf("vX value: %d\n", vX.Int())
	fmt.Printf("vY value: %q\n", vY.String())
	fmt.Printf("vZ length: %d\n", vZ.Len())
}

// InspectStruct shows how to inspect struct fields, types, and struct tags.
func InspectStruct(obj interface{}) {
	fmt.Println("\n=== Inspecting Structs with Reflection ===")

	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)

	// Since obj can be a pointer, we must dereference it using Elem() if it is.
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
		v = v.Elem()
	}

	if t.Kind() != reflect.Struct {
		fmt.Println("Passed object is not a struct!")
		return
	}

	fmt.Printf("Struct Type: %s\n", t.Name())
	fmt.Printf("Number of Fields: %d\n", t.NumField())

	// Iterate over the struct fields
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)      // returns reflect.StructField
		value := v.Field(i)      // returns reflect.Value
		
		fmt.Printf("Field %d: Name = %-6s, Type = %-8s, Value = %-15v\n", 
			i, field.Name, field.Type, value.Interface())

		// Access tags
		jsonTag := field.Tag.Get("json")
		dbTag := field.Tag.Get("db")
		validateTag := field.Tag.Get("validate")

		if jsonTag != "" {
			fmt.Printf("  -> json tag: %q\n", jsonTag)
		}
		if dbTag != "" {
			fmt.Printf("  -> db tag: %q\n", dbTag)
		}
		if validateTag != "" {
			fmt.Printf("  -> validate tag: %q\n", validateTag)
		}
	}
}
