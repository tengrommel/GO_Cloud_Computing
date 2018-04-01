package main

import (
	"fmt"
)

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
	Printer("hello")
}
func person1() {
	Printer("world")
}
