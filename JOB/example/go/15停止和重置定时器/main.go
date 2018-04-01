package main

import (
	"time"
	"fmt"
)

func main() {
	timer := time.NewTimer(3 * time.Second)

	go func() {
		<- timer.C
		fmt.Println("子协程可以打印了，因为定时器的时间到")
	}()

	timer.Stop()
	for{

	}
}
