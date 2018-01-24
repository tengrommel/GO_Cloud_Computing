package main

import (
	"net/http"
	"fmt"
)

func main() {
	checkStatus := func(done <-chan interface{},
											urls ...string,) <-chan *http.Response {
		responses := make(chan *http.Response)
		go func() {
			defer close(responses)
			for _, url := range urls{
				resp, err := http.Get(url)
				if err !=nil{
					fmt.Println(err)
					continue
				}
				select {
				case <-done:
					return
				case responses <- resp:
				}
			}
		}()
		return responses
	}

	done := make(chan interface{})
	defer close(done)

	urls := []string{"http://www.baidu.com", "https://badhost"}
	for response := range checkStatus(done, urls...)  {
		fmt.Printf("Response: %v\n", response.Status)
	}
}

/*
Here we see that the goroutine has been given no choice in the matter. It can’t simply swallow the error,
and so it does the only sensible thing: it prints the error and hopes something is paying attention.
Don’t put your goroutines in this awkward position. I suggest you separate your concerns:

in general, your concurrent processes should send their errors to another part of your program that has
complete information about the state of your program, and can make a more informed decision about
what to do.
 */