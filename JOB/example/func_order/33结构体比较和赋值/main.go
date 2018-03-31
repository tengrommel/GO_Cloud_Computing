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
	var s1 Student = Student{1, "mike", 'm', 18, "bj"}
	var s2 Student = Student{1, "mike", 'm', 18, "bj"}
	var s3 Student = Student{2, "mike", 'm', 18, "bj"}

	fmt.Println("s1 == s2", s1 == s2)
	fmt.Println("s1 == s3", s1 == s3)

	// 同类型的结构体变量可以相互赋值
	var tmp Student
	tmp = s3
	fmt.Println("tmp = ", tmp)
}
