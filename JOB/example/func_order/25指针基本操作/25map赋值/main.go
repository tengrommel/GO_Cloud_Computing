package main

import "fmt"

func main() {
	m1 := map[int]string{1:"mike", 2:"yoyo"}
	// 赋值，如果已经存在的key值，修改内容
	fmt.Println("m1 = ", m1)
	m1[1] = "c++"
	m1[2] = "go" // map底层自动扩容，和append类似
	fmt.Println("m1 = ", m1)
}

