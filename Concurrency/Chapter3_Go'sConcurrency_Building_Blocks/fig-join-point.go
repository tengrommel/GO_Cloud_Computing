package main

import (
	"sync"
	"fmt"
)

/*
	a goroutine is a function that is running concurrently(remember: not necessarily in parallel!)alongside other code.
 */


func main() {
	var wg sync.WaitGroup
	sayHello := func() {
		defer wg.Done()
		fmt.Println("hello")
	}
	wg.Add(1)
	go sayHello()
	wg.Wait()
}
