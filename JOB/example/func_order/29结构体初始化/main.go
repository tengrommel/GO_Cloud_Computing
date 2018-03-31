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
	// 顺序初始化，每个成员必须初始化
	var s1 Student = Student{
		1,"",'d',43,"bj",
	}
	fmt.Println(s1)

	// 指定成员初始化，没有初始化的成员，自动赋值为0
	s2 := Student{
		id:2,
	}
	fmt.Println(s2)
}
