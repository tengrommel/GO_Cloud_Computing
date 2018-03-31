package main

import "fmt"

// 实现两个数相加
// 面向过程
func Add01(a, b int) int {
	return a + b
}

// 面向对象，方法：给某个类型绑定一个函数
type long int

func (l long)Add02(other long) long {
	return l + other
}

func main() {
	var result int
	result = Add01(1, 1)
	fmt.Println("result = ", result)

	var ll long = 1
	var res long
	res = ll.Add02(3)
	fmt.Println(res)
	var a long
	r := a.Add02(3)
	fmt.Println(r)
}
