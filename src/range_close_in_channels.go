package main

import (
	"fmt"
)

func message(text string, c chan<- string) {
	c <- text
}

func main() {
	c := make(chan string, 2)
	c <- "Message 1"
	c <- "Message 2"

	fmt.Println(len(c), cap(c))

	//Range and close
	close(c)
	for message := range c {
		fmt.Println(message)
	}

	//select
	email_1 := make(chan string)
	email_2 := make(chan string)
	go message("Message 1", email_1)
	go message("Message 2", email_2)
	for i := 0; i < 2; i++ {
		select {
		case m1 := <-email_1:
			fmt.Println("Email received from email ", m1)
		case m2 := <-email_2:
			fmt.Println("Email received from email ", m2)
		}
	}
}
