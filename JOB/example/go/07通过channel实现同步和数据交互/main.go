package main

import (
	"fmt"
	"time"
)

func main() {
	var ok chan string
	ok = make(chan string)
	go func() {
		defer fmt.Println("子协程调用完毕")

		for i:=0;i<2;i++{
			fmt.Println("子协程 i = ", i)
			time.Sleep(time.Second)
		}
		//ok <- "ok"
		close(ok)
	}()
	//for i := range <- ok{
	//	fmt.Println(i)
	//}
	//for{
	//	for	i := range <-ok{
	//		fmt.Println("子协程：", i)
	//	}
	//}
	<-ok
}
