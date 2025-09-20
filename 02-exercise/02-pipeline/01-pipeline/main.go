package main

import (
	"fmt"
)

// TODO: Build a Pipeline
// generator() -> square() -> print

// generator - convertes a list of integers to a channel
func generator(nums ...int) <-chan int {

	ch := make(chan int)
	fmt.Println("init generator")
	go func() {
		for v := range nums {
			ch <- v
		}
		close(ch)
	}()
	return ch
}

// square - receive on inbound channel
// square the number
// output on outbound channel
func square(in <-chan int) <-chan int {
	out := make(chan int)
	fmt.Println("init square")
	go func() {

		for v := range in {
			out <- v * 2
		}
		close(out)

	}()
	return out
}

func main() {
	// set up the pipeline
	// ch := generator(1, 2, 3)

	// out := square(ch)

	// for v := range out {
	// 	fmt.Println("result:", v)
	// }

	// run the last stage of pipeline
	// receive the values from square stage
	// print each one, until channel is closed.
	// ch := make(chan int, 1)
	// go func() {
	// 	fmt.Println("Goroutine: trying to send 42")
	// 	ch <- 1
	// 	fmt.Println("Goroutine: sent 42")
	// }()

	// fmt.Println("Main: sleeping 2s before receiving...")
	// time.Sleep(time.Second * 3)
	// val := <-ch
	// fmt.Println("Main: received", val)

	ch := make(chan int) // buffered = 1

	fmt.Println("Main: sending 42 into buffered channel")
	ch <- 42 // KHÔNG block (buffer có chỗ trống)
	fmt.Println("Main: sent 42 without waiting")

	// val := <-ch // Lúc này mới lấy ra
	// fmt.Println("Main: received", val)

}
