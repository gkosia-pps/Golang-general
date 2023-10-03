package main

import "fmt"

type Shape interface {
	tell_your_kind()
}

type Rectangle struct {
}

type Circle struct {
}

func (c Circle) tell_your_kind() {
	print("I am a circle \n")
}

func (c Rectangle) tell_your_kind() {
	print("I am a rectangle \n")
}

func process_shapes(s Shape) {
	s.tell_your_kind()
}
func main_07() {

	process_shapes(Circle{})
	process_shapes(Rectangle{})

	// cast interface level to child struct: type assertion
	circle := Circle{}
	shape := Shape(circle)

	// ok will be true if shape it was Circle
	// c is the reference to the instance
	c, ok := shape.(Circle)

	if ok {
		c.tell_your_kind()
	}

	// OR for multiple checks
	switch t := shape.(type) {
	case Circle:
		t.tell_your_kind()
	case Rectangle:
		t.tell_your_kind()
	default:
		fmt.Println("Its not in excpected shapes")
	}

}
