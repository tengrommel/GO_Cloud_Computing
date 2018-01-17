package main

import (
	"sync"
	"fmt"
)

func main() {
	var wg sync.WaitGroup
	for _, salutation := range []string{"hello", "greetings", "good day"}{
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Println(salutation)
			/*
			Here we reference the loop variable salutation created by ranging over a string slice.
			The answer is trickier than most people expect, and is one of the surprising things in Go.
			 */
		}()
	}
	wg.Wait()
}
/*
good day
good day
good day

That's kind of surprising!Let's figure out what's going on here.
In this example, the goroutine is running a closure that has closed over the iteration variable salutation, which has a string.
As our loop iterates, salutation is being assigned to the next string value in the slice literal.
Because the goroutines being scheduled may run at any point in time in the future, it is undetermined what
values will exit before the goroutines are begun. This means the salutation variable falls
out of scope.
 */
