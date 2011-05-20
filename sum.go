package main

import (
	"fmt"
)

func Generate(channel chan int) {
	for i := 2; ; i++ {
		channel <- i
	}
}

func main() {
	channel := make(chan int)
	go Generate(channel)
	for i := 0; i < 100; i++ {
		number := <-channel
		fmt.Println(number)
	}
}
