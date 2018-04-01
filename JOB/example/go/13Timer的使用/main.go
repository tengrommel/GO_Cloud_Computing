package main

import (
	"time"
	"fmt"
)

func main() {
	// 创建一个定时器，设置时间为2s，2s后，往time通道写内容（当前时间）
	timer := time.NewTimer(2 * time.Second)
	fmt.Println("当前时间：", time.Now())
	t := <-timer.C
	fmt.Println("t = ", t)
}
