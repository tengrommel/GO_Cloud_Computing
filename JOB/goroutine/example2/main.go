package main

import (
	"runtime"
	"sync"
	"fmt"
)

func init()  {
	// Allocate one logical processor for the scheduler to use.
	runtime.GOMAXPROCS(1)
}

func main() {
	// wg is used to manage concurrency.
	var wg sync.WaitGroup
	wg.Add(2)

	// Create the first goroutine and manage its life
	go func() {
		printPrime("A")
		wg.Done()
	}()

	// Wait for the goroutine and manage its lifecycle here.
	go func() {
		printPrime("B")
		wg.Done()
	}()

	// Wait for the goroutines to finish
	fmt.Println("Waiting To Finish")
	wg.Wait()
	fmt.Println("Terminating Program")
}

// printPrime displays prime numbers for the first 5000 numbers.
func printPrime(prefix string)  {
	next:
		for outer:=2;outer<5000;outer++ {
			for inner:=2;inner<outer;inner++{
				if outer%inner == 0 {
					continue next
				}
			}
			fmt.Printf("%s:%d\n", prefix, outer)
		}
		fmt.Println("Completed", prefix)
}
