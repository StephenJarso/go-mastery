package pointers

// Pointers store the memory address of a value.
// - The "&" operator generates a pointer to its operand.
// - The "*" operator denotes the pointer's underlying value.
// Go has no pointer arithmetic.

// IncrementByValue attempts to increment an integer.
// Since Go is pass-by-value, this increments a local copy and has no effect on the caller's variable.
func IncrementByValue(val int) {
	val++
}

// IncrementByPointer increments the integer that the pointer points to.
// This directly modifies the caller's variable.
func IncrementByPointer(val *int) {
	if val != nil {
		*val++
	}
}

// SliceModify demonstrates slice header modification.
// Since slices are reference types (headers containing pointers to underlying arrays),
// modifying slice elements directly affects the caller's underlying array.
func ModifyFirstElement(slice []int, newVal int) {
	if len(slice) > 0 {
		slice[0] = newVal
	}
}
