package main

import (
	"sync"
	"fmt"
)

/*
You'll notice that we always call Unlock within a defer statement.
This is a very common idiom when utilizing a Mutex to ensure the call always happens, even when panicing.
 */


func main() {
	var count int
	var lock sync.Mutex

	increment := func() {
		lock.Lock()
		/*
		Here we request exclusive use of the critical
		section --- in this case the count variable --- guarded by a Mutex,lock.
		 */
		defer lock.Unlock()
		//Here we indicate that we're done with the critical section lock is guarding.
		count++
		fmt.Printf("Incrementing: %d\n", count)
	}

	decrement := func() {
		lock.Lock()
		defer lock.Unlock()
		count--
		fmt.Printf("Decrementing: %d\n", count)
	}

	// Increment
	var arithmetic sync.WaitGroup
	for i:=0;i<=5;i++ {
		arithmetic.Add(1)
		go func() {
			defer arithmetic.Done()
			increment()
		}()
	}

	// Decrement
	for i:=0;i<=5;i++ {
		arithmetic.Add(1)
		go func() {
			defer arithmetic.Done()
			decrement()
		}()
	}

	arithmetic.Wait()
	fmt.Println("Arithmetic complete.")
}
