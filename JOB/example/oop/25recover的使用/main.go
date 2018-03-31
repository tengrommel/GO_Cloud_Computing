package main

import "fmt"

func main() {
	teata()
	teatb(2)
	teatc()
}

func teata() {
	fmt.Println("aaaaaaaaaaaaaaaaaaaa")
}
func teatb(x int) {
	// 设置recover
	defer func() {
		//recover() // 可以打印panic的错误信息
		if err := recover(); err !=nil{
			fmt.Println(err)
		}
	}()// 调用匿名函数

	var a [10]int
	a[x] = 111 // 当选为20
}
func teatc() {
	fmt.Println("cccccccccccccccccccc")
}
