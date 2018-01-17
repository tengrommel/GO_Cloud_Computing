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
