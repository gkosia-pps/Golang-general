package main

import "fmt"

// first class function
func type_a(message string) string {
	return "!!" + message + "!!"
}

// first class function
func type_b(message string) string {
	return "&&" + message + "&&"
}

// higher order function: take as input a function
func formatter(msg string, type_format func(string) string) {
	fmt.Println(type_format(msg))
}

// currying function: take as input a function and extend it and return a new function (python decorators)
func currying(formatter func(string) string) func(string) {
	return func(x string) {

		// doesnt matter where to specify defer, it will executed just before the function close
		defer fmt.Println("This is the end of the function")
		fmt.Println("Start timing")
		msg := formatter(x)
		fmt.Println(msg)
		fmt.Println("End timing")
	}
}

// Closeres: return function that references veriables of parent scopr
// Variables initialized on parent fuction and used in returned function
func adder() func(int) int {
	sum := 0
	return func(num int) int {
		sum += num
		return sum
	}
}

func main_11() {

	// higher order functions are functions that accept as input other function
	formatter("Hi", type_a)
	formatter("Hi", type_b)

	// currying
	runcode := currying(type_a)
	runcode("Gab")

	// Closeres
	adder_ref := adder()
	adder_ref(5)
	adder_ref(5)
	adder_ref(5)
	adder_ref(5)
	fmt.Println(adder_ref(1))

	// anonymus functions: declared in a scope
	new_curring_passing_anonymous := currying(func(x string) string { return "$$" + x + "$$" })
	new_curring_passing_anonymous("yooo")

}
