package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

//TODO: identify the data race
// fix the issue.

func main() {
	resetCh := make(chan struct{})
	start := time.Now()
	var mu sync.Mutex

	go func() {
		t := time.NewTimer(randomDuration())
		for {
			select {
			case <-t.C:
				mu.Lock()
				fmt.Println("tick at", time.Now())
				fmt.Println(time.Now().Sub(start))
				t.Reset(randomDuration())
				mu.Unlock()
			case <-resetCh:
				fmt.Println("stop loop")
				return
			}
		}
	}()

	<-time.After(5 * time.Second)
	close(resetCh)
	fmt.Println("close channel")
}

func randomDuration() time.Duration {
	return time.Duration(rand.Int63n(1e9))
}

//----------------------------------------------------
// (main goroutine) -> t <- (time.AfterFunc goroutine)
//----------------------------------------------------
// (working condition)
// main goroutine..
// t = time.AfterFunc()  // returns a timer..

// AfterFunc goroutine
// t.Reset()        // timer reset
//----------------------------------------------------
// (race condition- random duration is very small)
// AfterFunc goroutine
// t.Reset() // t = nil

// main goroutine..
// t = time.AfterFunc()
//----------------------------------------------------
