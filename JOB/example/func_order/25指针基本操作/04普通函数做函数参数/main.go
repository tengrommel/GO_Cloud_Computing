package main

import "fmt"

func swap(a, b *int)  {
	*a, *b = *b, *a
	fmt.Println(*a, *b)
}

func main() {
	a, b := 10, 20
	// 通过一个函数交换a和b的内容
	swap(&a, &b) // 变量本身传递，值传递（站在变量角度）

}
