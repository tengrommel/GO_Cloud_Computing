package main

import (
	"fmt"
	"time"
)

/*
当接收方从通道接收到一个值类型时，对该值的修改就不会影响到发送方持有的那个值。
但对于引用类型的值来说，这个修改会同时影响收发双方持有的值。
*/

var mapChan = make(chan map[string]int, 1)

func main() {
	syncChan := make(chan struct{}, 2)
	go func() { // 用于演示接收操作
		for{
			if elem, ok := <-mapChan; ok{
				elem["count"]++
				fmt.Println(elem["count"])
			} else {
				break
			}
		}
		fmt.Println("Stopped. [receiver]")
		syncChan <- struct{}{}
	}()

	go func() {
		countMap := make(map[string]int)
		for i:=0; i<5;i++ {
			mapChan <- countMap
			time.Sleep(time.Millisecond)
			fmt.Printf("The count map: %v. [sender]\n", countMap)
		}
		close(mapChan)
		syncChan <- struct{}{}
	}()

	<- syncChan
	<- syncChan
}
