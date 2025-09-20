// generator() -> square() ->
//
//															-> merge -> print
//	            -> square() ->
package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func generator(done chan struct{}, nums ...int) <-chan int {
	out := make(chan int)

	go func() {
		for _, n := range nums {
			select {
			case val := <-done:
				fmt.Println("cancel generator: ", val)
				return
			case out <- n:
				fmt.Println("still send generator", n)
			}
		}
		close(out)
	}()
	return out
}

func square(done chan struct{}, in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			select {
			case val := <-done:
				fmt.Println("cancel square: ", val)
				return
			case out <- n * n:
				fmt.Println("still send square", out)
			}

		}
		defer close(out)
	}()
	return out
}

func merge(done chan struct{}, cs ...<-chan int) <-chan int {
	out := make(chan int)
	var wg sync.WaitGroup

	output := func(c <-chan int, i int) {
		defer wg.Done()
		for n := range c {
			select {
			case val := <-done:
				fmt.Println("cancel merge: ", val, "index: ", i)
				return
			case out <- n:
				fmt.Println("still merge: ", n)
			}
		}
	}

	wg.Add(len(cs))
	for i, c := range cs {
		go output(c, i)
	}

	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

func main() {
	done := make(chan struct{})

	c1 := square(done, generator(done, 2, 3))
	c2 := square(done, generator(done, 3, 5))

	out := merge(done, c1, c2)

	// TODO: cancel goroutines after receiving one value.

	v := <-out
	fmt.Println("first value:", v)
	close(done)

	time.Sleep(200 * time.Millisecond)
	g := runtime.NumGoroutine()
	fmt.Printf("number of goroutine active = %d\n", g)

}
