package main

import "fmt"

func main() {
	a := 10
	b := 20

	defer func(a, b int) {
		fmt.Printf("a = %d, b = %d\n", a, b)
	}(a, b) // () 代表调用此匿名函数，把参数传递过去，已经先传递参数，只是没有调用

	a = 111
	b = 222
	fmt.Printf("外部：a = %d, b = %d\n", a, b)

}
