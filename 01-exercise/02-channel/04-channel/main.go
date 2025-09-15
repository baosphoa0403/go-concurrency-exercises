package main

import "fmt"

// TODO: Implement relaying of message with Channel Direction

func genMsg(in chan<- string) {
	// send message on ch1
	in <- "hello world"
}

func relayMsg(in <-chan string, out chan<- string) {
	// recv message on ch1
	// send it on ch2
	out <- "channel 2" + <-in
}

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	// create ch1 and ch2

	// spine goroutine genMsg and relayMsg
	go genMsg(ch1)
	go relayMsg(ch1, ch2)
	fmt.Println("value channel: ", <-ch2)

	// recv message on ch2
}
