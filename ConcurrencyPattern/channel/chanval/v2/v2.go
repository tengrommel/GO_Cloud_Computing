package main

/*
有时被传递的值的类型不能简单地判定为值类型或引用类型。
例如，一个结构体类型的值中包含了类型为切片的字段。
 */

import (
	"fmt"
	"time"
)

// Counter代表计数器的类型
type Counter struct{
	count int
}

var mapChan = make(chan map[string]Counter, 1)

func main() {
	syncChan := make(chan  struct{}, 2)
	go func() { // 用于演示接收操作
		for{
			if elem, ok := <-mapChan; ok{
				counter := elem["count"]
				counter.count++
			} else {
				break
			}
		}
		fmt.Println("Stopped. [receiver]")
		syncChan<- struct{}{}
	}()

	go func() { // 用于演示发送操作。
		countMap := map[string]Counter{
			"count": Counter{},
		}
		for i:=0;i<5;i++{
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
