package main

import (
	"fmt"
	"sync"
)

func main() {
	ch := make(chan int, 6)
	var wg sync.WaitGroup
	go func() {
		for i := 0; i < 6; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				ch <- i
			}()
		}
	}()
	go func() {
		fmt.Println("zoo close")
		wg.Wait()
		close(ch)
	}()

	for v := range ch {
		fmt.Println("value: ", v)
	}

	// go func() {

	// }()

	// TODO: range over channel to recv values

}
