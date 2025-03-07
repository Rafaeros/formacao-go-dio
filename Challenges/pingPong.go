// Challenge 3s
// Use channels and goroutines so that your program displays, alternately, the words ping and pong.

package Challenges

import (
	"fmt"
	"time"
)

func Ping(ch chan string) {
	for {
		ch <- "Ping"
		time.Sleep(time.Second)
	}
}

func Pong(ch chan string) {
	for {
		ch <- "Pong"
		time.Sleep(time.Second)
	}
}

func Alternating(ch chan string) {
	for {
		select {
		case msg := <-ch:
			fmt.Println(msg)
		}
	}
}

func PingPong() {
	ch := make(chan string)

	go Ping(ch)
	go Pong(ch)
	go Alternating(ch)
	
	time.Sleep(5 * time.Second)
}