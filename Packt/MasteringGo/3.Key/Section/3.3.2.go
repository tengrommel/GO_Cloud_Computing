package main

import (
	"time"
	"fmt"
)

func main() {
	go func(s string) {
		fmt.Println(s)
	}("World")
	fmt.Println("Hello")
	time.Sleep(1*time.Millisecond)
}
