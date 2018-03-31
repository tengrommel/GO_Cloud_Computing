package main

import (
	"fmt"
	"runtime"
)

func main() {
	// 创建新的协程
	go func() {
		fmt.Println("aaaaaaaaaaaaaaaaaaaa")
		// 调用了别的函数
		test()
		fmt.Println("bbbbbbbbbbbbbbbbbbbb")
	}()

	for{

	}

}
func test() {
	defer fmt.Println("cccccccccccccccc")
	//return // 终止此函数
	runtime.Goexit() // 终止所在的协程
	fmt.Println("ddddddddddddddddddddd")
}
