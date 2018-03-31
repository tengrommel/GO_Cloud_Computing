package main

import "fmt"

func test(x int) {
	result := 100/x
	fmt.Println("result = ", result)
}

func main() {
	defer fmt.Println("aaaaaaaaaaaaaaaaaaaa")
	defer fmt.Println("bbbbbbbbbbbbbbbbbbbb")
	// 调用一个函数，导致内存出问题
	defer test(0)
	defer fmt.Println("cccccccccccccccccccc")
}
