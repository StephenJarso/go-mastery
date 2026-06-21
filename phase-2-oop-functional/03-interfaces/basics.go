package phase2oopfunctional

import "fmt"
//interfacebasics demonstrates declaring and implementing interfaces
type Speaker interface{
	speak() string
}
type Dog struct{
	Name string
}

type Cat struct{
	Name string
}

func(d Dog) speak()string{
	return d.Name+" says woof!"
}
func(c Cat) speak() string{
	return c.Name+" says meow!"
}
