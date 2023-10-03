package main

import (
	"fmt"
)

func main_02() {

	// int aliases to int32 or int64 depend of the operating system. To save memory use int8

	var name string = "gav"
	var age int = 33
	var salary float32 = 10.2
	var likeFootball bool = true

	fmt.Printf("%q %v %f %v", name, age, salary, likeFootball)

	// let compiler to decide the data type, short assignment operator
	fav_number := 4
	fmt.Println(fav_number)

	// declare multiple variables
	country, city := "cy", "lim"

	fmt.Println(country, city)

	// cast
	var my_var float64 = 88.8
	var my_var_int = int(my_var)

	fmt.Println(my_var_int)

	// constants: immutable, do not support short operator
	// constants values should be known at compile time
	const app_name = "MY_GO_APP"
	fmt.Println(app_name)

}
