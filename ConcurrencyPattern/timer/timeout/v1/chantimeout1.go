package main

import (
	"time"
	"fmt"
)
// 使用定时器，我们可以便捷地实现对接收操作的超时设定

func main() {
	intChan := make(chan int, 1)
	go func() {
		time.Sleep(time.Second)
		intChan <- 1
	}()
	select {
	case e := <- intChan:
		fmt.Printf("Received: %v\n", e)
	case <-time.NewTimer(time.Millisecond*500).C:
		fmt.Println("Timeout!")
	}
}
