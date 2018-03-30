package main

import (
	"os"
	"fmt"
)

func main() {
	// 接收用户传递的参数，都是以字符串方式传递
	list := os.Args
	n := len(list)
	fmt.Println("n = ", n)
}
