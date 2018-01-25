package main

import (
	"time"
	"fmt"
)

func main() {
	intChan := make(chan int, 1)
	go func() {
		for i:=0;i<5;i++ {
			time.Sleep(time.Second)
			intChan <- i
		}
		close(intChan)
	}()
	timeout := time.Millisecond * 500
	var timer *time.Timer
	for{
		if timer == nil{
			timer = time.NewTimer(timeout)
		} else {
			timer.Reset(timeout)
		}
		select {
		case e := <-intChan:
			fmt.Printf("Received: %v\n", e)
		case <-time.NewTimer(time.Millisecond * 500).C:
			fmt.Println("Timeout!")
		}
	}

}
