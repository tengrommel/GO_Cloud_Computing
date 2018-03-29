package main

import (
	"sync"
	"sync/atomic"
	"fmt"
)

var counter int64

func main() {
	// Number of goroutines to use.
	const grs = 2

	// wg is used to manage concurrency.
	var wg sync.WaitGroup
	wg.Add(grs)

	// Create two goroutines
	for i:=0;i<grs;i++ {
		go func() {
			for count:=0;count<2;count++ {
				atomic.AddInt64(&counter, 1)
			}
			wg.Done()
		}()
	}

	// Wait for the goroutines to finish.
	wg.Wait()

	// Display the final value.
	fmt.Println("Final Counter:", counter)
}
