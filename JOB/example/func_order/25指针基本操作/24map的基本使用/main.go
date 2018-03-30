package main

import "fmt"

func main() {
	// 定义一个变量，类型位map[int]string
	var m1 map[int]string
	fmt.Println("m1 = ", m1)
	// 对于map只有len，没有cap
	fmt.Println("len = ", len(m1))

	// 可以通过make创建
	m2 := make(map[int]string)
	fmt.Println("m2 = ", m2)
	fmt.Println("len = ", len(m2))

	// 可以通过make创建，可以指定长度,只是指定了容量，但是里面却是一个数据也没有
	m3 := make(map[int]string, 2)
	fmt.Println("m3 = ", m3)
	fmt.Println("len = ", len(m3))
	m3[1]= "mike" // 元素的操作
	m3[2]= "go"
	m3[3]= "c++"
	m3[1]= "c++"

	// 键值是唯一的
	fmt.Println("m3 = ", m3)
	fmt.Println("len = ", len(m3))

	m4 := map[int]string{
		1:"mike",2:"go",3:"c++",
	}
	fmt.Println("m4 = ", m4)
}
