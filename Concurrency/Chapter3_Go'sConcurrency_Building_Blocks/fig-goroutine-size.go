package main

import (
	"runtime"
	"sync"
	"fmt"
)

func main() {
	memConsumed := func() uint64 {
		runtime.GC()
		var s runtime.MemStats
		runtime.ReadMemStats(&s)
		return s.Sys
	}

	var c <- chan interface{}
	var wg sync.WaitGroup
	noop := func() {wg.Done(); <-c}

	/*
	We require a goroutine that will never exit so that we can keep a number of them in memory
	for measurement. Don't worry about how we're achieving this at this time;just know that this
	goroutine won't exit until the process is finished.
	 */

	const numGoroutines = 1e4
	/*
	Here we define the number of goroutines to create.We will use the law of large numbers to
	asymptotically approach the size of a goroutine.
	 */
	wg.Add(numGoroutines)
	before := memConsumed()
	// Here we measure the amount of memory consumed before creating our goroutines.
	for i:=numGoroutines; i>0; i--{
		go noop()
	}
	wg.Wait()
	after := memConsumed()
	// And here we measure the amount of memory consumed after creating our goroutines.
	fmt.Printf("%.3fkb", float64(after-before)/numGoroutines/1000)
}
