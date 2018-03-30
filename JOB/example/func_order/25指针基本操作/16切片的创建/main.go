package main

import "fmt"

func main() {
	// 自动推导类型，同时初始化
	s1 := []int{1,2,3,4}
	fmt.Println(s1)
	// 借助make 函数，格式make(切片类型, 长度, 容量)
	s2 := make([]int, 5, 10)
	fmt.Println(s2)

	// 没有指定容量则长度和容量一样
	s3 := make([]int, 5)
	fmt.Println(s3)
	// 切片和数组的区别
	// 数组[]里面的长度是固定的一个常量，数组不能修改长度，len和cap永远都是5
	a := [5]int{}
	fmt.Printf("len = %d, cap = %d\n", len(a), cap(a))
	// 切片，[]里面为空，或者为...，切片的长度或容量可以不固定
	s := []int{}
	fmt.Printf("len = %d, cap = %d\n", len(s), cap(s))
	s = append(s, 11) // 给切片末尾追加一个成员
	fmt.Printf("len = %d, cap = %d\n", len(s), cap(s))
}
