package main

import "fmt"

func modify(p *[5]int) {
	fmt.Println(*p)
	(*p)[0] = 666
	fmt.Println("modify a = ", *p)
}

func main() {
	a := [5]int{1, 2, 3, 4, 5} // 初始化
	modify(&a) // 数组地址传递过去
	fmt.Println("main: a = ", a)
}
