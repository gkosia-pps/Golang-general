package main

import (
	"fmt"
)

// %v: default, %s string, %d int, %.f float %.2f float two decimals

func main_03() {
	// print to output formated string
	fmt.Printf("Hi my name is %v \n", "gav")

	// build formatted string
	msg := fmt.Sprintf("This a message from %v", "gav")
	fmt.Println(msg)
}
