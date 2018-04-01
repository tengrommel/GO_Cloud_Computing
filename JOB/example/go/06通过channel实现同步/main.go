package main

import (
	"fmt"
)

var ch = make(chan int)

//定义一个打印机，参数为字符串，按每个字符打印
// 打印机是属于公共资源
func Printer(str string)  {
	for _, data := range str{
		fmt.Printf("%c", data)
	}
	fmt.Printf("\n")
}

func main() {
	// 新建2个协程，代表两个人，两个人同时使用打印机
	go person1()
	go person2()

	for {

	}
}
func person2() {
	<- ch // 从管道取数据，如果通道没有数据他就会阻塞
	Printer("hello")
}

// person1执行完后，才能到person2执行
func person1() {
	Printer("world")
	ch <- 666// 给管道写数据
}
