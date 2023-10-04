package main

import (
	"fmt"
	"sync"
	"time"
)

func increment_age(person map[string]int, lock *sync.Mutex) {

	lock.Lock()
	defer lock.Unlock()
	age := person["gab"]
	time.Sleep(time.Millisecond * 100)
	person["gab"] = age + 1
}

func main_16() {

	person_map := map[string]int{
		"gab": 33,
	}

	var lock *sync.Mutex = &sync.Mutex{}

	go increment_age(person_map, lock)
	go increment_age(person_map, lock)
	go increment_age(person_map, lock)
	go increment_age(person_map, lock)
	go increment_age(person_map, lock)
	go increment_age(person_map, lock)

	time.Sleep(time.Millisecond * 5000)

	fmt.Println(person_map["gab"])

}
