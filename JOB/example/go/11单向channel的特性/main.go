package main

import "fmt"

func main() {
	// 创建一个channel，双向的
	ch := make(chan int)
	// 双向channel能隐式转换为单向channel
	var writeCh chan<- int = ch // 只能写，不能读
	var readCh <-chan int = ch // 只能读，不能写
	go func() {
		writeCh <- 666 // 写
	}()
	fmt.Println(<-readCh)
}
