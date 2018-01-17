package main

import (
	"sync"
	"fmt"
)

func main() {
	var wg sync.WaitGroup
	for _, saluation := range []string{"hello", "greetings", "good day"}{
		wg.Add(1)
		go func(saluation string) {
			defer wg.Done()
			fmt.Println(saluation)
		}(saluation)
	}
	wg.Wait()
}

/*
Because goroutines operate within the same address space as each other,and simply host functions,
utilizing goroutines is a natural extension to writing nonconcurrent code.
 */