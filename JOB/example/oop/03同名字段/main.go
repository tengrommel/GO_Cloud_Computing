package main

import "fmt"

type Person struct {
	name string // 名字
	sex byte // 性别
	age int // 年龄
}

type Student struct {
	Person // 只有类型，没有名字，匿名字段，继承了Person的成员
	id int
	addr string
	name string // 和Person同名了
}

func main() {
	// 声明(定义一个变量)
	var s Student
	s.name = "mike" // 操作的是Student的name，还是Person的name?
	s.sex = 'm'
	s.age = 17
	s.addr= "bj"
	s.Person.name = "yoyo"
	fmt.Printf("s = %+v\n", s)
}
