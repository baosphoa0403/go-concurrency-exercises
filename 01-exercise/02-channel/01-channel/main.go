package main

import "fmt"

func main() {
	ch := make(chan int, 1)
	go func(a, b int) {
		c := a + b
		ch <- c
	}(1, 2)

	fmt.Printf("computed value %v\n", <-ch)

	// TODO: get the value computed from goroutine
	// fmt.Printf("computed value %v\n", c)
}
