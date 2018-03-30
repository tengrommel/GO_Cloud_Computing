package main

import "fmt"

func main() {

	slice := []int{0, 1,2,3,4,5 ,6,7,8,9}
	// [low:hig:max]取下标从low开始的元素，len=high -low, cap = max-low
	s1 := slice[:] // [0:len(array):len(slice)]不指定容量和长度一样
	fmt.Println("s1 = ", s1)
	fmt.Printf("len = %d, cap = %d\n", len(s1), cap(s1))
	// 操作某个元素，和数组操作方式一样
	data := slice[0]
	fmt.Println("data = ", data)

	s2 := slice[3:6:7] // a[3], a[4], a[5]  len = 6-3 = 3 cap = 7-3=4
	fmt.Println("s2 = ", s2)
	fmt.Printf("len = %d, cap = %d\n", len(s2), cap(s2))
	s3 := slice[:6] // 从0开始，取6个元素，容量也是6
	fmt.Printf("s3 = ", s3)
	fmt.Printf("len = %d, cap = %d\n", len(s3), cap(s3))

	s4 := slice[3:] // 从下标3开始到结尾
	fmt.Printf("s4 = ", s4)
	fmt.Printf("len = %d, cap = %d\n", len(s4), cap(s4))
}
