package main

import (
	"context"
	"fmt"
)

func main() {

	// TODO: generator -  generates integers in a separate goroutine and
	// sends them to the returned channel.
	// The callers of gen need to cancel the goroutine once
	// they consume 5th integer value
	// so that internal goroutine
	// started by gen is not leaked.
	ctx, cancel := context.WithCancel(context.Background())
	generator := func(ctx context.Context) <-chan int {
		ch := make(chan int, 0)
		go func() {
			defer close(ch)
			for i := 0; ; i++ {
				select {
				case <-ctx.Done():
					return
				case ch <- i:
				}
			}
		}()

		return ch
	}

	// go func() {
	// 	<-time.After(time.Second * 5)
	// 	cancel()
	// }()

	value := generator(ctx)
	for v := range value {
		if v == 5 {
			cancel()
		}
		fmt.Println("value: ", v)

	}

	// Create a context that is cancellable.

}
