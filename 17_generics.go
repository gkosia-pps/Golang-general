package main

import "fmt"

// generics allows to avoid datatype in order to reuse code
// use T insteam of datatype in functions parameters
func genericfunc[T any](s []T) ([]T, []T) {
	mid := int(len(s) / 2)

	return s[:mid], s[mid:]
}
func main_17() {

	x, y := genericfunc[int]([]int{1, 2, 3, 4, 5, 6, 7})
	fmt.Println(x)
	fmt.Println(y)

	z, l := genericfunc[string]([]string{"aaa", "bbb", "ccc"})
	fmt.Println(z)
	fmt.Println(l)
}
