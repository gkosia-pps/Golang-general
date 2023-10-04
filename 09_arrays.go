package main

import "fmt"

// arrays has fix size and store specific datatype

// ...int: slicer
func variadic_fun(nums ...int) {
	fmt.Println(nums)

}

func main_09() {

	my_array := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	for i := 0; i < len(my_array); i++ {
		fmt.Println(my_array[i])
	}

	// slices have dynamic size
	// slices are build on top of arrays
	// sclicers are passed by reference in the functions

	// create slicer ontop of array
	scl := my_array[0:5]
	fmt.Println(scl)

	// create a new slicer
	// make(datatype, initial size, capacity) capacity can be ignored
	// if pass the lenght it will copy the data to other memory place and expand
	my_scl := make([]int, 5, 10)
	fmt.Println(len(my_scl), cap(my_scl))

	// call variadic function
	// using variadic definition we can pass array, slicer or comma separated list
	variadic_fun(1, 1, 2, 3)

	// spread operator
	// used to spread the values of a slicer as input to a periodic function
	variadic_fun(scl...)

	// append values at the end of slicer
	scl = append(scl, 6, 7)
	fmt.Println(scl)

	// two dimension slice: slice of slices
	matrix := make([][]int, 0)
	matrix = append(matrix, []int{})
	fmt.Println(matrix)

	// loop over slicer
	for i, e := range scl {
		fmt.Println(i, e)
	}
}
