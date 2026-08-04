package reflection

import (
	"testing"
)

type Dummy struct {
	Name string
	Age  int
}

func (d Dummy) Add(x int) int {
	return d.Age + x
}

func TestGetStructFields(t *testing.T) {
	d := Dummy{Name: "Alice", Age: 25}
	fields := GetStructFields(d)
	if fields == nil {
		t.Skip("Exercise GetStructFields not implemented yet")
	}
	if len(fields) != 2 || fields[0] != "Name" || fields[1] != "Age" {
		t.Errorf("GetStructFields failed: %v", fields)
	}
}

func TestInvokeByName(t *testing.T) {
	d := Dummy{Age: 25}
	res, err := InvokeByName(d, "Add", 5)
	if res == 0 && err == nil {
		t.Skip("Exercise InvokeByName not implemented yet")
	}
	if err != nil || res != 30 {
		t.Errorf("InvokeByName failed: %d, %v", res, err)
	}
}

func TestSetStructField(t *testing.T) {
	d := Dummy{Name: "Alice", Age: 25}
	err := SetStructField(&d, "Name", "Bob")
	if d.Name == "Alice" && err == nil {
		t.Skip("Exercise SetStructField not implemented yet")
	}
	if err != nil || d.Name != "Bob" {
		t.Errorf("SetStructField failed: %v", err)
	}

	err = SetStructField(d, "Name", "Bob")
	if err == nil {
		t.Error("expected error for non-pointer")
	}
}
