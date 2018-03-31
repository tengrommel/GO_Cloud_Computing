package main

import (
	"math/rand"
	"time"
	"fmt"
)

func main() {
	var n = 10;
	// 创建一个切片，len为n
	s := make([]int, n)
	InitData(s) // 初始化数组
	fmt.Println("排序前：", s)
	BubbleSort(s) // 冒泡排序

}
func BubbleSort(s []int) {
	n := len(s)
	for i:=0;i<n;i++ {
		for j:=0;j<n-i-1 ;i++ {
			if s[j]>s[j+1]{
				s[j], s[j+1] = s[j+1], s[j]
			}
		}
	}
}

func InitData(s []int) {
	// 设置种子
	rand.Seed(time.Now().UnixNano())
	for i:=0;i<len(s);i++ {
		s[i] = rand.Intn(100) // 100 以内的随机数
	}
}
