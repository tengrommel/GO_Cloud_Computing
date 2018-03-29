package main

import (
	"sync"
	"fmt"
)

var counter int

func main() {
	// Number of goroutine to use.
	const grs = 2

	// wg is used to manage concurrency.
	var wg sync.WaitGroup
	wg.Add(grs)

	// Create two goroutines.
	for i:=0;i<grs;i++ {
		go func() {
			for count := 0;count < 2; count++{
				// Capture the value of Counter.
				value := counter
				// Increment our local value of Counter
				value++
				// Store the value back into Counter.
				counter =value
			}
			wg.Done()
		}()
	}
	// Wait for the goroutines to finish.
	wg.Wait()
	fmt.Println("Final Counter:", counter)
}
