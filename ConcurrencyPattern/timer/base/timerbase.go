package main

import (
	"time"
	"fmt"
)

/*
C 代表的是一个chan time.Time类型的带缓冲的接收通道，C的值原先为双向通道，
只不过在赋给字段C的时候被自动转换为了接收通道。
 */

func main() {
	timer := time.NewTimer(2 * time.Second)
	fmt.Printf("Present time: %v.\n", time.Now())
	expirationTime := <- timer.C
	fmt.Printf("Expiration time: %v.\n", expirationTime)
	fmt.Printf("Stop timer: %v.\n", timer.Stop())
}
