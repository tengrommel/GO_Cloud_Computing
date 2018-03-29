package main

import (
	"sync"
	"math/rand"
	"time"
	"fmt"
	"sync/atomic"
)

var data []string

var rwMutex sync.RWMutex

var readCount int64

func init()  {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	// wg is used to manage concurrency.
	var wg sync.WaitGroup
	wg.Add(1)

	// Create a writer goroutine.
	go func() {
		for i:=0;i<100;i++ {
			time.Sleep(time.Duration(rand.Intn(100))* time.Millisecond)
			writer(i)
		}
		wg.Done()
	}()

	// Create eight reader goroutines.
	for i:=0;i<8;i++ {
		go func() {
			for{
				reader(i)
			}
		}()
	}
	wg.Wait()
	fmt.Println("Program Complete")
}

// writer adds a new string to the slice in random intervals.
func writer(i int) {

	// Only allow one goroutine to read/write to the slice at a time.
	rwMutex.Lock()
	{
		// Capture the current read count.
		// Keep this safe though we can due without this call.
		rc := atomic.LoadInt64(&readCount)

		// Perform some work since we have a full lock.
		fmt.Printf("****> : Performing Write : RCount[%d]\n", rc)
		data = append(data, fmt.Sprintf("String: %d", i))
	}
	rwMutex.Unlock()
	// Release the lock.
}

// reader wakes up and iterates over the data slice.
func reader(id int) {

	// Any goroutine can read when no write operation is taking place.
	rwMutex.RLock()
	{
		// Increment the read count value by 1.
		rc := atomic.AddInt64(&readCount, 1)

		// Perform some read work and display values.
		time.Sleep(time.Duration(rand.Intn(10)) * time.Millisecond)
		fmt.Printf("%d : Performing Read : Length[%d] RCount[%d]\n", id, len(data), rc)

		// Decrement the read count value by 1.
		atomic.AddInt64(&readCount, -1)
	}
	rwMutex.RUnlock()
	// Release the read lock.
}