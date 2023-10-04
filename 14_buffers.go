package main

import (
	"fmt"
	"time"
)

// buffer has specific size
// writers are blocked when buffer is full
// readers are blocked when buffer is empty
// buffer is fifo

func produce_messages(msg_list []string) chan string {
	buff := make(chan string, 6)

	for i := range msg_list {
		buff <- msg_list[i]
		fmt.Printf("Produce %v \n", msg_list[i])
	}

	return buff
}

func consume_msg(ch chan string, batch_size int) {

	for {
		msg, ok := <-ch
		if !ok {
			fmt.Println("Exiting loop")
			break
		} else {
			fmt.Printf("Consume %v \n", msg)
		}
	}
}
func main_14() {
	msg := []string{"aaa", "bbb", "cccc", "dddd", "eeee", "ffff"}

	bf := produce_messages(msg)

	go consume_msg(bf, 6)

	time.Sleep(time.Millisecond * 3000)

	close(bf)

	time.Sleep(time.Millisecond * 3000)
}
