package main

import "fmt"

func main() {
	doWork := func(strings <-chan string) <-chan interface{} {
		completed := make(chan interface{})
		go func() {
			defer fmt.Println("doWork exited.")
			defer close(completed)
			for s:=range strings {
				// Do something interesting
				fmt.Println(s)
			}
		}()
		return completed
	}

	doWork(nil)
	// Perhaps more work is done here
	fmt.Println("Done.")
}

/*
Here we see that the main goroutine passes a nil channel into doWork .
Therefore, the strings channel will never actually gets any strings written onto it,
and the goroutine containing doWork will remain in memory for the lifetime of this process
(we would even deadlock if we joined the goroutine within doWork and the main goroutine).
 */