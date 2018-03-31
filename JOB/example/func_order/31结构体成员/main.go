package main

import "fmt"

type Student struct {
	id int
	name string
	sex byte
	age int
	addr string
}

func main() {
	// 定义一个结构体普通变量
	var s Student

	// 操作成员，需要使用点(.)运算符
	s.id = 1
	s.name = "mike"
	s.age = 34
	s.sex = 'm'
	s.addr ="bj"
	fmt.Println("s = ", s)
}
