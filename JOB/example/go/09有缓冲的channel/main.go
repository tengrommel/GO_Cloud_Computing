package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 3)
	fmt.Printf("len(ch) = %d, cap(ch) = %d\n", len(ch), cap(ch))

	go func() {
		for i:=0;i<3;i++{
			ch <- i // 往chan写内容
			fmt.Printf("子协程[%d]: len(ch) = %d, cap(ch) = %d\n",i, len(ch), cap(ch))
		}
	}()

	// 延时
	time.Sleep(2 * time.Second)
	for i:=0;i<3;i++ {
		num := <- ch // 读管道中内容，没有内容前，阻塞
		fmt.Println("num = ", num)
	}
}
