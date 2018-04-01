package main

import "fmt"

func main() {
	ch := make(chan int) // 数字通信
	quit := make(chan bool) // 程序是否结束

	// 消费者，从channel读取内容
	go func() {
		for i:=0;i<8;i++{
			num := <-ch
			fmt.Println(num)
		}// 可以停止
		quit <- true
	}()


	// 生产者，产生数字，写入channel
	fibonacci(ch, quit)
}

// ch只写
func fibonacci(ch chan<- int, quit <-chan bool) {
	x, y := 1, 1
	for {
		// 监听channel数据的流动
		select {
		case ch<-x:
			x, y = y , x+y
		case flag := <-quit:
			fmt.Println(flag)
			return
		}
	}
}

