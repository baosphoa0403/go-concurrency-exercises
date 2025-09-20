// generator() -> square() -> print

package main

import (
	"fmt"
	"sync"
)

func generator(nums ...int) <-chan int {
	out := make(chan int)

	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

func square(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}

func merge(cs ...<-chan int) <-chan int {
	// Implement fan-in
	// merge a list of channels to a single channel
	out := make(chan int)
	var wg sync.WaitGroup
	wg.Add(len(cs))

	for _, c := range cs {
		go func(c <-chan int) {
			for v := range c {
				out <- v
			}
			wg.Done()

		}(c)
	}

	go func() {
		fmt.Println("goroutine close")
		wg.Wait()
		fmt.Println("goroutine after wait")
		close(out)
	}()

	return out
}

func main() {
	// TODO: fan out square stage to run two instances.

	out := merge(square(generator(2, 3)), square(generator(2, 3)))

	// TODO: fan in the results of square stages.
	for v := range out {
		fmt.Println("result:", v)
	}
}
