package main

import (
	"fmt"
	"time"
)

func main() {

	c := make(chan string)

	go func() {
		for {
			c <- "ping"
			time.Sleep(time.Second)
		}
	}()

	go func() {
		for {
			c <- "pong"
			time.Sleep(time.Second)
		}
	}()

	for {
		msg := <-c
		fmt.Println(msg)
	}
}
