package main

import (
	"time"
	"fmt"
)

func main() {
	go SlowCounter(2)
	time.Sleep(15 * time.Second)
}

func SlowCounter(n int)  {
	i := 0
	for{
		// create a duration of n seconds
		d := time.Duration(n) * time.Second

		// create a timer for this duration
		t := time.NewTimer(d)
		<- t.C

		i++
		fmt.Println(i)
	}
}
