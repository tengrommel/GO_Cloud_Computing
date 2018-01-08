package main

import (
	"container/ring"
	"fmt"
)

func main() {
	ring := ring.New(3)
	for i:=1;i<=3;i++{
		ring.Value = i
		ring = ring.Next()
	}

	// 计算1+2+3
	s := 0
	ring.Do(func(i interface{}) {
		s += i.(int)
	})
	fmt.Println("sum is", s)
}
