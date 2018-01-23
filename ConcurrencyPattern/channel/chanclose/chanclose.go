package main

import "fmt"

/*
试图向一个已关闭的通道发送元素，会让发送操作引发运行时恐慌。
注意：无论怎样都不应该在接收端关闭通道。因为接收端无法判断是否还会向该通道发送元素。
 */

func main() {
	dataChan := make(chan int, 5)
	syncChan1 := make(chan struct{}, 1)
	syncChan2 := make(chan struct{}, 2)

	go func() { // 用于演示接收操作
		<- syncChan1
		for {
			if elem, ok := <-dataChan; ok{
				fmt.Printf("Received: %d [receiver]\n", elem)
			} else {
				break
			}
		}
		fmt.Println("Done. [receiver]")
		syncChan2 <- struct{}{}
	}()

	go func() { // 用于演示发送操作。
		for i:=0;i<5;i++ {
			dataChan <- i
			fmt.Printf("Sent: %d [sender]\n", i)
		}
		close(dataChan)
		syncChan1 <- struct{}{} // 唤醒演示
		fmt.Println("Done. [sender]")
		syncChan2 <- struct{}{}
	}()
	<- syncChan2
	<- syncChan2
}
