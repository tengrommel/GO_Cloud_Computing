package main

import (
	"sync"
	"fmt"
)

var counter int

var mutex sync.Mutex

func main() {
	// Number of goroutines to use.
	const grs = 2

	// wg is used to manage concurrency.
	var wg sync.WaitGroup
	wg.Add(grs)

	// Create two goroutines.
	for i:=0;i<grs;i++ {
		go func() {
			for count:=0;count<2;count++ {
				mutex.Lock()
				{
					// Capture the value of counter.
					value := counter
					// Increment our local value of counter.
					value++
					// Store the value back into counter.
					counter = value
				}
				mutex.Unlock()
			}
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Printf("Final Counter: %d\n", counter)
}
