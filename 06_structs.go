package main

import "fmt"

// struct can contain any structure
type Student struct {
	name string
	age  int
}

func main() {
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

}
