package methods



type Car struct {
	Brand string
	Model string
	Speed int
}

type CarBuilder struct {
	car Car
}

type Counter struct {
	val int
}

type IntSlice []int

// Exercise 1: Car Builder
// Implement method chaining on CarBuilder (Brand, Model, Speed, Build).
func NewCarBuilder() *CarBuilder {
	// TODO: Implement
	return nil
}

func (cb *CarBuilder) Brand(brand string) *CarBuilder {
	// TODO: Implement
	return cb
}

func (cb *CarBuilder) Model(model string) *CarBuilder {
	// TODO: Implement
	return cb
}

func (cb *CarBuilder) Speed(speed int) *CarBuilder {
	// TODO: Implement
	return cb
}

func (cb *CarBuilder) Build() Car {
	// TODO: Implement
	return Car{}
}

// Exercise 2: Value vs Pointer Receivers
// Increment increments Counter.val by 1 (pointer receiver).
// GetValue returns Counter.val (value receiver).
func (c *Counter) Increment() {
	// TODO: Implement
}

func (c Counter) GetValue() int {
	// TODO: Implement
	return 0
}

// Exercise 3: Methods on Custom Slices
// Sum returns the sum of elements in IntSlice.
func (s IntSlice) Sum() int {
	// TODO: Implement
	return 0
}
