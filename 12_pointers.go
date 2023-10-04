package main

import "fmt"

func increment_x(p *int) {
	*p++
}

func main_12() {

	x := 5
	y := x
	y = 6
	fmt.Println(x, y)

	// pointer to int value
	var x_pointer *int
	x_pointer = &x
	*x_pointer++
	increment_x(x_pointer)
	fmt.Println(x)

}
