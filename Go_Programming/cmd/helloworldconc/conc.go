package main

import (
	"fmt"
)

func main() {
	ch := make(chan string)
	// go starts a goroutine
	for i:=0;i<500;i++ {
		go printHelloworld(i, ch)
	}
	for {
		msg := <- ch
		fmt.Print(msg)
	}
}

func printHelloworld(i int, ch chan string)  {
	for  {
		ch <- fmt.Sprintf("Hello world from goroutine %d!\n", i)
	}
}