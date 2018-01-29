package main

import (
	"fmt"
	"time"
)

func main() {
	go waitAndSay("hello")
	fmt.Println("Hello")
	time.Sleep(3 * time.Second)
}

func waitAndSay(s string)  {
	time.Sleep(2 * time.Second)
	fmt.Println(s)
}
