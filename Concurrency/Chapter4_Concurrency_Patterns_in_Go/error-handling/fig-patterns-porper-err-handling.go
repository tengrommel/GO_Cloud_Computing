package main

import (
	"net/http"
	"fmt"
)

func main() {
	type Result struct {
		Error error
		Response *http.Response
	}
	/*
	Here we create a type that encompasses both the *http.Response and the error possible form
	an iteration of the loop within our goroutine.
	 */
	checkStatus := func(done <-chan interface{}, urls ...string) <-chan Result {
		// This line returns a channel that can be read from to retrieve results of an iteration of our loop.
		results := make(chan Result)
		go func() {
			defer close(results)
			for _, url := range urls {
				var result Result
				resp, err := http.Get(url)
				result = Result{Error:err, Response:resp}
				// Here we create a Result instance with the Error and Response fields set.
				select {
				case <-done:
					return
				case results <- result:
				}
			}
		}()
		return results
	}

	done := make(chan interface{})
	defer close(done)

	urls := []string{"http://www.baidu.com", "https://badhost"}
	for result := range checkStatus(done, urls...) {
		if result.Error != nil{
			/*
			Here, in our main goroutine, we are able to deal with errors coming out of the goroutine started
			by checkStatus intelligently, and with the full context of the larger program.
			 */
			fmt.Printf("error: %v", result.Error)
			continue
		}
		fmt.Printf("Response: %v\n", result.Response.Status)
	}
}
