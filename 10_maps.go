package main

import (
	"fmt"
)

func main_10() {

	// maps are passed by index

	// create map using make
	students := make(map[string]int)
	students["gab"] = 33
	students["zac"] = 31

	// or create map using short assignt operator
	ages := map[string]int{
		"gab": 33,
		"zac": 31,
	}

	fmt.Println(len(ages), students["gab"])

	// check if key exists and delete it
	if _, ok := students["blabla"]; !ok {
		fmt.Println("Student does not exists")
	}

	// value associate with key and a boolean that indicate if key exists
	if _, ok := students["zac"]; ok {
		// map, key to delete
		delete(students, "zac")

		for k, v := range students {
			fmt.Println(k, v)
		}
	}
}
