package main

import "fmt"

func main() {
	ch := make(chan int) // 创建一个无缓存channel
	// 新建一个goroutine
	go func() {
		for i:=0;i<5;i++ {
			ch <- i // 往通道写数据
		}
		// 不需要再写数据时，关闭channel
		close(ch)
	}()

	//for{
	//	// 如果ok为true，说明这个管道没有关闭
	//	if num, ok := <-ch;ok==true{
	//		fmt.Println("num = ", num)
	//	}else { // 管道关闭
	//		break
	//	}
	//}

	for num := range ch{
		fmt.Println("num = ", num)
	}

}
