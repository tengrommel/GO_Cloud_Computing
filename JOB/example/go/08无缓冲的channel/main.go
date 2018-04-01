package main

import (
	"fmt"
	"time"
)

func main() {
	// 创建一个无缓冲的channel
	ch := make(chan int)
	fmt.Printf("len(ch) = %d, cap(ch) = %d\n",len(ch), cap(ch))

	// 新建协程
	go func() {
		for i:=0;i<3;i++ {
			fmt.Println("子协程： i = ", i)
			ch <- i // 往chan 写内容
			fmt.Printf("len(ch) = %d, cap(ch) = %d\n",len(ch), cap(ch))
		}
	}()
	// 延时
	time.Sleep(2 * time.Second)
	for i:=0;i<3;i++{
		num := <- ch // 读管道中内容，没有内容前，阻塞
		fmt.Println("num = ", num)
	}
}
