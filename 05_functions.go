package main

import "fmt"

func my_func(num1 int8, num2 int8) int8 {
	return num1 + num2
}

// In go parameters are passed by value
func increment(num int) {
	num += 1
	fmt.Printf("From function increment see index= %v", num)
}

// function can return multiple values
func return_nultiple() (string, string, string) {
	return "first_string", "second_string", "third_string"
}

// func with named return values
func named_return() (x string, y string) {
	x = "x"
	y = "y"

	return
}
func main_05() {
	fmt.Println(my_func(6, 1))

	index := 5
	increment(index)
	fmt.Printf("From main  see index= %v \n", index)

	// use _ to ignore third parameter
	str1, str2, _ := return_nultiple()
	fmt.Println(str1, str2)

	// named return function
	fmt.Println(named_return())
}
