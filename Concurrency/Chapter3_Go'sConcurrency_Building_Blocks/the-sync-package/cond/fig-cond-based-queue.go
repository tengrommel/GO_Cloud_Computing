package main

import (
	"sync"
	"time"
	"fmt"
)

func main() {
	c := sync.NewCond(&sync.Mutex{})
	// First, we create our condition using a standard sync.Mutex as the Locker.
	queue := make([]interface{}, 0, 10)
	// Next, we create a slice with a length of zero.
	// Since we know we'll eventually add 10 items, we instantiate it with a capacity of 10.
	removeFromQueue := func(delay time.Duration) {
		time.Sleep(delay)
		c.L.Lock()
		// We once again enter the critical section for the condition so we can modify data pertinent
		// to the condition.
		queue = queue[1:]
		// Here we simulate dequeuing an item by reassigning the head of the slice to the second item.
		fmt.Println("Removed from queue")
		c.L.Unlock()
		c.Signal() // Here we let a goroutine waiting on the condition know that something has occurred.
	}
	for i:=0;i<10;i++ {
		c.L.Lock()
		for len(queue)==2 {
			/* Here we check the length of the queue in a loop.
			This is important because a signal on the condition doesn’t necessarily mean
			what you’ve been waiting for has occurred — only that something has occurred.
			*/
			c.Wait()
			// We call wait, which will suspend the main goroutine until a signal on the condition has been sent.
		}
		fmt.Println("Adding to queue")
		queue = append(queue, struct {}{})
		go removeFromQueue(1 * time.Second)
		// Here we create a new goroutine that will dequeue an element after one second.
		c.L.Unlock()
		// Here we exit the condition's critical section since we've successfully enqueued an item.
	}
}
