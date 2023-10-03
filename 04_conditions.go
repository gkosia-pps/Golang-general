package main

import "fmt"

// %v: default, %s string, %d int, %.f float %.2f float two decimals

func main_04() {

	// declare variable that will used only in condition
	if str_len := len("this is a string"); str_len > 10 {
		fmt.Printf("It is %v chars", str_len)
	}

}
