package solutions



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

func NewCarBuilder() *CarBuilder {
	return &CarBuilder{}
}

func (cb *CarBuilder) Brand(brand string) *CarBuilder {
	cb.car.Brand = brand
	return cb
}

func (cb *CarBuilder) Model(model string) *CarBuilder {
	cb.car.Model = model
	return cb
}

func (cb *CarBuilder) Speed(speed int) *CarBuilder {
	cb.car.Speed = speed
	return cb
}

func (cb *CarBuilder) Build() Car {
	return cb.car
}

func (c *Counter) Increment() {
	if c != nil {
		c.val++
	}
}

func (c Counter) GetValue() int {
	return c.val
}

func (s IntSlice) Sum() int {
	sum := 0
	for _, v := range s {
		sum += v
	}
	return sum
}
