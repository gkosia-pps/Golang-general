package main

import (
	"fmt"
)

// to run a process concurent use go keyword
// channels have only one place
// writers are blocked until the value readed, readers are blocked until a value is written
// ch    <- value: send to chanel
// value <- ch: read from chanel
func produce_to_chanel(ch chan int, num int) {
	go func() {
		for i := 0; i < num; i++ {
			ch <- i
			fmt.Printf("Produce : %v \n", i)
		}

	}()
}

func main_13() {

	channel := make(chan int)

	produce_to_chanel(channel, 10)

	for num_readed := 0; num_readed <= 9; {
		if num_readed == 9 {
			break
		}
		num_readed = <-channel
		fmt.Printf("Consume : %v \n", num_readed)
	}

}
