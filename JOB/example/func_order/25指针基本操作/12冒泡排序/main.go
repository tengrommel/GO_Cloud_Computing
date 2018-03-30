package main

import (
	"math/rand"
	"time"
	"fmt"
)

func main() {
	// 设置种子，只需一次
	// 如果种子参数一样，每次运行程序产生的随机数都一样
	rand.Seed(time.Now().UnixNano())
	var a [10]int
	n := len(a)
	for i:=0;i<n;i++ {
		a[i]= rand.Intn(100) // 100 以内的随机数
		fmt.Printf("%d, ", a[i])
	}
	fmt.Println()
	// 冒泡排序，挨着的2个元素比较，升序（大于则交换）
	for i:=0;i<n-1;i++ {
		for j:=0;j<n-1-i;j++{
			if a[j] > a[j+1]{
				a[j], a[j+1] = a[j+1], a[j]
			}
		}
	}
	fmt.Println(a)
}
