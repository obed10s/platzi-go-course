package main

import (
	"fmt"
	"time"
)

func pingFunction(pinger <-chan int, ponger chan<- int) {
	for {
		<-pinger
		fmt.Println("Ping")
		ponger <- 1
		time.Sleep(time.Second)
	}
}

func pongFunction(pinger chan<- int, ponger <-chan int) {
	for {
		<-ponger
		fmt.Println("Pong")
		pinger <- 1
		time.Sleep(time.Second)
	}

}

func main() {
	ping := make(chan int)
	pong := make(chan int)

	go pingFunction(ping, pong)
	go pongFunction(ping, pong)

	ping <- 1
	select {}
}
