package main

import (
	"fmt"
	"time"
)

func fun(s string) {
	for i := 0; i < 3; i++ {
		fmt.Println(s)
		time.Sleep(1 * time.Millisecond)
	}
}

func main() {
	// Direct call
	// fun("direct call")

	// TODO: write goroutine with different variants for function call.

	// goroutine function call
	// go fun("direct call")
	// goroutine with anonymous function
	go func() {
		fun("direct call")
	}()

	// goroutine with function value call

	// wait for goroutines to end
	time.Sleep(time.Millisecond * 100)
	fmt.Println("done..")
}
