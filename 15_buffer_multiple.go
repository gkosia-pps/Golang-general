package main

import (
	"fmt"
	"time"
)

// select allows to listen to multiple bufers

func produce_topic_a(topic_a chan string) {

	for i := 0; i < 10; i++ {
		topic_a <- fmt.Sprintf("topic_a_message_%v \n", i)
		time.Sleep(time.Millisecond * 500)
	}

	close(topic_a)
}

func produce_topic_b(topic_b chan string) {

	for i := 0; i < 30; i++ {
		topic_b <- fmt.Sprintf("topic_b_message_%v \n", i)
		time.Sleep(time.Millisecond * 300)
	}

	close(topic_b)
}

func main_15() {

	t_a := make(chan string, 10)
	t_b := make(chan string, 30)

	go produce_topic_a(t_a)
	go produce_topic_b(t_b)

	for {
		select {
		case msg, ok := <-t_a:
			if !ok {
				return
			}
			fmt.Printf("Consumed %v \n", msg)
		case msg, ok := <-t_b:
			if !ok {
				return
			}
			fmt.Printf("Consumed %v \n", msg)
		}
	}
}
