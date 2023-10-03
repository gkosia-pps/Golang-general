package main

import "fmt"

// struct can contain any structure
type Student struct {
	name string
	age  int
}

type car struct {
	model  string
	number string
}

type truck struct {
	car
	size int
}

// functions on structs
func (t truck) print_plates() {
	fmt.Println(t.number)
}

func main_06() {
	// initiate a struct instance, all elements will be initialized to deafult
	gab := Student{}
	gab.name = "gab"
	gab.age = 33

	// anonymus struct, when dont want to create more than one instance of the struct
	car := struct {
		model string
		age   int
	}{
		model: "kia",
		age:   1,
	}

	fmt.Println(car.model)

	// embedded structs: the struct that embed the other struct inherits its fields
	my_truck := truck{}
	my_truck.model = "suzuki"
	my_truck.number = "abc123"
	my_truck.print_plates()
}
