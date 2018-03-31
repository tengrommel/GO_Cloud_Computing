package main

import "fmt"

type Person struct {
	name string // 名字
	sex byte // 性别
	age int // 年龄
}



type Student struct {
	*Person
	id int
	addr string
}

func main() {
	s1 := Student{&Person{"teng", 'm', 34},18,"bj"}
	fmt.Printf("s1 = %+v\n", s1)
	// 先定义变量
	var s2 Student
	s2.Person = new(Person) // 分配空间
	s2.name = "yoyo"
	s2.age = 17
	s2.id = 23
	s2.addr = "sz"
	fmt.Println(s2.name, s2.sex, s2.age, s2.id, s2.addr)
}
