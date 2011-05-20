package main

import (
	"fmt"
)

func Generate(channel chan int) {
	for i := 2; ; i++ {
		channel <- i
	}
}

func Filter(inChannel, outChannel chan int, prime int) {
	for {
		i := <-inChannel
		if i % prime != 0 {
			outChannel <- i
		}
	}
}

func main() {
	channel := make(chan int)
	go Generate(channel)
	for i := 0; i < 100; i++ {
		number := <-channel
		fmt.Println(number)
		ch1 := make(chan int)
		go Filter(channel, ch1, number)
		channel = ch1
	}
}
