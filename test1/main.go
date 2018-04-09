package main

import (
	"time"
	"fmt"
)

var cha = make(chan int, 1)
var chb = make(chan int, 0)

func main() {
	go func() {
		time.Sleep(time.Second*6)
		chb <- 1
	}()

	go func() {
		time.Sleep(time.Second*5)
		cha <-1
	}()

	select {
	case a :=<-cha:
		fmt.Println("cha:", a)
	case b:=<-chb:
		fmt.Println("chb:", b)
	case <-time.After(3*time.Second):
		fmt.Println("timeout")
	}
	time.Sleep(10*time.Second)
}
