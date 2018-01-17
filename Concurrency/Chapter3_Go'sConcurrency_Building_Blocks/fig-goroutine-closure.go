package main

import (
	"sync"
	"fmt"
)
/*
This example will deterministically block the main goroutine until the goroutine hosting the sayHello function terminates.
We've been using a lot of anonymous functions in out examples to create quick goroutine examples.
Let's shift our attention to closures.Closures close around the lexical scope the are created in,
thereby capturing variables.
 */
func main() {
	var wg sync.WaitGroup
	salutation := "hello"
	wg.Add(1)
	go func() {
		defer wg.Done()
		salutation = "welcome"
	}()
	wg.Wait()
	fmt.Println(salutation)
}

/*
result: welcome

Interesting!It turns out that goroutines execute within the same address space they were created in,
and so our program prints out the word "welcome."
 */
