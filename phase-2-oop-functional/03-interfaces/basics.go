package phase2oopfunctional

import "fmt"

// interfacebasics demonstrates declaring and implementing interfaces
// In GO,interfaces are satisfied IMPLICITLY -  there's no "implements" keyword
// If a type has a right method it automatically satisfies the interface
// Speaker is a simple interface with one method
// Any type with speak() string satisfies this interface
type Speaker interface {
	speak() string
}

// Dog satisfies Speaker by having a speak method
type Dog struct {
	Name string
}

func (d Dog) speak() string {
	return d.Name + " says woof!"
}

// Cat also satisfies Speaker - completely independent of Dog
type Cat struct {
	Name string
}

func (c Cat) speak() string {
	return c.Name + " says meow!"
}

//basicInterfaceUsage shows that any Speaker can be used interchangeably
func BasicInterfaceUsage() {
	var s Speaker

	s = Dog{
		Name: "Rex",
	}
	fmt.Println(s.speak())
	s = Cat{Name: "Whiskers"}
	fmt.Println(s.speak())
}
