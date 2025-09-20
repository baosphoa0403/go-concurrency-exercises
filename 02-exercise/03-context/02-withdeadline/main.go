package main

import (
	"context"
	"fmt"
	"time"
)

type data struct {
	result string
}

func main() {

	// TODO: set deadline for goroutine to return computational result.

	compute := func(ctx context.Context) <-chan data {
		ch := make(chan data)
		go func() {
			defer close(ch)
			// Simulate work.
			// time.Sleep(50 * time.Millisecond)

			select {
			case <-ctx.Done():
				fmt.Println("context cancel by deadline", ctx.Err())
				break
			case <-time.After(time.Second * 10):
				ch <- data{"123"}
				break
			}
		}()
		return ch
	}
	start := time.Now()
	deadline := time.Now().Add(5 * time.Second)
	ctx, cancel := context.WithDeadline(context.Background(), deadline)
	defer cancel()

	ch := compute(ctx)
	// Nhận kết quả
	if d, ok := <-ch; ok {
		fmt.Printf("work complete: %s (elapsed %s)\n", d.result, time.Since(start))
	} else {
		fmt.Printf("no result (elapsed %s)\n", time.Since(start))
	}

}
