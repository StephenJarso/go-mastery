package solutions

import (
	"errors"
	"reflect"
)


func GetStructFields(s interface{}) []string {
	val := reflect.ValueOf(s)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}
	if val.Kind() != reflect.Struct {
		return nil
	}
	typ := val.Type()
	var fields []string
	for i := 0; i < val.NumField(); i++ {
		fields = append(fields, typ.Field(i).Name)
	}
	return fields
}

func InvokeByName(obj interface{}, methodName string, arg int) (int, error) {
	val := reflect.ValueOf(obj)
	method := val.MethodByName(methodName)
	if !method.IsValid() {
		return 0, errors.New("method not found")
	}
	in := []reflect.Value{reflect.ValueOf(arg)}
	out := method.Call(in)
	if len(out) == 0 {
		return 0, nil
	}
	return int(out[0].Int()), nil
}

func SetStructField(s interface{}, fieldName string, newVal interface{}) error {
	val := reflect.ValueOf(s)
	if val.Kind() != reflect.Ptr {
		return errors.New("must pass a pointer to struct")
	}
	elem := val.Elem()
	if elem.Kind() != reflect.Struct {
		return errors.New("must point to a struct")
	}
	field := elem.FieldByName(fieldName)
	if !field.IsValid() {
		return errors.New("field not found")
	}
	if !field.CanSet() {
		return errors.New("field cannot be set")
	}
	fieldVal := reflect.ValueOf(newVal)
	if field.Type() != fieldVal.Type() {
		return errors.New("type mismatch")
	}
	field.Set(fieldVal)
	return nil
}
