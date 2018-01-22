package main

import "time"

func main() {
	done := make(chan interface{})
	go func() {
		time.Sleep(5 * time.Second)
		close(done)
	}()
	workCounter := 0

loop:
	for  {
		select {
		case <-done:
			break loop
		default:
		}
		// Simulate work
		workCounter++
		time.Sleep(1 * time.Second)
	}

}
