package main

import (
	"time"
	"fmt"
)

func main() {
	// 延时2s后打印一句话
	timer := time.NewTimer(2 * time.Second)
	<-timer.C
	fmt.Println("时间到")

	<-time.After(2 * time.Second)
	fmt.Println("2s")
}
